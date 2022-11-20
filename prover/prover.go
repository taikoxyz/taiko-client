package prover

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/prover/producer"
	"github.com/urfave/cli/v2"
)

var (
	// Gas limit of TaikoL1.proveBlock and TaikoL1.proveBlockInvalid transactions.
	// TODO: tune this value based when the on-chain solidity verifier is available.
	proveBlocksGasLimit uint64 = 1000000
	// Time interval to fetch the unproved blocks and prove them, even if no BlockProposed
	// event received.
	forceTimerCycle = 1 * time.Minute
)

// Prover keep trying to prove new proposed blocks valid/invalid.
type Prover struct {
	// Configurations
	cfg *Config

	// Clients
	rpc *rpc.Client

	// Contract configurations
	txListValidator  *TxListValidator
	anchorGasLimit   uint64
	maxPendingBlocks uint64

	// States
	lastFinalizedHeader *types.Header

	// Subscriptions
	blockProposedCh   chan *bindings.TaikoL1ClientBlockProposed
	blockProposedSub  event.Subscription
	blockFinalizedCh  chan *bindings.TaikoL1ClientBlockFinalized
	blockFinalizedSub event.Subscription

	// Proof related
	proveValidProofCh   chan *producer.ProofWithHeader
	proveInvalidProofCh chan *producer.ProofWithHeader
	proofProducer       producer.ProofProducer

	ctx context.Context
	wg  sync.WaitGroup

	// For testing
	blockProposedEventsBuffer []*bindings.TaikoL1ClientBlockProposed
}

// New initializes the given prover instance based on the command line flags.
func (p *Prover) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return initFromConfig(ctx, p, cfg)
}

// initFromConfig initializes the prover instance based on the given configurations.
func initFromConfig(ctx context.Context, p *Prover, cfg *Config) (err error) {
	log.Debug("Prover configurations", "config", cfg)

	p.cfg = cfg
	p.ctx = ctx

	// Clients
	if p.rpc, err = rpc.NewClient(p.ctx, &rpc.ClientConfig{
		L1Endpoint:     cfg.L1Endpoint,
		L2Endpoint:     cfg.L2Endpoint,
		TaikoL1Address: cfg.TaikoL1Address,
		TaikoL2Address: cfg.TaikoL2Address,
	}); err != nil {
		return err
	}

	// Constants
	_, maxPendingBlocks, _, _, _,
		maxBlocksGasLimit, maxBlockNumTxs, _, maxTxlistBytes, minTxGasLimit,
		anchorGasLimit, _, _, err := p.rpc.TaikoL1.GetConstants(nil)
	if err != nil {
		return err
	}

	log.Info(
		"LibConstants configurations",
		"maxBlocksGasLimit", maxBlocksGasLimit,
		"maxBlockNumTxs", maxBlockNumTxs,
		"maxTxlistBytes", maxTxlistBytes,
		"maxPendingBlocks", maxPendingBlocks,
		"anchorGasLimit", anchorGasLimit,
	)

	p.txListValidator = NewTxListValidator(
		maxBlocksGasLimit.Uint64(),
		maxBlockNumTxs.Uint64(),
		maxTxlistBytes.Uint64(),
		minTxGasLimit.Uint64(),
		p.rpc.L2ChainID,
	)
	p.maxPendingBlocks = maxPendingBlocks.Uint64()
	p.anchorGasLimit = anchorGasLimit.Uint64()
	p.blockProposedCh = make(chan *bindings.TaikoL1ClientBlockProposed, 10)
	p.blockFinalizedCh = make(chan *bindings.TaikoL1ClientBlockFinalized, 10)
	p.proveValidProofCh = make(chan *producer.ProofWithHeader, 10)
	p.proveInvalidProofCh = make(chan *producer.ProofWithHeader, 10)

	if cfg.Dummy {
		p.proofProducer = new(producer.DummyProofProducer)
	} else {
		if p.proofProducer, err = producer.NewZkevmRpcdProducer(cfg.ZKEvmRpcdEndpoint); err != nil {
			return err
		}
	}

	return nil
}

