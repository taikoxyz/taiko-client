package selector

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	capacity "github.com/taikoxyz/taiko-client/prover/capacity_manager"
	"github.com/taikoxyz/taiko-client/prover/server"
	"github.com/taikoxyz/taiko-client/testutils"
	"github.com/taikoxyz/taiko-client/testutils/helper"
)

type ProverSelectorTestSuite struct {
	testutils.ClientSuite
	s               *ETHFeeEOASelector
	proverAddress   common.Address
	rpcClient       *rpc.Client
	proverEndpoints []*url.URL
	proverServer    *server.ProverServer
}

func (s *ProverSelectorTestSuite) SetupTest() {
	s.ClientSuite.SetupTest()
	jwtSecret, err := jwt.ParseSecretFromFile(testutils.JwtSecretFile)
	s.NoError(err)
	s.rpcClient, err = rpc.NewClient(context.Background(), &rpc.ClientConfig{
		L1Endpoint:        s.L1.WsEndpoint(),
		L2Endpoint:        s.L2.WsEndpoint(),
		TaikoL1Address:    s.L1.TaikoL1Address,
		TaikoTokenAddress: s.L1.TaikoL1TokenAddress,
		TaikoL2Address:    testutils.TaikoL2Address,
		L2EngineEndpoint:  s.L2.AuthEndpoint(),
		JwtSecret:         string(jwtSecret),
		RetryInterval:     backoff.DefaultMaxInterval,
	})
	s.NoError(err)
	l1ProverPrivKey := testutils.ProverPrivKey
	s.proverAddress = crypto.PubkeyToAddress(l1ProverPrivKey.PublicKey)

	protocolConfigs, err := s.rpcClient.TaikoL1.GetConfig(nil)
	s.Nil(err)
	s.proverEndpoints = []*url.URL{testutils.LocalRandomProverEndpoint()}
	s.proverServer, err = helper.NewFakeProver(s.L1.TaikoL1Address, &protocolConfigs, jwtSecret,
		s.rpcClient, testutils.ProverPrivKey, capacity.New(1024, 100*time.Second), s.proverEndpoints[0])
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
	s.proverServer.Shutdown(context.Background())
	s.rpcClient.Close()
	s.ClientSuite.TearDownTest()
}

func TestProverSelectorTestSuite(t *testing.T) {
	suite.Run(t, new(ProverSelectorTestSuite))
}
