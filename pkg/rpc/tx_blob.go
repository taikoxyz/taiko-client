package rpc

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/holiman/uint256"
)

func (c *EthClient) TransactBlobTx(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	// Sign the transaction and schedule it for execution
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}
	// Create blob tx.
	rawTx, err := c.createBlobTx(opts, data)
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

func (c *EthClient) createBlobTx(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	header, err := c.HeaderByNumber(opts.Context, nil)
	if err != nil {
		return nil, err
	}
	// Estimate TipCap
	gasTipCap := opts.GasTipCap
	if gasTipCap == nil {
		tip, err := c.SuggestGasTipCap(opts.Context)
		if err != nil {
			return nil, err
		}
		gasTipCap = tip
	}
	// Estimate FeeCap
	gasFeeCap := opts.GasFeeCap
	if gasFeeCap == nil {
		gasFeeCap = new(big.Int).Add(
			gasTipCap,
			new(big.Int).Mul(header.BaseFee, big.NewInt(2)),
		)
	}
	if gasFeeCap.Cmp(gasTipCap) < 0 {
		return nil, fmt.Errorf("maxFeePerGas (%v) < maxPriorityFeePerGas (%v)", gasFeeCap, gasTipCap)
	}
	// Estimate GasLimit
	gasLimit := opts.GasLimit
	if opts.GasLimit == 0 {
		var err error
		gasLimit, err = c.EstimateGas(opts.Context, ethereum.CallMsg{
			From:      opts.From,
			To:        nil,
			GasPrice:  nil,
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
			Value:     nil,
			Data:      nil,
		})
		if err != nil {
			return nil, err
		}
	}
	// create the transaction
	nonce, err := c.getNonce(opts)
	if err != nil {
		return nil, err
	}
	chainID, err := c.ChainID(opts.Context)
	if err != nil {
		return nil, err
	}

	sidecar, err := makeSidecarWithSingleBlob(data)
	if err != nil {
		return nil, err
	}

	var blobFeeCap uint64 = 100066
	if header.ExcessBlobGas != nil {
		blobFeeCap = *header.ExcessBlobGas
	}

	baseTx := &types.BlobTx{
		ChainID:    uint256.NewInt(chainID.Uint64()),
		Nonce:      nonce,
		GasTipCap:  uint256.NewInt(gasTipCap.Uint64()),
		GasFeeCap:  uint256.NewInt(gasFeeCap.Uint64()),
		Gas:        gasLimit,
		BlobFeeCap: uint256.MustFromBig(eip4844.CalcBlobFee(blobFeeCap)),
		BlobHashes: sidecar.BlobHashes(),
		Sidecar:    sidecar,
	}
	return types.NewTx(baseTx), nil
}

func (c *EthClient) getNonce(opts *bind.TransactOpts) (uint64, error) {
	if opts.Nonce == nil {
		return c.PendingNonceAt(opts.Context, opts.From)
	}
	return opts.Nonce.Uint64(), nil
}

func makeSidecarWithSingleBlob(data []byte) (*types.BlobTxSidecar, error) {
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
