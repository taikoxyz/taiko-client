package jwt

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSecretFromFile(t *testing.T) {
	_, err := ParseSecretFromFile(os.Getenv("JWT_SECRET"))

	require.Nil(t, err)
}
