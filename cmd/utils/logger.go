package utils

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/taikochain/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

func InitLogger(c *cli.Context) {
	var handler log.Handler
	if c.Bool(flags.LogJson.Name) {
		handler = log.LvlFilterHandler(
			log.Lvl(c.Int(flags.Verbosity.Name)),
			log.StreamHandler(os.Stdout, log.JSONFormat()),
		)
	} else {
		handler = log.LvlFilterHandler(
			log.Lvl(c.Int(flags.Verbosity.Name)),
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		)
	}

	log.Root().SetHandler(handler)
}
