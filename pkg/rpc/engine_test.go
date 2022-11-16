package rpc

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestForkchoiceUpdate(t *testing.T) {
	c := newTestClient(t)

	_, err := c.L2Engine.ForkchoiceUpdate(
		context.Background(),
		&beacon.ForkchoiceStateV1{},
		&beacon.PayloadAttributesV1{},
	)

	log.Error("ForkchoiceUpdate", "err", err)

	require.ErrorContains(t, err, "Forbidden")
}
