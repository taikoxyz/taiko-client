package capacity_manager

import (
	"sort"
	"time"
)

// BlockIDPlaceHolder is a special blockID which represents a temporarily holen slot with no blockID.
const BlockIDPlaceHolder = 0

// capacitySlot represents a block slot with an expired time.
type capacitySlot struct {
	blockID   uint64
	expiredAt *time.Time
}

// slotsManager manages all the block slots with a max capacity.
type slotsManager struct {
	slots    []*capacitySlot
	maxSlots uint64
}

// sort sorts the slots by expired time, if a slot has no expired time, it will be put at the end.
func (s *slotsManager) sort() {
	sort.Slice(s.slots, func(i, j int) bool {
		if s.slots[i].expiredAt == nil && s.slots[j].expiredAt == nil {
			return false
		}
		if s.slots[i].expiredAt == nil && s.slots[j].expiredAt != nil {
			return false
		}
		if s.slots[i].expiredAt != nil && s.slots[j].expiredAt == nil {
			return true
		}

		return s.slots[i].expiredAt.Before(*s.slots[j].expiredAt)
	})
}

// removeItemByBlockID removes a slot by blockID, if the blockID is BlockIDPlaceHolder,
// it will remove the first slot with expired time, otherwise it will remove the slot with the blockID.
func (s *slotsManager) removeItemByBlockID(id uint64) bool {
	defer s.sort()

	if len(s.slots) == 0 {
		return false
	}

	if id == BlockIDPlaceHolder {
		if s.slots[0].expiredAt != nil {
			s.slots[0] = s.slots[len(s.slots)-1]
			s.slots = s.slots[:len(s.slots)-1]
			return true
		}

		return false
	}

	for i := range s.slots {
		if s.slots[i].blockID == id {
			s.slots[i] = s.slots[len(s.slots)-1]
			s.slots = s.slots[:len(s.slots)-1]
			return true
		}
	}

	return false
}

// clearOneExpiredSlots tries to remove one expired slot.
func (s *slotsManager) clearOneExpiredSlots() {
	defer s.sort()

	for i := range s.slots {
		if s.slots[i].expiredAt != nil && s.slots[i].expiredAt.Before(time.Now()) {
			s.slots[i] = s.slots[len(s.slots)-1]
			s.slots = s.slots[:len(s.slots)-1]
			return
		}
	}
}

// HoldOneSlot holds one slot with an expired time.
func (s *slotsManager) HoldOneSlot(expiry time.Duration) bool {
	defer s.sort()

	s.clearOneExpiredSlots()

	if len(s.slots) >= int(s.maxSlots) {
		return false
	}

	expiredAt := time.Now().Add(expiry)
	s.slots = append(s.slots, &capacitySlot{blockID: BlockIDPlaceHolder, expiredAt: &expiredAt})
	return true
}

// TakeOneSlot tries to taken one holden slot (blockID == BlockIDPlaceHolder), if there is no holden slot,
// it will return false.
func (s *slotsManager) TakeOneSlot(blockID uint64) bool {
	defer s.sort()

	if ok := s.removeItemByBlockID(BlockIDPlaceHolder); !ok {
		return false
	}

	s.slots = append(s.slots, &capacitySlot{blockID: blockID})
	return true
}

// Len returns the current usage of the slots.
func (s *slotsManager) Len() uint64 {
	return uint64(len(s.slots))
}

// MaxSlots returns the max capacity of the slots.
func (s *slotsManager) MaxSlots() uint64 {
	return s.maxSlots
}
