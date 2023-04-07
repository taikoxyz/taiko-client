package producer

import (
	"context"
	"crypto/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/bindings"
)

func TestRequestProof(t *testing.T) {
	dummyProofProducer := &DummyProofProducer{}

	resCh := make(chan *ProofWithHeader, 1)

	blockID := common.Big32
	header := &types.Header{
		ParentHash:  randHash(),
		UncleHash:   randHash(),
		Coinbase:    common.HexToAddress("0x0000777735367b36bC9B61C50022d9D0700dB4Ec"),
		Root:        randHash(),
		TxHash:      randHash(),
		ReceiptHash: randHash(),
		Difficulty:  common.Big0,
		Number:      common.Big256,
		GasLimit:    1024,
		GasUsed:     1024,
		Time:        uint64(time.Now().Unix()),
		Extra:       randHash().Bytes(),
		MixDigest:   randHash(),
		Nonce:       types.BlockNonce{},
	}
	require.Nil(t, dummyProofProducer.RequestProof(
		context.Background(),
		&ProofRequestOptions{},
		blockID,
		&bindings.TaikoDataBlockMetadata{},
		header,
		resCh,
	))

	res := <-resCh
	require.Equal(t, res.BlockID, blockID)
	require.Equal(t, res.Header, header)
	require.NotEmpty(t, res.ZkProof)
}

func TestProofDelay(t *testing.T) {
	dummyProofProducer := &DummyProofProducer{}
	require.Equal(t, time.Duration(0), dummyProofProducer.proofDelay())

	var (
		delays    []time.Duration
		oneSecond = 1 * time.Second
		oneDay    = 24 * time.Hour
	)
	for i := 0; i < 1024; i++ {
		dummyProofProducer := &DummyProofProducer{
			RandomDummyProofDelayLowerBound: &oneSecond,
			RandomDummyProofDelayUpperBound: &oneDay,
		}

		delay := dummyProofProducer.proofDelay()

		require.LessOrEqual(t, delay, oneDay)
		require.Greater(t, delay, oneSecond)

		delays = append(delays, delay)
	}

	allSame := func(d []time.Duration) bool {
		for i := 1; i < len(d); i++ {
			if d[i] != d[0] {
				return false
			}
		}
		return true
	}

	require.False(t, allSame(delays))
}

func randHash() common.Hash {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		log.Crit("Failed to generate random bytes", err)
	}
	return common.BytesToHash(b)
}
