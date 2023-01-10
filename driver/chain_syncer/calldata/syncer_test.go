package calldata

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
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

func TestCalldataSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(CalldataSyncerTestSuite))
}
