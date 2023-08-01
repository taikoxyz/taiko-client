package proposer

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

var l1Endpoint = os.Getenv("L1_NODE_WS_ENDPOINT")
var l2Endpoint = os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
var taikoL1 = os.Getenv("TAIKO_L1_ADDRESS")
var taikoL2 = os.Getenv("TAIKO_L2_ADDRESS")
var proposeInterval = "10s"
var commitSlot = 1024
var rpcTimeout = 5 * time.Second

func (s *ProposerTestSuite) TestNewConfigFromCliContext() {
	goldenTouchAddress, err := s.RpcClient.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.Nil(err)

	goldenTouchPrivKey, err := s.RpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1Endpoint, c.L1Endpoint)
		s.Equal(l2Endpoint, c.L2Endpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(goldenTouchAddress, crypto.PubkeyToAddress(c.L1ProposerPrivKey.PublicKey))
		s.Equal(goldenTouchAddress, c.L2SuggestedFeeRecipient)
		s.Equal(float64(10), c.ProposeInterval.Seconds())
		s.Equal(uint64(commitSlot), c.CommitSlot)
		s.Equal(1, len(c.LocalAddresses))
		s.Equal(goldenTouchAddress, c.LocalAddresses[0])
		s.Equal(uint64(5), c.ProposeBlockTxReplacementMultiplier)
		s.Equal(rpcTimeout, *c.RPCTimeout)
		s.Nil(new(Proposer).InitFromCli(context.Background(), ctx))

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1WSEndpoint.Name, l1Endpoint,
		"-" + flags.L2HTTPEndpoint.Name, l2Endpoint,
		"-" + flags.TaikoL1Address.Name, taikoL1,
		"-" + flags.TaikoL2Address.Name, taikoL2,
		"-" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + flags.L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"-" + flags.ProposeInterval.Name, proposeInterval,
		"-" + flags.CommitSlot.Name, strconv.Itoa(commitSlot),
		"-" + flags.TxPoolLocals.Name, goldenTouchAddress.Hex(),
		"-" + flags.ProposeBlockTxReplacementMultiplier.Name, "5",
		"-" + flags.RPCTimeout.Name, "5",
		"-" + flags.ProposeBlockTxGasLimit.Name, "100000",
	}))
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextPrivKeyErr() {
	app := s.SetupApp()

	s.NotNil(app.Run([]string{
		"TestNewConfigFromCliContextPrivKeyErr",
		"-" + flags.L1ProposerPrivKey.Name, string(common.FromHex("0x")),
	}))
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextPropIntervalErr() {
	goldenTouchPrivKey, err := s.RpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.NotNil(app.Run([]string{
		"TestNewConfigFromCliContextProposeIntervalErr",
		"-" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + flags.ProposeInterval.Name, "",
	}))
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextEmptyPropoIntervalErr() {
	goldenTouchPrivKey, err := s.RpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.NotNil(app.Run([]string{
		"TestNewConfigFromCliContextEmptyProposalIntervalErr",
		"-" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + flags.ProposeInterval.Name, proposeInterval,
		"-" + flags.ProposeEmptyBlocksInterval.Name, "",
	}))
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextL2RecipErr() {
	goldenTouchPrivKey, err := s.RpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.NotNil(app.Run([]string{
		"TestNewConfigFromCliContextL2RecipErr",
		"-" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + flags.ProposeInterval.Name, proposeInterval,
		"-" + flags.ProposeEmptyBlocksInterval.Name, proposeInterval,
		"-" + flags.L2SuggestedFeeRecipient.Name, "notAnAddress",
	}))
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextTxPoolLocalsErr() {
	goldenTouchAddress, err := s.RpcClient.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.Nil(err)

	goldenTouchPrivKey, err := s.RpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.NotNil(app.Run([]string{
		"TestNewConfigFromCliContextTxPoolLocalsErr",
		"-" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + flags.ProposeInterval.Name, proposeInterval,
		"-" + flags.ProposeEmptyBlocksInterval.Name, proposeInterval,
		"-" + flags.L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"-" + flags.TxPoolLocals.Name, "notAnAddress",
	}))
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextReplMultErr() {
	goldenTouchAddress, err := s.RpcClient.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.Nil(err)

	goldenTouchPrivKey, err := s.RpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.NotNil(app.Run([]string{
		"TestNewConfigFromCliContextReplMultErr",
		"-" + flags.L1ProposerPrivKey.Name, common.Bytes2Hex(goldenTouchPrivKey.Bytes()),
		"-" + flags.ProposeInterval.Name, proposeInterval,
		"-" + flags.ProposeEmptyBlocksInterval.Name, proposeInterval,
		"-" + flags.L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"-" + flags.TxPoolLocals.Name, goldenTouchAddress.Hex(),
		"-" + flags.ProposeBlockTxReplacementMultiplier.Name, "0",
	}))
}

func (s *ProposerTestSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.L1ProposerPrivKey.Name},
		&cli.StringFlag{Name: flags.L2SuggestedFeeRecipient.Name},
		&cli.StringFlag{Name: flags.ProposeEmptyBlocksInterval.Name},
		&cli.StringFlag{Name: flags.ProposeInterval.Name},
		&cli.Uint64Flag{Name: flags.CommitSlot.Name},
		&cli.StringFlag{Name: flags.TxPoolLocals.Name},
		&cli.Uint64Flag{Name: flags.ProposeBlockTxReplacementMultiplier.Name},
		&cli.Uint64Flag{Name: flags.RPCTimeout.Name},
		&cli.Uint64Flag{Name: flags.ProposeBlockTxGasLimit.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		return err
	}
	return app
}
