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
		MetaHash:   randomHash(),
		BlockHash:  randomHash(),
		ParentHash: randomHash(),
		SignalRoot: randomHash(),
		Graffiti:   randomHash(),
		Prover:     common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
		Proofs:     randomHash().Big().Bytes(),
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
			MetaHash:   randomHash(),
			BlockHash:  randomHash(),
			ParentHash: randomHash(),
			SignalRoot: randomHash(),
			Graffiti:   randomHash(),
			Prover:     common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
			Proofs:     randomHash().Big().Bytes(),
		},
	)

	require.Nil(t, err)
	require.NotNil(t, encoded)
}

func TestEncodeProverAssignment(t *testing.T) {
	encoded, err := EncodeProverAssignment(
		&ProverAssignment{
			Prover: common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
			Data:   randomHash().Big().Bytes(),
			Expiry: 1024,
		},
	)

	require.Nil(t, err)
	require.NotNil(t, encoded)
}

func TestEncodeProveBlockInvalidInput(t *testing.T) {
	encoded, err := EncodeProveBlockInvalidInput(
		&TaikoL1Evidence{
			MetaHash:   randomHash(),
			BlockHash:  randomHash(),
			ParentHash: randomHash(),
			SignalRoot: randomHash(),
			Graffiti:   randomHash(),
			Prover:     common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
			Proofs:     randomHash().Big().Bytes(),
		},
		&testMeta,
		types.NewReceipt(randomHash().Bytes(), false, 1024),
	)

	require.Nil(t, err)
	require.NotNil(t, encoded)
}

func TestEncodeBlockMetadata(t *testing.T) {
	// since strings are right padded in solidity https://github.com/ethereum/solidity/issues/1340
	var abcdBytes [32]byte
	copy(abcdBytes[:], common.RightPadBytes([]byte("abcd"), 32))

	// Encode block metadata using EncodeBlockMetadata function
	encoded, err := EncodeBlockMetadata(&bindings.TaikoDataBlockMetadata{
		Id:                uint64(1),
		L1Height:          uint64(1),
		L1Hash:            abcdBytes,
		Proposer:          common.HexToAddress("0x10020FCb72e27650651B05eD2CEcA493bC807Ba4"),
		TxListHash:        abcdBytes,
		TxListByteStart:   big.NewInt(0),
		TxListByteEnd:     big.NewInt(1000),
		GasLimit:          1,
		MixHash:           abcdBytes,
		Timestamp:         uint64(1),
		DepositsProcessed: []bindings.TaikoDataEthDeposit{},
	})

	require.Nil(t, err)
	require.NotNil(t, encoded)

	kgv, err := hexutil.Decode("0x00000000000000000000000000000000000000000000000000000000" +
		"0000002000000000000000000000000000000000000000000000000000000000000000010000000000000" +
		"0000000000000000000000000000000000000000000000000010000000000000000000000000000000000" +
		"0000000000000000000000000000016162636400000000000000000000000000000000000000000000000" +
		"0000000006162636400000000000000000000000000000000000000000000000000000000616263640000" +
		"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
		"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
		"00000003e8000000000000000000000000000000000000000000000000000000000000000100000000000" +
		"000000000000010020fcb72e27650651b05ed2ceca493bc807ba400000000000000000000000000000000" +
		"0000000000000000000000000000016000000000000000000000000000000000000000000000000000000" +
		"00000000000")

	require.Nil(t, err)
	require.Equal(t, kgv, encoded)

	encoded2, err := EncodeBlockMetadata(&bindings.TaikoDataBlockMetadata{
		Id:              uint64(1),
		L1Height:        uint64(1),
		L1Hash:          abcdBytes,
		Proposer:        common.HexToAddress("0x10020FCb72e27650651B05eD2CEcA493bC807Ba4"),
		TxListHash:      abcdBytes,
		TxListByteStart: big.NewInt(0),
		TxListByteEnd:   big.NewInt(1000),
		GasLimit:        1,
		MixHash:         abcdBytes,
		Timestamp:       uint64(1),
		DepositsProcessed: []bindings.TaikoDataEthDeposit{
			{Recipient: common.HexToAddress("0x10020FCb72e27650651B05eD2CEcA493bC807Ba4"), Amount: big.NewInt(2), Id: uint64(1)},
		},
	})

	require.Nil(t, err)
	require.NotNil(t, encoded2)

	kgv2, err := hexutil.Decode("0x0000000000000000000000000000000000000000000000000000000" +
		"0000000200000000000000000000000000000000000000000000000000000000000000001000000000000" +
		"0000000000000000000000000000000000000000000000000001000000000000000000000000000000000" +
		"0000000000000000000000000000001616263640000000000000000000000000000000000000000000000" +
		"0000000000616263640000000000000000000000000000000000000000000000000000000061626364000" +
		"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
		"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
		"000000003e800000000000000000000000000000000000000000000000000000000000000010000000000" +
		"0000000000000010020fcb72e27650651b05ed2ceca493bc807ba40000000000000000000000000000000" +
		"0000000000000000000000000000001600000000000000000000000000000000000000000000000000000" +
		"00000000000100000000000000000000000010020fcb72e27650651b05ed2ceca493bc807ba4000000000" +
		"0000000000000000000000000000000000000000000000000000002000000000000000000000000000000" +
		"0000000000000000000000000000000001")

	require.Nil(t, err)
	require.Equal(t, kgv2, encoded2)
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
