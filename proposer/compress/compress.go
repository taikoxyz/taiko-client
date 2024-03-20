package compress

import (
	"bytes"
	"compress/zlib"
	"io"
)

// EncodeTxListBytes compresses the given txList bytes using zlib.
func EncodeTxListBytes(txList []byte) ([]byte, error) {
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

// DecodeTxListBytes decompresses the given txList bytes using zlib.
func DecodeTxListBytes(compressedTxList []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewBuffer(compressedTxList))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			return nil, err
		}
	}

	return b, nil
}
