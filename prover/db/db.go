package db

import "fmt"

var (
	BlockKeyPrefix = "blockid-"
)

func BuildBlockKey(blockID string) []byte {
	return []byte(fmt.Sprintf("%v%v", BlockKeyPrefix, blockID))
}
