package jwt

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSecretFromFile(t *testing.T) {
	_, err := ParseSecretFromFile(os.Getenv("JWT_SECRET"))
	require.Nil(t, err)

	secret, err := ParseSecretFromFile("")
	require.Nil(t, err)
	require.Nil(t, secret)

	// File not exists
	_, err = ParseSecretFromFile("TestParseSecretFromFile")
	require.NotNil(t, err)

	// Empty file
	file, err := os.CreateTemp("", "TestParseSecretFromFile")
	require.Nil(t, err)
	defer os.Remove(file.Name())

	_, err = ParseSecretFromFile(file.Name())
	require.ErrorContains(t, err, "cannot be empty")
}
