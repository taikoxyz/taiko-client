package auction

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings"
)

var _ Strategy = &AlwaysBidStrategy{}

// AlwaysBidStrategy is a bid strategy always bids, no matter what, to win a block if it can.
// it has no regard for profitably or caps on amounts.
type AlwaysBidStrategy struct {
	startingBid *big.Int
	deposit     uint64
}

func NewAlwaysBidStrategy() *AlwaysBidStrategy {
	return &AlwaysBidStrategy{
		startingBid: big.NewInt(10000000),
		deposit:     1000000000,
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
	// but do the minimum next bid, which should be 10 percent lower than the existing one
	var feePerGas *big.Int
	if currentBid.Prover == common.HexToAddress("0x0000000000000000000000000000000000000000") {
		feePerGas = s.startingBid
	} else {
		feePerGas = new(big.Int).Sub(currentBid.FeePerGas, (new(big.Int).Div(currentBid.FeePerGas, big.NewInt(10))))
	}
	return bindings.TaikoDataBid{
		Deposit:     s.deposit,
		FeePerGas:   feePerGas,
		ProofWindow: 2000,
	}, nil
}
