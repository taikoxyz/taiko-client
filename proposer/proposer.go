package proposer

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
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
	"golang.org/x/sync/errgroup"
)

var (
	errNoNewTxs               = errors.New("no new transactions")
	proposeEmptyBlockGasLimit = 500_000
)

// Proposer keep proposing new transactions from L2 execution engine's tx pool at a fixed interval.
type Proposer struct {
	// RPC clients
	rpc *rpc.Client

	// Private keys and account addresses
	l1ProposerPrivKey       *ecdsa.PrivateKey
	l2SuggestedFeeRecipient common.Address

	// Proposing configurations
	proposingInterval          *time.Duration
	proposeEmptyBlocksInterval *time.Duration
	proposingTimer             *time.Timer
	commitSlot                 uint64
	locals                     []common.Address
	maxProposedTxListsPerEpoch uint64

	// Protocol configurations
	protocolConfigs *bindings.TaikoDataConfig

	// Only for testing purposes
	CustomProposeOpHook func() error
	AfterCommitHook     func() error

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
	p.l1ProposerPrivKey = cfg.L1ProposerPrivKey
	p.l2SuggestedFeeRecipient = cfg.L2SuggestedFeeRecipient
	p.proposingInterval = cfg.ProposeInterval
	p.proposeEmptyBlocksInterval = cfg.ProposeEmptyBlocksInterval
	p.wg = sync.WaitGroup{}
	p.locals = cfg.LocalAddresses
	p.commitSlot = cfg.CommitSlot
	p.maxProposedTxListsPerEpoch = cfg.MaxProposedTxListsPerEpoch
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

	// Protocol configs
	protocolConfigs, err := p.rpc.TaikoL1.GetConfig(nil)
	if err != nil {
		return fmt.Errorf("failed to get protocol configs: %w", err)
	}
	p.protocolConfigs = &protocolConfigs

	log.Info("Protocol configs", "configs", p.protocolConfigs)

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
	defer func() {
		p.proposingTimer.Stop()
		p.wg.Done()
	}()

	var (
		lastNonEmptyBlockProposedAt        = time.Now()
		epoch                       uint64 = 0
	)
	for {
		p.updateProposingTicker()

		select {
		case <-p.ctx.Done():
			return
		case <-p.proposingTimer.C:
			metrics.ProposerProposeEpochCounter.Inc(1)
			epoch += 1

			if err := p.ProposeOp(p.ctx, epoch); err != nil {
				if !errors.Is(err, errNoNewTxs) {
					log.Error("Proposing operation error", "error", err)
					continue
				}

				if p.proposeEmptyBlocksInterval != nil {
					if time.Now().Before(lastNonEmptyBlockProposedAt.Add(*p.proposeEmptyBlocksInterval)) {
						continue
					}

					if err := p.ProposeEmptyBlockOp(p.ctx); err != nil {
						log.Error("Proposing an empty block operation error", "error", err)
					}

					lastNonEmptyBlockProposedAt = time.Now()
				}

				continue
			}

			lastNonEmptyBlockProposedAt = time.Now()
		}
	}
}

// Close closes the proposer instance.
func (p *Proposer) Close() {
	p.wg.Wait()
}

