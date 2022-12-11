package driver

import (
	"github.com/ethereum/go-ethereum"
)

func (s *DriverTestSuite) TestProgressed() {
	s.True(Progressed(nil, &ethereum.SyncProgress{}))
	s.False(Progressed(&ethereum.SyncProgress{}, &ethereum.SyncProgress{}))

	// Block
	s.True(Progressed(&ethereum.SyncProgress{CurrentBlock: 0}, &ethereum.SyncProgress{CurrentBlock: 1}))
	s.False(Progressed(&ethereum.SyncProgress{CurrentBlock: 0}, &ethereum.SyncProgress{CurrentBlock: 0}))
	s.False(Progressed(&ethereum.SyncProgress{CurrentBlock: 1}, &ethereum.SyncProgress{CurrentBlock: 1}))

	// Fast sync fields
	s.True(Progressed(&ethereum.SyncProgress{PulledStates: 0}, &ethereum.SyncProgress{PulledStates: 1}))

	// Snap sync fields
	s.True(Progressed(&ethereum.SyncProgress{SyncedAccounts: 0}, &ethereum.SyncProgress{SyncedAccounts: 1}))
	s.True(Progressed(&ethereum.SyncProgress{SyncedAccountBytes: 0}, &ethereum.SyncProgress{SyncedAccountBytes: 1}))
	s.True(Progressed(&ethereum.SyncProgress{SyncedBytecodes: 0}, &ethereum.SyncProgress{SyncedBytecodes: 1}))
	s.True(Progressed(&ethereum.SyncProgress{SyncedBytecodeBytes: 0}, &ethereum.SyncProgress{SyncedBytecodeBytes: 1}))
	s.True(Progressed(&ethereum.SyncProgress{SyncedStorage: 0}, &ethereum.SyncProgress{SyncedStorage: 1}))
	s.True(Progressed(&ethereum.SyncProgress{SyncedStorageBytes: 0}, &ethereum.SyncProgress{SyncedStorageBytes: 1}))
	s.True(Progressed(&ethereum.SyncProgress{HealedTrienodes: 0}, &ethereum.SyncProgress{HealedTrienodes: 1}))
	s.True(Progressed(&ethereum.SyncProgress{HealedTrienodeBytes: 0}, &ethereum.SyncProgress{HealedTrienodeBytes: 1}))
	s.True(Progressed(&ethereum.SyncProgress{HealedBytecodes: 0}, &ethereum.SyncProgress{HealedBytecodes: 1}))
	s.True(Progressed(&ethereum.SyncProgress{HealedBytecodeBytes: 0}, &ethereum.SyncProgress{HealedBytecodeBytes: 1}))
	s.True(Progressed(&ethereum.SyncProgress{HealingTrienodes: 0}, &ethereum.SyncProgress{HealingTrienodes: 1}))
	s.True(Progressed(&ethereum.SyncProgress{HealingBytecode: 0}, &ethereum.SyncProgress{HealingBytecode: 1}))
}
