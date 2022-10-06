// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/taikochain/taiko-client"
	"github.com/taikochain/taiko-client/accounts/abi"
	"github.com/taikochain/taiko-client/accounts/abi/bind"
	"github.com/taikochain/taiko-client/common"
	"github.com/taikochain/taiko-client/core/types"
	"github.com/taikochain/taiko-client/event"
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

// LibInvalidTxListMetaData contains all meta data concerning the LibInvalidTxList contract.
var LibInvalidTxListMetaData = &bind.MetaData{
	ABI: "[]",
}

// LibInvalidTxListABI is the input ABI used to generate the binding from.
// Deprecated: Use LibInvalidTxListMetaData.ABI instead.
var LibInvalidTxListABI = LibInvalidTxListMetaData.ABI

// LibInvalidTxList is an auto generated Go binding around an Ethereum contract.
type LibInvalidTxList struct {
	LibInvalidTxListCaller     // Read-only binding to the contract
	LibInvalidTxListTransactor // Write-only binding to the contract
	LibInvalidTxListFilterer   // Log filterer for contract events
}

// LibInvalidTxListCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibInvalidTxListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibInvalidTxListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibInvalidTxListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibInvalidTxListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibInvalidTxListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibInvalidTxListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibInvalidTxListSession struct {
	Contract     *LibInvalidTxList // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibInvalidTxListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibInvalidTxListCallerSession struct {
	Contract *LibInvalidTxListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LibInvalidTxListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibInvalidTxListTransactorSession struct {
	Contract     *LibInvalidTxListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LibInvalidTxListRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibInvalidTxListRaw struct {
	Contract *LibInvalidTxList // Generic contract binding to access the raw methods on
}

// LibInvalidTxListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibInvalidTxListCallerRaw struct {
	Contract *LibInvalidTxListCaller // Generic read-only contract binding to access the raw methods on
}

// LibInvalidTxListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibInvalidTxListTransactorRaw struct {
	Contract *LibInvalidTxListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibInvalidTxList creates a new instance of LibInvalidTxList, bound to a specific deployed contract.
func NewLibInvalidTxList(address common.Address, backend bind.ContractBackend) (*LibInvalidTxList, error) {
	contract, err := bindLibInvalidTxList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibInvalidTxList{LibInvalidTxListCaller: LibInvalidTxListCaller{contract: contract}, LibInvalidTxListTransactor: LibInvalidTxListTransactor{contract: contract}, LibInvalidTxListFilterer: LibInvalidTxListFilterer{contract: contract}}, nil
}

// NewLibInvalidTxListCaller creates a new read-only instance of LibInvalidTxList, bound to a specific deployed contract.
func NewLibInvalidTxListCaller(address common.Address, caller bind.ContractCaller) (*LibInvalidTxListCaller, error) {
	contract, err := bindLibInvalidTxList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibInvalidTxListCaller{contract: contract}, nil
}

// NewLibInvalidTxListTransactor creates a new write-only instance of LibInvalidTxList, bound to a specific deployed contract.
func NewLibInvalidTxListTransactor(address common.Address, transactor bind.ContractTransactor) (*LibInvalidTxListTransactor, error) {
	contract, err := bindLibInvalidTxList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibInvalidTxListTransactor{contract: contract}, nil
}

// NewLibInvalidTxListFilterer creates a new log filterer instance of LibInvalidTxList, bound to a specific deployed contract.
func NewLibInvalidTxListFilterer(address common.Address, filterer bind.ContractFilterer) (*LibInvalidTxListFilterer, error) {
	contract, err := bindLibInvalidTxList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibInvalidTxListFilterer{contract: contract}, nil
}

// bindLibInvalidTxList binds a generic wrapper to an already deployed contract.
func bindLibInvalidTxList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibInvalidTxListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibInvalidTxList *LibInvalidTxListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibInvalidTxList.Contract.LibInvalidTxListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibInvalidTxList *LibInvalidTxListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibInvalidTxList.Contract.LibInvalidTxListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibInvalidTxList *LibInvalidTxListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibInvalidTxList.Contract.LibInvalidTxListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibInvalidTxList *LibInvalidTxListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibInvalidTxList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibInvalidTxList *LibInvalidTxListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibInvalidTxList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibInvalidTxList *LibInvalidTxListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibInvalidTxList.Contract.contract.Transact(opts, method, params...)
}
