package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/prysmaticlabs/prysm/v4/io/file"
	"github.com/urfave/cli/v2"
)

// Flags used by driver.
var (
	L2AuthEndpoint = &cli.StringFlag{
		Name:     "l2.auth",
		Usage:    "Authenticated HTTP RPC endpoint of a L2 taiko-geth execution engine",
		Required: true,
		Category: driverCategory,
		Action: func(c *cli.Context, v string) error {
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
			jwtSecret, err := parseSecretFromFile(v)
			if err != nil {
				return err
			}
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

// Taken from: https://github.com/prysmaticlabs/prysm/blob/v2.1.4/cmd/beacon-chain/execution/options.go#L43
// Parses a JWT secret from a file path. This secret is required when connecting to execution nodes
// over HTTP, and must be the same one used in Prysm and the execution node server Prysm is connecting to.
// The engine API specification here https://github.com/ethereum/execution-apis/blob/main/src/engine/authentication.md
// Explains how we should validate this secret and the format of the file a user can specify.
//
// The secret must be stored as a hex-encoded string within a file in the filesystem.
func parseSecretFromFile(jwtSecretFile string) ([]byte, error) {
	if jwtSecretFile == "" {
		return nil, nil
	}
	enc, err := file.ReadFileAsBytes(jwtSecretFile)
	if err != nil {
		return nil, err
	}
	strData := strings.TrimSpace(string(enc))
	if len(strData) == 0 {
		return nil, fmt.Errorf("provided JWT secret in file %s cannot be empty", jwtSecretFile)
	}
	secret, err := hex.DecodeString(strings.TrimPrefix(strData, "0x"))
	if err != nil {
		return nil, err
	}
	if len(secret) < 32 {
		return nil, errors.New("provided JWT secret should be a hex string of at least 32 bytes")
	}
	return secret, nil
}
