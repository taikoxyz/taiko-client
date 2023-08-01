package prover

import (
	"context"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

var testFlags = []cli.Flag{
	&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
	&cli.StringFlag{Name: flags.L1HTTPEndpoint.Name},
	&cli.StringFlag{Name: flags.L2WSEndpoint.Name},
	&cli.StringFlag{Name: flags.L2HTTPEndpoint.Name},
	&cli.StringFlag{Name: flags.TaikoL1Address.Name},
	&cli.StringFlag{Name: flags.TaikoL2Address.Name},
	&cli.StringFlag{Name: flags.L1ProverPrivKey.Name},
	&cli.BoolFlag{Name: flags.Dummy.Name},
	&cli.StringFlag{Name: flags.RandomDummyProofDelay.Name},
	&cli.BoolFlag{Name: flags.OracleProver.Name},
	&cli.StringFlag{Name: flags.OracleProverPrivateKey.Name},
	&cli.BoolFlag{Name: flags.SystemProver.Name},
	&cli.StringFlag{Name: flags.SystemProverPrivateKey.Name},
	&cli.StringFlag{Name: flags.Graffiti.Name},
	&cli.Uint64Flag{Name: flags.RPCTimeout.Name},
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProver() {
	l1WsEndpoint := os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint := os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint := os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint := os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	taikoL1 := os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2 := os.Getenv("TAIKO_L2_ADDRESS")

	app := cli.NewApp()
	app.Flags = testFlags
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1WsEndpoint, c.L1WsEndpoint)
		s.Equal(l1HttpEndpoint, c.L1HttpEndpoint)
		s.Equal(l2WsEndpoint, c.L2WsEndpoint)
		s.Equal(l2HttpEndpoint, c.L2HttpEndpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(
			crypto.PubkeyToAddress(s.p.cfg.L1ProverPrivKey.PublicKey),
			crypto.PubkeyToAddress(c.L1ProverPrivKey.PublicKey),
		)
		s.Equal(30*time.Minute, *c.RandomDummyProofDelayLowerBound)
		s.Equal(time.Hour, *c.RandomDummyProofDelayUpperBound)
		s.True(c.Dummy)
		s.True(c.OracleProver)
		s.Equal(
			crypto.PubkeyToAddress(s.p.cfg.OracleProverPrivateKey.PublicKey),
			crypto.PubkeyToAddress(c.OracleProverPrivateKey.PublicKey),
		)
		s.Equal("", c.Graffiti)
		s.Nil(c.RPCTimeout)
		s.Nil(new(Prover).InitFromCli(context.Background(), ctx))

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1WSEndpoint.Name, l1WsEndpoint,
		"-" + flags.L1HTTPEndpoint.Name, l1HttpEndpoint,
		"-" + flags.L2WSEndpoint.Name, l2WsEndpoint,
		"-" + flags.L2HTTPEndpoint.Name, l2HttpEndpoint,
		"-" + flags.TaikoL1Address.Name, taikoL1,
		"-" + flags.TaikoL2Address.Name, taikoL2,
		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.Dummy.Name,
		"-" + flags.RandomDummyProofDelay.Name, "30m-1h",
		"-" + flags.OracleProver.Name,
		"-" + flags.OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.Graffiti.Name, "",
	}))
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_SystemProver() {
	l1WsEndpoint := os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint := os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint := os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint := os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	taikoL1 := os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2 := os.Getenv("TAIKO_L2_ADDRESS")

	app := cli.NewApp()
	app.Flags = testFlags
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1WsEndpoint, c.L1WsEndpoint)
		s.Equal(l1HttpEndpoint, c.L1HttpEndpoint)
		s.Equal(l2WsEndpoint, c.L2WsEndpoint)
		s.Equal(l2HttpEndpoint, c.L2HttpEndpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(
			crypto.PubkeyToAddress(s.p.cfg.L1ProverPrivKey.PublicKey),
			crypto.PubkeyToAddress(c.L1ProverPrivKey.PublicKey),
		)
		s.Equal(30*time.Minute, *c.RandomDummyProofDelayLowerBound)
		s.Equal(time.Hour, *c.RandomDummyProofDelayUpperBound)
		s.True(c.Dummy)
		s.True(c.SystemProver)
		s.Equal(
			crypto.PubkeyToAddress(s.p.cfg.SystemProverPrivateKey.PublicKey),
			crypto.PubkeyToAddress(c.SystemProverPrivateKey.PublicKey),
		)
		s.Equal("", c.Graffiti)
		s.Nil(new(Prover).InitFromCli(context.Background(), ctx))

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1WSEndpoint.Name, l1WsEndpoint,
		"-" + flags.L1HTTPEndpoint.Name, l1HttpEndpoint,
		"-" + flags.L2WSEndpoint.Name, l2WsEndpoint,
		"-" + flags.L2HTTPEndpoint.Name, l2HttpEndpoint,
		"-" + flags.TaikoL1Address.Name, taikoL1,
		"-" + flags.TaikoL2Address.Name, taikoL2,
		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.Dummy.Name,
		"-" + flags.RandomDummyProofDelay.Name, "30m-1h",
		"-" + flags.SystemProver.Name,
		"-" + flags.SystemProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.Graffiti.Name, "",
	}))
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProverError() {
	l1WsEndpoint := os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint := os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint := os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint := os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	taikoL1 := os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2 := os.Getenv("TAIKO_L2_ADDRESS")

	app := cli.NewApp()
	app.Flags = testFlags
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		s.NotNil(err)
		return err
	}

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1WSEndpoint.Name, l1WsEndpoint,
		"-" + flags.L1HTTPEndpoint.Name, l1HttpEndpoint,
		"-" + flags.L2WSEndpoint.Name, l2WsEndpoint,
		"-" + flags.L2HTTPEndpoint.Name, l2HttpEndpoint,
		"-" + flags.TaikoL1Address.Name, taikoL1,
		"-" + flags.TaikoL2Address.Name, taikoL2,
		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.Dummy.Name,
		"-" + flags.RandomDummyProofDelay.Name, "30m-1h",
		"-" + flags.OracleProver.Name,
		"-" + flags.Graffiti.Name, "",
		"-" + flags.RPCTimeout.Name, "5",
	}), "oracleProver flag set without oracleProverPrivateKey set")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_SystemProverError() {
	l1WsEndpoint := os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint := os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint := os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint := os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	taikoL1 := os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2 := os.Getenv("TAIKO_L2_ADDRESS")

	app := cli.NewApp()
	app.Flags = testFlags
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		s.NotNil(err)
		return err
	}

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1WSEndpoint.Name, l1WsEndpoint,
		"-" + flags.L1HTTPEndpoint.Name, l1HttpEndpoint,
		"-" + flags.L2WSEndpoint.Name, l2WsEndpoint,
		"-" + flags.L2HTTPEndpoint.Name, l2HttpEndpoint,
		"-" + flags.TaikoL1Address.Name, taikoL1,
		"-" + flags.TaikoL2Address.Name, taikoL2,
		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.Dummy.Name,
		"-" + flags.RandomDummyProofDelay.Name, "30m-1h",
		"-" + flags.SystemProver.Name,
		"-" + flags.Graffiti.Name, "",
	}), "systemProver flag set without systemProverPrivateKey set")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_SystemProverAndOracleProverBothSetError() {
	l1WsEndpoint := os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint := os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint := os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint := os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	taikoL1 := os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2 := os.Getenv("TAIKO_L2_ADDRESS")

	app := cli.NewApp()
	app.Flags = testFlags
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		s.NotNil(err)
		return err
	}

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1WSEndpoint.Name, l1WsEndpoint,
		"-" + flags.L1HTTPEndpoint.Name, l1HttpEndpoint,
		"-" + flags.L2WSEndpoint.Name, l2WsEndpoint,
		"-" + flags.L2HTTPEndpoint.Name, l2HttpEndpoint,
		"-" + flags.TaikoL1Address.Name, taikoL1,
		"-" + flags.TaikoL2Address.Name, taikoL2,
		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.Dummy.Name,
		"-" + flags.RandomDummyProofDelay.Name, "30m-1h",
		"-" + flags.OracleProver.Name,
		"-" + flags.SystemProver.Name,
		"-" + flags.Graffiti.Name, "",
	}), "cannot set both oracleProver and systemProver")
}
