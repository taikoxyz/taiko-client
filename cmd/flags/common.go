package flags

import (
	"github.com/urfave/cli/v2"
)

var (
	commonCategory   = "COMMON"
	loggingCategory  = "LOGGING"
	driverCategory   = "DRIVER"
	proposerCategory = "PROPOSER"
	proverCategory   = "PROVER"
)

// Required flags used by all client softwares.
var (
	L1NodeEndpoint = cli.StringFlag{
		Name:     "l1",
		Usage:    "Websocket RPC endpoint of a L1 ethereum node",
		Required: true,
		Category: commonCategory,
	}
	L2NodeEndpoint = cli.StringFlag{
		Name:     "l2",
		Usage:    "Websocket RPC endpoint of a L2 taiko-geth node",
		Required: true,
		Category: commonCategory,
	}
	TaikoL1Address = cli.StringFlag{
		Name:     "taikoL1",
		Usage:    "TaikoL1 contract address",
		Required: true,
		Category: commonCategory,
	}
	TaikoL2Address = cli.StringFlag{
		Name:     "taikoL2",
		Usage:    "TaikoL2 contract address",
		Required: true,
		Category: commonCategory,
	}
	// Optional flags used by all client softwares.
	// Logging
	Verbosity = &cli.IntFlag{
		Name:     "verbosity",
		Usage:    "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail",
		Value:    3,
		Category: loggingCategory,
	}
	LogJson = &cli.BoolFlag{
		Name:     "log.json",
		Usage:    "Format logs with JSON",
		Category: loggingCategory,
	}
	// Metrics
	MetricsEnabledFlag = cli.BoolFlag{
		Name:  "metrics",
		Usage: "Enable metrics collection and reporting",
		Value: false,
	}
	MetricsAddrFlag = cli.StringFlag{
		Name:  "metrics.addr",
		Usage: "Metrics reporting server listening address",
		Value: "0.0.0.0",
	}
	MetricsPortFlag = cli.IntFlag{
		Name:  "metrics.port",
		Usage: "Metrics reporting server listening port",
		Value: 6060,
	}
)

// All common flags.
var CommonFlags = []cli.Flag{
	// Required
	&L1NodeEndpoint,
	&L2NodeEndpoint,
	&TaikoL1Address,
	&TaikoL2Address,
	// Optional
	Verbosity,
	LogJson,
}

// MergeFlags merges the given flag slices.
func MergeFlags(groups ...[]cli.Flag) []cli.Flag {
	var merged []cli.Flag
	for _, group := range groups {
		merged = append(merged, group...)
	}
	return merged
}
