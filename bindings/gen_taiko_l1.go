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
)

// LibUtilsStateVariables is an auto generated low-level Go binding around an user-defined struct.
type LibUtilsStateVariables struct {
	GenesisHeight        uint64
	GenesisTimestamp     uint64
	StatusBits           uint64
	FeeBase              *big.Int
	NextBlockId          uint64
	LastProposedAt       uint64
	AvgBlockTime         uint64
	LatestVerifiedHeight uint64
	LatestVerifiedId     uint64
	AvgProofTime         uint64
}

// TaikoDataBlockMetadata is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataBlockMetadata struct {
	Id           *big.Int
	L1Height     *big.Int
	L1Hash       [32]byte
	Beneficiary  common.Address
	TxListHash   [32]byte
	MixHash      [32]byte
	ExtraData    []byte
	GasLimit     uint64
	Timestamp    uint64
	CommitHeight uint64
	CommitSlot   uint64
}

// TaikoDataConfig is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataConfig struct {
	ChainId                          *big.Int
	MaxNumBlocks                     *big.Int
	BlockHashHistory                 *big.Int
	ZkProofsPerBlock                 *big.Int
	MaxVerificationsPerTx            *big.Int
	CommitConfirmations              *big.Int
	MaxProofsPerForkChoice           *big.Int
	BlockMaxGasLimit                 *big.Int
	MaxTransactionsPerBlock          *big.Int
	MaxBytesPerTxList                *big.Int
	MinTxGasLimit                    *big.Int
	AnchorTxGasLimit                 *big.Int
	SlotSmoothingFactor              *big.Int
	RewardBurnBips                   *big.Int
	ProposerDepositPctg              *big.Int
	FeeBaseMAF                       *big.Int
	BlockTimeMAF                     *big.Int
	ProofTimeMAF                     *big.Int
	RewardMultiplierPctg             uint64
	FeeGracePeriodPctg               uint64
	FeeMaxPeriodPctg                 uint64
	BlockTimeCap                     uint64
	ProofTimeCap                     uint64
	BootstrapDiscountHalvingPeriod   uint64
	InitialUncleDelay                uint64
	ProverRewardRandomizedPercentage uint64
	EnableTokenomics                 bool
	EnablePublicInputsCheck          bool
	EnableProofValidation            bool
	EnableOracleProver               bool
}

// TaikoDataForkChoice is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataForkChoice struct {
	BlockHash [32]byte
	ProvenAt  uint64
	Provers   []common.Address
}

// TaikoDataProposedBlock is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataProposedBlock struct {
	MetaHash   [32]byte
	Deposit    *big.Int
	Proposer   common.Address
	ProposedAt uint64
}

