package driver

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestVerfiyL2Block(t *testing.T) {
	d := newTestDriver(t)

	genesis, err := d.rpc.L2.HeaderByNumber(context.Background(), common.Big0)

	require.Nil(t, err)
	require.Nil(t, d.state.VerfiyL2Block(context.Background(), common.Big0, genesis.Hash()))
}
