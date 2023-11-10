package prover

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	capacity "github.com/taikoxyz/taiko-client/prover/capacity_manager"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	proofSubmitter "github.com/taikoxyz/taiko-client/prover/proof_submitter"
	"github.com/taikoxyz/taiko-client/prover/server"
	"github.com/urfave/cli/v2"
)

var (
	errNoCapacity   = errors.New("no prover capacity available")
	errTierNotFound = errors.New("tier not found")
)

// Prover keep trying to prove new proposed blocks valid/invalid.
type Prover struct {
	// Configurations
	cfg           *Config
	proverAddress common.Address

	// Clients
	rpc *rpc.Client

	// Prover Server
	srv *server.ProverServer

	// Contract configurations
	protocolConfigs *bindings.TaikoDataConfig

	// States
	latestVerifiedL1Height uint64
	lastHandledBlockID     uint64
	genesisHeightL1        uint64
	l1Current              *types.Header
	reorgDetectedFlag      bool
	tiers                  []*rpc.TierProviderTierWithID

	// Proof submitters
	proofSubmitters []proofSubmitter.Submitter
	proofContester  proofSubmitter.Contester

	// Subscriptions
	blockProposedCh        chan *bindings.TaikoL1ClientBlockProposed
	blockProposedSub       event.Subscription
	transitionProvedCh     chan *bindings.TaikoL1ClientTransitionProved
	transitionProvedSub    event.Subscription
	transitionContestedCh  chan *bindings.TaikoL1ClientTransitionContested
	transitionContestedSub event.Subscription
	blockVerifiedCh        chan *bindings.TaikoL1ClientBlockVerified
	blockVerifiedSub       event.Subscription
	proofWindowExpiredCh   chan *bindings.TaikoL1ClientBlockProposed
	proveNotify            chan struct{}

	// Proof related
	proofGenerationCh chan *proofProducer.ProofWithHeader

	// Concurrency guards
	proposeConcurrencyGuard     chan struct{}
	submitProofConcurrencyGuard chan struct{}

	// capacity-related configs
	capacityManager *capacity.CapacityManager

	ctx context.Context
	wg  sync.WaitGroup
}

// InitFromCli initializes the given prover instance based on the command line flags.
func (p *Prover) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return InitFromConfig(ctx, p, cfg)
}

