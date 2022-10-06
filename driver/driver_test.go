package driver

import (
	"context"
	"crypto/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taikochain/client-mono/util"
	"github.com/taikochain/taiko-client/common"
)

func newTestDriver(t *testing.T) *Driver {
	jwtSecret, err := util.ParseJWTSecretFromFile(os.Getenv("JWT_SECRET"))
	require.Nil(t, err)
	require.NotEmpty(t, jwtSecret)

	d, err := New(context.Background(), &Config{
		L1Endpoint:     os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:     os.Getenv("L2_NODE_ENDPOINT"),
		L2AuthEndpoint: os.Getenv("L2_NODE_ENGINE_ENDPOINT"),
		TaikoL1Address: common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		JwtSecret:      string(jwtSecret),
	})
	require.Nil(t, err)

	return d
}

// randomHash generates a random blob of data and returns it as a hash.
func randomHash() common.Hash {
	var hash common.Hash
	if n, err := rand.Read(hash[:]); n != common.HashLength || err != nil {
		panic(err)
	}
	return hash
}
