package testutils

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/taikoxyz/taiko-client/cmd/utils"
	"github.com/taikoxyz/taiko-client/internal/sender"
)

type CalldataSyncer interface {
	ProcessL1Blocks(ctx context.Context, l1End *types.Header) error
}

type Proposer interface {
	utils.SubcommandApplication
	ProposeOp(ctx context.Context) error
	ProposeTxList(
		ctx context.Context,
		txList []*types.Transaction) error
	GetSender() *sender.Sender
}
