package proposer

import (
	"crypto/ecdsa"
	"math/big"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// Config contains all configurations to initialize a Taiko proposer.
type Config struct {
	L1Endpoint                          string
	L2Endpoint                          string
	TaikoL1Address                      common.Address
	TaikoL2Address                      common.Address
	TaikoTokenAddress                   common.Address
	L1ProposerPrivKey                   *ecdsa.PrivateKey
	L2SuggestedFeeRecipient             common.Address
	ProposeInterval                     *time.Duration
	LocalAddresses                      []common.Address
	LocalAddressesOnly                  bool
	ProposeEmptyBlocksInterval          *time.Duration
	MaxProposedTxListsPerEpoch          uint64
	ProposeBlockTxGasLimit              *uint64
	BackOffRetryInterval                time.Duration
	ProposeBlockTxReplacementMultiplier uint64
	RPCTimeout                          *time.Duration
	WaitReceiptTimeout                  time.Duration
	ProposeBlockTxGasTipCap             *big.Int
	ProverEndpoints                     []*url.URL
	BlockProposalFee                    *big.Int
	BlockProposalFeeIncreasePercentage  uint64
	BlockProposalFeeIterations          uint64
}

// Validate checks if the provided configuration is valid.
func (c *Config) Validate() error {
	if err := rpc.CheckURLScheme(c.L1Endpoint, "ws"); err != nil {
		return err
	}
	if err := rpc.CheckURLScheme(c.L2Endpoint, "http"); err != nil {
		return err
	}
	return nil
}
