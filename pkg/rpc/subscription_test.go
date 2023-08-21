package rpc

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/bindings"
)

func TestSubscribeEvent(t *testing.T) {
	require.NotNil(t, SubscribeEvent("test", func(ctx context.Context) (event.Subscription, error) {
		return event.NewSubscription(func(c <-chan struct{}) error { return nil }), nil
	}))
}

func TestSubscribeBlockVerified(t *testing.T) {
	require.NotNil(t, SubscribeBlockVerified(
		newTestClient(t).TaikoL1,
		make(chan *bindings.TaikoL1ClientBlockVerified, 1024)),
	)
}

func TestSubscribeBlockProposed(t *testing.T) {
	require.NotNil(t, SubscribeBlockProposed(
		newTestClient(t).TaikoL1,
		make(chan *bindings.TaikoL1ClientBlockProposed, 1024)),
	)
}

func TestSubscribeSubscribeXchainSynced(t *testing.T) {
	require.NotNil(t, SubscribeXchainSynced(
		newTestClient(t).TaikoL1,
		make(chan *bindings.TaikoL1ClientCrossChainSynced, 1024)),
	)
}

func TestSubscribeBlockProven(t *testing.T) {
	require.NotNil(t, SubscribeBlockProven(
		newTestClient(t).TaikoL1,
		make(chan *bindings.TaikoL1ClientBlockProven, 1024)),
	)
}

func TestSubscribeChainHead(t *testing.T) {
	require.NotNil(t, SubscribeChainHead(
		newTestClient(t).L1,
		make(chan *types.Header, 1024)),
	)
}
