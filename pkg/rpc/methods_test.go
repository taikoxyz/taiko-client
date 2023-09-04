package rpc

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	testAddress = common.HexToAddress("0x98f86166571FE624778203d87A8eD6fd84695B79")
)

func TestL2AccountNonce(t *testing.T) {
	client := newTestClientWithTimeout(t)

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

func TestIsJustSyncedByP2P(t *testing.T) {
	client := newTestClient(t)
	_, err := client.IsJustSyncedByP2P(context.Background())
	require.Nil(t, err)
}

func TestWaitTillL2ExecutionEngineSyncedNewClient(t *testing.T) {
	client := newTestClient(t)
	err := client.WaitTillL2ExecutionEngineSynced(context.Background())
	require.Nil(t, err)
}

func TestWaitTillL2ExecutionEngineSyncedContextErr(t *testing.T) {
	client := newTestClient(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	<-ctx.Done()

	err := client.WaitTillL2ExecutionEngineSynced(ctx)
	require.ErrorContains(t, err, "context canceled")
}

func TestGetPoolContentValid(t *testing.T) {
	client := newTestClient(t)
	configs, err := client.TaikoL1.GetConfig(&bind.CallOpts{Context: context.Background()})
	require.Nil(t, err)
	goldenTouchAddress, err := client.TaikoL2.GOLDENTOUCHADDRESS(nil)
	require.Nil(t, err)
	parent, err := client.L2.BlockByNumber(context.Background(), nil)
	require.Nil(t, err)
	baseFee, err := client.TaikoL2.GetBasefee(nil, 1, uint32(parent.GasUsed()))
	require.Nil(t, err)
	gasLimit := configs.BlockMaxGasLimit
	maxBytes := configs.BlockMaxTxListBytes

	txPools := []common.Address{goldenTouchAddress}

	_, err2 := client.GetPoolContent(
		context.Background(),
		goldenTouchAddress,
		baseFee,
		gasLimit,
		maxBytes.Uint64(),
		txPools,
		defaultMaxTransactionsPerBlock,
	)
	require.Nil(t, err2)
}

func TestGetStorageRootNewestBlock(t *testing.T) {
	client := newTestClient(t)
	_, err := client.GetStorageRoot(
		context.Background(),
		client.L1GethClient,
		common.HexToAddress(os.Getenv("L1_SIGNAL_SERVICE_CONTRACT_ADDRESS")),
		nil)
	require.Nil(t, err)
}
