package producer

import (
	"context"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

// DummyProofProducer always returns a dummy proof.
type DummyProofProducer struct {
	RandomDummyProofDelayLowerBound *time.Duration
	RandomDummyProofDelayUpperBound *time.Duration
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

	time.AfterFunc(d.proofDelay(), func() {
		resultCh <- &ProofWithHeader{
			BlockID: blockID,
			Meta:    meta,
			Header:  header,
			ZkProof: []byte{0xff},
			Degree:  CircuitsDegree10Txs,
			Opts:    opts,
		}
	})

	return nil
}

// proofDelay calculates a random proof delay between the bounds.
func (d *DummyProofProducer) proofDelay() time.Duration {
	if d.RandomDummyProofDelayLowerBound == nil ||
		d.RandomDummyProofDelayUpperBound == nil ||
		*d.RandomDummyProofDelayUpperBound == time.Duration(0) {
		return time.Duration(0)
	}

	lowerSeconds := int(d.RandomDummyProofDelayLowerBound.Seconds())
	upperSeconds := int(d.RandomDummyProofDelayUpperBound.Seconds())

	randomDurationSeconds := rand.Intn((upperSeconds - lowerSeconds)) + lowerSeconds
	delay := time.Duration(randomDurationSeconds) * time.Second

	log.Info("Random dummy proof delay", "delay", delay)

	return delay
}

// Cancel cancels an existing proof generation.
func (d *DummyProofProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	return nil
}
