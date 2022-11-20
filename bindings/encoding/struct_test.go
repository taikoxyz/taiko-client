package encoding

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/bindings"
)

var (
	testHeader = &types.Header{
		ParentHash:  randomHash(),
		UncleHash:   types.EmptyUncleHash,
		Coinbase:    common.BytesToAddress(randomHash().Bytes()),
		Root:        randomHash(),
		TxHash:      randomHash(),
		ReceiptHash: randomHash(),
		Bloom:       types.BytesToBloom(randomHash().Bytes()),
		Difficulty:  new(big.Int).SetUint64(rand.Uint64()),
		Number:      new(big.Int).SetUint64(rand.Uint64()),
		GasLimit:    rand.Uint64(),
		GasUsed:     rand.Uint64(),
		Time:        uint64(time.Now().Unix()),
		Extra:       randomHash().Bytes(),
		MixDigest:   randomHash(),
		Nonce:       types.EncodeNonce(rand.Uint64()),
		BaseFee:     new(big.Int).SetUint64(rand.Uint64()),
	}
	testMeta = bindings.LibDataBlockMetadata{
		Id:          new(big.Int).SetUint64(rand.Uint64()),
		L1Height:    new(big.Int).SetUint64(rand.Uint64()),
		L1Hash:      randomHash(),
		Beneficiary: common.BytesToAddress(randomHash().Bytes()),
		GasLimit:    rand.Uint64(),
		Timestamp:   uint64(time.Now().Unix()),
		TxListHash:  randomHash(),
		MixHash:     randomHash(),
		ExtraData:   randomHash().Bytes(),
	}
)

func TestFromGethHeader(t *testing.T) {
	header := FromGethHeader(testHeader)

	require.Equal(t, testHeader.ParentHash, common.BytesToHash(header.ParentHash[:]))
	require.Equal(t, testHeader.UncleHash, common.BytesToHash(header.OmmersHash[:]))
	require.Equal(t, testHeader.Coinbase, header.Beneficiary)
	require.Equal(t, testHeader.Root, common.BytesToHash(header.StateRoot[:]))
	require.Equal(t, testHeader.TxHash, common.BytesToHash(header.TransactionsRoot[:]))
	require.Equal(t, testHeader.ReceiptHash, common.BytesToHash(header.ReceiptsRoot[:]))
	require.Equal(t, BloomToBytes(testHeader.Bloom), header.LogsBloom)
	require.Equal(t, testHeader.Difficulty, header.Difficulty)
	require.Equal(t, testHeader.Number, header.Height)
	require.Equal(t, testHeader.GasLimit, header.GasLimit)
	require.Equal(t, testHeader.GasUsed, header.GasUsed)
	require.Equal(t, testHeader.Time, header.Timestamp)
	require.Equal(t, testHeader.Extra, header.ExtraData)
	require.Equal(t, testHeader.MixDigest, common.BytesToHash(header.MixHash[:]))
	require.Equal(t, testHeader.Nonce.Uint64(), header.Nonce)
	require.Equal(t, testHeader.BaseFee.Uint64(), header.BaseFeePerGas.Uint64())
}

func TestFromGethHeaderLegacyTx(t *testing.T) {
	testHeader.BaseFee = nil
	header := FromGethHeader(testHeader)

	require.Equal(t, testHeader.ParentHash, common.BytesToHash(header.ParentHash[:]))
	require.Equal(t, testHeader.UncleHash, common.BytesToHash(header.OmmersHash[:]))
	require.Equal(t, testHeader.Coinbase, header.Beneficiary)
	require.Equal(t, testHeader.Root, common.BytesToHash(header.StateRoot[:]))
	require.Equal(t, testHeader.TxHash, common.BytesToHash(header.TransactionsRoot[:]))
	require.Equal(t, testHeader.ReceiptHash, common.BytesToHash(header.ReceiptsRoot[:]))
	require.Equal(t, BloomToBytes(testHeader.Bloom), header.LogsBloom)
	require.Equal(t, testHeader.Difficulty, header.Difficulty)
	require.Equal(t, testHeader.Number, header.Height)
	require.Equal(t, testHeader.GasLimit, header.GasLimit)
	require.Equal(t, testHeader.GasUsed, header.GasUsed)
	require.Equal(t, testHeader.Time, header.Timestamp)
	require.Equal(t, testHeader.Extra, header.ExtraData)
	require.Equal(t, testHeader.MixDigest, common.BytesToHash(header.MixHash[:]))
	require.Equal(t, testHeader.Nonce.Uint64(), header.Nonce)
	require.Equal(t, new(big.Int).SetInt64(0).Uint64(), header.BaseFeePerGas.Uint64())
}

// randomHash generates a random blob of data and returns it as a hash.
func randomHash() common.Hash {
	var hash common.Hash
	if n, err := rand.Read(hash[:]); n != common.HashLength || err != nil {
		panic(err)
	}
	return hash
}
