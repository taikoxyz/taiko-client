// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TaikoDataAuction is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataAuction struct {
	Bid       TaikoDataBid
	BatchId   uint64
	StartedAt uint64
}

// TaikoDataBid is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataBid struct {
	Prover           common.Address
	Deposit          uint64
	FeePerGas        uint64
	BlockMaxGasLimit uint64
	ProofWindow      uint16
}

// TaikoDataBlockMetadata is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataBlockMetadata struct {
	Id                uint64
	Timestamp         uint64
	L1Height          uint64
	L1Hash            [32]byte
	MixHash           [32]byte
	TxListHash        [32]byte
	TxListByteStart   *big.Int
	TxListByteEnd     *big.Int
	GasLimit          uint32
	Beneficiary       common.Address
	Treasury          common.Address
	DepositsProcessed []TaikoDataEthDeposit
}

// TaikoDataConfig is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataConfig struct {
	ChainId                      *big.Int
	MaxNumProposedBlocks         *big.Int
	BlockRingBufferSize          *big.Int
	AuctionRingBufferSize        *big.Int
	MaxVerificationsPerTx        *big.Int
	BlockMaxGasLimit             uint64
	BlockFeeBaseGas              uint64
	MaxTransactionsPerBlock      uint64
	MaxBytesPerTxList            uint64
	TxListCacheExpiry            *big.Int
	ProofCooldownPeriod          *big.Int
	SystemProofCooldownPeriod    *big.Int
	EthDepositGas                *big.Int
	EthDepositMaxFee             *big.Int
	MinEthDepositsPerBlock       uint64
	MaxEthDepositsPerBlock       uint64
	MaxEthDepositAmount          *big.Int
	MinEthDepositAmount          *big.Int
	AuctionWindow                uint16
	AuctionProofWindowMultiplier uint64
	AuctionDepositMultipler      uint64
	AuctionMaxFeePerGasMultipler uint64
	AuctionBatchSize             uint16
	AuctionMaxAheadOfProposals   uint16
	AuctionMaxProofWindow        uint16
	RelaySignalRoot              bool
}

// TaikoDataEthDeposit is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataEthDeposit struct {
	Recipient common.Address
	Amount    *big.Int
	Id        uint64
}

// TaikoDataForkChoice is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataForkChoice struct {
	Key        [32]byte
	BlockHash  [32]byte
	SignalRoot [32]byte
	ProvenAt   uint64
	Prover     common.Address
	GasUsed    uint32
}

// TaikoDataStateVariables is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataStateVariables struct {
	FeePerGas               uint64
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	NumBlocks               uint64
	LastVerifiedBlockId     uint64
	NumAuctions             uint64
	NextEthDepositToProcess uint64
	NumEthDeposits          uint64
}

// TaikoL1ClientMetaData contains all meta data concerning the TaikoL1Client contract.
var TaikoL1ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BATCH_NOT_AUCTIONABLE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BATCH_NOT_AUCTIONABLE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_BID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_BID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF_OVERWRITE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF_OVERWRITE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_BETTER_BID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_BETTER_BID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_PROVEABLE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_PROVEABLE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_SPECIAL_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SAME_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SAME_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_UNAUTHORIZED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_UNAUTHORIZED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addressManager\",\"type\":\"address\"}],\"name\":\"AddressManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feePerGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"proofWindow\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structTaikoData.Bid\",\"name\":\"bid\",\"type\":\"tuple\"}],\"name\":\"BatchBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feePerGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"proofWindow\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structTaikoData.Bid\",\"name\":\"bid\",\"type\":\"tuple\"}],\"name\":\"BatchBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structTaikoData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structTaikoData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"BlockProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"BlockProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reward\",\"type\":\"uint64\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reward\",\"type\":\"uint64\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTaikoData.EthDeposit\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"EthDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTaikoData.EthDeposit\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"EthDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"adjustmentQuotient\",\"type\":\"uint16\"}],\"name\":\"ProofParamsChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"batchForBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feePerGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"proofWindow\",\"type\":\"uint16\"}],\"internalType\":\"structTaikoData.Bid\",\"name\":\"bid\",\"type\":\"tuple\"}],\"name\":\"bidForBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEtherToL2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositTaikoToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startBatchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getAuctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feePerGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"proofWindow\",\"type\":\"uint16\"}],\"internalType\":\"structTaikoData.Bid\",\"name\":\"bid\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.Auction[]\",\"name\":\"auctions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_metaHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_proposedAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"name\":\"getBlockFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumProposedBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockRingBufferSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auctionRingBufferSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVerificationsPerTx\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockFeeBaseGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTransactionsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxBytesPerTxList\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txListCacheExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"systemProofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositMaxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minEthDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxEthDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxEthDepositAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"minEthDepositAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint16\",\"name\":\"auctionWindow\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"auctionProofWindowMultiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionDepositMultipler\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionMaxFeePerGasMultipler\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"auctionBatchSize\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"auctionMaxAheadOfProposals\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"auctionMaxProofWindow\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"relaySignalRoot\",\"type\":\"bool\"}],\"internalType\":\"structTaikoData.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainSignalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"getForkChoice\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasUsed\",\"type\":\"uint32\"}],\"internalType\":\"structTaikoData.ForkChoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateVariables\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"feePerGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numAuctions\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numEthDeposits\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.StateVariables\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getTaikoTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"getVerifierName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_initFeePerGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_initAvgProofWindow\",\"type\":\"uint64\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"isBatchAuctionable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feePerGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"proofWindow\",\"type\":\"uint16\"}],\"internalType\":\"structTaikoData.Bid\",\"name\":\"newBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feePerGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"proofWindow\",\"type\":\"uint16\"}],\"internalType\":\"structTaikoData.Bid\",\"name\":\"oldBid\",\"type\":\"tuple\"}],\"name\":\"isBidBetter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"isBlockProvableBy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feePerGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"proofWindow\",\"type\":\"uint16\"}],\"internalType\":\"structTaikoData.Bid\",\"name\":\"bid\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.Auction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txList\",\"type\":\"bytes\"}],\"name\":\"proposeBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"}],\"internalType\":\"structTaikoData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"proveBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reserved70\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reserved71\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numAuctions\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reserved81\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feePerGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avgProofWindow\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"name\":\"verifyBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTaikoToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TaikoL1ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use TaikoL1ClientMetaData.ABI instead.
var TaikoL1ClientABI = TaikoL1ClientMetaData.ABI

// TaikoL1Client is an auto generated Go binding around an Ethereum contract.
type TaikoL1Client struct {
	TaikoL1ClientCaller     // Read-only binding to the contract
	TaikoL1ClientTransactor // Write-only binding to the contract
	TaikoL1ClientFilterer   // Log filterer for contract events
}

