package driver

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/cmd/utils"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Taiko driver.
type Config struct {
	*rpc.ClientConfig
	P2PSyncVerifiedBlocks bool
	P2PSyncTimeout        time.Duration
	RPCTimeout            time.Duration
	RetryInterval         time.Duration
	MaxExponent           uint64
	BlobServerEndpoint    *url.URL
}

// NewConfigFromCliContext creates a new config instance from
// the command line inputs.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	// Defaults config
	cfg := Config{
		ClientConfig: &rpc.ClientConfig{
			// TODO: To be confirmed whether taiko addresses in L1 or L2 are constant
			TaikoL1Address: common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
			TaikoL2Address: common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
			Timeout:        12 * time.Second,
		},
		RetryInterval:         backoff.DefaultMaxInterval,
		P2PSyncVerifiedBlocks: false,
		P2PSyncTimeout:        1 * time.Hour,
		RPCTimeout:            12 * time.Second,
		MaxExponent:           0,
	}

	// Load config file
	if file := c.String(flags.ConfigFile.Name); file != "" {
		if err := utils.LoadConfigFile(file, &cfg); err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	}
	// Apply flag value
	err := ApplyFlagValue(c, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ApplyFlagValue(c *cli.Context, cfg *Config) error {
	if c.IsSet(flags.L1WSEndpoint.Name) {
		cfg.ClientConfig.L1Endpoint = c.String(flags.L1WSEndpoint.Name)
	}

	if !c.IsSet(flags.L1BeaconEndpoint.Name) {
		return errors.New("empty L1 beacon endpoint")
	}
	cfg.ClientConfig.L1BeaconEndpoint = c.String(flags.L1BeaconEndpoint.Name)

	if c.IsSet(flags.L2WSEndpoint.Name) {
		cfg.ClientConfig.L2Endpoint = c.String(flags.L2WSEndpoint.Name)
	}

	if c.IsSet(flags.TaikoL1Address.Name) {
		cfg.ClientConfig.TaikoL1Address = common.HexToAddress(c.String(flags.TaikoL1Address.Name))
	}

	if c.IsSet(flags.TaikoL2Address.Name) {
		cfg.ClientConfig.TaikoL2Address = common.HexToAddress(c.String(flags.TaikoL2Address.Name))
	}

	if c.IsSet(flags.L2AuthEndpoint.Name) {
		cfg.ClientConfig.L2EngineEndpoint = c.String(flags.L2AuthEndpoint.Name)
	}

	if c.IsSet(flags.RPCTimeout.Name) {
		cfg.ClientConfig.Timeout = c.Duration(flags.RPCTimeout.Name)
		cfg.RPCTimeout = c.Duration(flags.RPCTimeout.Name)
	}

	if c.IsSet(flags.BackOffRetryInterval.Name) {
		cfg.RetryInterval = c.Duration(flags.BackOffRetryInterval.Name)
	}

	if c.IsSet(flags.P2PSyncTimeout.Name) {
		cfg.P2PSyncTimeout = c.Duration(flags.P2PSyncTimeout.Name)
	}

	if c.IsSet(flags.MaxExponent.Name) {
		cfg.MaxExponent = c.Uint64(flags.MaxExponent.Name)
	}

	if c.IsSet(flags.BlobServerEndpoint.Name) {
		blobServerEndpoint, err := url.Parse(
			c.String(flags.BlobServerEndpoint.Name),
		)
		if err != nil {
			return err
		}
		cfg.BlobServerEndpoint = blobServerEndpoint
	}

	jwtSecret, err := jwt.ParseSecretFromFile(c.String(flags.JWTSecret.Name))
	if err != nil {
		return fmt.Errorf("invalid JWT secret file: %w", err)
	}
	cfg.ClientConfig.JwtSecret = string(jwtSecret)

	// Must be defined via flags
	p2pSyncVerifiedBlocks := c.Bool(flags.P2PSyncVerifiedBlocks.Name)
	l2CheckPoint := c.String(flags.CheckPointSyncURL.Name)
	if p2pSyncVerifiedBlocks && len(l2CheckPoint) == 0 {
		return errors.New("empty L2 check point URL")
	}
	cfg.ClientConfig.L2CheckPoint = l2CheckPoint
	cfg.P2PSyncVerifiedBlocks = p2pSyncVerifiedBlocks

	return nil
}
