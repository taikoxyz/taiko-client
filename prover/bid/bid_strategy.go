package bid

import (
	"context"
	"math/big"
)

type BidStrategyOption string

const (
	BidStrategyMinimumAmount  BidStrategyOption = "minimum-amount"
	BidStrategyAlways         BidStrategyOption = "always-bid"
	BidStrategyStayProfitable BidStrategyOption = "stay-profitable"
)

var (
	BidStrategies = []BidStrategyOption{BidStrategyMinimumAmount, BidStrategyAlways, BidStrategyStayProfitable}
)

func IsValidBidStrategy(option BidStrategyOption) bool {
	for _, s := range BidStrategies {
		if s == option {
			return true
		}
	}

	return false
}

type BidStrategy interface {
	ShouldBid(ctx context.Context, currentBid *big.Int) (bool, error)
	NextBidAmount(ctx context.Context, currentBid *big.Int) (*big.Int, error)
}
