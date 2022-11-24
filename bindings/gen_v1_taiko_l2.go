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

// V1TaikoL2ClientMetaData contains all meta data concerning the V1TaikoL2Client contract.
var V1TaikoL2ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"}],\"name\":\"BlockInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcHash\",\"type\":\"bytes32\"}],\"name\":\"HeaderSynced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l1Height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"}],\"name\":\"anchor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConstants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestSyncedHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getSyncedHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txList\",\"type\":\"bytes\"},{\"internalType\":\"enumLibInvalidTxList.Reason\",\"name\":\"hint\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"txIdx\",\"type\":\"uint256\"}],\"name\":\"invalidateBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestSyncedHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// V1TaikoL2ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use V1TaikoL2ClientMetaData.ABI instead.
var V1TaikoL2ClientABI = V1TaikoL2ClientMetaData.ABI

// V1TaikoL2Client is an auto generated Go binding around an Ethereum contract.
type V1TaikoL2Client struct {
	V1TaikoL2ClientCaller     // Read-only binding to the contract
	V1TaikoL2ClientTransactor // Write-only binding to the contract
	V1TaikoL2ClientFilterer   // Log filterer for contract events
}

// V1TaikoL2ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type V1TaikoL2ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1TaikoL2ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V1TaikoL2ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1TaikoL2ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V1TaikoL2ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1TaikoL2ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V1TaikoL2ClientSession struct {
	Contract     *V1TaikoL2Client  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V1TaikoL2ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V1TaikoL2ClientCallerSession struct {
	Contract *V1TaikoL2ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// V1TaikoL2ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V1TaikoL2ClientTransactorSession struct {
	Contract     *V1TaikoL2ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// V1TaikoL2ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type V1TaikoL2ClientRaw struct {
	Contract *V1TaikoL2Client // Generic contract binding to access the raw methods on
}

// V1TaikoL2ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V1TaikoL2ClientCallerRaw struct {
	Contract *V1TaikoL2ClientCaller // Generic read-only contract binding to access the raw methods on
}

// V1TaikoL2ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V1TaikoL2ClientTransactorRaw struct {
	Contract *V1TaikoL2ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV1TaikoL2Client creates a new instance of V1TaikoL2Client, bound to a specific deployed contract.
