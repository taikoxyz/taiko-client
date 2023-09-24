package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-resty/resty/v2"
	echo "github.com/labstack/echo/v4"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	capacity "github.com/taikoxyz/taiko-client/prover/capacity_manager"
)

type ProverServerTestSuite struct {
	suite.Suite
	srv *ProverServer
}

func (s *ProverServerTestSuite) SetupTest() {
	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	timeout := 5 * time.Second
	rpcClient, err := rpc.NewClient(context.Background(), &rpc.ClientConfig{
		L1Endpoint:        os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:        os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:    common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:    common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		TaikoTokenAddress: common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		L2EngineEndpoint:  os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
		JwtSecret:         os.Getenv("JWT_SECRET"),
		RetryInterval:     backoff.DefaultMaxInterval,
		Timeout:           &timeout,
	})
	s.Nil(err)

	srv := &ProverServer{
		echo:             echo.New(),
		proverPrivateKey: l1ProverPrivKey,
		minProofFee:      common.Big1,
		maxExpiry:        24 * time.Hour,
		capacityManager:  capacity.New(1024),
		taikoL1Address:   common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		rpc:              rpcClient,
		bond:             common.Big0,
		isOracle:         false,
	}

	srv.echo.HideBanner = true
	srv.configureMiddleware()
	srv.configureRoutes()

	s.srv = srv
}

func (s *ProverServerTestSuite) TestHealth() {
	s.Equal(http.StatusOK, s.sendReq("/healthz").Code)
}

func (s *ProverServerTestSuite) TestRoot() {
	s.Equal(http.StatusOK, s.sendReq("/").Code)
}

func (s *ProverServerTestSuite) TestStartShutdown() {
	port, err := freeport.GetFreePort()
	s.Nil(err)

	url, err := url.Parse(fmt.Sprintf("http://localhost:%v", port))
	s.Nil(err)

	go func() {
		if err := s.srv.Start(fmt.Sprintf(":%v", port)); err != nil {
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

	s.Nil(s.srv.Shutdown(context.Background()))
}

func (s *ProverServerTestSuite) TearDownTest() {
	s.Nil(s.srv.Shutdown(context.Background()))
}

func TestProverServerTestSuite(t *testing.T) {
	suite.Run(t, new(ProverServerTestSuite))
}

func (s *ProverServerTestSuite) sendReq(path string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(echo.GET, path, nil)
	s.Nil(err)
	rec := httptest.NewRecorder()

	s.srv.ServeHTTP(rec, req)

	return rec
}
