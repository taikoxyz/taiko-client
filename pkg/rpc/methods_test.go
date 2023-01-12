package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

var (
	testAddress1 = common.HexToAddress("0xDA1Ea1362475997419D2055dD43390AEE34c6c37")
	testAddress2 = common.HexToAddress("0x9b557777Be33A8A2fE6aF93E017A0d139B439E5D")
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

	require.NotNil(t, poolContent.ToTxsByPriceAndNonce(newTestClient(t).L2ChainID))
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
