package driver

import (
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Config contains the configurations to initialize a Taiko driver.
type Config struct {
	L1Endpoint           string
	L2Endpoint           string
	L2EngineEndpoint     string
	TaikoL1Address       common.Address
	TaikoL2Address       common.Address
	JwtSecret            string
	BackOffRetryInterval time.Duration
	RPCTimeout           *time.Duration
	// 后面的有用
	L2CheckPoint          string
	P2PSyncVerifiedBlocks bool
	P2PSyncTimeout        time.Duration
}

func (c *Config) Check() error {
	if c.P2PSyncVerifiedBlocks && len(c.L2CheckPoint) == 0 {
		return errors.New("empty L2 check point URL")
	}
	return nil
}
