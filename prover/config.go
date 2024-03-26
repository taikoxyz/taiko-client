package prover

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-client/cmd/flags"
	pkgFlags "github.com/taikoxyz/taiko-client/pkg/flags"
)

// Config contains the configurations to initialize a Taiko prover.
type Config struct {
	L1WsEndpoint                            string
	L1HttpEndpoint                          string
	L1BeaconEndpoint                        string
	L2WsEndpoint                            string
	L2HttpEndpoint                          string
	TaikoL1Address                          common.Address
	TaikoL2Address                          common.Address
	TaikoTokenAddress                       common.Address
	AssignmentHookAddress                   common.Address
	L1ProverPrivKey                         *ecdsa.PrivateKey
	StartingBlockID                         *big.Int
	Dummy                                   bool
	GuardianProverAddress                   common.Address
	GuardianProofSubmissionDelay            time.Duration
	Graffiti                                string
	BackOffMaxRetrys                        uint64
	BackOffRetryInterval                    time.Duration
	ProveUnassignedBlocks                   bool
	ContesterMode                           bool
	EnableLivenessBondProof                 bool
	RPCTimeout                              time.Duration
	WaitReceiptTimeout                      time.Duration
	HTTPServerPort                          uint64
	Capacity                                uint64
	MinOptimisticTierFee                    *big.Int
	MinSgxTierFee                           *big.Int
	MinSgxAndZkVMTierFee                    *big.Int
	MinEthBalance                           *big.Int
	MinTaikoTokenBalance                    *big.Int
	MaxExpiry                               time.Duration
	MaxProposedIn                           uint64
	MaxBlockSlippage                        uint64
	Allowance                               *big.Int
	GuardianProverHealthCheckServerEndpoint *url.URL
	RaikoHostEndpoint                       string
	L1NodeVersion                           string
	L2NodeVersion                           string
	BlockConfirmations                      uint64
	TxmgrConfigs                            *txmgr.CLIConfig
}

