package driver

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/testutils"
)

func (s *DriverTestSuite) TestProcessL1Blocks() {
	l1Head1, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head1, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Nil(s.d.ChainSyncer().ProcessL1Blocks(context.Background(), l1Head1))

	// Propose an invalid L2 block
	testutils.ProposeAndInsertThrowawayBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer())

	l2Head2, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Equal(l2Head2.Number.Uint64(), l2Head1.Number.Uint64())

	// Propose a valid L2 block
	testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer())

	l2Head3, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(l2Head3.Number.Uint64(), l2Head2.Number.Uint64())

	// Empty blocks
	testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.d.ChainSyncer())
	s.Nil(err)

	l2Head4, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Equal(l2Head3.Number.Uint64()+2, l2Head4.Number.Uint64())

	for _, height := range []uint64{l2Head4.Number.Uint64(), l2Head4.Number.Uint64() - 1} {
		header, err := s.d.rpc.L2.HeaderByNumber(context.Background(), new(big.Int).SetUint64(height))
		s.Nil(err)

		txCount, err := s.d.rpc.L2.TransactionCount(context.Background(), header.Hash())
		s.Nil(err)
		s.Equal(uint(1), txCount)

		anchorTx, err := s.d.rpc.L2.TransactionInBlock(context.Background(), header.Hash(), 0)
		s.Nil(err)

		method, err := encoding.TaikoL2ABI.MethodById(anchorTx.Data())
		s.Nil(err)
		s.Equal("anchor", method.Name)
	}
}

func (s *DriverTestSuite) TestGetInvalidateBlockTxOpts() {
	opts, err := s.d.ChainSyncer().getInvalidateBlockTxOpts(context.Background(), common.Big0)

	s.Nil(err)
	s.True(opts.NoSend)
}
