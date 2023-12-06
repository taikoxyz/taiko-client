package db

import "fmt"

var (
	BlockKeyPrefix = "blockid-"
)

func BuildBlockKey(blockTimestamp uint64) []byte {
	return []byte(fmt.Sprintf("%v%v", BlockKeyPrefix, blockTimestamp))
}
