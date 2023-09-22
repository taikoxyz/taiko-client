package main

import (
	"fmt"
	"time"

	"github.com/taikoxyz/taiko-client/driver"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/urfave/cli/v2"
)

const (
	driverCmd = "driver"
)

var driverConf = &driver.Config{}

// Flags used by driver.
var (
	L2AuthEndpoint = &cli.StringFlag{
		Name:     "l2.auth",
		Usage:    "Authenticated HTTP RPC endpoint of a L2 taiko-geth execution engine",
		Required: true,
		Category: driverCategory,
		Action: func(c *cli.Context, v string) error {
			driverConf.L2EngineEndpoint = v
			endpointConf.L2EngineEndpoint = v
			return nil
		},
	}
	JWTSecret = &cli.StringFlag{
		Name:     "jwtSecret",
		Usage:    "Path to a JWT secret to use for authenticated RPC endpoints",
		Required: true,
		Category: driverCategory,
		Action: func(c *cli.Context, v string) error {
			jwtSecret, err := jwt.ParseSecretFromFile(v)
			if err != nil {
				return err
			}
			driverConf.JwtSecret = string(jwtSecret)
			endpointConf.JwtSecret = string(jwtSecret)
			return nil
		},
	}
)

// Optional flags used by driver.
var (
	P2PSyncVerifiedBlocks = &cli.BoolFlag{
		Name: "p2p.syncVerifiedBlocks",
		Usage: "Try P2P syncing verified blocks between L2 execution engines, " +
			"will be helpful to bring a new node online quickly",
		Value:    false,
		Category: driverCategory,
		Action: func(c *cli.Context, v bool) error {
			driverConf.P2PSyncVerifiedBlocks = v
			return nil
		},
	}
	P2PSyncTimeout = &cli.DurationFlag{
		Name: "p2p.syncTimeout",
		Usage: "P2P syncing timeout in `duration`, if no sync progress is made within this time span, " +
			"driver will stop the P2P sync and insert all remaining L2 blocks one by one",
		Value:    1800,
		Category: driverCategory,
		Action: func(c *cli.Context, v time.Duration) error {
			driverConf.P2PSyncTimeout = v
			return nil
		},
	}
	CheckPointSyncUrl = &cli.StringFlag{
		Name:     "p2p.checkPointSyncUrl",
		Usage:    "HTTP RPC endpoint of another synced L2 execution engine node",
		Category: driverCategory,
		Action: func(ctx *cli.Context, s string) error {
			driverConf.L2CheckPoint = s // 可能没用
			endpointConf.L2CheckPoint = s
			return nil
		},
	}
)

// All driver flags.
var driverFlags = MergeFlags(CommonFlags, []cli.Flag{
	L2WSEndpoint,
	L2AuthEndpoint,
	JWTSecret,
	P2PSyncVerifiedBlocks,
	P2PSyncTimeout,
	CheckPointSyncUrl,
})

func configDriver(c *cli.Context, ep *rpc.Client) error {
	if err := driverConf.Check(); err != nil {
		return err
	}
	peers, err := ep.L2.PeerCount(c.Context)
	if err != nil {
		return err
	}
	if driverConf.P2PSyncVerifiedBlocks && peers == 0 {
		fmt.Printf("P2P syncing verified blocks enabled, but no connected peer found in L2 execution engine")
	}
	d, err := driver.New(c.Context, ep, driverConf)
	if err != nil {
		return err
	}
	exec = d
	return nil
}
