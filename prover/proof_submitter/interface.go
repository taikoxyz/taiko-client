package submitter

import (
	"context"

	bindings "github.com/taikoxyz/taiko-client/bindings/taikol1"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
)

type Submitter interface {
	RequestProof(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error
	SubmitProof(ctx context.Context, proofWithHeader *proofProducer.ProofWithHeader) error
	Producer() proofProducer.ProofProducer
	Tier() uint16
}

type Contester interface {
	SubmitContest(
		ctx context.Context,
		blockProposedEvent *bindings.TaikoL1ClientBlockProposed,
		transitionProvedEvent *bindings.TaikoL1ClientTransitionProved,
	) error
}
