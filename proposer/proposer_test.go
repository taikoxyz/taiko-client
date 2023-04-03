package proposer

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
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/testutils"
)

type ProposerTestSuite struct {
	testutils.ClientTestSuite
	p      *Proposer
	cancel context.CancelFunc
}

func (s *ProposerTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	p := new(Proposer)

	ctx, cancel := context.WithCancel(context.Background())
	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.Nil(InitFromConfig(ctx, p, (&Config{
		L1Endpoint:              os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:              os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
		TaikoL1Address:          common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:          common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:       l1ProposerPrivKey,
		L2SuggestedFeeRecipient: common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:         &proposeInterval,
	})))

	s.p = p
	s.p.AfterCommitHook = s.MineL1Confirmations
	s.cancel = cancel
}

func (s *ProposerTestSuite) TestSumTxsGasLimit() {
	txs := []*types.Transaction{
		types.NewTransaction(0, common.Address{}, common.Big0, 1, common.Big0, []byte{}), // gasLimit: 1
		types.NewTransaction(0, common.Address{}, common.Big0, 2, common.Big0, []byte{}), // gasLimit: 2
		types.NewTransaction(0, common.Address{}, common.Big0, 3, common.Big0, []byte{}), // gasLimit: 3
	}

	s.Equal(uint64(1+2+3), sumTxsGasLimit(txs))
}

func (s *ProposerTestSuite) TestName() {
	s.Equal("proposer", s.p.Name())
}

func (s *ProposerTestSuite) TestProposeOp() {
	// Nothing to propose
	s.EqualError(errNoNewTxs, s.p.ProposeOp(context.Background(), 0).Error())

	// Propose txs in L2 execution engine's mempool
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := s.p.rpc.TaikoL1.WatchBlockProposed(nil, sink, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	nonce, err := s.p.rpc.L2.PendingNonceAt(context.Background(), s.TestAddr)
	s.Nil(err)

	tx := types.NewTransaction(
		nonce,
		common.BytesToAddress(testutils.RandomBytes(32)),
		common.Big1,
		100000,
		common.Big1,
		[]byte{},
	)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(s.p.rpc.L2ChainID), s.TestAddrPrivKey)
	s.Nil(err)
	s.Nil(s.p.rpc.L2.SendTransaction(context.Background(), signedTx))

	s.Nil(s.p.ProposeOp(context.Background(), 0))

	event := <-sink

	_, isPending, err := s.p.rpc.L1.TransactionByHash(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.False(isPending)
	s.Equal(s.p.l2SuggestedFeeRecipient, event.Meta.Beneficiary)

	receipt, err := s.p.rpc.L1.TransactionReceipt(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.Equal(types.ReceiptStatusSuccessful, receipt.Status)
}

func (s *ProposerTestSuite) TestProposeEmptyBlockOp() {
	s.Nil(s.p.ProposeEmptyBlockOp(context.Background()))
}

func (s *ProposerTestSuite) TestCommitTxList() {
	txListBytes := testutils.RandomBytes(1024)
	gasLimit := uint64(102400)

	meta, tx, err := s.p.CommitTxList(context.Background(), txListBytes, gasLimit, 0)
	s.Nil(err)
	s.Equal(meta.GasLimit, gasLimit)

	if s.p.protocolConfigs.CommitConfirmations.Cmp(common.Big0) > 0 {
		receipt, err := rpc.WaitReceipt(context.Background(), s.p.rpc.L1, tx)
		s.Nil(err)
		s.Equal(types.ReceiptStatusSuccessful, receipt.Status)
	}
}

func (s *ProposerTestSuite) TestCustomProposeOpHook() {
	flag := false

	s.p.CustomProposeOpHook = func() error {
		flag = true
		return nil
	}

	s.Nil(s.p.ProposeOp(context.Background(), 0))
	s.True(flag)
}

func (s *ProposerTestSuite) TestUpdateProposingTicker() {
	oneHour := 1 * time.Hour
	s.p.proposingInterval = &oneHour
	s.NotPanics(s.p.updateProposingTicker)

	s.p.proposingInterval = nil
	s.NotPanics(s.p.updateProposingTicker)
}

func (s *ProposerTestSuite) TestStartClose() {
	s.Nil(s.p.Start())
	s.cancel()
	s.NotPanics(s.p.Close)
}

func TestProposerTestSuite(t *testing.T) {
	suite.Run(t, new(ProposerTestSuite))
}
