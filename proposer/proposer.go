package proposer

import (
	"bytes"
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
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/internal/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	selector "github.com/taikoxyz/taiko-client/proposer/prover_selector"
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
	// RPC clients
	rpc *rpc.Client

	*Config
	// Private keys and account addresses
	proposerAddress common.Address

	proposingTimer *time.Timer

	tiers    []*rpc.TierProviderTierWithID
	tierFees []encoding.TierFee

	// Prover selector
	proverSelector selector.ProverSelector

	// Protocol configurations
	protocolConfigs *bindings.TaikoDataConfig

	// Only for testing purposes
	CustomProposeOpHook func() error
	AfterCommitHook     func() error

	ctx context.Context
	wg  sync.WaitGroup
}

// InitFromCli New initializes the given proposer instance based on the command line flags.
func (p *Proposer) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return p.InitFromConfig(ctx, cfg)
}

// InitFromConfig initializes the proposer instance based on the given configurations.
func (p *Proposer) InitFromConfig(ctx context.Context, cfg *Config) (err error) {
	p.proposerAddress = crypto.PubkeyToAddress(cfg.L1ProposerPrivKey.PublicKey)
	p.wg = sync.WaitGroup{}
	p.ctx = ctx
	p.Config = cfg

	// RPC clients
	if p.rpc, err = rpc.NewClient(p.ctx, cfg.ClientConfig); err != nil {
		return fmt.Errorf("initialize rpc clients error: %w", err)
	}

	// Protocol configs
	protocolConfigs, err := p.rpc.TaikoL1.GetConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get protocol configs: %w", err)
	}
	p.protocolConfigs = &protocolConfigs

	log.Info("Protocol configs", "configs", p.protocolConfigs)

	if p.tiers, err = p.rpc.GetTiers(ctx); err != nil {
		return err
	}
	if err := p.initTierFees(); err != nil {
		return err
	}

	if p.proverSelector, err = selector.NewETHFeeEOASelector(
		&protocolConfigs,
		p.rpc,
		cfg.TaikoL1Address,
		cfg.AssignmentHookAddress,
		p.tierFees,
		cfg.TierFeePriceBump,
		cfg.ProverEndpoints,
		cfg.MaxTierFeePriceBumps,
		proverAssignmentTimeout,
		requestProverServerTimeout,
	); err != nil {
		return err
	}

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

	var lastNonEmptyBlockProposedAt = time.Now()
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
				if p.ProposeEmptyBlocksInterval != nil {
					if time.Now().Before(lastNonEmptyBlockProposedAt.Add(*p.ProposeEmptyBlocksInterval)) {
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
		p.proposerAddress,
		baseFee,
		p.protocolConfigs.BlockMaxGasLimit,
		p.protocolConfigs.BlockMaxTxListBytes.Uint64(),
		p.LocalAddresses,
		p.MaxProposedTxListsPerEpoch,
	)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	if p.LocalAddressesOnly {
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

				for _, localAddress := range p.LocalAddresses {
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
		crypto.PubkeyToAddress(p.L1ProposerPrivKey.PublicKey),
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
				if i >= int(p.MaxProposedTxListsPerEpoch) {
					return nil
				}

				txListBytes, err := rlp.EncodeToBytes(txs)
				if err != nil {
					return fmt.Errorf("failed to encode transactions: %w", err)
				}

				txNonce := nonce + uint64(i)
				if err := p.ProposeTxList(ctx, txListBytes, uint(txs.Len()), &txNonce); err != nil {
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

func (p *Proposer) sendProposeBlockTxWithBlobHash(
	ctx context.Context,
	txListBytes []byte,
	nonce *uint64,
	isReplacement bool) (*types.Transaction, error) {
	// Make sidecar in order to get blob hash.
	sideCar, err := rpc.MakeSidecarWithSingleBlob(txListBytes)
	if err != nil {
		return nil, err
	}

	assignment, assignedProver, maxFee, err := p.proverSelector.AssignProver(
		ctx,
		p.tierFees,
		sideCar.BlobHashes()[0],
	)
	if err != nil {
		return nil, err
	}

	// Propose the transactions list
	opts, err := getTxOpts(ctx, p.rpc.L1, p.L1ProposerPrivKey, p.rpc.L1ChainID, maxFee)
	if err != nil {
		return nil, err
	}
	if nonce != nil {
		opts.Nonce = new(big.Int).SetUint64(*nonce)
	}
	opts.GasLimit = p.ProposeBlockTxGasLimit
	if isReplacement {
		if opts, err = rpc.IncreaseGasTipCap(
			ctx,
			p.rpc,
			opts,
			p.proposerAddress,
			new(big.Int).SetUint64(p.ProposeBlockTxReplacementMultiplier),
			p.ProposeBlockTxGasTipCap,
		); err != nil {
			return nil, err
		}
	}

	var parentMetaHash = [32]byte{}
	if p.IncludeParentMetaHash {
		state, err := p.rpc.TaikoL1.State(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, err
		}

		parent, err := p.rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, state.SlotB.NumBlocks-1)
		if err != nil {
			return nil, err
		}

		parentMetaHash = parent.MetaHash
	}

	hookCalls := make([]encoding.HookCall, 0)

	// initially just use the AssignmentHook default.
	// TODO: flag for additional hook addresses and data.
	hookInputData, err := encoding.EncodeAssignmentHookInput(&encoding.AssignmentHookInput{
		Assignment: assignment,
		Tip:        common.Big0, // TODO: flag for tip
	})
	if err != nil {
		return nil, err
	}

	hookCalls = append(hookCalls, encoding.HookCall{
		Hook: p.AssignmentHookAddress,
		Data: hookInputData,
	})
	encodedParams, err := encoding.EncodeBlockParams(&encoding.BlockParams{
		AssignedProver:    assignedProver,
		ExtraData:         rpc.StringToBytes32(p.ExtraData),
		TxListByteOffset:  common.Big0,
		TxListByteSize:    big.NewInt(int64(len(txListBytes))),
		BlobHash:          [32]byte{},
		CacheBlobForReuse: false,
		ParentMetaHash:    parentMetaHash,
		HookCalls:         hookCalls,
	})
	if err != nil {
		return nil, err
	}

	opts.NoSend = true
	rawTx, err := p.rpc.TaikoL1.ProposeBlock(
		opts,
		encodedParams,
		nil,
	)
	if err != nil {
		return nil, encoding.TryParsingCustomError(err)
	}

	// Create blob tx and send it.
	opts.NoSend = false
	proposeTx, err := p.rpc.L1.TransactBlobTx(opts, &p.TaikoL1Address, rawTx.Data(), sideCar)
	if err != nil {
		return nil, err
	}

	log.Info("Tx", "proposeBlockTx", proposeTx, "type", proposeTx.Type())

	return proposeTx, nil
}

// sendProposeBlockTx tries to send a TaikoL1.proposeBlock transaction.
func (p *Proposer) sendProposeBlockTx(
	ctx context.Context,
	txListBytes []byte,
	nonce *uint64,
	isReplacement bool,
) (*types.Transaction, error) {
	assignment, assignedProver, maxFee, err := p.proverSelector.AssignProver(
		ctx,
		p.tierFees,
		crypto.Keccak256Hash(txListBytes),
	)
	if err != nil {
		return nil, err
	}

	// Propose the transactions list
	opts, err := getTxOpts(ctx, p.rpc.L1, p.L1ProposerPrivKey, p.rpc.L1ChainID, maxFee)
	if err != nil {
		return nil, err
	}
	if nonce != nil {
		opts.Nonce = new(big.Int).SetUint64(*nonce)
	}
	opts.GasLimit = p.ProposeBlockTxGasLimit
	if isReplacement {
		if opts, err = rpc.IncreaseGasTipCap(
			ctx,
			p.rpc,
			opts,
			p.proposerAddress,
			new(big.Int).SetUint64(p.ProposeBlockTxReplacementMultiplier),
			p.ProposeBlockTxGasTipCap,
		); err != nil {
			return nil, err
		}
	}

	var parentMetaHash = [32]byte{}
	if p.IncludeParentMetaHash {
		state, err := p.rpc.TaikoL1.State(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, err
		}

		parent, err := p.rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, state.SlotB.NumBlocks-1)
		if err != nil {
			return nil, err
		}

		parentMetaHash = parent.MetaHash
	}

	hookCalls := make([]encoding.HookCall, 0)

	// initially just use the AssignmentHook default.
	// TODO: flag for additional hook addresses and data.
	hookInputData, err := encoding.EncodeAssignmentHookInput(&encoding.AssignmentHookInput{
		Assignment: assignment,
		Tip:        p.L1BlockBuilderTip,
	})
	if err != nil {
		return nil, err
	}

	hookCalls = append(hookCalls, encoding.HookCall{
		Hook: p.AssignmentHookAddress,
		Data: hookInputData,
	})

	encodedParams, err := encoding.EncodeBlockParams(&encoding.BlockParams{
		AssignedProver:    assignedProver,
		ExtraData:         rpc.StringToBytes32(p.ExtraData),
		TxListByteOffset:  common.Big0,
		TxListByteSize:    common.Big0,
		BlobHash:          [32]byte{},
		CacheBlobForReuse: false,
		ParentMetaHash:    parentMetaHash,
		HookCalls:         hookCalls,
	})
	if err != nil {
		return nil, err
	}

	proposeTx, err := p.rpc.TaikoL1.ProposeBlock(
		opts,
		encodedParams,
		txListBytes,
	)
	if err != nil {
		return nil, encoding.TryParsingCustomError(err)
	}

	return proposeTx, nil
}

// ProposeTxList proposes the given transactions list to TaikoL1 smart contract.
func (p *Proposer) ProposeTxList(
	ctx context.Context,
	txListBytes []byte,
	txNum uint,
	nonce *uint64,
) error {
	var (
		isReplacement bool
		tx            *types.Transaction
		err           error
	)
	if err = backoff.Retry(
		func() error {
			if ctx.Err() != nil {
				return nil
			}
			// Send tx list by blob tx.
			if p.BlobAllowed {
				tx, err = p.sendProposeBlockTxWithBlobHash(
					ctx,
					txListBytes,
					nonce,
					isReplacement,
				)
			} else {
				tx, err = p.sendProposeBlockTx(
					ctx,
					txListBytes,
					nonce,
					isReplacement,
				)
			}
			if err != nil {
				log.Warn("Failed to send taikoL1.proposeBlock transaction", "error", encoding.TryParsingCustomError(err))
				if strings.Contains(err.Error(), core.ErrNonceTooLow.Error()) {
					return nil
				}
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

	ctxWithTimeout, cancel := context.WithTimeout(ctx, p.WaitReceiptTimeout)
	defer cancel()
	if tx != nil {
		if _, err = rpc.WaitReceipt(ctxWithTimeout, p.rpc.L1, tx); err != nil {
			return err
		}
	}

	log.Info("📝 Propose transactions succeeded", "txs", txNum)

	metrics.ProposerProposedTxListsCounter.Inc(1)
	metrics.ProposerProposedTxsCounter.Inc(int64(txNum))

	return nil
}

// ProposeEmptyBlockOp performs a proposing one empty block operation.
func (p *Proposer) ProposeEmptyBlockOp(ctx context.Context) error {
	emptyTxListBytes, err := rlp.EncodeToBytes(types.Transactions{})
	if err != nil {
		return err
	}
	return p.ProposeTxList(ctx, emptyTxListBytes, 0, nil)
}

// updateProposingTicker updates the internal proposing timer.
func (p *Proposer) updateProposingTicker() {
	if p.proposingTimer != nil {
		p.proposingTimer.Stop()
	}

	var duration time.Duration
	if p.ProposeInterval != nil {
		duration = *p.ProposeInterval
	} else {
		// Random number between 12 - 120
		randomSeconds := rand.Intn(120-11) + 12 // nolint: gosec
		duration = time.Duration(randomSeconds) * time.Second
	}

	p.proposingTimer = time.NewTimer(duration)
}

// Name returns the application name.
func (p *Proposer) Name() string {
	return "proposer"
}

// initTierFees initializes the proving fees for every proof tier configured in the protocol for the proposer.
func (p *Proposer) initTierFees() error {
	for _, tier := range p.tiers {
		log.Info(
			"Protocol tier",
			"id", tier.ID,
			"name", string(bytes.TrimRight(tier.VerifierName[:], "\x00")),
			"validityBond", tier.ValidityBond,
			"contestBond", tier.ContestBond,
			"provingWindow", tier.ProvingWindow,
			"cooldownWindow", tier.CooldownWindow,
		)

		switch tier.ID {
		case encoding.TierOptimisticID:
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: p.OptimisticTierFee})
		case encoding.TierSgxID:
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: p.SgxTierFee})
		case encoding.TierPseZkevmID:
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: p.PseZkevmTierFee})
		case encoding.TierSgxAndPseZkevmID:
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: p.SgxAndPseZkevmTierFee})
		case encoding.TierGuardianID:
			// Guardian prover should not charge any fee.
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: common.Big0})
		default:
			return fmt.Errorf("unknown tier: %d", tier.ID)
		}
	}

	return nil
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
