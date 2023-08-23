package prover

import (
	"context"
	"fmt"
	"math/big"
	"strings"
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
	"github.com/taikoxyz/taiko-client/prover/http"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	proofSubmitter "github.com/taikoxyz/taiko-client/prover/proof_submitter"
	"github.com/urfave/cli/v2"
)

var (
	zeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

type cancelFunc func()

// Prover keep trying to prove new proposed blocks valid/invalid.
type Prover struct {
	// Configurations
	cfg                 *Config
	proverAddress       common.Address
	oracleProverAddress common.Address

	// Clients
	rpc *rpc.Client

	// HTTP Server
	srv *http.Server

	// Contract configurations
	protocolConfigs *bindings.TaikoDataConfig

	// States
	latestVerifiedL1Height uint64
	lastHandledBlockID     uint64
	genesisHeightL1        uint64
	l1Current              *types.Header
	reorgDetectedFlag      bool

	// Proof submitters
	validProofSubmitter proofSubmitter.ProofSubmitter

	// Subscriptions
	blockProposedCh  chan *bindings.TaikoL1ClientBlockProposed
	blockProposedSub event.Subscription
	blockProvenCh    chan *bindings.TaikoL1ClientBlockProven
	blockProvenSub   event.Subscription
	blockVerifiedCh  chan *bindings.TaikoL1ClientBlockVerified
	blockVerifiedSub event.Subscription
	proveNotify      chan struct{}

	// Proof related
	proofGenerationCh chan *proofProducer.ProofWithHeader

	// Concurrency guards
	proposeConcurrencyGuard     chan struct{}
	submitProofConcurrencyGuard chan struct{}
	submitProofTxMutex          *sync.Mutex

	currentBlocksBeingProven                map[uint64]cancelFunc
	currentBlocksBeingProvenMutex           *sync.Mutex
	currentBlocksWaitingForProofWindow      map[uint64]uint64 // l2BlockId : l1Height
	currentBlocksWaitingForProofWindowMutex *sync.Mutex

	// interval settings
	checkProofWindowExpiredInterval time.Duration

	// capacity-related configs
	maxCapacity uint64

	ctx context.Context
	wg  sync.WaitGroup
}

// New initializes the given prover instance based on the command line flags.
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
	p.currentBlocksBeingProven = make(map[uint64]cancelFunc)
	p.currentBlocksBeingProvenMutex = &sync.Mutex{}
	p.currentBlocksWaitingForProofWindow = make(map[uint64]uint64, 0)
	p.currentBlocksWaitingForProofWindowMutex = &sync.Mutex{}
	p.maxCapacity = cfg.Capacity

	if !p.cfg.OracleProver {
		p.srv, err = http.NewServer(http.NewServerOpts{
			ProverPrivateKey: p.cfg.L1ProverPrivKey,
			MaxCapacity:      p.cfg.Capacity,
			MinProofFee:      p.cfg.MinProofFee,
		})
		if err != nil {
			return err
		}
	}

	// Clients
	if p.rpc, err = rpc.NewClient(p.ctx, &rpc.ClientConfig{
		L1Endpoint:     cfg.L1WsEndpoint,
		L2Endpoint:     cfg.L2WsEndpoint,
		TaikoL1Address: cfg.TaikoL1Address,
		TaikoL2Address: cfg.TaikoL2Address,
		RetryInterval:  cfg.BackOffRetryInterval,
		Timeout:        cfg.RPCTimeout,
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

	p.submitProofTxMutex = &sync.Mutex{}
	p.proverAddress = crypto.PubkeyToAddress(p.cfg.L1ProverPrivKey.PublicKey)

	chBufferSize := p.protocolConfigs.BlockMaxProposals
	p.blockProposedCh = make(chan *bindings.TaikoL1ClientBlockProposed, chBufferSize)
	p.blockVerifiedCh = make(chan *bindings.TaikoL1ClientBlockVerified, chBufferSize)
	p.blockProvenCh = make(chan *bindings.TaikoL1ClientBlockProven, chBufferSize)
	p.proofGenerationCh = make(chan *proofProducer.ProofWithHeader, chBufferSize)
	p.proveNotify = make(chan struct{}, 1)
	if err := p.initL1Current(cfg.StartingBlockID); err != nil {
		return fmt.Errorf("initialize L1 current cursor error: %w", err)
	}

	// Concurrency guards
	p.proposeConcurrencyGuard = make(chan struct{}, cfg.MaxConcurrentProvingJobs)
	p.submitProofConcurrencyGuard = make(chan struct{}, cfg.MaxConcurrentProvingJobs)

	p.checkProofWindowExpiredInterval = p.cfg.CheckProofWindowExpiredInterval

	oracleProverAddress, err := p.rpc.TaikoL1.Resolve(
		&bind.CallOpts{Context: ctx},
		p.rpc.L1ChainID,
		rpc.StringToBytes32("oracle_prover"),
		true,
	)
	if err != nil {
		return err
	}

	p.oracleProverAddress = oracleProverAddress

	var producer proofProducer.ProofProducer

	if cfg.Dummy {
		producer = &proofProducer.DummyProofProducer{
			RandomDummyProofDelayLowerBound: p.cfg.RandomDummyProofDelayLowerBound,
			RandomDummyProofDelayUpperBound: p.cfg.RandomDummyProofDelayUpperBound,
		}
	} else {
		if producer, err = proofProducer.NewZkevmRpcdProducer(
			cfg.ZKEvmRpcdEndpoint,
			cfg.ZkEvmRpcdParamsPath,
			cfg.L1HttpEndpoint,
			cfg.L2HttpEndpoint,
			true,
			p.protocolConfigs,
		); err != nil {
			return err
		}
	}

	// Proof submitter
	if p.validProofSubmitter, err = proofSubmitter.NewValidProofSubmitter(
		p.rpc,
		producer,
		p.proofGenerationCh,
		p.cfg.TaikoL2Address,
		p.cfg.L1ProverPrivKey,
		p.submitProofTxMutex,
		p.cfg.OracleProver,
		p.cfg.Graffiti,
		p.cfg.ProofSubmissionMaxRetry,
		p.cfg.BackOffRetryInterval,
		p.cfg.WaitReceiptTimeout,
		p.cfg.ProveBlockGasLimit,
	); err != nil {
		return err
	}

	return nil
}

// Start starts the main loop of the L2 block prover.
func (p *Prover) Start() error {
	p.wg.Add(1)
	p.initSubscription()
	go func() {
		err := p.srv.Start(fmt.Sprintf(":%v", p.cfg.HTTPServerPort))
		if err != nil {
			log.Crit("error starting http server", "error", err)
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

	lastLatestVerifiedL1Height := p.latestVerifiedL1Height

	// If there is too many (TaikoData.Config.maxNumBlocks) pending blocks in TaikoL1 contract, there will be no new
	// BlockProposed temporarily, so except the BlockProposed subscription, we need another trigger to start
	// fetching the proposed blocks.
	forceProvingTicker := time.NewTicker(15 * time.Second)
	defer forceProvingTicker.Stop()

	// If there is no new block verification in `proofCooldownPeriod * 2` seconeds, and the current prover is
	// a special prover, we will go back to try proving the block whose id is `lastVerifiedBlockId + 1`.
	verificationCheckTicker := time.NewTicker(
		time.Duration(p.protocolConfigs.ProofRegularCooldown.Uint64()*2) * time.Second,
	)
	defer verificationCheckTicker.Stop()

	checkProofWindowExpiredTicker := time.After(p.checkProofWindowExpiredInterval)

	// Call reqProving() right away to catch up with the latest state.
	reqProving()

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-verificationCheckTicker.C:
			if err := backoff.Retry(
				func() error { return p.checkChainVerification(lastLatestVerifiedL1Height) },
				backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval),
			); err != nil {
				log.Error("Check chain verification error", "error", err)
			}
		case <-checkProofWindowExpiredTicker:
			func() {
				defer func() { checkProofWindowExpiredTicker = time.After(p.checkProofWindowExpiredInterval) }()
				if err := p.checkProofWindowsExpired(p.ctx); err != nil {
					log.Error("error checking proof window expired", "error", err)
				}
			}()
		case proofWithHeader := <-p.proofGenerationCh:
			p.submitProofOp(p.ctx, proofWithHeader)
		case <-p.proveNotify:
			if err := p.proveOp(); err != nil {
				log.Error("Prove new blocks error", "error", err)
			}
		case <-p.blockProposedCh:
			reqProving()
		case e := <-p.blockVerifiedCh:
			if err := p.onBlockVerified(p.ctx, e); err != nil {
				log.Error("Handle BlockVerified event error", "error", err)
			}
		case e := <-p.blockProvenCh:
			if err := p.onBlockProven(p.ctx, e); err != nil {
				log.Error("Handle BlockProven event error", "error", err)
			}
		case <-forceProvingTicker.C:
			reqProving()
		}
	}
}

// Close closes the prover instance.
func (p *Prover) Close(ctx context.Context) {
	p.closeSubscription()
	if err := p.srv.Shutdown(ctx); err != nil {
		log.Error("Error shutting down http server", "error", err)
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
	// If there is newly generated proofs, we need to submit them as soon as possible.
	if len(p.proofGenerationCh) > 0 {
		end()
		return nil
	}

	if _, err := p.rpc.WaitL1Origin(ctx, event.BlockId); err != nil {
		return fmt.Errorf("failed to wait L1Origin (eventID %d): %w", event.BlockId, err)
	}

	// Check whteher the L2 EE's recorded L1 info, to see if the L1 chain has been reorged.
	reorged, l1CurrentToReset, lastHandledBlockIDToReset, err := p.rpc.CheckL1ReorgFromL2EE(
		ctx,
		new(big.Int).Sub(event.BlockId, common.Big1),
	)
	if err != nil {
		return fmt.Errorf("failed to check whether L1 chain was reorged from L2EE (eventID %d): %w", event.BlockId, err)
	}

	// then check the l1Current cursor at first, to see if the L1 chain has been reorged.
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

	currentL1OriginHeader, err := p.rpc.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(event.Meta.L1Height))
	if err != nil {
		return fmt.Errorf("failed to get L1 header, height %d: %w", event.Meta.L1Height, err)
	}

	if currentL1OriginHeader.Hash() != event.Meta.L1Hash {
		log.Warn(
			"L1 block hash mismatch due to L1 reorg",
			"height", event.Meta.L1Height,
			"currentL1OriginHeader", currentL1OriginHeader.Hash(),
			"L1HashInEvent", event.Meta.L1Hash,
		)

		return fmt.Errorf(
			"L1 block hash mismatch due to L1 reorg: %s != %s",
			currentL1OriginHeader.Hash(),
			event.Meta.L1Hash,
		)
	}

	log.Info(
		"Proposed block",
		"L1Height", event.Raw.BlockNumber,
		"L1Hash", event.Raw.BlockHash,
		"BlockID", event.BlockId,
		"Removed", event.Raw.Removed,
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
		if !p.cfg.OracleProver {
			needNewProof, err := rpc.NeedNewProof(
				p.ctx,
				p.rpc,
				event.BlockId,
				p.proverAddress,
			)
			if err != nil {
				return fmt.Errorf("failed to check whether the L2 block needs a new proof: %w", err)
			}

			if !needNewProof {
				return nil
			}
		}

		// Check if the current prover has seen this block ID before, there was probably
		// a L1 reorg, we need to cancel that reorged block's proof generation task at first.
		if p.currentBlocksBeingProven[event.Meta.Id] != nil {
			p.cancelProof(ctx, event.Meta.Id)
		}

		block, err := p.rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, event.BlockId.Uint64())
		if err != nil {
			return err
		}

		log.Info(
			"Proposed block information",
			"blockID", event.BlockId,
			"prover", block.Prover,
			"proposedAt", block.ProposedAt,
		)

		var skipProofWindowExpiredCheck bool
		if p.cfg.OracleProver {
			shouldSkipProofWindowExpiredCheck := func() (bool, error) {
				parent, err := p.rpc.L2ParentByBlockId(ctx, event.BlockId)
				if err != nil {
					return false, err
				}

				// check if an invalid proof has been submitted, if so, we can skip proofWindowExpired check below
				// and always submit proof. otherwise, oracleProver follows same proof logic as regular.
				forkChoice, err := p.rpc.TaikoL1.GetForkChoice(
					&bind.CallOpts{Context: ctx},
					event.BlockId.Uint64(),
					parent.Hash(),
					uint32(parent.GasUsed),
				)
				if err != nil {
					if strings.Contains(encoding.TryParsingCustomError(err).Error(), "L1_FORK_CHOICE_NOT_FOUND") {
						// proof hasnt been submitted
						return false, nil
					} else {
						return false, err
					}
				}

				block, err := p.rpc.L2.BlockByNumber(ctx, event.BlockId)
				if err != nil {
					return false, err
				}

				// proof is invalid but has correct parents, oracle prover should skip
				// checking proofWindow expired, and simply force prove.
				if forkChoice.BlockHash != block.Hash() {
					log.Info(
						"Oracle prover forcing prove block due to invalid proof",
						"blockID", event.BlockId,
						"forkChoiceBlockHash", common.BytesToHash(forkChoice.BlockHash[:]).Hex(),
						"expectedBlockHash", block.Hash().Hex(),
					)

					return true, nil
				}

				return false, nil
			}

			if skipProofWindowExpiredCheck, err = shouldSkipProofWindowExpiredCheck(); err != nil {
				return err
			}
		}

		if !skipProofWindowExpiredCheck {
			proofWindowExpiresAt := block.ProposedAt + uint64(p.protocolConfigs.ProofWindow)
			proofWindowExpired := uint64(time.Now().Unix()) > proofWindowExpiresAt
			// zero address means anyone can prove, proofWindowExpired means anyone can prove even if not zero address
			if block.Prover != p.proverAddress && block.Prover != zeroAddress && !proofWindowExpired {
				log.Info(
					"Proposed block not proveable",
					"blockID",
					event.BlockId,
					"prover",
					block.Prover.Hex(),
					"proofWindowExpiresAt",
					proofWindowExpiresAt,
					"timeToExpire",
					proofWindowExpiresAt-uint64(time.Now().Unix()),
				)

				// if we cant prove it now, but config is set to wait and try to prove
				// expired proofs
				if p.cfg.ProveUnassignedBlocks {
					log.Info("Adding proposed block to wait for proof window expiration",
						"blockID",
						event.BlockId,
						"prover",
						block.Prover.Hex(),
						"proofWindowExpiresAt",
						proofWindowExpiresAt,
					)

					p.currentBlocksWaitingForProofWindowMutex.Lock()
					p.currentBlocksWaitingForProofWindow[event.Meta.Id] = event.Raw.BlockNumber
					p.currentBlocksWaitingForProofWindowMutex.Unlock()
				}

				return nil
			}

			// if set not to prove unassigned blocks, this block is still not provable
			// by us even though its open proving.
			if (block.Prover == zeroAddress || proofWindowExpired) && !p.cfg.ProveUnassignedBlocks {
				log.Info(
					"Skipping proposed open proving block, not assigned to us",
					"blockID", event.BlockId,
				)
				return nil
			}

			log.Info(
				"Proposed block is proveable",
				"blockID", event.BlockId,
				"prover", block.Prover.Hex(),
				"proofWindowExpired", proofWindowExpired,
			)

			metrics.ProverProofsAssigned.Inc(1)
		}

		ctx, cancelCtx := context.WithCancel(ctx)
		p.currentBlocksBeingProvenMutex.Lock()
		p.currentBlocksBeingProven[event.BlockId.Uint64()] = cancelFunc(func() {
			defer cancelCtx()
			if err := p.validProofSubmitter.CancelProof(ctx, event.BlockId); err != nil {
				log.Error("failed to cancel proof", "error", err, "blockID", event.BlockId)
			}
		})
		p.currentBlocksBeingProvenMutex.Unlock()

		return p.validProofSubmitter.RequestProof(ctx, event)
	}

	p.proposeConcurrencyGuard <- struct{}{}

	newL1Current, err := p.rpc.L1.HeaderByHash(ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}
	p.l1Current = newL1Current
	p.lastHandledBlockID = event.BlockId.Uint64()

	go func() {
		if err := backoff.Retry(
			func() error { return handleBlockProposedEvent() },
			backoff.WithMaxRetries(backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval), p.cfg.BackOffMaxRetrys),
		); err != nil {
			p.currentBlocksBeingProvenMutex.Lock()
			delete(p.currentBlocksBeingProven, event.BlockId.Uint64())
			p.currentBlocksBeingProvenMutex.Unlock()
			log.Error("Handle new BlockProposed event error", "error", err)
		}
	}()

	return nil
}

