package calldata

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	progressTracker "github.com/taikoxyz/taiko-client/driver/chain_syncer/progress_tracker"
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

	throwawayBlocksBuilderPrivKey, err := crypto.HexToECDSA(bindings.GoldenTouchPrivKey[2:])
	s.Nil(err)

	syncer, err := NewSyncer(
		context.Background(),
		s.RpcClient,
		state,
		progressTracker.New(s.RpcClient.L2, 1*time.Hour),
		throwawayBlocksBuilderPrivKey,
	)
	s.Nil(err)
	s.s = syncer
}

func (s *CalldataSyncerTestSuite) TestGetInvalidateBlockTxOpts() {
	opts, err := s.s.getInvalidateBlockTxOpts(context.Background(), common.Big0)

	s.Nil(err)
	s.True(opts.NoSend)
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
	_, rpcErr, payloadErr := s.s.insertNewHead(
		context.Background(),
		&bindings.TaikoL1ClientBlockProposed{
			Id: common.Big1,
			Meta: bindings.TaikoDataBlockMetadata{
				Id:          common.Big1,
				L1Height:    common.Big1,
				L1Hash:      testutils.RandomHash(),
				Beneficiary: common.BytesToAddress(testutils.RandomBytes(1024)),
				TxListHash:  testutils.RandomHash(),
				MixHash:     testutils.RandomHash(),
				ExtraData:   []byte{},
				GasLimit:    rand.Uint64(),
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

func (s *CalldataSyncerTestSuite) TestInsertThrowAwayBlock() {
	parent, err := s.s.rpc.L2.HeaderByNumber(context.Background(), common.Big0)
	s.Nil(err)
	txListBytes := testutils.RandomBytes(1024)
	_, rpcErr, payloadErr := s.s.insertThrowAwayBlock(
		context.Background(),
		&bindings.TaikoL1ClientBlockProposed{
			Id: common.Big1,
			Meta: bindings.TaikoDataBlockMetadata{
				Id:          common.Big1,
				L1Height:    common.Big1,
				L1Hash:      testutils.RandomHash(),
				Beneficiary: common.BytesToAddress(testutils.RandomBytes(1024)),
				TxListHash:  crypto.Keccak256Hash(txListBytes),
				MixHash:     testutils.RandomHash(),
				ExtraData:   []byte{},
				GasLimit:    rand.Uint64(),
				Timestamp:   uint64(time.Now().Unix()),
			},
		},
		parent,
		2, // BINARY_NOT_DECODABLE
		common.Big0,
		common.Big2,
		txListBytes,
		&rawdb.L1Origin{
			BlockID:       common.Big1,
			L1BlockHeight: common.Big1,
			L1BlockHash:   testutils.RandomHash(),
		},
	)
	s.Nil(rpcErr)
	s.NotNil(payloadErr)
}

func TestCalldataSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(CalldataSyncerTestSuite))
}
