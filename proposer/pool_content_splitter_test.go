package proposer

import (
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/testutils"
)

func (s *ProposerTestSuite) TestPoolContentSplit() {
	// Gas limit is smaller than the limit.
	splitter := &poolContentSplitter{
		chainID:       s.RpcClient.L2ChainID,
		minTxGasLimit: 21000,
	}

	splitted := splitter.split(rpc.PoolContent{
		common.BytesToAddress(testutils.RandomBytes(32)): {
			"0": types.NewTx(&types.LegacyTx{}),
		},
	})

	s.Empty(splitted)

	// Gas limit is larger than the limit.
	splitter = &poolContentSplitter{
		chainID:       s.RpcClient.L2ChainID,
		minTxGasLimit: 21000,
	}

	splitted = splitter.split(rpc.PoolContent{
		common.BytesToAddress(testutils.RandomBytes(32)): {
			"0": types.NewTx(&types.LegacyTx{Gas: 21001}),
		},
	})

	s.Empty(splitted)

	// Transaction's RLP encoded bytes is larger than the limit.
	txBytesTooLarge := types.NewTx(&types.LegacyTx{})

	bytes, err := rlp.EncodeToBytes(txBytesTooLarge)
	s.Nil(err)
	s.NotEmpty(bytes)

	splitter = &poolContentSplitter{
		chainID:           s.RpcClient.L2ChainID,
		maxBytesPerTxList: uint64(len(bytes) - 1),
		minTxGasLimit:     uint64(len(bytes) - 2),
	}

	splitted = splitter.split(rpc.PoolContent{
		common.BytesToAddress(testutils.RandomBytes(32)): {"0": txBytesTooLarge},
	})

	s.Empty(splitted)

	// Transactions that meet the limits
	goldenTouchPriKey, err := crypto.HexToECDSA(bindings.GoldenTouchPrivKey[2:])
	s.Nil(err)

	signer := types.LatestSignerForChainID(s.RpcClient.L2ChainID)
	tx1 := types.MustSignNewTx(goldenTouchPriKey, signer, &types.LegacyTx{Gas: 21001, Nonce: 1})
	tx2 := types.MustSignNewTx(goldenTouchPriKey, signer, &types.LegacyTx{Gas: 21001, Nonce: 2})

	bytes, err = rlp.EncodeToBytes(tx1)
	s.Nil(err)
	s.NotEmpty(bytes)

	splitter = &poolContentSplitter{
		chainID:                 s.RpcClient.L2ChainID,
		minTxGasLimit:           21000,
		maxBytesPerTxList:       uint64(len(bytes) + 1000),
		maxTransactionsPerBlock: 1,
		blockMaxGasLimit:        tx1.Gas() + 1000,
	}

	splitted = splitter.split(rpc.PoolContent{
		bindings.GoldenTouchAddress: {"1": tx1, "2": tx2},
	})

	s.Equal(2, len(splitted))
}

func (s *ProposerTestSuite) TestWeightedShuffle() {
	splitter := &poolContentSplitter{shufflePoolContent: true}

	txLists := make([]types.Transactions, 1024)

	for i := 0; i < len(txLists); i++ {
		var txList types.Transactions
		for j := 0; j < 1024; j++ {
			txList = append(txList, types.NewTx(&types.LegacyTx{Nonce: uint64(j), GasPrice: big.NewInt(int64(i))}))
		}
		txLists[i] = txList
	}

	shuffled := splitter.weightedShuffle(txLists)

	// Whether is sorted
	s.False(sort.SliceIsSorted(shuffled, func(i, j int) bool {
		var (
			gasA uint64 = 0
			gasB uint64 = 0
		)

		for _, tx := range shuffled[i] {
			gasA += tx.GasPrice().Uint64()
		}

		for _, tx := range shuffled[j] {
			gasB += tx.GasPrice().Uint64()
		}

		return gasA < gasB
	}))

	for _, txList := range shuffled {
		s.True(sort.IsSorted(types.TxByNonce(txList)))
	}
}
