package util

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/taikochain/taiko-client/core/types"
	"github.com/taikochain/taiko-client/ethclient"
	"github.com/taikochain/taiko-client/log"
)

// WaitNConfirmations won't return before N blocks confirmations have been seen
// on destination chain.
func WaitNConfirmations(ctx context.Context, client *ethclient.Client, confirmations uint64, begin uint64) error {
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
