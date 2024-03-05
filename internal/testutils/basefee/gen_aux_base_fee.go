// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basefee

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

// AuxBaseFeeMetaData contains all meta data concerning the AuxBaseFee contract.
var AuxBaseFeeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EIP1559_INVALID_PARAMS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasExcess_\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"basefeeAdjustmentQuotient\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"gasTargetPerL1Block\",\"type\":\"uint32\"}],\"name\":\"baseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"basefee_\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061059b8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063c82f638d1461002d575b5f80fd5b61004760048036038101906100429190610422565b61005d565b6040516100549190610481565b60405180910390f35b5f61007c848363ffffffff168560ff1661007791906104c7565b610085565b90509392505050565b5f8082036100bf576040517fc52de37200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81670de0b6b3a76400006100d385856100ef565b6100dd9190610535565b6100e79190610535565b905092915050565b5f8082670de0b6b3a76400008561010691906104c7565b6101109190610535565b9050680755bf798b4a1bf1e46fffffffffffffffffffffffffffffffff1681111561015457680755bf798b4a1bf1e46fffffffffffffffffffffffffffffffff1690505b61015d81610166565b91505092915050565b5f7ffffffffffffffffffffffffffffffffffffffffffffffffdb731c958f34d94c18213610196575f9050610377565b680755bf798b4a1bf1e582126101d8576040517f35278d1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6503782dace9d9604e83901b816101f2576101f1610508565b5b0591505f60606b8000000000000000000000006bb17217f7d1cf79abc9e3b398606086901b8161022557610224610508565b5b0501901d90506bb17217f7d1cf79abc9e3b3988102830392505f6c10fe68e7fd37d0007b713f7650840190506d02d16720577bd19bf614176fe9ea6060858302901d0190505f6d04a4fd9f2a8b96949216d2255a6c8583010390506e0587f503bb6ea29d25fcb7401964506060838302901d01905079d835ebba824c98fb31b83b2ca45c0000000000000000000000008582020190505f6c240c330e9fb2d9cbaf0fd5aafc860390506d0277594991cfc85f6e2461837cd96060878302901d0190506d1a521255e34f6a5061b25ef1c9c46060878302901d0390506db1bbb201f443cf962f1a1d3db4a56060878302901d0190506e02c72388d9f74f51a9331fed693f156060878302901d0390506e05180bb14799ab47a8a8cb2a527d576060878302901d01905080820594508360c30374029d9dc38563c32e5c2f6dc192ee70ef65f9978af38602901c9450505050505b919050565b5f80fd5b5f819050919050565b61039281610380565b811461039c575f80fd5b50565b5f813590506103ad81610389565b92915050565b5f60ff82169050919050565b6103c8816103b3565b81146103d2575f80fd5b50565b5f813590506103e3816103bf565b92915050565b5f63ffffffff82169050919050565b610401816103e9565b811461040b575f80fd5b50565b5f8135905061041c816103f8565b92915050565b5f805f606084860312156104395761043861037c565b5b5f6104468682870161039f565b9350506020610457868287016103d5565b92505060406104688682870161040e565b9150509250925092565b61047b81610380565b82525050565b5f6020820190506104945f830184610472565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6104d182610380565b91506104dc83610380565b92508282026104ea81610380565b915082820484148315176105015761050061049a565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61053f82610380565b915061054a83610380565b92508261055a57610559610508565b5b82820490509291505056fea2646970667358221220c9160a503fc937b1e6a83e2240c90e042cb64342a24c07f4222b2ddd6ae8f87564736f6c63430008180033",
}

// AuxBaseFeeABI is the input ABI used to generate the binding from.
// Deprecated: Use AuxBaseFeeMetaData.ABI instead.
var AuxBaseFeeABI = AuxBaseFeeMetaData.ABI

// AuxBaseFeeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuxBaseFeeMetaData.Bin instead.
var AuxBaseFeeBin = AuxBaseFeeMetaData.Bin

// DeployAuxBaseFee deploys a new Ethereum contract, binding an instance of AuxBaseFee to it.
func DeployAuxBaseFee(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AuxBaseFee, error) {
	parsed, err := AuxBaseFeeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuxBaseFeeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuxBaseFee{AuxBaseFeeCaller: AuxBaseFeeCaller{contract: contract}, AuxBaseFeeTransactor: AuxBaseFeeTransactor{contract: contract}, AuxBaseFeeFilterer: AuxBaseFeeFilterer{contract: contract}}, nil
}

