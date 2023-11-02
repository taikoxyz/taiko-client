package rpc

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

var (
	ZeroAddress                = common.HexToAddress("0x0000000000000000000000000000000000000000")
	waitReceiptPollingInterval = 3 * time.Second
	defaultWaitReceiptTimeout  = 1 * time.Minute
)

// GetProtocolStateVariables gets the protocol states from TaikoL1 contract.
func GetProtocolStateVariables(
	taikoL1Client *bindings.TaikoL1Client,
	opts *bind.CallOpts,
) (*struct {
	A bindings.TaikoDataSlotA
	B bindings.TaikoDataSlotB
}, error) {
	stateVars, err := taikoL1Client.GetStateVariables(opts)
	if err != nil {
		return nil, err
	}
	return &stateVars, nil
}

// CheckProverBalance checks if the prover has the necessary balance either in TaikoL1 token balances
// or, if not, then check allowance, as contract will attempt to burn directly after
// if it doesnt have the available token balance in-contract.
func CheckProverBalance(
	ctx context.Context,
	rpc *Client,
	prover common.Address,
	taikoL1Address common.Address,
	bond *big.Int,
) (bool, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	depositedBalance, err := rpc.TaikoL1.GetTaikoTokenBalance(&bind.CallOpts{Context: ctxWithTimeout}, prover)
	if err != nil {
		return false, err
	}

	log.Info("Prover's deposited taikoTokenBalance", "balance", depositedBalance.String(), "address", prover.Hex())

	if bond.Cmp(depositedBalance) > 0 {
		// Check allowance on taiko token contract
		allowance, err := rpc.TaikoToken.Allowance(&bind.CallOpts{Context: ctxWithTimeout}, prover, taikoL1Address)
		if err != nil {
			return false, err
		}

		log.Info("Prover allowance for TaikoL1 contract", "allowance", allowance.String(), "address", prover.Hex())

		// Check prover's taiko token balance
		balance, err := rpc.TaikoToken.BalanceOf(&bind.CallOpts{Context: ctxWithTimeout}, prover)
		if err != nil {
			return false, err
		}

		log.Info("Prover's wallet taiko token balance", "balance", balance.String(), "address", prover.Hex())

		if bond.Cmp(allowance) > 0 || bond.Cmp(balance) > 0 {
			log.Info(
				"Assigned prover does not have required on-chain token balance or allowance",
				"providedProver", prover.Hex(),
				"depositedBalance", depositedBalance.String(),
				"taikoTokenBalance", balance,
				"allowance", allowance.String(),
				"bond", bond,
			)
			return false, nil
		}
	}

	return true, nil
}

