package chainscanner

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

const (
	DefaultMaxBlocksReadPerEpoch uint64 = 1000
)

var (
	errContinue = errors.New("continue")
)

type UpdateCurrentFunc func(*types.Header)
type HandlerFunc func(
	ctx context.Context,
	start *types.Header,
	end *types.Header,
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
	handlerFunc           HandlerFunc
}

type Config struct {
	Client                *ethclient.Client
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	HandlerFunc           HandlerFunc
}

func New(ctx context.Context, cfg *Config) (*ChainScanner, error) {
	if cfg.Client == nil {
		return nil, errors.New("invalid RPC client")
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
		ctx:         ctx,
		client:      cfg.Client,
		chainID:     chainID,
		startHeight: cfg.StartHeight.Uint64(),
		current:     startHeader,
		handlerFunc: cfg.HandlerFunc,
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

func (cs *ChainScanner) Start() error {
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

	if err := cs.handlerFunc(cs.ctx, cs.current, endHeader, cs.updateCurrent); err != nil {
		return err
	}

	if !isLastEpoch {
		return errContinue
	}

	return io.EOF
}

func (cs *ChainScanner) updateCurrent(current *types.Header) {
	if current == nil {
		log.Warn("Receive a nil header as chainScanner.current")
		return
	}

	cs.current = current
}
