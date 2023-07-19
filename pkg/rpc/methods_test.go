package rpc

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

func TestL2ContentFrom(t *testing.T) {
	client := newTestClient(t)
	l2Head, err := client.L2.HeaderByNumber(context.Background(), nil)
	require.Nil(t, err)

	baseFee, err := client.TaikoL2.GetBasefee(nil, 0, 60000000, uint32(l2Head.GasUsed))
	require.Nil(t, err)

	testAddrPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	require.Nil(t, err)

	testAddr := crypto.PubkeyToAddress(testAddrPrivKey.PublicKey)

	nonce, err := client.L2.PendingNonceAt(context.Background(), testAddr)
	require.Nil(t, err)

	tx := types.NewTransaction(
		nonce,
		testAddr,
		common.Big1,
		100000,
		baseFee,
		[]byte{},
	)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(client.L2ChainID), testAddrPrivKey)
	require.Nil(t, err)
	require.Nil(t, client.L2.SendTransaction(context.Background(), signedTx))

	content, err := client.L2ContentFrom(context.Background(), testAddr)
	require.Nil(t, err)

	require.NotZero(t, len(content["pending"]))
}
