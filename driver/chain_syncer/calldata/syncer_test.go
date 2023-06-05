package calldata

import (
	"context"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
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
		L1Endpoint:                 os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:                 os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:             common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:             common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:          l1ProposerPrivKey,
		L2SuggestedFeeRecipient:    common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:            &proposeInterval,
		MaxProposedTxListsPerEpoch: 1,
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

func (s *CalldataSyncerTestSuite) TestTreasuryIncomeAllAnchors() {
	treasury := common.HexToAddress(os.Getenv("TREASURY"))
	s.NotZero(treasury.Big().Uint64())

	balance, err := s.RpcClient.L2.BalanceAt(context.Background(), treasury, nil)
	s.Nil(err)

	headBefore, err := s.RpcClient.L2.BlockNumber(context.Background())
	s.Nil(err)

	testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s)

	headAfter, err := s.RpcClient.L2.BlockNumber(context.Background())
	s.Nil(err)

	balanceAfter, err := s.RpcClient.L2.BalanceAt(context.Background(), treasury, nil)
	s.Nil(err)

	s.Greater(headAfter, headBefore)
	s.Zero(balanceAfter.Cmp(balance))
}

func (s *CalldataSyncerTestSuite) TestTreasuryIncome() {
	treasury := common.HexToAddress(os.Getenv("TREASURY"))
	s.NotZero(treasury.Big().Uint64())

	balance, err := s.RpcClient.L2.BalanceAt(context.Background(), treasury, nil)
	s.Nil(err)

	headBefore, err := s.RpcClient.L2.BlockNumber(context.Background())
	s.Nil(err)

	testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s)
	testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.s)

	headAfter, err := s.RpcClient.L2.BlockNumber(context.Background())
	s.Nil(err)

	balanceAfter, err := s.RpcClient.L2.BalanceAt(context.Background(), treasury, nil)
	s.Nil(err)

	s.Greater(headAfter, headBefore)
	s.True(balanceAfter.Cmp(balance) > 0)

	var hasNoneAnchorTxs bool
	for i := headBefore + 1; i <= headAfter; i++ {
		block, err := s.RpcClient.L2.BlockByNumber(context.Background(), new(big.Int).SetUint64(i))
		s.Nil(err)
		s.GreaterOrEqual(block.Transactions().Len(), 1)
		s.Greater(block.BaseFee().Uint64(), uint64(0))

		for j, tx := range block.Transactions() {
			if j == 0 {
				continue
			}

			hasNoneAnchorTxs = true
			receipt, err := s.RpcClient.L2.TransactionReceipt(context.Background(), tx.Hash())
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
