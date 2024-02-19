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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"modernc.org/mathutil"

	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type Config struct {
	// The gap number between a block be confirmed and the latest block.
	Confirmations uint64
	// The maximum gas price can be used to send transaction.
	MaxGasPrice uint64
	gasRate     uint64
	// The maximum number of pending transactions.
	maxPendTxs int

	retryTimes uint64
}

type TxConfirm struct {
	retryTimes uint64
	confirms   uint64

	TxID    uint64
	baseTx  *types.DynamicFeeTx
	Tx      *types.Transaction
	Receipt *types.Receipt

	Error error
}

type Sender struct {
	ctx context.Context
	*Config

	chainID *big.Int
	header  *types.Header
	client  *rpc.EthClient

	Opts *bind.TransactOpts

	globalTxID     uint64
	unconfirmedTxs map[uint64]*TxConfirm
	txConfirmCh    map[uint64]chan *TxConfirm

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
		chainID:        chainID,
		header:         header,
		client:         client,
		Opts:           opts,
		unconfirmedTxs: make(map[uint64]*TxConfirm, cfg.maxPendTxs),
		txConfirmCh:    make(map[uint64]chan *TxConfirm, cfg.maxPendTxs),
		stopCh:         make(chan struct{}),
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
func (s *Sender) WaitTxConfirm(txID uint64) (bool, <-chan *TxConfirm) {
	confirmCh, ok := s.txConfirmCh[txID]
	return ok, confirmCh
}

// SendTransaction sends a transaction to the target address.
func (s *Sender) SendTransaction(tx *types.Transaction) (uint64, error) {
	if len(s.unconfirmedTxs) >= s.maxPendTxs {
		return 0, fmt.Errorf("too many pending transactions")
	}
	txID := atomic.AddUint64(&s.globalTxID, 1)
	confirmTx := &TxConfirm{
		TxID: txID,
		baseTx: &types.DynamicFeeTx{
			ChainID:   s.chainID,
			To:        tx.To(),
			GasFeeCap: tx.GasFeeCap(),
			GasTipCap: tx.GasTipCap(),
			Gas:       tx.Gas(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		},
		Tx: tx,
	}
	err := s.sendTx(confirmTx)
	if err != nil {
		return 0, err
	}
	// Add the transaction to the unconfirmed transactions
	s.unconfirmedTxs[txID] = confirmTx
	s.txConfirmCh[txID] = make(chan *TxConfirm, 1)

	return txID, nil
}

func (s *Sender) sendTx(confirmTx *TxConfirm) error {
	confirmTx.retryTimes++
	tx := confirmTx.baseTx
	baseTx := &types.DynamicFeeTx{
		ChainID:   s.chainID,
		To:        tx.To,
		GasFeeCap: tx.GasFeeCap,
		GasTipCap: tx.GasTipCap,
		Gas:       tx.Gas,
		Value:     tx.Value,
		Data:      tx.Data,
	}
	var (
		rawTx *types.Transaction
	)
	// Try 3 retryTimes if nonce is not correct.
	for i := 0; i < 3; i++ {
		nonce, err := s.client.NonceAt(s.ctx, s.Opts.From, nil)
		if err != nil {
			return err
		}
		baseTx.Nonce = nonce
		rawTx, err = s.Opts.Signer(s.Opts.From, types.NewTx(baseTx))
		if err != nil {
			return err
		}
		confirmTx.Error = s.client.SendTransaction(s.ctx, rawTx)
		// Check if the error is nonce too low or nonce too high
		if strings.Contains(err.Error(), "nonce too low") ||
			strings.Contains(err.Error(), "nonce too high") {
			log.Warn("nonce is not correct, retry to send transaction", "tx_hash", rawTx.Hash().String(), "nonce", nonce, "err", err)
			time.Sleep(time.Millisecond * 500)
			continue
		} else if err != nil {
			log.Error("failed to send transaction", "tx_hash", rawTx.Hash().String(), "err", err)
			return err
		}

		confirmTx.baseTx = baseTx
		confirmTx.Tx = rawTx
		break
	}

	return nil
}

func (s *Sender) resendTx(confirmTx *TxConfirm) {
	if confirmTx.retryTimes >= s.retryTimes {
		// TODO: add the transaction to the failed transactions
	}

	// Increase the gas price.
	gas := confirmTx.baseTx.Gas
	confirmTx.baseTx.Gas = mathutil.MinUint64(s.MaxGasPrice, gas+gas/s.gasRate)

	confirmTx.Error = s.sendTx(confirmTx)
}

func (s *Sender) updateGasTipGasFee(head *types.Header) error {
	// Get the gas tip cap
	gasTipCap, err := s.client.SuggestGasTipCap(s.ctx)
	if err != nil {
		return err
	}
	s.Opts.GasTipCap = gasTipCap

	// Get the gas fee cap
	gasFeeCap := new(big.Int).Add(gasTipCap, new(big.Int).Mul(head.BaseFee, big.NewInt(2)))
	// Check if the gas fee cap is less than the gas tip cap
	if gasFeeCap.Cmp(gasTipCap) < 0 {
		return fmt.Errorf("maxFeePerGas (%v) < maxPriorityFeePerGas (%v)", gasFeeCap, gasTipCap)
	}
	s.Opts.GasFeeCap = gasFeeCap

	return nil
}

func (s *Sender) loop() {
	defer s.wg.Done()

	tick := time.NewTicker(time.Second * 3)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			if len(s.unconfirmedTxs) == 0 {
				continue
			}
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
			// Check the unconfirmed transactions
			s.checkPendingTransactions()

		case <-s.ctx.Done():
			return
		case <-s.stopCh:
			return
		}
	}
}

func (s *Sender) checkPendingTransactions() {
	for txID, txConfirm := range s.unconfirmedTxs {
		if txConfirm.Error != nil {
			s.resendTx(txConfirm)
			continue
		}
		if txConfirm.Receipt == nil {
			// Get the transaction receipt
			receipt, err := s.client.TransactionReceipt(s.ctx, txConfirm.Tx.Hash())
			if err != nil {
				log.Warn("failed to get the transaction receipt", "tx_hash", txConfirm.Tx.Hash().String(), "err", err)
				continue
			}
			txConfirm.Receipt = receipt
		}

		txConfirm.confirms = s.header.Number.Uint64() - txConfirm.Receipt.BlockNumber.Uint64()
		// Check if the transaction is confirmed
		if s.header.Number.Uint64()-txConfirm.confirms >= s.Confirmations {
			select {
			case s.txConfirmCh[txID] <- txConfirm:
			default:
			}
			// Remove the transaction from the unconfirmed transactions
			delete(s.unconfirmedTxs, txID)
			delete(s.txConfirmCh, txID)
		}
	}
}
