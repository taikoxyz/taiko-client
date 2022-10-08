package proposer

import (
	"github.com/taikochain/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

const (
	defaultProposeInterval      = "120s"
	defaultProduceInvalidBlocks = false
)

// Config contains all configurations to initialize a Taiko proposer.
type Config struct {
	L1Node                 string
	L2Node                 string
	TaikoL1Address         string
	TaikoL2Address         string
	L1ProposerPrivKey      string
	L2SuggestedFeeRecipien string
	ProposeInterval        string

	// Only for testing
	ProduceInvalidBlocks         bool
	ProduceInvalidBlocksInterval uint64
}

// NewConfigFromCliContext initializes a Config instance from
// command line flags.
func NewConfigFromCliContext(ctx *cli.Context) (*Config, error) {
	cfg := &Config{}

	if ctx.IsSet(flags.L1NodeEndpoint.Name) {
		cfg.L1Node = ctx.String(flags.L1NodeEndpoint.Name)
	}
	if ctx.IsSet(flags.L2NodeEndpoint.Name) {
		cfg.L2Node = ctx.String(flags.L2NodeEndpoint.Name)
	}
	if ctx.IsSet(flags.TaikoL1Address.Name) {
		cfg.TaikoL1Address = ctx.String(flags.TaikoL1Address.Name)
	}
	if ctx.IsSet(flags.TaikoL2Address.Name) {
		cfg.TaikoL2Address = ctx.String(flags.TaikoL2Address.Name)
	}
	if ctx.IsSet(flags.L1ProposerPrivKey.Name) {
		cfg.L1ProposerPrivKey = ctx.String(flags.L1ProposerPrivKey.Name)
	}
	if ctx.IsSet(flags.L2SuggestedFeeRecipient.Name) {
		cfg.L2SuggestedFeeRecipien = ctx.String(flags.L2SuggestedFeeRecipient.Name)
	}
	if ctx.IsSet(flags.ProposeInterval.Name) {
		cfg.ProposeInterval = ctx.String(flags.ProposeInterval.Name)
	} else {
		cfg.ProposeInterval = defaultProposeInterval
	}
	if ctx.IsSet(flags.ProduceInvalidBlocks.Name) {
		cfg.ProduceInvalidBlocks = ctx.Bool(flags.ProduceInvalidBlocks.Name)
	} else {
		cfg.ProduceInvalidBlocks = defaultProduceInvalidBlocks
	}
	if ctx.IsSet(flags.ProduceInvalidBlocksInterval.Name) {
		cfg.ProduceInvalidBlocksInterval = ctx.Uint64(
			flags.ProduceInvalidBlocksInterval.Name,
		)
	}

	return cfg, nil
}
