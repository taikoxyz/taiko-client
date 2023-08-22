package prover

import (
	"context"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

var (
	l1WsEndpoint   = os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint = os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint   = os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint = os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	taikoL1        = os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2        = os.Getenv("TAIKO_L2_ADDRESS")
	rpcTimeout     = 5 * time.Second
)

func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProver() {
	app := s.SetupApp()
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
		s.Equal(30*time.Second, c.CheckProofWindowExpiredInterval)
		s.Equal(true, c.ProveUnassignedBlocks)
		s.Equal(rpcTimeout, *c.RPCTimeout)
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
		"-" + flags.StartingBlockID.Name, "0",
		"-" + flags.RPCTimeout.Name, "5",
		"-" + flags.Dummy.Name,
		"-" + flags.RandomDummyProofDelay.Name, "30m-1h",
		"-" + flags.OracleProver.Name,
		"-" + flags.OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.Graffiti.Name, "",
		"-" + flags.CheckProofWindowExpiredInterval.Name, "30",
		"-" + flags.ProveUnassignedBlocks.Name, "true",
		"-" + flags.ProverCapacity.Name, "8",
	}))
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProverError() {
	app := s.SetupApp()

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

func (s *ProverTestSuite) TestNewConfigFromCliContext_ProverKeyError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1ProverPrivKey.Name, "0x",
	}), "invalid L1 prover private key")
}

// TODO: find case for ToECDSA failing
// func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProverKeyError() {
// 	app := s.SetupApp()

// 	s.ErrorContains(app.Run([]string{
// 		"TestNewConfigFromCliContext",
// 		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
// 		"-" + flags.OracleProverPrivateKey.Name, "0x",
// 	}), "invalid oracle private key")
// }

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.OracleProver.Name,
		"-" + flags.RandomDummyProofDelay.Name, "130m",
	}), "invalid random dummy proof delay value")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayErrorLower() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.OracleProver.Name,
		"-" + flags.RandomDummyProofDelay.Name, "30x-1h",
	}), "invalid random dummy proof delay value")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayErrorUpper() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.OracleProver.Name,
		"-" + flags.RandomDummyProofDelay.Name, "30m-1x",
	}), "invalid random dummy proof delay value")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayErrorOrder() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.OracleProver.Name,
		"-" + flags.RandomDummyProofDelay.Name, "1h-30m",
	}), "invalid random dummy proof delay value (lower > upper)")
}

func (s *ProverTestSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L1HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.L2WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.L1ProverPrivKey.Name},
		&cli.Uint64Flag{Name: flags.StartingBlockID.Name},
		&cli.BoolFlag{Name: flags.Dummy.Name},
		&cli.StringFlag{Name: flags.RandomDummyProofDelay.Name},
		&cli.BoolFlag{Name: flags.OracleProver.Name},
		&cli.StringFlag{Name: flags.OracleProverPrivateKey.Name},
		&cli.StringFlag{Name: flags.Graffiti.Name},
		&cli.Uint64Flag{Name: flags.CheckProofWindowExpiredInterval.Name},
		&cli.BoolFlag{Name: flags.ProveUnassignedBlocks.Name},
		&cli.Uint64Flag{Name: flags.RPCTimeout.Name},
		&cli.Uint64Flag{Name: flags.ProverCapacity.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		s.NotNil(err)
		return err
	}
	return app
}
