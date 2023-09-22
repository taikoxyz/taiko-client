package main

import (
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/urfave/cli/v2"
)

var (
	l2EEHTTP         = os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	proverEndpoints  = "http://localhost:9876,http://localhost:1234"
	taikoL1          = os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2          = os.Getenv("TAIKO_L2_ADDRESS")
	taikoToken       = os.Getenv("TAIKO_TOKEN_ADDRESS")
	blockProposalFee = "10000000000"
	proposeInterval  = "10s"
)

type ProposerTestSuite struct {
	suite.Suite
	RPC *rpc.Client
}

func (s *ProposerTestSuite) TestNewConfigFromCliContext() {
	goldenTouchAddress, err := s.RPC.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.NoError(err)

	goldenTouchPrivKey, err := s.RPC.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.NoError(err)

	app := s.SetupApp()

	app.Action = func(ctx *cli.Context) error {
		c := proposerConf
		s.Equal(l1EEWS, c.L1Endpoint)
		s.Equal(l2EEHTTP, c.L2Endpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(taikoToken, c.TaikoTokenAddress.String())
		s.Equal(goldenTouchAddress, crypto.PubkeyToAddress(c.L1ProposerPrivKey.PublicKey))
		s.Equal(goldenTouchAddress, c.L2SuggestedFeeRecipient)
		s.Equal(float64(10), c.ProposeInterval.Seconds())
		s.Equal(1, len(c.LocalAddresses))
		s.Equal(goldenTouchAddress, c.LocalAddresses[0])
		s.Equal(uint64(5), c.ProposeBlockTxReplacementMultiplier)
		s.Equal(rpcTimeout, *c.RPCTimeout)
		s.Equal(10*time.Second, c.WaitReceiptTimeout)
		for i, e := range strings.Split(proverEndpoints, ",") {
			s.Equal(c.ProverEndpoints[i].String(), e)
		}

		fee, _ := new(big.Int).SetString(blockProposalFee, 10)
		s.Equal(fee, c.BlockProposalFee)

		s.Equal(uint64(15), c.BlockProposalFeeIncreasePercentage.Uint64())
		s.Equal(uint64(5), c.BlockProposalFeeIterations)

		return err
	}

	s.NoError(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + L1WSEndpoint.Name, l1EEWS,
		"-" + L2HTTPEndpoint.Name, l2EEHTTP,
		"-" + TaikoL1Address.Name, taikoL1,
		"-" + TaikoL2Address.Name, taikoL2,
		"-" + TaikoTokenAddress.Name, taikoToken,
		"-" + L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"-" + ProposeInterval.Name, proposeInterval,
		"-" + TxPoolLocals.Name, goldenTouchAddress.Hex(),
		"-" + ProposeBlockTxReplacementMultiplier.Name, "5",
		"-" + RPCTimeout.Name, "5",
		"-" + WaitReceiptTimeout.Name, "10",
		"-" + ProposeBlockTxGasTipCap.Name, "100000",
		"-" + ProposeBlockTxGasLimit.Name, "100000",
		"-" + ProverEndpoints.Name, proverEndpoints,
		"-" + BlockProposalFee.Name, blockProposalFee,
		"-" + BlockProposalFeeIncreasePercentage.Name, "15",
		"-" + BlockProposalFeeIterations.Name, "5",
	}))
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextPrivKeyErr() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextPrivKeyErr",
		"-" + L1ProposerPrivKey.Name, string(common.FromHex("0x")),
	}), "invalid L1 proposer private key")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextPropIntervalErr() {
	goldenTouchPrivKey, err := s.RPC.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.NoError(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextProposeIntervalErr",
		"-" + L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + ProposeInterval.Name, "",
	}), "invalid proposing interval")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextEmptyPropoIntervalErr() {
	goldenTouchPrivKey, err := s.RPC.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.NoError(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextEmptyProposalIntervalErr",
		"-" + L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + ProposeInterval.Name, proposeInterval,
		"-" + ProposeEmptyBlocksInterval.Name, "",
	}), "invalid proposing empty blocks interval")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextL2RecipErr() {
	goldenTouchPrivKey, err := s.RPC.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.NoError(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextL2RecipErr",
		"-" + L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + ProposeInterval.Name, proposeInterval,
		"-" + ProposeEmptyBlocksInterval.Name, proposeInterval,
		"-" + L2SuggestedFeeRecipient.Name, "notAnAddress",
	}), "invalid L2 suggested fee recipient address")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextTxPoolLocalsErr() {
	goldenTouchAddress, err := s.RPC.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.NoError(err)

	goldenTouchPrivKey, err := s.RPC.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.NoError(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextTxPoolLocalsErr",
		"-" + L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + ProposeInterval.Name, proposeInterval,
		"-" + ProposeEmptyBlocksInterval.Name, proposeInterval,
		"-" + L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"-" + TxPoolLocals.Name, "notAnAddress",
	}), "invalid account in --txpool.locals")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextReplMultErr() {
	goldenTouchAddress, err := s.RPC.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.NoError(err)

	goldenTouchPrivKey, err := s.RPC.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.NoError(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextReplMultErr",
		"-" + L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + ProposeInterval.Name, proposeInterval,
		"-" + ProposeEmptyBlocksInterval.Name, proposeInterval,
		"-" + L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"-" + TxPoolLocals.Name, goldenTouchAddress.Hex(),
		"-" + ProposeBlockTxReplacementMultiplier.Name, "0",
	}), "invalid --proposeBlockTxReplacementMultiplier value")
}

func (s *ProposerTestSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: L1WSEndpoint.Name},
		&cli.StringFlag{Name: L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: TaikoL1Address.Name},
		&cli.StringFlag{Name: TaikoL2Address.Name},
		&cli.StringFlag{Name: TaikoTokenAddress.Name},
		&cli.StringFlag{Name: L1ProposerPrivKey.Name},
		&cli.StringFlag{Name: L2SuggestedFeeRecipient.Name},
		&cli.StringFlag{Name: ProposeEmptyBlocksInterval.Name},
		&cli.StringFlag{Name: ProposeInterval.Name},
		&cli.StringFlag{Name: TxPoolLocals.Name},
		&cli.StringFlag{Name: ProverEndpoints.Name},
		&cli.Uint64Flag{Name: BlockProposalFee.Name},
		&cli.Uint64Flag{Name: ProposeBlockTxReplacementMultiplier.Name},
		&cli.Uint64Flag{Name: RPCTimeout.Name},
		&cli.Uint64Flag{Name: WaitReceiptTimeout.Name},
		&cli.Uint64Flag{Name: ProposeBlockTxGasTipCap.Name},
		&cli.Uint64Flag{Name: ProposeBlockTxGasLimit.Name},
		&cli.Uint64Flag{Name: BlockProposalFeeIncreasePercentage.Name},
		&cli.Uint64Flag{Name: BlockProposalFeeIterations.Name},
	}
	app.Action = func(c *cli.Context) error {
		_, err := configProposer(c)
		s.NoError(err)
		return nil
	}
	return app
}
