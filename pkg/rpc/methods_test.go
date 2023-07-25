package rpc

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	testAddress = common.HexToAddress("0x98f86166571FE624778203d87A8eD6fd84695B79")
)

func TestL2AccountNonce(t *testing.T) {
	client := newTestClient(t)

	nonce, err := client.L2AccountNonce(context.Background(), testAddress, common.Big0)

	require.Nil(t, err)
	require.Zero(t, nonce)
}

func TestGetGenesisL1Header(t *testing.T) {
	client := newTestClient(t)

	header, err := client.GetGenesisL1Header(context.Background())

	require.Nil(t, err)
	require.NotZero(t, header.Number.Uint64())
}

func TestLatestL2KnownL1Header(t *testing.T) {
	client := newTestClient(t)

	header, err := client.LatestL2KnownL1Header(context.Background())

	require.Nil(t, err)
	require.NotZero(t, header.Number.Uint64())
}

func TestL2ParentByBlockId(t *testing.T) {
	client := newTestClient(t)

	header, err := client.L2ParentByBlockId(context.Background(), common.Big1)
	require.Nil(t, err)
	require.Zero(t, header.Number.Uint64())

	_, err = client.L2ParentByBlockId(context.Background(), common.Big2)
	require.NotNil(t, err)
}

func TestL2ExecutionEngineSyncProgress(t *testing.T) {
	client := newTestClient(t)

	progress, err := client.L2ExecutionEngineSyncProgress(context.Background())
	require.Nil(t, err)
	require.NotNil(t, progress)
}

func TestGetProtocolStateVariables(t *testing.T) {
	client := newTestClient(t)
	_, err := client.GetProtocolStateVariables(nil)
	require.Nil(t, err)
}

func TestCheckL1ReorgFromL1Cursor(t *testing.T) {
	client := newTestClient(t)

	l1Head, err := client.L1.HeaderByNumber(context.Background(), nil)
	require.Nil(t, err)

	_, newL1Current, _, err := client.CheckL1ReorgFromL1Cursor(context.Background(), l1Head, l1Head.Number.Uint64())
	require.Nil(t, err)

	require.Equal(t, l1Head.Number.Uint64(), newL1Current.Number.Uint64())

	stateVar, err := client.TaikoL1.GetStateVariables(nil)
	require.Nil(t, err)

	reorged, _, _, err := client.CheckL1ReorgFromL1Cursor(context.Background(), l1Head, stateVar.GenesisHeight)
	require.Nil(t, err)
	require.False(t, reorged)

	l1Head.BaseFee = new(big.Int).Add(l1Head.BaseFee, common.Big1)

	reorged, newL1Current, _, err = client.CheckL1ReorgFromL1Cursor(context.Background(), l1Head, stateVar.GenesisHeight)
	require.Nil(t, err)
	require.True(t, reorged)
	require.Equal(t, l1Head.ParentHash, newL1Current.Hash())
}
