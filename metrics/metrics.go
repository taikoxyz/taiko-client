package metrics

import (
	"context"
	"net"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/metrics/prometheus"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

// Metrics
var (
	// Driver
	DriverL1HeadHeightGauge     = metrics.NewRegisteredGauge("driver/l1Head/height", nil)
	DriverL2HeadHeightGauge     = metrics.NewRegisteredGauge("driver/l2Head/height", nil)
	DriverL1CurrentHeightGauge  = metrics.NewRegisteredGauge("driver/l1Current/height", nil)
	DriverL2HeadIDGauge         = metrics.NewRegisteredGauge("driver/l2Head/id", nil)
	DriverL2VerifiedHeightGauge = metrics.NewRegisteredGauge("driver/l2Verified/id", nil)

	// Proposer
	ProposerProposeEpochCounter    = metrics.NewRegisteredCounter("proposer/epoch", nil)
	ProposerProposedTxListsCounter = metrics.NewRegisteredCounter("proposer/proposed/txLists", nil)
	ProposerProposedTxsCounter     = metrics.NewRegisteredCounter("proposer/proposed/txs", nil)
	ProposerBaseFeeLenMetrics      = metrics.NewRegisteredGauge("proposer/proposed/baseFeeLen", nil)

	// Prover
	ProverPendingBlocksGauge          = metrics.NewRegisteredGauge("prover/pending/blocks", nil)
	ProverLatestVerifiedIDGauge       = metrics.NewRegisteredGauge("prover/latestVerified/id", nil)
	ProverLatestProvenBlockIDGauge    = metrics.NewRegisteredGauge("prover/latestProven/id", nil)
	ProverQueuedProofCounter          = metrics.NewRegisteredCounter("prover/proof/all/queued", nil)
	ProverQueuedValidProofCounter     = metrics.NewRegisteredCounter("prover/proof/valid/queued", nil)
	ProverQueuedInvalidProofCounter   = metrics.NewRegisteredCounter("prover/proof/invalid/queued", nil)
	ProverReceivedProofCounter        = metrics.NewRegisteredCounter("prover/proof/all/received", nil)
	ProverReceivedValidProofCounter   = metrics.NewRegisteredCounter("prover/proof/valid/received", nil)
	ProverReceivedInvalidProofCounter = metrics.NewRegisteredCounter("prover/proof/invalid/received", nil)
	ProverSentProofCounter            = metrics.NewRegisteredCounter("prover/proof/all/sent", nil)
	ProverSentValidProofCounter       = metrics.NewRegisteredCounter("prover/proof/valid/sent", nil)
	ProverSentInvalidProofCounter     = metrics.NewRegisteredCounter("prover/proof/invalid/sent", nil)
	ProverReceivedProposedBlockGauge  = metrics.NewRegisteredGauge("prover/proposed/received", nil)
)

// Serve starts the metrics server on the given address, will be closed when the given
// context is cancelled.
func Serve(ctx context.Context, c *cli.Context) error {
	if !c.Bool(flags.MetricsEnabled.Name) {
		return nil
	}

	address := net.JoinHostPort(
		c.String(flags.MetricsAddr.Name),
		strconv.Itoa(c.Int(flags.MetricsPort.Name)),
	)

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
