package producer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// DummyProofProducer always returns a dummy proof.
type DummyProofProducer struct{}

// RequestProof implements the ProofProducer interface.
func (d *DummyProofProducer) RequestProof(
	opts *ProofRequestOptions,
	blockID *big.Int,
	header *types.Header,
	resultCh chan *ProofWithHeader,
) error {
	log.Info(
		"Request dummy proof",
		"blockID", blockID,
		"height", header.Number,
		"hash", header.Hash(),
	)
	resultCh <- &ProofWithHeader{
		BlockID: blockID, Header: header, ZkProof: []byte{0xff},
	}
	return nil
}
