package capacity_manager

import (
	"github.com/stretchr/testify/suite"
)

var (
	testCapacity uint64 = 1024
)

type CapacityManagerTestSuite struct {
	suite.Suite
	m *CapacityManager
}

func (s *CapacityManagerTestSuite) SetupTest() {
	s.m = New(testCapacity)
}

func (s *CapacityManagerTestSuite) TestReadCapacity() {
	s.Equal(testCapacity, s.m.ReadCapacity())
}

func (s *CapacityManagerTestSuite) TestReleaseOneCapacity() {
	capacity, released := s.m.ReleaseOneCapacity()
	s.Equal(true, released)

	s.Equal(testCapacity+1, capacity)
	s.Equal(testCapacity+1, s.m.ReadCapacity())
}

func (s *CapacityManagerTestSuite) TestTakeOneCapacity() {
	capacity, ok := s.m.TakeOneCapacity()
	s.True(ok)
	s.Equal(testCapacity-1, capacity)
	s.Equal(testCapacity-1, s.m.ReadCapacity())
}
