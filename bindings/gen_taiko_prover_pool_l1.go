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

// ProverPoolProver is an auto generated low-level Go binding around an user-defined struct.
type ProverPoolProver struct {
	StakedAmount    uint32
	RewardPerGas    uint16
	CurrentCapacity uint16
}

// ProverPoolStaker is an auto generated low-level Go binding around an user-defined struct.
type ProverPoolStaker struct {
	ExitRequestedAt uint64
	ExitAmount      uint32
	MaxCapacity     uint16
	ProverId        uint8
}

// TaikoL1ProverPoolMetaData contains all meta data concerning the TaikoL1ProverPool contract.
var TaikoL1ProverPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"INVALID_PARAMS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NO_MATURE_EXIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PROVER_NOT_GOOD_ENOUGH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UNAUTHORIZED\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addressManager\",\"type\":\"address\"}],\"name\":\"AddressManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"Exited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rewardPerGas\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"currentCapacity\",\"type\":\"uint16\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EXIT_PERIOD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CAPACITY_LOWER_BOUND\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_PROVERS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_STAKE_PER_CAPACITY\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ONE_TKO\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_POINTS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"feePerGas\",\"type\":\"uint32\"}],\"name\":\"assignProver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"rewardPerGas\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProvers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"stakedAmount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"rewardPerGas\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"currentCapacity\",\"type\":\"uint16\"}],\"internalType\":\"structProverPool.Prover[]\",\"name\":\"_provers\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"_stakers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStaker\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"exitRequestedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"exitAmount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxCapacity\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"proverId\",\"type\":\"uint8\"}],\"internalType\":\"structProverPool.Staker\",\"name\":\"staker\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"stakedAmount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"rewardPerGas\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"currentCapacity\",\"type\":\"uint16\"}],\"internalType\":\"structProverPool.Prover\",\"name\":\"prover\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"feePerGas\",\"type\":\"uint32\"}],\"name\":\"getWeights\",\"outputs\":[{\"internalType\":\"uint256[32]\",\"name\":\"weights\",\"type\":\"uint256[32]\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"idToProver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"releaseProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"slashProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"rewardPerGas\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxCapacity\",\"type\":\"uint16\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"exitRequestedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"exitAmount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxCapacity\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"proverId\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TaikoL1ProverPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use TaikoL1ProverPoolMetaData.ABI instead.
var TaikoL1ProverPoolABI = TaikoL1ProverPoolMetaData.ABI

// TaikoL1ProverPool is an auto generated Go binding around an Ethereum contract.
type TaikoL1ProverPool struct {
	TaikoL1ProverPoolCaller     // Read-only binding to the contract
	TaikoL1ProverPoolTransactor // Write-only binding to the contract
	TaikoL1ProverPoolFilterer   // Log filterer for contract events
}

// TaikoL1ProverPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaikoL1ProverPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL1ProverPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaikoL1ProverPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL1ProverPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaikoL1ProverPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL1ProverPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaikoL1ProverPoolSession struct {
	Contract     *TaikoL1ProverPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TaikoL1ProverPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaikoL1ProverPoolCallerSession struct {
	Contract *TaikoL1ProverPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TaikoL1ProverPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaikoL1ProverPoolTransactorSession struct {
	Contract     *TaikoL1ProverPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TaikoL1ProverPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaikoL1ProverPoolRaw struct {
	Contract *TaikoL1ProverPool // Generic contract binding to access the raw methods on
}

// TaikoL1ProverPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaikoL1ProverPoolCallerRaw struct {
	Contract *TaikoL1ProverPoolCaller // Generic read-only contract binding to access the raw methods on
}

// TaikoL1ProverPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaikoL1ProverPoolTransactorRaw struct {
	Contract *TaikoL1ProverPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaikoL1ProverPool creates a new instance of TaikoL1ProverPool, bound to a specific deployed contract.
func NewTaikoL1ProverPool(address common.Address, backend bind.ContractBackend) (*TaikoL1ProverPool, error) {
	contract, err := bindTaikoL1ProverPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPool{TaikoL1ProverPoolCaller: TaikoL1ProverPoolCaller{contract: contract}, TaikoL1ProverPoolTransactor: TaikoL1ProverPoolTransactor{contract: contract}, TaikoL1ProverPoolFilterer: TaikoL1ProverPoolFilterer{contract: contract}}, nil
}

// NewTaikoL1ProverPoolCaller creates a new read-only instance of TaikoL1ProverPool, bound to a specific deployed contract.
func NewTaikoL1ProverPoolCaller(address common.Address, caller bind.ContractCaller) (*TaikoL1ProverPoolCaller, error) {
	contract, err := bindTaikoL1ProverPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolCaller{contract: contract}, nil
}

// NewTaikoL1ProverPoolTransactor creates a new write-only instance of TaikoL1ProverPool, bound to a specific deployed contract.
func NewTaikoL1ProverPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*TaikoL1ProverPoolTransactor, error) {
	contract, err := bindTaikoL1ProverPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolTransactor{contract: contract}, nil
}

