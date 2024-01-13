package utils

import "github.com/modern-go/reflect2"

// IsNil Check if the interface is empty.
func IsNil(i interface{}) bool {
	return i == nil || reflect2.IsNil(i)
}
