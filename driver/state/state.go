package state

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// HeightOrID contains a block height or a block ID.
type HeightOrID struct {
	Height *big.Int
	ID     *big.Int
}

// NotEmpty checks whether this is an empty struct.
func (h *HeightOrID) NotEmpty() bool {
	return h.Height != nil || h.ID != nil
}

// State contains all states which will be used by driver.
type State struct {
	// Subscriptions, will automatically resubscribe on errors
	l1HeadSub          event.Subscription // L1 new heads
	l2HeadSub          event.Subscription // L2 new heads
	l2BlockProvenSub   event.Subscription // TaikoL1.BlockProven events
	l2BlockVerifiedSub event.Subscription // TaikoL1.BlockVerified events
	l2BlockProposedSub event.Subscription // TaikoL1.BlockProposed events
	l2HeaderSyncedSub  event.Subscription // TaikoL1.HeaderSynced events

	l1HeadCh         chan *types.Header
	l2HeadCh         chan *types.Header
	blockProposedCh  chan *bindings.TaikoL1ClientBlockProposed
	blockProvenCh    chan *bindings.TaikoL1ClientBlockProven
	blockVerifiedCh  chan *bindings.TaikoL1ClientBlockVerified
	crossChainSynced chan *bindings.TaikoL1ClientCrossChainSynced

	// Feeds
	l1HeadsFeed event.Feed // L1 new heads notification feed

	l1Head         *atomic.Value // Latest known L1 head
	l2Head         *atomic.Value // Current L2 execution engine's local chain head
	l2HeadBlockID  *atomic.Value // Latest known L2 block ID
	l2VerifiedHead *atomic.Value // Latest known L2 verified head
	l1Current      *atomic.Value // Current L1 block sync cursor

	// Constants
	GenesisL1Height  *big.Int
	BlockDeadendHash common.Hash

	// RPC clients
	rpc *rpc.Client
}

// New creates a new driver state instance.
func New(ctx context.Context, rpc *rpc.Client) (*State, error) {
	s := &State{
		rpc:              rpc,
		l1Head:           new(atomic.Value),
		l2Head:           new(atomic.Value),
		l2HeadBlockID:    new(atomic.Value),
		l2VerifiedHead:   new(atomic.Value),
		l1Current:        new(atomic.Value),
		l1HeadCh:         make(chan *types.Header, 10),
		l2HeadCh:         make(chan *types.Header, 10),
		blockProposedCh:  make(chan *bindings.TaikoL1ClientBlockProposed, 10),
		blockProvenCh:    make(chan *bindings.TaikoL1ClientBlockProven, 10),
		blockVerifiedCh:  make(chan *bindings.TaikoL1ClientBlockVerified, 10),
		crossChainSynced: make(chan *bindings.TaikoL1ClientCrossChainSynced, 10),
		BlockDeadendHash: common.BigToHash(common.Big1),
	}

	if err := s.init(ctx); err != nil {
		return nil, err
	}

	s.startSubscriptions(ctx)

	return s, nil
}

// Close closes all inner subscriptions.
func (s *State) Close() {
	s.l1HeadSub.Unsubscribe()
	s.l2HeadSub.Unsubscribe()
	s.l2BlockVerifiedSub.Unsubscribe()
	s.l2BlockProposedSub.Unsubscribe()
	s.l2BlockProvenSub.Unsubscribe()
	s.l2HeaderSyncedSub.Unsubscribe()
}

// init fetches the latest status and initializes the state instance.
func (s *State) init(ctx context.Context) error {
	stateVars, err := s.rpc.GetProtocolStateVariables(nil)
	if err != nil {
		return err
	}

	log.Info("Genesis L1 height", "height", stateVars.GenesisHeight)
	s.GenesisL1Height = new(big.Int).SetUint64(stateVars.GenesisHeight)

	// Set the L2 head's latest known L1 origin as current L1 sync cursor.
	latestL2KnownL1Header, err := s.rpc.LatestL2KnownL1Header(ctx)
	if err != nil {
		return err
	}
	s.l1Current.Store(latestL2KnownL1Header)

	// L1 head
	l1Head, err := s.rpc.L1.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	s.setL1Head(l1Head)

	// L2 head
	l2Head, err := s.rpc.L2.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	log.Info("L2 execution engine head", "height", l2Head.Number, "hash", l2Head.Hash())
	s.setL2Head(l2Head)

	latestVerifiedBlockHash, err := s.rpc.TaikoL1.GetCrossChainBlockHash(
		nil,
		new(big.Int).SetUint64(stateVars.LastVerifiedBlockId),
	)
	if err != nil {
		return err
	}

	s.setLatestVerifiedBlockHash(
		new(big.Int).SetUint64(stateVars.LastVerifiedBlockId),
		new(big.Int).SetUint64(stateVars.LastVerifiedBlockId),
		latestVerifiedBlockHash,
	)
	s.setHeadBlockID(new(big.Int).SetUint64(stateVars.NumBlocks - 1))

	return nil
}