// TaikoL1ClientMetaData contains all meta data concerning the TaikoL1Client contract.
var TaikoL1ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L1_0_FEE_BASE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_CALLDATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_DEST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_GAS_LIMIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_ADDR\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_DATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_LOGS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_STATUS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_TOPICS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_SIG_R\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_SIG_S\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_TYPE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_NUMBER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_NUMBER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_CANNOT_BE_FIRST_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_CIRCUIT_LENGTH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_COMMITTED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_CONFLICT_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_DUP_PROVERS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_EXTRA_DATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_GAS_LIMIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_HALTED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_HALT_CONDITION\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_HALT_CONDITION\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INPUT_SIZE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_METADATA_FIELD\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_META_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_COMMITTED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_FIRST_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_PROOF_LENGTH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SOLO_PROPOSER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_LATE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_PROVERS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ZKP\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"commitSlot\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"name\":\"BlockCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"commitHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"commitSlot\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTaikoData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"BlockProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"halted\",\"type\":\"bool\"}],\"name\":\"Halted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcHash\",\"type\":\"bytes32\"}],\"name\":\"HeaderSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"commitSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"name\":\"commitBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockHashHistory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zkProofsPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVerificationsPerTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitConfirmations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxProofsPerForkChoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTransactionsPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBytesPerTxList\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"anchorTxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slotSmoothingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBurnBips\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposerDepositPctg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBaseMAF\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeMAF\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofTimeMAF\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"rewardMultiplierPctg\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feeGracePeriodPctg\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feeMaxPeriodPctg\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimeCap\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeCap\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"bootstrapDiscountHalvingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initialUncleDelay\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proverRewardRandomizedPercentage\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enableTokenomics\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enablePublicInputsCheck\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableProofValidation\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableOracleProver\",\"type\":\"bool\"}],\"internalType\":\"structTaikoData.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"}],\"name\":\"getForkChoice\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"provers\",\"type\":\"address[]\"}],\"internalType\":\"structTaikoData.ForkChoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestSyncedHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposedAt\",\"type\":\"uint64\"}],\"name\":\"getProofReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getProposedBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"metaHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"proposedAt\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.ProposedBlock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numProvers\",\"type\":\"uint256\"}],\"name\":\"getProverRewardBips\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateVariables\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"statusBits\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nextBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avgBlockTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestVerifiedHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestVerifiedId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avgProofTime\",\"type\":\"uint64\"}],\"internalType\":\"structLibUtils.StateVariables\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getSyncedHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getUncleProofDelay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"toHalt\",\"type\":\"bool\"}],\"name\":\"halt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_feeBase\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"}],\"name\":\"isBlockVerifiable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commitSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"name\":\"isCommitValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isHalted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"proposeBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"proveBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"proveBlockInvalid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"k\",\"type\":\"uint8\"}],\"name\":\"signWithGoldenTouch\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reservedA1\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"statusBits\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nextBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avgBlockTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__avgGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestVerifiedHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestVerifiedId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avgProofTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reservedC1\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"name\":\"verifyBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
	parsed, err := abi.JSON(strings.NewReader(TaikoL1ClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientCaller) GetBlockFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getBlockFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientSession) GetBlockFee() (*big.Int, error) {
	return _TaikoL1Client.Contract.GetBlockFee(&_TaikoL1Client.CallOpts)
}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetBlockFee() (*big.Int, error) {
	return _TaikoL1Client.Contract.GetBlockFee(&_TaikoL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,bool,bool,bool,bool))
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
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,bool,bool,bool,bool))
func (_TaikoL1Client *TaikoL1ClientSession) GetConfig() (TaikoDataConfig, error) {
	return _TaikoL1Client.Contract.GetConfig(&_TaikoL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,bool,bool,bool,bool))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetConfig() (TaikoDataConfig, error) {
	return _TaikoL1Client.Contract.GetConfig(&_TaikoL1Client.CallOpts)
}

// GetForkChoice is a free data retrieval call binding the contract method 0xe00ea1e1.
//
// Solidity: function getForkChoice(uint256 id, bytes32 parentHash) view returns((bytes32,uint64,address[]))
func (_TaikoL1Client *TaikoL1ClientCaller) GetForkChoice(opts *bind.CallOpts, id *big.Int, parentHash [32]byte) (TaikoDataForkChoice, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getForkChoice", id, parentHash)

	if err != nil {
		return *new(TaikoDataForkChoice), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoDataForkChoice)).(*TaikoDataForkChoice)

	return out0, err

}

// GetForkChoice is a free data retrieval call binding the contract method 0xe00ea1e1.
//
// Solidity: function getForkChoice(uint256 id, bytes32 parentHash) view returns((bytes32,uint64,address[]))
func (_TaikoL1Client *TaikoL1ClientSession) GetForkChoice(id *big.Int, parentHash [32]byte) (TaikoDataForkChoice, error) {
	return _TaikoL1Client.Contract.GetForkChoice(&_TaikoL1Client.CallOpts, id, parentHash)
}

// GetForkChoice is a free data retrieval call binding the contract method 0xe00ea1e1.
//
// Solidity: function getForkChoice(uint256 id, bytes32 parentHash) view returns((bytes32,uint64,address[]))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetForkChoice(id *big.Int, parentHash [32]byte) (TaikoDataForkChoice, error) {
	return _TaikoL1Client.Contract.GetForkChoice(&_TaikoL1Client.CallOpts, id, parentHash)
}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCaller) GetLatestSyncedHeader(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getLatestSyncedHeader")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientSession) GetLatestSyncedHeader() ([32]byte, error) {
	return _TaikoL1Client.Contract.GetLatestSyncedHeader(&_TaikoL1Client.CallOpts)
}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetLatestSyncedHeader() ([32]byte, error) {
	return _TaikoL1Client.Contract.GetLatestSyncedHeader(&_TaikoL1Client.CallOpts)
}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint256 reward)
func (_TaikoL1Client *TaikoL1ClientCaller) GetProofReward(opts *bind.CallOpts, provenAt uint64, proposedAt uint64) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getProofReward", provenAt, proposedAt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint256 reward)
func (_TaikoL1Client *TaikoL1ClientSession) GetProofReward(provenAt uint64, proposedAt uint64) (*big.Int, error) {
	return _TaikoL1Client.Contract.GetProofReward(&_TaikoL1Client.CallOpts, provenAt, proposedAt)
}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint256 reward)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetProofReward(provenAt uint64, proposedAt uint64) (*big.Int, error) {
	return _TaikoL1Client.Contract.GetProofReward(&_TaikoL1Client.CallOpts, provenAt, proposedAt)
}

