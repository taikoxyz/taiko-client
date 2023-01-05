package proofSubmitter

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/prover/producer"
)

func (s *ProofSubmitterTestSuite) TestProveBlockInvalidL1OriginTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s.ErrorContains(
		s.invalidProofSubmitter.RequestProof(ctx, &bindings.TaikoL1ClientBlockProposed{Id: common.Big256}),
		"context deadline exceeded",
	)
}

func (s *ProofSubmitterTestSuite) TestSubmitInvalidBlockProofThrowawayBlockNotFound() {
	s.Error(
		s.invalidProofSubmitter.SubmitProof(
			context.Background(), &producer.ProofWithHeader{
				BlockID: common.Big256,
				Meta:    &bindings.TaikoDataBlockMetadata{},
				Header:  &types.Header{},
				ZkProof: []byte{0xff},
			},
		),
	)
}
