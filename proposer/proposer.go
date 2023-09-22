package proposer

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	selector "github.com/taikoxyz/taiko-client/proposer/prover_selector"
	"golang.org/x/sync/errgroup"
)

var (
	errNoNewTxs                = errors.New("no new transactions")
	maxSendProposeBlockTxRetry = 10
	retryInterval              = 12 * time.Second
	proverAssignmentTimeout    = 30 * time.Minute
	requestProverServerTimeout = 12 * time.Second
)

// Proposer keep proposing new transactions from L2 execution engine's tx pool at a fixed interval.
type Proposer struct {
	// rpc clients
	rpc *rpc.Client

	// Private keys and account addresses
	l1ProposerPrivKey       *ecdsa.PrivateKey
	l1ProposerAddress       common.Address
	l2SuggestedFeeRecipient common.Address

	// Proposing configurations
	proposingInterval          *time.Duration
	proposeEmptyBlocksInterval *time.Duration
	proposingTimer             *time.Timer
	locals                     []common.Address
	localsOnly                 bool
	maxProposedTxListsPerEpoch uint64
	proposeBlockTxGasLimit     *uint64
	txReplacementTipMultiplier uint64
	proposeBlockTxGasTipCap    *big.Int

	// Prover selector
	proverSelector selector.ProverSelector

	// Protocol configurations
	protocolConfigs *bindings.TaikoDataConfig

	// Only for testing purposes
	CustomProposeOpHook func() error
	AfterCommitHook     func() error

	ctx context.Context
	wg  sync.WaitGroup

	waitReceiptTimeout time.Duration

	cfg *Config
}

