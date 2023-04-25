package oracle

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

// HashSignAndSetEvidenceForOracleProof hashes and signs the TaikoL1Evidence according to the
// protoco spec to generate an "oracle proof" via the signature and v value.
func HashSignAndSetEvidenceForOracleProof(
	evidence *encoding.TaikoL1Evidence,
	privateKey *ecdsa.PrivateKey,
) ([]byte, uint8, error) {
	evidence.VerifierId = 0
	evidence.Proof = nil

	inputToSign, err := encoding.EncodeProveBlockInput(evidence)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to encode TaikoL1.proveBlock inputs: %w", err)
	}

	hashed := crypto.Keccak256Hash(inputToSign)

	sig, err := crypto.Sign(hashed.Bytes(), privateKey)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to sign TaikoL1Evidence: %w", err)
	}

	// add 27 to be able to be ecrecover in solidity
	v := uint8(int(sig[64])) + 27

	evidence.VerifierId = uint16(v)
	evidence.Proof = sig

	return sig, v, nil
}
