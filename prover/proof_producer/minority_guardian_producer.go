package producer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

// MinorityGuardianProofProducer always returns an optimistic (dummy) proof.
type MinorityGuardianProofProducer struct {
	returnLivenessBond bool
	DummyProofProducer
}

func NewMinorityGuardianProofProducer(returnLivenessBond bool) *MinorityGuardianProofProducer {
	return &MinorityGuardianProofProducer{
		returnLivenessBond: returnLivenessBond,
	}
}

func (g *MinorityGuardianProofProducer) RequestProof(
	_ context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
) (*ProofWithHeader, error) {
	log.Info(
		"Request guardian proof",
		"blockID", blockID,
		"coinbase", meta.Coinbase,
		"height", header.Number,
		"hash", header.Hash(),
	)

	if g.returnLivenessBond {
		return &ProofWithHeader{
			BlockID: blockID,
			Meta:    meta,
			Header:  header,
			Proof:   crypto.Keccak256([]byte("RETURN_LIVENESS_BOND")),
			Opts:    opts,
			Tier:    g.Tier(),
		}, nil
	}

	return g.DummyProofProducer.RequestProof(opts, blockID, meta, header, g.Tier())
}

// Tier returns TierGuardianMinorityID
func (g *MinorityGuardianProofProducer) Tier() uint16 {
	return encoding.TierGuardianMinorityID
}
