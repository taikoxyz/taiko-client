package eventscanner

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/taiko-client/bindings"
	chainscanner "github.com/taikoxyz/taiko-client/pkg/chain_scanner"
)

type OnBlockProposedEvent func(context.Context, *bindings.TaikoL1ClientBlockProposed) error

type BlockProposedScanner struct {
	ctx         context.Context
	taikoL1     *bindings.TaikoL1Client
	cs          *chainscanner.ChainScanner
	filterQuery []*big.Int
}

type BlockProposedScannerConfig struct {
	Client                *ethclient.Client
	TaikoL1               *bindings.TaikoL1Client
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	FilterQuery           []*big.Int
	OnBlockProposedEvent  OnBlockProposedEvent
}

func NewBlockProposedScanner(ctx context.Context, cfg *BlockProposedScannerConfig) (*BlockProposedScanner, error) {
	if cfg.OnBlockProposedEvent == nil {
		return nil, errors.New("invalid callback")
	}

	s := &BlockProposedScanner{
		ctx:         ctx,
		taikoL1:     cfg.TaikoL1,
		filterQuery: cfg.FilterQuery,
	}

	cs, err := chainscanner.NewChainScanner(ctx, &chainscanner.ChainScannerConfig{
		Client:                cfg.Client,
		MaxBlocksReadPerEpoch: cfg.MaxBlocksReadPerEpoch,
		StartHeight:           cfg.StartHeight,
		EndHeight:             cfg.EndHeight,
		OnBlocksScanned: assembleBlockProposedScannerCallback(
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

func (s *BlockProposedScanner) Scan() error {
	return s.cs.Scan()
}

func assembleBlockProposedScannerCallback(
	client *ethclient.Client,
	taikoL1Client *bindings.TaikoL1Client,
	filterQuery []*big.Int,
	callback OnBlockProposedEvent,
) chainscanner.OnBlocksScannedFunc {
	return func(ctx context.Context, start, end *types.Header, updateCurrentFunc chainscanner.UpdateCurrentFunc) error {
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
