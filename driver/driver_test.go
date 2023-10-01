package driver

import (
	"context"
	"math/big"
	"net/url"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/prover/server"
	"github.com/taikoxyz/taiko-client/testutils"
	"github.com/taikoxyz/taiko-client/testutils/helper"
)

type DriverTestSuite struct {
	testutils.ClientTestSuite
	cancel          context.CancelFunc
	p               *proposer.Proposer
	d               *Driver
	rpcClient       *rpc.Client
	proverEndpoints []*url.URL
	proverServer    *server.ProverServer
}

func (s *DriverTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()
	jwtSecret, err := jwt.ParseSecretFromFile(testutils.JwtSecretFile)
	s.NoError(err)
	s.rpcClient = helper.NewWsRpcClient(&s.ClientTestSuite)
	// Init driver

	d := new(Driver)
	ctx, cancel := context.WithCancel(context.Background())
	s.Nil(InitFromConfig(ctx, d, &Config{
		L1Endpoint:       s.L1.WsEndpoint(),
		L2Endpoint:       s.L2.WsEndpoint(),
		L2EngineEndpoint: s.L2.AuthEndpoint(),
		TaikoL1Address:   s.L1.TaikoL1Address,
		TaikoL2Address:   testutils.TaikoL2Address,
		JwtSecret:        string(jwtSecret),
	}))
	s.d = d
	s.cancel = cancel
	s.proverEndpoints, s.proverServer, err = helper.DefaultFakeProver(&s.ClientTestSuite, s.rpcClient)
	s.NoError(err)
	// Init proposer
	p := new(proposer.Proposer)
	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.Nil(proposer.InitFromConfig(context.Background(), p, (&proposer.Config{
		L1Endpoint:                         s.L1.WsEndpoint(),
		L2Endpoint:                         s.L2.WsEndpoint(),
		TaikoL1Address:                     s.L1.TaikoL1Address,
		TaikoL2Address:                     testutils.TaikoL2Address,
		TaikoTokenAddress:                  s.L1.TaikoL1TokenAddress,
		L1ProposerPrivKey:                  testutils.ProposerPrivKey,
		L2SuggestedFeeRecipient:            testutils.L2SuggestedFeeRecipient,
		ProposeInterval:                    &proposeInterval,
		MaxProposedTxListsPerEpoch:         1,
		WaitReceiptTimeout:                 10 * time.Second,
		ProverEndpoints:                    s.proverEndpoints,
		BlockProposalFee:                   big.NewInt(1000),
		BlockProposalFeeIterations:         3,
		BlockProposalFeeIncreasePercentage: common.Big2,
	})))
	s.p = p
}

func (s *DriverTestSuite) TearDownTest() {
	s.d.Close(context.Background())
	s.p.Close(context.Background())
	_ = s.proverServer.Shutdown(context.Background())
	s.rpcClient.Close()
	s.ClientTestSuite.TearDownTest()
}

func (s *DriverTestSuite) TestName() {
	s.Equal("driver", s.d.Name())
}

func (s *DriverTestSuite) TestProcessL1Blocks() {
	l1Head1, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head1, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Nil(s.d.ChainSyncer().CalldataSyncer().ProcessL1Blocks(context.Background(), l1Head1))

	// Propose a valid L2 block
	helper.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())

	l2Head2, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(l2Head2.Number.Uint64(), l2Head1.Number.Uint64())

	// Empty blocks
	helper.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())
	s.Nil(err)

	l2Head3, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(l2Head3.Number.Uint64(), l2Head2.Number.Uint64())

	for _, height := range []uint64{l2Head3.Number.Uint64(), l2Head3.Number.Uint64() - 1} {
		header, err := s.d.rpc.L2.HeaderByNumber(context.Background(), new(big.Int).SetUint64(height))
		s.Nil(err)

		txCount, err := s.d.rpc.L2.TransactionCount(context.Background(), header.Hash())
		s.Nil(err)
		s.Equal(uint(1), txCount)

		anchorTx, err := s.d.rpc.L2.TransactionInBlock(context.Background(), header.Hash(), 0)
		s.Nil(err)

		method, err := encoding.TaikoL2ABI.MethodById(anchorTx.Data())
		s.Nil(err)
		s.Equal("anchor", method.Name)
	}
}

func (s *DriverTestSuite) TestCheckL1ReorgToHigherFork() {
	var testnetL1SnapshotID string
	s.Nil(s.rpcClient.L1RawRPC.CallContext(context.Background(), &testnetL1SnapshotID, "evm_snapshot"))
	s.NotEmpty(testnetL1SnapshotID)

	l1Head1, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l2Head1, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	// Propose two L2 blocks
	helper.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())
	helper.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())

	l1Head2, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l2Head2, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(l2Head2.Number.Uint64(), l2Head1.Number.Uint64())
	s.Greater(l1Head2.Number.Uint64(), l1Head1.Number.Uint64())

	reorged, _, _, err := s.rpcClient.CheckL1ReorgFromL2EE(context.Background(), l2Head2.Number)
	s.Nil(err)
	s.False(reorged)

	// Reorg back to l2Head1
	var revertRes bool
	s.Nil(s.rpcClient.L1RawRPC.CallContext(context.Background(), &revertRes, "evm_revert", testnetL1SnapshotID))
	s.True(revertRes)

	l1Head3, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Equal(l1Head3.Number.Uint64(), l1Head1.Number.Uint64())
	s.Equal(l1Head3.Hash(), l1Head1.Hash())

	// Propose ten blocks on another fork
	for i := 0; i < 10; i++ {
		helper.ProposeInvalidTxListBytes(&s.ClientTestSuite, s.p)
	}

	l1Head4, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(l1Head4.Number.Uint64(), l1Head2.Number.Uint64())

	s.Nil(s.d.ChainSyncer().CalldataSyncer().ProcessL1Blocks(context.Background(), l1Head4))

	l2Head3, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Equal(l2Head1.Number.Uint64()+10, l2Head3.Number.Uint64())

	parent, err := s.d.rpc.L2.HeaderByNumber(context.Background(), new(big.Int).SetUint64(l2Head1.Number.Uint64()+1))
	s.Nil(err)
	s.Equal(parent.ParentHash, l2Head1.Hash())
	s.NotEqual(parent.Hash(), l2Head2.ParentHash)
}

