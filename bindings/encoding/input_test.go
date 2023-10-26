package encoding

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

func TestEncodeEvidence(t *testing.T) {
	evidence := &BlockEvidence{
		MetaHash:   randomHash(),
		BlockHash:  randomHash(),
		ParentHash: randomHash(),
		SignalRoot: randomHash(),
		Graffiti:   randomHash(),
		Tier:       uint16(rand.Uint64()),
		Proof:      randomHash().Big().Bytes(),
	}

	b, err := EncodeEvidence(evidence)

	require.Nil(t, err)
	require.NotEmpty(t, b)
}

func TestEncodeProverAssignment(t *testing.T) {
	encoded, err := EncodeProverAssignment(
		&ProverAssignment{
			Prover:    common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
			FeeToken:  common.Address{},
			TierFees:  []TierFee{{Tier: 0, Fee: common.Big1}},
			Signature: randomHash().Big().Bytes(),
			Expiry:    1024,
		},
	)

	require.Nil(t, err)
	require.NotNil(t, encoded)
}

func TestEncodeProverAssignmentPayload(t *testing.T) {
	encoded, err := EncodeProverAssignmentPayload(
		common.BytesToHash(randomBytes(32)),
		common.BytesToAddress(randomBytes(20)),
		120,
		[]TierFee{{Tier: 0, Fee: common.Big1}},
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
