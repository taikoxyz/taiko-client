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

// TaikoL2Config is an auto generated low-level Go binding around an user-defined struct.
type TaikoL2Config struct {
	GasTargetPerL1Block       uint64
	BasefeeAdjustmentQuotient *big.Int
	BlockRewardPerL1Block     *big.Int
	BlockRewardPoolMax        *big.Int
	BlockRewardPoolPctg       uint8
}

// TaikoL2ClientMetaData contains all meta data concerning the TaikoL2Client contract.
var TaikoL2ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EIP1559_INVALID_PARAMS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_BASEFEE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_GAS_EXCESS_TOO_LARGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_CHAIN_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_GOLDEN_TOUCH_K\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_SENDER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_PUBLIC_INPUT_HASH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_TOO_LATE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_MANAGER\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"gasExcess\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"blockReward\",\"type\":\"uint128\"}],\"name\":\"Anchored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"srcHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ANCHOR_GAS_DEDUCT\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOLDEN_TOUCH_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOLDEN_TOUCH_PRIVATEKEY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedReward\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l1BlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l1SignalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"anchor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avgGasUsed\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasExcess\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"getBasefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"basefee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockId\",\"type\":\"uint64\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gasTargetPerL1Block\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"basefeeAdjustmentQuotient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockRewardPerL1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"blockRewardPoolMax\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"blockRewardPoolPctg\",\"type\":\"uint8\"}],\"internalType\":\"structTaikoL2.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockId\",\"type\":\"uint64\"}],\"name\":\"getSyncedSnippet\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structICrossChainSync.Snippet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_gasExcess\",\"type\":\"uint128\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"l2Hashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestSyncedL1Height\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"k\",\"type\":\"uint8\"}],\"name\":\"signAnchor\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l1height\",\"type\":\"uint256\"}],\"name\":\"snippets\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TaikoL2ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use TaikoL2ClientMetaData.ABI instead.
var TaikoL2ClientABI = TaikoL2ClientMetaData.ABI

// TaikoL2Client is an auto generated Go binding around an Ethereum contract.
type TaikoL2Client struct {
	TaikoL2ClientCaller     // Read-only binding to the contract
	TaikoL2ClientTransactor // Write-only binding to the contract
	TaikoL2ClientFilterer   // Log filterer for contract events
}

