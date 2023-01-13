package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

var (
	prviKey1, _  = crypto.ToECDSA(common.Hex2Bytes("b4851b82a544ba35f2ed8690beec93c9bd2ee8d95bb1255365e70a91065c38c1"))
	prviKey2, _  = crypto.ToECDSA(common.Hex2Bytes("6f20d1e42dfd1ca35c051f9a6730cd0c3003ce446e83b30bfcfff619d449b2cb"))
	testAddress1 = common.HexToAddress("0x98f86166571FE624778203d87A8eD6fd84695B79")
	testAddress2 = common.HexToAddress("0x283593Cd94F70EE3ded6eF92a46Da5Aa8803e7bf")
)

func TestL2PoolContent(t *testing.T) {
	client := newTestClient(t)

	_, _, err := client.L2PoolContent(context.Background())
	require.Nil(t, err)
}

func TestL2AccountNonce(t *testing.T) {
	client := newTestClient(t)

	nonce, err := client.L2AccountNonce(context.Background(), testAddress1, common.Big0)

	require.Nil(t, err)
	require.Zero(t, nonce)
}

func TestPoolContentLen(t *testing.T) {
	poolContent := &PoolContent{
		testAddress1: map[string]*types.Transaction{
			"6": types.NewTransaction(6, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
			"5": types.NewTransaction(5, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
			"7": types.NewTransaction(7, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
		},
		testAddress2: map[string]*types.Transaction{
			"2": types.NewTransaction(2, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
			"1": types.NewTransaction(1, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
		},
	}

	require.Equal(t, 5, poolContent.Len())
}

func TestToTxsByPriceAndNonce(t *testing.T) {
	localPrivKey, err := crypto.ToECDSA(
		common.Hex2Bytes("6977077fd38e2802ec39636bfed7444f1bf93e7ed55bd900a2dbd2caa3787cf5"),
	)
	require.Nil(t, err)
	localAddress := common.HexToAddress("0xc4c76aFa357f95fc35c1ddE165fafaB6449DB639")

	signer := types.LatestSignerForChainID(common.Big1)
	poolContent := &PoolContent{
		testAddress1: map[string]*types.Transaction{
			"6": types.MustSignNewTx(prviKey1, signer, &types.LegacyTx{Gas: 21000, Nonce: 6}),
			"5": types.MustSignNewTx(prviKey1, signer, &types.LegacyTx{Gas: 21000, Nonce: 5}),
			"7": types.MustSignNewTx(prviKey1, signer, &types.LegacyTx{Gas: 21000, Nonce: 7}),
		},
		testAddress2: map[string]*types.Transaction{
			"2": types.MustSignNewTx(prviKey2, signer, &types.LegacyTx{Gas: 21000, Nonce: 2}),
			"1": types.MustSignNewTx(prviKey2, signer, &types.LegacyTx{Gas: 21000, Nonce: 1}),
		},
		localAddress: map[string]*types.Transaction{
			"9":  types.MustSignNewTx(localPrivKey, signer, &types.LegacyTx{Gas: 21000, Nonce: 9}),
			"10": types.MustSignNewTx(localPrivKey, signer, &types.LegacyTx{Gas: 21000, Nonce: 10}),
		},
	}

	locals, remotes := poolContent.ToTxsByPriceAndNonce(common.Big1, []common.Address{localAddress})

	for i := 0; i < 2; i++ {
		require.NotNil(t, locals.Peek())
		locals.Shift()
	}

	require.Nil(t, locals.Peek())

	for i := 0; i < 5; i++ {
		require.NotNil(t, remotes.Peek())
		remotes.Shift()
	}

	require.Nil(t, remotes.Peek())
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
