package utils

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
)

type SubcommandApp interface {
	Name() string
	Start() error
	Close()
}

func RunSubcommand(app SubcommandApp) error {
	log.Info("Starting Taiko client application", "name", app.Name())

	if err := app.Start(); err != nil {
		log.Error("Starting application error", "name", app.Name(), "error", err)
		return err
	}

	defer func() {
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
