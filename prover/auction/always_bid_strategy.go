package auction

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings"
)

var _ Strategy = &AlwaysBidStrategy{}

// AlwaysBidStrategy is a bid strategy always bids, no matter what, to win a block if it can.
// it has no regard for profitably or caps on amounts.
type AlwaysBidStrategy struct {
}

func NewAlwaysBidStrategy() *AlwaysBidStrategy {
	return &AlwaysBidStrategy{}
}

func (s *AlwaysBidStrategy) ShouldBid(ctx context.Context, currentBid bindings.TaikoDataBid) (bool, error) {
	return true, nil
}

func (s *AlwaysBidStrategy) NextBid(ctx context.Context, proverAddress common.Address, currentBid bindings.TaikoDataBid) (bindings.TaikoDataBid, error) {
	return bindings.TaikoDataBid{}, nil
}
