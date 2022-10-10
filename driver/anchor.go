package driver

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikochain/taiko-client/driver/crypto"
)

// Address of the golden touch account.
var (
	goldenTouchAddress = common.HexToAddress("0x0000777735367b36bC9B61C50022d9D0700dB4Ec")
)

// assembleAnchorTx creates a signed TaikoL2.anchor transaction.
func (b *L2ChainInserter) assembleAnchorTx(
	ctx context.Context,
	l1Height *big.Int,
	l1Hash common.Hash,
	l2Height *big.Int,
) (*types.Transaction, error) {
	opts, err := b.newAnchorTransactor(ctx, l2Height)
	if err != nil {
		return nil, err
	}

	return b.rpc.TaikoL2.Anchor(opts, l1Height, l1Hash)
}

// newAnchorTransactor is a utility method to create some transact options using
// golden touch account's private key.
func (b *L2ChainInserter) newAnchorTransactor(ctx context.Context, height *big.Int) (*bind.TransactOpts, error) {
	signer := types.LatestSignerForChainID(b.chainID)

	// Get the nonce of golden touch account at the specified height.
	nonce, err := b.rpc.L2AccountNonce(ctx, goldenTouchAddress, height)
	if err != nil {
		return nil, err
	}

	return &bind.TransactOpts{
		From: goldenTouchAddress,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != goldenTouchAddress {
				return nil, bind.ErrNotAuthorized
			}
			signature, err := crypto.SignAnchor(signer.Hash(tx).Bytes())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
		Nonce:    new(big.Int).SetUint64(nonce),
		Context:  ctx,
		GasPrice: common.Big0,
		GasLimit: b.state.anchorTxGasLimit.Uint64(),
		NoSend:   true,
	}, nil
}