// InitFromConfig initializes the prover instance based on the given configurations.
func InitFromConfig(ctx context.Context, p *Prover, cfg *Config) (err error) {
	p.cfg = cfg
	p.ctx = ctx
	p.capacityManager = capacity.New(cfg.Capacity)

	// Clients
	if p.rpc, err = rpc.NewClient(p.ctx, &rpc.ClientConfig{
		L1Endpoint:            cfg.L1WsEndpoint,
		L2Endpoint:            cfg.L2WsEndpoint,
		TaikoL1Address:        cfg.TaikoL1Address,
		TaikoL2Address:        cfg.TaikoL2Address,
		TaikoTokenAddress:     cfg.TaikoTokenAddress,
		GuardianProverAddress: cfg.GuardianProverAddress,
		RetryInterval:         cfg.BackOffRetryInterval,
		Timeout:               cfg.RPCTimeout,
		BackOffMaxRetrys:      new(big.Int).SetUint64(p.cfg.BackOffMaxRetrys),
	}); err != nil {
		return err
	}

	// Configs
	protocolConfigs, err := p.rpc.TaikoL1.GetConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get protocol configs: %w", err)
	}
	p.protocolConfigs = &protocolConfigs

	log.Info("Protocol configs", "configs", p.protocolConfigs)

	p.proverAddress = crypto.PubkeyToAddress(p.cfg.L1ProverPrivKey.PublicKey)

	chBufferSize := p.protocolConfigs.BlockMaxProposals
	p.blockProposedCh = make(chan *bindings.TaikoL1ClientBlockProposed, chBufferSize)
	p.blockVerifiedCh = make(chan *bindings.TaikoL1ClientBlockVerified, chBufferSize)
	p.transitionProvedCh = make(chan *bindings.TaikoL1ClientTransitionProved, chBufferSize)
	p.transitionContestedCh = make(chan *bindings.TaikoL1ClientTransitionContested, chBufferSize)
	p.proofGenerationCh = make(chan *proofProducer.ProofWithHeader, chBufferSize)
	p.proofWindowExpiredCh = make(chan *bindings.TaikoL1ClientBlockProposed, chBufferSize)
	p.proveNotify = make(chan struct{}, 1)
	if err := p.initL1Current(cfg.StartingBlockID); err != nil {
		return fmt.Errorf("initialize L1 current cursor error: %w", err)
	}

	// Concurrency guards
	p.proposeConcurrencyGuard = make(chan struct{}, cfg.MaxConcurrentProvingJobs)
	p.submitProofConcurrencyGuard = make(chan struct{}, cfg.MaxConcurrentProvingJobs)

	// Protocol proof tiers
	if p.tiers, err = p.rpc.GetTiers(ctx); err != nil {
		return err
	}

	// Proof submitters
	for _, tier := range p.tiers {
		var (
			producer  proofProducer.ProofProducer
			submitter proofSubmitter.Submitter
		)
		switch tier.ID {
		case encoding.TierOptimisticID:
			producer = &proofProducer.OptimisticProofProducer{DummyProofProducer: new(proofProducer.DummyProofProducer)}
		case encoding.TierSgxID:
			producer = &proofProducer.SGXProofProducer{DummyProofProducer: new(proofProducer.DummyProofProducer)}
		case encoding.TierSgxAndPseZkevmID:
			zkEvmRpcdProducer, err := proofProducer.NewZkevmRpcdProducer(
				cfg.ZKEvmRpcdEndpoint,
				cfg.ZkEvmRpcdParamsPath,
				cfg.L1HttpEndpoint,
				cfg.L2HttpEndpoint,
				true,
				p.protocolConfigs,
			)
			if err != nil {
				return err
			}
			if p.cfg.Dummy {
				zkEvmRpcdProducer.DummyProofProducer = new(proofProducer.DummyProofProducer)
			}
			producer = zkEvmRpcdProducer
		case encoding.TierGuardianID:
			producer = &proofProducer.GuardianProofProducer{DummyProofProducer: new(proofProducer.DummyProofProducer)}
		}

		if submitter, err = proofSubmitter.New(
			p.rpc,
			producer,
			p.proofGenerationCh,
			p.cfg.TaikoL2Address,
			p.cfg.L1ProverPrivKey,
			p.cfg.Graffiti,
			p.cfg.ProofSubmissionMaxRetry,
			p.cfg.BackOffRetryInterval,
			p.cfg.WaitReceiptTimeout,
			p.cfg.ProveBlockGasLimit,
			p.cfg.ProveBlockTxReplacementMultiplier,
			p.cfg.ProveBlockMaxTxGasTipCap,
		); err != nil {
			return err
		}

		p.proofSubmitters = append(p.proofSubmitters, submitter)
	}

	// Proof contester
	p.proofContester, err = proofSubmitter.NewProofContester(
		p.rpc,
		p.cfg.L1ProverPrivKey,
		p.cfg.ProveBlockGasLimit,
		p.cfg.ProveBlockTxReplacementMultiplier,
		p.cfg.ProveBlockMaxTxGasTipCap,
		p.cfg.ProofSubmissionMaxRetry,
		p.cfg.BackOffRetryInterval,
		p.cfg.WaitReceiptTimeout,
		p.cfg.Graffiti,
	)
	if err != nil {
		return err
	}

	// Prover server
	proverServerOpts := &server.NewProverServerOpts{
		ProverPrivateKey:         p.cfg.L1ProverPrivKey,
		MinOptimisticTierFee:     p.cfg.MinOptimisticTierFee,
		MinSgxTierFee:            p.cfg.MinSgxTierFee,
		MinPseZkevmTierFee:       p.cfg.MinPseZkevmTierFee,
		MinSgxAndPseZkevmTierFee: p.cfg.MinSgxAndPseZkevmTierFee,
		MaxExpiry:                p.cfg.MaxExpiry,
		MaxBlockSlippage:         p.cfg.MaxBlockSlippage,
		CapacityManager:          p.capacityManager,
		TaikoL1Address:           p.cfg.TaikoL1Address,
		Rpc:                      p.rpc,
		LivenessBond:             protocolConfigs.LivenessBond,
		IsGuardian:               p.IsGuardianProver(),
	}
	if p.IsGuardianProver() {
		proverServerOpts.ProverPrivateKey = p.cfg.GuardianProverPrivateKey
	}
	if p.srv, err = server.New(proverServerOpts); err != nil {
		return err
	}

	return nil
}

// Start starts the main loop of the L2 block prover.
func (p *Prover) Start() error {
	p.wg.Add(1)
	p.initSubscription()
	go func() {
		if err := p.srv.Start(fmt.Sprintf(":%v", p.cfg.HTTPServerPort)); !errors.Is(err, http.ErrServerClosed) {
			log.Crit("Failed to start http server", "error", err)
		}
	}()
	go p.eventLoop()

	return nil
}