// NewTaikoL1ProverPoolFilterer creates a new log filterer instance of TaikoL1ProverPool, bound to a specific deployed contract.
func NewTaikoL1ProverPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*TaikoL1ProverPoolFilterer, error) {
	contract, err := bindTaikoL1ProverPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolFilterer{contract: contract}, nil
}

// bindTaikoL1ProverPool binds a generic wrapper to an already deployed contract.
func bindTaikoL1ProverPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaikoL1ProverPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoL1ProverPool *TaikoL1ProverPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoL1ProverPool.Contract.TaikoL1ProverPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoL1ProverPool *TaikoL1ProverPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.TaikoL1ProverPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoL1ProverPool *TaikoL1ProverPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.TaikoL1ProverPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoL1ProverPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.contract.Transact(opts, method, params...)
}

// EXITPERIOD is a free data retrieval call binding the contract method 0xc04b5f65.
//
// Solidity: function EXIT_PERIOD() view returns(uint64)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) EXITPERIOD(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "EXIT_PERIOD")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// EXITPERIOD is a free data retrieval call binding the contract method 0xc04b5f65.
//
// Solidity: function EXIT_PERIOD() view returns(uint64)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) EXITPERIOD() (uint64, error) {
	return _TaikoL1ProverPool.Contract.EXITPERIOD(&_TaikoL1ProverPool.CallOpts)
}

// EXITPERIOD is a free data retrieval call binding the contract method 0xc04b5f65.
//
// Solidity: function EXIT_PERIOD() view returns(uint64)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) EXITPERIOD() (uint64, error) {
	return _TaikoL1ProverPool.Contract.EXITPERIOD(&_TaikoL1ProverPool.CallOpts)
}

// MAXCAPACITYLOWERBOUND is a free data retrieval call binding the contract method 0x35acc933.
//
// Solidity: function MAX_CAPACITY_LOWER_BOUND() view returns(uint32)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) MAXCAPACITYLOWERBOUND(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "MAX_CAPACITY_LOWER_BOUND")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXCAPACITYLOWERBOUND is a free data retrieval call binding the contract method 0x35acc933.
//
// Solidity: function MAX_CAPACITY_LOWER_BOUND() view returns(uint32)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) MAXCAPACITYLOWERBOUND() (uint32, error) {
	return _TaikoL1ProverPool.Contract.MAXCAPACITYLOWERBOUND(&_TaikoL1ProverPool.CallOpts)
}

// MAXCAPACITYLOWERBOUND is a free data retrieval call binding the contract method 0x35acc933.
//
// Solidity: function MAX_CAPACITY_LOWER_BOUND() view returns(uint32)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) MAXCAPACITYLOWERBOUND() (uint32, error) {
	return _TaikoL1ProverPool.Contract.MAXCAPACITYLOWERBOUND(&_TaikoL1ProverPool.CallOpts)
}

// MAXNUMPROVERS is a free data retrieval call binding the contract method 0x62c0fd98.
//
// Solidity: function MAX_NUM_PROVERS() view returns(uint256)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) MAXNUMPROVERS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "MAX_NUM_PROVERS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNUMPROVERS is a free data retrieval call binding the contract method 0x62c0fd98.
//
// Solidity: function MAX_NUM_PROVERS() view returns(uint256)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) MAXNUMPROVERS() (*big.Int, error) {
	return _TaikoL1ProverPool.Contract.MAXNUMPROVERS(&_TaikoL1ProverPool.CallOpts)
}

