package main

import (
	"fmt"
	"os"

	"github.com/golang/gddo/log"
	"github.com/taikoxyz/taiko-client/driver"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/prover"
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
			Action:      startServer,
		},
		{
			Name:        proposerCmd,
			Flags:       proposerFlags,
			Usage:       "Starts the proposer software",
			Description: "Taiko proposer software",
			Action:      startServer,
		},
		{
			Name:        proverCmd,
			Flags:       proverFlags,
			Usage:       "Starts the prover software",
			Description: "Taiko prover software",
			Action:      startServer,
		},
	}

	app.Before = func(c *cli.Context) (err error) {
		ctx := c.Context
		ep, err := rpc.NewClient(ctx, endpointConf)
		if err != nil {
			return err
		}
		switch c.Command.Name {
		case driverCmd:
			s, err = prepareDriver(c, ep)
		case proposerCmd:
			s = &proposer.Proposer{
				RPC: ep,
			}
		case proverCmd:
			s = &prover.Prover{
				RPC: ep,
			}
		default:
			panic("Unknown command name")
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func prepareDriver(c *cli.Context, ep *rpc.Client) (*driver.Driver, error) {
	if err := driverConf.Check(); err != nil {
		return nil, err
	}
	peers, err := ep.L2.PeerCount(c.Context)
	if err != nil {
		return nil, err
	}
	if driverConf.P2PSyncVerifiedBlocks && peers == 0 {
		log.Warn("P2P syncing verified blocks enabled, but no connected peer found in L2 execution engine")
	}
	d, err := driver.New(c.Context, ep, driverConf)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func prepareProposer(c *cli.Context, ep *rpc.Client) (*driver.Driver, error) {
	return nil, nil
}
