package main

import (
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/urfave/cli/v2"
)

const proposerCmd = "proposer"

var proposerConf = &proposer.Config{}

// Required flags used by proposer.
var (
	L1ProposerPrivKeyFlag = &cli.StringFlag{
		Name:     "l1.proposerPrivKey",
		Usage:    "Private key of the L1 proposer, who will send TaikoL1.proposeBlock transactions",
		Required: true,
		Category: proposerCategory,
		Action: func(c *cli.Context, v string) error {
			k, err := crypto.ToECDSA(common.Hex2Bytes(v))
			if err != nil {
				return fmt.Errorf("invalid L1 proposer private key: %w", err)
			}
			proposerConf.L1ProposerPrivKey = k
			return nil
		},
	}
	L2SuggestedFeeRecipientFlag = &cli.StringFlag{
		Name:     "l2.suggestedFeeRecipient",
		Usage:    "Address of the proposed block's suggested fee recipient",
		Required: true,
		Category: proposerCategory,
		Action: func(c *cli.Context, v string) error {
			if !common.IsHexAddress(v) {
				return fmt.Errorf("invalid L2 suggested fee recipient address: %s", v)
			}
			proposerConf.L2SuggestedFeeRecipient = common.HexToAddress(v)
			return nil
		},
	}
	ProverEndpointsFlag = &cli.StringSliceFlag{
		Name:     "proverEndpoints",
		Usage:    "Comma-delineated list of prover endpoints proposer should query when attempting to propose a block",
		Category: proposerCategory,
		Action: func(c *cli.Context, v []string) error {
			for _, e := range v {
				endpoint, err := url.Parse(e)
				if err != nil {
					return err
				}
				proposerConf.ProverEndpoints = append(proposerConf.ProverEndpoints, endpoint)
			}
			return nil
		},
	}
	BlockProposalFeeFlag = &cli.StringFlag{
		Name:     "blockProposalFee",
		Usage:    "Initial block proposal fee (in wei) paid on block proposing",
		Category: proposerCategory,
		Action: func(c *cli.Context, v string) error {
			fee, ok := new(big.Int).SetString(v, 10)
			if !ok {
				return fmt.Errorf("invalid blockProposalFee: %v", v)
			}
			proposerConf.BlockProposalFee = fee
			return nil
		},
	}
	TaikoTokenAddressFlag = &cli.StringFlag{
		Name:     "taikoToken",
		Usage:    "TaikoToken contract address",
		Required: true,
		Category: proposerCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.TaikoTokenAddress = common.HexToAddress(v)
			return nil
		},
	}
)

