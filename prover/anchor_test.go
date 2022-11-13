package prover

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestValidateAnchorTx(t *testing.T) {
	p := newTestProver(t)

	// invalid To
	tx := types.NewTransaction(
		0, common.BytesToAddress(randBytes(1024)), common.Big0, common.Big0.Uint64(), common.Big0, []byte{},
	)

	require.NotNil(t, p.validateAnchorTx(context.Background(), tx))
}
