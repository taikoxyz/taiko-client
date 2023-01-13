package rpc

import (
	"context"
	"fmt"
	"math/big"
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

// Len returns the number of transactions in the PoolContent.
func (pc PoolContent) Len() int {
	len := 0
	for _, pendingTxs := range pc {
		for range pendingTxs {
			len += 1
		}
	}

	return len
}

// ToTxsByPriceAndNonce creates a transaction set that can retrieve price sorted transactions in a nonce-honouring way.
func (pc PoolContent) ToTxsByPriceAndNonce(
	chainID *big.Int,
	localAddresses []common.Address,
) (
	locals *types.TransactionsByPriceAndNonce,
	remotes *types.TransactionsByPriceAndNonce,
) {
	var (
		localTxs  = map[common.Address]types.Transactions{}
		remoteTxs = map[common.Address]types.Transactions{}
	)

	for address, txsWithNonce := range pc {
	out:
		for _, tx := range txsWithNonce {
			for _, localAddress := range localAddresses {
				if address == localAddress {
					localTxs[address] = append(localTxs[address], tx)
					continue out
				}
			}
			remoteTxs[address] = append(remoteTxs[address], tx)
		}
	}

	return types.NewTransactionsByPriceAndNonce(types.LatestSignerForChainID(chainID), localTxs, nil),
		types.NewTransactionsByPriceAndNonce(types.LatestSignerForChainID(chainID), remoteTxs, nil)
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

// GetProtocolStateVariables gets the protocol states from TaikoL1 contract.
func (c *Client) GetProtocolStateVariables(opts *bind.CallOpts) (*bindings.ProtocolStateVariables, error) {
	return GetProtocolStateVariables(c.TaikoL1, opts)
}
