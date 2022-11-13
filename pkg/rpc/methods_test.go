package rpc

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
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

func TestPoolContentToTxLists(t *testing.T) {
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

	txLists := poolContent.ToTxLists()

	require.Equal(t, 2, len(txLists))

	for _, txs := range txLists {
		switch len(txs) {
		case 2:
			require.Equal(t, uint64(1), txs[0].Nonce())
			require.Equal(t, uint64(2), txs[1].Nonce())
		case 3:
			require.Equal(t, uint64(5), txs[0].Nonce())
			require.Equal(t, uint64(6), txs[1].Nonce())
			require.Equal(t, uint64(7), txs[2].Nonce())
		default:
			log.Crit("Invalid txs length")
		}
	}

	require.Equal(t, 5, txLists.Len())
}