func NewV1TaikoL2Client(address common.Address, backend bind.ContractBackend) (*V1TaikoL2Client, error) {
	contract, err := bindV1TaikoL2Client(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V1TaikoL2Client{V1TaikoL2ClientCaller: V1TaikoL2ClientCaller{contract: contract}, V1TaikoL2ClientTransactor: V1TaikoL2ClientTransactor{contract: contract}, V1TaikoL2ClientFilterer: V1TaikoL2ClientFilterer{contract: contract}}, nil
}

// NewV1TaikoL2ClientCaller creates a new read-only instance of V1TaikoL2Client, bound to a specific deployed contract.
func NewV1TaikoL2ClientCaller(address common.Address, caller bind.ContractCaller) (*V1TaikoL2ClientCaller, error) {
	contract, err := bindV1TaikoL2Client(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V1TaikoL2ClientCaller{contract: contract}, nil
}

// NewV1TaikoL2ClientTransactor creates a new write-only instance of V1TaikoL2Client, bound to a specific deployed contract.
func NewV1TaikoL2ClientTransactor(address common.Address, transactor bind.ContractTransactor) (*V1TaikoL2ClientTransactor, error) {
	contract, err := bindV1TaikoL2Client(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V1TaikoL2ClientTransactor{contract: contract}, nil
}

// NewV1TaikoL2ClientFilterer creates a new log filterer instance of V1TaikoL2Client, bound to a specific deployed contract.
func NewV1TaikoL2ClientFilterer(address common.Address, filterer bind.ContractFilterer) (*V1TaikoL2ClientFilterer, error) {
	contract, err := bindV1TaikoL2Client(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V1TaikoL2ClientFilterer{contract: contract}, nil
}

// bindV1TaikoL2Client binds a generic wrapper to an already deployed contract.
func bindV1TaikoL2Client(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(V1TaikoL2ClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1TaikoL2Client *V1TaikoL2ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1TaikoL2Client.Contract.V1TaikoL2ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1TaikoL2Client *V1TaikoL2ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1TaikoL2Client.Contract.V1TaikoL2ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1TaikoL2Client *V1TaikoL2ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1TaikoL2Client.Contract.V1TaikoL2ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1TaikoL2Client *V1TaikoL2ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1TaikoL2Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1TaikoL2Client *V1TaikoL2ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1TaikoL2Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1TaikoL2Client *V1TaikoL2ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1TaikoL2Client.Contract.contract.Transact(opts, method, params...)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_V1TaikoL2Client *V1TaikoL2ClientCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1TaikoL2Client.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_V1TaikoL2Client *V1TaikoL2ClientSession) AddressManager() (common.Address, error) {
	return _V1TaikoL2Client.Contract.AddressManager(&_V1TaikoL2Client.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_V1TaikoL2Client *V1TaikoL2ClientCallerSession) AddressManager() (common.Address, error) {
	return _V1TaikoL2Client.Contract.AddressManager(&_V1TaikoL2Client.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCaller) GetBlockHash(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _V1TaikoL2Client.contract.Call(opts, &out, "getBlockHash", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientSession) GetBlockHash(number *big.Int) ([32]byte, error) {
	return _V1TaikoL2Client.Contract.GetBlockHash(&_V1TaikoL2Client.CallOpts, number)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCallerSession) GetBlockHash(number *big.Int) ([32]byte, error) {
	return _V1TaikoL2Client.Contract.GetBlockHash(&_V1TaikoL2Client.CallOpts, number)
}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256, bytes32, uint256, uint256, uint256, bytes4, bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCaller) GetConstants(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, [32]byte, *big.Int, *big.Int, *big.Int, [4]byte, [32]byte, error) {
	var out []interface{}
	err := _V1TaikoL2Client.contract.Call(opts, &out, "getConstants")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([32]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([4]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)
	out9 := *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	out10 := *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	out11 := *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	out12 := *abi.ConvertType(out[12], new([4]byte)).(*[4]byte)
	out13 := *abi.ConvertType(out[13], new([32]byte)).(*[32]byte)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, out11, out12, out13, err

}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256, bytes32, uint256, uint256, uint256, bytes4, bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientSession) GetConstants() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, [32]byte, *big.Int, *big.Int, *big.Int, [4]byte, [32]byte, error) {
	return _V1TaikoL2Client.Contract.GetConstants(&_V1TaikoL2Client.CallOpts)
}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256, bytes32, uint256, uint256, uint256, bytes4, bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCallerSession) GetConstants() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, [32]byte, *big.Int, *big.Int, *big.Int, [4]byte, [32]byte, error) {
	return _V1TaikoL2Client.Contract.GetConstants(&_V1TaikoL2Client.CallOpts)
}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCaller) GetLatestSyncedHeader(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V1TaikoL2Client.contract.Call(opts, &out, "getLatestSyncedHeader")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientSession) GetLatestSyncedHeader() ([32]byte, error) {
	return _V1TaikoL2Client.Contract.GetLatestSyncedHeader(&_V1TaikoL2Client.CallOpts)
}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCallerSession) GetLatestSyncedHeader() ([32]byte, error) {
	return _V1TaikoL2Client.Contract.GetLatestSyncedHeader(&_V1TaikoL2Client.CallOpts)
}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCaller) GetSyncedHeader(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _V1TaikoL2Client.contract.Call(opts, &out, "getSyncedHeader", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientSession) GetSyncedHeader(number *big.Int) ([32]byte, error) {
	return _V1TaikoL2Client.Contract.GetSyncedHeader(&_V1TaikoL2Client.CallOpts, number)
}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCallerSession) GetSyncedHeader(number *big.Int) ([32]byte, error) {
	return _V1TaikoL2Client.Contract.GetSyncedHeader(&_V1TaikoL2Client.CallOpts, number)
}

// LatestSyncedHeader is a free data retrieval call binding the contract method 0xc8d772ab.
//
// Solidity: function latestSyncedHeader() view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCaller) LatestSyncedHeader(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V1TaikoL2Client.contract.Call(opts, &out, "latestSyncedHeader")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestSyncedHeader is a free data retrieval call binding the contract method 0xc8d772ab.
//
// Solidity: function latestSyncedHeader() view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientSession) LatestSyncedHeader() ([32]byte, error) {
	return _V1TaikoL2Client.Contract.LatestSyncedHeader(&_V1TaikoL2Client.CallOpts)
}

