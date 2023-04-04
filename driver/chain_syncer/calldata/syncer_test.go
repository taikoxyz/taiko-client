package calldata

import (
	"context"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/beaconsync"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/testutils"
)

type CalldataSyncerTestSuite struct {
	testutils.ClientTestSuite
	s *Syncer
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

func TestCalldataSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(CalldataSyncerTestSuite))
}
