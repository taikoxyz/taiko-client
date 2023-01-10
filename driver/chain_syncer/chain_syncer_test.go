package chainSyncer

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
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

	throwawayBlocksBuilderPrivKey, err := crypto.HexToECDSA(bindings.GoldenTouchPrivKey[2:])
	s.Nil(err)

	syncer, err := New(
		context.Background(),
		s.RpcClient,
		state,
		throwawayBlocksBuilderPrivKey,
		false,
		1*time.Hour,
	)
	s.Nil(err)
	s.s = syncer
}

func (s *ChainSyncerTestSuite) TestGetInnerSyncers() {
	s.NotNil(s.s.BeaconSyncer())
	s.NotNil(s.s.CalldataSyncer())
}

func TestChainSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(ChainSyncerTestSuite))
}
