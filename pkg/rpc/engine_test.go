package rpc

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/stretchr/testify/require"
)

func TestL2EngineBorbidden(t *testing.T) {
	c := newTestClient(t)

	_, err := c.L2Engine.ForkchoiceUpdate(
		context.Background(),
		&engine.ForkchoiceStateV1{},
		&engine.PayloadAttributes{},
	)
	require.ErrorContains(t, err, "Unauthorized")

	_, err = c.L2Engine.NewPayload(
		context.Background(),
		&engine.ExecutableData{},
	)
	require.ErrorContains(t, err, "Unauthorized")

	_, err = c.L2Engine.GetPayload(
		context.Background(),
		&engine.PayloadID{},
	)
	require.ErrorContains(t, err, "Unauthorized")
}
