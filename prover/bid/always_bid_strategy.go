package bid

import (
	"context"
	"math/big"
	"net/rpc"
)

// AlwaysBidStrategy is a bid strategy always bids, no matter what, to win a block if it can.
// it has no regard for profitably or caps on amounts.
type AlwaysBidStrategy struct {
	rpc *rpc.Client
}

type NewAlwaysBidStrategyOpts struct {
	RPC *rpc.Client
}

func NewAlwaysBidStrategy(opts NewAlwaysBidStrategyOpts) *AlwaysBidStrategy {
	return &AlwaysBidStrategy{
		rpc: opts.RPC,
	}
}

func (b *AlwaysBidStrategy) ShouldBid(ctx context.Context, currentBid *big.Int) (bool, error) {
	return true, nil
}

func (b *AlwaysBidStrategy) NextBidAmount(ctx context.Context, currentBid *big.Int) (*big.Int, error) {
	return big.NewInt(4000), nil
}
