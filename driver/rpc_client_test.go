package driver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/common"
)

func TestLatestL2KnownL1Header(t *testing.T) {
	d := newTestDriver(t)

	header, err := d.rpc.LatestL2KnownL1Header(context.Background())

	require.Nil(t, err)
	require.NotEmpty(t, header.Root)
	require.GreaterOrEqual(t, header.Number.Uint64(), uint64(0))
}

func TestGetGenesisL1Header(t *testing.T) {
	d := newTestDriver(t)

	h, err := d.rpc.GetGenesisL1Header(context.Background())

	require.Nil(t, err)
	require.NotEmpty(t, h.Hash())
	require.True(t, h.Number.Cmp(common.Big0) > 0)
}