// NewConfigFromCliContext creates a new config instance from command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProverPrivKey, err := crypto.ToECDSA(common.FromHex(c.String(flags.L1ProverPrivKey.Name)))
	if err != nil {
		return nil, fmt.Errorf("invalid L1 prover private key: %w", err)
	}

	if !c.IsSet(flags.L1BeaconEndpoint.Name) {
		return nil, errors.New("empty L1 beacon endpoint")
	}

	var startingBlockID *big.Int
	if c.IsSet(flags.StartingBlockID.Name) {
		startingBlockID = new(big.Int).SetUint64(c.Uint64(flags.StartingBlockID.Name))
	}

	var allowance = common.Big0
	if c.IsSet(flags.Allowance.Name) {
		amt, ok := new(big.Int).SetString(c.String(flags.Allowance.Name), 10)
		if !ok {
			return nil, fmt.Errorf("invalid setting allowance config value: %v", c.String(flags.Allowance.Name))
		}

		allowance = amt
	}

	var guardianProverHealthCheckServerEndpoint *url.URL
	if c.IsSet(flags.GuardianProverHealthCheckServerEndpoint.Name) {
		if guardianProverHealthCheckServerEndpoint, err = url.Parse(
			c.String(flags.GuardianProverHealthCheckServerEndpoint.Name),
		); err != nil {
			return nil, err
		}
	}

	// If we are running a guardian prover, we need to prove unassigned blocks and run in contester mode by default.
	if c.IsSet(flags.GuardianProver.Name) {
		if err := c.Set(flags.ProveUnassignedBlocks.Name, "true"); err != nil {
			return nil, err
		}

		if err := c.Set(flags.ContesterMode.Name, "true"); err != nil {
			return nil, err
		}

		// l1 and l2 node version flags are required only if guardian prover
		if !c.IsSet(flags.L1NodeVersion.Name) {
			return nil, errors.New("L1NodeVersion is required if guardian prover is set")
		}

		if !c.IsSet(flags.L2NodeVersion.Name) {
			return nil, errors.New("L2NodeVersion is required if guardian prover is set")
		}
	}

	if !c.IsSet(flags.GuardianProver.Name) && !c.IsSet(flags.RaikoHostEndpoint.Name) {
		return nil, fmt.Errorf("raiko host not provided")
	}

	return &Config{
		L1WsEndpoint:                            c.String(flags.L1WSEndpoint.Name),
		L1HttpEndpoint:                          c.String(flags.L1HTTPEndpoint.Name),
		L1BeaconEndpoint:                        c.String(flags.L1BeaconEndpoint.Name),
		L2WsEndpoint:                            c.String(flags.L2WSEndpoint.Name),
		L2HttpEndpoint:                          c.String(flags.L2HTTPEndpoint.Name),
		TaikoL1Address:                          common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
		TaikoL2Address:                          common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
		TaikoTokenAddress:                       common.HexToAddress(c.String(flags.TaikoTokenAddress.Name)),
		AssignmentHookAddress:                   common.HexToAddress(c.String(flags.ProverAssignmentHookAddress.Name)),
		L1ProverPrivKey:                         l1ProverPrivKey,
		RaikoHostEndpoint:                       c.String(flags.RaikoHostEndpoint.Name),
		StartingBlockID:                         startingBlockID,
		Dummy:                                   c.Bool(flags.Dummy.Name),
		GuardianProverAddress:                   common.HexToAddress(c.String(flags.GuardianProver.Name)),
		GuardianProofSubmissionDelay:            c.Duration(flags.GuardianProofSubmissionDelay.Name),
		GuardianProverHealthCheckServerEndpoint: guardianProverHealthCheckServerEndpoint,
		Graffiti:                                c.String(flags.Graffiti.Name),
		BackOffMaxRetrys:                        c.Uint64(flags.BackOffMaxRetrys.Name),
		BackOffRetryInterval:                    c.Duration(flags.BackOffRetryInterval.Name),
		ProveUnassignedBlocks:                   c.Bool(flags.ProveUnassignedBlocks.Name),
		ContesterMode:                           c.Bool(flags.ContesterMode.Name),
		EnableLivenessBondProof:                 c.Bool(flags.EnableLivenessBondProof.Name),
		RPCTimeout:                              c.Duration(flags.RPCTimeout.Name),
		WaitReceiptTimeout:                      c.Duration(flags.WaitReceiptTimeout.Name),
		Capacity:                                c.Uint64(flags.ProverCapacity.Name),
		HTTPServerPort:                          c.Uint64(flags.ProverHTTPServerPort.Name),
		MinOptimisticTierFee:                    new(big.Int).SetUint64(c.Uint64(flags.MinOptimisticTierFee.Name)),
		MinSgxTierFee:                           new(big.Int).SetUint64(c.Uint64(flags.MinSgxTierFee.Name)),
		MinSgxAndZkVMTierFee:                    new(big.Int).SetUint64(c.Uint64(flags.MinSgxAndZkVMTierFee.Name)),
		MinEthBalance:                           new(big.Int).SetUint64(c.Uint64(flags.MinEthBalance.Name)),
		MinTaikoTokenBalance:                    new(big.Int).SetUint64(c.Uint64(flags.MinTaikoTokenBalance.Name)),
		MaxExpiry:                               c.Duration(flags.MaxExpiry.Name),
		MaxBlockSlippage:                        c.Uint64(flags.MaxAcceptableBlockSlippage.Name),
		MaxProposedIn:                           c.Uint64(flags.MaxProposedIn.Name),
		Allowance:                               allowance,
		L1NodeVersion:                           c.String(flags.L1NodeVersion.Name),
		L2NodeVersion:                           c.String(flags.L2NodeVersion.Name),
		BlockConfirmations:                      c.Uint64(flags.BlockConfirmations.Name),
		TxmgrConfigs: pkgFlags.InitTxmgrConfigsFromCli(
			c.String(flags.L1HTTPEndpoint.Name),
			l1ProverPrivKey,
			c,
		),
	}, nil
}

