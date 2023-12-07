package db

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildBlockKey(t *testing.T) {
	assert.Equal(t, []byte("block-1"), BuildBlockKey(1))
}

func Test_BuildBlockValue(t *testing.T) {
	v := BuildBlockValue([]byte("hash"), []byte("sig"))
	spl := strings.Split(string(v), "-")
	assert.Equal(t, "hash", spl[0])
	assert.Equal(t, "sig", spl[1])
}
