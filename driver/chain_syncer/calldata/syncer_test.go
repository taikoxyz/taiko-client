package calldata

import (
	"context"
	"math/big"
	"math/rand"
	"net/url"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/beaconsync"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/prover/server"
	"github.com/taikoxyz/taiko-client/testutils"
	"github.com/taikoxyz/taiko-client/testutils/helper"
)

type CalldataSyncerTestSuite struct {
	testutils.ClientTestSuite
	s               *Syncer
	p               testutils.Proposer
	rpcClient       *rpc.Client
	proverEndpoints []*url.URL
	proverServer    *server.ProverServer
}

func (s *CalldataSyncerTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()
	s.rpcClient = helper.NewWsRpcClient(&s.ClientTestSuite)
	state, err := state.New(context.Background(), s.rpcClient)
	s.Nil(err)

	syncer, err := NewSyncer(
		context.Background(),
		s.rpcClient,
		state,
		beaconsync.NewSyncProgressTracker(s.rpcClient.L2, 1*time.Hour),
		s.L1.TaikoL1SignalService,
	)
	s.Nil(err)
	s.s = syncer

	prop := new(proposer.Proposer)
	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests

	s.proverEndpoints, s.proverServer, err = helper.DefaultFakeProver(&s.ClientTestSuite, s.rpcClient)
	s.NoError(err)
	s.Nil(proposer.InitFromConfig(context.Background(), prop, (&proposer.Config{
		L1Endpoint:                         s.L1.WsEndpoint(),
		L2Endpoint:                         s.L2.WsEndpoint(),
		TaikoL1Address:                     s.L1.TaikoL1Address,
		TaikoL2Address:                     testutils.TaikoL2Address,
		TaikoTokenAddress:                  s.L1.TaikoL1TokenAddress,
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

	s.p = prop
}

func (s *CalldataSyncerTestSuite) TearDownTest() {
	s.proverServer.Shutdown(context.Background())
	s.p.Close(context.Background())
	s.rpcClient.Close()
	s.ClientTestSuite.TearDownTest()
}

func (s *CalldataSyncerTestSuite) TestCancelNewSyncer() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	syncer, err := NewSyncer(
		ctx,
		s.rpcClient,
		s.s.state,
		s.s.progressTracker,
		s.L1.TaikoL1SignalService,
	)
	s.Nil(syncer)
	s.NotNil(err)
}

func (s *CalldataSyncerTestSuite) TestProcessL1Blocks() {
	head, err := s.s.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Nil(s.s.ProcessL1Blocks(context.Background(), head))
}

func (s *CalldataSyncerTestSuite) TestProcessL1BlocksReorg() {
	head, err := s.s.rpc.L1.HeaderByNumber(context.Background(), nil)
	helper.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s)
	s.NoError(err)
	s.NoError(s.s.ProcessL1Blocks(context.Background(), head))
}

func (s *CalldataSyncerTestSuite) TestOnBlockProposed() {
	s.Nil(s.s.onBlockProposed(
		context.Background(),
		&bindings.TaikoL1ClientBlockProposed{BlockId: common.Big0},
		func() {},
	))
	s.NotNil(s.s.onBlockProposed(
		context.Background(),
		&bindings.TaikoL1ClientBlockProposed{BlockId: common.Big1},
		func() {},
	))
}

func (s *CalldataSyncerTestSuite) TestInsertNewHead() {
	parent, err := s.s.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l1Head, err := s.s.rpc.L1.BlockByNumber(context.Background(), nil)
	s.Nil(err)
	_, err = s.s.insertNewHead(
		context.Background(),
		&bindings.TaikoL1ClientBlockProposed{
			BlockId: common.Big1,
			Meta: bindings.TaikoDataBlockMetadata{
				Id:         1,
				L1Height:   l1Head.NumberU64(),
				L1Hash:     l1Head.Hash(),
				Proposer:   common.BytesToAddress(helper.RandomBytes(1024)),
				TxListHash: helper.RandomHash(),
				MixHash:    helper.RandomHash(),
				GasLimit:   rand.Uint32(),
				Timestamp:  uint64(time.Now().Unix()),
			},
		},
		parent,
		common.Big2,
		[]byte{},
		&rawdb.L1Origin{
			BlockID:       common.Big1,
			L1BlockHeight: common.Big1,
			L1BlockHash:   helper.RandomHash(),
		},
	)
	s.Nil(err)
}

func (s *CalldataSyncerTestSuite) TestTreasuryIncomeAllAnchors() {
	treasury := testutils.TreasuryAddress
	s.NotZero(treasury.Big().Uint64())

	balance, err := s.rpcClient.L2.BalanceAt(context.Background(), treasury, nil)
	s.Nil(err)

	headBefore, err := s.rpcClient.L2.BlockNumber(context.Background())
	s.Nil(err)

	helper.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s)

	headAfter, err := s.rpcClient.L2.BlockNumber(context.Background())
	s.Nil(err)

	balanceAfter, err := s.rpcClient.L2.BalanceAt(context.Background(), treasury, nil)
	s.Nil(err)

	s.Greater(headAfter, headBefore)
	s.Zero(balanceAfter.Cmp(balance))
}

func (s *CalldataSyncerTestSuite) TestTreasuryIncome() {
	treasury := testutils.TreasuryAddress
	s.NotZero(treasury.Big().Uint64())

	balance, err := s.rpcClient.L2.BalanceAt(context.Background(), treasury, nil)
	s.Nil(err)

	headBefore, err := s.rpcClient.L2.BlockNumber(context.Background())
	s.Nil(err)

	helper.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s)
	helper.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.s)

	headAfter, err := s.rpcClient.L2.BlockNumber(context.Background())
	s.Nil(err)

	balanceAfter, err := s.rpcClient.L2.BalanceAt(context.Background(), treasury, nil)
	s.Nil(err)

	s.Greater(headAfter, headBefore)
	s.True(balanceAfter.Cmp(balance) > 0)

	var hasNoneAnchorTxs bool
	for i := headBefore + 1; i <= headAfter; i++ {
		block, err := s.rpcClient.L2.BlockByNumber(context.Background(), new(big.Int).SetUint64(i))
		s.Nil(err)
		s.GreaterOrEqual(block.Transactions().Len(), 1)
		s.Greater(block.BaseFee().Uint64(), uint64(0))

		for j, tx := range block.Transactions() {
			if j == 0 {
				continue
			}

			hasNoneAnchorTxs = true
			receipt, err := s.rpcClient.L2.TransactionReceipt(context.Background(), tx.Hash())
			s.Nil(err)

			fee := new(big.Int).Mul(block.BaseFee(), new(big.Int).SetUint64(receipt.GasUsed))

			balance = new(big.Int).Add(balance, fee)
		}
	}

	s.True(hasNoneAnchorTxs)
	s.Zero(balanceAfter.Cmp(balance))
}

func TestCalldataSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(CalldataSyncerTestSuite))
}
