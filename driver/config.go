package driver

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Taiko driver.
type Config struct {
	L1Endpoint                    string
	L2Endpoint                    string
	L2EngineEndpoint              string
	TaikoL1Address                common.Address
	TaikoL2Address                common.Address
	ThrowawayBlocksBuilderPrivKey *ecdsa.PrivateKey
	JwtSecret                     string
	FromGenesis                   bool
}

// NewConfigFromCliContext creates a new config instance from
// the command line inputs.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	jwtSecret, err := jwt.ParseSecretFromFile(c.String(flags.JWTSecret.Name))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT secret file: %w", err)
	}

	throwawayBlocksBuilderPrivKey, err := crypto.HexToECDSA(c.String(flags.ThrowawayBlocksBuilderPrivKey.Name))
	if err != nil {
		return nil, fmt.Errorf("invalid throwaway blocks builder private key: %w", err)
	}

	return &Config{
		L1Endpoint:                    c.String(flags.L1NodeEndpoint.Name),
		L2Endpoint:                    c.String(flags.L2NodeEndpoint.Name),
		L2EngineEndpoint:              c.String(flags.L2NodeEngineEndpoint.Name),
		TaikoL1Address:                common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
		TaikoL2Address:                common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
		ThrowawayBlocksBuilderPrivKey: throwawayBlocksBuilderPrivKey,
		JwtSecret:                     string(jwtSecret),
	}, nil
}
