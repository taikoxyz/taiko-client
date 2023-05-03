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

// TaikoDataBlockMetadata is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataBlockMetadata struct {
	Id                uint64
	Timestamp         uint64
	L1Height          uint64
	L1Hash            [32]byte
	MixHash           [32]byte
	DepositsRoot      [32]byte
	TxListHash        [32]byte
	TxListByteStart   *big.Int
	TxListByteEnd     *big.Int
	GasLimit          uint32
	Beneficiary       common.Address
	CacheTxListInfo   uint8
	Treasure          common.Address
	DepositsProcessed []TaikoDataEthDeposit
}

// TaikoDataConfig is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataConfig struct {
	ChainId                 *big.Int
	MaxNumProposedBlocks    *big.Int
	RingBufferSize          *big.Int
	MaxVerificationsPerTx   *big.Int
	BlockMaxGasLimit        *big.Int
	MaxTransactionsPerBlock *big.Int
	MaxBytesPerTxList       *big.Int
	MinTxGasLimit           *big.Int
	TxListCacheExpiry       *big.Int
	ProofCooldownPeriod     *big.Int
	RealProofSkipSize       *big.Int
	EthDepositGas           *big.Int
	EthDepositMaxFee        *big.Int
	MinEthDepositsPerBlock  uint64
	MaxEthDepositsPerBlock  uint64
	MaxEthDepositAmount     *big.Int
	MinEthDepositAmount     *big.Int
	ProofTimeTarget         uint64
	AdjustmentQuotient      uint8
	RelaySignalRoot         bool
}

// TaikoDataEthDeposit is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataEthDeposit struct {
	Recipient common.Address
	Amount    *big.Int
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
	BlockFee                uint64
	AccBlockFees            uint64
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	NumBlocks               uint64
	ProofTimeIssued         uint64
	LastVerifiedBlockId     uint64
	AccProposedAt           uint64
	NextEthDepositToProcess uint64
	NumEthDeposits          uint64
}