// TaikoL2ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaikoL2ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL2ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaikoL2ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL2ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaikoL2ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL2ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaikoL2ClientSession struct {
	Contract     *TaikoL2Client    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaikoL2ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaikoL2ClientCallerSession struct {
	Contract *TaikoL2ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TaikoL2ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaikoL2ClientTransactorSession struct {
	Contract     *TaikoL2ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TaikoL2ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaikoL2ClientRaw struct {
	Contract *TaikoL2Client // Generic contract binding to access the raw methods on
}

// TaikoL2ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaikoL2ClientCallerRaw struct {
	Contract *TaikoL2ClientCaller // Generic read-only contract binding to access the raw methods on
}

// TaikoL2ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaikoL2ClientTransactorRaw struct {
	Contract *TaikoL2ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaikoL2Client creates a new instance of TaikoL2Client, bound to a specific deployed contract.
func NewTaikoL2Client(address common.Address, backend bind.ContractBackend) (*TaikoL2Client, error) {
	contract, err := bindTaikoL2Client(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaikoL2Client{TaikoL2ClientCaller: TaikoL2ClientCaller{contract: contract}, TaikoL2ClientTransactor: TaikoL2ClientTransactor{contract: contract}, TaikoL2ClientFilterer: TaikoL2ClientFilterer{contract: contract}}, nil
}

// NewTaikoL2ClientCaller creates a new read-only instance of TaikoL2Client, bound to a specific deployed contract.
func NewTaikoL2ClientCaller(address common.Address, caller bind.ContractCaller) (*TaikoL2ClientCaller, error) {
	contract, err := bindTaikoL2Client(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientCaller{contract: contract}, nil
}

// NewTaikoL2ClientTransactor creates a new write-only instance of TaikoL2Client, bound to a specific deployed contract.
func NewTaikoL2ClientTransactor(address common.Address, transactor bind.ContractTransactor) (*TaikoL2ClientTransactor, error) {
	contract, err := bindTaikoL2Client(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientTransactor{contract: contract}, nil
}

// NewTaikoL2ClientFilterer creates a new log filterer instance of TaikoL2Client, bound to a specific deployed contract.
func NewTaikoL2ClientFilterer(address common.Address, filterer bind.ContractFilterer) (*TaikoL2ClientFilterer, error) {
	contract, err := bindTaikoL2Client(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientFilterer{contract: contract}, nil
}

// bindTaikoL2Client binds a generic wrapper to an already deployed contract.
func bindTaikoL2Client(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaikoL2ClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoL2Client *TaikoL2ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoL2Client.Contract.TaikoL2ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoL2Client *TaikoL2ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.TaikoL2ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoL2Client *TaikoL2ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.TaikoL2ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoL2Client *TaikoL2ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoL2Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoL2Client *TaikoL2ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoL2Client *TaikoL2ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.contract.Transact(opts, method, params...)
}

// ANCHORGASDEDUCT is a free data retrieval call binding the contract method 0x51bf57b4.
//
// Solidity: function ANCHOR_GAS_DEDUCT() view returns(uint32)
func (_TaikoL2Client *TaikoL2ClientCaller) ANCHORGASDEDUCT(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "ANCHOR_GAS_DEDUCT")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ANCHORGASDEDUCT is a free data retrieval call binding the contract method 0x51bf57b4.
//
// Solidity: function ANCHOR_GAS_DEDUCT() view returns(uint32)
func (_TaikoL2Client *TaikoL2ClientSession) ANCHORGASDEDUCT() (uint32, error) {
	return _TaikoL2Client.Contract.ANCHORGASDEDUCT(&_TaikoL2Client.CallOpts)
}

// ANCHORGASDEDUCT is a free data retrieval call binding the contract method 0x51bf57b4.
//
// Solidity: function ANCHOR_GAS_DEDUCT() view returns(uint32)
func (_TaikoL2Client *TaikoL2ClientCallerSession) ANCHORGASDEDUCT() (uint32, error) {
	return _TaikoL2Client.Contract.ANCHORGASDEDUCT(&_TaikoL2Client.CallOpts)
}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCaller) GOLDENTOUCHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "GOLDEN_TOUCH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_TaikoL2Client *TaikoL2ClientSession) GOLDENTOUCHADDRESS() (common.Address, error) {
	return _TaikoL2Client.Contract.GOLDENTOUCHADDRESS(&_TaikoL2Client.CallOpts)
}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GOLDENTOUCHADDRESS() (common.Address, error) {
	return _TaikoL2Client.Contract.GOLDENTOUCHADDRESS(&_TaikoL2Client.CallOpts)
}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_TaikoL2Client *TaikoL2ClientCaller) GOLDENTOUCHPRIVATEKEY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "GOLDEN_TOUCH_PRIVATEKEY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_TaikoL2Client *TaikoL2ClientSession) GOLDENTOUCHPRIVATEKEY() (*big.Int, error) {
	return _TaikoL2Client.Contract.GOLDENTOUCHPRIVATEKEY(&_TaikoL2Client.CallOpts)
}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GOLDENTOUCHPRIVATEKEY() (*big.Int, error) {
	return _TaikoL2Client.Contract.GOLDENTOUCHPRIVATEKEY(&_TaikoL2Client.CallOpts)
}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint128)
func (_TaikoL2Client *TaikoL2ClientCaller) AccumulatedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "accumulatedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint128)
func (_TaikoL2Client *TaikoL2ClientSession) AccumulatedReward() (*big.Int, error) {
	return _TaikoL2Client.Contract.AccumulatedReward(&_TaikoL2Client.CallOpts)
}