// GetProposedBlock is a free data retrieval call binding the contract method 0x8972b10c.
//
// Solidity: function getProposedBlock(uint256 id) view returns((bytes32,uint256,address,uint64))
func (_TaikoL1Client *TaikoL1ClientCaller) GetProposedBlock(opts *bind.CallOpts, id *big.Int) (TaikoDataProposedBlock, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getProposedBlock", id)

	if err != nil {
		return *new(TaikoDataProposedBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoDataProposedBlock)).(*TaikoDataProposedBlock)

	return out0, err

}

// GetProposedBlock is a free data retrieval call binding the contract method 0x8972b10c.
//
// Solidity: function getProposedBlock(uint256 id) view returns((bytes32,uint256,address,uint64))
func (_TaikoL1Client *TaikoL1ClientSession) GetProposedBlock(id *big.Int) (TaikoDataProposedBlock, error) {
	return _TaikoL1Client.Contract.GetProposedBlock(&_TaikoL1Client.CallOpts, id)
}

// GetProposedBlock is a free data retrieval call binding the contract method 0x8972b10c.
//
// Solidity: function getProposedBlock(uint256 id) view returns((bytes32,uint256,address,uint64))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetProposedBlock(id *big.Int) (TaikoDataProposedBlock, error) {
	return _TaikoL1Client.Contract.GetProposedBlock(&_TaikoL1Client.CallOpts, id)
}

// GetProverRewardBips is a free data retrieval call binding the contract method 0xf22fdc66.
//
// Solidity: function getProverRewardBips(uint256 numProvers) view returns(uint256[])
func (_TaikoL1Client *TaikoL1ClientCaller) GetProverRewardBips(opts *bind.CallOpts, numProvers *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getProverRewardBips", numProvers)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetProverRewardBips is a free data retrieval call binding the contract method 0xf22fdc66.
//
// Solidity: function getProverRewardBips(uint256 numProvers) view returns(uint256[])
func (_TaikoL1Client *TaikoL1ClientSession) GetProverRewardBips(numProvers *big.Int) ([]*big.Int, error) {
	return _TaikoL1Client.Contract.GetProverRewardBips(&_TaikoL1Client.CallOpts, numProvers)
}

// GetProverRewardBips is a free data retrieval call binding the contract method 0xf22fdc66.
//
// Solidity: function getProverRewardBips(uint256 numProvers) view returns(uint256[])
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetProverRewardBips(numProvers *big.Int) ([]*big.Int, error) {
	return _TaikoL1Client.Contract.GetProverRewardBips(&_TaikoL1Client.CallOpts, numProvers)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint256,uint64,uint64,uint64,uint64,uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientCaller) GetStateVariables(opts *bind.CallOpts) (LibUtilsStateVariables, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getStateVariables")

	if err != nil {
		return *new(LibUtilsStateVariables), err
	}

	out0 := *abi.ConvertType(out[0], new(LibUtilsStateVariables)).(*LibUtilsStateVariables)

	return out0, err

}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint256,uint64,uint64,uint64,uint64,uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientSession) GetStateVariables() (LibUtilsStateVariables, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint256,uint64,uint64,uint64,uint64,uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetStateVariables() (LibUtilsStateVariables, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCaller) GetSyncedHeader(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getSyncedHeader", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientSession) GetSyncedHeader(number *big.Int) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetSyncedHeader(&_TaikoL1Client.CallOpts, number)
}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetSyncedHeader(number *big.Int) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetSyncedHeader(&_TaikoL1Client.CallOpts, number)
}

