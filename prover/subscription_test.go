package prover

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStartSubscription(t *testing.T) {
	p := newTestProver(t)
	require.NotPanics(t, p.startSubscription)
	require.NotPanics(t, p.closeSubscription)
}
