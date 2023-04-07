package submitter

import (
	"context"

	"github.com/taikoxyz/taiko-client/bindings"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
)

type ProofSubmitter interface {
	RequestProof(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error
	SubmitProof(ctx context.Context, proofWithHeader *proofProducer.ProofWithHeader, isOracle bool) error
}
