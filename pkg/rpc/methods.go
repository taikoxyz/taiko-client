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
	errSyncing = errors.New("syncing")
	// syncProgressRecheckDelay is the time delay of rechecking the L2 execution engine's sync progress again,
	// if the previous check failed.
	syncProgressRecheckDelay = 12 * time.Second
	minTxGasLimit            = 21000
)

// ensureGenesisMatched fetches the L2 genesis block from TaikoL1 contract,
// and checks whether the fetched genesis is same to the node local genesis.
func (c *Client) ensureGenesisMatched(ctx context.Context) error {
	stateVars, err := c.GetProtocolStateVariables(nil)
	if err != nil {
		return err
	}

	// Fetch the genesis `BlockVerified` event.
	iter, err := c.TaikoL1.FilterBlockVerified(
		&bind.FilterOpts{Start: stateVars.GenesisHeight, End: &stateVars.GenesisHeight},
		[]*big.Int{common.Big0},
	)
	if err != nil {
		return err
	}

	// Fetch the node's genesis block.
	nodeGenesis, err := c.L2.HeaderByNumber(ctx, common.Big0)
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

// WaitTillL2Synced keeps waiting until the L2 execution engine is fully synced.
func (c *Client) WaitTillL2Synced(ctx context.Context) error {
	return backoff.Retry(
		func() error {
			if ctx.Err() != nil {
				return nil
			}
			progress, err := c.L2ExecutionEngineSyncProgress(ctx)
			if err != nil {
				log.Error("Fetch L2 execution engine sync progress error", "error", err)
				return err
			}

			if progress.SyncProgress != nil ||
				progress.CurrentBlockID == nil ||
				progress.HighestBlockID == nil ||
				progress.CurrentBlockID.Cmp(progress.HighestBlockID) < 0 {
				log.Info("L2 execution engine is syncing", "progress", progress)
				return errSyncing
			}

			return nil
		},
		backoff.NewConstantBackOff(syncProgressRecheckDelay),
	)
}

// LatestL2KnownL1Header fetches the L2 execution engine's latest known L1 header.
func (c *Client) LatestL2KnownL1Header(ctx context.Context) (*types.Header, error) {
	headL1Origin, err := c.L2.HeadL1Origin(ctx)

	if err != nil {
		switch err.Error() {
		case ethereum.NotFound.Error():
			return c.GetGenesisL1Header(ctx)
		default:
			return nil, err
		}
	}

	if headL1Origin == nil {
		return c.GetGenesisL1Header(ctx)
	}

	header, err := c.L1.HeaderByHash(ctx, headL1Origin.L1BlockHash)
	if err != nil {
		switch err.Error() {
		case ethereum.NotFound.Error():
			log.Warn("Latest L2 known L1 header not found, use genesis instead", "hash", headL1Origin.L1BlockHash)
			return c.GetGenesisL1Header(ctx)
		default:
			return nil, err
		}
	}

	log.Info("Latest L2 known L1 header", "height", header.Number, "hash", header.Hash())

	return header, nil
}

// GetGenesisL1Header fetches the L1 header that including L2 genesis block.
func (c *Client) GetGenesisL1Header(ctx context.Context) (*types.Header, error) {
	stateVars, err := c.GetProtocolStateVariables(nil)
	if err != nil {
		return nil, err
	}

	return c.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(stateVars.GenesisHeight))
}

// L2ParentByBlockId fetches the block header from L2 execution engine with the largest block id that
// smaller than the given `blockId`.
func (c *Client) L2ParentByBlockId(ctx context.Context, blockID *big.Int) (*types.Header, error) {
	parentBlockId := new(big.Int).Sub(blockID, common.Big1)

	log.Debug("Get parent block by block ID", "parentBlockId", parentBlockId)

	if parentBlockId.Cmp(common.Big0) == 0 {
		return c.L2.HeaderByNumber(ctx, common.Big0)
	}

	l1Origin, err := c.L2.L1OriginByID(ctx, parentBlockId)
	if err != nil {
		return nil, err
	}

	log.Debug("Parent block L1 origin", "l1Origin", l1Origin, "parentBlockId", parentBlockId)

	return c.L2.HeaderByHash(ctx, l1Origin.L2BlockHash)
}

// WaitL1Origin keeps waiting until the L1Origin with given block ID appears on the L2 execution engine.
func (c *Client) WaitL1Origin(ctx context.Context, blockID *big.Int) (*rawdb.L1Origin, error) {
	var (
		l1Origin *rawdb.L1Origin
		err      error
	)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	log.Debug("Start fetching L1Origin from L2 execution engine", "blockID", blockID)

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			l1Origin, err = c.L2.L1OriginByID(ctx, blockID)
			if err != nil {
				log.Warn("Failed to fetch L1Origin from L2 execution engine", "blockID", blockID, "error", err)
				continue
			}

			if l1Origin == nil {
				continue
			}

			return l1Origin, nil
		}
	}
}

