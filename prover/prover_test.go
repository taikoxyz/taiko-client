package prover

import (
	"context"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/driver"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/proposer"
	producer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	"github.com/taikoxyz/taiko-client/testutils"
	"github.com/taikoxyz/taiko-client/testutils/fakeprover"
)

type ProverTestSuite struct {
	testutils.ClientSuite
	p        *Prover
	cancel   context.CancelFunc
	d        *driver.Driver
	proposer *proposer.Proposer
}

func (s *ProverTestSuite) SetupTest() {
	s.ClientSuite.SetupTest()

	// Init prover
	l1ProverPrivKey := testutils.ProverPrivKey

	proverServerUrl := testutils.LocalRandomProverEndpoint()
	port, err := strconv.Atoi(proverServerUrl.Port())
	s.Nil(err)

	ctx, cancel := context.WithCancel(context.Background())
	p := new(Prover)
	s.Nil(InitFromConfig(ctx, p, (&Config{
		L1WsEndpoint:                    s.L1.WsEndpoint(),
		L1HttpEndpoint:                  s.L1.HttpEndpoint(),
		L2WsEndpoint:                    s.L2.WsEndpoint(),
		L2HttpEndpoint:                  s.L2.HttpEndpoint(),
		TaikoL1Address:                  testutils.TaikoL1Address,
		TaikoL2Address:                  testutils.TaikoL2Address,
		L1ProverPrivKey:                 l1ProverPrivKey,
		OracleProverPrivateKey:          l1ProverPrivKey,
		OracleProver:                    false,
		Dummy:                           true,
		MaxConcurrentProvingJobs:        1,
		CheckProofWindowExpiredInterval: 5 * time.Second,
		ProveUnassignedBlocks:           true,
		Capacity:                        1024,
		MinProofFee:                     common.Big1,
		HTTPServerPort:                  uint64(port),
	})))
	jwtSecret, err := jwt.ParseSecretFromFile(testutils.JwtSecretFile)
	s.NoError(err)
	s.NotEmpty(jwtSecret)
	rpcClient, err := rpc.NewClient(context.Background(), &rpc.ClientConfig{
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
	protocolConfigs, err := rpcClient.TaikoL1.GetConfig(nil)
	s.NoError(err)
	p.srv, err = fakeprover.New(&protocolConfigs, jwtSecret, rpcClient, l1ProverPrivKey, p.capacityManager, proverServerUrl)
	s.NoError(err)
	s.p = p
	s.cancel = cancel

	// Init driver

	d := new(driver.Driver)
	s.Nil(driver.InitFromConfig(context.Background(), d, &driver.Config{
		L1Endpoint:       s.L1.WsEndpoint(),
		L2Endpoint:       s.L2.WsEndpoint(),
		L2EngineEndpoint: s.L2.AuthEndpoint(),
		TaikoL1Address:   testutils.TaikoL1Address,
		TaikoL2Address:   testutils.TaikoL2Address,
		JwtSecret:        string(jwtSecret),
	}))
	s.d = d

	// Init proposer
	l1ProposerPrivKey := testutils.ProposerPrivKey
	s.Nil(err)

	prop := new(proposer.Proposer)

	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.Nil(proposer.InitFromConfig(context.Background(), prop, (&proposer.Config{
		L1Endpoint:                         s.L1.WsEndpoint(),
		L2Endpoint:                         s.L2.WsEndpoint(),
		TaikoL1Address:                     testutils.TaikoL1Address,
		TaikoL2Address:                     testutils.TaikoL2Address,
		TaikoTokenAddress:                  testutils.TaikoL1TokenAddress,
		L1ProposerPrivKey:                  l1ProposerPrivKey,
		L2SuggestedFeeRecipient:            testutils.ProposerAddress,
		ProposeInterval:                    &proposeInterval,
		MaxProposedTxListsPerEpoch:         1,
		WaitReceiptTimeout:                 10 * time.Second,
		ProverEndpoints:                    []*url.URL{proverServerUrl},
		BlockProposalFee:                   common.Big256,
		BlockProposalFeeIterations:         3,
		BlockProposalFeeIncreasePercentage: common.Big2,
	})))

	s.proposer = prop
}

func (s *ProverTestSuite) TestName() {
	s.Equal("prover", s.p.Name())
}

func (s *ProverTestSuite) TestInitError() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	l1ProverPrivKey := testutils.ProverPrivKey

	p := new(Prover)
	// Error should be "context canceled", instead is "Dial ethclient error:"
	s.ErrorContains(InitFromConfig(ctx, p, (&Config{
		L1WsEndpoint:                      s.L1.WsEndpoint(),
		L1HttpEndpoint:                    s.L1.HttpEndpoint(),
		L2WsEndpoint:                      s.L2.WsEndpoint(),
		L2HttpEndpoint:                    s.L2.HttpEndpoint(),
		TaikoL1Address:                    testutils.TaikoL1Address,
		TaikoL2Address:                    testutils.TaikoL2Address,
		L1ProverPrivKey:                   l1ProverPrivKey,
		OracleProverPrivateKey:            l1ProverPrivKey,
		Dummy:                             true,
		MaxConcurrentProvingJobs:          1,
		CheckProofWindowExpiredInterval:   5 * time.Second,
		ProveUnassignedBlocks:             true,
		ProveBlockTxReplacementMultiplier: 2,
	})), "dial tcp:")
}

func (s *ProverTestSuite) TestOnBlockProposed() {
	s.p.cfg.OracleProver = true
	// Init prover
	l1ProverPrivKey := testutils.ProverPrivKey
	s.p.cfg.OracleProverPrivateKey = l1ProverPrivKey
	// Valid block
	e := proposer.ProposeAndInsertValidBlock(&s.ClientSuite, s.proposer, s.d.ChainSyncer().CalldataSyncer())
	s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))
	s.Nil(s.p.validProofSubmitter.SubmitProof(context.Background(), <-s.p.proofGenerationCh))

	// Empty blocks
	for _, e = range proposer.ProposeAndInsertEmptyBlocks(
		&s.ClientSuite,
		s.proposer,
		s.d.ChainSyncer().CalldataSyncer(),
	) {
		s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))

		s.Nil(s.p.validProofSubmitter.SubmitProof(context.Background(), <-s.p.proofGenerationCh))
	}
}

