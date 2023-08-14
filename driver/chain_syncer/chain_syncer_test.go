package chainSyncer

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
		L1Endpoint:                 os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:                 os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:             common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:             common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:          l1ProposerPrivKey,
		L2SuggestedFeeRecipient:    common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:            &proposeInterval,
		MaxProposedTxListsPerEpoch: 1,
		WaitReceiptTimeout:         10 * time.Second,
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

// func (s *ChainSyncerTestSuite) TestSyncTriggerBeaconSync() {
// 	s.s.p2pSyncVerifiedBlocks = true
//  NOTE: need to increase the verified block as one of the conditions to trigger
//         needBeaconSyncTriggered()
// 	s.s.state.setLatestVerifiedBlockHash(common.Hash{})
// }

func (s *ChainSyncerTestSuite) TestAheadOfProtocolVerifiedHead2() {
	s.TakeSnapshot()
	// propose a couple blocks
	testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.s.calldataSyncer)

	var result uint64
	s.RpcClient.L1RawRPC.CallContext(context.Background(), &result, "evm_increaseTime", 2000)
	s.NotNil(result)
	fmt.Printf("evm time increase: %v\n", result)

	head, err := s.RpcClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	fmt.Printf("L1HeaderByNumber head: %v\n", head.Number)
	fmt.Printf("LatestVerifiedBlock number: %v\n", s.s.state.GetLatestVerifiedBlock().ID.Uint64())
	fmt.Printf("LatestL2Head Number: %v\n", s.s.state.GetL2Head().Number)

	// NOTE: verify the block so that the state returns a value > 0
	// Can't figure out how to do this, all the listed methods below don't work / give
	// nil pointer referencing errors or just aren't relevant.

	// s.s.rpc.TaikoL1.TaikoL1ClientTransactor.VerifyBlocks(nil, common.Big1)
	// s.s.rpc.TaikoL1.VerifyBlocks(nil, common.Big1)
	// s.Nil(s.s.state.VerifyL2Block(context.Background(), head.Number, head.Hash()))
	// tx, err := s.s.rpc.TaikoL1.VerifyBlocks(nil, common.Big1)
	// s.Nil(err)
	// err2 := s.s.rpc.TaikoL1.TaikoL1ClientCaller.contract.Call(nil, &out, "getBlock", common.Big0)
	// s.Nil(err)
	// fmt.Printf("tx: %v\n", tx.Hash().Hex())

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
