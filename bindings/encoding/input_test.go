package encoding

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/bindings"
)

func TestEncodeEvidence(t *testing.T) {
	evidence := &TaikoL1Evidence{
		Meta: bindings.LibDataBlockMetadata{
			Id:          new(big.Int).SetUint64(rand.Uint64()),
			L1Height:    new(big.Int).SetUint64(rand.Uint64()),
			L1Hash:      randomHash(),
			Beneficiary: common.BytesToAddress(randomHash().Bytes()),
			GasLimit:    rand.Uint64(),
			Timestamp:   uint64(time.Now().Unix()),
			TxListHash:  randomHash(),
			MixHash:     randomHash(),
			ExtraData:   randomHash().Bytes(),
		},
		Header: *FromGethHeader(testHeader),
		Prover: common.BytesToAddress(randomHash().Bytes()),
		Proofs: [][]byte{randomHash().Bytes(), randomHash().Bytes(), randomHash().Bytes()},
	}

	b, err := EncodeEvidence(evidence)

	require.Nil(t, err)
	require.NotEmpty(t, b)
}
