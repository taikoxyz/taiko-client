package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/prover"
	"github.com/urfave/cli/v2"
)

const proverCmd = "prover"

var proverConf = &prover.Config{}

// Required flags used by prover.
var (
	ZkEvmRpcdEndpointFlag = &cli.StringFlag{
		Name:     "zkevmRpcdEndpoint",
		Usage:    "RPC endpoint of a ZKEVM RPCD service",
		Required: true,
		Category: proverCategory,
		Action: func(c *cli.Context, v string) error {
			proverConf.ZKEvmRpcdEndpoint = v
			return nil
		},
	}
	ZkEvmRpcdParamsPathFlag = &cli.StringFlag{
		Name:     "zkevmRpcdParamsPath",
		Usage:    "Path of ZKEVM parameters file to use",
		Required: true,
		Category: proverCategory,
		Action: func(c *cli.Context, v string) error {
			proverConf.ZkEvmRpcdParamsPath = v
			return nil
		},
	}
	L1ProverPrivKeyFlag = &cli.StringFlag{
		Name: "l1.proverPrivKey",
		Usage: "Private key of L1 prover, " +
			"who will send TaikoL1.proveBlock / TaikoL1.proveBlockInvalid transactions",
		Required: true,
		Category: proverCategory,
		Action: func(c *cli.Context, v string) error {
			k, err := crypto.ToECDSA(common.Hex2Bytes(v))
			if err != nil {
				return fmt.Errorf("invalid L1 prover private key: %w", err)
			}
			proverConf.L1ProverPrivKey = k
			return nil
		},
	}
	MinProofFeeFlag = &cli.StringFlag{
		Name:     "prover.minProofFee",
		Usage:    "Minimum accepted fee for accepting proving a block",
		Required: true,
		Category: proverCategory,
		Action: func(c *cli.Context, v string) error {
			fee, ok := new(big.Int).SetString(v, 10)
			if !ok {
				return fmt.Errorf("invalid prover.minProofFee: %v", v)
			}
			proverConf.MinProofFee = fee
			return nil
		},
	}
)

