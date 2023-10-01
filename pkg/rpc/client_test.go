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
	testutils.ClientTestSuite
	cli *Client
}

func (s *RpcTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()
	s.cli = s.newTestClient()
}

func (s *RpcTestSuite) TearDownTest() {
	s.ClientTestSuite.TearDownTest()
	s.cli.Close()
}

func (s *RpcTestSuite) newTestClient() *Client {
	cli, err := NewClient(context.Background(), &ClientConfig{
		L1Endpoint:        s.L1.WsEndpoint(),
		L2Endpoint:        s.L2.WsEndpoint(),
		TaikoL1Address:    s.L1.TaikoL1Address,
		TaikoL2Address:    testutils.TaikoL2Address,
		TaikoTokenAddress: s.L1.TaikoL1TokenAddress,
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
		TaikoL1Address:    s.L1.TaikoL1Address,
		TaikoL2Address:    testutils.TaikoL2Address,
		TaikoTokenAddress: s.L1.TaikoL1TokenAddress,
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
