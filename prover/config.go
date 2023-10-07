package prover

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Taiko prover.
type Config struct {
	L1WsEndpoint                      string
	L1HttpEndpoint                    string
	L2WsEndpoint                      string
	L2HttpEndpoint                    string
	TaikoL1Address                    common.Address
	TaikoL2Address                    common.Address
	TaikoTokenAddress                 common.Address
	L1ProverPrivKey                   *ecdsa.PrivateKey
	ZKEvmRpcdEndpoint                 string
	ZkEvmRpcdParamsPath               string
	StartingBlockID                   *big.Int
	MaxConcurrentProvingJobs          uint
	Dummy                             bool
	GuardianProver                    bool
	GuardianProverPrivateKey          *ecdsa.PrivateKey
	GuardianProofSubmissionDelay      time.Duration
	ProofSubmissionMaxRetry           uint64
	Graffiti                          string
	BackOffMaxRetrys                  uint64
	BackOffRetryInterval              time.Duration
	ProveUnassignedBlocks             bool
	RPCTimeout                        *time.Duration
	WaitReceiptTimeout                time.Duration
	ProveBlockGasLimit                *uint64
	ProveBlockTxReplacementMultiplier uint64
	ProveBlockMaxTxGasTipCap          *big.Int
	HTTPServerPort                    uint64
	Capacity                          uint64
	TempCapacityExpiresAt             time.Duration
	MinProofFee                       *big.Int
	MaxExpiry                         time.Duration
}

// NewConfigFromCliContext creates a new config instance from command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProverPrivKeyStr := c.String(flags.L1ProverPrivKey.Name)

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(l1ProverPrivKeyStr))
	if err != nil {
		return nil, fmt.Errorf("invalid L1 prover private key: %w", err)
	}

	var guardianProverPrivKey *ecdsa.PrivateKey
	if c.IsSet(flags.GuardianProver.Name) {
		if !c.IsSet(flags.GuardianProverPrivateKey.Name) {
			return nil, fmt.Errorf("guardianProver flag set without guardianProverPrivateKey set")
		}

		guardianProverPrivKey, err = crypto.ToECDSA(common.Hex2Bytes(c.String(flags.GuardianProverPrivateKey.Name)))
		if err != nil {
			return nil, fmt.Errorf("invalid guardian private key: %w", err)
		}
	}

	var startingBlockID *big.Int
	if c.IsSet(flags.StartingBlockID.Name) {
		startingBlockID = new(big.Int).SetUint64(c.Uint64(flags.StartingBlockID.Name))
	}

	var timeout *time.Duration
	if c.IsSet(flags.RPCTimeout.Name) {
		duration := c.Duration(flags.RPCTimeout.Name)
		timeout = &duration
	}

	var proveBlockTxGasLimit *uint64
	if c.IsSet(flags.ProveBlockTxGasLimit.Name) {
		gasLimit := c.Uint64(flags.ProveBlockTxGasLimit.Name)
		proveBlockTxGasLimit = &gasLimit
	}

	minProofFee, ok := new(big.Int).SetString(c.String(flags.MinProofFee.Name), 10)
	if !ok {
		return nil, fmt.Errorf("invalid minProofFee: %v", minProofFee)
	}

	proveBlockTxReplacementMultiplier := c.Uint64(flags.ProveBlockTxReplacementMultiplier.Name)
	if proveBlockTxReplacementMultiplier == 0 {
		return nil, fmt.Errorf(
			"invalid --proveBlockTxReplacementMultiplier value: %d",
			proveBlockTxReplacementMultiplier,
		)
	}

	var proveBlockMaxTxGasTipCap *big.Int
	if c.IsSet(flags.ProveBlockMaxTxGasTipCap.Name) {
		proveBlockMaxTxGasTipCap = new(big.Int).SetUint64(c.Uint64(flags.ProveBlockMaxTxGasTipCap.Name))
	}

	return &Config{
		L1WsEndpoint:                      c.String(flags.L1WSEndpoint.Name),
		L1HttpEndpoint:                    c.String(flags.L1HTTPEndpoint.Name),
		L2WsEndpoint:                      c.String(flags.L2WSEndpoint.Name),
		L2HttpEndpoint:                    c.String(flags.L2HTTPEndpoint.Name),
		TaikoL1Address:                    common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
		TaikoL2Address:                    common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
		TaikoTokenAddress:                 common.HexToAddress(c.String(flags.TaikoTokenAddress.Name)),
		L1ProverPrivKey:                   l1ProverPrivKey,
		ZKEvmRpcdEndpoint:                 c.String(flags.ZkEvmRpcdEndpoint.Name),
		ZkEvmRpcdParamsPath:               c.String(flags.ZkEvmRpcdParamsPath.Name),
		StartingBlockID:                   startingBlockID,
		MaxConcurrentProvingJobs:          c.Uint(flags.MaxConcurrentProvingJobs.Name),
		Dummy:                             c.Bool(flags.Dummy.Name),
		GuardianProver:                    c.Bool(flags.GuardianProver.Name),
		GuardianProverPrivateKey:          guardianProverPrivKey,
		GuardianProofSubmissionDelay:      c.Duration(flags.GuardianProofSubmissionDelay.Name),
		ProofSubmissionMaxRetry:           c.Uint64(flags.ProofSubmissionMaxRetry.Name),
		Graffiti:                          c.String(flags.Graffiti.Name),
		BackOffMaxRetrys:                  c.Uint64(flags.BackOffMaxRetrys.Name),
		BackOffRetryInterval:              c.Duration(flags.BackOffRetryInterval.Name),
		ProveUnassignedBlocks:             c.Bool(flags.ProveUnassignedBlocks.Name),
		RPCTimeout:                        timeout,
		WaitReceiptTimeout:                c.Duration(flags.WaitReceiptTimeout.Name),
		ProveBlockGasLimit:                proveBlockTxGasLimit,
		Capacity:                          c.Uint64(flags.ProverCapacity.Name),
		TempCapacityExpiresAt:             c.Duration(flags.TempCapacityExpiresAt.Name),
		ProveBlockTxReplacementMultiplier: proveBlockTxReplacementMultiplier,
		ProveBlockMaxTxGasTipCap:          proveBlockMaxTxGasTipCap,
		HTTPServerPort:                    c.Uint64(flags.ProverHTTPServerPort.Name),
		MinProofFee:                       minProofFee,
		MaxExpiry:                         c.Duration(flags.MaxExpiry.Name),
	}, nil
}
