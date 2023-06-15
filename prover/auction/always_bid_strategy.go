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
	startingBid uint64
}

func NewAlwaysBidStrategy() *AlwaysBidStrategy {
	return &AlwaysBidStrategy{
		startingBid: 1,
	}
}

func (s *AlwaysBidStrategy) ShouldBid(ctx context.Context, currentBid bindings.TaikoDataBid) (bool, error) {
	return true, nil
}

func (s *AlwaysBidStrategy) NextBid(
	ctx context.Context,
	proverAddress common.Address,
	currentBid bindings.TaikoDataBid,
) (bindings.TaikoDataBid, error) {
	// re-use existing bid deposit
	deposit := currentBid.Deposit
	// but do the minimum next bid, which should be 10 percent lower than the existing one
	var feePerGas uint64
	if currentBid.FeePerGas == 0 {
		feePerGas = 1
	} else {
		feePerGas = currentBid.FeePerGas - (currentBid.FeePerGas / 10)
	}
	return bindings.TaikoDataBid{
		Deposit:   deposit,
		FeePerGas: feePerGas,
	}, nil
}
