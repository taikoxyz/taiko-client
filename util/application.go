package util

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

type SubcommandApp interface {
	Name() string
	Start() error
	Close()
}

// MergeFlags merges the given flag slices.
func MergeFlags(groups ...[]cli.Flag) []cli.Flag {
	var ret []cli.Flag
	for _, group := range groups {
		ret = append(ret, group...)
	}
	return ret
}

func RunSubcommand(app SubcommandApp) error {
	log.Info("Starting Taiko client application", "name", app.Name())

	if err := app.Start(); err != nil {
		log.Error("Starting application error", "name", app.Name(), "error", err)
		return err
	}
	defer app.Close()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}...)
	<-quitCh

	log.Info("Application stopped", "name", app.Name())

	return nil
}
