package main

import (
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/urfave/cli/v2"
)

const proposerCmd = "proposer"

var proposerConf = &proposer.Config{}

// Required flags used by proposer.
var (
	L1ProposerPrivKey = &cli.StringFlag{
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
	L2SuggestedFeeRecipient = &cli.StringFlag{
		Name:     "l2.suggestedFeeRecipient",
		Usage:    "Address of the proposed block's suggested fee recipient",
		Required: true,
		Category: proposerCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.L2SuggestedFeeRecipient = common.HexToAddress(v)
			return nil
		},
	}
	ProverEndpoints = &cli.StringSliceFlag{
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
	BlockProposalFee = &cli.StringFlag{
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
	TaikoTokenAddress = &cli.StringFlag{
		Name:     "taikoToken",
		Usage:    "TaikoToken contract address",
		Required: true,
		Category: proposerCategory,
		Action: func(c *cli.Context, v string) error {
			proposerConf.TaikoTokenAddress = common.HexToAddress(v)
			endpointConf.TaikoTokenAddress = common.HexToAddress(v)
			return nil
		},
	}
)

// Optional flags used by proposer.
var (
	ProposeInterval = &cli.DurationFlag{
		Name:     "proposeInterval",
		Usage:    "Time interval to propose L2 pending transactions",
		Category: proposerCategory,
		Action: func(c *cli.Context, v time.Duration) error {
			proposerConf.ProposeInterval = &v
			return nil
		},
	}
	TxPoolLocals = &cli.StringSliceFlag{
		Name:     "txpool.locals",
		Usage:    "Comma separated accounts to treat as locals (priority inclusion)",
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
	TxPoolLocalsOnly = &cli.BoolFlag{
		Name:     "txpool.localsOnly",
		Usage:    "If set to true, proposer will only propose transactions of local accounts",
		Value:    false,
		Category: proposerCategory,
		Action: func(c *cli.Context, v bool) error {
			proposerConf.LocalAddressesOnly = v
			return nil
		},
	}
	ProposeEmptyBlocksInterval = &cli.DurationFlag{
		Name:     "proposeEmptyBlockInterval",
		Usage:    "Time interval to propose empty blocks",
		Category: proposerCategory,
		Action: func(c *cli.Context, v time.Duration) error {
			proposerConf.ProposeEmptyBlocksInterval = &v
			return nil
		},
	}
	MaxProposedTxListsPerEpoch = &cli.Uint64Flag{
		Name:     "maxProposedTxListsPerEpoch",
		Value:    1,
		Category: proposerCategory,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.MaxProposedTxListsPerEpoch = v
			return nil
		},
	}
	ProposeBlockTxGasLimit = &cli.Uint64Flag{
		Name:     "proposeBlockTxGasLimit",
		Usage:    "Gas limit will be used for TaikoL1.proposeBlock transactions",
		Category: proposerCategory,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.ProposeBlockTxGasLimit = &v
			return nil
		},
	}
	ProposeBlockTxReplacementMultiplier = &cli.Uint64Flag{
		Name:     "proposeBlockTxReplacementMultiplier",
		Value:    2,
		Usage:    "Gas tip multiplier when replacing a TaikoL1.proposeBlock transaction with same nonce",
		Category: proposerCategory,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.ProposeBlockTxReplacementMultiplier = v
			return nil
		},
	}
	ProposeBlockTxGasTipCap = &cli.Uint64Flag{
		Name:     "proposeBlockTxGasTipCap",
		Usage:    "Gas tip cap (in wei) for a TaikoL1.proposeBlock transaction when doing the transaction replacement",
		Category: proposerCategory,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.ProposeBlockTxGasTipCap = new(big.Int).SetUint64(v)
			return nil
		},
	}
	BlockProposalFeeIncreasePercentage = &cli.Uint64Flag{
		Name:     "blockProposalFeeIncreasePercentage",
		Usage:    "Increase fee by what percentage when no prover wants to accept the block at initial fee",
		Category: proposerCategory,
		Value:    10,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.BlockProposalFeeIncreasePercentage = new(big.Int).SetUint64(v)
			return nil
		},
	}
	BlockProposalFeeIterations = &cli.Uint64Flag{
		Name:     "blockProposalFeeIterations",
		Usage:    "If nobody accepts block at initial fee, how many iterations to increase fee before giving up",
		Category: proposerCategory,
		Value:    3,
		Action: func(c *cli.Context, v uint64) error {
			proposerConf.BlockProposalFeeIterations = v
			return nil
		},
	}
)

// All proposer flags.
var proposerFlags = MergeFlags(CommonFlags, []cli.Flag{
	L2HTTPEndpoint,
	L1ProposerPrivKey,
	L2SuggestedFeeRecipient,
	ProposeInterval,
	TxPoolLocals,
	TxPoolLocalsOnly,
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

func prepareProposer(c *cli.Context, ep *rpc.Client) (p *proposer.Proposer, err error) {
	return proposer.New(c.Context, ep, proposerConf)
}
