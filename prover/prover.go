package prover

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikochain/taiko-client/bindings"
	"github.com/taikochain/taiko-client/cmd/utils"
	"github.com/taikochain/taiko-client/pkg/rpc"
	"github.com/taikochain/taiko-client/prover/producer"
	"github.com/urfave/cli/v2"
)

var (
	// errInvalidProposeBlockTx is returned when the given `proposeBlock`
	// transaction is invalid.
	errInvalidProposeBlockTx = errors.New("invalid propose block transaction")
	// Gas limit of TaikoL1.proveBlock and TaikoL1.proveBlockInvalid transactions.
	// TODO: tune this value
	proveBlocksGasLimit uint64 = 1000000
)

// Action returns the main function that the subcommand should run.
func Action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		cfg, err := NewConfigFromCliContext(ctx)
		if err != nil {
			return err
		}

		prover, err := New(cfg)
		if err != nil {
			return err
		}

		return utils.RunSubcommand(prover)
	}
}

// Prover keep trying to prove new proposed blocks valid/invalid.
type Prover struct {
	// Configurations
	cfg *Config

	// Clients
	rpc        *rpc.Client
	taikoL1Abi *abi.ABI
	taikoL2Abi *abi.ABI

	// Contract configurations
	txListValidator  *TxListValidator
	anchorGasLimit   uint64
	maxPendingBlocks uint64
	chainID          *big.Int

	// States
	lastFinalizedHeight *big.Int
	lastFinalizedHeader *types.Header

	// Subscriptions
	blockProposedCh   chan *bindings.TaikoL1ClientBlockProposed
	blockProposedSub  event.Subscription
	blockFinalizedCh  chan *bindings.TaikoL1ClientBlockFinalized
	blockFinalizedSub event.Subscription

	// Prover related
	proveValidResultCh   chan *producer.ProofWithHeader
	proveInvalidResultCh chan *producer.ProofWithHeader
	proofProducer        producer.ProofProducer

	ctx   context.Context
	close context.CancelFunc
	wg    sync.WaitGroup

	// For testing
	blockProposedEventsBuffer []*bindings.TaikoL1ClientBlockProposed
}

// New initializes a new prover instance based on the given configurations.
func New(cfg *Config) (*Prover, error) {
	log.Info("Prover configurations", "config", cfg)
	var err error

	p := &Prover{cfg: cfg}
	p.ctx, p.close = context.WithCancel(context.Background())

	// Clients
	if p.rpc, err = rpc.NewClient(p.ctx, &rpc.ClientConfig{
		L1Endpoint:     cfg.L1Endpoint,
		L2Endpoint:     cfg.L2Endpoint,
		TaikoL1Address: cfg.TaikoL1Address,
		TaikoL2Address: cfg.TaikoL2Address,
	}); err != nil {
		return nil, err
	}

	if p.taikoL1Abi, err = bindings.TaikoL1ClientMetaData.GetAbi(); err != nil {
		return nil, err
	}

	if p.taikoL2Abi, err = bindings.V1TaikoL2ClientMetaData.GetAbi(); err != nil {
		return nil, err
	}

	// Constants
	chainID, maxPendingBlocks, _, _, _,
		maxBlocksGasLimit, maxBlockNumTxs, _, maxTxlistBytes, minTxGasLimit,
		anchorGasLimit, _, _, err := p.rpc.TaikoL1.GetConstants(nil)
	if err != nil {
		return nil, err
	}

	log.Info(
		"LibConstants configurations",
		"maxBlocksGasLimit", maxBlocksGasLimit,
		"maxBlockNumTxs", maxBlockNumTxs,
		"maxTxlistBytes", maxTxlistBytes,
		"maxPendingBlocks", maxPendingBlocks,
		"anchorGasLimit", anchorGasLimit,
		"chainID", chainID,
	)

	p.txListValidator = &TxListValidator{
		maxBlocksGasLimit: maxBlocksGasLimit.Uint64(),
		maxBlockNumTxs:    maxBlockNumTxs.Uint64(),
		maxTxlistBytes:    maxTxlistBytes.Uint64(),
		minTxGasLimit:     minTxGasLimit.Uint64(),
		chainID:           chainID,
	}
	p.maxPendingBlocks = maxPendingBlocks.Uint64()
	p.anchorGasLimit = anchorGasLimit.Uint64()
	p.chainID = chainID
	p.blockProposedCh = make(chan *bindings.TaikoL1ClientBlockProposed, p.maxPendingBlocks)
	p.blockFinalizedCh = make(chan *bindings.TaikoL1ClientBlockFinalized, p.maxPendingBlocks)
	p.proveValidResultCh = make(chan *producer.ProofWithHeader, p.maxPendingBlocks)
	p.proveInvalidResultCh = make(chan *producer.ProofWithHeader, p.maxPendingBlocks)

	if cfg.Dummy {
		p.proofProducer = new(producer.DummyProofProducer)
	} else {
		if p.proofProducer, err = producer.NewZkevmRpcdProducer(cfg.ZKEvmRpcdEndpoint); err != nil {
			return nil, err
		}
	}

	return p, nil
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
	ticker := time.NewTicker(time.Minute)
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
		case proofWithHeader := <-p.proveValidResultCh:
			if err := p.submitValidBlockProof(proofWithHeader); err != nil {
				log.Error("Prove valid block error", "error", err)
			}
		case proofWithHeader := <-p.proveInvalidResultCh:
			if err := p.submitInvalidBlockProof(proofWithHeader); err != nil {
				log.Error("Prove invalid block error", "error", err)
			}
		}
	}
}

