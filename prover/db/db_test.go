package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildBlockKey(t *testing.T) {
	assert.Equal(t, BuildBlockKey("1"), []byte("blockid-1"))
}
