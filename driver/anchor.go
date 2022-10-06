package driver

import (
	"context"
	"math/big"

	"github.com/taikochain/client-mono/driver/crypto"
	"github.com/taikochain/taiko-client/accounts/abi/bind"
	"github.com/taikochain/taiko-client/common"
	"github.com/taikochain/taiko-client/common/hexutil"
	"github.com/taikochain/taiko-client/core/types"
)

// Address of the Taiko gold finger account.
var goldenTouchAddress = common.HexToAddress("0x0000777735367b36bC9B61C50022d9D0700dB4Ec")

// newAnchorTransactor is a utility method to easily create a transaction signer
// from gold finger private key.
func (b *L2ChainInserter) newAnchorTransactor(ctx context.Context, height *big.Int) (*bind.TransactOpts, error) {
	signer := types.LatestSignerForChainID(b.chainID)

	// Get the nonce of gold finger account at the specified height
	nonce, err := b.getNonce(ctx, goldenTouchAddress, height)
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
		Context:  context.Background(),
		GasPrice: common.Big0,
		GasLimit: b.state.anchorTxGasLimit.Uint64(),
		NoSend:   true,
	}, nil
}

// prepareAnchorTx creates a signed TaikoL2.anchor transaction.
func (b *L2ChainInserter) prepareAnchorTx(ctx context.Context, l1Height *big.Int, l1Hash common.Hash, l2Height *big.Int) (*types.Transaction, error) {
	opts, err := b.newAnchorTransactor(ctx, l2Height)
	if err != nil {
		return nil, err
	}

	return b.rpc.taikoL2.Anchor(opts, l1Height, l1Hash)
}

// getNonce fetches the nonce of the given account at a specified height.
func (b *L2ChainInserter) getNonce(ctx context.Context, account common.Address, height *big.Int) (uint64, error) {
	var result hexutil.Uint64
	err := b.rpc.l2RawRPC.CallContext(ctx, &result, "eth_getTransactionCount", account, hexutil.EncodeBig(height))
	return uint64(result), err
}
