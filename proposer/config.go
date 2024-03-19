package proposer

import (
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// Config містить всі конфігурації для ініціалізації Taiko proposer.
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

// NewConfigFromCliContext ініціалізує екземпляр Config з командних рядків.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(c.String(flags.L1ProposerPrivKey.Name)))
	if err != nil {
		return nil, fmt.Errorf("неправильний приватний ключ L1 proposer: %w", err)
	}

	l2SuggestedFeeRecipient := common.HexToAddress(c.String(flags.L2SuggestedFeeRecipient.Name))

	var localAddresses []common.Address
	if c.IsSet(flags.TxPoolLocals.Name) {
		for _, account := range strings.Split(c.String(flags.TxPoolLocals.Name), ",") {
			trimmed := strings.TrimSpace(account)
			if !common.IsHexAddress(trimmed) {
				return nil, fmt.Errorf("неправильний рахунок в --txpool.locals: %s", trimmed)
			}
			localAddresses = append(localAddresses, common.HexToAddress(trimmed))
		}
	}

	proposeBlockTxReplacementMultiplier := c.Uint64(flags.ProposeBlockTxReplacementMultiplier.Name)
	if proposeBlockTxReplacementMultiplier == 0 {
		return nil, fmt.Errorf("неправильне значення --proposeBlockTxReplacementMultiplier: %d", proposeBlockTxReplacementMultiplier)
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
		L2SuggestedFeeRecipient:             l2SuggestedFeeRecipient,
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

