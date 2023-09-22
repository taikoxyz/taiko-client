package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-client/metrics"
)

// subCmd is the interface for the sub command.
type subCmd interface {
	Name() string
	Start() error
	Close(context.Context)
}

func startSubCmd(c *cli.Context) error {
	ctx, ctxClose := context.WithCancel(c.Context)
	defer func() { ctxClose() }()

	initLogger(logConf)
	cmd, err := cmdFromContext(c)
	if err != nil {
		return err
	}
	log.Info("Starting Taiko client application", "name", cmd.Name())

	if err := cmd.Start(); err != nil {
		log.Error("Starting application error", "name", cmd.Name(), "error", err)
		return err
	}

	if err := metrics.Serve(ctx, metricConf); err != nil {
		log.Error("Starting metrics server error", "error", err)
		return err
	}

	defer func() {
		ctxClose()
		cmd.Close(ctx)
		log.Info("Application stopped", "name", cmd.Name())
	}()

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

func cmdFromContext(c *cli.Context) (subCmd, error) {
	switch c.Command.Name {
	case driverCmd:
		return newDriver(c)
	case proposerCmd:
		return newProposer(c)
	case proverCmd:
		return newProver(c)
	default:
		panic("Unknown command name")
	}
}
