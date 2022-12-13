package driver

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var (
	syncProgressFetchInterval = 10 * time.Second
)

// BeaconSyncProgressTracker is responsible for tracking the L2 execution engine's sync progress, after
// a beacon sync is triggered in it, and check whether the L2 execution is out of sync (due to no connected peer
// or some other reasons).
type BeaconSyncProgressTracker struct {
	// RPC client
	client *ethclient.Client

	// Meta data
	triggered                     bool
	lastSyncedVerifiedBlockID     *big.Int
	lastSyncedVerifiedBlockHeight *big.Int
	lastSyncedVerifiedBlockHash   common.Hash

	// Out-of-sync check related
	lastSyncProgress   *ethereum.SyncProgress
	lastProgressedTime time.Time
	timeout            time.Duration
	outOfSync          bool
	ticker             *time.Ticker

	// Read-write mutex
	mutex sync.RWMutex
}

// NewBeaconSyncProgressTracker creates a new BeaconSyncProgressTracker instance.
func NewBeaconSyncProgressTracker(c *ethclient.Client, timeout time.Duration) *BeaconSyncProgressTracker {
	return &BeaconSyncProgressTracker{client: c, timeout: timeout, ticker: time.NewTicker(syncProgressFetchInterval)}
}

// Track starts the inner event loop, to monitor the sync progress.
func (s *BeaconSyncProgressTracker) Track(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-s.ticker.C:
			s.track(ctx)
		}
	}
}

// track is the internal implementation of MonitorSyncProgress, tries to
// track the L2 execution engine's beacon sync progress.
func (s *BeaconSyncProgressTracker) track(ctx context.Context) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.triggered {
		log.Debug("Beacon sync not triggered")
		return
	}

	if s.outOfSync {
		return
	}

	progress, err := s.client.SyncProgress(ctx)
	if err != nil {
		log.Error("Get L2 execution engine sync progress error", "error", err)
		return
	}

	log.Info(
		"L2 execution engine sync progress",
		"progress", progress,
		"lastProgressedTime", s.lastProgressedTime,
		"timeout", s.timeout,
	)

	if progress == nil {
		headHeight, err := s.client.BlockNumber(ctx)
		if err != nil {
			log.Error("Get L2 execution engine head height error", "error", err)
			return
		}

		if new(big.Int).SetUint64(headHeight).Cmp(s.lastSyncedVerifiedBlockHeight) >= 0 {
			s.lastProgressedTime = time.Now()
			log.Info("L2 execution engine has finished the P2P sync work, all verfiied blocks synced, "+
				"will switch to insert pending blocks ony be one",
				"lastSyncedVerifiedBlockID", s.lastSyncedVerifiedBlockID,
				"lastSyncedVerifiedBlockHeight", s.lastSyncedVerifiedBlockHeight,
				"lastSyncedVerifiedBlockHash", s.lastSyncedVerifiedBlockHash,
			)
			return
		}

		log.Warn("L2 execution engine has not started P2P syncing yet")
	}

	defer func() { s.lastSyncProgress = progress }()

	// Check whether the L2 execution engine has synced any new block through P2P since last event loop.
	if syncProgressed(s.lastSyncProgress, progress) {
		s.outOfSync = false
		s.lastProgressedTime = time.Now()
		return
	}

	// Has not synced any new block since last loop, check whether reaching the timeout.
	if time.Since(s.lastProgressedTime) > s.timeout {
		// Mark the L2 execution engine out of sync.
		s.outOfSync = true

		log.Warn(
			"L2 execution engine out of sync",
			"lastProgressedTime", s.lastProgressedTime,
			"timeout", s.timeout,
		)
	}
}

// UpdateMeta updates the inner beacon sync meta data.
func (s *BeaconSyncProgressTracker) UpdateMeta(id, height *big.Int, blockHash common.Hash) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Debug("Update sync progress tracker meta", "id", id, "height", height, "hash", blockHash)

	if !s.triggered {
		s.lastProgressedTime = time.Now()
	}

	s.triggered = true
	s.lastSyncedVerifiedBlockID = id
	s.lastSyncedVerifiedBlockHeight = height
	s.lastSyncedVerifiedBlockHash = blockHash
}

// ClearMeta cleans the inner beacon sync meta data.
func (s *BeaconSyncProgressTracker) ClearMeta() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Debug("Clear sync progress tracker meta")

	s.triggered = false
	s.lastSyncedVerifiedBlockID = nil
	s.lastSyncedVerifiedBlockHash = common.Hash{}
	s.outOfSync = false
}

// HeadChanged checks if a new beacon sync request will be needed.
func (s *BeaconSyncProgressTracker) HeadChanged(newID *big.Int) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.triggered {
		return true
	}

	return s.lastSyncedVerifiedBlockID != nil && s.lastSyncedVerifiedBlockID != newID
}

// OutOfSync tells whether the L2 execution engine is marked as out of sync.
func (s *BeaconSyncProgressTracker) OutOfSync() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.outOfSync
}

// Triggered returns tracker.triggered.
func (s *BeaconSyncProgressTracker) Triggered() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.triggered
}

// LastSyncedVerifiedBlockID returns tracker.lastSyncedVerifiedBlockID.
func (s *BeaconSyncProgressTracker) LastSyncedVerifiedBlockID() *big.Int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return new(big.Int).Set(s.lastSyncedVerifiedBlockID)
}

// LastSyncedVerifiedBlockHeight returns tracker.lastSyncedVerifiedBlockHeight.
func (s *BeaconSyncProgressTracker) LastSyncedVerifiedBlockHeight() *big.Int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return new(big.Int).Set(s.lastSyncedVerifiedBlockHeight)
}

// LastSyncedVerifiedBlockHash returns tracker.lastSyncedVerifiedBlockHash.
func (s *BeaconSyncProgressTracker) LastSyncedVerifiedBlockHash() common.Hash {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.lastSyncedVerifiedBlockHash
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
