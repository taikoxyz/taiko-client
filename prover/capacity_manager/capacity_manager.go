package capacity_manager

import (
	"sync"

	"github.com/ethereum/go-ethereum/log"
)

// CapacityManager manages the prover capacity concurrent-safely.
type CapacityManager struct {
	capacity uint64
	mutex    sync.RWMutex
}

// New creates a new CapacityManager instance.
func New(capacity uint64) *CapacityManager {
	return &CapacityManager{capacity: capacity}
}

// ReadCapacity reads the current capacity.
func (m *CapacityManager) ReadCapacity() uint64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	log.Info("reading capacity", "capacity", m.capacity)

	return m.capacity
}

// ReleaseCapacity releases one capacity.
func (m *CapacityManager) ReleaseOneCapacity() uint64 {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.capacity += 1

	log.Info("released capacity", "capacityAfterRelease", m.capacity)

	return m.capacity
}

// TakeOneCapacity takes one capacity.
func (m *CapacityManager) TakeOneCapacity() (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.capacity == 0 {
		log.Info("could not take one capacity", "capacity", m.capacity)

		return 0, false
	}

	m.capacity -= 1

	log.Info("took one capacity", "capacityAfterTaking", m.capacity)

	return m.capacity, true
}
