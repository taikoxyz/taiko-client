package chainscanner

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

type UpdateCurrentFunc func(*types.Header)
type OnBlocksScannedFunc func(
	ctx context.Context,
	start, end *types.Header,
	updateCurrentFunc UpdateCurrentFunc,
) error

type ChainScanner struct {
	ctx                   context.Context
	client                *ethclient.Client
	chainID               *big.Int
	maxBlocksReadPerEpoch uint64
	startHeight           uint64
	endHeight             *uint64
	current               *types.Header
	onBlocksScanned       OnBlocksScannedFunc
}

type ChainScannerConfig struct {
	Client                *ethclient.Client
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	OnBlocksScanned       OnBlocksScannedFunc
}

func NewChainScanner(ctx context.Context, cfg *ChainScannerConfig) (*ChainScanner, error) {
	if cfg.Client == nil {
		return nil, errors.New("invalid RPC client")
	}

	if cfg.OnBlocksScanned == nil {
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

	cs := &ChainScanner{
		ctx:             ctx,
		client:          cfg.Client,
		chainID:         chainID,
		startHeight:     cfg.StartHeight.Uint64(),
		current:         startHeader,
		onBlocksScanned: cfg.OnBlocksScanned,
	}

	if cfg.MaxBlocksReadPerEpoch != nil {
		cs.maxBlocksReadPerEpoch = *cfg.MaxBlocksReadPerEpoch
	} else {
		cs.maxBlocksReadPerEpoch = DefaultMaxBlocksReadPerEpoch
	}

	if cfg.EndHeight != nil {
		endHeightUint64 := cfg.EndHeight.Uint64()
		cs.endHeight = &endHeightUint64
	}

	return cs, nil
}

func (cs *ChainScanner) Scan() error {
	scanOp := func() error {
		for {
			if err := cs.scan(); err != nil {
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
		case <-cs.ctx.Done():
			return cs.ctx.Err()
		default:
			if err := backoff.Retry(scanOp, backoff.NewExponentialBackOff()); err != nil {
				return err
			}
		}
	}
}

func (cs *ChainScanner) scan() error {
	if err := cs.ensureCurrentNotReorged(); err != nil {
		return fmt.Errorf("failed to check whether chainScanner.current has been reorged: %w", err)
	}

	head, err := cs.client.HeaderByNumber(cs.ctx, nil)
	if err != nil {
		return err
	}

	if cs.current.Number.Cmp(head.Number) == 0 {
		return nil
	}

	var (
		endHeader   = head
		isLastEpoch = true
	)
	if cs.endHeight != nil {
		endHeight := cs.current.Number.Uint64() + cs.maxBlocksReadPerEpoch

		if endHeight < head.Number.Uint64() {
			if endHeader, err = cs.client.HeaderByNumber(cs.ctx, new(big.Int).SetUint64(endHeight)); err != nil {
				return err
			}
			isLastEpoch = false
		}
	}

	if err := cs.onBlocksScanned(cs.ctx, cs.current, endHeader, cs.updateCurrent); err != nil {
		return err
	}

	if !isLastEpoch {
		return errContinue
	}

	return io.EOF
}

// updateCurrent updates the scanner's current cursor.
func (cs *ChainScanner) updateCurrent(current *types.Header) {
	if current == nil {
		log.Warn("Receive a nil header as chainScanner.current")
		return
	}

	cs.current = current
}

// ensureCurrentNotReorged checks if the chainScanner.current was reorged, if was, will
// rewind back `ReorgRewindDepth` blocks.
func (cs *ChainScanner) ensureCurrentNotReorged() error {
	current, err := cs.client.HeaderByHash(cs.ctx, cs.current.Hash())
	if err != nil && err != ethereum.NotFound {
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

	cs.current, err = cs.client.HeaderByNumber(cs.ctx, new(big.Int).SetUint64(newCurrentHeight))
	return err
}
