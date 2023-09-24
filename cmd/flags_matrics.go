package main

import (
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/urfave/cli/v2"
)

var metricConf = &metrics.Config{}

var (
	MetricsEnabledFlag = &cli.BoolFlag{
		Name:        "metrics",
		Usage:       "Enable metrics collection and reporting",
		Category:    metricsCategory,
		Value:       false,
		Destination: &metricConf.Enabled,
		Action: func(c *cli.Context, v bool) error {
			metricConf.Enabled = v
			return nil
		},
	}
	MetricsAddrFlag = &cli.StringFlag{
		Name:        "metrics.addr",
		Usage:       "Metrics reporting server listening address",
		Category:    metricsCategory,
		Value:       "0.0.0.0:60660",
		Destination: &metricConf.Address,
		Action: func(c *cli.Context, v string) error {
			metricConf.Address = v
			return nil
		},
	}
)
