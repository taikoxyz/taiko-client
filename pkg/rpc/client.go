package rpc

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/taikoxyz/taiko-client/bindings"
)

// Client contains all L1/L2 RPC clients that a driver needs.
type Client struct {
	// Geth ethclient clients
	L1           *ethclient.Client
	L2           *ethclient.Client
	L2CheckPoint *ethclient.Client
	// Geth gethclient clients
	L1GethClient *gethclient.Client
	L2GethClient *gethclient.Client
	// Geth raw RPC clients
	L1RawRPC *rpc.Client
	L2RawRPC *rpc.Client
	// Geth Engine API clients
	L2Engine *EngineClient
	// Protocol contracts clients
	TaikoL1 *bindings.TaikoL1Client
	TaikoL2 *bindings.TaikoL2Client
	// Chain IDs
	L1ChainID *big.Int
	L2ChainID *big.Int
}

// ClientConfig contains all configs which will be used to initializing an
// RPC client. If not providing L2EngineEndpoint or JwtSecret, then the L2Engine client
// won't be initialized.
type ClientConfig struct {
	L1Endpoint       string
	L2Endpoint       string
	L2CheckPoint     string
	TaikoL1Address   common.Address
	TaikoL2Address   common.Address
	L2EngineEndpoint string
	JwtSecret        string
	RetryInterval    time.Duration
}

// NewClient initializes all RPC clients used by Taiko client softwares.
func NewClient(ctx context.Context, cfg *ClientConfig) (*Client, error) {
	l1RPC, err := DialClientWithBackoff(ctx, cfg.L1Endpoint, cfg.RetryInterval)
	if err != nil {
		return nil, err
	}

	taikoL1, err := bindings.NewTaikoL1Client(cfg.TaikoL1Address, l1RPC)
	if err != nil {
		return nil, err
	}

	l2RPC, err := DialClientWithBackoff(ctx, cfg.L2Endpoint, cfg.RetryInterval)
	if err != nil {
		return nil, err
	}

	taikoL2, err := bindings.NewTaikoL2Client(cfg.TaikoL2Address, l2RPC)
	if err != nil {
		return nil, err
	}

	l1RawRPC, err := rpc.Dial(cfg.L1Endpoint)
	if err != nil {
		return nil, err
	}

	l2RawRPC, err := rpc.Dial(cfg.L2Endpoint)
	if err != nil {
		return nil, err
	}

	l1ChainID, err := l1RPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	l2ChainID, err := l2RPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	// If not providing L2EngineEndpoint or JwtSecret, then the L2Engine client
	// won't be initialized.
	var l2AuthRPC *EngineClient
	if len(cfg.L2EngineEndpoint) != 0 && len(cfg.JwtSecret) != 0 {
		if l2AuthRPC, err = DialEngineClientWithBackoff(
			ctx,
			cfg.L2EngineEndpoint,
			cfg.JwtSecret,
			cfg.RetryInterval,
		); err != nil {
			return nil, err
		}
	}

	var l2CheckPoint *ethclient.Client
	if len(cfg.L2CheckPoint) != 0 {
		if l2CheckPoint, err = DialClientWithBackoff(ctx, cfg.L2CheckPoint, cfg.RetryInterval); err != nil {
			return nil, err
		}
	}

	client := &Client{
		L1:           l1RPC,
		L2:           l2RPC,
		L2CheckPoint: l2CheckPoint,
		L1RawRPC:     l1RawRPC,
		L2RawRPC:     l2RawRPC,
		L1GethClient: gethclient.New(l1RawRPC),
		L2GethClient: gethclient.New(l2RawRPC),
		L2Engine:     l2AuthRPC,
		TaikoL1:      taikoL1,
		TaikoL2:      taikoL2,
		L1ChainID:    l1ChainID,
		L2ChainID:    l2ChainID,
	}

	if err := client.ensureGenesisMatched(ctx); err != nil {
		return nil, err
	}

	return client, nil
}
