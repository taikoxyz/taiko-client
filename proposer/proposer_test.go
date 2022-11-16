package proposer

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	require.Nil(t, err)

	p := new(Proposer)

	require.Nil(t, initFromConfig(context.Background(), p, (&Config{
		L1Endpoint:              os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:              os.Getenv("L2_NODE_ENDPOINT"),
		TaikoL1Address:          common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:          common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:       l1ProposerPrivKey,
		L2SuggestedFeeRecipient: common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:         1024 * time.Hour, // No need to periodically propose transactions list in unit tests
	})))

	return p
}

func TestSumTxsGasLimit(t *testing.T) {
	txs := []*types.Transaction{
		types.NewTransaction(0, common.Address{}, common.Big0, 1, common.Big0, []byte{}), // gasLimit: 1
		types.NewTransaction(0, common.Address{}, common.Big0, 2, common.Big0, []byte{}), // gasLimit: 2
		types.NewTransaction(0, common.Address{}, common.Big0, 3, common.Big0, []byte{}), // gasLimit: 3
	}

	require.Equal(t, uint64(1+2+3), sumTxsGasLimit(txs))
}

func TestName(t *testing.T) {
	require.Equal(t, "proposer", newTestProposer(t).Name())
}

func TestProposeOp(t *testing.T) {
	require.ErrorContains(t, newTestProposer(t).proposeOp(context.Background()), "l2 node is syncing")
}
