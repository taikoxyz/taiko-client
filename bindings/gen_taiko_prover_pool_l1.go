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

// TaikoL1ProverPoolMetaData contains all meta data concerning the TaikoL1ProverPool contract.
var TaikoL1ProverPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addressManager\",\"type\":\"address\"}],\"name\":\"AddressManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"oldCapacity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newCapacity\",\"type\":\"uint32\"}],\"name\":\"ProverAdjustedCapacity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeMultiplier\",\"type\":\"uint256\"}],\"name\":\"ProverAdjustedFeeMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"capacity\",\"type\":\"uint64\"}],\"name\":\"ProverEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"ProverExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProverSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"}],\"name\":\"ProverStakedMoreTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProverWithdrawAwards\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_TKO_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newCapacity\",\"type\":\"uint32\"}],\"name\":\"adjustCapacity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newFeeMultiplier\",\"type\":\"uint8\"}],\"name\":\"adjustFeeMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"blockIdToProver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"capacity\",\"type\":\"uint32\"}],\"name\":\"enterProverPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getProver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_maxPoolSize\",\"type\":\"uint16\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPoolSize\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"pickRandomProver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"provers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proverAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakedTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBlockTsToBeProven\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"capacity\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numAssignedBlocks\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"feeMultiplier\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proversInPool\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeMoreTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"topProvers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// MAXMULTIPLIER is a free data retrieval call binding the contract method 0x5d6a618d.
//
// Solidity: function MAX_MULTIPLIER() view returns(uint8)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) MAXMULTIPLIER(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "MAX_MULTIPLIER")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXMULTIPLIER is a free data retrieval call binding the contract method 0x5d6a618d.
//
// Solidity: function MAX_MULTIPLIER() view returns(uint8)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) MAXMULTIPLIER() (uint8, error) {
	return _TaikoL1ProverPool.Contract.MAXMULTIPLIER(&_TaikoL1ProverPool.CallOpts)
}

// MAXMULTIPLIER is a free data retrieval call binding the contract method 0x5d6a618d.
//
// Solidity: function MAX_MULTIPLIER() view returns(uint8)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) MAXMULTIPLIER() (uint8, error) {
	return _TaikoL1ProverPool.Contract.MAXMULTIPLIER(&_TaikoL1ProverPool.CallOpts)
}

// MINMULTIPLIER is a free data retrieval call binding the contract method 0xed03e78c.
//
// Solidity: function MIN_MULTIPLIER() view returns(uint8)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) MINMULTIPLIER(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "MIN_MULTIPLIER")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MINMULTIPLIER is a free data retrieval call binding the contract method 0xed03e78c.
//
// Solidity: function MIN_MULTIPLIER() view returns(uint8)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) MINMULTIPLIER() (uint8, error) {
	return _TaikoL1ProverPool.Contract.MINMULTIPLIER(&_TaikoL1ProverPool.CallOpts)
}

// MINMULTIPLIER is a free data retrieval call binding the contract method 0xed03e78c.
//
// Solidity: function MIN_MULTIPLIER() view returns(uint8)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) MINMULTIPLIER() (uint8, error) {
	return _TaikoL1ProverPool.Contract.MINMULTIPLIER(&_TaikoL1ProverPool.CallOpts)
}

// MINTKOAMOUNT is a free data retrieval call binding the contract method 0x005eb8f8.
//
// Solidity: function MIN_TKO_AMOUNT() view returns(uint256)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) MINTKOAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "MIN_TKO_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTKOAMOUNT is a free data retrieval call binding the contract method 0x005eb8f8.
//
// Solidity: function MIN_TKO_AMOUNT() view returns(uint256)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) MINTKOAMOUNT() (*big.Int, error) {
	return _TaikoL1ProverPool.Contract.MINTKOAMOUNT(&_TaikoL1ProverPool.CallOpts)
}

// MINTKOAMOUNT is a free data retrieval call binding the contract method 0x005eb8f8.
//
// Solidity: function MIN_TKO_AMOUNT() view returns(uint256)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) MINTKOAMOUNT() (*big.Int, error) {
	return _TaikoL1ProverPool.Contract.MINTKOAMOUNT(&_TaikoL1ProverPool.CallOpts)
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