// MAXNUMPROVERS is a free data retrieval call binding the contract method 0x62c0fd98.
//
// Solidity: function MAX_NUM_PROVERS() view returns(uint256)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) MAXNUMPROVERS() (*big.Int, error) {
	return _TaikoL1ProverPool.Contract.MAXNUMPROVERS(&_TaikoL1ProverPool.CallOpts)
}

// MINSTAKEPERCAPACITY is a free data retrieval call binding the contract method 0x7d62c057.
//
// Solidity: function MIN_STAKE_PER_CAPACITY() view returns(uint32)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) MINSTAKEPERCAPACITY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "MIN_STAKE_PER_CAPACITY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MINSTAKEPERCAPACITY is a free data retrieval call binding the contract method 0x7d62c057.
//
// Solidity: function MIN_STAKE_PER_CAPACITY() view returns(uint32)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) MINSTAKEPERCAPACITY() (uint32, error) {
	return _TaikoL1ProverPool.Contract.MINSTAKEPERCAPACITY(&_TaikoL1ProverPool.CallOpts)
}

// MINSTAKEPERCAPACITY is a free data retrieval call binding the contract method 0x7d62c057.
//
// Solidity: function MIN_STAKE_PER_CAPACITY() view returns(uint32)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) MINSTAKEPERCAPACITY() (uint32, error) {
	return _TaikoL1ProverPool.Contract.MINSTAKEPERCAPACITY(&_TaikoL1ProverPool.CallOpts)
}

// ONETKO is a free data retrieval call binding the contract method 0xe22fcaaa.
//
// Solidity: function ONE_TKO() view returns(uint64)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) ONETKO(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "ONE_TKO")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ONETKO is a free data retrieval call binding the contract method 0xe22fcaaa.
//
// Solidity: function ONE_TKO() view returns(uint64)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) ONETKO() (uint64, error) {
	return _TaikoL1ProverPool.Contract.ONETKO(&_TaikoL1ProverPool.CallOpts)
}

// ONETKO is a free data retrieval call binding the contract method 0xe22fcaaa.
//
// Solidity: function ONE_TKO() view returns(uint64)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) ONETKO() (uint64, error) {
	return _TaikoL1ProverPool.Contract.ONETKO(&_TaikoL1ProverPool.CallOpts)
}

// SLASHPOINTS is a free data retrieval call binding the contract method 0xdd9fb65c.
//
// Solidity: function SLASH_POINTS() view returns(uint32)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) SLASHPOINTS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "SLASH_POINTS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SLASHPOINTS is a free data retrieval call binding the contract method 0xdd9fb65c.
//
// Solidity: function SLASH_POINTS() view returns(uint32)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) SLASHPOINTS() (uint32, error) {
	return _TaikoL1ProverPool.Contract.SLASHPOINTS(&_TaikoL1ProverPool.CallOpts)
}

// SLASHPOINTS is a free data retrieval call binding the contract method 0xdd9fb65c.
//
// Solidity: function SLASH_POINTS() view returns(uint32)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) SLASHPOINTS() (uint32, error) {
	return _TaikoL1ProverPool.Contract.SLASHPOINTS(&_TaikoL1ProverPool.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) AddressManager() (common.Address, error) {
	return _TaikoL1ProverPool.Contract.AddressManager(&_TaikoL1ProverPool.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) AddressManager() (common.Address, error) {
	return _TaikoL1ProverPool.Contract.AddressManager(&_TaikoL1ProverPool.CallOpts)
}

// GetCapacity is a free data retrieval call binding the contract method 0xc40000d4.
//
// Solidity: function getCapacity() view returns(uint256 capacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) GetCapacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "getCapacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCapacity is a free data retrieval call binding the contract method 0xc40000d4.
//
// Solidity: function getCapacity() view returns(uint256 capacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) GetCapacity() (*big.Int, error) {
	return _TaikoL1ProverPool.Contract.GetCapacity(&_TaikoL1ProverPool.CallOpts)
}

