package main

import (
	"fmt"
	"os"

	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/version"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Name = "Taiko Clients"
	app.Usage = "The taiko client softwares command line interface"
	app.Copyright = "Copyright 2021-2022 Taiko Labs"
	app.Version = version.VersionWithCommit()
	app.Description = "Client softwares implementation in Golang for Taiko protocol"
	app.Authors = []*cli.Author{{Name: "Taiko Labs", Email: "info@taiko.xyz"}}
	app.EnableBashCompletion = true

	// All supported sub commands.
	app.Commands = []*cli.Command{
		{
			Name:        driverCmd,
			Flags:       driverFlags,
			Usage:       "Starts the driver software",
			Description: "Taiko driver software",
			Action:      startApp,
		},
		{
			Name:        proposerCmd,
			Flags:       proposerFlags,
			Usage:       "Starts the proposer software",
			Description: "Taiko proposer software",
			Action:      startApp,
		},
		{
			Name:        proverCmd,
			Flags:       proverFlags,
			Usage:       "Starts the prover software",
			Description: "Taiko prover software",
			Action:      startApp,
		},
	}

	app.Before = func(c *cli.Context) error {
		ctx := c.Context
		ep, err := rpc.NewClient(ctx, endpointConf)
		if err != nil {
			return err
		}
		switch c.Command.Name {
		case driverCmd:
			return configDriver(c, ep)
		case proposerCmd:
			return configProposer(c, ep)
		case proverCmd:
			return configProver(c, ep)
		default:
			panic("Unknown command name")
		}
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
