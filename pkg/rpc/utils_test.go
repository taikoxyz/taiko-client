package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestWaitConfirmations(t *testing.T) {
	client := newTestClient(t)

	l1Head, err := client.L1.BlockNumber(context.Background())
	require.Nil(t, err)
	require.Nil(t, WaitConfirmations(context.Background(), client.L1, 4, l1Head))
}

func TestWaitReceiptTimeout(t *testing.T) {
	client := newTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := WaitReceipt(
		ctx, client.L1, types.NewTransaction(0, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
	)

	require.ErrorContains(t, err, "context deadline exceeded")
}

func TestGetReceiptsByBlock(t *testing.T) {
	client := newTestClient(t)

	l1Genesis, err := client.L1.BlockByNumber(context.Background(), common.Big0)
	require.Nil(t, err)

	receipts, err := GetReceiptsByBlock(context.Background(), client.L1, l1Genesis)
	require.Nil(t, err)
	require.Empty(t, receipts)
}
