package auction

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

var _ Strategy = &MinimumBidFeePerGasStrategy{}

// MinimumBidFeePerGasStrategy is a bid strategy that has a minimum amount you are willing to accept
// per wei. Once the bidding reaches that number, you will no longer bid on that block. It
// disregards profitability, and simply compares the minimum accepted fee you have said,
// and the current bid.
type MinimumBidFeePerGasStrategy struct {
	deposit             *big.Int
	minimumBidFeePerGas *big.Int
	rpc                 *rpc.Client
}

type NewMinimumBidFeePerGasStrategyOpts struct {
	MinimumBidFeePerGas *big.Int
	RPC                 *rpc.Client
	Deposit             *big.Int
}

func NewMinimumBidFeePerGasStrategy(opts NewMinimumBidFeePerGasStrategyOpts) *MinimumBidFeePerGasStrategy {
	return &MinimumBidFeePerGasStrategy{
		minimumBidFeePerGas: opts.MinimumBidFeePerGas,
		rpc:                 opts.RPC,
		deposit:             opts.Deposit,
	}
}

func (s *MinimumBidFeePerGasStrategy) ShouldBid(ctx context.Context, currentBid bindings.TaikoDataBid) (bool, error) {
	if currentBid.FeePerGas < s.minimumBidFeePerGas.Uint64() {
		return false, nil
	}

	return true, nil
}

func (s *MinimumBidFeePerGasStrategy) NextBid(
	ctx context.Context,
	proverAddress common.Address,
	currentBid bindings.TaikoDataBid,
) (bindings.TaikoDataBid, error) {
	return bindings.TaikoDataBid{
		Deposit: s.deposit.Uint64(),
	}, nil
}
