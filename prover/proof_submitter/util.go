package submitter

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

var (
	errUnretryable = errors.New("unretryable")
)

// isSubmitProofTxErrorRetryable checks whether the error returned by a proof submission transaction
// is retryable.
func isSubmitProofTxErrorRetryable(err error, blockID *big.Int) bool {
	if !strings.HasPrefix(err.Error(), "L1_") {
		return true
	}

	log.Warn("🤷 Unretryable proof submission error", "error", err, "blockID", blockID)
	metrics.ProverValidProofSubmissionErrorCounter.Inc(1)
	return false
}

// getProveBlocksTxOpts creates a bind.TransactOpts instance using the given private key.
// Used for creating TaikoL1.proveBlock and TaikoL1.proveBlockInvalid transactions.
func getProveBlocksTxOpts(
	ctx context.Context,
	cli *rpc.EthClient,
	chainID *big.Int,
	proverPrivKey *ecdsa.PrivateKey,
) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(proverPrivKey, chainID)
	if err != nil {
		return nil, err
	}
	gasTipCap, err := cli.SuggestGasTipCap(ctx)
	if err != nil {
		if rpc.IsMaxPriorityFeePerGasNotFoundError(err) {
			gasTipCap = rpc.FallbackGasTipCap
		} else {
			return nil, err
		}
	}

	opts.GasTipCap = gasTipCap

	return opts, nil
}

// sendTxWithBackoff tries to send the given proof submission transaction with a backoff policy.
func sendTxWithBackoff(
	ctx context.Context,
	cli *rpc.Client,
	blockID *big.Int,
	eventL1Hash common.Hash,
	proposedAt uint64,
	meta *bindings.TaikoDataBlockMetadata,
	sendTxFunc func(*big.Int) (*types.Transaction, error),
	retryInterval time.Duration,
	maxRetry *uint64,
	waitReceiptTimeout time.Duration,
) error {
	var (
		isUnretryableError bool
		nonce              *big.Int
		backOffPolicy      backoff.BackOff = backoff.NewConstantBackOff(retryInterval)
	)

	if maxRetry != nil {
		backOffPolicy = backoff.WithMaxRetries(backOffPolicy, *maxRetry)
	}

	if err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return nil
		}

		// Check if the corresponding L1 block is still in the canonical chain.
		l1Header, err := cli.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(meta.L1Height+1))
		if err != nil {
			log.Warn(
				"Failed to fetch L1 block",
				"blockID", blockID,
				"l1Height", meta.L1Height+1,
				"error", err,
			)
			return err
		}
		if l1Header.Hash() != eventL1Hash {
			log.Warn(
				"Reorg detected, skip the current proof submission",
				"blockID", blockID,
				"l1Height", meta.L1Height+1,
				"l1HashOld", eventL1Hash,
				"l1HashNew", l1Header.Hash(),
			)
			return nil
		}

		// check if latest verified head is ahead of this block proof
		stateVars, err := cli.GetProtocolStateVariables(&bind.CallOpts{Context: ctx})
		if err != nil {
			log.Warn(
				"Failed to fetch state variables",
				"blockID", blockID,
				"error", err,
			)
			return err
		}

		latestVerifiedId := stateVars.LastVerifiedBlockId
		if new(big.Int).SetUint64(latestVerifiedId).Cmp(blockID) >= 0 {
			log.Warn(
				"Block is already verified, skip current proof submission",
				"blockID", blockID.Uint64(),
				"latestVerifiedId", latestVerifiedId,
			)
			return nil
		}

		tx, err := sendTxFunc(nonce)
		if err != nil {
			err = encoding.TryParsingCustomError(err)
			if isSubmitProofTxErrorRetryable(err, blockID) {
				log.Info("Retry sending TaikoL1.proveBlock transaction", "blockID", blockID, "reason", err)
				if strings.Contains(err.Error(), core.ErrNonceTooLow.Error()) {
					nonce = nil
				}

				return err
			}

			isUnretryableError = true
			return nil
		}

		nonce = new(big.Int).SetUint64(tx.Nonce())
		ctxWithTimeout, cancel := context.WithTimeout(ctx, waitReceiptTimeout)
		defer cancel()

		if _, err := rpc.WaitReceipt(ctxWithTimeout, cli.L1, tx); err != nil {
			log.Warn(
				"Failed to wait till transaction executed",
				"blockID", blockID,
				"txHash", tx.Hash(),
				"nonce", nonce,
				"error", err,
			)
			return err
		}

		log.Info(
			"💰 Your block proof was accepted",
			"blockID", blockID,
			"txHash", tx.Hash(),
			"nonce", nonce,
			"proposedAt", proposedAt,
		)
		metrics.ProverValidProofSubmissionAcceptedCounter.Inc(1)
		return nil
	}, backOffPolicy); err != nil {
		if maxRetry != nil {
			log.Error("Failed to send TaikoL1.proveBlock transaction", "error", err, "maxRetry", *maxRetry)
			return errUnretryable
		}
		return fmt.Errorf("failed to send TaikoL1.proveBlock transaction: %w", err)
	}

	if isUnretryableError {
		return errUnretryable
	}

	return nil
}
