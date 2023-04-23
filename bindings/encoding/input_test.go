package encoding

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestEncodeEvidence(t *testing.T) {
	evidence := &TaikoL1Evidence{
		MetaHash:      randomHash(),
		BlockHash:     randomHash(),
		ParentHash:    randomHash(),
		SignalRoot:    randomHash(),
		Graffiti:      randomHash(),
		Prover:        common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
		ParentGasUsed: 1024,
		GasUsed:       1024,
		VerifierId:    1024,
		Proof:         randomHash().Big().Bytes(),
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
			MetaHash:      randomHash(),
			BlockHash:     randomHash(),
			ParentHash:    randomHash(),
			SignalRoot:    randomHash(),
			Graffiti:      randomHash(),
			Prover:        common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
			ParentGasUsed: 1024,
			GasUsed:       1024,
			VerifierId:    1024,
			Proof:         randomHash().Big().Bytes(),
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
			MetaHash:      randomHash(),
			BlockHash:     randomHash(),
			ParentHash:    randomHash(),
			SignalRoot:    randomHash(),
			Graffiti:      randomHash(),
			Prover:        common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
			ParentGasUsed: 1024,
			GasUsed:       1024,
			VerifierId:    1024,
			Proof:         randomHash().Big().Bytes(),
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
