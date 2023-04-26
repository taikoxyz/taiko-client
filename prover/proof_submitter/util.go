package submitter

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

var (
	errUnretryable = errors.New("unretryable")
)

// isSubmitProofTxErrorRetryable checks whether the error returned by a proof submission transaction
// is retryable.
func isSubmitProofTxErrorRetryable(err error, blockID *big.Int) bool {
	if strings.HasPrefix(err.Error(), "L1_NOT_ORACLE_PROVER") || !strings.HasPrefix(err.Error(), "L1_") {
		return true
	}

	log.Warn("🤷‍♂️ Unretryable proof submission error", "error", err, "blockID", blockID)
	return false
}

// getProveBlocksTxOpts creates a bind.TransactOpts instance using the given private key.
// Used for creating TaikoL1.proveBlock and TaikoL1.proveBlockInvalid transactions.
func getProveBlocksTxOpts(
	ctx context.Context,
	cli *ethclient.Client,
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
	sendTxFunc func() (*types.Transaction, error),
) error {
	var isUnretryableError bool
	if err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return nil
		}

		tx, err := sendTxFunc()
		if err != nil {
			err = encoding.TryParsingCustomError(err)
			if isSubmitProofTxErrorRetryable(err, blockID) {
				log.Info("Retry sending TaikoL1.proveBlock transaction", "reason", err)
				return err
			}

			isUnretryableError = true
			return nil
		}

		if _, err := rpc.WaitReceipt(ctx, cli.L1, tx); err != nil {
			log.Warn("Failed to wait till transaction executed", "blockID", blockID, "txHash", tx.Hash(), "error", err)
			return err
		}

		return nil
	}, backoff.NewExponentialBackOff()); err != nil {
		return fmt.Errorf("failed to send TaikoL1.proveBlock transaction: %w", err)
	}

	if isUnretryableError {
		return errUnretryable
	}

	return nil
}

// HashSignAndSetEvidenceForOracleProof hashes and signs the TaikoL1Evidence according to the
// protoco spec to generate an "oracle proof" via the signature and v value.
func hashSignAndSetEvidenceForOracleProof(
	evidence *encoding.TaikoL1Evidence,
	privateKey *ecdsa.PrivateKey,
) ([]byte, uint8, error) {
	evidence.VerifierId = 0
	evidence.Proof = nil

	inputToSign, err := encoding.EncodeProveBlockInput(evidence)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to encode TaikoL1.proveBlock inputs: %w", err)
	}

	hashed := crypto.Keccak256Hash(inputToSign)

	sig, err := crypto.Sign(hashed.Bytes(), privateKey)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to sign TaikoL1Evidence: %w", err)
	}

	// add 27 to be able to be ecrecover in solidity
	v := uint8(int(sig[64])) + 27

	evidence.VerifierId = uint16(v)
	evidence.Proof = sig

	return sig, v, nil
}
