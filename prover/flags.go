package prover

import (
	"github.com/urfave/cli/v2"
)

// Required flags.
var (
	L1NodeEndpointFlag = cli.StringFlag{
		Name:     "l1",
		Usage:    "RPC endpoint of a L1 ethereum node",
		Required: true,
	}
	L2NodeEndpointFlag = cli.StringFlag{
		Name:     "l2",
		Usage:    "RPC endpoint of a L2 ethereum node",
		Required: true,
	}
	TaikoL1AddressFlag = cli.StringFlag{
		Name:     "taikoL1",
		Usage:    "TaikoL1 contract address",
		Required: true,
	}
	TaikoL2AddressFlag = cli.StringFlag{
		Name:     "taikoL2",
		Usage:    "TaikoL2 contract address",
		Required: true,
	}
	ZkEvmRpcdEndpointFlag = cli.StringFlag{
		Name:     "zkevmRpcdEndpoint",
		Usage:    "RPC endpoint of a ZKEVM RPCD service",
		Required: true,
	}
	ZkEvmRpcdParamsPathFlag = cli.StringFlag{
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
	DummyFlag = cli.BoolFlag{
		Name:  "dummy",
		Usage: "Produce dummy proofs",
	}
	BatchSubmitFlag = cli.BoolFlag{
		Name:  "batchSubmit",
		Usage: "Batch submit proofs",
	}
)

// All flags.
var Flags = []cli.Flag{
	&L1NodeEndpointFlag,
	&L2NodeEndpointFlag,
	&TaikoL1AddressFlag,
	&TaikoL2AddressFlag,
	&ZkEvmRpcdEndpointFlag,
	&ZkEvmRpcdParamsPathFlag,
	&L1ProverPrivKeyFlag,
	&DummyFlag,
	&BatchSubmitFlag,
}
