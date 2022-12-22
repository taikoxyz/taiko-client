package prover

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/prover/producer"
)

func (s *ProverTestSuite) TestProveBlockInvalidL1OriginTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s.ErrorContains(
		s.p.proveBlockInvalid(ctx, &bindings.TaikoL1ClientBlockProposed{Id: common.Big256}, 0, 0),
		"context deadline exceeded",
	)
}

func (s *ProverTestSuite) TestSubmitInvalidBlockProofThrowawayBlockNotFound() {
	s.Error(
		s.p.submitInvalidBlockProof(
			context.Background(), &producer.ProofWithHeader{
				BlockID: common.Big256,
				Meta:    &bindings.LibDataBlockMetadata{},
				Header:  &types.Header{},
				ZkProof: []byte{0xff},
			},
		),
	)
}
