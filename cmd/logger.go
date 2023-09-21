package main

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

// initLogger initializes the root logger with the command line flags.
func initLogger(c *cli.Context) {
	var handler log.Handler
	if c.Bool(LogJson.Name) {
		handler = log.LvlFilterHandler(
			log.Lvl(c.Int(Verbosity.Name)),
			log.StreamHandler(os.Stdout, log.JSONFormat()),
		)
	} else {
		handler = log.LvlFilterHandler(
			log.Lvl(c.Int(Verbosity.Name)),
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		)
	}

	log.Root().SetHandler(handler)
}
