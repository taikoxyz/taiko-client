package compress

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/internal/testutils"
)

func TestEncodeDecodeTxListBytes(t *testing.T) {
	b := testutils.RandomBytes(1024)

	compressed, err := EncodeTxListBytes(b)
	require.Nil(t, err)
	require.NotEmpty(t, compressed)

	decompressed, err := DecodeTxListBytes(compressed)
	require.Nil(t, err)
	fmt.Println(1, decompressed)

	require.Equal(t, b, decompressed)
}
