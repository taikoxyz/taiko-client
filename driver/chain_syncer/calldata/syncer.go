package calldata

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	anchorTxConstructor "github.com/taikoxyz/taiko-client/driver/anchor_tx_constructor"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/beaconsync"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/metrics"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	txListValidator "github.com/taikoxyz/taiko-client/pkg/tx_list_validator"
)

// ParentBlockInfo is an abstraction between *types.Header and our code, to allow passing in
// a zero hash and other zero-value fields.
type ParentBlockInfo struct {
	Hash    common.Hash
	Number  *big.Int
	GasUsed uint64
}

// Syncer responsible for letting the L2 execution engine catching up with protocol's latest
// pending block through deriving L1 calldata.
type Syncer struct {
	ctx               context.Context
	rpc               *rpc.Client
	state             *state.State
	progressTracker   *beaconsync.SyncProgressTracker          // Sync progress tracker
	anchorConstructor *anchorTxConstructor.AnchorTxConstructor // TaikoL2.anchor transactions constructor
	txListValidator   *txListValidator.TxListValidator         // Transactions list validator
	// Used by BlockInserter
	lastInsertedBlockID *big.Int
}

// NewSyncer creates a new syncer instance.
func NewSyncer(
	ctx context.Context,
	rpc *rpc.Client,
	state *state.State,
	progressTracker *beaconsync.SyncProgressTracker,
	signalServiceAddress common.Address,
) (*Syncer, error) {
	configs, err := rpc.TaikoL1.GetConfig(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get protocol configs: %w", err)
	}

	constructor, err := anchorTxConstructor.New(rpc, signalServiceAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize anchor constructor: %w", err)
	}

	return &Syncer{
		ctx:               ctx,
		rpc:               rpc,
		state:             state,
		progressTracker:   progressTracker,
		anchorConstructor: constructor,
		txListValidator: txListValidator.NewTxListValidator(
			configs.BlockMaxGasLimit.Uint64(),
			configs.MaxTransactionsPerBlock.Uint64(),
			configs.MaxBytesPerTxList.Uint64(),
			configs.MinTxGasLimit.Uint64(),
			rpc.L2ChainID,
		),
	}, nil
}

// ProcessL1Blocks fetches all `TaikoL1.BlockProposed` events between given
// L1 block heights, and then tries inserting them into L2 execution engine's block chain.
func (s *Syncer) ProcessL1Blocks(ctx context.Context, l1End *types.Header) error {
	iter, err := eventIterator.NewBlockProposedIterator(ctx, &eventIterator.BlockProposedIteratorConfig{
		Client:               s.rpc.L1,
		TaikoL1:              s.rpc.TaikoL1,
		StartHeight:          s.state.GetL1Current().Number,
		EndHeight:            l1End.Number,
		FilterQuery:          nil,
		OnBlockProposedEvent: s.onBlockProposed,
	})
	if err != nil {
		return err
	}

	if err := iter.Iter(); err != nil {
		return err
	}

	s.state.SetL1Current(l1End)
	metrics.DriverL1CurrentHeightGauge.Update(s.state.GetL1Current().Number.Int64())

	return nil
}

