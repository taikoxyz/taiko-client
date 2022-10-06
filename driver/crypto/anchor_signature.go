package crypto

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/taikochain/taiko-client/common/hexutil"
	"github.com/taikochain/taiko-client/log"
)

const (
	compactSigSize = 65 // r(32) + s(32) + v(1)
)

var (
	// 32 zero bytes
	zero32 [32]byte

	// Private key of gold finger account
	goldFingerPrivateKey = func() *secp256k1.ModNScalar {
		b := hexutil.MustDecode("0x92954368afd3caa1f3ce3ead0069c1af414054aefe1ef9aeacc1bf426222ce38")
		var priv btcec.PrivateKey
		if overflow := priv.Key.SetByteSlice(b); overflow || priv.Key.IsZero() {
			log.Crit("invalid private key")
		}
		return &priv.Key
	}()

	// ECDSA signing methods when using fixed K = 1 / K = 2
	rs1 = fixedRAndS(1)
	rs2 = fixedRAndS(2)
)

// fixedRAndS signs the given hash using fixed K = 1 / K = 2.
func fixedRAndS(ui uint32) func(hash []byte) ([]byte, bool) {
	k := new(secp256k1.ModNScalar).SetInt(ui)
	// k*G
	var kG secp256k1.JacobianPoint
	secp256k1.ScalarBaseMultNonConst(k, &kG)
	kG.ToAffine()

	// r = kG.X mod N
	// r never zero
	r, overflow := fieldToModNScalar(&kG.X)
	pubKeyRecoveryCode := byte(overflow<<1) | byte(kG.Y.IsOddBit())

	kinv := new(secp256k1.ModNScalar).InverseValNonConst(k)
	_s := new(secp256k1.ModNScalar).Mul2(goldFingerPrivateKey, &r)

	return func(hash []byte) ([]byte, bool) {
		var e secp256k1.ModNScalar
		e.SetByteSlice(hash)
		// copy _s avoid modifying the original one
		_s := *_s
		s := _s.Add(&e).Mul(kinv)
		if s.IsZero() {
			return nil, false
		}
		// copy pubKeyRecoveryCode avoid modifying the original one
		pubKeyRecoveryCode := pubKeyRecoveryCode
		if s.IsOverHalfOrder() {
			s.Negate()

			pubKeyRecoveryCode ^= 0x01
		}

		var sig [compactSigSize]byte
		r.PutBytesUnchecked(sig[:32])
		s.PutBytesUnchecked(sig[32:64])
		sig[64] = pubKeyRecoveryCode
		return sig[:], true
	}
}

// SignAnchor only be used for gold finger account and anchor transaction
func SignAnchor(hash []byte) ([]byte, error) {
	// fixed hash length
	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}

	sig, ok := rs1(hash)
	if !ok {
		sig, ok = rs2(hash)
		if !ok {
			log.Crit("Failed to sign anchor, unrecoverable error")
		}
	}
	return sig[:], nil
}

func fieldToModNScalar(v *secp256k1.FieldVal) (secp256k1.ModNScalar, uint32) {
	var buf [32]byte
	v.PutBytes(&buf)
	var s secp256k1.ModNScalar
	overflow := s.SetBytes(&buf)
	// Clear buf here maybe for preventing memory theft(copy from source)
	zeroArray32(&buf)
	return s, overflow
}

func zeroArray32(b *[32]byte) {
	copy(b[:], zero32[:])
}
