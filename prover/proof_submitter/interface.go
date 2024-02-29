package submitter

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/taikoxyz/taiko-client/bindings"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
)

// GenerateProofRequest represents a request body to generate a proof.
type GenerateProofRequest struct {
	Tier  uint16
	Event *bindings.TaikoL1ClientBlockProposed
}

// Submitter is the interface for submitting proofs of the L2 blocks.
type Submitter interface {
	RequestProof(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error
	SubmitProof(ctx context.Context, proofWithHeader *proofProducer.ProofWithHeader) error
	Producer() proofProducer.ProofProducer
	Tier() uint16
}

// Contester is the interface for contesting proofs of the L2 blocks.
type Contester interface {
	SubmitContest(
		ctx context.Context,
		blockID *big.Int,
		proposedIn *big.Int,
		parentHash common.Hash,
		meta *bindings.TaikoDataBlockMetadata,
		tier uint16,
	) error
}
