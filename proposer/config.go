package proposer

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// Config contains all configurations to initialize a Taiko proposer.
type Config struct {
	*rpc.ClientConfig
	AssignmentHookAddress               common.Address
	L1ProposerPrivKey                   *ecdsa.PrivateKey
	L2SuggestedFeeRecipient             common.Address
	ExtraData                           string
	ProposeInterval                     time.Duration
	LocalAddresses                      []common.Address
	LocalAddressesOnly                  bool
	ProposeEmptyBlocksInterval          time.Duration
	MaxProposedTxListsPerEpoch          uint64
	ProposeBlockTxGasLimit              uint64
	ProposeBlockTxReplacementMultiplier uint64
	WaitReceiptTimeout                  time.Duration
	ProposeBlockTxGasTipCap             *big.Int
	ProverEndpoints                     []*url.URL
	OptimisticTierFee                   *big.Int
	SgxTierFee                          *big.Int
	TierFeePriceBump                    *big.Int
	MaxTierFeePriceBumps                uint64
	IncludeParentMetaHash               bool
	BlobAllowed                         bool
	L1BlockBuilderTip                   *big.Int
}

// NewConfigFromCliContext initializes a Config instance from
// command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProposerPrivKey, err := crypto.ToECDSA(
		common.Hex2Bytes(c.String(flags.L1ProposerPrivKey.Name)),
	)
	if err != nil {
		return nil, fmt.Errorf("invalid L1 proposer private key: %w", err)
	}

	l2SuggestedFeeRecipient := c.String(flags.L2SuggestedFeeRecipient.Name)
	if !common.IsHexAddress(l2SuggestedFeeRecipient) {
		return nil, fmt.Errorf("invalid L2 suggested fee recipient address: %s", l2SuggestedFeeRecipient)
	}

	var localAddresses []common.Address
	if c.IsSet(flags.TxPoolLocals.Name) {
		for _, account := range strings.Split(c.String(flags.TxPoolLocals.Name), ",") {
			if trimmed := strings.TrimSpace(account); !common.IsHexAddress(trimmed) {
				return nil, fmt.Errorf("invalid account in --txpool.locals: %s", trimmed)
			}
			localAddresses = append(localAddresses, common.HexToAddress(account))
		}
	}

	proposeBlockTxReplacementMultiplier := c.Uint64(flags.ProposeBlockTxReplacementMultiplier.Name)
	if proposeBlockTxReplacementMultiplier == 0 {
		return nil, fmt.Errorf(
			"invalid --proposeBlockTxReplacementMultiplier value: %d",
			proposeBlockTxReplacementMultiplier,
		)
	}

	var proposeBlockTxGasTipCap *big.Int
	if c.IsSet(flags.ProposeBlockTxGasTipCap.Name) {
		proposeBlockTxGasTipCap = new(big.Int).SetUint64(c.Uint64(flags.ProposeBlockTxGasTipCap.Name))
	}

	var proverEndpoints []*url.URL
	for _, e := range strings.Split(c.String(flags.ProverEndpoints.Name), ",") {
		endpoint, err := url.Parse(e)
		if err != nil {
			return nil, err
		}
		proverEndpoints = append(proverEndpoints, endpoint)
	}

	return &Config{
		ClientConfig: &rpc.ClientConfig{
			L1Endpoint:        c.String(flags.L1WSEndpoint.Name),
			L2Endpoint:        c.String(flags.L2HTTPEndpoint.Name),
			TaikoL1Address:    common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
			TaikoL2Address:    common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
			TaikoTokenAddress: common.HexToAddress(c.String(flags.TaikoTokenAddress.Name)),
			Timeout:           c.Duration(flags.RPCTimeout.Name),
		},
		AssignmentHookAddress:               common.HexToAddress(c.String(flags.ProposerAssignmentHookAddress.Name)),
		L1ProposerPrivKey:                   l1ProposerPrivKey,
		L2SuggestedFeeRecipient:             common.HexToAddress(l2SuggestedFeeRecipient),
		ExtraData:                           c.String(flags.ExtraData.Name),
		ProposeInterval:                     c.Duration(flags.ProposeInterval.Name),
		LocalAddresses:                      localAddresses,
		LocalAddressesOnly:                  c.Bool(flags.TxPoolLocalsOnly.Name),
		ProposeEmptyBlocksInterval:          c.Duration(flags.ProposeEmptyBlocksInterval.Name),
		MaxProposedTxListsPerEpoch:          c.Uint64(flags.MaxProposedTxListsPerEpoch.Name),
		ProposeBlockTxGasLimit:              c.Uint64(flags.ProposeBlockTxGasLimit.Name),
		ProposeBlockTxReplacementMultiplier: proposeBlockTxReplacementMultiplier,
		WaitReceiptTimeout:                  c.Duration(flags.WaitReceiptTimeout.Name),
		ProposeBlockTxGasTipCap:             proposeBlockTxGasTipCap,
		ProverEndpoints:                     proverEndpoints,
		OptimisticTierFee:                   new(big.Int).SetUint64(c.Uint64(flags.OptimisticTierFee.Name)),
		SgxTierFee:                          new(big.Int).SetUint64(c.Uint64(flags.SgxTierFee.Name)),
		TierFeePriceBump:                    new(big.Int).SetUint64(c.Uint64(flags.TierFeePriceBump.Name)),
		MaxTierFeePriceBumps:                c.Uint64(flags.MaxTierFeePriceBumps.Name),
		IncludeParentMetaHash:               c.Bool(flags.ProposeBlockIncludeParentMetaHash.Name),
		BlobAllowed:                         c.Bool(flags.BlobAllowed.Name),
		L1BlockBuilderTip:                   new(big.Int).SetUint64(c.Uint64(flags.L1BlockBuilderTip.Name)),
	}, nil
}