// GetUncleProofDelay is a free data retrieval call binding the contract method 0xf728abaf.
//
// Solidity: function getUncleProofDelay(uint256 blockId) view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientCaller) GetUncleProofDelay(opts *bind.CallOpts, blockId *big.Int) (uint64, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getUncleProofDelay", blockId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetUncleProofDelay is a free data retrieval call binding the contract method 0xf728abaf.
//
// Solidity: function getUncleProofDelay(uint256 blockId) view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientSession) GetUncleProofDelay(blockId *big.Int) (uint64, error) {
	return _TaikoL1Client.Contract.GetUncleProofDelay(&_TaikoL1Client.CallOpts, blockId)
}

// GetUncleProofDelay is a free data retrieval call binding the contract method 0xf728abaf.
//
// Solidity: function getUncleProofDelay(uint256 blockId) view returns(uint64)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetUncleProofDelay(blockId *big.Int) (uint64, error) {
	return _TaikoL1Client.Contract.GetUncleProofDelay(&_TaikoL1Client.CallOpts, blockId)
}

// IsBlockVerifiable is a free data retrieval call binding the contract method 0x46a06c27.
//
// Solidity: function isBlockVerifiable(uint256 blockId, bytes32 parentHash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) IsBlockVerifiable(opts *bind.CallOpts, blockId *big.Int, parentHash [32]byte) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "isBlockVerifiable", blockId, parentHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlockVerifiable is a free data retrieval call binding the contract method 0x46a06c27.
//
// Solidity: function isBlockVerifiable(uint256 blockId, bytes32 parentHash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) IsBlockVerifiable(blockId *big.Int, parentHash [32]byte) (bool, error) {
	return _TaikoL1Client.Contract.IsBlockVerifiable(&_TaikoL1Client.CallOpts, blockId, parentHash)
}

// IsBlockVerifiable is a free data retrieval call binding the contract method 0x46a06c27.
//
// Solidity: function isBlockVerifiable(uint256 blockId, bytes32 parentHash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) IsBlockVerifiable(blockId *big.Int, parentHash [32]byte) (bool, error) {
	return _TaikoL1Client.Contract.IsBlockVerifiable(&_TaikoL1Client.CallOpts, blockId, parentHash)
}

// IsCommitValid is a free data retrieval call binding the contract method 0x340d9599.
//
// Solidity: function isCommitValid(uint256 commitSlot, uint256 commitHeight, bytes32 commitHash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) IsCommitValid(opts *bind.CallOpts, commitSlot *big.Int, commitHeight *big.Int, commitHash [32]byte) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "isCommitValid", commitSlot, commitHeight, commitHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommitValid is a free data retrieval call binding the contract method 0x340d9599.
//
// Solidity: function isCommitValid(uint256 commitSlot, uint256 commitHeight, bytes32 commitHash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) IsCommitValid(commitSlot *big.Int, commitHeight *big.Int, commitHash [32]byte) (bool, error) {
	return _TaikoL1Client.Contract.IsCommitValid(&_TaikoL1Client.CallOpts, commitSlot, commitHeight, commitHash)
}

// IsCommitValid is a free data retrieval call binding the contract method 0x340d9599.
//
// Solidity: function isCommitValid(uint256 commitSlot, uint256 commitHeight, bytes32 commitHash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) IsCommitValid(commitSlot *big.Int, commitHeight *big.Int, commitHash [32]byte) (bool, error) {
	return _TaikoL1Client.Contract.IsCommitValid(&_TaikoL1Client.CallOpts, commitSlot, commitHeight, commitHash)
}

// IsHalted is a free data retrieval call binding the contract method 0xc7ff1584.
//
// Solidity: function isHalted() view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) IsHalted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "isHalted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHalted is a free data retrieval call binding the contract method 0xc7ff1584.
//
// Solidity: function isHalted() view returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) IsHalted() (bool, error) {
	return _TaikoL1Client.Contract.IsHalted(&_TaikoL1Client.CallOpts)
}

// IsHalted is a free data retrieval call binding the contract method 0xc7ff1584.
//
// Solidity: function isHalted() view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) IsHalted() (bool, error) {
	return _TaikoL1Client.Contract.IsHalted(&_TaikoL1Client.CallOpts)
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

