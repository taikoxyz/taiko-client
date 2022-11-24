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
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
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
	commitSlot        uint64

	// Constants in LibConstants
	commitDelayConfirmations uint64
	poolContentSplitter      *poolContentSplitter

	// Flags for testing
	produceInvalidBlocks         bool
	produceInvalidBlocksInterval uint64

	// Only for testing purposes
	AfterCommitHook func() error

	ctx context.Context
	wg  sync.WaitGroup
}

// New initializes the given proposer instance based on the command line flags.
func (p *Proposer) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return InitFromConfig(ctx, p, cfg)
}

// InitFromConfig initializes the proposer instance based on the given configurations.
func InitFromConfig(ctx context.Context, p *Proposer, cfg *Config) (err error) {
	log.Debug("Proposer configurations", "config", cfg)

	p.l1ProposerPrivKey = cfg.L1ProposerPrivKey
	p.l2SuggestedFeeRecipient = cfg.L2SuggestedFeeRecipient
	p.proposingInterval = cfg.ProposeInterval
	p.wg = sync.WaitGroup{}
	p.ctx = ctx

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
	_, _, _, _, commitDelayConfirmations, _,
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
		shufflePoolContent: cfg.ShufflePoolContent,
		maxTxPerBlock:      maxTxPerBlock.Uint64(),
		maxGasPerBlock:     maxGasPerBlock.Uint64(),
		maxTxBytesPerBlock: maxTxBytesPerBlock.Uint64(),
		minTxGasLimit:      minTxGasLimit.Uint64(),
	}
	p.commitSlot = cfg.CommitSlot

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
			metrics.ProposerProposeEpochCounter.Inc(1)

			if err := p.ProposeOp(p.ctx); err != nil {
				log.Error("Proposing operation error", "error", err)
				continue
			}

			// Only for testing purposes
			if p.produceInvalidBlocks && p.produceInvalidBlocksInterval > 0 {
				if err := p.ProposeInvalidBlocksOp(p.ctx, p.produceInvalidBlocksInterval); err != nil {
					log.Error("Proposing invalid blocks operation error", "error", err)
				}
			}
		}
	}
}

// Close closes the proposer instance.
func (p *Proposer) Close() {
	p.wg.Wait()
}

type commitTxListRes struct {
	meta        *bindings.LibDataBlockMetadata
	commitTx    *types.Transaction
	txListBytes []byte
	txNum       uint
}

// ProposeOp performs a proposing operation, fetching transactions
// from L2 node's tx pool, splitting them by proposing constraints,
// and then proposing them to TaikoL1 contract.
func (p *Proposer) ProposeOp(ctx context.Context) error {
	syncProgress, err := p.rpc.L2.SyncProgress(ctx)
	if err != nil || syncProgress != nil {
		return fmt.Errorf("l2 node is syncing: %w, syncProgress: %v", err, syncProgress)
	}

	log.Info("Start fetching L2 node's transaction pool content")

	pendingContent, _, err := p.rpc.L2PoolContent(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	log.Info("Fetching L2 pending transactions finished", "length", pendingContent.ToTxLists().Len())

	var commitTxListResQueue []*commitTxListRes
	for _, txs := range p.poolContentSplitter.split(pendingContent) {
		txListBytes, err := rlp.EncodeToBytes(txs)
		if err != nil {
			return fmt.Errorf("failed to encode transactions: %w", err)
		}

		meta, commitTx, err := p.CommitTxList(ctx, txListBytes, sumTxsGasLimit(txs))
		if err != nil {
			return fmt.Errorf("failed to commit transactions: %w", err)
		}

		commitTxListResQueue = append(commitTxListResQueue, &commitTxListRes{
			meta:        meta,
			commitTx:    commitTx,
			txListBytes: txListBytes,
			txNum:       uint(len(txs)),
		})
	}

	if p.AfterCommitHook != nil {
		if err := p.AfterCommitHook(); err != nil {
			log.Error("Run AfterCommitHook error", "error", err)
		}
	}

	for _, commitTxListRes := range commitTxListResQueue {
		if err := p.ProposeTxList(ctx, commitTxListRes); err != nil {
			return fmt.Errorf("failed to propose transactions: %w", err)
		}

		metrics.ProposerProposedTxListsCounter.Inc(1)
		metrics.ProposerProposedTxsCounter.Inc(int64(commitTxListRes.txNum))
	}

	return nil
}

func (p *Proposer) CommitTxList(ctx context.Context, txListBytes []byte, gasLimit uint64) (
	*bindings.LibDataBlockMetadata,
	*types.Transaction,
	error,
) {
	// Assemble the block context and commit the txList
	meta := &bindings.LibDataBlockMetadata{
		Id:          common.Big0,
		L1Height:    common.Big0,
		L1Hash:      common.Hash{},
		Beneficiary: p.l2SuggestedFeeRecipient,
		GasLimit:    gasLimit,
		TxListHash:  crypto.Keccak256Hash(txListBytes),
		CommitSlot:  p.commitSlot,
	}

	if p.commitDelayConfirmations == 0 {
		log.Debug("No commit delay confirmation, skip committing transactions list")
		return meta, nil, nil
	}

	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey, p.rpc.L1ChainID)
	if err != nil {
		return nil, nil, err
	}

	commitHash := common.BytesToHash(encoding.EncodeCommitHash(meta.Beneficiary, meta.TxListHash))

	commitTx, err := p.rpc.TaikoL1.CommitBlock(opts, p.commitSlot, commitHash)
	if err != nil {
		return nil, nil, err
	}

	return meta, commitTx, nil
}

func (p *Proposer) ProposeTxList(
	ctx context.Context,
	commitRes *commitTxListRes,
) error {
	if p.commitDelayConfirmations > 0 {
		receipt, err := rpc.WaitReceipt(ctx, p.rpc.L1, commitRes.commitTx)
		if err != nil {
			return err
		}

		log.Info(
			"Commit block finished, wait some L1 blocks confirmations before proposing",
			"commitHeight", receipt.BlockNumber,
			"confirmations", p.commitDelayConfirmations,
		)

		commitRes.meta.CommitHeight = receipt.BlockNumber.Uint64()

		if err := rpc.WaitConfirmations(
			ctx, p.rpc.L1, p.commitDelayConfirmations, receipt.BlockNumber.Uint64(),
		); err != nil {
			return fmt.Errorf("wait L1 blocks confirmations error, commitHash %s: %w", receipt.BlockNumber, err)
		}
	}

	// Propose the transactions list
	inputs, err := encoding.EncodeProposeBlockInput(commitRes.meta, commitRes.txListBytes)
	if err != nil {
		return err
	}

	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey, p.rpc.L1ChainID)
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
