package testutils

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/sync/errgroup"
)

func init() {
	// Don't change the following initialization order
	var g errgroup.Group
	g.Go(initLog)
	g.Go(initMonoPath)
	g.Go(initJwtSecret)
	g.Go(initTestAccount)
	g.Go(initProverAccount)
	if err := g.Wait(); err != nil {
		panic(err)
	}
	if err := startBaseContainer(context.Background()); err != nil {
		panic(fmt.Errorf("startBaseContainer: %v", err))
	}
}

func initLog() (err error) {
	level := log.LvlDebug
	if os.Getenv("LOG_LEVEL") != "" {
		level, err = log.LvlFromString(os.Getenv("LOG_LEVEL"))
		if err != nil {
			return fmt.Errorf("invalid log level: %v", os.Getenv("LOG_LEVEL"))
		}
	}
	log.Root().SetHandler(
		log.LvlFilterHandler(level, log.StreamHandler(os.Stdout, log.TerminalFormat(true))),
	)
	return nil
}