// Resolve is a free data retrieval call binding the contract method 0x0ca4dffd.
//
// Solidity: function resolve(string name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCaller) Resolve(opts *bind.CallOpts, name string, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "resolve", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x0ca4dffd.
//
// Solidity: function resolve(string name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientSession) Resolve(name string, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve(&_TaikoL1Client.CallOpts, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x0ca4dffd.
//
// Solidity: function resolve(string name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Resolve(name string, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve(&_TaikoL1Client.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0x1be2bfa7.
//
// Solidity: function resolve(uint256 chainId, string name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCaller) Resolve0(opts *bind.CallOpts, chainId *big.Int, name string, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "resolve0", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0x1be2bfa7.
//
// Solidity: function resolve(uint256 chainId, string name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientSession) Resolve0(chainId *big.Int, name string, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve0(&_TaikoL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0x1be2bfa7.
//
// Solidity: function resolve(uint256 chainId, string name, bool allowZeroAddress) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Resolve0(chainId *big.Int, name string, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve0(&_TaikoL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// SignWithGoldenTouch is a free data retrieval call binding the contract method 0xdadec12a.
//
// Solidity: function signWithGoldenTouch(bytes32 hash, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL1Client *TaikoL1ClientCaller) SignWithGoldenTouch(opts *bind.CallOpts, hash [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "signWithGoldenTouch", hash, k)

	outstruct := new(struct {
		V uint8
		R *big.Int
		S *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.S = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SignWithGoldenTouch is a free data retrieval call binding the contract method 0xdadec12a.
//
// Solidity: function signWithGoldenTouch(bytes32 hash, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL1Client *TaikoL1ClientSession) SignWithGoldenTouch(hash [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _TaikoL1Client.Contract.SignWithGoldenTouch(&_TaikoL1Client.CallOpts, hash, k)
}

// SignWithGoldenTouch is a free data retrieval call binding the contract method 0xdadec12a.
//
// Solidity: function signWithGoldenTouch(bytes32 hash, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL1Client *TaikoL1ClientCallerSession) SignWithGoldenTouch(hash [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _TaikoL1Client.Contract.SignWithGoldenTouch(&_TaikoL1Client.CallOpts, hash, k)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reservedA1, uint64 statusBits, uint256 feeBase, uint64 nextBlockId, uint64 lastProposedAt, uint64 avgBlockTime, uint64 __avgGasLimit, uint64 latestVerifiedHeight, uint64 latestVerifiedId, uint64 avgProofTime, uint64 __reservedC1)
func (_TaikoL1Client *TaikoL1ClientCaller) State(opts *bind.CallOpts) (struct {
	GenesisHeight        uint64
	GenesisTimestamp     uint64
	ReservedA1           uint64
	StatusBits           uint64
	FeeBase              *big.Int
	NextBlockId          uint64
	LastProposedAt       uint64
	AvgBlockTime         uint64
	AvgGasLimit          uint64
	LatestVerifiedHeight uint64
	LatestVerifiedId     uint64
	AvgProofTime         uint64
	ReservedC1           uint64
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		GenesisHeight        uint64
		GenesisTimestamp     uint64
		ReservedA1           uint64
		StatusBits           uint64
		FeeBase              *big.Int
		NextBlockId          uint64
		LastProposedAt       uint64
		AvgBlockTime         uint64
		AvgGasLimit          uint64
		LatestVerifiedHeight uint64
		LatestVerifiedId     uint64
		AvgProofTime         uint64
		ReservedC1           uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GenesisHeight = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.GenesisTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ReservedA1 = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.StatusBits = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.FeeBase = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.NextBlockId = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.LastProposedAt = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.AvgBlockTime = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.AvgGasLimit = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.LatestVerifiedHeight = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.LatestVerifiedId = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.AvgProofTime = *abi.ConvertType(out[11], new(uint64)).(*uint64)
	outstruct.ReservedC1 = *abi.ConvertType(out[12], new(uint64)).(*uint64)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reservedA1, uint64 statusBits, uint256 feeBase, uint64 nextBlockId, uint64 lastProposedAt, uint64 avgBlockTime, uint64 __avgGasLimit, uint64 latestVerifiedHeight, uint64 latestVerifiedId, uint64 avgProofTime, uint64 __reservedC1)
func (_TaikoL1Client *TaikoL1ClientSession) State() (struct {
	GenesisHeight        uint64
	GenesisTimestamp     uint64
	ReservedA1           uint64
	StatusBits           uint64
	FeeBase              *big.Int
	NextBlockId          uint64
	LastProposedAt       uint64
	AvgBlockTime         uint64
	AvgGasLimit          uint64
	LatestVerifiedHeight uint64
	LatestVerifiedId     uint64
	AvgProofTime         uint64
	ReservedC1           uint64
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reservedA1, uint64 statusBits, uint256 feeBase, uint64 nextBlockId, uint64 lastProposedAt, uint64 avgBlockTime, uint64 __avgGasLimit, uint64 latestVerifiedHeight, uint64 latestVerifiedId, uint64 avgProofTime, uint64 __reservedC1)
func (_TaikoL1Client *TaikoL1ClientCallerSession) State() (struct {
	GenesisHeight        uint64
	GenesisTimestamp     uint64
	ReservedA1           uint64
	StatusBits           uint64
	FeeBase              *big.Int
	NextBlockId          uint64
	LastProposedAt       uint64
	AvgBlockTime         uint64
	AvgGasLimit          uint64
	LatestVerifiedHeight uint64
	LatestVerifiedId     uint64
	AvgProofTime         uint64
	ReservedC1           uint64
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x7e7a262c.
//
// Solidity: function commitBlock(uint64 commitSlot, bytes32 commitHash) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) CommitBlock(opts *bind.TransactOpts, commitSlot uint64, commitHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "commitBlock", commitSlot, commitHash)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x7e7a262c.
//
// Solidity: function commitBlock(uint64 commitSlot, bytes32 commitHash) returns()
func (_TaikoL1Client *TaikoL1ClientSession) CommitBlock(commitSlot uint64, commitHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.CommitBlock(&_TaikoL1Client.TransactOpts, commitSlot, commitHash)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x7e7a262c.
//
// Solidity: function commitBlock(uint64 commitSlot, bytes32 commitHash) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) CommitBlock(commitSlot uint64, commitHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.CommitBlock(&_TaikoL1Client.TransactOpts, commitSlot, commitHash)
}

// Halt is a paid mutator transaction binding the contract method 0x96a38d22.
//
// Solidity: function halt(bool toHalt) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) Halt(opts *bind.TransactOpts, toHalt bool) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "halt", toHalt)
}

// Halt is a paid mutator transaction binding the contract method 0x96a38d22.
//
// Solidity: function halt(bool toHalt) returns()
func (_TaikoL1Client *TaikoL1ClientSession) Halt(toHalt bool) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Halt(&_TaikoL1Client.TransactOpts, toHalt)
}

// Halt is a paid mutator transaction binding the contract method 0x96a38d22.
//
// Solidity: function halt(bool toHalt) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) Halt(toHalt bool) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Halt(&_TaikoL1Client.TransactOpts, toHalt)
}

// Init is a paid mutator transaction binding the contract method 0x9c5e9f06.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint256 _feeBase) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _genesisBlockHash [32]byte, _feeBase *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "init", _addressManager, _genesisBlockHash, _feeBase)
}

// Init is a paid mutator transaction binding the contract method 0x9c5e9f06.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint256 _feeBase) returns()
func (_TaikoL1Client *TaikoL1ClientSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _feeBase *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Init(&_TaikoL1Client.TransactOpts, _addressManager, _genesisBlockHash, _feeBase)
}

// Init is a paid mutator transaction binding the contract method 0x9c5e9f06.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint256 _feeBase) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _feeBase *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Init(&_TaikoL1Client.TransactOpts, _addressManager, _genesisBlockHash, _feeBase)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xa043dbdf.
//
// Solidity: function proposeBlock(bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) ProposeBlock(opts *bind.TransactOpts, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proposeBlock", inputs)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xa043dbdf.
//
// Solidity: function proposeBlock(bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientSession) ProposeBlock(inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, inputs)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xa043dbdf.
//
// Solidity: function proposeBlock(bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProposeBlock(inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, inputs)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x8ed7b3be.
//
// Solidity: function proveBlock(uint256 blockId, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) ProveBlock(opts *bind.TransactOpts, blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proveBlock", blockId, inputs)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x8ed7b3be.
//
// Solidity: function proveBlock(uint256 blockId, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientSession) ProveBlock(blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlock(&_TaikoL1Client.TransactOpts, blockId, inputs)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x8ed7b3be.
//
// Solidity: function proveBlock(uint256 blockId, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProveBlock(blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlock(&_TaikoL1Client.TransactOpts, blockId, inputs)
}

