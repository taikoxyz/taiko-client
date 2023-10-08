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
		Name:     "l1.proverPrivKey",
		Usage:    "Private key of L1 prover, who will send TaikoL1.proveBlock transactions",
		Required: true,
		Category: proverCategory,
	}
	ProverCapacity = &cli.Uint64Flag{
		Name:     "prover.capacity",
		Usage:    "Capacity of prover",
		Required: true,
		Category: proverCategory,
	}
)

// Optional flags used by prover.
var (
	MinOptimisticTierFee = &cli.Uint64Flag{
		Name:     "minTierFee.optimistic",
		Usage:    "Minimum accepted fee for generating an optimistic proof",
		Category: proverCategory,
	}
	MinSgxTierFee = &cli.Uint64Flag{
		Name:     "minTierFee.sgx",
		Usage:    "Minimum accepted fee for generating a SGX proof",
		Category: proverCategory,
	}
	MinPseZkevmTierFee = &cli.Uint64Flag{
		Name:     "minTierFee.pseZKEvm",
		Usage:    "Minimum accepted fee for generating a PSE zkEVM proof",
		Category: proverCategory,
	}
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
	GuardianProver = &cli.BoolFlag{
		Name:     "guardianProver",
		Usage:    "Set whether prover should use guardian prover or not",
		Category: proverCategory,
	}
	GuardianProverPrivateKey = &cli.StringFlag{
		Name:     "guardianProverPrivateKey",
		Usage:    "Private key of guardian prover",
		Category: proverCategory,
	}
	GuardianProofSubmissionDelay = &cli.DurationFlag{
		Name:     "guardianProofSubmissionDelay",
		Usage:    "Guardian proof submission delay",
		Value:    0 * time.Second,
		Category: proverCategory,
	}
	ProofSubmissionMaxRetry = &cli.Uint64Flag{
		Name:     "proofSubmissionMaxRetry",
		Usage:    "Max retry counts for proof submission",
		Value:    3,
		Category: proverCategory,
	}
	Graffiti = &cli.StringFlag{
		Name:     "graffiti",
		Usage:    "When string is passed, adds additional graffiti info to proof evidence",
		Category: proverCategory,
		Value:    "",
	}
	ProveUnassignedBlocks = &cli.BoolFlag{
		Name:     "prover.proveUnassignedBlocks",
		Usage:    "Whether you want to prove unassigned blocks, or only work on assigned proofs",
		Category: proverCategory,
		Value:    false,
	}
	ContestControversialProofs = &cli.BoolFlag{
		Name:     "prover.contestControversialProofs",
		Usage:    "Whether you want to contest proofs with different L2 hashes with higher tier proofs",
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
	MaxExpiry = &cli.DurationFlag{
		Name:     "prover.maxExpiry",
		Usage:    "maximum accepted expiry in seconds for accepting proving a block",
		Value:    1 * time.Hour,
		Category: proverCategory,
	}
	TempCapacityExpiresAt = &cli.DurationFlag{
		Name:     "prover.tempCapacityExpiresAt",
		Usage:    "time in seconds temporary capacity lives for (format: 36s)",
		Value:    36 * time.Second,
		Category: proverCategory,
	}
	// Special flags for testing.
	Dummy = &cli.BoolFlag{
		Name:     "dummy",
		Usage:    "Produce dummy proofs, testing purposes only",
		Value:    false,
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
	MinOptimisticTierFee,
	MinSgxTierFee,
	MinPseZkevmTierFee,
	StartingBlockID,
	MaxConcurrentProvingJobs,
	Dummy,
	GuardianProver,
	GuardianProverPrivateKey,
	GuardianProofSubmissionDelay,
	ProofSubmissionMaxRetry,
	ProveBlockTxReplacementMultiplier,
	ProveBlockMaxTxGasTipCap,
	Graffiti,
	ProveUnassignedBlocks,
	ContestControversialProofs,
	ProveBlockTxGasLimit,
	ProverHTTPServerPort,
	ProverCapacity,
	MaxExpiry,
	TaikoTokenAddress,
	TempCapacityExpiresAt,
})
