package chainsyncer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/metrics"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	txListValidator "github.com/taikoxyz/taiko-client/pkg/tx_list_validator"
)

type State interface {
	GetConstants() struct {
		AnchorTxGasLimit  *big.Int
		MaxTxlistBytes    *big.Int
		MaxBlockNumTxs    *big.Int
		MaxBlocksGasLimit *big.Int
		MinTxGasLimit     *big.Int
	}
	GetL1Current() *types.Header
	GetHeadBlockID() *big.Int
	GetLastVerifiedBlock() struct {
		ID     *big.Int
		Hash   common.Hash
		Height *big.Int
	}
	SetL1Current(l1Current *types.Header)
	GetL1Head() *types.Header
	GetL2Head() *types.Header
	ResetL1Current(ctx context.Context, id *big.Int) error
}

type L2ChainSyncer struct {
	ctx                           context.Context
	state                         State                            // Driver's state
	rpc                           *rpc.Client                      // L1/L2 RPC clients
	throwawayBlocksBuilderPrivKey *ecdsa.PrivateKey                // Private key of L2 throwaway blocks builder
	txListValidator               *txListValidator.TxListValidator // Transactions list validator
	// Try P2P beacon-sync if current node is behind of  the protocol's latest verified block head
	p2pSyncVerifiedBlocks       bool
	lastSyncedVerifiedBlockHash common.Hash
	lastSyncedVerifiedBlockID   *big.Int
	beaconSyncTriggered         bool
}

// NewL2ChainSyncer creates a new chain syncer instance.
func NewL2ChainSyncer(
	ctx context.Context,
	rpc *rpc.Client,
	state State,
	throwawayBlocksBuilderPrivKey *ecdsa.PrivateKey,
	p2pSyncVerifiedBlocks bool,
) (*L2ChainSyncer, error) {
	return &L2ChainSyncer{
		ctx:                           ctx,
		rpc:                           rpc,
		state:                         state,
		throwawayBlocksBuilderPrivKey: throwawayBlocksBuilderPrivKey,
		txListValidator: txListValidator.NewTxListValidator(
			state.GetConstants().MaxBlocksGasLimit.Uint64(),
			state.GetConstants().MaxBlockNumTxs.Uint64(),
			state.GetConstants().MaxTxlistBytes.Uint64(),
			state.GetConstants().MinTxGasLimit.Uint64(),
			rpc.L2ChainID,
		),
		p2pSyncVerifiedBlocks: p2pSyncVerifiedBlocks,
	}, nil
}

// Sync performs a sync operation to L2 node's local chain.
func (s *L2ChainSyncer) Sync(l1End *types.Header) error {
	// If current L2 node's chain is behind of the protocol's latest verified block head, and the
	// `P2PSyncVerifiedBlocks` flag is set, try triggering a beacon-sync in L2 node to catch up the
	// latest verified block head.
	// TODO: check whether the engine is not syncing through P2P (eth_syncing)
	if s.p2pSyncVerifiedBlocks && s.state.GetLastVerifiedBlock().Height.Uint64() > 0 && !s.AheadOfProtocolVerifiedHead() {
		if err := s.TriggerBeaconSync(); err != nil {
			return fmt.Errorf("trigger beacon-sync error: %w", err)
		}

		return nil
	}

	if s.beaconSyncTriggered {
		l2Head, err := s.rpc.L2.HeaderByNumber(s.ctx, nil)
		if err != nil {
			return err
		}

		if l2Head.Hash() != s.lastSyncedVerifiedBlockHash {
			log.Crit(
				"Verified header mismatch, height: %d, hash: %s != %s",
				l2Head.Number, l2Head.Hash(), s.lastSyncedVerifiedBlockHash,
			)
		}

		if err := s.state.ResetL1Current(s.ctx, s.lastSyncedVerifiedBlockID); err != nil {
			return err
		}
	}

	return s.ProcessL1Blocks(s.ctx, l1End)
}

// AheadOfProtocolVerifiedHead checks whether the L2 chain is ahead of verified head in protocol.
func (s *L2ChainSyncer) AheadOfProtocolVerifiedHead() bool {
	return s.state.GetL2Head().Number.Cmp(s.state.GetLastVerifiedBlock().Height) >= 0
}

// ProcessL1Blocks fetches all `TaikoL1.BlockProposed` events between given
// L1 block heights, and then tries inserting them into L2 node's block chain.
func (s *L2ChainSyncer) ProcessL1Blocks(ctx context.Context, l1End *types.Header) error {
	iter, err := eventIterator.NewBlockProposedIterator(ctx, &eventIterator.BlockProposedIteratorConfig{
		Client:               s.rpc.L1,
		TaikoL1:              s.rpc.TaikoL1,
		StartHeight:          s.state.GetL1Current().Number,
		EndHeight:            l1End.Number,
		FilterQuery:          nil,
		OnBlockProposedEvent: s.onBlockProposed,
	})
	if err != nil {
		return err
	}

	if err := iter.Iter(); err != nil {
		return err
	}

	s.state.SetL1Current(l1End)
	metrics.DriverL1CurrentHeightGauge.Update(s.state.GetL1Current().Number.Int64())

	return nil
}
