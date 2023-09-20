package testutils

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/node"
)

type CalldataSyncer interface {
	ProcessL1Blocks(ctx context.Context, l1End *types.Header) error
}

type Proposer interface {
	node.Service
	ProposeOp(ctx context.Context) error
	ProposeEmptyBlockOp(ctx context.Context) error
	L2SuggestedFeeRecipient() common.Address
	ProposeTxList(
		ctx context.Context,
		meta *encoding.TaikoL1BlockMetadataInput,
		txListBytes []byte,
		txNum uint,
		nonce *uint64,
	) error
}