// Start starts the main loop of the L2 block prover.
func (p *Prover) Start() error {
	p.wg.Add(1)
	p.startSubscription()
	go p.eventLoop()

	return nil
}

// eventLoop starts the main loop of Taiko prover.
func (p *Prover) eventLoop() {
	ticker := time.NewTicker(forceTimerCycle)
	defer func() {
		ticker.Stop()
		p.wg.Done()
	}()

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-ticker.C:
			if err := p.onForceTimer(p.ctx); err != nil {
				log.Error("Handle forceTimer event error", "error", err)
			}
		case e := <-p.blockProposedCh:
			if p.cfg.BatchSubmit {
				if err := p.batchHandleBlockProposedEvents(p.ctx, e); err != nil {
					log.Error("Batch handling BlockProposed event error", "error", err)
				}
				continue
			}

			if err := p.onBlockProposed(p.ctx, e); err != nil {
				log.Error("Handle BlockProposed event error", "error", err)
			}
		case e := <-p.blockFinalizedCh:
			if err := p.onBlockFinalized(p.ctx, e); err != nil {
				log.Error("Handle BlockFinalized event error", "error", err)
			}
		case proofWithHeader := <-p.proveValidProofCh:
			if err := p.submitValidBlockProof(p.ctx, proofWithHeader); err != nil {
				log.Error("Prove valid block error", "error", err)
			}
		case proofWithHeader := <-p.proveInvalidProofCh:
			if err := p.submitInvalidBlockProof(p.ctx, proofWithHeader); err != nil {
				log.Error("Prove invalid block error", "error", err)
			}
		}
	}
}

// Close closes the prover instance.
func (p *Prover) Close() {
	p.closeSubscription()
	p.wg.Wait()
}

