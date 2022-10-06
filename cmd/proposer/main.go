package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/taikochain/client-mono/proposer"
	"github.com/taikochain/taiko-client/log"
	"github.com/urfave/cli/v2"
)

func main() {
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlDebug,
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		),
	)
	app := cli.NewApp()
	app.Action = runProposer
	app.Authors = []*cli.Author{
		{
			Name:  "Taiko Labs",
			Email: "info@taiko.xyz",
		},
	}
	app.Compiled = time.Now()
	app.Copyright = "Copyright 2022 Taiko Labs"
	app.Description = "The proposer service for Taiko protocol."
	app.EnableBashCompletion = true
	app.Flags = Flags
	app.Name = "proposer"
	app.Usage = "Taiko Proposer"
	app.Version = "0.0.1"
	if err := app.Run(os.Args); err != nil {
		log.Crit("Failed to start proposer", "err", err)
	}
}

// runProposer starts the main loop of taiko proposer.
func runProposer(ctx *cli.Context) error {
	if args := ctx.Args(); args.Len() > 0 {
		return fmt.Errorf("invalid command: %q", args.First())
	}

	log.Info("Starting proposer")

	cfg, err := NewConfigFromCliContext(ctx)
	if err != nil {
		return err
	}
	proposer, err := proposer.New(cfg)
	if err != nil {
		return err
	}

	proposer.Start()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}...)
	<-quitCh

	log.Info("Proposer stopped")

	return nil
}
