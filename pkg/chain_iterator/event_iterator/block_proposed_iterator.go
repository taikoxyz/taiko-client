package eventiterator

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/taiko-client/bindings"
	chainiterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator"
)

// OnBlockProposedEvent represents the callback function which will be called when a TaikoL1.BlockProposed event is
// iterated.
type OnBlockProposedEvent func(context.Context, *bindings.TaikoL1ClientBlockProposed) error

// BlockProposedIterator iterates the emitted TaikoL1.BlockProposed events in the chain,
// with the awareness of reorganization.
type BlockProposedIterator struct {
	ctx         context.Context
	taikoL1     *bindings.TaikoL1Client
	cs          *chainiterator.ChainIterator
	filterQuery []*big.Int
}

// BlockProposedIteratorConfig represents the configs of a BlockProposed event iterator.
type BlockProposedIteratorConfig struct {
	Client                *ethclient.Client
	TaikoL1               *bindings.TaikoL1Client
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	FilterQuery           []*big.Int
	OnBlockProposedEvent  OnBlockProposedEvent
}

// NewBlockProposedIterator creates a new instance of BlockProposed event iterator.
func NewBlockProposedIterator(ctx context.Context, cfg *BlockProposedIteratorConfig) (*BlockProposedIterator, error) {
	if cfg.OnBlockProposedEvent == nil {
		return nil, errors.New("invalid callback")
	}

	s := &BlockProposedIterator{
		ctx:         ctx,
		taikoL1:     cfg.TaikoL1,
		filterQuery: cfg.FilterQuery,
	}

	cs, err := chainiterator.NewChainIterator(ctx, &chainiterator.ChainIteratorConfig{
		Client:                cfg.Client,
		MaxBlocksReadPerEpoch: cfg.MaxBlocksReadPerEpoch,
		StartHeight:           cfg.StartHeight,
		EndHeight:             cfg.EndHeight,
		OnBlocks: assembleBlockProposedIteratorCallback(
			cfg.Client,
			cfg.TaikoL1,
			cfg.FilterQuery,
			cfg.OnBlockProposedEvent,
		),
	})
	if err != nil {
		return nil, err
	}

	s.cs = cs

	return s, nil
}

// Iter iterates the given chain between the given start and end heights,
// will call the callback when a BlockProposed event is iterated.
func (s *BlockProposedIterator) Iter() error {
	return s.cs.Iter()
}

// assembleBlockProposedIteratorCallback assembles the callback which will be used
// by a event iterator's inner chain iterator.
func assembleBlockProposedIteratorCallback(
	client *ethclient.Client,
	taikoL1Client *bindings.TaikoL1Client,
	filterQuery []*big.Int,
	callback OnBlockProposedEvent,
) chainiterator.OnBlocksFunc {
	return func(ctx context.Context, start, end *types.Header, updateCurrentFunc chainiterator.UpdateCurrentFunc) error {
		endHeight := end.Number.Uint64()
		iter, err := taikoL1Client.FilterBlockProposed(
			&bind.FilterOpts{Start: start.Number.Uint64(), End: &endHeight},
			filterQuery,
		)
		if err != nil {
			return err
		}

		for iter.Next() {
			if ctx.Err() != nil {
				return nil
			}

			event := iter.Event

			// Since we are not using eth_subscribe, this should not happen,
			// only check for safety.
			if event.Raw.Removed {
				continue
			}

			if err := callback(ctx, event); err != nil {
				return err
			}

			current, err := client.HeaderByHash(ctx, event.Raw.BlockHash)
			if err != nil {
				return err
			}

			updateCurrentFunc(current)
		}

		return nil
	}
}
