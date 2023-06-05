package flags

import (
	"math/rand"

	"github.com/urfave/cli/v2"
)

// Required flags used by proposer.
var (
	L1ProposerPrivKey = &cli.StringFlag{
		Name:     "l1.proposerPrivKey",
		Usage:    "Private key of the L1 proposer, who will send TaikoL1.proposeBlock transactions",
		Required: true,
		Category: proposerCategory,
	}
	L2SuggestedFeeRecipient = &cli.StringFlag{
		Name:     "l2.suggestedFeeRecipient",
		Usage:    "Address of the proposed block's suggested fee recipient",
		Required: true,
		Category: proposerCategory,
	}
)

// Optional flags used by proposer.
var (
	ProposeInterval = &cli.StringFlag{
		Name:     "proposeInterval",
		Usage:    "Time interval to propose L2 pending transactions",
		Category: proposerCategory,
	}
	CommitSlot = &cli.Uint64Flag{
		Name:     "commitSlot",
		Usage:    "The commit slot will be used by proposer, by default, a random number will be used",
		Value:    rand.Uint64(),
		Category: proposerCategory,
	}
	TxPoolLocals = &cli.StringFlag{
		Name:     "txpool.locals",
		Usage:    "Comma separated accounts to treat as locals (priority inclusion)",
		Category: proposerCategory,
	}
	ProposeEmptyBlocksInterval = &cli.StringFlag{
		Name:     "proposeEmptyBlockInterval",
		Usage:    "Time interval to propose empty blocks",
		Category: proposerCategory,
	}
	MinBlockGasLimit = &cli.Uint64Flag{
		Name:     "minimalBlockGasLimit",
		Usage:    "Minimal block gasLimit when proposing a block",
		Category: proposerCategory,
	}
	MaxProposedTxListsPerEpoch = &cli.Uint64Flag{
		Name:     "maxProposedTxListsPerEpoch",
		Value:    1,
		Category: proposerCategory,
	}
)

// All proposer flags.
var ProposerFlags = MergeFlags(CommonFlags, []cli.Flag{
	L2HTTPEndpoint,
	L1ProposerPrivKey,
	L2SuggestedFeeRecipient,
	ProposeInterval,
	CommitSlot,
	TxPoolLocals,
	ProposeEmptyBlocksInterval,
	MinBlockGasLimit,
	MaxProposedTxListsPerEpoch,
})