// OnBlockProposed is a `BlockProposed` event callback which responsible for
// inserting the proposed block one by one to the L2 execution engine.
func (s *Syncer) onBlockProposed(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	endIter eventIterator.EndBlockProposedEventIterFunc,
) error {
	// Ignore those already inserted blocks.
	if event.Id.Cmp(common.Big0) == 0 || (s.lastInsertedBlockID != nil && event.Id.Cmp(s.lastInsertedBlockID) <= 0) {
		return nil
	}

	log.Info(
		"New BlockProposed event",
		"L1Height", event.Raw.BlockNumber,
		"L1Hash", event.Raw.BlockHash,
		"BlockID", event.Id,
		"Removed", event.Raw.Removed,
	)

	// handle reorg
	if event.Raw.Removed {
		return s.handleReorg(ctx, event)
	}

	// Fetch the L2 parent block.
	var (
		parent *types.Header
		err    error
	)
	if s.progressTracker.Triggered() {
		// Already synced through beacon sync, just skip this event.
		if event.Id.Cmp(s.progressTracker.LastSyncedVerifiedBlockID()) <= 0 {
			return nil
		}

		parent, err = s.rpc.L2.HeaderByHash(ctx, s.progressTracker.LastSyncedVerifiedBlockHash())
	} else {
		parent, err = s.rpc.L2ParentByBlockId(ctx, event.Id)
	}

	if err != nil {
		return fmt.Errorf("failed to fetch L2 parent block: %w", err)
	}

	log.Debug("Parent block", "height", parent.Number, "hash", parent.Hash())

	tx, err := s.rpc.L1.TransactionInBlock(
		ctx,
		event.Raw.BlockHash,
		event.Raw.TxIndex,
	)
	if err != nil {
		return fmt.Errorf("failed to fetch original TaikoL1.proposeBlock transaction: %w", err)
	}

	// Check whether the transactions list is valid.
	txListBytes, hint, invalidTxIndex, err := s.txListValidator.ValidateTxList(event.Id, tx.Data())
	if err != nil {
		return fmt.Errorf("failed to validate transactions list: %w", err)
	}

	log.Info(
		"Validate transactions list",
		"blockID", event.Id,
		"hint", hint,
		"invalidTxIndex", invalidTxIndex,
	)

	l1Origin := &rawdb.L1Origin{
		BlockID:       event.Id,
		L2BlockHash:   common.Hash{}, // Will be set by taiko-geth.
		L1BlockHeight: new(big.Int).SetUint64(event.Raw.BlockNumber),
		L1BlockHash:   event.Raw.BlockHash,
	}

	if event.Meta.Timestamp > uint64(time.Now().Unix()) {
		log.Warn("Future L2 block, waiting", "L2BlockTimestamp", event.Meta.Timestamp, "now", time.Now().Unix())
		time.Sleep(time.Until(time.Unix(int64(event.Meta.Timestamp), 0)))
	}

	// If the transactions list is invalid, we simply insert an empty L2 block.
	if hint != txListValidator.HintOK {
		log.Info("Invalid transactions list, insert an empty L2 block instead", "blockID", event.Id)
		txListBytes = []byte{}
	}

	payloadData, rpcError, payloadError := s.insertNewHead(
		ctx,
		event,
		&ParentBlockInfo{
			Hash:    parent.Hash(),
			Number:  parent.Number,
			GasUsed: parent.GasUsed,
		},
		s.state.GetHeadBlockID(),
		txListBytes,
		l1Origin,
	)

	// RPC errors are recoverable.
	if rpcError != nil {
		return fmt.Errorf("failed to insert new head to L2 execution engine: %w", rpcError)
	}

	if payloadError != nil {
		log.Warn(
			"Ignore invalid block context", "blockID", event.Id, "payloadError", payloadError, "payloadData", payloadData,
		)
		return nil
	}

	log.Debug("Payload data", "hash", payloadData.BlockHash, "txs", len(payloadData.Transactions))

	log.Info(
		"🔗 New L2 block inserted",
		"blockID", event.Id,
		"height", payloadData.Number,
		"hash", payloadData.BlockHash,
		"latestVerifiedBlockHeight", s.state.GetLatestVerifiedBlock().Height,
		"latestVerifiedBlockHash", s.state.GetLatestVerifiedBlock().Hash,
		"transactions", len(payloadData.Transactions),
		"baseFee", payloadData.BaseFeePerGas,
		"withdrawals", len(payloadData.Withdrawals),
	)

	metrics.DriverL1CurrentHeightGauge.Update(int64(event.Raw.BlockNumber))
	s.lastInsertedBlockID = event.Id

	if s.progressTracker.Triggered() {
		s.progressTracker.ClearMeta()
	}

	return nil
}

