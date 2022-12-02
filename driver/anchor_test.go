package driver

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *DriverTestSuite) TestNewAnchorTransactor() {
	opts, err := s.d.ChainSyncer().newAnchorTransactor(context.Background(), common.Big0)
	s.Nil(err)
	s.Equal(true, opts.NoSend)
	s.Equal(common.Big0, opts.GasPrice)
	s.Equal(common.Big0, opts.Nonce)
	s.Equal(bindings.GoldenTouchAddress, opts.From)
}
