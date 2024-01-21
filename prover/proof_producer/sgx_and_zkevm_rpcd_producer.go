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
) (*ProofWithHeader, error) {
	log.Info(
		"Request SGX+PSE proof",
		"blockID", blockID,
		"coinbase", meta.Coinbase,
		"height", header.Number,
		"hash", header.Hash(),
	)

	proofs := make([][]byte, 2)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		res, err := o.SGXProofProducer.RequestProof(ctx, opts, blockID, meta, header)
		if err == nil {
			proofs[0] = res.Proof
		}
		return err
	})
	g.Go(func() error {
		res, err := o.ZkevmRpcdProducer.RequestProof(ctx, opts, blockID, meta, header)
		if err == nil {
			proofs[1] = res.Proof
		}
		return err
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &ProofWithHeader{
		BlockID: blockID,
		Meta:    meta,
		Header:  header,
		Proof:   append(proofs[0], proofs[1]...),
		Opts:    opts,
		Tier:    o.Tier(),
	}, nil
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
