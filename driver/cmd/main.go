package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/taikochain/client-mono/driver"
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
	app.Flags = driver.Flags
	app.Name = "taiko-driver"
	app.Usage = "Taiko L2 Driver"
	app.Description = "The driver for Taiko L2 node, sync blocks to L2 node by deriving them from L1"
	app.Action = RunDriver

	if err := app.Run(os.Args); err != nil {
		log.Crit("Failed to start driver", "error", err)
	}
}

// RunDriver starts the main loop of taiko driver.
func RunDriver(c *cli.Context) error {
	log.Info("Starting driver")

	cfg, err := driver.NewConfigFromCliContext(c)
	if err != nil {
		log.Error("Unable to parse driver configurations", "error", err)
		return err
	}

	driver, err := driver.New(context.Background(), cfg)
	if err != nil {
		log.Error("Initializes driver error", "error", err)
		return err
	}

	driver.Start()
	defer driver.Close()

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
