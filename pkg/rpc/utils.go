package rpc

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/neilotoole/errgroup"
)

// WaitConfirmations won't return before N blocks confirmations have been seen
// on destination chain.
func WaitConfirmations(ctx context.Context, client *ethclient.Client, confirmations uint64, begin uint64) error {
	for {
		if deadline, ok := ctx.Deadline(); ok {
			if time.Now().After(deadline) {
				return fmt.Errorf("wait N blocks confirmations timeout, deadline: %s", deadline)
			}
		}

		latest, err := client.BlockNumber(ctx)
		if err != nil {
			log.Error("Fetch latest block number error: %w", err)
			continue
		}

		if latest < begin+confirmations {
			continue
		}

		break
	}

	return nil
}

// WaitForTx keeps waiting until the given transaction has an execution
// receipt to know whether it reverted or not.
func WaitForTx(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (*big.Int, error) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var height *big.Int
	for range ticker.C {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			continue
		}

		if receipt.Status != types.ReceiptStatusSuccessful {
			return nil, fmt.Errorf("transaction reverted, hash: %s", tx.Hash())
		}

		height = receipt.BlockNumber
		break
	}

	return height, nil
}

// GetReceiptsByBlock fetches all transaction receipts in a block.
// TODO: fetch all receipts in one GraphQL call?
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