// AccumulatedReward is a free data retrieval call binding the contract method 0x80fad325.
//
// Solidity: function accumulatedReward() view returns(uint128)
func (_TaikoL2Client *TaikoL2ClientCallerSession) AccumulatedReward() (*big.Int, error) {
	return _TaikoL2Client.Contract.AccumulatedReward(&_TaikoL2Client.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL2Client *TaikoL2ClientSession) AddressManager() (common.Address, error) {
	return _TaikoL2Client.Contract.AddressManager(&_TaikoL2Client.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCallerSession) AddressManager() (common.Address, error) {
	return _TaikoL2Client.Contract.AddressManager(&_TaikoL2Client.CallOpts)
}

// AvgGasUsed is a free data retrieval call binding the contract method 0x6d6372c4.
//
// Solidity: function avgGasUsed() view returns(uint32)
func (_TaikoL2Client *TaikoL2ClientCaller) AvgGasUsed(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "avgGasUsed")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AvgGasUsed is a free data retrieval call binding the contract method 0x6d6372c4.
//
// Solidity: function avgGasUsed() view returns(uint32)
func (_TaikoL2Client *TaikoL2ClientSession) AvgGasUsed() (uint32, error) {
	return _TaikoL2Client.Contract.AvgGasUsed(&_TaikoL2Client.CallOpts)
}

// AvgGasUsed is a free data retrieval call binding the contract method 0x6d6372c4.
//
// Solidity: function avgGasUsed() view returns(uint32)
func (_TaikoL2Client *TaikoL2ClientCallerSession) AvgGasUsed() (uint32, error) {
	return _TaikoL2Client.Contract.AvgGasUsed(&_TaikoL2Client.CallOpts)
}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint128)
func (_TaikoL2Client *TaikoL2ClientCaller) GasExcess(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "gasExcess")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint128)
func (_TaikoL2Client *TaikoL2ClientSession) GasExcess() (*big.Int, error) {
	return _TaikoL2Client.Contract.GasExcess(&_TaikoL2Client.CallOpts)
}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint128)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GasExcess() (*big.Int, error) {
	return _TaikoL2Client.Contract.GasExcess(&_TaikoL2Client.CallOpts)
}

// GetBasefee is a free data retrieval call binding the contract method 0xa7e022d1.
//
// Solidity: function getBasefee(uint64 l1Height, uint32 parentGasUsed) view returns(uint256 basefee)
func (_TaikoL2Client *TaikoL2ClientCaller) GetBasefee(opts *bind.CallOpts, l1Height uint64, parentGasUsed uint32) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "getBasefee", l1Height, parentGasUsed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBasefee is a free data retrieval call binding the contract method 0xa7e022d1.
//
// Solidity: function getBasefee(uint64 l1Height, uint32 parentGasUsed) view returns(uint256 basefee)
func (_TaikoL2Client *TaikoL2ClientSession) GetBasefee(l1Height uint64, parentGasUsed uint32) (*big.Int, error) {
	return _TaikoL2Client.Contract.GetBasefee(&_TaikoL2Client.CallOpts, l1Height, parentGasUsed)
}

// GetBasefee is a free data retrieval call binding the contract method 0xa7e022d1.
//
// Solidity: function getBasefee(uint64 l1Height, uint32 parentGasUsed) view returns(uint256 basefee)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GetBasefee(l1Height uint64, parentGasUsed uint32) (*big.Int, error) {
	return _TaikoL2Client.Contract.GetBasefee(&_TaikoL2Client.CallOpts, l1Height, parentGasUsed)
}

// GetBlockHash is a free data retrieval call binding the contract method 0x23ac7136.
//
// Solidity: function getBlockHash(uint64 blockId) view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientCaller) GetBlockHash(opts *bind.CallOpts, blockId uint64) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "getBlockHash", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0x23ac7136.
//
// Solidity: function getBlockHash(uint64 blockId) view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientSession) GetBlockHash(blockId uint64) ([32]byte, error) {
	return _TaikoL2Client.Contract.GetBlockHash(&_TaikoL2Client.CallOpts, blockId)
}

