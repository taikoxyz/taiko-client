package handler

import (
	"context"

	"github.com/taikoxyz/taiko-client/bindings"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
)

type BlockProposedHandler interface {
	OnBlockProposed(
		ctx context.Context,
		event *bindings.TaikoL1ClientBlockProposed,
		end eventIterator.EndBlockProposedEventIterFunc,
	) error
}
