package proposer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
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
	"github.com/taikochain/taiko-client/rpc"
	"github.com/taikochain/taiko-client/util"
	"github.com/urfave/cli/v2"
)

// Action returns the main function that the subcommand should run.
func Action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		cfg, err := NewConfigFromCliContext(ctx)
		if err != nil {
			return err
		}

		proposer, err := New(context.Background(), cfg)
		if err != nil {
			return err
		}

		return util.RunSubcommand(proposer)
	}
}

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
}

// New initializes a new proposer instance based on the given configurations.
func New(ctx context.Context, cfg *Config) (*Proposer, error) {
	// RPC clients
	rpcClient, err := rpc.NewClient(ctx, &rpc.ClientConfig{
		L1Endpoint:     cfg.L1Endpoint,
		L2Endpoint:     cfg.L2Endpoint,
		TaikoL1Address: cfg.TaikoL1Address,
		TaikoL2Address: cfg.TaikoL2Address,
	})
	if err != nil {
		return nil, fmt.Errorf("initialize rpc clients error: %w", err)
	}

	// Protocol constants
	_, _, _, commitDelayConfirmations, _,
		maxGasPerBlock, maxTxPerBlock, _, maxTxBytesPerBlock, minTxGasLimit,
		_, _, _, err := rpcClient.TaikoL1.GetConstants(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get TaikoL1 constants: %w", err)
	}

	log.Info(
		"Protocol constants",
		"commitDelayConfirmations", commitDelayConfirmations,
		"maxTxPerBlock", maxTxPerBlock,
		"maxGasPerBlock", maxGasPerBlock,
		"maxTxBytesPerBlock", maxTxBytesPerBlock,
		"minTxGasLimit", minTxGasLimit,
	)

	return &Proposer{
		rpc:                      rpcClient,
		l1ProposerPrivKey:        cfg.L1ProposerPrivKey,
		l2SuggestedFeeRecipient:  cfg.L2SuggestedFeeRecipient,
		proposingInterval:        cfg.ProposeInterval,
		commitDelayConfirmations: commitDelayConfirmations.Uint64(),
		poolContentSplitter: &poolContentSplitter{
			maxTxPerBlock:      maxTxPerBlock.Uint64(),
			maxGasPerBlock:     maxGasPerBlock.Uint64(),
			maxTxBytesPerBlock: maxTxBytesPerBlock.Uint64(),
			minTxGasLimit:      minTxGasLimit.Uint64(),
		},
		// Configurations for testing
		produceInvalidBlocks:         cfg.ProduceInvalidBlocks,
		produceInvalidBlocksInterval: cfg.ProduceInvalidBlocksInterval,
	}, nil
}

// Start starts the proposer's main loop.
func (p *Proposer) Start() error {
	// TODO: make the top level context cancellable.
	go func() {
		ticker := time.NewTicker(p.proposingInterval)
		defer ticker.Stop()

		for range ticker.C {
			if err := p.proposeOp(context.Background()); err != nil {
				log.Error("Perform proposing operation error", "error", err)
			}

			// Only for testing purposes
			if p.produceInvalidBlocks {
				if err := p.proposeInvalidBlocksOp(
					context.Background(),
					p.produceInvalidBlocksInterval,
				); err != nil {
					log.Error("Perform proposing invalid blocks operation error", "error", err)
				}
			}
		}
	}()

	return nil
}

// Close closes the proposer instance.
// TODO: implement this method.
func (p *Proposer) Close() {}

// proposeOp performs a proposing operation, fetching transactions
// from L2 node's tx pool, splitting them by proposing constraints,
// and then proposing them to TaikoL1 contract.
func (p *Proposer) proposeOp(ctx context.Context) error {
	syncProgress, err := p.rpc.L2.SyncProgress(ctx)
	if err != nil || syncProgress != nil {
		return fmt.Errorf("l2 node is syncing: %w", err)
	}

	log.Info("Start fetching pending transactions from L2 node's tx pool")

	pendingContent, _, err := p.rpc.L2.TxPoolContent(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	log.Info("Fetching pending transactions finished",
		"length", len(pendingContent),
	)

	for _, txs := range p.poolContentSplitter.split(pendingContent) {
		txListBytes, err := rlp.EncodeToBytes(txs)
		if err != nil {
			return fmt.Errorf("failed to encode transactions: %w", err)
		}

		if err := p.commitAndPropose(
			ctx,
			txListBytes,
			sumTxsGasLimit(txs),
		); err != nil {
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
	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey)
	if err != nil {
		return err
	}

	commitHash := common.BytesToHash(
		encoding.EncodeCommitHash(meta.Beneficiary, meta.TxListHash),
	)

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
		commitHeight, err = util.WaitForTx(ctx, p.rpc.L1, commitTx)
		if err != nil {
			return err
		}
	}

	log.Info(
		"Commit block finished, wait some L1 blocks confirmations before proposing",
		"confirmations", p.commitDelayConfirmations,
	)

	if err := util.WaitNConfirmations(
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

	if _, err := util.WaitForTx(ctx, p.rpc.L1, proposeTx); err != nil {
		return err
	}

	log.Info("📝 Propose transactions succeeded")

	return nil
}

// Name returns the application name.
func (p *Proposer) Name() string {
	return "proposer"
}

// sumTxsGasLimit calculates the summarized gasLimit of the given transactions.
func sumTxsGasLimit(txs []*types.Transaction) uint64 {
	var total uint64
	for i := range txs {
		total += txs[i].Gas()
	}
	return total
}

// getTxOpts creates a bind.TransactOpts instance with the sender's signatures.
func getTxOpts(ctx context.Context, cli *ethclient.Client, privKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainID, err := cli.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get networkID: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate prepareBlock transaction options: %w", err)
	}

	gasTipCap, err := cli.SuggestGasTipCap(ctx)
	if err != nil {
		if util.IsMaxPriorityFeePerGasNotFoundError(err) {
			gasTipCap = util.FallbackGasTipCap
		} else {
			return nil, err
		}
	}

	opts.GasTipCap = gasTipCap

	return opts, nil
}
