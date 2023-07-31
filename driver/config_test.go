package driver

import (
	"context"
	"os"
	"time"

	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

func (s *DriverTestSuite) TestNewConfigFromCliContext() {
	l1Endpoint := os.Getenv("L1_NODE_WS_ENDPOINT")
	l2Endpoint := os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2EngineEndpoint := os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT")
	taikoL1 := os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2 := os.Getenv("TAIKO_L2_ADDRESS")
	rpcTimeout := 5 * time.Second

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2AuthEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.JWTSecret.Name},
		&cli.UintFlag{Name: flags.P2PSyncTimeout.Name},
		&cli.UintFlag{Name: flags.RPCTimeout.Name},
	}
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
	}))
}
