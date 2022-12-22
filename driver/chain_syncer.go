package driver

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/metrics"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	txListValidator "github.com/taikoxyz/taiko-client/pkg/tx_list_validator"
)

// HeightOrID contains a block height or a block ID.
type HeightOrID struct {
	Height *big.Int
	ID     *big.Int
}

// NotEmpty checks whether this is an empty struct.
func (h *HeightOrID) NotEmpty() bool {
	return h.Height != nil || h.ID != nil
}

// L2ChainSyncer is responsible for keeping the L2 execution engine's local chain in sync with the one
// in TaikoL1 contract.
type L2ChainSyncer struct {
	ctx                           context.Context
	state                         *State                           // Driver's state
	rpc                           *rpc.Client                      // L1/L2 RPC clients
	throwawayBlocksBuilderPrivKey *ecdsa.PrivateKey                // Private key of L2 throwaway blocks builder
	txListValidator               *txListValidator.TxListValidator // Transactions list validator
	anchorConstructor             *AnchorConstructor               // TaikoL2.anchor transactions constructor
	protocolConstants             *bindings.ProtocolConstants      // Protocol constants

	// If this flag is activated, will try P2P beacon sync if current node is behind of the protocol's
	// latest verified block head
	p2pSyncVerifiedBlocks bool
	// Monitor the L2 execution engine's sync progress
	syncProgressTracker *BeaconSyncProgressTracker

	// Used by BlockInserter
	lastInsertedBlockID *big.Int
}

// NewL2ChainSyncer creates a new chain syncer instance.
func NewL2ChainSyncer(
	ctx context.Context,
	rpc *rpc.Client,
	state *State,
	throwawayBlocksBuilderPrivKey *ecdsa.PrivateKey,
	p2pSyncVerifiedBlocks bool,
	p2pSyncTimeout time.Duration,
) (*L2ChainSyncer, error) {
	constants, err := rpc.GetProtocolConstants(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get protocol constants: %w", err)
	}

	anchorConstructor, err := NewAnchorConstructor(
		rpc,
		constants.AnchorTxGasLimit.Uint64(),
		bindings.GoldenTouchAddress,
		bindings.GoldenTouchPrivKey,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize anchor constructor: %w", err)
	}

	tracker := NewBeaconSyncProgressTracker(rpc.L2, p2pSyncTimeout)
	go tracker.Track(ctx)

	return &L2ChainSyncer{
		ctx:                           ctx,
		rpc:                           rpc,
		state:                         state,
		throwawayBlocksBuilderPrivKey: throwawayBlocksBuilderPrivKey,
		protocolConstants:             constants,
		txListValidator: txListValidator.NewTxListValidator(
			constants.BlockMaxGasLimit.Uint64(),
			constants.BlockMaxTxs.Uint64(),
			constants.TxListMaxBytes.Uint64(),
			constants.TxMinGasLimit.Uint64(),
			rpc.L2ChainID,
		),
		anchorConstructor:     anchorConstructor,
		p2pSyncVerifiedBlocks: p2pSyncVerifiedBlocks,
		syncProgressTracker:   tracker,
	}, nil
}

