package encoding

import (
	"errors"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/testutils"
)

type testJsonError struct{}

func (e *testJsonError) Error() string { return common.Bytes2Hex(testutils.RandomBytes(10)) }

func (e *testJsonError) ErrorData() interface{} { return "0xb6d363fd" }

func TestTryParseCustomError(t *testing.T) {
	randomErr := common.Bytes2Hex(testutils.RandomBytes(10))
	require.Equal(t, randomErr, TryParseCustomError(errors.New(randomErr)).Error())

	err := TryParseCustomError(errors.New(
		// L1_COMMITTED
		"VM Exception while processing transaction: reverted with an unrecognized custom error (return data: 0xb6d363fd)",
	))

	require.True(t, strings.HasPrefix(err.Error(), "L1_COMMITTED"))

	err = TryParseCustomError(&testJsonError{})

	require.True(t, strings.HasPrefix(err.Error(), "L1_COMMITTED"))
}
