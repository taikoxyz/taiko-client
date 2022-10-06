package driver

import (
	"github.com/urfave/cli/v2"
)

// Required flags
var (
	L1NodeEndpoint = cli.StringFlag{
		Name:     "l1NodeEndpoint",
		Usage:    "RPC endpoint of a L1 ethereum node",
		Required: true,
	}
	L2NodeEndpoint = cli.StringFlag{
		Name:     "l2NodeEndpoint",
		Usage:    "RPC endpoint of a L2 taiko-client node",
		Required: true,
	}
	L2NodeAuthEndpoint = cli.StringFlag{
		Name:     "l2NodeAuthEndpoint",
		Usage:    "RPC endpoint of a L2 taiko-client node authenticated RPC APIs",
		Required: true,
	}
	TaikoL1Address = cli.StringFlag{
		Name:     "taikoL1Address",
		Usage:    "TaikoL1 contract address",
		Required: true,
	}
	TaikoL2Address = cli.StringFlag{
		Name:     "taikoL2Address",
		Usage:    "TaikoL2 contract address",
		Required: true,
	}
	ThrowawayBlocksBuilderPrivKey = cli.StringFlag{
		Name:     "throwawayBlockBuilderPrivKey",
		Usage:    "Private key of L2 throwaway blocks builder",
		Required: true,
	}
	JWTSecret = cli.StringFlag{
		Name:     "jwtSecret",
		Usage:    "Path to a JWT secret to use for authenticated RPC endpoints",
		Required: true,
	}

	Flags = []cli.Flag{
		&L1NodeEndpoint,
		&L2NodeEndpoint,
		&L2NodeAuthEndpoint,
		&TaikoL1Address,
		&TaikoL2Address,
		&ThrowawayBlocksBuilderPrivKey,
		&JWTSecret,
	}
)
