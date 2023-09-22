package main

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
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

var endpointConf = &rpc.ClientConfig{}

// Required endpoint flags which are used by all client softwares.
var (
	L1WSEndpoint = &cli.StringFlag{
		Name:     "l1.ws",
		Usage:    "Websocket RPC endpoint of a L1 ethereum node",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.L1Endpoint = v
			proverConf.L1WsEndpoint = v
			driverConf.L1Endpoint = v
			endpointConf.L1Endpoint = v
			return nil
		},
	}
	L2WSEndpoint = &cli.StringFlag{
		Name:     "l2.ws",
		Usage:    "Websocket RPC endpoint of a L2 taiko-geth execution engine",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proverConf.L2WsEndpoint = v
			driverConf.L2Endpoint = v
			endpointConf.L2Endpoint = v
			return nil
		},
	}
	L1HTTPEndpoint = &cli.StringFlag{
		Name:     "l1.http",
		Usage:    "HTTP RPC endpoint of a L1 ethereum node",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proverConf.L1HttpEndpoint = v
			endpointConf.L1Endpoint = v
			return nil
		},
	}
	L2HTTPEndpoint = &cli.StringFlag{
		Name:     "l2.http",
		Usage:    "HTTP RPC endpoint of a L2 taiko-geth execution engine",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.L2Endpoint = v
			proverConf.L2HttpEndpoint = v
			endpointConf.L2Endpoint = v
			return nil
		},
	}
	TaikoL1Address = &cli.StringFlag{
		Name:     "taikoL1",
		Usage:    "TaikoL1 contract address",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.TaikoL1Address = common.HexToAddress(v)
			proverConf.TaikoL1Address = common.HexToAddress(v)
			driverConf.TaikoL1Address = common.HexToAddress(v)
			endpointConf.TaikoL1Address = common.HexToAddress(v)
			return nil
		},
	}
	TaikoL2Address = &cli.StringFlag{
		Name:     "taikoL2",
		Usage:    "TaikoL2 contract address",
		Required: true,
		Category: commonCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.TaikoL2Address = common.HexToAddress(v)
			proverConf.TaikoL2Address = common.HexToAddress(v)
			driverConf.TaikoL2Address = common.HexToAddress(v)
			endpointConf.TaikoL2Address = common.HexToAddress(v)
			return nil
		},
	}
)

var (
	// Required  flags which are used by all client softwares.
	BackOffMaxRetrys = &cli.Uint64Flag{
		Name:     "backoff.maxRetrys",
		Usage:    "Max retry times when there is an error",
		Category: commonCategory,
		Value:    10,
		Action: func(c *cli.Context, v uint64) error {
			proverConf.BackOffMaxRetrys = v
			return nil
		},
	}
	BackOffRetryInterval = &cli.DurationFlag{
		Name:     "backoff.retryInterval",
		Usage:    "Retry interval in `duration` when there is an error",
		Category: commonCategory,
		Value:    12,
		Action: func(c *cli.Context, v time.Duration) error {
			proposerConf.BackOffRetryInterval = v
			proverConf.BackOffRetryInterval = v
			driverConf.BackOffRetryInterval = v
			endpointConf.RetryInterval = v
			return nil
		},
	}
	RPCTimeout = &cli.DurationFlag{
		Name:     "rpc.timeout",
		Usage:    "Timeout in `duration` for RPC calls",
		Category: commonCategory,
		Action: func(c *cli.Context, v time.Duration) error {
			proposerConf.RPCTimeout = &v
			proverConf.RPCTimeout = &v
			driverConf.RPCTimeout = &v
			endpointConf.Timeout = &v
			return nil
		},
	}
	WaitReceiptTimeout = &cli.DurationFlag{
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
	L1WSEndpoint,
	TaikoL1Address,
	TaikoL2Address,
	// Optional
	Verbosity,
	LogJson,
	MetricsEnabled,
	MetricsAddr,
	BackOffMaxRetrys,
	BackOffRetryInterval,
	RPCTimeout,
	WaitReceiptTimeout,
}

// MergeFlags merges the given flag slices.
func MergeFlags(groups ...[]cli.Flag) []cli.Flag {
	var merged []cli.Flag
	for _, group := range groups {
		merged = append(merged, group...)
	}
	return merged
}
