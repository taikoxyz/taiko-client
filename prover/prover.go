package prover

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
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
	txListValidator "github.com/taikoxyz/taiko-client/pkg/tx_list_validator"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	proofSubmitter "github.com/taikoxyz/taiko-client/prover/proof_submitter"
	"github.com/urfave/cli/v2"
)

type cancelFunc func()

// Prover keep trying to prove new proposed blocks valid/invalid.
type Prover struct {
	// Configurations
	cfg                 *Config
	proverAddress       common.Address
	oracleProverAddress common.Address
	systemProverAddress common.Address

	// Clients
	rpc *rpc.Client

	// Contract configurations
	txListValidator *txListValidator.TxListValidator
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

	currentBlocksBeingProven      map[uint64]cancelFunc
	currentBlocksBeingProvenMutex *sync.Mutex

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
	protocolConfigs, err := p.rpc.TaikoL1.GetConfig(nil)
	if err != nil {
		return fmt.Errorf("failed to get protocol configs: %w", err)
	}
	p.protocolConfigs = &protocolConfigs

	log.Info("Protocol configs", "configs", p.protocolConfigs)

	p.submitProofTxMutex = &sync.Mutex{}
	p.txListValidator = txListValidator.NewTxListValidator(
		p.protocolConfigs.BlockMaxGasLimit,
		p.protocolConfigs.MaxTransactionsPerBlock,
		p.protocolConfigs.MaxBytesPerTxList,
		p.rpc.L2ChainID,
	)
	p.proverAddress = crypto.PubkeyToAddress(p.cfg.L1ProverPrivKey.PublicKey)

	chBufferSize := p.protocolConfigs.MaxNumProposedBlocks.Uint64()
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

	oracleProverAddress, err := p.rpc.TaikoL1.Resolve(nil, p.rpc.L1ChainID, rpc.StringToBytes32("oracle_prover"), true)
	if err != nil {
		return err
	}

	p.oracleProverAddress = oracleProverAddress

	systemProverAddress, err := p.rpc.TaikoL1.Resolve(nil, p.rpc.L1ChainID, rpc.StringToBytes32("system_prover"), true)
	if err != nil {
		return err
	}

	p.systemProverAddress = systemProverAddress

	var producer proofProducer.ProofProducer

	isSystemProver := cfg.SystemProver
	isOracleProver := cfg.OracleProver

	if isSystemProver || isOracleProver {
		var specialProverAddress common.Address
		var privateKey *ecdsa.PrivateKey
		if isSystemProver {
			specialProverAddress = systemProverAddress
			privateKey = p.cfg.SystemProverPrivateKey
		} else {
			specialProverAddress = oracleProverAddress
			privateKey = p.cfg.OracleProverPrivateKey
		}

		if producer, err = proofProducer.NewSpecialProofProducer(
			p.rpc,
			privateKey,
			p.cfg.TaikoL2Address,
			specialProverAddress,
			p.cfg.Graffiti,
			isSystemProver,
		); err != nil {
			return err
		}
	} else if cfg.Dummy {
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
		p.cfg.SystemProver,
		p.cfg.Graffiti,
		p.cfg.ExpectedReward,
		p.cfg.BackOffRetryInterval,
	); err != nil {
		return err
	}

	return nil
}

