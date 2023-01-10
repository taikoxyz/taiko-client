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
// BlockProven event with given blockID / blockHash.
func (s *State) ResetL1Current(ctx context.Context, heightOrID *HeightOrID) (*big.Int, error) {
	if !heightOrID.NotEmpty() {
		return nil, fmt.Errorf("empty input %v", heightOrID)
	}

	log.Info("Reset L1 current cursor", "heightOrID", heightOrID)

	var (
		l1CurrentHeight *big.Int
		err             error
	)

	if (heightOrID.ID != nil && heightOrID.ID.Cmp(common.Big0) == 0) ||
		(heightOrID.Height != nil && heightOrID.Height.Cmp(common.Big0) == 0) {
		l1Current, err := s.rpc.L1.HeaderByNumber(ctx, s.GenesisL1Height)
		if err != nil {
			return nil, err
		}
		s.SetL1Current(l1Current)
		return common.Big0, nil
	}

	// Need to find the block ID at first, before filtering the BlockProposed events.
	if heightOrID.ID == nil {
		header, err := s.rpc.L2.HeaderByNumber(context.Background(), heightOrID.Height)
		if err != nil {
			return nil, err
		}
		targetHash := header.Hash()

		iter, err := eventIterator.NewBlockProvenIterator(
			ctx,
			&eventIterator.BlockProvenIteratorConfig{
				Client:      s.rpc.L1,
				TaikoL1:     s.rpc.TaikoL1,
				StartHeight: s.GenesisL1Height,
				EndHeight:   s.GetL1Head().Number,
				FilterQuery: []*big.Int{},
				Reverse:     true,
				OnBlockProvenEvent: func(
					ctx context.Context,
					e *bindings.TaikoL1ClientBlockProven,
					end eventIterator.EndBlockProvenEventIterFunc,
				) error {
					log.Debug("Filtered BlockProven event", "ID", e.Id, "hash", common.Hash(e.BlockHash))
					if e.BlockHash == targetHash {
						heightOrID.ID = e.Id
						end()
					}

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

		if heightOrID.ID == nil {
			return nil, fmt.Errorf("BlockProven event not found, hash: %s", targetHash)
		}
	}

	iter, err := eventIterator.NewBlockProposedIterator(
		ctx,
		&eventIterator.BlockProposedIteratorConfig{
			Client:      s.rpc.L1,
			TaikoL1:     s.rpc.TaikoL1,
			StartHeight: s.GenesisL1Height,
			EndHeight:   s.GetL1Head().Number,
			FilterQuery: []*big.Int{heightOrID.ID},
			Reverse:     true,
			OnBlockProposedEvent: func(
				ctx context.Context,
				e *bindings.TaikoL1ClientBlockProposed,
				end eventIterator.EndBlockProposedEventIterFunc,
			) error {
				l1CurrentHeight = new(big.Int).SetUint64(e.Raw.BlockNumber)
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

	if l1CurrentHeight == nil {
		return nil, fmt.Errorf("BlockProposed event not found, blockID: %s", heightOrID.ID)
	}

	l1Current, err := s.rpc.L1.HeaderByNumber(ctx, l1CurrentHeight)
	if err != nil {
		return nil, err
	}
	s.SetL1Current(l1Current)

	log.Info("Reset L1 current cursor", "height", s.GetL1Current().Number)

	return heightOrID.ID, nil
}