// Sync performs a sync operation to L2 execution engine's local chain.
func (s *L2ChainSyncer) Sync(l1End *types.Header) error {
	// If current L2 execution engine's chain is behind of the protocol's latest verified block head, and the
	// `P2PSyncVerifiedBlocks` flag is set, try triggering a beacon sync in L2 execution engine to catch up the
	// latest verified block head.
	if s.p2pSyncVerifiedBlocks &&
		s.state.getLatestVerifiedBlock().Height.Uint64() > 0 &&
		!s.AheadOfProtocolVerifiedHead() &&
		!s.syncProgressTracker.OutOfSync() {
		if err := s.TriggerBeaconSync(); err != nil {
			return fmt.Errorf("trigger beacon sync error: %w", err)
		}

		return nil
	}

	// We have triggered at least a beacon sync in L2 execution engine, we should reset the L1Current
	// cursor at first, before start inserting pending L2 blocks one by one.
	if s.syncProgressTracker.Triggered() {
		log.Info(
			"Switch to insert pending blocks one by one",
			"p2pEnabled", s.p2pSyncVerifiedBlocks,
			"p2pOutOfSync", s.syncProgressTracker.OutOfSync(),
		)

		// Get the execution engine's chain head.
		l2Head, err := s.rpc.L2.HeaderByNumber(s.ctx, nil)
		if err != nil {
			return err
		}

		// Make sure the execution engine's chain head is recorded in protocol.
		l2HeadHash, err := s.rpc.TaikoL1.GetSyncedHeader(nil, l2Head.Number)
		if err != nil {
			return err
		}

		heightOrID := &HeightOrID{Height: l2Head.Number}
		// If there is a verified block hash mismatch, log the error and then try to re-sync from genesis one by one.
		if l2Head.Hash() != l2HeadHash {
			log.Error(
				"L2 block hash mismatch, re-sync from genesis",
				"height", l2Head.Number,
				"hash in protocol", common.Hash(l2HeadHash),
				"hash in execution engine", l2Head.Hash(),
			)

			heightOrID.ID = common.Big0
			heightOrID.Height = common.Big0
			if l2HeadHash, err = s.rpc.TaikoL1.GetSyncedHeader(nil, common.Big0); err != nil {
				return err
			}
		}

		// If the L2 execution engine has synced to latest verified block.
		if l2HeadHash == s.syncProgressTracker.LastSyncedVerifiedBlockHash() {
			heightOrID.ID = s.syncProgressTracker.LastSyncedVerifiedBlockID()
		}

		// Reset the L1Current cursor.
		blockID, err := s.state.resetL1Current(s.ctx, heightOrID)
		if err != nil {
			return err
		}

		// Reset to the latest L2 execution engine's chain status.
		s.syncProgressTracker.UpdateMeta(blockID, heightOrID.Height, l2HeadHash)
	}

	// Insert the proposed block one by one.
	return s.ProcessL1Blocks(s.ctx, l1End)
}

// AheadOfProtocolVerifiedHead checks whether the L2 chain is ahead of verified head in protocol.
func (s *L2ChainSyncer) AheadOfProtocolVerifiedHead() bool {
	verifiedHeightToCompare := s.state.getLatestVerifiedBlock().Height.Uint64()
	log.Info(
		"Checking whether the execution engine is head of protocol's verfiied head",
		"latestVerifiedBlock", verifiedHeightToCompare,
		"executionEngineHead", s.state.GetL2Head(),
	)
	if verifiedHeightToCompare > 0 {
		// If latest verified head height is equal to L2 execution engine's synced head height minus one,
		// we also mark the triggered P2P sync progress as finished to prevent a protenial `InsertBlockWithoutSetHead` in
		// execution engine, which may cause errors since we do not pass all transactions in ExecutePayload when calling
		// NewPayloadV1.
		verifiedHeightToCompare -= 1
	}

	if s.state.GetL2Head().Number.Uint64() < verifiedHeightToCompare {
		return false
	}

	if s.syncProgressTracker.LastSyncedVerifiedBlockHeight() != nil {
		return s.state.GetL2Head().Number.Uint64() >= s.syncProgressTracker.LastSyncedVerifiedBlockHeight().Uint64()
	}

	return true
}

// ProcessL1Blocks fetches all `TaikoL1.BlockProposed` events between given
// L1 block heights, and then tries inserting them into L2 execution engine's block chain.
func (s *L2ChainSyncer) ProcessL1Blocks(ctx context.Context, l1End *types.Header) error {
	iter, err := eventIterator.NewBlockProposedIterator(ctx, &eventIterator.BlockProposedIteratorConfig{
		Client:               s.rpc.L1,
		TaikoL1:              s.rpc.TaikoL1,
		StartHeight:          s.state.l1Current.Number,
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

	s.state.l1Current = l1End
	metrics.DriverL1CurrentHeightGauge.Update(s.state.l1Current.Number.Int64())

	return nil
}
