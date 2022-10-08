package driver

import (
	"context"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	gethRPC "github.com/ethereum/go-ethereum/rpc"
	"github.com/taikochain/taiko-client/bindings"
	"github.com/taikochain/taiko-client/rpc"
)

// RPCClient contains all L1/L2 RPC clients that a driver needs.
type RPCClient struct {
	l1       *ethclient.Client // L1 node to communicate with
	l2       *ethclient.Client // L2 node to communicate with
	l2RawRPC *gethRPC.Client
	l2Engine *rpc.EngineRPCClient      // L2 node's engine API
	taikoL1  *bindings.TaikoL1Client   // TaikoL1 contract client
	taikoL2  *bindings.V1TaikoL2Client // TaikoL2 contract client (used for anchor tx)
}

// NewRPCClient initializes all RPC clients to run a Taiko driver.
func NewRPCClient(ctx context.Context, cfg *Config) (*RPCClient, error) {
	l1RPC, err := rpc.DialClientWithBackoff(ctx, cfg.L1Endpoint)
	if err != nil {
		return nil, err
	}

	taikoL1, err := bindings.NewTaikoL1Client(cfg.TaikoL1Address, l1RPC)
	if err != nil {
		return nil, err
	}

	l2RPC, err := rpc.DialClientWithBackoff(ctx, cfg.L2Endpoint)
	if err != nil {
		return nil, err
	}

	l2RawRPC, err := gethRPC.Dial(cfg.L2Endpoint)
	if err != nil {
		return nil, err
	}

	l2AuthRPC, err := rpc.DialEngineClientWithBackoff(
		ctx,
		cfg.L2EngineEndpoint,
		cfg.JwtSecret,
	)
	if err != nil {
		return nil, err
	}

	taikoL2, err := bindings.NewV1TaikoL2Client(cfg.TaikoL2Address, l2RPC)
	if err != nil {
		return nil, err
	}

	client := &RPCClient{
		l1:       l1RPC,
		l2:       l2RPC,
		l2RawRPC: l2RawRPC,
		l2Engine: l2AuthRPC,
		taikoL1:  taikoL1,
		taikoL2:  taikoL2,
	}

	if err := client.ensureGenesisMatched(ctx); err != nil {
		return nil, err
	}

	return client, nil
}

// ensureGenesisMatched fetches the L2 genesis block from TaikoL1 contract,
// and checks whether the fetched genesis is same to the node local genesis.
func (c *RPCClient) ensureGenesisMatched(ctx context.Context) error {
	L1GenesisHeight, _, _, _, err := c.taikoL1.GetStateVariables(nil)
	if err != nil {
		return err
	}

	// Fetch the genesis `BlockFinalized` event.
	iter, err := c.taikoL1.FilterBlockFinalized(&bind.FilterOpts{
		Start: L1GenesisHeight,
		End:   &L1GenesisHeight,
	},
		[]*big.Int{common.Big0},
	)
	if err != nil {
		return err
	}

	// Fetch the node's genesis block.
	nodeGenesis, err := c.l2.HeaderByNumber(ctx, common.Big0)
	if err != nil {
		return err
	}

	for iter.Next() {
		l2GenesisHash := iter.Event.BlockHash

		log.Info("Genesis hash", "node", nodeGenesis.Hash(), "TaikoL1", common.BytesToHash(l2GenesisHash[:]))

		// Node's genesis header and TaikoL1 contract's genesis header must match.
		if common.BytesToHash(l2GenesisHash[:]) != nodeGenesis.Hash() {
			return genesisHashMismatchError{
				Node:    nodeGenesis.Hash(),
				TaikoL1: common.BytesToHash(l2GenesisHash[:]),
			}
		} else {
			return nil
		}
	}

	return errGenesisNotFound
}

// LatestL2KnownL1Header fetches the L2 node's latest known L1 header.
func (c *RPCClient) LatestL2KnownL1Header(ctx context.Context) (*types.Header, error) {
	headL1Origin, err := c.l2.HeadL1Origin(ctx)

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

	return c.l1.HeaderByHash(ctx, headL1Origin.L1BlockHash)
}

// GetGenesisL1Header fetches the L1 header that including L2 genesis block.
func (c *RPCClient) GetGenesisL1Header(ctx context.Context) (*types.Header, error) {
	genesisHeight, _, _, _, err := c.taikoL1.GetStateVariables(nil)
	if err != nil {
		return nil, err
	}

	return c.l1.HeaderByNumber(ctx, new(big.Int).SetUint64(genesisHeight))
}

// ParentByBlockId fetches the block header from L2 node with the largest block id that
// smaller than the given `blockId`.
func (c *RPCClient) ParentByBlockId(ctx context.Context, blockID *big.Int) (*types.Header, error) {
	parentBlockId := new(big.Int).Sub(blockID, common.Big1)

	log.Info("Get parent block by block ID", "parentBlockId", parentBlockId)

	for parentBlockId.Cmp(common.Big0) > 0 {
		l1Origin, err := c.l2.L1OriginByID(ctx, parentBlockId)
		if err != nil {
			return nil, err
		}

		log.Info("Parent block L1 origin", "l1Origin", l1Origin, "parentBlockId", parentBlockId)

		if l1Origin.Throwaway {
			parentBlockId = new(big.Int).Sub(parentBlockId, common.Big1)
			continue
		}

		return c.l2.HeaderByHash(ctx, l1Origin.L2BlockHash)
	}

	return c.l2.HeaderByNumber(ctx, common.Big0)
}
