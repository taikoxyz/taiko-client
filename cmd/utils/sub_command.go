package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/taikochain/taiko-client/cmd/logger"
	"github.com/taikochain/taiko-client/metrics"
	"github.com/urfave/cli/v2"
)

type Subcommand interface {
	InitFromCli(context.Context, *cli.Context) error
	Name() string
	Start() error
	Close()
}

func SubcommandAction(app Subcommand) cli.ActionFunc {
	return func(c *cli.Context) error {
		logger.InitLogger(c)

		ctx, ctxClose := context.WithCancel(context.Background())
		defer func() {
			if ctx.Err() != nil {
				ctxClose()
			}
		}()

		if err := app.InitFromCli(ctx, c); err != nil {
			return err
		}

		log.Info("Starting Taiko client application", "name", app.Name())

		if err := app.Start(); err != nil {
			log.Error("Starting application error", "name", app.Name(), "error", err)
			return err
		}

		metrics.Serve(ctx, c)

		defer func() {
			ctxClose()
			app.Close()
			log.Info("Application stopped", "name", app.Name())
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
}
