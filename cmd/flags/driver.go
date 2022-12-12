package flags

import (
	"github.com/urfave/cli/v2"
)

// Flags used by driver.
var (
	L2AuthEndpoint = cli.StringFlag{
		Name:     "l2.auth",
		Usage:    "Authenticated HTTP RPC endpoint of a L2 taiko-geth execution engine",
		Required: true,
		Category: driverCategory,
	}
	ThrowawayBlocksBuilderPrivKey = cli.StringFlag{
		Name: "l2.throwawayBlockBuilderPrivKey",
		Usage: "Private key of the L2 throwaway blocks builder," +
			"who will be the suggested fee recipient of L2 throwaway blocks",
		Required: true,
		Category: driverCategory,
	}
	JWTSecret = cli.StringFlag{
		Name:     "jwtSecret",
		Usage:    "Path to a JWT secret to use for authenticated RPC endpoints",
		Required: true,
		Category: driverCategory,
	}
)

// Optional flags used by driver.
var (
	P2PSyncVerifiedBlocks = cli.BoolFlag{
		Name: "p2p.syncVerifiedBlocks",
		Usage: "Try P2P syncing verified blocks between L2 execution engines, " +
			"will be helpful to bring a new node online quickly",
		Value:    false,
		Category: driverCategory,
	}
	P2PSyncTimeout = cli.UintFlag{
		Name: "p2p.syncTimeout",
		Usage: "P2P syncing timeout in seconds, if no sync progress is made within this time span, " +
			"driver will stop the P2P sync and insert all remaining L2 blocks one by one",
		Value:    120,
		Category: driverCategory,
	}
)

// All driver flags.
var DriverFlags = MergeFlags(CommonFlags, []cli.Flag{
	&L2AuthEndpoint,
	&ThrowawayBlocksBuilderPrivKey,
	&JWTSecret,
	&P2PSyncVerifiedBlocks,
	&P2PSyncTimeout,
})
