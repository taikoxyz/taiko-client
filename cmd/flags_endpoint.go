package main

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

var (
	commonCategory   = "COMMON"
	metricsCategory  = "METRICS"
	loggingCategory  = "LOGGING"
	driverCategory   = "DRIVER"
	proposerCategory = "PROPOSER"
	proverCategory   = "PROVER"
)

// Required endpoint flags which are used by all client softwares.
var (
	L1WSEndpointFlag = &cli.StringFlag{
		Name:     "l1.ws",
		Usage:    "Websocket RPC endpoint of a L1 ethereum node",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.L1Endpoint = v
			proverConf.L1WsEndpoint = v
			driverConf.L1Endpoint = v
			return nil
		},
	}
	L2WSEndpointFlag = &cli.StringFlag{
		Name:     "l2.ws",
		Usage:    "Websocket RPC endpoint of a L2 taiko-geth execution engine",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proverConf.L2WsEndpoint = v
			driverConf.L2Endpoint = v
			return nil
		},
	}
	L1HTTPEndpointFlag = &cli.StringFlag{
		Name:     "l1.http",
		Usage:    "HTTP RPC endpoint of a L1 ethereum node",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proverConf.L1HttpEndpoint = v
			return nil
		},
	}
	L2HTTPEndpointFlag = &cli.StringFlag{
		Name:     "l2.http",
		Usage:    "HTTP RPC endpoint of a L2 taiko-geth execution engine",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.L2Endpoint = v
			proverConf.L2HttpEndpoint = v
			return nil
		},
	}
	TaikoL1AddressFlag = &cli.StringFlag{
		Name:     "taikoL1",
		Usage:    "TaikoL1 contract address",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.TaikoL1Address = common.HexToAddress(v)
			proverConf.TaikoL1Address = common.HexToAddress(v)
			driverConf.TaikoL1Address = common.HexToAddress(v)
			return nil
		},
	}
	TaikoL2AddressFlag = &cli.StringFlag{
		Name:     "taikoL2",
		Usage:    "TaikoL2 contract address",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.TaikoL2Address = common.HexToAddress(v)
			proverConf.TaikoL2Address = common.HexToAddress(v)
			driverConf.TaikoL2Address = common.HexToAddress(v)
			return nil
		},
	}
)

var (
	// Required  flags which are used by all client softwares.
	BackOffMaxRetrysFlag = &cli.Uint64Flag{
		Name:     "backoff.maxRetrys",
		Usage:    "Max retry times when there is an error",
		Category: commonCategory,
		Value:    10,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.BackOffMaxRetrys = v
			return nil
		},
	}
	BackOffRetryIntervalFlag = &cli.DurationFlag{
		Name:     "backoff.retryInterval",
		Usage:    "Retry interval in `duration` when there is an error",
		Category: commonCategory,
		Value:    12,
		Action: func(c *cli.Context, v time.Duration) error {
			proposerConf.BackOffRetryInterval = v
			proverConf.BackOffRetryInterval = v
			driverConf.BackOffRetryInterval = v
			return nil
		},
	}
	RPCTimeoutFlag = &cli.DurationFlag{
		Name:     "rpc.timeout",
		Usage:    "Timeout in `duration` for RPC calls",
		Category: commonCategory,
		Action: func(c *cli.Context, v time.Duration) error {
			proposerConf.RPCTimeout = &v
			proverConf.RPCTimeout = &v
			driverConf.RPCTimeout = &v
			return nil
		},
	}
	WaitReceiptTimeoutFlag = &cli.DurationFlag{
		Name:     "rpc.waitReceiptTimeout",
		Usage:    "Timeout in `duration` for wait for receipts for RPC transactions",
		Category: commonCategory,
		Value:    60,
		Action: func(c *cli.Context, v time.Duration) error {
			proverConf.WaitReceiptTimeout = v
			proposerConf.WaitReceiptTimeout = v
			return nil
		},
	}
)

// All common flags.
var CommonFlags = []cli.Flag{
	// Required
	L1WSEndpointFlag,
	TaikoL1AddressFlag,
	TaikoL2AddressFlag,
	// Optional
	VerbosityFlag,
	LogJsonFlag,
	MetricsEnabledFlag,
	MetricsAddrFlag,
	BackOffMaxRetrysFlag,
	BackOffRetryIntervalFlag,
	RPCTimeoutFlag,
	WaitReceiptTimeoutFlag,
}

// MergeFlags merges the given flag slices.
func MergeFlags(groups ...[]cli.Flag) []cli.Flag {
	var merged []cli.Flag
	for _, group := range groups {
		merged = append(merged, group...)
	}
	return merged
}
