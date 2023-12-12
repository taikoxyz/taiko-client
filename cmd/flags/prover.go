package flags

import (
	"time"

	"github.com/urfave/cli/v2"
)

// Required flags used by prover.
var (
	ZkEvmRpcdEndpoint = &cli.StringFlag{
		Name:     "zkevm.rpcdEndpoint",
		Usage:    "RPC endpoint of a ZKEVM RPCD service",
		Required: true,
		Category: proverCategory,
	}
	ZkEvmRpcdParamsPath = &cli.StringFlag{
		Name:     "zkevm.rpcdParamsPath",
		Usage:    "Path of ZKEVM parameters file to use",
		Required: true,
		Category: proverCategory,
	}
	RaikoHostEndpoint = &cli.StringFlag{
		Name:     "raiko.hostEndpoint",
		Usage:    "RPC endpoint of a ZKEVM RPCD service",
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
	ProverAssignmentHookAddress = &cli.StringFlag{
		Name:     "assignmentHookAddress",
		Usage:    "Address of the AssignmentHook contract",
		Required: true,
		Category: proverCategory,
	}
)

// Optional flags used by prover.
var (
	StartingBlockID = &cli.Uint64Flag{
		Name:     "prover.startingBlockID",
		Usage:    "If set, prover will start proving blocks from the block with this ID",
		Category: proverCategory,
	}
	Graffiti = &cli.StringFlag{
		Name:     "prover.graffiti",
		Usage:    "When string is passed, adds additional graffiti info to proof evidence",
		Category: proverCategory,
		Value:    "",
	}
	// Proving strategy.
	ProveUnassignedBlocks = &cli.BoolFlag{
		Name:     "prover.proveUnassignedBlocks",
		Usage:    "Whether you want to prove unassigned blocks, or only work on assigned proofs",
		Category: proverCategory,
		Value:    false,
	}
	// Tier fee related.
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
	MinSgxAndPseZkevmTierFee = &cli.Uint64Flag{
		Name:     "minTierFee.sgxAndPseZKEvm",
		Usage:    "Minimum accepted fee for generating a SGX + PSE zkEVM proof",
		Category: proverCategory,
	}
	// Guardian prover related.
	GuardianProver = &cli.StringFlag{
		Name:     "guardianProver",
		Usage:    "GuardianProver contract `address`",
		Category: proverCategory,
	}
	GuardianProverPrivateKey = &cli.StringFlag{
		Name:     "guardian.privateKey",
		Usage:    "Private key of guardian prover",
		Category: proverCategory,
	}
	GuardianProofSubmissionDelay = &cli.DurationFlag{
		Name:     "guardian.submissionDelay",
		Usage:    "Guardian proof submission delay",
		Value:    0 * time.Second,
		Category: proverCategory,
	}
	// Transaction related.
	ProofSubmissionMaxRetry = &cli.Uint64Flag{
		Name:     "tx.submissionMaxRetry",
		Usage:    "Max retry counts for proof submission",
		Value:    3,
		Category: proverCategory,
	}
	ProveBlockTxGasLimit = &cli.Uint64Flag{
		Name:     "tx.gasLimit",
		Usage:    "Gas limit will be used for TaikoL1.proveBlock transactions",
		Category: proverCategory,
	}
	ProveBlockMaxTxGasTipCap = &cli.Uint64Flag{
		Name:     "tx.maxGasTipCap",
		Usage:    "Gas tip cap (in wei) for a TaikoL1.proveBlock transaction when doing the transaction replacement",
		Category: proverCategory,
	}
	ProveBlockTxReplacementMultiplier = &cli.Uint64Flag{
		Name:     "tx.replacementMultiplier",
		Value:    2,
		Usage:    "Gas tip multiplier when replacing a TaikoL1.proveBlock transaction with same nonce",
		Category: proverCategory,
	}
	// Running mode
	ContesterMode = &cli.BoolFlag{
		Name:     "mode.contester",
		Usage:    "Whether you want to contest wrong transitions with higher tier proofs",
		Category: proverCategory,
		Value:    false,
	}
	// HTTP server related.
	ProverHTTPServerPort = &cli.Uint64Flag{
		Name:     "http.port",
		Usage:    "Port to expose for http server",
		Category: proverCategory,
		Value:    9876,
	}
	MaxExpiry = &cli.DurationFlag{
		Name:     "http.maxExpiry",
		Usage:    "Maximum accepted expiry in seconds for accepting proving a block",
		Value:    1 * time.Hour,
		Category: proverCategory,
	}
	// Special flags for testing.
	Dummy = &cli.BoolFlag{
		Name:     "prover.dummy",
		Usage:    "Produce dummy proofs, testing purposes only",
		Value:    false,
		Category: proverCategory,
	}
	// Max slippage allowed
	MaxAcceptableBlockSlippage = &cli.Uint64Flag{
		Name:     "prover.blockSlippage",
		Usage:    "Maximum accepted slippage difference for blockID for accepting proving a block",
		Value:    1024,
		Category: proverCategory,
	}
	// Max amount of L1 blocks that can pass before block is invalid
	MaxProposedIn = &cli.Uint64Flag{
		Name:     "prover.maxProposedIn",
		Usage:    "Maximum amount of L1 blocks that can pass before block can not be proposed. 0 means no limit.",
		Value:    0,
		Category: proverCategory,
	}
	// DB file location
	// Required for guardian prover only
	DatabasePath = &cli.StringFlag{
		Name:     "db.path",
		Usage:    "Database file location",
		Category: proverCategory,
	}
	DatabaseCacheSize = &cli.Uint64Flag{
		Name:     "db.cacheSize",
		Usage:    "Database cache size in megabytes",
		Value:    16,
		Category: proverCategory,
	}
	Allowance = &cli.StringFlag{
		Name:     "prover.allowance",
		Usage:    "Amount to approve TaikoL1 contract for TaikoToken usage",
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
	RaikoHostEndpoint,
	L1ProverPrivKey,
	MinOptimisticTierFee,
	MinSgxTierFee,
	MinPseZkevmTierFee,
	MinSgxAndPseZkevmTierFee,
	StartingBlockID,
	Dummy,
	GuardianProver,
	GuardianProverPrivateKey,
	GuardianProofSubmissionDelay,
	ProofSubmissionMaxRetry,
	ProveBlockTxReplacementMultiplier,
	ProveBlockMaxTxGasTipCap,
	Graffiti,
	ProveUnassignedBlocks,
	ContesterMode,
	ProveBlockTxGasLimit,
	ProverHTTPServerPort,
	ProverCapacity,
	MaxExpiry,
	MaxProposedIn,
	TaikoTokenAddress,
	MaxAcceptableBlockSlippage,
	DatabasePath,
	DatabaseCacheSize,
	ProverAssignmentHookAddress,
	Allowance,
})
