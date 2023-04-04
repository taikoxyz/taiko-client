package encoding

import (
	"errors"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type testJsonError struct{}

func (e *testJsonError) Error() string { return common.Bytes2Hex(randomBytes(10)) }

func (e *testJsonError) ErrorData() interface{} { return "0x3b67b808" }

func TestTryParsingCustomError(t *testing.T) {
	randomErr := common.Bytes2Hex(randomBytes(10))
	require.Equal(t, randomErr, TryParsingCustomError(errors.New(randomErr)).Error())

	err := TryParsingCustomError(errors.New(
		// L1_FORK_CHOICE_NOT_FOUND
		"VM Exception while processing transaction: reverted with an unrecognized custom error (return data: 0x3b67b808)",
	))

	require.True(t, strings.HasPrefix(err.Error(), "L1_FORK_CHOICE_NOT_FOUND"))

	err = TryParsingCustomError(&testJsonError{})

	require.True(t, strings.HasPrefix(err.Error(), "L1_FORK_CHOICE_NOT_FOUND"))
}
