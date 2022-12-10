package producer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

// DummyProofProducer always returns a dummy proof.
type DummyProofProducer struct{}

// RequestProof implements the ProofProducer interface.
func (d *DummyProofProducer) RequestProof(
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.LibDataBlockMetadata,
	header *types.Header,
	resultCh chan *ProofWithHeader,
) error {
	log.Info(
		"Request dummy proof",
		"blockID", blockID,
		"meta", meta,
		"height", header.Number,
		"hash", header.Hash(),
	)
	resultCh <- &ProofWithHeader{
		BlockID: blockID, Meta: meta, Header: header, ZkProof: []byte{0xff},
	}
	return nil
}
