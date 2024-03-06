package handler

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	proofSubmitter "github.com/taikoxyz/taiko-client/prover/proof_submitter"
)

// TransitionContestedEventHandler is responsible for handling the TransitionContested event.
type TransitionContestedEventHandler struct {
	rpc               *rpc.Client
	proofSubmissionCh chan *proofSubmitter.ProofRequestBody
	contesterMode     bool
}

// NewTransitionContestedEventHandler creates a new TransitionContestedEventHandler instance.
func NewTransitionContestedEventHandler(
	rpc *rpc.Client,
	proofSubmissionCh chan *proofSubmitter.ProofRequestBody,
	contesterMode bool,
) *TransitionContestedEventHandler {
	return &TransitionContestedEventHandler{rpc, proofSubmissionCh, contesterMode}
}

// Handle implements the TransitionContestedHandler interface.
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

	// If the proof is invalid, we contest it.
	blockInfo, err := h.rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, e.BlockId.Uint64())
	if err != nil {
		return err
	}

	blockProposedEvent, err := getBlockProposedEventFromBlockID(
		ctx,
		h.rpc,
		e.BlockId,
		new(big.Int).SetUint64(blockInfo.Blk.ProposedIn),
	)
	if err != nil {
		return err
	}

	go func() {
		h.proofSubmissionCh <- &proofSubmitter.ProofRequestBody{
			Tier:  e.Tier + 1,
			Event: blockProposedEvent,
		}
	}()

	return nil
}
