package rpc

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
)

func TestDialEngineClientWithBackoff(t *testing.T) {
	jwtSecret, err := jwt.ParseSecretFromFile(os.Getenv("JWT_SECRET"))

	require.Nil(t, err)
	require.NotEmpty(t, jwtSecret)

	client, err := DialEngineClientWithBackoff(
		context.Background(),
		os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
		string(jwtSecret),
		12*time.Second,
	)

	require.Nil(t, err)

	var result engine.ExecutableData
	err = client.CallContext(context.Background(), &result, "engine_getPayloadV1", engine.PayloadID{})

	require.Equal(t, engine.UnknownPayload.Error(), err.Error())
}

func TestDialClientWithBackoff(t *testing.T) {
	client, err := DialClientWithBackoff(
		context.Background(),
		os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		12*time.Second,
	)
	require.Nil(t, err)

	genesis, err := client.HeaderByNumber(context.Background(), common.Big0)
	require.Nil(t, err)

	require.Equal(t, common.Big0.Uint64(), genesis.Number.Uint64())
}
