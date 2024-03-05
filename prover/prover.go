package prover

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/internal/sender"
	"github.com/taikoxyz/taiko-client/internal/version"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	handler "github.com/taikoxyz/taiko-client/prover/event_handler"
	guardianproversender "github.com/taikoxyz/taiko-client/prover/guardian_prover_sender"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	proofSubmitter "github.com/taikoxyz/taiko-client/prover/proof_submitter"
	"github.com/taikoxyz/taiko-client/prover/proof_submitter/transaction"
	"github.com/taikoxyz/taiko-client/prover/server"
	state "github.com/taikoxyz/taiko-client/prover/shared_state"
)

var (
	heartbeatInterval = 12 * time.Second
)

// Prover keeps trying to prove newly proposed blocks.
type Prover struct {
	// Configurations
	cfg              *Config
	proverPrivateKey *ecdsa.PrivateKey

	// Clients
	rpc *rpc.Client

	// Guardian prover related
	srv                  *server.ProverServer
	guardianProverSender guardianproversender.BlockSenderHeartbeater

	// Contract configurations
	protocolConfigs *bindings.TaikoDataConfig

	// States
	sharedState     *state.SharedState
	genesisHeightL1 uint64

	// Event handlers
	blockProposedHandler       handler.BlockProposedHandler
	blockVerifiedHandler       handler.BlockVerifiedHandler
	transitionContestedHandler handler.TransitionContestedHandler
	transitionProvedHandler    handler.TransitionProvedHandler
	assignmentExpiredHandler   handler.AssignmentExpiredHandler

	// Proof submitters
	proofSubmitters []proofSubmitter.Submitter
	proofContester  proofSubmitter.Contester

	assignmentExpiredCh chan *bindings.TaikoL1ClientBlockProposed
	proveNotify         chan struct{}

	// Proof related channels
	proofSubmissionCh chan *proofSubmitter.ProofRequestBody
	proofContestCh    chan *proofSubmitter.ContestRequestBody
	proofGenerationCh chan *proofProducer.ProofWithHeader

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
	p.proverPrivateKey = cfg.L1ProverPrivKey
	p.sharedState = new(state.SharedState)

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
		BackOffMaxRetries:     cfg.BackOffMaxRetrys,
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

	proverAddress := crypto.PubkeyToAddress(p.cfg.L1ProverPrivKey.PublicKey)

	chBufferSize := p.protocolConfigs.BlockMaxProposals
	p.proofGenerationCh = make(chan *proofProducer.ProofWithHeader, chBufferSize)
	p.assignmentExpiredCh = make(chan *bindings.TaikoL1ClientBlockProposed, chBufferSize)
	p.proofSubmissionCh = make(chan *proofSubmitter.ProofRequestBody, p.cfg.Capacity)
	p.proofContestCh = make(chan *proofSubmitter.ContestRequestBody, p.cfg.Capacity)
	p.proveNotify = make(chan struct{}, 1)

	if err := p.initL1Current(cfg.StartingBlockID); err != nil {
		return fmt.Errorf("initialize L1 current cursor error: %w", err)
	}

	// Protocol proof tiers
	tiers, err := p.rpc.GetTiers(ctx)
	if err != nil {
		return err
	}
	p.sharedState.SetTiers(tiers)

	senderCfg := &sender.Config{
		ConfirmationDepth: 0,
		MaxRetrys:         p.cfg.ProofSubmissionMaxRetry,
		GasGrowthRate:     p.cfg.ProveBlockTxReplacementGasGrowthRate,
	}
	if p.cfg.ProveBlockGasLimit != nil {
		senderCfg.GasLimit = *p.cfg.ProveBlockGasLimit
	}
	if p.cfg.ProveBlockMaxTxGasFeeCap != nil {
		senderCfg.MaxGasFee = p.cfg.ProveBlockMaxTxGasFeeCap.Uint64()
	}
	// For guaridan provers we always simply keep retrying for each its proof submission.
	if p.IsGuardianProver() && senderCfg.MaxRetrys != 0 {
		senderCfg.MaxRetrys = 0
	}

	txSender, err := sender.NewSender(p.ctx, senderCfg, p.rpc.L1, p.proverPrivateKey)
	if err != nil {
		return err
	}

	txBuilder := transaction.NewProveBlockTxBuilder(
		p.rpc,
		p.proverPrivateKey,
	)

	// Proof submitters
	if err := p.initProofSubmitters(p.ctx, txSender, txBuilder); err != nil {
		return err
	}

	// Proof contester
	p.proofContester, err = proofSubmitter.NewProofContester(
		p.ctx,
		p.rpc,
		p.cfg.L1ProverPrivKey,
		txSender,
		p.cfg.Graffiti,
		txBuilder,
	)
	if err != nil {
		return err
	}

	// levelDB
	var db ethdb.KeyValueStore
	if cfg.DatabasePath != "" {
		if db, err = leveldb.New(
			cfg.DatabasePath,
			int(cfg.DatabaseCacheSize),
			16, // Minimum number of files handles is 16 in leveldb.
			"taiko",
			false,
		); err != nil {
			return err
		}
	}

	// Prover server
	proverServerOpts := &server.NewProverServerOpts{
		ProverPrivateKey:      p.cfg.L1ProverPrivKey,
		MinOptimisticTierFee:  p.cfg.MinOptimisticTierFee,
		MinSgxTierFee:         p.cfg.MinSgxTierFee,
		MaxExpiry:             p.cfg.MaxExpiry,
		MaxBlockSlippage:      p.cfg.MaxBlockSlippage,
		TaikoL1Address:        p.cfg.TaikoL1Address,
		AssignmentHookAddress: p.cfg.AssignmentHookAddress,
		RPC:                   p.rpc,
		ProtocolConfigs:       &protocolConfigs,
		LivenessBond:          protocolConfigs.LivenessBond,
		IsGuardian:            p.IsGuardianProver(),
		DB:                    db,
	}
	if p.srv, err = server.New(proverServerOpts); err != nil {
		return err
	}

	// Guardian prover heartbeat sender
	if p.IsGuardianProver() {
		// Check guardian prover contract address is correct.
		if _, err := p.rpc.GuardianProver.MinGuardians(&bind.CallOpts{Context: ctx}); err != nil {
			return fmt.Errorf("failed to get MinGuardians from guardian prover contract: %w", err)
		}

		p.guardianProverSender = guardianproversender.New(
			p.cfg.L1ProverPrivKey,
			p.cfg.GuardianProverHealthCheckServerEndpoint,
			db,
			p.rpc,
			proverAddress,
		)
	}

	// Initialize event handlers.
	p.initEventHandlers()

	return nil
}

