package auction

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings"
)

type Option string

const (
	StrategyMinimumBidFeePerGas Option = "minimum-amount"
	StrategyAlwaysBid           Option = "always-bid"
	StrategyStayProfitable      Option = "stay-profitable"
)

var (
	Strategies = []Option{StrategyMinimumBidFeePerGas, StrategyAlwaysBid, StrategyStayProfitable}
)

func IsValidStrategy(option Option) bool {
	for _, s := range Strategies {
		if s == option {
			return true
		}
	}

	return false
}

type Strategy interface {
	ShouldBid(ctx context.Context, currentBid bindings.TaikoDataBid) (bool, error)
	NextBid(
		ctx context.Context,
		proverAddress common.Address,
		currentBid bindings.TaikoDataBid,
	) (bindings.TaikoDataBid, error)
}
