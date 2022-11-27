package driver

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sync/atomic"

	"github.com/cenkalti/backoff/v4"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type VerifiedHeader struct {
	ID     *big.Int
	Hash   common.Hash
	Height *big.Int
}

type State struct {
	// Subscriptions, will automatically resubscribe on errors
	l1HeadSub          event.Subscription // L1 new heads
	l2HeadSub          event.Subscription // L2 new heads
	l2BlockVerifiedSub event.Subscription // TaikoL1.BlockVerified events
	l2BlockProposedSub event.Subscription // TaikoL1.BlockProposed events

	// Feeds
	l1HeadsFeed event.Feed // L1 new heads notification feed

	l1Head         *atomic.Value // Latest known L1 head
	l2Head         *atomic.Value // Current L2 node head
	l2HeadBlockID  *atomic.Value // Latest known L2 block ID
	l2VerifiedHead *atomic.Value // Latest known L2 verified head
	l1Current      *types.Header // Current L1 block sync cursor

	// Constants
	anchorTxGasLimit  *big.Int
	maxTxlistBytes    *big.Int
	maxBlockNumTxs    *big.Int
	maxBlocksGasLimit *big.Int
	minTxGasLimit     *big.Int

	rpc *rpc.Client // L1/L2 RPC clients
}

// NewState creates a new driver state instance.
func NewState(ctx context.Context, rpc *rpc.Client) (*State, error) {
	// Set the L2 head's latest known L1 origin as current L1 sync cursor.
	latestL2KnownL1Header, err := rpc.LatestL2KnownL1Header(ctx)
	if err != nil {
		return nil, err
	}

	_, _, _, _, _, _,
		maxBlocksGasLimit, maxBlockNumTxs, _, maxTxlistBytes, minTxGasLimit,
		anchorTxGasLimit, _, _, err := rpc.TaikoL1.GetConstants(nil)
	if err != nil {
		return nil, err
	}

	log.Info(
		"Taiko protocol constants",
		"anchorTxGasLimit", anchorTxGasLimit,
		"maxTxlistBytes", maxTxlistBytes,
		"maxBlockNumTxs", maxBlockNumTxs,
		"maxBlocksGasLimit", maxBlocksGasLimit,
		"minTxGasLimit", minTxGasLimit,
	)

	s := &State{
		rpc:               rpc,
		anchorTxGasLimit:  anchorTxGasLimit,
		maxTxlistBytes:    maxTxlistBytes,
		maxBlockNumTxs:    maxBlockNumTxs,
		maxBlocksGasLimit: maxBlocksGasLimit,
		minTxGasLimit:     minTxGasLimit,
		l1Head:            new(atomic.Value),
		l2HeadBlockID:     new(atomic.Value),
		l2VerifiedHead:    new(atomic.Value),
		l1Current:         latestL2KnownL1Header,
	}

	if err := s.initSubscriptions(ctx); err != nil {
		return nil, err
	}

	return s, nil
}

// Close closes all inner subscriptions.
func (s *State) Close() {
	s.l1HeadSub.Unsubscribe()
	s.l2BlockVerifiedSub.Unsubscribe()
	s.l2BlockProposedSub.Unsubscribe()
}

// initSubscriptions initializes all subscriptions in the given state instance.
func (s *State) initSubscriptions(ctx context.Context) error {
	// L1 head
	l1Head, err := s.rpc.L1.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	s.setL1Head(l1Head)

	s.l1HeadSub = event.ResubscribeErr(
		backoff.DefaultMaxInterval,
		func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				log.Warn("Failed to subscribe L1 head, try resubscribing", "error", err)
			}

			return s.watchL1Head(ctx)
		},
	)

	// L2 head
	l2Head, err := s.rpc.L2.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	s.setL2Head(l2Head)

	s.l2HeadSub = event.ResubscribeErr(
		backoff.DefaultMaxInterval,
		func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				log.Warn("Failed to subscribe L2 head, try resubscribing", "error", err)
			}

			return s.watchL2Head(ctx)
		},
	)

	// TaikoL1.BlockVerified events
	_, lastVerifiedHeight, lastVerifiedId, _, err := s.rpc.TaikoL1.GetStateVariables(nil)
	if err != nil {
		return err
	}

	lastVerifiedBlockHash, err := s.rpc.TaikoL1.GetSyncedHeader(
		nil,
		new(big.Int).SetUint64(lastVerifiedHeight),
	)
	if err != nil {
		return err
	}

	s.setLastVerifiedBlockHash(
		new(big.Int).SetUint64(lastVerifiedId),
		new(big.Int).SetUint64(lastVerifiedHeight),
		lastVerifiedBlockHash,
	)

	s.l2BlockVerifiedSub = event.ResubscribeErr(
		backoff.DefaultMaxInterval,
		func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				log.Warn("Failed to subscribe TaikoL1.BlockVerifiedId events, try resubscribing", "error", err)
			}

			return s.watchBlockVerified(ctx)
		},
	)

	// TaikoL1.BlockProposed events
	_, _, _, nextPendingID, err := s.rpc.TaikoL1.GetStateVariables(nil)
	if err != nil {
		return err
	}

	s.setHeadBlockID(new(big.Int).SetUint64(nextPendingID - 1))

	s.l2BlockProposedSub = event.ResubscribeErr(
		backoff.DefaultMaxInterval,
		func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				log.Warn("Failed to subscribe TaikoL1.BlockProposed events, try resubscribing", "error", err)
			}

			return s.watchBlockProposed(ctx)
		},
	)

	return nil
}

