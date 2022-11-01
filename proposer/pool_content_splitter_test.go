package proposer

import (
	"math/big"
	"sort"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/pkg/rpc"
	"golang.org/x/exp/slices"
)

func TestPoolContentSplit(t *testing.T) {
	// Gas limit is smaller than the limit.
	splitter := &poolContentSplitter{minTxGasLimit: 21000}

	splitted := splitter.split(rpc.PoolContent{
		common.BytesToAddress(randomBytes(32)): {
			"0": types.NewTx(&types.LegacyTx{}),
		},
	})

	require.Empty(t, splitted)

	// Gas limit is larger than the limit.
	splitter = &poolContentSplitter{minTxGasLimit: 21000}

	splitted = splitter.split(rpc.PoolContent{
		common.BytesToAddress(randomBytes(32)): {
			"0": types.NewTx(&types.LegacyTx{Gas: 21001}),
		},
	})

	require.Empty(t, splitted)

	// Transaction's RLP encoded bytes is larger than the limit.
	txBytesTooLarge := types.NewTx(&types.LegacyTx{})

	bytes, err := rlp.EncodeToBytes(txBytesTooLarge)
	require.Nil(t, err)
	require.NotEmpty(t, bytes)

	splitter = &poolContentSplitter{
		maxTxBytesPerBlock: uint64(len(bytes) - 1),
		minTxGasLimit:      uint64(len(bytes) - 2),
	}

	splitted = splitter.split(rpc.PoolContent{
		common.BytesToAddress(randomBytes(32)): {"0": txBytesTooLarge},
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

	splitted = splitter.split(rpc.PoolContent{
		common.BytesToAddress(randomBytes(32)): {"0": tx, "1": tx},
	})

	require.Equal(t, 2, len(splitted))
}

func TestWeightedShuffle(t *testing.T) {
	splitter := &poolContentSplitter{shufflePoolContent: true}

	txs := make(types.Transactions, 1024)

	for i := 0; i < txs.Len(); i++ {
		txs[i] = types.NewTx(&types.LegacyTx{GasPrice: big.NewInt(int64(i))})
	}

	shuffled := splitter.weightedShuffle(txs)

	// Whether is sorted
	require.False(t, sort.SliceIsSorted(shuffled, func(i, j int) bool {
		return shuffled[i].GasPrice().Cmp(shuffled[j].GasPrice()) < 0
	}))

	// Whether contains duplicated elements
	buffer := []uint64{}
	for _, tx := range shuffled {
		require.Equal(t, -1, slices.Index(buffer, tx.GasPrice().Uint64()))
		buffer = append(buffer, tx.GasPrice().Uint64())
	}
}
