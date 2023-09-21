package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/prover"
	"github.com/urfave/cli/v2"
)

const proverCmd = "prover"

var proverConf = &prover.Config{}

// Required flags used by prover.
var (
	ZkEvmRpcdEndpoint = &cli.StringFlag{
		Name:     "zkevmRpcdEndpoint",
		Usage:    "RPC endpoint of a ZKEVM RPCD service",
		Required: true,
		Category: proverCategory,
		Action: func(c *cli.Context, v string) error {
			proverConf.ZKEvmRpcdEndpoint = v
			return nil
		},
	}
	ZkEvmRpcdParamsPath = &cli.StringFlag{
		Name:     "zkevmRpcdParamsPath",
		Usage:    "Path of ZKEVM parameters file to use",
		Required: true,
		Category: proverCategory,
		Action: func(c *cli.Context, v string) error {
			proverConf.ZkEvmRpcdParamsPath = v
			return nil
		},
	}
	L1ProverPrivKey = &cli.StringFlag{
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
	MinProofFee = &cli.StringFlag{
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
	StartingBlockID = &cli.Uint64Flag{
		Name:     "startingBlockID",
		Usage:    "If set, prover will start proving blocks from the block with this ID",
		Category: proverCategory,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.StartingBlockID = new(big.Int).SetUint64(v)
			return nil
		},
	}
	MaxConcurrentProvingJobs = &cli.UintFlag{
		Name:     "maxConcurrentProvingJobs",
		Usage:    "Limits the number of concurrent proving blocks jobs",
		Value:    1,
		Category: proverCategory,
		Action: func(c *cli.Context, v uint) error {
			proverConf.MaxConcurrentProvingJobs = v
			return nil
		},
	}
	// Special flags for testing.
	Dummy = &cli.BoolFlag{
		Name:     "dummy",
		Usage:    "Produce dummy proofs, testing purposes only",
		Value:    false,
		Category: proverCategory,
		Action: func(c *cli.Context, v bool) error {
			proverConf.Dummy = v
			return nil
		},
	}
	RandomDummyProofDelay = &cli.StringFlag{
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
	OracleProver = &cli.BoolFlag{
		Name:     "oracleProver",
		Usage:    "Set whether prover should use oracle prover or not",
		Category: proverCategory,
		Action: func(c *cli.Context, v bool) error {
			proverConf.OracleProver = v
			return nil
		},
	}
	OracleProverPrivateKey = &cli.StringFlag{
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
	OracleProofSubmissionDelay = &cli.DurationFlag{
		Name:     "oracleProofSubmissionDelay",
		Usage:    "Oracle proof submission delay in `duration`",
		Value:    0,
		Category: proverCategory,
		Action: func(c *cli.Context, v time.Duration) error {
			proverConf.OracleProofSubmissionDelay = v
			return nil
		},
	}
	ProofSubmissionMaxRetry = &cli.Uint64Flag{
		Name:     "proofSubmissionMaxRetry",
		Usage:    "Max retry counts for proof submission",
		Value:    0,
		Category: proverCategory,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.ProofSubmissionMaxRetry = v
			return nil
		},
	}
	Graffiti = &cli.StringFlag{
		Name:     "graffiti",
		Usage:    "When string is passed, adds additional graffiti info to proof evidence",
		Category: proverCategory,
		Value:    "",
		Action: func(c *cli.Context, v string) error {
			proverConf.Graffiti = v
			return nil
		},
	}
	CheckProofWindowExpiredInterval = &cli.DurationFlag{
		Name:     "prover.checkProofWindowExpiredInterval",
		Usage:    "Interval in `duration` to check for expired proof windows from other provers",
		Category: proverCategory,
		Value:    15 * time.Second,
		Action: func(c *cli.Context, v time.Duration) error {
			proverConf.CheckProofWindowExpiredInterval = v
			return nil
		},
	}
	ProveUnassignedBlocks = &cli.BoolFlag{
		Name:     "prover.proveUnassignedBlocks",
		Usage:    "Whether you want to prove unassigned blocks, or only work on assigned proofs",
		Category: proverCategory,
		Value:    false,
		Action: func(c *cli.Context, v bool) error {
			proverConf.ProveUnassignedBlocks = v
			return nil
		},
	}
	ProveBlockTxGasLimit = &cli.Uint64Flag{
		Name:     "prover.proveBlockTxGasLimit",
		Usage:    "Gas limit will be used for TaikoL1.proveBlock transactions",
		Category: proverCategory,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.ProveBlockGasLimit = &v
			return nil
		},
	}
	ProverHTTPServerPort = &cli.Uint64Flag{
		Name:     "prover.httpServerPort",
		Usage:    "Port to expose for http server",
		Category: proverCategory,
		Value:    9876,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.HTTPServerPort = v
			return nil
		},
	}
	ProverCapacity = &cli.Uint64Flag{
		Name:     "prover.capacity",
		Usage:    "Capacity of prover, required if oracleProver is false",
		Category: proverCategory,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.Capacity = v
			return nil
		},
	}
	MaxExpiry = &cli.DurationFlag{
		Name:     "prover.maxExpiry",
		Usage:    "maximum accepted expiry in `duration` for accepting proving a block",
		Value:    time.Hour,
		Category: proverCategory,
		Action: func(c *cli.Context, v time.Duration) error {
			proverConf.MaxExpiry = v
			return nil
		},
	}
)

// All prover flags.
var proverFlags = MergeFlags(CommonFlags, []cli.Flag{
	L1HTTPEndpoint,
	L2WSEndpoint,
	L2HTTPEndpoint,
	ZkEvmRpcdEndpoint,
	ZkEvmRpcdParamsPath,
	L1ProverPrivKey,
	MinProofFee,
	StartingBlockID,
	MaxConcurrentProvingJobs,
	Dummy,
	RandomDummyProofDelay,
	OracleProver,
	OracleProverPrivateKey,
	OracleProofSubmissionDelay,
	ProofSubmissionMaxRetry,
	Graffiti,
	CheckProofWindowExpiredInterval,
	ProveUnassignedBlocks,
	ProveBlockTxGasLimit,
	ProverHTTPServerPort,
	ProverCapacity,
	MaxExpiry,
})

func prepareProver(c *cli.Context, ep *rpc.Client) (p *prover.Prover, err error) {
	if err := proverConf.Check(); err != nil {
		return nil, err
	}
	return prover.New(c.Context, ep, proverConf)
}
