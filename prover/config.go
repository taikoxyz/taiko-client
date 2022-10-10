package prover

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikochain/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Taiko prover.
type Config struct {
	L1Endpoint          string
	L2Endpoint          string
	TaikoL1Address      common.Address
	TaikoL2Address      common.Address
	L1ProverPrivKey     *ecdsa.PrivateKey
	ZKEvmRpcdEndpoint   string
	ZkEvmRpcdParamsPath string
	Dummy               bool
	// For testing
	BatchSubmit bool
}

// NewConfigFromCliContext creates a new config instance from command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProverPrivKeyStr := c.String(flags.L1ProverPrivKeyFlag.Name)

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(l1ProverPrivKeyStr))
	if err != nil {
		return nil, fmt.Errorf("invalid L1 prover private key: %w", err)
	}

	return &Config{
		L1Endpoint:          c.String(flags.L1NodeEndpoint.Name),
		L2Endpoint:          c.String(flags.L2NodeEndpoint.Name),
		TaikoL1Address:      common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
		TaikoL2Address:      common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
		L1ProverPrivKey:     l1ProverPrivKey,
		ZKEvmRpcdEndpoint:   c.String(flags.ZkEvmRpcdEndpoint.Name),
		ZkEvmRpcdParamsPath: c.String(flags.ZkEvmRpcdParamsPath.Name),
		Dummy:               c.Bool(flags.Dummy.Name),
		BatchSubmit:         c.Bool(flags.BatchSubmit.Name),
	}, nil
}