// BlockIdToProver is a free data retrieval call binding the contract method 0xf33ed81e.
//
// Solidity: function blockIdToProver(uint256 blockId) view returns(address prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) BlockIdToProver(opts *bind.CallOpts, blockId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "blockIdToProver", blockId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockIdToProver is a free data retrieval call binding the contract method 0xf33ed81e.
//
// Solidity: function blockIdToProver(uint256 blockId) view returns(address prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) BlockIdToProver(blockId *big.Int) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.BlockIdToProver(&_TaikoL1ProverPool.CallOpts, blockId)
}

// BlockIdToProver is a free data retrieval call binding the contract method 0xf33ed81e.
//
// Solidity: function blockIdToProver(uint256 blockId) view returns(address prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) BlockIdToProver(blockId *big.Int) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.BlockIdToProver(&_TaikoL1ProverPool.CallOpts, blockId)
}

// GetProver is a free data retrieval call binding the contract method 0xe02f1931.
//
// Solidity: function getProver(uint256 blockId) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) GetProver(opts *bind.CallOpts, blockId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "getProver", blockId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProver is a free data retrieval call binding the contract method 0xe02f1931.
//
// Solidity: function getProver(uint256 blockId) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) GetProver(blockId *big.Int) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.GetProver(&_TaikoL1ProverPool.CallOpts, blockId)
}

// GetProver is a free data retrieval call binding the contract method 0xe02f1931.
//
// Solidity: function getProver(uint256 blockId) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) GetProver(blockId *big.Int) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.GetProver(&_TaikoL1ProverPool.CallOpts, blockId)
}

// MaxPoolSize is a free data retrieval call binding the contract method 0xc5579dc0.
//
// Solidity: function maxPoolSize() view returns(uint16)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) MaxPoolSize(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "maxPoolSize")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxPoolSize is a free data retrieval call binding the contract method 0xc5579dc0.
//
// Solidity: function maxPoolSize() view returns(uint16)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) MaxPoolSize() (uint16, error) {
	return _TaikoL1ProverPool.Contract.MaxPoolSize(&_TaikoL1ProverPool.CallOpts)
}

// MaxPoolSize is a free data retrieval call binding the contract method 0xc5579dc0.
//
// Solidity: function maxPoolSize() view returns(uint16)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) MaxPoolSize() (uint16, error) {
	return _TaikoL1ProverPool.Contract.MaxPoolSize(&_TaikoL1ProverPool.CallOpts)
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

// Provers is a free data retrieval call binding the contract method 0x1dec844b.
//
// Solidity: function provers(address ) view returns(address proverAddress, uint256 stakedTokens, uint256 rewards, uint256 healthScore, uint256 lastBlockTsToBeProven, uint32 capacity, uint32 numAssignedBlocks, uint8 feeMultiplier)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) Provers(opts *bind.CallOpts, arg0 common.Address) (struct {
	ProverAddress         common.Address
	StakedTokens          *big.Int
	Rewards               *big.Int
	HealthScore           *big.Int
	LastBlockTsToBeProven *big.Int
	Capacity              uint32
	NumAssignedBlocks     uint32
	FeeMultiplier         uint8
}, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "provers", arg0)

	outstruct := new(struct {
		ProverAddress         common.Address
		StakedTokens          *big.Int
		Rewards               *big.Int
		HealthScore           *big.Int
		LastBlockTsToBeProven *big.Int
		Capacity              uint32
		NumAssignedBlocks     uint32
		FeeMultiplier         uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProverAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StakedTokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Rewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.HealthScore = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastBlockTsToBeProven = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Capacity = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.NumAssignedBlocks = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.FeeMultiplier = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// Provers is a free data retrieval call binding the contract method 0x1dec844b.
//
// Solidity: function provers(address ) view returns(address proverAddress, uint256 stakedTokens, uint256 rewards, uint256 healthScore, uint256 lastBlockTsToBeProven, uint32 capacity, uint32 numAssignedBlocks, uint8 feeMultiplier)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Provers(arg0 common.Address) (struct {
	ProverAddress         common.Address
	StakedTokens          *big.Int
	Rewards               *big.Int
	HealthScore           *big.Int
	LastBlockTsToBeProven *big.Int
	Capacity              uint32
	NumAssignedBlocks     uint32
	FeeMultiplier         uint8
}, error) {
	return _TaikoL1ProverPool.Contract.Provers(&_TaikoL1ProverPool.CallOpts, arg0)
}