// GetBlockHash is a free data retrieval call binding the contract method 0x23ac7136.
//
// Solidity: function getBlockHash(uint64 blockId) view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GetBlockHash(blockId uint64) ([32]byte, error) {
	return _TaikoL2Client.Contract.GetBlockHash(&_TaikoL2Client.CallOpts, blockId)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint64,uint256,uint256,uint128,uint8) config)
func (_TaikoL2Client *TaikoL2ClientCaller) GetConfig(opts *bind.CallOpts) (TaikoL2Config, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(TaikoL2Config), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoL2Config)).(*TaikoL2Config)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint64,uint256,uint256,uint128,uint8) config)
func (_TaikoL2Client *TaikoL2ClientSession) GetConfig() (TaikoL2Config, error) {
	return _TaikoL2Client.Contract.GetConfig(&_TaikoL2Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint64,uint256,uint256,uint128,uint8) config)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GetConfig() (TaikoL2Config, error) {
	return _TaikoL2Client.Contract.GetConfig(&_TaikoL2Client.CallOpts)
}

// GetSyncedSnippet is a free data retrieval call binding the contract method 0x8cfb0459.
//
// Solidity: function getSyncedSnippet(uint64 blockId) view returns((bytes32,bytes32))
func (_TaikoL2Client *TaikoL2ClientCaller) GetSyncedSnippet(opts *bind.CallOpts, blockId uint64) (ICrossChainSyncSnippet, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "getSyncedSnippet", blockId)

	if err != nil {
		return *new(ICrossChainSyncSnippet), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossChainSyncSnippet)).(*ICrossChainSyncSnippet)

	return out0, err

}

// GetSyncedSnippet is a free data retrieval call binding the contract method 0x8cfb0459.
//
// Solidity: function getSyncedSnippet(uint64 blockId) view returns((bytes32,bytes32))
func (_TaikoL2Client *TaikoL2ClientSession) GetSyncedSnippet(blockId uint64) (ICrossChainSyncSnippet, error) {
	return _TaikoL2Client.Contract.GetSyncedSnippet(&_TaikoL2Client.CallOpts, blockId)
}

// GetSyncedSnippet is a free data retrieval call binding the contract method 0x8cfb0459.
//
// Solidity: function getSyncedSnippet(uint64 blockId) view returns((bytes32,bytes32))
func (_TaikoL2Client *TaikoL2ClientCallerSession) GetSyncedSnippet(blockId uint64) (ICrossChainSyncSnippet, error) {
	return _TaikoL2Client.Contract.GetSyncedSnippet(&_TaikoL2Client.CallOpts, blockId)
}

