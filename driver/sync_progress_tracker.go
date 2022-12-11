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
	triggered                   bool
	lastSyncedVerifiedBlockID   *big.Int
	lastSyncedVerifiedBlockHash common.Hash

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
	return &BeaconSyncProgressTracker{client: c, ticker: time.NewTicker(syncProgressFetchInterval)}
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

	progress, err := s.client.SyncProgress(ctx)
	if err != nil {
		log.Error("Get L2 execution engine sync progress error", "error", err)
		return
	}

	if progress == nil {
		log.Info("L2 execution engine has finished the P2P sync work")
		return
	}

	log.Info("L2 execution engine sync progress", "progress", progress)

	defer func() { s.lastSyncProgress = progress }()

	// Check whether the L2 execution engine has synced any new block through P2P since last event loop.
	if Progressed(s.lastSyncProgress, progress) {
		s.outOfSync = false
		s.lastProgressedTime = time.Now()
		return
	}

	// Has not synced any new block since last loop, check whether reaching the timeout.
	if time.Since(s.lastProgressedTime) > s.timeout {
		// Mark the L2 execution engine out of sync.
		s.outOfSync = true
	}
}

// UpdateMeta updates the inner beacon sync meta data.
func (s *BeaconSyncProgressTracker) UpdateMeta(id *big.Int, blockHash common.Hash) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.triggered = true
	s.lastSyncedVerifiedBlockID = id
	s.lastSyncedVerifiedBlockHash = blockHash
}

// ClearMeta cleans the inner beacon sync meta data.
func (s *BeaconSyncProgressTracker) ClearMeta() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.triggered = false
	s.lastSyncedVerifiedBlockID = nil
	s.lastSyncedVerifiedBlockHash = common.Hash{}
}

// HeadChanged checks if a new beacon sync request will be needed.
func (s *BeaconSyncProgressTracker) HeadChanged(newID *big.Int) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.triggered {
		return false
	}

	return s.lastSyncedVerifiedBlockID != nil && s.lastSyncedVerifiedBlockID != newID
}

// OutOfSync tells whether the L2 execution engine is marked as out of sync.
func (s *BeaconSyncProgressTracker) OutOfSync() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.outOfSync
}

// Progressed checks whether there is any new progress since last sync progress check.
func Progressed(last *ethereum.SyncProgress, new *ethereum.SyncProgress) bool {
	if last == nil {
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
