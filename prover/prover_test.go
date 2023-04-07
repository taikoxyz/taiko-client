package prover

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/driver"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/proposer"
	producer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	"github.com/taikoxyz/taiko-client/testutils"
)

type ProverTestSuite struct {
	testutils.ClientTestSuite
	p        *Prover
	cancel   context.CancelFunc
	d        *driver.Driver
	proposer *proposer.Proposer
}

func (s *ProverTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	// Init prover
	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	ctx, cancel := context.WithCancel(context.Background())
	p := new(Prover)
	s.Nil(InitFromConfig(ctx, p, (&Config{
		L1WsEndpoint:             os.Getenv("L1_NODE_WS_ENDPOINT"),
		L1HttpEndpoint:           os.Getenv("L1_NODE_HTTP_ENDPOINT"),
		L2WsEndpoint:             os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		L2HttpEndpoint:           os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
		TaikoL1Address:           common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:           common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProverPrivKey:          l1ProverPrivKey,
		Dummy:                    true,
		MaxConcurrentProvingJobs: 1,
	})))
	s.p = p
	s.cancel = cancel

	// Init driver
	jwtSecret, err := jwt.ParseSecretFromFile(os.Getenv("JWT_SECRET"))
	s.Nil(err)
	s.NotEmpty(jwtSecret)

	throwawayBlocksBuilderPrivKey, err := crypto.ToECDSA(
		common.Hex2Bytes(os.Getenv("THROWAWAY_BLOCKS_BUILDER_PRIV_KEY")),
	)
	s.Nil(err)

	d := new(driver.Driver)
	s.Nil(driver.InitFromConfig(context.Background(), d, &driver.Config{
		L1Endpoint:                    os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:                    os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		L2EngineEndpoint:              os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
		TaikoL1Address:                common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:                common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		ThrowawayBlocksBuilderPrivKey: throwawayBlocksBuilderPrivKey,
		JwtSecret:                     string(jwtSecret),
	}))
	s.d = d

	// Init proposer
	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	prop := new(proposer.Proposer)

	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.Nil(proposer.InitFromConfig(context.Background(), prop, (&proposer.Config{
		L1Endpoint:              os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:              os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:          common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:          common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:       l1ProposerPrivKey,
		L2SuggestedFeeRecipient: common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:         &proposeInterval, // No need to periodically propose transactions list in unit tests
	})))

	s.proposer = prop
	s.proposer.AfterCommitHook = s.MineL1Confirmations
}

func (s *ProverTestSuite) TestName() {
	s.Equal("prover", s.p.Name())
}

func (s *ProverTestSuite) TestOnBlockProposed() {
	// Valid block
	e := testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.proposer, s.d.ChainSyncer().CalldataSyncer())
	s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))
	s.Nil(s.p.validProofSubmitter.SubmitProof(context.Background(), <-s.p.proveValidProofCh, false))

	// Empty blocks
	for _, e = range testutils.ProposeAndInsertEmptyBlocks(
		&s.ClientTestSuite,
		s.proposer,
		s.d.ChainSyncer().CalldataSyncer(),
	) {
		s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))
		s.Nil(s.p.validProofSubmitter.SubmitProof(context.Background(), <-s.p.proveValidProofCh, false))
	}

	// Invalid block
	e = testutils.ProposeAndInsertThrowawayBlock(&s.ClientTestSuite, s.proposer, s.d.ChainSyncer().CalldataSyncer())
	s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))
	s.Nil(s.p.invalidProofSubmitter.SubmitProof(context.Background(), <-s.p.proveInvalidProofCh, false))
}

func (s *ProverTestSuite) TestOnBlockVerifiedEmptyBlockHash() {
	s.Nil(s.p.onBlockVerified(context.Background(), &bindings.TaikoL1ClientBlockVerified{
		Id:        common.Big1,
		BlockHash: common.Hash{}},
	))
}

func (s *ProverTestSuite) TestSubmitProofOp() {
	s.NotPanics(func() {
		s.p.submitProofOp(context.Background(), &producer.ProofWithHeader{
			BlockID: common.Big1,
			Meta:    &bindings.TaikoDataBlockMetadata{},
			Header:  &types.Header{},
			ZkProof: []byte{},
		}, true)
	})
	s.NotPanics(func() {
		s.p.submitProofOp(context.Background(), &producer.ProofWithHeader{
			BlockID: common.Big1,
			Meta:    &bindings.TaikoDataBlockMetadata{},
			Header:  &types.Header{},
			ZkProof: []byte{},
		}, false)
	})
}

func (s *ProverTestSuite) TestStartSubscription() {
	s.NotPanics(s.p.initSubscription)
	s.NotPanics(s.p.closeSubscription)
}

func (s *ProverTestSuite) TestStartClose() {
	s.Nil(s.p.Start())
	s.cancel()
	s.NotPanics(s.p.Close)
}

func TestProverTestSuite(t *testing.T) {
	suite.Run(t, new(ProverTestSuite))
}
