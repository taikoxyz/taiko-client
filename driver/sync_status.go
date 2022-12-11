package driver

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type BeaconSyncStatus struct {
	triggered                      bool
	lastSyncedVerifiedBlockID      *big.Int
	lastSyncedVerifiedBlockHash    common.Hash
	lastExecutionEngineHeadChanged time.Time
	mutex                          sync.RWMutex
}

func (s *BeaconSyncStatus) Init() error {
	// monitor sync progress
	return nil
}

func (s *BeaconSyncStatus) Refresh(id *big.Int, blockHash common.Hash) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.triggered = true
	s.lastSyncedVerifiedBlockID = id
	s.lastSyncedVerifiedBlockHash = blockHash
	s.lastExecutionEngineHeadChanged = time.Now()
}

func (s *BeaconSyncStatus) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.triggered = false
	s.lastSyncedVerifiedBlockID = nil
	s.lastSyncedVerifiedBlockHash = common.Hash{}
}

func (s *BeaconSyncStatus) HeadChanged(newID *big.Int) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.triggered {
		return false
	}

	return s.lastSyncedVerifiedBlockID != nil && s.lastSyncedVerifiedBlockID != newID
}
