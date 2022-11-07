package metrics

import (
	"context"
	"net"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/metrics/prometheus"
)

func Setup(ctx context.Context, hostname string, port int) error {
	address := net.JoinHostPort(hostname, strconv.Itoa(port))
	server := &http.Server{
		Addr:    address,
		Handler: prometheus.Handler(metrics.DefaultRegistry),
	}

	go func() {
		<-ctx.Done()
		if err := server.Close(); err != nil {
			log.Error("Failed to close metrics server", "error", err)
		}
	}()

	log.Info("Starting metrics server", "address", address)

	return server.ListenAndServe()
}
