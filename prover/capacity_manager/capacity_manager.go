package capacity_manager

import (
	"sync"

	"github.com/ethereum/go-ethereum/log"
)

// CapacityManager manages the prover capacity concurrent-safely.
type CapacityManager struct {
	capacity    uint64
	maxCapacity uint64
	mutex       sync.RWMutex
}

// New creates a new CapacityManager instance.
func New(capacity uint64) *CapacityManager {
	return &CapacityManager{capacity: capacity, maxCapacity: capacity}
}

// ReadCapacity reads the current capacity.
func (m *CapacityManager) ReadCapacity() uint64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	log.Info("Reading capacity", "capacity", m.capacity)

	return m.capacity
}

// ReleaseOneCapacity releases one capacity.
func (m *CapacityManager) ReleaseOneCapacity() (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.capacity+1 > m.maxCapacity {
		log.Info("Can not release capacity", "currentCapacity", m.capacity, "maxCapacity", m.maxCapacity)
		return m.capacity, false
	}

	m.capacity += 1

	log.Info("Released capacity", "capacityAfterRelease", m.capacity)

	return m.capacity, true
}

// TakeOneCapacity takes one capacity.
func (m *CapacityManager) TakeOneCapacity() (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.capacity == 0 {
		log.Info("Could not take one capacity", "capacity", m.capacity)
		return 0, false
	}

	m.capacity -= 1

	log.Info("Took one capacity", "capacityAfterTaking", m.capacity)

	return m.capacity, true
}
