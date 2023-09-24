package main

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

type logConfig struct {
	Verbosity    int
	IsJsonFormat bool
}

var logConf = &logConfig{}

var (
	// Optional flags for logging which are used by all client softwares.
	VerbosityFlag = &cli.IntFlag{
		Name:        "verbosity",
		Usage:       "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail",
		Value:       3,
		Category:    loggingCategory,
		Destination: &logConf.Verbosity,
		Action: func(c *cli.Context, v int) error {
			logConf.Verbosity = v
			return nil
		},
	}
	LogJsonFlag = &cli.BoolFlag{
		Name:     "log.json",
		Usage:    "Format logs with JSON",
		Category: loggingCategory,
		Action: func(c *cli.Context, v bool) error {
			logConf.IsJsonFormat = v
			return nil
		},
	}
)

func initLog(conf *logConfig) {
	var handler log.Handler
	if conf.IsJsonFormat {
		handler = log.LvlFilterHandler(
			log.Lvl(conf.Verbosity),
			log.StreamHandler(os.Stdout, log.JSONFormat()),
		)
	} else {
		handler = log.LvlFilterHandler(
			log.Lvl(conf.Verbosity),
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		)
	}
	log.Root().SetHandler(handler)
}
