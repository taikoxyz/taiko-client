package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildBlockKey(t *testing.T) {
	assert.Equal(t, []byte("block-1"), BuildBlockKey(1))
}
