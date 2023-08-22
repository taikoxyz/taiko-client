package http

import (
	"context"
	"crypto/ecdsa"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/echo/v4/middleware"

	echo "github.com/labstack/echo/v4"
)

type Server struct {
	echo             *echo.Echo
	proverPrivateKey *ecdsa.PrivateKey
	proverAddress    common.Address

	// capacity related configs
	maxCapacity     uint64
	currentCapacity uint64
}

type NewServerOpts struct {
	ProverPrivateKey *ecdsa.PrivateKey
	MaxCapacity      uint64
}

func NewServer(opts NewServerOpts) (*Server, error) {
	address := crypto.PubkeyToAddress(opts.ProverPrivateKey.PublicKey)
	srv := &Server{
		proverPrivateKey: opts.ProverPrivateKey,
		proverAddress:    address,
		echo:             echo.New(),
		maxCapacity:      opts.MaxCapacity,
		currentCapacity:  0,
	}

	srv.configureMiddleware()
	srv.configureRoutes()

	return srv, nil
}

// Start starts the HTTP server
func (srv *Server) Start(address string) error {
	return srv.echo.Start(address)
}

// Shutdown shuts down the HTTP server
func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.echo.Shutdown(ctx)
}

// ServeHTTP implements the `http.Handler` interface which serves HTTP requests
func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.echo.ServeHTTP(w, r)
}

// Health endpoints for probes
func (srv *Server) Health(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

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

func (srv *Server) configureMiddleware() {
	srv.echo.Use(middleware.RequestID())

	srv.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: LogSkipper,
		Format: `{"time":"${time_rfc3339_nano}","level":"INFO","message":{"id":"${id}","remote_ip":"${remote_ip}",` + //nolint:lll
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` + //nolint:lll
			`"response_status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}",` +
			`"bytes_in":${bytes_in},"bytes_out":${bytes_out}}}` + "\n",
		Output: os.Stdout,
	}))
}
