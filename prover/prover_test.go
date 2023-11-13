package prover

import (
	"context"
	"math/big"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
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

	proverServerUrl := testutils.LocalRandomProverEndpoint()
	port, err := strconv.Atoi(proverServerUrl.Port())
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
		GuardianProverAddress:    common.HexToAddress(os.Getenv("GUARDIAN_PROVER_CONTRACT_ADDRESS")),
		L1ProverPrivKey:          l1ProverPrivKey,
		GuardianProverPrivateKey: l1ProverPrivKey,
		Dummy:                    true,
		MaxConcurrentProvingJobs: 1,
		ProveUnassignedBlocks:    true,
		Capacity:                 1024,
		MinOptimisticTierFee:     common.Big1,
		MinSgxTierFee:            common.Big1,
		MinPseZkevmTierFee:       common.Big1,
		MinSgxAndPseZkevmTierFee: common.Big1,
		HTTPServerPort:           uint64(port),
		WaitReceiptTimeout:       12 * time.Second,
	})))
	p.srv = testutils.NewTestProverServer(
		&s.ClientTestSuite,
		l1ProverPrivKey,
		p.capacityManager,
		proverServerUrl,
	)
	s.p = p
	s.cancel = cancel

	// Init driver
	jwtSecret, err := jwt.ParseSecretFromFile(os.Getenv("JWT_SECRET"))
	s.Nil(err)
	s.NotEmpty(jwtSecret)

	d := new(driver.Driver)
	s.Nil(driver.InitFromConfig(context.Background(), d, &driver.Config{
		L1Endpoint:       os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:       os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		L2EngineEndpoint: os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
		TaikoL1Address:   common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:   common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		JwtSecret:        string(jwtSecret),
	}))
	s.d = d

	// Init proposer
	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	prop := new(proposer.Proposer)

	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.Nil(proposer.InitFromConfig(context.Background(), prop, (&proposer.Config{
		L1Endpoint:                 os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:                 os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:             common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:             common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		TaikoTokenAddress:          common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		L1ProposerPrivKey:          l1ProposerPrivKey,
		ProposeInterval:            &proposeInterval,
		MaxProposedTxListsPerEpoch: 1,
		WaitReceiptTimeout:         12 * time.Second,
		ProverEndpoints:            []*url.URL{proverServerUrl},
		OptimisticTierFee:          common.Big256,
		SgxTierFee:                 common.Big256,
		PseZkevmTierFee:            common.Big256,
		SgxAndPseZkevmTierFee:      common.Big256,
		MaxTierFeePriceBumps:       3,
		TierFeePriceBump:           common.Big2,
	})))

	s.proposer = prop
}

func (s *ProverTestSuite) TestName() {
	s.Equal("prover", s.p.Name())
}

func (s *ProverTestSuite) TestInitError() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	p := new(Prover)
	// Error should be "context canceled", instead is "Dial ethclient error:"
	s.ErrorContains(InitFromConfig(ctx, p, (&Config{
		L1WsEndpoint:                      os.Getenv("L1_NODE_WS_ENDPOINT"),
		L1HttpEndpoint:                    os.Getenv("L1_NODE_HTTP_ENDPOINT"),
		L2WsEndpoint:                      os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		L2HttpEndpoint:                    os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
		TaikoL1Address:                    common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:                    common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProverPrivKey:                   l1ProverPrivKey,
		GuardianProverPrivateKey:          l1ProverPrivKey,
		Dummy:                             true,
		MaxConcurrentProvingJobs:          1,
		ProveUnassignedBlocks:             true,
		ProveBlockTxReplacementMultiplier: 2,
	})), "dial tcp:")
}

func (s *ProverTestSuite) TestOnBlockProposed() {
	// Init prover
	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)
	s.p.cfg.GuardianProverPrivateKey = l1ProverPrivKey
	// Valid block
	e := testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.proposer, s.d.ChainSyncer().CalldataSyncer())
	s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))
	s.Nil(s.p.selectSubmitter(e.Meta.MinTier).SubmitProof(context.Background(), <-s.p.proofGenerationCh))

	// Empty blocks
	for _, e = range testutils.ProposeAndInsertEmptyBlocks(
		&s.ClientTestSuite,
		s.proposer,
		s.d.ChainSyncer().CalldataSyncer(),
	) {
		s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))

		s.Nil(s.p.selectSubmitter(e.Meta.MinTier).SubmitProof(context.Background(), <-s.p.proofGenerationCh))
	}
}