func (s *DriverTestSuite) TestCheckL1ReorgToLowerFork() {
	var testnetL1SnapshotID string
	s.Nil(s.rpcClient.L1RawRPC.CallContext(context.Background(), &testnetL1SnapshotID, "evm_snapshot"))
	s.NotEmpty(testnetL1SnapshotID)

	l1Head1, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l2Head1, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	// Propose two L2 blocks
	helper.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())
	time.Sleep(3 * time.Second)
	helper.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())

	l1Head2, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l2Head2, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(l2Head2.Number.Uint64(), l2Head1.Number.Uint64())
	s.Greater(l1Head2.Number.Uint64(), l1Head1.Number.Uint64())

	reorged, _, _, err := s.rpcClient.CheckL1ReorgFromL2EE(context.Background(), l2Head2.Number)
	s.Nil(err)
	s.False(reorged)

	// Reorg back to l2Head1
	var revertRes bool
	s.Nil(s.rpcClient.L1RawRPC.CallContext(context.Background(), &revertRes, "evm_revert", testnetL1SnapshotID))
	s.True(revertRes)

	l1Head3, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Equal(l1Head3.Number.Uint64(), l1Head1.Number.Uint64())
	s.Equal(l1Head3.Hash(), l1Head1.Hash())

	// Propose one blocks on another fork
	helper.ProposeInvalidTxListBytes(&s.ClientTestSuite, s.p)

	l1Head4, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(l1Head4.Number.Uint64(), l1Head3.Number.Uint64())
	s.Less(l1Head4.Number.Uint64(), l1Head2.Number.Uint64())

	s.Nil(s.d.ChainSyncer().CalldataSyncer().ProcessL1Blocks(context.Background(), l1Head4))

	l2Head3, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	parent, err := s.d.rpc.L2.HeaderByHash(context.Background(), l2Head3.ParentHash)
	s.Nil(err)
	s.Equal(l2Head3.Number.Uint64(), l2Head2.Number.Uint64()-1)
	s.Equal(parent.Hash(), l2Head1.Hash())
}

func (s *DriverTestSuite) TestCheckL1ReorgToSameHeightFork() {
	var testnetL1SnapshotID string
	s.Nil(s.rpcClient.L1RawRPC.CallContext(context.Background(), &testnetL1SnapshotID, "evm_snapshot"))
	s.NotEmpty(testnetL1SnapshotID)

	l1Head1, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l2Head1, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	// Propose two L2 blocks
	helper.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())
	time.Sleep(3 * time.Second)
	helper.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())

	l1Head2, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l2Head2, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(l2Head2.Number.Uint64(), l2Head1.Number.Uint64())
	s.Greater(l1Head2.Number.Uint64(), l1Head1.Number.Uint64())

	reorged, _, _, err := s.rpcClient.CheckL1ReorgFromL2EE(context.Background(), l2Head2.Number)
	s.Nil(err)
	s.False(reorged)

	// Reorg back to l2Head1
	var revertRes bool
	s.Nil(s.rpcClient.L1RawRPC.CallContext(context.Background(), &revertRes, "evm_revert", testnetL1SnapshotID))
	s.True(revertRes)

	l1Head3, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Equal(l1Head3.Number.Uint64(), l1Head1.Number.Uint64())
	s.Equal(l1Head3.Hash(), l1Head1.Hash())

	// Propose two blocks on another fork
	helper.ProposeInvalidTxListBytes(&s.ClientTestSuite, s.p)
	time.Sleep(3 * time.Second)
	helper.ProposeInvalidTxListBytes(&s.ClientTestSuite, s.p)

	l1Head4, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(l1Head4.Number.Uint64(), l1Head3.Number.Uint64())
	s.Equal(l1Head4.Number.Uint64(), l1Head2.Number.Uint64())

	s.Nil(s.d.ChainSyncer().CalldataSyncer().ProcessL1Blocks(context.Background(), l1Head4))

	l2Head3, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	parent, err := s.d.rpc.L2.HeaderByHash(context.Background(), l2Head3.ParentHash)
	s.Nil(err)
	s.Equal(l2Head3.Number.Uint64(), l2Head2.Number.Uint64())
	s.NotEqual(l2Head3.Hash(), l2Head2.Hash())
	s.Equal(parent.ParentHash, l2Head1.Hash())
}

func (s *DriverTestSuite) TestDoSyncNoNewL2Blocks() {
	s.Nil(s.d.doSync())
}

func (s *DriverTestSuite) TestStartClose() {
	s.Nil(s.d.Start())
	s.cancel()
	s.d.Close(context.Background())
}

func (s *DriverTestSuite) TestL1Current() {
	// propose and insert a block
	helper.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())
	// reset L1 current with increased height
	_, id, err := s.d.state.ResetL1Current(s.d.ctx, &state.HeightOrID{ID: common.Big1})
	s.Equal(common.Big1, id)
	s.Nil(err)
}

func TestDriverTestSuite(t *testing.T) {
	suite.Run(t, new(DriverTestSuite))
}
