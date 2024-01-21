package rpc

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings"
)

const (
	defaultTimeout = 10 * time.Minute
)

// Client contains all L1Client/L2Client RPC clients that a driver needs.
type Client struct {
	// Geth ethclient clients
	L1Client     *EthClient
	L2Client     *EthClient
	L2CheckPoint *EthClient
	// Geth Engine API clients
	L2AuthClient *EngineClient
	// Protocol contracts clients
	TaikoL1        *bindings.TaikoL1Client
	TaikoL2        *bindings.TaikoL2Client
	TaikoToken     *bindings.TaikoToken
	GuardianProver *bindings.GuardianProver
	// Chain IDs
	L1ChainID *big.Int
	L2ChainID *big.Int
}

// ClientConfig contains all configs which will be used to initializing an
// RPC client. If not providing L2EngineEndpoint or JwtSecret, then the L2AuthClient client
// won't be initialized.
type ClientConfig struct {
	L1Endpoint            string
	L2Endpoint            string
	L2CheckPoint          string
	TaikoL1Address        common.Address
	TaikoL2Address        common.Address
	TaikoTokenAddress     common.Address
	GuardianProverAddress common.Address
	L2EngineEndpoint      string
	JwtSecret             string
	RetryInterval         time.Duration
	Timeout               time.Duration
	BackOffMaxRetries     uint64
}

// NewClient initializes all RPC clients used by Taiko client software.
func NewClient(ctx context.Context, cfg *ClientConfig) (*Client, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	L1Client, err := NewEthClient(cfg.L1Endpoint, cfg.Timeout)
	if err != nil {
		return nil, err
	}

	L2Client, err := NewEthClient(cfg.L2Endpoint, cfg.Timeout)
	if err != nil {
		return nil, err
	}

	l1ChainID, err := L1Client.ChainID(ctxWithTimeout)
	if err != nil {
		return nil, err
	}

	l2ChainID, err := L2Client.ChainID(ctxWithTimeout)
	if err != nil {
		return nil, err
	}

	taikoL1, err := bindings.NewTaikoL1Client(cfg.TaikoL1Address, L1Client)
	if err != nil {
		return nil, err
	}

	taikoL2, err := bindings.NewTaikoL2Client(cfg.TaikoL2Address, L2Client)
	if err != nil {
		return nil, err
	}

	var (
		taikoToken     *bindings.TaikoToken
		guardianProver *bindings.GuardianProver
	)
	if cfg.TaikoTokenAddress.Hex() != ZeroAddress.Hex() {
		if taikoToken, err = bindings.NewTaikoToken(cfg.TaikoTokenAddress, L1Client); err != nil {
			return nil, err
		}
	}
	if cfg.GuardianProverAddress.Hex() != ZeroAddress.Hex() {
		if guardianProver, err = bindings.NewGuardianProver(cfg.GuardianProverAddress, L1Client); err != nil {
			return nil, err
		}
	}

	stateVars, err := taikoL1.GetStateVariables(&bind.CallOpts{Context: ctxWithTimeout})
	if err != nil {
		return nil, err
	}
	isArchive, err := IsArchiveNode(ctxWithTimeout, L1Client, stateVars.A.GenesisHeight)
	if err != nil {
		return nil, err
	}
	if !isArchive {
		return nil, fmt.Errorf("error with RPC endpoint: node (%s) must be archive node", cfg.L1Endpoint)
	}

	// If not providing L2EngineEndpoint or JwtSecret, then the L2AuthClient client
	// won't be initialized.
	var l2AuthClient *EngineClient
	if len(cfg.L2EngineEndpoint) != 0 && len(cfg.JwtSecret) != 0 {
		l2AuthClient, err = NewJWTEngineClient(cfg.L2EngineEndpoint, cfg.JwtSecret)
		if err != nil {
			return nil, err
		}
	}

	var l2CheckPoint *EthClient
	if cfg.L2CheckPoint != "" {
		l2CheckPoint, err = NewEthClient(cfg.L2CheckPoint, cfg.Timeout)
		if err != nil {
			return nil, err
		}
	}

	client := &Client{
		L1Client:       L1Client,
		L2Client:       L2Client,
		L2CheckPoint:   l2CheckPoint,
		L2AuthClient:   l2AuthClient,
		TaikoL1:        taikoL1,
		TaikoL2:        taikoL2,
		TaikoToken:     taikoToken,
		GuardianProver: guardianProver,
		L1ChainID:      l1ChainID,
		L2ChainID:      l2ChainID,
	}

	if err := client.ensureGenesisMatched(ctxWithTimeout); err != nil {
		return nil, err
	}

	return client, nil
}
