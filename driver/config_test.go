package driver

import (
	"context"
	"time"

	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/testutils"
	"github.com/urfave/cli/v2"
)

var rpcTimeout = 5 * time.Second

func (s *DriverTestSuite) TestNewConfigFromCliContext() {
	app := s.SetupApp()

	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.NoError(err)
		s.Equal(s.L1.WsEndpoint(), c.L1Endpoint)
		s.Equal(s.L1.WsEndpoint(), c.L2Endpoint)
		s.Equal(s.L2.AuthEndpoint(), c.L2EngineEndpoint)
		s.Equal(testutils.TaikoL1Address, c.TaikoL1Address.String())
		s.Equal(testutils.TaikoL2Address, c.TaikoL2Address.String())
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
		"--" + flags.L1WSEndpoint.Name, s.L1.WsEndpoint(),
		"--" + flags.L2WSEndpoint.Name, s.L2.WsEndpoint(),
		"--" + flags.L2AuthEndpoint.Name, s.L2.AuthEndpoint(),
		"--" + flags.TaikoL1Address.Name, testutils.TaikoL1Address.Hex(),
		"--" + flags.TaikoL2Address.Name, testutils.TaikoL2Address.Hex(),
		"--" + flags.JWTSecret.Name, testutils.JwtSecretFile,
		"--" + flags.P2PSyncTimeout.Name, "120",
		"--" + flags.RPCTimeout.Name, "5",
		"--" + flags.P2PSyncVerifiedBlocks.Name,
		"--" + flags.CheckPointSyncUrl.Name, "http://localhost:8545",
	}))
}

func (s *DriverTestSuite) TestNewConfigFromCliContextJWTError() {
	app := s.SetupApp()
	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.JWTSecret.Name, "wrongsecretfile.txt",
	}), "invalid JWT secret file")
}

func (s *DriverTestSuite) TestNewConfigFromCliContextEmptyL2CheckPoint() {
	app := s.SetupApp()
	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.JWTSecret.Name, testutils.JwtSecretFile,
		"--" + flags.P2PSyncVerifiedBlocks.Name,
		"--" + flags.L2WSEndpoint.Name, "",
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
