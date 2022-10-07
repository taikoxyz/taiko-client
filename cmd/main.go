package main

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/taikochain/client-mono/cmd/flags"
	"github.com/taikochain/client-mono/driver"
	"github.com/taikochain/client-mono/proposer"
	"github.com/taikochain/client-mono/prover"
	"github.com/urfave/cli/v2"
)

func main() {
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlInfo,
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		),
	)

	app := cli.NewApp()

	app.Version = "0.0.1"
	app.Name = "Taiko Clients"
	app.Description = "Entrypoint of Taiko Clients"
	app.Authors = []*cli.Author{{Name: "Taiko Labs", Email: "info@taiko.xyz"}}

	app.Commands = []*cli.Command{
		{
			Name:        "driver",
			Flags:       flags.DriverFlags,
			Description: "Taiko Driver software",
			Action:      driver.Action(),
		},
		{
			Name:        "proposer",
			Flags:       flags.ProposerFlags,
			Description: "Taiko Proposer software",
			Action:      proposer.Action(),
		},
		{
			Name:        "prover",
			Flags:       flags.ProverFlags,
			Description: "Taiko Prover software",
			Action:      prover.Action(),
		},
	}

	app.Action = func(ctx *cli.Context) error {
		log.Crit("Expected driver/proposer/prover subcommands")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Crit("Failed to start Taiko client", "error", err)
	}
}