// GetCapacity is a free data retrieval call binding the contract method 0xc40000d4.
//
// Solidity: function getCapacity() view returns(uint256 capacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) GetCapacity() (*big.Int, error) {
	return _TaikoL1ProverPool.Contract.GetCapacity(&_TaikoL1ProverPool.CallOpts)
}

// GetProvers is a free data retrieval call binding the contract method 0xc0bfd036.
//
// Solidity: function getProvers() view returns((uint32,uint16,uint16)[] _provers, address[] _stakers)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) GetProvers(opts *bind.CallOpts) (struct {
	Provers []ProverPoolProver
	Stakers []common.Address
}, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "getProvers")

	outstruct := new(struct {
		Provers []ProverPoolProver
		Stakers []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Provers = *abi.ConvertType(out[0], new([]ProverPoolProver)).(*[]ProverPoolProver)
	outstruct.Stakers = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetProvers is a free data retrieval call binding the contract method 0xc0bfd036.
//
// Solidity: function getProvers() view returns((uint32,uint16,uint16)[] _provers, address[] _stakers)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) GetProvers() (struct {
	Provers []ProverPoolProver
	Stakers []common.Address
}, error) {
	return _TaikoL1ProverPool.Contract.GetProvers(&_TaikoL1ProverPool.CallOpts)
}

// GetProvers is a free data retrieval call binding the contract method 0xc0bfd036.
//
// Solidity: function getProvers() view returns((uint32,uint16,uint16)[] _provers, address[] _stakers)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) GetProvers() (struct {
	Provers []ProverPoolProver
	Stakers []common.Address
}, error) {
	return _TaikoL1ProverPool.Contract.GetProvers(&_TaikoL1ProverPool.CallOpts)
}

// GetStaker is a free data retrieval call binding the contract method 0xa23c44b1.
//
// Solidity: function getStaker(address addr) view returns((uint64,uint32,uint16,uint8) staker, (uint32,uint16,uint16) prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) GetStaker(opts *bind.CallOpts, addr common.Address) (struct {
	Staker ProverPoolStaker
	Prover ProverPoolProver
}, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "getStaker", addr)

	outstruct := new(struct {
		Staker ProverPoolStaker
		Prover ProverPoolProver
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Staker = *abi.ConvertType(out[0], new(ProverPoolStaker)).(*ProverPoolStaker)
	outstruct.Prover = *abi.ConvertType(out[1], new(ProverPoolProver)).(*ProverPoolProver)

	return *outstruct, err

}

// GetStaker is a free data retrieval call binding the contract method 0xa23c44b1.
//
// Solidity: function getStaker(address addr) view returns((uint64,uint32,uint16,uint8) staker, (uint32,uint16,uint16) prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) GetStaker(addr common.Address) (struct {
	Staker ProverPoolStaker
	Prover ProverPoolProver
}, error) {
	return _TaikoL1ProverPool.Contract.GetStaker(&_TaikoL1ProverPool.CallOpts, addr)
}

// GetStaker is a free data retrieval call binding the contract method 0xa23c44b1.
//
// Solidity: function getStaker(address addr) view returns((uint64,uint32,uint16,uint8) staker, (uint32,uint16,uint16) prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) GetStaker(addr common.Address) (struct {
	Staker ProverPoolStaker
	Prover ProverPoolProver
}, error) {
	return _TaikoL1ProverPool.Contract.GetStaker(&_TaikoL1ProverPool.CallOpts, addr)
}

// GetWeights is a free data retrieval call binding the contract method 0xc2082974.
//
// Solidity: function getWeights(uint32 feePerGas) view returns(uint256[32] weights, uint256 totalWeight)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) GetWeights(opts *bind.CallOpts, feePerGas uint32) (struct {
	Weights     [32]*big.Int
	TotalWeight *big.Int
}, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "getWeights", feePerGas)

	outstruct := new(struct {
		Weights     [32]*big.Int
		TotalWeight *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weights = *abi.ConvertType(out[0], new([32]*big.Int)).(*[32]*big.Int)
	outstruct.TotalWeight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetWeights is a free data retrieval call binding the contract method 0xc2082974.
//
// Solidity: function getWeights(uint32 feePerGas) view returns(uint256[32] weights, uint256 totalWeight)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) GetWeights(feePerGas uint32) (struct {
	Weights     [32]*big.Int
	TotalWeight *big.Int
}, error) {
	return _TaikoL1ProverPool.Contract.GetWeights(&_TaikoL1ProverPool.CallOpts, feePerGas)
}