func (s *ProverTestSuite) TestOnBlockVerifiedEmptyBlockHash() {
	s.Nil(s.p.onBlockVerified(context.Background(), &bindings.TaikoL1ClientBlockVerified{
		BlockId:   common.Big1,
		BlockHash: common.Hash{},
	},
	))
}

func (s *ProverTestSuite) TestSubmitProofOp() {
	s.NotPanics(func() {
		s.p.submitProofOp(context.Background(), &producer.ProofWithHeader{
			BlockID: common.Big1,
			Meta:    &bindings.TaikoDataBlockMetadata{},
			Header:  &types.Header{},
			ZkProof: []byte{},
		})
	})
	s.NotPanics(func() {
		s.p.submitProofOp(context.Background(), &producer.ProofWithHeader{
			BlockID: common.Big1,
			Meta:    &bindings.TaikoDataBlockMetadata{},
			Header:  &types.Header{},
			ZkProof: []byte{},
		})
	})
}

func (s *ProverTestSuite) TestStartSubscription() {
	s.NotPanics(s.p.initSubscription)
	s.NotPanics(s.p.closeSubscription)
}

func (s *ProverTestSuite) TestCheckChainVerification() {
	s.Nil(s.p.checkChainVerification(0))
	s.p.latestVerifiedL1Height = 1024
	s.Nil(s.p.checkChainVerification(1024))
}

func TestProverTestSuite(t *testing.T) {
	suite.Run(t, new(ProverTestSuite))
}
