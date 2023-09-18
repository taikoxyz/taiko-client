package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	capacity "github.com/taikoxyz/taiko-client/prover/capacity_manager"
)

type ProverServerTestSuite struct {
	suite.Suite
	ps *ProverServer
	ws *httptest.Server // web server
}

func (s *ProverServerTestSuite) SetupTest() {
	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	p := &ProverServer{
		echo:             echo.New(),
		proverPrivateKey: l1ProverPrivKey,
		minProofFee:      common.Big1,
		maxExpiry:        24 * time.Hour,
		capacityManager:  capacity.New(1024),
	}

	p.echo.HideBanner = true
	p.configureMiddleware()
	p.configureRoutes()
	s.ps = p
	s.ws = httptest.NewServer(p.echo)
}

func (s *ProverServerTestSuite) TestHealth() {
	s.Equal(http.StatusOK, s.sendReq("/healthz").StatusCode)
}

func (s *ProverServerTestSuite) TestRoot() {
	s.Equal(http.StatusOK, s.sendReq("/").StatusCode)
}

func (s *ProverServerTestSuite) TearDownTest() {
	s.ws.Close()
}

func TestProverServerTestSuite(t *testing.T) {
	suite.Run(t, new(ProverServerTestSuite))
}

func (s *ProverServerTestSuite) sendReq(path string) *http.Response {
	resp, err := http.Get(s.ws.URL + path)
	s.Nil(err)
	return resp
}
