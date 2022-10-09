package proposer

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
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

func newTestProposer(t *testing.T) *Proposer {
	proposer, err := New(context.Background(), &Config{
		L1Endpoint:              os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:              os.Getenv("L2_NODE_ENDPOINT"),
		TaikoL1Address:          os.Getenv("TAIKO_L1_ADDRESS"),
		TaikoL2Address:          os.Getenv("TAIKO_L2_ADDRESS"),
		L1ProposerPrivKey:       os.Getenv("L1_PROPOSER_PRIVATE_KEY"),
		L2SuggestedFeeRecipient: os.Getenv("L2_SUGGESTED_FEE_RECIPIENT"),
		ProposeInterval:         "1024h", // No need to periodically propose transactions list in unit tests
	})

	require.Nil(t, err)

	return proposer
}
