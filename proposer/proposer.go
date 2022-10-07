package proposer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/taikochain/client-mono/bindings"
	"github.com/taikochain/client-mono/bindings/encoding"
	"github.com/taikochain/client-mono/rpc"
	"github.com/taikochain/client-mono/util"
	"github.com/taikochain/taiko-client/accounts/abi/bind"
	"github.com/taikochain/taiko-client/common"
	"github.com/taikochain/taiko-client/core/types"
	"github.com/taikochain/taiko-client/crypto"
	"github.com/taikochain/taiko-client/ethclient"
	"github.com/taikochain/taiko-client/log"
	"github.com/taikochain/taiko-client/rlp"
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
	l1Node  *ethclient.Client
	l2Node  *ethclient.Client
	taikoL1 *bindings.TaikoL1Client
	taikoL2 *bindings.V1TaikoL2Client

	// Private keys and account addresses
	l1ProposerPrivKey       *ecdsa.PrivateKey
	l1ProposerAddress       common.Address
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
	p := &Proposer{}
	var err error

	// RPC clients
	if p.l1Node, err = rpc.DialClientWithBackoff(
		ctx,
		cfg.L1Node,
	); err != nil {
		return nil, fmt.Errorf("failed to connect to L1 node: %w", err)
	}

	if p.l2Node, err = rpc.DialClientWithBackoff(
		ctx,
		cfg.L2Node,
	); err != nil {
		return nil, fmt.Errorf("failed to connect to L2 node: %w", err)
	}

	if p.taikoL1, err = bindings.NewTaikoL1Client(
		common.HexToAddress(cfg.TaikoL1Address),
		p.l1Node,
	); err != nil {
		return nil, fmt.Errorf("failed to create TaikoL1 client: %w", err)
	}

	if p.taikoL2, err = bindings.NewV1TaikoL2Client(
		common.HexToAddress(cfg.TaikoL2Address),
		p.l2Node,
	); err != nil {
		return nil, fmt.Errorf("failed to create TaikoL2 client: %w", err)
	}

	// Private keys and account addresses
	if p.l1ProposerPrivKey, err = crypto.ToECDSA(
		common.Hex2Bytes(cfg.L1ProposerPrivKey),
	); err != nil {
		return nil, fmt.Errorf("invalid L1 proposer private key: %w", err)
	}

	p.l1ProposerAddress = crypto.PubkeyToAddress(p.l1ProposerPrivKey.PublicKey)
	p.l2SuggestedFeeRecipient = common.HexToAddress(cfg.L2SuggestedFeeRecipien)

	// Proposing configuration
	if p.proposingInterval, err = time.ParseDuration(cfg.ProposeInterval); err != nil {
		return nil, fmt.Errorf("invalid propose interval: %w", err)
	}

	// Protocol constants
	_, _, _, commitDelayConfirmations, _,
		maxGasPerBlock, maxTxPerBlock, _, maxTxBytesPerBlock, minTxGasLimit,
		_, _, _, err := p.taikoL1.GetConstants(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get TaikoL1 constants: %w", err)
	}

	p.commitDelayConfirmations = commitDelayConfirmations.Uint64()
	p.poolContentSplitter = &poolContentSplitter{
		maxTxPerBlock:      maxTxPerBlock.Uint64(),
		maxGasPerBlock:     maxGasPerBlock.Uint64(),
		maxTxBytesPerBlock: maxTxBytesPerBlock.Uint64(),
		minTxGasLimit:      minTxGasLimit.Uint64(),
	}

	log.Info(
		"Protocol constants",
		"commitDelayConfirmations", commitDelayConfirmations,
		"maxTxPerBlock", maxTxPerBlock,
		"maxGasPerBlock", maxGasPerBlock,
		"maxTxBytesPerBlock", maxTxBytesPerBlock,
		"minTxGasLimit", minTxGasLimit,
	)

	// Flags for testing
	p.produceInvalidBlocks = cfg.ProduceInvalidBlocks
	p.produceInvalidBlocksInterval = cfg.ProduceInvalidBlocksInterval

	return p, nil
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
	syncProgress, err := p.l2Node.SyncProgress(ctx)
	if err != nil || syncProgress != nil {
		return fmt.Errorf("l2 node is syncing: %w", err)
	}

	log.Info("Start fetching pending transactions from L2 node's tx pool")

	pendingContent, _, err := p.l2Node.TxPoolContent(ctx)
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
	opts, err := getTxOpts(ctx, p.l1Node, p.l1ProposerPrivKey)
	if err != nil {
		return err
	}

	commitHash := common.BytesToHash(
		encoding.EncodeCommitHash(meta.Beneficiary, meta.TxListHash),
	)

	// Check if the transactions list has been committed before.
	commitHeight, err := p.taikoL1.GetCommitHeight(nil, commitHash)
	if err != nil {
		return fmt.Errorf("check whether a commitHash %s has been committed error: %w", commitHash, err)
	}

	if commitHeight.Cmp(common.Big0) == 0 {
		log.Info("Transactions list has never been committed before", "hash", commitHash)
		commitTx, err := p.taikoL1.CommitBlock(opts, commitHash)
		if err != nil {
			return err
		}
		commitHeight, err = util.WaitForTx(ctx, p.l1Node, commitTx)
		if err != nil {
			return err
		}
	}

	log.Info(
		"Commit block finished, wait some L1 blocks confirmations before proposing",
		"confirmations", p.commitDelayConfirmations,
	)

	if err := util.WaitNConfirmations(
		ctx, p.l1Node, p.commitDelayConfirmations, commitHeight.Uint64(),
	); err != nil {
		return fmt.Errorf("wait L1 blocks confirmations error, commitHash %s: %w", commitHash, err)
	}

	// Propose the transactions list
	inputs, err := encoding.EncodeProposeBlockInput(meta, txListBytes)
	if err != nil {
		return err
	}

	proposeTx, err := p.taikoL1.ProposeBlock(opts, inputs)
	if err != nil {
		return err
	}

	if _, err := util.WaitForTx(ctx, p.l1Node, proposeTx); err != nil {
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