// TaikoL1ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaikoL1ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL1ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaikoL1ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL1ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaikoL1ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL1ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaikoL1ClientSession struct {
	Contract     *TaikoL1Client    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaikoL1ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaikoL1ClientCallerSession struct {
	Contract *TaikoL1ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TaikoL1ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaikoL1ClientTransactorSession struct {
	Contract     *TaikoL1ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TaikoL1ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaikoL1ClientRaw struct {
	Contract *TaikoL1Client // Generic contract binding to access the raw methods on
}

// TaikoL1ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaikoL1ClientCallerRaw struct {
	Contract *TaikoL1ClientCaller // Generic read-only contract binding to access the raw methods on
}

// TaikoL1ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaikoL1ClientTransactorRaw struct {
	Contract *TaikoL1ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaikoL1Client creates a new instance of TaikoL1Client, bound to a specific deployed contract.
func NewTaikoL1Client(address common.Address, backend bind.ContractBackend) (*TaikoL1Client, error) {
	contract, err := bindTaikoL1Client(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaikoL1Client{TaikoL1ClientCaller: TaikoL1ClientCaller{contract: contract}, TaikoL1ClientTransactor: TaikoL1ClientTransactor{contract: contract}, TaikoL1ClientFilterer: TaikoL1ClientFilterer{contract: contract}}, nil
}

// NewTaikoL1ClientCaller creates a new read-only instance of TaikoL1Client, bound to a specific deployed contract.
func NewTaikoL1ClientCaller(address common.Address, caller bind.ContractCaller) (*TaikoL1ClientCaller, error) {
	contract, err := bindTaikoL1Client(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientCaller{contract: contract}, nil
}

// NewTaikoL1ClientTransactor creates a new write-only instance of TaikoL1Client, bound to a specific deployed contract.
func NewTaikoL1ClientTransactor(address common.Address, transactor bind.ContractTransactor) (*TaikoL1ClientTransactor, error) {
	contract, err := bindTaikoL1Client(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTransactor{contract: contract}, nil
}

// NewTaikoL1ClientFilterer creates a new log filterer instance of TaikoL1Client, bound to a specific deployed contract.
func NewTaikoL1ClientFilterer(address common.Address, filterer bind.ContractFilterer) (*TaikoL1ClientFilterer, error) {
	contract, err := bindTaikoL1Client(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientFilterer{contract: contract}, nil
}

// bindTaikoL1Client binds a generic wrapper to an already deployed contract.
func bindTaikoL1Client(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaikoL1ClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoL1Client *TaikoL1ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoL1Client.Contract.TaikoL1ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoL1Client *TaikoL1ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.TaikoL1ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoL1Client *TaikoL1ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.TaikoL1ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoL1Client *TaikoL1ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoL1Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoL1Client *TaikoL1ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoL1Client *TaikoL1ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.contract.Transact(opts, method, params...)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL1Client *TaikoL1ClientCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL1Client *TaikoL1ClientSession) AddressManager() (common.Address, error) {
	return _TaikoL1Client.Contract.AddressManager(&_TaikoL1Client.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL1Client *TaikoL1ClientCallerSession) AddressManager() (common.Address, error) {
	return _TaikoL1Client.Contract.AddressManager(&_TaikoL1Client.CallOpts)
}

// BatchForBlock is a free data retrieval call binding the contract method 0xba9bd83c.
//
// Solidity: function batchForBlock(uint256 blockId) pure returns(uint256)
func (_TaikoL1Client *TaikoL1ClientCaller) BatchForBlock(opts *bind.CallOpts, blockId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "batchForBlock", blockId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchForBlock is a free data retrieval call binding the contract method 0xba9bd83c.
//
// Solidity: function batchForBlock(uint256 blockId) pure returns(uint256)
func (_TaikoL1Client *TaikoL1ClientSession) BatchForBlock(blockId *big.Int) (*big.Int, error) {
	return _TaikoL1Client.Contract.BatchForBlock(&_TaikoL1Client.CallOpts, blockId)
}

// BatchForBlock is a free data retrieval call binding the contract method 0xba9bd83c.
//
// Solidity: function batchForBlock(uint256 blockId) pure returns(uint256)
func (_TaikoL1Client *TaikoL1ClientCallerSession) BatchForBlock(blockId *big.Int) (*big.Int, error) {
	return _TaikoL1Client.Contract.BatchForBlock(&_TaikoL1Client.CallOpts, blockId)
}

// GetAuctions is a free data retrieval call binding the contract method 0xceb6a22f.
//
// Solidity: function getAuctions(uint256 startBatchId, uint256 count) view returns(uint256 currentTime, ((address,uint64,uint64,uint64,uint16),uint64,uint64)[] auctions)
func (_TaikoL1Client *TaikoL1ClientCaller) GetAuctions(opts *bind.CallOpts, startBatchId *big.Int, count *big.Int) (struct {
	CurrentTime *big.Int
	Auctions    []TaikoDataAuction
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getAuctions", startBatchId, count)

	outstruct := new(struct {
		CurrentTime *big.Int
		Auctions    []TaikoDataAuction
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Auctions = *abi.ConvertType(out[1], new([]TaikoDataAuction)).(*[]TaikoDataAuction)

	return *outstruct, err

}

// GetAuctions is a free data retrieval call binding the contract method 0xceb6a22f.
//
// Solidity: function getAuctions(uint256 startBatchId, uint256 count) view returns(uint256 currentTime, ((address,uint64,uint64,uint64,uint16),uint64,uint64)[] auctions)
func (_TaikoL1Client *TaikoL1ClientSession) GetAuctions(startBatchId *big.Int, count *big.Int) (struct {
	CurrentTime *big.Int
	Auctions    []TaikoDataAuction
}, error) {
	return _TaikoL1Client.Contract.GetAuctions(&_TaikoL1Client.CallOpts, startBatchId, count)
}

// GetAuctions is a free data retrieval call binding the contract method 0xceb6a22f.
//
// Solidity: function getAuctions(uint256 startBatchId, uint256 count) view returns(uint256 currentTime, ((address,uint64,uint64,uint64,uint16),uint64,uint64)[] auctions)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetAuctions(startBatchId *big.Int, count *big.Int) (struct {
	CurrentTime *big.Int
	Auctions    []TaikoDataAuction
}, error) {
	return _TaikoL1Client.Contract.GetAuctions(&_TaikoL1Client.CallOpts, startBatchId, count)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, address _proposer, uint64 _proposedAt)
func (_TaikoL1Client *TaikoL1ClientCaller) GetBlock(opts *bind.CallOpts, blockId *big.Int) (struct {
	MetaHash   [32]byte
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getBlock", blockId)

	outstruct := new(struct {
		MetaHash   [32]byte
		Proposer   common.Address
		ProposedAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MetaHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Proposer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ProposedAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, address _proposer, uint64 _proposedAt)
func (_TaikoL1Client *TaikoL1ClientSession) GetBlock(blockId *big.Int) (struct {
	MetaHash   [32]byte
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	return _TaikoL1Client.Contract.GetBlock(&_TaikoL1Client.CallOpts, blockId)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, address _proposer, uint64 _proposedAt)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetBlock(blockId *big.Int) (struct {
	MetaHash   [32]byte
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	return _TaikoL1Client.Contract.GetBlock(&_TaikoL1Client.CallOpts, blockId)
}

// GetBlockFee is a free data retrieval call binding the contract method 0xffb2784f.
//
// Solidity: function getBlockFee(uint32 gasLimit) view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientCaller) GetBlockFee(opts *bind.CallOpts, gasLimit uint32) (uint64, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getBlockFee", gasLimit)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBlockFee is a free data retrieval call binding the contract method 0xffb2784f.
//
// Solidity: function getBlockFee(uint32 gasLimit) view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientSession) GetBlockFee(gasLimit uint32) (uint64, error) {
	return _TaikoL1Client.Contract.GetBlockFee(&_TaikoL1Client.CallOpts, gasLimit)
}

// GetBlockFee is a free data retrieval call binding the contract method 0xffb2784f.
//
// Solidity: function getBlockFee(uint32 gasLimit) view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetBlockFee(gasLimit uint32) (uint64, error) {
	return _TaikoL1Client.Contract.GetBlockFee(&_TaikoL1Client.CallOpts, gasLimit)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,uint16,uint64,uint64,uint64,uint16,uint16,uint16,bool))
func (_TaikoL1Client *TaikoL1ClientCaller) GetConfig(opts *bind.CallOpts) (TaikoDataConfig, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(TaikoDataConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoDataConfig)).(*TaikoDataConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,uint16,uint64,uint64,uint64,uint16,uint16,uint16,bool))
func (_TaikoL1Client *TaikoL1ClientSession) GetConfig() (TaikoDataConfig, error) {
	return _TaikoL1Client.Contract.GetConfig(&_TaikoL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,uint16,uint64,uint64,uint64,uint16,uint16,uint16,bool))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetConfig() (TaikoDataConfig, error) {
	return _TaikoL1Client.Contract.GetConfig(&_TaikoL1Client.CallOpts)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCaller) GetCrossChainBlockHash(opts *bind.CallOpts, blockId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getCrossChainBlockHash", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientSession) GetCrossChainBlockHash(blockId *big.Int) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetCrossChainBlockHash(&_TaikoL1Client.CallOpts, blockId)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetCrossChainBlockHash(blockId *big.Int) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetCrossChainBlockHash(&_TaikoL1Client.CallOpts, blockId)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCaller) GetCrossChainSignalRoot(opts *bind.CallOpts, blockId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getCrossChainSignalRoot", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientSession) GetCrossChainSignalRoot(blockId *big.Int) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetCrossChainSignalRoot(&_TaikoL1Client.CallOpts, blockId)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetCrossChainSignalRoot(blockId *big.Int) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetCrossChainSignalRoot(&_TaikoL1Client.CallOpts, blockId)
}

// GetForkChoice is a free data retrieval call binding the contract method 0x7163e0ed.
//
// Solidity: function getForkChoice(uint256 blockId, bytes32 parentHash, uint32 parentGasUsed) view returns((bytes32,bytes32,bytes32,uint64,address,uint32))
func (_TaikoL1Client *TaikoL1ClientCaller) GetForkChoice(opts *bind.CallOpts, blockId *big.Int, parentHash [32]byte, parentGasUsed uint32) (TaikoDataForkChoice, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getForkChoice", blockId, parentHash, parentGasUsed)

	if err != nil {
		return *new(TaikoDataForkChoice), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoDataForkChoice)).(*TaikoDataForkChoice)

	return out0, err

}

// GetForkChoice is a free data retrieval call binding the contract method 0x7163e0ed.
//
// Solidity: function getForkChoice(uint256 blockId, bytes32 parentHash, uint32 parentGasUsed) view returns((bytes32,bytes32,bytes32,uint64,address,uint32))
func (_TaikoL1Client *TaikoL1ClientSession) GetForkChoice(blockId *big.Int, parentHash [32]byte, parentGasUsed uint32) (TaikoDataForkChoice, error) {
	return _TaikoL1Client.Contract.GetForkChoice(&_TaikoL1Client.CallOpts, blockId, parentHash, parentGasUsed)
}

// GetForkChoice is a free data retrieval call binding the contract method 0x7163e0ed.
//
// Solidity: function getForkChoice(uint256 blockId, bytes32 parentHash, uint32 parentGasUsed) view returns((bytes32,bytes32,bytes32,uint64,address,uint32))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetForkChoice(blockId *big.Int, parentHash [32]byte, parentGasUsed uint32) (TaikoDataForkChoice, error) {
	return _TaikoL1Client.Contract.GetForkChoice(&_TaikoL1Client.CallOpts, blockId, parentHash, parentGasUsed)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientCaller) GetStateVariables(opts *bind.CallOpts) (TaikoDataStateVariables, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getStateVariables")

	if err != nil {
		return *new(TaikoDataStateVariables), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoDataStateVariables)).(*TaikoDataStateVariables)

	return out0, err

}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientSession) GetStateVariables() (TaikoDataStateVariables, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetStateVariables() (TaikoDataStateVariables, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetTaikoTokenBalance is a free data retrieval call binding the contract method 0x8dff9cea.
//
// Solidity: function getTaikoTokenBalance(address addr) view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientCaller) GetTaikoTokenBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getTaikoTokenBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTaikoTokenBalance is a free data retrieval call binding the contract method 0x8dff9cea.
//
// Solidity: function getTaikoTokenBalance(address addr) view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientSession) GetTaikoTokenBalance(addr common.Address) (*big.Int, error) {
	return _TaikoL1Client.Contract.GetTaikoTokenBalance(&_TaikoL1Client.CallOpts, addr)
}

// GetTaikoTokenBalance is a free data retrieval call binding the contract method 0x8dff9cea.
//
// Solidity: function getTaikoTokenBalance(address addr) view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetTaikoTokenBalance(addr common.Address) (*big.Int, error) {
	return _TaikoL1Client.Contract.GetTaikoTokenBalance(&_TaikoL1Client.CallOpts, addr)
}

// GetVerifierName is a free data retrieval call binding the contract method 0x0372303d.
//
// Solidity: function getVerifierName(uint16 id) pure returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCaller) GetVerifierName(opts *bind.CallOpts, id uint16) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getVerifierName", id)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifierName is a free data retrieval call binding the contract method 0x0372303d.
//
// Solidity: function getVerifierName(uint16 id) pure returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientSession) GetVerifierName(id uint16) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetVerifierName(&_TaikoL1Client.CallOpts, id)
}

// GetVerifierName is a free data retrieval call binding the contract method 0x0372303d.
//
// Solidity: function getVerifierName(uint16 id) pure returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetVerifierName(id uint16) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetVerifierName(&_TaikoL1Client.CallOpts, id)
}

// IsBatchAuctionable is a free data retrieval call binding the contract method 0xd1a8f4ba.
//
// Solidity: function isBatchAuctionable(uint256 batchId) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) IsBatchAuctionable(opts *bind.CallOpts, batchId *big.Int) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "isBatchAuctionable", batchId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchAuctionable is a free data retrieval call binding the contract method 0xd1a8f4ba.
//
// Solidity: function isBatchAuctionable(uint256 batchId) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) IsBatchAuctionable(batchId *big.Int) (bool, error) {
	return _TaikoL1Client.Contract.IsBatchAuctionable(&_TaikoL1Client.CallOpts, batchId)
}

// IsBatchAuctionable is a free data retrieval call binding the contract method 0xd1a8f4ba.
//
// Solidity: function isBatchAuctionable(uint256 batchId) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) IsBatchAuctionable(batchId *big.Int) (bool, error) {
	return _TaikoL1Client.Contract.IsBatchAuctionable(&_TaikoL1Client.CallOpts, batchId)
}

