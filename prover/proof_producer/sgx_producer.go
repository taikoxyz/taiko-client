package producer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

// SGXProofProducer generates a SGX proof for the given block.
type SGXProofProducer struct{ *DummyProofProducer }

// RequestProof implements the ProofProducer interface.
func (s *SGXProofProducer) RequestProof(
	ctx context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
	resultCh chan *ProofWithHeader,
) error {
	log.Warn(
		"SGX proof producer is unimplemented, will return a dummy proof instead",
		"blockID", blockID,
		"coinbase", meta.Coinbase,
		"height", header.Number,
		"hash", header.Hash(),
	)

	return s.DummyProofProducer.RequestProof(ctx, opts, blockID, meta, header, s.Tier(), resultCh)
}

// Tier implements the ProofProducer interface.
func (s *SGXProofProducer) Tier() uint16 {
	return encoding.TierSgxID
}

// Cancel cancels an existing proof generation.
func (s *SGXProofProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	return nil
}
