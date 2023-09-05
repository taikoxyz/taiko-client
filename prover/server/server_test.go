package server

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func newTestServer(url string) *ProverServer {
	l1ProverPrivKey, _ := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))

	srv := &ProverServer{
		echo:                     echo.New(),
		proverPrivateKey:         l1ProverPrivKey,
		minProofFee:              big.NewInt(1),
		requestCurrentCapacityCh: make(chan struct{}, 1024),
		receiveCurrentCapacityCh: make(chan uint64, 1024),
	}

	srv.configureMiddleware()
	srv.configureRoutes()

	return srv
}

func TestHealth(t *testing.T) {
	srv := newTestServer("")

	req, _ := http.NewRequest(echo.GET, "/healthz", nil)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Test_Health expected code %v, got %v", http.StatusOK, rec.Code)
	}
}

func TestRoot(t *testing.T) {
	srv := newTestServer("")

	req, _ := http.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Test_Root expected code %v, got %v", http.StatusOK, rec.Code)
	}
}

func TestStartShutdown(t *testing.T) {
	srv := newTestServer("")

	go func() {
		_ = srv.Start(":3928")
	}()
	assert.Nil(t, srv.Shutdown(context.Background()))
}
