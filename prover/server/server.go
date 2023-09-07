package server

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	capacity "github.com/taikoxyz/taiko-client/prover/capacity_manager"
)

// ProverServer represents a prover server instance.
type ProverServer struct {
	echo             *echo.Echo
	proverPrivateKey *ecdsa.PrivateKey
	proverAddress    common.Address

	// capacity-related configs
	capacityManager *capacity.CapacityManager
	minProofFee     *big.Int
}

// NewProverServerOpts contains all configurations for creating a prover server instance.
type NewProverServerOpts struct {
	ProverPrivateKey *ecdsa.PrivateKey
	MinProofFee      *big.Int
	CapacityManager  *capacity.CapacityManager
}

// New creates a new prover server instance.
func New(opts *NewProverServerOpts) (*ProverServer, error) {
	address := crypto.PubkeyToAddress(opts.ProverPrivateKey.PublicKey)
	srv := &ProverServer{
		proverPrivateKey: opts.ProverPrivateKey,
		proverAddress:    address,
		echo:             echo.New(),
		minProofFee:      opts.MinProofFee,
		capacityManager:  opts.CapacityManager,
	}

	srv.echo.HideBanner = true
	srv.configureMiddleware()
	srv.configureRoutes()

	return srv, nil
}

// Start starts the HTTP server.
func (srv *ProverServer) Start(address string) error {
	return srv.echo.Start(address)
}

// Shutdown shuts down the HTTP server.
func (srv *ProverServer) Shutdown(ctx context.Context) error {
	return srv.echo.Shutdown(ctx)
}

// ServeHTTP implements the `http.Handler` interface which serves HTTP requests.
func (srv *ProverServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.echo.ServeHTTP(w, r)
}

// Health endpoints for probes.
func (srv *ProverServer) Health(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// LogSkipper implements the `middleware.Skipper` interface.
func LogSkipper(c echo.Context) bool {
	switch c.Request().URL.Path {
	case "/healthz":
		return true
	case "/metrics":
		return true
	default:
		return true
	}
}

// configureMiddleware configures the server middlewares.
func (srv *ProverServer) configureMiddleware() {
	srv.echo.Use(middleware.RequestID())

	srv.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: LogSkipper,
		Format: `{"time":"${time_rfc3339_nano}","level":"INFO","message":{"id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"response_status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}",` +
			`"bytes_in":${bytes_in},"bytes_out":${bytes_out}}}` + "\n",
		Output: os.Stdout,
	}))
}

// configureRoutes contains all routes which will be used by prover server.
func (srv *ProverServer) configureRoutes() {
	srv.echo.GET("/", srv.Health)
	srv.echo.GET("/healthz", srv.Health)
	srv.echo.POST("/proposeBlock", srv.ProposeBlock)
}