// Close closes the prover instance.
func (p *Prover) Close() {
	if p.close != nil {
		p.close()
	}
	p.wg.Wait()
}

// onBlockProposed tries to prove that the newly proposed block is valid/invalid.
func (p *Prover) onBlockProposed(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error {
	log.Info("New proposed block", "blockID", event.Id)

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

	p.lastFinalizedHeight = l2BlockHeader.Number
	p.lastFinalizedHeader = l2BlockHeader

	return nil
}

// onForceTimer fetches the oldest unproved block and tries to prove it.
func (p *Prover) onForceTimer(ctx context.Context) error {
	_, _, latestFinalizedID, nextBlockID, err := p.rpc.TaikoL1.GetStateVariables(nil)
	if err != nil {
		return fmt.Errorf("failed to get TaikoL1 state variables: %w", err)
	}

	log.Info("TaikoL1 state variables",
		"latestFinalizedID", latestFinalizedID,
		"nextBlockID", nextBlockID,
	)

	var oldestUnprovedBlockID *big.Int
	for i := latestFinalizedID + 1; i < nextBlockID; i++ {
		proposedBlock, err := p.rpc.TaikoL1.GetProposedBlock(nil, new(big.Int).SetUint64(i))
		if err != nil {
			return fmt.Errorf("failed to get proposed block: %w", err)
		}

		// enum EverProven {
		// 	_NO, //=0
		// 	NO, //=1
		// 	YES //=2
		// }
		if proposedBlock.EverProven == 2 {
			oldestUnprovedBlockID = new(big.Int).SetUint64(i)
			break
		}
	}

	if oldestUnprovedBlockID == nil {
		log.Info("All proposed blocks are proved")
		return nil
	}

	log.Info("Oldest unproved block ID", "blockID", oldestUnprovedBlockID)

	l1Origin, err := p.rpc.L1.L1OriginByID(ctx, oldestUnprovedBlockID)
	if err != nil {
		return fmt.Errorf("failed to get L1Origin: %w", err)
	}

	blockProposedL1Height := l1Origin.L1BlockHeight.Uint64()
	iter, err := p.rpc.TaikoL1.FilterBlockProposed(
		&bind.FilterOpts{
			Start: blockProposedL1Height,
			End:   &blockProposedL1Height,
		},
		[]*big.Int{oldestUnprovedBlockID},
	)
	if err != nil {
		return fmt.Errorf("failed to filter BlockProposed events: %w", err)
	}

	for iter.Next() {
		return p.onBlockProposed(ctx, iter.Event)
	}

	return fmt.Errorf("BlockProposed events not found, blockID: %d, l1Height: %d",
		oldestUnprovedBlockID, blockProposedL1Height,
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
		log.Info("New BlockProposed event buffered", "blockID", newEvent.Id, "size", len(p.blockProposedEventsBuffer))
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(p.blockProposedEventsBuffer), func(i, j int) {
		p.blockProposedEventsBuffer[i], p.blockProposedEventsBuffer[j] =
			p.blockProposedEventsBuffer[j], p.blockProposedEventsBuffer[i]
	})

	for i := 0; i < len(p.blockProposedEventsBuffer); i++ {
		if err := p.onBlockProposed(p.ctx, p.blockProposedEventsBuffer[i]); err != nil {
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

// getProveBlocksTxOpts creates a bind.TransactOpts instance with the sender's signatures
// for TaikoL1.proveBlock and TaikoL1.proveBlockInvalid transactions.
func (p *Prover) getProveBlocksTxOpts(ctx context.Context) (*bind.TransactOpts, error) {
	networkID, err := p.rpc.L1.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(p.cfg.L1ProverPrivKey, networkID)
	if err != nil {
		return nil, err
	}

	opts.GasLimit = proveBlocksGasLimit

	return opts, nil
}
