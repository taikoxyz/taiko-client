package sender

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type Config struct {
	// The gap number between a block be confirmed and the latest block.
	Confirmations uint64
	// The maximum gas price can be used to send transaction.
	MaxGasPrice uint64
}

type TxConfirm struct {
	TxID    uint64
	Confirm uint64
	Tx      *types.Transaction
	Receipt *types.Receipt
}

type Sender struct {
	ctx context.Context
	cfg *Config

	chainID *big.Int
	client  *rpc.EthClient

	Opts *bind.TransactOpts

	unconfirmedTxs map[uint64]*types.Transaction
	txConfirmCh    map[uint64]chan<- *TxConfirm

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
	// Create a new transactor
	opts, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	if err != nil {
		return nil, err
	}
	opts.NoSend = true

	sender := &Sender{
		ctx:     ctx,
		cfg:     cfg,
		chainID: chainID,
		client:  client,
		Opts:    opts,
	}

	sender.wg.Add(1)
	go sender.loop()

	return sender, nil
}

func (s *Sender) Stop() {
	close(s.stopCh)
	s.wg.Wait()
}

func (s *Sender) WaitTxConfirm(txID uint64) <-chan *TxConfirm {
	s.txConfirmCh[txID] <- s.updateGasTipGasFee[txID]
	return ch
}

// SendTransaction sends a transaction to the target address.
func (s *Sender) SendTransaction(tx *types.Transaction) error {
	baseTx := &types.DynamicFeeTx{
		ChainID:   s.chainID,
		To:        tx.To(),
		Nonce:     s.Opts.Nonce.Uint64(),
		GasFeeCap: new(big.Int).Set(s.Opts.GasFeeCap),
		GasTipCap: new(big.Int).Set(s.Opts.GasTipCap),
		Gas:       tx.Gas(),
		Value:     tx.Value(),
		Data:      tx.Data(),
	}
	rawTx, err := s.Opts.Signer(s.Opts.From, types.NewTx(baseTx))
	if err != nil {
		return err
	}
	// Send the transaction
	err = s.client.SendTransaction(s.ctx, rawTx)
	// Check if the error is nonce too low or nonce too high
	if err == nil ||
		strings.Contains(err.Error(), "nonce too low") ||
		strings.Contains(err.Error(), "nonce too high") {

		return nil
	}

	return err
}

func (s *Sender) setNonce() error {
	// Get the nonce
	nonce, err := s.client.PendingNonceAt(s.ctx, s.Opts.From)
	if err != nil {
		return err
	}
	s.Opts.Nonce = new(big.Int).SetUint64(nonce)

	return nil
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

	headCh := make(chan *types.Header, 3)
	// Subscribe new head
	sub, err := s.client.SubscribeNewHead(s.ctx, headCh)
	if err != nil {
		log.Crit("failed to subscribe new head", "err", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case head := <-headCh:
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