// L2Hashes is a free data retrieval call binding the contract method 0x8551f41e.
//
// Solidity: function l2Hashes(uint256 blockId) view returns(bytes32 blockHash)
func (_TaikoL2Client *TaikoL2ClientCaller) L2Hashes(opts *bind.CallOpts, blockId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "l2Hashes", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2Hashes is a free data retrieval call binding the contract method 0x8551f41e.
//
// Solidity: function l2Hashes(uint256 blockId) view returns(bytes32 blockHash)
func (_TaikoL2Client *TaikoL2ClientSession) L2Hashes(blockId *big.Int) ([32]byte, error) {
	return _TaikoL2Client.Contract.L2Hashes(&_TaikoL2Client.CallOpts, blockId)
}

// L2Hashes is a free data retrieval call binding the contract method 0x8551f41e.
//
// Solidity: function l2Hashes(uint256 blockId) view returns(bytes32 blockHash)
func (_TaikoL2Client *TaikoL2ClientCallerSession) L2Hashes(blockId *big.Int) ([32]byte, error) {
	return _TaikoL2Client.Contract.L2Hashes(&_TaikoL2Client.CallOpts, blockId)
}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientCaller) LatestSyncedL1Height(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "latestSyncedL1Height")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientSession) LatestSyncedL1Height() (uint64, error) {
	return _TaikoL2Client.Contract.LatestSyncedL1Height(&_TaikoL2Client.CallOpts)
}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientCallerSession) LatestSyncedL1Height() (uint64, error) {
	return _TaikoL2Client.Contract.LatestSyncedL1Height(&_TaikoL2Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL2Client *TaikoL2ClientSession) Owner() (common.Address, error) {
	return _TaikoL2Client.Contract.Owner(&_TaikoL2Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCallerSession) Owner() (common.Address, error) {
	return _TaikoL2Client.Contract.Owner(&_TaikoL2Client.CallOpts)
}

// ParentProposer is a free data retrieval call binding the contract method 0xfca19902.
//
// Solidity: function parentProposer() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCaller) ParentProposer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "parentProposer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParentProposer is a free data retrieval call binding the contract method 0xfca19902.
//
// Solidity: function parentProposer() view returns(address)
func (_TaikoL2Client *TaikoL2ClientSession) ParentProposer() (common.Address, error) {
	return _TaikoL2Client.Contract.ParentProposer(&_TaikoL2Client.CallOpts)
}

// ParentProposer is a free data retrieval call binding the contract method 0xfca19902.
//
// Solidity: function parentProposer() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCallerSession) ParentProposer() (common.Address, error) {
	return _TaikoL2Client.Contract.ParentProposer(&_TaikoL2Client.CallOpts)
}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientCaller) PublicInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "publicInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientSession) PublicInputHash() ([32]byte, error) {
	return _TaikoL2Client.Contract.PublicInputHash(&_TaikoL2Client.CallOpts)
}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientCallerSession) PublicInputHash() ([32]byte, error) {
	return _TaikoL2Client.Contract.PublicInputHash(&_TaikoL2Client.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientCaller) Resolve(opts *bind.CallOpts, chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL2Client.Contract.Resolve(&_TaikoL2Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientCallerSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL2Client.Contract.Resolve(&_TaikoL2Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientCaller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL2Client.Contract.Resolve0(&_TaikoL2Client.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientCallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL2Client.Contract.Resolve0(&_TaikoL2Client.CallOpts, name, allowZeroAddress)
}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL2Client *TaikoL2ClientCaller) SignAnchor(opts *bind.CallOpts, digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "signAnchor", digest, k)

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

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL2Client *TaikoL2ClientSession) SignAnchor(digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _TaikoL2Client.Contract.SignAnchor(&_TaikoL2Client.CallOpts, digest, k)
}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL2Client *TaikoL2ClientCallerSession) SignAnchor(digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _TaikoL2Client.Contract.SignAnchor(&_TaikoL2Client.CallOpts, digest, k)
}

// Snippets is a free data retrieval call binding the contract method 0xe8e2c5fb.
//
// Solidity: function snippets(uint256 l1height) view returns(bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientCaller) Snippets(opts *bind.CallOpts, l1height *big.Int) (struct {
	BlockHash  [32]byte
	SignalRoot [32]byte
}, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "snippets", l1height)

	outstruct := new(struct {
		BlockHash  [32]byte
		SignalRoot [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.SignalRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Snippets is a free data retrieval call binding the contract method 0xe8e2c5fb.
//
// Solidity: function snippets(uint256 l1height) view returns(bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientSession) Snippets(l1height *big.Int) (struct {
	BlockHash  [32]byte
	SignalRoot [32]byte
}, error) {
	return _TaikoL2Client.Contract.Snippets(&_TaikoL2Client.CallOpts, l1height)
}

// Snippets is a free data retrieval call binding the contract method 0xe8e2c5fb.
//
// Solidity: function snippets(uint256 l1height) view returns(bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientCallerSession) Snippets(l1height *big.Int) (struct {
	BlockHash  [32]byte
	SignalRoot [32]byte
}, error) {
	return _TaikoL2Client.Contract.Snippets(&_TaikoL2Client.CallOpts, l1height)
}

