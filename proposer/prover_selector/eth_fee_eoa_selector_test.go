package selector

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/prover/server"
	"github.com/taikoxyz/taiko-client/testutils"
	"github.com/taikoxyz/taiko-client/testutils/helper"
)

type ProverSelectorTestSuite struct {
	testutils.ClientTestSuite
	s               *ETHFeeEOASelector
	proverAddress   common.Address
	rpcClient       *rpc.Client
	proverEndpoints []*url.URL
	proverServer    *server.ProverServer
}

func (s *ProverSelectorTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()
	s.rpcClient = helper.NewWsRpcClient(&s.ClientTestSuite)
	s.proverAddress = crypto.PubkeyToAddress(testutils.ProverPrivKey.PublicKey)
	protocolConfigs, err := s.rpcClient.TaikoL1.GetConfig(nil)
	s.Nil(err)
	s.proverEndpoints, s.proverServer, err = helper.DefaultFakeProver(&s.ClientTestSuite, s.rpcClient)
	s.NoError(err)
	s.s, err = NewETHFeeEOASelector(
		&protocolConfigs,
		s.rpcClient,
		s.L1.TaikoL1Address,
		common.Big256,
		common.Big2,
		[]*url.URL{s.proverEndpoints[0]},
		32,
		1*time.Minute,
		1*time.Minute,
	)
	s.Nil(err)
}

func (s *ProverSelectorTestSuite) TearDownTest() {
	_ = s.proverServer.Shutdown(context.Background())
	s.rpcClient.Close()
	s.ClientTestSuite.TearDownTest()
}

func TestProverSelectorTestSuite(t *testing.T) {
	suite.Run(t, new(ProverSelectorTestSuite))
}
