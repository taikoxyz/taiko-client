package prover

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/bindings"
	"github.com/taikochain/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

func TestNewConfigFromCliContext(t *testing.T) {
	l1Endpoint := randomHash().Hex()
	l2Endpoint := randomHash().Hex()
	taikoL1 := "0x2414c6826978EfB054DD4479e5C0Af0C1549Fe59"
	taikoL2 := "0xDf08F82De32B8d460adbE8D72043E3a7e25A3B39"

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1NodeEndpoint.Name},
		&cli.StringFlag{Name: flags.L2NodeEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.L1ProverPrivKey.Name},
		&cli.BoolFlag{Name: flags.Dummy.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		require.Nil(t, err)
		require.Equal(t, l1Endpoint, c.L1Endpoint)
		require.Equal(t, l2Endpoint, c.L2Endpoint)
		require.Equal(t, taikoL1, c.TaikoL1Address.String())
		require.Equal(t, taikoL2, c.TaikoL2Address.String())
		require.Equal(t, bindings.GoldenTouchAddress, crypto.PubkeyToAddress(c.L1ProverPrivKey.PublicKey))
		require.True(t, c.Dummy)

		return err
	}

	require.Nil(t, app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1NodeEndpoint.Name, l1Endpoint,
		"-" + flags.L2NodeEndpoint.Name, l2Endpoint,
		"-" + flags.TaikoL1Address.Name, taikoL1,
		"-" + flags.TaikoL2Address.Name, taikoL2,
		"-" + flags.L1ProverPrivKey.Name, bindings.GoldenTouchPrivKey[2:],
		"-" + flags.Dummy.Name,
	}))
}
