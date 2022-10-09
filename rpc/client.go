package rpc

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/taikochain/taiko-client/bindings"
)

// Client contains all L1/L2 RPC clients that a driver needs.
type Client struct {
	// Geth ethclient clients
	L1 *ethclient.Client
	L2 *ethclient.Client
	// Geth raw RPC clients
	L2RawRPC *rpc.Client
	// geth engine RPC clients
	L2Engine *EngineRPCClient
	// Protocol contracts clients
	TaikoL1 *bindings.TaikoL1Client
	TaikoL2 *bindings.V1TaikoL2Client
}

// ClientConfig contains all configs used by initializing an
// RPC client. If L2EngineEndpoint or JwtSecret not provided, the L2Engine client
// wont be initialized.
type ClientConfig struct {
	L1Endpoint       string
	L2Endpoint       string
	TaikoL1Address   common.Address
	TaikoL2Address   common.Address
	L2EngineEndpoint string
	JwtSecret        string
}

// NewClient initializes all RPC clients to run a Taiko client.
func NewClient(ctx context.Context, cfg *ClientConfig) (*Client, error) {
	l1RPC, err := DialClientWithBackoff(ctx, cfg.L1Endpoint)
	if err != nil {
		return nil, err
	}

	taikoL1, err := bindings.NewTaikoL1Client(cfg.TaikoL1Address, l1RPC)
	if err != nil {
		return nil, err
	}

	l2RPC, err := DialClientWithBackoff(ctx, cfg.L2Endpoint)
	if err != nil {
		return nil, err
	}

	taikoL2, err := bindings.NewV1TaikoL2Client(cfg.TaikoL2Address, l2RPC)
	if err != nil {
		return nil, err
	}

	l2RawRPC, err := rpc.Dial(cfg.L2Endpoint)
	if err != nil {
		return nil, err
	}

	var l2AuthRPC *EngineRPCClient
	if len(cfg.L2EngineEndpoint) != 0 && len(cfg.JwtSecret) != 0 {
		l2AuthRPC, err = DialEngineClientWithBackoff(
			ctx,
			cfg.L2EngineEndpoint,
			cfg.JwtSecret,
		)
		if err != nil {
			return nil, err
		}
	}

	client := &Client{
		L1:       l1RPC,
		L2:       l2RPC,
		L2RawRPC: l2RawRPC,
		L2Engine: l2AuthRPC,
		TaikoL1:  taikoL1,
		TaikoL2:  taikoL2,
	}

	if err := client.ensureGenesisMatched(ctx); err != nil {
		return nil, err
	}

	return client, nil
}

// ensureGenesisMatched fetches the L2 genesis block from TaikoL1 contract,
// and checks whether the fetched genesis is same to the node local genesis.
func (c *Client) ensureGenesisMatched(ctx context.Context) error {
	L1GenesisHeight, _, _, _, err := c.TaikoL1.GetStateVariables(nil)
	if err != nil {
		return err
	}

	// Fetch the genesis `BlockFinalized` event.
	iter, err := c.TaikoL1.FilterBlockFinalized(&bind.FilterOpts{
		Start: L1GenesisHeight,
		End:   &L1GenesisHeight,
	},
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

		log.Info("Genesis hash", "node", nodeGenesis.Hash(), "TaikoL1", common.BytesToHash(l2GenesisHash[:]))

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

// LatestL2KnownL1Header fetches the L2 node's latest known L1 header.
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

	return c.L1.HeaderByHash(ctx, headL1Origin.L1BlockHash)
}

// GetGenesisL1Header fetches the L1 header that including L2 genesis block.
func (c *Client) GetGenesisL1Header(ctx context.Context) (*types.Header, error) {
	genesisHeight, _, _, _, err := c.TaikoL1.GetStateVariables(nil)
	if err != nil {
		return nil, err
	}

	return c.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(genesisHeight))
}

// ParentByBlockId fetches the block header from L2 node with the largest block id that
// smaller than the given `blockId`.
func (c *Client) ParentByBlockId(ctx context.Context, blockID *big.Int) (*types.Header, error) {
	parentBlockId := new(big.Int).Sub(blockID, common.Big1)

	log.Info("Get parent block by block ID", "parentBlockId", parentBlockId)

	for parentBlockId.Cmp(common.Big0) > 0 {
		l1Origin, err := c.L2.L1OriginByID(ctx, parentBlockId)
		if err != nil {
			return nil, err
		}

		log.Info("Parent block L1 origin", "l1Origin", l1Origin, "parentBlockId", parentBlockId)

		if l1Origin.Throwaway {
			parentBlockId = new(big.Int).Sub(parentBlockId, common.Big1)
			continue
		}

		return c.L2.HeaderByHash(ctx, l1Origin.L2BlockHash)
	}

	return c.L2.HeaderByNumber(ctx, common.Big0)
}
