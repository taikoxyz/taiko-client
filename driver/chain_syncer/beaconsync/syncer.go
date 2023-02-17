package beaconsync

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/driver/state"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// Syncer responsible for letting the L2 execution engine catching up with protocol's latest
// verified block through P2P beacon sync.
type Syncer struct {
	ctx             context.Context
	rpc             *rpc.Client
	state           *state.State
	progressTracker *SyncProgressTracker // Sync progress tracker
}

// NewSyncer creates a new syncer instance.
func NewSyncer(
	ctx context.Context,
	rpc *rpc.Client,
	state *state.State,
	progressTracker *SyncProgressTracker,
) *Syncer {
	return &Syncer{ctx, rpc, state, progressTracker}
}

// TriggerBeaconSync triggers the L2 execution engine to start performing a beacon sync.
func (s *Syncer) TriggerBeaconSync() error {
	blockID, latestVerifiedHeadPayload, err := s.getVerifiedBlockPayload(s.ctx)
	if err != nil {
		return err
	}

	if !s.progressTracker.HeadChanged(blockID) {
		log.Debug("Verified head has not changed", "blockID", blockID, "hash", latestVerifiedHeadPayload.BlockHash)
		return nil
	}

	status, err := s.rpc.L2Engine.NewPayload(
		s.ctx,
		latestVerifiedHeadPayload,
	)
	if err != nil {
		return err
	}

	if status.Status != beacon.SYNCING && status.Status != beacon.VALID {
		return fmt.Errorf("unexpected NewPayload response status: %s", status.Status)
	}

	fcRes, err := s.rpc.L2Engine.ForkchoiceUpdate(s.ctx, &beacon.ForkchoiceStateV1{
		HeadBlockHash:      latestVerifiedHeadPayload.BlockHash,
		SafeBlockHash:      latestVerifiedHeadPayload.BlockHash,
		FinalizedBlockHash: latestVerifiedHeadPayload.BlockHash,
	}, nil)
	if err != nil {
		return err
	}
	if fcRes.PayloadStatus.Status != beacon.SYNCING {
		return fmt.Errorf("unexpected ForkchoiceUpdate response status: %s", status.Status)
	}

	// Update sync status.
	s.progressTracker.UpdateMeta(
		blockID,
		new(big.Int).SetUint64(latestVerifiedHeadPayload.Number),
		latestVerifiedHeadPayload.BlockHash,
	)

	log.Info(
		"⛓️ Beacon sync triggered",
		"newHeadID", blockID,
		"newHeadHeight", s.progressTracker.LastSyncedVerifiedBlockHeight(),
		"newHeadHash", s.progressTracker.LastSyncedVerifiedBlockHash(),
	)

	return nil
}

// getVerifiedBlockPayload fetches the latest verified block's header, and converts it to an Engine API executable data,
// which will be used to let the node to start beacon syncing.
func (s *Syncer) getVerifiedBlockPayload(ctx context.Context) (*big.Int, *beacon.ExecutableDataV1, error) {
	var (
		proveBlockTxHash    common.Hash
		latestVerifiedBlock = s.state.GetLatestVerifiedBlock()
	)

	// Get the latest verified block's corresponding BlockProven event.
	iter, err := eventIterator.NewBlockProvenIterator(s.ctx, &eventIterator.BlockProvenIteratorConfig{
		Client:      s.rpc.L1,
		TaikoL1:     s.rpc.TaikoL1,
		StartHeight: s.state.GenesisL1Height,
		EndHeight:   s.state.GetL1Head().Number,
		FilterQuery: []*big.Int{latestVerifiedBlock.ID},
		Reverse:     true,
		OnBlockProvenEvent: func(
			ctx context.Context,
			e *bindings.TaikoL1ClientBlockProven,
			endIter eventIterator.EndBlockProvenEventIterFunc,
		) error {
			if bytes.Equal(e.BlockHash[:], latestVerifiedBlock.Hash.Bytes()) {
				log.Info(
					"Latest verified block's BlockProven event found",
					"height", e.Raw.BlockNumber,
					"txHash", e.Raw.TxHash,
				)
				proveBlockTxHash = e.Raw.TxHash
				endIter()
			}
			return nil
		},
	})

	if err != nil {
		return nil, nil, err
	}

	if err := iter.Iter(); err != nil {
		return nil, nil, err
	}

	if proveBlockTxHash == (common.Hash{}) {
		return nil, nil, fmt.Errorf(
			"failed to find L1 height of latest verified block's ProveBlock transaction, id: %s",
			latestVerifiedBlock.ID,
		)
	}

	// Get the latest verified block's header, then convert it to ExecutableDataV1.
	proveBlockTx, _, err := s.rpc.L1.TransactionByHash(s.ctx, proveBlockTxHash)
	if err != nil {
		return nil, nil, err
	}

	evidenceHeader, err := encoding.UnpackEvidenceHeader(proveBlockTx.Data())
	if err != nil {
		return nil, nil, err
	}

	header := encoding.ToGethHeader(evidenceHeader)

	if header.Hash() != latestVerifiedBlock.Hash {
		return nil, nil, fmt.Errorf(
			"latest verified block hash mismatch: %s != %s", header.Hash(), latestVerifiedBlock.Hash,
		)
	}

	log.Info("Latest verified block header retrieved", "hash", header.Hash())

	return latestVerifiedBlock.ID, encoding.ToExecutableDataV1(header), nil
}
