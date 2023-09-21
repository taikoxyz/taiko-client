package proposer

import (
	"crypto/ecdsa"
	"math/big"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Config contains all configurations to initialize a Taiko proposer.
type Config struct {
	L1Endpoint                          string
	L2Endpoint                          string
	TaikoL1Address                      common.Address
	TaikoL2Address                      common.Address
	TaikoTokenAddress                   common.Address
	L1ProposerPrivKey                   *ecdsa.PrivateKey
	L2SuggestedFeeRecipient             common.Address
	ProposeInterval                     *time.Duration
	LocalAddresses                      []common.Address
	LocalAddressesOnly                  bool
	ProposeEmptyBlocksInterval          *time.Duration
	MaxProposedTxListsPerEpoch          uint64
	ProposeBlockTxGasLimit              *uint64
	BackOffRetryInterval                time.Duration
	ProposeBlockTxReplacementMultiplier uint64
	RPCTimeout                          *time.Duration
	WaitReceiptTimeout                  time.Duration
	ProposeBlockTxGasTipCap             *big.Int
	ProverEndpoints                     []*url.URL
	BlockProposalFee                    *big.Int
	BlockProposalFeeIncreasePercentage  *big.Int
	BlockProposalFeeIterations          uint64
}