func (s *ProverTestSuite) TestOnBlockVerifiedEmptyBlockHash() {
	s.Nil(s.p.onBlockVerified(context.Background(), &bindings.TaikoL1ClientBlockVerified{
		BlockId:   common.Big1,
		BlockHash: common.Hash{},
	}))
}

func (s *ProverTestSuite) TestSubmitProofOp() {
	s.NotPanics(func() {
		s.p.submitProofOp(context.Background(), &producer.ProofWithHeader{
			BlockID: common.Big1,
			Meta:    &bindings.TaikoDataBlockMetadata{},
			Header:  &types.Header{},
			Proof:   []byte{},
			Tier:    encoding.TierOptimisticID,
			Opts:    &producer.ProofRequestOptions{},
		})
	})
	s.NotPanics(func() {
		s.p.submitProofOp(context.Background(), &producer.ProofWithHeader{
			BlockID: common.Big1,
			Meta:    &bindings.TaikoDataBlockMetadata{},
			Header:  &types.Header{},
			Proof:   []byte{},
			Tier:    encoding.TierOptimisticID,
			Opts:    &producer.ProofRequestOptions{},
		})
	})
}

func (s *ProverTestSuite) TestOnBlockVerified() {
	id := testutils.RandomHash().Big().Uint64()
	s.Nil(s.p.onBlockVerified(context.Background(), &bindings.TaikoL1ClientBlockVerified{
		BlockId: testutils.RandomHash().Big(),
		Raw: types.Log{
			BlockHash:   testutils.RandomHash(),
			BlockNumber: id,
		},
	}))
	s.Equal(id, s.p.latestVerifiedL1Height)
}

func (s *ProverTestSuite) TestContestWrongBlocks() {
	s.p.cfg.ContesterMode = false
	e := testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.proposer, s.d.ChainSyncer().CalldataSyncer())
	s.Nil(s.p.onTransitionProved(context.Background(), &bindings.TaikoL1ClientTransitionProved{
		BlockId: e.BlockId,
		Tier:    e.Meta.MinTier,
	}))
	s.p.cfg.ContesterMode = true

	// Submit a wrong proof at first.
	sink := make(chan *bindings.TaikoL1ClientTransitionProved)
	header, err := s.p.rpc.L2.HeaderByNumber(context.Background(), e.BlockId)
	s.Nil(err)

	sub, err := s.p.rpc.TaikoL1.WatchTransitionProved(nil, sink, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	s.Nil(s.p.proveOp())
	proofWithHeader := <-s.p.proofGenerationCh
	proofWithHeader.Opts.BlockHash = testutils.RandomHash()
	s.Nil(s.p.selectSubmitter(e.Meta.MinTier).SubmitProof(context.Background(), proofWithHeader))

	event := <-sink
	s.Equal(header.Number.Uint64(), event.BlockId.Uint64())
	s.Equal(common.BytesToHash(proofWithHeader.Opts.BlockHash[:]), common.BytesToHash(event.Tran.BlockHash[:]))
	s.NotEqual(header.Hash(), common.BytesToHash(event.Tran.BlockHash[:]))
	s.Equal(header.ParentHash, common.BytesToHash(event.Tran.ParentHash[:]))

	// Contest the transition.
	contestedSink := make(chan *bindings.TaikoL1ClientTransitionContested)
	contestedSub, err := s.p.rpc.TaikoL1.WatchTransitionContested(nil, contestedSink, nil)
	s.Nil(err)
	defer func() {
		contestedSub.Unsubscribe()
		close(contestedSink)
	}()

	s.Greater(header.Number.Uint64(), uint64(0))
	s.Nil(s.p.onTransitionProved(context.Background(), event))

	contestedEvent := <-contestedSink
	s.Equal(header.Number.Uint64(), contestedEvent.BlockId.Uint64())
	s.Equal(header.Hash(), common.BytesToHash(contestedEvent.Tran.BlockHash[:]))
	s.Equal(header.ParentHash, common.BytesToHash(contestedEvent.Tran.ParentHash[:]))

	s.Nil(s.p.onTransitionContested(context.Background(), contestedEvent))

	if contestedEvent.Tier >= encoding.TierSgxAndPseZkevmID {
		approvedSink := make(chan *bindings.GuardianProverApproved)
		approvedSub, err := s.p.rpc.GuardianProver.WatchApproved(nil, approvedSink, []uint64{})
		s.Nil(err)
		defer func() {
			approvedSub.Unsubscribe()
			close(approvedSink)
		}()

		s.Nil(s.p.selectSubmitter(encoding.TierGuardianID).SubmitProof(context.Background(), <-s.p.proofGenerationCh))
		approvedEvent := <-approvedSink

		s.Equal(header.Number.Uint64(), approvedEvent.BlockId)
		return
	}

	s.Nil(s.p.selectSubmitter(contestedEvent.Tier+1).SubmitProof(context.Background(), <-s.p.proofGenerationCh))
	event = <-sink
	s.Equal(header.Number.Uint64(), event.BlockId.Uint64())
}