// Anchor is a paid mutator transaction binding the contract method 0xda69d3db.
//
// Solidity: function anchor(bytes32 l1BlockHash, bytes32 l1SignalRoot, uint64 l1Height, uint32 parentGasUsed) returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) Anchor(opts *bind.TransactOpts, l1BlockHash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint32) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "anchor", l1BlockHash, l1SignalRoot, l1Height, parentGasUsed)
}

// Anchor is a paid mutator transaction binding the contract method 0xda69d3db.
//
// Solidity: function anchor(bytes32 l1BlockHash, bytes32 l1SignalRoot, uint64 l1Height, uint32 parentGasUsed) returns()
func (_TaikoL2Client *TaikoL2ClientSession) Anchor(l1BlockHash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint32) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Anchor(&_TaikoL2Client.TransactOpts, l1BlockHash, l1SignalRoot, l1Height, parentGasUsed)
}

// Anchor is a paid mutator transaction binding the contract method 0xda69d3db.
//
// Solidity: function anchor(bytes32 l1BlockHash, bytes32 l1SignalRoot, uint64 l1Height, uint32 parentGasUsed) returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) Anchor(l1BlockHash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint32) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Anchor(&_TaikoL2Client.TransactOpts, l1BlockHash, l1SignalRoot, l1Height, parentGasUsed)
}

// Init is a paid mutator transaction binding the contract method 0xcce03bf3.
//
// Solidity: function init(address _addressManager, uint128 _gasExcess) returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _gasExcess *big.Int) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "init", _addressManager, _gasExcess)
}

// Init is a paid mutator transaction binding the contract method 0xcce03bf3.
//
// Solidity: function init(address _addressManager, uint128 _gasExcess) returns()
func (_TaikoL2Client *TaikoL2ClientSession) Init(_addressManager common.Address, _gasExcess *big.Int) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Init(&_TaikoL2Client.TransactOpts, _addressManager, _gasExcess)
}

// Init is a paid mutator transaction binding the contract method 0xcce03bf3.
//
// Solidity: function init(address _addressManager, uint128 _gasExcess) returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) Init(_addressManager common.Address, _gasExcess *big.Int) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Init(&_TaikoL2Client.TransactOpts, _addressManager, _gasExcess)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL2Client *TaikoL2ClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoL2Client.Contract.RenounceOwnership(&_TaikoL2Client.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoL2Client.Contract.RenounceOwnership(&_TaikoL2Client.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL2Client *TaikoL2ClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.TransferOwnership(&_TaikoL2Client.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.TransferOwnership(&_TaikoL2Client.TransactOpts, newOwner)
}

