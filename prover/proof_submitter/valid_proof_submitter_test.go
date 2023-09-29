package submitter

import (
	"bytes"
	"context"
	"math/big"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/beaconsync"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/calldata"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/proposer"
	capacity "github.com/taikoxyz/taiko-client/prover/capacity_manager"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	"github.com/taikoxyz/taiko-client/prover/server"
	"github.com/taikoxyz/taiko-client/testutils"
	"github.com/taikoxyz/taiko-client/testutils/fakeprover"
)

type ProofSubmitterTestSuite struct {
	testutils.ClientSuite
	validProofSubmitter *ValidProofSubmitter
	calldataSyncer      *calldata.Syncer
	proposer            *proposer.Proposer
	validProofCh        chan *proofProducer.ProofWithHeader
	invalidProofCh      chan *proofProducer.ProofWithHeader
	rpcClient           *rpc.Client
	proverEndpoints     []*url.URL
	proverServer        *server.ProverServer
}

func (s *ProofSubmitterTestSuite) SetupTest() {
	s.ClientSuite.SetupTest()
	jwtSecret, err := jwt.ParseSecretFromFile(testutils.JwtSecretFile)
	s.NoError(err)
	s.rpcClient, err = rpc.NewClient(context.Background(), &rpc.ClientConfig{
		L1Endpoint:        s.L1.WsEndpoint(),
		L2Endpoint:        s.L2.WsEndpoint(),
		TaikoL1Address:    testutils.TaikoL1Address,
		TaikoTokenAddress: testutils.TaikoL1TokenAddress,
		TaikoL2Address:    testutils.TaikoL2Address,
		L2EngineEndpoint:  s.L2.AuthEndpoint(),
		JwtSecret:         string(jwtSecret),
		RetryInterval:     backoff.DefaultMaxInterval,
	})
	s.NoError(err)

	s.validProofCh = make(chan *proofProducer.ProofWithHeader, 1024)
	s.invalidProofCh = make(chan *proofProducer.ProofWithHeader, 1024)

	s.validProofSubmitter, err = NewValidProofSubmitter(
		s.rpcClient,
		&proofProducer.DummyProofProducer{},
		s.validProofCh,
		testutils.TaikoL2Address,
		testutils.ProverPrivKey,
		&sync.Mutex{},
		false,
		"test",
		1,
		12*time.Second,
		10*time.Second,
		nil,
		2,
		nil,
	)
	s.Nil(err)

	// Init calldata syncer
	testState, err := state.New(context.Background(), s.rpcClient)
	s.Nil(err)

	tracker := beaconsync.NewSyncProgressTracker(s.rpcClient.L2, 30*time.Second)

	s.calldataSyncer, err = calldata.NewSyncer(
		context.Background(),
		s.rpcClient,
		testState,
		tracker,
		testutils.TaikoL1SignalService,
	)
	s.Nil(err)

	// Init proposer
	prop := new(proposer.Proposer)
	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.proverEndpoints = []*url.URL{testutils.LocalRandomProverEndpoint()}
	protocolConfigs, err := s.rpcClient.TaikoL1.GetConfig(nil)
	s.NoError(err)
	s.proverServer, err = fakeprover.New(&protocolConfigs, jwtSecret, s.rpcClient, testutils.ProverPrivKey, capacity.New(1024, 100*time.Second), s.proverEndpoints[0])
	s.NoError(err)
	s.Nil(proposer.InitFromConfig(context.Background(), prop, (&proposer.Config{
		L1Endpoint:                         s.L1.WsEndpoint(),
		L2Endpoint:                         s.L2.WsEndpoint(),
		TaikoL1Address:                     testutils.TaikoL1Address,
		TaikoL2Address:                     testutils.TaikoL2Address,
		TaikoTokenAddress:                  testutils.TaikoL1TokenAddress,
		L1ProposerPrivKey:                  testutils.ProposerPrivKey,
		L2SuggestedFeeRecipient:            testutils.ProposerAddress,
		ProposeInterval:                    &proposeInterval,
		MaxProposedTxListsPerEpoch:         1,
		WaitReceiptTimeout:                 10 * time.Second,
		ProverEndpoints:                    s.proverEndpoints,
		BlockProposalFee:                   big.NewInt(1000),
		BlockProposalFeeIterations:         3,
		BlockProposalFeeIncreasePercentage: common.Big2,
	})))

	s.proposer = prop
}

func (s *ProofSubmitterTestSuite) TearDownTest() {
	s.rpcClient.Close()
	s.ClientSuite.TearDownTest()
}

func (s *ProofSubmitterTestSuite) TestValidProofSubmitterRequestProofDeadlineExceeded() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s.ErrorContains(
		s.validProofSubmitter.RequestProof(
			ctx, &bindings.TaikoL1ClientBlockProposed{BlockId: common.Big256}), "context deadline exceeded",
	)
}

func (s *ProofSubmitterTestSuite) TestValidProofSubmitterSubmitProofMetadataNotFound() {
	s.Error(
		s.validProofSubmitter.SubmitProof(
			context.Background(), &proofProducer.ProofWithHeader{
				BlockID: common.Big256,
				Meta:    &bindings.TaikoDataBlockMetadata{},
				Header:  &types.Header{},
				ZkProof: bytes.Repeat([]byte{0xff}, 100),
			},
		),
	)
}

func (s *ProofSubmitterTestSuite) TestValidSubmitProofs() {
	events := proposer.ProposeAndInsertEmptyBlocks(&s.ClientSuite, s.proposer, s.calldataSyncer)

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
			ctx, &bindings.TaikoL1ClientBlockProposed{BlockId: common.Big256}), "context canceled",
	)
}

func TestProofSubmitterTestSuite(t *testing.T) {
	suite.Run(t, new(ProofSubmitterTestSuite))
}
