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
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type VerifiedHeaderInfo struct {
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
	l2Head         *atomic.Value // Current L2 execution engine's local chain head
	l2HeadBlockID  *atomic.Value // Latest known L2 block ID
	l2VerifiedHead *atomic.Value // Latest known L2 verified head
	l1Current      *types.Header // Current L1 block sync cursor

	// Constants
	genesisL1Height *big.Int

	// RPC clients
	rpc *rpc.Client
}

// NewState creates a new driver state instance.
func NewState(ctx context.Context, rpc *rpc.Client) (*State, error) {
	// Set the L2 head's latest known L1 origin as current L1 sync cursor.
	latestL2KnownL1Header, err := rpc.LatestL2KnownL1Header(ctx)
	if err != nil {
		return nil, err
	}

	stateVars, err := rpc.GetProtocolStateVariables(nil)
	if err != nil {
		return nil, err
	}

	log.Info("Genesis L1 height", "height", stateVars.GenesisHeight)

	s := &State{
		rpc:             rpc,
		genesisL1Height: new(big.Int).SetUint64(stateVars.GenesisHeight),
		l1Head:          new(atomic.Value),
		l2Head:          new(atomic.Value),
		l2HeadBlockID:   new(atomic.Value),
		l2VerifiedHead:  new(atomic.Value),
		l1Current:       latestL2KnownL1Header,
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
	stateVars, err := s.rpc.GetProtocolStateVariables(nil)
	if err != nil {
		return err
	}

	latestVerifiedBlockHash, err := s.rpc.TaikoL1.GetSyncedHeader(
		nil,
		new(big.Int).SetUint64(stateVars.LatestVerifiedHeight),
	)
	if err != nil {
		return err
	}

	s.setLatestVerifiedBlockHash(
		new(big.Int).SetUint64(stateVars.LatestVerifiedID),
		new(big.Int).SetUint64(stateVars.LatestVerifiedHeight),
		latestVerifiedBlockHash,
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
	s.setHeadBlockID(new(big.Int).SetUint64(stateVars.NextBlockID - 1))

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
			// L2 execution has not synced that verified block yet.
			if s.GetL2Head().Number.Cmp(e.Height) < 0 {
				continue
			}
			if err := s.VerifyL2Block(ctx, e.SrcHash); err != nil {
				log.Error("Check new verified L2 block error", "error", err)
				continue
			}
			id, err := s.getSyncedHeaderID(e.Raw.BlockNumber, e.SrcHash)
			if err != nil {
				log.Error("Get synced header block ID error", "error", err)
				continue
			}
			s.setLatestVerifiedBlockHash(id, e.SrcHeight, e.SrcHash)
		case err := <-sub.Err():
			return sub, err
		case <-ctx.Done():
			return sub, nil
		}
	}
}

// setLatestVerifiedBlockHash sets the latest verified L2 block hash concurrent safely.
func (s *State) setLatestVerifiedBlockHash(id *big.Int, height *big.Int, hash common.Hash) {
	log.Debug("New verified block", "height", height, "hash", hash)
	metrics.DriverL2VerifiedHeightGauge.Update(height.Int64())
	s.l2VerifiedHead.Store(&VerifiedHeaderInfo{ID: id, Height: height, Hash: hash})
}

// getLatestVerifiedBlock reads the latest verified L2 block concurrent safely.
func (s *State) getLatestVerifiedBlock() *VerifiedHeaderInfo {
	return s.l2VerifiedHead.Load().(*VerifiedHeaderInfo)
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

// getHeadBlockID reads the last pending block ID concurrent safely.
func (s *State) getHeadBlockID() *big.Int {
	return s.l2HeadBlockID.Load().(*big.Int)
}

// SubL1HeadsFeed registers a subscription of new L1 heads.
func (s *State) SubL1HeadsFeed(ch chan *types.Header) event.Subscription {
	return s.l1HeadsFeed.Subscribe(ch)
}

// VerifyL2Block checks whether the given block is in L2 execution engine's local chain.
func (s *State) VerifyL2Block(ctx context.Context, protocolBlockHash common.Hash) error {
	header, err := s.rpc.L2.HeaderByHash(ctx, protocolBlockHash)
	if err != nil {
		return err
	}

	if header.Hash() != protocolBlockHash {
		log.Crit(
			"Verified block hash mismatch",
			"protocolBlockHash", protocolBlockHash,
			"block number in L2 execution engine", header.Number,
			"block hash in L2 execution engine", header.Hash(),
		)
	}

	return nil
}

// resetL1Current resets the l1Current cursor to the L1 height which emitted a
// BlockProven event with given blockID / blockHash.
func (s *State) resetL1Current(ctx context.Context, heightOrID *HeightOrID) (*big.Int, error) {
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
		s.l1Current, err = s.rpc.L1.HeaderByNumber(ctx, s.genesisL1Height)
		return common.Big0, err
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
				StartHeight: s.genesisL1Height,
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
			StartHeight: s.genesisL1Height,
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
		return nil, fmt.Errorf("BlockProprosed event not found, blockID: %s", heightOrID.ID)
	}

	if s.l1Current, err = s.rpc.L1.HeaderByNumber(ctx, l1CurrentHeight); err != nil {
		return nil, err
	}

	log.Info("Reset L1 current cursor", "height", s.l1Current.Number)

	return heightOrID.ID, nil
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

		if !bytes.Equal(e.BlockHash[:], hash.Bytes()) {
			continue
		}

		return e.Id, nil
	}

	return nil, fmt.Errorf("verified block %s BlockVerified event not found", hash)
}
