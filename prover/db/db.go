package db

import (
	"bytes"
	"strconv"
)

var (
	BlockKeyPrefix = "blockid-"
)

func BuildBlockKey(blockTimestamp uint64) []byte {
	strconv.Itoa(int(blockTimestamp))
	return bytes.Join(
		[][]byte{
			[]byte(BlockKeyPrefix),
			[]byte(strconv.Itoa(int(blockTimestamp))),
		}, []byte{})
}
