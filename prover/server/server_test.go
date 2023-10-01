package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-resty/resty/v2"
	echo "github.com/labstack/echo/v4"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	capacity "github.com/taikoxyz/taiko-client/prover/capacity_manager"
	"github.com/taikoxyz/taiko-client/testutils"
)

type ProverServerTestSuite struct {
	testutils.ClientTestSuite
	ps        *ProverServer
	ws        *httptest.Server // web server
	rpcClient *rpc.Client
}

func (s *ProverServerTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()
	l1ProverPrivKey := testutils.ProverPrivKey
	var err error
	timeout := 5 * time.Second
	s.rpcClient, err = rpc.NewClient(context.Background(), &rpc.ClientConfig{
		L1Endpoint:        s.L1.WsEndpoint(),
		L2Endpoint:        s.L2.WsEndpoint(),
		TaikoL1Address:    s.L1.TaikoL1Address,
		TaikoL2Address:    testutils.TaikoL2Address,
		TaikoTokenAddress: s.L1.TaikoL1TokenAddress,
		L2EngineEndpoint:  s.L2.AuthEndpoint(),
		JwtSecret:         testutils.JwtSecretFile,
		RetryInterval:     backoff.DefaultMaxInterval,
		Timeout:           &timeout,
	})
	s.Nil(err)

	p := &ProverServer{
		echo:             echo.New(),
		proverPrivateKey: l1ProverPrivKey,
		minProofFee:      common.Big1,
		maxExpiry:        24 * time.Hour,
		capacityManager:  capacity.New(1024, 100*time.Second),
		taikoL1Address:   s.L1.TaikoL1Address,
		rpc:              s.rpcClient,
		bond:             common.Big0,
		isOracle:         false,
	}

	p.echo.HideBanner = true
	p.configureMiddleware()
	p.configureRoutes()
	s.ps = p
	s.ws = httptest.NewServer(p.echo)
}

func (s *ProverServerTestSuite) TearDownTest() {
	s.rpcClient.Close()
	s.ws.Close()
	s.ClientTestSuite.TearDownTest()
}

func (s *ProverServerTestSuite) TestHealth() {
	s.Equal(http.StatusOK, s.sendReq("/healthz").StatusCode)
}

func (s *ProverServerTestSuite) TestRoot() {
	s.Equal(http.StatusOK, s.sendReq("/").StatusCode)
}

func (s *ProverServerTestSuite) TestStartShutdown() {
	port, err := freeport.GetFreePort()
	s.Nil(err)

	url, err := url.Parse(fmt.Sprintf("http://localhost:%v", port))
	s.Nil(err)

	go func() {
		if err := s.ps.Start(fmt.Sprintf(":%v", port)); err != nil {
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

	s.Nil(s.ps.Shutdown(context.Background()))
}

func TestProverServerTestSuite(t *testing.T) {
	suite.Run(t, new(ProverServerTestSuite))
}

func (s *ProverServerTestSuite) sendReq(path string) *http.Response {
	resp, err := http.Get(s.ws.URL + path)
	s.Nil(err)
	return resp
}
