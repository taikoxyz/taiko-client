package sender

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
	"github.com/pborman/uuid"
	"modernc.org/mathutil"

	"github.com/taikoxyz/taiko-client/pkg/rpc"
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
			ChainID:    s.client.ChainID,
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
			ChainID:    uint256.MustFromBig(s.client.ChainID),
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

func (s *Sender) handleReorgTransactions() { // nolint: unused
	content, err := rpc.Content(s.ctx, s.client)
	if err != nil {
		log.Warn("failed to get the unconfirmed transactions", "address", s.Opts.From.String(), "err", err)
		return
	}
	if len(content) == 0 {
		return
	}

	txs := map[common.Hash]*types.Transaction{}
	for _, txMapStatus := range content {
		for key, txMapNonce := range txMapStatus {
			addr := common.HexToAddress(key)
			if addr != s.Opts.From {
				continue
			}
			for _, tx := range txMapNonce {
				txs[tx.Hash()] = tx
			}
		}
	}
	// Remove the already handled transactions.
	for _, confirm := range s.unconfirmedTxs.Items() {
		delete(txs, confirm.Tx.Hash())
	}
	for _, tx := range txs {
		baseTx, err := s.makeTxData(tx)
		if err != nil {
			log.Warn("failed to make the transaction data when handle reorg txs", "tx_hash", tx.Hash().String(), "err", err)
			return
		}
		txID := uuid.New()
		confirm := &TxConfirm{
			TxID:   txID,
			Tx:     tx,
			baseTx: baseTx,
		}
		s.unconfirmedTxs.Set(txID, confirm)
		s.txConfirmCh.Set(txID, make(chan *TxConfirm, 1))
		log.Info("handle reorg tx", "tx_hash", tx.Hash().String(), "tx_id", txID)
	}
}

func setDefault[T uint64 | time.Duration](src, dest T) T {
	if src == 0 {
		return dest
	}
	return src
}

func setConfig(config *Config) *Config {
	if config == nil {
		return DefaultConfig
	}
	return &Config{
		Confirm:        setDefault(config.Confirm, DefaultConfig.Confirm),
		RetryTimes:     setDefault(config.RetryTimes, DefaultConfig.RetryTimes),
		MaxWaitingTime: setDefault(config.MaxWaitingTime, DefaultConfig.MaxWaitingTime),
		GasLimit:       setDefault(config.GasLimit, DefaultConfig.GasLimit),
		GasGrowthRate:  setDefault(config.GasGrowthRate, DefaultConfig.GasGrowthRate),
		MaxGasFee:      setDefault(config.MaxGasFee, DefaultConfig.MaxGasFee),
		MaxBlobFee:     setDefault(config.MaxBlobFee, DefaultConfig.MaxBlobFee),
	}
}
