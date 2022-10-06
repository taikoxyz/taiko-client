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

// LibConstantsMetaData contains all meta data concerning the LibConstants contract.
var LibConstantsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"TAIKO_BLOCK_DEADEND_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAIKO_BLOCK_MAX_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAIKO_BLOCK_MAX_TXS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAIKO_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAIKO_COMMIT_DELAY_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAIKO_MAX_FINALIZATIONS_PER_TX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAIKO_MAX_PROOFS_PER_FORK_CHOICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAIKO_MAX_PROPOSED_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAIKO_TXLIST_MAX_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAIKO_TX_MIN_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"V1_ANCHOR_TX_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"V1_ANCHOR_TX_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"V1_INVALIDATE_BLOCK_LOG_TOPIC\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LibConstantsABI is the input ABI used to generate the binding from.
// Deprecated: Use LibConstantsMetaData.ABI instead.
var LibConstantsABI = LibConstantsMetaData.ABI

// LibConstants is an auto generated Go binding around an Ethereum contract.
type LibConstants struct {
	LibConstantsCaller     // Read-only binding to the contract
	LibConstantsTransactor // Write-only binding to the contract
	LibConstantsFilterer   // Log filterer for contract events
}

// LibConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibConstantsSession struct {
	Contract     *LibConstants     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibConstantsCallerSession struct {
	Contract *LibConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LibConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibConstantsTransactorSession struct {
	Contract     *LibConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LibConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibConstantsRaw struct {
	Contract *LibConstants // Generic contract binding to access the raw methods on
}

// LibConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibConstantsCallerRaw struct {
	Contract *LibConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// LibConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibConstantsTransactorRaw struct {
	Contract *LibConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibConstants creates a new instance of LibConstants, bound to a specific deployed contract.
func NewLibConstants(address common.Address, backend bind.ContractBackend) (*LibConstants, error) {
	contract, err := bindLibConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibConstants{LibConstantsCaller: LibConstantsCaller{contract: contract}, LibConstantsTransactor: LibConstantsTransactor{contract: contract}, LibConstantsFilterer: LibConstantsFilterer{contract: contract}}, nil
}

// NewLibConstantsCaller creates a new read-only instance of LibConstants, bound to a specific deployed contract.
func NewLibConstantsCaller(address common.Address, caller bind.ContractCaller) (*LibConstantsCaller, error) {
	contract, err := bindLibConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibConstantsCaller{contract: contract}, nil
}

// NewLibConstantsTransactor creates a new write-only instance of LibConstants, bound to a specific deployed contract.
func NewLibConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*LibConstantsTransactor, error) {
	contract, err := bindLibConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibConstantsTransactor{contract: contract}, nil
}

// NewLibConstantsFilterer creates a new log filterer instance of LibConstants, bound to a specific deployed contract.
func NewLibConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*LibConstantsFilterer, error) {
	contract, err := bindLibConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibConstantsFilterer{contract: contract}, nil
}

// bindLibConstants binds a generic wrapper to an already deployed contract.
func bindLibConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibConstantsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibConstants *LibConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibConstants.Contract.LibConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibConstants *LibConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibConstants.Contract.LibConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibConstants *LibConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibConstants.Contract.LibConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibConstants *LibConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibConstants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibConstants *LibConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibConstants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibConstants *LibConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibConstants.Contract.contract.Transact(opts, method, params...)
}

// TAIKOBLOCKDEADENDHASH is a free data retrieval call binding the contract method 0xc5d4233c.
//
// Solidity: function TAIKO_BLOCK_DEADEND_HASH() view returns(bytes32)
func (_LibConstants *LibConstantsCaller) TAIKOBLOCKDEADENDHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "TAIKO_BLOCK_DEADEND_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TAIKOBLOCKDEADENDHASH is a free data retrieval call binding the contract method 0xc5d4233c.
//
// Solidity: function TAIKO_BLOCK_DEADEND_HASH() view returns(bytes32)
func (_LibConstants *LibConstantsSession) TAIKOBLOCKDEADENDHASH() ([32]byte, error) {
	return _LibConstants.Contract.TAIKOBLOCKDEADENDHASH(&_LibConstants.CallOpts)
}

// TAIKOBLOCKDEADENDHASH is a free data retrieval call binding the contract method 0xc5d4233c.
//
// Solidity: function TAIKO_BLOCK_DEADEND_HASH() view returns(bytes32)
func (_LibConstants *LibConstantsCallerSession) TAIKOBLOCKDEADENDHASH() ([32]byte, error) {
	return _LibConstants.Contract.TAIKOBLOCKDEADENDHASH(&_LibConstants.CallOpts)
}

