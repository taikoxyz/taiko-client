package driver

import (
	"context"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	chainSyncer "github.com/taikoxyz/taiko-client/driver/chain_syncer"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/urfave/cli/v2"
)

const (
	protocolStatusReportInterval     = 30 * time.Second
	exchangeTransitionConfigTimeout  = 30 * time.Second
	exchangeTransitionConfigInterval = 1 * time.Minute
)

// Driver keeps the L2 execution engine's local block chain in sync with the TaikoL1
// contract.
type Driver struct {
	rpc           *rpc.Client
	l2ChainSyncer *chainSyncer.L2ChainSyncer
	state         *state.State

	l1HeadCh   chan *types.Header
	l1HeadSub  event.Subscription
	syncNotify chan struct{}

	backOffRetryInterval time.Duration
	ctx                  context.Context
	wg                   sync.WaitGroup
}

// New initializes the given driver instance based on the command line flags.
func (d *Driver) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return InitFromConfig(ctx, d, cfg)
}

// InitFromConfig initializes the driver instance based on the given configurations.
func InitFromConfig(ctx context.Context, d *Driver, cfg *Config) (err error) {
	d.l1HeadCh = make(chan *types.Header, 1024)
	d.wg = sync.WaitGroup{}
	d.syncNotify = make(chan struct{}, 1)
	d.ctx = ctx
	d.backOffRetryInterval = cfg.BackOffRetryInterval

	if d.rpc, err = rpc.NewClient(d.ctx, &rpc.ClientConfig{
		L1Endpoint:       cfg.L1Endpoint,
		L2Endpoint:       cfg.L2Endpoint,
		L2CheckPoint:     cfg.L2CheckPoint,
		TaikoL1Address:   cfg.TaikoL1Address,
		TaikoL2Address:   cfg.TaikoL2Address,
		L2EngineEndpoint: cfg.L2EngineEndpoint,
		JwtSecret:        cfg.JwtSecret,
		RetryInterval:    cfg.BackOffRetryInterval,
	}); err != nil {
		return err
	}

	if d.state, err = state.New(d.ctx, d.rpc); err != nil {
		return err
	}

	peers, err := d.rpc.L2.PeerCount(d.ctx)
	if err != nil {
		return err
	}

	if cfg.P2PSyncVerifiedBlocks && peers == 0 {
		log.Warn("P2P syncing verified blocks enabled, but no connected peer found in L2 execution engine")
	}

	signalServiceAddress, err := d.rpc.TaikoL1.Resolve0(nil, rpc.StringToBytes32("signal_service"), false)
	if err != nil {
		return err
	}

	if d.l2ChainSyncer, err = chainSyncer.New(
		d.ctx,
		d.rpc,
		d.state,
		cfg.P2PSyncVerifiedBlocks,
		cfg.P2PSyncTimeout,
		signalServiceAddress,
	); err != nil {
		return err
	}

	d.l1HeadSub = d.state.SubL1HeadsFeed(d.l1HeadCh)

	return nil
}

// Start starts the driver instance.
func (d *Driver) Start() error {
	d.wg.Add(3)
	go d.eventLoop()
	go d.reportProtocolStatus()
	go d.exchangeTransitionConfigLoop()

	return nil
}

// Close closes the driver instance.
func (d *Driver) Close() {
	d.state.Close()
	d.wg.Wait()
}

// eventLoop starts the main loop of a L2 execution engine's driver.
func (d *Driver) eventLoop() {
	defer d.wg.Done()

	// reqSync requests performing a synchronising operation, won't block
	// if we are already synchronising.
	reqSync := func() {
		select {
		case d.syncNotify <- struct{}{}:
		default:
		}
	}

	// doSyncWithBackoff performs a synchronising operation with a backoff strategy.
	doSyncWithBackoff := func() {
		if err := backoff.Retry(d.doSync, backoff.NewConstantBackOff(d.backOffRetryInterval)); err != nil {
			log.Error("Sync L2 execution engine's block chain error", "error", err)
		}
	}

	// Call doSync() right away to catch up with the latest known L1 head.
	doSyncWithBackoff()

	for {
		select {
		case <-d.ctx.Done():
			return
		case <-d.syncNotify:
			doSyncWithBackoff()
		case <-d.l1HeadCh:
			reqSync()
		}
	}
}

// doSync fetches all `BlockProposed` events emitted from local
// L1 sync cursor to the L1 head, and then applies all corresponding
// L2 blocks into node's local block chain.
func (d *Driver) doSync() error {
	// Check whether the application is closing.
	if d.ctx.Err() != nil {
		log.Warn("Driver context error", "error", d.ctx.Err())
		return nil
	}

	l1Head := d.state.GetL1Head()

	if err := d.l2ChainSyncer.Sync(l1Head); err != nil {
		log.Error("Process new L1 blocks error", "error", err)
		return err
	}

	return nil
}

// ChainSyncer returns the driver's chain syncer.
func (d *Driver) ChainSyncer() *chainSyncer.L2ChainSyncer {
	return d.l2ChainSyncer
}

// reportProtocolStatus reports some protocol status intervally.
func (d *Driver) reportProtocolStatus() {
	ticker := time.NewTicker(protocolStatusReportInterval)
	defer func() {
		ticker.Stop()
		d.wg.Done()
	}()

	var maxNumBlocks uint64
	if err := backoff.Retry(
		func() error {
			configs, err := d.rpc.TaikoL1.GetConfig(nil)
			if err != nil {
				return err
			}

			maxNumBlocks = configs.MaxNumProposedBlocks.Uint64()
			return nil
		},
		backoff.NewConstantBackOff(d.backOffRetryInterval),
	); err != nil {
		log.Error("Failed to get protocol state variables", "error", err)
		return
	}

	for {
		select {
		case <-d.ctx.Done():
			return
		case <-ticker.C:
			vars, err := d.rpc.GetProtocolStateVariables(nil)
			if err != nil {
				log.Error("Failed to get protocol state variables", "error", err)
				continue
			}

			log.Info(
				"📖 Protocol status",
				"lastVerifiedBlockId", vars.LastVerifiedBlockId,
				"pendingBlocks", vars.NumBlocks-vars.LastVerifiedBlockId-1,
				"availableSlots", vars.LastVerifiedBlockId+maxNumBlocks-vars.NumBlocks,
			)
		}
	}
}

// exchangeTransitionConfigLoop keeps exchanging transition configs with the
// L2 execution engine.
func (d *Driver) exchangeTransitionConfigLoop() {
	ticker := time.NewTicker(exchangeTransitionConfigInterval)
	defer func() {
		ticker.Stop()
		d.wg.Done()
	}()

	for {
		select {
		case <-d.ctx.Done():
			return
		case <-ticker.C:
			func() {
				ctx, cancel := context.WithTimeout(d.ctx, exchangeTransitionConfigTimeout)
				defer cancel()

				tc, err := d.rpc.L2Engine.ExchangeTransitionConfiguration(ctx, &engine.TransitionConfigurationV1{
					TerminalTotalDifficulty: (*hexutil.Big)(common.Big0),
					TerminalBlockHash:       common.Hash{},
					TerminalBlockNumber:     0,
				})
				if err != nil {
					log.Error("Failed to exchange Transition Configuration", "error", err)
					return
				}

				log.Debug(
					"Exchanged transition config",
					"transitionConfig", tc,
				)
			}()
		}
	}
}

// Name returns the application name.
func (d *Driver) Name() string {
	return "driver"
}