// WaitReceipt keeps waiting until the given transaction has an execution
// receipt to know whether it was reverted or not.
func WaitReceipt(
	ctx context.Context,
	client *EthClient,
	tx *types.Transaction,
) (*types.Receipt, error) {
	ticker := time.NewTicker(waitReceiptPollingInterval)
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultWaitReceiptTimeout)

	defer func() {
		cancel()
		ticker.Stop()
	}()

	for {
		select {
		case <-ctxWithTimeout.Done():
			return nil, ctxWithTimeout.Err()
		case <-ticker.C:
			receipt, err := client.TransactionReceipt(ctxWithTimeout, tx.Hash())
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

// NeedNewProof checks whether the L2 block still needs a new proof.
func NeedNewProof(
	ctx context.Context,
	cli *Client,
	id *big.Int,
	proverAddress common.Address,
) (bool, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	var parent *types.Header
	if id.Cmp(common.Big1) == 0 {
		header, err := cli.L2.HeaderByNumber(ctxWithTimeout, common.Big0)
		if err != nil {
			return false, err
		}

		parent = header
	} else {
		parentL1Origin, err := cli.WaitL1Origin(ctxWithTimeout, new(big.Int).Sub(id, common.Big1))
		if err != nil {
			return false, err
		}

		if parent, err = cli.L2.HeaderByHash(ctxWithTimeout, parentL1Origin.L2BlockHash); err != nil {
			return false, err
		}
	}

	transition, err := cli.TaikoL1.GetTransition(
		&bind.CallOpts{Context: ctxWithTimeout},
		id.Uint64(),
		parent.Hash(),
	)
	if err != nil {
		if !strings.Contains(encoding.TryParsingCustomError(err).Error(), "L1_TRANSITION_NOT_FOUND") {
			return false, encoding.TryParsingCustomError(err)
		}

		return true, nil
	}

	l1Origin, err := cli.WaitL1Origin(ctxWithTimeout, id)
	if err != nil {
		return false, err
	}

	if l1Origin.L2BlockHash != transition.BlockHash {
		log.Info(
			"Different blockhash detected, try submitting a proof",
			"local", common.BytesToHash(l1Origin.L2BlockHash[:]),
			"protocol", common.BytesToHash(transition.BlockHash[:]),
		)
		return true, nil
	}

	if proverAddress == transition.Prover {
		log.Info("📬 Block's proof has already been submitted by current prover", "blockID", id)
		return false, nil
	}

	log.Info(
		"📬 Block's proof has already been submitted by another prover",
		"blockID", id,
		"prover", transition.Prover,
		"timestamp", transition.Timestamp,
	)

	return false, nil
}

type AccountPoolContent map[string]map[string]*types.Transaction

// ContentFrom fetches a given account's transactions list from a node's transactions pool.
func ContentFrom(
	ctx context.Context,
	rawRPC *rpc.Client,
	address common.Address,
) (AccountPoolContent, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	var result AccountPoolContent
	return result, rawRPC.CallContext(
		ctxWithTimeout,
		&result,
		"txpool_contentFrom",
		address,
	)
}

// IncreaseGasTipCap tries to increase the given transaction's gasTipCap.
func IncreaseGasTipCap(
	ctx context.Context,
	cli *Client,
	opts *bind.TransactOpts,
	address common.Address,
	txReplacementTipMultiplier *big.Int,
	maxGasTipCap *big.Int,
) (*bind.TransactOpts, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	log.Info("Try replacing a transaction with same nonce", "sender", address, "nonce", opts.Nonce)

	originalTx, err := GetPendingTxByNonce(ctxWithTimeout, cli, address, opts.Nonce.Uint64())
	if err != nil || originalTx == nil {
		log.Warn(
			"Original transaction not found",
			"sender", address,
			"nonce", opts.Nonce,
			"error", err,
		)

		opts.GasTipCap = new(big.Int).Mul(opts.GasTipCap, txReplacementTipMultiplier)
	} else {
		log.Info(
			"Original transaction to replace",
			"sender", address,
			"nonce", opts.Nonce,
			"gasTipCap", originalTx.GasTipCap(),
			"gasFeeCap", originalTx.GasFeeCap(),
		)

		opts.GasTipCap = new(big.Int).Mul(originalTx.GasTipCap(), txReplacementTipMultiplier)
	}

	if maxGasTipCap != nil && opts.GasTipCap.Cmp(maxGasTipCap) > 0 {
		log.Info(
			"New gasTipCap exceeds limit, keep waiting",
			"multiplier", txReplacementTipMultiplier,
			"newGasTipCap", opts.GasTipCap,
			"maxTipCap", maxGasTipCap,
		)
		return nil, txpool.ErrReplaceUnderpriced
	}

	return opts, nil
}

// GetPendingTxByNonce tries to retrieve a pending transaction with a given nonce in a node's mempool.
func GetPendingTxByNonce(
	ctx context.Context,
	cli *Client,
	address common.Address,
	nonce uint64,
) (*types.Transaction, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	content, err := ContentFrom(ctxWithTimeout, cli.L1RawRPC, address)
	if err != nil {
		return nil, err
	}

	for _, txMap := range content {
		for txNonce, tx := range txMap {
			if txNonce == strconv.Itoa(int(nonce)) {
				return tx, nil
			}
		}
	}

	return nil, nil
}

// SetHead makes a `debug_setHead` RPC call to set the chain's head, should only be used
// for testing purpose.
func SetHead(ctx context.Context, rpc *rpc.Client, headNum *big.Int) error {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	return gethclient.New(rpc).SetHead(ctxWithTimeout, headNum)
}

// StringToBytes32 converts the given string to [32]byte.
func StringToBytes32(str string) [32]byte {
	var b [32]byte
	copy(b[:], []byte(str))

	return b
}

// IsArchiveNode checks if the given node is an archive node.
func IsArchiveNode(ctx context.Context, client *EthClient, l2GenesisHeight uint64) (bool, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	if _, err := client.BalanceAt(ctxWithTimeout, ZeroAddress, new(big.Int).SetUint64(l2GenesisHeight)); err != nil {
		if strings.Contains(err.Error(), "missing trie node") {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// ctxWithTimeoutOrDefault sets a context timeout if the deadline has not passed or is not set,
// and otherwise returns the context as passed in. cancel func is always set to an empty function
// so is safe to defer the cancel.
func ctxWithTimeoutOrDefault(ctx context.Context, defaultTimeout time.Duration) (context.Context, context.CancelFunc) {
	var (
		ctxWithTimeout                    = ctx
		cancel         context.CancelFunc = func() {}
	)
	if _, ok := ctx.Deadline(); !ok {
		ctxWithTimeout, cancel = context.WithTimeout(ctx, defaultTimeout)
	}

	return ctxWithTimeout, cancel
}
