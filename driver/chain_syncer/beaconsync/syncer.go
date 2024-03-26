package beaconsync

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/driver/state"
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

// TriggerBeaconSync triggers the L2 execution engine to start performing a beacon sync, if the
// latest verified block has changed.
func (s *Syncer) TriggerBeaconSync(number uint64) error {
	blockID := new(big.Int).SetUint64(number)
	header, err := s.rpc.L2CheckPoint.HeaderByNumber(s.ctx, blockID)
	if err != nil {
		return err
	}

	beaconHeadPayload := encoding.ToExecutableData(header)

	if !s.progressTracker.HeadChanged(blockID) {
		log.Debug("beacon sync head has not changed", "blockID", blockID, "hash", beaconHeadPayload.BlockHash)
		return nil
	}

	if s.progressTracker.Triggered() {
		if s.progressTracker.lastSyncProgress == nil {
			log.Info(
				"Syncing beacon headers, please check L2 execution engine logs for progress",
				"currentSyncHead", s.progressTracker.LastSyncedBlockID(),
				"newBlockID", blockID,
			)
		}
	}

	status, err := s.rpc.L2Engine.NewPayload(s.ctx, beaconHeadPayload)
	if err != nil {
		return err
	}

	if status.Status != engine.SYNCING && status.Status != engine.VALID {
		return fmt.Errorf("unexpected NewPayload response status: %s", status.Status)
	}

	fcRes, err := s.rpc.L2Engine.ForkchoiceUpdate(s.ctx, &engine.ForkchoiceStateV1{
		HeadBlockHash: beaconHeadPayload.BlockHash,
	}, nil)
	if err != nil {
		return err
	}
	if fcRes.PayloadStatus.Status != engine.SYNCING {
		return fmt.Errorf("unexpected ForkchoiceUpdate response status: %s", fcRes.PayloadStatus.Status)
	}

	// Update sync status.
	s.progressTracker.UpdateMeta(blockID, beaconHeadPayload.BlockHash)

	log.Info(
		"⛓️ Beacon sync triggered",
		"newHeadID", blockID,
		"newHeadHash", s.progressTracker.LastSyncedBlockHash(),
	)

	return nil
}
