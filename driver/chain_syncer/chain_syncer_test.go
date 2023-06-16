package chainSyncer

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/testutils"
)

type ChainSyncerTestSuite struct {
	testutils.ClientTestSuite
	s *L2ChainSyncer
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

func TestChainSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(ChainSyncerTestSuite))
}
