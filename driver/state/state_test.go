package state

import (
	"context"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/testutils"
	"github.com/taikoxyz/taiko-client/testutils/helper"
)

type DriverStateTestSuite struct {
	testutils.ClientTestSuite
	s         *State
	rpcClient *rpc.Client
}

func (s *DriverStateTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()
	s.rpcClient = helper.NewWsRpcClient(&s.ClientTestSuite)
	state, err := New(context.Background(), s.rpcClient)
	s.Nil(err)
	s.s = state
}

func (s *DriverStateTestSuite) TearDownTest() {
	s.s.Close()
	s.rpcClient.Close()
	s.ClientTestSuite.TearDownTest()
}

func (s *DriverStateTestSuite) TestVerifyL2Block() {
	head, err := s.rpcClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Nil(s.s.VerifyL2Block(context.Background(), head.Number, head.Hash()))
}

func (s *DriverStateTestSuite) TestGetL1Head() {
	l1Head := s.s.GetL1Head()
	s.NotNil(l1Head)
}

func (s *DriverStateTestSuite) TestGetLatestVerifiedBlock() {
	b := s.s.GetLatestVerifiedBlock()
	s.NotNil(b.Hash)
}

func (s *DriverStateTestSuite) TestGetHeadBlockID() {
	s.Equal(uint64(0), s.s.GetHeadBlockID().Uint64())
}

func (s *DriverStateTestSuite) TestHeightOrIDNotEmpty() {
	s.False((&HeightOrID{}).NotEmpty())
	s.True((&HeightOrID{Height: common.Big0}).NotEmpty())
	s.True((&HeightOrID{ID: common.Big0}).NotEmpty())
}

func (s *DriverStateTestSuite) TestClose() {
	s.NotPanics(s.s.Close)
}

func (s *DriverStateTestSuite) TestGetL2Head() {
	testHeight := rand.Uint64()

	s.s.setL2Head(nil)
	s.s.setL2Head(&types.Header{Number: new(big.Int).SetUint64(testHeight)})
	h := s.s.GetL2Head()
	s.Equal(testHeight, h.Number.Uint64())
}

func (s *DriverStateTestSuite) TestSubL1HeadsFeed() {
	s.NotNil(s.s.SubL1HeadsFeed(make(chan *types.Header)))
}

func (s *DriverStateTestSuite) TestGetSyncedHeaderID() {
	l2Genesis, err := s.rpcClient.L2.BlockByNumber(context.Background(), common.Big0)
	s.Nil(err)

	id, err := s.s.getSyncedHeaderID(context.Background(), s.s.GenesisL1Height.Uint64(), l2Genesis.Hash())
	s.Nil(err)
	s.Zero(id.Uint64())
}

func (s *DriverStateTestSuite) TestNewDriverContextErr() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	state, err := New(ctx, s.rpcClient)
	s.Nil(state)
	s.ErrorContains(err, "context canceled")
}

func (s *DriverStateTestSuite) TestDriverInitContextErr() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := s.s.init(ctx)
	s.ErrorContains(err, "context canceled")
}

func TestDriverStateTestSuite(t *testing.T) {
	suite.Run(t, new(DriverStateTestSuite))
}