// ProveBlockInvalid is a paid mutator transaction binding the contract method 0xa279cec7.
//
// Solidity: function proveBlockInvalid(uint256 blockId, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) ProveBlockInvalid(opts *bind.TransactOpts, blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proveBlockInvalid", blockId, inputs)
}

// ProveBlockInvalid is a paid mutator transaction binding the contract method 0xa279cec7.
//
// Solidity: function proveBlockInvalid(uint256 blockId, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientSession) ProveBlockInvalid(blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlockInvalid(&_TaikoL1Client.TransactOpts, blockId, inputs)
}

// ProveBlockInvalid is a paid mutator transaction binding the contract method 0xa279cec7.
//
// Solidity: function proveBlockInvalid(uint256 blockId, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProveBlockInvalid(blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlockInvalid(&_TaikoL1Client.TransactOpts, blockId, inputs)
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

// TaikoL1ClientBlockCommittedIterator is returned from FilterBlockCommitted and is used to iterate over the raw logs and unpacked data for BlockCommitted events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockCommittedIterator struct {
	Event *TaikoL1ClientBlockCommitted // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientBlockCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockCommitted)
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
		it.Event = new(TaikoL1ClientBlockCommitted)
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
func (it *TaikoL1ClientBlockCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockCommitted represents a BlockCommitted event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockCommitted struct {
	CommitSlot uint64
	CommitHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockCommitted is a free log retrieval operation binding the contract event 0x51264991e22d808f3bcbb1cbffa82b752eae327c24055259a5c455c0aa5b7136.
//
// Solidity: event BlockCommitted(uint64 commitSlot, bytes32 commitHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockCommitted(opts *bind.FilterOpts) (*TaikoL1ClientBlockCommittedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockCommitted")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockCommittedIterator{contract: _TaikoL1Client.contract, event: "BlockCommitted", logs: logs, sub: sub}, nil
}

// WatchBlockCommitted is a free log subscription operation binding the contract event 0x51264991e22d808f3bcbb1cbffa82b752eae327c24055259a5c455c0aa5b7136.
//
// Solidity: event BlockCommitted(uint64 commitSlot, bytes32 commitHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockCommitted(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockCommitted) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockCommitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockCommitted)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockCommitted", log); err != nil {
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

// ParseBlockCommitted is a log parse operation binding the contract event 0x51264991e22d808f3bcbb1cbffa82b752eae327c24055259a5c455c0aa5b7136.
//
// Solidity: event BlockCommitted(uint64 commitSlot, bytes32 commitHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockCommitted(log types.Log) (*TaikoL1ClientBlockCommitted, error) {
	event := new(TaikoL1ClientBlockCommitted)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockCommitted", log); err != nil {
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
	Id   *big.Int
	Meta TaikoDataBlockMetadata
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0x344fc5d5f80c4a29bd7e06ad7c28a0c8c9c08d682129da3a31936d5982e4f044.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint256,uint256,bytes32,address,bytes32,bytes32,bytes,uint64,uint64,uint64,uint64) meta)
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