// TAIKOBLOCKMAXGASLIMIT is a free data retrieval call binding the contract method 0xa45d500c.
//
// Solidity: function TAIKO_BLOCK_MAX_GAS_LIMIT() view returns(uint256)
func (_LibConstants *LibConstantsCaller) TAIKOBLOCKMAXGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "TAIKO_BLOCK_MAX_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAIKOBLOCKMAXGASLIMIT is a free data retrieval call binding the contract method 0xa45d500c.
//
// Solidity: function TAIKO_BLOCK_MAX_GAS_LIMIT() view returns(uint256)
func (_LibConstants *LibConstantsSession) TAIKOBLOCKMAXGASLIMIT() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOBLOCKMAXGASLIMIT(&_LibConstants.CallOpts)
}

// TAIKOBLOCKMAXGASLIMIT is a free data retrieval call binding the contract method 0xa45d500c.
//
// Solidity: function TAIKO_BLOCK_MAX_GAS_LIMIT() view returns(uint256)
func (_LibConstants *LibConstantsCallerSession) TAIKOBLOCKMAXGASLIMIT() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOBLOCKMAXGASLIMIT(&_LibConstants.CallOpts)
}

// TAIKOBLOCKMAXTXS is a free data retrieval call binding the contract method 0xf36df3c9.
//
// Solidity: function TAIKO_BLOCK_MAX_TXS() view returns(uint256)
func (_LibConstants *LibConstantsCaller) TAIKOBLOCKMAXTXS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "TAIKO_BLOCK_MAX_TXS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAIKOBLOCKMAXTXS is a free data retrieval call binding the contract method 0xf36df3c9.
//
// Solidity: function TAIKO_BLOCK_MAX_TXS() view returns(uint256)
func (_LibConstants *LibConstantsSession) TAIKOBLOCKMAXTXS() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOBLOCKMAXTXS(&_LibConstants.CallOpts)
}

// TAIKOBLOCKMAXTXS is a free data retrieval call binding the contract method 0xf36df3c9.
//
// Solidity: function TAIKO_BLOCK_MAX_TXS() view returns(uint256)
func (_LibConstants *LibConstantsCallerSession) TAIKOBLOCKMAXTXS() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOBLOCKMAXTXS(&_LibConstants.CallOpts)
}

// TAIKOCHAINID is a free data retrieval call binding the contract method 0xf77e0914.
//
// Solidity: function TAIKO_CHAIN_ID() view returns(uint256)
func (_LibConstants *LibConstantsCaller) TAIKOCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "TAIKO_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAIKOCHAINID is a free data retrieval call binding the contract method 0xf77e0914.
//
// Solidity: function TAIKO_CHAIN_ID() view returns(uint256)
func (_LibConstants *LibConstantsSession) TAIKOCHAINID() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOCHAINID(&_LibConstants.CallOpts)
}

// TAIKOCHAINID is a free data retrieval call binding the contract method 0xf77e0914.
//
// Solidity: function TAIKO_CHAIN_ID() view returns(uint256)
func (_LibConstants *LibConstantsCallerSession) TAIKOCHAINID() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOCHAINID(&_LibConstants.CallOpts)
}

// TAIKOCOMMITDELAYCONFIRMATIONS is a free data retrieval call binding the contract method 0x81d83d27.
//
// Solidity: function TAIKO_COMMIT_DELAY_CONFIRMATIONS() view returns(uint256)
func (_LibConstants *LibConstantsCaller) TAIKOCOMMITDELAYCONFIRMATIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "TAIKO_COMMIT_DELAY_CONFIRMATIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAIKOCOMMITDELAYCONFIRMATIONS is a free data retrieval call binding the contract method 0x81d83d27.
//
// Solidity: function TAIKO_COMMIT_DELAY_CONFIRMATIONS() view returns(uint256)
func (_LibConstants *LibConstantsSession) TAIKOCOMMITDELAYCONFIRMATIONS() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOCOMMITDELAYCONFIRMATIONS(&_LibConstants.CallOpts)
}

// TAIKOCOMMITDELAYCONFIRMATIONS is a free data retrieval call binding the contract method 0x81d83d27.
//
// Solidity: function TAIKO_COMMIT_DELAY_CONFIRMATIONS() view returns(uint256)
func (_LibConstants *LibConstantsCallerSession) TAIKOCOMMITDELAYCONFIRMATIONS() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOCOMMITDELAYCONFIRMATIONS(&_LibConstants.CallOpts)
}