// GetWeights is a free data retrieval call binding the contract method 0xc2082974.
//
// Solidity: function getWeights(uint32 feePerGas) view returns(uint256[32] weights, uint256 totalWeight)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) GetWeights(feePerGas uint32) (struct {
	Weights     [32]*big.Int
	TotalWeight *big.Int
}, error) {
	return _TaikoL1ProverPool.Contract.GetWeights(&_TaikoL1ProverPool.CallOpts, feePerGas)
}

// IdToProver is a free data retrieval call binding the contract method 0x2f88a7fe.
//
// Solidity: function idToProver(uint256 id) view returns(address prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) IdToProver(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "idToProver", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdToProver is a free data retrieval call binding the contract method 0x2f88a7fe.
//
// Solidity: function idToProver(uint256 id) view returns(address prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) IdToProver(id *big.Int) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.IdToProver(&_TaikoL1ProverPool.CallOpts, id)
}

// IdToProver is a free data retrieval call binding the contract method 0x2f88a7fe.
//
// Solidity: function idToProver(uint256 id) view returns(address prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) IdToProver(id *big.Int) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.IdToProver(&_TaikoL1ProverPool.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Owner() (common.Address, error) {
	return _TaikoL1ProverPool.Contract.Owner(&_TaikoL1ProverPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) Owner() (common.Address, error) {
	return _TaikoL1ProverPool.Contract.Owner(&_TaikoL1ProverPool.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) Resolve(opts *bind.CallOpts, chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.Resolve(&_TaikoL1ProverPool.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.Resolve(&_TaikoL1ProverPool.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.Resolve0(&_TaikoL1ProverPool.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.Resolve0(&_TaikoL1ProverPool.CallOpts, name, allowZeroAddress)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(uint64 exitRequestedAt, uint32 exitAmount, uint16 maxCapacity, uint8 proverId)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) Stakers(opts *bind.CallOpts, staker common.Address) (struct {
	ExitRequestedAt uint64
	ExitAmount      uint32
	MaxCapacity     uint16
	ProverId        uint8
}, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "stakers", staker)

	outstruct := new(struct {
		ExitRequestedAt uint64
		ExitAmount      uint32
		MaxCapacity     uint16
		ProverId        uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExitRequestedAt = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ExitAmount = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.MaxCapacity = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ProverId = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(uint64 exitRequestedAt, uint32 exitAmount, uint16 maxCapacity, uint8 proverId)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Stakers(staker common.Address) (struct {
	ExitRequestedAt uint64
	ExitAmount      uint32
	MaxCapacity     uint16
	ProverId        uint8
}, error) {
	return _TaikoL1ProverPool.Contract.Stakers(&_TaikoL1ProverPool.CallOpts, staker)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(uint64 exitRequestedAt, uint32 exitAmount, uint16 maxCapacity, uint8 proverId)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) Stakers(staker common.Address) (struct {
	ExitRequestedAt uint64
	ExitAmount      uint32
	MaxCapacity     uint16
	ProverId        uint8
}, error) {
	return _TaikoL1ProverPool.Contract.Stakers(&_TaikoL1ProverPool.CallOpts, staker)
}

// AssignProver is a paid mutator transaction binding the contract method 0xbd849fe9.
//
// Solidity: function assignProver(uint64 blockId, uint32 feePerGas) returns(address prover, uint32 rewardPerGas)
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) AssignProver(opts *bind.TransactOpts, blockId uint64, feePerGas uint32) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "assignProver", blockId, feePerGas)
}

// AssignProver is a paid mutator transaction binding the contract method 0xbd849fe9.
//
// Solidity: function assignProver(uint64 blockId, uint32 feePerGas) returns(address prover, uint32 rewardPerGas)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) AssignProver(blockId uint64, feePerGas uint32) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.AssignProver(&_TaikoL1ProverPool.TransactOpts, blockId, feePerGas)
}

// AssignProver is a paid mutator transaction binding the contract method 0xbd849fe9.
//
// Solidity: function assignProver(uint64 blockId, uint32 feePerGas) returns(address prover, uint32 rewardPerGas)
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) AssignProver(blockId uint64, feePerGas uint32) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.AssignProver(&_TaikoL1ProverPool.TransactOpts, blockId, feePerGas)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Exit() (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Exit(&_TaikoL1ProverPool.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) Exit() (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Exit(&_TaikoL1ProverPool.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _addressManager) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "init", _addressManager)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _addressManager) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Init(_addressManager common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Init(&_TaikoL1ProverPool.TransactOpts, _addressManager)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _addressManager) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) Init(_addressManager common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Init(&_TaikoL1ProverPool.TransactOpts, _addressManager)
}

