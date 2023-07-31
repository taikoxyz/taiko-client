package driver

import (
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Taiko driver.
type Config struct {
	L1Endpoint            string
	L2Endpoint            string
	L2EngineEndpoint      string
	L2CheckPoint          string
	TaikoL1Address        common.Address
	TaikoL2Address        common.Address
	JwtSecret             string
	P2PSyncVerifiedBlocks bool
	P2PSyncTimeout        time.Duration
	BackOffRetryInterval  time.Duration
	RPCTimeout            *time.Duration
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
		l2CheckPoint          = c.String(flags.CheckPointSyncUrl.Name)
	)

	if p2pSyncVerifiedBlocks && len(l2CheckPoint) == 0 {
		return nil, errors.New("empty L2 check point URL")
	}

	var timeout *time.Duration

	if c.IsSet(flags.RPCTimeout.Name) {
		duration := time.Duration(c.Uint64(flags.RPCTimeout.Name)) * time.Second
		timeout = &duration
	}

	return &Config{
		L1Endpoint:            c.String(flags.L1WSEndpoint.Name),
		L2Endpoint:            c.String(flags.L2WSEndpoint.Name),
		L2EngineEndpoint:      c.String(flags.L2AuthEndpoint.Name),
		L2CheckPoint:          l2CheckPoint,
		TaikoL1Address:        common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
		TaikoL2Address:        common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
		JwtSecret:             string(jwtSecret),
		P2PSyncVerifiedBlocks: p2pSyncVerifiedBlocks,
		P2PSyncTimeout:        time.Duration(int64(time.Second) * int64(c.Uint(flags.P2PSyncTimeout.Name))),
		BackOffRetryInterval:  time.Duration(c.Uint64(flags.BackOffRetryInterval.Name)) * time.Second,
		RPCTimeout:            timeout,
	}, nil
}
