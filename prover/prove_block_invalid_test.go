package prover

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/prover/producer"
)

func TestProveBlockInvalidL1OriginTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	p := newTestProver(t)

	err := p.proveBlockInvalid(ctx, &bindings.TaikoL1ClientBlockProposed{Id: common.Big256}, 0, 0)

	require.ErrorContains(t, err, "context deadline exceeded")
}

func TestSubmitInvalidBlockProofThrowawayBlockNotFound(t *testing.T) {
	p := newTestProver(t)

	require.ErrorContains(t,
		p.submitInvalidBlockProof(
			context.Background(), &producer.ProofWithHeader{
				BlockID: common.Big256,
				Header:  &types.Header{},
				ZkProof: []byte{0xff},
			},
		), "failed to fetch throwaway block",
	)
}
