package db

import (
	"bytes"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

var (
	BlockKeyPrefix = "block-"
	separator      = "++"
)

type SignedBlockData struct {
	BlockID   *big.Int
	BlockHash common.Hash
	Signature string
}

// BuildBlockKey will build a block key for a signed block
func BuildBlockKey(blockTimestamp uint64) []byte {
	return bytes.Join(
		[][]byte{
			[]byte(BlockKeyPrefix),
			[]byte(strconv.Itoa(int(blockTimestamp))),
		}, []byte{})
}

// BuildBlockValue will build a block value for a signed block
func BuildBlockValue(hash []byte, signature []byte, blockID *big.Int) []byte {
	return bytes.Join(
		[][]byte{
			hash,
			signature,
			blockID.Bytes(),
		}, []byte(separator))
}

func SignedBlockDataFromValue(val []byte) SignedBlockData {
	v := bytes.Split(val, []byte(separator))

	return SignedBlockData{
		BlockID:   new(big.Int).SetBytes(v[2]),
		BlockHash: common.BytesToHash(v[0]),
		Signature: common.Bytes2Hex(v[1]),
	}
}
