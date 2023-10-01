package helper

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/testutils"
)

func NewWsRpcClientConfig(s *testutils.ClientTestSuite) *rpc.ClientConfig {
	timeout := 5 * time.Second
	jwtSecret, err := jwt.ParseSecretFromFile(testutils.JwtSecretFile)
	s.NoError(err)
	return &rpc.ClientConfig{
		L1Endpoint:        s.L1.WsEndpoint(),
		L2Endpoint:        s.L2.WsEndpoint(),
		TaikoL1Address:    s.L1.TaikoL1Address,
		TaikoTokenAddress: s.L1.TaikoL1TokenAddress,
		TaikoL2Address:    testutils.TaikoL2Address,
		L2EngineEndpoint:  s.L2.AuthEndpoint(),
		JwtSecret:         string(jwtSecret),
		RetryInterval:     backoff.DefaultMaxInterval,
		Timeout:           &timeout,
	}
}

func NewWsRpcClient(s *testutils.ClientTestSuite) *rpc.Client {
	cli, err := rpc.NewClient(context.Background(), NewWsRpcClientConfig(s))
	s.NoError(err)
	return cli
}
