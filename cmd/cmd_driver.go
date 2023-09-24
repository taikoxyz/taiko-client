package main

import (
	"fmt"
	"time"

	"github.com/taikoxyz/taiko-client/driver"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/urfave/cli/v2"
)

const (
	driverCmd = "driver"
)

var driverConf = &driver.Config{}

// Required flags used by driver.
var (
	L2AuthEndpointFlag = &cli.StringFlag{
		Name:     "l2.auth",
		Usage:    "Authenticated HTTP RPC endpoint of a L2 taiko-geth execution engine",
		Required: true,
		Category: driverCategory,
		Action: func(c *cli.Context, v string) error {
			driverConf.L2EngineEndpoint = v
			return nil
		},
	}
	JWTSecretFlag = &cli.StringFlag{
		Name:     "jwtSecret",
		Usage:    "Path to a JWT secret to use for authenticated RPC endpoints",
		Required: true,
		Category: driverCategory,
		Action: func(c *cli.Context, v string) error {
			jwtSecret, err := jwt.ParseSecretFromFile(v)
			if err != nil {
				return fmt.Errorf("invalid JWT secret file: %w", err)
			}
			driverConf.JwtSecret = string(jwtSecret)
			return nil
		},
	}
)

// Optional flags used by driver.
var (
	P2PSyncVerifiedBlocksFlag = &cli.BoolFlag{
		Name: "p2p.syncVerifiedBlocks",
		Usage: "Try P2P syncing verified blocks between L2 execution engines, " +
			"will be helpful to bring a new node online quickly",
		Value:       false,
		Category:    driverCategory,
		Destination: &driverConf.P2PSyncVerifiedBlocks,
		Action: func(c *cli.Context, v bool) error {
			driverConf.P2PSyncVerifiedBlocks = v
			return nil
		},
	}
	P2PSyncTimeoutFlag = &cli.DurationFlag{
		Name: "p2p.syncTimeout",
		Usage: "P2P syncing timeout in `duration`, if no sync progress is made within this time span, " +
			"driver will stop the P2P sync and insert all remaining L2 blocks one by one",
		Value:       1800 * time.Second,
		Category:    driverCategory,
		Destination: &driverConf.P2PSyncTimeout,
		Action: func(c *cli.Context, v time.Duration) error {
			driverConf.P2PSyncTimeout = v
			return nil
		},
	}
	CheckPointSyncUrlFlag = &cli.StringFlag{
		Name:     "p2p.checkPointSyncUrl",
		Usage:    "HTTP RPC endpoint of another synced L2 execution engine node",
		Category: driverCategory,
		Action: func(ctx *cli.Context, s string) error {
			driverConf.L2CheckPoint = s
			return nil
		},
	}
)

// All driver flags.
var driverFlags = MergeFlags(CommonFlags, []cli.Flag{
	L2WSEndpointFlag,
	L2AuthEndpointFlag,
	JWTSecretFlag,
	P2PSyncVerifiedBlocksFlag,
	P2PSyncTimeoutFlag,
	CheckPointSyncUrlFlag,
})

func newDriver(c *cli.Context) (*driver.Driver, error) {
	if err := driverConf.Validate(c.Context); err != nil {
		return nil, err
	}
	return driver.New(c.Context, driverConf)
}