func NewConfigFromConfigFile(c *cli.Context, path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, fmt.Errorf("error loading .env config: %w", err)
	}

	if os.Getenv("L1_BEACON_ENDPOINT") == "" {
		return nil, errors.New("empty L1 beacon endpoint")
	}

	l1ProverPrivKey, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	if err != nil {
		return nil, fmt.Errorf("invalid L1 prover private key: %w", err)
	}

	var startingBlockID *big.Int
	if os.Getenv("STARTING_BLOCK_ID") != "" {
		id, err := strconv.ParseUint(os.Getenv("STARTING_BLOCK_ID"), 0, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing STARTING_BLOCK_ID %w", err)
		}
		startingBlockID = new(big.Int).SetUint64(id)
	}

	dummy, err := strconv.ParseBool(os.Getenv("DUMMY"))
	if err != nil {
		return nil, fmt.Errorf("error parsing DUMMY %w", err)
	}

	guardianProofSubDelay, err := time.ParseDuration(os.Getenv("GUARDIAN_PROOF_SUB_DELAY"))
	if err != nil {
		return nil, fmt.Errorf("error parsing GUARDIAN_PROOF_SUB_DELAY: %w", err)
	}

	var guardianProverHealthCheckServerEndpoint *url.URL
	if os.Getenv("GUARDIAN_PROVER_HEALTH_PORT") != "" {
		if guardianProverHealthCheckServerEndpoint, err = url.Parse(
			os.Getenv("GUARDIAN_PROVER_HEALTH_PORT"),
		); err != nil {
			return nil, err
		}
	}

	backoffMaxRetry, err := strconv.ParseUint(os.Getenv("BACKOFF_MAX_RETRY"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing BACKOFF_MAX_RETRY: %w", err)
	}
	retryInterval, err := time.ParseDuration(os.Getenv("RETRY_INTERVAL"))
	if err != nil {
		return nil, fmt.Errorf("error parsing RETRY_INTERVAL: %w", err)
	}
	livenessBond, err := strconv.ParseBool(os.Getenv("ENABLE_LIVENESS_BOND_PROOF"))
	if err != nil {
		return nil, fmt.Errorf("error parsing ENABLE_LIVENESS_BOND_PROOF: %w", err)
	}
	timeout, err := time.ParseDuration(os.Getenv("RPC_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("error parsing RPC_TIMEOUT: %w", err)
	}
	c.Set(flags.RPCTimeout.Name, timeout.String())
	waitReceiptTimeout, err := time.ParseDuration(os.Getenv("WAIT_RECEIPT_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("error parsing WAIT_RECEIPT_TIMEOUT: %w", err)
	}

	capacity, err := strconv.ParseUint(os.Getenv("PROVER_CAPACITY"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing PROVER_CAPACITY: %w", err)
	}

	serverPort, err := strconv.ParseUint(os.Getenv("PROVER_HTTP_PORT"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing PROVER_HTTP_PORT %w", err)
	}

	var (
		minOptimisticTierFee *big.Int
		minSgxTierFee        *big.Int
		minSgxAndZkVMTierFee *big.Int
		minEthBalance        *big.Int
		minTaikoTokenBalance *big.Int
	)
	op, err := strconv.ParseUint(os.Getenv("MIN_OPTIMISTIC_TIER_FEE"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing MIN_OPTIMISTIC_TIER_FEE %w", err)
	}
	minOptimisticTierFee = new(big.Int).SetUint64(op)
	sgx, err := strconv.ParseUint(os.Getenv("MIN_SGX_TIER_FEE"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing MIN_SGX_TIER_FEE %w", err)
	}
	minSgxTierFee = new(big.Int).SetUint64(sgx)
	sgxZKVM, err := strconv.ParseUint(os.Getenv("MIN_SGX_ZKVM_TIER_FEE"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing MIN_SGX_ZKVM_TIER_FEE %w", err)
	}
	minSgxAndZkVMTierFee = new(big.Int).SetUint64(sgxZKVM)

	eth, err := strconv.ParseUint(os.Getenv("MIN_ETH_BALANCE"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing MIN_ETH_BALANCE %w", err)
	}
	minEthBalance = new(big.Int).SetUint64(eth)

	tko, err := strconv.ParseUint(os.Getenv("MIN_TAIKO_TOKEN_BALANCE"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing MIN_TAIKO_TOKEN_BALANCE %w", err)
	}
	minTaikoTokenBalance = new(big.Int).SetUint64(tko)

	maxExpiry, err := time.ParseDuration(os.Getenv("MAX_EXPIRY"))
	if err != nil {
		return nil, fmt.Errorf("error parsing MAX_EXPIRY: %w", err)
	}

	maxSlippage, err := strconv.ParseUint(os.Getenv("MAX_BLOCK_SLIPPAGE"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing MAX_BLOCK_SLIPPAGE: %w", err)
	}
	maxProposedIn, err := strconv.ParseUint(os.Getenv("MAX_PROPOSED_IN"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing MAX_PROPOSED_IN: %w", err)
	}

	var allowance = common.Big0
	if os.Getenv("TOKEN_ALLOWANCE") != "" {
		amt, ok := new(big.Int).SetString(os.Getenv("TOKEN_ALLOWANCE"), 10)
		if !ok {
			return nil, fmt.Errorf("invalid setting allowance config value: %v", os.Getenv("TOKEN_ALLOWANCE"))
		}
		allowance = amt
	}

	blockConfirmations, err := strconv.ParseUint(os.Getenv("BLOCK_CONFIRMATIONS"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing BLOCK_CONFIRMATIONS: %w", err)
	}

	var (
		proveUnassignedBlocks bool
		contesterMode         bool
	)
	proveUnassignedBlocks, err = strconv.ParseBool(os.Getenv("PROVE_UNASSIGNED_BLOCKS"))
	if err != nil {
		return nil, fmt.Errorf("error parsing PROVE_UNASSIGNED_BLOCKS %w", err)
	}
	contesterMode, err = strconv.ParseBool(os.Getenv("CONTESTER_MODE"))
	if err != nil {
		return nil, fmt.Errorf("error parsing CONTESTER_MODE %w", err)
	}

	// If we are running a guardian prover, we need to prove unassigned blocks and run in contester mode by default.
	if os.Getenv("GUARDIAN_PROVER_CONTRACT_ADDRESS") != "" {
		proveUnassignedBlocks = true
		contesterMode = true

		// l1 and l2 node version flags are required only if guardian prover
		if os.Getenv("L1_NODE_VERSION") == "" {
			return nil, errors.New("L1NodeVersion is required if guardian prover is set")
		}

		if os.Getenv("L2_NODE_VERSION") == "" {
			return nil, errors.New("L2NodeVersion is required if guardian prover is set")
		}
	}

	if os.Getenv("GUARDIAN_PROVER_CONTRACT_ADDRESS") == "" && os.Getenv("SGX_RAIKO_HOST") == "" {
		return nil, fmt.Errorf("raiko host not provided")
	}

	return &Config{
		L1WsEndpoint:                            os.Getenv("L1_NODE_WS_ENDPOINT"),
		L1HttpEndpoint:                          os.Getenv("L1_NODE_HTTP_ENDPOINT"),
		L1BeaconEndpoint:                        os.Getenv("L1_BEACON_ENDPOINT"),
		L2WsEndpoint:                            os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		L2HttpEndpoint:                          os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
		TaikoL1Address:                          common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:                          common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		TaikoTokenAddress:                       common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		AssignmentHookAddress:                   common.HexToAddress(os.Getenv("ASSIGNMENT_HOOK_ADDRESS")),
		L1ProverPrivKey:                         l1ProverPrivKey,
		RaikoHostEndpoint:                       os.Getenv("SGX_RAIKO_HOST"),
		StartingBlockID:                         startingBlockID,
		Dummy:                                   dummy,
		GuardianProverAddress:                   common.HexToAddress(os.Getenv("GUARDIAN_PROVER_CONTRACT_ADDRESS")),
		GuardianProofSubmissionDelay:            guardianProofSubDelay,
		GuardianProverHealthCheckServerEndpoint: guardianProverHealthCheckServerEndpoint,
		Graffiti:                                os.Getenv("GRAFFITI"),
		BackOffMaxRetrys:                        backoffMaxRetry,
		BackOffRetryInterval:                    retryInterval,
		ProveUnassignedBlocks:                   proveUnassignedBlocks,
		ContesterMode:                           contesterMode,
		EnableLivenessBondProof:                 livenessBond,
		RPCTimeout:                              timeout,
		WaitReceiptTimeout:                      waitReceiptTimeout,
		Capacity:                                capacity,
		HTTPServerPort:                          serverPort,
		MinOptimisticTierFee:                    minOptimisticTierFee,
		MinSgxTierFee:                           minSgxTierFee,
		MinSgxAndZkVMTierFee:                    minSgxAndZkVMTierFee,
		MinEthBalance:                           minEthBalance,
		MinTaikoTokenBalance:                    minTaikoTokenBalance,
		MaxExpiry:                               maxExpiry,
		MaxBlockSlippage:                        maxSlippage,
		MaxProposedIn:                           maxProposedIn,
		Allowance:                               allowance,
		L1NodeVersion:                           os.Getenv("L1_NODE_VERSION"),
		L2NodeVersion:                           os.Getenv("L2_NODE_VERSION"),
		BlockConfirmations:                      blockConfirmations,
		TxmgrConfigs: pkgFlags.InitTxmgrConfigsFromCli(
			os.Getenv("L1_NODE_HTTP_ENDPOINT"),
			l1ProverPrivKey,
			c,
		),
	}, nil
}
