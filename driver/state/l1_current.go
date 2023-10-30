package state

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
)

// GetL1Current reads the L1 current cursor concurrent safely.
func (s *State) GetL1Current() *types.Header {
	return s.l1Current.Load().(*types.Header)
}

// SetL1Current sets the L1 current cursor concurrent safely.
func (s *State) SetL1Current(h *types.Header) {
	if h == nil {
		log.Warn("Empty l1 current cursor")
		return
	}
	log.Debug("Set L1 current cursor", "number", h.Number)
	s.l1Current.Store(h)
}

// ResetL1Current resets the l1Current cursor to the L1 height which emitted a
// BlockProposed event with given blockID / blockHash.
func (s *State) ResetL1Current(
	ctx context.Context,
	blockID *big.Int,
) (*bindings.TaikoL1ClientBlockProposed, error) {
	if blockID == nil {
		return nil, fmt.Errorf("empty block ID")
	}

	log.Info("Reset L1 current cursor", "blockID", blockID)

	if blockID.Cmp(common.Big0) == 0 {
		l1Current, err := s.rpc.L1.HeaderByNumber(ctx, s.GenesisL1Height)
		if err != nil {
			return nil, err
		}
		s.SetL1Current(l1Current)
		return nil, nil
	}

	var event *bindings.TaikoL1ClientBlockProposed
	iter, err := eventIterator.NewBlockProposedIterator(
		ctx,
		&eventIterator.BlockProposedIteratorConfig{
			Client:      s.rpc.L1,
			TaikoL1:     s.rpc.TaikoL1,
			StartHeight: s.GenesisL1Height,
			EndHeight:   s.GetL1Head().Number,
			FilterQuery: []*big.Int{blockID},
			Reverse:     true,
			OnBlockProposedEvent: func(
				ctx context.Context,
				e *bindings.TaikoL1ClientBlockProposed,
				end eventIterator.EndBlockProposedEventIterFunc,
			) error {
				event = e
				end()
				return nil
			},
		},
	)
	if err != nil {
		return nil, err
	}

	if err := iter.Iter(); err != nil {
		return nil, err
	}

	if event == nil {
		return nil, fmt.Errorf("BlockProposed event not found, blockID: %s", blockID)
	}

	l1Current, err := s.rpc.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(event.Raw.BlockNumber))
	if err != nil {
		return nil, err
	}
	s.SetL1Current(l1Current)

	log.Info("Reset L1 current cursor", "height", s.GetL1Current().Number)

	return event, nil
}
