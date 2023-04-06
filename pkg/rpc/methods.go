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
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"golang.org/x/sync/errgroup"
)

var (
	// errSyncing is returned when the L2 execution engine is syncing.
	errSyncing = errors.New("syncing")
	// syncProgressRecheckDelay is the time delay of rechecking the L2 execution engine's sync progress again,
	// if the previous check failed.
	syncProgressRecheckDelay = 3 * time.Second
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

	return fmt.Errorf("genesis block not found in TaikoL1")
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

	for parentBlockId.Cmp(common.Big0) > 0 {
		l1Origin, err := c.L2.L1OriginByID(ctx, parentBlockId)
		if err != nil {
			return nil, err
		}

		log.Debug("Parent block L1 origin", "l1Origin", l1Origin, "parentBlockId", parentBlockId)

		if l1Origin.Throwaway {
			parentBlockId = new(big.Int).Sub(parentBlockId, common.Big1)
			continue
		}

		return c.L2.HeaderByHash(ctx, l1Origin.L2BlockHash)
	}

	return c.L2.HeaderByNumber(ctx, common.Big0)
}

// WaitL1Origin keeps waiting until the L1Origin with given block ID appears on the L2 execution engine.
func (c *Client) WaitL1Origin(ctx context.Context, blockID *big.Int) (*rawdb.L1Origin, error) {
	var (
		l1Origin *rawdb.L1Origin
		err      error
	)

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	log.Debug("Start fetching L1Origin from L2 execution engine", "blockID", blockID)

	if _, ok := ctx.Deadline(); !ok {
		log.Debug("No deadline set, set a default deadline")
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 45*time.Second)
		defer cancel()
		ctx = ctxWithTimeout
	}

	for {
		select {
		case <-time.After(45 * time.Second):
			return nil, fmt.Errorf("timeout waiting for L1Origin with block ID %s", blockID)
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			l1Origin, err = c.L2.L1OriginByID(ctx, blockID)
			if err != nil {
				log.Warn("Failed to fetch L1Origin from L2 execution engine", "blockID", blockID, "error", err)
				continue
			}

			if l1Origin == nil {
				log.Debug("L1Origin not found", "blockID", blockID)
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
	maxTransactionsPerBlock *big.Int,
	blockMaxGasLimit *big.Int,
	maxBytesPerTxList *big.Int,
	minTxGasLimit *big.Int,
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
		maxTransactionsPerBlock.Uint64(),
		blockMaxGasLimit.Uint64(),
		maxBytesPerTxList.Uint64(),
		minTxGasLimit.Uint64(),
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
		progress.HighestBlockID = new(big.Int).SetUint64(stateVars.NextBlockId - 1)
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
func (c *Client) GetProtocolStateVariables(opts *bind.CallOpts) (*bindings.LibUtilsStateVariables, error) {
	return GetProtocolStateVariables(c.TaikoL1, opts)
}
