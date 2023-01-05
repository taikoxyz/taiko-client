package proofSubmitter

import (
	"context"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/prover/producer"
)

type ProofSubmitter interface {
	RequestProof(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error
	SubmitProof(ctx context.Context, proofWithHeader *producer.ProofWithHeader) error
}
