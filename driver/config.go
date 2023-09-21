package driver

import (
	"errors"
	"time"
)

// Config contains the configurations to initialize a Taiko driver.
type Config struct {
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
