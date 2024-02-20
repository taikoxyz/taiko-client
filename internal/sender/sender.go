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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	cmap "github.com/orcaman/concurrent-map/v2"
	"modernc.org/mathutil"

	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type Config struct {
	// The gap number between a block be confirmed and the latest block.
	Confirmations uint64
	// The maximum gas price can be used to send transaction.
	MaxGasPrice uint64
	// The gas rate to increase the gas price.
	GasRate uint64
	// The maximum number of pending transactions.
	MaxPendTxs int
	// The maximum retry times to send transaction.
	RetryTimes uint64
}

type TxConfirm struct {
	retryTimes uint64
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

	chainID *big.Int
	header  *types.Header
	client  *rpc.EthClient

	opts *bind.TransactOpts

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
		chainID:        chainID,
		header:         header,
		client:         client,
		opts:           opts,
		unconfirmedTxs: cmap.New[*TxConfirm](),
		txConfirmCh:    cmap.New[chan *TxConfirm](),
		stopCh:         make(chan struct{}),
	}
	// Set the nonce
	sender.setNonce()
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
func (s *Sender) SendRaw(target *common.Address, value *big.Int, data []byte) (string, error) {
	gasLimit, err := s.client.EstimateGas(s.ctx, ethereum.CallMsg{
		From:      s.opts.From,
		To:        target,
		Value:     value,
		Data:      data,
		GasFeeCap: s.opts.GasFeeCap,
		GasTipCap: s.opts.GasTipCap,
	})
	if err != nil {
		return "", err
	}
	return s.SendTransaction(types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.chainID,
		To:        target,
		GasFeeCap: s.opts.GasFeeCap,
		GasTipCap: s.opts.GasTipCap,
		Gas:       gasLimit,
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
	// Try 3 RetryTimes if nonce is not correct.
	for i := 0; i < 3; i++ {
		baseTx.Nonce = s.opts.Nonce.Uint64()
		rawTx, err := s.opts.Signer(s.opts.From, types.NewTx(baseTx))
		if err != nil {
			return err
		}
		err = s.client.SendTransaction(s.ctx, rawTx)
		confirmTx.Error = err
		// Check if the error is nonce too low or nonce too high
		if err != nil {
			if strings.Contains(err.Error(), "nonce too low") ||
				strings.Contains(err.Error(), "nonce too high") {
				s.setNonce()
				log.Warn("nonce is not correct, retry to send transaction", "tx_hash", rawTx.Hash().String(), "err", err)
				time.Sleep(time.Millisecond * 500)
				continue
			}
			log.Error("failed to send transaction", "tx_hash", rawTx.Hash().String(), "err", err)
			return err
		}
		s.opts.Nonce = big.NewInt(s.opts.Nonce.Int64() + 1)

		confirmTx.baseTx = baseTx
		confirmTx.Tx = rawTx
		break
	}

	return nil
}

func (s *Sender) resendTx(confirmTx *TxConfirm) {
	if confirmTx.retryTimes >= s.RetryTimes {
		s.unconfirmedTxs.Remove(confirmTx.TxID)
		s.txConfirmCh.Remove(confirmTx.TxID)
		return
	}

	// Increase the gas price.
	gas := confirmTx.baseTx.Gas
	confirmTx.baseTx.Gas = mathutil.MinUint64(s.MaxGasPrice, gas+gas/s.GasRate)

	confirmTx.Error = s.sendTx(confirmTx)
}

func (s *Sender) setNonce() {
	nonce, err := s.client.PendingNonceAt(s.ctx, s.opts.From)
	if err != nil {
		log.Warn("failed to get the nonce", "from", s.opts.From, "err", err)
		return
	}
	s.opts.Nonce = new(big.Int).SetUint64(nonce)
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
	s.opts.GasTipCap = gasTipCap
	s.opts.GasFeeCap = gasFeeCap

	return nil
}

func (s *Sender) loop() {
	defer s.wg.Done()

	tick := time.NewTicker(time.Second * 3)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			if s.unconfirmedTxs.Count() == 0 {
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
	for txID, txConfirm := range s.unconfirmedTxs.Items() {
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
			confirmCh, _ := s.txConfirmCh.Get(txID)
			select {
			case confirmCh <- txConfirm:
			default:
			}
			// Remove the transaction from the unconfirmed transactions
			s.unconfirmedTxs.Remove(txID)
			s.txConfirmCh.Remove(txID)
		}
	}
}
