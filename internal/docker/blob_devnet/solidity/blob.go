// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity

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

// BallotTestMetaData contains all meta data concerning the BallotTest contract.
var BallotTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L1_BLOB_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"storeBlobHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"11762da2": "storeBlobHash()",
	},
	Bin: "0x608060405234801561000f575f80fd5b5060ac8061001c5f395ff3fe6080604052348015600e575f80fd5b50600436106026575f3560e01c806311762da214602a575b5f80fd5b60306032565b005b5f495f819055505f801b5f54036074576040517f9e7e2ddd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b56fea2646970667358221220a02f78c0536b7c1e0f20718b4ef6dd26efec58dc427ce9b93baaf44f316463ac64736f6c63430008180033",
}

// BallotTestABI is the input ABI used to generate the binding from.
// Deprecated: Use BallotTestMetaData.ABI instead.
var BallotTestABI = BallotTestMetaData.ABI

// Deprecated: Use BallotTestMetaData.Sigs instead.
// BallotTestFuncSigs maps the 4-byte function signature to its string representation.
var BallotTestFuncSigs = BallotTestMetaData.Sigs

// BallotTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BallotTestMetaData.Bin instead.
var BallotTestBin = BallotTestMetaData.Bin

// DeployBallotTest deploys a new Ethereum contract, binding an instance of BallotTest to it.
func DeployBallotTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BallotTest, error) {
	parsed, err := BallotTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BallotTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BallotTest{BallotTestCaller: BallotTestCaller{contract: contract}, BallotTestTransactor: BallotTestTransactor{contract: contract}, BallotTestFilterer: BallotTestFilterer{contract: contract}}, nil
}

// BallotTest is an auto generated Go binding around an Ethereum contract.
type BallotTest struct {
	BallotTestCaller     // Read-only binding to the contract
	BallotTestTransactor // Write-only binding to the contract
	BallotTestFilterer   // Log filterer for contract events
}

// BallotTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type BallotTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BallotTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BallotTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BallotTestSession struct {
	Contract     *BallotTest       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BallotTestCallerSession struct {
	Contract *BallotTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BallotTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BallotTestTransactorSession struct {
	Contract     *BallotTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BallotTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type BallotTestRaw struct {
	Contract *BallotTest // Generic contract binding to access the raw methods on
}

// BallotTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BallotTestCallerRaw struct {
	Contract *BallotTestCaller // Generic read-only contract binding to access the raw methods on
}

// BallotTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BallotTestTransactorRaw struct {
	Contract *BallotTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBallotTest creates a new instance of BallotTest, bound to a specific deployed contract.
func NewBallotTest(address common.Address, backend bind.ContractBackend) (*BallotTest, error) {
	contract, err := bindBallotTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BallotTest{BallotTestCaller: BallotTestCaller{contract: contract}, BallotTestTransactor: BallotTestTransactor{contract: contract}, BallotTestFilterer: BallotTestFilterer{contract: contract}}, nil
}

// NewBallotTestCaller creates a new read-only instance of BallotTest, bound to a specific deployed contract.
func NewBallotTestCaller(address common.Address, caller bind.ContractCaller) (*BallotTestCaller, error) {
	contract, err := bindBallotTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BallotTestCaller{contract: contract}, nil
}

// NewBallotTestTransactor creates a new write-only instance of BallotTest, bound to a specific deployed contract.
func NewBallotTestTransactor(address common.Address, transactor bind.ContractTransactor) (*BallotTestTransactor, error) {
	contract, err := bindBallotTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BallotTestTransactor{contract: contract}, nil
}

// NewBallotTestFilterer creates a new log filterer instance of BallotTest, bound to a specific deployed contract.
func NewBallotTestFilterer(address common.Address, filterer bind.ContractFilterer) (*BallotTestFilterer, error) {
	contract, err := bindBallotTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BallotTestFilterer{contract: contract}, nil
}

// bindBallotTest binds a generic wrapper to an already deployed contract.
func bindBallotTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BallotTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BallotTest *BallotTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BallotTest.Contract.BallotTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BallotTest *BallotTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotTest.Contract.BallotTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BallotTest *BallotTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BallotTest.Contract.BallotTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BallotTest *BallotTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BallotTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BallotTest *BallotTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BallotTest *BallotTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BallotTest.Contract.contract.Transact(opts, method, params...)
}

// StoreBlobHash is a paid mutator transaction binding the contract method 0x11762da2.
//
// Solidity: function storeBlobHash() returns()
func (_BallotTest *BallotTestTransactor) StoreBlobHash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BallotTest.contract.Transact(opts, "storeBlobHash")
}

// StoreBlobHash is a paid mutator transaction binding the contract method 0x11762da2.
//
// Solidity: function storeBlobHash() returns()
func (_BallotTest *BallotTestSession) StoreBlobHash() (*types.Transaction, error) {
	return _BallotTest.Contract.StoreBlobHash(&_BallotTest.TransactOpts)
}

// StoreBlobHash is a paid mutator transaction binding the contract method 0x11762da2.
//
// Solidity: function storeBlobHash() returns()
func (_BallotTest *BallotTestTransactorSession) StoreBlobHash() (*types.Transaction, error) {
	return _BallotTest.Contract.StoreBlobHash(&_BallotTest.TransactOpts)
}