// eventLoop starts the main loop of Taiko prover.
func (p *Prover) eventLoop() {
	defer func() {
		p.wg.Done()
	}()

	// reqProving requests performing a proving operation, won't block
	// if we are already proving.
	reqProving := func() {
		select {
		case p.proveNotify <- struct{}{}:
		default:
		}
	}

	// If there is too many (TaikoData.Config.blockMaxProposals) pending blocks in TaikoL1 contract, there will be no new
	// BlockProposed temporarily, so except the BlockProposed subscription, we need another trigger to start
	// fetching the proposed blocks.
	forceProvingTicker := time.NewTicker(15 * time.Second)
	defer forceProvingTicker.Stop()

	// Call reqProving() right away to catch up with the latest state.
	reqProving()

	for {
		select {
		case <-p.ctx.Done():
			return
		case proofWithHeader := <-p.proofGenerationCh:
			p.submitProofOp(p.ctx, proofWithHeader)
		case <-p.proveNotify:
			if err := p.proveOp(); err != nil {
				log.Error("Prove new blocks error", "error", err)
			}
		case e := <-p.blockVerifiedCh:
			if err := p.onBlockVerified(p.ctx, e); err != nil {
				log.Error("Handle BlockVerified event error", "error", err)
			}
		case e := <-p.transitionProvedCh:
			if err := p.onTransitionProved(p.ctx, e); err != nil {
				log.Error("Handle TransitionProved event error", "error", err)
			}
		case e := <-p.transitionContestedCh:
			if err := p.onTransitionContested(p.ctx, e); err != nil {
				log.Error("Handle TransitionContested event error", "error", err)
			}
		case e := <-p.proofWindowExpiredCh:
			if err := p.onProvingWindowExpired(p.ctx, e); err != nil {
				log.Error("Handle provingWindow expired event error", "error", err)
			}
		case <-p.blockProposedCh:
			reqProving()
		case <-forceProvingTicker.C:
			reqProving()
		}
	}
}

// Close closes the prover instance.
func (p *Prover) Close(ctx context.Context) {
	p.closeSubscription()
	if err := p.srv.Shutdown(ctx); err != nil {
		log.Error("Failed to shut down prover server", "error", err)
	}
	p.wg.Wait()
}

// proveOp performs a proving operation, find current unproven blocks, then
// request generating proofs for them.
func (p *Prover) proveOp() error {
	firstTry := true

	for firstTry || p.reorgDetectedFlag {
		p.reorgDetectedFlag = false
		firstTry = false

		iter, err := eventIterator.NewBlockProposedIterator(p.ctx, &eventIterator.BlockProposedIteratorConfig{
			Client:               p.rpc.L1,
			TaikoL1:              p.rpc.TaikoL1,
			StartHeight:          new(big.Int).SetUint64(p.l1Current.Number.Uint64()),
			OnBlockProposedEvent: p.onBlockProposed,
		})
		if err != nil {
			log.Error("Failed to start event iterator", "event", "BlockProposed", "error", err)
			return err
		}

		if err := iter.Iter(); err != nil {
			return err
		}
	}

	return nil
}

