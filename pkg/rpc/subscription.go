package rpc

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

// SubscribeEvent creates a event subscription, will retry if the established subscription failed.
func SubscribeEvent(
	eventName string,
	handler func(ctx context.Context) (event.Subscription, error),
) event.Subscription {
	return event.ResubscribeErr(
		backoff.DefaultMaxInterval,
		func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				log.Warn("Failed to subscribe protocol event, try resubscribing", "event", eventName, "error", err)
			}

			return handler(ctx)
		},
	)
}

// SubscribeBlockVerified subscribes the protocol's BlockVerified events.
func SubscribeBlockVerified(
	taikoL1 *bindings.TaikoL1Client,
	ch chan *bindings.TaikoL1ClientBlockVerified,
) event.Subscription {
	return SubscribeEvent("BlockVerified", func(ctx context.Context) (event.Subscription, error) {
		sub, err := taikoL1.WatchBlockVerified(nil, ch, nil)
		if err != nil {
			log.Error("Create TaikoL1.BlockVerified subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// SubscribeBlockProposed subscribes the protocol's BlockProposed events.
func SubscribeBlockProposed(
	taikoL1 *bindings.TaikoL1Client,
	ch chan *bindings.TaikoL1ClientBlockProposed,
) event.Subscription {
	return SubscribeEvent("BlockProposed", func(ctx context.Context) (event.Subscription, error) {
		sub, err := taikoL1.WatchBlockProposed(nil, ch, nil)
		if err != nil {
			log.Error("Create TaikoL1.BlockProposed subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// SubscribeXchainSynced subscribes the protocol's XchainSynced events.
func SubscribeXchainSynced(
	taikoL1 *bindings.TaikoL1Client,
	ch chan *bindings.TaikoL1ClientCrossChainSynced,
) event.Subscription {
	return SubscribeEvent("CrossChainSynced", func(ctx context.Context) (event.Subscription, error) {
		sub, err := taikoL1.WatchCrossChainSynced(nil, ch, nil)
		if err != nil {
			log.Error("Create TaikoL1.XchainSynced subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// SubscribeBlockProven subscribes the protocol's BlockProven events.
func SubscribeBlockProven(
	taikoL1 *bindings.TaikoL1Client,
	ch chan *bindings.TaikoL1ClientBlockProven,
) event.Subscription {
	return SubscribeEvent("BlockProven", func(ctx context.Context) (event.Subscription, error) {
		sub, err := taikoL1.WatchBlockProven(nil, ch, nil)
		if err != nil {
			log.Error("Create TaikoL1.BlockProven subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// SubscribeChainHead subscribes the new chain heads.
func SubscribeChainHead(
	client *EthClient,
	ch chan *types.Header,
) event.Subscription {
	return SubscribeEvent("ChainHead", func(ctx context.Context) (event.Subscription, error) {
		sub, err := client.SubscribeNewHead(ctx, ch)
		if err != nil {
			log.Error("Create chain head subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// waitSubErr keeps waiting until the given subscription failed.
func waitSubErr(ctx context.Context, sub event.Subscription) (event.Subscription, error) {
	for {
		select {
		case err := <-sub.Err():
			return sub, err
		case <-ctx.Done():
			return sub, nil
		}
	}
}
