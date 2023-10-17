package producer

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	bindings "github.com/taikoxyz/taiko-client/bindings/taikol1"
)

// OptimisticProofProducer always returns a dummy proof.
type DummyProofProducer struct{}

// RequestProof returns a dummy proof to the result channel.
func (o *DummyProofProducer) RequestProof(
	ctx context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
	tier uint16,
	resultCh chan *ProofWithHeader,
) error {
	resultCh <- &ProofWithHeader{
		BlockID: blockID,
		Meta:    meta,
		Header:  header,
		Proof:   bytes.Repeat([]byte{0xff}, 100),
		Degree:  CircuitsIdx,
		Opts:    opts,
		Tier:    tier,
	}

	return nil
}
