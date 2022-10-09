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
	}
	ZkEvmRpcdParamsPath = cli.StringFlag{
		Name:     "zkevmRpcdParamsPath",
		Usage:    "Path of ZKEVM parameters file to use",
		Required: true,
	}
	L1ProverPrivKeyFlag = cli.StringFlag{
		Name: "l1.proverPrivKey",
		Usage: "Private key of L1 prover, " +
			"who will send TaikoL1.proveBlock / TaikoL1.proveBlockInvalid transactions to the L1 node",
		Required: true,
	}
)

// Special flags for testing.
var (
	Dummy = cli.BoolFlag{
		Name:  "dummy",
		Usage: "Produce dummy proofs",
	}
	BatchSubmit = cli.BoolFlag{
		Name:  "batchSubmit",
		Usage: "Batch submit proofs",
	}
)

// All prover flags.
var ProverFlags = MergeFlags(CommonFlags, []cli.Flag{
	&ZkEvmRpcdEndpoint,
	&ZkEvmRpcdParamsPath,
	&L1ProverPrivKeyFlag,
	&Dummy,
	&BatchSubmit,
})
