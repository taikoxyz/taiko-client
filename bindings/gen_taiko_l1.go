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

// ITierProviderTier is an auto generated low-level Go binding around an user-defined struct.
type ITierProviderTier struct {
	VerifierName   [32]byte
	ValidityBond   *big.Int
	ContestBond    *big.Int
	CooldownWindow *big.Int
	ProvingWindow  uint16
}

// TaikoDataBlock is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataBlock struct {
	MetaHash             [32]byte
	AssignedProver       common.Address
	LivenessBond         *big.Int
	BlockId              uint64
	ProposedAt           uint64
	NextTransitionId     uint32
	VerifiedTransitionId uint32
	MinTier              uint16
	Reserved             [7][32]byte
}

// TaikoDataBlockMetadata is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataBlockMetadata struct {
	L1Hash            [32]byte
	Difficulty        [32]byte
	TxListHash        [32]byte
	ExtraData         [32]byte
	Id                uint64
	Timestamp         uint64
	L1Height          uint64
	GasLimit          uint32
	Coinbase          common.Address
	DepositsProcessed []TaikoDataEthDeposit
}

// TaikoDataConfig is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataConfig struct {
	ChainId                    *big.Int
	RelaySignalRoot            bool
	BlockMaxProposals          uint64
	BlockRingBufferSize        uint64
	BlockMaxVerificationsPerTx uint64
	BlockMaxGasLimit           uint32
	BlockFeeBaseGas            uint32
	BlockMaxTxListBytes        *big.Int
	ProposerRewardPerL1Block   *big.Int
	ProposerRewardMax          *big.Int
	ProposerRewardPoolPctg     uint8
	LivenessBond               *big.Int
	EthDepositRingBufferSize   *big.Int
	EthDepositMinCountPerBlock uint64
	EthDepositMaxCountPerBlock uint64
	EthDepositMinAmount        *big.Int
	EthDepositMaxAmount        *big.Int
	EthDepositGas              *big.Int
	EthDepositMaxFee           *big.Int
}

// TaikoDataEthDeposit is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataEthDeposit struct {
	Recipient common.Address
	Amount    *big.Int
	Id        uint64
}

// TaikoDataSlotA is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataSlotA struct {
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	NumEthDeposits          uint64
	NextEthDepositToProcess uint64
}

// TaikoDataSlotB is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataSlotB struct {
	NumBlocks               uint64
	NextEthDepositToProcess uint64
	LastVerifiedAt          uint64
	LastVerifiedBlockId     uint64
}

// TaikoDataSlotC is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataSlotC struct {
	AccumulatedReward  *big.Int
	LastProposedHeight uint64
}

// TaikoDataStateVariables is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataStateVariables struct {
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	NextEthDepositToProcess uint64
	NumEthDeposits          uint64
	NumBlocks               uint64
	LastVerifiedBlockId     uint64
}

// TaikoDataTransition is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataTransition struct {
	Key          [32]byte
	BlockHash    [32]byte
	SignalRoot   [32]byte
	Prover       common.Address
	ValidityBond *big.Int
	Contester    common.Address
	ContestBond  *big.Int
	Timestamp    uint64
	Tier         uint16
	Reserved     [4][32]byte
}

// TaikoL1ClientMetaData contains all meta data concerning the TaikoL1Client contract.
var TaikoL1ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L1_ALREADY_CONTESTED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ALREADY_CONTESTED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ALREADY_PROVED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ALREADY_PROVED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ASSIGNED_PROVER_NOT_ALLOWED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ASSIGNED_PROVER_NOT_ALLOWED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ASSIGNMENT_EXPIRED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ASSIGNMENT_EXPIRED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ASSIGNMENT_INSUFFICIENT_FEE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ASSIGNMENT_INSUFFICIENT_FEE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ASSIGNMENT_INVALID_PARAMS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ASSIGNMENT_INVALID_PARAMS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ASSIGNMENT_INVALID_SIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ASSIGNMENT_INVALID_SIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROPOSER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROVER_SIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_TIER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_TIER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_ASSIGNED_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_ASSIGNED_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TIER_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TIER_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TRANSITION_ID_ZERO\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TRANSITION_ID_ZERO\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TRANSITION_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TRANSITION_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TXLIST_INVALID_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TXLIST_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TXLIST_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TXLIST_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TXLIST_TOO_LARGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TXLIST_TOO_LARGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_UNAUTHORIZED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_UNAUTHORIZED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_UNEXPECTED_TRANSITION_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_UNEXPECTED_TRANSITION_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_UNEXPECTED_TRANSITION_TIER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_UNEXPECTED_TRANSITION_TIER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addressManager\",\"type\":\"address\"}],\"name\":\"AddressManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignedProver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"livenessBond\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proverFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minTier\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"difficulty\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraData\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structTaikoData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignedProver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"livenessBond\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proverFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minTier\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"difficulty\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraData\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structTaikoData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignedProver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assignedProver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"srcHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"srcHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTaikoData.EthDeposit\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"EthDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTaikoData.EthDeposit\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"EthDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"minted\",\"type\":\"bool\"}],\"name\":\"TokenCredited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"minted\",\"type\":\"bool\"}],\"name\":\"TokenCredited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fromLocalBalance\",\"type\":\"bool\"}],\"name\":\"TokenDebited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fromLocalBalance\",\"type\":\"bool\"}],\"name\":\"TokenDebited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawnByOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawnByOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"contestBond\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"TransitionContested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"contestBond\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"TransitionContested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"validityBond\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"TransitionProved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"validityBond\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"TransitionProved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"canDepositEthToL2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositEtherToL2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositTaikoToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockId\",\"type\":\"uint64\"}],\"name\":\"getBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"metaHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"assignedProver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"livenessBond\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"blockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"nextTransitionId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"verifiedTransitionId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"minTier\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[7]\",\"name\":\"__reserved\",\"type\":\"bytes32[7]\"}],\"internalType\":\"structTaikoData.Block\",\"name\":\"blk\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"relaySignalRoot\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxProposals\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockRingBufferSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxVerificationsPerTx\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockFeeBaseGas\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"blockMaxTxListBytes\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"proposerRewardPerL1Block\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"proposerRewardMax\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"proposerRewardPoolPctg\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"livenessBond\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositRingBufferSize\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"ethDepositMinCountPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ethDepositMaxCountPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"ethDepositMinAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"ethDepositMaxAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositMaxFee\",\"type\":\"uint256\"}],\"internalType\":\"structTaikoData.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockId\",\"type\":\"uint64\"}],\"name\":\"getCrossChainBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockId\",\"type\":\"uint64\"}],\"name\":\"getCrossChainSignalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rand\",\"type\":\"uint256\"}],\"name\":\"getMinTier\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateVariables\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numEthDeposits\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.StateVariables\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getTaikoTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tierId\",\"type\":\"uint16\"}],\"name\":\"getTier\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"verifierName\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"validityBond\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"contestBond\",\"type\":\"uint96\"},{\"internalType\":\"uint24\",\"name\":\"cooldownWindow\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"provingWindow\",\"type\":\"uint16\"}],\"internalType\":\"structITierProvider.Tier\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTierIds\",\"outputs\":[{\"internalType\":\"uint16[]\",\"name\":\"\",\"type\":\"uint16[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"}],\"name\":\"getTransition\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"validityBond\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"contester\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"contestBond\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[4]\",\"name\":\"__reserved\",\"type\":\"bytes32[4]\"}],\"internalType\":\"structTaikoData.Transition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ownerWithdrawTaikoToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"assignment\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txList\",\"type\":\"bytes\"}],\"name\":\"proposeBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"difficulty\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraData\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"}],\"internalType\":\"structTaikoData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"proveBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numEthDeposits\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.SlotA\",\"name\":\"slotA\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.SlotB\",\"name\":\"slotB\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"accumulatedReward\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"lastProposedHeight\",\"type\":\"uint64\"}],\"internalType\":\"structTaikoData.SlotC\",\"name\":\"slotC\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"maxBlocks\",\"type\":\"uint64\"}],\"name\":\"verifyBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTaikoToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// CanDepositEthToL2 is a free data retrieval call binding the contract method 0xcf151d9a.
//
// Solidity: function canDepositEthToL2(uint256 amount) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) CanDepositEthToL2(opts *bind.CallOpts, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "canDepositEthToL2", amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDepositEthToL2 is a free data retrieval call binding the contract method 0xcf151d9a.
//
// Solidity: function canDepositEthToL2(uint256 amount) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) CanDepositEthToL2(amount *big.Int) (bool, error) {
	return _TaikoL1Client.Contract.CanDepositEthToL2(&_TaikoL1Client.CallOpts, amount)
}

