package flags

import (
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
	ProverEndpoints = &cli.StringFlag{
		Name:     "proverEndpoints",
		Usage:    "Comma-delineated list of prover endpoints proposer should query when attempting to propose a block",
		Required: true,
		Category: proposerCategory,
	}
	BlockProposalFee = &cli.StringFlag{
		Name:     "blockProposalFee",
		Usage:    "Initial block proposal fee (in wei) paid on block proposing",
		Required: true,
		Category: proposerCategory,
	}
	TaikoTokenAddress = &cli.StringFlag{
		Name:     "taikoToken",
		Usage:    "TaikoToken contract address",
		Required: true,
		Category: proposerCategory,
	}
)

// Optional flags used by proposer.
var (
	ExtraData = &cli.StringFlag{
		Name:     "extraData",
		Usage:    "Block extra data set by the proposer",
		Value:    "",
		Category: proposerCategory,
	}
	ProposeInterval = &cli.DurationFlag{
		Name:     "proposeInterval",
		Usage:    "Time interval to propose L2 pending transactions",
		Category: proposerCategory,
	}
	TxPoolLocals = &cli.StringSliceFlag{
		Name:     "txpool.locals",
		Usage:    "Comma separated accounts to treat as locals (priority inclusion)",
		Category: proposerCategory,
	}
	TxPoolLocalsOnly = &cli.BoolFlag{
		Name:     "txpool.localsOnly",
		Usage:    "If set to true, proposer will only propose transactions of local accounts",
		Value:    false,
		Category: proposerCategory,
	}
	ProposeEmptyBlocksInterval = &cli.DurationFlag{
		Name:     "proposeEmptyBlockInterval",
		Usage:    "Time interval to propose empty blocks",
		Category: proposerCategory,
	}
	MaxProposedTxListsPerEpoch = &cli.Uint64Flag{
		Name:     "maxProposedTxListsPerEpoch",
		Value:    1,
		Category: proposerCategory,
	}
	ProposeBlockTxGasLimit = &cli.Uint64Flag{
		Name:     "proposeBlockTxGasLimit",
		Usage:    "Gas limit will be used for TaikoL1.proposeBlock transactions",
		Category: proposerCategory,
	}
	ProposeBlockTxReplacementMultiplier = &cli.Uint64Flag{
		Name:     "proposeBlockTxReplacementMultiplier",
		Value:    2,
		Usage:    "Gas tip multiplier when replacing a TaikoL1.proposeBlock transaction with same nonce",
		Category: proposerCategory,
	}
	ProposeBlockTxGasTipCap = &cli.Uint64Flag{
		Name:     "proposeBlockTxGasTipCap",
		Usage:    "Gas tip cap (in wei) for a TaikoL1.proposeBlock transaction when doing the transaction replacement",
		Category: proposerCategory,
	}
	BlockProposalFeeIncreasePercentage = &cli.Uint64Flag{
		Name:     "blockProposalFeeIncreasePercentage",
		Usage:    "Increase fee by what percentage when no prover wants to accept the block at initial fee",
		Category: proposerCategory,
		Value:    10,
	}
	BlockProposalFeeIterations = &cli.Uint64Flag{
		Name:     "blockProposalFeeIterations",
		Usage:    "If nobody accepts block at initial fee, how many iterations to increase fee before giving up",
		Category: proposerCategory,
		Value:    3,
	}
)

// All proposer flags.
var ProposerFlags = MergeFlags(CommonFlags, []cli.Flag{
	L2HTTPEndpoint,
	L1ProposerPrivKey,
	L2SuggestedFeeRecipient,
	ProposeInterval,
	TxPoolLocals,
	TxPoolLocalsOnly,
	ExtraData,
	ProposeEmptyBlocksInterval,
	MaxProposedTxListsPerEpoch,
	ProposeBlockTxGasLimit,
	ProposeBlockTxReplacementMultiplier,
	ProposeBlockTxGasTipCap,
	ProverEndpoints,
	BlockProposalFee,
	BlockProposalFeeIncreasePercentage,
	BlockProposalFeeIterations,
	TaikoTokenAddress,
})
