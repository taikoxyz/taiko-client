package handler

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/internal/metrics"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	proofSubmitter "github.com/taikoxyz/taiko-client/prover/proof_submitter"
	state "github.com/taikoxyz/taiko-client/prover/shared_state"
)

var (
	errL1Reorged         = errors.New("L1 reorged")
	proofExpirationDelay = 1 * time.Minute
)

// BlockProposedEventHandler is responsible for handling the BlockProposed event as a prover.
type BlockProposedEventHandler struct {
	sharedState             *state.SharedState
	proverAddress           common.Address
	genesisHeightL1         uint64
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

// NewBlockProposedEventHandler creates a new BlockProposedEventHandler instance.
func NewBlockProposedEventHandler(
	sharedState *state.SharedState,
	proverAddress common.Address,
	genesisHeightL1 uint64,
	rpc *rpc.Client,
	proofGenerationCh chan *proofProducer.ProofWithHeader,
	proofWindowExpiredCh chan *bindings.TaikoL1ClientBlockProposed,
	proofSubmissionCh chan *proofSubmitter.GenerateProofRequest,
	proposeConcurrencyGuard chan struct{},
	BackOffRetryInterval time.Duration,
	backOffMaxRetrys uint64,
	isGuardian bool,
	contesterMode bool,
	proveUnassignedBlocks bool,
) BlockProposedHandler {
	handler := &BlockProposedEventHandler{
		sharedState,
		proverAddress,
		genesisHeightL1,
		rpc,
		proofGenerationCh,
		proofWindowExpiredCh,
		proofSubmissionCh,
		proposeConcurrencyGuard,
		BackOffRetryInterval,
		backOffMaxRetrys,
		contesterMode,
		proveUnassignedBlocks,
	}

	if !isGuardian {
		return handler
	}

	return &BlockProposedGuaridanEventHandler{*handler}
}

func (h *BlockProposedEventHandler) Handle(
	ctx context.Context,
	e *bindings.TaikoL1ClientBlockProposed,
	end eventIterator.EndBlockProposedEventIterFunc,
) error {
	// If there are newly generated proofs, we need to submit them as soon as possible.
	if len(h.proofGenerationCh) > 0 {
		log.Info("onBlockProposed callback early return", "proofGenerationChannelLength", len(h.proofGenerationCh))
		end()
		return nil
	}

	// Wait for the corresponding L2 block being mined in node.
	if _, err := h.rpc.WaitL1Origin(ctx, e.BlockId); err != nil {
		return fmt.Errorf("failed to wait L1Origin (eventID %d): %w", e.BlockId, err)
	}

	// Check if the L1 chain has reorged at first.
	if err := h.checkL1Reorg(ctx, e); err != nil {
		if errors.Is(err, errL1Reorged) {
			end()
			return nil
		}

		return err
	}

	// If the current block is handled, just skip it.
	if e.BlockId.Uint64() <= h.sharedState.GetLastHandledBlockID() {
		return nil
	}

	log.Info(
		"Proposed block",
		"l1Height", e.Raw.BlockNumber,
		"l1Hash", e.Raw.BlockHash,
		"blockID", e.BlockId,
		"removed", e.Raw.Removed,
		"assignedProver", e.AssignedProver,
		"livenessBond", e.LivenessBond,
		"minTier", e.Meta.MinTier,
	)
	metrics.ProverReceivedProposedBlockGauge.Update(e.BlockId.Int64())

	// Move l1Current cursor.
	newL1Current, err := h.rpc.L1.HeaderByHash(ctx, e.Raw.BlockHash)
	if err != nil {
		return err
	}
	h.sharedState.SetL1Current(newL1Current)
	h.sharedState.SetLastHandledBlockID(e.BlockId.Uint64())

	// Try generating a proof for the proposed block with the given backoff policy.
	go func() {
		if err := backoff.Retry(
			func() error {
				h.proposeConcurrencyGuard <- struct{}{}
				defer func() { <-h.proposeConcurrencyGuard }()

				if err := h.checkExpirationAndSubmitProof(ctx, e); err != nil {
					log.Error(
						"Failed to check proof status and submit proof",
						"error", err,
						"blockID", e.BlockId,
						"minTier", e.Meta.MinTier,
						"maxRetrys", h.backOffMaxRetrys,
					)
					return err
				}
				return nil
			},
			backoff.WithMaxRetries(backoff.NewConstantBackOff(h.BackOffRetryInterval), h.backOffMaxRetrys),
		); err != nil {
			log.Error("Handle new BlockProposed event error", "error", err)
		}
	}()

	return nil
}

// checkL1Reorg checks whether the L1 chain has been reorged.
func (h *BlockProposedEventHandler) checkL1Reorg(
	ctx context.Context,
	e *bindings.TaikoL1ClientBlockProposed,
) error {
	// Check whether the L2 EE's anchored L1 info, to see if the L1 chain has been reorged.
	reorged, l1CurrentToReset, lastHandledBlockIDToReset, err := h.rpc.CheckL1ReorgFromL2EE(
		ctx,
		new(big.Int).Sub(e.BlockId, common.Big1),
	)
	if err != nil {
		return fmt.Errorf("failed to check whether L1 chain was reorged from L2EE (eventID %d): %w", e.BlockId, err)
	}

	// Then check the l1Current cursor at first, to see if the L1 chain has been reorged.
	if !reorged {
		if reorged, l1CurrentToReset, lastHandledBlockIDToReset, err = h.rpc.CheckL1ReorgFromL1Cursor(
			ctx,
			h.sharedState.GetL1Current(),
			h.genesisHeightL1,
		); err != nil {
			return fmt.Errorf(
				"failed to check whether L1 chain was reorged from l1Current (eventID %d): %w",
				e.BlockId,
				err,
			)
		}
	}

	if reorged {
		log.Info(
			"Reset L1Current cursor due to reorg",
			"l1CurrentHeightOld", h.sharedState.GetL1Current().Number,
			"l1CurrentHeightNew", l1CurrentToReset.Number,
			"lastHandledBlockIDOld", h.sharedState.GetLastHandledBlockID(),
			"lastHandledBlockIDNew", lastHandledBlockIDToReset,
		)
		h.sharedState.SetL1Current(l1CurrentToReset)
		if lastHandledBlockIDToReset == nil {
			h.sharedState.SetLastHandledBlockID(0)
		} else {
			h.sharedState.SetLastHandledBlockID(lastHandledBlockIDToReset.Uint64())
		}
		h.sharedState.SetReorgDetectedFlag(true)
		return errL1Reorged
	}

	lastL1OriginHeader, err := h.rpc.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(e.Meta.L1Height))
	if err != nil {
		return fmt.Errorf("failed to get L1 header, height %d: %w", e.Meta.L1Height, err)
	}

	if lastL1OriginHeader.Hash() != e.Meta.L1Hash {
		log.Warn(
			"L1 block hash mismatch due to L1 reorg",
			"height", e.Meta.L1Height,
			"lastL1OriginHeader", lastL1OriginHeader.Hash(),
			"l1HashInEvent", e.Meta.L1Hash,
		)

		return fmt.Errorf(
			"L1 block hash mismatch due to L1 reorg: %s != %s",
			lastL1OriginHeader.Hash(),
			e.Meta.L1Hash,
		)
	}

	return nil
}

