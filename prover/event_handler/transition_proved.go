package handler

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/internal/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	state "github.com/taikoxyz/taiko-client/prover/shared_state"
)

type TransitionProvedEventHandler struct {
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

func (h *TransitionProvedEventHandler) Handle(
	ctx context.Context,
	e *bindings.TaikoL1ClientTransitionProved,
) error {
	metrics.ProverReceivedProvenBlockGauge.Update(e.BlockId.Int64())

	// If this prover is in contest mode, we check the validity of this proof and if it's invalid,
	// contest it with a higher tier proof.
	if !h.contesterMode {
		return nil
	}

	isValidProof, err := isValidProof(
		ctx,
		h.rpc,
		e.BlockId,
		e.Tran.ParentHash,
		e.Tran.BlockHash,
		e.Tran.StateRoot,
	)
	if err != nil {
		return err
	}
	// If the proof is valid, we simply return.
	if isValidProof {
		return nil
	}

	// If the proof is invalid, we contest it.
	blockInfo, err := h.rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, e.BlockId.Uint64())
	if err != nil {
		return err
	}

	log.Info(
		"Contest a proven transition",
		"blockID", e.BlockId,
		"l1Height", blockInfo.Blk.ProposedIn,
		"tier", e.Tier,
		"parentHash", common.Bytes2Hex(e.Tran.ParentHash[:]),
		"blockHash", common.Bytes2Hex(e.Tran.BlockHash[:]),
		"stateRoot", common.Bytes2Hex(e.Tran.StateRoot[:]),
	)

	return nil
	// return p.requestProofByBlockID(e.BlockId, new(big.Int).SetUint64(blockInfo.Blk.ProposedIn), e.Tier, e)
}