// TAIKOMAXFINALIZATIONSPERTX is a free data retrieval call binding the contract method 0x63ad5776.
//
// Solidity: function TAIKO_MAX_FINALIZATIONS_PER_TX() view returns(uint256)
func (_LibConstants *LibConstantsCaller) TAIKOMAXFINALIZATIONSPERTX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "TAIKO_MAX_FINALIZATIONS_PER_TX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAIKOMAXFINALIZATIONSPERTX is a free data retrieval call binding the contract method 0x63ad5776.
//
// Solidity: function TAIKO_MAX_FINALIZATIONS_PER_TX() view returns(uint256)
func (_LibConstants *LibConstantsSession) TAIKOMAXFINALIZATIONSPERTX() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOMAXFINALIZATIONSPERTX(&_LibConstants.CallOpts)
}

// TAIKOMAXFINALIZATIONSPERTX is a free data retrieval call binding the contract method 0x63ad5776.
//
// Solidity: function TAIKO_MAX_FINALIZATIONS_PER_TX() view returns(uint256)
func (_LibConstants *LibConstantsCallerSession) TAIKOMAXFINALIZATIONSPERTX() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOMAXFINALIZATIONSPERTX(&_LibConstants.CallOpts)
}

// TAIKOMAXPROOFSPERFORKCHOICE is a free data retrieval call binding the contract method 0xc0043906.
//
// Solidity: function TAIKO_MAX_PROOFS_PER_FORK_CHOICE() view returns(uint256)
func (_LibConstants *LibConstantsCaller) TAIKOMAXPROOFSPERFORKCHOICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "TAIKO_MAX_PROOFS_PER_FORK_CHOICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAIKOMAXPROOFSPERFORKCHOICE is a free data retrieval call binding the contract method 0xc0043906.
//
// Solidity: function TAIKO_MAX_PROOFS_PER_FORK_CHOICE() view returns(uint256)
func (_LibConstants *LibConstantsSession) TAIKOMAXPROOFSPERFORKCHOICE() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOMAXPROOFSPERFORKCHOICE(&_LibConstants.CallOpts)
}

// TAIKOMAXPROOFSPERFORKCHOICE is a free data retrieval call binding the contract method 0xc0043906.
//
// Solidity: function TAIKO_MAX_PROOFS_PER_FORK_CHOICE() view returns(uint256)
func (_LibConstants *LibConstantsCallerSession) TAIKOMAXPROOFSPERFORKCHOICE() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOMAXPROOFSPERFORKCHOICE(&_LibConstants.CallOpts)
}

// TAIKOMAXPROPOSEDBLOCKS is a free data retrieval call binding the contract method 0x49fbd976.
//
// Solidity: function TAIKO_MAX_PROPOSED_BLOCKS() view returns(uint256)
func (_LibConstants *LibConstantsCaller) TAIKOMAXPROPOSEDBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "TAIKO_MAX_PROPOSED_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAIKOMAXPROPOSEDBLOCKS is a free data retrieval call binding the contract method 0x49fbd976.
//
// Solidity: function TAIKO_MAX_PROPOSED_BLOCKS() view returns(uint256)
func (_LibConstants *LibConstantsSession) TAIKOMAXPROPOSEDBLOCKS() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOMAXPROPOSEDBLOCKS(&_LibConstants.CallOpts)
}

// TAIKOMAXPROPOSEDBLOCKS is a free data retrieval call binding the contract method 0x49fbd976.
//
// Solidity: function TAIKO_MAX_PROPOSED_BLOCKS() view returns(uint256)
func (_LibConstants *LibConstantsCallerSession) TAIKOMAXPROPOSEDBLOCKS() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOMAXPROPOSEDBLOCKS(&_LibConstants.CallOpts)
}

// TAIKOTXLISTMAXBYTES is a free data retrieval call binding the contract method 0x1577ef66.
//
// Solidity: function TAIKO_TXLIST_MAX_BYTES() view returns(uint256)
func (_LibConstants *LibConstantsCaller) TAIKOTXLISTMAXBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "TAIKO_TXLIST_MAX_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAIKOTXLISTMAXBYTES is a free data retrieval call binding the contract method 0x1577ef66.
//
// Solidity: function TAIKO_TXLIST_MAX_BYTES() view returns(uint256)
func (_LibConstants *LibConstantsSession) TAIKOTXLISTMAXBYTES() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOTXLISTMAXBYTES(&_LibConstants.CallOpts)
}

// TAIKOTXLISTMAXBYTES is a free data retrieval call binding the contract method 0x1577ef66.
//
// Solidity: function TAIKO_TXLIST_MAX_BYTES() view returns(uint256)
func (_LibConstants *LibConstantsCallerSession) TAIKOTXLISTMAXBYTES() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOTXLISTMAXBYTES(&_LibConstants.CallOpts)
}