// Start starts the main loop of the L2 block prover.
func (p *Prover) Start() error {
	// 1. Set approval amount for the contracts.
	for _, contract := range []common.Address{p.cfg.TaikoL1Address, p.cfg.AssignmentHookAddress} {
		if err := p.setApprovalAmount(p.ctx, contract); err != nil {
			log.Crit("Failed to set approval amount", "contract", contract, "error", err)
		}
	}

	// 2. Start the prover server.
	go func() {
		if err := p.srv.Start(fmt.Sprintf(":%v", p.cfg.HTTPServerPort)); !errors.Is(err, http.ErrServerClosed) {
			log.Crit("Failed to start http server", "error", err)
		}
	}()

	// 3. Start the guardian prover heartbeat sender if the current prover is a guardian prover.
	if p.IsGuardianProver() {
		if err := p.guardianProverSender.SendStartup(
			p.ctx,
			version.CommitVersion(),
			version.CommitVersion(),
			p.cfg.L1NodeVersion,
			p.cfg.L2NodeVersion,
		); err != nil {
			log.Crit("Failed to send guardian prover startup", "error", err)
		}

		go p.gurdianProverHeartbeatLoop(p.ctx)
	}

	// 4. Start the main event loop of the prover.
	go p.eventLoop()

	return nil
}