// WatchBlockProposed is a free log subscription operation binding the contract event 0x344fc5d5f80c4a29bd7e06ad7c28a0c8c9c08d682129da3a31936d5982e4f044.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint256,uint256,bytes32,address,bytes32,bytes32,bytes,uint64,uint64,uint64,uint64) meta)
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

// ParseBlockProposed is a log parse operation binding the contract event 0x344fc5d5f80c4a29bd7e06ad7c28a0c8c9c08d682129da3a31936d5982e4f044.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint256,uint256,bytes32,address,bytes32,bytes32,bytes,uint64,uint64,uint64,uint64) meta)
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
	Timestamp  uint64
	ProvenAt   uint64
	Prover     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockProven is a free log retrieval operation binding the contract event 0xf66e10e397da1b4c87429bd76e1d3b85fd417c38c5d43c8ea3e2ec45935bada0.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, uint64 timestamp, uint64 provenAt, address prover)
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

// WatchBlockProven is a free log subscription operation binding the contract event 0xf66e10e397da1b4c87429bd76e1d3b85fd417c38c5d43c8ea3e2ec45935bada0.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, uint64 timestamp, uint64 provenAt, address prover)
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

// ParseBlockProven is a log parse operation binding the contract event 0xf66e10e397da1b4c87429bd76e1d3b85fd417c38c5d43c8ea3e2ec45935bada0.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, uint64 timestamp, uint64 provenAt, address prover)
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

