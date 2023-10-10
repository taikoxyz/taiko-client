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
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
)

var (
	errUnretryable = errors.New("unretryable")
)

type TxAssembler func(*big.Int) (*types.Transaction, error)

// TxSender is responsible for sending proof submission transactions with a backoff policy.
type TxSender struct {
	rpc                *rpc.Client
	backOffPolicy      backoff.BackOff
	maxRetry           *uint64
	waitReceiptTimeout time.Duration
}

// NewTxSender creates a new TxSender instance.
func NewTxSender(
	cli *rpc.Client,
	retryInterval time.Duration,
	maxRetry *uint64,
	waitReceiptTimeout time.Duration,
) *TxSender {
	var backOffPolicy backoff.BackOff = backoff.NewConstantBackOff(retryInterval)
	if maxRetry != nil {
		backOffPolicy = backoff.WithMaxRetries(backOffPolicy, *maxRetry)
	}

	return &TxSender{
		rpc:                cli,
		backOffPolicy:      backOffPolicy,
		maxRetry:           maxRetry,
		waitReceiptTimeout: waitReceiptTimeout,
	}
}

// Send sends the given proof to the TaikoL1 smart contract with a backoff policy.
func (s *TxSender) Send(
	ctx context.Context,
	proofWithHeader *proofProducer.ProofWithHeader,
	txAssembler TxAssembler,
) error {
	var (
		isUnretryableError bool
		nonce              *big.Int
	)

	if err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return nil
		}

		// Check if the corresponding L1 block is still in the canonical chain.
		l1Header, err := s.rpc.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(proofWithHeader.Meta.L1Height+1))
		if err != nil {
			log.Warn(
				"Failed to fetch L1 block",
				"blockID", proofWithHeader.BlockID,
				"l1Height", proofWithHeader.Meta.L1Height+1,
				"error", err,
			)
			return err
		}
		if l1Header.Hash() != proofWithHeader.Opts.EventL1Hash {
			log.Warn(
				"Reorg detected, skip the current proof submission",
				"blockID", proofWithHeader.BlockID,
				"l1Height", proofWithHeader.Meta.L1Height+1,
				"l1HashOld", proofWithHeader.Opts.EventL1Hash,
				"l1HashNew", l1Header.Hash(),
			)
			return nil
		}

		// check if latest verified head is ahead of this block proof
		stateVars, err := s.rpc.GetProtocolStateVariables(&bind.CallOpts{Context: ctx})
		if err != nil {
			log.Warn(
				"Failed to fetch state variables",
				"blockID", proofWithHeader.BlockID,
				"error", err,
			)
			return err
		}

		latestVerifiedId := stateVars.LastVerifiedBlockId
		if new(big.Int).SetUint64(latestVerifiedId).Cmp(proofWithHeader.BlockID) >= 0 {
			log.Info(
				"Block is already verified, skip current proof submission",
				"blockID", proofWithHeader.BlockID.Uint64(),
				"latestVerifiedId", latestVerifiedId,
			)
			return nil
		}

		tx, err := txAssembler(nonce)
		if err != nil {
			err = encoding.TryParsingCustomError(err)
			if isSubmitProofTxErrorRetryable(err, proofWithHeader.BlockID) {
				log.Info("Retry sending TaikoL1.proveBlock transaction", "blockID", proofWithHeader.BlockID, "reason", err)
				if strings.Contains(err.Error(), core.ErrNonceTooLow.Error()) {
					nonce = nil
				}

				return err
			}

			isUnretryableError = true
			return nil
		}

		ctxWithTimeout, cancel := context.WithTimeout(ctx, s.waitReceiptTimeout)
		defer cancel()

		if _, err := rpc.WaitReceipt(ctxWithTimeout, s.rpc.L1, tx); err != nil {
			log.Warn(
				"Failed to wait till transaction executed",
				"blockID", proofWithHeader.BlockID,
				"txHash", tx.Hash(),
				"error", err,
			)
			return err
		}

		log.Info(
			"💰 Your block proof was accepted",
			"blockID", proofWithHeader.BlockID,
			"txHash", tx.Hash(),
		)

		return nil
	}, s.backOffPolicy); err != nil {
		if s.maxRetry != nil {
			log.Error("Failed to send TaikoL1.proveBlock transaction", "error", err, "maxRetry", *s.maxRetry)
			return errUnretryable
		}
		return fmt.Errorf("failed to send TaikoL1.proveBlock transaction: %w", err)
	}

	if isUnretryableError {
		return errUnretryable
	}

	return nil
}

// isSubmitProofTxErrorRetryable checks whether the error returned by a proof submission transaction
// is retryable.
func isSubmitProofTxErrorRetryable(err error, blockID *big.Int) bool {
	if !strings.HasPrefix(err.Error(), "L1_") {
		return true
	}

	log.Warn("🤷 Unretryable proof submission error", "error", err, "blockID", blockID)
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
