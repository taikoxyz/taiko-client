package rpc

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/stretchr/testify/require"
)

func TestL2EngineBorbidden(t *testing.T) {
	c := newTestClient(t)

	_, err := c.L2Engine.ForkchoiceUpdate(
		context.Background(),
		&beacon.ForkchoiceStateV1{},
		&beacon.PayloadAttributesV1{},
	)
	require.ErrorContains(t, err, "Forbidden")

	_, err = c.L2Engine.NewPayload(
		context.Background(),
		&beacon.ExecutableDataV1{},
	)
	require.ErrorContains(t, err, "Forbidden")

	_, err = c.L2Engine.GetPayload(
		context.Background(),
		&beacon.PayloadID{},
	)
	require.ErrorContains(t, err, "Forbidden")
}
