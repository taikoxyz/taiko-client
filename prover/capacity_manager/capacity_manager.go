package capacity_manager

import (
	"sync"
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

	return m.capacity
}

// ReleaseCapacity releases one capacity.
func (m *CapacityManager) ReleaseOneCapacity() uint64 {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.capacity += 1
	return m.capacity
}

// TakeOneCapacity takes one capacit·y.
func (m *CapacityManager) TakeOneCapacity() (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.capacity == 0 {
		return 0, false
	}

	m.capacity -= 1
	return m.capacity, true
}