// IsBidBetter is a free data retrieval call binding the contract method 0x2b94e559.
//
// Solidity: function isBidBetter((address,uint64,uint64,uint64,uint16) newBid, (address,uint64,uint64,uint64,uint16) oldBid) pure returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) IsBidBetter(opts *bind.CallOpts, newBid TaikoDataBid, oldBid TaikoDataBid) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "isBidBetter", newBid, oldBid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBidBetter is a free data retrieval call binding the contract method 0x2b94e559.
//
// Solidity: function isBidBetter((address,uint64,uint64,uint64,uint16) newBid, (address,uint64,uint64,uint64,uint16) oldBid) pure returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) IsBidBetter(newBid TaikoDataBid, oldBid TaikoDataBid) (bool, error) {
	return _TaikoL1Client.Contract.IsBidBetter(&_TaikoL1Client.CallOpts, newBid, oldBid)
}

// IsBidBetter is a free data retrieval call binding the contract method 0x2b94e559.
//
// Solidity: function isBidBetter((address,uint64,uint64,uint64,uint16) newBid, (address,uint64,uint64,uint64,uint16) oldBid) pure returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) IsBidBetter(newBid TaikoDataBid, oldBid TaikoDataBid) (bool, error) {
	return _TaikoL1Client.Contract.IsBidBetter(&_TaikoL1Client.CallOpts, newBid, oldBid)
}

