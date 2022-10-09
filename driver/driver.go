package driver

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikochain/taiko-client/rpc"
	"github.com/taikochain/taiko-client/util"
	"github.com/urfave/cli/v2"
)

const (
	// Time to wait before the next try, when receiving subscription errors.
	RetryDelay         = 10 * time.Second
	MaxReorgDepth      = 500
	ReorgRollbackDepth = 20
)

// Driver keeps the L2 node's local block chain in sync with the TaikoL1
// contract.
type Driver struct {
	rpc             *rpc.Client
	l2ChainInserter *L2ChainInserter
	state           *State

	l1HeadCh   chan *types.Header
	l1HeadSub  event.Subscription
	syncNotify chan struct{}

	ctx      context.Context
	ctxClose context.CancelFunc
	wg       sync.WaitGroup
}

// Action returns the main function that the subcommand should run.
func Action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		cfg, err := NewConfigFromCliContext(ctx)
		if err != nil {
			return err
		}

		driver, err := New(context.Background(), cfg)
		if err != nil {
			return err
		}

		return util.RunSubcommand(driver)
	}
}

// New initializes a new driver instance based on the given configurations.
func New(ctx context.Context, cfg *Config) (*Driver, error) {
	l1HeadCh := make(chan *types.Header)

	rpc, err := rpc.NewClient(ctx, &rpc.ClientConfig{
		L1Endpoint:       cfg.L1Endpoint,
		L2Endpoint:       cfg.L2Endpoint,
		TaikoL1Address:   cfg.TaikoL1Address,
		TaikoL2Address:   cfg.TaikoL2Address,
		L2EngineEndpoint: cfg.L2EngineEndpoint,
		JwtSecret:        cfg.JwtSecret,
	})
	if err != nil {
		return nil, err
	}

	state, err := NewState(ctx, rpc)
	if err != nil {
		return nil, err
	}

	blockInserter, err := NewL2ChainInserter(
		ctx,
		rpc,
		state,
		cfg.ThrowawayBlocksBuilderPrivKey,
	)
	if err != nil {
		return nil, err
	}

	withCancelCtx, cancel := context.WithCancel(ctx)

	return &Driver{
		rpc:             rpc,
		l2ChainInserter: blockInserter,
		state:           state,
		l1HeadCh:        l1HeadCh,
		l1HeadSub:       state.SubL1HeadsFeed(l1HeadCh),
		syncNotify:      make(chan struct{}, 1),
		ctx:             withCancelCtx,
		ctxClose:        cancel,
		wg:              sync.WaitGroup{},
	}, nil
}

// Start starts the driver instance.
func (d *Driver) Start() error {
	d.wg.Add(1)
	go d.eventLoop()

	return nil
}

// Close closes the driver instance.
func (d *Driver) Close() {
	d.ctxClose()
	d.wg.Wait()
	if d.state != nil {
		d.state.Close()
	}
	if d.l1HeadSub != nil {
		d.l1HeadSub.Unsubscribe()
	}
}

// eventLoop starts the main loop of L2 node's chain driver.
func (d *Driver) eventLoop() {
	defer d.wg.Done()
	exponentialBackoff := backoff.NewExponentialBackOff()

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
		if err := backoff.Retry(d.doSync, exponentialBackoff); err != nil {
			log.Error("Sync L2 node block chain error", "error", err)
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
	l1Head := d.state.getL1Head()

	if err := d.l2ChainInserter.ProcessL1Blocks(
		d.ctx,
		l1Head,
	); err != nil {
		log.Error("Process new L1 blocks error", "error", err)
		if errors.Is(err, context.Canceled) {
			return nil
		}
		return err
	}

	return nil
}

// Name returns the application name.
func (d *Driver) Name() string {
	return "driver"
}
