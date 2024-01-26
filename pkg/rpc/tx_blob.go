package rpc

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/holiman/uint256"
)

// TransactBlobTx create, sign and send blob tx.
func (c *EthClient) TransactBlobTx(
	opts *bind.TransactOpts,
	contract *common.Address,
	input []byte,
	sidecar *types.BlobTxSidecar,
) (*types.Transaction, error) {
	// Sign the transaction and schedule it for execution
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}
	// Create blob tx.
	rawTx, err := c.createBlobTx(opts, contract, input, sidecar)
	if err != nil {
		return nil, err
	}
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	if opts.NoSend {
		return signedTx, nil
	}
	if err := c.SendTransaction(opts.Context, signedTx); err != nil {
		return nil, err
	}
	return signedTx, nil
}

func (c *EthClient) createBlobTx(
	opts *bind.TransactOpts,
	contract *common.Address,
	input []byte,
	sidecar *types.BlobTxSidecar,
) (*types.Transaction, error) {
	// Get nonce.
	var nonce *hexutil.Uint64
	if opts.Nonce != nil {
		curNonce := hexutil.Uint64(opts.Nonce.Uint64())
		nonce = &curNonce
	}

	if input == nil {
		input = []byte{}
	}

	if contract == nil {
		contract = &common.Address{}
	}

	rawTx, err := c.FillTransaction(opts.Context, &TransactionArgs{
		From:                 &opts.From,
		To:                   contract,
		Gas:                  (*hexutil.Uint64)(&opts.GasLimit),
		GasPrice:             (*hexutil.Big)(opts.GasPrice),
		MaxFeePerGas:         (*hexutil.Big)(opts.GasFeeCap),
		MaxPriorityFeePerGas: (*hexutil.Big)(opts.GasTipCap),
		Value:                (*hexutil.Big)(opts.Value),
		Nonce:                nonce,
		Data:                 (*hexutil.Bytes)(&input),
		AccessList:           nil,
		ChainID:              nil,
		BlobFeeCap:           nil,
		BlobHashes:           sidecar.BlobHashes(),
	})
	if err != nil {
		return nil, err
	}
	if rawTx.Type() != types.BlobTxType {
		return nil, fmt.Errorf("expect tx type: %d, actual tx type: %d", types.BlobTxType, rawTx.Type())
	}

	blobTx := &types.BlobTx{
		ChainID:    uint256.MustFromBig(rawTx.ChainId()),
		Nonce:      rawTx.Nonce(),
		GasTipCap:  uint256.MustFromBig(rawTx.GasTipCap()),
		GasFeeCap:  uint256.MustFromBig(rawTx.GasFeeCap()),
		Gas:        rawTx.Gas(),
		To:         *rawTx.To(),
		Value:      uint256.MustFromBig(rawTx.Value()),
		Data:       rawTx.Data(),
		AccessList: rawTx.AccessList(),
		BlobFeeCap: uint256.MustFromBig(rawTx.BlobGasFeeCap()),
		BlobHashes: rawTx.BlobHashes(),
		Sidecar:    sidecar,
	}

	return types.NewTx(blobTx), nil
}

// MakeSidecarWithSingleBlob make a sidecar that just include one blob.
func MakeSidecarWithSingleBlob(data []byte) (*types.BlobTxSidecar, error) {
	if len(data) > BlobBytes {
		return nil, fmt.Errorf("data is bigger than 128k")
	}
	blob := kzg4844.Blob{}
	copy(blob[:], data)
	commitment, err := kzg4844.BlobToCommitment(blob)
	if err != nil {
		return nil, err
	}
	proof, err := kzg4844.ComputeBlobProof(blob, commitment)
	if err != nil {
		return nil, err
	}
	return &types.BlobTxSidecar{
		Blobs:       []kzg4844.Blob{blob},
		Commitments: []kzg4844.Commitment{commitment},
		Proofs:      []kzg4844.Proof{proof},
	}, nil
}
