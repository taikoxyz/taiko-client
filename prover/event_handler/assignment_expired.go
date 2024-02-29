package handler

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	state "github.com/taikoxyz/taiko-client/prover/shared_state"
)

// AssignmentExpiredEventHandler is responsible for handling the expiration of proof assignments.
type AssignmentExpiredEventHandler struct {
	sharedState             *state.SharedState
	proverAddress           common.Address
	rpc                     *rpc.Client
	proofGenerationCh       chan *proofProducer.ProofWithHeader
	proofWindowExpiredCh    chan *bindings.TaikoL1ClientBlockProposed
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

		// return p.handleInvalidProof(
		// 	ctx,
		// 	e.BlockId,
		// 	new(big.Int).SetUint64(e.Raw.BlockNumber),
		// 	proofStatus.ParentHeader.Hash(),
		// 	proofStatus.CurrentTransitionState.Contester,
		// 	&e.Meta,
		// 	proofStatus.CurrentTransitionState.Tier,
		// )
		return nil
	}

	// return p.requestProofByBlockID(e.BlockId, new(big.Int).SetUint64(e.Raw.BlockNumber), e.Meta.MinTier, nil)
	// return requestProofByBlockID(ctx, e.BlockId, h.rpc, e, nil)
	return nil
}
