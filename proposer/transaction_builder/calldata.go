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
		state, err := b.rpc.TaikoL1.State(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, err
		}

		parent, err := b.rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, state.SlotB.NumBlocks-1)
		if err != nil {
			return nil, err
		}

		parentMetaHash = parent.Blk.MetaHash
	}
	hookCalls := make([]encoding.HookCall, 0)

	// Initially just use the AssignmentHook default.
	hookInputData, err := encoding.EncodeAssignmentHookInput(&encoding.AssignmentHookInput{
		Assignment: assignment,
		Tip:        b.l1BlockBuilderTip,
	})
	if err != nil {
		return nil, err
	}

	hookCalls = append(hookCalls, encoding.HookCall{
		Hook: b.assignmentHookAddress,
		Data: hookInputData,
	})

	encodedParams, err := encoding.EncodeBlockParams(&encoding.BlockParams{
		AssignedProver:    assignedProver,
		Coinbase:          b.l2SuggestedFeeRecipient,
		ExtraData:         rpc.StringToBytes32(b.extraData),
		TxListByteOffset:  common.Big0,
		TxListByteSize:    common.Big0,
		BlobHash:          [32]byte{},
		CacheBlobForReuse: false,
		ParentMetaHash:    parentMetaHash,
		HookCalls:         hookCalls,
	})
	if err != nil {
		return nil, err
	}

	proposeTx, err := b.rpc.TaikoL1.ProposeBlock(
		opts,
		encodedParams,
		txListBytes,
	)
	if err != nil {
		return nil, encoding.TryParsingCustomError(err)
	}

	return proposeTx, nil
}
