package guardianproversender

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/beaconsync"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/calldata"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/testutils"
	"golang.org/x/sync/errgroup"
)

type GuardianProverSenderTestSuite struct {
	testutils.ClientTestSuite
	proposer          *proposer.Proposer
	calldataSyncer    *calldata.Syncer
	healthCheckServer *httptest.Server
	sender            *GuardianProverBlockSender
}

func (s *GuardianProverSenderTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	// Init sender
	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	s.healthCheckServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	healthCheckServerUrl, err := url.Parse(s.healthCheckServer.URL)
	s.Nil(err)

	s.sender = New(
		l1ProverPrivKey,
		healthCheckServerUrl,
		memorydb.New(),
		s.RpcClient,
		crypto.PubkeyToAddress(l1ProverPrivKey.PublicKey),
	)

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
		L1Endpoint:            os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:            os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:        common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:        common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		TaikoTokenAddress:     common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		AssignmentHookAddress: common.HexToAddress(os.Getenv("ASSIGNMENT_HOOK_ADDRESS")),

		L1ProposerPrivKey:          l1ProposerPrivKey,
		ProposeInterval:            &proposeInterval,
		MaxProposedTxListsPerEpoch: 1,
		WaitReceiptTimeout:         12 * time.Second,
		ProverEndpoints:            s.ProverEndpoints,
		OptimisticTierFee:          common.Big256,
		SgxTierFee:                 common.Big256,
		PseZkevmTierFee:            common.Big256,
		SgxAndPseZkevmTierFee:      common.Big256,
		MaxTierFeePriceBumps:       3,
		TierFeePriceBump:           common.Big2,
	})))

	s.proposer = prop
}

func (s *GuardianProverSenderTestSuite) TearDownTest() {
	s.Nil(s.sender.Close())
	s.healthCheckServer.Close()
}

func (s *GuardianProverSenderTestSuite) TestPost() {
	s.Nil(s.sender.post(context.Background(), "healthCheck", &healthCheckReq{}))
}

func (s *GuardianProverSenderTestSuite) TestSign() {
	events := testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.proposer, s.calldataSyncer)
	s.NotEmpty(events)

	h, err := s.RpcClient.L2.HeaderByNumber(context.Background(), common.Big1)
	s.Nil(err)

	sig, header, err := s.sender.sign(context.Background(), common.Big1)
	s.Nil(err)
	s.NotEmpty(sig)
	s.Equal(h.Hash(), header.Hash())

	head, err := s.RpcClient.L2.BlockNumber(context.Background())
	s.Nil(err)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		newHead := head + 1
		sig, header, err = s.sender.sign(ctx, new(big.Int).SetUint64(newHead))
		s.Nil(err)

		h, err := s.RpcClient.L2.HeaderByNumber(context.Background(), new(big.Int).SetUint64(newHead))
		s.Nil(err)
		s.NotEmpty(sig)
		s.Equal(h.Hash(), header.Hash())

		return err
	})
	g.Go(func() error {
		events := testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.proposer, s.calldataSyncer)
		s.NotEmpty(events)
		return nil
	})

	s.Nil(g.Wait())
}

func (s *GuardianProverSenderTestSuite) TestSendHeartbeat() {
	s.Nil(s.sender.SendHeartbeat(context.Background()))
}

func (s *GuardianProverSenderTestSuite) TestSignAndSendBlock() {
	events := testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.proposer, s.calldataSyncer)
	s.NotEmpty(events)

	s.Nil(s.sender.SignAndSendBlock(context.Background(), common.Big1))
}

func TestGuardianProverSenderTestSuite(t *testing.T) {
	suite.Run(t, new(GuardianProverSenderTestSuite))
}
