package utils

import (
	"crypto/rand"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/modern-go/reflect2"
)

func RandUint64(max *big.Int) uint64 {
	if max == nil {
		max = new(big.Int)
		max.SetUint64(math.MaxUint64)
	}
	num, _ := rand.Int(rand.Reader, max)

	return num.Uint64()
}

func RandUint32(max *big.Int) uint32 {
	if max == nil {
		max = new(big.Int)
		max.SetUint64(math.MaxUint32)
	}
	num, _ := rand.Int(rand.Reader, max)
	return uint32(num.Uint64())
}

// IsNil checks if the interface is empty.
func IsNil(i interface{}) bool {
	return i == nil || reflect2.IsNil(i)
}
