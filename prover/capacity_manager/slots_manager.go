package capacity_manager

import (
	"sort"
	"time"
)

const (
	BlockIDPlaceHolder = 0
)

type capacitySlot struct {
	blockID   uint64
	expiredAt *time.Time
}

type slotsManager struct {
	slots    []*capacitySlot
	maxSlots uint64
}

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

func (s *slotsManager) TakeOneSlot(blockID uint64) bool {
	defer s.sort()

	if ok := s.removeItemByBlockID(BlockIDPlaceHolder); !ok {
		return false
	}

	s.slots = append(s.slots, &capacitySlot{blockID: blockID})
	return true
}

func (s *slotsManager) Len() uint64 {
	return uint64(len(s.slots))
}

func (s *slotsManager) MaxSlots() uint64 {
	return s.maxSlots
}
