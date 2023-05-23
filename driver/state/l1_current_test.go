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
}

func (s *DriverStateTestSuite) TestResetL1CurrentEmptyHeight() {
	_, l1Current, err := s.s.ResetL1Current(context.Background(), &HeightOrID{ID: common.Big0})
	s.Nil(err)
	s.Zero(l1Current.Uint64())

	_, _, err = s.s.ResetL1Current(context.Background(), &HeightOrID{Height: common.Big0})
	s.Nil(err)
}

func (s *DriverStateTestSuite) TestResetL1CurrentEmptyID() {
	_, _, err := s.s.ResetL1Current(context.Background(), &HeightOrID{Height: common.Big1})
	s.NotNil(err)
}