// Optional flags used by proposer.
var (
	ProposeIntervalFlag = &cli.DurationFlag{
		Name:     "proposeInterval",
		Usage:    "Time interval in `duration` to propose L2 pending transactions",
		Category: proposerCategory,
		Action: func(c *cli.Context, v time.Duration) error {
			proposerConf.ProposeInterval = &v
			return nil
		},
	}
	TxPoolLocalsFlag = &cli.StringSliceFlag{
		Name:     "txpool.locals",
		Usage:    "Comma separated `accounts` to treat as locals (priority inclusion)",
		Category: proposerCategory,
		Action: func(c *cli.Context, v []string) error {
			for _, account := range v {
				if trimmed := strings.TrimSpace(account); !common.IsHexAddress(trimmed) {
					return fmt.Errorf("invalid account in --txpool.locals: %s", trimmed)
				} else {
					proposerConf.LocalAddresses = append(proposerConf.LocalAddresses, common.HexToAddress(account))
				}
			}
			return nil
		},
	}
	TxPoolLocalsOnlyFlag = &cli.BoolFlag{
		Name:        "txpool.localsOnly",
		Usage:       "If set to true, proposer will only propose transactions of local accounts",
		Value:       false,
		Category:    proposerCategory,
		Destination: &proposerConf.LocalAddressesOnly,
		Action: func(c *cli.Context, v bool) error {
			proposerConf.LocalAddressesOnly = v
			return nil
		},
	}
	ProposeEmptyBlocksIntervalFlag = &cli.DurationFlag{
		Name:     "proposeEmptyBlockInterval",
		Usage:    "Time interval in `duration` to propose empty blocks",
		Category: proposerCategory,
		Action: func(c *cli.Context, v time.Duration) error {
			proposerConf.ProposeEmptyBlocksInterval = &v
			return nil
		},
	}
	MaxProposedTxListsPerEpochFlag = &cli.Uint64Flag{
		Name:        "maxProposedTxListsPerEpoch",
		Value:       1,
		Category:    proposerCategory,
		Destination: &proposerConf.MaxProposedTxListsPerEpoch,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.MaxProposedTxListsPerEpoch = v
			return nil
		},
	}
	ProposeBlockTxGasLimitFlag = &cli.Uint64Flag{
		Name:     "proposeBlockTxGasLimit",
		Usage:    "Gas limit will be used for TaikoL1.proposeBlock transactions",
		Category: proposerCategory,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.ProposeBlockTxGasLimit = &v
			return nil
		},
	}
	ProposeBlockTxReplacementMultiplierFlag = &cli.Uint64Flag{
		Name:        "proposeBlockTxReplacementMultiplier",
		Value:       2,
		Usage:       "Gas tip multiplier when replacing a TaikoL1.proposeBlock transaction with same nonce",
		Category:    proposerCategory,
		Destination: &proposerConf.ProposeBlockTxReplacementMultiplier,
		Action: func(c *cli.Context, v uint64) error {
			if v == 0 {
				return fmt.Errorf("invalid --proposeBlockTxReplacementMultiplier value: %d", v)
			}
			proposerConf.ProposeBlockTxReplacementMultiplier = v
			return nil
		},
	}
	ProposeBlockTxGasTipCapFlag = &cli.Uint64Flag{
		Name:     "proposeBlockTxGasTipCap",
		Usage:    "Gas tip cap (in wei) for a TaikoL1.proposeBlock transaction when doing the transaction replacement",
		Category: proposerCategory,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.ProposeBlockTxGasTipCap = new(big.Int).SetUint64(v)
			return nil
		},
	}
	BlockProposalFeeIncreasePercentageFlag = &cli.Uint64Flag{
		Name:        "blockProposalFeeIncreasePercentage",
		Usage:       "Increase fee by what percentage when no prover wants to accept the block at initial fee",
		Category:    proposerCategory,
		Value:       10,
		Destination: &proposerConf.BlockProposalFeeIncreasePercentage,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.BlockProposalFeeIncreasePercentage = v
			return nil
		},
	}
	BlockProposalFeeIterationsFlag = &cli.Uint64Flag{
		Name:        "blockProposalFeeIterations",
		Usage:       "If nobody accepts block at initial fee, how many iterations to increase fee before giving up",
		Category:    proposerCategory,
		Value:       3,
		Destination: &proposerConf.BlockProposalFeeIterations,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.BlockProposalFeeIterations = v
			return nil
		},
	}
)

// All proposer flags.
var proposerFlags = MergeFlags(CommonFlags, []cli.Flag{
	L2HTTPEndpointFlag,
	L1ProposerPrivKeyFlag,
	L2SuggestedFeeRecipientFlag,
	ProposeIntervalFlag,
	TxPoolLocalsFlag,
	TxPoolLocalsOnlyFlag,
	ProposeEmptyBlocksIntervalFlag,
	MaxProposedTxListsPerEpochFlag,
	ProposeBlockTxGasLimitFlag,
	ProposeBlockTxReplacementMultiplierFlag,
	ProposeBlockTxGasTipCapFlag,
	ProverEndpointsFlag,
	BlockProposalFeeFlag,
	BlockProposalFeeIncreasePercentageFlag,
	BlockProposalFeeIterationsFlag,
	TaikoTokenAddressFlag,
})

func newProposer(c *cli.Context) (*proposer.Proposer, error) {
	if err := proposerConf.Validate(); err != nil {
		return nil, err
	}
	return proposer.New(c.Context, proposerConf)
}
