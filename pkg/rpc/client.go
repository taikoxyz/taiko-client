package rpc

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	// Geth Engine API clients
	L2Engine *EngineClient
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

	var l2AuthRPC *EngineClient
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
