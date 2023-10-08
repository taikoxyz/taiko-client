package chainSyncer

import (
	"bytes"
	"context"

	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/testutils"
)

type ChainSyncerTestSuite struct {
	testutils.ClientTestSuite
	s          *L2ChainSyncer
	snapshotID string
	p          testutils.Proposer
}

func (s *ChainSyncerTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	state, err := state.New(context.Background(), s.RpcClient)
	s.Nil(err)

	syncer, err := New(
		context.Background(),
		s.RpcClient,
		state,
		false,
		1*time.Hour,
		common.HexToAddress(os.Getenv("L1_SIGNAL_SERVICE_CONTRACT_ADDRESS")),
	)
	s.Nil(err)
	s.s = syncer

	prop := new(proposer.Proposer)
	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)
	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests

	s.Nil(proposer.InitFromConfig(context.Background(), prop, (&proposer.Config{
		L1Endpoint:                    os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:                    os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:                common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:                common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		TaikoTokenAddress:             common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		L1ProposerPrivKey:             l1ProposerPrivKey,
		L2SuggestedFeeRecipient:       common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:               &proposeInterval,
		MaxProposedTxListsPerEpoch:    1,
		WaitReceiptTimeout:            10 * time.Second,
		ProverEndpoints:               s.ProverEndpoints,
		OptimisticTierFee:             common.Big256,
		SgxTierFee:                    common.Big256,
		PseZkevmTierFee:               common.Big256,
		MaxTierFeePriceBumpIterations: 3,
		TierFeePriceBump:              common.Big2,
		ExtraData:                     "test",
	})))

	s.p = prop
}

func (s *ChainSyncerTestSuite) TestGetInnerSyncers() {
	s.NotNil(s.s.BeaconSyncer())
	s.NotNil(s.s.CalldataSyncer())
}

func (s *ChainSyncerTestSuite) TestSync() {
	head, err := s.RpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Nil(s.s.Sync(head))
}

func (s *ChainSyncerTestSuite) TestAheadOfProtocolVerifiedHead2() {
	s.TakeSnapshot()
	// propose a couple blocks
	testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s.calldataSyncer)

	// NOTE: need to prove the proposed blocks to be verified, writing helper function
	// generate transactopts to interact with TaikoL1 contract with.
	privKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)
	opts, err := bind.NewKeyedTransactorWithChainID(privKey, s.RpcClient.L1ChainID)
	s.Nil(err)

	head, err := s.RpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head, err := s.RpcClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Equal("test", string(bytes.TrimRight(l2Head.Extra, "\x00")))
	log.Info("L1HeaderByNumber head", "number", head.Number)
	// (equiv to s.state.GetL2Head().Number)
	log.Info("L2HeaderByNumber head", "number", l2Head.Number)
	log.Info("LatestVerifiedBlock number", "number", s.s.state.GetLatestVerifiedBlock().ID.Uint64())

	// increase evm time to make blocks verifiable.
	var result uint64
	s.Nil(s.RpcClient.L1RawRPC.CallContext(
		context.Background(),
		&result,
		"evm_increaseTime",
		(1024 * time.Hour).Seconds(),
	))
	s.NotNil(result)
	log.Info("EVM time increase", "number", result)

	// interact with TaikoL1 contract to allow for verification of L2 blocks
	tx, err := s.s.rpc.TaikoL1.VerifyBlocks(opts, uint64(3))
	s.Nil(err)
	s.NotNil(tx)

	head2, err := s.RpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head2, err := s.RpcClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	log.Info("L1HeaderByNumber head2", "number", head2.Number)
	log.Info("L2HeaderByNumber head", "number", l2Head2.Number)
	log.Info("LatestVerifiedBlock number", "number", s.s.state.GetLatestVerifiedBlock().ID.Uint64())

	s.RevertSnapshot()
}

func TestChainSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(ChainSyncerTestSuite))
}

func (s *ChainSyncerTestSuite) TakeSnapshot() {
	// record snapshot state to revert to before changes
	s.Nil(s.RpcClient.L1RawRPC.CallContext(context.Background(), &s.snapshotID, "evm_snapshot"))
}

func (s *ChainSyncerTestSuite) RevertSnapshot() {
	// revert to the snapshot state so protocol configs are unaffected
	var revertRes bool
	s.Nil(s.RpcClient.L1RawRPC.CallContext(context.Background(), &revertRes, "evm_revert", s.snapshotID))
	s.True(revertRes)
	s.Nil(rpc.SetHead(context.Background(), s.RpcClient.L2RawRPC, common.Big0))
}

func (s *ChainSyncerTestSuite) TestAheadOfProtocolVerifiedHead() {
	s.True(s.s.AheadOfProtocolVerifiedHead())
}
