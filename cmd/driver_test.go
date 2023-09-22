package main

import (
	"os"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/urfave/cli/v2"
)

var (
	l1EEWS           = os.Getenv("L1_NODE_WS_ENDPOINT")
	l2EEWS           = os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2EngineEndpoint = os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT")
	rpcTimeout       = 5 * time.Second
)

type DriverTestSuite struct {
	suite.Suite
}

func (s *DriverTestSuite) TestNewConfigFromCliContext() {
	app := s.SetupApp()

	app.Action = func(ctx *cli.Context) error {
		c := driverConf
		s.Equal(l1EEWS, c.L1Endpoint)
		s.Equal(l2EEWS, c.L2EngineEndpoint)
		s.Equal(l2EngineEndpoint, c.L2EngineEndpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(120*time.Second, c.P2PSyncTimeout)
		s.Equal(rpcTimeout, *c.RPCTimeout)
		s.NotEmpty(c.JwtSecret)
		return nil
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1WSEndpoint.Name, l1EEWS,
		"-" + L2WSEndpoint.Name, l2EEWS,
		"-" + L2AuthEndpoint.Name, l2EngineEndpoint,
		"-" + TaikoL1Address.Name, taikoL1,
		"-" + TaikoL2Address.Name, taikoL2,
		"-" + JWTSecret.Name, os.Getenv("JWT_SECRET"),
		"-" + P2PSyncTimeout.Name, "120",
		"-" + RPCTimeout.Name, "5",
	}))
}

func (s *DriverTestSuite) TestNewConfigFromCliContextJWTError() {
	app := s.SetupApp()
	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + JWTSecret.Name, "wrongsecretfile.txt",
	}), "invalid JWT secret file")
}

func (s *DriverTestSuite) TestNewConfigFromCliContextEmptyL2CheckPoint() {
	app := s.SetupApp()
	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + JWTSecret.Name, os.Getenv("JWT_SECRET"),
		"-" + P2PSyncVerifiedBlocks.Name, "true",
		"-" + L2WSEndpoint.Name, "",
	}), "empty L2 check point URL")
}

func (s *DriverTestSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: L1WSEndpoint.Name},
		&cli.StringFlag{Name: L2WSEndpoint.Name},
		&cli.StringFlag{Name: L2AuthEndpoint.Name},
		&cli.StringFlag{Name: TaikoL1Address.Name},
		&cli.StringFlag{Name: TaikoL2Address.Name},
		&cli.StringFlag{Name: JWTSecret.Name},
		&cli.BoolFlag{Name: P2PSyncVerifiedBlocks.Name},
		&cli.UintFlag{Name: P2PSyncTimeout.Name},
		&cli.UintFlag{Name: RPCTimeout.Name},
	}
	app.Action = func(c *cli.Context) error {
		ep, err := rpc.NewClient(c.Context, endpointConf)
		s.NoError(err)
		s.NoError(configDriver(c, ep))
		return nil
	}
	return app
}
