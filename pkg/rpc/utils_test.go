package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestWaitReceiptTimeout(t *testing.T) {
	client := newTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := WaitReceipt(
		ctx, client.L1, types.NewTransaction(0, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
	)

	require.ErrorContains(t, err, "context deadline exceeded")
}

func TestSetHead(t *testing.T) {
	require.Nil(t, SetHead(context.Background(), newTestClient(t).L2RawRPC, common.Big0))
}

func TestStringToBytes32(t *testing.T) {
	require.Equal(t, [32]byte{}, StringToBytes32(""))
	require.Equal(t, [32]byte{0x61, 0x62, 0x63}, StringToBytes32("abc"))
}
