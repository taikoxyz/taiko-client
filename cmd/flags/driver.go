package flags

import (
	"time"

	"github.com/urfave/cli/v2"
)

// Define a function to create flags
func newStringFlag(name, usage string, required bool) *cli.StringFlag {
	return &cli.StringFlag{
		Name:     name,
		Usage:    usage,
		Required: required,
		Category: driverCategory,
	}
}

func newBoolFlag(name, usage string, value bool) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:     name,
		Usage:    usage,
		Value:    value,
		Category: driverCategory,
	}
}

func newDurationFlag(name, usage string, value time.Duration) *cli.DurationFlag {
	return &cli.DurationFlag{
		Name:     name,
		Usage:    usage,
		Value:    value,
		Category: driverCategory,
	}
}

// Flags used by driver.
var (
	L2AuthEndpoint = newStringFlag(
		"l2.auth",
		"Authenticated HTTP RPC endpoint of a L2 taiko-geth execution engine",
		true,
	)

	JWTSecret = newStringFlag(
		"jwtSecret",
		"Path to a JWT secret to use for authenticated RPC endpoints",
		true,
	)
)

// Optional flags used by driver.
var (
	P2PSyncVerifiedBlocks = newBoolFlag(
		"p2p.syncVerifiedBlocks",
		"Try P2P syncing verified blocks between L2 execution engines, will be helpful to bring a new node online quickly",
		false,
	)

	P2PSyncTimeout = newDurationFlag(
		"p2p.syncTimeout",
		"P2P syncing timeout, if no sync progress is made within this time span, driver will stop the P2P sync and insert all remaining L2 blocks one by one",
		1*time.Hour,
	)

	CheckPointSyncURL = newStringFlag(
		"p2p.checkPointSyncUrl",
		"HTTP RPC endpoint of another synced L2 execution engine node",
		false,
	)
)

// DriverFlags All driver flags.
var DriverFlags = MergeFlags(CommonFlags, []cli.Flag{
	L1BeaconEndpoint,
	L2WSEndpoint,
	L2AuthEndpoint,
	JWTSecret,
	P2PSyncVerifiedBlocks,
	P2PSyncTimeout,
	CheckPointSyncURL,
})
