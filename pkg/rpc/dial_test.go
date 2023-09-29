package rpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/testutils"
)

func (s *RpcTestSuite) TestDialEngineClientWithBackoff() {
	jwtSecret, err := jwt.ParseSecretFromFile(testutils.JwtSecretFile)
	s.NoError(err)
	s.NotEmpty(jwtSecret)

	client, err := DialEngineClientWithBackoff(
		context.Background(),
		s.L2.AuthEndpoint(),
		string(jwtSecret),
		12*time.Second,
		new(big.Int).SetUint64(10),
	)

	s.NoError(err)

	var result engine.ExecutableData
	err = client.CallContext(context.Background(), &result, "engine_getPayloadV1", engine.PayloadID{})

	s.Equal(engine.UnknownPayload.Error(), err.Error())
	client.Close()
}

func (s *RpcTestSuite) TestDialClientWithBackoff() {
	client, err := DialClientWithBackoff(
		context.Background(),
		s.L2.WsEndpoint(),
		12*time.Second,
		new(big.Int).SetUint64(10),
	)
	s.NoError(err)

	genesis, err := client.HeaderByNumber(context.Background(), common.Big0)
	s.NoError(err)

	s.Equal(common.Big0.Uint64(), genesis.Number.Uint64())
	client.Close()
}

func (s *RpcTestSuite) TestDialClientWithBackoff_CtxError() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := DialClientWithBackoff(
		ctx,
		"invalid",
		-1,
		new(big.Int).SetUint64(10),
	)
	s.Error(err)
}

func (s *RpcTestSuite) TestDialEngineClientWithBackoff_CtxError() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	jwtSecret, err := jwt.ParseSecretFromFile(testutils.JwtSecretFile)
	s.NoError(err)
	s.NotEmpty(jwtSecret)

	_, err2 := DialEngineClientWithBackoff(
		ctx,
		"invalid",
		string(jwtSecret),
		-1,
		new(big.Int).SetUint64(10),
	)
	s.Error(err2)
}

func (s *RpcTestSuite) TestDialEngineClient_UrlError() {
	_, err := DialEngineClient(context.Background(), "invalid", "invalid")
	s.Error(err)
}
