package capacity_manager

import (
	"time"

	"github.com/stretchr/testify/suite"
)

var (
	testCapacity          uint64        = 5
	tempCapacityExpiresAt time.Duration = 5 * time.Second
)

type CapacityManagerTestSuite struct {
	suite.Suite
	m *CapacityManager
}

func (s *CapacityManagerTestSuite) SetupTest() {
	s.m = New(testCapacity, tempCapacityExpiresAt)
}

func (s *CapacityManagerTestSuite) TestReadCapacity() {
	s.Equal(testCapacity, s.m.ReadCapacity())
}

func (s *CapacityManagerTestSuite) TestReleaseOneCapacity() {
	var blockID uint64 = 1
	_, released := s.m.ReleaseOneCapacity(blockID)
	s.Equal(false, released)

	_, ok := s.m.TakeOneCapacity(blockID)

	s.Equal(true, ok)

	capacity, released := s.m.ReleaseOneCapacity(blockID)
	s.Equal(true, released)

	s.Equal(testCapacity+1, capacity)
	s.Equal(testCapacity+1, s.m.ReadCapacity())
}

func (s *CapacityManagerTestSuite) TestTakeOneCapacity() {
	var blockID uint64 = 1

	capacity, ok := s.m.TakeOneCapacity(blockID)
	s.True(ok)
	s.Equal(testCapacity-1, capacity)
	s.Equal(testCapacity-1, s.m.ReadCapacity())
}

func (s *CapacityManagerTestSuite) TestTakeOneTempCapacity() {
	// take 3 actual capacity
	var sl []uint64 = []uint64{1, 2, 3}

	for _, c := range sl {
		_, ok := s.m.TakeOneCapacity(c)
		s.True(ok)
	}

	// should be 2 temp capacity left to take
	capacity, ok := s.m.TakeOneTempCapacity()
	s.True(ok)
	s.Equal(int(testCapacity)-len(sl)-1, capacity)

	capacity, ok = s.m.TakeOneTempCapacity()
	s.True(ok)
	s.Equal(int(testCapacity)-len(sl)-2, capacity)

	// now it should fail, 3 capacity + 2 temp capacity
	capacity, ok = s.m.TakeOneTempCapacity()
	s.False(ok)
	s.Equal(int(testCapacity)-len(sl)-2, capacity)

	// wait until they expire
	time.Sleep(s.m.tempCapacityExpiresAt)

	// both should be expired, we should be able to take two more
	capacity, ok = s.m.TakeOneTempCapacity()
	s.True(ok)
	s.Equal(int(testCapacity)-len(sl)-1, capacity)

	capacity, ok = s.m.TakeOneTempCapacity()
	s.True(ok)
	s.Equal(int(testCapacity)-len(sl)-2, capacity)

	// now remove one actual capacity, simulate "block done being proven"
	capacity, ok = s.m.ReleaseOneCapacity(sl[0])
	s.True(ok)
	s.Equal(int(testCapacity)-len(sl)-1, capacity)

	// and we should be able to take another temp capacity
	capacity, ok = s.m.TakeOneTempCapacity()
	s.True(ok)
	s.Equal(int(testCapacity)-len(sl)-2, capacity)
}
