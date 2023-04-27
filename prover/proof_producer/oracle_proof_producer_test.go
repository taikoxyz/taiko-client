package producer

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
	"math/rand"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

// randomHash generates a random blob of data and returns it as a hash.
func randomHash() common.Hash {
	var hash common.Hash
	if n, err := rand.Read(hash[:]); n != common.HashLength || err != nil {
		panic(err)
	}
	return hash
}

func TestHashAndSignOracleProof(t *testing.T) {
	evidence := &encoding.TaikoL1Evidence{
		MetaHash:      randomHash(),
		BlockHash:     randomHash(),
		ParentHash:    randomHash(),
		SignalRoot:    randomHash(),
		Graffiti:      randomHash(),
		Prover:        common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
		ParentGasUsed: 1024,
		GasUsed:       1024,
		VerifierId:    0,
		Proof:         nil,
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("L1_PROVER_PRIVATE_KEY"))
	require.Nil(t, err)

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)

	input, err := encoding.EncodeProveBlockInput(evidence)
	require.Nil(t, err)

	hash := crypto.Keccak256Hash(input)

	sig, _, err := hashAndSignForOracleProof(evidence, privateKey)
	require.Nil(t, err)

	pubKey, err := crypto.Ecrecover(hash.Bytes(), sig)
	require.Nil(t, err)
	isValid := crypto.VerifySignature(pubKey, hash.Bytes(), sig[:64])
	require.True(t, isValid)

	require.Equal(t, elliptic.Marshal(publicKeyECDSA, publicKeyECDSA.X, publicKeyECDSA.Y), pubKey)
}
