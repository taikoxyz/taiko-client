package prover

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/bindings"
)

func TestProveBlockInvalidL1OriginTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	p := newTestProver(t)

	err := p.proveBlockInvalid(ctx, &bindings.TaikoL1ClientBlockProposed{Id: common.Big256}, 0, 0)

	require.Error(t, err, "context deadline exceeded")
}