// CanDepositEthToL2 is a free data retrieval call binding the contract method 0xcf151d9a.
//
// Solidity: function canDepositEthToL2(uint256 amount) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) CanDepositEthToL2(amount *big.Int) (bool, error) {
	return _TaikoL1Client.Contract.CanDepositEthToL2(&_TaikoL1Client.CallOpts, amount)
}

// GetBlock is a free data retrieval call binding the contract method 0x5fa15e79.
//
// Solidity: function getBlock(uint64 blockId) view returns((bytes32,address,uint96,uint64,uint64,uint32,uint32,uint16,bytes32[7]) blk)
func (_TaikoL1Client *TaikoL1ClientCaller) GetBlock(opts *bind.CallOpts, blockId uint64) (TaikoDataBlock, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getBlock", blockId)

	if err != nil {
		return *new(TaikoDataBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoDataBlock)).(*TaikoDataBlock)

	return out0, err

}

// GetBlock is a free data retrieval call binding the contract method 0x5fa15e79.
//
// Solidity: function getBlock(uint64 blockId) view returns((bytes32,address,uint96,uint64,uint64,uint32,uint32,uint16,bytes32[7]) blk)
func (_TaikoL1Client *TaikoL1ClientSession) GetBlock(blockId uint64) (TaikoDataBlock, error) {
	return _TaikoL1Client.Contract.GetBlock(&_TaikoL1Client.CallOpts, blockId)
}

// GetBlock is a free data retrieval call binding the contract method 0x5fa15e79.
//
// Solidity: function getBlock(uint64 blockId) view returns((bytes32,address,uint96,uint64,uint64,uint32,uint32,uint16,bytes32[7]) blk)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetBlock(blockId uint64) (TaikoDataBlock, error) {
	return _TaikoL1Client.Contract.GetBlock(&_TaikoL1Client.CallOpts, blockId)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,bool,uint64,uint64,uint64,uint32,uint32,uint24,uint128,uint128,uint8,uint96,uint256,uint64,uint64,uint96,uint96,uint256,uint256))
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
// Solidity: function getConfig() pure returns((uint256,bool,uint64,uint64,uint64,uint32,uint32,uint24,uint128,uint128,uint8,uint96,uint256,uint64,uint64,uint96,uint96,uint256,uint256))
func (_TaikoL1Client *TaikoL1ClientSession) GetConfig() (TaikoDataConfig, error) {
	return _TaikoL1Client.Contract.GetConfig(&_TaikoL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,bool,uint64,uint64,uint64,uint32,uint32,uint24,uint128,uint128,uint8,uint96,uint256,uint64,uint64,uint96,uint96,uint256,uint256))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetConfig() (TaikoDataConfig, error) {
	return _TaikoL1Client.Contract.GetConfig(&_TaikoL1Client.CallOpts)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbdd6bc36.
//
// Solidity: function getCrossChainBlockHash(uint64 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCaller) GetCrossChainBlockHash(opts *bind.CallOpts, blockId uint64) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getCrossChainBlockHash", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbdd6bc36.
//
// Solidity: function getCrossChainBlockHash(uint64 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientSession) GetCrossChainBlockHash(blockId uint64) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetCrossChainBlockHash(&_TaikoL1Client.CallOpts, blockId)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbdd6bc36.
//
// Solidity: function getCrossChainBlockHash(uint64 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetCrossChainBlockHash(blockId uint64) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetCrossChainBlockHash(&_TaikoL1Client.CallOpts, blockId)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0x0599d294.
//
// Solidity: function getCrossChainSignalRoot(uint64 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCaller) GetCrossChainSignalRoot(opts *bind.CallOpts, blockId uint64) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getCrossChainSignalRoot", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0x0599d294.
//
// Solidity: function getCrossChainSignalRoot(uint64 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientSession) GetCrossChainSignalRoot(blockId uint64) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetCrossChainSignalRoot(&_TaikoL1Client.CallOpts, blockId)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0x0599d294.
//
// Solidity: function getCrossChainSignalRoot(uint64 blockId) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetCrossChainSignalRoot(blockId uint64) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetCrossChainSignalRoot(&_TaikoL1Client.CallOpts, blockId)
}

// GetMinTier is a free data retrieval call binding the contract method 0x59ab4e23.
//
// Solidity: function getMinTier(uint256 rand) view returns(uint16)
func (_TaikoL1Client *TaikoL1ClientCaller) GetMinTier(opts *bind.CallOpts, rand *big.Int) (uint16, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getMinTier", rand)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetMinTier is a free data retrieval call binding the contract method 0x59ab4e23.
//
// Solidity: function getMinTier(uint256 rand) view returns(uint16)
func (_TaikoL1Client *TaikoL1ClientSession) GetMinTier(rand *big.Int) (uint16, error) {
	return _TaikoL1Client.Contract.GetMinTier(&_TaikoL1Client.CallOpts, rand)
}

// GetMinTier is a free data retrieval call binding the contract method 0x59ab4e23.
//
// Solidity: function getMinTier(uint256 rand) view returns(uint16)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetMinTier(rand *big.Int) (uint16, error) {
	return _TaikoL1Client.Contract.GetMinTier(&_TaikoL1Client.CallOpts, rand)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64))
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
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientSession) GetStateVariables() (TaikoDataStateVariables, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetStateVariables() (TaikoDataStateVariables, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetTaikoTokenBalance is a free data retrieval call binding the contract method 0x8dff9cea.
//
// Solidity: function getTaikoTokenBalance(address user) view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientCaller) GetTaikoTokenBalance(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getTaikoTokenBalance", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTaikoTokenBalance is a free data retrieval call binding the contract method 0x8dff9cea.
//
// Solidity: function getTaikoTokenBalance(address user) view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientSession) GetTaikoTokenBalance(user common.Address) (*big.Int, error) {
	return _TaikoL1Client.Contract.GetTaikoTokenBalance(&_TaikoL1Client.CallOpts, user)
}

// GetTaikoTokenBalance is a free data retrieval call binding the contract method 0x8dff9cea.
//
// Solidity: function getTaikoTokenBalance(address user) view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetTaikoTokenBalance(user common.Address) (*big.Int, error) {
	return _TaikoL1Client.Contract.GetTaikoTokenBalance(&_TaikoL1Client.CallOpts, user)
}

