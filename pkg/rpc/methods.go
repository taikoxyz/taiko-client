package rpc

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/cenkalti/backoff/v4"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"golang.org/x/sync/errgroup"
)

var (
	// errSyncing is returned when the L2 execution engine is syncing.
	errSyncing        = errors.New("syncing")
	errEmptyTiersList = errors.New("empty proof tiers list in protocol")
	// syncProgressRecheckDelay is the time delay of rechecking the L2 execution engine's sync progress again,
	// if the previous check failed.
	syncProgressRecheckDelay       = 12 * time.Second
	waitL1OriginPollingInterval    = 3 * time.Second
	defaultWaitL1OriginTimeout     = 3 * time.Minute
	defaultMaxTransactionsPerBlock = uint64(150)
)

// ensureGenesisMatched fetches the L2 genesis block from TaikoL1 contract,
// and checks whether the fetched genesis is same to the node local genesis.
func (c *Client) ensureGenesisMatched(ctx context.Context) error {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	stateVars, err := c.GetProtocolStateVariables(&bind.CallOpts{Context: ctxWithTimeout})
	if err != nil {
		return err
	}

	// Fetch the genesis `BlockVerified` event.
	iter, err := c.TaikoL1.FilterBlockVerified(
		&bind.FilterOpts{Start: stateVars.A.GenesisHeight, End: &stateVars.A.GenesisHeight, Context: ctxWithTimeout},
		[]*big.Int{common.Big0},
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	// Fetch the node's genesis block.
	nodeGenesis, err := c.L2.HeaderByNumber(ctxWithTimeout, common.Big0)
	if err != nil {
		return err
	}

	for iter.Next() {
		l2GenesisHash := iter.Event.BlockHash

		log.Debug("Genesis hash", "node", nodeGenesis.Hash(), "TaikoL1", common.BytesToHash(l2GenesisHash[:]))

		// Node's genesis header and TaikoL1 contract's genesis header must match.
		if common.BytesToHash(l2GenesisHash[:]) != nodeGenesis.Hash() {
			return fmt.Errorf(
				"genesis header hash mismatch, node: %s, TaikoL1 contract: %s",
				nodeGenesis.Hash(),
				common.BytesToHash(l2GenesisHash[:]),
			)
		} else {
			return nil
		}
	}

	log.Warn("Genesis block not found in TaikoL1")

	return nil
}

// WaitTillL2ExecutionEngineSynced keeps waiting until the L2 execution engine is fully synced.
func (c *Client) WaitTillL2ExecutionEngineSynced(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return backoff.Retry(
		func() error {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			progress, err := c.L2ExecutionEngineSyncProgress(ctx)
			if err != nil {
				log.Error("Fetch L2 execution engine sync progress error", "error", err)
				return err
			}

			if progress.isSyncing() {
				log.Info("L2 execution engine is syncing", "CurrentBlockID", progress.CurrentBlockID,
					"HighestBlockID", progress.HighestBlockID, "progress", progress.SyncProgress)
				return errSyncing
			}

			return nil
		},
		backoff.WithMaxRetries(backoff.NewConstantBackOff(syncProgressRecheckDelay), 10),
	)
}

// LatestL2KnownL1Header fetches the L2 execution engine's latest known L1 header.
func (c *Client) LatestL2KnownL1Header(ctx context.Context) (*types.Header, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	headL1Origin, err := c.L2.HeadL1Origin(ctxWithTimeout)
	if err != nil {
		switch err.Error() {
		case ethereum.NotFound.Error():
			return c.GetGenesisL1Header(ctxWithTimeout)
		default:
			return nil, err
		}
	}

	if headL1Origin == nil {
		return c.GetGenesisL1Header(ctxWithTimeout)
	}

	header, err := c.L1.HeaderByHash(ctxWithTimeout, headL1Origin.L1BlockHash)
	if err != nil {
		switch err.Error() {
		case ethereum.NotFound.Error():
			log.Warn("Latest L2 known L1 header not found, use genesis instead", "hash", headL1Origin.L1BlockHash)
			return c.GetGenesisL1Header(ctxWithTimeout)
		default:
			return nil, err
		}
	}

	log.Info("Latest L2 known L1 header", "height", header.Number, "hash", header.Hash())

	return header, nil
}

// GetGenesisL1Header fetches the L1 header that including L2 genesis block.
func (c *Client) GetGenesisL1Header(ctx context.Context) (*types.Header, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	stateVars, err := c.GetProtocolStateVariables(&bind.CallOpts{Context: ctxWithTimeout})
	if err != nil {
		return nil, err
	}

	return c.L1.HeaderByNumber(ctxWithTimeout, new(big.Int).SetUint64(stateVars.A.GenesisHeight))
}

// L2ParentByBlockId fetches the block header from L2 execution engine with the largest block id that
// smaller than the given `blockId`.
func (c *Client) L2ParentByBlockId(ctx context.Context, blockID *big.Int) (*types.Header, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	parentBlockId := new(big.Int).Sub(blockID, common.Big1)

	log.Debug("Get parent block by block ID", "parentBlockId", parentBlockId)

	if parentBlockId.Cmp(common.Big0) == 0 {
		return c.L2.HeaderByNumber(ctxWithTimeout, common.Big0)
	}

	l1Origin, err := c.L2.L1OriginByID(ctxWithTimeout, parentBlockId)
	if err != nil {
		return nil, err
	}

	log.Debug("Parent block L1 origin", "l1Origin", l1Origin, "parentBlockId", parentBlockId)

	return c.L2.HeaderByHash(ctxWithTimeout, l1Origin.L2BlockHash)
}

// WaitL1Origin keeps waiting until the L1Origin with given block ID appears on the L2 execution engine.
func (c *Client) WaitL1Origin(ctx context.Context, blockID *big.Int) (*rawdb.L1Origin, error) {
	var (
		l1Origin *rawdb.L1Origin
		err      error
	)

	ticker := time.NewTicker(waitL1OriginPollingInterval)
	defer ticker.Stop()

	var (
		ctxWithTimeout = ctx
		cancel         context.CancelFunc
	)
	if _, ok := ctx.Deadline(); !ok {
		ctxWithTimeout, cancel = context.WithTimeout(ctx, defaultWaitL1OriginTimeout)
		defer cancel()
	}

	log.Debug("Start fetching L1Origin from L2 execution engine", "blockID", blockID)
	for ; true; <-ticker.C {
		if ctxWithTimeout.Err() != nil {
			return nil, ctxWithTimeout.Err()
		}

		l1Origin, err = c.L2.L1OriginByID(ctxWithTimeout, blockID)
		if err != nil {
			log.Debug("L1Origin from L2 execution engine not found, keep retrying", "blockID", blockID, "error", err)
			continue
		}

		if l1Origin == nil {
			continue
		}

		return l1Origin, nil
	}

	return nil, fmt.Errorf("failed to fetch L1Origin from L2 execution engine, blockID: %d", blockID)
}

// GetPoolContent fetches the transactions list from L2 execution engine's transactions pool with given
// upper limit.
func (c *Client) GetPoolContent(
	ctx context.Context,
	beneficiary common.Address,
	baseFee *big.Int,
	blockMaxGasLimit uint32,
	maxBytesPerTxList uint64,
	locals []common.Address,
	maxTransactionsLists uint64,
) ([]types.Transactions, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	var localsArg []string
	for _, local := range locals {
		localsArg = append(localsArg, local.Hex())
	}

	var result []types.Transactions
	err := c.L2RawRPC.CallContext(
		ctxWithTimeout,
		&result,
		"taiko_txPoolContent",
		beneficiary,
		baseFee,
		defaultMaxTransactionsPerBlock,
		blockMaxGasLimit,
		maxBytesPerTxList,
		localsArg,
		maxTransactionsLists,
	)

	return result, err
}

// L2AccountNonce fetches the nonce of the given L2 account at a specified height.
func (c *Client) L2AccountNonce(
	ctx context.Context,
	account common.Address,
	height *big.Int,
) (uint64, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	var result hexutil.Uint64
	err := c.L2RawRPC.CallContext(ctxWithTimeout, &result, "eth_getTransactionCount", account, hexutil.EncodeBig(height))
	return uint64(result), err
}

// L2SyncProgress represents the sync progress of a L2 execution engine, `ethereum.SyncProgress` is used to check
// the sync progress of verified blocks, and block IDs are used to check the sync progress of pending blocks.
type L2SyncProgress struct {
	*ethereum.SyncProgress
	CurrentBlockID *big.Int
	HighestBlockID *big.Int
}

// isSyncing returns true if the L2 execution engine is syncing with L1.
func (p *L2SyncProgress) isSyncing() bool {
	return p.SyncProgress != nil ||
		p.CurrentBlockID == nil ||
		p.HighestBlockID == nil ||
		p.CurrentBlockID.Cmp(p.HighestBlockID) < 0
}

// L2ExecutionEngineSyncProgress fetches the sync progress of the given L2 execution engine.
func (c *Client) L2ExecutionEngineSyncProgress(ctx context.Context) (*L2SyncProgress, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	var (
		progress = new(L2SyncProgress)
		err      error
	)
	g, ctx := errgroup.WithContext(ctxWithTimeout)

	g.Go(func() error {
		progress.SyncProgress, err = c.L2.SyncProgress(ctx)
		return err
	})
	g.Go(func() error {
		stateVars, err := c.GetProtocolStateVariables(&bind.CallOpts{Context: ctx})
		if err != nil {
			return err
		}
		progress.HighestBlockID = new(big.Int).SetUint64(stateVars.B.NumBlocks - 1)
		return nil
	})
	g.Go(func() error {
		headL1Origin, err := c.L2.HeadL1Origin(ctx)
		if err != nil {
			switch err.Error() {
			case ethereum.NotFound.Error():
				// There is only genesis block in the L2 execution engine, or it has not started
				// syncing the pending blocks yet.
				progress.CurrentBlockID = common.Big0
				return nil
			default:
				return err
			}
		}
		progress.CurrentBlockID = headL1Origin.BlockID
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return progress, nil
}

// GetProtocolStateVariables gets the protocol states from TaikoL1 contract.
func (c *Client) GetProtocolStateVariables(opts *bind.CallOpts) (*struct {
	A bindings.TaikoDataSlotA
	B bindings.TaikoDataSlotB
}, error) {
	var (
		ctxWithTimeout context.Context
		cancel         context.CancelFunc
	)
	if opts != nil && opts.Context != nil {
		if _, ok := opts.Context.Deadline(); !ok {
			ctxWithTimeout, cancel = context.WithTimeout(opts.Context, defaultWaitReceiptTimeout)
			defer cancel()
			opts.Context = ctxWithTimeout
		}
	} else {
		ctxWithTimeout, cancel = context.WithTimeout(context.Background(), defaultWaitReceiptTimeout)
		defer cancel()
		opts = &bind.CallOpts{Context: ctxWithTimeout}
	}

	return GetProtocolStateVariables(c.TaikoL1, opts)
}

// GetStorageRoot returns a contract's storage root at the given height.
func (c *Client) GetStorageRoot(
	ctx context.Context,
	gethclient *gethclient.Client,
	contract common.Address,
	height *big.Int,
) (common.Hash, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	proof, err := gethclient.GetProof(
		ctxWithTimeout,
		contract,
		[]string{"0x0000000000000000000000000000000000000000000000000000000000000000"},
		height,
	)
	if err != nil {
		return common.Hash{}, err
	}

	return proof.StorageHash, nil
}

// CheckL1ReorgFromL2EE checks whether the L1 chain has been reorged from the L1Origin records in L2 EE,
// if so, returns the l1Current cursor and L2 blockID that need to reset to.
func (c *Client) CheckL1ReorgFromL2EE(ctx context.Context, blockID *big.Int) (bool, *types.Header, *big.Int, error) {
	var (
		reorged          bool
		l1CurrentToReset *types.Header
		blockIDToReset   *big.Int
	)
	for {
		ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
		defer cancel()

		if blockID.Cmp(common.Big0) == 0 {
			stateVars, err := c.TaikoL1.GetStateVariables(&bind.CallOpts{Context: ctxWithTimeout})
			if err != nil {
				return false, nil, nil, err
			}

			if l1CurrentToReset, err = c.L1.HeaderByNumber(
				ctxWithTimeout,
				new(big.Int).SetUint64(stateVars.A.GenesisHeight),
			); err != nil {
				return false, nil, nil, err
			}

			blockIDToReset = blockID
			break
		}

		l1Origin, err := c.L2.L1OriginByID(ctxWithTimeout, blockID)
		if err != nil {
			if err.Error() == ethereum.NotFound.Error() {
				log.Info("L1Origin not found", "blockID", blockID)

				// If the L2 EE is just synced through P2P, there is a chance that the EE do not have
				// the chain head L1Origin information recorded.
				justSyncedByP2P, err := c.IsJustSyncedByP2P(ctxWithTimeout)
				if err != nil {
					return false,
						nil,
						nil,
						fmt.Errorf("failed to check whether the L2 execution engine has just finished a P2P sync: %w", err)
				}

				log.Info(
					"Check whether the L2 execution engine has just finished a P2P sync",
					"justSyncedByP2P",
					justSyncedByP2P,
				)

				if justSyncedByP2P {
					return false, nil, nil, nil
				}

				log.Info("Reorg detected due to L1Origin not found", "blockID", blockID)
				reorged = true
				blockID = new(big.Int).Sub(blockID, common.Big1)
				continue
			}
			return false, nil, nil, err
		}

		l1Header, err := c.L1.HeaderByNumber(ctxWithTimeout, l1Origin.L1BlockHeight)
		if err != nil {
			if err.Error() == ethereum.NotFound.Error() {
				continue
			}
			return false, nil, nil, fmt.Errorf("failed to fetch L1 header (%d): %w", l1Origin.L1BlockHeight, err)
		}

		if l1Header.Hash() != l1Origin.L1BlockHash {
			log.Info(
				"Reorg detected",
				"blockID", blockID,
				"l1Height", l1Origin.L1BlockHeight,
				"l1HashOld", l1Origin.L1BlockHash,
				"l1HashNew", l1Header.Hash(),
			)
			reorged = true
			blockID = new(big.Int).Sub(blockID, common.Big1)
			continue
		}

		l1CurrentToReset = l1Header
		blockIDToReset = l1Origin.BlockID
		break
	}

	log.Debug(
		"Check L1 reorg from L2 EE",
		"reorged", reorged,
		"l1CurrentToResetNumber", l1CurrentToReset.Number,
		"l1CurrentToResetHash", l1CurrentToReset.Hash(),
		"blockIDToReset", blockIDToReset,
	)

	return reorged, l1CurrentToReset, blockIDToReset, nil
}

// CheckL1ReorgFromL1Cursor checks whether the L1 chain has been reorged from the given l1Current cursor,
// if so, returns the l1Current cursor that need to reset to.
func (c *Client) CheckL1ReorgFromL1Cursor(
	ctx context.Context,
	l1Current *types.Header,
	genesisHeightL1 uint64,
) (bool, *types.Header, *big.Int, error) {
	var (
		reorged          bool
		l1CurrentToReset *types.Header
	)
	for {
		ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
		defer cancel()

		if l1Current.Number.Uint64() <= genesisHeightL1 {
			newL1Current, err := c.L1.HeaderByNumber(ctxWithTimeout, new(big.Int).SetUint64(genesisHeightL1))
			if err != nil {
				return false, nil, nil, err
			}

			l1CurrentToReset = newL1Current
			break
		}

		l1Header, err := c.L1.BlockByNumber(ctxWithTimeout, l1Current.Number)
		if err != nil {
			if err.Error() == ethereum.NotFound.Error() {
				continue
			}

			return false, nil, nil, err
		}

		if l1Header.Hash() != l1Current.Hash() {
			log.Info(
				"Reorg detected",
				"l1Height", l1Current.Number,
				"l1HashOld", l1Current.Hash(),
				"l1HashNew", l1Header.Hash(),
			)
			reorged = true
			if l1Current, err = c.L1.HeaderByHash(ctxWithTimeout, l1Current.ParentHash); err != nil {
				return false, nil, nil, err
			}
			continue
		}

		l1CurrentToReset = l1Current
		break
	}

	log.Debug(
		"Check L1 reorg from l1Current cursor",
		"reorged", reorged,
		"l1CurrentToResetNumber", l1CurrentToReset.Number,
		"l1CurrentToResetHash", l1CurrentToReset.Hash(),
	)

	return reorged, l1CurrentToReset, nil, nil
}

// IsJustSyncedByP2P checks whether the given L2 execution engine has just finished a P2P
// sync.
func (c *Client) IsJustSyncedByP2P(ctx context.Context) (bool, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	l2Head, err := c.L2.HeaderByNumber(ctxWithTimeout, nil)
	if err != nil {
		return false, err
	}

	if _, err = c.L2.L1OriginByID(ctxWithTimeout, l2Head.Number); err != nil {
		if err.Error() == ethereum.NotFound.Error() {
			return true, nil
		}

		return false, err
	}

	return false, nil
}

// TierProviderTierWithID wraps protocol ITierProviderTier struct with an ID.
type TierProviderTierWithID struct {
	ID uint16
	bindings.ITierProviderTier
}

// GetTiers fetches all protocol supported tiers.
func (c *Client) GetTiers(ctx context.Context) ([]*TierProviderTierWithID, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	ids, err := c.TaikoL1.GetTierIds(&bind.CallOpts{Context: ctxWithTimeout})
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errEmptyTiersList
	}

	var tiers []*TierProviderTierWithID
	for _, id := range ids {
		tier, err := c.TaikoL1.GetTier(&bind.CallOpts{Context: ctxWithTimeout}, id)
		if err != nil {
			return nil, err
		}
		tiers = append(tiers, &TierProviderTierWithID{ID: id, ITierProviderTier: tier})
	}

	return tiers, nil
}
