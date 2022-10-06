package client

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taikochain/client-mono/util"
	"github.com/taikochain/taiko-client/core/beacon"
)

func TestNewRPCClientWithAuth(t *testing.T) {
	jwtSecret, err := util.ParseJWTSecretFromFile(os.Getenv("JWT_SECRET"))

	require.Nil(t, err)
	require.NotEmpty(t, jwtSecret)

	client, err := DialEngineClientWithBackoff(
		context.Background(),
		os.Getenv("L2_NODE_AUTH_ENDPOINT"),
		string(jwtSecret),
	)

	require.Nil(t, err)

	var result beacon.ExecutableDataV1
	err = client.CallContext(context.Background(), &result, "engine_getPayloadV1", beacon.PayloadID{})

	require.Equal(t, beacon.UnknownPayload.Error(), err.Error())
}
