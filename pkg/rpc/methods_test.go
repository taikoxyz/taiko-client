package rpc

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestPoolContentFaltten(t *testing.T) {
	poolContent := &PoolContent{
		common.HexToAddress("0xDA1Ea1362475997419D2055dD43390AEE34c6c37"): map[string]*types.Transaction{
			"6": types.NewTransaction(6, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
			"5": types.NewTransaction(5, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
			"7": types.NewTransaction(7, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
		},
		common.HexToAddress("0x9b557777Be33A8A2fE6aF93E017A0d139B439E5D"): map[string]*types.Transaction{
			"2": types.NewTransaction(2, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
			"1": types.NewTransaction(1, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
		},
	}

	txs := poolContent.Faltten()

	require.Equal(t, 5, txs.Len())
	require.Equal(t, uint64(5), txs[0].Nonce())
	require.Equal(t, uint64(6), txs[1].Nonce())
	require.Equal(t, uint64(7), txs[2].Nonce())
	require.Equal(t, uint64(1), txs[3].Nonce())
	require.Equal(t, uint64(2), txs[4].Nonce())
}