// ReleaseProver is a paid mutator transaction binding the contract method 0xcba0414f.
//
// Solidity: function releaseProver(address addr) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) ReleaseProver(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "releaseProver", addr)
}

// ReleaseProver is a paid mutator transaction binding the contract method 0xcba0414f.
//
// Solidity: function releaseProver(address addr) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) ReleaseProver(addr common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.ReleaseProver(&_TaikoL1ProverPool.TransactOpts, addr)
}

// ReleaseProver is a paid mutator transaction binding the contract method 0xcba0414f.
//
// Solidity: function releaseProver(address addr) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) ReleaseProver(addr common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.ReleaseProver(&_TaikoL1ProverPool.TransactOpts, addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.RenounceOwnership(&_TaikoL1ProverPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.RenounceOwnership(&_TaikoL1ProverPool.TransactOpts)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) SetAddressManager(opts *bind.TransactOpts, newAddressManager common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "setAddressManager", newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.SetAddressManager(&_TaikoL1ProverPool.TransactOpts, newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.SetAddressManager(&_TaikoL1ProverPool.TransactOpts, newAddressManager)
}

// SlashProver is a paid mutator transaction binding the contract method 0xcd362a5b.
//
// Solidity: function slashProver(address addr) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) SlashProver(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "slashProver", addr)
}

// SlashProver is a paid mutator transaction binding the contract method 0xcd362a5b.
//
// Solidity: function slashProver(address addr) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) SlashProver(addr common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.SlashProver(&_TaikoL1ProverPool.TransactOpts, addr)
}

// SlashProver is a paid mutator transaction binding the contract method 0xcd362a5b.
//
// Solidity: function slashProver(address addr) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) SlashProver(addr common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.SlashProver(&_TaikoL1ProverPool.TransactOpts, addr)
}

// Stake is a paid mutator transaction binding the contract method 0x3ca316f1.
//
// Solidity: function stake(uint32 amount, uint16 rewardPerGas, uint16 maxCapacity) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) Stake(opts *bind.TransactOpts, amount uint32, rewardPerGas uint16, maxCapacity uint16) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "stake", amount, rewardPerGas, maxCapacity)
}

// Stake is a paid mutator transaction binding the contract method 0x3ca316f1.
//
// Solidity: function stake(uint32 amount, uint16 rewardPerGas, uint16 maxCapacity) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Stake(amount uint32, rewardPerGas uint16, maxCapacity uint16) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Stake(&_TaikoL1ProverPool.TransactOpts, amount, rewardPerGas, maxCapacity)
}

// Stake is a paid mutator transaction binding the contract method 0x3ca316f1.
//
// Solidity: function stake(uint32 amount, uint16 rewardPerGas, uint16 maxCapacity) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) Stake(amount uint32, rewardPerGas uint16, maxCapacity uint16) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Stake(&_TaikoL1ProverPool.TransactOpts, amount, rewardPerGas, maxCapacity)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.TransferOwnership(&_TaikoL1ProverPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.TransferOwnership(&_TaikoL1ProverPool.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Withdraw() (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Withdraw(&_TaikoL1ProverPool.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) Withdraw() (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Withdraw(&_TaikoL1ProverPool.TransactOpts)
}