// watchL1Head watches new L1 head blocks and keep updating current
// driver state.
func (s *State) watchL1Head(ctx context.Context) (event.Subscription, error) {
	newL1HeadCh := make(chan *types.Header, 10)

	sub, err := s.rpc.L1.SubscribeNewHead(ctx, newL1HeadCh)
	if err != nil {
		log.Error("Create L1 head subscription error", "error", err)
		return nil, err
	}

	defer sub.Unsubscribe()

	for {
		select {
		case newHead := <-newL1HeadCh:
			s.setL1Head(newHead)
			s.l1HeadsFeed.Send(newHead)
		case err := <-sub.Err():
			return sub, err
		case <-ctx.Done():
			return sub, nil
		}
	}
}

// setL1Head sets the L1 head concurrent safely.
func (s *State) setL1Head(l1Head *types.Header) {
	if l1Head == nil {
		log.Warn("Empty new L1 head")
		return
	}

	log.Debug("New L1 head", "height", l1Head.Number, "hash", l1Head.Hash(), "timestamp", l1Head.Time)
	metrics.DriverL1HeadHeightGauge.Update(l1Head.Number.Int64())

	s.l1Head.Store(l1Head)
}

// GetL1Head reads the L1 head concurrent safely.
func (s *State) GetL1Head() *types.Header {
	return s.l1Head.Load().(*types.Header)
}

// watchL2Head watches new L2 head blocks and keep updating current
// driver state.
func (s *State) watchL2Head(ctx context.Context) (event.Subscription, error) {
	newL2HeadCh := make(chan *types.Header, 10)

	sub, err := s.rpc.L2.SubscribeNewHead(ctx, newL2HeadCh)
	if err != nil {
		log.Error("Create L2 head subscription error", "error", err)
		return nil, err
	}

	defer sub.Unsubscribe()

	for {
		select {
		case newHead := <-newL2HeadCh:
			s.setL2Head(newHead)
		case err := <-sub.Err():
			return sub, err
		case <-ctx.Done():
			return sub, nil
		}
	}
}

// setL1Head sets the L2 head concurrent safely.
func (s *State) setL2Head(l2Head *types.Header) {
	if l2Head == nil {
		log.Warn("Empty new L2 head")
		return
	}

	log.Debug("New L2 head", "height", l2Head.Number, "hash", l2Head.Hash(), "timestamp", l2Head.Time)
	metrics.DriverL2HeadHeightGauge.Update(l2Head.Number.Int64())

	s.l2Head.Store(l2Head)
}

// GetL2Head reads the L2 head concurrent safely.
func (s *State) GetL2Head() *types.Header {
	return s.l2Head.Load().(*types.Header)
}

// watchBlockVerified watches newly verified blocks and keep updating current
// driver state.
func (s *State) watchBlockVerified(ctx context.Context) (ethereum.Subscription, error) {
	newHeaderSyncedCh := make(chan *bindings.TaikoL1ClientHeaderSynced, 10)

	sub, err := s.rpc.TaikoL1.WatchHeaderSynced(nil, newHeaderSyncedCh, nil, nil)
	if err != nil {
		log.Error("Create TaikoL1.HeaderSynced subscription error", "error", err)
		return nil, err
	}

	defer sub.Unsubscribe()

	for {
		select {
		case e := <-newHeaderSyncedCh:
			if err := s.VerfiyL2Block(ctx, e.SrcHash); err != nil {
				log.Error("Check new verified L2 block error", "error", err)
				continue
			}
			id, err := s.getSyncedHeaderID(e.Raw.BlockNumber, e.SrcHash)
			if err != nil {
				log.Error("Get synced header block ID error", "error", err)
				continue
			}
			s.setLastVerifiedBlockHash(id, e.SrcHeight, e.SrcHash)
		case err := <-sub.Err():
			return sub, err
		case <-ctx.Done():
			return sub, nil
		}
	}
}

