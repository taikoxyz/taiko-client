package main

import (
	"github.com/urfave/cli/v2"
)

// Required flags.
var (
	L1NodeFlag = cli.StringFlag{
		Name:     "l1",
		Usage:    "RPC endpoint of a L1 ethereum node",
		Required: true,
	}
	L2NodeFlag = cli.StringFlag{
		Name:     "l2",
		Usage:    "RPC endpoint of a L2 ethereum node",
		Required: true,
	}
	TaikoL1AddressFlag = cli.StringFlag{
		Name:     "taikoL1",
		Usage:    "TaikoL1 contract address",
		Required: true,
	}
	TaikoL2AddressFlag = cli.StringFlag{
		Name:     "taikoL2",
		Usage:    "TaikoL2 contract address",
		Required: true,
	}
	L1ProposerPrivKeyFlag = cli.StringFlag{
		Name:     "l1.proposerPrivKey",
		Usage:    "Private key for L1 proposer, who will send TaikoL1.proposeBlock transactions to the L1 node",
		Required: true,
	}
	L2SuggestedFeeRecipientFlag = cli.StringFlag{
		Name:     "l2.suggestedFeeRecipient",
		Usage:    "Address of the proposed block's suggested fee recipient",
		Required: true,
	}
	ProposeIntervalFlag = cli.StringFlag{
		Name:     "proposeInterval",
		Usage:    "Interval for proposing L2 node's new pending transactions",
		Required: true,
	}
)

// Special flags for testing.
var (
	ProduceInvalidBlocksFlag = cli.BoolFlag{
		Name:   "produceInvalidBlocks",
		Usage:  "Special flag for testnet testing, if activated, the proposer will start producing bad blocks",
		Hidden: true,
	}
	ProduceInvalidBlocksInterval = cli.Uint64Flag{
		Name:   "produceInvalidBlocksInterval",
		Usage:  "Special flag for testnet testing, if activated, bad blocks will be produced every N valid blocks",
		Hidden: true,
	}
)

// All flags.
var Flags = []cli.Flag{
	&L1NodeFlag,
	&L2NodeFlag,
	&TaikoL1AddressFlag,
	&TaikoL2AddressFlag,
	&L1ProposerPrivKeyFlag,
	&L2SuggestedFeeRecipientFlag,
	&ProposeIntervalFlag,
	&ProduceInvalidBlocksFlag,
	&ProduceInvalidBlocksInterval,
}
