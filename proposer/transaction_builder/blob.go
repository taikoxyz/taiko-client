package builder

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	selector "github.com/taikoxyz/taiko-client/proposer/prover_selector"
)

// BlobTransactionBuilder is responsible for building a TaikoL1.proposeBlock transaction with txList
// bytes saved in blob.
type BlobTransactionBuilder struct {
	rpc                     *rpc.Client
	proverSelector          selector.ProverSelector
	l1BlockBuilderTip       *big.Int
	taikoL1Address          common.Address
	l2SuggestedFeeRecipient common.Address
	assignmentHookAddress   common.Address
	extraData               string
}

// NewBlobTransactionBuilder creates a new BlobTransactionBuilder instance based on giving configurations.

// Build implements the ProposeBlockTransactionBuilder interface.
func (b *BlobTransactionBuilder) Build(
	ctx context.Context,
	tierFees []encoding.TierFee,
	opts *bind.TransactOpts,
	includeParentMetaHash bool,
	txListBytes []byte,
) (*types.Transaction, error) {
	// Make sidecar in order to get blob hash.
	sideCar, err := rpc.MakeSidecar(txListBytes)
	if err != nil {
		return nil, err
	}

	assignment, assignedProver, maxFee, err := b.proverSelector.AssignProver(
		ctx,
		tierFees,
		sideCar.BlobHashes()[0],
	)
	if err != nil {
		return nil, err
	}

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

	// Initially just use the AssignmentHook default.
	hookInputData, err := encoding.EncodeAssignmentHookInput(&encoding.AssignmentHookInput{
		Assignment: assignment,
		Tip:        b.l1BlockBuilderTip,
	})
	if err != nil {
		return nil, err
	}

	encodedParams, err := encoding.EncodeBlockParams(&encoding.BlockParams{
		AssignedProver:    assignedProver,
		ExtraData:         rpc.StringToBytes32(b.extraData),
		TxListByteOffset:  common.Big0,
		TxListByteSize:    big.NewInt(int64(len(txListBytes))),
		BlobHash:          [32]byte{},
		CacheBlobForReuse: false,
		Coinbase:          b.l2SuggestedFeeRecipient,
		ParentMetaHash:    parentMetaHash,
		HookCalls: []encoding.HookCall{{
			Hook: b.assignmentHookAddress,
			Data: hookInputData,
		}},
	})
	if err != nil {
		return nil, err
	}

	opts.Value = maxFee
	rawTx, err := b.rpc.TaikoL1.ProposeBlock(
		opts,
		encodedParams,
		nil,
	)
	if err != nil {
		return nil, encoding.TryParsingCustomError(err)
	}

	proposeTx, err := b.rpc.L1.TransactBlobTx(opts, b.taikoL1Address, rawTx.Data(), sideCar)
	if err != nil {
		return nil, err
	}

	log.Debug("Transaction", " nonce", proposeTx.Nonce(), "type", proposeTx.Type())

	return proposeTx, nil
}