func NewConfigFromConfigFile(c *cli.Context, path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, fmt.Errorf("error loading .env config: %w", err)
	}

	timeout, err := time.ParseDuration(os.Getenv("RPC_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("error parsing RPC_TIMEOUT: %w", err)
	}

	l1ProposerPrivKey, err := crypto.ToECDSA(
		common.FromHex(os.Getenv("L1_PROPOSER_PRIVATE_KEY")),
	)
	if err != nil {
		return nil, fmt.Errorf("config invalid L1 proposer private key: %w, %s", err, os.Getenv("L1_PROPOSER_PRIVATE_KEY"))
	}

	l2SuggestedFeeRecipient := os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")
	if !common.IsHexAddress(l2SuggestedFeeRecipient) {
		return nil, fmt.Errorf("invalid L2 suggested fee recipient address: %s", l2SuggestedFeeRecipient)
	}

	proposeInterval, err := time.ParseDuration(os.Getenv("PROPOSE_INTERVAL"))
	if err != nil {
		return nil, fmt.Errorf("error setting propose interval: %w", err)
	}

	var localAddresses []common.Address
	localsOnly, err := strconv.ParseBool(os.Getenv("LOCAL_ADDRESSES_ONLY"))
	if err != nil {
		return nil, fmt.Errorf("error loading local_addresses_only: %w", err)
	}
	if localsOnly {
		for _, account := range strings.Split(os.Getenv("TXPOOL_LOCALS"), ",") {
			if trimmed := strings.TrimSpace(account); !common.IsHexAddress(trimmed) {
				return nil, fmt.Errorf("invalid account in --txpool.locals: %s", trimmed)
			}
			localAddresses = append(localAddresses, common.HexToAddress(account))
		}
	}

	proposeEmptyBlocksInteval, err := time.ParseDuration(os.Getenv("PROPOSE_EMPTY_BLOCKS_INTERVAL"))
	if err != nil {
		return nil, fmt.Errorf("error setting propose empty blocks interval: %w", err)
	}

	maxProposedTxListsPerEpoch, err := strconv.ParseUint(os.Getenv("MAX_PROPOSED_TX_LISTS_PER_EPOCH"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error setting max proposed tx lists per epoch: %w", err)
	}

	var proposeBlockTxGasLimit uint64
	gasLimit := os.Getenv("PROPOSE_BLOCK_TX_GAS_LIMIT")
	if gasLimit != "-1" {
		proposeBlockTxGasLimit, _ = strconv.ParseUint(gasLimit, 0, 64)
	} else {
		proposeBlockTxGasLimit = 0
	}

	proposeBlockTxReplacementMultiplier, err := strconv.ParseUint(
		os.Getenv("PROPOSE_BLOCK_TX_REPLACEMENT_MULTIPLIER"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error converting proposeBlockTxReplacementMultiplier: %w", err)
	}
	if proposeBlockTxReplacementMultiplier == 0 {
		return nil, fmt.Errorf(
			"invalid --proposeBlockTxReplacementMultiplier value: %d",
			proposeBlockTxReplacementMultiplier,
		)
	}

	waitReceiptTimeout, err := time.ParseDuration(os.Getenv("WAIT_RECEIPT_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	var proposeBlockTxGasTipCap *big.Int
	isSet := os.Getenv("PROPOSE_BLOCK_TX_GAS_TIP_CAP")
	if isSet != "" {
		tmp, _ := strconv.ParseUint(os.Getenv("PROPOSE_BLOCK_TX_GAS_TIP_CAP"), 0, 64)
		proposeBlockTxGasTipCap = new(big.Int).SetUint64(tmp)
	}

	var proverEndpoints []*url.URL
	for _, e := range strings.Split(os.Getenv("PROVER_ENDPOINTS"), ",") {
		endpoint, err := url.Parse(e)
		if err != nil {
			return nil, err
		}
		proverEndpoints = append(proverEndpoints, endpoint)
	}

	var optimisticTierFee *big.Int
	optimistic := os.Getenv("OPTIMISTIC_TIER_FEE")
	if optimistic != "" {
		tmp, _ := strconv.ParseUint(optimistic, 0, 64)
		optimisticTierFee = new(big.Int).SetUint64(tmp)
	} else {
		optimisticTierFee = common.Big0
	}

	var sgxTierFee *big.Int
	sgx := os.Getenv("SGX_TIER_FEE")
	if sgx != "" {
		tmp, _ := strconv.ParseUint(sgx, 0, 64)
		sgxTierFee = new(big.Int).SetUint64(tmp)
	} else {
		sgxTierFee = common.Big0
	}

	tierFee, err := strconv.ParseUint(os.Getenv("TIER_FEE_PRICE_BUMP"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error converting tierFeePriceBump: %w", err)
	}
	tierFeePriceBump := new(big.Int).SetUint64(tierFee)

	maxTierFeePriceBumps, err := strconv.ParseUint(os.Getenv("MAX_TIER_FEE_PRICE_BUMPS"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error converting maxTierFeePriceBumps: %w", err)
	}

	includeMetaHash, err := strconv.ParseBool(os.Getenv("INCLUDE_PARENT_METAHASH"))
	if err != nil {
		return nil, fmt.Errorf("error converting includeParentMetahash: %w", err)
	}

	blobAllowed, err := strconv.ParseBool(os.Getenv("BLOB_ALLOWED"))
	if err != nil {
		return nil, fmt.Errorf("error converting blobAllowed: %w", err)
	}

	tip, err := strconv.ParseUint(os.Getenv("L1_BLOCK_BUILDER_TIP"), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("error converting l1BlockBuilderTip: %w", err)
	}
	l1BlockBuilderTip := new(big.Int).SetUint64(tip)

	return &Config{
		ClientConfig: &rpc.ClientConfig{
			L1Endpoint:        os.Getenv("L1_NODE_WS_ENDPOINT"),
			L2Endpoint:        os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
			TaikoL1Address:    common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
			TaikoL2Address:    common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
			TaikoTokenAddress: common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
			Timeout:           timeout,
		},
		AssignmentHookAddress:               common.HexToAddress(os.Getenv("ASSIGNMENT_HOOK_ADDRESS")),
		L1ProposerPrivKey:                   l1ProposerPrivKey,
		L2SuggestedFeeRecipient:             common.HexToAddress(l2SuggestedFeeRecipient),
		ExtraData:                           c.String(flags.ExtraData.Name),
		ProposeInterval:                     proposeInterval,
		LocalAddresses:                      localAddresses,
		LocalAddressesOnly:                  localsOnly,
		ProposeEmptyBlocksInterval:          proposeEmptyBlocksInteval,
		MaxProposedTxListsPerEpoch:          maxProposedTxListsPerEpoch,
		ProposeBlockTxGasLimit:              proposeBlockTxGasLimit,
		ProposeBlockTxReplacementMultiplier: proposeBlockTxReplacementMultiplier,
		WaitReceiptTimeout:                  waitReceiptTimeout,
		ProposeBlockTxGasTipCap:             proposeBlockTxGasTipCap,
		ProverEndpoints:                     proverEndpoints,
		OptimisticTierFee:                   optimisticTierFee,
		SgxTierFee:                          sgxTierFee,
		TierFeePriceBump:                    tierFeePriceBump,
		MaxTierFeePriceBumps:                maxTierFeePriceBumps,
		IncludeParentMetaHash:               includeMetaHash,
		BlobAllowed:                         blobAllowed,
		L1BlockBuilderTip:                   l1BlockBuilderTip,
	}, nil
}
