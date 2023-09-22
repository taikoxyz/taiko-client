package main

import (
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/urfave/cli/v2"
)

var (
	l1WsEndpoint   = os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint = os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint   = os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint = os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	minProofFee    = "1024"
)

type ProverTestSuite struct {
	suite.Suite
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProver() {
	app := s.SetupApp()
	app.Action = func(ctx *cli.Context) error {
		c := proverConf
		s.Equal(l1WsEndpoint, c.L1WsEndpoint)
		s.Equal(l1HttpEndpoint, c.L1HttpEndpoint)
		s.Equal(l2WsEndpoint, c.L2WsEndpoint)
		s.Equal(l2HttpEndpoint, c.L2HttpEndpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(
			crypto.PubkeyToAddress(proverConf.L1ProverPrivKey.PublicKey),
			crypto.PubkeyToAddress(c.L1ProverPrivKey.PublicKey),
		)
		s.Equal(30*time.Minute, *c.RandomDummyProofDelayLowerBound)
		s.Equal(time.Hour, *c.RandomDummyProofDelayUpperBound)
		s.True(c.Dummy)
		s.True(c.OracleProver)
		s.Equal(
			crypto.PubkeyToAddress(proverConf.OracleProverPrivateKey.PublicKey),
			crypto.PubkeyToAddress(c.OracleProverPrivateKey.PublicKey),
		)
		s.Equal("", c.Graffiti)
		s.Equal(30*time.Second, c.CheckProofWindowExpiredInterval)
		s.Equal(true, c.ProveUnassignedBlocks)
		s.Equal(rpcTimeout, *c.RPCTimeout)
		s.Equal(uint64(8), c.Capacity)
		s.Equal(minProofFee, c.MinProofFee.String())

		return nil
	}

	s.NoError(app.Run([]string{
		"TestNewConfigFromCliContext_OracleProver",
		"-" + L1WSEndpoint.Name, l1WsEndpoint,
		"-" + L1HTTPEndpoint.Name, l1HttpEndpoint,
		"-" + L2WSEndpoint.Name, l2WsEndpoint,
		"-" + L2HTTPEndpoint.Name, l2HttpEndpoint,
		"-" + TaikoL1Address.Name, taikoL1,
		"-" + TaikoL2Address.Name, taikoL2,
		"-" + L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + StartingBlockID.Name, "0",
		"-" + RPCTimeout.Name, "5",
		"-" + ProveBlockTxGasLimit.Name, "100000",
		"-" + Dummy.Name,
		"-" + RandomDummyProofDelay.Name, "30m-1h",
		"-" + MinProofFee.Name, minProofFee,
		"-" + ProverCapacity.Name, "8",
		"-" + OracleProver.Name,
		"-" + OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + Graffiti.Name, "",
		"-" + CheckProofWindowExpiredInterval.Name, "30",
		"-" + ProveUnassignedBlocks.Name, "true",
	}))
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProverError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1WSEndpoint.Name, l1WsEndpoint,
		"-" + L1HTTPEndpoint.Name, l1HttpEndpoint,
		"-" + L2WSEndpoint.Name, l2WsEndpoint,
		"-" + L2HTTPEndpoint.Name, l2HttpEndpoint,
		"-" + TaikoL1Address.Name, taikoL1,
		"-" + TaikoL2Address.Name, taikoL2,
		"-" + L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + Dummy.Name,
		"-" + RandomDummyProofDelay.Name, "30m-1h",
		"-" + OracleProver.Name,
		"-" + Graffiti.Name, "",
		"-" + RPCTimeout.Name, "5",
		"-" + MinProofFee.Name, minProofFee,
	}), "oracleProver flag set without oracleProverPrivateKey set")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_ProverKeyError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1ProverPrivKey.Name, "0x",
	}), "invalid L1 prover private key")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_OracleProverKeyError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + OracleProver.Name,
		"-" + OracleProverPrivateKey.Name, "",
	}), "invalid oracle private key")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + OracleProver.Name,
		"-" + RandomDummyProofDelay.Name, "130m",
		"-" + MinProofFee.Name, minProofFee,
	}), "invalid random dummy proof delay value")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayErrorLower() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + OracleProver.Name,
		"-" + RandomDummyProofDelay.Name, "30x-1h",
		"-" + MinProofFee.Name, minProofFee,
	}), "invalid random dummy proof delay value")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayErrorUpper() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + OracleProver.Name,
		"-" + RandomDummyProofDelay.Name, "30m-1x",
		"-" + MinProofFee.Name, minProofFee,
	}), "invalid random dummy proof delay value")
}

func (s *ProverTestSuite) TestNewConfigFromCliContext_RandomDelayErrorOrder() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + OracleProverPrivateKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + OracleProver.Name,
		"-" + RandomDummyProofDelay.Name, "1h-30m",
		"-" + MinProofFee.Name, minProofFee,
	}), "invalid random dummy proof delay value (lower > upper)")
}

func (s *ProverTestSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: L1WSEndpoint.Name},
		&cli.StringFlag{Name: L1HTTPEndpoint.Name},
		&cli.StringFlag{Name: L2WSEndpoint.Name},
		&cli.StringFlag{Name: L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: TaikoL1Address.Name},
		&cli.StringFlag{Name: TaikoL2Address.Name},
		&cli.StringFlag{Name: L1ProverPrivKey.Name},
		&cli.Uint64Flag{Name: StartingBlockID.Name},
		&cli.BoolFlag{Name: Dummy.Name},
		&cli.StringFlag{Name: RandomDummyProofDelay.Name},
		&cli.BoolFlag{Name: OracleProver.Name},
		&cli.StringFlag{Name: OracleProverPrivateKey.Name},
		&cli.StringFlag{Name: Graffiti.Name},
		&cli.Uint64Flag{Name: CheckProofWindowExpiredInterval.Name},
		&cli.BoolFlag{Name: ProveUnassignedBlocks.Name},
		&cli.Uint64Flag{Name: RPCTimeout.Name},
		&cli.Uint64Flag{Name: ProverCapacity.Name},
		&cli.Uint64Flag{Name: MinProofFee.Name},
		&cli.Uint64Flag{Name: ProveBlockTxGasLimit.Name},
	}
	app.Action = func(c *cli.Context) error {
		ep, err := rpc.NewClient(c.Context, endpointConf)
		s.NoError(err)
		s.NoError(configProver(c, ep))
		return nil
	}
	return app
}
