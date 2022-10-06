package proposer

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/common"
	"github.com/taikochain/taiko-client/core/types"
	"github.com/taikochain/taiko-client/rlp"
)

func TestPoolContentSplit(t *testing.T) {
	// Gas limit is smaller than the limit.
	splitter := &poolContentSplitter{minTxGasLimit: 21000}

	splitted := splitter.split(map[common.Address]map[*big.Int]*types.Transaction{
		common.BytesToAddress(randomBytes(32)): {
			common.Big0: types.NewTx(&types.LegacyTx{}),
		},
	})

	require.Empty(t, splitted)

	// Gas limit is larger than the limit.
	splitter = &poolContentSplitter{minTxGasLimit: 21000}

	splitted = splitter.split(map[common.Address]map[*big.Int]*types.Transaction{
		common.BytesToAddress(randomBytes(32)): {
			common.Big0: types.NewTx(&types.LegacyTx{Gas: 21001}),
		},
	})

	require.Empty(t, splitted)

	// Transaction's RLP encoded bytes is larger than the limit.
	txBytesTooLarge := types.NewTx(&types.LegacyTx{})

	bytes, err := rlp.EncodeToBytes(txBytesTooLarge)
	require.Nil(t, err)
	require.NotEmpty(t, bytes)

	splitter = &poolContentSplitter{maxTxBytesPerBlock: uint64(len(bytes) - 1)}

	splitted = splitter.split(map[common.Address]map[*big.Int]*types.Transaction{
		common.BytesToAddress(randomBytes(32)): {common.Big0: txBytesTooLarge},
	})

	require.Empty(t, splitted)

	// Transactions that meet the limits
	tx := types.NewTx(&types.LegacyTx{Gas: 21001})

	bytes, err = rlp.EncodeToBytes(tx)
	require.Nil(t, err)
	require.NotEmpty(t, bytes)

	splitter = &poolContentSplitter{
		minTxGasLimit:      21000,
		maxTxBytesPerBlock: uint64(len(bytes) + 1),
		maxTxPerBlock:      1,
		maxGasPerBlock:     tx.Gas() + 1,
	}

	splitted = splitter.split(map[common.Address]map[*big.Int]*types.Transaction{
		common.BytesToAddress(randomBytes(32)): {common.Big0: tx, common.Big1: tx},
	})

	require.Equal(t, 2, len(splitted))
}
