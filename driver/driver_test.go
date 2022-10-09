package driver

import (
	"crypto/rand"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/pkg/jwt"
)

func TestMain(m *testing.M) {
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlDebug,
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		),
	)
	os.Exit(m.Run())
}

func newTestDriver(t *testing.T) *Driver {
	jwtSecret, err := jwt.ParseSecretFromFile(os.Getenv("JWT_SECRET"))
	require.Nil(t, err)
	require.NotEmpty(t, jwtSecret)

	d, err := New(&Config{
		L1Endpoint:       os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:       os.Getenv("L2_NODE_ENDPOINT"),
		L2EngineEndpoint: os.Getenv("L2_NODE_ENGINE_ENDPOINT"),
		TaikoL1Address:   common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		JwtSecret:        string(jwtSecret),
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