// GetTier is a free data retrieval call binding the contract method 0x576c3de7.
//
// Solidity: function getTier(uint16 tierId) view returns((bytes32,uint96,uint96,uint24,uint16))
func (_TaikoL1Client *TaikoL1ClientCaller) GetTier(opts *bind.CallOpts, tierId uint16) (ITierProviderTier, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getTier", tierId)

	if err != nil {
		return *new(ITierProviderTier), err
	}

	out0 := *abi.ConvertType(out[0], new(ITierProviderTier)).(*ITierProviderTier)

	return out0, err

}

// GetTier is a free data retrieval call binding the contract method 0x576c3de7.
//
// Solidity: function getTier(uint16 tierId) view returns((bytes32,uint96,uint96,uint24,uint16))
func (_TaikoL1Client *TaikoL1ClientSession) GetTier(tierId uint16) (ITierProviderTier, error) {
	return _TaikoL1Client.Contract.GetTier(&_TaikoL1Client.CallOpts, tierId)
}

// GetTier is a free data retrieval call binding the contract method 0x576c3de7.
//
// Solidity: function getTier(uint16 tierId) view returns((bytes32,uint96,uint96,uint24,uint16))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetTier(tierId uint16) (ITierProviderTier, error) {
	return _TaikoL1Client.Contract.GetTier(&_TaikoL1Client.CallOpts, tierId)
}

// GetTierIds is a free data retrieval call binding the contract method 0xd8cde1c6.
//
// Solidity: function getTierIds() view returns(uint16[])
func (_TaikoL1Client *TaikoL1ClientCaller) GetTierIds(opts *bind.CallOpts) ([]uint16, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getTierIds")

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// GetTierIds is a free data retrieval call binding the contract method 0xd8cde1c6.
//
// Solidity: function getTierIds() view returns(uint16[])
func (_TaikoL1Client *TaikoL1ClientSession) GetTierIds() ([]uint16, error) {
	return _TaikoL1Client.Contract.GetTierIds(&_TaikoL1Client.CallOpts)
}

// GetTierIds is a free data retrieval call binding the contract method 0xd8cde1c6.
//
// Solidity: function getTierIds() view returns(uint16[])
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetTierIds() ([]uint16, error) {
	return _TaikoL1Client.Contract.GetTierIds(&_TaikoL1Client.CallOpts)
}

// GetTransition is a free data retrieval call binding the contract method 0xfd257e29.
//
// Solidity: function getTransition(uint64 blockId, bytes32 parentHash) view returns((bytes32,bytes32,bytes32,address,uint96,address,uint96,uint64,uint16,bytes32[4]))
func (_TaikoL1Client *TaikoL1ClientCaller) GetTransition(opts *bind.CallOpts, blockId uint64, parentHash [32]byte) (TaikoDataTransition, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getTransition", blockId, parentHash)

	if err != nil {
		return *new(TaikoDataTransition), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoDataTransition)).(*TaikoDataTransition)

	return out0, err

}

// GetTransition is a free data retrieval call binding the contract method 0xfd257e29.
//
// Solidity: function getTransition(uint64 blockId, bytes32 parentHash) view returns((bytes32,bytes32,bytes32,address,uint96,address,uint96,uint64,uint16,bytes32[4]))
func (_TaikoL1Client *TaikoL1ClientSession) GetTransition(blockId uint64, parentHash [32]byte) (TaikoDataTransition, error) {
	return _TaikoL1Client.Contract.GetTransition(&_TaikoL1Client.CallOpts, blockId, parentHash)
}

