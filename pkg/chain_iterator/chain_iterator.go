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
	DefaultMaxBlocksReadPerEpoch = 1000
	ReorgRewindDepth             = 20
)

var (
	errContinue = errors.New("continue")
)

// OnBlocksFunc represents the callback function which will be called when blocks in chain are
// iterated.
type OnBlocksFunc func(
	ctx context.Context,
	start, end *types.Header,
	updateCurrentFunc UpdateCurrentFunc,
) error

// UpdateCurrentFunc updates the current block cursor in iterator.
type UpdateCurrentFunc func(*types.Header)

// ChainIterator iterates the blocks between the given start and end heights in a chain,
// with the awareness of reorganization.
type ChainIterator struct {
	ctx                   context.Context
	client                *ethclient.Client
	chainID               *big.Int
	maxBlocksReadPerEpoch uint64
	startHeight           uint64
	endHeight             *uint64
	current               *types.Header
	onBlocks              OnBlocksFunc
}

// ChainIteratorConfig represents the configs of a chain iterator.
type ChainIteratorConfig struct {
	Client                *ethclient.Client
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	OnBlocks              OnBlocksFunc
}

// NewChainIterator creates a new chain iterator instance.
func NewChainIterator(ctx context.Context, cfg *ChainIteratorConfig) (*ChainIterator, error) {
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

	ci := &ChainIterator{
		ctx:         ctx,
		client:      cfg.Client,
		chainID:     chainID,
		startHeight: cfg.StartHeight.Uint64(),
		current:     startHeader,
		onBlocks:    cfg.OnBlocks,
	}

	if cfg.MaxBlocksReadPerEpoch != nil {
		ci.maxBlocksReadPerEpoch = *cfg.MaxBlocksReadPerEpoch
	} else {
		ci.maxBlocksReadPerEpoch = DefaultMaxBlocksReadPerEpoch
	}

	if cfg.EndHeight != nil {
		endHeightUint64 := cfg.EndHeight.Uint64()
		ci.endHeight = &endHeightUint64
	}

	return ci, nil
}

// Iter iterates the given chain between the given start and end heights,
// will call the callback when blocks in chain are iterated.
func (ci *ChainIterator) Iter() error {
	iterOp := func() error {
		for {
			if err := ci.iter(); err != nil {
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
		case <-ci.ctx.Done():
			return ci.ctx.Err()
		default:
			return backoff.Retry(iterOp, backoff.NewExponentialBackOff())
		}
	}
}

func (ci *ChainIterator) iter() error {
	if err := ci.ensureCurrentNotReorged(); err != nil {
		return fmt.Errorf("failed to check whether chainScanner.current has been reorged: %w", err)
	}

	head, err := ci.client.HeaderByNumber(ci.ctx, nil)
	if err != nil {
		return err
	}

	if ci.current.Number.Cmp(head.Number) == 0 {
		return nil
	}

	var (
		endHeader   = head
		isLastEpoch = true
	)
	if ci.endHeight != nil {
		endHeight := ci.current.Number.Uint64() + ci.maxBlocksReadPerEpoch

		if endHeight < head.Number.Uint64() {
			if endHeader, err = ci.client.HeaderByNumber(ci.ctx, new(big.Int).SetUint64(endHeight)); err != nil {
				return err
			}
			isLastEpoch = false
		}
	}

	if err := ci.onBlocks(ci.ctx, ci.current, endHeader, ci.updateCurrent); err != nil {
		return err
	}

	ci.current = endHeader

	if !isLastEpoch {
		return errContinue
	}

	return io.EOF
}

// updateCurrent updates the scanner's current cursor.
func (ci *ChainIterator) updateCurrent(current *types.Header) {
	if current == nil {
		log.Warn("Receive a nil header as chainScanner.current")
		return
	}

	ci.current = current
}

// ensureCurrentNotReorged checks if the chainScanner.current was reorged, if was, will
// rewind back `ReorgRewindDepth` blocks.
func (ci *ChainIterator) ensureCurrentNotReorged() error {
	current, err := ci.client.HeaderByHash(ci.ctx, ci.current.Hash())
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

	ci.current, err = ci.client.HeaderByNumber(ci.ctx, new(big.Int).SetUint64(newCurrentHeight))
	return err
}
