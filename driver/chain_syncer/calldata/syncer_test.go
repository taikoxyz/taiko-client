package calldata

import (
	"context"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/beaconsync"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/testutils"
)

type CalldataSyncerTestSuite struct {
	testutils.ClientTestSuite
	s *Syncer
	p testutils.Proposer
}

func (s *CalldataSyncerTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	state, err := state.New(context.Background(), s.RpcClient)
	s.Nil(err)

	syncer, err := NewSyncer(
		context.Background(),
		s.RpcClient,
		state,
		beaconsync.NewSyncProgressTracker(s.RpcClient.L2, 1*time.Hour),
		common.HexToAddress(os.Getenv("L1_SIGNAL_SERVICE_CONTRACT_ADDRESS")),
	)
	s.Nil(err)
	s.s = syncer

	prop := new(proposer.Proposer)
	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)
	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.Nil(proposer.InitFromConfig(context.Background(), prop, (&proposer.Config{
		L1Endpoint:              os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:              os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:          common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:          common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:       l1ProposerPrivKey,
		L2SuggestedFeeRecipient: common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:         &proposeInterval,
	})))

	s.p = prop
}

func (s *CalldataSyncerTestSuite) TestProcessL1Blocks() {
	head, err := s.s.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Nil(s.s.ProcessL1Blocks(context.Background(), head))
}

func (s *CalldataSyncerTestSuite) TestOnBlockProposed() {
	s.Nil(s.s.onBlockProposed(context.Background(), &bindings.TaikoL1ClientBlockProposed{Id: common.Big0}, func() {}))
	s.NotNil(s.s.onBlockProposed(context.Background(), &bindings.TaikoL1ClientBlockProposed{Id: common.Big1}, func() {}))
}

func (s *CalldataSyncerTestSuite) TestInsertNewHead() {
	parent, err := s.s.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l1Head, err := s.s.rpc.L1.BlockByNumber(context.Background(), nil)
	s.Nil(err)
	_, rpcErr, payloadErr := s.s.insertNewHead(
		context.Background(),
		&bindings.TaikoL1ClientBlockProposed{
			Id: common.Big1,
			Meta: bindings.TaikoDataBlockMetadata{
				Id:          1,
				L1Height:    l1Head.NumberU64(),
				L1Hash:      l1Head.Hash(),
				Beneficiary: common.BytesToAddress(testutils.RandomBytes(1024)),
				TxListHash:  testutils.RandomHash(),
				MixHash:     testutils.RandomHash(),
				GasLimit:    rand.Uint32(),
				Timestamp:   uint64(time.Now().Unix()),
			},
		},
		parent,
		common.Big2,
		[]byte{},
		&rawdb.L1Origin{
			BlockID:       common.Big1,
			L1BlockHeight: common.Big1,
			L1BlockHash:   testutils.RandomHash(),
		},
	)
	s.Nil(rpcErr)
	s.Nil(payloadErr)
}

func (s *CalldataSyncerTestSuite) TestHandleReorgToGenesis() {
	testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s)

	l2Head1, err := s.s.rpc.L2.BlockByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(l2Head1.NumberU64(), uint64(0))
	s.NotZero(s.s.lastInsertedBlockID.Uint64())
	s.s.lastInsertedBlockID = common.Big0 // let the chain reorg to genesis

	s.Nil(s.s.handleReorg(context.Background(), &bindings.TaikoL1ClientBlockProposed{
		Id:  l2Head1.Number(),
		Raw: types.Log{Removed: true},
	}))

	l2Head2, err := s.s.rpc.L2.BlockByNumber(context.Background(), nil)
	s.Nil(err)
	s.Equal(uint64(0), l2Head2.NumberU64())
}

func (s *CalldataSyncerTestSuite) TestHandleReorgToNoneGenesis() {
	testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s)

	l2Head1, err := s.s.rpc.L2.BlockByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(l2Head1.NumberU64(), uint64(0))
	s.NotZero(s.s.lastInsertedBlockID.Uint64())
	s.s.lastInsertedBlockID = common.Big1 // let the chain reorg to height 1

	s.Nil(s.s.handleReorg(context.Background(), &bindings.TaikoL1ClientBlockProposed{
		Id:  l2Head1.Number(),
		Raw: types.Log{Removed: true},
	}))

	l2Head2, err := s.s.rpc.L2.BlockByNumber(context.Background(), nil)
	s.Nil(err)
	s.Equal(uint64(1), l2Head2.NumberU64())

	testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s)
	l2Head3, err := s.s.rpc.L2.BlockByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(l2Head3.NumberU64(), l2Head2.NumberU64())
	s.Greater(s.s.lastInsertedBlockID.Uint64(), uint64(1))
}

func (s *CalldataSyncerTestSuite) TestWithdrawRootCalculation() {
	events := testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s)

	for _, e := range events {
		header, err := s.s.rpc.L2.HeaderByNumber(context.Background(), e.Id)
		s.Nil(err)
		s.NotEmpty(e.Meta.DepositsRoot)
		s.Equal(common.BytesToHash(e.Meta.DepositsRoot[:]), *header.WithdrawalsHash)
	}
}

func TestCalldataSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(CalldataSyncerTestSuite))
}