// Provers is a free data retrieval call binding the contract method 0x1dec844b.
//
// Solidity: function provers(address ) view returns(address proverAddress, uint256 stakedTokens, uint256 rewards, uint256 healthScore, uint256 lastBlockTsToBeProven, uint32 capacity, uint32 numAssignedBlocks, uint8 feeMultiplier)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) Provers(arg0 common.Address) (struct {
	ProverAddress         common.Address
	StakedTokens          *big.Int
	Rewards               *big.Int
	HealthScore           *big.Int
	LastBlockTsToBeProven *big.Int
	Capacity              uint32
	NumAssignedBlocks     uint32
	FeeMultiplier         uint8
}, error) {
	return _TaikoL1ProverPool.Contract.Provers(&_TaikoL1ProverPool.CallOpts, arg0)
}

// ProversInPool is a free data retrieval call binding the contract method 0xb0aa5fcc.
//
// Solidity: function proversInPool() view returns(uint16)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) ProversInPool(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "proversInPool")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ProversInPool is a free data retrieval call binding the contract method 0xb0aa5fcc.
//
// Solidity: function proversInPool() view returns(uint16)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) ProversInPool() (uint16, error) {
	return _TaikoL1ProverPool.Contract.ProversInPool(&_TaikoL1ProverPool.CallOpts)
}

// ProversInPool is a free data retrieval call binding the contract method 0xb0aa5fcc.
//
// Solidity: function proversInPool() view returns(uint16)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) ProversInPool() (uint16, error) {
	return _TaikoL1ProverPool.Contract.ProversInPool(&_TaikoL1ProverPool.CallOpts)
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

// TopProvers is a free data retrieval call binding the contract method 0xd70e4ea6.
//
// Solidity: function topProvers(uint256 ) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCaller) TopProvers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1ProverPool.contract.Call(opts, &out, "topProvers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TopProvers is a free data retrieval call binding the contract method 0xd70e4ea6.
//
// Solidity: function topProvers(uint256 ) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) TopProvers(arg0 *big.Int) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.TopProvers(&_TaikoL1ProverPool.CallOpts, arg0)
}

// TopProvers is a free data retrieval call binding the contract method 0xd70e4ea6.
//
// Solidity: function topProvers(uint256 ) view returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolCallerSession) TopProvers(arg0 *big.Int) (common.Address, error) {
	return _TaikoL1ProverPool.Contract.TopProvers(&_TaikoL1ProverPool.CallOpts, arg0)
}

// AdjustCapacity is a paid mutator transaction binding the contract method 0x426759e3.
//
// Solidity: function adjustCapacity(uint32 newCapacity) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) AdjustCapacity(opts *bind.TransactOpts, newCapacity uint32) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "adjustCapacity", newCapacity)
}

// AdjustCapacity is a paid mutator transaction binding the contract method 0x426759e3.
//
// Solidity: function adjustCapacity(uint32 newCapacity) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) AdjustCapacity(newCapacity uint32) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.AdjustCapacity(&_TaikoL1ProverPool.TransactOpts, newCapacity)
}

// AdjustCapacity is a paid mutator transaction binding the contract method 0x426759e3.
//
// Solidity: function adjustCapacity(uint32 newCapacity) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) AdjustCapacity(newCapacity uint32) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.AdjustCapacity(&_TaikoL1ProverPool.TransactOpts, newCapacity)
}

// AdjustFeeMultiplier is a paid mutator transaction binding the contract method 0x4e07a0ce.
//
// Solidity: function adjustFeeMultiplier(uint8 newFeeMultiplier) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) AdjustFeeMultiplier(opts *bind.TransactOpts, newFeeMultiplier uint8) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "adjustFeeMultiplier", newFeeMultiplier)
}

// AdjustFeeMultiplier is a paid mutator transaction binding the contract method 0x4e07a0ce.
//
// Solidity: function adjustFeeMultiplier(uint8 newFeeMultiplier) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) AdjustFeeMultiplier(newFeeMultiplier uint8) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.AdjustFeeMultiplier(&_TaikoL1ProverPool.TransactOpts, newFeeMultiplier)
}

