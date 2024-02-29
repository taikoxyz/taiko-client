package handler

import (
	"context"
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
	state "github.com/taikoxyz/taiko-client/prover/shared_state"
)

// BlockProposedEventHandler is responsible for handling the BlockProposed event as a prover.
type BlockProposedEventHandler struct {
	sharedState             *state.SharedState
	genesisHeightL1         uint64
	rpc                     *rpc.Client
	proofGenerationCh       chan *proofProducer.ProofWithHeader
	proposeConcurrencyGuard chan struct{}
	BackOffRetryInterval    time.Duration
	backOffMaxRetrys        uint64
}

// NewBlockProposedEventHandler creates a new BlockProposedEventHandler instance.
func NewBlockProposedEventHandler(
	sharedState *state.SharedState,
	genesisHeightL1 uint64,
	rpc *rpc.Client,
	proofGenerationCh chan *proofProducer.ProofWithHeader,
	proposeConcurrencyGuard chan struct{},
	BackOffRetryInterval time.Duration,
	backOffMaxRetrys uint64,
	isGuardian bool,
) BlockProposedHandler {
	handler := &BlockProposedEventHandler{
		sharedState,
		genesisHeightL1,
		rpc,
		proofGenerationCh,
		proposeConcurrencyGuard,
		BackOffRetryInterval,
		backOffMaxRetrys,
	}

	if !isGuardian {
		return handler
	}

	return &BlockProposedGuaridanEventHandler{*handler}
}

func (h *BlockProposedEventHandler) OnBlockProposed(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	end eventIterator.EndBlockProposedEventIterFunc,
) error {
	// If there are newly generated proofs, we need to submit them as soon as possible.
	if len(h.proofGenerationCh) > 0 {
		log.Info("onBlockProposed callback early return", "proofGenerationChannelLength", len(h.proofGenerationCh))
		end()
		return nil
	}

	// Wait for the corresponding L2 block being mined in node.
	if _, err := h.rpc.WaitL1Origin(ctx, event.BlockId); err != nil {
		return fmt.Errorf("failed to wait L1Origin (eventID %d): %w", event.BlockId, err)
	}

	// Check whether the L2 EE's anchored L1 info, to see if the L1 chain has been reorged.
	reorged, l1CurrentToReset, lastHandledBlockIDToReset, err := h.rpc.CheckL1ReorgFromL2EE(
		ctx,
		new(big.Int).Sub(event.BlockId, common.Big1),
	)
	if err != nil {
		return fmt.Errorf("failed to check whether L1 chain was reorged from L2EE (eventID %d): %w", event.BlockId, err)
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
				event.BlockId,
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
		end()
		return nil
	}

	if event.BlockId.Uint64() <= h.sharedState.GetLastHandledBlockID() {
		return nil
	}

	lastL1OriginHeader, err := h.rpc.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(event.Meta.L1Height))
	if err != nil {
		return fmt.Errorf("failed to get L1 header, height %d: %w", event.Meta.L1Height, err)
	}

	if lastL1OriginHeader.Hash() != event.Meta.L1Hash {
		log.Warn(
			"L1 block hash mismatch due to L1 reorg",
			"height", event.Meta.L1Height,
			"lastL1OriginHeader", lastL1OriginHeader.Hash(),
			"l1HashInEvent", event.Meta.L1Hash,
		)

		return fmt.Errorf(
			"L1 block hash mismatch due to L1 reorg: %s != %s",
			lastL1OriginHeader.Hash(),
			event.Meta.L1Hash,
		)
	}

	log.Info(
		"Proposed block",
		"l1Height", event.Raw.BlockNumber,
		"l1Hash", event.Raw.BlockHash,
		"blockID", event.BlockId,
		"removed", event.Raw.Removed,
		"assignedProver", event.AssignedProver,
		"livenessBond", event.LivenessBond,
		"minTier", event.Meta.MinTier,
	)
	metrics.ProverReceivedProposedBlockGauge.Update(event.BlockId.Int64())

	// Move l1Current cursor.
	newL1Current, err := h.rpc.L1.HeaderByHash(ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}
	h.sharedState.SetL1Current(newL1Current)
	h.sharedState.SetLastHandledBlockID(event.BlockId.Uint64())

	// Try generating a proof for the proposed block with the given backoff policy.
	go func() {
		if err := backoff.Retry(
			func() error {
				h.proposeConcurrencyGuard <- struct{}{}
				defer func() { <-h.proposeConcurrencyGuard }()

				// TODO(David): fix this
				// if err := p.handleNewBlockProposedEvent(ctx, event); err != nil {
				// 	log.Error(
				// 		"Failed to handle BlockProposed event",
				// 		"error", err,
				// 		"blockID", event.BlockId,
				// 		"minTier", event.Meta.MinTier,
				// 		"maxRetrys", p.cfg.BackOffMaxRetrys,
				// 	)
				// 	return err
				// }
				return nil
			},
			backoff.WithMaxRetries(backoff.NewConstantBackOff(h.BackOffRetryInterval), h.backOffMaxRetrys),
		); err != nil {
			log.Error("Handle new BlockProposed event error", "error", err)
		}
	}()

	return nil
}

type BlockProposedGuaridanEventHandler struct {
	BlockProposedEventHandler
}

func (h *BlockProposedGuaridanEventHandler) OnBlockProposed(
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
	return h.BlockProposedEventHandler.OnBlockProposed(ctx, event, end)
}
