package chainiterator

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/testutils"
)

type BlockBatchIteratorTestSuite struct {
	testutils.ClientTestSuite
}

func (s *BlockBatchIteratorTestSuite) TestIter() {
	var maxBlocksReadPerEpoch uint64 = 2

	headHeight, err := s.RpcClient.L1.BlockNumber(context.Background())
	s.Nil(err)
	s.Greater(headHeight, uint64(0))

	lastEnd := common.Big0

	iter, err := NewBlockBatchIterator(context.Background(), &BlockBatchIteratorConfig{
		Client:                s.RpcClient.L1,
		MaxBlocksReadPerEpoch: &maxBlocksReadPerEpoch,
		StartHeight:           common.Big0,
		EndHeight:             new(big.Int).SetUint64(headHeight),
		OnBlocks: func(
			ctx context.Context,
			start, end *types.Header,
			updateCurrentFunc UpdateCurrentFunc,
			onReorgFunc OnReorgFunc,
			endIterFunc EndIterFunc,
		) (bool, error) {
			s.Equal(lastEnd.Uint64(), start.Number.Uint64())
			lastEnd = end.Number
			return false, nil
		},
	})

	s.Nil(err)
	s.Nil(iter.Iter())
	s.Equal(headHeight, lastEnd.Uint64())
}

func (s *BlockBatchIteratorTestSuite) TestIterReverse() {
	var (
		maxBlocksReadPerEpoch uint64 = 2
		startHeight           uint64 = 0
	)

	headHeight, err := s.RpcClient.L1.BlockNumber(context.Background())
	s.Nil(err)
	s.Greater(headHeight, startHeight)

	lastStart := new(big.Int).SetUint64(headHeight)

	iter, err := NewBlockBatchIterator(context.Background(), &BlockBatchIteratorConfig{
		Client:                s.RpcClient.L1,
		MaxBlocksReadPerEpoch: &maxBlocksReadPerEpoch,
		StartHeight:           new(big.Int).SetUint64(startHeight),
		EndHeight:             new(big.Int).SetUint64(headHeight),
		Reverse:               true,
		OnBlocks: func(
			ctx context.Context,
			start, end *types.Header,
			updateCurrentFunc UpdateCurrentFunc,
			onReorgFunc OnReorgFunc,
			endIterFunc EndIterFunc,
		) (bool, error) {
			s.Equal(lastStart.Uint64(), end.Number.Uint64())
			lastStart = start.Number
			return false, nil
		},
	})

	s.Nil(err)
	s.Nil(iter.Iter())
	s.Equal(startHeight, lastStart.Uint64())
}

func (s *BlockBatchIteratorTestSuite) TestIterEndFunc() {
	var maxBlocksReadPerEpoch uint64 = 2

	headHeight, err := s.RpcClient.L1.BlockNumber(context.Background())
	s.Nil(err)
	s.Greater(headHeight, maxBlocksReadPerEpoch)

	lastEnd := common.Big0

	iter, err := NewBlockBatchIterator(context.Background(), &BlockBatchIteratorConfig{
		Client:                s.RpcClient.L1,
		MaxBlocksReadPerEpoch: &maxBlocksReadPerEpoch,
		StartHeight:           common.Big0,
		EndHeight:             new(big.Int).SetUint64(headHeight),
		OnBlocks: func(
			ctx context.Context,
			start, end *types.Header,
			updateCurrentFunc UpdateCurrentFunc,
			onReorgFunc OnReorgFunc,
			endIterFunc EndIterFunc,
		) (bool, error) {
			s.Equal(lastEnd.Uint64(), start.Number.Uint64())
			lastEnd = end.Number
			endIterFunc()
			return false, nil
		},
	})

	s.Nil(err)
	s.Nil(iter.Iter())
	s.Equal(lastEnd.Uint64(), maxBlocksReadPerEpoch)
}

func (s *BlockBatchIteratorTestSuite) TestIter_ReorgEncountered() {
	var maxBlocksReadPerEpoch uint64 = 1
	var reorgRewindDepth uint64 = 1
	var rewindEveryNBlocks uint64 = 2
	var lastBlockReorged bool = false

	headHeight, err := s.RpcClient.L1.BlockNumber(context.Background())
	s.Nil(err)
	s.Greater(headHeight, uint64(0))

	lastEnd := common.Big0

	iter, err := NewBlockBatchIterator(context.Background(), &BlockBatchIteratorConfig{
		Client:                s.RpcClient.L1,
		MaxBlocksReadPerEpoch: &maxBlocksReadPerEpoch,
		StartHeight:           common.Big0,
		EndHeight:             new(big.Int).SetUint64(headHeight),
		ReorgRewindDepth:      &reorgRewindDepth,
		OnReorg: func() error {
			if lastEnd.Uint64() < reorgRewindDepth {
				lastEnd = common.Big0
			} else {
				lastEnd = new(big.Int).Sub(lastEnd, new(big.Int).SetUint64(reorgRewindDepth))
			}
			lastBlockReorged = true
			return nil
		},
		OnBlocks: func(
			ctx context.Context,
			start, end *types.Header,
			updateCurrentFunc UpdateCurrentFunc,
			onReorgFunc OnReorgFunc,
			endIterFunc EndIterFunc,
		) (bool, error) {
			// reorg every 2 blocks but not the first block
			if lastEnd != common.Big0 && lastEnd.Uint64()%rewindEveryNBlocks == 0 {
				return true, onReorgFunc()
			}

			if lastBlockReorged {
				s.Equal(start.Number.Uint64(), lastEnd.Uint64()+reorgRewindDepth)
			} else {
				s.Equal(start.Number.Uint64(), lastEnd.Uint64())
			}

			lastEnd = end.Number
			lastBlockReorged = false
			return false, nil
		},
	})

	s.Nil(err)
	s.Nil(iter.Iter())
}

func TestBlockBatchIteratorTestSuite(t *testing.T) {
	suite.Run(t, new(BlockBatchIteratorTestSuite))
}
