package main

import (
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/urfave/cli/v2"
)

var (
	proverEndpoints  = "http://localhost:9876,http://localhost:1234"
	taikoL1          = os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2          = os.Getenv("TAIKO_L2_ADDRESS")
	taikoToken       = os.Getenv("TAIKO_TOKEN_ADDRESS")
	blockProposalFee = "10000000000"
	proposeInterval  = "10s"
)

type proposerCmdSuite struct {
	cmdSuit
}

func (s *proposerCmdSuite) TestFlags() {
	s.app.After = func(ctx *cli.Context) error {
		s.Equal(l1WSEndpoint, proposerConf.L1Endpoint)
		s.Equal(l2HttpEndpoint, proposerConf.L2Endpoint)
		s.Equal(taikoL1, proposerConf.TaikoL1Address.String())
		s.Equal(taikoL2, proposerConf.TaikoL2Address.String())
		s.Equal(taikoToken, proposerConf.TaikoTokenAddress.String())
		s.Equal(float64(10), proposerConf.ProposeInterval.Seconds())
		s.Equal(1, len(proposerConf.LocalAddresses))
		s.Equal(uint64(5), proposerConf.ProposeBlockTxReplacementMultiplier)
		s.Equal(rpcTimeout, *proposerConf.RPCTimeout)
		s.Equal(10*time.Second, proposerConf.WaitReceiptTimeout)
		for i, e := range strings.Split(proverEndpoints, ",") {
			s.Equal(proposerConf.ProverEndpoints[i].String(), e)
		}
		fee, _ := new(big.Int).SetString(blockProposalFee, 10)
		s.Equal(fee, proposerConf.BlockProposalFee)
		s.Equal(uint64(10), proposerConf.BlockProposalFeeIncreasePercentage)
		s.Equal(uint64(100), proposerConf.BlockProposalFeeIterations)
		return nil
	}
	s.NoError(s.app.Run(flagsFromArgs(s.T(), s.args)))
	s.app.After = nil
}

func (s *proposerCmdSuite) TestPrivKeyErr() {
	s.args[L1ProposerPrivKeyFlag.Name] = "0x"
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid L1 proposer private key")
}

func (s *proposerCmdSuite) TestL2RecipErr() {
	s.args[L2SuggestedFeeRecipientFlag.Name] = "notAnAddress"
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid L2 suggested fee recipient address")
}

func (s *proposerCmdSuite) TestTxPoolLocalsErr() {
	s.args[TxPoolLocalsFlag.Name] = "notAnAddress"
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid account in --txpool.locals")
}

func (s *proposerCmdSuite) TestBlockTxReplacementMultiplier() {
	s.args[ProposeBlockTxReplacementMultiplierFlag.Name] = "0"
	s.ErrorContains(s.app.Run(flagsFromArgs(s.T(), s.args)), "invalid --proposeBlockTxReplacementMultiplier value")
}

func (s *proposerCmdSuite) SetupTest() {
	proposerConf = &proposer.Config{}
	s.app = cli.NewApp()
	s.app.Flags = proposerFlags
	s.app.Action = func(c *cli.Context) error {
		parseMultiUsedFlags()
		return proposerConf.Validate()
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
		// proposer flags
		L2HTTPEndpointFlag.Name:                      os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
		L1ProposerPrivKeyFlag.Name:                   os.Getenv("L1_PROPOSER_PRIVATE_KEY"),
		L2SuggestedFeeRecipientFlag.Name:             os.Getenv("L2_SUGGESTED_FEE_RECIPIENT"),
		ProposeIntervalFlag.Name:                     proposeInterval,
		TxPoolLocalsFlag.Name:                        "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		TxPoolLocalsOnlyFlag.Name:                    false,
		ProposeEmptyBlocksIntervalFlag.Name:          proposeInterval,
		MaxProposedTxListsPerEpochFlag.Name:          1,
		ProposeBlockTxGasLimitFlag.Name:              "100000",
		ProposeBlockTxReplacementMultiplierFlag.Name: "5",
		ProposeBlockTxGasTipCapFlag.Name:             "100000",
		ProverEndpointsFlag.Name:                     proverEndpoints,
		BlockProposalFeeFlag.Name:                    blockProposalFee,
		BlockProposalFeeIncreasePercentageFlag.Name:  "10",
		BlockProposalFeeIterationsFlag.Name:          "100",
		TaikoTokenAddressFlag.Name:                   os.Getenv("TAIKO_TOKEN_ADDRESS"),
	}
}

func TestProposerCmdSuit(t *testing.T) {
	suite.Run(t, new(proposerCmdSuite))
}
