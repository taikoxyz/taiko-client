package driver

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestL2ChainInserter_newAnchorTransactor(t *testing.T) {
	d := newTestDriver(t)
	opts, err := d.l2ChainInserter.newAnchorTransactor(context.Background(), common.Big0)
	require.Nil(t, err)
	require.Equal(t, true, opts.NoSend)
	require.Equal(t, common.Big0, opts.GasPrice)
	require.Equal(t, common.Big0, opts.Nonce)
	require.Equal(t, goldenTouchAddress, opts.From)
}