// AdjustFeeMultiplier is a paid mutator transaction binding the contract method 0x4e07a0ce.
//
// Solidity: function adjustFeeMultiplier(uint8 newFeeMultiplier) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) AdjustFeeMultiplier(newFeeMultiplier uint8) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.AdjustFeeMultiplier(&_TaikoL1ProverPool.TransactOpts, newFeeMultiplier)
}

// EnterProverPool is a paid mutator transaction binding the contract method 0xd12037ff.
//
// Solidity: function enterProverPool(uint256 amount, uint256 feeMultiplier, uint32 capacity) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) EnterProverPool(opts *bind.TransactOpts, amount *big.Int, feeMultiplier *big.Int, capacity uint32) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "enterProverPool", amount, feeMultiplier, capacity)
}

// EnterProverPool is a paid mutator transaction binding the contract method 0xd12037ff.
//
// Solidity: function enterProverPool(uint256 amount, uint256 feeMultiplier, uint32 capacity) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) EnterProverPool(amount *big.Int, feeMultiplier *big.Int, capacity uint32) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.EnterProverPool(&_TaikoL1ProverPool.TransactOpts, amount, feeMultiplier, capacity)
}

// EnterProverPool is a paid mutator transaction binding the contract method 0xd12037ff.
//
// Solidity: function enterProverPool(uint256 amount, uint256 feeMultiplier, uint32 capacity) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) EnterProverPool(amount *big.Int, feeMultiplier *big.Int, capacity uint32) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.EnterProverPool(&_TaikoL1ProverPool.TransactOpts, amount, feeMultiplier, capacity)
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

// Init is a paid mutator transaction binding the contract method 0xd2fd155e.
//
// Solidity: function init(address _addressManager, uint16 _maxPoolSize) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _maxPoolSize uint16) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "init", _addressManager, _maxPoolSize)
}

// Init is a paid mutator transaction binding the contract method 0xd2fd155e.
//
// Solidity: function init(address _addressManager, uint16 _maxPoolSize) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Init(_addressManager common.Address, _maxPoolSize uint16) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Init(&_TaikoL1ProverPool.TransactOpts, _addressManager, _maxPoolSize)
}

// Init is a paid mutator transaction binding the contract method 0xd2fd155e.
//
// Solidity: function init(address _addressManager, uint16 _maxPoolSize) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) Init(_addressManager common.Address, _maxPoolSize uint16) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Init(&_TaikoL1ProverPool.TransactOpts, _addressManager, _maxPoolSize)
}

// PickRandomProver is a paid mutator transaction binding the contract method 0xc9117290.
//
// Solidity: function pickRandomProver(uint256 randomNumber, uint256 blockId) returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) PickRandomProver(opts *bind.TransactOpts, randomNumber *big.Int, blockId *big.Int) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "pickRandomProver", randomNumber, blockId)
}

// PickRandomProver is a paid mutator transaction binding the contract method 0xc9117290.
//
// Solidity: function pickRandomProver(uint256 randomNumber, uint256 blockId) returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) PickRandomProver(randomNumber *big.Int, blockId *big.Int) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.PickRandomProver(&_TaikoL1ProverPool.TransactOpts, randomNumber, blockId)
}

// PickRandomProver is a paid mutator transaction binding the contract method 0xc9117290.
//
// Solidity: function pickRandomProver(uint256 randomNumber, uint256 blockId) returns(address)
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) PickRandomProver(randomNumber *big.Int, blockId *big.Int) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.PickRandomProver(&_TaikoL1ProverPool.TransactOpts, randomNumber, blockId)
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

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address prover) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) Slash(opts *bind.TransactOpts, prover common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "slash", prover)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address prover) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) Slash(prover common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Slash(&_TaikoL1ProverPool.TransactOpts, prover)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address prover) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) Slash(prover common.Address) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.Slash(&_TaikoL1ProverPool.TransactOpts, prover)
}

// StakeMoreTokens is a paid mutator transaction binding the contract method 0x9e8008d1.
//
// Solidity: function stakeMoreTokens(uint256 amount) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) StakeMoreTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "stakeMoreTokens", amount)
}

// StakeMoreTokens is a paid mutator transaction binding the contract method 0x9e8008d1.
//
// Solidity: function stakeMoreTokens(uint256 amount) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) StakeMoreTokens(amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.StakeMoreTokens(&_TaikoL1ProverPool.TransactOpts, amount)
}

