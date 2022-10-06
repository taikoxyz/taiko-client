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

// LibTxDecoderTx is an auto generated low-level Go binding around an user-defined struct.
type LibTxDecoderTx struct {
	TxType      uint8
	Destination common.Address
	Data        []byte
	GasLimit    *big.Int
	V           uint8
	R           *big.Int
	S           *big.Int
	TxData      []byte
}

// LibTxDecoderTxList is an auto generated low-level Go binding around an user-defined struct.
type LibTxDecoderTxList struct {
	Items []LibTxDecoderTx
}

// LibTxDecoderMetaData contains all meta data concerning the LibTxDecoder contract.
var LibTxDecoderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"decodeTx\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"}],\"internalType\":\"structLibTxDecoder.Tx\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"decodeTxList\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"}],\"internalType\":\"structLibTxDecoder.Tx[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"internalType\":\"structLibTxDecoder.TxList\",\"name\":\"txList\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// LibTxDecoderABI is the input ABI used to generate the binding from.
// Deprecated: Use LibTxDecoderMetaData.ABI instead.
var LibTxDecoderABI = LibTxDecoderMetaData.ABI

// LibTxDecoder is an auto generated Go binding around an Ethereum contract.
type LibTxDecoder struct {
	LibTxDecoderCaller     // Read-only binding to the contract
	LibTxDecoderTransactor // Write-only binding to the contract
	LibTxDecoderFilterer   // Log filterer for contract events
}

// LibTxDecoderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibTxDecoderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibTxDecoderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibTxDecoderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibTxDecoderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibTxDecoderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibTxDecoderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibTxDecoderSession struct {
	Contract     *LibTxDecoder     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibTxDecoderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibTxDecoderCallerSession struct {
	Contract *LibTxDecoderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LibTxDecoderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibTxDecoderTransactorSession struct {
	Contract     *LibTxDecoderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LibTxDecoderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibTxDecoderRaw struct {
	Contract *LibTxDecoder // Generic contract binding to access the raw methods on
}

// LibTxDecoderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibTxDecoderCallerRaw struct {
	Contract *LibTxDecoderCaller // Generic read-only contract binding to access the raw methods on
}

// LibTxDecoderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibTxDecoderTransactorRaw struct {
	Contract *LibTxDecoderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibTxDecoder creates a new instance of LibTxDecoder, bound to a specific deployed contract.
func NewLibTxDecoder(address common.Address, backend bind.ContractBackend) (*LibTxDecoder, error) {
	contract, err := bindLibTxDecoder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibTxDecoder{LibTxDecoderCaller: LibTxDecoderCaller{contract: contract}, LibTxDecoderTransactor: LibTxDecoderTransactor{contract: contract}, LibTxDecoderFilterer: LibTxDecoderFilterer{contract: contract}}, nil
}

// NewLibTxDecoderCaller creates a new read-only instance of LibTxDecoder, bound to a specific deployed contract.
func NewLibTxDecoderCaller(address common.Address, caller bind.ContractCaller) (*LibTxDecoderCaller, error) {
	contract, err := bindLibTxDecoder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibTxDecoderCaller{contract: contract}, nil
}

// NewLibTxDecoderTransactor creates a new write-only instance of LibTxDecoder, bound to a specific deployed contract.
func NewLibTxDecoderTransactor(address common.Address, transactor bind.ContractTransactor) (*LibTxDecoderTransactor, error) {
	contract, err := bindLibTxDecoder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibTxDecoderTransactor{contract: contract}, nil
}

// NewLibTxDecoderFilterer creates a new log filterer instance of LibTxDecoder, bound to a specific deployed contract.
func NewLibTxDecoderFilterer(address common.Address, filterer bind.ContractFilterer) (*LibTxDecoderFilterer, error) {
	contract, err := bindLibTxDecoder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibTxDecoderFilterer{contract: contract}, nil
}

// bindLibTxDecoder binds a generic wrapper to an already deployed contract.
func bindLibTxDecoder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibTxDecoderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibTxDecoder *LibTxDecoderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibTxDecoder.Contract.LibTxDecoderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibTxDecoder *LibTxDecoderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibTxDecoder.Contract.LibTxDecoderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibTxDecoder *LibTxDecoderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibTxDecoder.Contract.LibTxDecoderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibTxDecoder *LibTxDecoderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibTxDecoder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibTxDecoder *LibTxDecoderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibTxDecoder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibTxDecoder *LibTxDecoderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibTxDecoder.Contract.contract.Transact(opts, method, params...)
}

