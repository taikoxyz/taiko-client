package capacity_manager

import (
	"sync"

	"github.com/ethereum/go-ethereum/log"
)

// CapacityManager manages the prover capacity concurrent-safely.
type CapacityManager struct {
	capacity    map[uint64]bool
	maxCapacity uint64
	mutex       sync.RWMutex
}

// New creates a new CapacityManager instance.
func New(capacity uint64) *CapacityManager {
	return &CapacityManager{capacity: make(map[uint64]bool), maxCapacity: capacity}
}

// ReadCapacity reads the current capacity.
func (m *CapacityManager) ReadCapacity() uint64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	log.Info("Reading capacity", "capacity", len(m.capacity))

	return uint64(len(m.capacity))
}

// ReleaseOneCapacity releases one capacity.
func (m *CapacityManager) ReleaseOneCapacity(blockID uint64) (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.capacity[blockID]; !ok {
		log.Info("Can not release capacity", "blockID", blockID, "currentCapacity", m.capacity, "maxCapacity", m.maxCapacity)
		return uint64(len(m.capacity)), false
	}

	delete(m.capacity, blockID)

	log.Info("Released capacity", "capacityAfterRelease", len(m.capacity))

	return uint64(len(m.capacity)), true
}

// TakeOneCapacity takes one capacity.
func (m *CapacityManager) TakeOneCapacity(blockID uint64) (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if len(m.capacity) == int(m.maxCapacity) {
		log.Info("Could not take one capacity", "capacity", len(m.capacity))
		return 0, false
	}

	m.capacity[blockID] = true

	log.Info("Took one capacity", "blockID", blockID, "capacityAfterTaking", m.capacity)

	return uint64(len(m.capacity)), true
}
