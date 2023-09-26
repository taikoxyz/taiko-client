package driver

import (
	"context"
	"os"
	"time"

	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

var (
	l1Endpoint       = os.Getenv("L1_NODE_WS_ENDPOINT")
	l2Endpoint       = os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2EngineEndpoint = os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT")
	taikoL1          = os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2          = os.Getenv("TAIKO_L2_ADDRESS")
	rpcTimeout       = 5 * time.Second
)

func (s *DriverTestSuite) TestNewConfigFromCliContext() {
	app := s.SetupApp()

	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1Endpoint, c.L1Endpoint)
		s.Equal(l2Endpoint, c.L2Endpoint)
		s.Equal(l2EngineEndpoint, c.L2EngineEndpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(120*time.Second, c.P2PSyncTimeout)
		s.Equal(rpcTimeout, *c.RPCTimeout)
		s.NotEmpty(c.JwtSecret)
		s.Nil(new(Driver).InitFromCli(context.Background(), ctx))
		s.True(c.P2PSyncVerifiedBlocks)
		s.Equal("http://localhost:8545", c.L2CheckPoint)

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1WSEndpoint.Name, l1Endpoint,
		"-" + flags.L2WSEndpoint.Name, l2Endpoint,
		"-" + flags.L2AuthEndpoint.Name, l2EngineEndpoint,
		"-" + flags.TaikoL1Address.Name, taikoL1,
		"-" + flags.TaikoL2Address.Name, taikoL2,
		"-" + flags.JWTSecret.Name, os.Getenv("JWT_SECRET"),
		"-" + flags.P2PSyncTimeout.Name, "120",
		"-" + flags.RPCTimeout.Name, "5",
		"--" + flags.P2PSyncVerifiedBlocks.Name,
		"-" + flags.CheckPointSyncUrl.Name, "http://localhost:8545",
	}))
}

func (s *DriverTestSuite) TestNewConfigFromCliContextJWTError() {
	app := s.SetupApp()
	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.JWTSecret.Name, "wrongsecretfile.txt",
	}), "invalid JWT secret file")
}

func (s *DriverTestSuite) TestNewConfigFromCliContextEmptyL2CheckPoint() {
	app := s.SetupApp()
	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.JWTSecret.Name, os.Getenv("JWT_SECRET"),
		"--" + flags.P2PSyncVerifiedBlocks.Name,
		"-" + flags.L2WSEndpoint.Name, "",
	}), "empty L2 check point URL")
}

func (s *DriverTestSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2AuthEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.JWTSecret.Name},
		&cli.BoolFlag{Name: flags.P2PSyncVerifiedBlocks.Name},
		&cli.UintFlag{Name: flags.P2PSyncTimeout.Name},
		&cli.UintFlag{Name: flags.RPCTimeout.Name},
		&cli.StringFlag{Name: flags.CheckPointSyncUrl.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		return err
	}
	return app
}
