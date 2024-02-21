package sender

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	cmap "github.com/orcaman/concurrent-map/v2"

	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type Config struct {
	// The gap number between a block be confirmed and the latest block.
	Confirmations uint64
	// The maximum gas price can be used to send transaction.
	MaxGasPrice *big.Int
	// The gas rate to increase the gas price.
	GasRate uint64
	// The maximum number of pending transactions.
	MaxPendTxs int
	// The maximum retry times to send transaction.
	RetryTimes uint64
}

type TxConfirm struct {
	RetryTimes uint64
	confirms   uint64

	TxID string

	baseTx  *types.DynamicFeeTx
	Tx      *types.Transaction
	Receipt *types.Receipt

	Error error
}

type Sender struct {
	ctx context.Context
	*Config

	header *types.Header
	client *rpc.EthClient

	ChainID *big.Int
	Opts    *bind.TransactOpts

	globalTxID     uint64
	unconfirmedTxs cmap.ConcurrentMap[string, *TxConfirm] //uint64]*TxConfirm
	txConfirmCh    cmap.ConcurrentMap[string, chan *TxConfirm]

	mu     sync.Mutex
	wg     sync.WaitGroup
	stopCh chan struct{}
}

// NewSender returns a new instance of Sender.
func NewSender(ctx context.Context, cfg *Config, client *rpc.EthClient, priv *ecdsa.PrivateKey) (*Sender, error) {
	// Get the chain ID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Create a new transactor
	opts, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	if err != nil {
		return nil, err
	}
	// Do not automatically send transactions
	opts.NoSend = true

	sender := &Sender{
		ctx:            ctx,
		Config:         cfg,
		ChainID:        chainID,
		header:         header,
		client:         client,
		Opts:           opts,
		unconfirmedTxs: cmap.New[*TxConfirm](),
		txConfirmCh:    cmap.New[chan *TxConfirm](),
		stopCh:         make(chan struct{}),
	}
	// Set the nonce
	sender.adjustNonce(nil)
	// Update the gas tip and gas fee.
	err = sender.updateGasTipGasFee(header)
	if err != nil {
		return nil, err
	}

	sender.wg.Add(1)
	go sender.loop()

	return sender, nil
}

func (s *Sender) Stop() {
	close(s.stopCh)
	s.wg.Wait()
}

// WaitTxConfirm returns a channel to receive the transaction confirmation.
func (s *Sender) WaitTxConfirm(txID string) (<-chan *TxConfirm, bool) {
	confirmCh, ok := s.txConfirmCh.Get(txID)
	return confirmCh, ok
}

// SendRaw sends a transaction to the target address.
func (s *Sender) SendRaw(nonce uint64, target *common.Address, value *big.Int, data []byte) (string, error) {
	return s.SendTransaction(types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.ChainID,
		To:        target,
		Nonce:     nonce,
		GasFeeCap: s.Opts.GasFeeCap,
		GasTipCap: s.Opts.GasTipCap,
		Gas:       1000000,
		Value:     value,
		Data:      data,
	}))
}

// SendTransaction sends a transaction to the target address.
func (s *Sender) SendTransaction(tx *types.Transaction) (string, error) {
	if s.unconfirmedTxs.Count() >= s.MaxPendTxs {
		return "", fmt.Errorf("too many pending transactions")
	}
	txID := fmt.Sprint(atomic.AddUint64(&s.globalTxID, 1))
	confirmTx := &TxConfirm{
		TxID: txID,
		baseTx: &types.DynamicFeeTx{
			ChainID:   s.ChainID,
			To:        tx.To(),
			Nonce:     tx.Nonce(),
			GasFeeCap: s.Opts.GasFeeCap,
			GasTipCap: s.Opts.GasTipCap,
			Gas:       tx.Gas(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		},
		Tx: tx,
	}
	err := s.sendTx(confirmTx)
	if err != nil && !strings.Contains(err.Error(), "replacement transaction") {
		log.Error("failed to send transaction", "tx_id", txID, "tx_hash", tx.Hash().String(), "err", err)
		return "", err
	}
	// Add the transaction to the unconfirmed transactions
	s.unconfirmedTxs.Set(txID, confirmTx)
	s.txConfirmCh.Set(txID, make(chan *TxConfirm, 1))

	return txID, nil
}

func (s *Sender) sendTx(confirmTx *TxConfirm) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	baseTx := confirmTx.baseTx

	// Try 3 RetryTimes if nonce is not correct.
	rawTx, err := s.Opts.Signer(s.Opts.From, types.NewTx(baseTx))
	if err != nil {
		return err
	}
	confirmTx.Tx = rawTx
	err = s.client.SendTransaction(s.ctx, rawTx)
	confirmTx.Error = err
	// Check if the error is nonce too low.
	if err != nil {
		if strings.Contains(err.Error(), "nonce too low") {
			s.adjustNonce(baseTx)
			log.Warn("nonce is not correct, retry to send transaction", "tx_hash", rawTx.Hash().String(), "err", err)
			return nil
		}
		if err.Error() == "replacement transaction underpriced" {
			s.adjustGas(baseTx)
			log.Warn("replacement transaction underpriced", "tx_hash", rawTx.Hash().String(), "err", err)
			return nil
		}
		log.Error("failed to send transaction", "tx_hash", rawTx.Hash().String(), "err", err)
		return err
	}
	s.Opts.Nonce = big.NewInt(s.Opts.Nonce.Int64() + 1)

	return nil
}

