package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildBlockKey(t *testing.T) {
	assert.Equal(t, BuildBlockKey("1"), []byte("blockid-1"))
}
