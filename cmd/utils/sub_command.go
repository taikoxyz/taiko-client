package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/cmd/logger"
	"github.com/taikoxyz/taiko-client/driver"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/node"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/prover"
	"github.com/urfave/cli/v2"
)

func StartServer(c *cli.Context) error {
	logger.InitLogger(c)

	s := getServer(c)
	ctx, ctxClose := context.WithCancel(context.Background())
	defer func() { ctxClose() }()

	if err := s.InitFromCli(ctx, c); err != nil {
		return err
	}

	log.Info("Starting Taiko client application", "name", s.Name())

	if err := s.Start(); err != nil {
		log.Error("Starting application error", "name", s.Name(), "error", err)
		return err
	}

	if err := metrics.Serve(ctx, c); err != nil {
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

// getServer returns an instance of node.Server based on the command name provided in the cli.Context parameter.
func getServer(ctx *cli.Context) node.Service {
	switch ctx.Command.Name {
	case "driver":
		return new(driver.Driver)
	case "proposer":
		return new(proposer.Proposer)
	case "prover":
		return new(prover.Prover)
	default:
		panic("Unknown command name")
	}
}