// AuxBaseFee is an auto generated Go binding around an Ethereum contract.
type AuxBaseFee struct {
	AuxBaseFeeCaller     // Read-only binding to the contract
	AuxBaseFeeTransactor // Write-only binding to the contract
	AuxBaseFeeFilterer   // Log filterer for contract events
}

// AuxBaseFeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuxBaseFeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuxBaseFeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuxBaseFeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuxBaseFeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuxBaseFeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuxBaseFeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuxBaseFeeSession struct {
	Contract     *AuxBaseFee       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuxBaseFeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuxBaseFeeCallerSession struct {
	Contract *AuxBaseFeeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AuxBaseFeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuxBaseFeeTransactorSession struct {
	Contract     *AuxBaseFeeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AuxBaseFeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuxBaseFeeRaw struct {
	Contract *AuxBaseFee // Generic contract binding to access the raw methods on
}

// AuxBaseFeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuxBaseFeeCallerRaw struct {
	Contract *AuxBaseFeeCaller // Generic read-only contract binding to access the raw methods on
}

// AuxBaseFeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuxBaseFeeTransactorRaw struct {
	Contract *AuxBaseFeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuxBaseFee creates a new instance of AuxBaseFee, bound to a specific deployed contract.
func NewAuxBaseFee(address common.Address, backend bind.ContractBackend) (*AuxBaseFee, error) {
	contract, err := bindAuxBaseFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuxBaseFee{AuxBaseFeeCaller: AuxBaseFeeCaller{contract: contract}, AuxBaseFeeTransactor: AuxBaseFeeTransactor{contract: contract}, AuxBaseFeeFilterer: AuxBaseFeeFilterer{contract: contract}}, nil
}

// NewAuxBaseFeeCaller creates a new read-only instance of AuxBaseFee, bound to a specific deployed contract.
func NewAuxBaseFeeCaller(address common.Address, caller bind.ContractCaller) (*AuxBaseFeeCaller, error) {
	contract, err := bindAuxBaseFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuxBaseFeeCaller{contract: contract}, nil
}

// NewAuxBaseFeeTransactor creates a new write-only instance of AuxBaseFee, bound to a specific deployed contract.
func NewAuxBaseFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*AuxBaseFeeTransactor, error) {
	contract, err := bindAuxBaseFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuxBaseFeeTransactor{contract: contract}, nil
}

// NewAuxBaseFeeFilterer creates a new log filterer instance of AuxBaseFee, bound to a specific deployed contract.
func NewAuxBaseFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*AuxBaseFeeFilterer, error) {
	contract, err := bindAuxBaseFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuxBaseFeeFilterer{contract: contract}, nil
}

// bindAuxBaseFee binds a generic wrapper to an already deployed contract.
func bindAuxBaseFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuxBaseFeeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuxBaseFee *AuxBaseFeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuxBaseFee.Contract.AuxBaseFeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuxBaseFee *AuxBaseFeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuxBaseFee.Contract.AuxBaseFeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuxBaseFee *AuxBaseFeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuxBaseFee.Contract.AuxBaseFeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuxBaseFee *AuxBaseFeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuxBaseFee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuxBaseFee *AuxBaseFeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuxBaseFee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuxBaseFee *AuxBaseFeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuxBaseFee.Contract.contract.Transact(opts, method, params...)
}

