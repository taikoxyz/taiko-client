package proposer

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/testutils"
	"github.com/urfave/cli/v2"
)

var (
	proverEndpoints  = "http://localhost:9876,http://localhost:1234"
	blockProposalFee = "10000000000"
	proposeInterval  = "10s"
	rpcTimeout       = 5 * time.Second
)

func (s *ProposerTestSuite) TestNewConfigFromCliContext() {
	goldenTouchAddress, err := s.rpcClient.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.Nil(err)

	goldenTouchPrivKey, err := s.rpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(s.L1.WsEndpoint(), c.L1Endpoint)
		s.Equal(s.L2.HttpEndpoint(), c.L2Endpoint)
		s.Equal(s.L1.TaikoL1Address, c.TaikoL1Address)
		s.Equal(testutils.TaikoL2Address, c.TaikoL2Address)
		s.Equal(s.L1.TaikoL1TokenAddress, c.TaikoTokenAddress)
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
		s.Nil(new(Proposer).InitFromCli(context.Background(), ctx))

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1WSEndpoint.Name, s.L1.WsEndpoint(),
		"--" + flags.L2HTTPEndpoint.Name, s.L2.HttpEndpoint(),
		"--" + flags.TaikoL1Address.Name, s.L1.TaikoL1Address.Hex(),
		"--" + flags.TaikoL2Address.Name, testutils.TaikoL2Address.Hex(),
		"--" + flags.TaikoTokenAddress.Name, s.L1.TaikoL1TokenAddress.Hex(),
		"--" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"--" + flags.L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"--" + flags.ProposeInterval.Name, proposeInterval,
		"--" + flags.TxPoolLocals.Name, goldenTouchAddress.Hex(),
		"--" + flags.ProposeBlockTxReplacementMultiplier.Name, "5",
		"--" + flags.RPCTimeout.Name, "5",
		"--" + flags.WaitReceiptTimeout.Name, "10",
		"--" + flags.ProposeBlockTxGasTipCap.Name, "100000",
		"--" + flags.ProposeBlockTxGasLimit.Name, "100000",
		"--" + flags.ProverEndpoints.Name, proverEndpoints,
		"--" + flags.BlockProposalFee.Name, blockProposalFee,
		"--" + flags.BlockProposalFeeIncreasePercentage.Name, "15",
		"--" + flags.BlockProposalFeeIterations.Name, "5",
	}))
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextPrivKeyErr() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextPrivKeyErr",
		"--" + flags.L1ProposerPrivKey.Name, string(common.FromHex("0x")),
	}), "invalid L1 proposer private key")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextPropIntervalErr() {
	goldenTouchPrivKey, err := s.rpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextProposeIntervalErr",
		"--" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"--" + flags.ProposeInterval.Name, "",
	}), "invalid proposing interval")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextEmptyPropoIntervalErr() {
	goldenTouchPrivKey, err := s.rpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextEmptyProposalIntervalErr",
		"--" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"--" + flags.ProposeInterval.Name, proposeInterval,
		"--" + flags.ProposeEmptyBlocksInterval.Name, "",
	}), "invalid proposing empty blocks interval")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextL2RecipErr() {
	goldenTouchPrivKey, err := s.rpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextL2RecipErr",
		"--" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"--" + flags.ProposeInterval.Name, proposeInterval,
		"--" + flags.ProposeEmptyBlocksInterval.Name, proposeInterval,
		"--" + flags.L2SuggestedFeeRecipient.Name, "notAnAddress",
	}), "invalid L2 suggested fee recipient address")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextTxPoolLocalsErr() {
	goldenTouchAddress, err := s.rpcClient.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.Nil(err)

	goldenTouchPrivKey, err := s.rpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextTxPoolLocalsErr",
		"--" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"--" + flags.ProposeInterval.Name, proposeInterval,
		"--" + flags.ProposeEmptyBlocksInterval.Name, proposeInterval,
		"--" + flags.L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"--" + flags.TxPoolLocals.Name, "notAnAddress",
	}), "invalid account in --txpool.locals")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextReplMultErr() {
	goldenTouchAddress, err := s.rpcClient.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.Nil(err)

	goldenTouchPrivKey, err := s.rpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextReplMultErr",
		"--" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"--" + flags.ProposeInterval.Name, proposeInterval,
		"--" + flags.ProposeEmptyBlocksInterval.Name, proposeInterval,
		"--" + flags.L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"--" + flags.TxPoolLocals.Name, goldenTouchAddress.Hex(),
		"--" + flags.ProposeBlockTxReplacementMultiplier.Name, "0",
	}), "invalid --proposeBlockTxReplacementMultiplier value")
}

func (s *ProposerTestSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.TaikoTokenAddress.Name},
		&cli.StringFlag{Name: flags.L1ProposerPrivKey.Name},
		&cli.StringFlag{Name: flags.L2SuggestedFeeRecipient.Name},
		&cli.StringFlag{Name: flags.ProposeEmptyBlocksInterval.Name},
		&cli.StringFlag{Name: flags.ProposeInterval.Name},
		&cli.StringFlag{Name: flags.TxPoolLocals.Name},
		&cli.StringFlag{Name: flags.ProverEndpoints.Name},
		&cli.Uint64Flag{Name: flags.BlockProposalFee.Name},
		&cli.Uint64Flag{Name: flags.ProposeBlockTxReplacementMultiplier.Name},
		&cli.Uint64Flag{Name: flags.RPCTimeout.Name},
		&cli.Uint64Flag{Name: flags.WaitReceiptTimeout.Name},
		&cli.Uint64Flag{Name: flags.ProposeBlockTxGasTipCap.Name},
		&cli.Uint64Flag{Name: flags.ProposeBlockTxGasLimit.Name},
		&cli.Uint64Flag{Name: flags.BlockProposalFeeIncreasePercentage.Name},
		&cli.Uint64Flag{Name: flags.BlockProposalFeeIterations.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		return err
	}
	return app
}