// onBlockProposed tries to prove that the newly proposed block is valid/invalid.
func (p *Prover) onBlockProposed(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	end eventIterator.EndBlockProposedEventIterFunc,
) error {
	// If there are newly generated proofs, we need to submit them as soon as possible.
	if len(p.proofGenerationCh) > 0 {
		log.Info("onBlockProposed early return", "proofGenerationChannelLength", len(p.proofGenerationCh))
		end()
		return nil
	}

	if _, err := p.rpc.WaitL1Origin(ctx, event.BlockId); err != nil {
		return fmt.Errorf("failed to wait L1Origin (eventID %d): %w", event.BlockId, err)
	}

	// Check whether the L2 EE's anchored L1 info, to see if the L1 chain has been reorged.
	reorged, l1CurrentToReset, lastHandledBlockIDToReset, err := p.rpc.CheckL1ReorgFromL2EE(
		ctx,
		new(big.Int).Sub(event.BlockId, common.Big1),
	)
	if err != nil {
		return fmt.Errorf("failed to check whether L1 chain was reorged from L2EE (eventID %d): %w", event.BlockId, err)
	}

	// Then check the l1Current cursor at first, to see if the L1 chain has been reorged.
	if !reorged {
		if reorged, l1CurrentToReset, lastHandledBlockIDToReset, err = p.rpc.CheckL1ReorgFromL1Cursor(
			ctx,
			p.l1Current,
			p.genesisHeightL1,
		); err != nil {
			return fmt.Errorf(
				"failed to check whether L1 chain was reorged from l1Current (eventID %d): %w",
				event.BlockId,
				err,
			)
		}
	}

	if reorged {
		log.Info(
			"Reset L1Current cursor due to reorg",
			"l1CurrentHeightOld", p.l1Current,
			"l1CurrentHeightNew", l1CurrentToReset.Number,
			"lastHandledBlockIDOld", p.lastHandledBlockID,
			"lastHandledBlockIDNew", lastHandledBlockIDToReset,
		)
		p.l1Current = l1CurrentToReset
		if lastHandledBlockIDToReset == nil {
			p.lastHandledBlockID = 0
		} else {
			p.lastHandledBlockID = lastHandledBlockIDToReset.Uint64()
		}
		p.reorgDetectedFlag = true
		end()
		return nil
	}

	if event.BlockId.Uint64() <= p.lastHandledBlockID {
		return nil
	}

	lastL1OriginHeader, err := p.rpc.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(event.Meta.L1Height))
	if err != nil {
		return fmt.Errorf("failed to get L1 header, height %d: %w", event.Meta.L1Height, err)
	}

	if lastL1OriginHeader.Hash() != event.Meta.L1Hash {
		log.Warn(
			"L1 block hash mismatch due to L1 reorg",
			"height", event.Meta.L1Height,
			"lastL1OriginHeader", lastL1OriginHeader.Hash(),
			"l1HashInEvent", event.Meta.L1Hash,
		)

		return fmt.Errorf(
			"L1 block hash mismatch due to L1 reorg: %s != %s",
			lastL1OriginHeader.Hash(),
			event.Meta.L1Hash,
		)
	}

	log.Info(
		"Proposed block",
		"l1Height", event.Raw.BlockNumber,
		"l1Hash", event.Raw.BlockHash,
		"blockID", event.BlockId,
		"removed", event.Raw.Removed,
	)
	metrics.ProverReceivedProposedBlockGauge.Update(event.BlockId.Int64())

	handleBlockProposedEvent := func() error {
		defer func() { <-p.proposeConcurrencyGuard }()

		// Check whether the block has been verified.
		isVerified, err := p.isBlockVerified(event.BlockId)
		if err != nil {
			return fmt.Errorf("failed to check if the current L2 block is verified: %w", err)
		}
		if isVerified {
			log.Info("📋 Block has been verified", "blockID", event.BlockId)
			return nil
		}

		// Check whether the block's proof is still needed.
		proofStatus, err := rpc.GetBlockProofStatus(
			p.ctx,
			p.rpc,
			event.BlockId,
			p.proverAddress,
		)
		if err != nil {
			return fmt.Errorf("failed to check whether the L2 block needs a new proof: %w", err)
		}

		if proofStatus.IsSubmitted {
			// If there is already a proof submitted and there is no need to contest
			// it, we skip proving this block here.
			if !proofStatus.Invalid || !p.cfg.ContesterMode {
				return nil
			}

			// The proof submitted to protocol is invalid.
			return p.handleInvalidProof(
				ctx, event.BlockId,
				new(big.Int).SetUint64(event.Raw.BlockNumber),
				proofStatus.ParentBlock.Hash(),
				proofStatus.CurrentTransitionState.Contester,
				&event.Meta,
				proofStatus.CurrentTransitionState.Tier,
			)
		}

		log.Info(
			"Proposed block information",
			"blockID", event.BlockId,
			"assignedProver", event.AssignedProver,
			"minTier", event.Meta.MinTier,
		)

		provingWindow, err := p.getProvingWindow(event)
		if err != nil {
			return fmt.Errorf("failed to get proving window: %w", err)
		}

		var (
			now                    = uint64(time.Now().Unix())
			provingWindowExpiresAt = event.Meta.Timestamp + uint64(provingWindow.Seconds())
			provingWindowExpired   = now > provingWindowExpiresAt
			timeToExpire           = time.Duration(provingWindowExpiresAt-now) * time.Second
		)

		if provingWindowExpired {
			// If the proving window is expired, we need to check if the current prover is the assigned prover
			// at first, if yes, we should skip proving this block, if no, then we check if the current prover
			// wants to prove unassigned blocks.
			log.Info(
				"Proposed block's proving window has expired",
				"blockID", event.BlockId,
				"prover", event.AssignedProver,
				"expiresAt", provingWindowExpiresAt,
			)
			if event.AssignedProver == p.proverAddress {
				log.Warn(
					"Assigned prover is the current prover, but the proving window has expired, skip proving",
					"blockID", event.BlockId,
					"prover", event.AssignedProver,
					"expiresAt", provingWindowExpiresAt,
				)
				return nil
			}
			if !p.cfg.ProveUnassignedBlocks {
				log.Info(
					"Skip proving expired blocks",
					"blockID", event.BlockId,
					"prover", event.AssignedProver,
					"expiresAt", provingWindowExpiresAt,
				)
				return nil
			}
		} else {
			// If the proving window is not expired, we need to check if the current prover is the assigned prover,
			// if no and the current prover wants to prove unassigned blocks, then we should wait for its expiration.
			if event.AssignedProver != p.proverAddress {
				log.Info(
					"Proposed block is not provable",
					"blockID", event.BlockId,
					"prover", event.AssignedProver,
					"expiresAt", provingWindowExpiresAt,
					"timeToExpire", timeToExpire,
				)

				if p.cfg.ProveUnassignedBlocks {
					log.Info(
						"Add proposed block to wait for proof window expiration",
						"blockID", event.BlockId,
					)
					time.AfterFunc(
						// Add another 12 seconds, to ensure one more L1 block will be mined before the proof submission
						timeToExpire+12*time.Second,
						func() { p.proofWindowExpiredCh <- event },
					)
				}

				return nil
			}
		}

		log.Info(
			"Proposed block is provable",
			"blockID", event.BlockId,
			"prover", event.AssignedProver,
			"expiresAt", provingWindowExpiresAt,
			"minTier", event.Meta.MinTier,
		)

		metrics.ProverProofsAssigned.Inc(1)

		// Make sure to take a capacity before requesting proof.
		if err := p.takeOneCapacity(event.BlockId); err != nil {
			return err
		}

		if proofSubmitter := p.selectSubmitter(event.Meta.MinTier); proofSubmitter != nil {
			return proofSubmitter.RequestProof(ctx, event)
		}

		return nil
	}

	// Move l1Current cursor.
	newL1Current, err := p.rpc.L1.HeaderByHash(ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}
	p.l1Current = newL1Current
	p.lastHandledBlockID = event.BlockId.Uint64()

	// Try generating a proof for the proposed block with the given backoff policy.
	go func() {
		if err := backoff.Retry(
			func() error {
				p.proposeConcurrencyGuard <- struct{}{}

				if err := handleBlockProposedEvent(); err != nil {
					log.Error(
						"Failed to handle BlockProposed event",
						"error", err,
						"blockID", event.BlockId,
						"minTier", event.Meta.MinTier,
						"maxRetrys", p.cfg.BackOffMaxRetrys,
					)
					return err
				}
				return nil
			},
			backoff.WithMaxRetries(backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval), p.cfg.BackOffMaxRetrys),
		); err != nil {
			log.Error("Handle new BlockProposed event error", "error", err)
		}
	}()

	return nil
}

