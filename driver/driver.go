package driver

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	chainSyncer "github.com/taikoxyz/taiko-client/driver/chain_syncer"
	"github.com/taikoxyz/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/urfave/cli/v2"
)

const (
	protocolStatusReportInterval     = 30 * time.Second
	exchangeTransitionConfigInterval = 1 * time.Minute
)

// Driver keeps the L2 execution engine's local block chain in sync with the TaikoL1
// contract.
type Driver struct {
	*Config
	rpc           *rpc.Client
	l2ChainSyncer *chainSyncer.L2ChainSyncer
	state         *state.State

	syncNotify chan struct{}

	maxNumBlocks uint64

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// InitFromCli initializes the given driver instance based on the command line flags.
func (d *Driver) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return d.InitFromConfig(ctx, cfg)
}

// InitFromConfig initializes the driver instance based on the given configurations.
func (d *Driver) InitFromConfig(ctx context.Context, cfg *Config) (err error) {
	d.syncNotify = make(chan struct{}, 1)
	d.ctx, d.cancel = context.WithCancel(ctx)
	d.Config = cfg

	if d.rpc, err = rpc.NewClient(d.ctx, cfg.ClientConfig); err != nil {
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

	if d.l2ChainSyncer, err = chainSyncer.New(
		d.rpc,
		d.state,
		cfg.P2PSyncVerifiedBlocks,
		cfg.P2PSyncTimeout,
		cfg.MaxExponent,
	); err != nil {
		return err
	}

	configs, err := d.rpc.TaikoL1.GetConfig(&bind.CallOpts{Context: d.ctx})
	if err != nil {
		log.Error("Failed to get protocol state variables", "error", err)
		return err
	}
	d.maxNumBlocks = configs.BlockMaxProposals

	return nil
}

// Start starts the driver instance.
func (d *Driver) Start() error {
	go d.eventLoop()
	go d.reportProtocolStatus()
	go d.exchangeTransitionConfigLoop()

	return nil
}

// Close closes the driver instance.
func (d *Driver) Close(_ context.Context) {
	if d.cancel != nil {
		d.cancel()
	}
	d.wg.Wait()
	d.l2ChainSyncer.Close()
	d.state.Close()
	d.rpc.Close()
}

// eventLoop starts the main loop of a L2 execution engine's driver.
func (d *Driver) eventLoop() {
	d.wg.Add(1)
	defer d.wg.Done()

	// reqSync requests performing a synchronising operation, won't block
	// if we are already synchronising.
	reqSync := func() {
		select {
		case d.syncNotify <- struct{}{}:
		default:
		}
	}

	ctx, cancel := context.WithCancel(d.ctx)
	defer cancel()

	// doSyncWithBackoff performs a synchronising operation with a backoff strategy.
	doSyncWithBackoff := func() {
		if err := d.l2ChainSyncer.Sync(ctx, d.state.GetL1Head()); err != nil {
			log.Error("Process new L1 blocks error", "error", err)
		}
	}

	// Call doSync() right away to catch up with the latest known L1 head.
	doSyncWithBackoff()

	l1HeadCh := make(chan *types.Header, 1024)
	l1HeadSub := d.state.SubL1HeadsFeed(l1HeadCh)
	defer l1HeadSub.Unsubscribe()

	for {
		select {
		case <-d.ctx.Done():
			return
		case <-d.syncNotify:
			doSyncWithBackoff()
		case <-l1HeadCh:
			reqSync()
		}
	}
}

// ChainSyncer returns the driver's chain syncer, this method
// should only be used for testing.
func (d *Driver) ChainSyncer() *chainSyncer.L2ChainSyncer {
	return d.l2ChainSyncer
}

// reportProtocolStatus reports some protocol status intervally.
func (d *Driver) reportProtocolStatus() {
	d.wg.Add(1)
	defer d.wg.Done()

	var ticker = time.NewTicker(protocolStatusReportInterval)
	defer ticker.Stop()

	subCtx, cancel := context.WithCancel(d.ctx)
	defer cancel()

	for {
		select {
		case <-d.ctx.Done():
			return
		case <-ticker.C:
			vars, err := d.rpc.GetProtocolStateVariables(&bind.CallOpts{Context: subCtx})
			if err != nil {
				log.Error("Failed to get protocol state variables", "error", err)
				continue
			}

			log.Info(
				"📖 Protocol status",
				"lastVerifiedBlockId", vars.B.LastVerifiedBlockId,
				"pendingBlocks", vars.B.NumBlocks-vars.B.LastVerifiedBlockId-1,
				"availableSlots", vars.B.LastVerifiedBlockId+d.maxNumBlocks-vars.B.NumBlocks,
			)
		}
	}
}

// exchangeTransitionConfigLoop keeps exchanging transition configs with the
// L2 execution engine.
func (d *Driver) exchangeTransitionConfigLoop() {
	d.wg.Add(1)
	defer d.wg.Done()

	ticker := time.NewTicker(exchangeTransitionConfigInterval)
	defer ticker.Stop()

	subCtx, cancel := context.WithCancel(d.ctx)
	defer cancel()

	for {
		select {
		case <-d.ctx.Done():
			return
		case <-ticker.C:
			func() {
				tc, err := d.rpc.L2Engine.ExchangeTransitionConfiguration(subCtx, &engine.TransitionConfigurationV1{
					TerminalTotalDifficulty: (*hexutil.Big)(common.Big0),
					TerminalBlockHash:       common.Hash{},
					TerminalBlockNumber:     0,
				})
				if err != nil {
					log.Error("Failed to exchange Transition Configuration", "error", err)
				} else {
					log.Debug("Exchanged transition config", "transitionConfig", tc)
				}
			}()
		}
	}
}

// Name returns the application name.
func (d *Driver) Name() string {
	return "driver"
}
