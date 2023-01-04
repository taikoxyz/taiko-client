package rpc

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/taikoxyz/taiko-client/bindings"
)

// GetProtocolStateVariables gets the protocol states from TaikoL1 contract.
func GetProtocolStateVariables(
	taikoL1Client *bindings.TaikoL1Client,
	opts *bind.CallOpts,
) (*bindings.ProtocolStateVariables, error) {
	var (
		stateVars = new(bindings.ProtocolStateVariables)
		err       error
	)

	stateVars.GenesisHeight,
		stateVars.GenesisTimestamp,
		stateVars.StatusBits,
		stateVars.FeeBase,
		stateVars.NextBlockID,
		stateVars.LastProposedAt,
		stateVars.AvgBlockTime,
		stateVars.LatestVerifiedHeight,
		stateVars.LatestVerifiedID,
		stateVars.AvgProofTime,
		err = taikoL1Client.GetStateVariables(opts)

	return stateVars, err
}

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
func GetReceiptsByBlock(ctx context.Context, cli *rpc.Client, block *types.Block) (types.Receipts, error) {
	reqs := make([]rpc.BatchElem, block.Transactions().Len())
	receipts := make([]*types.Receipt, block.Transactions().Len())

	for i, tx := range block.Transactions() {
		reqs[i] = rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args:   []interface{}{tx.Hash()},
			Result: &receipts[i],
		}
	}

	if err := cli.BatchCallContext(ctx, reqs); err != nil {
		return nil, err
	}

	for i := range reqs {
		if reqs[i].Error != nil {
			return nil, reqs[i].Error
		}

		if receipts[i] == nil {
			return nil, fmt.Errorf("got null receipt of transaction %s", block.Transactions()[i].Hash())
		}
	}

	return receipts, nil
}

// SetHead makes a `debug_setHead` RPC call to set the chain's head, should only be used
// for testing purpose.
func SetHead(ctx context.Context, rpc *rpc.Client, headNum *big.Int) error {
	return gethclient.New(rpc).SetHead(ctx, headNum)
}
