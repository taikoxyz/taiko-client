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
	ABI: "[{\"inputs\":[],\"name\":\"EIP1559_INVALID_PARAMS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_gasTargetPerL1Block\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_adjustmentQuotient\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gasExcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gasIssuance\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"calc1559BaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"basefee_\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasExcess_\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506106fb8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c80635d958a851461002d575b5f80fd5b610047600480360381019061004291906104cf565b61005e565b60405161005592919061056d565b60405180910390f35b5f805f8363ffffffff168667ffffffffffffffff1661007d91906105c1565b90508467ffffffffffffffff1681116100975760016100ae565b8467ffffffffffffffff16816100ad91906105f4565b5b90506100cd67ffffffffffffffff80168261011090919063ffffffff16565b91506100f78267ffffffffffffffff168963ffffffff168960ff166100f29190610627565b610128565b92505f830361010557600192505b509550959350505050565b5f81831161011e5782610120565b815b905092915050565b5f808203610162576040517fc52de37200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81670de0b6b3a76400006101768585610192565b6101809190610695565b61018a9190610695565b905092915050565b5f8082670de0b6b3a7640000856101a99190610627565b6101b39190610695565b9050680755bf798b4a1bf1e46fffffffffffffffffffffffffffffffff168111156101f757680755bf798b4a1bf1e46fffffffffffffffffffffffffffffffff1690505b61020081610209565b91505092915050565b5f7ffffffffffffffffffffffffffffffffffffffffffffffffdb731c958f34d94c18213610239575f905061041a565b680755bf798b4a1bf1e5821261027b576040517f35278d1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6503782dace9d9604e83901b8161029557610294610668565b5b0591505f60606b8000000000000000000000006bb17217f7d1cf79abc9e3b398606086901b816102c8576102c7610668565b5b0501901d90506bb17217f7d1cf79abc9e3b3988102830392505f6c10fe68e7fd37d0007b713f7650840190506d02d16720577bd19bf614176fe9ea6060858302901d0190505f6d04a4fd9f2a8b96949216d2255a6c8583010390506e0587f503bb6ea29d25fcb7401964506060838302901d01905079d835ebba824c98fb31b83b2ca45c0000000000000000000000008582020190505f6c240c330e9fb2d9cbaf0fd5aafc860390506d0277594991cfc85f6e2461837cd96060878302901d0190506d1a521255e34f6a5061b25ef1c9c46060878302901d0390506db1bbb201f443cf962f1a1d3db4a56060878302901d0190506e02c72388d9f74f51a9331fed693f156060878302901d0390506e05180bb14799ab47a8a8cb2a527d576060878302901d01905080820594508360c30374029d9dc38563c32e5c2f6dc192ee70ef65f9978af38602901c9450505050505b919050565b5f80fd5b5f63ffffffff82169050919050565b61043b81610423565b8114610445575f80fd5b50565b5f8135905061045681610432565b92915050565b5f60ff82169050919050565b6104718161045c565b811461047b575f80fd5b50565b5f8135905061048c81610468565b92915050565b5f67ffffffffffffffff82169050919050565b6104ae81610492565b81146104b8575f80fd5b50565b5f813590506104c9816104a5565b92915050565b5f805f805f60a086880312156104e8576104e761041f565b5b5f6104f588828901610448565b95505060206105068882890161047e565b9450506040610517888289016104bb565b9350506060610528888289016104bb565b925050608061053988828901610448565b9150509295509295909350565b5f819050919050565b61055881610546565b82525050565b61056781610492565b82525050565b5f6040820190506105805f83018561054f565b61058d602083018461055e565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6105cb82610546565b91506105d683610546565b92508282019050808211156105ee576105ed610594565b5b92915050565b5f6105fe82610546565b915061060983610546565b925082820390508181111561062157610620610594565b5b92915050565b5f61063182610546565b915061063c83610546565b925082820261064a81610546565b9150828204841483151761066157610660610594565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61069f82610546565b91506106aa83610546565b9250826106ba576106b9610668565b5b82820490509291505056fea26469706673582212205b99fe84c6a2e5cd8720f09e509edde79d128ed074ef72c1e98981acf78bcda764736f6c63430008180033",
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