// New initializes the proposer instance based on the given configurations.
func New(ctx context.Context, cfg *Config) (p *Proposer, err error) {
	p = &Proposer{}
	p.rpc, err = EndpointFromConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	p.l1ProposerPrivKey = cfg.L1ProposerPrivKey
	p.l1ProposerAddress = crypto.PubkeyToAddress(p.l1ProposerPrivKey.PublicKey)
	p.l2SuggestedFeeRecipient = cfg.L2SuggestedFeeRecipient
	p.proposingInterval = cfg.ProposeInterval
	p.proposeEmptyBlocksInterval = cfg.ProposeEmptyBlocksInterval
	p.proposeBlockTxGasLimit = cfg.ProposeBlockTxGasLimit
	p.wg = sync.WaitGroup{}
	p.locals = cfg.LocalAddresses
	p.localsOnly = cfg.LocalAddressesOnly
	p.maxProposedTxListsPerEpoch = cfg.MaxProposedTxListsPerEpoch
	p.txReplacementTipMultiplier = cfg.ProposeBlockTxReplacementMultiplier
	p.proposeBlockTxGasTipCap = cfg.ProposeBlockTxGasTipCap
	p.ctx = ctx
	p.waitReceiptTimeout = cfg.WaitReceiptTimeout
	p.cfg = cfg

	// Protocol configs
	protocolConfigs, err := p.rpc.TaikoL1.GetConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get protocol configs: %w", err)
	}
	p.protocolConfigs = &protocolConfigs

	log.Info("Protocol configs", "configs", p.protocolConfigs)

	if p.proverSelector, err = selector.NewETHFeeEOASelector(
		&protocolConfigs,
		p.rpc,
		cfg.TaikoL1Address,
		cfg.BlockProposalFee,
		cfg.BlockProposalFeeIncreasePercentage,
		cfg.ProverEndpoints,
		cfg.BlockProposalFeeIterations,
		proverAssignmentTimeout,
		requestProverServerTimeout,
	); err != nil {
		return nil, err
	}

	return p, nil
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

	lastNonEmptyBlockProposedAt := time.Now()
	for {
		p.updateProposingTicker()

		select {
		case <-p.ctx.Done():
			return
		// proposing interval timer has been reached
		case <-p.proposingTimer.C:
			metrics.ProposerProposeEpochCounter.Inc(1)
			// attempt propose operation
			if err := p.ProposeOp(p.ctx); err != nil {
				if !errors.Is(err, errNoNewTxs) {
					log.Error("Proposing operation error", "error", err)
					continue
				}
				// if no new transactions and empty block interval has passed, propose an empty block
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
func (p *Proposer) Close(ctx context.Context) {
	p.wg.Wait()
}

// ProposeOp performs a proposing operation, fetching transactions
// from L2 execution engine's tx pool, splitting them by proposing constraints,
// and then proposing them to TaikoL1 contract.
func (p *Proposer) ProposeOp(ctx context.Context) error {
	if p.CustomProposeOpHook != nil {
		return p.CustomProposeOpHook()
	}

	// Wait until L2 execution engine is synced at first.
	if err := p.rpc.WaitTillL2ExecutionEngineSynced(ctx); err != nil {
		return fmt.Errorf("failed to wait until L2 execution engine synced: %w", err)
	}

	log.Info("Start fetching L2 execution engine's transaction pool content")

	l2Head, err := p.rpc.L2.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	baseFee, err := p.rpc.TaikoL2.GetBasefee(
		&bind.CallOpts{Context: ctx},
		uint64(time.Now().Unix())-l2Head.Time,
		uint32(l2Head.GasUsed),
	)
	if err != nil {
		return err
	}

	log.Info("Current base fee", "fee", baseFee)

	txLists, err := p.rpc.GetPoolContent(
		ctx,
		p.L2SuggestedFeeRecipient(),
		baseFee,
		p.protocolConfigs.BlockMaxGasLimit,
		p.protocolConfigs.BlockMaxTxListBytes.Uint64(),
		p.locals,
		p.maxProposedTxListsPerEpoch,
	)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	if p.localsOnly {
		var (
			localTxsLists []types.Transactions
			signer        = types.LatestSignerForChainID(p.rpc.L2ChainID)
		)
		for _, txs := range txLists {
			var filtered types.Transactions
			for _, tx := range txs {
				sender, err := types.Sender(signer, tx)
				if err != nil {
					return err
				}

				for _, localAddress := range p.locals {
					if sender == localAddress {
						filtered = append(filtered, tx)
					}
				}
			}

			if filtered.Len() != 0 {
				localTxsLists = append(localTxsLists, filtered)
			}
		}
		txLists = localTxsLists
	}

	log.Info("Transactions lists count", "count", len(txLists))

	if len(txLists) == 0 {
		return errNoNewTxs
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
	for i, txs := range txLists {
		func(i int, txs types.Transactions) {
			g.Go(func() error {
				if i >= int(p.maxProposedTxListsPerEpoch) {
					return nil
				}

				txListBytes, err := rlp.EncodeToBytes(txs)
				if err != nil {
					return fmt.Errorf("failed to encode transactions: %w", err)
				}

				txNonce := nonce + uint64(i)
				if err := p.ProposeTxList(ctx, &encoding.TaikoL1BlockMetadataInput{
					Proposer:        p.l2SuggestedFeeRecipient,
					TxListHash:      crypto.Keccak256Hash(txListBytes),
					TxListByteStart: common.Big0,
					TxListByteEnd:   new(big.Int).SetUint64(uint64(len(txListBytes))),
					CacheTxListInfo: false,
				}, txListBytes, uint(txs.Len()), &txNonce); err != nil {
					return fmt.Errorf("failed to propose transactions: %w", err)
				}

				return nil
			})
		}(i, txs)
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to propose transactions: %w", err)
	}

	if p.AfterCommitHook != nil {
		if err := p.AfterCommitHook(); err != nil {
			log.Error("Run AfterCommitHook error", "error", err)
		}
	}

	return nil
}

// sendProposeBlockTx tries to send a TaikoL1.proposeBlock transaction.
func (p *Proposer) sendProposeBlockTx(
	ctx context.Context,
	meta *encoding.TaikoL1BlockMetadataInput,
	txListBytes []byte,
	nonce *uint64,
	assignment []byte,
	fee *big.Int,
	isReplacement bool,
) (*types.Transaction, error) {
	// Propose the transactions list
	inputs, err := encoding.EncodeProposeBlockInput(meta)
	if err != nil {
		return nil, err
	}
	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey, p.rpc.L1ChainID, fee)
	if err != nil {
		return nil, err
	}
	if nonce != nil {
		opts.Nonce = new(big.Int).SetUint64(*nonce)
	}
	if p.proposeBlockTxGasLimit != nil {
		opts.GasLimit = *p.proposeBlockTxGasLimit
	}
	if isReplacement {
		log.Info("Try replacing a transaction with same nonce", "sender", p.l1ProposerAddress, "nonce", nonce)
		originalTx, err := rpc.GetPendingTxByNonce(ctx, p.rpc, p.l1ProposerAddress, *nonce)
		if err != nil || originalTx == nil {
			log.Warn(
				"Original transaction not found",
				"sender", p.l1ProposerAddress,
				"nonce", nonce,
				"error", err,
			)

			opts.GasTipCap = new(big.Int).Mul(opts.GasTipCap, new(big.Int).SetUint64(p.txReplacementTipMultiplier))
		} else {
			log.Info(
				"Original transaction to replace",
				"sender", p.l1ProposerAddress,
				"nonce", nonce,
				"gasTipCap", originalTx.GasTipCap(),
				"gasFeeCap", originalTx.GasFeeCap(),
			)

			opts.GasTipCap = new(big.Int).Mul(
				originalTx.GasTipCap(),
				new(big.Int).SetUint64(p.txReplacementTipMultiplier),
			)
		}

		if p.proposeBlockTxGasTipCap != nil && opts.GasTipCap.Cmp(p.proposeBlockTxGasTipCap) > 0 {
			log.Info(
				"New gasTipCap exceeds limit, keep waiting",
				"multiplier", p.txReplacementTipMultiplier,
				"newGasTipCap", opts.GasTipCap,
				"maxTipCap", p.proposeBlockTxGasTipCap,
			)
			return nil, txpool.ErrReplaceUnderpriced
		}
	}

	proposeTx, err := p.rpc.TaikoL1.ProposeBlock(opts, inputs, assignment, txListBytes)
	if err != nil {
		return nil, encoding.TryParsingCustomError(err)
	}

	return proposeTx, nil
}

// ProposeTxList proposes the given transactions list to TaikoL1 smart contract.
func (p *Proposer) ProposeTxList(
	ctx context.Context,
	meta *encoding.TaikoL1BlockMetadataInput,
	txListBytes []byte,
	txNum uint,
	nonce *uint64,
) error {
	assignment, fee, err := p.proverSelector.AssignProver(ctx, meta)
	if err != nil {
		return err
	}

	var (
		isReplacement bool
		tx            *types.Transaction
	)
	if err := backoff.Retry(
		func() error {
			if ctx.Err() != nil {
				return nil
			}
			if tx, err = p.sendProposeBlockTx(ctx, meta, txListBytes, nonce, assignment, fee, isReplacement); err != nil {
				log.Warn("Failed to send propose block transaction, retrying", "error", encoding.TryParsingCustomError(err))
				if strings.Contains(err.Error(), txpool.ErrReplaceUnderpriced.Error()) {
					isReplacement = true
				} else {
					isReplacement = false
				}
				return err
			}

			return nil
		},
		backoff.WithMaxRetries(
			backoff.NewConstantBackOff(retryInterval),
			uint64(maxSendProposeBlockTxRetry),
		),
	); err != nil {
		return err
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if err != nil {
		return err
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, p.waitReceiptTimeout)
	defer cancel()

	if _, err := rpc.WaitReceipt(ctxWithTimeout, p.rpc.L1, tx); err != nil {
		return err
	}

	log.Info("📝 Propose transactions succeeded", "txs", txNum)

	metrics.ProposerProposedTxListsCounter.Inc(1)
	metrics.ProposerProposedTxsCounter.Inc(int64(txNum))

	return nil
}

// ProposeEmptyBlockOp performs a proposing one empty block operation.
func (p *Proposer) ProposeEmptyBlockOp(ctx context.Context) error {
	return p.ProposeTxList(ctx, &encoding.TaikoL1BlockMetadataInput{
		TxListHash:      crypto.Keccak256Hash([]byte{}),
		Proposer:        p.L2SuggestedFeeRecipient(),
		TxListByteStart: common.Big0,
		TxListByteEnd:   common.Big0,
		CacheTxListInfo: false,
	}, []byte{}, 0, nil)
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
		// Random number between 12 - 120
		randomSeconds := rand.Intn(120-11) + 12
		duration = time.Duration(randomSeconds) * time.Second
	}

	p.proposingTimer = time.NewTimer(duration)
}

// Name returns the application name.
func (p *Proposer) Name() string {
	return "proposer"
}

// L2SuggestedFeeRecipient returns the L2 suggested fee recipient of the current proposer.
func (p *Proposer) L2SuggestedFeeRecipient() common.Address {
	return p.l2SuggestedFeeRecipient
}

// getTxOpts creates a bind.TransactOpts instance using the given private key.
func getTxOpts(
	ctx context.Context,
	cli *rpc.EthClient,
	privKey *ecdsa.PrivateKey,
	chainID *big.Int,
	fee *big.Int,
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

	opts.Value = fee

	return opts, nil
}

// EndpointFromConfig generates an RPC client from a given configuration.
func EndpointFromConfig(ctx context.Context, cfg *Config) (*rpc.Client, error) {
	return rpc.NewClient(ctx, &rpc.ClientConfig{
		L1Endpoint:        cfg.L1Endpoint,
		L2Endpoint:        cfg.L2Endpoint,
		TaikoL1Address:    cfg.TaikoL1Address,
		TaikoL2Address:    cfg.TaikoL2Address,
		TaikoTokenAddress: cfg.TaikoTokenAddress,
		RetryInterval:     cfg.BackOffRetryInterval,
		Timeout:           cfg.RPCTimeout,
	})
}
