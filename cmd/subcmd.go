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

// application is the interface for the server that the node runs.
type application interface {
	Name() string
	Start() error
	Close(context.Context)
}

var exec application

func startApp(c *cli.Context) error {
	ctx, ctxClose := context.WithCancel(context.Background())
	defer func() { ctxClose() }()

	log.Info("Starting Taiko client application", "name", exec.Name())

	if err := exec.Start(); err != nil {
		log.Error("Starting application error", "name", exec.Name(), "error", err)
		return err
	}

	if err := metrics.Serve(ctx, metricConf); err != nil {
		log.Error("Starting metrics server error", "error", err)
		return err
	}

	defer func() {
		ctxClose()
		exec.Close(ctx)
		log.Info("Application stopped", "name", exec.Name())
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
