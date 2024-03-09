package builder

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	selector "github.com/taikoxyz/taiko-client/proposer/prover_selector"
)

// CalldataTransactionBuilder is responsible for building a TaikoL1.proposeBlock transaction with txList
// bytes saved in calldata.
type CalldataTransactionBuilder struct {
	rpc                     *rpc.Client
	proverSelector          selector.ProverSelector
	l1BlockBuilderTip       *big.Int
	l2SuggestedFeeRecipient common.Address
	assignmentHookAddress   common.Address
	extraData               string
}

// NewCalldataTransactionBuilder creates a new CalldataTransactionBuilder instance based on giving configurations.
func NewCalldataTransactionBuilder(
	rpc *rpc.Client,
	proverSelector selector.ProverSelector,
	l1BlockBuilderTip *big.Int,
	l2SuggestedFeeRecipient common.Address,
	assignmentHookAddress common.Address,
	extraData string,
) *CalldataTransactionBuilder {
	return &CalldataTransactionBuilder{
		rpc,
		proverSelector,
		l1BlockBuilderTip,
		l2SuggestedFeeRecipient,
		assignmentHookAddress,
		extraData,
	}
}

// Build implements the ProposeBlockTransactionBuilder interface.
func (b *CalldataTransactionBuilder) Build(
	ctx context.Context,
	tierFees []encoding.TierFee,
	opts *bind.TransactOpts,
	includeParentMetaHash bool,
	txListBytes []byte,
) (*types.Transaction, error) {
	// Try to assign a prover.
	assignment, assignedProver, maxFee, err := b.proverSelector.AssignProver(
		ctx,
		tierFees,
		crypto.Keccak256Hash(txListBytes),
	)
	if err != nil {
		return nil, err
	}
	opts.Value = maxFee

	var parentMetaHash = [32]byte{}
	if includeParentMetaHash {
		if parentMetaHash, err = getParentMetaHash(ctx, b.rpc); err != nil {
			return nil, err
		}
	}

	// Initially just use the AssignmentHook default.
	hookInputData, err := encoding.EncodeAssignmentHookInput(&encoding.AssignmentHookInput{
		Assignment: assignment,
		Tip:        b.l1BlockBuilderTip,
	})
	if err != nil {
		return nil, err
	}

	// ABI encode the TaikoL1.ProposeBlock parameters.
	encodedParams, err := encoding.EncodeBlockParams(&encoding.BlockParams{
		AssignedProver:    assignedProver,
		Coinbase:          b.l2SuggestedFeeRecipient,
		ExtraData:         rpc.StringToBytes32(b.extraData),
		TxListByteOffset:  common.Big0,
		TxListByteSize:    common.Big0,
		BlobHash:          [32]byte{},
		CacheBlobForReuse: false,
		ParentMetaHash:    parentMetaHash,
		HookCalls:         []encoding.HookCall{{Hook: b.assignmentHookAddress, Data: hookInputData}},
	})
	if err != nil {
		return nil, err
	}

	// Send the transaction to the L1 node.
	proposeTx, err := b.rpc.TaikoL1.ProposeBlock(opts, encodedParams, txListBytes)
	if err != nil {
		return nil, encoding.TryParsingCustomError(err)
	}

	return proposeTx, nil
}