// StakeMoreTokens is a paid mutator transaction binding the contract method 0x9e8008d1.
//
// Solidity: function stakeMoreTokens(uint256 amount) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) StakeMoreTokens(amount *big.Int) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.StakeMoreTokens(&_TaikoL1ProverPool.TransactOpts, amount)
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

// WithdrawRewards is a paid mutator transaction binding the contract method 0x45c69831.
//
// Solidity: function withdrawRewards(uint64 amount) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactor) WithdrawRewards(opts *bind.TransactOpts, amount uint64) (*types.Transaction, error) {
	return _TaikoL1ProverPool.contract.Transact(opts, "withdrawRewards", amount)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x45c69831.
//
// Solidity: function withdrawRewards(uint64 amount) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolSession) WithdrawRewards(amount uint64) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.WithdrawRewards(&_TaikoL1ProverPool.TransactOpts, amount)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x45c69831.
//
// Solidity: function withdrawRewards(uint64 amount) returns()
func (_TaikoL1ProverPool *TaikoL1ProverPoolTransactorSession) WithdrawRewards(amount uint64) (*types.Transaction, error) {
	return _TaikoL1ProverPool.Contract.WithdrawRewards(&_TaikoL1ProverPool.TransactOpts, amount)
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

// TaikoL1ProverPoolProverAdjustedCapacityIterator is returned from FilterProverAdjustedCapacity and is used to iterate over the raw logs and unpacked data for ProverAdjustedCapacity events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverAdjustedCapacityIterator struct {
	Event *TaikoL1ProverPoolProverAdjustedCapacity // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolProverAdjustedCapacityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolProverAdjustedCapacity)
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
		it.Event = new(TaikoL1ProverPoolProverAdjustedCapacity)
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
func (it *TaikoL1ProverPoolProverAdjustedCapacityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolProverAdjustedCapacityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolProverAdjustedCapacity represents a ProverAdjustedCapacity event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverAdjustedCapacity struct {
	Prover      common.Address
	OldCapacity uint32
	NewCapacity uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProverAdjustedCapacity is a free log retrieval operation binding the contract event 0xa426d28c428ee95eedf18d1c934f2069b32d38457a5f7f87127aeb8ea0729fb2.
//
// Solidity: event ProverAdjustedCapacity(address prover, uint32 oldCapacity, uint32 newCapacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterProverAdjustedCapacity(opts *bind.FilterOpts) (*TaikoL1ProverPoolProverAdjustedCapacityIterator, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "ProverAdjustedCapacity")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolProverAdjustedCapacityIterator{contract: _TaikoL1ProverPool.contract, event: "ProverAdjustedCapacity", logs: logs, sub: sub}, nil
}

// WatchProverAdjustedCapacity is a free log subscription operation binding the contract event 0xa426d28c428ee95eedf18d1c934f2069b32d38457a5f7f87127aeb8ea0729fb2.
//
// Solidity: event ProverAdjustedCapacity(address prover, uint32 oldCapacity, uint32 newCapacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchProverAdjustedCapacity(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolProverAdjustedCapacity) (event.Subscription, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "ProverAdjustedCapacity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolProverAdjustedCapacity)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverAdjustedCapacity", log); err != nil {
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

// ParseProverAdjustedCapacity is a log parse operation binding the contract event 0xa426d28c428ee95eedf18d1c934f2069b32d38457a5f7f87127aeb8ea0729fb2.
//
// Solidity: event ProverAdjustedCapacity(address prover, uint32 oldCapacity, uint32 newCapacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseProverAdjustedCapacity(log types.Log) (*TaikoL1ProverPoolProverAdjustedCapacity, error) {
	event := new(TaikoL1ProverPoolProverAdjustedCapacity)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverAdjustedCapacity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolProverAdjustedFeeMultiplierIterator is returned from FilterProverAdjustedFeeMultiplier and is used to iterate over the raw logs and unpacked data for ProverAdjustedFeeMultiplier events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverAdjustedFeeMultiplierIterator struct {
	Event *TaikoL1ProverPoolProverAdjustedFeeMultiplier // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolProverAdjustedFeeMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolProverAdjustedFeeMultiplier)
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
		it.Event = new(TaikoL1ProverPoolProverAdjustedFeeMultiplier)
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
func (it *TaikoL1ProverPoolProverAdjustedFeeMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolProverAdjustedFeeMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolProverAdjustedFeeMultiplier represents a ProverAdjustedFeeMultiplier event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverAdjustedFeeMultiplier struct {
	Prover           common.Address
	OldFeeMultiplier *big.Int
	NewFeeMultiplier *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProverAdjustedFeeMultiplier is a free log retrieval operation binding the contract event 0xe2ab56ce94312a84cb6aa7b090126fd821c5fad200bf22a7c55f675ec668db9e.
//
// Solidity: event ProverAdjustedFeeMultiplier(address prover, uint256 oldFeeMultiplier, uint256 newFeeMultiplier)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterProverAdjustedFeeMultiplier(opts *bind.FilterOpts) (*TaikoL1ProverPoolProverAdjustedFeeMultiplierIterator, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "ProverAdjustedFeeMultiplier")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolProverAdjustedFeeMultiplierIterator{contract: _TaikoL1ProverPool.contract, event: "ProverAdjustedFeeMultiplier", logs: logs, sub: sub}, nil
}

// WatchProverAdjustedFeeMultiplier is a free log subscription operation binding the contract event 0xe2ab56ce94312a84cb6aa7b090126fd821c5fad200bf22a7c55f675ec668db9e.
//
// Solidity: event ProverAdjustedFeeMultiplier(address prover, uint256 oldFeeMultiplier, uint256 newFeeMultiplier)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchProverAdjustedFeeMultiplier(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolProverAdjustedFeeMultiplier) (event.Subscription, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "ProverAdjustedFeeMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolProverAdjustedFeeMultiplier)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverAdjustedFeeMultiplier", log); err != nil {
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

// ParseProverAdjustedFeeMultiplier is a log parse operation binding the contract event 0xe2ab56ce94312a84cb6aa7b090126fd821c5fad200bf22a7c55f675ec668db9e.
//
// Solidity: event ProverAdjustedFeeMultiplier(address prover, uint256 oldFeeMultiplier, uint256 newFeeMultiplier)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseProverAdjustedFeeMultiplier(log types.Log) (*TaikoL1ProverPoolProverAdjustedFeeMultiplier, error) {
	event := new(TaikoL1ProverPoolProverAdjustedFeeMultiplier)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverAdjustedFeeMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolProverEnteredIterator is returned from FilterProverEntered and is used to iterate over the raw logs and unpacked data for ProverEntered events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverEnteredIterator struct {
	Event *TaikoL1ProverPoolProverEntered // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolProverEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolProverEntered)
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
		it.Event = new(TaikoL1ProverPoolProverEntered)
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
func (it *TaikoL1ProverPoolProverEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolProverEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolProverEntered represents a ProverEntered event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverEntered struct {
	Prover        common.Address
	Amount        *big.Int
	FeeMultiplier *big.Int
	Capacity      uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProverEntered is a free log retrieval operation binding the contract event 0x1b14a9db786d0b879f833a7bd117d424ee00758960ef3c2e812b2abf0a51f892.
//
// Solidity: event ProverEntered(address prover, uint256 amount, uint256 feeMultiplier, uint64 capacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterProverEntered(opts *bind.FilterOpts) (*TaikoL1ProverPoolProverEnteredIterator, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "ProverEntered")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolProverEnteredIterator{contract: _TaikoL1ProverPool.contract, event: "ProverEntered", logs: logs, sub: sub}, nil
}

// WatchProverEntered is a free log subscription operation binding the contract event 0x1b14a9db786d0b879f833a7bd117d424ee00758960ef3c2e812b2abf0a51f892.
//
// Solidity: event ProverEntered(address prover, uint256 amount, uint256 feeMultiplier, uint64 capacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchProverEntered(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolProverEntered) (event.Subscription, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "ProverEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolProverEntered)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverEntered", log); err != nil {
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

// ParseProverEntered is a log parse operation binding the contract event 0x1b14a9db786d0b879f833a7bd117d424ee00758960ef3c2e812b2abf0a51f892.
//
// Solidity: event ProverEntered(address prover, uint256 amount, uint256 feeMultiplier, uint64 capacity)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseProverEntered(log types.Log) (*TaikoL1ProverPoolProverEntered, error) {
	event := new(TaikoL1ProverPoolProverEntered)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolProverExitedIterator is returned from FilterProverExited and is used to iterate over the raw logs and unpacked data for ProverExited events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverExitedIterator struct {
	Event *TaikoL1ProverPoolProverExited // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolProverExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolProverExited)
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
		it.Event = new(TaikoL1ProverPoolProverExited)
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
func (it *TaikoL1ProverPoolProverExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolProverExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolProverExited represents a ProverExited event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverExited struct {
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProverExited is a free log retrieval operation binding the contract event 0x2815fc337451500d2c4aa22628a7584582edde5bf78b2ba9caa6efbd6cce4a8e.
//
// Solidity: event ProverExited(address prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterProverExited(opts *bind.FilterOpts) (*TaikoL1ProverPoolProverExitedIterator, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "ProverExited")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolProverExitedIterator{contract: _TaikoL1ProverPool.contract, event: "ProverExited", logs: logs, sub: sub}, nil
}

// WatchProverExited is a free log subscription operation binding the contract event 0x2815fc337451500d2c4aa22628a7584582edde5bf78b2ba9caa6efbd6cce4a8e.
//
// Solidity: event ProverExited(address prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchProverExited(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolProverExited) (event.Subscription, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "ProverExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolProverExited)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverExited", log); err != nil {
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

// ParseProverExited is a log parse operation binding the contract event 0x2815fc337451500d2c4aa22628a7584582edde5bf78b2ba9caa6efbd6cce4a8e.
//
// Solidity: event ProverExited(address prover)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseProverExited(log types.Log) (*TaikoL1ProverPoolProverExited, error) {
	event := new(TaikoL1ProverPoolProverExited)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolProverSlashedIterator is returned from FilterProverSlashed and is used to iterate over the raw logs and unpacked data for ProverSlashed events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverSlashedIterator struct {
	Event *TaikoL1ProverPoolProverSlashed // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolProverSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolProverSlashed)
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
		it.Event = new(TaikoL1ProverPoolProverSlashed)
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
func (it *TaikoL1ProverPoolProverSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolProverSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolProverSlashed represents a ProverSlashed event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverSlashed struct {
	Prover common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProverSlashed is a free log retrieval operation binding the contract event 0x664b84a0f5b173c1d62371e87f48268f943748c0fe5805d64ebfab28af48e17b.
//
// Solidity: event ProverSlashed(address prover, uint256 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterProverSlashed(opts *bind.FilterOpts) (*TaikoL1ProverPoolProverSlashedIterator, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "ProverSlashed")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolProverSlashedIterator{contract: _TaikoL1ProverPool.contract, event: "ProverSlashed", logs: logs, sub: sub}, nil
}

// WatchProverSlashed is a free log subscription operation binding the contract event 0x664b84a0f5b173c1d62371e87f48268f943748c0fe5805d64ebfab28af48e17b.
//
// Solidity: event ProverSlashed(address prover, uint256 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchProverSlashed(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolProverSlashed) (event.Subscription, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "ProverSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolProverSlashed)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverSlashed", log); err != nil {
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

// ParseProverSlashed is a log parse operation binding the contract event 0x664b84a0f5b173c1d62371e87f48268f943748c0fe5805d64ebfab28af48e17b.
//
// Solidity: event ProverSlashed(address prover, uint256 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseProverSlashed(log types.Log) (*TaikoL1ProverPoolProverSlashed, error) {
	event := new(TaikoL1ProverPoolProverSlashed)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolProverStakedMoreTokensIterator is returned from FilterProverStakedMoreTokens and is used to iterate over the raw logs and unpacked data for ProverStakedMoreTokens events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverStakedMoreTokensIterator struct {
	Event *TaikoL1ProverPoolProverStakedMoreTokens // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolProverStakedMoreTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolProverStakedMoreTokens)
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
		it.Event = new(TaikoL1ProverPoolProverStakedMoreTokens)
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
func (it *TaikoL1ProverPoolProverStakedMoreTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolProverStakedMoreTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolProverStakedMoreTokens represents a ProverStakedMoreTokens event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverStakedMoreTokens struct {
	Prover      common.Address
	Amount      *big.Int
	TotalStaked *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProverStakedMoreTokens is a free log retrieval operation binding the contract event 0xede799219aee9e2d2e6ce762bb15b4cdfb387b1618280f67b3404dc12463c382.
//
// Solidity: event ProverStakedMoreTokens(address prover, uint256 amount, uint256 totalStaked)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterProverStakedMoreTokens(opts *bind.FilterOpts) (*TaikoL1ProverPoolProverStakedMoreTokensIterator, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "ProverStakedMoreTokens")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolProverStakedMoreTokensIterator{contract: _TaikoL1ProverPool.contract, event: "ProverStakedMoreTokens", logs: logs, sub: sub}, nil
}

// WatchProverStakedMoreTokens is a free log subscription operation binding the contract event 0xede799219aee9e2d2e6ce762bb15b4cdfb387b1618280f67b3404dc12463c382.
//
// Solidity: event ProverStakedMoreTokens(address prover, uint256 amount, uint256 totalStaked)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchProverStakedMoreTokens(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolProverStakedMoreTokens) (event.Subscription, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "ProverStakedMoreTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolProverStakedMoreTokens)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverStakedMoreTokens", log); err != nil {
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

// ParseProverStakedMoreTokens is a log parse operation binding the contract event 0xede799219aee9e2d2e6ce762bb15b4cdfb387b1618280f67b3404dc12463c382.
//
// Solidity: event ProverStakedMoreTokens(address prover, uint256 amount, uint256 totalStaked)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseProverStakedMoreTokens(log types.Log) (*TaikoL1ProverPoolProverStakedMoreTokens, error) {
	event := new(TaikoL1ProverPoolProverStakedMoreTokens)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverStakedMoreTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ProverPoolProverWithdrawAwardsIterator is returned from FilterProverWithdrawAwards and is used to iterate over the raw logs and unpacked data for ProverWithdrawAwards events raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverWithdrawAwardsIterator struct {
	Event *TaikoL1ProverPoolProverWithdrawAwards // Event containing the contract specifics and raw log

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
func (it *TaikoL1ProverPoolProverWithdrawAwardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ProverPoolProverWithdrawAwards)
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
		it.Event = new(TaikoL1ProverPoolProverWithdrawAwards)
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
func (it *TaikoL1ProverPoolProverWithdrawAwardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ProverPoolProverWithdrawAwardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ProverPoolProverWithdrawAwards represents a ProverWithdrawAwards event raised by the TaikoL1ProverPool contract.
type TaikoL1ProverPoolProverWithdrawAwards struct {
	Prover common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProverWithdrawAwards is a free log retrieval operation binding the contract event 0x78e57d846371816e02bd37310e6cc83bb7a979a78e86a5f47214fa6f90518be4.
//
// Solidity: event ProverWithdrawAwards(address prover, uint256 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) FilterProverWithdrawAwards(opts *bind.FilterOpts) (*TaikoL1ProverPoolProverWithdrawAwardsIterator, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.FilterLogs(opts, "ProverWithdrawAwards")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ProverPoolProverWithdrawAwardsIterator{contract: _TaikoL1ProverPool.contract, event: "ProverWithdrawAwards", logs: logs, sub: sub}, nil
}

// WatchProverWithdrawAwards is a free log subscription operation binding the contract event 0x78e57d846371816e02bd37310e6cc83bb7a979a78e86a5f47214fa6f90518be4.
//
// Solidity: event ProverWithdrawAwards(address prover, uint256 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) WatchProverWithdrawAwards(opts *bind.WatchOpts, sink chan<- *TaikoL1ProverPoolProverWithdrawAwards) (event.Subscription, error) {

	logs, sub, err := _TaikoL1ProverPool.contract.WatchLogs(opts, "ProverWithdrawAwards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ProverPoolProverWithdrawAwards)
				if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverWithdrawAwards", log); err != nil {
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

// ParseProverWithdrawAwards is a log parse operation binding the contract event 0x78e57d846371816e02bd37310e6cc83bb7a979a78e86a5f47214fa6f90518be4.
//
// Solidity: event ProverWithdrawAwards(address prover, uint256 amount)
func (_TaikoL1ProverPool *TaikoL1ProverPoolFilterer) ParseProverWithdrawAwards(log types.Log) (*TaikoL1ProverPoolProverWithdrawAwards, error) {
	event := new(TaikoL1ProverPoolProverWithdrawAwards)
	if err := _TaikoL1ProverPool.contract.UnpackLog(event, "ProverWithdrawAwards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
