package bid

import (
	"context"
	"math/big"
	"net/rpc"
)

// MinimumAmountBidStrategy is a bid strategy that has a minimum amount you are willing to accept
// per wei. Once the bidding reaches that number, you will no longer bid on that block. It
// disregards profitability, and simply compares the minimum accepted fee you have said,
// and the current bid.
type MinimumAmountBidStrategy struct {
	minimumAmount *big.Int
	rpc           *rpc.Client
}

type NewMinimumAmountBidStrategyOpts struct {
	MinimumAmount *big.Int
	RPC           *rpc.Client
}

func NewMinimumAmountBidStrategy(opts NewMinimumAmountBidStrategyOpts) *MinimumAmountBidStrategy {
	return &MinimumAmountBidStrategy{
		minimumAmount: opts.MinimumAmount,
		rpc:           opts.RPC,
	}
}

func (b *MinimumAmountBidStrategy) ShouldBid(ctx context.Context, currentBid *big.Int) (bool, error) {
	return true, nil
}

func (b *MinimumAmountBidStrategy) NextBidAmount(ctx context.Context, currentBid *big.Int) (*big.Int, error) {
	return big.NewInt(4000), nil
}
