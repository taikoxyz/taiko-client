package db

import (
	"bytes"
	"strconv"
)

var (
	BlockKeyPrefix = "blockid-"
)

func BuildBlockKey(blockTimestamp uint64) []byte {
	return bytes.Join(
		[][]byte{
			[]byte(BlockKeyPrefix),
			[]byte(strconv.Itoa(int(blockTimestamp))),
		}, []byte{})
}

func BuildBlockValue(hash []byte, signature []byte) []byte {
	return bytes.Join(
		[][]byte{
			hash,
			signature,
		}, []byte("-"))
}
