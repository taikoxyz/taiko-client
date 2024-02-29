package handler

import (
	"context"

	"github.com/taikoxyz/taiko-client/bindings"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
)

type BlockProposedGuaridanEventHandler struct {
	proofGenerationCh chan struct{}
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
	return nil
}
