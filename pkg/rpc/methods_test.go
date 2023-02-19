package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

var (
	prviKey1, _  = crypto.ToECDSA(common.Hex2Bytes("b4851b82a544ba35f2ed8690beec93c9bd2ee8d95bb1255365e70a91065c38c1"))
	prviKey2, _  = crypto.ToECDSA(common.Hex2Bytes("6f20d1e42dfd1ca35c051f9a6730cd0c3003ce446e83b30bfcfff619d449b2cb"))
	testAddress1 = common.HexToAddress("0x98f86166571FE624778203d87A8eD6fd84695B79")
	testAddress2 = common.HexToAddress("0x283593Cd94F70EE3ded6eF92a46Da5Aa8803e7bf")
)

func TestL2AccountNonce(t *testing.T) {
	client := newTestClient(t)

	nonce, err := client.L2AccountNonce(context.Background(), testAddress1, common.Big0)

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

func TestWaitL1OriginTimeout(t *testing.T) {
	client := newTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.WaitL1Origin(ctx, common.Big1)
	require.Nil(t, err)
}

func TestGetProtocolStateVariables(t *testing.T) {
	client := newTestClient(t)
	_, err := client.GetProtocolStateVariables(nil)
	require.Nil(t, err)
}
