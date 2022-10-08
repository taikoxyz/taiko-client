package flags

import (
	"github.com/taikochain/taiko-client/util"
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
		Name:     "l1ProverPrivKey",
		Usage:    "Private key for L1 prover",
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
var ProverFlags = util.MergeFlags(CommonFlags, []cli.Flag{
	&ZkEvmRpcdEndpoint,
	&ZkEvmRpcdParamsPath,
	&L1ProverPrivKeyFlag,
	&Dummy,
	&BatchSubmit,
})
