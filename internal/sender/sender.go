package sender

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pborman/uuid"

	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

var (
	rootSender = map[common.Address]*Sender{}
)

type Config struct {
	// The minimum confirmations to consider the transaction is confirmed.
	Confirm uint64
	// The maximum retry times to send transaction.
	RetryTimes uint64
	// The maximum waiting time for transaction in mempool.
	MaxWaitingTime time.Duration

	// The gas limit for raw transaction.
	GasLimit uint64
	// The gas rate to increase the gas price, 20 means 20% gas growth rate.
	GasGrowthRate uint64
	// The maximum gas price can be used to send transaction.
	MaxGasFee  uint64
	MaxBlobFee uint64
}

type TxConfirm struct {
	RetryTimes uint64
	confirm    uint64

	TxID string

	baseTx  types.TxData
	Tx      *types.Transaction
	Receipt *types.Receipt

	Err error
}

type Sender struct {
	ctx context.Context
	*Config

	header *types.Header
	client *rpc.EthClient

	ChainID *big.Int
	Opts    *bind.TransactOpts

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

	// Set default MaxWaitingTime.
	if cfg.MaxWaitingTime == 0 {
		cfg.MaxWaitingTime = time.Minute * 5
	}
	if cfg.GasLimit == 0 {
		cfg.GasLimit = 21000
	}

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
	sender.AdjustNonce(nil)
	// Update the gas tip and gas fee.
	err = sender.updateGasTipGasFee(header)
	if err != nil {
		return nil, err
	}
	// Add the sender to the root sender.
	if rootSender[opts.From] != nil {
		return nil, fmt.Errorf("sender already exists")
	}
	if os.Getenv("RUN_TESTS") == "" {
		rootSender[opts.From] = sender
	}

	sender.wg.Add(1)
	go sender.loop()

	return sender, nil
}

func (s *Sender) Close() {
	close(s.stopCh)
	s.wg.Wait()
}

// ConfirmChannel returns a channel to receive the transaction confirmation.
func (s *Sender) ConfirmChannel(txID string) <-chan *TxConfirm {
	confirmCh, ok := s.txConfirmCh.Get(txID)
	if !ok {
		log.Warn("transaction not found", "tx_id", txID)
	}
	return confirmCh
}

// ConfirmChannels returns all the transaction confirmation channels.
func (s *Sender) ConfirmChannels() map[string]<-chan *TxConfirm {
	channels := map[string]<-chan *TxConfirm{}
	for txID, confirmCh := range s.txConfirmCh.Items() {
		channels[txID] = confirmCh
	}
	return channels
}

// GetUnconfirmedTx returns the unconfirmed transaction by the transaction ID.
func (s *Sender) GetUnconfirmedTx(txID string) *types.Transaction {
	txConfirm, ok := s.unconfirmedTxs.Get(txID)
	if !ok {
		return nil
	}
	return txConfirm.Tx
}

// SendRaw sends a transaction to the target address.
func (s *Sender) SendRaw(nonce uint64, target *common.Address, value *big.Int, data []byte) (string, error) {
	return s.SendTransaction(types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.ChainID,
		To:        target,
		Nonce:     nonce,
		GasFeeCap: s.Opts.GasFeeCap,
		GasTipCap: s.Opts.GasTipCap,
		Gas:       s.GasLimit,
		Value:     value,
		Data:      data,
	}))
}

// SendTransaction sends a transaction to the target address.
func (s *Sender) SendTransaction(tx *types.Transaction) (string, error) {
	if s.unconfirmedTxs.Count() >= 100 {
		return "", fmt.Errorf("too many pending transactions")
	}

	txData, err := s.makeTxData(tx)
	if err != nil {
		return "", err
	}
	txID := uuid.New()
	confirmTx := &TxConfirm{
		TxID:   txID,
		baseTx: txData,
		Tx:     tx,
	}
	err = s.sendTx(confirmTx)
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

	for i := 0; i < 3; i++ {
		// Try 3 RetryTimes if nonce is not correct.
		rawTx, err := s.Opts.Signer(s.Opts.From, types.NewTx(baseTx))
		if err != nil {
			return err
		}
		confirmTx.Tx = rawTx
		err = s.client.SendTransaction(s.ctx, rawTx)
		confirmTx.Err = err
		// Check if the error is nonce too low.
		if err != nil {
			if strings.Contains(err.Error(), "nonce too low") {
				s.AdjustNonce(baseTx)
				log.Warn("nonce is not correct, retry to send transaction", "tx_hash", rawTx.Hash().String(), "err", err)
				continue
			}
			if err.Error() == "replacement transaction underpriced" {
				s.adjustGas(baseTx)
				log.Warn("replacement transaction underpriced", "tx_hash", rawTx.Hash().String(), "err", err)
				continue
			}
			log.Error("failed to send transaction", "tx_hash", rawTx.Hash().String(), "err", err)
			return err
		}
		s.Opts.Nonce = big.NewInt(s.Opts.Nonce.Int64() + 1)
		break
	}
	return nil
}

func (s *Sender) loop() {
	defer s.wg.Done()

	headCh := make(chan *types.Header, 2)
	sub, err := s.client.SubscribeNewHead(s.ctx, headCh)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	tick := time.NewTicker(time.Second * 2)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			s.resendTransaction()
		case header := <-headCh:
			// If chain appear reorg then handle mempool transactions.
			// TODO: handle reorg transactions
			s.header = header
			// Update the gas tip and gas fee
			err = s.updateGasTipGasFee(header)
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

func (s *Sender) resendTransaction() {
	for txID, txConfirm := range s.unconfirmedTxs.Items() {
		if txConfirm.Err == nil {
			continue
		}
		txConfirm.RetryTimes++
		if s.RetryTimes != 0 && txConfirm.RetryTimes >= s.RetryTimes {
			s.releaseConfirm(txID)
			continue
		}
		_ = s.sendTx(txConfirm)
	}
}

func (s *Sender) checkPendingTransactions() {
	for txID, txConfirm := range s.unconfirmedTxs.Items() {
		if txConfirm.Err != nil {
			continue
		}
		if txConfirm.Receipt == nil {
			// Ignore the transaction if it is pending.
			tx, isPending, err := s.client.TransactionByHash(s.ctx, txConfirm.Tx.Hash())
			if err != nil {
				continue
			}
			if isPending {
				// If the transaction is in mempool for too long, replace it.
				if waitTime := time.Since(tx.Time()); waitTime > s.MaxWaitingTime {
					txConfirm.Err = fmt.Errorf("transaction in mempool for too long")
				}
				continue
			}
			// Get the transaction receipt.
			receipt, err := s.client.TransactionReceipt(s.ctx, txConfirm.Tx.Hash())
			if err != nil {
				if err.Error() == "not found" {
					txConfirm.Err = err
					s.releaseConfirm(txID)
				}
				log.Warn("failed to get the transaction receipt", "tx_hash", txConfirm.Tx.Hash().String(), "err", err)
				continue
			}
			txConfirm.Receipt = receipt
			if receipt.Status != types.ReceiptStatusSuccessful {
				txConfirm.Err = fmt.Errorf("transaction reverted, hash: %s", receipt.TxHash.String())
				s.releaseConfirm(txID)
				continue
			}
		}
		txConfirm.confirm = s.header.Number.Uint64() - txConfirm.Receipt.BlockNumber.Uint64()
		if txConfirm.confirm >= s.Confirm {
			s.releaseConfirm(txID)
		}
	}
}

func (s *Sender) releaseConfirm(txID string) {
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