// LatestSyncedHeader is a free data retrieval call binding the contract method 0xc8d772ab.
//
// Solidity: function latestSyncedHeader() view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCallerSession) LatestSyncedHeader() ([32]byte, error) {
	return _V1TaikoL2Client.Contract.LatestSyncedHeader(&_V1TaikoL2Client.CallOpts)
}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCaller) PublicInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V1TaikoL2Client.contract.Call(opts, &out, "publicInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientSession) PublicInputHash() ([32]byte, error) {
	return _V1TaikoL2Client.Contract.PublicInputHash(&_V1TaikoL2Client.CallOpts)
}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_V1TaikoL2Client *V1TaikoL2ClientCallerSession) PublicInputHash() ([32]byte, error) {
	return _V1TaikoL2Client.Contract.PublicInputHash(&_V1TaikoL2Client.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string name) view returns(address)
func (_V1TaikoL2Client *V1TaikoL2ClientCaller) Resolve(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _V1TaikoL2Client.contract.Call(opts, &out, "resolve", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string name) view returns(address)
func (_V1TaikoL2Client *V1TaikoL2ClientSession) Resolve(name string) (common.Address, error) {
	return _V1TaikoL2Client.Contract.Resolve(&_V1TaikoL2Client.CallOpts, name)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string name) view returns(address)
func (_V1TaikoL2Client *V1TaikoL2ClientCallerSession) Resolve(name string) (common.Address, error) {
	return _V1TaikoL2Client.Contract.Resolve(&_V1TaikoL2Client.CallOpts, name)
}

// Resolve0 is a free data retrieval call binding the contract method 0xf16c7934.
//
// Solidity: function resolve(uint256 chainId, string name) view returns(address)
func (_V1TaikoL2Client *V1TaikoL2ClientCaller) Resolve0(opts *bind.CallOpts, chainId *big.Int, name string) (common.Address, error) {
	var out []interface{}
	err := _V1TaikoL2Client.contract.Call(opts, &out, "resolve0", chainId, name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xf16c7934.
//
// Solidity: function resolve(uint256 chainId, string name) view returns(address)
func (_V1TaikoL2Client *V1TaikoL2ClientSession) Resolve0(chainId *big.Int, name string) (common.Address, error) {
	return _V1TaikoL2Client.Contract.Resolve0(&_V1TaikoL2Client.CallOpts, chainId, name)
}

// Resolve0 is a free data retrieval call binding the contract method 0xf16c7934.
//
// Solidity: function resolve(uint256 chainId, string name) view returns(address)
func (_V1TaikoL2Client *V1TaikoL2ClientCallerSession) Resolve0(chainId *big.Int, name string) (common.Address, error) {
	return _V1TaikoL2Client.Contract.Resolve0(&_V1TaikoL2Client.CallOpts, chainId, name)
}

// Anchor is a paid mutator transaction binding the contract method 0xa0ca2d08.
//
// Solidity: function anchor(uint256 l1Height, bytes32 l1Hash) returns()
func (_V1TaikoL2Client *V1TaikoL2ClientTransactor) Anchor(opts *bind.TransactOpts, l1Height *big.Int, l1Hash [32]byte) (*types.Transaction, error) {
	return _V1TaikoL2Client.contract.Transact(opts, "anchor", l1Height, l1Hash)
}

// Anchor is a paid mutator transaction binding the contract method 0xa0ca2d08.
//
// Solidity: function anchor(uint256 l1Height, bytes32 l1Hash) returns()
func (_V1TaikoL2Client *V1TaikoL2ClientSession) Anchor(l1Height *big.Int, l1Hash [32]byte) (*types.Transaction, error) {
	return _V1TaikoL2Client.Contract.Anchor(&_V1TaikoL2Client.TransactOpts, l1Height, l1Hash)
}

// Anchor is a paid mutator transaction binding the contract method 0xa0ca2d08.
//
// Solidity: function anchor(uint256 l1Height, bytes32 l1Hash) returns()
func (_V1TaikoL2Client *V1TaikoL2ClientTransactorSession) Anchor(l1Height *big.Int, l1Hash [32]byte) (*types.Transaction, error) {
	return _V1TaikoL2Client.Contract.Anchor(&_V1TaikoL2Client.TransactOpts, l1Height, l1Hash)
}

// InvalidateBlock is a paid mutator transaction binding the contract method 0x975e09a0.
//
// Solidity: function invalidateBlock(bytes txList, uint8 hint, uint256 txIdx) returns()
func (_V1TaikoL2Client *V1TaikoL2ClientTransactor) InvalidateBlock(opts *bind.TransactOpts, txList []byte, hint uint8, txIdx *big.Int) (*types.Transaction, error) {
	return _V1TaikoL2Client.contract.Transact(opts, "invalidateBlock", txList, hint, txIdx)
}

// InvalidateBlock is a paid mutator transaction binding the contract method 0x975e09a0.
//
// Solidity: function invalidateBlock(bytes txList, uint8 hint, uint256 txIdx) returns()
func (_V1TaikoL2Client *V1TaikoL2ClientSession) InvalidateBlock(txList []byte, hint uint8, txIdx *big.Int) (*types.Transaction, error) {
	return _V1TaikoL2Client.Contract.InvalidateBlock(&_V1TaikoL2Client.TransactOpts, txList, hint, txIdx)
}

// InvalidateBlock is a paid mutator transaction binding the contract method 0x975e09a0.
//
// Solidity: function invalidateBlock(bytes txList, uint8 hint, uint256 txIdx) returns()
func (_V1TaikoL2Client *V1TaikoL2ClientTransactorSession) InvalidateBlock(txList []byte, hint uint8, txIdx *big.Int) (*types.Transaction, error) {
	return _V1TaikoL2Client.Contract.InvalidateBlock(&_V1TaikoL2Client.TransactOpts, txList, hint, txIdx)
}

// V1TaikoL2ClientBlockInvalidatedIterator is returned from FilterBlockInvalidated and is used to iterate over the raw logs and unpacked data for BlockInvalidated events raised by the V1TaikoL2Client contract.
type V1TaikoL2ClientBlockInvalidatedIterator struct {
	Event *V1TaikoL2ClientBlockInvalidated // Event containing the contract specifics and raw log

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
func (it *V1TaikoL2ClientBlockInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1TaikoL2ClientBlockInvalidated)
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
		it.Event = new(V1TaikoL2ClientBlockInvalidated)
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
func (it *V1TaikoL2ClientBlockInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1TaikoL2ClientBlockInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1TaikoL2ClientBlockInvalidated represents a BlockInvalidated event raised by the V1TaikoL2Client contract.
type V1TaikoL2ClientBlockInvalidated struct {
	TxListHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockInvalidated is a free log retrieval operation binding the contract event 0x64b299ff9f8ba674288abb53380419048a4271dda03b837ecba6b40e6ddea4a2.
//
// Solidity: event BlockInvalidated(bytes32 indexed txListHash)
func (_V1TaikoL2Client *V1TaikoL2ClientFilterer) FilterBlockInvalidated(opts *bind.FilterOpts, txListHash [][32]byte) (*V1TaikoL2ClientBlockInvalidatedIterator, error) {

	var txListHashRule []interface{}
	for _, txListHashItem := range txListHash {
		txListHashRule = append(txListHashRule, txListHashItem)
	}

	logs, sub, err := _V1TaikoL2Client.contract.FilterLogs(opts, "BlockInvalidated", txListHashRule)
	if err != nil {
		return nil, err
	}
	return &V1TaikoL2ClientBlockInvalidatedIterator{contract: _V1TaikoL2Client.contract, event: "BlockInvalidated", logs: logs, sub: sub}, nil
}

// WatchBlockInvalidated is a free log subscription operation binding the contract event 0x64b299ff9f8ba674288abb53380419048a4271dda03b837ecba6b40e6ddea4a2.
//
// Solidity: event BlockInvalidated(bytes32 indexed txListHash)
func (_V1TaikoL2Client *V1TaikoL2ClientFilterer) WatchBlockInvalidated(opts *bind.WatchOpts, sink chan<- *V1TaikoL2ClientBlockInvalidated, txListHash [][32]byte) (event.Subscription, error) {

	var txListHashRule []interface{}
	for _, txListHashItem := range txListHash {
		txListHashRule = append(txListHashRule, txListHashItem)
	}

	logs, sub, err := _V1TaikoL2Client.contract.WatchLogs(opts, "BlockInvalidated", txListHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1TaikoL2ClientBlockInvalidated)
				if err := _V1TaikoL2Client.contract.UnpackLog(event, "BlockInvalidated", log); err != nil {
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

// ParseBlockInvalidated is a log parse operation binding the contract event 0x64b299ff9f8ba674288abb53380419048a4271dda03b837ecba6b40e6ddea4a2.
//
// Solidity: event BlockInvalidated(bytes32 indexed txListHash)
func (_V1TaikoL2Client *V1TaikoL2ClientFilterer) ParseBlockInvalidated(log types.Log) (*V1TaikoL2ClientBlockInvalidated, error) {
	event := new(V1TaikoL2ClientBlockInvalidated)
	if err := _V1TaikoL2Client.contract.UnpackLog(event, "BlockInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1TaikoL2ClientHeaderSyncedIterator is returned from FilterHeaderSynced and is used to iterate over the raw logs and unpacked data for HeaderSynced events raised by the V1TaikoL2Client contract.
type V1TaikoL2ClientHeaderSyncedIterator struct {
	Event *V1TaikoL2ClientHeaderSynced // Event containing the contract specifics and raw log

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
func (it *V1TaikoL2ClientHeaderSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1TaikoL2ClientHeaderSynced)
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
		it.Event = new(V1TaikoL2ClientHeaderSynced)
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
func (it *V1TaikoL2ClientHeaderSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1TaikoL2ClientHeaderSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1TaikoL2ClientHeaderSynced represents a HeaderSynced event raised by the V1TaikoL2Client contract.
type V1TaikoL2ClientHeaderSynced struct {
	Height    *big.Int
	SrcHeight *big.Int
	SrcHash   [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHeaderSynced is a free log retrieval operation binding the contract event 0x930c750845026c7bb04c0e3d9111d512b4c86981713c4944a35a10a4a7a854f3.
//
// Solidity: event HeaderSynced(uint256 indexed height, uint256 indexed srcHeight, bytes32 srcHash)
func (_V1TaikoL2Client *V1TaikoL2ClientFilterer) FilterHeaderSynced(opts *bind.FilterOpts, height []*big.Int, srcHeight []*big.Int) (*V1TaikoL2ClientHeaderSyncedIterator, error) {

	var heightRule []interface{}
	for _, heightItem := range height {
		heightRule = append(heightRule, heightItem)
	}
	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _V1TaikoL2Client.contract.FilterLogs(opts, "HeaderSynced", heightRule, srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &V1TaikoL2ClientHeaderSyncedIterator{contract: _V1TaikoL2Client.contract, event: "HeaderSynced", logs: logs, sub: sub}, nil
}

// WatchHeaderSynced is a free log subscription operation binding the contract event 0x930c750845026c7bb04c0e3d9111d512b4c86981713c4944a35a10a4a7a854f3.
//
// Solidity: event HeaderSynced(uint256 indexed height, uint256 indexed srcHeight, bytes32 srcHash)
func (_V1TaikoL2Client *V1TaikoL2ClientFilterer) WatchHeaderSynced(opts *bind.WatchOpts, sink chan<- *V1TaikoL2ClientHeaderSynced, height []*big.Int, srcHeight []*big.Int) (event.Subscription, error) {

	var heightRule []interface{}
	for _, heightItem := range height {
		heightRule = append(heightRule, heightItem)
	}
	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _V1TaikoL2Client.contract.WatchLogs(opts, "HeaderSynced", heightRule, srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1TaikoL2ClientHeaderSynced)
				if err := _V1TaikoL2Client.contract.UnpackLog(event, "HeaderSynced", log); err != nil {
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

// ParseHeaderSynced is a log parse operation binding the contract event 0x930c750845026c7bb04c0e3d9111d512b4c86981713c4944a35a10a4a7a854f3.
//
// Solidity: event HeaderSynced(uint256 indexed height, uint256 indexed srcHeight, bytes32 srcHash)
func (_V1TaikoL2Client *V1TaikoL2ClientFilterer) ParseHeaderSynced(log types.Log) (*V1TaikoL2ClientHeaderSynced, error) {
	event := new(V1TaikoL2ClientHeaderSynced)
	if err := _V1TaikoL2Client.contract.UnpackLog(event, "HeaderSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
