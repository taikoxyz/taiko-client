package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/testutils"
)

type RpcTestSuite struct {
	testutils.ClientSuite
}

func (s *RpcTestSuite) SetupTest() {
	s.ClientSuite.SetupTest()
}

func (s *RpcTestSuite) TearDownTest() {
	s.ClientSuite.TearDownTest()
}

func (s *RpcTestSuite) newTestClient() *Client {
	cli, err := NewClient(context.Background(), &ClientConfig{
		L1Endpoint:        s.L1.WsEndpoint(),
		L2Endpoint:        s.L2.WsEndpoint(),
		TaikoL1Address:    testutils.TaikoL1Address,
		TaikoL2Address:    testutils.TaikoL2Address,
		TaikoTokenAddress: testutils.TaikoL1TokenAddress,
		L2EngineEndpoint:  s.L2.AuthEndpoint(),
		JwtSecret:         testutils.JwtSecretFile,
		RetryInterval:     backoff.DefaultMaxInterval,
	})
	s.NoError(err)
	s.NotNil(cli)
	return cli
}

func (s *RpcTestSuite) newTestClientWithTimeout() *Client {
	timeout := 5 * time.Second
	cli, err := NewClient(context.Background(), &ClientConfig{
		L1Endpoint:        s.L1.WsEndpoint(),
		L2Endpoint:        s.L2.WsEndpoint(),
		TaikoL1Address:    testutils.TaikoL1Address,
		TaikoL2Address:    testutils.TaikoL2Address,
		TaikoTokenAddress: testutils.TaikoL1TokenAddress,
		L2EngineEndpoint:  s.L2.AuthEndpoint(),
		JwtSecret:         testutils.JwtSecretFile,
		RetryInterval:     backoff.DefaultMaxInterval,
		Timeout:           &timeout,
	})

	s.NoError(err)
	s.NotNil(cli)

	return cli
}

func TestRPCTestSuite(t *testing.T) {
	suite.Run(t, new(RpcTestSuite))
}