// IsBlockProvableBy is a free data retrieval call binding the contract method 0x5e060b62.
//
// Solidity: function isBlockProvableBy(uint256 blockId, address prover) view returns(bool, ((address,uint64,uint64,uint64,uint16),uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientCaller) IsBlockProvableBy(opts *bind.CallOpts, blockId *big.Int, prover common.Address) (bool, TaikoDataAuction, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "isBlockProvableBy", blockId, prover)

	if err != nil {
		return *new(bool), *new(TaikoDataAuction), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(TaikoDataAuction)).(*TaikoDataAuction)

	return out0, out1, err

}

// IsBlockProvableBy is a free data retrieval call binding the contract method 0x5e060b62.
//
// Solidity: function isBlockProvableBy(uint256 blockId, address prover) view returns(bool, ((address,uint64,uint64,uint64,uint16),uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientSession) IsBlockProvableBy(blockId *big.Int, prover common.Address) (bool, TaikoDataAuction, error) {
	return _TaikoL1Client.Contract.IsBlockProvableBy(&_TaikoL1Client.CallOpts, blockId, prover)
}

// IsBlockProvableBy is a free data retrieval call binding the contract method 0x5e060b62.
//
// Solidity: function isBlockProvableBy(uint256 blockId, address prover) view returns(bool, ((address,uint64,uint64,uint64,uint16),uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientCallerSession) IsBlockProvableBy(blockId *big.Int, prover common.Address) (bool, TaikoDataAuction, error) {
	return _TaikoL1Client.Contract.IsBlockProvableBy(&_TaikoL1Client.CallOpts, blockId, prover)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL1Client *TaikoL1ClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL1Client *TaikoL1ClientSession) Owner() (common.Address, error) {
	return _TaikoL1Client.Contract.Owner(&_TaikoL1Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Owner() (common.Address, error) {
	return _TaikoL1Client.Contract.Owner(&_TaikoL1Client.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCaller) Resolve(opts *bind.CallOpts, chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve(&_TaikoL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve(&_TaikoL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCaller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve0(&_TaikoL1Client.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve0(&_TaikoL1Client.CallOpts, name, allowZeroAddress)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reserved70, uint64 __reserved71, uint64 numAuctions, uint64 __reserved81, uint64 numBlocks, uint64 nextEthDepositToProcess, uint64 lastVerifiedAt, uint64 feePerGas, uint64 lastVerifiedBlockId, uint64 avgProofWindow)
func (_TaikoL1Client *TaikoL1ClientCaller) State(opts *bind.CallOpts) (struct {
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	Reserved70              uint64
	Reserved71              uint64
	NumAuctions             uint64
	Reserved81              uint64
	NumBlocks               uint64
	NextEthDepositToProcess uint64
	LastVerifiedAt          uint64
	FeePerGas               uint64
	LastVerifiedBlockId     uint64
	AvgProofWindow          uint64
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		GenesisHeight           uint64
		GenesisTimestamp        uint64
		Reserved70              uint64
		Reserved71              uint64
		NumAuctions             uint64
		Reserved81              uint64
		NumBlocks               uint64
		NextEthDepositToProcess uint64
		LastVerifiedAt          uint64
		FeePerGas               uint64
		LastVerifiedBlockId     uint64
		AvgProofWindow          uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GenesisHeight = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.GenesisTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Reserved70 = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Reserved71 = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.NumAuctions = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.Reserved81 = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.NumBlocks = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.NextEthDepositToProcess = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.LastVerifiedAt = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.FeePerGas = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.LastVerifiedBlockId = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.AvgProofWindow = *abi.ConvertType(out[11], new(uint64)).(*uint64)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reserved70, uint64 __reserved71, uint64 numAuctions, uint64 __reserved81, uint64 numBlocks, uint64 nextEthDepositToProcess, uint64 lastVerifiedAt, uint64 feePerGas, uint64 lastVerifiedBlockId, uint64 avgProofWindow)
func (_TaikoL1Client *TaikoL1ClientSession) State() (struct {
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	Reserved70              uint64
	Reserved71              uint64
	NumAuctions             uint64
	Reserved81              uint64
	NumBlocks               uint64
	NextEthDepositToProcess uint64
	LastVerifiedAt          uint64
	FeePerGas               uint64
	LastVerifiedBlockId     uint64
	AvgProofWindow          uint64
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reserved70, uint64 __reserved71, uint64 numAuctions, uint64 __reserved81, uint64 numBlocks, uint64 nextEthDepositToProcess, uint64 lastVerifiedAt, uint64 feePerGas, uint64 lastVerifiedBlockId, uint64 avgProofWindow)
func (_TaikoL1Client *TaikoL1ClientCallerSession) State() (struct {
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	Reserved70              uint64
	Reserved71              uint64
	NumAuctions             uint64
	Reserved81              uint64
	NumBlocks               uint64
	NextEthDepositToProcess uint64
	LastVerifiedAt          uint64
	FeePerGas               uint64
	LastVerifiedBlockId     uint64
	AvgProofWindow          uint64
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// BidForBatch is a paid mutator transaction binding the contract method 0x16b54e64.
//
// Solidity: function bidForBatch(uint64 batchId, (address,uint64,uint64,uint64,uint16) bid) payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) BidForBatch(opts *bind.TransactOpts, batchId uint64, bid TaikoDataBid) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "bidForBatch", batchId, bid)
}

// BidForBatch is a paid mutator transaction binding the contract method 0x16b54e64.
//
// Solidity: function bidForBatch(uint64 batchId, (address,uint64,uint64,uint64,uint16) bid) payable returns()
func (_TaikoL1Client *TaikoL1ClientSession) BidForBatch(batchId uint64, bid TaikoDataBid) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.BidForBatch(&_TaikoL1Client.TransactOpts, batchId, bid)
}

// BidForBatch is a paid mutator transaction binding the contract method 0x16b54e64.
//
// Solidity: function bidForBatch(uint64 batchId, (address,uint64,uint64,uint64,uint16) bid) payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) BidForBatch(batchId uint64, bid TaikoDataBid) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.BidForBatch(&_TaikoL1Client.TransactOpts, batchId, bid)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0xa22f7670.
//
// Solidity: function depositEtherToL2() payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) DepositEtherToL2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "depositEtherToL2")
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0xa22f7670.
//
// Solidity: function depositEtherToL2() payable returns()
func (_TaikoL1Client *TaikoL1ClientSession) DepositEtherToL2() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.DepositEtherToL2(&_TaikoL1Client.TransactOpts)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0xa22f7670.
//
// Solidity: function depositEtherToL2() payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) DepositEtherToL2() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.DepositEtherToL2(&_TaikoL1Client.TransactOpts)
}

// DepositTaikoToken is a paid mutator transaction binding the contract method 0x98f39aba.
//
// Solidity: function depositTaikoToken(uint256 amount) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) DepositTaikoToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "depositTaikoToken", amount)
}

// DepositTaikoToken is a paid mutator transaction binding the contract method 0x98f39aba.
//
// Solidity: function depositTaikoToken(uint256 amount) returns()
func (_TaikoL1Client *TaikoL1ClientSession) DepositTaikoToken(amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.DepositTaikoToken(&_TaikoL1Client.TransactOpts, amount)
}

// DepositTaikoToken is a paid mutator transaction binding the contract method 0x98f39aba.
//
// Solidity: function depositTaikoToken(uint256 amount) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) DepositTaikoToken(amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.DepositTaikoToken(&_TaikoL1Client.TransactOpts, amount)
}

// Init is a paid mutator transaction binding the contract method 0x578c65a4.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initFeePerGas, uint64 _initAvgProofWindow) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _genesisBlockHash [32]byte, _initFeePerGas uint64, _initAvgProofWindow uint64) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "init", _addressManager, _genesisBlockHash, _initFeePerGas, _initAvgProofWindow)
}

// Init is a paid mutator transaction binding the contract method 0x578c65a4.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initFeePerGas, uint64 _initAvgProofWindow) returns()
func (_TaikoL1Client *TaikoL1ClientSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _initFeePerGas uint64, _initAvgProofWindow uint64) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Init(&_TaikoL1Client.TransactOpts, _addressManager, _genesisBlockHash, _initFeePerGas, _initAvgProofWindow)
}

// Init is a paid mutator transaction binding the contract method 0x578c65a4.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initFeePerGas, uint64 _initAvgProofWindow) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _initFeePerGas uint64, _initAvgProofWindow uint64) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Init(&_TaikoL1Client.TransactOpts, _addressManager, _genesisBlockHash, _initFeePerGas, _initAvgProofWindow)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes input, bytes txList) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientTransactor) ProposeBlock(opts *bind.TransactOpts, input []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proposeBlock", input, txList)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes input, bytes txList) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientSession) ProposeBlock(input []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, input, txList)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes input, bytes txList) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProposeBlock(input []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, input, txList)
}

// ProveBlock is a paid mutator transaction binding the contract method 0xf3840f60.
//
// Solidity: function proveBlock(uint256 blockId, bytes input) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) ProveBlock(opts *bind.TransactOpts, blockId *big.Int, input []byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proveBlock", blockId, input)
}

// ProveBlock is a paid mutator transaction binding the contract method 0xf3840f60.
//
// Solidity: function proveBlock(uint256 blockId, bytes input) returns()
func (_TaikoL1Client *TaikoL1ClientSession) ProveBlock(blockId *big.Int, input []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlock(&_TaikoL1Client.TransactOpts, blockId, input)
}

// ProveBlock is a paid mutator transaction binding the contract method 0xf3840f60.
//
// Solidity: function proveBlock(uint256 blockId, bytes input) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProveBlock(blockId *big.Int, input []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlock(&_TaikoL1Client.TransactOpts, blockId, input)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL1Client *TaikoL1ClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.RenounceOwnership(&_TaikoL1Client.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.RenounceOwnership(&_TaikoL1Client.TransactOpts)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) SetAddressManager(opts *bind.TransactOpts, newAddressManager common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "setAddressManager", newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_TaikoL1Client *TaikoL1ClientSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.SetAddressManager(&_TaikoL1Client.TransactOpts, newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.SetAddressManager(&_TaikoL1Client.TransactOpts, newAddressManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL1Client *TaikoL1ClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.TransferOwnership(&_TaikoL1Client.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.TransferOwnership(&_TaikoL1Client.TransactOpts, newOwner)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) VerifyBlocks(opts *bind.TransactOpts, maxBlocks *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "verifyBlocks", maxBlocks)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_TaikoL1Client *TaikoL1ClientSession) VerifyBlocks(maxBlocks *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.VerifyBlocks(&_TaikoL1Client.TransactOpts, maxBlocks)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) VerifyBlocks(maxBlocks *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.VerifyBlocks(&_TaikoL1Client.TransactOpts, maxBlocks)
}

