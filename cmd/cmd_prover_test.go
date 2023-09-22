package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/prover"
	"github.com/urfave/cli/v2"
)

var (
	l1WsEndpoint   = os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint = os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint   = os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint = os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	minProofFee    = "1024"
)

type proverCmdSuite struct {
	cmdSuit
}

func (s *proverCmdSuite) TestOracleProver() {
	s.app.After = func(ctx *cli.Context) error {
		s.Equal(l1WsEndpoint, proverConf.L1WsEndpoint)
		s.Equal(l1HttpEndpoint, proverConf.L1HttpEndpoint)
		s.Equal(l2WsEndpoint, proverConf.L2WsEndpoint)
		s.Equal(l2HttpEndpoint, proverConf.L2HttpEndpoint)
		s.Equal(taikoL1, proverConf.TaikoL1Address.String())
		s.Equal(taikoL2, proverConf.TaikoL2Address.String())
		s.Equal(30*time.Minute, *proverConf.RandomDummyProofDelayLowerBound)
		s.Equal(time.Hour, *proverConf.RandomDummyProofDelayUpperBound)
		s.True(proverConf.Dummy)
		s.True(proverConf.OracleProver)
		s.Equal("", proverConf.Graffiti)
		s.Equal(30*time.Second, proverConf.CheckProofWindowExpiredInterval)
		s.Equal(true, proverConf.ProveUnassignedBlocks)
		s.Equal(rpcTimeout, *proverConf.RPCTimeout)
		s.Equal(uint64(8), proverConf.Capacity)
		s.Equal(minProofFee, proverConf.MinProofFee.String())
		return nil
	}

	s.NoError(s.app.Run(flagsFromArgs(s.T(), s.args)))
	s.app.After = nil
}

func (s *proverCmdSuite) TestOracleProverError() {
	s.args[OracleProverFlag.Name] = true
	delete(s.args, OracleProverPrivateKeyFlag.Name)
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "oracleProver flag set without oracleProverPrivateKey set")
}

func (s *proverCmdSuite) TestProverKeyError() {
	s.args[L1ProverPrivKeyFlag.Name] = "0x"
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid L1 prover private key")
}

func (s *proverCmdSuite) TestOracleProverKeyError() {
	s.args[OracleProverFlag.Name] = true
	s.args[OracleProverPrivateKeyFlag.Name] = ""
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid oracle private key")
}

func (s *proverCmdSuite) TestNewConfigFromCliContext_RandomDelayError() {
	s.args[RandomDummyProofDelayFlag.Name] = "130m"
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid random dummy proof delay value")
}

func (s *proverCmdSuite) TestNewConfigFromCliContext_RandomDelayErrorLower() {
	s.args[RandomDummyProofDelayFlag.Name] = "30x-1h"
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid random dummy proof delay value")
}

func (s *proverCmdSuite) TestRandomDelayErrorUpper() {
	s.args[RandomDummyProofDelayFlag.Name] = "30m-1x"
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid random dummy proof delay value")
}

func (s *proverCmdSuite) TestRandomDelayErrorOrder() {
	s.args[RandomDummyProofDelayFlag.Name] = "1h-30m"
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid random dummy proof delay value (lower > upper)")
}

func (s *proverCmdSuite) SetupTest() {
	proverConf = &prover.Config{}
	s.app = cli.NewApp()
	s.app.Flags = proverFlags
	s.app.Action = func(c *cli.Context) error {
		return proverConf.Validate()
	}
	s.args = map[string]interface{}{
		// common flags
		L1WSEndpointFlag.Name:       os.Getenv("L1_NODE_WS_ENDPOINT"),
		TaikoL1AddressFlag.Name:     os.Getenv("TAIKO_L1_ADDRESS"),
		TaikoL2AddressFlag.Name:     os.Getenv("TAIKO_L2_ADDRESS"),
		VerbosityFlag.Name:          "0",
		LogJsonFlag.Name:            "false",
		MetricsEnabledFlag.Name:     "false",
		MetricsAddrFlag.Name:        "",
		BackOffMaxRetrysFlag.Name:   "10",
		RPCTimeoutFlag.Name:         rpcTimeout.String(),
		WaitReceiptTimeoutFlag.Name: "10s",
		// 		// proposer flags
		L1HTTPEndpointFlag.Name:                  os.Getenv("L1_NODE_HTTP_ENDPOINT"),
		L2WSEndpointFlag.Name:                    os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		L2HTTPEndpointFlag.Name:                  os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
		ZkEvmRpcdEndpointFlag.Name:               os.Getenv("ZK_EVM_RPCD_ENDPOINT"),
		ZkEvmRpcdParamsPathFlag.Name:             os.Getenv("ZK_EVM_RPCD_PARAMS_PATH"),
		L1ProverPrivKeyFlag.Name:                 os.Getenv("L1_PROVER_PRIVATE_KEY"),
		MinProofFeeFlag.Name:                     minProofFee,
		StartingBlockIDFlag.Name:                 0,
		MaxConcurrentProvingJobsFlag.Name:        1,
		DummyFlag.Name:                           true,
		RandomDummyProofDelayFlag.Name:           "30m-1h",
		OracleProverFlag.Name:                    true,
		OracleProverPrivateKeyFlag.Name:          os.Getenv("L1_PROVER_PRIVATE_KEY"),
		OracleProofSubmissionDelayFlag.Name:      "10s",
		ProofSubmissionMaxRetryFlag.Name:         3,
		GraffitiFlag.Name:                        "",
		CheckProofWindowExpiredIntervalFlag.Name: "30s",
		ProveUnassignedBlocksFlag.Name:           true,
		ProveBlockTxGasLimitFlag.Name:            "100000",
		ProverCapacityFlag.Name:                  8,
		MaxExpiryFlag.Name:                       "30m",
	}
}

func TestProverCmdSuit(t *testing.T) {
	suite.Run(t, new(proverCmdSuite))
}
