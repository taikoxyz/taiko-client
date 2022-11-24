package proposer

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

// Config contains all configurations to initialize a Taiko proposer.
type Config struct {
	L1Endpoint              string
	L2Endpoint              string
	TaikoL1Address          common.Address
	TaikoL2Address          common.Address
	L1ProposerPrivKey       *ecdsa.PrivateKey
	L2SuggestedFeeRecipient common.Address
	ProposeInterval         time.Duration
	ShufflePoolContent      bool
	CommitSlot              uint64

	// Only for testing
	ProduceInvalidBlocks         bool
	ProduceInvalidBlocksInterval uint64
}

// NewConfigFromCliContext initializes a Config instance from
// command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProposerPrivKey, err := crypto.ToECDSA(
		common.Hex2Bytes(c.String(flags.L1ProposerPrivKey.Name)),
	)
	if err != nil {
		return nil, fmt.Errorf("invalid L1 proposer private key: %w", err)
	}

	// Proposing configuration
	proposingInterval, err := time.ParseDuration(c.String(flags.ProposeInterval.Name))
	if err != nil {
		return nil, fmt.Errorf("invalid proposing interval: %w", err)
	}

	l2SuggestedFeeRecipient := c.String(flags.L2SuggestedFeeRecipient.Name)
	if !common.IsHexAddress(l2SuggestedFeeRecipient) {
		return nil, fmt.Errorf("invalid L2 suggested fee recipient address: %s", l2SuggestedFeeRecipient)
	}

	return &Config{
		L1Endpoint:                   c.String(flags.L1NodeEndpoint.Name),
		L2Endpoint:                   c.String(flags.L2NodeEndpoint.Name),
		TaikoL1Address:               common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
		TaikoL2Address:               common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
		L1ProposerPrivKey:            l1ProposerPrivKey,
		L2SuggestedFeeRecipient:      common.HexToAddress(l2SuggestedFeeRecipient),
		ProposeInterval:              proposingInterval,
		ShufflePoolContent:           c.Bool(flags.ShufflePoolContent.Name),
		CommitSlot:                   c.Uint64(flags.CommitSlot.Name),
		ProduceInvalidBlocks:         c.Bool(flags.ProduceInvalidBlocks.Name),
		ProduceInvalidBlocksInterval: c.Uint64(flags.ProduceInvalidBlocksInterval.Name),
	}, nil
}
