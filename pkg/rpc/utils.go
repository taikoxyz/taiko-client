package rpc

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

// GetProtocolStateVariables gets the protocol states from TaikoL1 contract.
func GetProtocolStateVariables(
	taikoL1Client *bindings.TaikoL1Client,
	opts *bind.CallOpts,
) (*bindings.TaikoDataStateVariables, error) {
	stateVars, err := taikoL1Client.GetStateVariables(opts)
	if err != nil {
		return nil, err
	}
	return &stateVars, nil
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

// NeedNewProof checks whether the L2 block still needs a new proof.
func NeedNewProof(
	ctx context.Context,
	cli *Client,
	id *big.Int,
	proverAddress common.Address,
	realProofSkipSize *big.Int,
) (bool, error) {
	if realProofSkipSize != nil && id.Uint64()%realProofSkipSize.Uint64() != 0 {
		log.Info(
			"Skipping valid block proof",
			"blockID", id.Uint64(),
			"skipSize", realProofSkipSize.Uint64(),
		)

		return false, nil
	}

	var parent *types.Header
	if id.Cmp(common.Big1) == 0 {
		header, err := cli.L2.HeaderByNumber(ctx, common.Big0)
		if err != nil {
			return false, err
		}

		parent = header
	} else {
		parentL1Origin, err := cli.WaitL1Origin(ctx, new(big.Int).Sub(id, common.Big1))
		if err != nil {
			return false, err
		}

		if parent, err = cli.L2.HeaderByHash(ctx, parentL1Origin.L2BlockHash); err != nil {
			return false, err
		}
	}

	fc, err := cli.TaikoL1.GetForkChoice(nil, id, parent.Hash(), uint32(parent.GasUsed))
	if err != nil {
		if !strings.Contains(encoding.TryParsingCustomError(err).Error(), "L1_FORK_CHOICE_NOT_FOUND") {
			return false, encoding.TryParsingCustomError(err)
		}

		return true, nil
	}

	if fc.Prover == encoding.OracleProverAddress || fc.Prover == encoding.SystemProverAddress {
		return true, nil
	}

	if proverAddress == fc.Prover {
		log.Info("📬 Block's proof has already been submitted by current prover", "blockID", id)
		return false, nil
	}

	log.Info(
		"📬 Block's proof has already been submitted by another prover",
		"blockID", id,
		"prover", fc.Prover,
		"provenAt", fc.ProvenAt,
	)

	return false, nil
}

// SetHead makes a `debug_setHead` RPC call to set the chain's head, should only be used
// for testing purpose.
func SetHead(ctx context.Context, rpc *rpc.Client, headNum *big.Int) error {
	return gethclient.New(rpc).SetHead(ctx, headNum)
}

// StringToBytes32 converts the given string to [32]byte.
func StringToBytes32(str string) [32]byte {
	var b [32]byte
	copy(b[:], []byte(str))

	return b
}
