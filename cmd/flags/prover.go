package flags

import (
	"time"

	"github.com/urfave/cli/v2"
)

// Required flags used by prover.
var (
	L1ProverPrivKey = &cli.StringFlag{
		Name:     "l1.proverPrivKey",
		Usage:    "Private key of L1 prover, who will send TaikoL1.proveBlock transactions",
		Required: true,
		Category: proverCategory,
		EnvVars:  []string{"L1_PROVER_PRIV_KEY"},
	}
	ProverCapacity = &cli.Uint64Flag{
		Name:     "prover.capacity",
		Usage:    "Capacity of prover",
		Required: true,
		Category: proverCategory,
		EnvVars:  []string{"PROVER_CAPACITY"},
	}
)

// Optional flags used by prover.
var (
	RaikoHostEndpoint = &cli.StringFlag{
		Name:     "raiko.hostEndpoint",
		Usage:    "RPC endpoint of a Raiko host service",
		Category: proverCategory,
		EnvVars:  []string{"RAIKO_HOST_ENDPOINT"},
	}
	RaikoL1Endpoint = &cli.StringFlag{
		Name:     "raiko.l1Endpoint",
		Usage:    "L1 RPC endpoint which will be sent to the Raiko service",
		Category: proverCategory,
		EnvVars:  []string{"RAIKO_L1_ENDPOINT"},
	}
	RaikoL1BeaconEndpoint = &cli.StringFlag{
		Name:     "raiko.l1BeaconEndpoint",
		Usage:    "L1 beacon RPC endpoint which will be sent to the Raiko service",
		Category: proverCategory,
		EnvVars:  []string{"RAIKO_L1_BEACON_ENDPOINT"},
	}
	RaikoL2Endpoint = &cli.StringFlag{
		Name:     "raiko.l2Endpoint",
		Usage:    "L2 RPC endpoint which will be sent to the Raiko service",
		Category: proverCategory,
		EnvVars:  []string{"RAIKO_L2_ENDPOINT"},
	}
	StartingBlockID = &cli.Uint64Flag{
		Name:     "prover.startingBlockID",
		Usage:    "If set, prover will start proving blocks from the block with this ID",
		Category: proverCategory,
		EnvVars:  []string{"STARTING_BLOCK_ID"},
	}
	Graffiti = &cli.StringFlag{
		Name:     "prover.graffiti",
		Usage:    "When string is passed, adds additional graffiti info to proof evidence",
		Category: proverCategory,
		Value:    "",
		EnvVars:  []string{"GRAFFITI"},
	}
	// Proving strategy.
	ProveUnassignedBlocks = &cli.BoolFlag{
		Name:     "prover.proveUnassignedBlocks",
		Usage:    "Whether you want to prove unassigned blocks, or only work on assigned proofs",
		Category: proverCategory,
		Value:    false,
		EnvVars:  []string{"PROVE_UNASSIGNED_BLOCKS"},
	}
	MinEthBalance = &cli.Float64Flag{
		Name:     "prover.minEthBalance",
		Usage:    "Minimum ETH balance (in Ether) a prover wants to keep",
		Category: proverCategory,
		Value:    0,
		EnvVars:  []string{"MIN_ETH_BALANCE"},
	}
	MinTaikoTokenBalance = &cli.Float64Flag{
		Name:     "prover.minTaikoTokenBalance",
		Usage:    "Minimum Taiko token balance (in Ether) a prover wants to keep",
		Category: proverCategory,
		Value:    0,
		EnvVars:  []string{"MIN_TAIKO_TOKEN_BALANCE"},
	}
	// Tier fee related.
	MinOptimisticTierFee = &cli.Uint64Flag{
		Name:     "minTierFee.optimistic",
		Usage:    "Minimum accepted fee for generating an optimistic proof",
		Category: proverCategory,
		EnvVars:  []string{"MIN_OPTIMISTIC_TIER_FEE"},
	}
	MinSgxTierFee = &cli.Uint64Flag{
		Name:     "minTierFee.sgx",
		Usage:    "Minimum accepted fee for generating a SGX proof",
		Category: proverCategory,
		EnvVars:  []string{"MIN_SGX_TIER_FEE"},
	}
	MinSgxAndZkVMTierFee = &cli.Uint64Flag{
		Name:     "minTierFee.sgxAndZkvm",
		Usage:    "Minimum accepted fee for generating a SGX + zkVM proof",
		Category: proverCategory,
		EnvVars:  []string{"MIN_SGX_AND_ZKVM_TIER_FEE"},
	}
	// Guardian prover related.
	GuardianProver = &cli.StringFlag{
		Name:     "guardianProver",
		Usage:    "GuardianProver contract `address`",
		Category: proverCategory,
		EnvVars:  []string{"GUARDIAN_PROVER"},
	}
	GuardianProofSubmissionDelay = &cli.DurationFlag{
		Name:     "guardian.submissionDelay",
		Usage:    "Guardian proof submission delay",
		Value:    0 * time.Second,
		Category: proverCategory,
		EnvVars:  []string{"GUARDIAN_PROOF_SUBMISSION_DELAY"},
	}
	// Running mode
	ContesterMode = &cli.BoolFlag{
		Name:     "mode.contester",
		Usage:    "Whether you want to contest wrong transitions with higher tier proofs",
		Category: proverCategory,
		Value:    false,
		EnvVars:  []string{"CONTESTER_MODE"},
	}
	// HTTP server related.
	ProverHTTPServerPort = &cli.Uint64Flag{
		Name:     "http.port",
		Usage:    "Port to expose for http server",
		Category: proverCategory,
		Value:    9876,
		EnvVars:  []string{"PROVER_HTTP_SERVER_PORT"},
	}
	MaxExpiry = &cli.DurationFlag{
		Name:     "http.maxExpiry",
		Usage:    "Maximum accepted expiry in seconds for accepting proving a block",
		Value:    1 * time.Hour,
		Category: proverCategory,
		EnvVars:  []string{"MAX_EXPIRY"},
	}
	// Special flags for testing.
	Dummy = &cli.BoolFlag{
		Name:     "prover.dummy",
		Usage:    "Produce dummy proofs, testing purposes only",
		Value:    false,
		Category: proverCategory,
		EnvVars:  []string{"DUMMY"},
	}
	// Max slippage allowed
	MaxAcceptableBlockSlippage = &cli.Uint64Flag{
		Name:     "prover.blockSlippage",
		Usage:    "Maximum accepted slippage difference for blockID for accepting proving a block",
		Value:    1024,
		Category: proverCategory,
		EnvVars:  []string{"MAX_ACCEPTABLE_BLOCK_SLIPPAGE"},
	}
	// Max amount of L1 blocks that can pass before block is invalid
	MaxProposedIn = &cli.Uint64Flag{
		Name:     "prover.maxProposedIn",
		Usage:    "Maximum amount of L1 blocks that can pass before block can not be proposed. 0 means no limit.",
		Value:    0,
		Category: proverCategory,
		EnvVars:  []string{"MAX_PROPOSED_IN"},
	}
	Allowance = &cli.Float64Flag{
		Name:     "prover.allowance",
		Usage:    "Amount (in Ether) to approve AssignmentHook contract for TaikoToken usage",
		Category: proverCategory,
		EnvVars:  []string{"ALLOWANCE"},
	}
	GuardianProverHealthCheckServerEndpoint = &cli.StringFlag{
		Name:     "prover.guardianProverHealthCheckServerEndpoint",
		Usage:    "HTTP endpoint for main guardian prover health check server",
		Category: proverCategory,
		EnvVars:  []string{"GUARDIAN_PROVER_HEALTH_CHECK_SERVER_ENDPOINT"},
	}
	// Guardian prover specific flag
	EnableLivenessBondProof = &cli.BoolFlag{
		Name:     "prover.enableLivenessBondProof",
		Usage:    "Toggles whether the proof is a dummy proof or returns keccak256(RETURN_LIVENESS_BOND) as proof",
		Value:    false,
		Category: proverCategory,
		EnvVars:  []string{"ENABLE_LIVENESS_BOND_PROOF"},
	}
	L1NodeVersion = &cli.StringFlag{
		Name:     "prover.l1NodeVersion",
		Usage:    "Version or tag or the L1 Node Version used as an L1 RPC Url by this guardian prover",
		Category: proverCategory,
		EnvVars:  []string{"L1_NODE_VERSION"},
	}
	L2NodeVersion = &cli.StringFlag{
		Name:     "prover.l2NodeVersion",
		Usage:    "Version or tag or the L2 Node Version used as an L2 RPC Url by this guardian prover",
		Category: proverCategory,
		EnvVars:  []string{"L2_NODE_VERSION"},
	}
	// Confirmations specific flag
	BlockConfirmations = &cli.Uint64Flag{
		Name:     "prover.blockConfirmations",
		Usage:    "Confirmations to the latest L1 block before submitting a proof for a L2 block",
		Value:    6,
		Category: proverCategory,
		EnvVars:  []string{"BLOCK_CONFIRMATIONS"},
	}
)

// ProverFlags All prover flags.
var ProverFlags = MergeFlags(CommonFlags, []cli.Flag{
	L1HTTPEndpoint,
	L1BeaconEndpoint,
	L2WSEndpoint,
	L2HTTPEndpoint,
	RaikoHostEndpoint,
	RaikoL1Endpoint,
	RaikoL1BeaconEndpoint,
	RaikoL2Endpoint,
	L1ProverPrivKey,
	MinOptimisticTierFee,
	MinSgxTierFee,
	MinSgxAndZkVMTierFee,
	MinEthBalance,
	MinTaikoTokenBalance,
	StartingBlockID,
	Dummy,
	GuardianProver,
	GuardianProofSubmissionDelay,
	GuardianProverHealthCheckServerEndpoint,
	Graffiti,
	ProveUnassignedBlocks,
	ContesterMode,
	ProverHTTPServerPort,
	ProverCapacity,
	MaxExpiry,
	MaxProposedIn,
	TaikoTokenAddress,
	MaxAcceptableBlockSlippage,
	AssignmentHookAddress,
	Allowance,
	L1NodeVersion,
	L2NodeVersion,
	BlockConfirmations,
}, TxmgrFlags)
