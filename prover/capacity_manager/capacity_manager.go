package capacity_manager

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// CapacityManager manages the prover capacity concurrent-safely.
type CapacityManager struct {
	slotsManager *slotsManager
	mutex        sync.RWMutex
}

// New creates a new CapacityManager instance.
func New(capacity uint64) *CapacityManager {
	return &CapacityManager{
		slotsManager: &slotsManager{[]*capacitySlot{}, capacity},
	}
}

// ReadCapacity reads the current capacity.
func (m *CapacityManager) ReadCapacity() uint64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	log.Info(
		"Reading capacity",
		"maxCapacity", m.slotsManager.MaxSlots(),
		"currentCapacity", m.slotsManager.MaxSlots()-m.slotsManager.Len(),
		"currentUsage", m.slotsManager.Len(),
	)

	return m.slotsManager.MaxSlots() - m.slotsManager.Len()
}

// ReleaseOneCapacity releases one capacity.
func (m *CapacityManager) ReleaseOneCapacity(blockID uint64) (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if blockID == BlockIDPlaceHolder {
		return m.slotsManager.MaxSlots() - m.slotsManager.Len(), false
	}

	if ok := m.slotsManager.removeItemByBlockID(blockID); !ok {
		log.Info(
			"Can not release capacity",
			"blockID", blockID,
			"maxCapacity", m.slotsManager.MaxSlots(),
			"currentCapacity", m.slotsManager.MaxSlots()-m.slotsManager.Len(),
			"currentUsage", m.slotsManager.Len(),
		)
		return m.slotsManager.MaxSlots() - m.slotsManager.Len(), false
	}

	log.Info(
		"Released capacity",
		"blockID", blockID,
		"maxCapacity", m.slotsManager.MaxSlots(),
		"currentCapacity", m.slotsManager.MaxSlots()-m.slotsManager.Len(),
		"currentUsage", m.slotsManager.Len(),
	)

	return m.slotsManager.MaxSlots() - m.slotsManager.Len(), true
}

func (m *CapacityManager) HoldOneCapacity(expiry time.Duration) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.slotsManager.HoldOneSlot(expiry)
}

// TakeOneCapacity takes one capacity.
func (m *CapacityManager) TakeOneCapacity(blockID uint64) (uint64, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if ok := m.slotsManager.TakeOneSlot(blockID); !ok {
		log.Info(
			"Could not take one capacity",
			"blockID", blockID,
			"maxCapacity", m.slotsManager.MaxSlots(),
			"currentCapacity", m.slotsManager.MaxSlots()-m.slotsManager.Len(),
			"currentUsage", m.slotsManager.Len(),
		)

		return m.slotsManager.MaxSlots() - m.slotsManager.Len(), true
	}

	log.Info(
		"Took one capacity",
		"blockID", blockID,
		"maxCapacity", m.slotsManager.MaxSlots(),
		"currentCapacity", m.slotsManager.MaxSlots()-m.slotsManager.Len(),
		"currentUsage", m.slotsManager.Len(),
	)

	return m.slotsManager.MaxSlots() - m.slotsManager.Len(), true
}
