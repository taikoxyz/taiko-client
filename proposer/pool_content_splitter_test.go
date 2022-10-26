package proposer

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/pkg/rpc"
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