// onBlockProposed tries to prove that the newly proposed block is valid/invalid.
func (p *Prover) onBlockProposed(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error {
	log.Info("New proposed block", "blockID", event.Id)
	metrics.ProverReceivedProposedBlockGauge.Update(event.Id.Int64())

	proposeBlockTx, err := p.rpc.L1.TransactionInBlock(ctx, event.Raw.BlockHash, event.Raw.TxIndex)
	if err != nil {
		return err
	}

	// Check whether the transactions list is valid.
	hint, invalidTxIndex, err := p.txListValidator.ValidateTxList(event.Id, proposeBlockTx.Data())
	if err != nil {
		return err
	}

	// Prove the proposed block valid.
	if hint == HintOK {
		return p.proveBlockValid(ctx, event)
	}

	// Prove the proposed block invalid.
	return p.proveBlockInvalid(ctx, event, hint, invalidTxIndex)
}

// onBlockFinalized update the lastFinalized block in current state.
func (p *Prover) onBlockFinalized(ctx context.Context, event *bindings.TaikoL1ClientBlockFinalized) error {
	if event.BlockHash == (common.Hash{}) {
		log.Info("Ignore BlockFinalized event of invalid transaction list", "blockID", event.Id)
		return nil
	}

	metrics.ProverLatestFinalizedIDGauge.Update(event.Id.Int64())

	l2BlockHeader, err := p.rpc.L2.HeaderByHash(ctx, event.BlockHash)
	if err != nil {
		return fmt.Errorf("failed to find L2 block with hash %s: %w", common.BytesToHash(event.BlockHash[:]), err)
	}

	log.Info(
		"New finalized block",
		"blockID", event.Id,
		"height", l2BlockHeader.Number,
		"hash", common.BytesToHash(event.BlockHash[:]),
	)
	p.lastFinalizedHeader = l2BlockHeader

	return nil
}

// onForceTimer fetches the oldest unfinalized block and if it is still not proven, then prove it.
func (p *Prover) onForceTimer(ctx context.Context) error {
	_, _, latestFinalizedID, nextBlockID, err := p.rpc.TaikoL1.GetStateVariables(nil)
	if err != nil {
		return fmt.Errorf("failed to get TaikoL1 state variables: %w", err)
	}

	log.Debug("TaikoL1 state variables", "latestFinalizedID", latestFinalizedID, "nextBlockID", nextBlockID)

	if latestFinalizedID == nextBlockID-1 {
		log.Info("All proposed blocks are finalized")
		return nil
	}

	// Check whether the oldest unfinalized block is still unproved.
	oldestUnfinalizedBlockID := new(big.Int).SetUint64(latestFinalizedID + 1)

	proposedBlock, err := p.rpc.TaikoL1.GetProposedBlock(nil, oldestUnfinalizedBlockID)
	if err != nil {
		return fmt.Errorf("failed to get proposed block: %w", err)
	}

	commitHeight, err := p.rpc.TaikoL1.GetCommitHeight(nil, proposedBlock.MetaHash)
	if err != nil {
		return fmt.Errorf("failed to get proposed block commit height: %w", err)
	}

	blockProvenIter, err := p.rpc.TaikoL1.FilterBlockProven(
		&bind.FilterOpts{Start: commitHeight.Uint64()},
		[]*big.Int{oldestUnfinalizedBlockID},
	)
	if err != nil {
		return fmt.Errorf("failed to filter BlockProven events: %w", err)
	}

	if blockProvenIter.Next() {
		return nil
	}

	log.Info("Oldest unproved block ID", "blockID", oldestUnfinalizedBlockID)

	l1Origin, err := p.rpc.L2.L1OriginByID(ctx, oldestUnfinalizedBlockID)
	if err != nil {
		return fmt.Errorf("failed to get L1Origin: %w", err)
	}

	filterHeight := l1Origin.L1BlockHeight.Uint64()
	blockProposedIter, err := p.rpc.TaikoL1.FilterBlockProposed(
		&bind.FilterOpts{Start: filterHeight, End: &filterHeight},
		[]*big.Int{oldestUnfinalizedBlockID},
	)
	if err != nil {
		return fmt.Errorf("failed to filter BlockProposed events: %w", err)
	}

	for blockProposedIter.Next() {
		return p.onBlockProposed(ctx, blockProposedIter.Event)
	}

	return fmt.Errorf("BlockProposed events not found, blockID: %d, l1Height: %d",
		oldestUnfinalizedBlockID, filterHeight,
	)
}

// batchHandleBlockProposedEvents will randomly handle buffered BlockProposed
// events if the buffer size reaches `maxPendingBlocks`.
func (p *Prover) batchHandleBlockProposedEvents(
	ctx context.Context,
	newEvent *bindings.TaikoL1ClientBlockProposed,
) error {
	p.blockProposedEventsBuffer = append(p.blockProposedEventsBuffer, newEvent)

	if len(p.blockProposedEventsBuffer) < int(p.maxPendingBlocks) {
		log.Debug("New BlockProposed event buffered", "blockID", newEvent.Id, "size", len(p.blockProposedEventsBuffer))
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(p.blockProposedEventsBuffer), func(i, j int) {
		p.blockProposedEventsBuffer[i], p.blockProposedEventsBuffer[j] =
			p.blockProposedEventsBuffer[j], p.blockProposedEventsBuffer[i]
	})

	for i := 0; i < len(p.blockProposedEventsBuffer); i++ {
		if err := p.onBlockProposed(ctx, p.blockProposedEventsBuffer[i]); err != nil {
			return err
		}
	}

	p.blockProposedEventsBuffer = []*bindings.TaikoL1ClientBlockProposed{}

	return nil
}

// Name returns the application name.
func (p *Prover) Name() string {
	return "prover"
}

// getProveBlocksTxOpts creates a bind.TransactOpts instance using the given private key.
// Used for creating TaikoL1.proveBlock and TaikoL1.proveBlockInvalid transactions.
func (p *Prover) getProveBlocksTxOpts(ctx context.Context) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(p.cfg.L1ProverPrivKey, p.rpc.L1ChainID)
	if err != nil {
		return nil, err
	}

	opts.GasLimit = proveBlocksGasLimit

	return opts, nil
}
