package proposer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikochain/taiko-client/bindings"
	"github.com/taikochain/taiko-client/bindings/encoding"
	"github.com/taikochain/taiko-client/pkg/rpc"
	"github.com/urfave/cli/v2"
)

// Proposer keep proposing new transactions from L2 node's tx pool at a fixed interval.
type Proposer struct {
	// RPC clients
	rpc *rpc.Client

	// Private keys and account addresses
	l1ProposerPrivKey       *ecdsa.PrivateKey
	l2SuggestedFeeRecipient common.Address

	// Proposing configuration
	proposingInterval time.Duration

	// Constants in LibConstants
	commitDelayConfirmations uint64
	poolContentSplitter      *poolContentSplitter

	// Flags for testing
	produceInvalidBlocks         bool
	produceInvalidBlocksInterval uint64

	ctx   context.Context
	close context.CancelFunc
	wg    sync.WaitGroup
}

// New initializes the given proposer instance based on the command line flags.
func (p *Proposer) InitFromCli(c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return initFromConfig(p, cfg)
}

// initFromConfig initializes the proposer instance based on the given configurations.
func initFromConfig(p *Proposer, cfg *Config) (err error) {
	log.Debug("Proposer configurations", "config", cfg)

	p.l1ProposerPrivKey = cfg.L1ProposerPrivKey
	p.l2SuggestedFeeRecipient = cfg.L2SuggestedFeeRecipient
	p.proposingInterval = cfg.ProposeInterval
	p.wg = sync.WaitGroup{}

	p.ctx, p.close = context.WithCancel(context.Background())

	// RPC clients
	if p.rpc, err = rpc.NewClient(p.ctx, &rpc.ClientConfig{
		L1Endpoint:     cfg.L1Endpoint,
		L2Endpoint:     cfg.L2Endpoint,
		TaikoL1Address: cfg.TaikoL1Address,
		TaikoL2Address: cfg.TaikoL2Address,
	}); err != nil {
		return fmt.Errorf("initialize rpc clients error: %w", err)
	}

	// Protocol constants
	_, _, _, commitDelayConfirmations, _,
		maxGasPerBlock, maxTxPerBlock, _, maxTxBytesPerBlock, minTxGasLimit,
		_, _, _, err := p.rpc.TaikoL1.GetConstants(nil)
	if err != nil {
		return fmt.Errorf("failed to get TaikoL1 constants: %w", err)
	}

	log.Info(
		"Protocol constants",
		"commitDelayConfirmations", commitDelayConfirmations,
		"maxTxPerBlock", maxTxPerBlock,
		"maxGasPerBlock", maxGasPerBlock,
		"maxTxBytesPerBlock", maxTxBytesPerBlock,
		"minTxGasLimit", minTxGasLimit,
	)

	p.commitDelayConfirmations = commitDelayConfirmations.Uint64()
	p.poolContentSplitter = &poolContentSplitter{
		maxTxPerBlock:      maxTxPerBlock.Uint64(),
		maxGasPerBlock:     maxGasPerBlock.Uint64(),
		maxTxBytesPerBlock: maxTxBytesPerBlock.Uint64(),
		minTxGasLimit:      minTxGasLimit.Uint64(),
	}

	// Configurations for testing
	p.produceInvalidBlocks = cfg.ProduceInvalidBlocks
	p.produceInvalidBlocksInterval = cfg.ProduceInvalidBlocksInterval

	return nil
}

// Start starts the proposer's main loop.
func (p *Proposer) Start() error {
	p.wg.Add(1)
	go p.eventLoop()
	return nil
}

// eventLoop starts the main loop of Taiko proposer.
func (p *Proposer) eventLoop() {
	ticker := time.NewTicker(p.proposingInterval)
	defer func() {
		ticker.Stop()
		p.wg.Done()
	}()

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-ticker.C:
			if err := p.proposeOp(p.ctx); err != nil {
				log.Error("Proposing operation error", "error", err)
				continue
			}

			// Only for testing purposes
			if p.produceInvalidBlocks && p.produceInvalidBlocksInterval > 0 {
				if err := p.proposeInvalidBlocksOp(p.ctx, p.produceInvalidBlocksInterval); err != nil {
					log.Error("Proposing invalid blocks operation error", "error", err)
				}
			}
		}
	}
}

// Close closes the proposer instance.
func (p *Proposer) Close() {
	if p.close != nil {
		p.close()
	}
	p.wg.Wait()
}