// handleInvalidProof handles the case when the proof submitted to protocol is invalid.
func (p *Prover) handleInvalidProof(
	ctx context.Context,
	blockID *big.Int,
	proposedIn *big.Int,
	parentHash common.Hash,
	contester common.Address,
	meta *bindings.TaikoDataBlockMetadata,
	tier uint16,
) error {
	// The proof submitted to protocol is invalid.
	log.Info(
		"Invalid proof detected",
		"blockID", blockID,
		"parent", parentHash,
	)

	// If there is no contester, we submit a contest to protocol.
	if contester == rpc.ZeroAddress {
		log.Info(
			"Try submitting a contest",
			"blockID", blockID,
			"parent", parentHash,
		)

		return p.proofContester.SubmitContest(ctx, blockID, proposedIn, parentHash, meta, tier)
	}

	log.Info(
		"Try submitting a higher tier proof",
		"blockID", blockID,
		"parent", parentHash,
	)

	// If there is already a contester, we try submitting a proof with a higher tier here.
	return p.requestProofByBlockID(blockID, proposedIn, tier+1, nil)
}

// submitProofOp performs a proof submission operation.
func (p *Prover) submitProofOp(ctx context.Context, proofWithHeader *proofProducer.ProofWithHeader) {
	go func() {
		p.submitProofConcurrencyGuard <- struct{}{}

		defer func() {
			<-p.submitProofConcurrencyGuard
			p.releaseOneCapacity(proofWithHeader.BlockID)
		}()

		if err := backoff.Retry(
			func() error {
				proofSubmitter := p.getSubmitterByTier(proofWithHeader.Tier)
				if proofSubmitter == nil {
					return nil
				}

				if err := proofSubmitter.SubmitProof(p.ctx, proofWithHeader); err != nil {
					log.Error("Submit proof error", "error", err)
					return err
				}

				return nil
			},
			backoff.WithMaxRetries(backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval), p.cfg.BackOffMaxRetrys),
		); err != nil {
			log.Error("Submit proof error", "error", err)
		}
	}()
}

// onTransitionContested tries to submit a higher tier proof for the contested transition.
func (p *Prover) onTransitionContested(ctx context.Context, e *bindings.TaikoL1ClientTransitionContested) error {
	log.Info(
		"🗡 Transition contested",
		"blockID", e.BlockId,
		"parentHash", common.Bytes2Hex(e.Tran.ParentHash[:]),
		"hash", common.Bytes2Hex(e.Tran.BlockHash[:]),
		"signalRoot", common.BytesToHash(e.Tran.SignalRoot[:]),
		"contester", e.Contester,
		"bond", e.ContestBond,
	)

	// If this prover is not in contester mode, we simply output a log and return.
	if !p.cfg.ContesterMode {
		return nil
	}

	contestedTransition, err := p.rpc.TaikoL1.GetTransition(
		&bind.CallOpts{Context: ctx},
		e.BlockId.Uint64(),
		e.Tran.ParentHash,
	)
	if err != nil {
		return err
	}

	// Compare the contested transition to the block in local L2 canonical chain.
	isValidProof, err := p.isValidProof(
		ctx,
		e.BlockId,
		e.Tran.ParentHash,
		contestedTransition.BlockHash,
		contestedTransition.SignalRoot,
	)
	if err != nil {
		return err
	}
	if isValidProof {
		log.Info(
			"Contested transition is valid to local canonical chain, ignore the contest",
			"blockID", e.BlockId,
			"parentHash", common.Bytes2Hex(e.Tran.ParentHash[:]),
			"hash", common.Bytes2Hex(contestedTransition.BlockHash[:]),
			"signalRoot", common.BytesToHash(contestedTransition.SignalRoot[:]),
			"contester", e.Contester,
			"bond", e.ContestBond,
		)
		return nil
	}

	blockInfo, err := p.rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, e.BlockId.Uint64())
	if err != nil {
		return err
	}

	return p.requestProofByBlockID(e.BlockId, new(big.Int).SetUint64(blockInfo.ProposedIn), e.Tier+1, nil)
}

