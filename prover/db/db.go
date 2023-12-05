package db

import "fmt"

var (
	BlockKeyPrefix = "blockid-"
)

// BuildBlockKey returns the database key for the given block ID.
func BuildBlockKey(blockID string) []byte {
	return []byte(fmt.Sprintf("%v%v", BlockKeyPrefix, blockID))
}