// TaikoL2ClientAnchoredIterator is returned from FilterAnchored and is used to iterate over the raw logs and unpacked data for Anchored events raised by the TaikoL2Client contract.
type TaikoL2ClientAnchoredIterator struct {
	Event *TaikoL2ClientAnchored // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientAnchoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientAnchored)
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
		it.Event = new(TaikoL2ClientAnchored)
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
func (it *TaikoL2ClientAnchoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientAnchoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientAnchored represents a Anchored event raised by the TaikoL2Client contract.
type TaikoL2ClientAnchored struct {
	ParentHash  [32]byte
	GasExcess   *big.Int
	BlockReward *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAnchored is a free log retrieval operation binding the contract event 0xf3a1ab67cc789a3e3961363ffd80ff59c13f394ff689b2fb85531a45dd0aa80d.
//
// Solidity: event Anchored(bytes32 parentHash, uint128 gasExcess, uint128 blockReward)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterAnchored(opts *bind.FilterOpts) (*TaikoL2ClientAnchoredIterator, error) {

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "Anchored")
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientAnchoredIterator{contract: _TaikoL2Client.contract, event: "Anchored", logs: logs, sub: sub}, nil
}

// WatchAnchored is a free log subscription operation binding the contract event 0xf3a1ab67cc789a3e3961363ffd80ff59c13f394ff689b2fb85531a45dd0aa80d.
//
// Solidity: event Anchored(bytes32 parentHash, uint128 gasExcess, uint128 blockReward)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchAnchored(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientAnchored) (event.Subscription, error) {

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "Anchored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientAnchored)
				if err := _TaikoL2Client.contract.UnpackLog(event, "Anchored", log); err != nil {
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

// ParseAnchored is a log parse operation binding the contract event 0xf3a1ab67cc789a3e3961363ffd80ff59c13f394ff689b2fb85531a45dd0aa80d.
//
// Solidity: event Anchored(bytes32 parentHash, uint128 gasExcess, uint128 blockReward)
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseAnchored(log types.Log) (*TaikoL2ClientAnchored, error) {
	event := new(TaikoL2ClientAnchored)
	if err := _TaikoL2Client.contract.UnpackLog(event, "Anchored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientCrossChainSyncedIterator is returned from FilterCrossChainSynced and is used to iterate over the raw logs and unpacked data for CrossChainSynced events raised by the TaikoL2Client contract.
type TaikoL2ClientCrossChainSyncedIterator struct {
	Event *TaikoL2ClientCrossChainSynced // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientCrossChainSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientCrossChainSynced)
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
		it.Event = new(TaikoL2ClientCrossChainSynced)
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
func (it *TaikoL2ClientCrossChainSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientCrossChainSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientCrossChainSynced represents a CrossChainSynced event raised by the TaikoL2Client contract.
type TaikoL2ClientCrossChainSynced struct {
	SrcHeight  uint64
	BlockHash  [32]byte
	SignalRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced is a free log retrieval operation binding the contract event 0x004ce985b8852a486571d0545799251fd671adcf33b7854a5f0f6a6a2431a555.
//
// Solidity: event CrossChainSynced(uint64 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterCrossChainSynced(opts *bind.FilterOpts, srcHeight []uint64) (*TaikoL2ClientCrossChainSyncedIterator, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientCrossChainSyncedIterator{contract: _TaikoL2Client.contract, event: "CrossChainSynced", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced is a free log subscription operation binding the contract event 0x004ce985b8852a486571d0545799251fd671adcf33b7854a5f0f6a6a2431a555.
//
// Solidity: event CrossChainSynced(uint64 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchCrossChainSynced(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientCrossChainSynced, srcHeight []uint64) (event.Subscription, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientCrossChainSynced)
				if err := _TaikoL2Client.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
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
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseCrossChainSynced(log types.Log) (*TaikoL2ClientCrossChainSynced, error) {
	event := new(TaikoL2ClientCrossChainSynced)
	if err := _TaikoL2Client.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaikoL2Client contract.
type TaikoL2ClientInitializedIterator struct {
	Event *TaikoL2ClientInitialized // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientInitialized)
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
		it.Event = new(TaikoL2ClientInitialized)
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
func (it *TaikoL2ClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientInitialized represents a Initialized event raised by the TaikoL2Client contract.
type TaikoL2ClientInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaikoL2ClientInitializedIterator, error) {

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientInitializedIterator{contract: _TaikoL2Client.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientInitialized) (event.Subscription, error) {

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientInitialized)
				if err := _TaikoL2Client.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseInitialized(log types.Log) (*TaikoL2ClientInitialized, error) {
	event := new(TaikoL2ClientInitialized)
	if err := _TaikoL2Client.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaikoL2Client contract.
type TaikoL2ClientOwnershipTransferredIterator struct {
	Event *TaikoL2ClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientOwnershipTransferred)
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
		it.Event = new(TaikoL2ClientOwnershipTransferred)
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
func (it *TaikoL2ClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientOwnershipTransferred represents a OwnershipTransferred event raised by the TaikoL2Client contract.
type TaikoL2ClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaikoL2ClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientOwnershipTransferredIterator{contract: _TaikoL2Client.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientOwnershipTransferred)
				if err := _TaikoL2Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseOwnershipTransferred(log types.Log) (*TaikoL2ClientOwnershipTransferred, error) {
	event := new(TaikoL2ClientOwnershipTransferred)
	if err := _TaikoL2Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
