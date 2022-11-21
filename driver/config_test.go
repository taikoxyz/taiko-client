package driver

import (
	"context"
	"os"

	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

func (s *DriverTestSuite) TestNewConfigFromCliContext() {
	l1Endpoint := os.Getenv("L1_NODE_ENDPOINT")
	l2Endpoint := os.Getenv("L2_NODE_ENDPOINT")
	l2EngineEndpoint := os.Getenv("L2_NODE_ENGINE_ENDPOINT")
	taikoL1 := os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2 := os.Getenv("TAIKO_L2_ADDRESS")
	throwawayBlocksBuilderPrivKey := os.Getenv("THROWAWAY_BLOCKS_BUILDER_PRIV_KEY")

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1NodeEndpoint.Name},
		&cli.StringFlag{Name: flags.L2NodeEndpoint.Name},
		&cli.StringFlag{Name: flags.L2NodeEngineEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.ThrowawayBlocksBuilderPrivKey.Name},
		&cli.StringFlag{Name: flags.JWTSecret.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1Endpoint, c.L1Endpoint)
		s.Equal(l2Endpoint, c.L2Endpoint)
		s.Equal(l2EngineEndpoint, c.L2EngineEndpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.NotEmpty(c.JwtSecret)
		s.Nil(new(Driver).InitFromCli(context.Background(), ctx))

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1NodeEndpoint.Name, l1Endpoint,
		"-" + flags.L2NodeEndpoint.Name, l2Endpoint,
		"-" + flags.L2NodeEngineEndpoint.Name, l2EngineEndpoint,
		"-" + flags.TaikoL1Address.Name, taikoL1,
		"-" + flags.TaikoL2Address.Name, taikoL2,
		"-" + flags.ThrowawayBlocksBuilderPrivKey.Name, throwawayBlocksBuilderPrivKey,
		"-" + flags.JWTSecret.Name, os.Getenv("JWT_SECRET"),
	}))
}
