package state

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/testutils"
)

type DriverStateTestSuite struct {
	testutils.ClientTestSuite
	s *State
}

func (s *DriverStateTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()
	state, err := New(context.Background(), s.RpcClient)
	s.Nil(err)
	s.s = state
}

func (s *DriverStateTestSuite) TestVerifyL2Block() {
	head, err := s.RpcClient.L2.HeaderByNumber(context.Background(), nil)

	s.Nil(err)
	s.Nil(s.s.VerifyL2Block(context.Background(), head.Hash()))
}

func (s *DriverStateTestSuite) TestGetL1Head() {
	l1Head := s.s.GetL1Head()
	s.NotNil(l1Head)
}

func (s *DriverStateTestSuite) TestGetLatestVerifiedBlock() {
	b := s.s.GetLatestVerifiedBlock()
	s.NotNil(b.Hash)
}

func (s *DriverStateTestSuite) TestGetHeadBlockID() {
	s.Equal(uint64(0), s.s.GetHeadBlockID().Uint64())
}

func TestDriverStateTestSuite(t *testing.T) {
	suite.Run(t, new(DriverStateTestSuite))
}
