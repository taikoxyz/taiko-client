package prover

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	"github.com/taikochain/taiko-client/event"
	"github.com/taikochain/taiko-client/log"
)

// startSubscription initializes all subscriptions in current prover instance.
func (p *Prover) startSubscription() {
	p.blockProposedSub = event.ResubscribeErr(backoff.DefaultMaxInterval, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			log.Warn("Failed to subscribe TaikoL1.BlockProposed, try resubscribing", "error", err)
		}

		return p.watchBlockProposed(ctx)
	})

	p.blockFinalizedSub = event.ResubscribeErr(backoff.DefaultMaxInterval, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			log.Warn("Failed to subscribe TaikoL1.BlockFinalized, try resubscribing", "error", err)
		}

		return p.watchBlockFinalized(ctx)
	})

	// TODO: whether these are necessary?
	go func() {
		err, ok := <-p.blockProposedSub.Err()
		if !ok {
			return
		}
		log.Error("Subscribe TaikoL1.BlockProposed error", "error", err)
	}()
	go func() {
		err, ok := <-p.blockFinalizedSub.Err()
		if !ok {
			return
		}
		log.Error("Subscribe TaikoL1.BlockFinalized error", "error", err)
	}()
}

// watchBlockFinalized watches newly finalized blocks from TaikoL1 contract.
func (p *Prover) watchBlockFinalized(ctx context.Context) (event.Subscription, error) {
	sub, err := p.taikoL1.WatchBlockFinalized(nil, p.blockFinalizedCh, nil)
	if err != nil {
		log.Error("Create TaikoL1.BlockFinalized subscription error", "error", err)
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// watchBlockProposed watches newly proposed blocks from TaikoL1 contract.
func (p *Prover) watchBlockProposed(ctx context.Context) (event.Subscription, error) {
	sub, err := p.taikoL1.WatchBlockProposed(nil, p.blockProposedCh, nil)
	if err != nil {
		log.Error("Create TaikoL1.BlockProposed subscription error", "error", err)
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
