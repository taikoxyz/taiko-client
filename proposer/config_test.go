package proposer

import (
	"context"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

func (s *ProposerTestSuite) TestNewConfigFromCliContext() {
	l1Endpoint := os.Getenv("L1_NODE_ENDPOINT")
	l2Endpoint := os.Getenv("L2_NODE_ENDPOINT")
	taikoL1 := os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2 := os.Getenv("TAIKO_L2_ADDRESS")
	proposeInterval := "10s"
	commitSlot := 1024

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1NodeEndpoint.Name},
		&cli.StringFlag{Name: flags.L2NodeEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.L1ProposerPrivKey.Name},
		&cli.StringFlag{Name: flags.L2SuggestedFeeRecipient.Name},
		&cli.StringFlag{Name: flags.ProposeInterval.Name},
		&cli.Uint64Flag{Name: flags.CommitSlot.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1Endpoint, c.L1Endpoint)
		s.Equal(l2Endpoint, c.L2Endpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(bindings.GoldenTouchAddress, crypto.PubkeyToAddress(c.L1ProposerPrivKey.PublicKey))
		s.Equal(bindings.GoldenTouchAddress, c.L2SuggestedFeeRecipient)
		s.Equal(float64(10), c.ProposeInterval.Seconds())
		s.Equal(uint64(commitSlot), c.CommitSlot)
		s.Nil(new(Proposer).InitFromCli(context.Background(), ctx))

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1NodeEndpoint.Name, l1Endpoint,
		"-" + flags.L2NodeEndpoint.Name, l2Endpoint,
		"-" + flags.TaikoL1Address.Name, taikoL1,
		"-" + flags.TaikoL2Address.Name, taikoL2,
		"-" + flags.L1ProposerPrivKey.Name, bindings.GoldenTouchPrivKey[2:],
		"-" + flags.L2SuggestedFeeRecipient.Name, bindings.GoldenTouchAddress.Hex(),
		"-" + flags.ProposeInterval.Name, proposeInterval,
		"-" + flags.CommitSlot.Name, strconv.Itoa(commitSlot),
	}))
}