// Optional flags used by prover.
var (
	StartingBlockIDFlag = &cli.Uint64Flag{
		Name:     "startingBlockID",
		Usage:    "If set, prover will start proving blocks from the block with this ID",
		Category: proverCategory,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.StartingBlockID = new(big.Int).SetUint64(v)
			return nil
		},
	}
	MaxConcurrentProvingJobsFlag = &cli.UintFlag{
		Name:        "maxConcurrentProvingJobs",
		Usage:       "Limits the number of concurrent proving blocks jobs",
		Value:       1,
		Category:    proverCategory,
		Destination: &proverConf.MaxConcurrentProvingJobs,
		Action: func(c *cli.Context, v uint) error {
			proverConf.MaxConcurrentProvingJobs = v
			return nil
		},
	}
	// Special flags for testing.
	DummyFlag = &cli.BoolFlag{
		Name:        "dummy",
		Usage:       "Produce dummy proofs, testing purposes only",
		Value:       false,
		Category:    proverCategory,
		Destination: &proverConf.Dummy,
		Action: func(c *cli.Context, v bool) error {
			proverConf.Dummy = v
			return nil
		},
	}
	RandomDummyProofDelayFlag = &cli.StringFlag{
		Name: "randomDummyProofDelay",
		Usage: "Set the random dummy proof delay between the bounds using the format: " +
			"`lowerBound-upperBound` (e.g. `30m,1h`), testing purposes only",
		Category: proverCategory,
		Action: func(c *cli.Context, s string) error {
			v := strings.Split(s, "-")
			if len(v) != 2 {
				return fmt.Errorf("invalid random dummy proof delay value: %s", v)
			}

			lower, err := time.ParseDuration(v[0])
			if err != nil {
				return fmt.Errorf("invalid random dummy proof delay value: %s, err: %w", v, err)
			}
			upper, err := time.ParseDuration(v[1])
			if err != nil {
				return fmt.Errorf("invalid random dummy proof delay value: %s, err: %w", v, err)
			}
			if lower > upper {
				return fmt.Errorf("invalid random dummy proof delay value (lower > upper): %s", v)
			}

			if upper != time.Duration(0) {
				proverConf.RandomDummyProofDelayLowerBound = &lower
				proverConf.RandomDummyProofDelayUpperBound = &upper
			}
			return nil
		},
	}
	OracleProverFlag = &cli.BoolFlag{
		Name:     "oracleProver",
		Usage:    "Set whether prover should use oracle prover or not",
		Category: proverCategory,
		Action: func(c *cli.Context, v bool) error {
			proverConf.OracleProver = v
			return nil
		},
	}
	OracleProverPrivateKeyFlag = &cli.StringFlag{
		Name:     "oracleProverPrivateKey",
		Usage:    "Private key of oracle prover",
		Category: proverCategory,
		Action: func(c *cli.Context, v string) error {
			k, err := crypto.ToECDSA(common.Hex2Bytes(v))
			if err != nil {
				return fmt.Errorf("invalid oracle private key: %w", err)
			}
			proverConf.OracleProverPrivateKey = k
			return nil
		},
	}
	OracleProofSubmissionDelayFlag = &cli.DurationFlag{
		Name:        "oracleProofSubmissionDelay",
		Usage:       "Oracle proof submission delay in `duration`",
		Value:       0,
		Category:    proverCategory,
		Destination: &proverConf.OracleProofSubmissionDelay,
		Action: func(c *cli.Context, v time.Duration) error {
			proverConf.OracleProofSubmissionDelay = v
			return nil
		},
	}
	ProofSubmissionMaxRetryFlag = &cli.Uint64Flag{
		Name:        "proofSubmissionMaxRetry",
		Usage:       "Max retry counts for proof submission",
		Value:       0,
		Category:    proverCategory,
		Destination: &proverConf.ProofSubmissionMaxRetry,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.ProofSubmissionMaxRetry = v
			return nil
		},
	}
	GraffitiFlag = &cli.StringFlag{
		Name:        "graffiti",
		Usage:       "When string is passed, adds additional graffiti info to proof evidence",
		Category:    proverCategory,
		Value:       "",
		Destination: &proverConf.Graffiti,
		Action: func(c *cli.Context, v string) error {
			proverConf.Graffiti = v
			return nil
		},
	}
	CheckProofWindowExpiredIntervalFlag = &cli.DurationFlag{
		Name:        "prover.checkProofWindowExpiredInterval",
		Usage:       "Interval in `duration` to check for expired proof windows from other provers",
		Category:    proverCategory,
		Value:       15 * time.Second,
		Destination: &proverConf.CheckProofWindowExpiredInterval,
		Action: func(c *cli.Context, v time.Duration) error {
			proverConf.CheckProofWindowExpiredInterval = v
			return nil
		},
	}
	ProveUnassignedBlocksFlag = &cli.BoolFlag{
		Name:        "prover.proveUnassignedBlocks",
		Usage:       "Whether you want to prove unassigned blocks, or only work on assigned proofs",
		Category:    proverCategory,
		Value:       false,
		Destination: &proverConf.ProveUnassignedBlocks,
		Action: func(c *cli.Context, v bool) error {
			proverConf.ProveUnassignedBlocks = v
			return nil
		},
	}
	ProveBlockTxGasLimitFlag = &cli.Uint64Flag{
		Name:     "prover.proveBlockTxGasLimit",
		Usage:    "Gas limit will be used for TaikoL1.proveBlock transactions",
		Category: proverCategory,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.ProveBlockGasLimit = &v
			return nil
		},
	}
	ProverHTTPServerPortFlag = &cli.Uint64Flag{
		Name:        "prover.httpServerPort",
		Usage:       "Port to expose for http server",
		Category:    proverCategory,
		Value:       9876,
		Destination: &proverConf.HTTPServerPort,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.HTTPServerPort = v
			return nil
		},
	}
	ProverCapacityFlag = &cli.Uint64Flag{
		Name:     "prover.capacity",
		Usage:    "Capacity of prover, required if oracleProver is false",
		Category: proverCategory,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.Capacity = v
			return nil
		},
	}
	MaxExpiryFlag = &cli.DurationFlag{
		Name:        "prover.maxExpiry",
		Usage:       "maximum accepted expiry in `duration` for accepting proving a block",
		Value:       time.Hour,
		Category:    proverCategory,
		Destination: &proverConf.MaxExpiry,
		Action: func(c *cli.Context, v time.Duration) error {
			proverConf.MaxExpiry = v
			return nil
		},
	}
)

// All prover flags.
var proverFlags = MergeFlags(CommonFlags, []cli.Flag{
	L1HTTPEndpointFlag,
	L2WSEndpointFlag,
	L2HTTPEndpointFlag,
	ZkEvmRpcdEndpointFlag,
	ZkEvmRpcdParamsPathFlag,
	L1ProverPrivKeyFlag,
	MinProofFeeFlag,
	StartingBlockIDFlag,
	MaxConcurrentProvingJobsFlag,
	DummyFlag,
	RandomDummyProofDelayFlag,
	OracleProverFlag,
	OracleProverPrivateKeyFlag,
	OracleProofSubmissionDelayFlag,
	ProofSubmissionMaxRetryFlag,
	GraffitiFlag,
	CheckProofWindowExpiredIntervalFlag,
	ProveUnassignedBlocksFlag,
	ProveBlockTxGasLimitFlag,
	ProverHTTPServerPortFlag,
	ProverCapacityFlag,
	MaxExpiryFlag,
})

func newProver(c *cli.Context) (*prover.Prover, error) {
	if err := proverConf.Validate(); err != nil {
		return nil, err
	}
	return prover.New(c.Context, proverConf)
}