// proposeOp performs a proposing operation, fetching transactions
// from L2 node's tx pool, splitting them by proposing constraints,
// and then proposing them to TaikoL1 contract.
func (p *Proposer) proposeOp(ctx context.Context) error {
	syncProgress, err := p.rpc.L2.SyncProgress(ctx)
	if err != nil || syncProgress != nil {
		return fmt.Errorf("l2 node is syncing: %w", err)
	}

	log.Info("Start fetching L2 node's transaction pool content")

	pendingContent, _, err := p.rpc.L2PoolContent(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	log.Info("Fetching L2 pending transactions finished", "length", len(pendingContent))

	for _, txs := range p.poolContentSplitter.split(pendingContent) {
		txListBytes, err := rlp.EncodeToBytes(txs)
		if err != nil {
			return fmt.Errorf("failed to encode transactions: %w", err)
		}

		if err := p.commitAndPropose(ctx, txListBytes, sumTxsGasLimit(txs)); err != nil {
			return fmt.Errorf("failed to commit and propose transactions: %w", err)
		}
	}

	return nil
}

// commitAndPropose proposes new transactions to TaikoL1 contract by committing
// them firstly.
func (p *Proposer) commitAndPropose(ctx context.Context, txListBytes []byte, gasLimit uint64) error {
	// Assemble the block context and commit the txList
	meta := &bindings.LibDataBlockMetadata{
		Id:          common.Big0,
		L1Height:    common.Big0,
		L1Hash:      common.Hash{},
		Beneficiary: p.l2SuggestedFeeRecipient,
		GasLimit:    gasLimit,
		TxListHash:  crypto.Keccak256Hash(txListBytes),
	}
	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey, p.rpc.L1ChainID)
	if err != nil {
		return err
	}

	commitHash := common.BytesToHash(encoding.EncodeCommitHash(meta.Beneficiary, meta.TxListHash))

	// Check if the transactions list has been committed before.
	commitHeight, err := p.rpc.TaikoL1.GetCommitHeight(nil, commitHash)
	if err != nil {
		return fmt.Errorf("check whether a commitHash %s has been committed error: %w", commitHash, err)
	}

	if commitHeight.Cmp(common.Big0) == 0 {
		log.Info("Transactions list has never been committed before", "hash", commitHash)

		commitTx, err := p.rpc.TaikoL1.CommitBlock(opts, commitHash)
		if err != nil {
			return err
		}

		receipt, err := rpc.WaitReceipt(ctx, p.rpc.L1, commitTx)
		if err != nil {
			return err
		}

		commitHeight = receipt.BlockNumber
	}

	log.Info(
		"Commit block finished, wait some L1 blocks confirmations before proposing",
		"commitHeight", commitHeight,
		"confirmations", p.commitDelayConfirmations,
	)

	if err := rpc.WaitConfirmations(
		ctx, p.rpc.L1, p.commitDelayConfirmations, commitHeight.Uint64(),
	); err != nil {
		return fmt.Errorf("wait L1 blocks confirmations error, commitHash %s: %w", commitHash, err)
	}

	// Propose the transactions list
	inputs, err := encoding.EncodeProposeBlockInput(meta, txListBytes)
	if err != nil {
		return err
	}

	proposeTx, err := p.rpc.TaikoL1.ProposeBlock(opts, inputs)
	if err != nil {
		return err
	}

	if _, err := rpc.WaitReceipt(ctx, p.rpc.L1, proposeTx); err != nil {
		return err
	}

	log.Info("📝 Propose transactions succeeded")

	return nil
}

// Name returns the application name.
func (p *Proposer) Name() string {
	return "proposer"
}

// sumTxsGasLimit calculates the accumulated gas limit of all transactions in the list.
func sumTxsGasLimit(txs []*types.Transaction) uint64 {
	var total uint64
	for i := range txs {
		total += txs[i].Gas()
	}
	return total
}

// getTxOpts creates a bind.TransactOpts instance using the given private key.
func getTxOpts(
	ctx context.Context,
	cli *ethclient.Client,
	privKey *ecdsa.PrivateKey,
	chainID *big.Int,
) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate prepareBlock transaction options: %w", err)
	}

	gasTipCap, err := cli.SuggestGasTipCap(ctx)
	if err != nil {
		if rpc.IsMaxPriorityFeePerGasNotFoundError(err) {
			gasTipCap = rpc.FallbackGasTipCap
		} else {
			return nil, err
		}
	}

	opts.GasTipCap = gasTipCap

	return opts, nil
}
