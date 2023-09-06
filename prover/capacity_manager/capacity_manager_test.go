package capacity_manager

import (
	"github.com/taikoxyz/taiko-client/testutils"
)

var (
	testCapacity uint64 = 1024
)

type CapacityManagerTestSuite struct {
	testutils.ClientTestSuite
	m *CapacityManager
}

func (s *CapacityManagerTestSuite) SetupTest() {
	s.m = New(testCapacity)
}

func (s *CapacityManagerTestSuite) TestReadCapacity() {
	s.Equal(testCapacity, s.m.ReadCapacity())
}

func (s *CapacityManagerTestSuite) TestReleaseOneCapacity() {
	s.Equal(testCapacity+1, s.m.ReleaseOneCapacity())
	s.Equal(testCapacity+1, s.m.ReadCapacity())
}

func (s *CapacityManagerTestSuite) TestTakeOneCapacity() {
	s.Equal(testCapacity-1, s.m.TakeOneCapacity())
	s.Equal(testCapacity-1, s.m.ReadCapacity())
}