// eventLoop starts the main loop of Taiko prover.
func (p *Prover) eventLoop() {
	p.wg.Add(1)
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
	// Call reqProving() right away to catch up with the latest state.
	reqProving()

	// If there is too many (TaikoData.Config.blockMaxProposals) pending blocks in TaikoL1 contract, there will be no new
	// BlockProposed event temporarily, so except the BlockProposed subscription, we need another trigger to start
	// fetching the proposed blocks.
	forceProvingTicker := time.NewTicker(15 * time.Second)
	defer forceProvingTicker.Stop()

	// Channels
	chBufferSize := p.protocolConfigs.BlockMaxProposals
	blockProposedCh := make(chan *bindings.TaikoL1ClientBlockProposed, chBufferSize)
	blockVerifiedCh := make(chan *bindings.TaikoL1ClientBlockVerified, chBufferSize)
	transitionProvedCh := make(chan *bindings.TaikoL1ClientTransitionProved, chBufferSize)
	transitionContestedCh := make(chan *bindings.TaikoL1ClientTransitionContested, chBufferSize)
	// Subscriptions
	blockProposedSub := rpc.SubscribeBlockProposed(p.rpc.TaikoL1, blockProposedCh)
	blockVerifiedSub := rpc.SubscribeBlockVerified(p.rpc.TaikoL1, blockVerifiedCh)
	transitionProvedSub := rpc.SubscribeTransitionProved(p.rpc.TaikoL1, transitionProvedCh)
	transitionContestedSub := rpc.SubscribeTransitionContested(p.rpc.TaikoL1, transitionContestedCh)
	defer func() {
		blockProposedSub.Unsubscribe()
		blockVerifiedSub.Unsubscribe()
		transitionProvedSub.Unsubscribe()
		transitionContestedSub.Unsubscribe()
	}()

	for {
		select {
		case <-p.ctx.Done():
			return
		case proofWithHeader := <-p.proofGenerationCh:
			p.submitProofOp(p.ctx, proofWithHeader)
		case req := <-p.proofSubmissionCh:
			p.requestProofOp(p.ctx, req.Event, req.Tier)
		case req := <-p.proofContestCh:
			p.contestProofOp(p.ctx, req)
		case <-p.proveNotify:
			if err := p.proveOp(); err != nil {
				log.Error("Prove new blocks error", "error", err)
			}
		case e := <-blockVerifiedCh:
			p.blockVerifiedHandler.Handle(e)
		case e := <-transitionProvedCh:
			go func() {
				if err := p.withRetry(func() error { return p.transitionProvedHandler.Handle(p.ctx, e) }); err != nil {
					log.Error("Handle TaikoL1.TransitionProved event error", "error", err)
				}
			}()
		case e := <-transitionContestedCh:
			go func() {
				if err := p.withRetry(func() error { return p.transitionContestedHandler.Handle(p.ctx, e) }); err != nil {
					log.Error("Handle TaikoL1.TransitionContested event error", "error", err)
				}
			}()
		case e := <-p.assignmentExpiredCh:
			go func() {
				if err := p.withRetry(func() error { return p.assignmentExpiredHandler.Handle(p.ctx, e) }); err != nil {
					log.Error("Handle proof window expired event error", "error", err)
				}
			}()
		case <-blockProposedCh:
			reqProving()
		case <-forceProvingTicker.C:
			reqProving()
		}
	}
}

// Close closes the prover instance.
func (p *Prover) Close(ctx context.Context) {
	if p.guardianProverSender != nil {
		if err := p.guardianProverSender.Close(); err != nil {
			log.Error("failed to close database connection", "error", err)
		}
	}

	if err := p.srv.Shutdown(ctx); err != nil {
		log.Error("Failed to shut down prover server", "error", err)
	}
	p.wg.Wait()
}

