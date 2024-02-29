package handler

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	proofSubmitter "github.com/taikoxyz/taiko-client/prover/proof_submitter"
	state "github.com/taikoxyz/taiko-client/prover/shared_state"
)

type TransitionContestedEventHandler struct {
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

func (h *TransitionContestedEventHandler) Handle(
	ctx context.Context,
	e *bindings.TaikoL1ClientTransitionContested,
) error {
	log.Info(
		"🗡 Transition contested",
		"blockID", e.BlockId,
		"parentHash", common.Bytes2Hex(e.Tran.ParentHash[:]),
		"hash", common.Bytes2Hex(e.Tran.BlockHash[:]),
		"stateRoot", common.BytesToHash(e.Tran.StateRoot[:]),
		"contester", e.Contester,
		"bond", e.ContestBond,
	)

	// If this prover is not in contester mode, we simply output a log and return.
	if !h.contesterMode {
		return nil
	}

	contestedTransition, err := h.rpc.TaikoL1.GetTransition(
		&bind.CallOpts{Context: ctx},
		e.BlockId.Uint64(),
		e.Tran.ParentHash,
	)
	if err != nil {
		return err
	}

	// Compare the contested transition to the block in local L2 canonical chain.
	isValidProof, err := isValidProof(
		ctx,
		h.rpc,
		e.BlockId,
		e.Tran.ParentHash,
		contestedTransition.BlockHash,
		contestedTransition.StateRoot,
	)
	if err != nil {
		return err
	}
	if isValidProof {
		log.Info(
			"Contested transition is valid to local canonical chain, ignore the contest",
			"blockID", e.BlockId,
			"parentHash", common.Bytes2Hex(e.Tran.ParentHash[:]),
			"hash", common.Bytes2Hex(contestedTransition.BlockHash[:]),
			"stateRoot", common.BytesToHash(contestedTransition.StateRoot[:]),
			"contester", e.Contester,
			"bond", e.ContestBond,
		)
		return nil
	}

	h.proofSubmissionCh <- &proofSubmitter.GenerateProofRequest{
		Tier:  e.Tier + 1,
		Event: nil, // TODO
	}

	return nil
}