// TaikoL1ProverPoolAddressManagerChangedIterator is returned from FilterAddressManagerChanged and is used to iterate over the raw logs and unpacked data for AddressManagerChanged events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolAddressManagerChangedIterator struct {
	Event *TaikoL1ProverPoolAddressManagerChanged // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolAddressManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolAddressManagerChanged)
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
		it.Event = new(TaikoL1ProverPoolAddressManagerChanged)
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
func (it *TaikoL1ProverPoolAddressManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolAddressManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolAddressManagerChanged represents a AddressManagerChanged event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolAddressManagerChanged struct {
	AddressManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressManagerChanged is a free log retrieval operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterAddressManagerChanged(opts *bind.FilterOpts) (*TaikoL1ProverPoolAddressManagerChangedIterator, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "AddressManagerChanged")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolAddressManagerChangedIterator{contract: _TaikoL1ProverPool.contract, event: "AddressManagerChanged", logs: logs, sub: sub}, nil
}

// WatchAddressManagerChanged is a free log subscription operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchAddressManagerChanged(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolAddressManagerChanged) (event.Subscription, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "AddressManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolAddressManagerChanged)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
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
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseAddressManagerChanged(log types.Log) (*TaikoL1ProverPoolAddressManagerChanged, error) {
	event := new(TaikoL1ProverPoolAddressManagerChanged)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolExitedIterator is returned from FilterExited and is used to iterate over the raw logs and unpacked data for Exited events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolExitedIterator struct {
	Event *TaikoL1ProverPoolExited // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolExited)
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
		it.Event = new(TaikoL1ProverPoolExited)
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
func (it *TaikoL1ProverPoolExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolExited represents a Exited event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolExited struct {
	Addr   common.Address
	Amount uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExited is a free log retrieval operation binding the contract event 0xfa715c30456da257872290458c962d9495b3a623e7df06cacb5b8dd68668554d.
//
// Solidity: event Exited(address indexed addr, uint32 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterExited(opts *bind.FilterOpts, addr []common.Address) (*TaikoL1ProverPoolExitedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "Exited", addrRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolExitedIterator{contract: _TaikoL1ProverPool.contract, event: "Exited", logs: logs, sub: sub}, nil
}

// WatchExited is a free log subscription operation binding the contract event 0xfa715c30456da257872290458c962d9495b3a623e7df06cacb5b8dd68668554d.
//
// Solidity: event Exited(address indexed addr, uint32 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchExited(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolExited, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "Exited", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolExited)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "Exited", log); err != nil {
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

// ParseExited is a log parse operation binding the contract event 0xfa715c30456da257872290458c962d9495b3a623e7df06cacb5b8dd68668554d.
//
// Solidity: event Exited(address indexed addr, uint32 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseExited(log types.Log) (*TaikoL1ProverPoolExited, error) {
	event := new(TaikoL1ProverPoolExited)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "Exited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolInitializedIterator struct {
	Event *TaikoL1ProverPoolInitialized // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolInitialized)
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
		it.Event = new(TaikoL1ProverPoolInitialized)
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
func (it *TaikoL1ProverPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolInitialized represents a Initialized event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaikoL1ProverPoolInitializedIterator, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolInitializedIterator{contract: _TaikoL1ProverPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolInitialized)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseInitialized(log types.Log) (*TaikoL1ProverPoolInitialized, error) {
	event := new(TaikoL1ProverPoolInitialized)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolOwnershipTransferredIterator struct {
	Event *TaikoL1ProverPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolOwnershipTransferred)
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
		it.Event = new(TaikoL1ProverPoolOwnershipTransferred)
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
func (it *TaikoL1ProverPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolOwnershipTransferred represents a OwnershipTransferred event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaikoL1ProverPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolOwnershipTransferredIterator{contract: _TaikoL1ProverPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolOwnershipTransferred)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseOwnershipTransferred(log types.Log) (*TaikoL1ProverPoolOwnershipTransferred, error) {
	event := new(TaikoL1ProverPoolOwnershipTransferred)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolSlashedIterator struct {
	Event *TaikoL1ProverPoolSlashed // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolSlashed)
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
		it.Event = new(TaikoL1ProverPoolSlashed)
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
func (it *TaikoL1ProverPoolSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolSlashed represents a Slashed event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolSlashed struct {
	Addr   common.Address
	Amount uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0xf2ca0cd8f45e853110fce6e92f38f84a31350c1168dce2735a782f9d6424e648.
//
// Solidity: event Slashed(address indexed addr, uint32 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterSlashed(opts *bind.FilterOpts, addr []common.Address) (*TaikoL1ProverPoolSlashedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "Slashed", addrRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolSlashedIterator{contract: _TaikoL1ProverPool.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0xf2ca0cd8f45e853110fce6e92f38f84a31350c1168dce2735a782f9d6424e648.
//
// Solidity: event Slashed(address indexed addr, uint32 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolSlashed, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "Slashed", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolSlashed)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0xf2ca0cd8f45e853110fce6e92f38f84a31350c1168dce2735a782f9d6424e648.
//
// Solidity: event Slashed(address indexed addr, uint32 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseSlashed(log types.Log) (*TaikoL1ProverPoolSlashed, error) {
	event := new(TaikoL1ProverPoolSlashed)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolStakedIterator struct {
	Event *TaikoL1ProverPoolStaked // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolStaked)
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
		it.Event = new(TaikoL1ProverPoolStaked)
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
func (it *TaikoL1ProverPoolStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolStaked represents a Staked event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolStaked struct {
	Addr            common.Address
	Amount          uint32
	RewardPerGas    uint16
	CurrentCapacity uint16
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xe2232b717d57a877785c68aad0aae9fa2caea5035fa13efd62421b4e0ac23dfe.
//
// Solidity: event Staked(address indexed addr, uint32 amount, uint16 rewardPerGas, uint16 currentCapacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterStaked(opts *bind.FilterOpts, addr []common.Address) (*TaikoL1ProverPoolStakedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "Staked", addrRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolStakedIterator{contract: _TaikoL1ProverPool.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xe2232b717d57a877785c68aad0aae9fa2caea5035fa13efd62421b4e0ac23dfe.
//
// Solidity: event Staked(address indexed addr, uint32 amount, uint16 rewardPerGas, uint16 currentCapacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolStaked, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "Staked", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolStaked)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xe2232b717d57a877785c68aad0aae9fa2caea5035fa13efd62421b4e0ac23dfe.
//
// Solidity: event Staked(address indexed addr, uint32 amount, uint16 rewardPerGas, uint16 currentCapacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseStaked(log types.Log) (*TaikoL1ProverPoolStaked, error) {
	event := new(TaikoL1ProverPoolStaked)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolWithdrawnIterator struct {
	Event *TaikoL1ProverPoolWithdrawn // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolWithdrawn)
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
		it.Event = new(TaikoL1ProverPoolWithdrawn)
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
func (it *TaikoL1ProverPoolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolWithdrawn represents a Withdrawn event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolWithdrawn struct {
	Addr   common.Address
	Amount uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x28f8f642e0ef98c9c0d14e1bd28a687ea91da388d69e6ca60165b772ef408c7c.
//
// Solidity: event Withdrawn(address indexed addr, uint32 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterWithdrawn(opts *bind.FilterOpts, addr []common.Address) (*TaikoL1ProverPoolWithdrawnIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "Withdrawn", addrRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolWithdrawnIterator{contract: _TaikoL1ProverPool.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x28f8f642e0ef98c9c0d14e1bd28a687ea91da388d69e6ca60165b772ef408c7c.
//
// Solidity: event Withdrawn(address indexed addr, uint32 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolWithdrawn, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "Withdrawn", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolWithdrawn)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x28f8f642e0ef98c9c0d14e1bd28a687ea91da388d69e6ca60165b772ef408c7c.
//
// Solidity: event Withdrawn(address indexed addr, uint32 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseWithdrawn(log types.Log) (*TaikoL1ProverPoolWithdrawn, error) {
	event := new(TaikoL1ProverPoolWithdrawn)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
