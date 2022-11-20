package crypto

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

var (
	// 32 zero bytes.
	zero32 [32]byte

	// Account address and private key of golden touch account.
	GoldenTouchPrivKey = func() *secp256k1.ModNScalar {
		b := hexutil.MustDecode(bindings.GoldenTouchPrivKey)
		var priv btcec.PrivateKey
		if overflow := priv.Key.SetByteSlice(b); overflow || priv.Key.IsZero() {
			log.Crit("Invalid private key")
		}
		return &priv.Key
	}()
)

// SignAnchor calculates an ECDSA signature for a V1TaikoL2.anchor transaction.
// ref: https://github.com/taikoxyz/taiko-mono/blob/main/packages/protocol/contracts/libs/LibAnchorSignature.sol
func SignAnchor(hash []byte) ([]byte, error) {
	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}

	sig, ok := signWithK(new(secp256k1.ModNScalar).SetInt(1))(hash)
	if !ok {
		sig, ok = signWithK(new(secp256k1.ModNScalar).SetInt(2))(hash)
		if !ok {
			log.Crit("Failed to sign V1TaikoL2.anchor transaction using K = 1 and K = 2")
		}
	}

	return sig[:], nil
}

// signWithK signs the given hash using fixed K.
func signWithK(k *secp256k1.ModNScalar) func(hash []byte) ([]byte, bool) {
	// k * G
	var kG secp256k1.JacobianPoint
	secp256k1.ScalarBaseMultNonConst(k, &kG)
	kG.ToAffine()

	// r = kG.X mod N
	// r != 0
	r, overflow := fieldToModNScalar(&kG.X)
	pubKeyRecoveryCode := byte(overflow<<1) | byte(kG.Y.IsOddBit())

	kinv := new(secp256k1.ModNScalar).InverseValNonConst(k)
	_s := new(secp256k1.ModNScalar).Mul2(GoldenTouchPrivKey, &r)

	return func(hash []byte) ([]byte, bool) {
		var e secp256k1.ModNScalar
		e.SetByteSlice(hash)
		// copy _s here to avoid modifying the original one.
		_s := *_s
		s := _s.Add(&e).Mul(kinv)
		if s.IsZero() {
			return nil, false
		}
		// copy pubKeyRecoveryCode here to avoid modifying the original one.
		pubKeyRecoveryCode := pubKeyRecoveryCode
		if s.IsOverHalfOrder() {
			s.Negate()

			pubKeyRecoveryCode ^= 0x01
		}

		var sig [65]byte // r(32) + s(32) + v(1)
		r.PutBytesUnchecked(sig[:32])
		s.PutBytesUnchecked(sig[32:64])
		sig[64] = pubKeyRecoveryCode
		return sig[:], true
	}
}

// fieldToModNScalar converts a `secp256k1.FieldVal` to `secp256k1.ModNScalar`.
func fieldToModNScalar(v *secp256k1.FieldVal) (secp256k1.ModNScalar, uint32) {
	var buf [32]byte
	v.PutBytes(&buf)
	var s secp256k1.ModNScalar
	overflow := s.SetBytes(&buf)
	// Clear buf here maybe for preventing memory theft (copy from source)
	resetBuffer(&buf)
	return s, overflow
}

// resetBuffer resets the given buffer.
func resetBuffer(b *[32]byte) {
	copy(b[:], zero32[:])
}
