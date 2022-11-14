package flags

import (
	"github.com/urfave/cli/v2"
)

// Required flags used by prover.
var (
	ZkEvmRpcdEndpoint = cli.StringFlag{
		Name:     "zkevmRpcdEndpoint",
		Usage:    "RPC endpoint of a ZKEVM RPCD service",
		Required: true,
		Category: proverCategory,
	}
	ZkEvmRpcdParamsPath = cli.StringFlag{
		Name:     "zkevmRpcdParamsPath",
		Usage:    "Path of ZKEVM parameters file to use",
		Required: true,
		Category: proverCategory,
	}
	L1ProverPrivKey = cli.StringFlag{
		Name: "l1.proverPrivKey",
		Usage: "Private key of L1 prover, " +
			"who will send TaikoL1.proveBlock / TaikoL1.proveBlockInvalid transactions",
		Required: true,
		Category: proverCategory,
	}
)

// Special flags for testing.
var (
	Dummy = cli.BoolFlag{
		Name:     "dummy",
		Usage:    "Produce dummy proofs",
		Value:    false,
		Category: proverCategory,
	}
	BatchSubmit = cli.BoolFlag{
		Name:     "batchSubmit",
		Usage:    "Batch submit proofs",
		Value:    false,
		Hidden:   true,
		Category: proverCategory,
	}
)

// All prover flags.
var ProverFlags = MergeFlags(CommonFlags, []cli.Flag{
	&ZkEvmRpcdEndpoint,
	&ZkEvmRpcdParamsPath,
	&L1ProverPrivKey,
	&Dummy,
	&BatchSubmit,
})