// TAIKOTXMINGASLIMIT is a free data retrieval call binding the contract method 0xbe8d42e0.
//
// Solidity: function TAIKO_TX_MIN_GAS_LIMIT() view returns(uint256)
func (_LibConstants *LibConstantsCaller) TAIKOTXMINGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "TAIKO_TX_MIN_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAIKOTXMINGASLIMIT is a free data retrieval call binding the contract method 0xbe8d42e0.
//
// Solidity: function TAIKO_TX_MIN_GAS_LIMIT() view returns(uint256)
func (_LibConstants *LibConstantsSession) TAIKOTXMINGASLIMIT() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOTXMINGASLIMIT(&_LibConstants.CallOpts)
}

// TAIKOTXMINGASLIMIT is a free data retrieval call binding the contract method 0xbe8d42e0.
//
// Solidity: function TAIKO_TX_MIN_GAS_LIMIT() view returns(uint256)
func (_LibConstants *LibConstantsCallerSession) TAIKOTXMINGASLIMIT() (*big.Int, error) {
	return _LibConstants.Contract.TAIKOTXMINGASLIMIT(&_LibConstants.CallOpts)
}

// V1ANCHORTXGASLIMIT is a free data retrieval call binding the contract method 0xc3676578.
//
// Solidity: function V1_ANCHOR_TX_GAS_LIMIT() view returns(uint256)
func (_LibConstants *LibConstantsCaller) V1ANCHORTXGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "V1_ANCHOR_TX_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// V1ANCHORTXGASLIMIT is a free data retrieval call binding the contract method 0xc3676578.
//
// Solidity: function V1_ANCHOR_TX_GAS_LIMIT() view returns(uint256)
func (_LibConstants *LibConstantsSession) V1ANCHORTXGASLIMIT() (*big.Int, error) {
	return _LibConstants.Contract.V1ANCHORTXGASLIMIT(&_LibConstants.CallOpts)
}

// V1ANCHORTXGASLIMIT is a free data retrieval call binding the contract method 0xc3676578.
//
// Solidity: function V1_ANCHOR_TX_GAS_LIMIT() view returns(uint256)
func (_LibConstants *LibConstantsCallerSession) V1ANCHORTXGASLIMIT() (*big.Int, error) {
	return _LibConstants.Contract.V1ANCHORTXGASLIMIT(&_LibConstants.CallOpts)
}

// V1ANCHORTXSELECTOR is a free data retrieval call binding the contract method 0x72dc8810.
//
// Solidity: function V1_ANCHOR_TX_SELECTOR() view returns(bytes4)
func (_LibConstants *LibConstantsCaller) V1ANCHORTXSELECTOR(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "V1_ANCHOR_TX_SELECTOR")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// V1ANCHORTXSELECTOR is a free data retrieval call binding the contract method 0x72dc8810.
//
// Solidity: function V1_ANCHOR_TX_SELECTOR() view returns(bytes4)
func (_LibConstants *LibConstantsSession) V1ANCHORTXSELECTOR() ([4]byte, error) {
	return _LibConstants.Contract.V1ANCHORTXSELECTOR(&_LibConstants.CallOpts)
}

// V1ANCHORTXSELECTOR is a free data retrieval call binding the contract method 0x72dc8810.
//
// Solidity: function V1_ANCHOR_TX_SELECTOR() view returns(bytes4)
func (_LibConstants *LibConstantsCallerSession) V1ANCHORTXSELECTOR() ([4]byte, error) {
	return _LibConstants.Contract.V1ANCHORTXSELECTOR(&_LibConstants.CallOpts)
}

// V1INVALIDATEBLOCKLOGTOPIC is a free data retrieval call binding the contract method 0x93a0ca2f.
//
// Solidity: function V1_INVALIDATE_BLOCK_LOG_TOPIC() view returns(bytes32)
func (_LibConstants *LibConstantsCaller) V1INVALIDATEBLOCKLOGTOPIC(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LibConstants.contract.Call(opts, &out, "V1_INVALIDATE_BLOCK_LOG_TOPIC")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// V1INVALIDATEBLOCKLOGTOPIC is a free data retrieval call binding the contract method 0x93a0ca2f.
//
// Solidity: function V1_INVALIDATE_BLOCK_LOG_TOPIC() view returns(bytes32)
func (_LibConstants *LibConstantsSession) V1INVALIDATEBLOCKLOGTOPIC() ([32]byte, error) {
	return _LibConstants.Contract.V1INVALIDATEBLOCKLOGTOPIC(&_LibConstants.CallOpts)
}

// V1INVALIDATEBLOCKLOGTOPIC is a free data retrieval call binding the contract method 0x93a0ca2f.
//
// Solidity: function V1_INVALIDATE_BLOCK_LOG_TOPIC() view returns(bytes32)
func (_LibConstants *LibConstantsCallerSession) V1INVALIDATEBLOCKLOGTOPIC() ([32]byte, error) {
	return _LibConstants.Contract.V1INVALIDATEBLOCKLOGTOPIC(&_LibConstants.CallOpts)
}
