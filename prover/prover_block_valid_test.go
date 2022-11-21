package prover

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/prover/producer"
)

func (s *ProverTestSuite) TestProveBlockValidL1OriginTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s.ErrorContains(
		s.p.proveBlockValid(ctx, &bindings.TaikoL1ClientBlockProposed{Id: common.Big256}), "context deadline exceeded",
	)
}

func (s *ProverTestSuite) TestSubmitValidBlockProofMetadataNotFound() {
	s.ErrorContains(
		s.p.submitValidBlockProof(
			context.Background(), &producer.ProofWithHeader{
				BlockID: common.Big256,
				Header:  &types.Header{},
				ZkProof: []byte{0xff},
			},
		), "failed to fetch L2 block with given block ID",
	)
}