// startSubscriptions initializes all subscriptions in the given state instance.
func (s *State) startSubscriptions(ctx context.Context) {
	s.l1HeadSub = rpc.SubscribeChainHead(s.rpc.L1, s.l1HeadCh)
	s.l2HeadSub = rpc.SubscribeChainHead(s.rpc.L2, s.l2HeadCh)
	s.l2HeaderSyncedSub = rpc.SubscribeXchainSynced(s.rpc.TaikoL1, s.crossChainSynced)
	s.l2BlockVerifiedSub = rpc.SubscribeBlockVerified(s.rpc.TaikoL1, s.blockVerifiedCh)
	s.l2BlockProposedSub = rpc.SubscribeBlockProposed(s.rpc.TaikoL1, s.blockProposedCh)
	s.l2BlockProvenSub = rpc.SubscribeBlockProven(s.rpc.TaikoL1, s.blockProvenCh)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case e := <-s.blockProposedCh:
				s.setHeadBlockID(e.Id)
			case e := <-s.blockProvenCh:
				if e.Prover != encoding.SystemProverAddress && e.Prover != encoding.OracleProverAddress {
					log.Info("✅ Block proven", "blockID", e.Id, "hash", common.Hash(e.BlockHash), "prover", e.Prover)
				}
			case e := <-s.blockVerifiedCh:
				log.Info("📈 Block verified", "blockID", e.Id, "hash", common.Hash(e.BlockHash), "reward", e.Reward)
			case e := <-s.crossChainSynced:
				// Verify the protocol synced block, check if it exists in
				// L2 execution engine.
				if s.GetL2Head().Number.Cmp(e.SrcHeight) >= 0 {
					if err := s.VerifyL2Block(ctx, e.SrcHeight, e.BlockHash); err != nil {
						log.Error("Check new verified L2 block error", "error", err)
						continue
					}
				}
				id, err := s.getSyncedHeaderID(e.Raw.BlockNumber, e.BlockHash)
				if err != nil {
					log.Error("Get synced header block ID error", "error", err)
					continue
				}
				s.setLatestVerifiedBlockHash(id, e.SrcHeight, e.BlockHash)
			case newHead := <-s.l1HeadCh:
				s.setL1Head(newHead)
				s.l1HeadsFeed.Send(newHead)
			case newHead := <-s.l2HeadCh:
				s.setL2Head(newHead)
			}
		}
	}()
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

// VerifiedHeaderInfo contains information about a verified L2 block header.
type VerifiedHeaderInfo struct {
	ID   *big.Int
	Hash common.Hash
}

// setLatestVerifiedBlockHash sets the latest verified L2 block hash concurrent safely.
func (s *State) setLatestVerifiedBlockHash(id *big.Int, height *big.Int, hash common.Hash) {
	log.Debug("New verified block", "height", height, "hash", hash)
	metrics.DriverL2VerifiedHeightGauge.Update(height.Int64())
	s.l2VerifiedHead.Store(&VerifiedHeaderInfo{ID: id, Hash: hash})
}

// GetLatestVerifiedBlock reads the latest verified L2 block concurrent safely.
func (s *State) GetLatestVerifiedBlock() *VerifiedHeaderInfo {
	return s.l2VerifiedHead.Load().(*VerifiedHeaderInfo)
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

// VerifyL2Block checks whether the given block is in L2 execution engine's local chain.
func (s *State) VerifyL2Block(ctx context.Context, height *big.Int, hash common.Hash) error {
	header, err := s.rpc.L2.HeaderByNumber(ctx, height)
	if err != nil {
		return err
	}

	if header.Hash() != hash {
		// TODO(david): do not exit but re-sync from genesis?
		log.Crit(
			"Verified block hash mismatch",
			"protocolBlockHash", hash,
			"block number in L2 execution engine", header.Number,
			"block hash in L2 execution engine", header.Hash(),
		)
	}
	return nil
}

// getSyncedHeaderID fetches the block ID of the synced L2 header.
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