// Start starts the main loop of the L2 block prover.
func (p *Prover) Start() error {
	p.wg.Add(1)
	p.initSubscription()
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
		time.Duration(p.protocolConfigs.ProofCooldownPeriod.Uint64()*2) * time.Second,
	)
	defer verificationCheckTicker.Stop()

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
func (p *Prover) Close() {
	p.closeSubscription()
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

	if _, err := p.rpc.WaitL1Origin(ctx, event.Id); err != nil {
		return fmt.Errorf("failed to wait L1Origin (eventID %d): %w", event.Id, err)
	}

	// Check whteher the L2 EE's recorded L1 info, to see if the L1 chain has been reorged.
	reorged, l1CurrentToReset, lastHandledBlockIDToReset, err := p.rpc.CheckL1ReorgFromL2EE(
		ctx,
		new(big.Int).Sub(event.Id, common.Big1),
	)
	if err != nil {
		return fmt.Errorf("failed to check whether L1 chain was reorged from L2EE (eventID %d): %w", event.Meta.Id, err)
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
				event.Meta.Id,
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

	if event.Id.Uint64() <= p.lastHandledBlockID {
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
		"BlockID", event.Id,
		"BlockFee", event.BlockFee,
		"Removed", event.Raw.Removed,
	)
	metrics.ProverReceivedProposedBlockGauge.Update(event.Id.Int64())

	handleBlockProposedEvent := func() error {
		defer func() { <-p.proposeConcurrencyGuard }()

		// Check whether the block has been verified.
		isVerified, err := p.isBlockVerified(event.Id)
		if err != nil {
			return fmt.Errorf("failed to check if the current L2 block is verified: %w", err)
		}

		if isVerified {
			log.Info("📋 Block has been verified", "blockID", event.Id)
			return nil
		}

		// Check whether the block's proof is still needed.
		if !p.cfg.OracleProver && !p.cfg.SystemProver {
			needNewProof, err := rpc.NeedNewProof(
				p.ctx,
				p.rpc,
				event.Id,
				p.proverAddress,
				p.protocolConfigs.RealProofSkipSize,
			)
			if err != nil {
				return fmt.Errorf("failed to check whether the L2 block needs a new proof: %w", err)
			}

			if !needNewProof {
				return nil
			}
		}

		if p.cfg.SystemProver {
			needNewSystemProof, err := rpc.NeedNewSystemProof(ctx, p.rpc, event.Id, p.protocolConfigs.RealProofSkipSize)
			if err != nil {
				return fmt.Errorf("failed to check whether the L2 block needs a new system proof: %w", err)
			}

			if !needNewSystemProof {
				return nil
			}
		}

		// Check if the current prover has seen this block ID before, there was probably
		// a L1 reorg, we need to cancel that reorged block's proof generation task at first.
		if p.currentBlocksBeingProven[event.Meta.Id] != nil {
			p.cancelProof(ctx, event.Meta.Id)
		}

		ctx, cancelCtx := context.WithCancel(ctx)
		p.currentBlocksBeingProvenMutex.Lock()
		p.currentBlocksBeingProven[event.Id.Uint64()] = cancelFunc(func() {
			defer cancelCtx()
			if err := p.validProofSubmitter.CancelProof(ctx, event.Id); err != nil {
				log.Error("failed to cancel proof", "error", err, "blockID", event.Id)
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
	p.lastHandledBlockID = event.Meta.Id

	go func() {
		if err := backoff.Retry(
			func() error { return handleBlockProposedEvent() },
			backoff.WithMaxRetries(backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval), p.cfg.BackOffMaxRetrys),
		); err != nil {
			p.currentBlocksBeingProvenMutex.Lock()
			delete(p.currentBlocksBeingProven, event.Id.Uint64())
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
			func() error { return p.validProofSubmitter.SubmitProof(p.ctx, proofWithHeader) },
			backoff.WithMaxRetries(backoff.NewConstantBackOff(p.cfg.BackOffRetryInterval), p.cfg.BackOffMaxRetrys),
		); err != nil {
			log.Error("Submit proof error", "error", err)
		}
	}()
}

// onBlockVerified update the latestVerified block in current state, and cancels
// the block being proven if it's verified.
func (p *Prover) onBlockVerified(ctx context.Context, event *bindings.TaikoL1ClientBlockVerified) error {
	metrics.ProverLatestVerifiedIDGauge.Update(event.Id.Int64())

	isNormalProof := p.protocolConfigs.RealProofSkipSize == nil ||
		(p.protocolConfigs.RealProofSkipSize != nil && event.Id.Uint64()%p.protocolConfigs.RealProofSkipSize.Uint64() == 0)
	if event.Reward > math.MaxInt64 {
		metrics.ProverAllProofRewardGauge.Update(math.MaxInt64)
		if isNormalProof {
			metrics.ProverNormalProofRewardGauge.Update(math.MaxInt64)
		}
	} else {
		metrics.ProverAllProofRewardGauge.Update(int64(event.Reward))
		if isNormalProof {
			metrics.ProverNormalProofRewardGauge.Update(int64(event.Reward))
		}
	}

	p.latestVerifiedL1Height = event.Raw.BlockNumber

	log.Info(
		"New verified block",
		"blockID", event.Id,
		"hash", common.BytesToHash(event.BlockHash[:]),
		"reward", event.Reward,
	)

	// cancel any proofs being generated for this block
	p.cancelProof(ctx, event.Id.Uint64())

	return nil
}

// onBlockProven cancels proof generation if the proof is being generated by this prover,
// and the proof is not the oracle proof address.
func (p *Prover) onBlockProven(ctx context.Context, event *bindings.TaikoL1ClientBlockProven) error {
	metrics.ProverReceivedProvenBlockGauge.Update(event.Id.Int64())
	// if this proof is submitted by an oracle prover or a system prover, don't cancel proof.
	if event.Prover == p.oracleProverAddress ||
		event.Prover == p.systemProverAddress ||
		event.Prover == encoding.SystemProverAddress ||
		event.Prover == encoding.OracleProverAddress {
		return nil
	}

	// cancel any proofs being generated for this block
	if err := p.cancelProofIfValid(ctx, event.Id.Uint64(), uint64(event.ParentGasUsed), event.ParentHash); err != nil {
		return err
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

	stateVars, err := p.rpc.GetProtocolStateVariables(nil)
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
	stateVars, err := p.rpc.GetProtocolStateVariables(nil)
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
	if (!p.cfg.SystemProver && !p.cfg.OracleProver) || lastLatestVerifiedL1Height != p.latestVerifiedL1Height {
		return nil
	}

	log.Warn(
		"No new block verification in `proofCooldownPeriod * 2` seconeds",
		"latestVerifiedL1Height", p.latestVerifiedL1Height,
		"proofCooldownPeriod", p.protocolConfigs.ProofCooldownPeriod,
	)

	stateVar, err := p.rpc.TaikoL1.GetStateVariables(nil)
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

// cancelProofIfValid cancels proof only if the parentGasUsed and parentHash in the proof match what
// is expected
func (p *Prover) cancelProofIfValid(
	ctx context.Context,
	blockID uint64,
	parentGasUsed uint64,
	parentHash common.Hash,
) error {
	parent, err := p.rpc.L2ParentByBlockId(ctx, new(big.Int).SetUint64(blockID))
	if err != nil {
		return err
	}

	if parent.GasUsed == parentGasUsed && parent.Hash() == parentHash {
		p.cancelProof(ctx, blockID)
	}

	return nil
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