// TaikoL1ClientHaltedIterator is returned from FilterHalted and is used to iterate over the raw logs and unpacked data for Halted events raised by the TaikoL1Client contract.
type TaikoL1ClientHaltedIterator struct {
	Event *TaikoL1ClientHalted // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientHaltedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientHalted)
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
		it.Event = new(TaikoL1ClientHalted)
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
func (it *TaikoL1ClientHaltedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientHaltedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientHalted represents a Halted event raised by the TaikoL1Client contract.
type TaikoL1ClientHalted struct {
	Halted bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHalted is a free log retrieval operation binding the contract event 0x92333b0b676476985757350034668cb9ee247674ac7a7479de10cd761381f733.
//
// Solidity: event Halted(bool halted)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterHalted(opts *bind.FilterOpts) (*TaikoL1ClientHaltedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "Halted")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientHaltedIterator{contract: _TaikoL1Client.contract, event: "Halted", logs: logs, sub: sub}, nil
}

// WatchHalted is a free log subscription operation binding the contract event 0x92333b0b676476985757350034668cb9ee247674ac7a7479de10cd761381f733.
//
// Solidity: event Halted(bool halted)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchHalted(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientHalted) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "Halted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientHalted)
				if err := _TaikoL1Client.contract.UnpackLog(event, "Halted", log); err != nil {
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

// ParseHalted is a log parse operation binding the contract event 0x92333b0b676476985757350034668cb9ee247674ac7a7479de10cd761381f733.
//
// Solidity: event Halted(bool halted)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseHalted(log types.Log) (*TaikoL1ClientHalted, error) {
	event := new(TaikoL1ClientHalted)
	if err := _TaikoL1Client.contract.UnpackLog(event, "Halted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientHeaderSyncedIterator is returned from FilterHeaderSynced and is used to iterate over the raw logs and unpacked data for HeaderSynced events raised by the TaikoL1Client contract.
type TaikoL1ClientHeaderSyncedIterator struct {
	Event *TaikoL1ClientHeaderSynced // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientHeaderSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientHeaderSynced)
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
		it.Event = new(TaikoL1ClientHeaderSynced)
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
func (it *TaikoL1ClientHeaderSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientHeaderSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientHeaderSynced represents a HeaderSynced event raised by the TaikoL1Client contract.
type TaikoL1ClientHeaderSynced struct {
	SrcHeight *big.Int
	SrcHash   [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHeaderSynced is a free log retrieval operation binding the contract event 0x58313b60ec6c5bfc381e52f0de3ede0faac3cdffea26f7d6bcc3d09b61018691.
//
// Solidity: event HeaderSynced(uint256 indexed srcHeight, bytes32 srcHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterHeaderSynced(opts *bind.FilterOpts, srcHeight []*big.Int) (*TaikoL1ClientHeaderSyncedIterator, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "HeaderSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientHeaderSyncedIterator{contract: _TaikoL1Client.contract, event: "HeaderSynced", logs: logs, sub: sub}, nil
}

// WatchHeaderSynced is a free log subscription operation binding the contract event 0x58313b60ec6c5bfc381e52f0de3ede0faac3cdffea26f7d6bcc3d09b61018691.
//
// Solidity: event HeaderSynced(uint256 indexed srcHeight, bytes32 srcHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchHeaderSynced(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientHeaderSynced, srcHeight []*big.Int) (event.Subscription, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "HeaderSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientHeaderSynced)
				if err := _TaikoL1Client.contract.UnpackLog(event, "HeaderSynced", log); err != nil {
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

// ParseHeaderSynced is a log parse operation binding the contract event 0x58313b60ec6c5bfc381e52f0de3ede0faac3cdffea26f7d6bcc3d09b61018691.
//
// Solidity: event HeaderSynced(uint256 indexed srcHeight, bytes32 srcHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseHeaderSynced(log types.Log) (*TaikoL1ClientHeaderSynced, error) {
	event := new(TaikoL1ClientHeaderSynced)
	if err := _TaikoL1Client.contract.UnpackLog(event, "HeaderSynced", log); err != nil {
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
