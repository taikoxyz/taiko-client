package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v2"
)

var (
	l1WSEndpoint     = os.Getenv("L1_NODE_WS_ENDPOINT")
	l2WSEndpoint     = os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2AuthorEndpoint = os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT")
	rpcTimeout       = 5 * time.Second
)

type cmdSuit struct {
	suite.Suite
	app  *cli.App
	args map[string]interface{}
}
type driverCmdSuite struct {
	cmdSuit
}

func (s *driverCmdSuite) TestParseConfig() {
	s.app.After = func(ctx *cli.Context) error {
		s.Equal(l1WSEndpoint, driverConf.L1Endpoint)
		s.Equal(l2WSEndpoint, driverConf.L2Endpoint)
		s.Equal(l2AuthorEndpoint, driverConf.L2EngineEndpoint)
		s.Equal(taikoL1, driverConf.TaikoL1Address.String())
		s.Equal(taikoL2, driverConf.TaikoL2Address.String())
		s.Equal(120*time.Second, driverConf.P2PSyncTimeout)
		s.Equal(rpcTimeout, *driverConf.RPCTimeout)
		s.NotEmpty(driverConf.JwtSecret)
		return nil
	}
	s.NoError(s.app.Run(flagsFromArgs(s.T(), s.args)))
}

func (s *driverCmdSuite) TestJWTError() {
	s.args[JWTSecretFlag.Name] = "wrongsecretfile.txt"
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid JWT secret file")
}

func (s *driverCmdSuite) TestEmptyL2CheckPoint() {
	delete(s.args, CheckPointSyncUrlFlag.Name)
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "empty L2 check point URL")
}

func (s *driverCmdSuite) SetupTest() {
	s.app = cli.NewApp()
	s.app.Flags = driverFlags
	s.app.Action = func(ctx *cli.Context) error {
		return driverConf.Validate(context.Background())
	}
	jwtPath, _ := filepath.Abs("../" + os.Getenv("JWT_SECRET"))
	s.args = map[string]interface{}{
		L1WSEndpointFlag.Name:       os.Getenv("L1_NODE_WS_ENDPOINT"),
		TaikoL1AddressFlag.Name:     os.Getenv("TAIKO_L1_ADDRESS"),
		TaikoL2AddressFlag.Name:     os.Getenv("TAIKO_L2_ADDRESS"),
		VerbosityFlag.Name:          "0",
		LogJsonFlag.Name:            "false",
		MetricsEnabledFlag.Name:     "false",
		MetricsAddrFlag.Name:        "",
		BackOffMaxRetrysFlag.Name:   "10",
		RPCTimeoutFlag.Name:         rpcTimeout.String(),
		WaitReceiptTimeoutFlag.Name: "10s",

		L2WSEndpointFlag.Name:          os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		L2AuthEndpointFlag.Name:        os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
		JWTSecretFlag.Name:             jwtPath,
		P2PSyncVerifiedBlocksFlag.Name: true,
		P2PSyncTimeoutFlag.Name:        "120s",
		CheckPointSyncUrlFlag.Name:     os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
	}
}

func flagsFromArgs(t *testing.T, args map[string]interface{}) []string {
	flags := []string{t.Name()}
	for k, v := range args {
		flags = append(flags, fmt.Sprintf("--%s=%v", k, v))
	}
	return flags
}

func TestDriverCmdSuit(t *testing.T) {
	suite.Run(t, new(driverCmdSuite))
}
