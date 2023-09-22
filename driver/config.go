package driver

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
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

// Validate checks the configuration settings.
func (c *Config) Validate(ctx context.Context) error {
	if err := rpc.CheckURLScheme(c.L1Endpoint, "ws"); err != nil {
		return err
	}
	if err := rpc.CheckURLScheme(c.L2Endpoint, "ws"); err != nil {
		return err
	}
	if err := rpc.CheckURLScheme(c.L2EngineEndpoint, "http"); err != nil {
		return err
	}

	if c.P2PSyncVerifiedBlocks && len(c.L2CheckPoint) == 0 {
		return errors.New("empty L2 check point URL")
	}

	return nil
}
