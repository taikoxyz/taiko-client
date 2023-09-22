package main

import (
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v2"
)

var (
	l1WsEndpoint   = os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint = os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint   = os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint = os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	minProofFee    = "1024"
)

type ProverCmdSuite struct {
	suite.Suite
}

func (s *ProverCmdSuite) TestNewConfigFromCliContext_OracleProver() {
	app := s.SetupApp()
	app.Action = func(ctx *cli.Context) error {
		s.Equal(l1WsEndpoint, proverConf.L1WsEndpoint)
		s.Equal(l1HttpEndpoint, proverConf.L1HttpEndpoint)
		s.Equal(l2WsEndpoint, proverConf.L2WsEndpoint)
		s.Equal(l2HttpEndpoint, proverConf.L2HttpEndpoint)
		s.Equal(taikoL1, proverConf.TaikoL1Address.String())
		s.Equal(taikoL2, proverConf.TaikoL2Address.String())
		s.Equal(
			crypto.PubkeyToAddress(proverConf.L1ProverPrivKey.PublicKey),
			crypto.PubkeyToAddress(proverConf.L1ProverPrivKey.PublicKey),
		)
		s.Equal(30*time.Minute, *proverConf.RandomDummyProofDelayLowerBound)
		s.Equal(time.Hour, *proverConf.RandomDummyProofDelayUpperBound)
		s.True(proverConf.Dummy)
		s.True(proverConf.OracleProver)
		s.Equal(
			crypto.PubkeyToAddress(proverConf.OracleProverPrivateKey.PublicKey),
			crypto.PubkeyToAddress(proverConf.OracleProverPrivateKey.PublicKey),
		)
		s.Equal("", proverConf.Graffiti)
		s.Equal(30*time.Second, proverConf.CheckProofWindowExpiredInterval)
		s.Equal(true, proverConf.ProveUnassignedBlocks)
		s.Equal(rpcTimeout, *proverConf.RPCTimeout)
		s.Equal(uint64(8), proverConf.Capacity)
		s.Equal(minProofFee, proverConf.MinProofFee.String())

		return nil
	}

	s.NoError(app.Run([]string{
		"TestNewConfigFromCliContext_OracleProver",
		"-" + L1WSEndpointFlag.Name, l1WsEndpoint,
		"-" + L1HTTPEndpoint.Name, l1HttpEndpoint,
		"-" + L2WSEndpointFlag.Name, l2WsEndpoint,
		"-" + L2HTTPEndpoint.Name, l2HttpEndpoint,
		"-" + TaikoL1AddressFlag.Name, taikoL1,
		"-" + TaikoL2AddressFlag.Name, taikoL2,
		"-" + L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + StartingBlockID.Name, "0",
		"-" + RPCTimeoutFlag.Name, "5",
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

func (s *ProverCmdSuite) TestNewConfigFromCliContext_OracleProverError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1WSEndpointFlag.Name, l1WsEndpoint,
		"-" + L1HTTPEndpoint.Name, l1HttpEndpoint,
		"-" + L2WSEndpointFlag.Name, l2WsEndpoint,
		"-" + L2HTTPEndpoint.Name, l2HttpEndpoint,
		"-" + TaikoL1AddressFlag.Name, taikoL1,
		"-" + TaikoL2AddressFlag.Name, taikoL2,
		"-" + L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + Dummy.Name,
		"-" + RandomDummyProofDelay.Name, "30m-1h",
		"-" + OracleProver.Name,
		"-" + Graffiti.Name, "",
		"-" + RPCTimeoutFlag.Name, "5",
		"-" + MinProofFee.Name, minProofFee,
	}), "oracleProver flag set without oracleProverPrivateKey set")
}

func (s *ProverCmdSuite) TestNewConfigFromCliContext_ProverKeyError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1ProverPrivKey.Name, "0x",
	}), "invalid L1 prover private key")
}

func (s *ProverCmdSuite) TestNewConfigFromCliContext_OracleProverKeyError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + OracleProver.Name,
		"-" + OracleProverPrivateKey.Name, "",
	}), "invalid oracle private key")
}

func (s *ProverCmdSuite) TestNewConfigFromCliContext_RandomDelayError() {
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

func (s *ProverCmdSuite) TestNewConfigFromCliContext_RandomDelayErrorLower() {
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

func (s *ProverCmdSuite) TestNewConfigFromCliContext_RandomDelayErrorUpper() {
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

func (s *ProverCmdSuite) TestNewConfigFromCliContext_RandomDelayErrorOrder() {
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

func (s *ProverCmdSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: L1WSEndpointFlag.Name},
		&cli.StringFlag{Name: L1HTTPEndpoint.Name},
		&cli.StringFlag{Name: L2WSEndpointFlag.Name},
		&cli.StringFlag{Name: L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: TaikoL1AddressFlag.Name},
		&cli.StringFlag{Name: TaikoL2AddressFlag.Name},
		&cli.StringFlag{Name: L1ProverPrivKey.Name},
		&cli.Uint64Flag{Name: StartingBlockID.Name},
		&cli.BoolFlag{Name: Dummy.Name},
		&cli.StringFlag{Name: RandomDummyProofDelay.Name},
		&cli.BoolFlag{Name: OracleProver.Name},
		&cli.StringFlag{Name: OracleProverPrivateKey.Name},
		&cli.StringFlag{Name: Graffiti.Name},
		&cli.Uint64Flag{Name: CheckProofWindowExpiredInterval.Name},
		&cli.BoolFlag{Name: ProveUnassignedBlocks.Name},
		&cli.Uint64Flag{Name: RPCTimeoutFlag.Name},
		&cli.Uint64Flag{Name: ProverCapacity.Name},
		&cli.Uint64Flag{Name: MinProofFee.Name},
		&cli.Uint64Flag{Name: ProveBlockTxGasLimit.Name},
	}
	app.Action = func(c *cli.Context) error {
		_, err := configProver(c)
		s.NoError(err)
		return nil
	}
	return app
}

func TestProverCmdSuit(t *testing.T) {
	suite.Run(t, new(ProverCmdSuite))
}