func (s *ProverTestSuite) TestProveExpiredUnassignedBlock() {
	e := testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.proposer, s.d.ChainSyncer().CalldataSyncer())
	sink := make(chan *bindings.TaikoL1ClientTransitionProved)

	header, err := s.p.rpc.L2.HeaderByNumber(context.Background(), e.BlockId)
	s.Nil(err)

	sub, err := s.p.rpc.TaikoL1.WatchTransitionProved(nil, sink, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	e.AssignedProver = common.BytesToAddress(testutils.RandomHash().Bytes())
	s.Nil(s.p.onProvingWindowExpired(context.Background(), e))
	s.Nil(s.p.selectSubmitter(e.Meta.MinTier).SubmitProof(context.Background(), <-s.p.proofGenerationCh))

	event := <-sink
	s.Equal(header.Number.Uint64(), event.BlockId.Uint64())
	s.Equal(header.Hash(), common.BytesToHash(event.Tran.BlockHash[:]))
	s.Equal(header.ParentHash, common.BytesToHash(event.Tran.ParentHash[:]))
}

func (s *ProverTestSuite) TestSelectSubmitter() {
	submitter := s.p.selectSubmitter(encoding.TierGuardianID - 1)
	s.NotNil(submitter)
	s.Equal(encoding.TierGuardianID, submitter.Tier())
}

func (s *ProverTestSuite) TestSelectSubmitterNotFound() {
	submitter := s.p.selectSubmitter(encoding.TierGuardianID + 1)
	s.Nil(submitter)
}

func (s *ProverTestSuite) TestGetSubmitterByTier() {
	submitter := s.p.getSubmitterByTier(encoding.TierGuardianID)
	s.NotNil(submitter)
	s.Equal(encoding.TierGuardianID, submitter.Tier())
	s.Nil(s.p.getSubmitterByTier(encoding.TierGuardianID + 1))
}

func (s *ProverTestSuite) TestGetProvingWindowNotFound() {
	_, err := s.p.getProvingWindow(&bindings.TaikoL1ClientBlockProposed{
		Meta: bindings.TaikoDataBlockMetadata{
			MinTier: encoding.TierGuardianID + 1,
		},
	})
	s.ErrorIs(err, errTierNotFound)
}

func (s *ProverTestSuite) TestIsBlockVerified() {
	vars, err := s.p.rpc.TaikoL1.GetStateVariables(nil)
	s.Nil(err)

	verified, err := s.p.isBlockVerified(new(big.Int).SetUint64(vars.B.LastVerifiedBlockId))
	s.Nil(err)
	s.True(verified)

	verified, err = s.p.isBlockVerified(new(big.Int).SetUint64(vars.B.LastVerifiedBlockId + 1))
	s.Nil(err)
	s.False(verified)
}

func (s *ProverTestSuite) TestProveOp() {
	e := testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.proposer, s.d.ChainSyncer().CalldataSyncer())
	sink := make(chan *bindings.TaikoL1ClientTransitionProved)

	header, err := s.p.rpc.L2.HeaderByNumber(context.Background(), e.BlockId)
	s.Nil(err)

	sub, err := s.p.rpc.TaikoL1.WatchTransitionProved(nil, sink, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	s.Nil(s.p.proveOp())
	s.Nil(s.p.selectSubmitter(e.Meta.MinTier).SubmitProof(context.Background(), <-s.p.proofGenerationCh))

	event := <-sink
	s.Equal(header.Number.Uint64(), event.BlockId.Uint64())
	s.Equal(header.Hash(), common.BytesToHash(event.Tran.BlockHash[:]))
	s.Equal(header.ParentHash, common.BytesToHash(event.Tran.ParentHash[:]))
}

func (s *ProverTestSuite) TestReleaseOneCapacity() {
	s.NotPanics(func() { s.p.releaseOneCapacity(common.Big1) })
}

func (s *ProverTestSuite) TestStartSubscription() {
	s.NotPanics(s.p.initSubscription)
	s.NotPanics(s.p.closeSubscription)
}

func TestProverTestSuite(t *testing.T) {
	suite.Run(t, new(ProverTestSuite))
}
