package capacity_manager

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SlotsManagerTestSuite struct {
	suite.Suite
	m *slotsManager
}

func (s *SlotsManagerTestSuite) SetupTest() {
	now := time.Now()
	oneHourLater := time.Now().Add(1 * time.Hour)

	s.m = &slotsManager{[]*capacitySlot{
		{blockID: BlockIDPlaceHolder, expiredAt: &now},
		{blockID: BlockIDPlaceHolder, expiredAt: &oneHourLater},
		{blockID: 3, expiredAt: nil},
	}, testCapacity}
}

func (s *SlotsManagerTestSuite) TestSort() {
	now := time.Now()
	oneHourLater := time.Now().Add(1 * time.Hour)

	s.m.slots = []*capacitySlot{
		{blockID: 1, expiredAt: nil},
		{blockID: 2, expiredAt: &oneHourLater},
		{blockID: 3, expiredAt: &now},
	}

	s.m.sort()
	s.Equal(3, int(s.m.slots[0].blockID))
	s.Equal(2, int(s.m.slots[1].blockID))
	s.Equal(1, int(s.m.slots[2].blockID))

	s.m.slots = []*capacitySlot{
		{blockID: 1, expiredAt: &now},
		{blockID: 2, expiredAt: &oneHourLater},
		{blockID: 3, expiredAt: nil},
	}

	s.m.sort()
	s.Equal(1, int(s.m.slots[0].blockID))
	s.Equal(2, int(s.m.slots[1].blockID))
	s.Equal(3, int(s.m.slots[2].blockID))
}

func (s *SlotsManagerTestSuite) TestRemoveItemByBlockID() {
	s.True(s.m.removeItemByBlockID(3))
	s.Equal(2, len(s.m.slots))
	s.Equal(BlockIDPlaceHolder, int(s.m.slots[0].blockID))
	s.Equal(BlockIDPlaceHolder, int(s.m.slots[1].blockID))
}

func (s *SlotsManagerTestSuite) TestRemoveItemByBlockIDWithPlaceHolder() {
	s.True(s.m.removeItemByBlockID(BlockIDPlaceHolder))
	s.Equal(2, len(s.m.slots))
	s.Equal(BlockIDPlaceHolder, int(s.m.slots[0].blockID))
	s.Equal(3, int(s.m.slots[1].blockID))
}

func (s *SlotsManagerTestSuite) TestClearOneExpiredSlots() {
	oneHourBefore := time.Now().Add(-1 * time.Hour)

	s.m.slots = []*capacitySlot{
		{blockID: 1, expiredAt: &oneHourBefore},
		{blockID: 2, expiredAt: nil},
		{blockID: 3, expiredAt: nil},
	}

	s.m.clearOneExpiredSlots()

	s.Equal(2, len(s.m.slots))
	s.Equal(3, int(s.m.slots[0].blockID))
	s.Equal(2, int(s.m.slots[1].blockID))
}

func (s *SlotsManagerTestSuite) TestHoldOneSlot() {
	oneHourLater := time.Now().Add(1 * time.Hour)
	twoHoursLater := time.Now().Add(2 * time.Hour)

	s.m = &slotsManager{[]*capacitySlot{
		{blockID: BlockIDPlaceHolder, expiredAt: &oneHourLater},
		{blockID: BlockIDPlaceHolder, expiredAt: &twoHoursLater},
		{blockID: 3, expiredAt: nil},
	}, testCapacity}

	s.m.HoldOneSlot(30 * time.Minute)

	s.Equal(uint64(4), s.m.Len())
}

func (s *SlotsManagerTestSuite) TestTakeOneSlot() {
	oneHourLater := time.Now().Add(1 * time.Hour)
	twoHoursLater := time.Now().Add(2 * time.Hour)

	s.m = &slotsManager{[]*capacitySlot{
		{blockID: BlockIDPlaceHolder, expiredAt: &oneHourLater},
		{blockID: BlockIDPlaceHolder, expiredAt: &twoHoursLater},
		{blockID: 3, expiredAt: nil},
	}, testCapacity}

	s.m.TakeOneSlot(4)

	s.Equal(uint64(3), s.m.Len())

	for _, slot := range s.m.slots {
		if slot.blockID == 4 {
			return
		}
	}

	s.FailNow("slot has not been taken correctly")
}

func TestSlotsManagerTestSuite(t *testing.T) {
	suite.Run(t, new(SlotsManagerTestSuite))
}