// onBlockVerified update the latestVerified block in current state, and cancels
// the block being proven if it's verified.
func (p *Prover) onBlockVerified(ctx context.Context, e *bindings.TaikoL1ClientBlockVerified) error {
	metrics.ProverLatestVerifiedIDGauge.Update(e.BlockId.Int64())

	p.latestVerifiedL1Height = e.Raw.BlockNumber

	log.Info(
		"New verified block",
		"blockID", e.BlockId,
		"hash", common.BytesToHash(e.BlockHash[:]),
		"signalRoot", common.BytesToHash(e.SignalRoot[:]),
		"assignedProver", e.AssignedProver,
		"prover", e.Prover,
	)

	return nil
}

// onTransitionProved verifies the proven block hash and will try contesting it if the block hash is wrong.
func (p *Prover) onTransitionProved(ctx context.Context, event *bindings.TaikoL1ClientTransitionProved) error {
	metrics.ProverReceivedProvenBlockGauge.Update(event.BlockId.Int64())

	// If the proof generation is cancellable, cancel it and release the capacity.
	proofSubmitter := p.getSubmitterByTier(event.Tier)
	if proofSubmitter != nil && proofSubmitter.Producer().Cancellable() {
		if err := proofSubmitter.Producer().Cancel(ctx, event.BlockId); err != nil {
			return err
		}
		// No need to check if the release is successful here, since this L2 block might
		// be assigned to other provers.
		p.capacityManager.ReleaseOneCapacity(event.BlockId.Uint64())
	}

	// If this prover is in contest mode, we check the validity of this proof and if it's invalid,
	// contest it with a higher tier proof.
	if !p.cfg.ContesterMode {
		return nil
	}

	isValidProof, err := p.isValidProof(
		ctx,
		event.BlockId,
		event.Tran.ParentHash,
		event.Tran.BlockHash,
		event.Tran.SignalRoot,
	)
	if err != nil {
		return err
	}
	if isValidProof {
		return nil
	}

	blockInfo, err := p.rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, event.BlockId.Uint64())
	if err != nil {
		return err
	}

	log.Info(
		"Contest a proven transition",
		"blockID", event.BlockId,
		"l1Height", blockInfo.ProposedIn,
		"tier", event.Tier,
		"parentHash", common.Bytes2Hex(event.Tran.ParentHash[:]),
		"blockHash", common.Bytes2Hex(event.Tran.BlockHash[:]),
		"signalRoot", common.Bytes2Hex(event.Tran.SignalRoot[:]),
	)

	return p.requestProofByBlockID(event.BlockId, new(big.Int).SetUint64(blockInfo.ProposedIn), event.Tier, event)
}

// Name returns the application name.
func (p *Prover) Name() string {
	return "prover"
}

// initL1Current initializes prover's L1Current cursor.
func (p *Prover) initL1Current(startingBlockID *big.Int) error {
	if err := p.rpc.WaitTillL2ExecutionEngineSynced(p.ctx); err != nil {
		return err
	}

	stateVars, err := p.rpc.GetProtocolStateVariables(&bind.CallOpts{Context: p.ctx})
	if err != nil {
		return err
	}
	p.genesisHeightL1 = stateVars.A.GenesisHeight

	if startingBlockID == nil {
		if stateVars.B.LastVerifiedBlockId == 0 {
			genesisL1Header, err := p.rpc.L1.HeaderByNumber(p.ctx, new(big.Int).SetUint64(stateVars.A.GenesisHeight))
			if err != nil {
				return err
			}

			p.l1Current = genesisL1Header
			return nil
		}

		startingBlockID = new(big.Int).SetUint64(stateVars.B.LastVerifiedBlockId)
	}

	log.Info("Init L1Current cursor", "startingBlockID", startingBlockID)

	latestVerifiedHeaderL1Origin, err := p.rpc.L2.L1OriginByID(p.ctx, startingBlockID)
	if err != nil {
		if err.Error() == ethereum.NotFound.Error() {
			log.Warn("Failed to find L1Origin for blockID, use latest L1 head instead", "blockID", startingBlockID)
			l1Head, err := p.rpc.L1.HeaderByNumber(p.ctx, nil)
			if err != nil {
				return err
			}

			p.l1Current = l1Head
			return nil
		}
		return err
	}

	if p.l1Current, err = p.rpc.L1.HeaderByHash(p.ctx, latestVerifiedHeaderL1Origin.L1BlockHash); err != nil {
		return err
	}

	return nil
}

