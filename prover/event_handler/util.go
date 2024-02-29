package handler

import (
	"errors"
	"fmt"
	"time"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

var (
	errTierNotFound = errors.New("tier not found")
)

// getProvingWindow returns the provingWindow of the given proposed block.
func getProvingWindow(
	e *bindings.TaikoL1ClientBlockProposed,
	tiers []*rpc.TierProviderTierWithID,
) (time.Duration, error) {
	for _, t := range tiers {
		if e.Meta.MinTier == t.ID {
			return time.Duration(t.ProvingWindow) * time.Minute, nil
		}
	}

	return 0, errTierNotFound
}

// isProvingWindowExpired returns true if the assigned prover proving window of
// the given proposed block is expired.
func isProvingWindowExpired(
	e *bindings.TaikoL1ClientBlockProposed,
	tiers []*rpc.TierProviderTierWithID,
) (bool, time.Duration, error) {
	provingWindow, err := getProvingWindow(e, tiers)
	if err != nil {
		return false, 0, fmt.Errorf("failed to get proving window: %w", err)
	}

	var (
		now       = uint64(time.Now().Unix())
		exipredAt = e.Meta.Timestamp + uint64(provingWindow.Seconds())
	)

	return now > exipredAt, time.Duration(exipredAt-now) * time.Second, nil
}
