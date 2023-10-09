package capacity_manager

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

var (
	testCapacity uint64 = 5
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

func (s *CapacityManagerTestSuite) TestTakeOneCapacity() {
	s.True(s.m.HoldOneCapacity(1 * time.Minute))
	capacity, ok := s.m.TakeOneCapacity(1)
	s.True(ok)
	s.Equal(testCapacity-1, capacity)
	s.Equal(testCapacity-1, s.m.ReadCapacity())
}

func (s *CapacityManagerTestSuite) TestReleaseOneCapacity() {
	var blockID uint64 = 1
	_, released := s.m.ReleaseOneCapacity(blockID)
	s.Equal(false, released)

	s.True(s.m.HoldOneCapacity(1 * time.Minute))
	_, ok := s.m.TakeOneCapacity(blockID)
	s.True(ok)

	capacity, released := s.m.ReleaseOneCapacity(blockID)
	s.True(released)

	s.Equal(testCapacity, capacity)
	s.Equal(testCapacity, s.m.ReadCapacity())
}

func TestCapacityManagerTestSuite(t *testing.T) {
	suite.Run(t, new(CapacityManagerTestSuite))
}
