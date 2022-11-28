package chainiterator

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

const (
	DefaultBlocksReadPerEpoch = 1000
	ReorgRewindDepth          = 20
)

var (
	errContinue = errors.New("continue")
)

// OnBlocksFunc represents the callback function which will be called when a batch of blocks in chain are
// iterated.
type OnBlocksFunc func(
	ctx context.Context,
	start, end *types.Header,
	updateCurrentFunc UpdateCurrentFunc,
) error

// UpdateCurrentFunc updates the iterator.current cursor in the iterator.
type UpdateCurrentFunc func(*types.Header)

// BlockBatchIterator iterates the blocks in batches between the given start and end heights,
// with the awareness of reorganization.
type BlockBatchIterator struct {
	ctx                context.Context
	client             *ethclient.Client
	chainID            *big.Int
	blocksReadPerEpoch uint64
	startHeight        uint64
	endHeight          *uint64
	current            *types.Header
	onBlocks           OnBlocksFunc
}

// BlockBatchIteratorConfig represents the configs of a block batch iterator.
type BlockBatchIteratorConfig struct {
	Client                *ethclient.Client
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	OnBlocks              OnBlocksFunc
}

// NewBlockBatchIterator creates a new block batch iterator instance.
func NewBlockBatchIterator(ctx context.Context, cfg *BlockBatchIteratorConfig) (*BlockBatchIterator, error) {
	if cfg.Client == nil {
		return nil, errors.New("invalid RPC client")
	}

	if cfg.OnBlocks == nil {
		return nil, errors.New("invalid callback")
	}

	chainID, err := cfg.Client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID, error: %w", err)
	}

	if cfg.StartHeight == nil {
		return nil, errors.New("invalid start height")
	}

	if cfg.EndHeight != nil && cfg.StartHeight.Cmp(cfg.EndHeight) > 0 {
		return nil, fmt.Errorf("start height (%d) > end height (%d)", cfg.StartHeight, cfg.EndHeight)
	}

	startHeader, err := cfg.Client.HeaderByNumber(ctx, cfg.StartHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get start header, height: %s, error: %w", cfg.StartHeight, err)
	}

	iterator := &BlockBatchIterator{
		ctx:         ctx,
		client:      cfg.Client,
		chainID:     chainID,
		startHeight: cfg.StartHeight.Uint64(),
		current:     startHeader,
		onBlocks:    cfg.OnBlocks,
	}

	if cfg.MaxBlocksReadPerEpoch != nil {
		iterator.blocksReadPerEpoch = *cfg.MaxBlocksReadPerEpoch
	} else {
		iterator.blocksReadPerEpoch = DefaultBlocksReadPerEpoch
	}

	if cfg.EndHeight != nil {
		endHeightUint64 := cfg.EndHeight.Uint64()
		iterator.endHeight = &endHeightUint64
	}

	return iterator, nil
}

// Iter iterates the given chain between the given start and end heights,
// will call the callback when a batch of blocks in chain are iterated.
func (i *BlockBatchIterator) Iter() error {
	iterOp := func() error {
		for {
			if err := i.iter(); err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				if errors.Is(err, errContinue) {
					continue
				}
				return err
			}
		}
		return nil
	}

	for {
		select {
		case <-i.ctx.Done():
			return i.ctx.Err()
		default:
			return backoff.Retry(iterOp, backoff.NewExponentialBackOff())
		}
	}
}

// iter is the internal implementation of Iter, which performs the iteration.
func (i *BlockBatchIterator) iter() (err error) {
	if err := i.ensureCurrentNotReorged(); err != nil {
		return fmt.Errorf("failed to check whether iterator.current cursor has been reorged: %w", err)
	}

	var (
		endHeight   uint64
		endHeader   *types.Header
		destHeight  uint64
		isLastEpoch bool
	)

	if i.endHeight != nil {
		destHeight = *i.endHeight
	} else {
		if destHeight, err = i.client.BlockNumber(i.ctx); err != nil {
			return err
		}
	}

	if i.current.Number.Uint64() >= destHeight {
		return io.EOF
	}

	endHeight = i.current.Number.Uint64() + i.blocksReadPerEpoch

	if endHeight >= destHeight {
		endHeight = destHeight
		isLastEpoch = true
	}

	if endHeader, err = i.client.HeaderByNumber(i.ctx, new(big.Int).SetUint64(endHeight)); err != nil {
		return err
	}

	if err := i.onBlocks(i.ctx, i.current, endHeader, i.updateCurrent); err != nil {
		return err
	}

	i.current = endHeader

	if !isLastEpoch {
		return errContinue
	}

	return io.EOF
}

// updateCurrent updates the iterator's current cursor.
func (i *BlockBatchIterator) updateCurrent(current *types.Header) {
	if current == nil {
		log.Warn("Receive a nil header as iterator.current cursor")
		return
	}

	i.current = current
}

// ensureCurrentNotReorged checks if the iterator.current cursor was reorged, if was, will
// rewind back `ReorgRewindDepth` blocks.
func (i *BlockBatchIterator) ensureCurrentNotReorged() error {
	current, err := i.client.HeaderByHash(i.ctx, i.current.Hash())
	if err != nil && !errors.Is(err, ethereum.NotFound) {
		return err
	}

	// Not reorged
	if current != nil {
		return nil
	}

	// Reorg detected, rewind back `ReorgRewindDepth` blocks
	var newCurrentHeight uint64
	if current.Number.Uint64() <= ReorgRewindDepth {
		newCurrentHeight = 0
	} else {
		newCurrentHeight = current.Number.Uint64() - ReorgRewindDepth
	}

	i.current, err = i.client.HeaderByNumber(i.ctx, new(big.Int).SetUint64(newCurrentHeight))
	return err
}
