package prover

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
)

// startSubscription initializes all subscriptions in current prover instance.
func (p *Prover) startSubscription() {
	p.blockProposedSub = event.ResubscribeErr(
		backoff.DefaultMaxInterval,
		func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				log.Warn("Failed to subscribe TaikoL1.BlockProposed, try resubscribing", "error", err)
			}

			return p.watchBlockProposed(ctx)
		},
	)

	p.blockFinalizedSub = event.ResubscribeErr(
		backoff.DefaultMaxInterval,
		func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				log.Warn("Failed to subscribe TaikoL1.BlockFinalized, try resubscribing", "error", err)
			}

			return p.watchBlockFinalized(ctx)
		},
	)
}

// closeSubscription closes all subscriptions.
func (p *Prover) closeSubscription() {
	p.blockFinalizedSub.Unsubscribe()
	p.blockProposedSub.Unsubscribe()
}

// watchBlockFinalized watches newly finalized blocks from TaikoL1 contract.
func (p *Prover) watchBlockFinalized(ctx context.Context) (event.Subscription, error) {
	sub, err := p.rpc.TaikoL1.WatchBlockFinalized(nil, p.blockFinalizedCh, nil)
	if err != nil {
		log.Error("Create TaikoL1.BlockFinalized subscription error", "error", err)
		return nil, err
	}

	defer sub.Unsubscribe()

	select {
	case err := <-sub.Err():
		return sub, err
	case <-ctx.Done():
		return sub, nil
	}
}

// watchBlockProposed watches newly proposed blocks from TaikoL1 contract.
func (p *Prover) watchBlockProposed(ctx context.Context) (event.Subscription, error) {
	sub, err := p.rpc.TaikoL1.WatchBlockProposed(nil, p.blockProposedCh, nil)
	if err != nil {
		log.Error("Create TaikoL1.BlockProposed subscription error", "error", err)
		return nil, err
	}

	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			return sub, err
		case <-ctx.Done():
			return sub, nil
		}
	}
}