// submitProofOp performs a proof submission operation.
func (p *Prover) submitProofOp(ctx context.Context, proofWithHeader *proofProducer.ProofWithHeader) {
	p.submitProofConcurrencyGuard <- struct{}{}
	go func() {
		defer func() {
			<-p.submitProofConcurrencyGuard
			p.currentBlocksBeingProvenMutex.Lock()
			delete(p.currentBlocksBeingProven, proofWithHeader.Meta.Id)
			p.currentBlocksBeingProvenMutex.Unlock()
		}()

		if err := backoff.Retry(
			func() error {
				err := p.validProofSubmitter.SubmitProof(p.ctx, proofWithHeader)
				if err != nil {
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

// onBlockVerified update the latestVerified block in current state, and cancels
// the block being proven if it's verified.
func (p *Prover) onBlockVerified(ctx context.Context, event *bindings.TaikoL1ClientBlockVerified) error {
	metrics.ProverLatestVerifiedIDGauge.Update(event.BlockId.Int64())

	p.latestVerifiedL1Height = event.Raw.BlockNumber

	log.Info(
		"New verified block",
		"blockID", event.BlockId,
		"hash", common.BytesToHash(event.BlockHash[:]),
		"prover", event.Prover,
	)

	// cancel any proofs being generated for this block
	p.cancelProof(ctx, event.BlockId.Uint64())

	return nil
}

// onBlockProven cancels proof generation if the proof is being generated by this prover,
// and the proof is not the oracle proof address.
func (p *Prover) onBlockProven(ctx context.Context, event *bindings.TaikoL1ClientBlockProven) error {
	metrics.ProverReceivedProvenBlockGauge.Update(event.BlockId.Int64())
	// if this proof is submitted by an oracle prover or a system prover, don't cancel proof.
	if event.Prover == p.oracleProverAddress ||
		event.Prover == encoding.OracleProverAddress {
		return nil
	}

	// cancel any proofs being generated for this block
	isValidProof, err := p.isValidProof(
		ctx,
		event.BlockId.Uint64(),
		uint64(event.ParentGasUsed),
		event.ParentHash,
		event.BlockHash,
	)

	if err != nil {
		return err
	}

	if isValidProof {
		p.cancelProof(ctx, event.BlockId.Uint64())
	} else {
		// generate oracle proof if oracle prover, proof is invalid
		if p.cfg.OracleProver {
			return p.requestProofForBlockId(event.BlockId, new(big.Int).SetUint64(event.Raw.BlockNumber))
		}
	}

	return nil
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
	p.genesisHeightL1 = stateVars.GenesisHeight

	if startingBlockID == nil {
		if stateVars.LastVerifiedBlockId == 0 {
			genesisL1Header, err := p.rpc.L1.HeaderByNumber(p.ctx, new(big.Int).SetUint64(stateVars.GenesisHeight))
			if err != nil {
				return err
			}

			p.l1Current = genesisL1Header
			return nil
		}

		startingBlockID = new(big.Int).SetUint64(stateVars.LastVerifiedBlockId)
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

// isBlockVerified checks whether the given block has been verified by other provers.
func (p *Prover) isBlockVerified(id *big.Int) (bool, error) {
	stateVars, err := p.rpc.GetProtocolStateVariables(&bind.CallOpts{Context: p.ctx})
	if err != nil {
		return false, err
	}

	return id.Uint64() <= stateVars.LastVerifiedBlockId, nil
}

// initSubscription initializes all subscriptions in current prover instance.
func (p *Prover) initSubscription() {
	p.blockProposedSub = rpc.SubscribeBlockProposed(p.rpc.TaikoL1, p.blockProposedCh)
	p.blockVerifiedSub = rpc.SubscribeBlockVerified(p.rpc.TaikoL1, p.blockVerifiedCh)
	p.blockProvenSub = rpc.SubscribeBlockProven(p.rpc.TaikoL1, p.blockProvenCh)
}

// closeSubscription closes all subscriptions.
func (p *Prover) closeSubscription() {
	p.blockVerifiedSub.Unsubscribe()
	p.blockProposedSub.Unsubscribe()
}

// checkChainVerification checks if there is no new block verification in protocol, if so,
// it will let current sepecial prover to go back to try proving the block whose id is `lastVerifiedBlockId + 1`.
func (p *Prover) checkChainVerification(lastLatestVerifiedL1Height uint64) error {
	if (!p.cfg.OracleProver) || lastLatestVerifiedL1Height != p.latestVerifiedL1Height {
		return nil
	}

	log.Warn(
		"No new block verification in `proofCooldownPeriod * 2` seconeds",
		"latestVerifiedL1Height", p.latestVerifiedL1Height,
		"proofCooldownPeriod", p.protocolConfigs.ProofRegularCooldown,
	)

	stateVar, err := p.rpc.TaikoL1.GetStateVariables(&bind.CallOpts{Context: p.ctx})
	if err != nil {
		log.Error("Failed to get protocol state variables", "error", err)
		return err
	}

	if err := p.initL1Current(new(big.Int).SetUint64(stateVar.LastVerifiedBlockId)); err != nil {
		return err
	}
	p.lastHandledBlockID = stateVar.LastVerifiedBlockId

	return nil
}

// isValidProof cancels proof only if the parentGasUsed and parentHash in the proof match what
// is expected
func (p *Prover) isValidProof(
	ctx context.Context,
	blockID uint64,
	parentGasUsed uint64,
	parentHash common.Hash,
	blockHash common.Hash,
) (bool, error) {
	parent, err := p.rpc.L2ParentByBlockId(ctx, new(big.Int).SetUint64(blockID))
	if err != nil {
		return false, err
	}

	block, err := p.rpc.L2.BlockByNumber(ctx, new(big.Int).SetUint64(blockID))
	if err != nil {
		return false, err
	}

	if parent.GasUsed == parentGasUsed && parent.Hash() == parentHash && blockHash == block.Hash() {
		return true, nil
	}

	return false, nil
}

// cancelProof cancels local proof generation
func (p *Prover) cancelProof(ctx context.Context, blockID uint64) {
	p.currentBlocksBeingProvenMutex.Lock()
	defer p.currentBlocksBeingProvenMutex.Unlock()

	if cancel, ok := p.currentBlocksBeingProven[blockID]; ok {
		cancel()
		delete(p.currentBlocksBeingProven, blockID)
	}
}

// checkProofWindowsExpired iterates through the current blocks waiting for proof window to expire,
// which are blocks that have been proposed, but we were not selected as the prover. if the proof window
// has expired, we can start generating a proof for them.
func (p *Prover) checkProofWindowsExpired(ctx context.Context) error {
	p.currentBlocksWaitingForProofWindowMutex.Lock()
	defer p.currentBlocksWaitingForProofWindowMutex.Unlock()

	for blockId, l1Height := range p.currentBlocksWaitingForProofWindow {
		if err := p.checkProofWindowExpired(ctx, l1Height, blockId); err != nil {
			return err
		}
	}

	return nil
}

// checkProofWindowExpired checks a single instance of a block to see if its proof winodw has expired
// and the proof is now able to be submitted by anyone, not just the blocks assigned prover.
func (p *Prover) checkProofWindowExpired(ctx context.Context, l1Height, blockId uint64) error {
	block, err := p.rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, blockId)
	if err != nil {
		return encoding.TryParsingCustomError(err)
	}

	isExpired := time.Now().Unix() > int64(block.ProposedAt)+int64(p.protocolConfigs.ProofWindow)

	if isExpired {
		log.Debug(
			"Block proof window is expired", "blockID", blockId, "l1Height", l1Height)

		// we should remove this block from being watched regardless of whether the block
		// has a valid proof
		delete(p.currentBlocksWaitingForProofWindow, blockId)

		// we can see if a fork choice with correct parentHash/gasUsed has come in.
		// if it hasnt, we can start to generate a proof for this.
		parent, err := p.rpc.L2ParentByBlockId(ctx, new(big.Int).SetUint64(blockId))
		if err != nil {
			return err
		}

		forkChoice, err := p.rpc.TaikoL1.GetForkChoice(
			&bind.CallOpts{Context: ctx},
			blockId,
			parent.Hash(),
			uint32(parent.GasUsed),
		)

		if err != nil && !strings.Contains(encoding.TryParsingCustomError(err).Error(), "L1_FORK_CHOICE_NOT_FOUND") {
			return encoding.TryParsingCustomError(err)
		}

		if forkChoice.Prover == zeroAddress {
			log.Info(
				"Proof window for proof not assigned to us expired, requesting proof",
				"blockID",
				blockId,
				"l1Height",
				l1Height,
			)
			// we can generate the proof, no proof came in by proof window expiring
			if err := p.requestProofForBlockId(new(big.Int).SetUint64(blockId), new(big.Int).SetUint64(l1Height)); err != nil {
				return err
			}
		} else {
			// we need to check the block hash vs the proof's blockHash to see
			// if the proof is valid or not
			block, err := p.rpc.L2.BlockByNumber(ctx, new(big.Int).SetUint64(blockId))
			if err != nil {
				return err
			}

			// if the hashes dont match, we can generate proof even though
			// a proof came in before proofwindow expired.
			if block.Hash() != forkChoice.BlockHash {
				log.Info(
					"Invalid proof detected while watching for proof window expiration, requesting proof",
					"blockID",
					blockId,
					"l1Height",
					l1Height,
					"expectedBlockHash",
					block.Hash(),
					"forkChoiceBlockHash",
					common.Bytes2Hex(forkChoice.BlockHash[:]),
				)
				// we can generate the proof, the proof is incorrect since blockHash does not match
				// the correct one but parentHash/gasUsed are correct.
				if err := p.requestProofForBlockId(new(big.Int).SetUint64(blockId), new(big.Int).SetUint64(l1Height)); err != nil {
					return err
				}
			}
		}
	}

	// otherwise, keep it in the map and check again next iteration
	return nil
}

// proveOp performs a proving operation, find current unproven blocks, then
// request generating proofs for them.
func (p *Prover) requestProofForBlockId(blockId *big.Int, l1Height *big.Int) error {
	onBlockProposed := func(
		ctx context.Context,
		event *bindings.TaikoL1ClientBlockProposed,
		end eventIterator.EndBlockProposedEventIterFunc,
	) error {
		// only filter for exact blockID we want
		if event.BlockId.Cmp(blockId) != 0 {
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

		ctx, cancelCtx := context.WithCancel(ctx)
		p.currentBlocksBeingProvenMutex.Lock()
		p.currentBlocksBeingProven[event.BlockId.Uint64()] = cancelFunc(func() {
			defer cancelCtx()
			if err := p.validProofSubmitter.CancelProof(ctx, event.BlockId); err != nil {
				log.Error("Failed to cancel proof", "error", err, "blockID", event.BlockId)
			}
		})
		p.currentBlocksBeingProvenMutex.Unlock()

		p.proposeConcurrencyGuard <- struct{}{}

		if err := p.validProofSubmitter.RequestProof(ctx, event); err != nil {
			return err
		}

		return nil
	}

	handleBlockProposedEvent := func() error {
		defer func() { <-p.proposeConcurrencyGuard }()

		iter, err := eventIterator.NewBlockProposedIterator(p.ctx, &eventIterator.BlockProposedIteratorConfig{
			Client:               p.rpc.L1,
			TaikoL1:              p.rpc.TaikoL1,
			StartHeight:          l1Height,
			EndHeight:            new(big.Int).Add(l1Height, common.Big1),
			OnBlockProposedEvent: onBlockProposed,
			FilterQuery:          []*big.Int{blockId},
		})
		if err != nil {
			return err
		}

		return iter.Iter()
	}

	go func() {
		if err := backoff.Retry(
			func() error {
				return handleBlockProposedEvent()
			},
			backoff.WithMaxRetries(backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval), p.cfg.BackOffMaxRetrys),
		); err != nil {
			p.currentBlocksBeingProvenMutex.Lock()
			defer p.currentBlocksBeingProvenMutex.Unlock()
			delete(p.currentBlocksBeingProven, blockId.Uint64())
			log.Error("Request proof with a given block ID", "blockID", blockId, "error", err)
		}
	}()

	return nil
}
