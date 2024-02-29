package state

import (
	"sync/atomic"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// State represents the internal state of a prover.
type State struct {
	lastHandledBlockID atomic.Uint64
	l1Current          *types.Header
	reorgDetectedFlag  bool
	tiers              []*rpc.TierProviderTierWithID
}

func (s *State) GetLastHandledBlockID() uint64 {
	return s.lastHandledBlockID.Load()
}

func (s *State) SetLastHandledBlockID(blockID uint64) {
	s.lastHandledBlockID.Store(blockID)
}

func (s *State) GetL1Current() *types.Header {
	return s.l1Current
}

func (s *State) SetL1Current(header *types.Header) {
	s.l1Current = header
}

func (s *State) GetReorgDetectedFlag() bool {
	return s.reorgDetectedFlag
}

func (s *State) SetReorgDetectedFlag(flag bool) {
	s.reorgDetectedFlag = flag
}

func (s *State) GetTiers() []*rpc.TierProviderTierWithID {
	return s.tiers
}

func (s *State) SetTiers(tiers []*rpc.TierProviderTierWithID) {
	s.tiers = tiers
}
