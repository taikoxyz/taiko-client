package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/taikochain/client-mono/prover"
	"github.com/taikochain/taiko-client/log"
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
	app.Flags = prover.Flags
	app.Name = "prover"
	app.Usage = "Taiko Prover"
	app.Description = "The prover service for Taiko bindings."
	app.Action = RunProver

	if err := app.Run(os.Args); err != nil {
		log.Crit("Failed to start prover", "error", err)
	}
}

// RunProver starts the main loop of taiko prover.
func RunProver(c *cli.Context) error {
	log.Info("Starting prover")

	cfg, err := prover.NewConfigFromCliContext(c)
	if err != nil {
		log.Error("Unable to parse prover configurations", "error", err)
		return err
	}

	prover, err := prover.New(context.Background(), cfg)
	if err != nil {
		log.Error("Initializes prover error", "error", err)
		return err
	}

	prover.Start()
	defer prover.Close()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}...)
	<-quitCh

	return nil
}
