package rpc

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestPoolContentFaltten(t *testing.T) {
	address1 := common.HexToAddress("0xDA1Ea1362475997419D2055dD43390AEE34c6c37")
	address2 := common.HexToAddress("0x9b557777Be33A8A2fE6aF93E017A0d139B439E5D")

	poolContent := &PoolContent{
		address1: map[string]*types.Transaction{
			"6": types.NewTransaction(6, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
			"5": types.NewTransaction(5, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
			"7": types.NewTransaction(7, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
		},
		address2: map[string]*types.Transaction{
			"2": types.NewTransaction(2, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
			"1": types.NewTransaction(1, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
		},
	}

	txs := poolContent.ToTxLists()

	require.Equal(t, 2, len(txs))

	require.Equal(t, uint64(5), txs[0][0].Nonce())
	require.Equal(t, uint64(6), txs[0][1].Nonce())
	require.Equal(t, uint64(7), txs[0][2].Nonce())
	require.Equal(t, uint64(1), txs[1][0].Nonce())
	require.Equal(t, uint64(2), txs[1][1].Nonce())
}
