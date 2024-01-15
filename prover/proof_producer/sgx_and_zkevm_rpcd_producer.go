package producer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/sync/errgroup"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

// SGXAndZkevmRpcdProducer generates a SGX + PSE ZKEVM proof for the given block.
type SGXAndZkevmRpcdProducer struct {
	*SGXProofProducer
	*ZkevmRpcdProducer
}

// RequestProof implements the ProofProducer interface.
func (o *SGXAndZkevmRpcdProducer) RequestProof(
	ctx context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
	resultCh chan *ProofWithHeader,
) error {
	log.Info(
		"Request SGX+PSE proof",
		"blockID", blockID,
		"coinbase", meta.Coinbase,
		"height", header.Number,
		"hash", header.Hash(),
	)

	sgxProofCh := make(chan *ProofWithHeader, 1)
	pseZkEvmProofCh := make(chan *ProofWithHeader, 1)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return o.SGXProofProducer.RequestProof(ctx, opts, blockID, meta, header, sgxProofCh)
	})
	g.Go(func() error {
		return o.ZkevmRpcdProducer.RequestProof(ctx, opts, blockID, meta, header, pseZkEvmProofCh)
	})
	if err := g.Wait(); err != nil {
		return err
	}

	resultCh <- &ProofWithHeader{
		BlockID: blockID,
		Meta:    meta,
		Header:  header,
		Proof:   append((<-sgxProofCh).Proof, (<-pseZkEvmProofCh).Proof...),
		Opts:    opts,
		Tier:    o.Tier(),
	}

	return nil
}

// Tier implements the ProofProducer interface.
func (o *SGXAndZkevmRpcdProducer) Tier() uint16 {
	return encoding.TierSgxAndPseZkevmID
}

// Cancellable implements the ProofProducer interface.
func (o *SGXAndZkevmRpcdProducer) Cancellable() bool {
	return false
}

// Cancel cancels an existing proof generation.
func (o *SGXAndZkevmRpcdProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	return nil
}
