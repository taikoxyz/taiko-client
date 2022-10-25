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

// LibDataBlockMetadata is an auto generated low-level Go binding around an user-defined struct.
type LibDataBlockMetadata struct {
	Id          *big.Int
	L1Height    *big.Int
	L1Hash      [32]byte
	Beneficiary common.Address
	GasLimit    uint64
	Timestamp   uint64
	TxListHash  [32]byte
	MixHash     [32]byte
	ExtraData   []byte
}

// LibDataProposedBlock is an auto generated low-level Go binding around an user-defined struct.
type LibDataProposedBlock struct {
	MetaHash [32]byte
}

// TaikoL1ClientMetaData contains all meta data concerning the TaikoL1Client contract.
var TaikoL1ClientMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validSince\",\"type\":\"uint256\"}],\"name\":\"BlockCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structLibData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"BlockProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcHash\",\"type\":\"bytes32\"}],\"name\":\"HeaderSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"name\":\"commitBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"name\":\"finalizeBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"name\":\"getCommitHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConstants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getProposedBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"metaHash\",\"type\":\"bytes32\"}],\"internalType\":\"structLibData.ProposedBlock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateVariables\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getSyncedHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isCommitValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"proposeBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"proveBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"proveBlockInvalid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"k\",\"type\":\"uint8\"}],\"name\":\"signWithGoldFinger\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"genesisHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nextBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"parentTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestFinalizedHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestFinalizedId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
	parsed, err := abi.JSON(strings.NewReader(TaikoL1ClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// GetCommitHeight is a free data retrieval call binding the contract method 0xb60c400a.
//
// Solidity: function getCommitHeight(bytes32 commitHash) view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientCaller) GetCommitHeight(opts *bind.CallOpts, commitHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getCommitHeight", commitHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommitHeight is a free data retrieval call binding the contract method 0xb60c400a.
//
// Solidity: function getCommitHeight(bytes32 commitHash) view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientSession) GetCommitHeight(commitHash [32]byte) (*big.Int, error) {
	return _TaikoL1Client.Contract.GetCommitHeight(&_TaikoL1Client.CallOpts, commitHash)
}

// GetCommitHeight is a free data retrieval call binding the contract method 0xb60c400a.
//
// Solidity: function getCommitHeight(bytes32 commitHash) view returns(uint256)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetCommitHeight(commitHash [32]byte) (*big.Int, error) {
	return _TaikoL1Client.Contract.GetCommitHeight(&_TaikoL1Client.CallOpts, commitHash)
}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, bytes32, uint256, uint256, uint256, bytes4, bytes32)
func (_TaikoL1Client *TaikoL1ClientCaller) GetConstants(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, [32]byte, *big.Int, *big.Int, *big.Int, [4]byte, [32]byte, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getConstants")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([32]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([4]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new([32]byte)).(*[32]byte)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	out9 := *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	out10 := *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	out11 := *abi.ConvertType(out[11], new([4]byte)).(*[4]byte)
	out12 := *abi.ConvertType(out[12], new([32]byte)).(*[32]byte)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, out11, out12, err

}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, bytes32, uint256, uint256, uint256, bytes4, bytes32)
func (_TaikoL1Client *TaikoL1ClientSession) GetConstants() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, [32]byte, *big.Int, *big.Int, *big.Int, [4]byte, [32]byte, error) {
	return _TaikoL1Client.Contract.GetConstants(&_TaikoL1Client.CallOpts)
}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, bytes32, uint256, uint256, uint256, bytes4, bytes32)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetConstants() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, [32]byte, *big.Int, *big.Int, *big.Int, [4]byte, [32]byte, error) {
	return _TaikoL1Client.Contract.GetConstants(&_TaikoL1Client.CallOpts)
}

