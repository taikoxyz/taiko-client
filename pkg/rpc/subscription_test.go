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
	client := s.newTestClient()
	defer client.Close()
	s.NotNil(SubscribeBlockVerified(
		client.TaikoL1,
		make(chan *bindings.TaikoL1ClientBlockVerified, 1024)),
	)
}

func (s *RpcTestSuite) TestSubscribeBlockProposed() {
	client := s.newTestClient()
	defer client.Close()
	s.NotNil(SubscribeBlockProposed(
		client.TaikoL1,
		make(chan *bindings.TaikoL1ClientBlockProposed, 1024)),
	)
}

func (s *RpcTestSuite) TestSubscribeSubscribeXchainSynced() {
	client := s.newTestClient()
	defer client.Close()
	s.NotNil(SubscribeXchainSynced(
		client.TaikoL1,
		make(chan *bindings.TaikoL1ClientCrossChainSynced, 1024)),
	)
}

func (s *RpcTestSuite) TestSubscribeBlockProven() {
	client := s.newTestClient()
	defer client.Close()
	s.NotNil(SubscribeBlockProven(
		client.TaikoL1,
		make(chan *bindings.TaikoL1ClientBlockProven, 1024)),
	)
}

func (s *RpcTestSuite) TestSubscribeChainHead() {
	client := s.newTestClient()
	defer client.Close()
	s.NotNil(SubscribeChainHead(
		client.L1,
		make(chan *types.Header, 1024)),
	)
}
