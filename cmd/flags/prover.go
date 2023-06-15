package flags

import (
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
	OracleProver = &cli.BoolFlag{
		Name:     "oracleProver",
		Usage:    "Set whether prover should use oracle prover or not",
		Category: proverCategory,
	}
	SystemProver = &cli.BoolFlag{
		Name:     "systemProver",
		Usage:    "Set whether prover should use system prover or not",
		Category: proverCategory,
	}
	OracleProverPrivateKey = &cli.StringFlag{
		Name:     "oracleProverPrivateKey",
		Usage:    "Private key of oracle prover",
		Category: proverCategory,
	}
	SystemProverPrivateKey = &cli.StringFlag{
		Name:     "systemProverPrivateKey",
		Usage:    "Private key of system prover",
		Category: proverCategory,
	}
	Graffiti = &cli.StringFlag{
		Name:     "graffiti",
		Usage:    "When string is passed, adds additional graffiti info to proof evidence",
		Category: proverCategory,
		Value:    "",
	}
	BidStrategy = &cli.StringFlag{
		Name:     "bidStrategy",
		Usage:    "Which strategy to use for bidding on proposed blocks",
		Category: proverCategory,
	}
	MinimumBidFeePerGas = &cli.StringFlag{
		Name:     "bidMinFeePerGas",
		Usage:    "Minimum amount in wei per gas you are willing to bid if bidStategy is MinimumBidFeePerGas",
		Category: proverCategory,
	}
	BidDeposit = &cli.StringFlag{
		Name:     "bidDeposit",
		Usage:    "Deposit to use for bids",
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
	StartingBlockID,
	MaxConcurrentProvingJobs,
	Dummy,
	OracleProver,
	SystemProver,
	OracleProverPrivateKey,
	SystemProverPrivateKey,
	Graffiti,
	BidStrategy,
	MinimumBidFeePerGas,
	BidDeposit,
})