// GetProposedBlock is a free data retrieval call binding the contract method 0x8972b10c.
//
// Solidity: function getProposedBlock(uint256 id) view returns((bytes32))
func (_TaikoL1Client *TaikoL1ClientCaller) GetProposedBlock(opts *bind.CallOpts, id *big.Int) (LibDataProposedBlock, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getProposedBlock", id)

	if err != nil {
		return *new(LibDataProposedBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(LibDataProposedBlock)).(*LibDataProposedBlock)

	return out0, err

}

// GetProposedBlock is a free data retrieval call binding the contract method 0x8972b10c.
//
// Solidity: function getProposedBlock(uint256 id) view returns((bytes32))
func (_TaikoL1Client *TaikoL1ClientSession) GetProposedBlock(id *big.Int) (LibDataProposedBlock, error) {
	return _TaikoL1Client.Contract.GetProposedBlock(&_TaikoL1Client.CallOpts, id)
}

// GetProposedBlock is a free data retrieval call binding the contract method 0x8972b10c.
//
// Solidity: function getProposedBlock(uint256 id) view returns((bytes32))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetProposedBlock(id *big.Int) (LibDataProposedBlock, error) {
	return _TaikoL1Client.Contract.GetProposedBlock(&_TaikoL1Client.CallOpts, id)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns(uint256, uint64, uint64, uint64)
func (_TaikoL1Client *TaikoL1ClientCaller) GetStateVariables(opts *bind.CallOpts) (*big.Int, uint64, uint64, uint64, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getStateVariables")

	if err != nil {
		return *new(*big.Int), *new(uint64), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return out0, out1, out2, out3, err

}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns(uint256, uint64, uint64, uint64)
func (_TaikoL1Client *TaikoL1ClientSession) GetStateVariables() (*big.Int, uint64, uint64, uint64, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns(uint256, uint64, uint64, uint64)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetStateVariables() (*big.Int, uint64, uint64, uint64, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCaller) GetSyncedHeader(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getSyncedHeader", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientSession) GetSyncedHeader(number *big.Int) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetSyncedHeader(&_TaikoL1Client.CallOpts, number)
}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetSyncedHeader(number *big.Int) ([32]byte, error) {
	return _TaikoL1Client.Contract.GetSyncedHeader(&_TaikoL1Client.CallOpts, number)
}

// IsCommitValid is a free data retrieval call binding the contract method 0xea818740.
//
// Solidity: function isCommitValid(bytes32 hash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) IsCommitValid(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "isCommitValid", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommitValid is a free data retrieval call binding the contract method 0xea818740.
//
// Solidity: function isCommitValid(bytes32 hash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) IsCommitValid(hash [32]byte) (bool, error) {
	return _TaikoL1Client.Contract.IsCommitValid(&_TaikoL1Client.CallOpts, hash)
}

// IsCommitValid is a free data retrieval call binding the contract method 0xea818740.
//
// Solidity: function isCommitValid(bytes32 hash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) IsCommitValid(hash [32]byte) (bool, error) {
	return _TaikoL1Client.Contract.IsCommitValid(&_TaikoL1Client.CallOpts, hash)
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

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string name) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCaller) Resolve(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "resolve", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string name) view returns(address)
func (_TaikoL1Client *TaikoL1ClientSession) Resolve(name string) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve(&_TaikoL1Client.CallOpts, name)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string name) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Resolve(name string) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve(&_TaikoL1Client.CallOpts, name)
}

// Resolve0 is a free data retrieval call binding the contract method 0xf16c7934.
//
// Solidity: function resolve(uint256 chainId, string name) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCaller) Resolve0(opts *bind.CallOpts, chainId *big.Int, name string) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "resolve0", chainId, name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xf16c7934.
//
// Solidity: function resolve(uint256 chainId, string name) view returns(address)
func (_TaikoL1Client *TaikoL1ClientSession) Resolve0(chainId *big.Int, name string) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve0(&_TaikoL1Client.CallOpts, chainId, name)
}

// Resolve0 is a free data retrieval call binding the contract method 0xf16c7934.
//
// Solidity: function resolve(uint256 chainId, string name) view returns(address)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Resolve0(chainId *big.Int, name string) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve0(&_TaikoL1Client.CallOpts, chainId, name)
}

// SignWithGoldFinger is a free data retrieval call binding the contract method 0x80aa0158.
//
// Solidity: function signWithGoldFinger(bytes32 hash, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL1Client *TaikoL1ClientCaller) SignWithGoldFinger(opts *bind.CallOpts, hash [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "signWithGoldFinger", hash, k)

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

// SignWithGoldFinger is a free data retrieval call binding the contract method 0x80aa0158.
//
// Solidity: function signWithGoldFinger(bytes32 hash, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL1Client *TaikoL1ClientSession) SignWithGoldFinger(hash [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _TaikoL1Client.Contract.SignWithGoldFinger(&_TaikoL1Client.CallOpts, hash, k)
}