func (s *Sender) adjustGas(baseTx *types.DynamicFeeTx) {
	rate := big.NewInt(int64(100 + s.GasRate))
	baseTx.GasFeeCap = new(big.Int).Mul(baseTx.GasFeeCap, rate)
	baseTx.GasFeeCap.Div(baseTx.GasFeeCap, big.NewInt(100))
	if s.MaxGasPrice.Cmp(baseTx.GasFeeCap) < 0 {
		baseTx.GasFeeCap = new(big.Int).Set(s.MaxGasPrice)
	}

	baseTx.GasTipCap = new(big.Int).Mul(baseTx.GasTipCap, rate)
	baseTx.GasTipCap.Div(baseTx.GasTipCap, big.NewInt(100))
	if baseTx.GasTipCap.Cmp(baseTx.GasFeeCap) > 0 {
		baseTx.GasTipCap = new(big.Int).Set(baseTx.GasFeeCap)
	}
}

func (s *Sender) adjustNonce(baseTx *types.DynamicFeeTx) {
	nonce, err := s.client.NonceAt(s.ctx, s.Opts.From, nil)
	if err != nil {
		log.Warn("failed to get the nonce", "from", s.Opts.From, "err", err)
		return
	}
	s.Opts.Nonce = new(big.Int).SetUint64(nonce)
	if baseTx != nil {
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
	if gasFeeCap.Cmp(s.MaxGasPrice) > 0 {
		gasFeeCap = new(big.Int).Set(s.MaxGasPrice)
		gasTipCap = new(big.Int).Set(s.MaxGasPrice)
	}

	s.Opts.GasTipCap = gasTipCap
	s.Opts.GasFeeCap = gasFeeCap

	return nil
}

func (s *Sender) loop() {
	defer s.wg.Done()

	tickHead := time.NewTicker(time.Second * 3)
	defer tickHead.Stop()

	tickResend := time.NewTicker(time.Second * 2)
	defer tickResend.Stop()

	for {
		select {
		case <-tickResend.C:
			if s.unconfirmedTxs.Count() == 0 {
				continue
			}
			s.resendTransaction()
			// Check the unconfirmed transactions
			s.checkPendingTransactions()
		case <-tickHead.C:
			head, err := s.client.HeaderByNumber(s.ctx, nil)
			if err != nil {
				log.Warn("failed to get the latest header", "err", err)
				continue
			}
			if s.header.Hash() == head.Hash() {
				continue
			}
			s.header = head

			// Update the gas tip and gas fee
			err = s.updateGasTipGasFee(head)
			if err != nil {
				log.Warn("failed to update gas tip and gas fee", "err", err)
			}
		case <-s.ctx.Done():
			return
		case <-s.stopCh:
			return
		}
	}
}

func (s *Sender) resendTransaction() {
	for txID, txConfirm := range s.unconfirmedTxs.Items() {
		if txConfirm.Error == nil {
			continue
		}
		txConfirm.RetryTimes++
		if s.RetryTimes != 0 && txConfirm.RetryTimes >= s.RetryTimes {
			s.unconfirmedTxs.Remove(txID)
			s.txConfirmCh.Remove(txID)
			continue
		}
		_ = s.sendTx(txConfirm)
	}
}

func (s *Sender) checkPendingTransactions() {
	for txID, txConfirm := range s.unconfirmedTxs.Items() {
		if txConfirm.Error != nil {
			continue
		}
		if txConfirm.Receipt == nil {
			// Ignore the transaction if it is pending.
			_, isPending, err := s.client.TransactionByHash(s.ctx, txConfirm.Tx.Hash())
			if err != nil || isPending {
				continue
			}
			// Get the transaction receipt.
			receipt, err := s.client.TransactionReceipt(s.ctx, txConfirm.Tx.Hash())
			if err != nil {
				if err.Error() == "not found" {
					txConfirm.Error = err
					s.releaseConfirmCh(txID)
				}
				log.Warn("failed to get the transaction receipt", "tx_hash", txConfirm.Tx.Hash().String(), "err", err)
				continue
			}
			txConfirm.Receipt = receipt
		}

		txConfirm.confirms = s.header.Number.Uint64() - txConfirm.Receipt.BlockNumber.Uint64()
		// Check if the transaction is confirmed
		if s.header.Number.Uint64()-txConfirm.confirms >= s.Confirmations {
			s.releaseConfirmCh(txID)
		}
	}
}

func (s *Sender) releaseConfirmCh(txID string) {
	txConfirm, _ := s.unconfirmedTxs.Get(txID)
	confirmCh, _ := s.txConfirmCh.Get(txID)
	select {
	case confirmCh <- txConfirm:
	default:
	}
	// Remove the transaction from the unconfirmed transactions
	s.unconfirmedTxs.Remove(txID)
	s.txConfirmCh.Remove(txID)
}
