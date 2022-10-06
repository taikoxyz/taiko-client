package main

import (
	"github.com/taikochain/client-mono/proposer"
	"github.com/urfave/cli/v2"
)

const (
	defaultProposeInterval      = "120s"
	defaultProduceInvalidBlocks = false
)

// NewConfigFromCliContext initializes a Config instance from
// command line flags.
func NewConfigFromCliContext(ctx *cli.Context) (*proposer.Config, error) {
	cfg := &proposer.Config{}

	if ctx.IsSet(L1NodeFlag.Name) {
		cfg.L1Node = ctx.String(L1NodeFlag.Name)
	}
	if ctx.IsSet(L2NodeFlag.Name) {
		cfg.L2Node = ctx.String(L2NodeFlag.Name)
	}
	if ctx.IsSet(TaikoL1AddressFlag.Name) {
		cfg.TaikoL1Address = ctx.String(TaikoL1AddressFlag.Name)
	}
	if ctx.IsSet(TaikoL2AddressFlag.Name) {
		cfg.TaikoL2Address = ctx.String(TaikoL2AddressFlag.Name)
	}
	if ctx.IsSet(L1ProposerPrivKeyFlag.Name) {
		cfg.L1ProposerPrivKey = ctx.String(L1ProposerPrivKeyFlag.Name)
	}
	if ctx.IsSet(L2SuggestedFeeRecipientFlag.Name) {
		cfg.L2SuggestedFeeRecipien = ctx.String(L2SuggestedFeeRecipientFlag.Name)
	}
	if ctx.IsSet(ProposeIntervalFlag.Name) {
		cfg.ProposeInterval = ctx.String(ProposeIntervalFlag.Name)
	} else {
		cfg.ProposeInterval = defaultProposeInterval
	}
	if ctx.IsSet(ProduceInvalidBlocksFlag.Name) {
		cfg.ProduceInvalidBlocks = ctx.Bool(ProduceInvalidBlocksFlag.Name)
	} else {
		cfg.ProduceInvalidBlocks = defaultProduceInvalidBlocks
	}
	if ctx.IsSet(ProduceInvalidBlocksInterval.Name) {
		cfg.ProduceInvalidBlocksInterval = ctx.Uint64(ProduceInvalidBlocksInterval.Name)
	}

	return cfg, nil
}
