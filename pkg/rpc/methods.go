package rpc

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
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

// PoolContent represents a response body of a `txpool_content` RPC call.
type PoolContent map[common.Address]map[string]*types.Transaction

type TxLists []types.Transactions

// ToTxLists flattens all transactions in pool content into transactions lists,
// each list contains transactions from a single account sorted by nonce.
func (pc PoolContent) ToTxLists() TxLists {
	txLists := make([]types.Transactions, 0)

	for _, pendingTxs := range pc {
		var txsByNonce types.TxByNonce

		for _, pendingTx := range pendingTxs {
			txsByNonce = append(txsByNonce, pendingTx)
		}

		sort.Sort(txsByNonce)

		txLists = append(txLists, types.Transactions(txsByNonce))
	}

	return txLists
}

// Len returns the number of transactions inside the transactions lists.
func (t TxLists) Len() int {
	var length = 0
	for _, pendingTxs := range t {
		length += len(pendingTxs)
	}
	return length
}

// L2PoolContent fetches the transaction pool content from a L2 execution engine.
func (c *Client) L2PoolContent(ctx context.Context) (pending PoolContent, queued PoolContent, err error) {
	var res map[string]PoolContent
	if err := c.L2RawRPC.CallContext(ctx, &res, "txpool_content"); err != nil {
		return nil, nil, err
	}

	return res["pending"], res["queued"], nil
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

// IsProverWhitelisted checks whether the given prover is whitelisted.
func (c *Client) IsProverWhitelisted(prover common.Address) (bool, error) {
	whitelisted, err := c.TaikoL1.IsProverWhitelisted(nil, prover)
	if err != nil {
		if strings.Contains(err.Error(), "reverted") {
			return true, nil
		}

		return false, err
	}

	return whitelisted, nil
}

// IsProposerWhitelisted checks whether the given proposer is whitelisted.
func (c *Client) IsProposerWhitelisted(proposer common.Address) (bool, error) {
	whitelisted, err := c.TaikoL1.IsProposerWhitelisted(nil, proposer)
	if err != nil {
		if strings.Contains(err.Error(), "reverted") {
			return true, nil
		}

		return false, err
	}

	return whitelisted, nil
}

// GetProtocolConstants gets the protocol constants from TaikoL1 contract.
func (c *Client) GetProtocolConstants(opts *bind.CallOpts) (*bindings.ProtocolConstants, error) {
	return GetProtocolConstants(c.TaikoL1, opts)
}

// GetProtocolStateVariables gets the protocol states from TaikoL1 contract.
func (c *Client) GetProtocolStateVariables(opts *bind.CallOpts) (*bindings.ProtocolStateVariables, error) {
	return GetProtocolStateVariables(c.TaikoL1, opts)
}
