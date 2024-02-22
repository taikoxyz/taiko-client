package sender

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
	"modernc.org/mathutil"
)

func (s *Sender) adjustGas(txData types.TxData) {
	rate := s.GasGrowthRate + 100
	switch baseTx := txData.(type) {
	case *types.DynamicFeeTx:
		gasFeeCap := baseTx.GasFeeCap.Int64()
		gasFeeCap = gasFeeCap / 100 * int64(rate)
		gasFeeCap = mathutil.MinInt64(gasFeeCap, int64(s.MaxGasFee))
		baseTx.GasFeeCap = big.NewInt(gasFeeCap)

		gasTipCap := baseTx.GasTipCap.Int64()
		gasTipCap = gasTipCap / 100 * int64(rate)
		gasTipCap = mathutil.MinInt64(gasFeeCap, mathutil.MinInt64(gasTipCap, int64(s.MaxGasFee)))
		baseTx.GasTipCap = big.NewInt(gasTipCap)
	case *types.BlobTx:
		gasFeeCap := baseTx.GasFeeCap.Uint64()
		gasFeeCap = gasFeeCap / 100 * rate
		gasFeeCap = mathutil.MinUint64(gasFeeCap, s.MaxGasFee)
		baseTx.GasFeeCap = uint256.NewInt(gasFeeCap)

		gasTipCap := baseTx.GasTipCap.Uint64()
		gasTipCap = gasTipCap / 100 * rate
		gasTipCap = mathutil.MinUint64(gasFeeCap, mathutil.MinUint64(gasTipCap, s.MaxGasFee))
		baseTx.GasTipCap = uint256.NewInt(gasTipCap)

		blobFeeCap := baseTx.BlobFeeCap.Uint64()
		blobFeeCap = blobFeeCap / 100 * rate
		blobFeeCap = mathutil.MinUint64(blobFeeCap, s.MaxBlobFee)
		baseTx.BlobFeeCap = uint256.NewInt(blobFeeCap)
	}
}

func (s *Sender) AdjustNonce(txData types.TxData) {
	nonce, err := s.client.NonceAt(s.ctx, s.Opts.From, nil)
	if err != nil {
		log.Warn("failed to get the nonce", "from", s.Opts.From, "err", err)
		return
	}
	s.Opts.Nonce = new(big.Int).SetUint64(nonce)

	switch baseTx := txData.(type) {
	case *types.DynamicFeeTx:
		baseTx.Nonce = nonce
	case *types.BlobTx:
		baseTx.Nonce = nonce
	}
}

func (s *Sender) updateGasTipGasFee(head *types.Header) error {
	// Get the gas tip cap
	gasTipCap, err := s.client.SuggestGasTipCap(s.ctx)
	if err != nil {
		return err
	}

	// Get the gas fee cap
	gasFeeCap := new(big.Int).Add(gasTipCap, new(big.Int).Mul(head.BaseFee, big.NewInt(2)))
	// Check if the gas fee cap is less than the gas tip cap
	if gasFeeCap.Cmp(gasTipCap) < 0 {
		return fmt.Errorf("maxFeePerGas (%v) < maxPriorityFeePerGas (%v)", gasFeeCap, gasTipCap)
	}
	maxGasFee := big.NewInt(int64(s.MaxGasFee))
	if gasFeeCap.Cmp(maxGasFee) > 0 {
		gasFeeCap = new(big.Int).Set(maxGasFee)
		gasTipCap = new(big.Int).Set(maxGasFee)
	}

	s.Opts.GasTipCap = gasTipCap
	s.Opts.GasFeeCap = gasFeeCap

	return nil
}

func (s *Sender) makeTxData(tx *types.Transaction) (types.TxData, error) {
	switch tx.Type() {
	case types.DynamicFeeTxType:
		return &types.DynamicFeeTx{
			ChainID:    s.ChainID,
			To:         tx.To(),
			Nonce:      tx.Nonce(),
			GasFeeCap:  s.Opts.GasFeeCap,
			GasTipCap:  s.Opts.GasTipCap,
			Gas:        tx.Gas(),
			Value:      tx.Value(),
			Data:       tx.Data(),
			AccessList: tx.AccessList(),
		}, nil
	case types.BlobTxType:
		var to common.Address
		if tx.To() != nil {
			to = *tx.To()
		}
		return &types.BlobTx{
			ChainID:    uint256.MustFromBig(s.ChainID),
			To:         to,
			Nonce:      tx.Nonce(),
			GasFeeCap:  uint256.MustFromBig(s.Opts.GasFeeCap),
			GasTipCap:  uint256.MustFromBig(s.Opts.GasTipCap),
			Gas:        tx.Gas(),
			Value:      uint256.MustFromBig(tx.Value()),
			Data:       tx.Data(),
			AccessList: tx.AccessList(),
			BlobFeeCap: uint256.MustFromBig(tx.BlobGasFeeCap()),
			BlobHashes: tx.BlobHashes(),
			Sidecar:    tx.BlobTxSidecar(),
		}, nil
	default:
		return nil, fmt.Errorf("unsupported transaction type: %v", tx.Type())
	}
}