// isBlockVerified checks whether the given L2 block has been verified.
func (p *Prover) isBlockVerified(id *big.Int) (bool, error) {
	stateVars, err := p.rpc.GetProtocolStateVariables(&bind.CallOpts{Context: p.ctx})
	if err != nil {
		return false, err
	}

	return id.Uint64() <= stateVars.B.LastVerifiedBlockId, nil
}

// initSubscription initializes all subscriptions in current prover instance.
func (p *Prover) initSubscription() {
	p.blockProposedSub = rpc.SubscribeBlockProposed(p.rpc.TaikoL1, p.blockProposedCh)
	p.blockVerifiedSub = rpc.SubscribeBlockVerified(p.rpc.TaikoL1, p.blockVerifiedCh)
	p.transitionProvedSub = rpc.SubscribeTransitionProved(p.rpc.TaikoL1, p.transitionProvedCh)
	p.transitionContestedSub = rpc.SubscribeTransitionContested(p.rpc.TaikoL1, p.transitionContestedCh)
}

// closeSubscription closes all subscriptions.
func (p *Prover) closeSubscription() {
	p.blockVerifiedSub.Unsubscribe()
	p.blockProposedSub.Unsubscribe()
	p.transitionProvedSub.Unsubscribe()
	p.transitionContestedSub.Unsubscribe()
}

// isValidProof checks if the given proof is a valid one, comparing to current L2 node canonical chain.
func (p *Prover) isValidProof(
	ctx context.Context,
	blockID *big.Int,
	parentHash common.Hash,
	blockHash common.Hash,
	signalRoot common.Hash,
) (bool, error) {
	parent, err := p.rpc.L2ParentByBlockId(ctx, blockID)
	if err != nil {
		return false, err
	}

	block, err := p.rpc.L2.BlockByNumber(ctx, blockID)
	if err != nil {
		return false, err
	}

	l2SignalService, err := p.rpc.TaikoL2.SignalService(
		&bind.CallOpts{Context: ctx, BlockNumber: blockID},
	)
	if err != nil {
		return false, err
	}
	root, err := p.rpc.GetStorageRoot(
		ctx,
		p.rpc.L2GethClient,
		l2SignalService,
		blockID,
	)
	if err != nil {
		return false, err
	}

	return parent.Hash() == parentHash &&
		block.Hash() == blockHash &&
		root == signalRoot, nil
}

// requestProofByBlockID performs a proving operation for the given block.
func (p *Prover) requestProofByBlockID(
	blockID *big.Int,
	l1Height *big.Int,
	minTier uint16,
	// If this event is not nil, then the prover will try contesting the transition.
	transitionProvedEvent *bindings.TaikoL1ClientTransitionProved,
) error {
	// NOTE: since this callback function will only be called after a L2 block's proving window is expired,
	// or a wrong proof's submission, so we won't check if L1 chain has been reorged here.
	onBlockProposed := func(
		ctx context.Context,
		event *bindings.TaikoL1ClientBlockProposed,
		end eventIterator.EndBlockProposedEventIterFunc,
	) error {
		// Only filter for exact blockID we want.
		if event.BlockId.Cmp(blockID) != 0 {
			return nil
		}

		// Check whether the block has been verified.
		isVerified, err := p.isBlockVerified(event.BlockId)
		if err != nil {
			return fmt.Errorf("failed to check if the current L2 block is verified: %w", err)
		}
		if isVerified {
			log.Info("📋 Block has been verified", "blockID", event.BlockId)
			return nil
		}

		// Make sure to take a capacity before requesting proof.
		if err := p.takeOneCapacity(event.BlockId); err != nil {
			return err
		}

		if transitionProvedEvent != nil {
			return p.proofContester.SubmitContest(
				ctx,
				event.BlockId,
				new(big.Int).SetUint64(event.Raw.BlockNumber),
				transitionProvedEvent.Tran.ParentHash,
				&event.Meta,
				transitionProvedEvent.Tier,
			)
		}

		// If there is no proof submitter selected, skip proving it.
		if proofSubmitter := p.selectSubmitter(minTier); proofSubmitter != nil {
			return proofSubmitter.RequestProof(ctx, event)
		}

		return nil
	}

	handleBlockProposedEvent := func() error {
		defer func() { <-p.proposeConcurrencyGuard }()

		// Make sure `end` height is less than the latest L1 head.
		l1Head, err := p.rpc.L1.BlockNumber(p.ctx)
		if err != nil {
			log.Error("Failed to get L1 block head", "error", err)
			return err
		}
		end := new(big.Int).Add(l1Height, common.Big1)
		if end.Uint64() > l1Head {
			end = new(big.Int).SetUint64(l1Head)
		}

		iter, err := eventIterator.NewBlockProposedIterator(p.ctx, &eventIterator.BlockProposedIteratorConfig{
			Client:               p.rpc.L1,
			TaikoL1:              p.rpc.TaikoL1,
			StartHeight:          new(big.Int).Sub(l1Height, common.Big1),
			EndHeight:            end,
			OnBlockProposedEvent: onBlockProposed,
		})
		if err != nil {
			log.Error("Failed to start event iterator", "event", "BlockProposed", "error", err)
			return err
		}

		return iter.Iter()
	}

	go func() {
		if err := backoff.Retry(
			func() error {
				if err := handleBlockProposedEvent(); err != nil {
					log.Error(
						"Failed to handle BlockProposed event",
						"error", err,
						"blockID", blockID,
						"maxRetrys", p.cfg.BackOffMaxRetrys,
					)
					return err
				}
				return nil
			},
			backoff.WithMaxRetries(backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval), p.cfg.BackOffMaxRetrys),
		); err != nil {
			log.Error("Failed to request proof with a given block ID", "blockID", blockID, "error", err)
		}
	}()

	return nil
}

