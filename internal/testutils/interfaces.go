package testutils

import (
	"context"

	"github.com/taikoxyz/taiko-client/cmd/utils"
)

type BlobSyncer interface {
	ProcessL1Blocks(ctx context.Context) error
}

type Proposer interface {
	utils.SubcommandApplication
	ProposeOp(ctx context.Context) error
	ProposeTxLists(ctx context.Context, txListsBytes [][]byte) []error
}
