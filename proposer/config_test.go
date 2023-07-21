package proposer

import (
	"context"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

func (s *ProposerTestSuite) TestNewConfigFromCliContext() {
	l1Endpoint := os.Getenv("L1_NODE_WS_ENDPOINT")
	l2Endpoint := os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	taikoL1 := os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2 := os.Getenv("TAIKO_L2_ADDRESS")
	proposeInterval := "10s"
	commitSlot := 1024

	goldenTouchAddress, err := s.RpcClient.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.Nil(err)

	goldenTouchPrivKey, err := s.RpcClient.TaikoL2.GOLDENTOUCHPRIVATEKEY(nil)
	s.Nil(err)

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.L1ProposerPrivKey.Name},
		&cli.StringFlag{Name: flags.L2SuggestedFeeRecipient.Name},
		&cli.StringFlag{Name: flags.ProposeInterval.Name},
		&cli.Uint64Flag{Name: flags.CommitSlot.Name},
		&cli.StringFlag{Name: flags.TxPoolLocals.Name},
		&cli.Uint64Flag{Name: flags.ProposeBlockTxReplacementMultiplier.Name},
	}
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
	}))
}
