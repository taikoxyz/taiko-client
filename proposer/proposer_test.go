package proposer

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
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
		L1Endpoint:                          os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:                          os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
		TaikoL1Address:                      common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:                      common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		TaikoTokenAddress:                   common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		L1ProposerPrivKey:                   l1ProposerPrivKey,
		L2SuggestedFeeRecipient:             common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:                     &proposeInterval,
		MaxProposedTxListsPerEpoch:          1,
		ProposeBlockTxReplacementMultiplier: 2,
		WaitReceiptTimeout:                  10 * time.Second,
		ProverEndpoints:                     s.ProverEndpoints,
		OptimisticTierFee:                   common.Big256,
		SgxTierFee:                          common.Big256,
		PseZkevmTierFee:                     common.Big256,
		TierFeePriceBump:                    common.Big2,
		MaxTierFeePriceBumpIterations:       3,
		ExtraData:                           "test",
	})))

	s.p = p
	s.cancel = cancel
}

func (s *ProposerTestSuite) TestName() {
	s.Equal("proposer", s.p.Name())
}

func (s *ProposerTestSuite) TestProposeOp() {
	// Propose txs in L2 execution engine's mempool
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := s.p.rpc.TaikoL1.WatchBlockProposed(nil, sink, nil, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	nonce, err := s.p.rpc.L2.PendingNonceAt(context.Background(), s.TestAddr)
	s.Nil(err)

	gaslimit := 21000

	parent, err := s.p.rpc.L2.BlockByNumber(context.Background(), nil)
	s.Nil(err)

	baseFee, err := s.p.rpc.TaikoL2.GetBasefee(nil, 1, uint32(parent.GasUsed()))
	s.Nil(err)

	to := common.BytesToAddress(testutils.RandomBytes(32))
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.RpcClient.L2ChainID,
		Nonce:     nonce,
		GasTipCap: common.Big0,
		GasFeeCap: new(big.Int).SetUint64(baseFee.Uint64() * 2),
		Gas:       uint64(gaslimit),
		To:        &to,
		Value:     common.Big1,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(s.p.rpc.L2ChainID), s.TestAddrPrivKey)
	s.Nil(err)
	s.Nil(s.p.rpc.L2.SendTransaction(context.Background(), signedTx))

	s.Nil(s.p.ProposeOp(context.Background()))

	event := <-sink

	_, isPending, err := s.p.rpc.L1.TransactionByHash(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.False(isPending)

	receipt, err := s.p.rpc.L1.TransactionReceipt(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.Equal(types.ReceiptStatusSuccessful, receipt.Status)
}

func (s *ProposerTestSuite) TestProposeEmptyBlockOp() {
	s.Nil(s.p.ProposeEmptyBlockOp(context.Background()))
}

func (s *ProposerTestSuite) TestCustomProposeOpHook() {
	flag := false

	s.p.CustomProposeOpHook = func() error {
		flag = true
		return nil
	}

	s.Nil(s.p.ProposeOp(context.Background()))
	s.True(flag)
}

func (s *ProposerTestSuite) TestSendProposeBlockTx() {
	fee := big.NewInt(10000)
	opts, err := getTxOpts(
		context.Background(),
		s.p.rpc.L1,
		s.p.proposerPrivKey,
		s.RpcClient.L1ChainID,
		fee,
	)
	s.Nil(err)
	s.Greater(opts.GasTipCap.Uint64(), uint64(0))

	nonce, err := s.RpcClient.L1.PendingNonceAt(context.Background(), s.p.proposerAddress)
	s.Nil(err)

	tx := types.NewTransaction(
		nonce,
		common.BytesToAddress([]byte{}),
		common.Big1,
		100000,
		opts.GasTipCap,
		[]byte{},
	)

	s.SetL1Automine(false)
	defer s.SetL1Automine(true)

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(s.RpcClient.L1ChainID), s.p.proposerPrivKey)
	s.Nil(err)
	s.Nil(s.RpcClient.L1.SendTransaction(context.Background(), signedTx))

	var emptyTxs []types.Transaction
	encoded, err := rlp.EncodeToBytes(emptyTxs)
	s.Nil(err)

	signedAssignment, fee, err := s.p.proverSelector.AssignProver(
		context.Background(),
		s.p.tierFees,
		crypto.Keccak256Hash(encoded),
	)
	s.Nil(err)

	newTx, err := s.p.sendProposeBlockTx(
		context.Background(),
		encoded,
		&nonce,
		signedAssignment,
		fee,
		true,
	)
	s.Nil(err)
	s.Greater(newTx.GasTipCap().Uint64(), tx.GasTipCap().Uint64())
}

func (s *ProposerTestSuite) TestAssignProverSuccessFirstRound() {
	s.SetL1Automine(false)
	defer s.SetL1Automine(true)

	_, fee, err := s.p.proverSelector.AssignProver(context.Background(), s.p.tierFees, testutils.RandomHash())

	s.Nil(err)
	s.Equal(fee.Uint64(), s.p.cfg.OptimisticTierFee.Uint64())
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
	s.NotPanics(func() { s.p.Close(context.Background()) })
}

func TestProposerTestSuite(t *testing.T) {
	suite.Run(t, new(ProposerTestSuite))
}
