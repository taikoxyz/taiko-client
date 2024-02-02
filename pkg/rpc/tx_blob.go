package rpc

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/params"
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

	var gas *hexutil.Uint64
	if opts.GasLimit != 0 {
		var gasVal = hexutil.Uint64(opts.GasLimit)
		gas = &gasVal
	}

	rawTx, err := c.FillTransaction(opts.Context, &TransactionArgs{
		From:                 &opts.From,
		To:                   contract,
		Gas:                  gas,
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
		BlobHashes: sidecar.BlobHashes(),
		Sidecar:    sidecar,
	}

	return types.NewTx(blobTx), nil
}

// MakeSidecarWithSingleBlob make a sidecar that just include one blob.
func MakeSidecarWithSingleBlob(data []byte) (*types.BlobTxSidecar, error) {
	if len(data) > BlobBytes {
		return nil, fmt.Errorf("data is bigger than 128k")
	}
	blob := EncodeBlobs(data)[0]
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

// EncodeBlobs encode bytes into Blob type.
func EncodeBlobs(data []byte) []kzg4844.Blob {
	blobs := []kzg4844.Blob{{}}
	blobIndex := 0
	fieldIndex := -1
	numOfElems := BlobBytes / 32
	for i := 0; i < len(data); i += 31 {
		fieldIndex++
		if fieldIndex == numOfElems {
			if blobIndex >= 1 {
				break
			}
			blobs = append(blobs, kzg4844.Blob{})
			blobIndex++
			fieldIndex = 0
		}
		max := i + 31
		if max > len(data) {
			max = len(data)
		}
		copy(blobs[blobIndex][fieldIndex*32+1:], data[i:max])
	}
	return blobs
}

// DecodeBlob decode blob data.
func DecodeBlob(blob []byte) []byte {
	if len(blob) != params.BlobTxFieldElementsPerBlob*32 {
		panic("invalid blob encoding")
	}
	var data []byte
	for i, j := 0, 0; i < params.BlobTxFieldElementsPerBlob; i++ {
		data = append(data, blob[j:j+31]...)
		j += 32
	}

	i := len(data) - 1
	for ; i >= 0; i-- {
		if data[i] != 0x00 {
			break
		}
	}
	data = data[:i+1]
	return data
}
