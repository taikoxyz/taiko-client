package prover

import (
	"context"
	"math/rand"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/bindings"
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

func newTestProver(t *testing.T) *Prover {
	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	require.Nil(t, err)

	p := new(Prover)

	require.Nil(t, initFromConfig(context.Background(), p, &Config{
		L1Endpoint:      os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:      os.Getenv("L2_NODE_ENDPOINT"),
		TaikoL1Address:  common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:  common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProverPrivKey: l1ProverPrivKey,
		Dummy:           true,
	}))

	return p
}
func TestName(t *testing.T) {
	require.Equal(t, "prover", newTestProver(t).Name())
}

func TestGetProveBlocksTxOpts(t *testing.T) {
	opts, err := newTestProver(t).getProveBlocksTxOpts(context.Background())
	require.Nil(t, err)
	require.Equal(t, proveBlocksGasLimit, opts.GasLimit)
}

func TestBatchHandleBlockProposedEventsBuffered(t *testing.T) {
	require.Nil(
		t, newTestProver(t).batchHandleBlockProposedEvents(context.Background(), &bindings.TaikoL1ClientBlockProposed{}),
	)
}

func TestOnForceTimerEventNotFound(t *testing.T) {
	require.ErrorContains(t, newTestProver(t).onForceTimer(context.Background()), "BlockProposed events not found")
}

func TestOnBlockFinalizedEmptyBlockHash(t *testing.T) {
	require.Nil(
		t,
		newTestProver(t).
			onBlockFinalized(context.Background(), &bindings.TaikoL1ClientBlockFinalized{BlockHash: common.Hash{}}),
	)
}

// randomHash generates a random blob of data and returns it as a hash.
func randomHash() common.Hash {
	var hash common.Hash
	if n, err := rand.Read(hash[:]); n != common.HashLength || err != nil {
		panic(err)
	}
	return hash
}