// WithdrawTaikoToken is a paid mutator transaction binding the contract method 0x5043f059.
//
// Solidity: function withdrawTaikoToken(uint256 amount) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) WithdrawTaikoToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "withdrawTaikoToken", amount)
}

// WithdrawTaikoToken is a paid mutator transaction binding the contract method 0x5043f059.
//
// Solidity: function withdrawTaikoToken(uint256 amount) returns()
func (_TaikoL1Client *TaikoL1ClientSession) WithdrawTaikoToken(amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.WithdrawTaikoToken(&_TaikoL1Client.TransactOpts, amount)
}

// WithdrawTaikoToken is a paid mutator transaction binding the contract method 0x5043f059.
//
// Solidity: function withdrawTaikoToken(uint256 amount) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) WithdrawTaikoToken(amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.WithdrawTaikoToken(&_TaikoL1Client.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaikoL1Client *TaikoL1ClientSession) Receive() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Receive(&_TaikoL1Client.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) Receive() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Receive(&_TaikoL1Client.TransactOpts)
}

// TaikoL1ClientAddressManagerChangedIterator is returned from FilterAddressManagerChanged and is used to iterate over the raw logs and unpacked data for AddressManagerChanged events raised by the TaikoL1Client contract.
type TaikoL1ClientAddressManagerChangedIterator struct {
	Event *TaikoL1ClientAddressManagerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientAddressManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientAddressManagerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientAddressManagerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientAddressManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientAddressManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientAddressManagerChanged represents a AddressManagerChanged event raised by the TaikoL1Client contract.
type TaikoL1ClientAddressManagerChanged struct {
	AddressManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressManagerChanged is a free log retrieval operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterAddressManagerChanged(opts *bind.FilterOpts) (*TaikoL1ClientAddressManagerChangedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "AddressManagerChanged")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientAddressManagerChangedIterator{contract: _TaikoL1Client.contract, event: "AddressManagerChanged", logs: logs, sub: sub}, nil
}

// WatchAddressManagerChanged is a free log subscription operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchAddressManagerChanged(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientAddressManagerChanged) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "AddressManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientAddressManagerChanged)
				if err := _TaikoL1Client.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddressManagerChanged is a log parse operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseAddressManagerChanged(log types.Log) (*TaikoL1ClientAddressManagerChanged, error) {
	event := new(TaikoL1ClientAddressManagerChanged)
	if err := _TaikoL1Client.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBatchBidIterator is returned from FilterBatchBid and is used to iterate over the raw logs and unpacked data for BatchBid events raised by the TaikoL1Client contract.
type TaikoL1ClientBatchBidIterator struct {
	Event *TaikoL1ClientBatchBid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBatchBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBatchBid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBatchBid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBatchBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBatchBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBatchBid represents a BatchBid event raised by the TaikoL1Client contract.
type TaikoL1ClientBatchBid struct {
	BatchId   uint64
	StartedAt uint64
	Bid       TaikoDataBid
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchBid is a free log retrieval operation binding the contract event 0xbc48648c22f0e34368831e00c57c0569c3fb7bd53d7b5c5cdfca260a7b9444d3.
//
// Solidity: event BatchBid(uint64 indexed batchId, uint64 startedAt, (address,uint64,uint64,uint64,uint16) bid)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBatchBid(opts *bind.FilterOpts, batchId []uint64) (*TaikoL1ClientBatchBidIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BatchBid", batchIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBatchBidIterator{contract: _TaikoL1Client.contract, event: "BatchBid", logs: logs, sub: sub}, nil
}

// WatchBatchBid is a free log subscription operation binding the contract event 0xbc48648c22f0e34368831e00c57c0569c3fb7bd53d7b5c5cdfca260a7b9444d3.
//
// Solidity: event BatchBid(uint64 indexed batchId, uint64 startedAt, (address,uint64,uint64,uint64,uint16) bid)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBatchBid(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBatchBid, batchId []uint64) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BatchBid", batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBatchBid)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BatchBid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchBid is a log parse operation binding the contract event 0xbc48648c22f0e34368831e00c57c0569c3fb7bd53d7b5c5cdfca260a7b9444d3.
//
// Solidity: event BatchBid(uint64 indexed batchId, uint64 startedAt, (address,uint64,uint64,uint64,uint16) bid)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBatchBid(log types.Log) (*TaikoL1ClientBatchBid, error) {
	event := new(TaikoL1ClientBatchBid)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BatchBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBatchBid0Iterator is returned from FilterBatchBid0 and is used to iterate over the raw logs and unpacked data for BatchBid0 events raised by the TaikoL1Client contract.
type TaikoL1ClientBatchBid0Iterator struct {
	Event *TaikoL1ClientBatchBid0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBatchBid0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBatchBid0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBatchBid0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBatchBid0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBatchBid0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBatchBid0 represents a BatchBid0 event raised by the TaikoL1Client contract.
type TaikoL1ClientBatchBid0 struct {
	BatchId   uint64
	StartedAt uint64
	Bid       TaikoDataBid
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchBid0 is a free log retrieval operation binding the contract event 0xbc48648c22f0e34368831e00c57c0569c3fb7bd53d7b5c5cdfca260a7b9444d3.
//
// Solidity: event BatchBid(uint64 indexed batchId, uint64 startedAt, (address,uint64,uint64,uint64,uint16) bid)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBatchBid0(opts *bind.FilterOpts, batchId []uint64) (*TaikoL1ClientBatchBid0Iterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BatchBid0", batchIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBatchBid0Iterator{contract: _TaikoL1Client.contract, event: "BatchBid0", logs: logs, sub: sub}, nil
}

// WatchBatchBid0 is a free log subscription operation binding the contract event 0xbc48648c22f0e34368831e00c57c0569c3fb7bd53d7b5c5cdfca260a7b9444d3.
//
// Solidity: event BatchBid(uint64 indexed batchId, uint64 startedAt, (address,uint64,uint64,uint64,uint16) bid)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBatchBid0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBatchBid0, batchId []uint64) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BatchBid0", batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBatchBid0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BatchBid0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchBid0 is a log parse operation binding the contract event 0xbc48648c22f0e34368831e00c57c0569c3fb7bd53d7b5c5cdfca260a7b9444d3.
//
// Solidity: event BatchBid(uint64 indexed batchId, uint64 startedAt, (address,uint64,uint64,uint64,uint16) bid)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBatchBid0(log types.Log) (*TaikoL1ClientBatchBid0, error) {
	event := new(TaikoL1ClientBatchBid0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BatchBid0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockProposedIterator is returned from FilterBlockProposed and is used to iterate over the raw logs and unpacked data for BlockProposed events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProposedIterator struct {
	Event *TaikoL1ClientBlockProposed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBlockProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockProposed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBlockProposed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBlockProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockProposed represents a BlockProposed event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProposed struct {
	Id       *big.Int
	Meta     TaikoDataBlockMetadata
	BlockFee uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0x555304f14500c230922dd951cfdbb74012005afbcd7350b4b9fd27dc12d710fe.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[]) meta, uint64 blockFee)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockProposed(opts *bind.FilterOpts, id []*big.Int) (*TaikoL1ClientBlockProposedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockProposed", idRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockProposedIterator{contract: _TaikoL1Client.contract, event: "BlockProposed", logs: logs, sub: sub}, nil
}

// WatchBlockProposed is a free log subscription operation binding the contract event 0x555304f14500c230922dd951cfdbb74012005afbcd7350b4b9fd27dc12d710fe.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[]) meta, uint64 blockFee)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockProposed(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockProposed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockProposed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockProposed)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProposed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockProposed is a log parse operation binding the contract event 0x555304f14500c230922dd951cfdbb74012005afbcd7350b4b9fd27dc12d710fe.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[]) meta, uint64 blockFee)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockProposed(log types.Log) (*TaikoL1ClientBlockProposed, error) {
	event := new(TaikoL1ClientBlockProposed)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockProposed0Iterator is returned from FilterBlockProposed0 and is used to iterate over the raw logs and unpacked data for BlockProposed0 events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProposed0Iterator struct {
	Event *TaikoL1ClientBlockProposed0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBlockProposed0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockProposed0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBlockProposed0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBlockProposed0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockProposed0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockProposed0 represents a BlockProposed0 event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProposed0 struct {
	Id       *big.Int
	Meta     TaikoDataBlockMetadata
	BlockFee uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed0 is a free log retrieval operation binding the contract event 0x555304f14500c230922dd951cfdbb74012005afbcd7350b4b9fd27dc12d710fe.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[]) meta, uint64 blockFee)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockProposed0(opts *bind.FilterOpts, id []*big.Int) (*TaikoL1ClientBlockProposed0Iterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockProposed0", idRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockProposed0Iterator{contract: _TaikoL1Client.contract, event: "BlockProposed0", logs: logs, sub: sub}, nil
}

// WatchBlockProposed0 is a free log subscription operation binding the contract event 0x555304f14500c230922dd951cfdbb74012005afbcd7350b4b9fd27dc12d710fe.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[]) meta, uint64 blockFee)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockProposed0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockProposed0, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockProposed0", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockProposed0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProposed0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockProposed0 is a log parse operation binding the contract event 0x555304f14500c230922dd951cfdbb74012005afbcd7350b4b9fd27dc12d710fe.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[]) meta, uint64 blockFee)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockProposed0(log types.Log) (*TaikoL1ClientBlockProposed0, error) {
	event := new(TaikoL1ClientBlockProposed0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProposed0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockProvenIterator is returned from FilterBlockProven and is used to iterate over the raw logs and unpacked data for BlockProven events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProvenIterator struct {
	Event *TaikoL1ClientBlockProven // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBlockProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockProven)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBlockProven)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBlockProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockProven represents a BlockProven event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProven struct {
	Id            *big.Int
	ParentHash    [32]byte
	BlockHash     [32]byte
	SignalRoot    [32]byte
	Prover        common.Address
	ParentGasUsed uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBlockProven is a free log retrieval operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockProven(opts *bind.FilterOpts, id []*big.Int) (*TaikoL1ClientBlockProvenIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockProven", idRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockProvenIterator{contract: _TaikoL1Client.contract, event: "BlockProven", logs: logs, sub: sub}, nil
}

// WatchBlockProven is a free log subscription operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockProven(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockProven, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockProven", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockProven)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProven", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockProven is a log parse operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockProven(log types.Log) (*TaikoL1ClientBlockProven, error) {
	event := new(TaikoL1ClientBlockProven)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockProven0Iterator is returned from FilterBlockProven0 and is used to iterate over the raw logs and unpacked data for BlockProven0 events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProven0Iterator struct {
	Event *TaikoL1ClientBlockProven0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBlockProven0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockProven0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBlockProven0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBlockProven0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockProven0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockProven0 represents a BlockProven0 event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProven0 struct {
	Id            *big.Int
	ParentHash    [32]byte
	BlockHash     [32]byte
	SignalRoot    [32]byte
	Prover        common.Address
	ParentGasUsed uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBlockProven0 is a free log retrieval operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockProven0(opts *bind.FilterOpts, id []*big.Int) (*TaikoL1ClientBlockProven0Iterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockProven0", idRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockProven0Iterator{contract: _TaikoL1Client.contract, event: "BlockProven0", logs: logs, sub: sub}, nil
}

// WatchBlockProven0 is a free log subscription operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockProven0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockProven0, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockProven0", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockProven0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProven0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockProven0 is a log parse operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockProven0(log types.Log) (*TaikoL1ClientBlockProven0, error) {
	event := new(TaikoL1ClientBlockProven0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProven0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockVerifiedIterator is returned from FilterBlockVerified and is used to iterate over the raw logs and unpacked data for BlockVerified events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockVerifiedIterator struct {
	Event *TaikoL1ClientBlockVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBlockVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBlockVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBlockVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockVerified represents a BlockVerified event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockVerified struct {
	Id        *big.Int
	BlockHash [32]byte
	Reward    uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified is a free log retrieval operation binding the contract event 0x71eb3f26e4ff6e88e66a79521590634475f8ddb0c9e30e0400de0c7f2b28fb83.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash, uint64 reward)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockVerified(opts *bind.FilterOpts, id []*big.Int) (*TaikoL1ClientBlockVerifiedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockVerified", idRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockVerifiedIterator{contract: _TaikoL1Client.contract, event: "BlockVerified", logs: logs, sub: sub}, nil
}

// WatchBlockVerified is a free log subscription operation binding the contract event 0x71eb3f26e4ff6e88e66a79521590634475f8ddb0c9e30e0400de0c7f2b28fb83.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash, uint64 reward)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockVerified(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockVerified, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockVerified", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockVerified)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockVerified is a log parse operation binding the contract event 0x71eb3f26e4ff6e88e66a79521590634475f8ddb0c9e30e0400de0c7f2b28fb83.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash, uint64 reward)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockVerified(log types.Log) (*TaikoL1ClientBlockVerified, error) {
	event := new(TaikoL1ClientBlockVerified)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockVerified0Iterator is returned from FilterBlockVerified0 and is used to iterate over the raw logs and unpacked data for BlockVerified0 events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockVerified0Iterator struct {
	Event *TaikoL1ClientBlockVerified0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBlockVerified0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockVerified0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBlockVerified0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBlockVerified0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockVerified0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockVerified0 represents a BlockVerified0 event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockVerified0 struct {
	Id        *big.Int
	BlockHash [32]byte
	Reward    uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified0 is a free log retrieval operation binding the contract event 0x71eb3f26e4ff6e88e66a79521590634475f8ddb0c9e30e0400de0c7f2b28fb83.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash, uint64 reward)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockVerified0(opts *bind.FilterOpts, id []*big.Int) (*TaikoL1ClientBlockVerified0Iterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockVerified0", idRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockVerified0Iterator{contract: _TaikoL1Client.contract, event: "BlockVerified0", logs: logs, sub: sub}, nil
}

// WatchBlockVerified0 is a free log subscription operation binding the contract event 0x71eb3f26e4ff6e88e66a79521590634475f8ddb0c9e30e0400de0c7f2b28fb83.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash, uint64 reward)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockVerified0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockVerified0, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockVerified0", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockVerified0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockVerified0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockVerified0 is a log parse operation binding the contract event 0x71eb3f26e4ff6e88e66a79521590634475f8ddb0c9e30e0400de0c7f2b28fb83.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash, uint64 reward)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockVerified0(log types.Log) (*TaikoL1ClientBlockVerified0, error) {
	event := new(TaikoL1ClientBlockVerified0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockVerified0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientCrossChainSyncedIterator is returned from FilterCrossChainSynced and is used to iterate over the raw logs and unpacked data for CrossChainSynced events raised by the TaikoL1Client contract.
type TaikoL1ClientCrossChainSyncedIterator struct {
	Event *TaikoL1ClientCrossChainSynced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientCrossChainSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientCrossChainSynced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientCrossChainSynced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientCrossChainSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientCrossChainSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientCrossChainSynced represents a CrossChainSynced event raised by the TaikoL1Client contract.
type TaikoL1ClientCrossChainSynced struct {
	SrcHeight  *big.Int
	BlockHash  [32]byte
	SignalRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced is a free log retrieval operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterCrossChainSynced(opts *bind.FilterOpts, srcHeight []*big.Int) (*TaikoL1ClientCrossChainSyncedIterator, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientCrossChainSyncedIterator{contract: _TaikoL1Client.contract, event: "CrossChainSynced", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced is a free log subscription operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchCrossChainSynced(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientCrossChainSynced, srcHeight []*big.Int) (event.Subscription, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientCrossChainSynced)
				if err := _TaikoL1Client.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainSynced is a log parse operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseCrossChainSynced(log types.Log) (*TaikoL1ClientCrossChainSynced, error) {
	event := new(TaikoL1ClientCrossChainSynced)
	if err := _TaikoL1Client.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientCrossChainSynced0Iterator is returned from FilterCrossChainSynced0 and is used to iterate over the raw logs and unpacked data for CrossChainSynced0 events raised by the TaikoL1Client contract.
type TaikoL1ClientCrossChainSynced0Iterator struct {
	Event *TaikoL1ClientCrossChainSynced0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientCrossChainSynced0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientCrossChainSynced0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientCrossChainSynced0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientCrossChainSynced0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientCrossChainSynced0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientCrossChainSynced0 represents a CrossChainSynced0 event raised by the TaikoL1Client contract.
type TaikoL1ClientCrossChainSynced0 struct {
	SrcHeight  *big.Int
	BlockHash  [32]byte
	SignalRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced0 is a free log retrieval operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterCrossChainSynced0(opts *bind.FilterOpts, srcHeight []*big.Int) (*TaikoL1ClientCrossChainSynced0Iterator, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "CrossChainSynced0", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientCrossChainSynced0Iterator{contract: _TaikoL1Client.contract, event: "CrossChainSynced0", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced0 is a free log subscription operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchCrossChainSynced0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientCrossChainSynced0, srcHeight []*big.Int) (event.Subscription, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "CrossChainSynced0", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientCrossChainSynced0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "CrossChainSynced0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainSynced0 is a log parse operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseCrossChainSynced0(log types.Log) (*TaikoL1ClientCrossChainSynced0, error) {
	event := new(TaikoL1ClientCrossChainSynced0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "CrossChainSynced0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientEthDepositedIterator is returned from FilterEthDeposited and is used to iterate over the raw logs and unpacked data for EthDeposited events raised by the TaikoL1Client contract.
type TaikoL1ClientEthDepositedIterator struct {
	Event *TaikoL1ClientEthDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientEthDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientEthDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientEthDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientEthDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientEthDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientEthDeposited represents a EthDeposited event raised by the TaikoL1Client contract.
type TaikoL1ClientEthDeposited struct {
	Deposit TaikoDataEthDeposit
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEthDeposited is a free log retrieval operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterEthDeposited(opts *bind.FilterOpts) (*TaikoL1ClientEthDepositedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "EthDeposited")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientEthDepositedIterator{contract: _TaikoL1Client.contract, event: "EthDeposited", logs: logs, sub: sub}, nil
}

// WatchEthDeposited is a free log subscription operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchEthDeposited(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientEthDeposited) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "EthDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientEthDeposited)
				if err := _TaikoL1Client.contract.UnpackLog(event, "EthDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEthDeposited is a log parse operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseEthDeposited(log types.Log) (*TaikoL1ClientEthDeposited, error) {
	event := new(TaikoL1ClientEthDeposited)
	if err := _TaikoL1Client.contract.UnpackLog(event, "EthDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientEthDeposited0Iterator is returned from FilterEthDeposited0 and is used to iterate over the raw logs and unpacked data for EthDeposited0 events raised by the TaikoL1Client contract.
type TaikoL1ClientEthDeposited0Iterator struct {
	Event *TaikoL1ClientEthDeposited0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientEthDeposited0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientEthDeposited0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientEthDeposited0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientEthDeposited0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientEthDeposited0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientEthDeposited0 represents a EthDeposited0 event raised by the TaikoL1Client contract.
type TaikoL1ClientEthDeposited0 struct {
	Deposit TaikoDataEthDeposit
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEthDeposited0 is a free log retrieval operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterEthDeposited0(opts *bind.FilterOpts) (*TaikoL1ClientEthDeposited0Iterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "EthDeposited0")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientEthDeposited0Iterator{contract: _TaikoL1Client.contract, event: "EthDeposited0", logs: logs, sub: sub}, nil
}

// WatchEthDeposited0 is a free log subscription operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchEthDeposited0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientEthDeposited0) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "EthDeposited0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientEthDeposited0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "EthDeposited0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEthDeposited0 is a log parse operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseEthDeposited0(log types.Log) (*TaikoL1ClientEthDeposited0, error) {
	event := new(TaikoL1ClientEthDeposited0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "EthDeposited0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaikoL1Client contract.
type TaikoL1ClientInitializedIterator struct {
	Event *TaikoL1ClientInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientInitialized represents a Initialized event raised by the TaikoL1Client contract.
type TaikoL1ClientInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaikoL1ClientInitializedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientInitializedIterator{contract: _TaikoL1Client.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientInitialized) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientInitialized)
				if err := _TaikoL1Client.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseInitialized(log types.Log) (*TaikoL1ClientInitialized, error) {
	event := new(TaikoL1ClientInitialized)
	if err := _TaikoL1Client.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaikoL1Client contract.
type TaikoL1ClientOwnershipTransferredIterator struct {
	Event *TaikoL1ClientOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientOwnershipTransferred represents a OwnershipTransferred event raised by the TaikoL1Client contract.
type TaikoL1ClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaikoL1ClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientOwnershipTransferredIterator{contract: _TaikoL1Client.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientOwnershipTransferred)
				if err := _TaikoL1Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseOwnershipTransferred(log types.Log) (*TaikoL1ClientOwnershipTransferred, error) {
	event := new(TaikoL1ClientOwnershipTransferred)
	if err := _TaikoL1Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientProofParamsChangedIterator is returned from FilterProofParamsChanged and is used to iterate over the raw logs and unpacked data for ProofParamsChanged events raised by the TaikoL1Client contract.
type TaikoL1ClientProofParamsChangedIterator struct {
	Event *TaikoL1ClientProofParamsChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientProofParamsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientProofParamsChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientProofParamsChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientProofParamsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientProofParamsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientProofParamsChanged represents a ProofParamsChanged event raised by the TaikoL1Client contract.
type TaikoL1ClientProofParamsChanged struct {
	ProofTimeTarget    uint64
	ProofTimeIssued    uint64
	BlockFee           uint64
	AdjustmentQuotient uint16
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProofParamsChanged is a free log retrieval operation binding the contract event 0x565e5aa69c99d81e441dd3bb8535d888585683743f3c6a3bf49e5e1b227bd8f9.
//
// Solidity: event ProofParamsChanged(uint64 proofTimeTarget, uint64 proofTimeIssued, uint64 blockFee, uint16 adjustmentQuotient)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterProofParamsChanged(opts *bind.FilterOpts) (*TaikoL1ClientProofParamsChangedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "ProofParamsChanged")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientProofParamsChangedIterator{contract: _TaikoL1Client.contract, event: "ProofParamsChanged", logs: logs, sub: sub}, nil
}

// WatchProofParamsChanged is a free log subscription operation binding the contract event 0x565e5aa69c99d81e441dd3bb8535d888585683743f3c6a3bf49e5e1b227bd8f9.
//
// Solidity: event ProofParamsChanged(uint64 proofTimeTarget, uint64 proofTimeIssued, uint64 blockFee, uint16 adjustmentQuotient)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchProofParamsChanged(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientProofParamsChanged) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "ProofParamsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientProofParamsChanged)
				if err := _TaikoL1Client.contract.UnpackLog(event, "ProofParamsChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProofParamsChanged is a log parse operation binding the contract event 0x565e5aa69c99d81e441dd3bb8535d888585683743f3c6a3bf49e5e1b227bd8f9.
//
// Solidity: event ProofParamsChanged(uint64 proofTimeTarget, uint64 proofTimeIssued, uint64 blockFee, uint16 adjustmentQuotient)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseProofParamsChanged(log types.Log) (*TaikoL1ClientProofParamsChanged, error) {
	event := new(TaikoL1ClientProofParamsChanged)
	if err := _TaikoL1Client.contract.UnpackLog(event, "ProofParamsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
