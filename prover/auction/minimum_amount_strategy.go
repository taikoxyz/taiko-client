package auction

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

var _ Strategy = &MinimumAmountStrategy{}

// MinimumAmountStrategy is a bid strategy that has a minimum amount you are willing to accept
// per wei. Once the bidding reaches that number, you will no longer bid on that block. It
// disregards profitability, and simply compares the minimum accepted fee you have said,
// and the current bid.
type MinimumAmountStrategy struct {
	minimumAmount *big.Int
	rpc           *rpc.Client
}

type NewMinimumAmountStrategyOpts struct {
	MinimumAmount *big.Int
	RPC           *rpc.Client
}

func NewMinimumAmountStrategy(opts NewMinimumAmountStrategyOpts) *MinimumAmountStrategy {
	return &MinimumAmountStrategy{
		minimumAmount: opts.MinimumAmount,
		rpc:           opts.RPC,
	}
}

func (s *MinimumAmountStrategy) ShouldBid(ctx context.Context, currentBid bindings.TaikoDataBid) (bool, error) {
	return false, nil
}

func (s *MinimumAmountStrategy) NextBid(ctx context.Context, proverAddress common.Address, currentBid bindings.TaikoDataBid) (bindings.TaikoDataBid, error) {
	return bindings.TaikoDataBid{}, nil
}
