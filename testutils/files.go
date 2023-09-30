package testutils

import (
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/log"
)

const (
	stateFile      = "/Users/lsl/go/src/github/taikoxyz/taiko-client/testutils/testdata/state.json"
	defMonoPath    = "/Users/lsl/go/src/github/taikoxyz/taiko-mono"
	defJwtFilePath = "/Users/lsl/go/src/github/taikoxyz/taiko-client/integration_test/nodes/jwt.hex"
)

// variables need to be initialized
var (
	JwtSecretFile string
	monoPath      string
)

func initMonoPath() (err error) {
	path := defMonoPath
	if os.Getenv("TAIKO_MONO") != "" {
		path = os.Getenv("TAIKO_MONO")
	}
	monoPath, err = filepath.Abs(path)
	if err != nil {
		return err
	}
	log.Info("Init", "TAIKO_MONO", monoPath)
	return nil
}

func initJwtSecret() (err error) {
	path := defJwtFilePath
	if os.Getenv("JWT_SECRET") != "" {
		path = os.Getenv("JWT_SECRET")
	}
	JwtSecretFile, err = filepath.Abs(path)
	if err != nil {
		return err
	}
	log.Info("Init", "JWT_SECRET", JwtSecretFile)
	return nil
}