// ProposeOp performs a proposing operation, fetching transactions
// from L2 execution engine's tx pool, splitting them by proposing constraints,
// and then proposing them to TaikoL1 contract.
func (p *Proposer) ProposeOp(ctx context.Context, epoch uint64) error {
	if p.CustomProposeOpHook != nil {
		return p.CustomProposeOpHook()
	}

	// Wait until L2 execution engine is synced at first.
	if err := p.rpc.WaitTillL2Synced(ctx); err != nil {
		return fmt.Errorf("failed to wait until L2 execution engine synced: %w", err)
	}

	stateVars, err := p.rpc.TaikoL1.GetStateVariables(nil)
	if err != nil {
		return fmt.Errorf("failed to get TaikoL1 status variables: %w", err)
	}

	var (
		maxTransactionsPerBlock = new(big.Int).SetUint64(10)
		pendings                = stateVars.NextBlockId - stateVars.LatestVerifiedId - 1
	)

	if epoch%2 == 0 && len(stateVars.FeeBase.String()) < 36 && pendings < 50 {
		maxTransactionsPerBlock = p.protocolConfigs.MaxTransactionsPerBlock
	}

	log.Info(
		"Start fetching L2 execution engine's transaction pool content",
		"epoch", epoch,
		"maxTransactionsPerBlock", maxTransactionsPerBlock,
		"feeBase", stateVars.FeeBase,
		"feeBaseLen", len(stateVars.FeeBase.String()),
		"pendings", pendings,
	)

	metrics.ProposerBaseFeeLenMetrics.Update(int64(len(stateVars.FeeBase.String())))

	txLists, err := p.rpc.GetPoolContent(
		ctx,
		maxTransactionsPerBlock,
		p.protocolConfigs.BlockMaxGasLimit,
		p.protocolConfigs.MaxBytesPerTxList,
		p.protocolConfigs.MinTxGasLimit,
		p.locals,
	)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	log.Info("Transactions lists count", "count", len(txLists))

	if len(txLists) == 0 {
		return errNoNewTxs
	}

	var commitTxListResQueue []*commitTxListRes
	for i, txs := range txLists {
		txListBytes, err := rlp.EncodeToBytes(txs)
		if err != nil {
			return fmt.Errorf("failed to encode transactions: %w", err)
		}

		meta, commitTx, err := p.CommitTxList(ctx, txListBytes, sumTxsGasLimit(txs), i)
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

	head, err := p.rpc.L1.BlockNumber(ctx)
	if err != nil {
		return err
	}
	nonce, err := p.rpc.L1.NonceAt(
		ctx,
		crypto.PubkeyToAddress(p.l1ProposerPrivKey.PublicKey),
		new(big.Int).SetUint64(head),
	)
	if err != nil {
		return err
	}

	log.Info("Proposer account information", "chainHead", head, "nonce", nonce)

	g := new(errgroup.Group)

	for i, res := range commitTxListResQueue {
		func(i int, res *commitTxListRes) {
			g.Go(func() error {
				if i >= int(p.maxProposedTxListsPerEpoch) {
					return nil
				}

				return p.ProposeTxListWithNonce(ctx, res.meta, res.commitTx, res.txListBytes, res.txNum, nonce+uint64(i))
			})
		}(i, res)
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to propose transactions: %w", err)
	}

	return nil
}

// commitTxListRes represents the result of a transactions list committing, will be used when proposing
// the corresponding transactions list.
type commitTxListRes struct {
	meta        *bindings.TaikoDataBlockMetadata
	commitTx    *types.Transaction
	txListBytes []byte
	txNum       uint
}

// CommitTxList submits a given transactions list's corresponding commit hash to TaikoL1 smart contract, then
// after `protocolConfigs.CommitConfirmations` L1 blocks delay, the given transactions list can be proposed.
func (p *Proposer) CommitTxList(ctx context.Context, txListBytes []byte, gasLimit uint64, splittedIdx int) (
	*bindings.TaikoDataBlockMetadata,
	*types.Transaction,
	error,
) {
	// Assemble the block context and commit the txList
	meta := &bindings.TaikoDataBlockMetadata{
		Id:          common.Big0,
		L1Height:    common.Big0,
		L1Hash:      common.Hash{},
		Beneficiary: p.l2SuggestedFeeRecipient,
		GasLimit:    gasLimit,
		TxListHash:  crypto.Keccak256Hash(txListBytes),
		CommitSlot:  p.commitSlot + uint64(splittedIdx),
	}

	if p.protocolConfigs.CommitConfirmations.Cmp(common.Big0) == 0 {
		log.Debug("No commit confirmation delay, skip committing transactions list")
		return meta, nil, nil
	}

	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey, p.rpc.L1ChainID)
	if err != nil {
		return nil, nil, err
	}

	commitHash := common.BytesToHash(encoding.EncodeCommitHash(meta.Beneficiary, meta.TxListHash))

	commitTx, err := p.rpc.TaikoL1.CommitBlock(opts, meta.CommitSlot, commitHash)
	if err != nil {
		return nil, nil, encoding.TryParsingCustomError(err)
	}

	return meta, commitTx, nil
}

