package testutils

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	capacity "github.com/taikoxyz/taiko-client/prover/capacity_manager"
	"github.com/taikoxyz/taiko-client/prover/server"
)

const (
	premintTokenAmount = "92233720368547758070000000000000"
	proposerPrivKey    = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	proverPrivKey      = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
)

var (
	taikoL2Address      = common.HexToAddress("0x1000777700000000000000000000000000000001")
	oracleProverAddress = common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
)

// variables need to be initialized
var (
	TestPrivKey   *ecdsa.PrivateKey
	TestAddr      common.Address
	ProverPrivKey *ecdsa.PrivateKey
	ProverAddr    common.Address
)

type ClientSuite struct {
	suite.Suite
	l1Container     *gethContainer
	l2Container     *gethContainer
	RpcClient       *rpc.Client
	ProverEndpoints []*url.URL
	proverServer    *server.ProverServer
}

func (s *ClientSuite) SetupTest() {
	var err error
	name := strings.ReplaceAll(s.T().Name(), "/", "_")
	s.l1Container, err = newL1Container("L1_" + name)
	s.NoError(err)

	s.l2Container, err = newL2Container("L2_" + name)
	s.NoError(err)

	s.RpcClient, err = rpc.NewClient(context.Background(), &rpc.ClientConfig{
		L1Endpoint:        s.l1Container.WsEndpoint(),
		L2Endpoint:        s.l2Container.WsEndpoint(),
		TaikoL1Address:    TaikoL1Address,
		TaikoTokenAddress: TaikoTokenAddress,
		TaikoL2Address:    taikoL2Address,
		L2EngineEndpoint:  s.l2Container.AuthEndpoint(),
		JwtSecret:         string(jwtSecret),
		RetryInterval:     backoff.DefaultMaxInterval,
	})
	s.NoError(err)
	s.ProverEndpoints = []*url.URL{LocalRandomProverEndpoint()}
	s.proverServer = fakeProverServer(s, ProverPrivKey, capacity.New(1024, 100*time.Second), s.ProverEndpoints[0])
}

// fakeProverServer starts a new prover server that has channel listeners to respond and react
// to requests for capacity, which provers can call.
func fakeProverServer(
	s *ClientSuite,
	proverPrivKey *ecdsa.PrivateKey,
	capacityManager *capacity.CapacityManager,
	url *url.URL,
) *server.ProverServer {
	protocolConfig, err := s.RpcClient.TaikoL1.GetConfig(nil)
	s.Nil(err)

	srv, err := server.New(&server.NewProverServerOpts{
		ProverPrivateKey: proverPrivKey,
		MinProofFee:      common.Big1,
		MaxExpiry:        24 * time.Hour,
		CapacityManager:  capacityManager,
		TaikoL1Address:   TaikoL1Address,
		Rpc:              s.RpcClient,
		Bond:             protocolConfig.ProofBond,
		IsOracle:         true,
	})
	s.NoError(err)

	go func() {
		if err := srv.Start(fmt.Sprintf(":%v", url.Port())); !errors.Is(err, http.ErrServerClosed) {
			log.Error("Failed to start prover server", "error", err)
		}
	}()

	// Wait till the server fully started.
	s.Nil(backoff.Retry(func() error {
		res, err := resty.New().R().Get(url.String() + "/healthz")
		if err != nil {
			return err
		}
		if !res.IsSuccess() {
			return fmt.Errorf("invalid response status code: %d", res.StatusCode())
		}

		return nil
	}, backoff.NewExponentialBackOff()))

	return srv
}

func (s *ClientSuite) TearDownTest() {
	s.NoError(s.l1Container.Stop())
	s.NoError(s.l2Container.Stop())
}