// setLastVerifiedBlockHash sets the last verified L2 block hash concurrent safely.
func (s *State) setLastVerifiedBlockHash(id *big.Int, height *big.Int, hash common.Hash) {
	log.Debug("New verified block", "height", height, "hash", hash)
	metrics.DriverL2VerifiedHeightGauge.Update(height.Int64())
	s.l2VerifiedHead.Store(&VerifiedHeader{ID: id, Height: height, Hash: hash})
}

// GetLastVerifiedBlockHash reads the last verified L2 block concurrent safely.
// TODO: use a defined type.
func (s *State) GetLastVerifiedBlock() struct {
	ID     *big.Int
	Hash   common.Hash
	Height *big.Int
} {
	h := s.l2VerifiedHead.Load().(*VerifiedHeader)

	return struct {
		ID     *big.Int
		Hash   common.Hash
		Height *big.Int
	}{
		Hash:   h.Hash,
		Height: h.Height,
	}
}

// watchBlockProposed watches newly proposed blocks and keeps updating current
// driver state.
func (s *State) watchBlockProposed(ctx context.Context) (ethereum.Subscription, error) {
	newBlockProposedCh := make(chan *bindings.TaikoL1ClientBlockProposed, 10)
	sub, err := s.rpc.TaikoL1.WatchBlockProposed(nil, newBlockProposedCh, nil)
	if err != nil {
		log.Error("Create TaikoL1.BlockProposed subscription error", "error", err)
		return nil, err
	}

	defer sub.Unsubscribe()

	for {
		select {
		case e := <-newBlockProposedCh:
			s.setHeadBlockID(e.Id)
		case err := <-sub.Err():
			return sub, err
		case <-ctx.Done():
			return sub, nil
		}
	}
}

// setHeadBlockID sets the last pending block ID concurrent safely.
func (s *State) setHeadBlockID(id *big.Int) {
	log.Debug("New head block ID", "ID", id)
	metrics.DriverL2HeadIDGauge.Update(id.Int64())
	s.l2HeadBlockID.Store(id)
}

// GetHeadBlockID reads the last pending block ID concurrent safely.
func (s *State) GetHeadBlockID() *big.Int {
	return s.l2HeadBlockID.Load().(*big.Int)
}

// SubL1HeadsFeed registers a subscription of new L1 heads.
func (s *State) SubL1HeadsFeed(ch chan *types.Header) event.Subscription {
	return s.l1HeadsFeed.Subscribe(ch)
}

// VerfiyL2Block checks whether the given block is in L2 node's local chain.
func (s *State) VerfiyL2Block(ctx context.Context, protocolBlockHash common.Hash) error {
	header, err := s.rpc.L2.HeaderByHash(ctx, protocolBlockHash)
	if err != nil {
		return err
	}

	if header.Hash() != protocolBlockHash {
		log.Crit(
			"Verified block hash mismatch",
			"protocolBlockHash", protocolBlockHash,
			"L2 node block number", header.Number,
			"L2 node block hash", header.Hash(),
		)
	}

	return nil
}

// TODO: use atomic.Value
func (s *State) SetL1Current(l1Current *types.Header) {
	s.l1Current = l1Current
}

func (s *State) GetL1Current() *types.Header {
	return s.l1Current
}

func (s *State) ResetL1Current() error {
	// l2Head, err := s.rpc.L2.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// GetConstants returns state's all constants.
func (s *State) GetConstants() struct {
	AnchorTxGasLimit  *big.Int
	MaxTxlistBytes    *big.Int
	MaxBlockNumTxs    *big.Int
	MaxBlocksGasLimit *big.Int
	MinTxGasLimit     *big.Int
} {
	return struct {
		AnchorTxGasLimit  *big.Int
		MaxTxlistBytes    *big.Int
		MaxBlockNumTxs    *big.Int
		MaxBlocksGasLimit *big.Int
		MinTxGasLimit     *big.Int
	}{
		AnchorTxGasLimit:  s.anchorTxGasLimit,
		MaxTxlistBytes:    s.maxTxlistBytes,
		MaxBlockNumTxs:    s.maxBlockNumTxs,
		MaxBlocksGasLimit: s.maxBlocksGasLimit,
		MinTxGasLimit:     s.minTxGasLimit,
	}
}

func (s *State) getSyncedHeaderID(l1Height uint64, hash common.Hash) (*big.Int, error) {
	iter, err := s.rpc.TaikoL1.FilterBlockVerified(&bind.FilterOpts{
		Start: l1Height,
		End:   &l1Height,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to filter BlockVerified event: %w", err)
	}

	for iter.Next() {
		e := iter.Event

		if bytes.Compare(e.BlockHash[:], hash.Bytes()) != 0 {
			continue
		}

		return e.Id, nil
	}

	return nil, fmt.Errorf("verified block %s BlockVerified event not found", hash)
}
