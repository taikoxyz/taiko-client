package prover

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/bindings"
	"github.com/taikochain/taiko-client/prover/producer"
)

func TestProveBlockValidL1OriginTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	p := newTestProver(t)

	err := p.proveBlockValid(ctx, &bindings.TaikoL1ClientBlockProposed{Id: common.Big256})

	require.ErrorContains(t, err, "context deadline exceeded")
}

func TestSubmitValidBlockProofMetadataNotFound(t *testing.T) {
	p := newTestProver(t)

	require.ErrorContains(t,
		p.submitValidBlockProof(
			context.Background(), &producer.ProofWithHeader{
				BlockID: common.Big256,
				Header:  &types.Header{},
				ZkProof: []byte{0xff},
			},
		), "failed to fetch L2 block with given block ID",
	)
}
