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

// ERC20VotesUpgradeableCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type ERC20VotesUpgradeableCheckpoint struct {
	FromBlock uint32
	Votes     *big.Int
}

// TaikoTokenMetaData contains all meta data concerning the TaikoToken contract.
var TaikoTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TKO_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TKO_INVALID_PREMINT_PARAMS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addressManager\",\"type\":\"address\"}],\"name\":\"AddressManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"pos\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint224\",\"name\":\"votes\",\"type\":\"uint224\"}],\"internalType\":\"structERC20VotesUpgradeable.Checkpoint\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPastTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPastVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_premintRecipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_premintAmounts\",\"type\":\"uint256[]\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TaikoTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TaikoTokenMetaData.ABI instead.
var TaikoTokenABI = TaikoTokenMetaData.ABI

// TaikoToken is an auto generated Go binding around an Ethereum contract.
type TaikoToken struct {
	TaikoTokenCaller     // Read-only binding to the contract
	TaikoTokenTransactor // Write-only binding to the contract
	TaikoTokenFilterer   // Log filterer for contract events
}

// TaikoTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaikoTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaikoTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaikoTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaikoTokenSession struct {
	Contract     *TaikoToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaikoTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaikoTokenCallerSession struct {
	Contract *TaikoTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TaikoTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaikoTokenTransactorSession struct {
	Contract     *TaikoTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TaikoTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaikoTokenRaw struct {
	Contract *TaikoToken // Generic contract binding to access the raw methods on
}

// TaikoTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaikoTokenCallerRaw struct {
	Contract *TaikoTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TaikoTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaikoTokenTransactorRaw struct {
	Contract *TaikoTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaikoToken creates a new instance of TaikoToken, bound to a specific deployed contract.
func NewTaikoToken(address common.Address, backend bind.ContractBackend) (*TaikoToken, error) {
	contract, err := bindTaikoToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaikoToken{TaikoTokenCaller: TaikoTokenCaller{contract: contract}, TaikoTokenTransactor: TaikoTokenTransactor{contract: contract}, TaikoTokenFilterer: TaikoTokenFilterer{contract: contract}}, nil
}

// NewTaikoTokenCaller creates a new read-only instance of TaikoToken, bound to a specific deployed contract.
func NewTaikoTokenCaller(address common.Address, caller bind.ContractCaller) (*TaikoTokenCaller, error) {
	contract, err := bindTaikoToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoTokenCaller{contract: contract}, nil
}

// NewTaikoTokenTransactor creates a new write-only instance of TaikoToken, bound to a specific deployed contract.
func NewTaikoTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TaikoTokenTransactor, error) {
	contract, err := bindTaikoToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoTokenTransactor{contract: contract}, nil
}

// NewTaikoTokenFilterer creates a new log filterer instance of TaikoToken, bound to a specific deployed contract.
func NewTaikoTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TaikoTokenFilterer, error) {
	contract, err := bindTaikoToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaikoTokenFilterer{contract: contract}, nil
}

// bindTaikoToken binds a generic wrapper to an already deployed contract.
func bindTaikoToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaikoTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoToken *TaikoTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoToken.Contract.TaikoTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoToken *TaikoTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoToken.Contract.TaikoTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoToken *TaikoTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoToken.Contract.TaikoTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoToken *TaikoTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoToken *TaikoTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoToken *TaikoTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TaikoToken *TaikoTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TaikoToken *TaikoTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _TaikoToken.Contract.DOMAINSEPARATOR(&_TaikoToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TaikoToken *TaikoTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _TaikoToken.Contract.DOMAINSEPARATOR(&_TaikoToken.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoToken *TaikoTokenCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoToken *TaikoTokenSession) AddressManager() (common.Address, error) {
	return _TaikoToken.Contract.AddressManager(&_TaikoToken.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoToken *TaikoTokenCallerSession) AddressManager() (common.Address, error) {
	return _TaikoToken.Contract.AddressManager(&_TaikoToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TaikoToken *TaikoTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TaikoToken *TaikoTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TaikoToken.Contract.Allowance(&_TaikoToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TaikoToken *TaikoTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TaikoToken.Contract.Allowance(&_TaikoToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TaikoToken *TaikoTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TaikoToken *TaikoTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TaikoToken.Contract.BalanceOf(&_TaikoToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TaikoToken *TaikoTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TaikoToken.Contract.BalanceOf(&_TaikoToken.CallOpts, account)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_TaikoToken *TaikoTokenCaller) BalanceOfAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "balanceOfAt", account, snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_TaikoToken *TaikoTokenSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _TaikoToken.Contract.BalanceOfAt(&_TaikoToken.CallOpts, account, snapshotId)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_TaikoToken *TaikoTokenCallerSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _TaikoToken.Contract.BalanceOfAt(&_TaikoToken.CallOpts, account, snapshotId)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_TaikoToken *TaikoTokenCaller) Checkpoints(opts *bind.CallOpts, account common.Address, pos uint32) (ERC20VotesUpgradeableCheckpoint, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "checkpoints", account, pos)

	if err != nil {
		return *new(ERC20VotesUpgradeableCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20VotesUpgradeableCheckpoint)).(*ERC20VotesUpgradeableCheckpoint)

	return out0, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_TaikoToken *TaikoTokenSession) Checkpoints(account common.Address, pos uint32) (ERC20VotesUpgradeableCheckpoint, error) {
	return _TaikoToken.Contract.Checkpoints(&_TaikoToken.CallOpts, account, pos)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_TaikoToken *TaikoTokenCallerSession) Checkpoints(account common.Address, pos uint32) (ERC20VotesUpgradeableCheckpoint, error) {
	return _TaikoToken.Contract.Checkpoints(&_TaikoToken.CallOpts, account, pos)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TaikoToken *TaikoTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TaikoToken *TaikoTokenSession) Decimals() (uint8, error) {
	return _TaikoToken.Contract.Decimals(&_TaikoToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TaikoToken *TaikoTokenCallerSession) Decimals() (uint8, error) {
	return _TaikoToken.Contract.Decimals(&_TaikoToken.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_TaikoToken *TaikoTokenCaller) Delegates(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "delegates", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_TaikoToken *TaikoTokenSession) Delegates(account common.Address) (common.Address, error) {
	return _TaikoToken.Contract.Delegates(&_TaikoToken.CallOpts, account)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_TaikoToken *TaikoTokenCallerSession) Delegates(account common.Address) (common.Address, error) {
	return _TaikoToken.Contract.Delegates(&_TaikoToken.CallOpts, account)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_TaikoToken *TaikoTokenCaller) GetPastTotalSupply(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "getPastTotalSupply", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_TaikoToken *TaikoTokenSession) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	return _TaikoToken.Contract.GetPastTotalSupply(&_TaikoToken.CallOpts, blockNumber)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_TaikoToken *TaikoTokenCallerSession) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	return _TaikoToken.Contract.GetPastTotalSupply(&_TaikoToken.CallOpts, blockNumber)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_TaikoToken *TaikoTokenCaller) GetPastVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "getPastVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_TaikoToken *TaikoTokenSession) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _TaikoToken.Contract.GetPastVotes(&_TaikoToken.CallOpts, account, blockNumber)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_TaikoToken *TaikoTokenCallerSession) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _TaikoToken.Contract.GetPastVotes(&_TaikoToken.CallOpts, account, blockNumber)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_TaikoToken *TaikoTokenCaller) GetVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_TaikoToken *TaikoTokenSession) GetVotes(account common.Address) (*big.Int, error) {
	return _TaikoToken.Contract.GetVotes(&_TaikoToken.CallOpts, account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_TaikoToken *TaikoTokenCallerSession) GetVotes(account common.Address) (*big.Int, error) {
	return _TaikoToken.Contract.GetVotes(&_TaikoToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TaikoToken *TaikoTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TaikoToken *TaikoTokenSession) Name() (string, error) {
	return _TaikoToken.Contract.Name(&_TaikoToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TaikoToken *TaikoTokenCallerSession) Name() (string, error) {
	return _TaikoToken.Contract.Name(&_TaikoToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_TaikoToken *TaikoTokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_TaikoToken *TaikoTokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _TaikoToken.Contract.Nonces(&_TaikoToken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_TaikoToken *TaikoTokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _TaikoToken.Contract.Nonces(&_TaikoToken.CallOpts, owner)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_TaikoToken *TaikoTokenCaller) NumCheckpoints(opts *bind.CallOpts, account common.Address) (uint32, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "numCheckpoints", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_TaikoToken *TaikoTokenSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _TaikoToken.Contract.NumCheckpoints(&_TaikoToken.CallOpts, account)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_TaikoToken *TaikoTokenCallerSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _TaikoToken.Contract.NumCheckpoints(&_TaikoToken.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoToken *TaikoTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoToken *TaikoTokenSession) Owner() (common.Address, error) {
	return _TaikoToken.Contract.Owner(&_TaikoToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoToken *TaikoTokenCallerSession) Owner() (common.Address, error) {
	return _TaikoToken.Contract.Owner(&_TaikoToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TaikoToken *TaikoTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TaikoToken *TaikoTokenSession) Paused() (bool, error) {
	return _TaikoToken.Contract.Paused(&_TaikoToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TaikoToken *TaikoTokenCallerSession) Paused() (bool, error) {
	return _TaikoToken.Contract.Paused(&_TaikoToken.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoToken *TaikoTokenCaller) Resolve(opts *bind.CallOpts, chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoToken *TaikoTokenSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoToken.Contract.Resolve(&_TaikoToken.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoToken *TaikoTokenCallerSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoToken.Contract.Resolve(&_TaikoToken.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoToken *TaikoTokenCaller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoToken *TaikoTokenSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoToken.Contract.Resolve0(&_TaikoToken.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoToken *TaikoTokenCallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoToken.Contract.Resolve0(&_TaikoToken.CallOpts, name, allowZeroAddress)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TaikoToken *TaikoTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TaikoToken *TaikoTokenSession) Symbol() (string, error) {
	return _TaikoToken.Contract.Symbol(&_TaikoToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TaikoToken *TaikoTokenCallerSession) Symbol() (string, error) {
	return _TaikoToken.Contract.Symbol(&_TaikoToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TaikoToken *TaikoTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TaikoToken *TaikoTokenSession) TotalSupply() (*big.Int, error) {
	return _TaikoToken.Contract.TotalSupply(&_TaikoToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TaikoToken *TaikoTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TaikoToken.Contract.TotalSupply(&_TaikoToken.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_TaikoToken *TaikoTokenCaller) TotalSupplyAt(opts *bind.CallOpts, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TaikoToken.contract.Call(opts, &out, "totalSupplyAt", snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_TaikoToken *TaikoTokenSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _TaikoToken.Contract.TotalSupplyAt(&_TaikoToken.CallOpts, snapshotId)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_TaikoToken *TaikoTokenCallerSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _TaikoToken.Contract.TotalSupplyAt(&_TaikoToken.CallOpts, snapshotId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TaikoToken *TaikoTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TaikoToken *TaikoTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Approve(&_TaikoToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TaikoToken *TaikoTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Approve(&_TaikoToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_TaikoToken *TaikoTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_TaikoToken *TaikoTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Burn(&_TaikoToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_TaikoToken *TaikoTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Burn(&_TaikoToken.TransactOpts, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_TaikoToken *TaikoTokenTransactor) Burn0(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "burn0", from, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_TaikoToken *TaikoTokenSession) Burn0(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Burn0(&_TaikoToken.TransactOpts, from, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_TaikoToken *TaikoTokenTransactorSession) Burn0(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Burn0(&_TaikoToken.TransactOpts, from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_TaikoToken *TaikoTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_TaikoToken *TaikoTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.BurnFrom(&_TaikoToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_TaikoToken *TaikoTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.BurnFrom(&_TaikoToken.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TaikoToken *TaikoTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TaikoToken *TaikoTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.DecreaseAllowance(&_TaikoToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TaikoToken *TaikoTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.DecreaseAllowance(&_TaikoToken.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_TaikoToken *TaikoTokenTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_TaikoToken *TaikoTokenSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _TaikoToken.Contract.Delegate(&_TaikoToken.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_TaikoToken *TaikoTokenTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _TaikoToken.Contract.Delegate(&_TaikoToken.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_TaikoToken *TaikoTokenTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_TaikoToken *TaikoTokenSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TaikoToken.Contract.DelegateBySig(&_TaikoToken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_TaikoToken *TaikoTokenTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TaikoToken.Contract.DelegateBySig(&_TaikoToken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TaikoToken *TaikoTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TaikoToken *TaikoTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.IncreaseAllowance(&_TaikoToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TaikoToken *TaikoTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.IncreaseAllowance(&_TaikoToken.TransactOpts, spender, addedValue)
}

// Init is a paid mutator transaction binding the contract method 0xc3a57123.
//
// Solidity: function init(address _addressManager, string _name, string _symbol, address[] _premintRecipients, uint256[] _premintAmounts) returns()
func (_TaikoToken *TaikoTokenTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _name string, _symbol string, _premintRecipients []common.Address, _premintAmounts []*big.Int) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "init", _addressManager, _name, _symbol, _premintRecipients, _premintAmounts)
}

// Init is a paid mutator transaction binding the contract method 0xc3a57123.
//
// Solidity: function init(address _addressManager, string _name, string _symbol, address[] _premintRecipients, uint256[] _premintAmounts) returns()
func (_TaikoToken *TaikoTokenSession) Init(_addressManager common.Address, _name string, _symbol string, _premintRecipients []common.Address, _premintAmounts []*big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Init(&_TaikoToken.TransactOpts, _addressManager, _name, _symbol, _premintRecipients, _premintAmounts)
}

// Init is a paid mutator transaction binding the contract method 0xc3a57123.
//
// Solidity: function init(address _addressManager, string _name, string _symbol, address[] _premintRecipients, uint256[] _premintAmounts) returns()
func (_TaikoToken *TaikoTokenTransactorSession) Init(_addressManager common.Address, _name string, _symbol string, _premintRecipients []common.Address, _premintAmounts []*big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Init(&_TaikoToken.TransactOpts, _addressManager, _name, _symbol, _premintRecipients, _premintAmounts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TaikoToken *TaikoTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TaikoToken *TaikoTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Mint(&_TaikoToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TaikoToken *TaikoTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Mint(&_TaikoToken.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TaikoToken *TaikoTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TaikoToken *TaikoTokenSession) Pause() (*types.Transaction, error) {
	return _TaikoToken.Contract.Pause(&_TaikoToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TaikoToken *TaikoTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _TaikoToken.Contract.Pause(&_TaikoToken.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TaikoToken *TaikoTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TaikoToken *TaikoTokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TaikoToken.Contract.Permit(&_TaikoToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TaikoToken *TaikoTokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TaikoToken.Contract.Permit(&_TaikoToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoToken *TaikoTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoToken *TaikoTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoToken.Contract.RenounceOwnership(&_TaikoToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoToken *TaikoTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoToken.Contract.RenounceOwnership(&_TaikoToken.TransactOpts)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_TaikoToken *TaikoTokenTransactor) SetAddressManager(opts *bind.TransactOpts, newAddressManager common.Address) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "setAddressManager", newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_TaikoToken *TaikoTokenSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _TaikoToken.Contract.SetAddressManager(&_TaikoToken.TransactOpts, newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_TaikoToken *TaikoTokenTransactorSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _TaikoToken.Contract.SetAddressManager(&_TaikoToken.TransactOpts, newAddressManager)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_TaikoToken *TaikoTokenTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_TaikoToken *TaikoTokenSession) Snapshot() (*types.Transaction, error) {
	return _TaikoToken.Contract.Snapshot(&_TaikoToken.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_TaikoToken *TaikoTokenTransactorSession) Snapshot() (*types.Transaction, error) {
	return _TaikoToken.Contract.Snapshot(&_TaikoToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TaikoToken *TaikoTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TaikoToken *TaikoTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Transfer(&_TaikoToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TaikoToken *TaikoTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.Transfer(&_TaikoToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TaikoToken *TaikoTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TaikoToken *TaikoTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.TransferFrom(&_TaikoToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TaikoToken *TaikoTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaikoToken.Contract.TransferFrom(&_TaikoToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoToken *TaikoTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoToken *TaikoTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoToken.Contract.TransferOwnership(&_TaikoToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoToken *TaikoTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoToken.Contract.TransferOwnership(&_TaikoToken.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TaikoToken *TaikoTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TaikoToken *TaikoTokenSession) Unpause() (*types.Transaction, error) {
	return _TaikoToken.Contract.Unpause(&_TaikoToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TaikoToken *TaikoTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _TaikoToken.Contract.Unpause(&_TaikoToken.TransactOpts)
}

// TaikoTokenAddressManagerChangedIterator is returned from FilterAddressManagerChanged and is used to iterate over the raw logs and unpacked data for AddressManagerChanged events raised by the TaikoToken contract.
type TaikoTokenAddressManagerChangedIterator struct {
	Event *TaikoTokenAddressManagerChanged // Event containing the contract specifics and raw log

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
func (it *TaikoTokenAddressManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoTokenAddressManagerChanged)
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
		it.Event = new(TaikoTokenAddressManagerChanged)
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
func (it *TaikoTokenAddressManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoTokenAddressManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoTokenAddressManagerChanged represents a AddressManagerChanged event raised by the TaikoToken contract.
type TaikoTokenAddressManagerChanged struct {
	AddressManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressManagerChanged is a free log retrieval operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address indexed addressManager)
func (_TaikoToken *TaikoTokenFilterer) FilterAddressManagerChanged(opts *bind.FilterOpts, addressManager []common.Address) (*TaikoTokenAddressManagerChangedIterator, error) {

	var addressManagerRule []interface{}
	for _, addressManagerItem := range addressManager {
		addressManagerRule = append(addressManagerRule, addressManagerItem)
	}

	logs, sub, err := _TaikoToken.contract.FilterLogs(opts, "AddressManagerChanged", addressManagerRule)
	if err != nil {
		return nil, err
	}
	return &TaikoTokenAddressManagerChangedIterator{contract: _TaikoToken.contract, event: "AddressManagerChanged", logs: logs, sub: sub}, nil
}

// WatchAddressManagerChanged is a free log subscription operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address indexed addressManager)
func (_TaikoToken *TaikoTokenFilterer) WatchAddressManagerChanged(opts *bind.WatchOpts, sink chan<- *TaikoTokenAddressManagerChanged, addressManager []common.Address) (event.Subscription, error) {

	var addressManagerRule []interface{}
	for _, addressManagerItem := range addressManager {
		addressManagerRule = append(addressManagerRule, addressManagerItem)
	}

	logs, sub, err := _TaikoToken.contract.WatchLogs(opts, "AddressManagerChanged", addressManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoTokenAddressManagerChanged)
				if err := _TaikoToken.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
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
func (_TaikoToken *TaikoTokenFilterer) ParseAddressManagerChanged(log types.Log) (*TaikoTokenAddressManagerChanged, error) {
	event := new(TaikoTokenAddressManagerChanged)
	if err := _TaikoToken.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TaikoToken contract.
type TaikoTokenApprovalIterator struct {
	Event *TaikoTokenApproval // Event containing the contract specifics and raw log

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
func (it *TaikoTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoTokenApproval)
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
		it.Event = new(TaikoTokenApproval)
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
func (it *TaikoTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoTokenApproval represents a Approval event raised by the TaikoToken contract.
type TaikoTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TaikoToken *TaikoTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TaikoTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TaikoToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TaikoTokenApprovalIterator{contract: _TaikoToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TaikoToken *TaikoTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TaikoTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TaikoToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoTokenApproval)
				if err := _TaikoToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TaikoToken *TaikoTokenFilterer) ParseApproval(log types.Log) (*TaikoTokenApproval, error) {
	event := new(TaikoTokenApproval)
	if err := _TaikoToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoTokenDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the TaikoToken contract.
type TaikoTokenDelegateChangedIterator struct {
	Event *TaikoTokenDelegateChanged // Event containing the contract specifics and raw log

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
func (it *TaikoTokenDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoTokenDelegateChanged)
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
		it.Event = new(TaikoTokenDelegateChanged)
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
func (it *TaikoTokenDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoTokenDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoTokenDelegateChanged represents a DelegateChanged event raised by the TaikoToken contract.
type TaikoTokenDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_TaikoToken *TaikoTokenFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*TaikoTokenDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _TaikoToken.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &TaikoTokenDelegateChangedIterator{contract: _TaikoToken.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_TaikoToken *TaikoTokenFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *TaikoTokenDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _TaikoToken.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoTokenDelegateChanged)
				if err := _TaikoToken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_TaikoToken *TaikoTokenFilterer) ParseDelegateChanged(log types.Log) (*TaikoTokenDelegateChanged, error) {
	event := new(TaikoTokenDelegateChanged)
	if err := _TaikoToken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoTokenDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the TaikoToken contract.
type TaikoTokenDelegateVotesChangedIterator struct {
	Event *TaikoTokenDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *TaikoTokenDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoTokenDelegateVotesChanged)
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
		it.Event = new(TaikoTokenDelegateVotesChanged)
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
func (it *TaikoTokenDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoTokenDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoTokenDelegateVotesChanged represents a DelegateVotesChanged event raised by the TaikoToken contract.
type TaikoTokenDelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_TaikoToken *TaikoTokenFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*TaikoTokenDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _TaikoToken.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &TaikoTokenDelegateVotesChangedIterator{contract: _TaikoToken.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_TaikoToken *TaikoTokenFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *TaikoTokenDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _TaikoToken.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoTokenDelegateVotesChanged)
				if err := _TaikoToken.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_TaikoToken *TaikoTokenFilterer) ParseDelegateVotesChanged(log types.Log) (*TaikoTokenDelegateVotesChanged, error) {
	event := new(TaikoTokenDelegateVotesChanged)
	if err := _TaikoToken.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoTokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaikoToken contract.
type TaikoTokenInitializedIterator struct {
	Event *TaikoTokenInitialized // Event containing the contract specifics and raw log

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
func (it *TaikoTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoTokenInitialized)
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
		it.Event = new(TaikoTokenInitialized)
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
func (it *TaikoTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoTokenInitialized represents a Initialized event raised by the TaikoToken contract.
type TaikoTokenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoToken *TaikoTokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaikoTokenInitializedIterator, error) {

	logs, sub, err := _TaikoToken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaikoTokenInitializedIterator{contract: _TaikoToken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoToken *TaikoTokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaikoTokenInitialized) (event.Subscription, error) {

	logs, sub, err := _TaikoToken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoTokenInitialized)
				if err := _TaikoToken.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TaikoToken *TaikoTokenFilterer) ParseInitialized(log types.Log) (*TaikoTokenInitialized, error) {
	event := new(TaikoTokenInitialized)
	if err := _TaikoToken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaikoToken contract.
type TaikoTokenOwnershipTransferredIterator struct {
	Event *TaikoTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaikoTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoTokenOwnershipTransferred)
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
		it.Event = new(TaikoTokenOwnershipTransferred)
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
func (it *TaikoTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoTokenOwnershipTransferred represents a OwnershipTransferred event raised by the TaikoToken contract.
type TaikoTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoToken *TaikoTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaikoTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaikoTokenOwnershipTransferredIterator{contract: _TaikoToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoToken *TaikoTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaikoTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoTokenOwnershipTransferred)
				if err := _TaikoToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaikoToken *TaikoTokenFilterer) ParseOwnershipTransferred(log types.Log) (*TaikoTokenOwnershipTransferred, error) {
	event := new(TaikoTokenOwnershipTransferred)
	if err := _TaikoToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoTokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TaikoToken contract.
type TaikoTokenPausedIterator struct {
	Event *TaikoTokenPaused // Event containing the contract specifics and raw log

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
func (it *TaikoTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoTokenPaused)
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
		it.Event = new(TaikoTokenPaused)
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
func (it *TaikoTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoTokenPaused represents a Paused event raised by the TaikoToken contract.
type TaikoTokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TaikoToken *TaikoTokenFilterer) FilterPaused(opts *bind.FilterOpts) (*TaikoTokenPausedIterator, error) {

	logs, sub, err := _TaikoToken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TaikoTokenPausedIterator{contract: _TaikoToken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TaikoToken *TaikoTokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TaikoTokenPaused) (event.Subscription, error) {

	logs, sub, err := _TaikoToken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoTokenPaused)
				if err := _TaikoToken.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TaikoToken *TaikoTokenFilterer) ParsePaused(log types.Log) (*TaikoTokenPaused, error) {
	event := new(TaikoTokenPaused)
	if err := _TaikoToken.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoTokenSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the TaikoToken contract.
type TaikoTokenSnapshotIterator struct {
	Event *TaikoTokenSnapshot // Event containing the contract specifics and raw log

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
func (it *TaikoTokenSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoTokenSnapshot)
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
		it.Event = new(TaikoTokenSnapshot)
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
func (it *TaikoTokenSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoTokenSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoTokenSnapshot represents a Snapshot event raised by the TaikoToken contract.
type TaikoTokenSnapshot struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_TaikoToken *TaikoTokenFilterer) FilterSnapshot(opts *bind.FilterOpts) (*TaikoTokenSnapshotIterator, error) {

	logs, sub, err := _TaikoToken.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &TaikoTokenSnapshotIterator{contract: _TaikoToken.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_TaikoToken *TaikoTokenFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *TaikoTokenSnapshot) (event.Subscription, error) {

	logs, sub, err := _TaikoToken.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoTokenSnapshot)
				if err := _TaikoToken.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_TaikoToken *TaikoTokenFilterer) ParseSnapshot(log types.Log) (*TaikoTokenSnapshot, error) {
	event := new(TaikoTokenSnapshot)
	if err := _TaikoToken.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TaikoToken contract.
type TaikoTokenTransferIterator struct {
	Event *TaikoTokenTransfer // Event containing the contract specifics and raw log

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
func (it *TaikoTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoTokenTransfer)
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
		it.Event = new(TaikoTokenTransfer)
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
func (it *TaikoTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoTokenTransfer represents a Transfer event raised by the TaikoToken contract.
type TaikoTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TaikoToken *TaikoTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TaikoTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TaikoToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TaikoTokenTransferIterator{contract: _TaikoToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TaikoToken *TaikoTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TaikoTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TaikoToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoTokenTransfer)
				if err := _TaikoToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TaikoToken *TaikoTokenFilterer) ParseTransfer(log types.Log) (*TaikoTokenTransfer, error) {
	event := new(TaikoTokenTransfer)
	if err := _TaikoToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoTokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TaikoToken contract.
type TaikoTokenUnpausedIterator struct {
	Event *TaikoTokenUnpaused // Event containing the contract specifics and raw log

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
func (it *TaikoTokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoTokenUnpaused)
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
		it.Event = new(TaikoTokenUnpaused)
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
func (it *TaikoTokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoTokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoTokenUnpaused represents a Unpaused event raised by the TaikoToken contract.
type TaikoTokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TaikoToken *TaikoTokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TaikoTokenUnpausedIterator, error) {

	logs, sub, err := _TaikoToken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TaikoTokenUnpausedIterator{contract: _TaikoToken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TaikoToken *TaikoTokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TaikoTokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _TaikoToken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoTokenUnpaused)
				if err := _TaikoToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TaikoToken *TaikoTokenFilterer) ParseUnpaused(log types.Log) (*TaikoTokenUnpaused, error) {
	event := new(TaikoTokenUnpaused)
	if err := _TaikoToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
