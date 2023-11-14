package flags

import (
	"github.com/taikoxyz/taiko-client/version"
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
	ProverEndpoints = &cli.StringFlag{
		Name:     "proverEndpoints",
		Usage:    "Comma-delineated list of prover endpoints proposer should query when attempting to propose a block",
		Required: true,
		Category: proposerCategory,
	}
)

// Optional flags used by proposer.
var (
	// Tier fee related.
	OptimisticTierFee = &cli.Uint64Flag{
		Name:     "tierFee.optimistic",
		Usage:    "Initial tier fee (in wei) paid to prover to generate an optimistic proofs",
		Category: proposerCategory,
	}
	SgxTierFee = &cli.Uint64Flag{
		Name:     "tierFee.sgx",
		Usage:    "Initial tier fee (in wei) paid to prover to generate a SGX proofs",
		Category: proposerCategory,
	}
	PseZkevmTierFee = &cli.Uint64Flag{
		Name:     "tierFee.pseZKEvm",
		Usage:    "Initial tier fee (in wei) paid to prover to generate a PSE zkEVM proofs",
		Category: proposerCategory,
	}
	SgxAndPseZkevmTierFee = &cli.Uint64Flag{
		Name:     "tierFee.sgxAndPseZKEvm",
		Usage:    "Initial tier fee (in wei) paid to prover to generate a SGX + PSE zkEVM proofs",
		Category: proposerCategory,
	}
	TierFeePriceBump = &cli.Uint64Flag{
		Name:     "tierFee.pricebump",
		Usage:    "Price bump percentage when no prover wants to accept the block at initial fee",
		Value:    10,
		Category: proposerCategory,
	}
	MaxTierFeePriceBumps = &cli.Uint64Flag{
		Name:     "tierFee.maxPriceBumps",
		Usage:    "If nobody accepts block at initial tier fee, how many iterations to increase tier fee before giving up",
		Category: proposerCategory,
		Value:    3,
	}
	// Proposing epoch related.
	ProposeInterval = &cli.DurationFlag{
		Name:     "epoch.interval",
		Usage:    "Time interval to propose L2 pending transactions",
		Category: proposerCategory,
	}
	ProposeEmptyBlocksInterval = &cli.DurationFlag{
		Name:     "epoch.emptyBlockInterval",
		Usage:    "Time interval to propose empty blocks",
		Category: proposerCategory,
	}
	// Proposing metadata related.
	ExtraData = &cli.StringFlag{
		Name:     "extraData",
		Usage:    "Block extra data set by the proposer (default = client version)",
		Value:    version.VersionWithCommit(),
		Category: proposerCategory,
	}
	// Transactions pool related.
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
	MaxProposedTxListsPerEpoch = &cli.Uint64Flag{
		Name:     "txpool.maxTxListsPerEpoch",
		Usage:    "Maximum number of transaction lists which will be proposed inside one proposing epoch",
		Value:    1,
		Category: proposerCategory,
	}
	// Transaction related.
	ProposeBlockTxGasLimit = &cli.Uint64Flag{
		Name:     "tx.gasLimit",
		Usage:    "Gas limit will be used for TaikoL1.proposeBlock transactions",
		Category: proposerCategory,
	}
	ProposeBlockTxReplacementMultiplier = &cli.Uint64Flag{
		Name:     "tx.replacementMultiplier",
		Value:    2,
		Usage:    "Gas tip multiplier when replacing a TaikoL1.proposeBlock transaction with same nonce",
		Category: proposerCategory,
	}
	ProposeBlockTxGasTipCap = &cli.Uint64Flag{
		Name:     "tx.gasTipCap",
		Usage:    "Gas tip cap (in wei) for a TaikoL1.proposeBlock transaction when doing the transaction replacement",
		Category: proposerCategory,
	}
	ProposeBlockIncludeParentMetaHash = &cli.BoolFlag{
		Name:     "includeParentMetaHash",
		Usage:    "Include parent meta hash when proposing block",
		Value:    false,
		Category: proposerCategory,
	}
)

// All proposer flags.
var ProposerFlags = MergeFlags(CommonFlags, []cli.Flag{
	L2HTTPEndpoint,
	TaikoTokenAddress,
	L1ProposerPrivKey,
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
	OptimisticTierFee,
	SgxTierFee,
	PseZkevmTierFee,
	SgxAndPseZkevmTierFee,
	TierFeePriceBump,
	MaxTierFeePriceBumps,
	ProposeBlockIncludeParentMetaHash,
})
