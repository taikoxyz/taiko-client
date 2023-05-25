package bid

import (
	"context"
	"math/big"
)

// AlwaysBidStrategy is a bid strategy always bids, no matter what, to win a block if it can.
// it has no regard for profitably or caps on amounts.
type AlwaysBidStrategy struct {
}

func NewAlwaysBidStrategy() *AlwaysBidStrategy {
	return &AlwaysBidStrategy{}
}

func (b *AlwaysBidStrategy) ShouldBid(ctx context.Context, currentBid *big.Int) (bool, error) {
	return true, nil
}

func (b *AlwaysBidStrategy) NextBidAmount(ctx context.Context, currentBid *big.Int) (*big.Int, error) {
	return currentBid.Sub(currentBid, big.NewInt(1000000)), nil
}
