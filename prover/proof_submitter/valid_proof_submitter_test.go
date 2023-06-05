package submitter

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/beaconsync"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/calldata"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/proposer"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	"github.com/taikoxyz/taiko-client/testutils"
)

type ProofSubmitterTestSuite struct {
	testutils.ClientTestSuite
	validProofSubmitter *ValidProofSubmitter
	calldataSyncer      *calldata.Syncer
	proposer            *proposer.Proposer
	validProofCh        chan *proofProducer.ProofWithHeader
	invalidProofCh      chan *proofProducer.ProofWithHeader
}

func (s *ProofSubmitterTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	s.validProofCh = make(chan *proofProducer.ProofWithHeader, 1024)
	s.invalidProofCh = make(chan *proofProducer.ProofWithHeader, 1024)

	s.validProofSubmitter, err = NewValidProofSubmitter(
		s.RpcClient,
		&proofProducer.DummyProofProducer{},
		s.validProofCh,
		common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		l1ProverPrivKey,
		&sync.Mutex{},
		false,
		false,
		"test",
		0,
	)
	s.Nil(err)

	// Init calldata syncer
	testState, err := state.New(context.Background(), s.RpcClient)
	s.Nil(err)

	tracker := beaconsync.NewSyncProgressTracker(s.RpcClient.L2, 30*time.Second)

	s.calldataSyncer, err = calldata.NewSyncer(
		context.Background(),
		s.RpcClient,
		testState,
		tracker,
		common.HexToAddress(os.Getenv("L1_SIGNAL_SERVICE_CONTRACT_ADDRESS")),
	)
	s.Nil(err)

	// Init proposer
	prop := new(proposer.Proposer)
	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)
	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.Nil(proposer.InitFromConfig(context.Background(), prop, (&proposer.Config{
		L1Endpoint:                 os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:                 os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:             common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:             common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:          l1ProposerPrivKey,
		L2SuggestedFeeRecipient:    common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:            &proposeInterval, // No need to periodically propose transactions list in unit tests
		MaxProposedTxListsPerEpoch: 1,
	})))

	s.proposer = prop
}

func (s *ProofSubmitterTestSuite) TestValidProofSubmitterRequestProofDeadlineExceeded() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s.ErrorContains(
		s.validProofSubmitter.RequestProof(
			ctx, &bindings.TaikoL1ClientBlockProposed{Id: common.Big256}), "context deadline exceeded",
	)
}

func (s *ProofSubmitterTestSuite) TestValidProofSubmitterSubmitProofMetadataNotFound() {
	s.Error(
		s.validProofSubmitter.SubmitProof(
			context.Background(), &proofProducer.ProofWithHeader{
				BlockID: common.Big256,
				Meta:    &bindings.TaikoDataBlockMetadata{},
				Header:  &types.Header{},
				ZkProof: []byte{0xff},
			},
		),
	)
}

func (s *ProofSubmitterTestSuite) TestValidSubmitProofs() {
	events := testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.proposer, s.calldataSyncer)

	for _, e := range events {
		s.Nil(s.validProofSubmitter.RequestProof(context.Background(), e))
		proofWithHeader := <-s.validProofCh
		s.Nil(s.validProofSubmitter.SubmitProof(context.Background(), proofWithHeader))
	}
}

func (s *ProofSubmitterTestSuite) TestValidProofSubmitterRequestProofCancelled() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.AfterFunc(2*time.Second, func() {
			cancel()
		})
	}()

	s.ErrorContains(
		s.validProofSubmitter.RequestProof(
			ctx, &bindings.TaikoL1ClientBlockProposed{Id: common.Big256}), "context canceled",
	)
}

func TestProofSubmitterTestSuite(t *testing.T) {
	suite.Run(t, new(ProofSubmitterTestSuite))
}
