package testutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/log"
)

// variables need to be initialized
var (
	JwtSecretFile string
	monoPath      string
)

func initMonoPath() (err error) {
	switch {
	case os.Getenv("TAIKO_MONO") != "":
		monoPath = os.Getenv("TAIKO_MONO")
	default:
		monoPath, err = defaultMonoPath()
		if err != nil {
			return err
		}
	}
	log.Info("Init", "TAIKO_MONO", monoPath)
	return nil
}

func defaultMonoPath() (string, error) {
	p, err := defaultProjectPath()
	if err != nil {
		return "", err
	}
	return filepath.Abs(p + "/taiko-mono")
}

func defaultProjectPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	idx := strings.LastIndex(cwd, "taiko-client")
	if idx == -1 {
		return "", fmt.Errorf("not found taiko-client in %s", cwd)
	}
	return cwd[:idx], nil
}

func defaultJwtFilePath() (string, error) {
	p, err := defaultProjectPath()
	if err != nil {
		return "", err
	}
	return filepath.Abs(p + "taiko-client/testutils/testdata/jwt.hex")
}

func initJwtSecret() (err error) {
	switch {
	case os.Getenv("JWT_SECRET") != "":
		JwtSecretFile = os.Getenv("JWT_SECRET")
	default:
		JwtSecretFile, err = defaultJwtFilePath()
		if err != nil {
			return err
		}
	}
	log.Info("Init", "JWT_SECRET", JwtSecretFile)
	return nil
}
