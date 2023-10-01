package rpc

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *RpcTestSuite) TestSubscribeEvent() {
	s.NotNil(SubscribeEvent("test", func(ctx context.Context) (event.Subscription, error) {
		return event.NewSubscription(func(c <-chan struct{}) error { return nil }), nil
	}))
}

func (s *RpcTestSuite) TestSubscribeBlockVerified() {
	s.NotNil(SubscribeBlockVerified(
		s.cli.TaikoL1,
		make(chan *bindings.TaikoL1ClientBlockVerified, 1024)),
	)
}

func (s *RpcTestSuite) TestSubscribeBlockProposed() {
	s.NotNil(SubscribeBlockProposed(
		s.cli.TaikoL1,
		make(chan *bindings.TaikoL1ClientBlockProposed, 1024)),
	)
}

func (s *RpcTestSuite) TestSubscribeSubscribeXchainSynced() {
	s.NotNil(SubscribeXchainSynced(
		s.cli.TaikoL1,
		make(chan *bindings.TaikoL1ClientCrossChainSynced, 1024)),
	)
}

func (s *RpcTestSuite) TestSubscribeBlockProven() {
	s.NotNil(SubscribeBlockProven(
		s.cli.TaikoL1,
		make(chan *bindings.TaikoL1ClientBlockProven, 1024)),
	)
}

func (s *RpcTestSuite) TestSubscribeChainHead() {
	s.NotNil(SubscribeChainHead(
		s.cli.L1,
		make(chan *types.Header, 1024)),
	)
}