// onProvingWindowExpired tries to submit a proof for an expired block.
func (p *Prover) onProvingWindowExpired(ctx context.Context, e *bindings.TaikoL1ClientBlockProposed) error {
	log.Info(
		"Block proving window is expired",
		"blockID", e.BlockId,
		"assignedProver", e.AssignedProver,
		"minTier", e.Meta.MinTier,
	)
	// If proving window is expired, then the assigned prover can not submit new proofs for it anymore.
	if p.proverAddress == e.AssignedProver {
		return nil
	}
	// Check if we still need to generate a new proof for that block.
	proofStatus, err := rpc.GetBlockProofStatus(ctx, p.rpc, e.BlockId, p.proverAddress)
	if err != nil {
		return err
	}
	if proofStatus.IsSubmitted {
		// If there is already a proof submitted and there is no need to contest
		// it, we skip proving this block here.
		if !proofStatus.Invalid || !p.cfg.ContesterMode {
			return nil
		}

		return p.handleInvalidProof(
			ctx,
			e.BlockId,
			new(big.Int).SetUint64(e.Raw.BlockNumber),
			proofStatus.ParentBlock.Hash(),
			proofStatus.CurrentTransitionState.Contester,
			&e.Meta,
			proofStatus.CurrentTransitionState.Tier,
		)
	}

	return p.requestProofByBlockID(e.BlockId, new(big.Int).SetUint64(e.Raw.BlockNumber), e.Meta.MinTier, nil)
}

// getProvingWindow returns the provingWindow of the given proposed block.
func (p *Prover) getProvingWindow(e *bindings.TaikoL1ClientBlockProposed) (time.Duration, error) {
	for _, t := range p.tiers {
		if e.Meta.MinTier == t.ID {
			return time.Duration(t.ProvingWindow) * time.Second, nil
		}
	}

	return 0, errTierNotFound
}

// selectSubmitter returns the proof submitter with the given minTier.
func (p *Prover) selectSubmitter(minTier uint16) proofSubmitter.Submitter {
	for _, s := range p.proofSubmitters {
		if s.Tier() >= minTier {
			return s
		}
	}

	log.Warn("No proof producer / submitter found for the given minTier", "minTier", minTier)

	return nil
}

// getSubmitterByTier returns the proof submitter with the given tier.
func (p *Prover) getSubmitterByTier(tier uint16) proofSubmitter.Submitter {
	for _, s := range p.proofSubmitters {
		if s.Tier() == tier {
			return s
		}
	}

	log.Warn("No proof producer / submitter found for the given tier", "tier", tier)

	return nil
}

// IsGuardianProver reutrns true if the current prover is a guardian prover.
func (p *Prover) IsGuardianProver() bool {
	return p.cfg.GuardianProverAddress != common.Address{}
}

// takeOneCapacity takes one capacity from the capacity manager.
func (p *Prover) takeOneCapacity(blockID *big.Int) error {
	if !p.IsGuardianProver() {
		if _, ok := p.capacityManager.TakeOneCapacity(blockID.Uint64()); !ok {
			return errNoCapacity
		}
	}

	return nil
}

// releaseOneCapacity releases one capacity to the capacity manager.
func (p *Prover) releaseOneCapacity(blockID *big.Int) {
	if !p.IsGuardianProver() {
		if _, released := p.capacityManager.ReleaseOneCapacity(blockID.Uint64()); !released {
			log.Error("Failed to release capacity", "id", blockID)
		}
	}
}
