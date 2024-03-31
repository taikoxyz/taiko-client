package beaconsync

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

var (
	syncProgressCheckInterval = 12 * time.Second
)

// SyncProgressTracker is responsible for tracking the L2 execution engine's sync progress, after
// a beacon sync is triggered, and check whether the L2 execution is not able to sync through P2P (due to no
// connected peer or some other reasons).
type SyncProgressTracker struct {
	// RPC client
	client *rpc.EthClient

	// Meta data
	triggered                   atomic.Bool
	lastSyncedVerifiedBlockID   atomic.Uint64
	lastSyncedVerifiedBlockHash atomic.Value

	// Out-of-sync check related
	lastSyncProgress   *ethereum.SyncProgress
	lastProgressedTime time.Time
	timeout            time.Duration
	outOfSync          atomic.Bool
	ticker             *time.Ticker
}

// NewSyncProgressTracker creates a new SyncProgressTracker instance.
func NewSyncProgressTracker(c *rpc.EthClient, timeout time.Duration) *SyncProgressTracker {
	return &SyncProgressTracker{client: c, timeout: timeout, ticker: time.NewTicker(syncProgressCheckInterval)}
}

// Track starts the inner event loop, to monitor the sync progress.
func (t *SyncProgressTracker) Track(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.ticker.C:
			t.track(ctx)
		}
	}
}

// track is the internal implementation of MonitorSyncProgress, tries to
// track the L2 execution engine's beacon sync progress.
func (t *SyncProgressTracker) track(ctx context.Context) {
	if !t.triggered.Load() {
		log.Debug("Beacon sync not triggered")
		return
	}

	if t.outOfSync.Load() {
		return
	}

	progress, err := t.client.SyncProgress(ctx)
	if err != nil {
		log.Error("Get L2 execution engine sync progress error", "error", err)
		return
	}

	if progress != nil {
		log.Info(
			"L2 execution engine sync progress",
			"progress", progress,
			"lastProgressedTime", t.lastProgressedTime,
			"timeout", t.timeout,
		)
	}

	if progress == nil {
		headHeight, err := t.client.BlockNumber(ctx)
		if err != nil {
			log.Error("Get L2 execution engine head height error", "error", err)
			return
		}

		if headHeight >= t.lastSyncedVerifiedBlockID.Load() {
			t.lastProgressedTime = time.Now()
			log.Info(
				"L2 execution engine has finished the P2P sync work, all verified blocks synced, "+
					"will switch to insert pending blocks one by one",
				"lastSyncedVerifiedBlockID", t.lastSyncedVerifiedBlockID.Load(),
				"lastSyncedVerifiedBlockHash", t.lastSyncedVerifiedBlockHash,
			)
			return
		}

		log.Info("L2 execution engine has not started P2P syncing yet", "timeout", t.timeout)
	}

	defer func() { t.lastSyncProgress = progress }()

	// Check whether the L2 execution engine has synced any new block through P2P since last event loop.
	if syncProgressed(t.lastSyncProgress, progress) {
		t.outOfSync.Store(false)
		t.lastProgressedTime = time.Now()
		return
	}

	// Has not synced any new block since last loop, check whether reaching the timeout.
	if time.Since(t.lastProgressedTime) > t.timeout {
		// Mark the L2 execution engine out of sync.
		t.outOfSync.Store(true)

		log.Warn(
			"L2 execution engine is not able to sync through P2P",
			"lastProgressedTime", t.lastProgressedTime,
			"timeout", t.timeout,
		)
	}
}

// UpdateMeta updates the inner beacon sync metadata.
func (t *SyncProgressTracker) UpdateMeta(id uint64, blockHash common.Hash) {
	log.Debug("Update sync progress tracker meta", "id", id, "hash", blockHash)

	if !t.triggered.Load() {
		t.lastProgressedTime = time.Now()
	}

	t.triggered.Store(true)
	t.lastSyncedVerifiedBlockID.Store(id)
	t.lastSyncedVerifiedBlockHash.Store(blockHash)
}

// ClearMeta cleans the inner beacon sync metadata.
func (t *SyncProgressTracker) ClearMeta() {
	log.Debug("Clear sync progress tracker meta")

	t.triggered.Store(false)
	t.lastSyncedVerifiedBlockID.Store(0)
	t.lastSyncedVerifiedBlockHash.Store(common.Hash{})
	t.outOfSync.Store(false)
}

// HeadChanged checks if a new beacon sync request will be needed.
func (t *SyncProgressTracker) HeadChanged(newID uint64) bool {
	if !t.triggered.Load() {
		return true
	}

	return t.lastSyncedVerifiedBlockID.Load() != newID
}

// OutOfSync tells whether the L2 execution engine is marked as out of sync.
func (t *SyncProgressTracker) OutOfSync() bool {
	return t.outOfSync.Load()
}

// Triggered returns tracker.triggered.
func (t *SyncProgressTracker) Triggered() bool {
	return t.triggered.Load()
}

// LastSyncedVerifiedBlockID returns tracker.lastSyncedVerifiedBlockID.
func (t *SyncProgressTracker) LastSyncedVerifiedBlockID() uint64 {
	return t.lastSyncedVerifiedBlockID.Load()
}

// LastSyncedVerifiedBlockHash returns tracker.lastSyncedVerifiedBlockHash.
func (t *SyncProgressTracker) LastSyncedVerifiedBlockHash() common.Hash {
	val := t.lastSyncedVerifiedBlockHash.Load()
	if val == nil {
		return common.Hash{}
	}
	return val.(common.Hash)
}

// syncProgressed checks whether there is any new progress since last sync progress check.
func syncProgressed(last *ethereum.SyncProgress, new *ethereum.SyncProgress) bool {
	if last == nil {
		return false
	}

	if new == nil {
		return true
	}

	// Block
	if new.CurrentBlock > last.CurrentBlock {
		return true
	}

	// Fast sync fields
	if new.PulledStates > last.PulledStates {
		return true
	}

	// Snap sync fields
	if new.SyncedAccounts > last.SyncedAccounts ||
		new.SyncedAccountBytes > last.SyncedAccountBytes ||
		new.SyncedBytecodes > last.SyncedBytecodes ||
		new.SyncedBytecodeBytes > last.SyncedBytecodeBytes ||
		new.SyncedStorage > last.SyncedStorage ||
		new.SyncedStorageBytes > last.SyncedStorageBytes ||
		new.HealedTrienodes > last.HealedTrienodes ||
		new.HealedTrienodeBytes > last.HealedTrienodeBytes ||
		new.HealedBytecodes > last.HealedBytecodes ||
		new.HealedBytecodeBytes > last.HealedBytecodeBytes ||
		new.HealingTrienodes > last.HealingTrienodes ||
		new.HealingBytecode > last.HealingBytecode {
		return true
	}

	return false
}
