package producer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

// DummyProofProducer always returns a dummy proof.
type DummyProofProducer struct {
}

// RequestProof implements the ProofProducer interface.
func (d *DummyProofProducer) RequestProof(
	ctx context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
	resultCh chan *ProofWithHeader,
) error {
	log.Info(
		"Request dummy proof",
		"blockID", blockID,
		"beneficiary", meta.Beneficiary,
		"height", header.Number,
		"hash", header.Hash(),
	)

	resultCh <- &ProofWithHeader{
		BlockID: blockID,
		Meta:    meta,
		Header:  header,
		ZkProof: []byte{0xff},
		Degree:  CircuitsDegree10Txs,
		Opts:    opts,
	}

	return nil
}

// Cancel cancels an existing proof generation.
func (d *DummyProofProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	return nil
}