// GetTransition is a free data retrieval call binding the contract method 0xfd257e29.
//
// Solidity: function getTransition(uint64 blockId, bytes32 parentHash) view returns((bytes32,bytes32,bytes32,address,uint96,address,uint96,uint64,uint16,bytes32[4]))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetTransition(blockId uint64, parentHash [32]byte) (TaikoDataTransition, error) {
	return _TaikoL1Client.Contract.GetTransition(&_TaikoL1Client.CallOpts, blockId, parentHash)
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
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
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
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL1Client *TaikoL1ClientSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve(&_TaikoL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve(&_TaikoL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
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
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL1Client *TaikoL1ClientSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve0(&_TaikoL1Client.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve0(&_TaikoL1Client.CallOpts, name, allowZeroAddress)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns((uint64,uint64,uint64,uint64) slotA, (uint64,uint64,uint64,uint64) slotB, (uint128,uint64) slotC)
func (_TaikoL1Client *TaikoL1ClientCaller) State(opts *bind.CallOpts) (struct {
	SlotA TaikoDataSlotA
	SlotB TaikoDataSlotB
	SlotC TaikoDataSlotC
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		SlotA TaikoDataSlotA
		SlotB TaikoDataSlotB
		SlotC TaikoDataSlotC
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SlotA = *abi.ConvertType(out[0], new(TaikoDataSlotA)).(*TaikoDataSlotA)
	outstruct.SlotB = *abi.ConvertType(out[1], new(TaikoDataSlotB)).(*TaikoDataSlotB)
	outstruct.SlotC = *abi.ConvertType(out[2], new(TaikoDataSlotC)).(*TaikoDataSlotC)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns((uint64,uint64,uint64,uint64) slotA, (uint64,uint64,uint64,uint64) slotB, (uint128,uint64) slotC)
func (_TaikoL1Client *TaikoL1ClientSession) State() (struct {
	SlotA TaikoDataSlotA
	SlotB TaikoDataSlotB
	SlotC TaikoDataSlotC
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns((uint64,uint64,uint64,uint64) slotA, (uint64,uint64,uint64,uint64) slotB, (uint128,uint64) slotC)
func (_TaikoL1Client *TaikoL1ClientCallerSession) State() (struct {
	SlotA TaikoDataSlotA
	SlotB TaikoDataSlotB
	SlotC TaikoDataSlotC
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0x047a289d.
//
// Solidity: function depositEtherToL2(address recipient) payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) DepositEtherToL2(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "depositEtherToL2", recipient)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0x047a289d.
//
// Solidity: function depositEtherToL2(address recipient) payable returns()
func (_TaikoL1Client *TaikoL1ClientSession) DepositEtherToL2(recipient common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.DepositEtherToL2(&_TaikoL1Client.TransactOpts, recipient)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0x047a289d.
//
// Solidity: function depositEtherToL2(address recipient) payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) DepositEtherToL2(recipient common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.DepositEtherToL2(&_TaikoL1Client.TransactOpts, recipient)
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

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _genesisBlockHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "init", _addressManager, _genesisBlockHash)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash) returns()
func (_TaikoL1Client *TaikoL1ClientSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Init(&_TaikoL1Client.TransactOpts, _addressManager, _genesisBlockHash)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Init(&_TaikoL1Client.TransactOpts, _addressManager, _genesisBlockHash)
}

// OwnerWithdrawTaikoToken is a paid mutator transaction binding the contract method 0x6edda419.
//
// Solidity: function ownerWithdrawTaikoToken(address to, uint256 amount) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) OwnerWithdrawTaikoToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "ownerWithdrawTaikoToken", to, amount)
}

// OwnerWithdrawTaikoToken is a paid mutator transaction binding the contract method 0x6edda419.
//
// Solidity: function ownerWithdrawTaikoToken(address to, uint256 amount) returns()
func (_TaikoL1Client *TaikoL1ClientSession) OwnerWithdrawTaikoToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.OwnerWithdrawTaikoToken(&_TaikoL1Client.TransactOpts, to, amount)
}

// OwnerWithdrawTaikoToken is a paid mutator transaction binding the contract method 0x6edda419.
//
// Solidity: function ownerWithdrawTaikoToken(address to, uint256 amount) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) OwnerWithdrawTaikoToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.OwnerWithdrawTaikoToken(&_TaikoL1Client.TransactOpts, to, amount)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0x04dc4c8b.
//
// Solidity: function proposeBlock(bytes32 txListHash, bytes32 extraData, bytes assignment, bytes txList) payable returns((bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint32,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientTransactor) ProposeBlock(opts *bind.TransactOpts, txListHash [32]byte, extraData [32]byte, assignment []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proposeBlock", txListHash, extraData, assignment, txList)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0x04dc4c8b.
//
// Solidity: function proposeBlock(bytes32 txListHash, bytes32 extraData, bytes assignment, bytes txList) payable returns((bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint32,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientSession) ProposeBlock(txListHash [32]byte, extraData [32]byte, assignment []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, txListHash, extraData, assignment, txList)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0x04dc4c8b.
//
// Solidity: function proposeBlock(bytes32 txListHash, bytes32 extraData, bytes assignment, bytes txList) payable returns((bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint32,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProposeBlock(txListHash [32]byte, extraData [32]byte, assignment []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, txListHash, extraData, assignment, txList)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x10d008bd.
//
// Solidity: function proveBlock(uint64 blockId, bytes input) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) ProveBlock(opts *bind.TransactOpts, blockId uint64, input []byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proveBlock", blockId, input)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x10d008bd.
//
// Solidity: function proveBlock(uint64 blockId, bytes input) returns()
func (_TaikoL1Client *TaikoL1ClientSession) ProveBlock(blockId uint64, input []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlock(&_TaikoL1Client.TransactOpts, blockId, input)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x10d008bd.
//
// Solidity: function proveBlock(uint64 blockId, bytes input) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProveBlock(blockId uint64, input []byte) (*types.Transaction, error) {
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

// VerifyBlocks is a paid mutator transaction binding the contract method 0x8778209d.
//
// Solidity: function verifyBlocks(uint64 maxBlocks) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) VerifyBlocks(opts *bind.TransactOpts, maxBlocks uint64) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "verifyBlocks", maxBlocks)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x8778209d.
//
// Solidity: function verifyBlocks(uint64 maxBlocks) returns()
func (_TaikoL1Client *TaikoL1ClientSession) VerifyBlocks(maxBlocks uint64) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.VerifyBlocks(&_TaikoL1Client.TransactOpts, maxBlocks)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x8778209d.
//
// Solidity: function verifyBlocks(uint64 maxBlocks) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) VerifyBlocks(maxBlocks uint64) (*types.Transaction, error) {
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
// Solidity: event AddressManagerChanged(address indexed addressManager)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterAddressManagerChanged(opts *bind.FilterOpts, addressManager []common.Address) (*TaikoL1ClientAddressManagerChangedIterator, error) {

	var addressManagerRule []interface{}
	for _, addressManagerItem := range addressManager {
		addressManagerRule = append(addressManagerRule, addressManagerItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "AddressManagerChanged", addressManagerRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientAddressManagerChangedIterator{contract: _TaikoL1Client.contract, event: "AddressManagerChanged", logs: logs, sub: sub}, nil
}

// WatchAddressManagerChanged is a free log subscription operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address indexed addressManager)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchAddressManagerChanged(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientAddressManagerChanged, addressManager []common.Address) (event.Subscription, error) {

	var addressManagerRule []interface{}
	for _, addressManagerItem := range addressManager {
		addressManagerRule = append(addressManagerRule, addressManagerItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "AddressManagerChanged", addressManagerRule)
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
// Solidity: event AddressManagerChanged(address indexed addressManager)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseAddressManagerChanged(log types.Log) (*TaikoL1ClientAddressManagerChanged, error) {
	event := new(TaikoL1ClientAddressManagerChanged)
	if err := _TaikoL1Client.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
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
	BlockId        *big.Int
	AssignedProver common.Address
	LivenessBond   *big.Int
	ProverFee      *big.Int
	Reward         *big.Int
	MinTier        uint16
	Meta           TaikoDataBlockMetadata
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0x3bc85eed257804b26e2232e767692931a620954ad0ae6ca5c58adb22be553418.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, uint256 proverFee, uint256 reward, uint16 minTier, (bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint32,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockProposed(opts *bind.FilterOpts, blockId []*big.Int, assignedProver []common.Address) (*TaikoL1ClientBlockProposedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockProposed", blockIdRule, assignedProverRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockProposedIterator{contract: _TaikoL1Client.contract, event: "BlockProposed", logs: logs, sub: sub}, nil
}

// WatchBlockProposed is a free log subscription operation binding the contract event 0x3bc85eed257804b26e2232e767692931a620954ad0ae6ca5c58adb22be553418.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, uint256 proverFee, uint256 reward, uint16 minTier, (bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint32,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockProposed(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockProposed, blockId []*big.Int, assignedProver []common.Address) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockProposed", blockIdRule, assignedProverRule)
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

// ParseBlockProposed is a log parse operation binding the contract event 0x3bc85eed257804b26e2232e767692931a620954ad0ae6ca5c58adb22be553418.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, uint256 proverFee, uint256 reward, uint16 minTier, (bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint32,address,(address,uint96,uint64)[]) meta)
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
	BlockId        *big.Int
	AssignedProver common.Address
	LivenessBond   *big.Int
	ProverFee      *big.Int
	Reward         *big.Int
	MinTier        uint16
	Meta           TaikoDataBlockMetadata
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed0 is a free log retrieval operation binding the contract event 0x3bc85eed257804b26e2232e767692931a620954ad0ae6ca5c58adb22be553418.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, uint256 proverFee, uint256 reward, uint16 minTier, (bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint32,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockProposed0(opts *bind.FilterOpts, blockId []*big.Int, assignedProver []common.Address) (*TaikoL1ClientBlockProposed0Iterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockProposed0", blockIdRule, assignedProverRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockProposed0Iterator{contract: _TaikoL1Client.contract, event: "BlockProposed0", logs: logs, sub: sub}, nil
}

// WatchBlockProposed0 is a free log subscription operation binding the contract event 0x3bc85eed257804b26e2232e767692931a620954ad0ae6ca5c58adb22be553418.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, uint256 proverFee, uint256 reward, uint16 minTier, (bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint32,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockProposed0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockProposed0, blockId []*big.Int, assignedProver []common.Address) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockProposed0", blockIdRule, assignedProverRule)
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

// ParseBlockProposed0 is a log parse operation binding the contract event 0x3bc85eed257804b26e2232e767692931a620954ad0ae6ca5c58adb22be553418.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, uint256 proverFee, uint256 reward, uint16 minTier, (bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint32,address,(address,uint96,uint64)[]) meta)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockProposed0(log types.Log) (*TaikoL1ClientBlockProposed0, error) {
	event := new(TaikoL1ClientBlockProposed0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProposed0", log); err != nil {
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
	BlockId        *big.Int
	AssignedProver common.Address
	Prover         common.Address
	BlockHash      [32]byte
	SignalRoot     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified is a free log retrieval operation binding the contract event 0x1433d0a0bbdef88a556faadc3c8481900a1bd0941743e979ddbe046e27ce4299.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockVerified(opts *bind.FilterOpts, blockId []*big.Int, assignedProver []common.Address, prover []common.Address) (*TaikoL1ClientBlockVerifiedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockVerified", blockIdRule, assignedProverRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockVerifiedIterator{contract: _TaikoL1Client.contract, event: "BlockVerified", logs: logs, sub: sub}, nil
}

// WatchBlockVerified is a free log subscription operation binding the contract event 0x1433d0a0bbdef88a556faadc3c8481900a1bd0941743e979ddbe046e27ce4299.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockVerified(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockVerified, blockId []*big.Int, assignedProver []common.Address, prover []common.Address) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockVerified", blockIdRule, assignedProverRule, proverRule)
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

// ParseBlockVerified is a log parse operation binding the contract event 0x1433d0a0bbdef88a556faadc3c8481900a1bd0941743e979ddbe046e27ce4299.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot)
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
	BlockId        *big.Int
	AssignedProver common.Address
	Prover         common.Address
	BlockHash      [32]byte
	SignalRoot     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified0 is a free log retrieval operation binding the contract event 0x1433d0a0bbdef88a556faadc3c8481900a1bd0941743e979ddbe046e27ce4299.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockVerified0(opts *bind.FilterOpts, blockId []*big.Int, assignedProver []common.Address, prover []common.Address) (*TaikoL1ClientBlockVerified0Iterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockVerified0", blockIdRule, assignedProverRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockVerified0Iterator{contract: _TaikoL1Client.contract, event: "BlockVerified0", logs: logs, sub: sub}, nil
}

// WatchBlockVerified0 is a free log subscription operation binding the contract event 0x1433d0a0bbdef88a556faadc3c8481900a1bd0941743e979ddbe046e27ce4299.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockVerified0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockVerified0, blockId []*big.Int, assignedProver []common.Address, prover []common.Address) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockVerified0", blockIdRule, assignedProverRule, proverRule)
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

// ParseBlockVerified0 is a log parse operation binding the contract event 0x1433d0a0bbdef88a556faadc3c8481900a1bd0941743e979ddbe046e27ce4299.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot)
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
	SrcHeight  uint64
	BlockHash  [32]byte
	SignalRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced is a free log retrieval operation binding the contract event 0x004ce985b8852a486571d0545799251fd671adcf33b7854a5f0f6a6a2431a555.
//
// Solidity: event CrossChainSynced(uint64 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterCrossChainSynced(opts *bind.FilterOpts, srcHeight []uint64) (*TaikoL1ClientCrossChainSyncedIterator, error) {

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

// WatchCrossChainSynced is a free log subscription operation binding the contract event 0x004ce985b8852a486571d0545799251fd671adcf33b7854a5f0f6a6a2431a555.
//
// Solidity: event CrossChainSynced(uint64 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchCrossChainSynced(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientCrossChainSynced, srcHeight []uint64) (event.Subscription, error) {

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

// ParseCrossChainSynced is a log parse operation binding the contract event 0x004ce985b8852a486571d0545799251fd671adcf33b7854a5f0f6a6a2431a555.
//
// Solidity: event CrossChainSynced(uint64 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
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
	SrcHeight  uint64
	BlockHash  [32]byte
	SignalRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced0 is a free log retrieval operation binding the contract event 0x004ce985b8852a486571d0545799251fd671adcf33b7854a5f0f6a6a2431a555.
//
// Solidity: event CrossChainSynced(uint64 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterCrossChainSynced0(opts *bind.FilterOpts, srcHeight []uint64) (*TaikoL1ClientCrossChainSynced0Iterator, error) {

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

// WatchCrossChainSynced0 is a free log subscription operation binding the contract event 0x004ce985b8852a486571d0545799251fd671adcf33b7854a5f0f6a6a2431a555.
//
// Solidity: event CrossChainSynced(uint64 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchCrossChainSynced0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientCrossChainSynced0, srcHeight []uint64) (event.Subscription, error) {

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

// ParseCrossChainSynced0 is a log parse operation binding the contract event 0x004ce985b8852a486571d0545799251fd671adcf33b7854a5f0f6a6a2431a555.
//
// Solidity: event CrossChainSynced(uint64 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
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

// TaikoL1ClientTokenCreditedIterator is returned from FilterTokenCredited and is used to iterate over the raw logs and unpacked data for TokenCredited events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenCreditedIterator struct {
	Event *TaikoL1ClientTokenCredited // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTokenCreditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenCredited)
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
		it.Event = new(TaikoL1ClientTokenCredited)
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
func (it *TaikoL1ClientTokenCreditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenCreditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenCredited represents a TokenCredited event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenCredited struct {
	Amount *big.Int
	Minted bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenCredited is a free log retrieval operation binding the contract event 0xeb0ac084e2fa6cfaebe54184190c7137c2fcc63c16e1dfaf6442ae95c01c44ad.
//
// Solidity: event TokenCredited(uint256 amount, bool minted)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenCredited(opts *bind.FilterOpts) (*TaikoL1ClientTokenCreditedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenCredited")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenCreditedIterator{contract: _TaikoL1Client.contract, event: "TokenCredited", logs: logs, sub: sub}, nil
}

// WatchTokenCredited is a free log subscription operation binding the contract event 0xeb0ac084e2fa6cfaebe54184190c7137c2fcc63c16e1dfaf6442ae95c01c44ad.
//
// Solidity: event TokenCredited(uint256 amount, bool minted)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenCredited(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenCredited) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenCredited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenCredited)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenCredited", log); err != nil {
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

// ParseTokenCredited is a log parse operation binding the contract event 0xeb0ac084e2fa6cfaebe54184190c7137c2fcc63c16e1dfaf6442ae95c01c44ad.
//
// Solidity: event TokenCredited(uint256 amount, bool minted)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenCredited(log types.Log) (*TaikoL1ClientTokenCredited, error) {
	event := new(TaikoL1ClientTokenCredited)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenCredited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenCredited0Iterator is returned from FilterTokenCredited0 and is used to iterate over the raw logs and unpacked data for TokenCredited0 events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenCredited0Iterator struct {
	Event *TaikoL1ClientTokenCredited0 // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTokenCredited0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenCredited0)
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
		it.Event = new(TaikoL1ClientTokenCredited0)
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
func (it *TaikoL1ClientTokenCredited0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenCredited0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenCredited0 represents a TokenCredited0 event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenCredited0 struct {
	Amount *big.Int
	Minted bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenCredited0 is a free log retrieval operation binding the contract event 0xeb0ac084e2fa6cfaebe54184190c7137c2fcc63c16e1dfaf6442ae95c01c44ad.
//
// Solidity: event TokenCredited(uint256 amount, bool minted)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenCredited0(opts *bind.FilterOpts) (*TaikoL1ClientTokenCredited0Iterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenCredited0")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenCredited0Iterator{contract: _TaikoL1Client.contract, event: "TokenCredited0", logs: logs, sub: sub}, nil
}

// WatchTokenCredited0 is a free log subscription operation binding the contract event 0xeb0ac084e2fa6cfaebe54184190c7137c2fcc63c16e1dfaf6442ae95c01c44ad.
//
// Solidity: event TokenCredited(uint256 amount, bool minted)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenCredited0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenCredited0) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenCredited0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenCredited0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenCredited0", log); err != nil {
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

// ParseTokenCredited0 is a log parse operation binding the contract event 0xeb0ac084e2fa6cfaebe54184190c7137c2fcc63c16e1dfaf6442ae95c01c44ad.
//
// Solidity: event TokenCredited(uint256 amount, bool minted)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenCredited0(log types.Log) (*TaikoL1ClientTokenCredited0, error) {
	event := new(TaikoL1ClientTokenCredited0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenCredited0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenDebitedIterator is returned from FilterTokenDebited and is used to iterate over the raw logs and unpacked data for TokenDebited events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDebitedIterator struct {
	Event *TaikoL1ClientTokenDebited // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTokenDebitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenDebited)
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
		it.Event = new(TaikoL1ClientTokenDebited)
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
func (it *TaikoL1ClientTokenDebitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenDebitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenDebited represents a TokenDebited event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDebited struct {
	Amount           *big.Int
	FromLocalBalance bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenDebited is a free log retrieval operation binding the contract event 0x8012eed335fcd2a9b93e688e5146be66d3b48c2e40578e713dc5f254d7fea32a.
//
// Solidity: event TokenDebited(uint256 amount, bool fromLocalBalance)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenDebited(opts *bind.FilterOpts) (*TaikoL1ClientTokenDebitedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenDebited")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenDebitedIterator{contract: _TaikoL1Client.contract, event: "TokenDebited", logs: logs, sub: sub}, nil
}

// WatchTokenDebited is a free log subscription operation binding the contract event 0x8012eed335fcd2a9b93e688e5146be66d3b48c2e40578e713dc5f254d7fea32a.
//
// Solidity: event TokenDebited(uint256 amount, bool fromLocalBalance)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenDebited(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenDebited) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenDebited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenDebited)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDebited", log); err != nil {
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

// ParseTokenDebited is a log parse operation binding the contract event 0x8012eed335fcd2a9b93e688e5146be66d3b48c2e40578e713dc5f254d7fea32a.
//
// Solidity: event TokenDebited(uint256 amount, bool fromLocalBalance)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenDebited(log types.Log) (*TaikoL1ClientTokenDebited, error) {
	event := new(TaikoL1ClientTokenDebited)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDebited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenDebited0Iterator is returned from FilterTokenDebited0 and is used to iterate over the raw logs and unpacked data for TokenDebited0 events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDebited0Iterator struct {
	Event *TaikoL1ClientTokenDebited0 // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTokenDebited0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenDebited0)
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
		it.Event = new(TaikoL1ClientTokenDebited0)
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
func (it *TaikoL1ClientTokenDebited0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenDebited0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenDebited0 represents a TokenDebited0 event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDebited0 struct {
	Amount           *big.Int
	FromLocalBalance bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenDebited0 is a free log retrieval operation binding the contract event 0x8012eed335fcd2a9b93e688e5146be66d3b48c2e40578e713dc5f254d7fea32a.
//
// Solidity: event TokenDebited(uint256 amount, bool fromLocalBalance)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenDebited0(opts *bind.FilterOpts) (*TaikoL1ClientTokenDebited0Iterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenDebited0")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenDebited0Iterator{contract: _TaikoL1Client.contract, event: "TokenDebited0", logs: logs, sub: sub}, nil
}

// WatchTokenDebited0 is a free log subscription operation binding the contract event 0x8012eed335fcd2a9b93e688e5146be66d3b48c2e40578e713dc5f254d7fea32a.
//
// Solidity: event TokenDebited(uint256 amount, bool fromLocalBalance)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenDebited0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenDebited0) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenDebited0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenDebited0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDebited0", log); err != nil {
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

// ParseTokenDebited0 is a log parse operation binding the contract event 0x8012eed335fcd2a9b93e688e5146be66d3b48c2e40578e713dc5f254d7fea32a.
//
// Solidity: event TokenDebited(uint256 amount, bool fromLocalBalance)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenDebited0(log types.Log) (*TaikoL1ClientTokenDebited0, error) {
	event := new(TaikoL1ClientTokenDebited0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDebited0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenDepositedIterator is returned from FilterTokenDeposited and is used to iterate over the raw logs and unpacked data for TokenDeposited events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDepositedIterator struct {
	Event *TaikoL1ClientTokenDeposited // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTokenDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenDeposited)
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
		it.Event = new(TaikoL1ClientTokenDeposited)
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
func (it *TaikoL1ClientTokenDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenDeposited represents a TokenDeposited event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDeposited struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposited is a free log retrieval operation binding the contract event 0x26a49ee784523ce049bcbe276a63c7c9dbd9f428b1aa53633e679c9c046e8858.
//
// Solidity: event TokenDeposited(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenDeposited(opts *bind.FilterOpts) (*TaikoL1ClientTokenDepositedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenDeposited")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenDepositedIterator{contract: _TaikoL1Client.contract, event: "TokenDeposited", logs: logs, sub: sub}, nil
}

// WatchTokenDeposited is a free log subscription operation binding the contract event 0x26a49ee784523ce049bcbe276a63c7c9dbd9f428b1aa53633e679c9c046e8858.
//
// Solidity: event TokenDeposited(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenDeposited(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenDeposited) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenDeposited)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
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

// ParseTokenDeposited is a log parse operation binding the contract event 0x26a49ee784523ce049bcbe276a63c7c9dbd9f428b1aa53633e679c9c046e8858.
//
// Solidity: event TokenDeposited(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenDeposited(log types.Log) (*TaikoL1ClientTokenDeposited, error) {
	event := new(TaikoL1ClientTokenDeposited)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenDeposited0Iterator is returned from FilterTokenDeposited0 and is used to iterate over the raw logs and unpacked data for TokenDeposited0 events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDeposited0Iterator struct {
	Event *TaikoL1ClientTokenDeposited0 // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTokenDeposited0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenDeposited0)
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
		it.Event = new(TaikoL1ClientTokenDeposited0)
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
func (it *TaikoL1ClientTokenDeposited0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenDeposited0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenDeposited0 represents a TokenDeposited0 event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDeposited0 struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposited0 is a free log retrieval operation binding the contract event 0x26a49ee784523ce049bcbe276a63c7c9dbd9f428b1aa53633e679c9c046e8858.
//
// Solidity: event TokenDeposited(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenDeposited0(opts *bind.FilterOpts) (*TaikoL1ClientTokenDeposited0Iterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenDeposited0")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenDeposited0Iterator{contract: _TaikoL1Client.contract, event: "TokenDeposited0", logs: logs, sub: sub}, nil
}

// WatchTokenDeposited0 is a free log subscription operation binding the contract event 0x26a49ee784523ce049bcbe276a63c7c9dbd9f428b1aa53633e679c9c046e8858.
//
// Solidity: event TokenDeposited(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenDeposited0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenDeposited0) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenDeposited0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenDeposited0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDeposited0", log); err != nil {
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

// ParseTokenDeposited0 is a log parse operation binding the contract event 0x26a49ee784523ce049bcbe276a63c7c9dbd9f428b1aa53633e679c9c046e8858.
//
// Solidity: event TokenDeposited(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenDeposited0(log types.Log) (*TaikoL1ClientTokenDeposited0, error) {
	event := new(TaikoL1ClientTokenDeposited0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDeposited0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenWithdrawnIterator struct {
	Event *TaikoL1ClientTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenWithdrawn)
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
		it.Event = new(TaikoL1ClientTokenWithdrawn)
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
func (it *TaikoL1ClientTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenWithdrawn represents a TokenWithdrawn event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenWithdrawn struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0xc172f6497c150fc242267f743e8e4034b31b16ee123408d6d5f75a81128de114.
//
// Solidity: event TokenWithdrawn(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts) (*TaikoL1ClientTokenWithdrawnIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenWithdrawn")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenWithdrawnIterator{contract: _TaikoL1Client.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0xc172f6497c150fc242267f743e8e4034b31b16ee123408d6d5f75a81128de114.
//
// Solidity: event TokenWithdrawn(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenWithdrawn) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenWithdrawn)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0xc172f6497c150fc242267f743e8e4034b31b16ee123408d6d5f75a81128de114.
//
// Solidity: event TokenWithdrawn(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenWithdrawn(log types.Log) (*TaikoL1ClientTokenWithdrawn, error) {
	event := new(TaikoL1ClientTokenWithdrawn)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenWithdrawn0Iterator is returned from FilterTokenWithdrawn0 and is used to iterate over the raw logs and unpacked data for TokenWithdrawn0 events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenWithdrawn0Iterator struct {
	Event *TaikoL1ClientTokenWithdrawn0 // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTokenWithdrawn0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenWithdrawn0)
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
		it.Event = new(TaikoL1ClientTokenWithdrawn0)
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
func (it *TaikoL1ClientTokenWithdrawn0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenWithdrawn0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenWithdrawn0 represents a TokenWithdrawn0 event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenWithdrawn0 struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn0 is a free log retrieval operation binding the contract event 0xc172f6497c150fc242267f743e8e4034b31b16ee123408d6d5f75a81128de114.
//
// Solidity: event TokenWithdrawn(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenWithdrawn0(opts *bind.FilterOpts) (*TaikoL1ClientTokenWithdrawn0Iterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenWithdrawn0")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenWithdrawn0Iterator{contract: _TaikoL1Client.contract, event: "TokenWithdrawn0", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn0 is a free log subscription operation binding the contract event 0xc172f6497c150fc242267f743e8e4034b31b16ee123408d6d5f75a81128de114.
//
// Solidity: event TokenWithdrawn(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenWithdrawn0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenWithdrawn0) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenWithdrawn0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenWithdrawn0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenWithdrawn0", log); err != nil {
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

// ParseTokenWithdrawn0 is a log parse operation binding the contract event 0xc172f6497c150fc242267f743e8e4034b31b16ee123408d6d5f75a81128de114.
//
// Solidity: event TokenWithdrawn(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenWithdrawn0(log types.Log) (*TaikoL1ClientTokenWithdrawn0, error) {
	event := new(TaikoL1ClientTokenWithdrawn0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenWithdrawn0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenWithdrawnByOwnerIterator is returned from FilterTokenWithdrawnByOwner and is used to iterate over the raw logs and unpacked data for TokenWithdrawnByOwner events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenWithdrawnByOwnerIterator struct {
	Event *TaikoL1ClientTokenWithdrawnByOwner // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTokenWithdrawnByOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenWithdrawnByOwner)
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
		it.Event = new(TaikoL1ClientTokenWithdrawnByOwner)
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
func (it *TaikoL1ClientTokenWithdrawnByOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenWithdrawnByOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenWithdrawnByOwner represents a TokenWithdrawnByOwner event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenWithdrawnByOwner struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawnByOwner is a free log retrieval operation binding the contract event 0xc4f72d4b841abf7e4afe6a797fbc9d92a5186aa6332c2ff33c0e327f1c3e9a4c.
//
// Solidity: event TokenWithdrawnByOwner(address to, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenWithdrawnByOwner(opts *bind.FilterOpts) (*TaikoL1ClientTokenWithdrawnByOwnerIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenWithdrawnByOwner")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenWithdrawnByOwnerIterator{contract: _TaikoL1Client.contract, event: "TokenWithdrawnByOwner", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawnByOwner is a free log subscription operation binding the contract event 0xc4f72d4b841abf7e4afe6a797fbc9d92a5186aa6332c2ff33c0e327f1c3e9a4c.
//
// Solidity: event TokenWithdrawnByOwner(address to, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenWithdrawnByOwner(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenWithdrawnByOwner) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenWithdrawnByOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenWithdrawnByOwner)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenWithdrawnByOwner", log); err != nil {
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

// ParseTokenWithdrawnByOwner is a log parse operation binding the contract event 0xc4f72d4b841abf7e4afe6a797fbc9d92a5186aa6332c2ff33c0e327f1c3e9a4c.
//
// Solidity: event TokenWithdrawnByOwner(address to, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenWithdrawnByOwner(log types.Log) (*TaikoL1ClientTokenWithdrawnByOwner, error) {
	event := new(TaikoL1ClientTokenWithdrawnByOwner)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenWithdrawnByOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenWithdrawnByOwner0Iterator is returned from FilterTokenWithdrawnByOwner0 and is used to iterate over the raw logs and unpacked data for TokenWithdrawnByOwner0 events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenWithdrawnByOwner0Iterator struct {
	Event *TaikoL1ClientTokenWithdrawnByOwner0 // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTokenWithdrawnByOwner0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenWithdrawnByOwner0)
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
		it.Event = new(TaikoL1ClientTokenWithdrawnByOwner0)
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
func (it *TaikoL1ClientTokenWithdrawnByOwner0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenWithdrawnByOwner0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenWithdrawnByOwner0 represents a TokenWithdrawnByOwner0 event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenWithdrawnByOwner0 struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawnByOwner0 is a free log retrieval operation binding the contract event 0xc4f72d4b841abf7e4afe6a797fbc9d92a5186aa6332c2ff33c0e327f1c3e9a4c.
//
// Solidity: event TokenWithdrawnByOwner(address to, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenWithdrawnByOwner0(opts *bind.FilterOpts) (*TaikoL1ClientTokenWithdrawnByOwner0Iterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenWithdrawnByOwner0")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenWithdrawnByOwner0Iterator{contract: _TaikoL1Client.contract, event: "TokenWithdrawnByOwner0", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawnByOwner0 is a free log subscription operation binding the contract event 0xc4f72d4b841abf7e4afe6a797fbc9d92a5186aa6332c2ff33c0e327f1c3e9a4c.
//
// Solidity: event TokenWithdrawnByOwner(address to, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenWithdrawnByOwner0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenWithdrawnByOwner0) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenWithdrawnByOwner0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenWithdrawnByOwner0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenWithdrawnByOwner0", log); err != nil {
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

// ParseTokenWithdrawnByOwner0 is a log parse operation binding the contract event 0xc4f72d4b841abf7e4afe6a797fbc9d92a5186aa6332c2ff33c0e327f1c3e9a4c.
//
// Solidity: event TokenWithdrawnByOwner(address to, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenWithdrawnByOwner0(log types.Log) (*TaikoL1ClientTokenWithdrawnByOwner0, error) {
	event := new(TaikoL1ClientTokenWithdrawnByOwner0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenWithdrawnByOwner0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTransitionContestedIterator is returned from FilterTransitionContested and is used to iterate over the raw logs and unpacked data for TransitionContested events raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionContestedIterator struct {
	Event *TaikoL1ClientTransitionContested // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTransitionContestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTransitionContested)
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
		it.Event = new(TaikoL1ClientTransitionContested)
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
func (it *TaikoL1ClientTransitionContestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTransitionContestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTransitionContested represents a TransitionContested event raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionContested struct {
	BlockId     *big.Int
	ParentHash  [32]byte
	BlockHash   [32]byte
	SignalRoot  [32]byte
	Contester   common.Address
	ContestBond *big.Int
	Tier        uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransitionContested is a free log retrieval operation binding the contract event 0x135961061f7c7cad9a389080483d14014ba6fe9ebcf6d3d826185c69dd71ce75.
//
// Solidity: event TransitionContested(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address contester, uint96 contestBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTransitionContested(opts *bind.FilterOpts, blockId []*big.Int) (*TaikoL1ClientTransitionContestedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TransitionContested", blockIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTransitionContestedIterator{contract: _TaikoL1Client.contract, event: "TransitionContested", logs: logs, sub: sub}, nil
}

// WatchTransitionContested is a free log subscription operation binding the contract event 0x135961061f7c7cad9a389080483d14014ba6fe9ebcf6d3d826185c69dd71ce75.
//
// Solidity: event TransitionContested(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address contester, uint96 contestBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTransitionContested(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTransitionContested, blockId []*big.Int) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TransitionContested", blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTransitionContested)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionContested", log); err != nil {
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

// ParseTransitionContested is a log parse operation binding the contract event 0x135961061f7c7cad9a389080483d14014ba6fe9ebcf6d3d826185c69dd71ce75.
//
// Solidity: event TransitionContested(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address contester, uint96 contestBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTransitionContested(log types.Log) (*TaikoL1ClientTransitionContested, error) {
	event := new(TaikoL1ClientTransitionContested)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionContested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTransitionContested0Iterator is returned from FilterTransitionContested0 and is used to iterate over the raw logs and unpacked data for TransitionContested0 events raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionContested0Iterator struct {
	Event *TaikoL1ClientTransitionContested0 // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTransitionContested0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTransitionContested0)
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
		it.Event = new(TaikoL1ClientTransitionContested0)
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
func (it *TaikoL1ClientTransitionContested0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTransitionContested0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTransitionContested0 represents a TransitionContested0 event raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionContested0 struct {
	BlockId     *big.Int
	ParentHash  [32]byte
	BlockHash   [32]byte
	SignalRoot  [32]byte
	Contester   common.Address
	ContestBond *big.Int
	Tier        uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransitionContested0 is a free log retrieval operation binding the contract event 0x135961061f7c7cad9a389080483d14014ba6fe9ebcf6d3d826185c69dd71ce75.
//
// Solidity: event TransitionContested(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address contester, uint96 contestBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTransitionContested0(opts *bind.FilterOpts, blockId []*big.Int) (*TaikoL1ClientTransitionContested0Iterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TransitionContested0", blockIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTransitionContested0Iterator{contract: _TaikoL1Client.contract, event: "TransitionContested0", logs: logs, sub: sub}, nil
}

// WatchTransitionContested0 is a free log subscription operation binding the contract event 0x135961061f7c7cad9a389080483d14014ba6fe9ebcf6d3d826185c69dd71ce75.
//
// Solidity: event TransitionContested(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address contester, uint96 contestBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTransitionContested0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTransitionContested0, blockId []*big.Int) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TransitionContested0", blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTransitionContested0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionContested0", log); err != nil {
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

// ParseTransitionContested0 is a log parse operation binding the contract event 0x135961061f7c7cad9a389080483d14014ba6fe9ebcf6d3d826185c69dd71ce75.
//
// Solidity: event TransitionContested(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address contester, uint96 contestBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTransitionContested0(log types.Log) (*TaikoL1ClientTransitionContested0, error) {
	event := new(TaikoL1ClientTransitionContested0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionContested0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTransitionProvedIterator is returned from FilterTransitionProved and is used to iterate over the raw logs and unpacked data for TransitionProved events raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionProvedIterator struct {
	Event *TaikoL1ClientTransitionProved // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTransitionProvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTransitionProved)
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
		it.Event = new(TaikoL1ClientTransitionProved)
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
func (it *TaikoL1ClientTransitionProvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTransitionProvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTransitionProved represents a TransitionProved event raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionProved struct {
	BlockId      *big.Int
	ParentHash   [32]byte
	BlockHash    [32]byte
	SignalRoot   [32]byte
	Prover       common.Address
	ValidityBond *big.Int
	Tier         uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransitionProved is a free log retrieval operation binding the contract event 0x9eaac66787b7d7e72981c31449ddd9bde78ffdb8cb3e6543ba2b9a4bb653b5e3.
//
// Solidity: event TransitionProved(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint96 validityBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTransitionProved(opts *bind.FilterOpts, blockId []*big.Int) (*TaikoL1ClientTransitionProvedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TransitionProved", blockIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTransitionProvedIterator{contract: _TaikoL1Client.contract, event: "TransitionProved", logs: logs, sub: sub}, nil
}

// WatchTransitionProved is a free log subscription operation binding the contract event 0x9eaac66787b7d7e72981c31449ddd9bde78ffdb8cb3e6543ba2b9a4bb653b5e3.
//
// Solidity: event TransitionProved(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint96 validityBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTransitionProved(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTransitionProved, blockId []*big.Int) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TransitionProved", blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTransitionProved)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionProved", log); err != nil {
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

// ParseTransitionProved is a log parse operation binding the contract event 0x9eaac66787b7d7e72981c31449ddd9bde78ffdb8cb3e6543ba2b9a4bb653b5e3.
//
// Solidity: event TransitionProved(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint96 validityBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTransitionProved(log types.Log) (*TaikoL1ClientTransitionProved, error) {
	event := new(TaikoL1ClientTransitionProved)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionProved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTransitionProved0Iterator is returned from FilterTransitionProved0 and is used to iterate over the raw logs and unpacked data for TransitionProved0 events raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionProved0Iterator struct {
	Event *TaikoL1ClientTransitionProved0 // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientTransitionProved0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTransitionProved0)
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
		it.Event = new(TaikoL1ClientTransitionProved0)
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
func (it *TaikoL1ClientTransitionProved0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTransitionProved0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTransitionProved0 represents a TransitionProved0 event raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionProved0 struct {
	BlockId      *big.Int
	ParentHash   [32]byte
	BlockHash    [32]byte
	SignalRoot   [32]byte
	Prover       common.Address
	ValidityBond *big.Int
	Tier         uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransitionProved0 is a free log retrieval operation binding the contract event 0x9eaac66787b7d7e72981c31449ddd9bde78ffdb8cb3e6543ba2b9a4bb653b5e3.
//
// Solidity: event TransitionProved(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint96 validityBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTransitionProved0(opts *bind.FilterOpts, blockId []*big.Int) (*TaikoL1ClientTransitionProved0Iterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TransitionProved0", blockIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTransitionProved0Iterator{contract: _TaikoL1Client.contract, event: "TransitionProved0", logs: logs, sub: sub}, nil
}

// WatchTransitionProved0 is a free log subscription operation binding the contract event 0x9eaac66787b7d7e72981c31449ddd9bde78ffdb8cb3e6543ba2b9a4bb653b5e3.
//
// Solidity: event TransitionProved(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint96 validityBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTransitionProved0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTransitionProved0, blockId []*big.Int) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TransitionProved0", blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTransitionProved0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionProved0", log); err != nil {
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

// ParseTransitionProved0 is a log parse operation binding the contract event 0x9eaac66787b7d7e72981c31449ddd9bde78ffdb8cb3e6543ba2b9a4bb653b5e3.
//
// Solidity: event TransitionProved(uint256 indexed blockId, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint96 validityBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTransitionProved0(log types.Log) (*TaikoL1ClientTransitionProved0, error) {
	event := new(TaikoL1ClientTransitionProved0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionProved0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
