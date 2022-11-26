package driver

import (
	"context"
)

func (s *DriverTestSuite) TestVerfiyL2Block() {
	head, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)

	s.Nil(err)
	s.Nil(s.d.state.VerfiyL2Block(context.Background(), head.Number, head.Hash()))
}

func (s *DriverTestSuite) TestGetL1Head() {
	l1Head := s.d.state.GetL1Head()
	s.NotNil(l1Head)
}

func (s *DriverTestSuite) TestGetLastVerifiedBlockHash() {
	hash := s.d.state.GetLastVerifiedBlockHash()
	s.NotNil(hash)
}

func (s *DriverTestSuite) TestGetHeadBlockID() {
	s.Equal(uint64(0), s.d.state.GetHeadBlockID().Uint64())
}
