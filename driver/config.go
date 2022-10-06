package driver

import (
	"crypto/ecdsa"

	"github.com/taikochain/client-mono/util"
	"github.com/taikochain/taiko-client/common"
	"github.com/taikochain/taiko-client/crypto"
	"github.com/taikochain/taiko-client/log"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Taiko driver.
type Config struct {
	L1Endpoint                    string
	L2Endpoint                    string
	L2AuthEndpoint                string
	TaikoL1Address                common.Address
	TaikoL2Address                common.Address
	ThrowawayBlocksBuilderPrivKey *ecdsa.PrivateKey
	JwtSecret                     string
}

// NewConfigFromCliContext creates a new config instance from
// the command line inputs.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	jwtSecret, err := util.ParseJWTSecretFromFile(c.String(JWTSecret.Name))
	if err != nil {
		log.Crit("Parse JWT secret from file error", "error", err)
	}

	throwawayBlocksBuilderPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(c.String(ThrowawayBlocksBuilderPrivKey.Name)))
	if err != nil {
		log.Crit("Parse throwaway blocks builder private key error", "error", err)
	}

	return &Config{
		L1Endpoint:                    c.String(L1NodeEndpoint.Name),
		L2Endpoint:                    c.String(L2NodeEndpoint.Name),
		L2AuthEndpoint:                c.String(L2NodeAuthEndpoint.Name),
		TaikoL1Address:                common.HexToAddress(c.String(TaikoL1Address.Name)),
		TaikoL2Address:                common.HexToAddress(c.String(TaikoL2Address.Name)),
		ThrowawayBlocksBuilderPrivKey: throwawayBlocksBuilderPrivKey,
		JwtSecret:                     string(jwtSecret),
	}, nil
}
