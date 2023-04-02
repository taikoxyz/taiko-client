package encoding

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/bindings"
)

func TestEncodeEvidence(t *testing.T) {
	evidence := &TaikoL1Evidence{
		Meta:   testMeta,
		Header: *FromGethHeader(testHeader),
		Prover: common.BytesToAddress(randomHash().Bytes()),
		Proofs: [][]byte{randomHash().Bytes(), randomHash().Bytes(), randomHash().Bytes()},
	}

	b, err := EncodeEvidence(evidence)

	require.Nil(t, err)
	require.NotEmpty(t, b)
}

func TestEncodeCommitHash(t *testing.T) {
	require.NotEmpty(t, EncodeCommitHash(common.BytesToAddress(randomHash().Bytes()), randomHash()))
}

func TestEncodeProposeBlockInput(t *testing.T) {
	encoded, err := EncodeProposeBlockInput(&testMetaInput)

	require.Nil(t, err)
	require.NotNil(t, encoded)
}

func TestEncodeProveBlockInput(t *testing.T) {
	encoded, err := EncodeProveBlockInput(
		&TaikoL1Evidence{
			Meta:   testMeta,
			Header: *FromGethHeader(testHeader),
			Prover: common.BytesToAddress(randomHash().Bytes()),
		},
		types.NewTransaction(
			0,
			common.BytesToAddress(randomHash().Bytes()),
			common.Big0,
			0,
			common.Big0,
			randomHash().Bytes(),
		),
		types.NewReceipt(randomHash().Bytes(), false, 1024),
	)

	require.Nil(t, err)
	require.NotNil(t, encoded)
}

func TestEncodeProveBlockInvalidInput(t *testing.T) {
	encoded, err := EncodeProveBlockInvalidInput(
		&TaikoL1Evidence{
			Meta:   testMeta,
			Header: *FromGethHeader(testHeader),
			Prover: common.BytesToAddress(randomHash().Bytes()),
		},
		&testMeta,
		types.NewReceipt(randomHash().Bytes(), false, 1024),
	)

	require.Nil(t, err)
	require.NotNil(t, encoded)
}

func TestUnpackTxListBytes(t *testing.T) {
	_, err := UnpackTxListBytes(randomBytes(1024))
	require.NotNil(t, err)

	_, err = UnpackTxListBytes(
		hexutil.MustDecode(
			"0xa0ca2d080000000000000000000000000000000000000000000000000000000000000" +
				"aa8e2b9725cce28787e99447c383d95a9ba83125fe31a9ffa9cbb2c504da86926ab",
		),
	)
	require.ErrorContains(t, err, "no method with id")
}

func TestDecodeEvidenceHeader(t *testing.T) {
	_, err := UnpackEvidenceHeader(randomBytes(1024))
	require.NotNil(t, err)

	_, err = decodeEvidenceHeader(randomBytes(1024))
	require.NotNil(t, err)

	b, err := EncodeEvidence(&TaikoL1Evidence{
		Meta: bindings.TaikoDataBlockMetadata{
			Id:          rand.Uint64(),
			L1Height:    rand.Uint64(),
			L1Hash:      randomHash(),
			Beneficiary: common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
			TxListHash:  randomHash(),
			MixHash:     randomHash(),
			GasLimit:    rand.Uint32(),
			Timestamp:   rand.Uint64(),
		},
		Header: *FromGethHeader(testHeader),
		Prover: common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
		Proofs: [][]byte{randomBytes(1024)},
	})
	require.Nil(t, err)

	header, err := decodeEvidenceHeader(b)
	require.Nil(t, err)
	require.Equal(t, FromGethHeader(testHeader), header)
}
