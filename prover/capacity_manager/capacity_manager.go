package capacity_manager

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// CapacityManager manages the prover capacity concurrent-safely.
type CapacityManager struct {
	capacity              map[uint64]bool
	tempCapacity          []time.Time
	tempCapacityExpiresAt time.Duration
	maxCapacity           uint64
	mutex                 sync.RWMutex
}

// New creates a new CapacityManager instance.
func New(capacity uint64, tempCapacityExpiresAt time.Duration) *CapacityManager {
	return &CapacityManager{
		capacity:              make(map[uint64]bool),
		maxCapacity:           capacity,
		tempCapacity:          make([]time.Time, 0),
		tempCapacityExpiresAt: tempCapacityExpiresAt,
	}
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
		log.Info("Can not release capacity",
			"blockID", blockID,
			"currentCapacity", len(m.capacity),
			"maxCapacity", m.maxCapacity)
		return uint64(len(m.capacity)), false
	}

	delete(m.capacity, blockID)

	log.Info("Released capacity", "blockID", blockID, "capacityAfterRelease", len(m.capacity))

	return uint64(len(m.capacity)), true
}

// TakeOneCapacity takes one capacity.
func (m *CapacityManager) TakeOneCapacity(blockID uint64) (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if len(m.capacity) == int(m.maxCapacity) {
		log.Info("Could not take one capacity", "blockID", blockID, "capacity", len(m.capacity))
		return 0, false
	}

	m.capacity[blockID] = true

	log.Info("Took one capacity", "blockID", blockID, "capacityAfterTaking", len(m.capacity))

	return uint64(len(m.capacity)), true
}

func (m *CapacityManager) TakeOneTempCapacity() (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// clear expired tempCapacities

	m.clearExpiredTempCapacities()

	if len(m.capacity)+len(m.tempCapacity) >= int(m.maxCapacity) {
		log.Info("Could not take one temp capacity", "capacity", len(m.capacity), "tempCapacity", len(m.tempCapacity))
		return 0, false
	}

	m.tempCapacity = append(m.tempCapacity, time.Now().UTC())

	return uint64(len(m.tempCapacity)), true
}

func (m *CapacityManager) clearExpiredTempCapacities() {
	for i, c := range m.tempCapacity {
		if time.Now().UTC().Sub(c) > m.tempCapacityExpiresAt {
			m.tempCapacity = append(m.tempCapacity[:i], m.tempCapacity[i+1:]...)
		}
	}
}
