package chainSyncer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

func (s *ChainSyncerTestSuite) TestGetInvalidateBlockTxOpts() {
	opts, err := s.s.getInvalidateBlockTxOpts(context.Background(), common.Big0)

	s.Nil(err)
	s.True(opts.NoSend)
}
