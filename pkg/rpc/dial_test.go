package rpc

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/pkg/jwt"
)

func TestDialEngineClient(t *testing.T) {
	jwtSecret, err := jwt.ParseSecretFromFile(os.Getenv("JWT_SECRET"))

	require.Nil(t, err)
	require.NotEmpty(t, jwtSecret)

	client, err := DialEngineClientWithBackoff(
		context.Background(),
		os.Getenv("L2_NODE_ENGINE_ENDPOINT"),
		string(jwtSecret),
	)

	require.Nil(t, err)

	var result beacon.ExecutableDataV1
	err = client.CallContext(context.Background(), &result, "engine_getPayloadV1", beacon.PayloadID{})

	require.Equal(t, beacon.UnknownPayload.Error(), err.Error())
}
