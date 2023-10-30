package state

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/testutils"
)

func (s *DriverStateTestSuite) TestGetL1Current() {
	s.NotNil(s.s.GetL1Current())
}

func (s *DriverStateTestSuite) TestSetL1Current() {
	h := &types.Header{ParentHash: testutils.RandomHash()}
	s.s.SetL1Current(h)
	s.Equal(h.Hash(), s.s.GetL1Current().Hash())

	// should warn, but not panic
	s.NotPanics(func() {
		s.s.SetL1Current(nil)
	})
}

func (s *DriverStateTestSuite) TestResetL1CurrentEmptyHeight() {
	_, err := s.s.ResetL1Current(context.Background(), common.Big0)
	s.Nil(err)

	_, err = s.s.ResetL1Current(context.Background(), common.Big0)
	s.Nil(err)
}

func (s *DriverStateTestSuite) TestResetL1CurrentEmptyID() {
	_, err := s.s.ResetL1Current(context.Background(), common.Big1)
	s.ErrorContains(err, "not found")
}

func (s *DriverStateTestSuite) TestResetL1CurrentCtxErr() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := s.s.ResetL1Current(ctx, common.Big0)
	s.ErrorContains(err, "context canceled")
}