// BaseFee is a free data retrieval call binding the contract method 0xc82f638d.
//
// Solidity: function baseFee(uint256 gasExcess_, uint8 basefeeAdjustmentQuotient, uint32 gasTargetPerL1Block) pure returns(uint256 basefee_)
func (_AuxBaseFee *AuxBaseFeeCaller) BaseFee(opts *bind.CallOpts, gasExcess_ *big.Int, basefeeAdjustmentQuotient uint8, gasTargetPerL1Block uint32) (*big.Int, error) {
	var out []interface{}
	err := _AuxBaseFee.contract.Call(opts, &out, "baseFee", gasExcess_, basefeeAdjustmentQuotient, gasTargetPerL1Block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseFee is a free data retrieval call binding the contract method 0xc82f638d.
//
// Solidity: function baseFee(uint256 gasExcess_, uint8 basefeeAdjustmentQuotient, uint32 gasTargetPerL1Block) pure returns(uint256 basefee_)
func (_AuxBaseFee *AuxBaseFeeSession) BaseFee(gasExcess_ *big.Int, basefeeAdjustmentQuotient uint8, gasTargetPerL1Block uint32) (*big.Int, error) {
	return _AuxBaseFee.Contract.BaseFee(&_AuxBaseFee.CallOpts, gasExcess_, basefeeAdjustmentQuotient, gasTargetPerL1Block)
}

// BaseFee is a free data retrieval call binding the contract method 0xc82f638d.
//
// Solidity: function baseFee(uint256 gasExcess_, uint8 basefeeAdjustmentQuotient, uint32 gasTargetPerL1Block) pure returns(uint256 basefee_)
func (_AuxBaseFee *AuxBaseFeeCallerSession) BaseFee(gasExcess_ *big.Int, basefeeAdjustmentQuotient uint8, gasTargetPerL1Block uint32) (*big.Int, error) {
	return _AuxBaseFee.Contract.BaseFee(&_AuxBaseFee.CallOpts, gasExcess_, basefeeAdjustmentQuotient, gasTargetPerL1Block)
}

// Lib1559MathMetaData contains all meta data concerning the Lib1559Math contract.
var Lib1559MathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EIP1559_INVALID_PARAMS\",\"type\":\"error\"}]",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220960da3e51a46826b2eb22395911a4f88abf5e55724f4bdf277fc83b864c731dc64736f6c63430008180033",
}

// Lib1559MathABI is the input ABI used to generate the binding from.
// Deprecated: Use Lib1559MathMetaData.ABI instead.
var Lib1559MathABI = Lib1559MathMetaData.ABI

// Lib1559MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Lib1559MathMetaData.Bin instead.
var Lib1559MathBin = Lib1559MathMetaData.Bin

// DeployLib1559Math deploys a new Ethereum contract, binding an instance of Lib1559Math to it.
func DeployLib1559Math(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lib1559Math, error) {
	parsed, err := Lib1559MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Lib1559MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lib1559Math{Lib1559MathCaller: Lib1559MathCaller{contract: contract}, Lib1559MathTransactor: Lib1559MathTransactor{contract: contract}, Lib1559MathFilterer: Lib1559MathFilterer{contract: contract}}, nil
}

// Lib1559Math is an auto generated Go binding around an Ethereum contract.
type Lib1559Math struct {
	Lib1559MathCaller     // Read-only binding to the contract
	Lib1559MathTransactor // Write-only binding to the contract
	Lib1559MathFilterer   // Log filterer for contract events
}

// Lib1559MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type Lib1559MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lib1559MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Lib1559MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lib1559MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Lib1559MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lib1559MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Lib1559MathSession struct {
	Contract     *Lib1559Math      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Lib1559MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Lib1559MathCallerSession struct {
	Contract *Lib1559MathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Lib1559MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Lib1559MathTransactorSession struct {
	Contract     *Lib1559MathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Lib1559MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type Lib1559MathRaw struct {
	Contract *Lib1559Math // Generic contract binding to access the raw methods on
}

// Lib1559MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Lib1559MathCallerRaw struct {
	Contract *Lib1559MathCaller // Generic read-only contract binding to access the raw methods on
}

// Lib1559MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Lib1559MathTransactorRaw struct {
	Contract *Lib1559MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLib1559Math creates a new instance of Lib1559Math, bound to a specific deployed contract.
func NewLib1559Math(address common.Address, backend bind.ContractBackend) (*Lib1559Math, error) {
	contract, err := bindLib1559Math(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lib1559Math{Lib1559MathCaller: Lib1559MathCaller{contract: contract}, Lib1559MathTransactor: Lib1559MathTransactor{contract: contract}, Lib1559MathFilterer: Lib1559MathFilterer{contract: contract}}, nil
}

// NewLib1559MathCaller creates a new read-only instance of Lib1559Math, bound to a specific deployed contract.
func NewLib1559MathCaller(address common.Address, caller bind.ContractCaller) (*Lib1559MathCaller, error) {
	contract, err := bindLib1559Math(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Lib1559MathCaller{contract: contract}, nil
}

// NewLib1559MathTransactor creates a new write-only instance of Lib1559Math, bound to a specific deployed contract.
func NewLib1559MathTransactor(address common.Address, transactor bind.ContractTransactor) (*Lib1559MathTransactor, error) {
	contract, err := bindLib1559Math(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Lib1559MathTransactor{contract: contract}, nil
}

// NewLib1559MathFilterer creates a new log filterer instance of Lib1559Math, bound to a specific deployed contract.
func NewLib1559MathFilterer(address common.Address, filterer bind.ContractFilterer) (*Lib1559MathFilterer, error) {
	contract, err := bindLib1559Math(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Lib1559MathFilterer{contract: contract}, nil
}

// bindLib1559Math binds a generic wrapper to an already deployed contract.
func bindLib1559Math(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Lib1559MathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lib1559Math *Lib1559MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lib1559Math.Contract.Lib1559MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lib1559Math *Lib1559MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lib1559Math.Contract.Lib1559MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lib1559Math *Lib1559MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lib1559Math.Contract.Lib1559MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lib1559Math *Lib1559MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lib1559Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lib1559Math *Lib1559MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lib1559Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lib1559Math *Lib1559MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lib1559Math.Contract.contract.Transact(opts, method, params...)
}

// LibFixedPointMathMetaData contains all meta data concerning the LibFixedPointMath contract.
var LibFixedPointMathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MAX_EXP_INPUT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SCALING_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61014261004e600b8282823980515f1a607314610042577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061003f575f3560e01c8063bb0a0aa514610043578063ef4cadc514610061575b5f80fd5b61004b61007f565b60405161005891906100c2565b60405180910390f35b61006961008c565b60405161007691906100f3565b60405180910390f35b680755bf798b4a1bf1e481565b670de0b6b3a764000081565b5f6fffffffffffffffffffffffffffffffff82169050919050565b6100bc81610098565b82525050565b5f6020820190506100d55f8301846100b3565b92915050565b5f819050919050565b6100ed816100db565b82525050565b5f6020820190506101065f8301846100e4565b9291505056fea26469706673582212208e75079756512a1458dfcde94a74ab0163eadfdac7848a4c95b2505cd04fd19464736f6c63430008180033",
}

// LibFixedPointMathABI is the input ABI used to generate the binding from.
// Deprecated: Use LibFixedPointMathMetaData.ABI instead.
var LibFixedPointMathABI = LibFixedPointMathMetaData.ABI

// LibFixedPointMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LibFixedPointMathMetaData.Bin instead.
var LibFixedPointMathBin = LibFixedPointMathMetaData.Bin

// DeployLibFixedPointMath deploys a new Ethereum contract, binding an instance of LibFixedPointMath to it.
func DeployLibFixedPointMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibFixedPointMath, error) {
	parsed, err := LibFixedPointMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LibFixedPointMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibFixedPointMath{LibFixedPointMathCaller: LibFixedPointMathCaller{contract: contract}, LibFixedPointMathTransactor: LibFixedPointMathTransactor{contract: contract}, LibFixedPointMathFilterer: LibFixedPointMathFilterer{contract: contract}}, nil
}

// LibFixedPointMath is an auto generated Go binding around an Ethereum contract.
type LibFixedPointMath struct {
	LibFixedPointMathCaller     // Read-only binding to the contract
	LibFixedPointMathTransactor // Write-only binding to the contract
	LibFixedPointMathFilterer   // Log filterer for contract events
}

// LibFixedPointMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibFixedPointMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibFixedPointMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibFixedPointMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibFixedPointMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibFixedPointMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibFixedPointMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibFixedPointMathSession struct {
	Contract     *LibFixedPointMath // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LibFixedPointMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibFixedPointMathCallerSession struct {
	Contract *LibFixedPointMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LibFixedPointMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibFixedPointMathTransactorSession struct {
	Contract     *LibFixedPointMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LibFixedPointMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibFixedPointMathRaw struct {
	Contract *LibFixedPointMath // Generic contract binding to access the raw methods on
}

// LibFixedPointMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibFixedPointMathCallerRaw struct {
	Contract *LibFixedPointMathCaller // Generic read-only contract binding to access the raw methods on
}

// LibFixedPointMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibFixedPointMathTransactorRaw struct {
	Contract *LibFixedPointMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibFixedPointMath creates a new instance of LibFixedPointMath, bound to a specific deployed contract.
func NewLibFixedPointMath(address common.Address, backend bind.ContractBackend) (*LibFixedPointMath, error) {
	contract, err := bindLibFixedPointMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibFixedPointMath{LibFixedPointMathCaller: LibFixedPointMathCaller{contract: contract}, LibFixedPointMathTransactor: LibFixedPointMathTransactor{contract: contract}, LibFixedPointMathFilterer: LibFixedPointMathFilterer{contract: contract}}, nil
}

// NewLibFixedPointMathCaller creates a new read-only instance of LibFixedPointMath, bound to a specific deployed contract.
func NewLibFixedPointMathCaller(address common.Address, caller bind.ContractCaller) (*LibFixedPointMathCaller, error) {
	contract, err := bindLibFixedPointMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibFixedPointMathCaller{contract: contract}, nil
}

// NewLibFixedPointMathTransactor creates a new write-only instance of LibFixedPointMath, bound to a specific deployed contract.
func NewLibFixedPointMathTransactor(address common.Address, transactor bind.ContractTransactor) (*LibFixedPointMathTransactor, error) {
	contract, err := bindLibFixedPointMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibFixedPointMathTransactor{contract: contract}, nil
}

// NewLibFixedPointMathFilterer creates a new log filterer instance of LibFixedPointMath, bound to a specific deployed contract.
func NewLibFixedPointMathFilterer(address common.Address, filterer bind.ContractFilterer) (*LibFixedPointMathFilterer, error) {
	contract, err := bindLibFixedPointMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibFixedPointMathFilterer{contract: contract}, nil
}

// bindLibFixedPointMath binds a generic wrapper to an already deployed contract.
func bindLibFixedPointMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LibFixedPointMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibFixedPointMath *LibFixedPointMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibFixedPointMath.Contract.LibFixedPointMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibFixedPointMath *LibFixedPointMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibFixedPointMath.Contract.LibFixedPointMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibFixedPointMath *LibFixedPointMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibFixedPointMath.Contract.LibFixedPointMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibFixedPointMath *LibFixedPointMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibFixedPointMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibFixedPointMath *LibFixedPointMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibFixedPointMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibFixedPointMath *LibFixedPointMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibFixedPointMath.Contract.contract.Transact(opts, method, params...)
}

// MAXEXPINPUT is a free data retrieval call binding the contract method 0xbb0a0aa5.
//
// Solidity: function MAX_EXP_INPUT() view returns(uint128)
func (_LibFixedPointMath *LibFixedPointMathCaller) MAXEXPINPUT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibFixedPointMath.contract.Call(opts, &out, "MAX_EXP_INPUT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXEXPINPUT is a free data retrieval call binding the contract method 0xbb0a0aa5.
//
// Solidity: function MAX_EXP_INPUT() view returns(uint128)
func (_LibFixedPointMath *LibFixedPointMathSession) MAXEXPINPUT() (*big.Int, error) {
	return _LibFixedPointMath.Contract.MAXEXPINPUT(&_LibFixedPointMath.CallOpts)
}

// MAXEXPINPUT is a free data retrieval call binding the contract method 0xbb0a0aa5.
//
// Solidity: function MAX_EXP_INPUT() view returns(uint128)
func (_LibFixedPointMath *LibFixedPointMathCallerSession) MAXEXPINPUT() (*big.Int, error) {
	return _LibFixedPointMath.Contract.MAXEXPINPUT(&_LibFixedPointMath.CallOpts)
}

// SCALINGFACTOR is a free data retrieval call binding the contract method 0xef4cadc5.
//
// Solidity: function SCALING_FACTOR() view returns(uint256)
func (_LibFixedPointMath *LibFixedPointMathCaller) SCALINGFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LibFixedPointMath.contract.Call(opts, &out, "SCALING_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCALINGFACTOR is a free data retrieval call binding the contract method 0xef4cadc5.
//
// Solidity: function SCALING_FACTOR() view returns(uint256)
func (_LibFixedPointMath *LibFixedPointMathSession) SCALINGFACTOR() (*big.Int, error) {
	return _LibFixedPointMath.Contract.SCALINGFACTOR(&_LibFixedPointMath.CallOpts)
}

// SCALINGFACTOR is a free data retrieval call binding the contract method 0xef4cadc5.
//
// Solidity: function SCALING_FACTOR() view returns(uint256)
func (_LibFixedPointMath *LibFixedPointMathCallerSession) SCALINGFACTOR() (*big.Int, error) {
	return _LibFixedPointMath.Contract.SCALINGFACTOR(&_LibFixedPointMath.CallOpts)
}
