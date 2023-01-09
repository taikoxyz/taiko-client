package progressTracker

import (
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/stretchr/testify/require"
)

func TestSyncProgressed(t *testing.T) {
	require.False(t, syncProgressed(nil, &ethereum.SyncProgress{}), nil)
	require.False(t, syncProgressed(&ethereum.SyncProgress{}, &ethereum.SyncProgress{}))

	// Block
	require.True(t, syncProgressed(&ethereum.SyncProgress{CurrentBlock: 0}, &ethereum.SyncProgress{CurrentBlock: 1}))
	require.False(t, syncProgressed(&ethereum.SyncProgress{CurrentBlock: 0}, &ethereum.SyncProgress{CurrentBlock: 0}))
	require.False(t, syncProgressed(&ethereum.SyncProgress{CurrentBlock: 1}, &ethereum.SyncProgress{CurrentBlock: 1}))

	// Fast sync fields
	require.True(t, syncProgressed(&ethereum.SyncProgress{PulledStates: 0}, &ethereum.SyncProgress{PulledStates: 1}))

	// Snap sync fields
	require.True(t, syncProgressed(&ethereum.SyncProgress{SyncedAccounts: 0}, &ethereum.SyncProgress{SyncedAccounts: 1}))
	require.True(t, syncProgressed(
		&ethereum.SyncProgress{SyncedAccountBytes: 0}, &ethereum.SyncProgress{SyncedAccountBytes: 1}),
	)
	require.True(
		t, syncProgressed(&ethereum.SyncProgress{SyncedBytecodes: 0}, &ethereum.SyncProgress{SyncedBytecodes: 1}),
	)
	require.True(
		t, syncProgressed(&ethereum.SyncProgress{SyncedBytecodeBytes: 0}, &ethereum.SyncProgress{SyncedBytecodeBytes: 1}),
	)
	require.True(t, syncProgressed(&ethereum.SyncProgress{SyncedStorage: 0}, &ethereum.SyncProgress{SyncedStorage: 1}))
	require.True(
		t, syncProgressed(&ethereum.SyncProgress{SyncedStorageBytes: 0}, &ethereum.SyncProgress{SyncedStorageBytes: 1}),
	)
	require.True(
		t, syncProgressed(&ethereum.SyncProgress{HealedTrienodes: 0}, &ethereum.SyncProgress{HealedTrienodes: 1}),
	)
	require.True(
		t, syncProgressed(&ethereum.SyncProgress{HealedTrienodeBytes: 0}, &ethereum.SyncProgress{HealedTrienodeBytes: 1}),
	)
	require.True(
		t, syncProgressed(&ethereum.SyncProgress{HealedBytecodes: 0}, &ethereum.SyncProgress{HealedBytecodes: 1}),
	)
	require.True(
		t, syncProgressed(&ethereum.SyncProgress{HealedBytecodeBytes: 0}, &ethereum.SyncProgress{HealedBytecodeBytes: 1}),
	)
	require.True(
		t, syncProgressed(&ethereum.SyncProgress{HealingTrienodes: 0}, &ethereum.SyncProgress{HealingTrienodes: 1}),
	)
	require.True(
		t, syncProgressed(&ethereum.SyncProgress{HealingBytecode: 0}, &ethereum.SyncProgress{HealingBytecode: 1}),
	)
}
