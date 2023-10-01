package testutils

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/sync/errgroup"
)

func init() {
	if err := initLog(); err != nil {
		panic(err)
	}
	var g errgroup.Group
	g.Go(initMonoPath)
	g.Go(initJwtSecret)
	g.Go(initTestAccount)
	g.Go(initProverAccount)
	if err := g.Wait(); err != nil {
		panic(err)
	}
}

func initLog() (err error) {
	level := log.LvlInfo
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