// checkExpirationAndSubmitProof checks whether the proposed block's proving window is expired,
// and submits a new proof if necessary.
func (h *BlockProposedEventHandler) checkExpirationAndSubmitProof(
	ctx context.Context,
	e *bindings.TaikoL1ClientBlockProposed,
) error {
	// Check whether the block has been verified.
	isVerified, err := isBlockVerified(ctx, h.rpc, e.BlockId)
	if err != nil {
		return fmt.Errorf("failed to check if the current L2 block is verified: %w", err)
	}
	if isVerified {
		log.Info("📋 Block has been verified", "blockID", e.BlockId)
		return nil
	}

	// Check whether the block's proof is still needed.
	proofStatus, err := rpc.GetBlockProofStatus(
		ctx,
		h.rpc,
		e.BlockId,
		h.proverAddress,
	)
	if err != nil {
		return fmt.Errorf("failed to check whether the L2 block needs a new proof: %w", err)
	}

	// If there is already a proof submitted on chain.
	if proofStatus.IsSubmitted {
		// If there is no need to contest the submitted proof, we skip proving this block here.
		if !proofStatus.Invalid {
			log.Info(
				"A valid proof has been submitted, skip proving",
				"blockID", e.BlockId,
				"parent", proofStatus.ParentHeader.Hash(),
			)
			return nil
		}

		// If there is an invalid proof, but current prover is not in contest mode, we skip proving this block.
		if !h.contesterMode {
			log.Info(
				"An invalid proof has been submitted, but current prover is not in contest mode, skip proving",
				"blockID", e.BlockId,
				"parent", proofStatus.ParentHeader.Hash(),
			)
			return nil
		}

		// The proof submitted to protocol is invalid.
		// TODO: Add contesting logic here.
		return nil
	}

	windowExpired, timeToExpire, err := isProvingWindowExpired(e, h.sharedState.GetTiers())
	if err != nil {
		return fmt.Errorf("failed to check if the proving window is expired: %w", err)
	}

	if windowExpired {
		// If the proving window is expired, we need to check if the current prover is the assigned prover
		// at first, if yes, we should skip proving this block, if no, then we check if the current prover
		// wants to prove unassigned blocks.
		log.Info(
			"Proposed block's proving window has expired",
			"blockID", e.BlockId,
			"prover", e.AssignedProver,
			"expiresAt", timeToExpire,
			"minTier", e.Meta.MinTier,
		)
		if e.AssignedProver == h.proverAddress {
			log.Warn(
				"Assigned prover is the current prover, but the proving window has expired, skip proving",
				"blockID", e.BlockId,
				"prover", e.AssignedProver,
			)
			return nil
		}
		// If the current prover doesn't want to prove unassigned blocks, we should skip proving this block.
		if !h.proveUnassignedBlocks {
			log.Info(
				"Skip proving expired blocks",
				"blockID", e.BlockId,
				"prover", e.AssignedProver,
			)
			return nil
		}
	} else {
		// If the proving window is not expired, we need to check if the current prover is the assigned prover,
		// if no and the current prover wants to prove unassigned blocks, then we should wait for its expiration.
		if e.AssignedProver != h.proverAddress {
			log.Info(
				"Proposed block is not provable by current prover at the moment",
				"blockID", e.BlockId,
				"prover", e.AssignedProver,
				"timeToExpire", timeToExpire,
			)

			if h.proveUnassignedBlocks {
				log.Info(
					"Add proposed block to wait for proof window expiration",
					"blockID", e.BlockId,
					"prover", e.AssignedProver,
					"timeToExpire", timeToExpire,
				)
				time.AfterFunc(
					// Add another 60 seconds, to ensure one more L1 block will be mined before the proof submission
					proofExpirationDelay,
					func() { h.proofWindowExpiredCh <- e },
				)
			}

			return nil
		}
	}

	tier := e.Meta.MinTier

	log.Info(
		"Proposed block is provable",
		"blockID", e.BlockId,
		"prover", e.AssignedProver,
		"minTier", e.Meta.MinTier,
		"currentTier", tier,
	)

	metrics.ProverProofsAssigned.Inc(1)

	h.proofSubmissionCh <- &proofSubmitter.GenerateProofRequest{Tier: tier, Event: e}

	return nil
}

type BlockProposedGuaridanEventHandler struct {
	BlockProposedEventHandler
}

func (h *BlockProposedGuaridanEventHandler) Handle(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	end eventIterator.EndBlockProposedEventIterFunc,
) error {
	// If we are operating as a guardian prover,
	// we should sign all seen proposed blocks as soon as possible.
	// go func() {
	// 	if err := p.guardianProverSender.SignAndSendBlock(ctx, event.BlockId); err != nil {
	// 		log.Error("Guardian prover unable to sign block", "blockID", event.BlockId, "error", err)
	// 	}
	// }()
	return h.BlockProposedEventHandler.Handle(ctx, event, end)
}
