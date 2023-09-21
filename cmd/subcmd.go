package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/node"
)

var s node.Service

func startServer(c *cli.Context) error {
	ctx, ctxClose := context.WithCancel(context.Background())
	defer func() { ctxClose() }()

	log.Info("Starting Taiko client application", "name", s.Name())

	if err := s.Start(); err != nil {
		log.Error("Starting application error", "name", s.Name(), "error", err)
		return err
	}

	if err := metrics.Serve(ctx, metricConf); err != nil {
		log.Error("Starting metrics server error", "error", err)
		return err
	}

	defer func() {
		ctxClose()
		s.Close(ctx)
		log.Info("Application stopped", "name", s.Name())
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
