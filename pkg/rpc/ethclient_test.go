package rpc

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestBlockByHash(t *testing.T) {
	client := newTestClientWithTimeout(t)

	head, err := client.L1Client.HeaderByNumber(context.Background(), nil)
	require.Nil(t, err)

	block, err := client.L1Client.BlockByHash(context.Background(), head.Hash())

	require.Nil(t, err)
	require.Equal(t, head.Hash(), block.Hash())
}

func TestBlockNumber(t *testing.T) {
	client := newTestClientWithTimeout(t)

	head, err := client.L1Client.BlockNumber(context.Background())
	require.Nil(t, err)
	require.Greater(t, head, uint64(0))
}

func TestPeerCount(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.PeerCount(context.Background())
	require.NotNil(t, err)
}

func TestTransactionByHash(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, _, err := client.L1Client.TransactionByHash(context.Background(), common.Hash{})
	require.NotNil(t, err)
}

func TestTransactionSender(t *testing.T) {
	client := newTestClientWithTimeout(t)

	block, err := client.L1Client.BlockByNumber(context.Background(), nil)
	require.Nil(t, err)
	require.NotZero(t, block.Transactions().Len())

	sender, err := client.L1Client.TransactionSender(context.Background(), block.Transactions()[0], block.Hash(), 0)
	require.Nil(t, err)
	require.NotEqual(t, common.Address{}, sender)
}

func TestTransactionCount(t *testing.T) {
	client := newTestClientWithTimeout(t)

	block, err := client.L1Client.BlockByNumber(context.Background(), nil)
	require.Nil(t, err)
	require.NotZero(t, block.Transactions().Len())

	c, err := client.L1Client.TransactionCount(context.Background(), block.Hash())
	require.Nil(t, err)
	require.NotZero(t, c)
}

func TestTransactionInBlock(t *testing.T) {
	client := newTestClientWithTimeout(t)

	block, err := client.L1Client.BlockByNumber(context.Background(), nil)
	require.Nil(t, err)
	require.NotZero(t, block.Transactions().Len())

	tx, err := client.L1Client.TransactionInBlock(context.Background(), block.Hash(), 0)
	require.Nil(t, err)
	require.NotEqual(t, common.Hash{}, tx.Hash())
}

func TestNetworkID(t *testing.T) {
	client := newTestClientWithTimeout(t)

	networkID, err := client.L1Client.NetworkID(context.Background())
	require.Nil(t, err)
	require.NotEqual(t, common.Big0.Uint64(), networkID.Uint64())
}

func TestStorageAt(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.StorageAt(context.Background(), common.Address{}, common.Hash{}, nil)
	require.Nil(t, err)
}

func TestCodeAt(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.CodeAt(context.Background(), common.Address{}, nil)
	require.Nil(t, err)
}

func TestNonceAt(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.NonceAt(context.Background(), common.Address{}, nil)
	require.Nil(t, err)
}

func TestPendingBalanceAt(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.PendingBalanceAt(context.Background(), common.Address{})
	require.Nil(t, err)
}

func TestPendingStorageAt(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.PendingStorageAt(context.Background(), common.Address{}, common.Hash{})
	require.Nil(t, err)
}

func TestPendingCodeAt(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.PendingCodeAt(context.Background(), common.Address{})
	require.Nil(t, err)
}

func TestPendingTransactionCount(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.PendingTransactionCount(context.Background())
	require.Nil(t, err)
}

func TestCallContractAtHash(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.CallContractAtHash(context.Background(), ethereum.CallMsg{}, common.Hash{})
	require.NotNil(t, err)
}

func TestPendingCallContract(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.PendingCallContract(context.Background(), ethereum.CallMsg{})
	require.Nil(t, err)
}

func TestSuggestGasPrice(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.SuggestGasPrice(context.Background())
	require.Nil(t, err)
}

func TestSuggestGasTipCap(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.SuggestGasTipCap(context.Background())
	require.Nil(t, err)
}

func TestFeeHistory(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.FeeHistory(context.Background(), 1, nil, []float64{})
	require.Nil(t, err)
}

func TestEstimateGas(t *testing.T) {
	client := newTestClientWithTimeout(t)

	_, err := client.L1Client.EstimateGas(context.Background(), ethereum.CallMsg{})
	require.Nil(t, err)
}