// proveOp iterates through BlockProposed events
func (p *Prover) proveOp() error {
	firstTry := true

	for firstTry || p.sharedState.GetReorgDetectedFlag() {
		p.sharedState.SetReorgDetectedFlag(false)
		firstTry = false

		iter, err := eventIterator.NewBlockProposedIterator(p.ctx, &eventIterator.BlockProposedIteratorConfig{
			Client:               p.rpc.L1,
			TaikoL1:              p.rpc.TaikoL1,
			StartHeight:          new(big.Int).SetUint64(p.sharedState.GetL1Current().Number.Uint64()),
			OnBlockProposedEvent: p.blockProposedHandler.Handle,
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

// contestProofOp performs a proof contest operation.
func (p *Prover) contestProofOp(ctx context.Context, req *proofSubmitter.ContestRequestBody) {
	go func() {
		if err := backoff.Retry(func() error {
			if err := p.proofContester.SubmitContest(
				p.ctx,
				req.BlockID,
				req.ProposedIn,
				req.ParentHash,
				req.Meta,
				req.Tier,
			); err != nil {
				log.Error("Request new proof contest error", "blockID", req.BlockID, "error", err)
				return err
			}

			return nil
		}, backoff.WithMaxRetries(
			backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval),
			p.cfg.BackOffMaxRetrys,
		)); err != nil {
			log.Error("Request new proof contest error", "blockID", req.BlockID, "error", err)
		}
	}()
}

// requestProofOp requests a new proof generation operation.
func (p *Prover) requestProofOp(ctx context.Context, e *bindings.TaikoL1ClientBlockProposed, minTier uint16) {
	go func() {
		if p.IsGuardianProver() {
			minTier = encoding.TierGuardianID
		}

		if err := backoff.Retry(func() error {
			if ctx.Err() != nil {
				log.Error("Context is done, aborting requestProofOp", "blockID", e.BlockId, "error", ctx.Err())
				return nil
			}

			if proofSubmitter := p.selectSubmitter(minTier); proofSubmitter != nil {
				if err := proofSubmitter.RequestProof(ctx, e); err != nil {
					log.Error("Request new proof error", "blockID", e.BlockId, "error", err)
					return err
				}

				return nil
			}

			log.Error("Failed to find proof submitter", "blockID", e.BlockId, "minTier", minTier)
			return nil
		}, backoff.WithMaxRetries(
			backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval),
			p.cfg.BackOffMaxRetrys,
		)); err != nil {
			log.Error("Request new proof error", "blockID", e.BlockId, "error", err)
		}
	}()
}

// submitProofOp performs a proof submission operation.
func (p *Prover) submitProofOp(ctx context.Context, proofWithHeader *proofProducer.ProofWithHeader) {
	go func() {
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

// Name returns the application name.
func (p *Prover) Name() string {
	return "prover"
}

// selectSubmitter returns the proof submitter with the given minTier.
func (p *Prover) selectSubmitter(minTier uint16) proofSubmitter.Submitter {
	for _, s := range p.proofSubmitters {
		if s.Tier() >= minTier {
			log.Debug("Proof submitter selected", "tier", s.Tier(), "minTier", minTier)
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

// IsGuardianProver returns true if the current prover is a guardian prover.
func (p *Prover) IsGuardianProver() bool {
	return p.cfg.GuardianProverAddress != common.Address{}
}

// ProverAddress returns the current prover account address.
func (p *Prover) ProverAddress() common.Address {
	return crypto.PubkeyToAddress(p.proverPrivateKey.PublicKey)
}

// withRetry retries the given function with prover backoff policy.
func (p *Prover) withRetry(f func() error) error {
	return backoff.Retry(
		func() error {
			if p.ctx.Err() != nil {
				log.Error("Context is done, aborting", "error", p.ctx.Err())
				return nil
			}
			return f()
		},
		backoff.WithMaxRetries(backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval), p.cfg.BackOffMaxRetrys),
	)
}