// ProposeTxList proposes the given transactions list to TaikoL1 smart contract.
func (p *Proposer) ProposeTxListWithNonce(
	ctx context.Context,
	meta *bindings.TaikoDataBlockMetadata,
	commitTx *types.Transaction,
	txListBytes []byte,
	txNum uint,
	nonce uint64,
) error {
	if p.protocolConfigs.CommitConfirmations.Cmp(common.Big0) > 0 {
		receipt, err := rpc.WaitReceipt(ctx, p.rpc.L1, commitTx)
		if err != nil {
			return err
		}

		if receipt.Status != types.ReceiptStatusSuccessful {
			log.Error("Failed to commit transactions list", "txHash", receipt.TxHash)
			return nil
		}

		log.Info(
			"Commit block finished, wait some L1 blocks confirmations before proposing",
			"commitHeight", receipt.BlockNumber,
			"commitConfirmations", p.protocolConfigs.CommitConfirmations,
		)

		meta.CommitHeight = receipt.BlockNumber.Uint64()

		if err := rpc.WaitConfirmations(
			ctx, p.rpc.L1, p.protocolConfigs.CommitConfirmations.Uint64(), receipt.BlockNumber.Uint64(),
		); err != nil {
			return fmt.Errorf("wait L1 blocks confirmations error, commitHash %s: %w", receipt.BlockNumber, err)
		}
	}

	// Propose the transactions list
	inputs, err := encoding.EncodeProposeBlockInput(meta, txListBytes)
	if err != nil {
		return err
	}

	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey, p.rpc.L1ChainID)
	if err != nil {
		return err
	}

	if len(txListBytes) == 0 {
		opts.GasLimit = uint64(proposeEmptyBlockGasLimit)
	} else {
		opts.GasLimit = 1_000_000
	}

	opts.Nonce = new(big.Int).SetUint64(nonce)

	proposeTx, err := p.rpc.TaikoL1.ProposeBlock(opts, inputs)
	if err != nil {
		return encoding.TryParsingCustomError(err)
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	if _, err := rpc.WaitReceipt(ctxWithTimeout, p.rpc.L1, proposeTx); err != nil {
		return err
	}

	log.Info("📝 Propose transactions succeeded", "txs", txNum)

	metrics.ProposerProposedTxListsCounter.Inc(1)
	metrics.ProposerProposedTxsCounter.Inc(int64(txNum))

	return nil
}

// ProposeTxList proposes the given transactions list to TaikoL1 smart contract.
func (p *Proposer) ProposeTxList(
	ctx context.Context,
	meta *bindings.TaikoDataBlockMetadata,
	commitTx *types.Transaction,
	txListBytes []byte,
	txNum uint,
) error {
	if p.protocolConfigs.CommitConfirmations.Cmp(common.Big0) > 0 {
		receipt, err := rpc.WaitReceipt(ctx, p.rpc.L1, commitTx)
		if err != nil {
			return err
		}

		if receipt.Status != types.ReceiptStatusSuccessful {
			log.Error("Failed to commit transactions list", "txHash", receipt.TxHash)
			return nil
		}

		log.Info(
			"Commit block finished, wait some L1 blocks confirmations before proposing",
			"commitHeight", receipt.BlockNumber,
			"commitConfirmations", p.protocolConfigs.CommitConfirmations,
		)

		meta.CommitHeight = receipt.BlockNumber.Uint64()

		if err := rpc.WaitConfirmations(
			ctx, p.rpc.L1, p.protocolConfigs.CommitConfirmations.Uint64(), receipt.BlockNumber.Uint64(),
		); err != nil {
			return fmt.Errorf("wait L1 blocks confirmations error, commitHash %s: %w", receipt.BlockNumber, err)
		}
	}

	// Propose the transactions list
	inputs, err := encoding.EncodeProposeBlockInput(meta, txListBytes)
	if err != nil {
		return err
	}

	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey, p.rpc.L1ChainID)
	if err != nil {
		return err
	}

	if len(txListBytes) == 0 {
		opts.GasLimit = uint64(proposeEmptyBlockGasLimit)
	}

	proposeTx, err := p.rpc.TaikoL1.ProposeBlock(opts, inputs)
	if err != nil {
		return encoding.TryParsingCustomError(err)
	}

	if _, err := rpc.WaitReceipt(ctx, p.rpc.L1, proposeTx); err != nil {
		return err
	}

	log.Info("📝 Propose transactions succeeded", "txs", txNum)

	metrics.ProposerProposedTxListsCounter.Inc(1)
	metrics.ProposerProposedTxsCounter.Inc(int64(txNum))

	return nil
}

// ProposeEmptyBlockOp performs a proposing one empty block operation.
func (p *Proposer) ProposeEmptyBlockOp(ctx context.Context) error {
	meta, commitTx, err := p.CommitTxList(ctx, []byte{}, 21000, 0)
	if err != nil {
		return fmt.Errorf("failed to commit an empty block: %w", err)
	}

	if err := p.ProposeTxList(ctx, meta, commitTx, []byte{}, 0); err != nil {
		return fmt.Errorf("failed to propose an empty block: %w", err)
	}

	return nil
}

// updateProposingTicker updates the internal proposing timer.
func (p *Proposer) updateProposingTicker() {
	if p.proposingTimer != nil {
		p.proposingTimer.Stop()
	}

	var duration time.Duration
	if p.proposingInterval != nil {
		duration = *p.proposingInterval
	} else {
		// Random number between 12 - 60
		randomSeconds := rand.Intn(60-11) + 12
		duration = time.Duration(randomSeconds) * time.Second
	}

	p.proposingTimer = time.NewTimer(duration)
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
