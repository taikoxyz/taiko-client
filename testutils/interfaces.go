package testutils

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/cmd/utils"
)

type L2ChainSyncer interface {
	ProcessL1Blocks(ctx context.Context, l1End *types.Header) error
}

type Proposer interface {
	utils.SubcommandApplication
	ProposeOp(ctx context.Context) error
	ProposeInvalidTxListBytes(ctx context.Context) error
}
