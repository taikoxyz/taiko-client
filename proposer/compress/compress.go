package compress

import (
	"bytes"
	"compress/zlib"
	"errors"
	"io"
)

// CompressTxListBytes compresses the given txList bytes using zlib.
func CompressTxListBytes(txList []byte) ([]byte, error) { // nolint: revive
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	defer w.Close()

	if _, err := w.Write(txList); err != nil {
		return nil, err
	}

	if err := w.Flush(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// DecompressTxListBytes decompresses the given txList bytes using zlib.
func DecompressTxListBytes(compressedTxList []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewBuffer(compressedTxList))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		if !errors.Is(err, io.EOF) && !errors.Is(err, io.ErrUnexpectedEOF) {
			return nil, err
		}
	}

	return b, nil
}