// GetPoolContent fetches the transactions list from L2 execution engine's transactions pool with given
// upper limit.
func (c *Client) GetPoolContent(
	ctx context.Context,
	maxTransactionsPerBlock uint64,
	blockMaxGasLimit uint64,
	maxBytesPerTxList uint64,
	locals []common.Address,
) ([]types.Transactions, error) {
	var localsArg []string
	for _, local := range locals {
		localsArg = append(localsArg, local.Hex())
	}

	var result []types.Transactions
	err := c.L2RawRPC.CallContext(
		ctx,
		&result,
		"taiko_txPoolContent",
		maxTransactionsPerBlock,
		blockMaxGasLimit,
		maxBytesPerTxList,
		minTxGasLimit,
		localsArg,
	)

	return result, err
}

// L2AccountNonce fetches the nonce of the given L2 account at a specified height.
func (c *Client) L2AccountNonce(
	ctx context.Context,
	account common.Address,
	height *big.Int,
) (uint64, error) {
	var result hexutil.Uint64
	err := c.L2RawRPC.CallContext(ctx, &result, "eth_getTransactionCount", account, hexutil.EncodeBig(height))
	return uint64(result), err
}

// L2SyncProgress represents the sync progress of a L2 execution engine, `ethereum.SyncProgress` is used to check
// the sync progress of verified blocks, and block IDs are used to check the sync progress of pending blocks.
type L2SyncProgress struct {
	*ethereum.SyncProgress
	CurrentBlockID *big.Int
	HighestBlockID *big.Int
}

// L2ExecutionEngineSyncProgress fetches the sync progress of the given L2 execution engine.
func (c *Client) L2ExecutionEngineSyncProgress(ctx context.Context) (*L2SyncProgress, error) {
	var (
		progress = new(L2SyncProgress)
		err      error
	)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		progress.SyncProgress, err = c.L2.SyncProgress(ctx)
		return err
	})

	g.Go(func() error {
		stateVars, err := c.GetProtocolStateVariables(nil)
		if err != nil {
			return err
		}
		progress.HighestBlockID = new(big.Int).SetUint64(stateVars.NumBlocks - 1)
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
func (c *Client) GetProtocolStateVariables(opts *bind.CallOpts) (*bindings.TaikoDataStateVariables, error) {
	return GetProtocolStateVariables(c.TaikoL1, opts)
}

// GetStorageRoot returns a contract's storage root at the given height.
func (c *Client) GetStorageRoot(
	ctx context.Context,
	gethclient *gethclient.Client,
	contract common.Address,
	height *big.Int,
) (common.Hash, error) {
	proof, err := gethclient.GetProof(
		ctx,
		contract,
		[]string{"0x0000000000000000000000000000000000000000000000000000000000000000"},
		height,
	)
	if err != nil {
		return common.Hash{}, err
	}

	return proof.StorageHash, nil
}

// CheckL1Reorg checks whether the L1 chain has been reorged, if so, returns the l1Current cursor and L2 blockID
// that need to reset to.
func (c *Client) CheckL1Reorg(ctx context.Context, blockID *big.Int) (bool, *types.Header, *big.Int, error) {
	var (
		reorged          bool
		l1CurrentToReset *types.Header
		blockIDToReset   *big.Int
	)
	for {
		if blockID.Cmp(common.Big0) == 0 {
			stateVars, err := c.TaikoL1.GetStateVariables(nil)
			if err != nil {
				return false, nil, nil, err
			}

			if l1CurrentToReset, err = c.L1.HeaderByNumber(
				ctx,
				new(big.Int).SetUint64(stateVars.GenesisHeight),
			); err != nil {
				return false, nil, nil, err
			}

			blockIDToReset = blockID
			break
		}

		l1Origin, err := c.L2.L1OriginByID(ctx, blockID)
		if err != nil {
			return false, nil, nil, err
		}

		l1Header, err := c.L1.HeaderByNumber(ctx, l1Origin.L1BlockHeight)
		if err != nil {
			if errors.Is(err, ethereum.NotFound) {
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
		"Check L1 reorg",
		"reorged", reorged,
		"l1CurrentToResetNumber", l1CurrentToReset.Number,
		"l1CurrentToResetHash", l1CurrentToReset.Hash(),
		"blockIDToReset", blockIDToReset,
	)

	return reorged, l1CurrentToReset, blockIDToReset, nil
}
