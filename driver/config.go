package driver

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// Config contains the configurations to initialize a Taiko driver.
type Config struct {
	*rpc.ClientConfig
	P2PSyncVerifiedBlocks bool
	P2PSyncTimeout        time.Duration
	RPCTimeout            time.Duration
	RetryInterval         time.Duration
}

// NewConfigFromCliContext creates a new config instance from
// the command line inputs.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	jwtSecret, err := jwt.ParseSecretFromFile(c.String(flags.JWTSecret.Name))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT secret file: %w", err)
	}

	var (
		p2pSyncVerifiedBlocks = c.Bool(flags.P2PSyncVerifiedBlocks.Name)
		l2CheckPoint          = c.String(flags.CheckPointSyncURL.Name)
	)

	if p2pSyncVerifiedBlocks && len(l2CheckPoint) == 0 {
		return nil, errors.New("empty L2 check point URL")
	}

	if !c.IsSet(flags.L1BeaconEndpoint.Name) {
		return nil, errors.New("empty L1 beacon endpoint")
	}

	var timeout = c.Duration(flags.RPCTimeout.Name)
	return &Config{
		ClientConfig: &rpc.ClientConfig{
			L1Endpoint:       c.String(flags.L1WSEndpoint.Name),
			L1BeaconEndpoint: c.String(flags.L1BeaconEndpoint.Name),
			L2Endpoint:       c.String(flags.L2WSEndpoint.Name),
			L2CheckPoint:     l2CheckPoint,
			TaikoL1Address:   common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
			TaikoL2Address:   common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
			L2EngineEndpoint: c.String(flags.L2AuthEndpoint.Name),
			JwtSecret:        string(jwtSecret),
			Timeout:          timeout,
		},
		RetryInterval:         c.Duration(flags.BackOffRetryInterval.Name),
		P2PSyncVerifiedBlocks: p2pSyncVerifiedBlocks,
		P2PSyncTimeout:        c.Duration(flags.P2PSyncTimeout.Name),
		RPCTimeout:            timeout,
	}, nil
}

func NewConfigFromConfigFile(c *cli.Context, path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, fmt.Errorf("error loading .env config: %w", err)
	}

	jwtSecret, err := jwt.ParseSecretFromFile(os.Getenv("JWT_SECRET"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT secret file: %w", err)
	}

	p2pSyncVerifiedBlocks, err := strconv.ParseBool(os.Getenv("P2P_SYNC_VERIFIED_BLOCKS"))
	if err != nil {
		return nil, fmt.Errorf("error parsing P2P sync verified blocks: %w", err)
	}
	l2CheckPoint := os.Getenv("CHECKPOINT_SYNC_URL")

	if p2pSyncVerifiedBlocks && len(l2CheckPoint) == 0 {
		return nil, errors.New("empty L2 check point URL")
	}

	if os.Getenv("L1_NODE_HTTP_ENDPOINT") == "" {
		return nil, errors.New("empty L1 beacon endpoint")
	}

	timeout, err := time.ParseDuration(os.Getenv("RPC_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("error parsing RPC_TIMEOUT: %w", err)
	}

	retryInterval, err := time.ParseDuration(os.Getenv("RETRY_INTERVAL"))
	if err != nil {
		return nil, fmt.Errorf("error parsing RETRY_INTERVAL: %w", err)
	}

	syncTimeout, err := time.ParseDuration(os.Getenv("P2P_SYNC_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("error parsing P2P_SYNC_TIMEOUT: %w", err)
	}
	return &Config{
		ClientConfig: &rpc.ClientConfig{
			L1Endpoint:       os.Getenv("L1_NODE_WS_ENDPOINT"),
			L1BeaconEndpoint: os.Getenv("L1_NODE_HTTP_ENDPOINT"),
			L2Endpoint:       os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
			L2CheckPoint:     l2CheckPoint,
			TaikoL1Address:   common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
			TaikoL2Address:   common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
			L2EngineEndpoint: os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
			JwtSecret:        string(jwtSecret),
			Timeout:          timeout,
		},
		RetryInterval:         retryInterval,
		P2PSyncVerifiedBlocks: p2pSyncVerifiedBlocks,
		P2PSyncTimeout:        syncTimeout,
		RPCTimeout:            timeout,
	}, nil
}
