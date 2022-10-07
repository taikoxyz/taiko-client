package flags

import (
	"github.com/urfave/cli/v2"
)

// Flags used by all client softwares.
var (
	L1NodeEndpoint = cli.StringFlag{
		Name:     "l1",
		Usage:    "RPC endpoint of a L1 ethereum node",
		Required: true,
	}
	L2NodeEndpoint = cli.StringFlag{
		Name:     "l2",
		Usage:    "RPC endpoint of a L2 ethereum node",
		Required: true,
	}
	TaikoL1Address = cli.StringFlag{
		Name:     "taikoL1",
		Usage:    "TaikoL1 contract address",
		Required: true,
	}
	TaikoL2Address = cli.StringFlag{
		Name:     "taikoL2",
		Usage:    "TaikoL2 contract address",
		Required: true,
	}

	CommonFlags = []cli.Flag{
		&L1NodeEndpoint,
		&L2NodeEndpoint,
		&TaikoL1Address,
		&TaikoL2Address,
	}
)
