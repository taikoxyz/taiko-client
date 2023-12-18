package prover

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Taiko prover.
type Config struct {
	L1WsEndpoint                            string
	L1HttpEndpoint                          string
	L2WsEndpoint                            string
	L2HttpEndpoint                          string
	TaikoL1Address                          common.Address
	TaikoL2Address                          common.Address
	TaikoTokenAddress                       common.Address
	AssignmentHookAddress                   common.Address
	L1ProverPrivKey                         *ecdsa.PrivateKey
	ZKEvmRpcdEndpoint                       string
	ZkEvmRpcdParamsPath                     string
	StartingBlockID                         *big.Int
	Dummy                                   bool
	GuardianProverAddress                   common.Address
	GuardianProofSubmissionDelay            time.Duration
	ProofSubmissionMaxRetry                 uint64
	Graffiti                                string
	BackOffMaxRetrys                        uint64
	BackOffRetryInterval                    time.Duration
	ProveUnassignedBlocks                   bool
	ContesterMode                           bool
	RPCTimeout                              *time.Duration
	WaitReceiptTimeout                      time.Duration
	ProveBlockGasLimit                      *uint64
	ProveBlockTxReplacementMultiplier       uint64
	ProveBlockMaxTxGasTipCap                *big.Int
	HTTPServerPort                          uint64
	Capacity                                uint64
	MinOptimisticTierFee                    *big.Int
	MinSgxTierFee                           *big.Int
	MinPseZkevmTierFee                      *big.Int
	MinSgxAndPseZkevmTierFee                *big.Int
	MaxExpiry                               time.Duration
	MaxProposedIn                           uint64
	MaxBlockSlippage                        uint64
	DatabasePath                            string
	DatabaseCacheSize                       uint64
	Allowance                               *big.Int
	GuardianProverHealthCheckServerEndpoint *url.URL
	RaikoHostEndpoint                       string
}