// handleReorg detects reorg and rewinds the chain by 1 until we find a block that is still in the chain,
// then inserts that block as the new head.
func (s *Syncer) handleReorg(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error {
	log.Info(
		"Reorg detected",
		"L1Height", event.Raw.BlockNumber,
		"L1Hash", event.Raw.BlockHash,
		"BlockID", event.Id,
		"Removed", event.Raw.Removed,
	)

	// rewind chain by 1 until we find a block that is still in the chain
	var lastKnownGoodBlockId *big.Int
	var blockId *big.Int = s.lastInsertedBlockID
	var block *types.Block
	var err error

	stateVars, err := s.rpc.GetProtocolStateVariables()
	if err != nil {
		return fmt.Errorf("failed to get state variables: %w", err)
	}

	for {
		if blockId.Cmp(big.NewInt(0)) == 0 {
			lastKnownGoodBlockId = new(big.Int).SetUint64(0)
			break
		}

		block, err = s.rpc.L2.BlockByNumber(ctx, blockId)
		if err != nil && !errors.Is(err, ethereum.NotFound) {
			return err
		}

		if block != nil && blockId.Uint64() < stateVars.NumBlocks {
			// block exists, we can rewind to this block
			lastKnownGoodBlockId = blockId
			break
		} else {
			// otherwise, sub 1 from blockId and try again
			blockId = new(big.Int).Sub(s.lastInsertedBlockID, big.NewInt(1))
		}
	}

	// shouldn't be able to reach this error because of the 0 check above
	// but just in case
	if lastKnownGoodBlockId == nil {
		return fmt.Errorf("failed to find last known good block ID after reorg")
	}

	log.Info(
		"🔗 Last known good block ID before reorg found",
		"blockID", lastKnownGoodBlockId,
	)

	var parentBlockInfo *ParentBlockInfo

	if lastKnownGoodBlockId.Cmp(common.Big0) == 0 {
		parentBlockInfo = &ParentBlockInfo{
			Hash:    common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
			Number:  big.NewInt(0),
			GasUsed: 0,
		}
	} else {
		parent, err := s.rpc.L2ParentByBlockId(ctx, lastKnownGoodBlockId)
		if err != nil {
			return fmt.Errorf("error getting l2 parent by block id: %w", err)
		}

		parentBlockInfo = &ParentBlockInfo{
			Hash:    parent.Hash(),
			Number:  parent.Number,
			GasUsed: parent.GasUsed,
		}
	}

	// reset l1 current to when the last known good block was inserted, and return the event.
	blockProposedEvent, _, err := s.state.ResetL1Current(ctx, &state.HeightOrID{Height: block.Number()})
	if err != nil {
		return fmt.Errorf("failed to reset l1 current: %w", err)
	}

	tx, err := s.rpc.L1.TransactionInBlock(
		ctx,
		block.Hash(),
		blockProposedEvent.Raw.TxIndex,
	)
	if err != nil {
		return fmt.Errorf("failed to fetch original TaikoL1.proposeBlock transaction: %w", err)
	}

	txListBytes, hint, _, err := s.txListValidator.ValidateTxList(block.Number(), tx.Data())
	if err != nil {
		return fmt.Errorf("failed to validate transactions list: %w", err)
	}

	if hint != txListValidator.HintOK {
		log.Info("Invalid transactions list, insert an empty L2 block instead", "blockID", block.NumberU64())
		txListBytes = []byte{}
	}

	l1Origin := &rawdb.L1Origin{
		BlockID:       block.Number(),
		L2BlockHash:   common.Hash{}, // Will be set by taiko-geth.
		L1BlockHeight: new(big.Int).SetUint64(blockProposedEvent.Raw.BlockNumber),
		L1BlockHash:   blockProposedEvent.Raw.BlockHash,
	}

	payloadData, rpcError, payloadError := s.insertNewHead(
		ctx,
		blockProposedEvent,
		parentBlockInfo,
		s.state.GetHeadBlockID(),
		txListBytes,
		l1Origin,
	)

	if rpcError != nil {
		return fmt.Errorf("failed to insert new head to L2 execution engine: %w", rpcError)
	}

	if payloadError != nil {
		log.Warn(
			"Ignore invalid block context", "blockID", event.Id, "payloadError", payloadError, "payloadData", payloadData,
		)
		return nil
	}

	log.Debug("Payload data", "hash", payloadData.BlockHash, "txs", len(payloadData.Transactions))

	log.Info(
		"🔗 Rewound chain and inserted last known good block as new head",
		"blockID", event.Id,
		"height", payloadData.Number,
		"hash", payloadData.BlockHash,
		"latestVerifiedBlockHeight", s.state.GetLatestVerifiedBlock().Height,
		"latestVerifiedBlockHash", s.state.GetLatestVerifiedBlock().Hash,
		"transactions", len(payloadData.Transactions),
		"baseFee", payloadData.BaseFeePerGas,
		"withdrawals", len(payloadData.Withdrawals),
	)

	metrics.DriverL1CurrentHeightGauge.Update(int64(event.Raw.BlockNumber))
	s.lastInsertedBlockID = block.Number()

	if s.progressTracker.Triggered() {
		s.progressTracker.ClearMeta()
	}

	return nil
}

// insertNewHead tries to insert a new head block to the L2 execution engine's local
// block chain through Engine APIs.
func (s *Syncer) insertNewHead(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	parentBlockInfo *ParentBlockInfo,
	headBlockID *big.Int,
	txListBytes []byte,
	l1Origin *rawdb.L1Origin,
) (*engine.ExecutableData, error, error) {
	log.Debug(
		"Try to insert a new L2 head block",
		"parentNumber", parentBlockInfo.Number,
		"parentHash", parentBlockInfo.Hash,
		"headBlockID", headBlockID,
		"l1Origin", l1Origin,
	)

	// Insert a TaikoL2.anchor transaction at transactions list head
	var txList []*types.Transaction
	if len(txListBytes) != 0 {
		if err := rlp.DecodeBytes(txListBytes, &txList); err != nil {
			log.Info("Ignore invalid txList bytes", "blockID", event.Id)
			return nil, nil, err
		}
	}

	parentTimestamp, err := s.rpc.TaikoL2.ParentTimestamp(&bind.CallOpts{BlockNumber: parentBlockInfo.Number})
	if err != nil {
		return nil, nil, err
	}

	// Get L2 baseFee
	baseFee, err := s.rpc.TaikoL2.GetBasefee(
		&bind.CallOpts{BlockNumber: parentBlockInfo.Number},
		uint32(event.Meta.Timestamp-parentTimestamp),
		uint64(event.Meta.GasLimit+uint32(s.anchorConstructor.GasLimit())),
		parentBlockInfo.GasUsed,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get L2 baseFee: %w", encoding.TryParsingCustomError(err))
	}

	log.Debug(
		"GetBasefee",
		"baseFee", baseFee,
		"timeSinceParent", uint32(event.Meta.Timestamp-parentTimestamp),
		"gasLimit", uint64(event.Meta.GasLimit+uint32(s.anchorConstructor.GasLimit())),
		"parentGasUsed", parentBlockInfo.GasUsed,
	)

	// Get withdrawals
	withdrawals := make(types.Withdrawals, len(event.Meta.DepositsProcessed))
	for i, d := range event.Meta.DepositsProcessed {
		withdrawals[i] = &types.Withdrawal{Address: d.Recipient, Amount: d.Amount.Uint64()}
	}

	// Assemble a TaikoL2.anchor transaction
	anchorTx, err := s.anchorConstructor.AssembleAnchorTx(
		ctx,
		new(big.Int).SetUint64(event.Meta.L1Height),
		event.Meta.L1Hash,
		new(big.Int).Add(parentBlockInfo.Number, common.Big1),
		baseFee,
		parentBlockInfo.GasUsed,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create TaikoL2.anchor transaction: %w", err)
	}

	txList = append([]*types.Transaction{anchorTx}, txList...)

	if txListBytes, err = rlp.EncodeToBytes(txList); err != nil {
		log.Warn("Encode txList error", "blockID", event.Id, "error", err)
		return nil, nil, err
	}

	payload, rpcErr, payloadErr := s.createExecutionPayloads(
		ctx,
		event,
		parentBlockInfo.Hash,
		l1Origin,
		headBlockID,
		txListBytes,
		baseFee,
		withdrawals,
	)

	if rpcErr != nil || payloadErr != nil {
		return nil, rpcErr, payloadErr
	}

	fc := &engine.ForkchoiceStateV1{HeadBlockHash: parentBlockInfo.Hash}

	// Update the fork choice
	fc.HeadBlockHash = payload.BlockHash
	fcRes, err := s.rpc.L2Engine.ForkchoiceUpdate(ctx, fc, nil)
	if err != nil {
		return nil, err, nil
	}
	if fcRes.PayloadStatus.Status != engine.VALID {
		return nil, nil, fmt.Errorf("unexpected ForkchoiceUpdate response status: %s", fcRes.PayloadStatus.Status)
	}

	return payload, nil, nil
}

// createExecutionPayloads creates a new execution payloads through
// Engine APIs.
func (s *Syncer) createExecutionPayloads(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	parentHash common.Hash,
	l1Origin *rawdb.L1Origin,
	headBlockID *big.Int,
	txListBytes []byte,
	baseFeee *big.Int,
	withdrawals types.Withdrawals,
) (payloadData *engine.ExecutableData, rpcError error, payloadError error) {
	fc := &engine.ForkchoiceStateV1{HeadBlockHash: parentHash}
	attributes := &engine.PayloadAttributes{
		Timestamp:             event.Meta.Timestamp,
		Random:                event.Meta.MixHash,
		SuggestedFeeRecipient: event.Meta.Beneficiary,
		Withdrawals:           withdrawals,
		BlockMetadata: &engine.BlockMetadata{
			HighestBlockID: headBlockID,
			Beneficiary:    event.Meta.Beneficiary,
			GasLimit:       uint64(event.Meta.GasLimit) + s.anchorConstructor.GasLimit(),
			Timestamp:      event.Meta.Timestamp,
			TxList:         txListBytes,
			MixHash:        event.Meta.MixHash,
			ExtraData:      []byte{},
		},
		BaseFeePerGas: baseFeee,
		L1Origin:      l1Origin,
	}

	log.Debug("PayloadAttributes", "attributes", attributes, "meta", attributes.BlockMetadata)

	// Step 1, prepare a payload
	fcRes, err := s.rpc.L2Engine.ForkchoiceUpdate(ctx, fc, attributes)
	if err != nil {
		return nil, err, nil
	}
	if fcRes.PayloadStatus.Status != engine.VALID {
		return nil, nil, fmt.Errorf("unexpected ForkchoiceUpdate response status: %s", fcRes.PayloadStatus.Status)
	}
	if fcRes.PayloadID == nil {
		return nil, nil, errors.New("empty payload ID")
	}

	// Step 2, get the payload
	payload, err := s.rpc.L2Engine.GetPayload(ctx, fcRes.PayloadID)
	if err != nil {
		return nil, err, nil
	}

	log.Debug("Payload", "payload", payload)

	// Step 3, execute the payload
	execStatus, err := s.rpc.L2Engine.NewPayload(ctx, payload)
	if err != nil {
		return nil, err, nil
	}
	if execStatus.Status != engine.VALID {
		return nil, nil, fmt.Errorf("unexpected NewPayload response status: %s", execStatus.Status)
	}

	return payload, nil, nil
}
