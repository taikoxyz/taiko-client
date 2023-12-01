package prover

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
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
	OracleProver                      bool
	OracleProverPrivateKey            *ecdsa.PrivateKey
	OracleProofSubmissionDelay        time.Duration
	ProofSubmissionMaxRetry           uint64
	Graffiti                          string
	RandomDummyProofDelayLowerBound   *time.Duration
	RandomDummyProofDelayUpperBound   *time.Duration
	BackOffMaxRetrys                  uint64
	BackOffRetryInterval              time.Duration
	CheckProofWindowExpiredInterval   time.Duration
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
	Allowance                         *big.Int
}

// NewConfigFromCliContext creates a new config instance from command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProverPrivKeyStr := c.String(flags.L1ProverPrivKey.Name)

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(l1ProverPrivKeyStr))
	if err != nil {
		return nil, fmt.Errorf("invalid L1 prover private key: %w", err)
	}

	oracleProverSet := c.IsSet(flags.OracleProver.Name)

	var oracleProverPrivKey *ecdsa.PrivateKey
	if oracleProverSet {
		if !c.IsSet(flags.OracleProverPrivateKey.Name) {
			return nil, fmt.Errorf("oracleProver flag set without oracleProverPrivateKey set")
		}

		oracleProverPrivKey, err = crypto.ToECDSA(common.Hex2Bytes(c.String(flags.OracleProverPrivateKey.Name)))
		if err != nil {
			return nil, fmt.Errorf("invalid oracle private key: %w", err)
		}
	} else {
		if !c.IsSet(flags.ProverCapacity.Name) {
			return nil, fmt.Errorf("capacity is required if oracleProver is not set to true")
		}
	}

	var (
		randomDummyProofDelayLowerBound *time.Duration
		randomDummyProofDelayUpperBound *time.Duration
	)
	if c.IsSet(flags.RandomDummyProofDelay.Name) {
		flagValue := c.String(flags.RandomDummyProofDelay.Name)
		splitted := strings.Split(flagValue, "-")
		if len(splitted) != 2 {
			return nil, fmt.Errorf("invalid random dummy proof delay value: %s", flagValue)
		}

		lower, err := time.ParseDuration(splitted[0])
		if err != nil {
			return nil, fmt.Errorf("invalid random dummy proof delay value: %s, err: %w", flagValue, err)
		}
		upper, err := time.ParseDuration(splitted[1])
		if err != nil {
			return nil, fmt.Errorf("invalid random dummy proof delay value: %s, err: %w", flagValue, err)
		}
		if lower > upper {
			return nil, fmt.Errorf("invalid random dummy proof delay value (lower > upper): %s", flagValue)
		}

		if upper != time.Duration(0) {
			randomDummyProofDelayLowerBound = &lower
			randomDummyProofDelayUpperBound = &upper
		}
	}

	var startingBlockID *big.Int
	if c.IsSet(flags.StartingBlockID.Name) {
		startingBlockID = new(big.Int).SetUint64(c.Uint64(flags.StartingBlockID.Name))
	}

	var timeout *time.Duration
	if c.IsSet(flags.RPCTimeout.Name) {
		duration := time.Duration(c.Uint64(flags.RPCTimeout.Name)) * time.Second
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

	var allowance *big.Int = common.Big0
	if c.IsSet(flags.Allowance.Name) {
		amt, ok := new(big.Int).SetString(c.String(flags.Allowance.Name), 10)
		if !ok {
			return nil, fmt.Errorf("error setting allowance config value: %v", c.String(flags.Allowance.Name))
		}

		allowance = amt
	}

	return &Config{
		L1WsEndpoint:                    c.String(flags.L1WSEndpoint.Name),
		L1HttpEndpoint:                  c.String(flags.L1HTTPEndpoint.Name),
		L2WsEndpoint:                    c.String(flags.L2WSEndpoint.Name),
		L2HttpEndpoint:                  c.String(flags.L2HTTPEndpoint.Name),
		TaikoL1Address:                  common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
		TaikoL2Address:                  common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
		TaikoTokenAddress:               common.HexToAddress(c.String(flags.TaikoTokenAddress.Name)),
		L1ProverPrivKey:                 l1ProverPrivKey,
		ZKEvmRpcdEndpoint:               c.String(flags.ZkEvmRpcdEndpoint.Name),
		ZkEvmRpcdParamsPath:             c.String(flags.ZkEvmRpcdParamsPath.Name),
		StartingBlockID:                 startingBlockID,
		MaxConcurrentProvingJobs:        c.Uint(flags.MaxConcurrentProvingJobs.Name),
		Dummy:                           c.Bool(flags.Dummy.Name),
		OracleProver:                    c.Bool(flags.OracleProver.Name),
		OracleProverPrivateKey:          oracleProverPrivKey,
		OracleProofSubmissionDelay:      time.Duration(c.Uint64(flags.OracleProofSubmissionDelay.Name)) * time.Second,
		ProofSubmissionMaxRetry:         c.Uint64(flags.ProofSubmissionMaxRetry.Name),
		Graffiti:                        c.String(flags.Graffiti.Name),
		RandomDummyProofDelayLowerBound: randomDummyProofDelayLowerBound,
		RandomDummyProofDelayUpperBound: randomDummyProofDelayUpperBound,
		BackOffMaxRetrys:                c.Uint64(flags.BackOffMaxRetrys.Name),
		BackOffRetryInterval:            time.Duration(c.Uint64(flags.BackOffRetryInterval.Name)) * time.Second,
		CheckProofWindowExpiredInterval: time.Duration(
			c.Uint64(flags.CheckProofWindowExpiredInterval.Name),
		) * time.Second,
		ProveUnassignedBlocks:             c.Bool(flags.ProveUnassignedBlocks.Name),
		RPCTimeout:                        timeout,
		WaitReceiptTimeout:                time.Duration(c.Uint64(flags.WaitReceiptTimeout.Name)) * time.Second,
		ProveBlockGasLimit:                proveBlockTxGasLimit,
		Capacity:                          c.Uint64(flags.ProverCapacity.Name),
		TempCapacityExpiresAt:             c.Duration(flags.TempCapacityExpiresAt.Name),
		ProveBlockTxReplacementMultiplier: proveBlockTxReplacementMultiplier,
		ProveBlockMaxTxGasTipCap:          proveBlockMaxTxGasTipCap,
		HTTPServerPort:                    c.Uint64(flags.ProverHTTPServerPort.Name),
		MinProofFee:                       minProofFee,
		MaxExpiry:                         time.Duration(c.Uint64(flags.MaxExpiry.Name)) * time.Second,
		Allowance:                         allowance,
	}, nil
}