// DecodeTx is a free data retrieval call binding the contract method 0xdae029d3.
//
// Solidity: function decodeTx(bytes txBytes) pure returns((uint8,address,bytes,uint256,uint8,uint256,uint256,bytes) _tx)
func (_LibTxDecoder *LibTxDecoderCaller) DecodeTx(opts *bind.CallOpts, txBytes []byte) (LibTxDecoderTx, error) {
	var out []interface{}
	err := _LibTxDecoder.contract.Call(opts, &out, "decodeTx", txBytes)

	if err != nil {
		return *new(LibTxDecoderTx), err
	}

	out0 := *abi.ConvertType(out[0], new(LibTxDecoderTx)).(*LibTxDecoderTx)

	return out0, err

}

// DecodeTx is a free data retrieval call binding the contract method 0xdae029d3.
//
// Solidity: function decodeTx(bytes txBytes) pure returns((uint8,address,bytes,uint256,uint8,uint256,uint256,bytes) _tx)
func (_LibTxDecoder *LibTxDecoderSession) DecodeTx(txBytes []byte) (LibTxDecoderTx, error) {
	return _LibTxDecoder.Contract.DecodeTx(&_LibTxDecoder.CallOpts, txBytes)
}

// DecodeTx is a free data retrieval call binding the contract method 0xdae029d3.
//
// Solidity: function decodeTx(bytes txBytes) pure returns((uint8,address,bytes,uint256,uint8,uint256,uint256,bytes) _tx)
func (_LibTxDecoder *LibTxDecoderCallerSession) DecodeTx(txBytes []byte) (LibTxDecoderTx, error) {
	return _LibTxDecoder.Contract.DecodeTx(&_LibTxDecoder.CallOpts, txBytes)
}

// DecodeTxList is a free data retrieval call binding the contract method 0x2cb6101a.
//
// Solidity: function decodeTxList(bytes encoded) pure returns(((uint8,address,bytes,uint256,uint8,uint256,uint256,bytes)[]) txList)
func (_LibTxDecoder *LibTxDecoderCaller) DecodeTxList(opts *bind.CallOpts, encoded []byte) (LibTxDecoderTxList, error) {
	var out []interface{}
	err := _LibTxDecoder.contract.Call(opts, &out, "decodeTxList", encoded)

	if err != nil {
		return *new(LibTxDecoderTxList), err
	}

	out0 := *abi.ConvertType(out[0], new(LibTxDecoderTxList)).(*LibTxDecoderTxList)

	return out0, err

}

// DecodeTxList is a free data retrieval call binding the contract method 0x2cb6101a.
//
// Solidity: function decodeTxList(bytes encoded) pure returns(((uint8,address,bytes,uint256,uint8,uint256,uint256,bytes)[]) txList)
func (_LibTxDecoder *LibTxDecoderSession) DecodeTxList(encoded []byte) (LibTxDecoderTxList, error) {
	return _LibTxDecoder.Contract.DecodeTxList(&_LibTxDecoder.CallOpts, encoded)
}

// DecodeTxList is a free data retrieval call binding the contract method 0x2cb6101a.
//
// Solidity: function decodeTxList(bytes encoded) pure returns(((uint8,address,bytes,uint256,uint8,uint256,uint256,bytes)[]) txList)
func (_LibTxDecoder *LibTxDecoderCallerSession) DecodeTxList(encoded []byte) (LibTxDecoderTxList, error) {
	return _LibTxDecoder.Contract.DecodeTxList(&_LibTxDecoder.CallOpts, encoded)
}