// Calc1559BaseFee is a free data retrieval call binding the contract method 0x5d958a85.
//
// Solidity: function calc1559BaseFee(uint32 _gasTargetPerL1Block, uint8 _adjustmentQuotient, uint64 _gasExcess, uint64 _gasIssuance, uint32 _parentGasUsed) pure returns(uint256 basefee_, uint64 gasExcess_)
func (_AuxBaseFee *AuxBaseFeeCaller) Calc1559BaseFee(opts *bind.CallOpts, _gasTargetPerL1Block uint32, _adjustmentQuotient uint8, _gasExcess uint64, _gasIssuance uint64, _parentGasUsed uint32) (struct {
	Basefee   *big.Int
	GasExcess uint64
}, error) {
	var out []interface{}
	err := _AuxBaseFee.contract.Call(opts, &out, "calc1559BaseFee", _gasTargetPerL1Block, _adjustmentQuotient, _gasExcess, _gasIssuance, _parentGasUsed)

	outstruct := new(struct {
		Basefee   *big.Int
		GasExcess uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Basefee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasExcess = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// Calc1559BaseFee is a free data retrieval call binding the contract method 0x5d958a85.
//
// Solidity: function calc1559BaseFee(uint32 _gasTargetPerL1Block, uint8 _adjustmentQuotient, uint64 _gasExcess, uint64 _gasIssuance, uint32 _parentGasUsed) pure returns(uint256 basefee_, uint64 gasExcess_)
func (_AuxBaseFee *AuxBaseFeeSession) Calc1559BaseFee(_gasTargetPerL1Block uint32, _adjustmentQuotient uint8, _gasExcess uint64, _gasIssuance uint64, _parentGasUsed uint32) (struct {
	Basefee   *big.Int
	GasExcess uint64
}, error) {
	return _AuxBaseFee.Contract.Calc1559BaseFee(&_AuxBaseFee.CallOpts, _gasTargetPerL1Block, _adjustmentQuotient, _gasExcess, _gasIssuance, _parentGasUsed)
}

// Calc1559BaseFee is a free data retrieval call binding the contract method 0x5d958a85.
//
// Solidity: function calc1559BaseFee(uint32 _gasTargetPerL1Block, uint8 _adjustmentQuotient, uint64 _gasExcess, uint64 _gasIssuance, uint32 _parentGasUsed) pure returns(uint256 basefee_, uint64 gasExcess_)
func (_AuxBaseFee *AuxBaseFeeCallerSession) Calc1559BaseFee(_gasTargetPerL1Block uint32, _adjustmentQuotient uint8, _gasExcess uint64, _gasIssuance uint64, _parentGasUsed uint32) (struct {
	Basefee   *big.Int
	GasExcess uint64
}, error) {
	return _AuxBaseFee.Contract.Calc1559BaseFee(&_AuxBaseFee.CallOpts, _gasTargetPerL1Block, _adjustmentQuotient, _gasExcess, _gasIssuance, _parentGasUsed)
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

// LibMathMetaData contains all meta data concerning the LibMath contract.
var LibMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212207b2d6a4d256fd37413ff5bb2963c8b217029784aa6bc50bd5915658915fe79c864736f6c63430008180033",
}

// LibMathABI is the input ABI used to generate the binding from.
// Deprecated: Use LibMathMetaData.ABI instead.
var LibMathABI = LibMathMetaData.ABI

// LibMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LibMathMetaData.Bin instead.
var LibMathBin = LibMathMetaData.Bin

// DeployLibMath deploys a new Ethereum contract, binding an instance of LibMath to it.
func DeployLibMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibMath, error) {
	parsed, err := LibMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LibMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibMath{LibMathCaller: LibMathCaller{contract: contract}, LibMathTransactor: LibMathTransactor{contract: contract}, LibMathFilterer: LibMathFilterer{contract: contract}}, nil
}

// LibMath is an auto generated Go binding around an Ethereum contract.
type LibMath struct {
	LibMathCaller     // Read-only binding to the contract
	LibMathTransactor // Write-only binding to the contract
	LibMathFilterer   // Log filterer for contract events
}

// LibMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibMathSession struct {
	Contract     *LibMath          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibMathCallerSession struct {
	Contract *LibMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LibMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibMathTransactorSession struct {
	Contract     *LibMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LibMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibMathRaw struct {
	Contract *LibMath // Generic contract binding to access the raw methods on
}

// LibMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibMathCallerRaw struct {
	Contract *LibMathCaller // Generic read-only contract binding to access the raw methods on
}

// LibMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibMathTransactorRaw struct {
	Contract *LibMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibMath creates a new instance of LibMath, bound to a specific deployed contract.
func NewLibMath(address common.Address, backend bind.ContractBackend) (*LibMath, error) {
	contract, err := bindLibMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibMath{LibMathCaller: LibMathCaller{contract: contract}, LibMathTransactor: LibMathTransactor{contract: contract}, LibMathFilterer: LibMathFilterer{contract: contract}}, nil
}

// NewLibMathCaller creates a new read-only instance of LibMath, bound to a specific deployed contract.
func NewLibMathCaller(address common.Address, caller bind.ContractCaller) (*LibMathCaller, error) {
	contract, err := bindLibMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibMathCaller{contract: contract}, nil
}

// NewLibMathTransactor creates a new write-only instance of LibMath, bound to a specific deployed contract.
func NewLibMathTransactor(address common.Address, transactor bind.ContractTransactor) (*LibMathTransactor, error) {
	contract, err := bindLibMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibMathTransactor{contract: contract}, nil
}

// NewLibMathFilterer creates a new log filterer instance of LibMath, bound to a specific deployed contract.
func NewLibMathFilterer(address common.Address, filterer bind.ContractFilterer) (*LibMathFilterer, error) {
	contract, err := bindLibMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibMathFilterer{contract: contract}, nil
}

// bindLibMath binds a generic wrapper to an already deployed contract.
func bindLibMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LibMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibMath *LibMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibMath.Contract.LibMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibMath *LibMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibMath.Contract.LibMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibMath *LibMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibMath.Contract.LibMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibMath *LibMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibMath *LibMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibMath *LibMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibMath.Contract.contract.Transact(opts, method, params...)
}
