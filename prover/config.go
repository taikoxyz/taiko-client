package prover

import (
	"crypto/ecdsa"

	"github.com/taikochain/taiko-client/common"
	"github.com/taikochain/taiko-client/crypto"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Taiko prover.
type Config struct {
	L1Endpoint          string
	L2Endpoint          string
	TaikoL1Address      common.Address
	TaikoL2Address      common.Address
	L2TxlistValidator   common.Address
	L1ProverPrivKey     ecdsa.PrivateKey
	ZKEvmRpcdEndpoint   string
	ZkEvmRpcdParamsPath string
	Dummy               bool
	// For testing
	BatchSubmit bool
}

// NewConfigFromCliContext creates a new config instance from command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProverPrivKeyStr := c.String(L1ProverPrivKeyFlag.Name)

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(l1ProverPrivKeyStr))
	if err != nil {
		return nil, err
	}

	return &Config{
		L1Endpoint:          c.String(L1NodeEndpointFlag.Name),
		L2Endpoint:          c.String(L2NodeEndpointFlag.Name),
		TaikoL1Address:      common.HexToAddress(c.String(TaikoL1AddressFlag.Name)),
		TaikoL2Address:      common.HexToAddress(c.String(TaikoL2AddressFlag.Name)),
		L1ProverPrivKey:     *l1ProverPrivKey,
		ZKEvmRpcdEndpoint:   c.String(ZkEvmRpcdEndpointFlag.Name),
		ZkEvmRpcdParamsPath: c.String(ZkEvmRpcdParamsPathFlag.Name),
		Dummy:               c.Bool(DummyFlag.Name),
		BatchSubmit:         c.Bool(BatchSubmitFlag.Name),
	}, nil
}
