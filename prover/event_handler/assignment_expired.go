package handler

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	proofSubmitter "github.com/taikoxyz/taiko-client/prover/proof_submitter"
	state "github.com/taikoxyz/taiko-client/prover/shared_state"
)

// AssignmentExpiredEventHandler is responsible for handling the expiration of proof assignments.
type AssignmentExpiredEventHandler struct {
	sharedState             *state.SharedState
	proverAddress           common.Address
	rpc                     *rpc.Client
	proofGenerationCh       chan *proofProducer.ProofWithHeader
	proofWindowExpiredCh    chan *bindings.TaikoL1ClientBlockProposed
	proofSubmissionCh       chan *proofSubmitter.GenerateProofRequest
	proposeConcurrencyGuard chan struct{}
	BackOffRetryInterval    time.Duration
	backOffMaxRetrys        uint64
	contesterMode           bool
	proveUnassignedBlocks   bool
}

func (h *BlockProposedEventHandler) OnAssignmentExpired(
	ctx context.Context,
	e *bindings.TaikoL1ClientBlockProposed,
) error {
	log.Info(
		"Proof assignment window is expired",
		"blockID", e.BlockId,
		"assignedProver", e.AssignedProver,
		"minTier", e.Meta.MinTier,
	)
	// If Proof assignment window is expired, then the assigned prover can not submit new proofs for it anymore.
	if h.proverAddress == e.AssignedProver {
		return nil
	}
	// Check if we still need to generate a new proof for that block.
	proofStatus, err := rpc.GetBlockProofStatus(ctx, h.rpc, e.BlockId, h.proverAddress)
	if err != nil {
		return err
	}
	if proofStatus.IsSubmitted {
		// If there is already a proof submitted and there is no need to contest
		// it, we skip proving this block here.
		if !proofStatus.Invalid || !h.contesterMode {
			return nil
		}

		// If there is no contester, we submit a contest to protocol.
		if proofStatus.CurrentTransitionState.Contester == rpc.ZeroAddress {
			// TODO
			return nil
		}

		h.proofSubmissionCh <- &proofSubmitter.GenerateProofRequest{
			Tier:  proofStatus.CurrentTransitionState.Tier + 1,
			Event: e,
		}
		return nil
	}

	h.proofSubmissionCh <- &proofSubmitter.GenerateProofRequest{Tier: e.Meta.MinTier, Event: e}
	return nil
}
