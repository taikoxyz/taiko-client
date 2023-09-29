package prover

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/testutils"
	"github.com/urfave/cli/v2"
)

var (
	rpcTimeout  = 5 * time.Second
	minProofFee = "1024"
)

func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProver() {
	app := s.SetupApp()
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(s.L1.HttpEndpoint(), c.L1HttpEndpoint)
		s.Equal(s.L2.HttpEndpoint(), c.L2HttpEndpoint)
		s.Equal(testutils.TaikoL1Address, c.TaikoL1Address.String())
		s.Equal(testutils.TaikoL2Address, c.TaikoL2Address.String())
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
		s.Equal(uint64(8), c.Capacity)
		s.Equal(minProofFee, c.MinProofFee.String())
		s.Equal(uint64(3), c.ProveBlockTxReplacementMultiplier)
		s.Equal(uint64(256), c.ProveBlockMaxTxGasTipCap.Uint64())
		s.Equal(15*time.Second, c.TempCapacityExpiresAt)
		s.Nil(new(Prover).InitFromCli(context.Background(), ctx))
		s.True(c.ProveUnassignedBlocks)

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext_OracleProver",
		"--" + flags.L1WSEndpoint.Name, s.L1.WsEndpoint(),
		"--" + flags.L1HTTPEndpoint.Name, s.L1.HttpEndpoint(),
		"--" + flags.L2WSEndpoint.Name, s.L2.WsEndpoint(),
		"--" + flags.L2HTTPEndpoint.Name, s.L2.HttpEndpoint(),
		"--" + flags.TaikoL1Address.Name, testutils.TaikoL1Address.Hex(),
		"--" + flags.TaikoL2Address.Name, testutils.TaikoL2Address.Hex(),
		"--" + flags.L1ProverPrivKey.Name, testutils.ProposerPrivateKey,
		"--" + flags.StartingBlockID.Name, "0",
		"--" + flags.RPCTimeout.Name, "5",
		"--" + flags.ProveBlockTxGasLimit.Name, "100000",
		"--" + flags.Dummy.Name,
		"--" + flags.RandomDummyProofDelay.Name, "30m-1h",
		"--" + flags.MinProofFee.Name, minProofFee,
		"--" + flags.ProverCapacity.Name, "8",
		"--" + flags.OracleProver.Name,
		"--" + flags.ProveBlockTxReplacementMultiplier.Name, "3",
		"--" + flags.ProveBlockMaxTxGasTipCap.Name, "256",
		"--" + flags.OracleProverPrivateKey.Name, testutils.ProverPrivateKey,
		"--" + flags.Graffiti.Name, "",
		"--" + flags.CheckProofWindowExpiredInterval.Name, "30",
		"--" + flags.TempCapacityExpiresAt.Name, "15s",
		"--" + flags.ProveUnassignedBlocks.Name,
	}))
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProverError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1WSEndpoint.Name, s.L1.WsEndpoint(),
		"--" + flags.L1HTTPEndpoint.Name, s.L1.HttpEndpoint(),
		"--" + flags.L2WSEndpoint.Name, s.L2.WsEndpoint(),
		"--" + flags.L2HTTPEndpoint.Name, s.L2.HttpEndpoint(),
		"--" + flags.TaikoL1Address.Name, testutils.TaikoL1Address.Hex(),
		"--" + flags.TaikoL2Address.Name, testutils.TaikoL2Address.Hex(),
		"--" + flags.L1ProverPrivKey.Name, testutils.ProposerPrivateKey,
		"--" + flags.Dummy.Name,
		"--" + flags.RandomDummyProofDelay.Name, "30m-1h",
		"--" + flags.OracleProver.Name,
		"--" + flags.Graffiti.Name, "",
		"--" + flags.RPCTimeout.Name, "5",
		"--" + flags.MinProofFee.Name, minProofFee,
	}), "oracleProver flag set without oracleProverPrivateKey set")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_ProverKeyError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1ProverPrivKey.Name, "0x",
	}), "invalid L1 prover private key")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProverKeyError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1ProverPrivKey.Name, testutils.ProposerPrivateKey,
		"--" + flags.OracleProver.Name,
		"--" + flags.OracleProverPrivateKey.Name, "",
	}), "invalid oracle private key")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1ProverPrivKey.Name, testutils.ProposerPrivateKey,
		"--" + flags.OracleProverPrivateKey.Name, testutils.ProverPrivateKey,
		"--" + flags.OracleProver.Name,
		"--" + flags.RandomDummyProofDelay.Name, "130m",
		"--" + flags.MinProofFee.Name, minProofFee,
	}), "invalid random dummy proof delay value")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayErrorLower() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1ProverPrivKey.Name, testutils.ProposerPrivateKey,
		"--" + flags.OracleProverPrivateKey.Name, testutils.ProverPrivateKey,
		"--" + flags.OracleProver.Name,
		"--" + flags.RandomDummyProofDelay.Name, "30x-1h",
		"--" + flags.MinProofFee.Name, minProofFee,
	}), "invalid random dummy proof delay value")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayErrorUpper() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1ProverPrivKey.Name, testutils.ProposerPrivateKey,
		"--" + flags.OracleProverPrivateKey.Name, testutils.ProverPrivateKey,
		"--" + flags.OracleProver.Name,
		"--" + flags.RandomDummyProofDelay.Name, "30m-1x",
		"--" + flags.MinProofFee.Name, minProofFee,
	}), "invalid random dummy proof delay value")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayErrorOrder() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1ProverPrivKey.Name, testutils.ProposerPrivateKey,
		"--" + flags.OracleProverPrivateKey.Name, testutils.ProverPrivateKey,
		"--" + flags.OracleProver.Name,
		"--" + flags.RandomDummyProofDelay.Name, "1h-30m",
		"--" + flags.MinProofFee.Name, minProofFee,
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
		&cli.Uint64Flag{Name: flags.ProveBlockTxReplacementMultiplier.Name},
		&cli.Uint64Flag{Name: flags.ProveBlockMaxTxGasTipCap.Name},
		&cli.Uint64Flag{Name: flags.RPCTimeout.Name},
		&cli.Uint64Flag{Name: flags.ProverCapacity.Name},
		&cli.Uint64Flag{Name: flags.MinProofFee.Name},
		&cli.Uint64Flag{Name: flags.ProveBlockTxGasLimit.Name},
		&cli.DurationFlag{Name: flags.TempCapacityExpiresAt.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		s.NotNil(err)
		return err
	}
	return app
}
