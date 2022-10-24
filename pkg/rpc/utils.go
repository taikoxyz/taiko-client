package rpc

import (
	"context"
	"fmt"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/sync/errgroup"
)

// WaitConfirmations won't return before N blocks confirmations have been seen
// on destination chain.
func WaitConfirmations(ctx context.Context, client *ethclient.Client, confirmations uint64, begin uint64) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			latest, err := client.BlockNumber(ctx)
			if err != nil {
				log.Error("Fetch latest block number error: %w", err)
				continue
			}

			if latest < begin+confirmations {
				continue
			}

			return nil
		}
	}
}

// WaitReceipt keeps waiting until the given transaction has an execution
// receipt to know whether it was reverted or not.
func WaitReceipt(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			receipt, err := client.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				continue
			}

			if receipt.Status != types.ReceiptStatusSuccessful {
				return nil, fmt.Errorf("transaction reverted, hash: %s", tx.Hash())
			}

			return receipt, nil
		}
	}
}

// GetReceiptsByBlock fetches all transaction receipts in a block.
// TODO: fetch all receipts in one GraphQL call.
func GetReceiptsByBlock(ctx context.Context, cli *ethclient.Client, block *types.Block) (types.Receipts, error) {
	g, ctx := errgroup.WithContext(ctx)

	receipts := make(types.Receipts, block.Transactions().Len())
	for i := range block.Transactions() {
		func(i int) {
			g.Go(func() (err error) {
				receipts[i], err = cli.TransactionReceipt(ctx, block.Transactions()[i].Hash())
				return err
			})
		}(i)
	}

	return receipts, g.Wait()
}

// TxWithEstimatedGasUsed represents a transaction with an estimated gasUsed information.
type TxWithEstimatedGasUsed struct {
	Tx               *types.Transaction
	EstimatedGasUsed uint64
}

// PoolContent represents a response body of a `txpool_content` RPC call.
type PoolContent map[common.Address]map[string]*types.Transaction

// ToTxsWithEstimatedGasUsed converts a PoolContent to a []TxWithEstimatedGasUsed.
func (pc PoolContent) ToTxsWithEstimatedGasUsed(
	ctx context.Context,
	signer types.Signer,
	client *ethclient.Client,
) ([]TxWithEstimatedGasUsed, error) {
	txsWithEstimatedGasUsed := []TxWithEstimatedGasUsed{}
	for _, txs := range pc {
		for _, tx := range txs {
			sender, err := signer.Sender(tx)
			if err != nil {
				return nil, err
			}

			gasUsed, err := client.EstimateGas(ctx, ethereum.CallMsg{
				From:       sender,
				To:         tx.To(),
				Gas:        tx.Gas(),
				GasPrice:   tx.GasPrice(),
				GasFeeCap:  tx.GasFeeCap(),
				GasTipCap:  tx.GasTipCap(),
				Value:      tx.Value(),
				Data:       tx.Data(),
				AccessList: tx.AccessList(),
			})
			// TODO: if this is only an execution reverted, we should let it pass.
			if err != nil {
				return nil, err
			}

			txsWithEstimatedGasUsed = append(txsWithEstimatedGasUsed, TxWithEstimatedGasUsed{tx, gasUsed})
		}
	}

	return txsWithEstimatedGasUsed, nil
}
