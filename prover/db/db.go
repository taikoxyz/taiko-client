package db

import (
	"bytes"
	"strconv"
)

var (
	BlockKeyPrefix = "blockid-"
)

// BuildBlockKey will build a block key for a signed block
func BuildBlockKey(blockTimestamp uint64) []byte {
	return bytes.Join(
		[][]byte{
			[]byte(BlockKeyPrefix),
			[]byte(strconv.Itoa(int(blockTimestamp))),
		}, []byte{})
}

// BuildBlockValue will build a block value for a signed block
func BuildBlockValue(hash []byte, signature []byte) []byte {
	return bytes.Join(
		[][]byte{
			hash,
			signature,
		}, []byte("-"))
}