// TaikoL1ClientMetaData contains all meta data concerning the TaikoL1Client contract.
var TaikoL1ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_ORACLE_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_ORACLE_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ORACLE_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ORACLE_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"depositsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"cacheTxListInfo\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"treasure\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structTaikoData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structTaikoData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"BlockProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structTaikoData.EthDeposit\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"EthDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEtherToL2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositTaikoToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_metaHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_deposit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_proposedAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumProposedBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ringBufferSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVerificationsPerTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTransactionsPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBytesPerTxList\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txListCacheExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"realProofSkipSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositMaxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minEthDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxEthDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxEthDepositAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"minEthDepositAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"adjustmentQuotient\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"relaySignalRoot\",\"type\":\"bool\"}],\"internalType\":\"structTaikoData.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainSignalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"getForkChoice\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasUsed\",\"type\":\"uint32\"}],\"internalType\":\"structTaikoData.ForkChoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposedAt\",\"type\":\"uint64\"}],\"name\":\"getProofReward\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateVariables\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accBlockFees\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numEthDeposits\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.StateVariables\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getTaikoTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"getVerifierName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_initBlockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_initProofTimeIssued\",\"type\":\"uint64\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txList\",\"type\":\"bytes\"}],\"name\":\"proposeBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"depositsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"cacheTxListInfo\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"treasure\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structTaikoData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"}],\"internalType\":\"structTaikoData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"proveBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"staticRefs\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reserved71\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reserved72\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accBlockFees\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reserved91\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"name\":\"verifyBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTaikoToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, uint256 _deposit, address _proposer, uint64 _proposedAt)
func (_TaikoL1Client *TaikoL1ClientCaller) GetBlock(opts *bind.CallOpts, blockId *big.Int) (struct {
	MetaHash   [32]byte
	Deposit    *big.Int
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getBlock", blockId)

	outstruct := new(struct {
		MetaHash   [32]byte
		Deposit    *big.Int
		Proposer   common.Address
		ProposedAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MetaHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Deposit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Proposer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ProposedAt = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, uint256 _deposit, address _proposer, uint64 _proposedAt)
func (_TaikoL1Client *TaikoL1ClientSession) GetBlock(blockId *big.Int) (struct {
	MetaHash   [32]byte
	Deposit    *big.Int
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	return _TaikoL1Client.Contract.GetBlock(&_TaikoL1Client.CallOpts, blockId)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, uint256 _deposit, address _proposer, uint64 _proposedAt)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetBlock(blockId *big.Int) (struct {
	MetaHash   [32]byte
	Deposit    *big.Int
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	return _TaikoL1Client.Contract.GetBlock(&_TaikoL1Client.CallOpts, blockId)
}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientCaller) GetBlockFee(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getBlockFee")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientSession) GetBlockFee() (uint64, error) {
	return _TaikoL1Client.Contract.GetBlockFee(&_TaikoL1Client.CallOpts)
}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetBlockFee() (uint64, error) {
	return _TaikoL1Client.Contract.GetBlockFee(&_TaikoL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,uint64,uint8,bool))
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
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,uint64,uint8,bool))
func (_TaikoL1Client *TaikoL1ClientSession) GetConfig() (TaikoDataConfig, error) {
	return _TaikoL1Client.Contract.GetConfig(&_TaikoL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,uint64,uint8,bool))
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

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientCaller) GetProofReward(opts *bind.CallOpts, provenAt uint64, proposedAt uint64) (uint64, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getProofReward", provenAt, proposedAt)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientSession) GetProofReward(provenAt uint64, proposedAt uint64) (uint64, error) {
	return _TaikoL1Client.Contract.GetProofReward(&_TaikoL1Client.CallOpts, provenAt, proposedAt)
}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetProofReward(provenAt uint64, proposedAt uint64) (uint64, error) {
	return _TaikoL1Client.Contract.GetProofReward(&_TaikoL1Client.CallOpts, provenAt, proposedAt)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
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
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientSession) GetStateVariables() (TaikoDataStateVariables, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
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
// Solidity: function state() view returns(bytes32 staticRefs, uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reserved71, uint64 __reserved72, uint64 accProposedAt, uint64 accBlockFees, uint64 numBlocks, uint64 nextEthDepositToProcess, uint64 blockFee, uint64 proofTimeIssued, uint64 lastVerifiedBlockId, uint64 __reserved91)
func (_TaikoL1Client *TaikoL1ClientCaller) State(opts *bind.CallOpts) (struct {
	StaticRefs              [32]byte
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	Reserved71              uint64
	Reserved72              uint64
	AccProposedAt           uint64
	AccBlockFees            uint64
	NumBlocks               uint64
	NextEthDepositToProcess uint64
	BlockFee                uint64
	ProofTimeIssued         uint64
	LastVerifiedBlockId     uint64
	Reserved91              uint64
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		StaticRefs              [32]byte
		GenesisHeight           uint64
		GenesisTimestamp        uint64
		Reserved71              uint64
		Reserved72              uint64
		AccProposedAt           uint64
		AccBlockFees            uint64
		NumBlocks               uint64
		NextEthDepositToProcess uint64
		BlockFee                uint64
		ProofTimeIssued         uint64
		LastVerifiedBlockId     uint64
		Reserved91              uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StaticRefs = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.GenesisHeight = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.GenesisTimestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Reserved71 = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Reserved72 = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.AccProposedAt = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.AccBlockFees = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.NumBlocks = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.NextEthDepositToProcess = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.BlockFee = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.ProofTimeIssued = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.LastVerifiedBlockId = *abi.ConvertType(out[11], new(uint64)).(*uint64)
	outstruct.Reserved91 = *abi.ConvertType(out[12], new(uint64)).(*uint64)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(bytes32 staticRefs, uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reserved71, uint64 __reserved72, uint64 accProposedAt, uint64 accBlockFees, uint64 numBlocks, uint64 nextEthDepositToProcess, uint64 blockFee, uint64 proofTimeIssued, uint64 lastVerifiedBlockId, uint64 __reserved91)
func (_TaikoL1Client *TaikoL1ClientSession) State() (struct {
	StaticRefs              [32]byte
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	Reserved71              uint64
	Reserved72              uint64
	AccProposedAt           uint64
	AccBlockFees            uint64
	NumBlocks               uint64
	NextEthDepositToProcess uint64
	BlockFee                uint64
	ProofTimeIssued         uint64
	LastVerifiedBlockId     uint64
	Reserved91              uint64
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(bytes32 staticRefs, uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reserved71, uint64 __reserved72, uint64 accProposedAt, uint64 accBlockFees, uint64 numBlocks, uint64 nextEthDepositToProcess, uint64 blockFee, uint64 proofTimeIssued, uint64 lastVerifiedBlockId, uint64 __reserved91)
func (_TaikoL1Client *TaikoL1ClientCallerSession) State() (struct {
	StaticRefs              [32]byte
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	Reserved71              uint64
	Reserved72              uint64
	AccProposedAt           uint64
	AccBlockFees            uint64
	NumBlocks               uint64
	NextEthDepositToProcess uint64
	BlockFee                uint64
	ProofTimeIssued         uint64
	LastVerifiedBlockId     uint64
	Reserved91              uint64
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
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
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initBlockFee, uint64 _initProofTimeIssued) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _genesisBlockHash [32]byte, _initBlockFee uint64, _initProofTimeIssued uint64) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "init", _addressManager, _genesisBlockHash, _initBlockFee, _initProofTimeIssued)
}

// Init is a paid mutator transaction binding the contract method 0x578c65a4.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initBlockFee, uint64 _initProofTimeIssued) returns()
func (_TaikoL1Client *TaikoL1ClientSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _initBlockFee uint64, _initProofTimeIssued uint64) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Init(&_TaikoL1Client.TransactOpts, _addressManager, _genesisBlockHash, _initBlockFee, _initProofTimeIssued)
}

// Init is a paid mutator transaction binding the contract method 0x578c65a4.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initBlockFee, uint64 _initProofTimeIssued) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _initBlockFee uint64, _initProofTimeIssued uint64) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Init(&_TaikoL1Client.TransactOpts, _addressManager, _genesisBlockHash, _initBlockFee, _initProofTimeIssued)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes input, bytes txList) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
func (_TaikoL1Client *TaikoL1ClientTransactor) ProposeBlock(opts *bind.TransactOpts, input []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proposeBlock", input, txList)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes input, bytes txList) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
func (_TaikoL1Client *TaikoL1ClientSession) ProposeBlock(input []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, input, txList)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes input, bytes txList) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
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
	Id   *big.Int
	Meta TaikoDataBlockMetadata
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0x31cc3eebb5aea55796fe5ba252fe1833fd17d063cb6e83634758fa58e7181535.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
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

// WatchBlockProposed is a free log subscription operation binding the contract event 0x31cc3eebb5aea55796fe5ba252fe1833fd17d063cb6e83634758fa58e7181535.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
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

// ParseBlockProposed is a log parse operation binding the contract event 0x31cc3eebb5aea55796fe5ba252fe1833fd17d063cb6e83634758fa58e7181535.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,uint8,address,(address,uint96)[]) meta)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockProposed(log types.Log) (*TaikoL1ClientBlockProposed, error) {
	event := new(TaikoL1ClientBlockProposed)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProposed", log); err != nil {
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
	Id         *big.Int
	ParentHash [32]byte
	BlockHash  [32]byte
	SignalRoot [32]byte
	Prover     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockProven is a free log retrieval operation binding the contract event 0xd93fde3ea1bb11dcd7a4e66320a05fc5aa63983b6447eff660084c4b1b1b499b.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover)
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

// WatchBlockProven is a free log subscription operation binding the contract event 0xd93fde3ea1bb11dcd7a4e66320a05fc5aa63983b6447eff660084c4b1b1b499b.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover)
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

// ParseBlockProven is a log parse operation binding the contract event 0xd93fde3ea1bb11dcd7a4e66320a05fc5aa63983b6447eff660084c4b1b1b499b.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockProven(log types.Log) (*TaikoL1ClientBlockProven, error) {
	event := new(TaikoL1ClientBlockProven)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProven", log); err != nil {
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
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified is a free log retrieval operation binding the contract event 0x68b82650828a9621868d09dc161400acbe189fa002e3fb7cf9dea5c2c6f928ee.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash)
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

// WatchBlockVerified is a free log subscription operation binding the contract event 0x68b82650828a9621868d09dc161400acbe189fa002e3fb7cf9dea5c2c6f928ee.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash)
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

// ParseBlockVerified is a log parse operation binding the contract event 0x68b82650828a9621868d09dc161400acbe189fa002e3fb7cf9dea5c2c6f928ee.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockVerified(log types.Log) (*TaikoL1ClientBlockVerified, error) {
	event := new(TaikoL1ClientBlockVerified)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockVerified", log); err != nil {
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

// FilterEthDeposited is a free log retrieval operation binding the contract event 0x1c146ddfc652d3b24f1c37ed1cabbb690bf0e198aea624f49211500467182eba.
//
// Solidity: event EthDeposited((address,uint96) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterEthDeposited(opts *bind.FilterOpts) (*TaikoL1ClientEthDepositedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "EthDeposited")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientEthDepositedIterator{contract: _TaikoL1Client.contract, event: "EthDeposited", logs: logs, sub: sub}, nil
}

// WatchEthDeposited is a free log subscription operation binding the contract event 0x1c146ddfc652d3b24f1c37ed1cabbb690bf0e198aea624f49211500467182eba.
//
// Solidity: event EthDeposited((address,uint96) deposit)
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

// ParseEthDeposited is a log parse operation binding the contract event 0x1c146ddfc652d3b24f1c37ed1cabbb690bf0e198aea624f49211500467182eba.
//
// Solidity: event EthDeposited((address,uint96) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseEthDeposited(log types.Log) (*TaikoL1ClientEthDeposited, error) {
	event := new(TaikoL1ClientEthDeposited)
	if err := _TaikoL1Client.contract.UnpackLog(event, "EthDeposited", log); err != nil {
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