// SignWithGoldFinger is a free data retrieval call binding the contract method 0x80aa0158.
//
// Solidity: function signWithGoldFinger(bytes32 hash, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL1Client *TaikoL1ClientCallerSession) SignWithGoldFinger(hash [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _TaikoL1Client.Contract.SignWithGoldFinger(&_TaikoL1Client.CallOpts, hash, k)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256 genesisHeight, uint64 nextBlockId, uint64 parentTimestamp, uint64 latestFinalizedHeight, uint64 latestFinalizedId)
func (_TaikoL1Client *TaikoL1ClientCaller) State(opts *bind.CallOpts) (struct {
	GenesisHeight         *big.Int
	NextBlockId           uint64
	ParentTimestamp       uint64
	LatestFinalizedHeight uint64
	LatestFinalizedId     uint64
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		GenesisHeight         *big.Int
		NextBlockId           uint64
		ParentTimestamp       uint64
		LatestFinalizedHeight uint64
		LatestFinalizedId     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GenesisHeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NextBlockId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ParentTimestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.LatestFinalizedHeight = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.LatestFinalizedId = *abi.ConvertType(out[4], new(uint64)).(*uint64)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256 genesisHeight, uint64 nextBlockId, uint64 parentTimestamp, uint64 latestFinalizedHeight, uint64 latestFinalizedId)
func (_TaikoL1Client *TaikoL1ClientSession) State() (struct {
	GenesisHeight         *big.Int
	NextBlockId           uint64
	ParentTimestamp       uint64
	LatestFinalizedHeight uint64
	LatestFinalizedId     uint64
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256 genesisHeight, uint64 nextBlockId, uint64 parentTimestamp, uint64 latestFinalizedHeight, uint64 latestFinalizedId)
func (_TaikoL1Client *TaikoL1ClientCallerSession) State() (struct {
	GenesisHeight         *big.Int
	NextBlockId           uint64
	ParentTimestamp       uint64
	LatestFinalizedHeight uint64
	LatestFinalizedId     uint64
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x74f6c562.
//
// Solidity: function commitBlock(bytes32 commitHash) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) CommitBlock(opts *bind.TransactOpts, commitHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "commitBlock", commitHash)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x74f6c562.
//
// Solidity: function commitBlock(bytes32 commitHash) returns()
func (_TaikoL1Client *TaikoL1ClientSession) CommitBlock(commitHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.CommitBlock(&_TaikoL1Client.TransactOpts, commitHash)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x74f6c562.
//
// Solidity: function commitBlock(bytes32 commitHash) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) CommitBlock(commitHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.CommitBlock(&_TaikoL1Client.TransactOpts, commitHash)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x64eed0c8.
//
// Solidity: function finalizeBlocks(uint256 maxBlocks) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) FinalizeBlocks(opts *bind.TransactOpts, maxBlocks *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "finalizeBlocks", maxBlocks)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x64eed0c8.
//
// Solidity: function finalizeBlocks(uint256 maxBlocks) returns()
func (_TaikoL1Client *TaikoL1ClientSession) FinalizeBlocks(maxBlocks *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.FinalizeBlocks(&_TaikoL1Client.TransactOpts, maxBlocks)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x64eed0c8.
//
// Solidity: function finalizeBlocks(uint256 maxBlocks) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) FinalizeBlocks(maxBlocks *big.Int) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.FinalizeBlocks(&_TaikoL1Client.TransactOpts, maxBlocks)
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

// ProposeBlock is a paid mutator transaction binding the contract method 0xa043dbdf.
//
// Solidity: function proposeBlock(bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) ProposeBlock(opts *bind.TransactOpts, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proposeBlock", inputs)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xa043dbdf.
//
// Solidity: function proposeBlock(bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientSession) ProposeBlock(inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, inputs)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xa043dbdf.
//
// Solidity: function proposeBlock(bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProposeBlock(inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, inputs)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x8ed7b3be.
//
// Solidity: function proveBlock(uint256 blockIndex, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) ProveBlock(opts *bind.TransactOpts, blockIndex *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proveBlock", blockIndex, inputs)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x8ed7b3be.
//
// Solidity: function proveBlock(uint256 blockIndex, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientSession) ProveBlock(blockIndex *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlock(&_TaikoL1Client.TransactOpts, blockIndex, inputs)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x8ed7b3be.
//
// Solidity: function proveBlock(uint256 blockIndex, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProveBlock(blockIndex *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlock(&_TaikoL1Client.TransactOpts, blockIndex, inputs)
}

// ProveBlockInvalid is a paid mutator transaction binding the contract method 0xa279cec7.
//
// Solidity: function proveBlockInvalid(uint256 blockIndex, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) ProveBlockInvalid(opts *bind.TransactOpts, blockIndex *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proveBlockInvalid", blockIndex, inputs)
}

// ProveBlockInvalid is a paid mutator transaction binding the contract method 0xa279cec7.
//
// Solidity: function proveBlockInvalid(uint256 blockIndex, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientSession) ProveBlockInvalid(blockIndex *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlockInvalid(&_TaikoL1Client.TransactOpts, blockIndex, inputs)
}

// ProveBlockInvalid is a paid mutator transaction binding the contract method 0xa279cec7.
//
// Solidity: function proveBlockInvalid(uint256 blockIndex, bytes[] inputs) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProveBlockInvalid(blockIndex *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlockInvalid(&_TaikoL1Client.TransactOpts, blockIndex, inputs)
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

// TaikoL1ClientBlockCommittedIterator is returned from FilterBlockCommitted and is used to iterate over the raw logs and unpacked data for BlockCommitted events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockCommittedIterator struct {
	Event *TaikoL1ClientBlockCommitted // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientBlockCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockCommitted)
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
		it.Event = new(TaikoL1ClientBlockCommitted)
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
func (it *TaikoL1ClientBlockCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockCommitted represents a BlockCommitted event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockCommitted struct {
	Hash       [32]byte
	ValidSince *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockCommitted is a free log retrieval operation binding the contract event 0x7617cbebd9ef36093e4020cd06e961fee8dcabf3702906570a571690f40afaa9.
//
// Solidity: event BlockCommitted(bytes32 hash, uint256 validSince)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockCommitted(opts *bind.FilterOpts) (*TaikoL1ClientBlockCommittedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockCommitted")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockCommittedIterator{contract: _TaikoL1Client.contract, event: "BlockCommitted", logs: logs, sub: sub}, nil
}

// WatchBlockCommitted is a free log subscription operation binding the contract event 0x7617cbebd9ef36093e4020cd06e961fee8dcabf3702906570a571690f40afaa9.
//
// Solidity: event BlockCommitted(bytes32 hash, uint256 validSince)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockCommitted(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockCommitted) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockCommitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockCommitted)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockCommitted", log); err != nil {
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

// ParseBlockCommitted is a log parse operation binding the contract event 0x7617cbebd9ef36093e4020cd06e961fee8dcabf3702906570a571690f40afaa9.
//
// Solidity: event BlockCommitted(bytes32 hash, uint256 validSince)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockCommitted(log types.Log) (*TaikoL1ClientBlockCommitted, error) {
	event := new(TaikoL1ClientBlockCommitted)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockFinalizedIterator struct {
	Event *TaikoL1ClientBlockFinalized // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientBlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockFinalized)
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
		it.Event = new(TaikoL1ClientBlockFinalized)
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
func (it *TaikoL1ClientBlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockFinalized represents a BlockFinalized event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockFinalized struct {
	Id        *big.Int
	BlockHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xf2c535759092d16e9334a11dd9b52eca543f1d9cca5ba9d16c472aef009de432.
//
// Solidity: event BlockFinalized(uint256 indexed id, bytes32 blockHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockFinalized(opts *bind.FilterOpts, id []*big.Int) (*TaikoL1ClientBlockFinalizedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockFinalized", idRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockFinalizedIterator{contract: _TaikoL1Client.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xf2c535759092d16e9334a11dd9b52eca543f1d9cca5ba9d16c472aef009de432.
//
// Solidity: event BlockFinalized(uint256 indexed id, bytes32 blockHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockFinalized, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockFinalized", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockFinalized)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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

// ParseBlockFinalized is a log parse operation binding the contract event 0xf2c535759092d16e9334a11dd9b52eca543f1d9cca5ba9d16c472aef009de432.
//
// Solidity: event BlockFinalized(uint256 indexed id, bytes32 blockHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockFinalized(log types.Log) (*TaikoL1ClientBlockFinalized, error) {
	event := new(TaikoL1ClientBlockFinalized)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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
	Id   *big.Int
	Meta LibDataBlockMetadata
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0x43ddf03c6701dc2f6c4fe864c9dabc9011e2d898c5e2ef046ca1bfdd6cfd2242.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint256,uint256,bytes32,address,uint64,uint64,bytes32,bytes32,bytes) meta)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockProposed(opts *bind.FilterOpts, id []*big.Int) (*TaikoL1ClientBlockProposedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockProposed", idRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockProposedIterator{contract: _TaikoL1Client.contract, event: "BlockProposed", logs: logs, sub: sub}, nil
}

// WatchBlockProposed is a free log subscription operation binding the contract event 0x43ddf03c6701dc2f6c4fe864c9dabc9011e2d898c5e2ef046ca1bfdd6cfd2242.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint256,uint256,bytes32,address,uint64,uint64,bytes32,bytes32,bytes) meta)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockProposed(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockProposed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockProposed", idRule)
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

// ParseBlockProposed is a log parse operation binding the contract event 0x43ddf03c6701dc2f6c4fe864c9dabc9011e2d898c5e2ef046ca1bfdd6cfd2242.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint256,uint256,bytes32,address,uint64,uint64,bytes32,bytes32,bytes) meta)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockProposed(log types.Log) (*TaikoL1ClientBlockProposed, error) {
	event := new(TaikoL1ClientBlockProposed)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockProvenIterator is returned from FilterBlockProven and is used to iterate over the raw logs and unpacked data for BlockProven events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProvenIterator struct {
	Event *TaikoL1ClientBlockProven // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientBlockProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockProven)
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
		it.Event = new(TaikoL1ClientBlockProven)
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
func (it *TaikoL1ClientBlockProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockProven represents a BlockProven event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProven struct {
	Id         *big.Int
	ParentHash [32]byte
	BlockHash  [32]byte
	Timestamp  uint64
	ProvenAt   uint64
	Prover     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockProven is a free log retrieval operation binding the contract event 0xf66e10e397da1b4c87429bd76e1d3b85fd417c38c5d43c8ea3e2ec45935bada0.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, uint64 timestamp, uint64 provenAt, address prover)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockProven(opts *bind.FilterOpts, id []*big.Int) (*TaikoL1ClientBlockProvenIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockProven", idRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockProvenIterator{contract: _TaikoL1Client.contract, event: "BlockProven", logs: logs, sub: sub}, nil
}

// WatchBlockProven is a free log subscription operation binding the contract event 0xf66e10e397da1b4c87429bd76e1d3b85fd417c38c5d43c8ea3e2ec45935bada0.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, uint64 timestamp, uint64 provenAt, address prover)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockProven(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockProven, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockProven", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockProven)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProven", log); err != nil {
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

// ParseBlockProven is a log parse operation binding the contract event 0xf66e10e397da1b4c87429bd76e1d3b85fd417c38c5d43c8ea3e2ec45935bada0.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, uint64 timestamp, uint64 provenAt, address prover)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockProven(log types.Log) (*TaikoL1ClientBlockProven, error) {
	event := new(TaikoL1ClientBlockProven)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientHeaderSyncedIterator is returned from FilterHeaderSynced and is used to iterate over the raw logs and unpacked data for HeaderSynced events raised by the TaikoL1Client contract.
type TaikoL1ClientHeaderSyncedIterator struct {
	Event *TaikoL1ClientHeaderSynced // Event containing the contract specifics and raw log

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
func (it *TaikoL1ClientHeaderSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientHeaderSynced)
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
		it.Event = new(TaikoL1ClientHeaderSynced)
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
func (it *TaikoL1ClientHeaderSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientHeaderSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientHeaderSynced represents a HeaderSynced event raised by the TaikoL1Client contract.
type TaikoL1ClientHeaderSynced struct {
	Height    *big.Int
	SrcHeight *big.Int
	SrcHash   [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHeaderSynced is a free log retrieval operation binding the contract event 0x930c750845026c7bb04c0e3d9111d512b4c86981713c4944a35a10a4a7a854f3.
//
// Solidity: event HeaderSynced(uint256 indexed height, uint256 indexed srcHeight, bytes32 srcHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterHeaderSynced(opts *bind.FilterOpts, height []*big.Int, srcHeight []*big.Int) (*TaikoL1ClientHeaderSyncedIterator, error) {

	var heightRule []interface{}
	for _, heightItem := range height {
		heightRule = append(heightRule, heightItem)
	}
	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "HeaderSynced", heightRule, srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientHeaderSyncedIterator{contract: _TaikoL1Client.contract, event: "HeaderSynced", logs: logs, sub: sub}, nil
}

// WatchHeaderSynced is a free log subscription operation binding the contract event 0x930c750845026c7bb04c0e3d9111d512b4c86981713c4944a35a10a4a7a854f3.
//
// Solidity: event HeaderSynced(uint256 indexed height, uint256 indexed srcHeight, bytes32 srcHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchHeaderSynced(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientHeaderSynced, height []*big.Int, srcHeight []*big.Int) (event.Subscription, error) {

	var heightRule []interface{}
	for _, heightItem := range height {
		heightRule = append(heightRule, heightItem)
	}
	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "HeaderSynced", heightRule, srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientHeaderSynced)
				if err := _TaikoL1Client.contract.UnpackLog(event, "HeaderSynced", log); err != nil {
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
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseHeaderSynced(log types.Log) (*TaikoL1ClientHeaderSynced, error) {
	event := new(TaikoL1ClientHeaderSynced)
	if err := _TaikoL1Client.contract.UnpackLog(event, "HeaderSynced", log); err != nil {
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
