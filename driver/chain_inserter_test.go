package driver

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/testutils"
)

func (s *DriverTestSuite) TestProcessL1Blocks() {
	l1Head1, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head1, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Nil(s.d.l2ChainInserter.ProcessL1Blocks(context.Background(), l1Head1))

	// Propose an invalid L2 block
	testutils.ProposeAndInsertThrowawayBlock(&s.ClientTestSuite, s.p, s.d.ChainInserter())

	l2Head2, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Equal(l2Head2.Number.Uint64(), l2Head1.Number.Uint64())

	// Propose a valid L2 block
	testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainInserter())

	l2Head3, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(l2Head3.Number.Uint64(), l2Head2.Number.Uint64())
}

func (s *DriverTestSuite) TestGetInvalidateBlockTxOpts() {
	opts, err := s.d.l2ChainInserter.getInvalidateBlockTxOpts(context.Background(), common.Big0)

	s.Nil(err)
	s.True(opts.NoSend)
}