// NewConfigFromCliContext creates a new config instance from command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProverPrivKeyStr := c.String(flags.L1ProverPrivKey.Name)

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(l1ProverPrivKeyStr))
	if err != nil {
		return nil, fmt.Errorf("invalid L1 prover private key: %w", err)
	}

	var startingBlockID *big.Int
	if c.IsSet(flags.StartingBlockID.Name) {
		startingBlockID = new(big.Int).SetUint64(c.Uint64(flags.StartingBlockID.Name))
	}

	var timeout *time.Duration
	if c.IsSet(flags.RPCTimeout.Name) {
		duration := c.Duration(flags.RPCTimeout.Name)
		timeout = &duration
	}

	var proveBlockTxGasLimit *uint64
	if c.IsSet(flags.ProveBlockTxGasLimit.Name) {
		gasLimit := c.Uint64(flags.ProveBlockTxGasLimit.Name)
		proveBlockTxGasLimit = &gasLimit
	}

	proveBlockTxReplacementMultiplier := c.Uint64(flags.ProveBlockTxReplacementMultiplier.Name)
	if proveBlockTxReplacementMultiplier == 0 {
		return nil, fmt.Errorf(
			"invalid --proveBlockTxReplacementMultiplier value: %d",
			proveBlockTxReplacementMultiplier,
		)
	}

	var proveBlockMaxTxGasTipCap *big.Int
	if c.IsSet(flags.ProveBlockMaxTxGasTipCap.Name) {
		proveBlockMaxTxGasTipCap = new(big.Int).SetUint64(c.Uint64(flags.ProveBlockMaxTxGasTipCap.Name))
	}

	var allowance *big.Int = common.Big0
	if c.IsSet(flags.Allowance.Name) {
		amt, ok := new(big.Int).SetString(c.String(flags.Allowance.Name), 10)
		if !ok {
			return nil, fmt.Errorf("invalid setting allowance config value: %v", c.String(flags.Allowance.Name))
		}

		allowance = amt
	}

	var guardianProverHealthCheckServerEndpoint *url.URL
	if c.IsSet(flags.GuardianProverHealthCheckServerEndpoint.Name) {
		guardianProverHealthCheckServerEndpoint, err = url.Parse(c.String(flags.GuardianProverHealthCheckServerEndpoint.Name))
		if err != nil {
			return nil, err
		}
	}

	if !c.IsSet(c.String(flags.GuardianProver.Name)) &&
		!c.IsSet(c.String(flags.RaikoHostEndpoint.Name)) {
		return nil, fmt.Errorf("raiko host not provided")
	}

	return &Config{
		L1WsEndpoint:                            c.String(flags.L1WSEndpoint.Name),
		L1HttpEndpoint:                          c.String(flags.L1HTTPEndpoint.Name),
		L2WsEndpoint:                            c.String(flags.L2WSEndpoint.Name),
		L2HttpEndpoint:                          c.String(flags.L2HTTPEndpoint.Name),
		TaikoL1Address:                          common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
		TaikoL2Address:                          common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
		TaikoTokenAddress:                       common.HexToAddress(c.String(flags.TaikoTokenAddress.Name)),
		AssignmentHookAddress:                   common.HexToAddress(c.String(flags.ProverAssignmentHookAddress.Name)),
		L1ProverPrivKey:                         l1ProverPrivKey,
		ZKEvmRpcdEndpoint:                       c.String(flags.ZkEvmRpcdEndpoint.Name),
		ZkEvmRpcdParamsPath:                     c.String(flags.ZkEvmRpcdParamsPath.Name),
		RaikoHostEndpoint:                       c.String(flags.RaikoHostEndpoint.Name),
		StartingBlockID:                         startingBlockID,
		Dummy:                                   c.Bool(flags.Dummy.Name),
		GuardianProverAddress:                   common.HexToAddress(c.String(flags.GuardianProver.Name)),
		GuardianProofSubmissionDelay:            c.Duration(flags.GuardianProofSubmissionDelay.Name),
		GuardianProverHealthCheckServerEndpoint: guardianProverHealthCheckServerEndpoint,
		ProofSubmissionMaxRetry:                 c.Uint64(flags.ProofSubmissionMaxRetry.Name),
		Graffiti:                                c.String(flags.Graffiti.Name),
		BackOffMaxRetrys:                        c.Uint64(flags.BackOffMaxRetrys.Name),
		BackOffRetryInterval:                    c.Duration(flags.BackOffRetryInterval.Name),
		ProveUnassignedBlocks:                   c.Bool(flags.ProveUnassignedBlocks.Name),
		ContesterMode:                           c.Bool(flags.ContesterMode.Name),
		RPCTimeout:                              timeout,
		WaitReceiptTimeout:                      c.Duration(flags.WaitReceiptTimeout.Name),
		ProveBlockGasLimit:                      proveBlockTxGasLimit,
		Capacity:                                c.Uint64(flags.ProverCapacity.Name),
		ProveBlockTxReplacementMultiplier:       proveBlockTxReplacementMultiplier,
		ProveBlockMaxTxGasTipCap:                proveBlockMaxTxGasTipCap,
		HTTPServerPort:                          c.Uint64(flags.ProverHTTPServerPort.Name),
		MinOptimisticTierFee:                    new(big.Int).SetUint64(c.Uint64(flags.MinOptimisticTierFee.Name)),
		MinSgxTierFee:                           new(big.Int).SetUint64(c.Uint64(flags.MinSgxTierFee.Name)),
		MinPseZkevmTierFee:                      new(big.Int).SetUint64(c.Uint64(flags.MinPseZkevmTierFee.Name)),
		MinSgxAndPseZkevmTierFee:                new(big.Int).SetUint64(c.Uint64(flags.MinSgxAndPseZkevmTierFee.Name)),
		MaxExpiry:                               c.Duration(flags.MaxExpiry.Name),
		MaxBlockSlippage:                        c.Uint64(flags.MaxAcceptableBlockSlippage.Name),
		MaxProposedIn:                           c.Uint64(flags.MaxProposedIn.Name),
		DatabasePath:                            c.String(flags.DatabasePath.Name),
		DatabaseCacheSize:                       c.Uint64(flags.DatabaseCacheSize.Name),
		Allowance:                               allowance,
	}, nil
}
