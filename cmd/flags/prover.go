package flags

import (
	"time"

	"github.com/urfave/cli/v2"
)

// Required flags used by prover.
var (
	ZkEvmRpcdEndpoint = &cli.StringFlag{
		Name:     "zkevmRpcdEndpoint",
		Usage:    "RPC endpoint of a ZKEVM RPCD service",
		Required: true,
		Category: proverCategory,
	}
	ZkEvmRpcdParamsPath = &cli.StringFlag{
		Name:     "zkevmRpcdParamsPath",
		Usage:    "Path of ZKEVM parameters file to use",
		Required: true,
		Category: proverCategory,
	}
	L1ProverPrivKey = &cli.StringFlag{
		Name: "l1.proverPrivKey",
		Usage: "Private key of L1 prover, " +
			"who will send TaikoL1.proveBlock / TaikoL1.proveBlockInvalid transactions",
		Required: true,
		Category: proverCategory,
	}
	MinProofFee = &cli.StringFlag{
		Name:     "prover.minProofFee",
		Usage:    "Minimum accepted fee for accepting proving a block",
		Required: true,
		Category: proverCategory,
	}
)

// Optional flags used by prover.
var (
	StartingBlockID = &cli.Uint64Flag{
		Name:     "startingBlockID",
		Usage:    "If set, prover will start proving blocks from the block with this ID",
		Category: proverCategory,
	}
	MaxConcurrentProvingJobs = &cli.UintFlag{
		Name:     "maxConcurrentProvingJobs",
		Usage:    "Limits the number of concurrent proving blocks jobs",
		Value:    1,
		Category: proverCategory,
	}
	// Special flags for testing.
	Dummy = &cli.BoolFlag{
		Name:     "dummy",
		Usage:    "Produce dummy proofs, testing purposes only",
		Value:    false,
		Category: proverCategory,
	}
	RandomDummyProofDelay = &cli.StringFlag{
		Name: "randomDummyProofDelay",
		Usage: "Set the random dummy proof delay between the bounds using the format: " +
			"`lowerBound-upperBound` (e.g. `30m-1h`), testing purposes only",
		Category: proverCategory,
	}
	OracleProver = &cli.BoolFlag{
		Name:     "oracleProver",
		Usage:    "Set whether prover should use oracle prover or not",
		Category: proverCategory,
	}
	OracleProverPrivateKey = &cli.StringFlag{
		Name:     "oracleProverPrivateKey",
		Usage:    "Private key of oracle prover",
		Category: proverCategory,
	}
	OracleProofSubmissionDelay = &cli.Uint64Flag{
		Name:     "oracleProofSubmissionDelay",
		Usage:    "Oracle proof submission delay in seconds",
		Value:    0,
		Category: proverCategory,
	}
	ProofSubmissionMaxRetry = &cli.Uint64Flag{
		Name:     "proofSubmissionMaxRetry",
		Usage:    "Max retry counts for proof submission",
		Value:    0,
		Category: proverCategory,
	}
	Graffiti = &cli.StringFlag{
		Name:     "graffiti",
		Usage:    "When string is passed, adds additional graffiti info to proof evidence",
		Category: proverCategory,
		Value:    "",
	}
	CheckProofWindowExpiredInterval = &cli.Uint64Flag{
		Name:     "prover.checkProofWindowExpiredInterval",
		Usage:    "Interval in seconds to check for expired proof windows from other provers",
		Category: proverCategory,
		Value:    15,
	}
	ProveUnassignedBlocks = &cli.BoolFlag{
		Name:     "prover.proveUnassignedBlocks",
		Usage:    "Whether you want to prove unassigned blocks, or only work on assigned proofs",
		Category: proverCategory,
		Value:    false,
	}
	ProveBlockTxGasLimit = &cli.Uint64Flag{
		Name:     "prover.proveBlockTxGasLimit",
		Usage:    "Gas limit will be used for TaikoL1.proveBlock transactions",
		Category: proverCategory,
	}
	ProveBlockTxReplacementMultiplier = &cli.Uint64Flag{
		Name:     "proveBlockTxReplacementMultiplier",
		Value:    2,
		Usage:    "Gas tip multiplier when replacing a TaikoL1.proveBlock transaction with same nonce",
		Category: proverCategory,
	}
	ProveBlockMaxTxGasTipCap = &cli.Uint64Flag{
		Name:     "proveBlockMaxTxGasTipCap",
		Usage:    "Gas tip cap (in wei) for a TaikoL1.proveBlock transaction when doing the transaction replacement",
		Category: proverCategory,
	}
	ProverHTTPServerPort = &cli.Uint64Flag{
		Name:     "prover.httpServerPort",
		Usage:    "Port to expose for http server",
		Category: proverCategory,
		Value:    9876,
	}
	ProverCapacity = &cli.Uint64Flag{
		Name:     "prover.capacity",
		Usage:    "Capacity of prover, required if oracleProver is false",
		Category: proverCategory,
	}
	MaxExpiry = &cli.Uint64Flag{
		Name:     "prover.maxExpiry",
		Usage:    "maximum accepted expiry in seconds for accepting proving a block",
		Value:    uint64(time.Hour.Seconds()),
		Category: proverCategory,
	}
	TempCapacityExpiresAt = &cli.DurationFlag{
		Name:     "prover.tempCapacityExpiresAt",
		Usage:    "time in seconds temporary capacity lives for, format: 12s",
		Value:    12 * time.Second,
		Category: proverCategory,
	}
)

// All prover flags.
var ProverFlags = MergeFlags(CommonFlags, []cli.Flag{
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
	ProveBlockTxReplacementMultiplier,
	ProveBlockMaxTxGasTipCap,
	Graffiti,
	CheckProofWindowExpiredInterval,
	ProveUnassignedBlocks,
	ProveBlockTxGasLimit,
	ProverHTTPServerPort,
	ProverCapacity,
	MaxExpiry,
	TaikoTokenAddress,
	TempCapacityExpiresAt,
})
