package metrics

import (
	"context"

	opMetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/opio"
	txmgrMetrics "github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	"github.com/ethereum/go-ethereum/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-client/cmd/flags"
)

// Metrics
var (
	registry = opMetrics.NewRegistry()
	factory  = opMetrics.With(registry)

	// Driver
	DriverL1HeadHeightGauge     = factory.NewGauge(prometheus.GaugeOpts{Name: "driver/l1Head/height"})
	DriverL2HeadHeightGauge     = factory.NewGauge(prometheus.GaugeOpts{Name: "driver/l2Head/height"})
	DriverL1CurrentHeightGauge  = factory.NewGauge(prometheus.GaugeOpts{Name: "driver/l1Current/height"})
	DriverL2HeadIDGauge         = factory.NewGauge(prometheus.GaugeOpts{Name: "driver/l2Head/id"})
	DriverL2VerifiedHeightGauge = factory.NewGauge(prometheus.GaugeOpts{Name: "driver/l2Verified/id"})

	// Proposer
	ProposerProposeEpochCounter    = factory.NewCounter(prometheus.CounterOpts{Name: "proposer/epoch"})
	ProposerProposedTxListsCounter = factory.NewCounter(prometheus.CounterOpts{Name: "proposer/proposed/txLists"})
	ProposerProposedTxsCounter     = factory.NewCounter(prometheus.CounterOpts{Name: "proposer/proposed/txs"})

	// Prover
	ProverLatestVerifiedIDGauge      = factory.NewGauge(prometheus.GaugeOpts{Name: "prover/latestVerified/id"})
	ProverLatestProvenBlockIDGauge   = factory.NewGauge(prometheus.GaugeOpts{Name: "prover/latestProven/id"})
	ProverQueuedProofCounter         = factory.NewCounter(prometheus.CounterOpts{Name: "prover/proof/all/queued"})
	ProverReceivedProofCounter       = factory.NewCounter(prometheus.CounterOpts{Name: "prover/proof/all/received"})
	ProverSentProofCounter           = factory.NewCounter(prometheus.CounterOpts{Name: "prover/proof/all/sent"})
	ProverProofsAssigned             = factory.NewCounter(prometheus.CounterOpts{Name: "prover/proof/assigned"})
	ProverReceivedProposedBlockGauge = factory.NewGauge(prometheus.GaugeOpts{Name: "prover/proposed/received"})
	ProverReceivedProvenBlockGauge   = factory.NewGauge(prometheus.GaugeOpts{Name: "prover/proven/received"})
	ProverSubmissionAcceptedCounter  = factory.NewCounter(prometheus.CounterOpts{
		Name: "prover/proof/submission/accepted",
	})
	ProverSubmissionErrorCounter = factory.NewCounter(prometheus.CounterOpts{
		Name: "prover/proof/submission/error",
	})
	ProverSgxProofGeneratedCounter = factory.NewCounter(prometheus.CounterOpts{
		Name: "prover/proof/sgx/generated",
	})
	ProverSubmissionRevertedCounter = factory.NewCounter(prometheus.CounterOpts{
		Name: "prover/proof/submission/reverted",
	})

	// TxManager
	TxMgrMetrics = txmgrMetrics.MakeTxMetrics("client", factory)
)

// Serve starts the metrics server on the given address, will be closed when the given
// context is cancelled.
func Serve(ctx context.Context, c *cli.Context) error {
	if !c.Bool(flags.MetricsEnabled.Name) {
		return nil
	}

	log.Info(
		"Starting metrics server",
		"host", c.String(flags.MetricsAddr.Name),
		"port", c.Int(flags.MetricsPort.Name),
	)

	server, err := opMetrics.StartServer(
		registry,
		c.String(flags.MetricsAddr.Name),
		c.Int(flags.MetricsPort.Name),
	)
	if err != nil {
		return err
	}

	defer func() {
		if err := server.Stop(ctx); err != nil {
			log.Error("Failed to close metrics server", "error", err)
		}
	}()

	opio.BlockOnInterruptsContext(ctx)

	return nil
}
