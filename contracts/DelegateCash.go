// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package delegatecash

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

// IDelegationRegistryContractDelegation is an auto generated low-level Go binding around an user-defined struct.
type IDelegationRegistryContractDelegation struct {
	Contract common.Address
	Delegate common.Address
}

// IDelegationRegistryDelegationInfo is an auto generated low-level Go binding around an user-defined struct.
type IDelegationRegistryDelegationInfo struct {
	Type     uint8
	Vault    common.Address
	Delegate common.Address
	Contract common.Address
	TokenId  *big.Int
}

// IDelegationRegistryTokenDelegation is an auto generated low-level Go binding around an user-defined struct.
type IDelegationRegistryTokenDelegation struct {
	Contract common.Address
	TokenId  *big.Int
	Delegate common.Address
}

// DelegatecashMetaData contains all meta data concerning the Delegatecash contract.
var DelegatecashMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"DelegateForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"DelegateForContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"DelegateForToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"RevokeAllDelegates\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"RevokeDelegate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"checkDelegateForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"name\":\"checkDelegateForContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"checkDelegateForToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"delegateForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"delegateForContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"delegateForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getContractLevelDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"internalType\":\"structIDelegationRegistry.ContractDelegation[]\",\"name\":\"contractDelegations\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getDelegatesForAll\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"delegates\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"name\":\"getDelegatesForContract\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"delegates\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDelegatesForToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"delegates\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"getDelegationsByDelegate\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIDelegationRegistry.DelegationType\",\"name\":\"type_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIDelegationRegistry.DelegationInfo[]\",\"name\":\"info\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getTokenLevelDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"internalType\":\"structIDelegationRegistry.TokenDelegation[]\",\"name\":\"tokenDelegations\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeAllDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"revokeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"revokeSelf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DelegatecashABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegatecashMetaData.ABI instead.
var DelegatecashABI = DelegatecashMetaData.ABI

// Delegatecash is an auto generated Go binding around an Ethereum contract.
type Delegatecash struct {
	DelegatecashCaller     // Read-only binding to the contract
	DelegatecashTransactor // Write-only binding to the contract
	DelegatecashFilterer   // Log filterer for contract events
}

// DelegatecashCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegatecashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatecashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegatecashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatecashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegatecashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatecashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegatecashSession struct {
	Contract     *Delegatecash     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegatecashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegatecashCallerSession struct {
	Contract *DelegatecashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DelegatecashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegatecashTransactorSession struct {
	Contract     *DelegatecashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DelegatecashRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegatecashRaw struct {
	Contract *Delegatecash // Generic contract binding to access the raw methods on
}

// DelegatecashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegatecashCallerRaw struct {
	Contract *DelegatecashCaller // Generic read-only contract binding to access the raw methods on
}

// DelegatecashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegatecashTransactorRaw struct {
	Contract *DelegatecashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegatecash creates a new instance of Delegatecash, bound to a specific deployed contract.
func NewDelegatecash(address common.Address, backend bind.ContractBackend) (*Delegatecash, error) {
	contract, err := bindDelegatecash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Delegatecash{DelegatecashCaller: DelegatecashCaller{contract: contract}, DelegatecashTransactor: DelegatecashTransactor{contract: contract}, DelegatecashFilterer: DelegatecashFilterer{contract: contract}}, nil
}

// NewDelegatecashCaller creates a new read-only instance of Delegatecash, bound to a specific deployed contract.
func NewDelegatecashCaller(address common.Address, caller bind.ContractCaller) (*DelegatecashCaller, error) {
	contract, err := bindDelegatecash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatecashCaller{contract: contract}, nil
}

// NewDelegatecashTransactor creates a new write-only instance of Delegatecash, bound to a specific deployed contract.
func NewDelegatecashTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegatecashTransactor, error) {
	contract, err := bindDelegatecash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatecashTransactor{contract: contract}, nil
}

// NewDelegatecashFilterer creates a new log filterer instance of Delegatecash, bound to a specific deployed contract.
func NewDelegatecashFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegatecashFilterer, error) {
	contract, err := bindDelegatecash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegatecashFilterer{contract: contract}, nil
}

// bindDelegatecash binds a generic wrapper to an already deployed contract.
func bindDelegatecash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelegatecashMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegatecash *DelegatecashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delegatecash.Contract.DelegatecashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegatecash *DelegatecashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegatecash.Contract.DelegatecashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegatecash *DelegatecashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegatecash.Contract.DelegatecashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegatecash *DelegatecashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delegatecash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegatecash *DelegatecashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegatecash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegatecash *DelegatecashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegatecash.Contract.contract.Transact(opts, method, params...)
}

// CheckDelegateForAll is a free data retrieval call binding the contract method 0x9c395bc2.
//
// Solidity: function checkDelegateForAll(address delegate, address vault) view returns(bool)
func (_Delegatecash *DelegatecashCaller) CheckDelegateForAll(opts *bind.CallOpts, delegate common.Address, vault common.Address) (bool, error) {
	var out []interface{}
	err := _Delegatecash.contract.Call(opts, &out, "checkDelegateForAll", delegate, vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDelegateForAll is a free data retrieval call binding the contract method 0x9c395bc2.
//
// Solidity: function checkDelegateForAll(address delegate, address vault) view returns(bool)
func (_Delegatecash *DelegatecashSession) CheckDelegateForAll(delegate common.Address, vault common.Address) (bool, error) {
	return _Delegatecash.Contract.CheckDelegateForAll(&_Delegatecash.CallOpts, delegate, vault)
}

// CheckDelegateForAll is a free data retrieval call binding the contract method 0x9c395bc2.
//
// Solidity: function checkDelegateForAll(address delegate, address vault) view returns(bool)
func (_Delegatecash *DelegatecashCallerSession) CheckDelegateForAll(delegate common.Address, vault common.Address) (bool, error) {
	return _Delegatecash.Contract.CheckDelegateForAll(&_Delegatecash.CallOpts, delegate, vault)
}

// CheckDelegateForContract is a free data retrieval call binding the contract method 0x90c9a2d0.
//
// Solidity: function checkDelegateForContract(address delegate, address vault, address contract_) view returns(bool)
func (_Delegatecash *DelegatecashCaller) CheckDelegateForContract(opts *bind.CallOpts, delegate common.Address, vault common.Address, contract_ common.Address) (bool, error) {
	var out []interface{}
	err := _Delegatecash.contract.Call(opts, &out, "checkDelegateForContract", delegate, vault, contract_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDelegateForContract is a free data retrieval call binding the contract method 0x90c9a2d0.
//
// Solidity: function checkDelegateForContract(address delegate, address vault, address contract_) view returns(bool)
func (_Delegatecash *DelegatecashSession) CheckDelegateForContract(delegate common.Address, vault common.Address, contract_ common.Address) (bool, error) {
	return _Delegatecash.Contract.CheckDelegateForContract(&_Delegatecash.CallOpts, delegate, vault, contract_)
}

// CheckDelegateForContract is a free data retrieval call binding the contract method 0x90c9a2d0.
//
// Solidity: function checkDelegateForContract(address delegate, address vault, address contract_) view returns(bool)
func (_Delegatecash *DelegatecashCallerSession) CheckDelegateForContract(delegate common.Address, vault common.Address, contract_ common.Address) (bool, error) {
	return _Delegatecash.Contract.CheckDelegateForContract(&_Delegatecash.CallOpts, delegate, vault, contract_)
}

// CheckDelegateForToken is a free data retrieval call binding the contract method 0xaba69cf8.
//
// Solidity: function checkDelegateForToken(address delegate, address vault, address contract_, uint256 tokenId) view returns(bool)
func (_Delegatecash *DelegatecashCaller) CheckDelegateForToken(opts *bind.CallOpts, delegate common.Address, vault common.Address, contract_ common.Address, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Delegatecash.contract.Call(opts, &out, "checkDelegateForToken", delegate, vault, contract_, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDelegateForToken is a free data retrieval call binding the contract method 0xaba69cf8.
//
// Solidity: function checkDelegateForToken(address delegate, address vault, address contract_, uint256 tokenId) view returns(bool)
func (_Delegatecash *DelegatecashSession) CheckDelegateForToken(delegate common.Address, vault common.Address, contract_ common.Address, tokenId *big.Int) (bool, error) {
	return _Delegatecash.Contract.CheckDelegateForToken(&_Delegatecash.CallOpts, delegate, vault, contract_, tokenId)
}

// CheckDelegateForToken is a free data retrieval call binding the contract method 0xaba69cf8.
//
// Solidity: function checkDelegateForToken(address delegate, address vault, address contract_, uint256 tokenId) view returns(bool)
func (_Delegatecash *DelegatecashCallerSession) CheckDelegateForToken(delegate common.Address, vault common.Address, contract_ common.Address, tokenId *big.Int) (bool, error) {
	return _Delegatecash.Contract.CheckDelegateForToken(&_Delegatecash.CallOpts, delegate, vault, contract_, tokenId)
}

// GetContractLevelDelegations is a free data retrieval call binding the contract method 0xf956cf94.
//
// Solidity: function getContractLevelDelegations(address vault) view returns((address,address)[] contractDelegations)
func (_Delegatecash *DelegatecashCaller) GetContractLevelDelegations(opts *bind.CallOpts, vault common.Address) ([]IDelegationRegistryContractDelegation, error) {
	var out []interface{}
	err := _Delegatecash.contract.Call(opts, &out, "getContractLevelDelegations", vault)

	if err != nil {
		return *new([]IDelegationRegistryContractDelegation), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelegationRegistryContractDelegation)).(*[]IDelegationRegistryContractDelegation)

	return out0, err

}

// GetContractLevelDelegations is a free data retrieval call binding the contract method 0xf956cf94.
//
// Solidity: function getContractLevelDelegations(address vault) view returns((address,address)[] contractDelegations)
func (_Delegatecash *DelegatecashSession) GetContractLevelDelegations(vault common.Address) ([]IDelegationRegistryContractDelegation, error) {
	return _Delegatecash.Contract.GetContractLevelDelegations(&_Delegatecash.CallOpts, vault)
}

// GetContractLevelDelegations is a free data retrieval call binding the contract method 0xf956cf94.
//
// Solidity: function getContractLevelDelegations(address vault) view returns((address,address)[] contractDelegations)
func (_Delegatecash *DelegatecashCallerSession) GetContractLevelDelegations(vault common.Address) ([]IDelegationRegistryContractDelegation, error) {
	return _Delegatecash.Contract.GetContractLevelDelegations(&_Delegatecash.CallOpts, vault)
}

// GetDelegatesForAll is a free data retrieval call binding the contract method 0x1b61f675.
//
// Solidity: function getDelegatesForAll(address vault) view returns(address[] delegates)
func (_Delegatecash *DelegatecashCaller) GetDelegatesForAll(opts *bind.CallOpts, vault common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Delegatecash.contract.Call(opts, &out, "getDelegatesForAll", vault)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDelegatesForAll is a free data retrieval call binding the contract method 0x1b61f675.
//
// Solidity: function getDelegatesForAll(address vault) view returns(address[] delegates)
func (_Delegatecash *DelegatecashSession) GetDelegatesForAll(vault common.Address) ([]common.Address, error) {
	return _Delegatecash.Contract.GetDelegatesForAll(&_Delegatecash.CallOpts, vault)
}

// GetDelegatesForAll is a free data retrieval call binding the contract method 0x1b61f675.
//
// Solidity: function getDelegatesForAll(address vault) view returns(address[] delegates)
func (_Delegatecash *DelegatecashCallerSession) GetDelegatesForAll(vault common.Address) ([]common.Address, error) {
	return _Delegatecash.Contract.GetDelegatesForAll(&_Delegatecash.CallOpts, vault)
}

// GetDelegatesForContract is a free data retrieval call binding the contract method 0xed4b878e.
//
// Solidity: function getDelegatesForContract(address vault, address contract_) view returns(address[] delegates)
func (_Delegatecash *DelegatecashCaller) GetDelegatesForContract(opts *bind.CallOpts, vault common.Address, contract_ common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Delegatecash.contract.Call(opts, &out, "getDelegatesForContract", vault, contract_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDelegatesForContract is a free data retrieval call binding the contract method 0xed4b878e.
//
// Solidity: function getDelegatesForContract(address vault, address contract_) view returns(address[] delegates)
func (_Delegatecash *DelegatecashSession) GetDelegatesForContract(vault common.Address, contract_ common.Address) ([]common.Address, error) {
	return _Delegatecash.Contract.GetDelegatesForContract(&_Delegatecash.CallOpts, vault, contract_)
}

// GetDelegatesForContract is a free data retrieval call binding the contract method 0xed4b878e.
//
// Solidity: function getDelegatesForContract(address vault, address contract_) view returns(address[] delegates)
func (_Delegatecash *DelegatecashCallerSession) GetDelegatesForContract(vault common.Address, contract_ common.Address) ([]common.Address, error) {
	return _Delegatecash.Contract.GetDelegatesForContract(&_Delegatecash.CallOpts, vault, contract_)
}

// GetDelegatesForToken is a free data retrieval call binding the contract method 0x1221156b.
//
// Solidity: function getDelegatesForToken(address vault, address contract_, uint256 tokenId) view returns(address[] delegates)
func (_Delegatecash *DelegatecashCaller) GetDelegatesForToken(opts *bind.CallOpts, vault common.Address, contract_ common.Address, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Delegatecash.contract.Call(opts, &out, "getDelegatesForToken", vault, contract_, tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDelegatesForToken is a free data retrieval call binding the contract method 0x1221156b.
//
// Solidity: function getDelegatesForToken(address vault, address contract_, uint256 tokenId) view returns(address[] delegates)
func (_Delegatecash *DelegatecashSession) GetDelegatesForToken(vault common.Address, contract_ common.Address, tokenId *big.Int) ([]common.Address, error) {
	return _Delegatecash.Contract.GetDelegatesForToken(&_Delegatecash.CallOpts, vault, contract_, tokenId)
}

// GetDelegatesForToken is a free data retrieval call binding the contract method 0x1221156b.
//
// Solidity: function getDelegatesForToken(address vault, address contract_, uint256 tokenId) view returns(address[] delegates)
func (_Delegatecash *DelegatecashCallerSession) GetDelegatesForToken(vault common.Address, contract_ common.Address, tokenId *big.Int) ([]common.Address, error) {
	return _Delegatecash.Contract.GetDelegatesForToken(&_Delegatecash.CallOpts, vault, contract_, tokenId)
}

// GetDelegationsByDelegate is a free data retrieval call binding the contract method 0x4fc69282.
//
// Solidity: function getDelegationsByDelegate(address delegate) view returns((uint8,address,address,address,uint256)[] info)
func (_Delegatecash *DelegatecashCaller) GetDelegationsByDelegate(opts *bind.CallOpts, delegate common.Address) ([]IDelegationRegistryDelegationInfo, error) {
	var out []interface{}
	err := _Delegatecash.contract.Call(opts, &out, "getDelegationsByDelegate", delegate)

	if err != nil {
		return *new([]IDelegationRegistryDelegationInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelegationRegistryDelegationInfo)).(*[]IDelegationRegistryDelegationInfo)

	return out0, err

}

// GetDelegationsByDelegate is a free data retrieval call binding the contract method 0x4fc69282.
//
// Solidity: function getDelegationsByDelegate(address delegate) view returns((uint8,address,address,address,uint256)[] info)
func (_Delegatecash *DelegatecashSession) GetDelegationsByDelegate(delegate common.Address) ([]IDelegationRegistryDelegationInfo, error) {
	return _Delegatecash.Contract.GetDelegationsByDelegate(&_Delegatecash.CallOpts, delegate)
}

// GetDelegationsByDelegate is a free data retrieval call binding the contract method 0x4fc69282.
//
// Solidity: function getDelegationsByDelegate(address delegate) view returns((uint8,address,address,address,uint256)[] info)
func (_Delegatecash *DelegatecashCallerSession) GetDelegationsByDelegate(delegate common.Address) ([]IDelegationRegistryDelegationInfo, error) {
	return _Delegatecash.Contract.GetDelegationsByDelegate(&_Delegatecash.CallOpts, delegate)
}

// GetTokenLevelDelegations is a free data retrieval call binding the contract method 0x6f007d87.
//
// Solidity: function getTokenLevelDelegations(address vault) view returns((address,uint256,address)[] tokenDelegations)
func (_Delegatecash *DelegatecashCaller) GetTokenLevelDelegations(opts *bind.CallOpts, vault common.Address) ([]IDelegationRegistryTokenDelegation, error) {
	var out []interface{}
	err := _Delegatecash.contract.Call(opts, &out, "getTokenLevelDelegations", vault)

	if err != nil {
		return *new([]IDelegationRegistryTokenDelegation), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelegationRegistryTokenDelegation)).(*[]IDelegationRegistryTokenDelegation)

	return out0, err

}

// GetTokenLevelDelegations is a free data retrieval call binding the contract method 0x6f007d87.
//
// Solidity: function getTokenLevelDelegations(address vault) view returns((address,uint256,address)[] tokenDelegations)
func (_Delegatecash *DelegatecashSession) GetTokenLevelDelegations(vault common.Address) ([]IDelegationRegistryTokenDelegation, error) {
	return _Delegatecash.Contract.GetTokenLevelDelegations(&_Delegatecash.CallOpts, vault)
}

// GetTokenLevelDelegations is a free data retrieval call binding the contract method 0x6f007d87.
//
// Solidity: function getTokenLevelDelegations(address vault) view returns((address,uint256,address)[] tokenDelegations)
func (_Delegatecash *DelegatecashCallerSession) GetTokenLevelDelegations(vault common.Address) ([]IDelegationRegistryTokenDelegation, error) {
	return _Delegatecash.Contract.GetTokenLevelDelegations(&_Delegatecash.CallOpts, vault)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Delegatecash *DelegatecashCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Delegatecash.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Delegatecash *DelegatecashSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Delegatecash.Contract.SupportsInterface(&_Delegatecash.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Delegatecash *DelegatecashCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Delegatecash.Contract.SupportsInterface(&_Delegatecash.CallOpts, interfaceId)
}

// DelegateForAll is a paid mutator transaction binding the contract method 0x685ee3e8.
//
// Solidity: function delegateForAll(address delegate, bool value) returns()
func (_Delegatecash *DelegatecashTransactor) DelegateForAll(opts *bind.TransactOpts, delegate common.Address, value bool) (*types.Transaction, error) {
	return _Delegatecash.contract.Transact(opts, "delegateForAll", delegate, value)
}

// DelegateForAll is a paid mutator transaction binding the contract method 0x685ee3e8.
//
// Solidity: function delegateForAll(address delegate, bool value) returns()
func (_Delegatecash *DelegatecashSession) DelegateForAll(delegate common.Address, value bool) (*types.Transaction, error) {
	return _Delegatecash.Contract.DelegateForAll(&_Delegatecash.TransactOpts, delegate, value)
}

// DelegateForAll is a paid mutator transaction binding the contract method 0x685ee3e8.
//
// Solidity: function delegateForAll(address delegate, bool value) returns()
func (_Delegatecash *DelegatecashTransactorSession) DelegateForAll(delegate common.Address, value bool) (*types.Transaction, error) {
	return _Delegatecash.Contract.DelegateForAll(&_Delegatecash.TransactOpts, delegate, value)
}

// DelegateForContract is a paid mutator transaction binding the contract method 0x49c95d29.
//
// Solidity: function delegateForContract(address delegate, address contract_, bool value) returns()
func (_Delegatecash *DelegatecashTransactor) DelegateForContract(opts *bind.TransactOpts, delegate common.Address, contract_ common.Address, value bool) (*types.Transaction, error) {
	return _Delegatecash.contract.Transact(opts, "delegateForContract", delegate, contract_, value)
}

// DelegateForContract is a paid mutator transaction binding the contract method 0x49c95d29.
//
// Solidity: function delegateForContract(address delegate, address contract_, bool value) returns()
func (_Delegatecash *DelegatecashSession) DelegateForContract(delegate common.Address, contract_ common.Address, value bool) (*types.Transaction, error) {
	return _Delegatecash.Contract.DelegateForContract(&_Delegatecash.TransactOpts, delegate, contract_, value)
}

// DelegateForContract is a paid mutator transaction binding the contract method 0x49c95d29.
//
// Solidity: function delegateForContract(address delegate, address contract_, bool value) returns()
func (_Delegatecash *DelegatecashTransactorSession) DelegateForContract(delegate common.Address, contract_ common.Address, value bool) (*types.Transaction, error) {
	return _Delegatecash.Contract.DelegateForContract(&_Delegatecash.TransactOpts, delegate, contract_, value)
}

// DelegateForToken is a paid mutator transaction binding the contract method 0x537a5c3d.
//
// Solidity: function delegateForToken(address delegate, address contract_, uint256 tokenId, bool value) returns()
func (_Delegatecash *DelegatecashTransactor) DelegateForToken(opts *bind.TransactOpts, delegate common.Address, contract_ common.Address, tokenId *big.Int, value bool) (*types.Transaction, error) {
	return _Delegatecash.contract.Transact(opts, "delegateForToken", delegate, contract_, tokenId, value)
}

// DelegateForToken is a paid mutator transaction binding the contract method 0x537a5c3d.
//
// Solidity: function delegateForToken(address delegate, address contract_, uint256 tokenId, bool value) returns()
func (_Delegatecash *DelegatecashSession) DelegateForToken(delegate common.Address, contract_ common.Address, tokenId *big.Int, value bool) (*types.Transaction, error) {
	return _Delegatecash.Contract.DelegateForToken(&_Delegatecash.TransactOpts, delegate, contract_, tokenId, value)
}

// DelegateForToken is a paid mutator transaction binding the contract method 0x537a5c3d.
//
// Solidity: function delegateForToken(address delegate, address contract_, uint256 tokenId, bool value) returns()
func (_Delegatecash *DelegatecashTransactorSession) DelegateForToken(delegate common.Address, contract_ common.Address, tokenId *big.Int, value bool) (*types.Transaction, error) {
	return _Delegatecash.Contract.DelegateForToken(&_Delegatecash.TransactOpts, delegate, contract_, tokenId, value)
}

// RevokeAllDelegates is a paid mutator transaction binding the contract method 0x36137872.
//
// Solidity: function revokeAllDelegates() returns()
func (_Delegatecash *DelegatecashTransactor) RevokeAllDelegates(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegatecash.contract.Transact(opts, "revokeAllDelegates")
}

// RevokeAllDelegates is a paid mutator transaction binding the contract method 0x36137872.
//
// Solidity: function revokeAllDelegates() returns()
func (_Delegatecash *DelegatecashSession) RevokeAllDelegates() (*types.Transaction, error) {
	return _Delegatecash.Contract.RevokeAllDelegates(&_Delegatecash.TransactOpts)
}

// RevokeAllDelegates is a paid mutator transaction binding the contract method 0x36137872.
//
// Solidity: function revokeAllDelegates() returns()
func (_Delegatecash *DelegatecashTransactorSession) RevokeAllDelegates() (*types.Transaction, error) {
	return _Delegatecash.Contract.RevokeAllDelegates(&_Delegatecash.TransactOpts)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0xfa352c00.
//
// Solidity: function revokeDelegate(address delegate) returns()
func (_Delegatecash *DelegatecashTransactor) RevokeDelegate(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _Delegatecash.contract.Transact(opts, "revokeDelegate", delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0xfa352c00.
//
// Solidity: function revokeDelegate(address delegate) returns()
func (_Delegatecash *DelegatecashSession) RevokeDelegate(delegate common.Address) (*types.Transaction, error) {
	return _Delegatecash.Contract.RevokeDelegate(&_Delegatecash.TransactOpts, delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0xfa352c00.
//
// Solidity: function revokeDelegate(address delegate) returns()
func (_Delegatecash *DelegatecashTransactorSession) RevokeDelegate(delegate common.Address) (*types.Transaction, error) {
	return _Delegatecash.Contract.RevokeDelegate(&_Delegatecash.TransactOpts, delegate)
}

// RevokeSelf is a paid mutator transaction binding the contract method 0x219044b0.
//
// Solidity: function revokeSelf(address vault) returns()
func (_Delegatecash *DelegatecashTransactor) RevokeSelf(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Delegatecash.contract.Transact(opts, "revokeSelf", vault)
}

// RevokeSelf is a paid mutator transaction binding the contract method 0x219044b0.
//
// Solidity: function revokeSelf(address vault) returns()
func (_Delegatecash *DelegatecashSession) RevokeSelf(vault common.Address) (*types.Transaction, error) {
	return _Delegatecash.Contract.RevokeSelf(&_Delegatecash.TransactOpts, vault)
}

// RevokeSelf is a paid mutator transaction binding the contract method 0x219044b0.
//
// Solidity: function revokeSelf(address vault) returns()
func (_Delegatecash *DelegatecashTransactorSession) RevokeSelf(vault common.Address) (*types.Transaction, error) {
	return _Delegatecash.Contract.RevokeSelf(&_Delegatecash.TransactOpts, vault)
}

// DelegatecashDelegateForAllIterator is returned from FilterDelegateForAll and is used to iterate over the raw logs and unpacked data for DelegateForAll events raised by the Delegatecash contract.
type DelegatecashDelegateForAllIterator struct {
	Event *DelegatecashDelegateForAll // Event containing the contract specifics and raw log

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
func (it *DelegatecashDelegateForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatecashDelegateForAll)
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
		it.Event = new(DelegatecashDelegateForAll)
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
func (it *DelegatecashDelegateForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatecashDelegateForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatecashDelegateForAll represents a DelegateForAll event raised by the Delegatecash contract.
type DelegatecashDelegateForAll struct {
	Vault    common.Address
	Delegate common.Address
	Value    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateForAll is a free log retrieval operation binding the contract event 0x58781eab4a0743ab1c285a238be846a235f06cdb5b968030573a635e5f8c92fa.
//
// Solidity: event DelegateForAll(address vault, address delegate, bool value)
func (_Delegatecash *DelegatecashFilterer) FilterDelegateForAll(opts *bind.FilterOpts) (*DelegatecashDelegateForAllIterator, error) {

	logs, sub, err := _Delegatecash.contract.FilterLogs(opts, "DelegateForAll")
	if err != nil {
		return nil, err
	}
	return &DelegatecashDelegateForAllIterator{contract: _Delegatecash.contract, event: "DelegateForAll", logs: logs, sub: sub}, nil
}

// WatchDelegateForAll is a free log subscription operation binding the contract event 0x58781eab4a0743ab1c285a238be846a235f06cdb5b968030573a635e5f8c92fa.
//
// Solidity: event DelegateForAll(address vault, address delegate, bool value)
func (_Delegatecash *DelegatecashFilterer) WatchDelegateForAll(opts *bind.WatchOpts, sink chan<- *DelegatecashDelegateForAll) (event.Subscription, error) {

	logs, sub, err := _Delegatecash.contract.WatchLogs(opts, "DelegateForAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatecashDelegateForAll)
				if err := _Delegatecash.contract.UnpackLog(event, "DelegateForAll", log); err != nil {
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

// ParseDelegateForAll is a log parse operation binding the contract event 0x58781eab4a0743ab1c285a238be846a235f06cdb5b968030573a635e5f8c92fa.
//
// Solidity: event DelegateForAll(address vault, address delegate, bool value)
func (_Delegatecash *DelegatecashFilterer) ParseDelegateForAll(log types.Log) (*DelegatecashDelegateForAll, error) {
	event := new(DelegatecashDelegateForAll)
	if err := _Delegatecash.contract.UnpackLog(event, "DelegateForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatecashDelegateForContractIterator is returned from FilterDelegateForContract and is used to iterate over the raw logs and unpacked data for DelegateForContract events raised by the Delegatecash contract.
type DelegatecashDelegateForContractIterator struct {
	Event *DelegatecashDelegateForContract // Event containing the contract specifics and raw log

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
func (it *DelegatecashDelegateForContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatecashDelegateForContract)
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
		it.Event = new(DelegatecashDelegateForContract)
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
func (it *DelegatecashDelegateForContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatecashDelegateForContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatecashDelegateForContract represents a DelegateForContract event raised by the Delegatecash contract.
type DelegatecashDelegateForContract struct {
	Vault    common.Address
	Delegate common.Address
	Contract common.Address
	Value    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateForContract is a free log retrieval operation binding the contract event 0x8d6b2f5255b8d815cc368855b2251146e003bf4e2fcccaec66145fff5c174b4f.
//
// Solidity: event DelegateForContract(address vault, address delegate, address contract_, bool value)
func (_Delegatecash *DelegatecashFilterer) FilterDelegateForContract(opts *bind.FilterOpts) (*DelegatecashDelegateForContractIterator, error) {

	logs, sub, err := _Delegatecash.contract.FilterLogs(opts, "DelegateForContract")
	if err != nil {
		return nil, err
	}
	return &DelegatecashDelegateForContractIterator{contract: _Delegatecash.contract, event: "DelegateForContract", logs: logs, sub: sub}, nil
}

// WatchDelegateForContract is a free log subscription operation binding the contract event 0x8d6b2f5255b8d815cc368855b2251146e003bf4e2fcccaec66145fff5c174b4f.
//
// Solidity: event DelegateForContract(address vault, address delegate, address contract_, bool value)
func (_Delegatecash *DelegatecashFilterer) WatchDelegateForContract(opts *bind.WatchOpts, sink chan<- *DelegatecashDelegateForContract) (event.Subscription, error) {

	logs, sub, err := _Delegatecash.contract.WatchLogs(opts, "DelegateForContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatecashDelegateForContract)
				if err := _Delegatecash.contract.UnpackLog(event, "DelegateForContract", log); err != nil {
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

// ParseDelegateForContract is a log parse operation binding the contract event 0x8d6b2f5255b8d815cc368855b2251146e003bf4e2fcccaec66145fff5c174b4f.
//
// Solidity: event DelegateForContract(address vault, address delegate, address contract_, bool value)
func (_Delegatecash *DelegatecashFilterer) ParseDelegateForContract(log types.Log) (*DelegatecashDelegateForContract, error) {
	event := new(DelegatecashDelegateForContract)
	if err := _Delegatecash.contract.UnpackLog(event, "DelegateForContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatecashDelegateForTokenIterator is returned from FilterDelegateForToken and is used to iterate over the raw logs and unpacked data for DelegateForToken events raised by the Delegatecash contract.
type DelegatecashDelegateForTokenIterator struct {
	Event *DelegatecashDelegateForToken // Event containing the contract specifics and raw log

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
func (it *DelegatecashDelegateForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatecashDelegateForToken)
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
		it.Event = new(DelegatecashDelegateForToken)
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
func (it *DelegatecashDelegateForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatecashDelegateForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatecashDelegateForToken represents a DelegateForToken event raised by the Delegatecash contract.
type DelegatecashDelegateForToken struct {
	Vault    common.Address
	Delegate common.Address
	Contract common.Address
	TokenId  *big.Int
	Value    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateForToken is a free log retrieval operation binding the contract event 0xe89c6ba1e8957285aed22618f52aa1dcb9d5bb64e1533d8b55136c72fcf5aa5d.
//
// Solidity: event DelegateForToken(address vault, address delegate, address contract_, uint256 tokenId, bool value)
func (_Delegatecash *DelegatecashFilterer) FilterDelegateForToken(opts *bind.FilterOpts) (*DelegatecashDelegateForTokenIterator, error) {

	logs, sub, err := _Delegatecash.contract.FilterLogs(opts, "DelegateForToken")
	if err != nil {
		return nil, err
	}
	return &DelegatecashDelegateForTokenIterator{contract: _Delegatecash.contract, event: "DelegateForToken", logs: logs, sub: sub}, nil
}

// WatchDelegateForToken is a free log subscription operation binding the contract event 0xe89c6ba1e8957285aed22618f52aa1dcb9d5bb64e1533d8b55136c72fcf5aa5d.
//
// Solidity: event DelegateForToken(address vault, address delegate, address contract_, uint256 tokenId, bool value)
func (_Delegatecash *DelegatecashFilterer) WatchDelegateForToken(opts *bind.WatchOpts, sink chan<- *DelegatecashDelegateForToken) (event.Subscription, error) {

	logs, sub, err := _Delegatecash.contract.WatchLogs(opts, "DelegateForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatecashDelegateForToken)
				if err := _Delegatecash.contract.UnpackLog(event, "DelegateForToken", log); err != nil {
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

// ParseDelegateForToken is a log parse operation binding the contract event 0xe89c6ba1e8957285aed22618f52aa1dcb9d5bb64e1533d8b55136c72fcf5aa5d.
//
// Solidity: event DelegateForToken(address vault, address delegate, address contract_, uint256 tokenId, bool value)
func (_Delegatecash *DelegatecashFilterer) ParseDelegateForToken(log types.Log) (*DelegatecashDelegateForToken, error) {
	event := new(DelegatecashDelegateForToken)
	if err := _Delegatecash.contract.UnpackLog(event, "DelegateForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatecashRevokeAllDelegatesIterator is returned from FilterRevokeAllDelegates and is used to iterate over the raw logs and unpacked data for RevokeAllDelegates events raised by the Delegatecash contract.
type DelegatecashRevokeAllDelegatesIterator struct {
	Event *DelegatecashRevokeAllDelegates // Event containing the contract specifics and raw log

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
func (it *DelegatecashRevokeAllDelegatesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatecashRevokeAllDelegates)
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
		it.Event = new(DelegatecashRevokeAllDelegates)
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
func (it *DelegatecashRevokeAllDelegatesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatecashRevokeAllDelegatesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatecashRevokeAllDelegates represents a RevokeAllDelegates event raised by the Delegatecash contract.
type DelegatecashRevokeAllDelegates struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRevokeAllDelegates is a free log retrieval operation binding the contract event 0x32d74befd0b842e19694e3e3af46263e18bcce41352c8b600ff0002b49edf662.
//
// Solidity: event RevokeAllDelegates(address vault)
func (_Delegatecash *DelegatecashFilterer) FilterRevokeAllDelegates(opts *bind.FilterOpts) (*DelegatecashRevokeAllDelegatesIterator, error) {

	logs, sub, err := _Delegatecash.contract.FilterLogs(opts, "RevokeAllDelegates")
	if err != nil {
		return nil, err
	}
	return &DelegatecashRevokeAllDelegatesIterator{contract: _Delegatecash.contract, event: "RevokeAllDelegates", logs: logs, sub: sub}, nil
}

// WatchRevokeAllDelegates is a free log subscription operation binding the contract event 0x32d74befd0b842e19694e3e3af46263e18bcce41352c8b600ff0002b49edf662.
//
// Solidity: event RevokeAllDelegates(address vault)
func (_Delegatecash *DelegatecashFilterer) WatchRevokeAllDelegates(opts *bind.WatchOpts, sink chan<- *DelegatecashRevokeAllDelegates) (event.Subscription, error) {

	logs, sub, err := _Delegatecash.contract.WatchLogs(opts, "RevokeAllDelegates")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatecashRevokeAllDelegates)
				if err := _Delegatecash.contract.UnpackLog(event, "RevokeAllDelegates", log); err != nil {
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

// ParseRevokeAllDelegates is a log parse operation binding the contract event 0x32d74befd0b842e19694e3e3af46263e18bcce41352c8b600ff0002b49edf662.
//
// Solidity: event RevokeAllDelegates(address vault)
func (_Delegatecash *DelegatecashFilterer) ParseRevokeAllDelegates(log types.Log) (*DelegatecashRevokeAllDelegates, error) {
	event := new(DelegatecashRevokeAllDelegates)
	if err := _Delegatecash.contract.UnpackLog(event, "RevokeAllDelegates", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatecashRevokeDelegateIterator is returned from FilterRevokeDelegate and is used to iterate over the raw logs and unpacked data for RevokeDelegate events raised by the Delegatecash contract.
type DelegatecashRevokeDelegateIterator struct {
	Event *DelegatecashRevokeDelegate // Event containing the contract specifics and raw log

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
func (it *DelegatecashRevokeDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatecashRevokeDelegate)
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
		it.Event = new(DelegatecashRevokeDelegate)
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
func (it *DelegatecashRevokeDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatecashRevokeDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatecashRevokeDelegate represents a RevokeDelegate event raised by the Delegatecash contract.
type DelegatecashRevokeDelegate struct {
	Vault    common.Address
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRevokeDelegate is a free log retrieval operation binding the contract event 0x3e34a3ee53064fb79c0ee57448f03774a627a9270b0c41286efb7d8e32dcde93.
//
// Solidity: event RevokeDelegate(address vault, address delegate)
func (_Delegatecash *DelegatecashFilterer) FilterRevokeDelegate(opts *bind.FilterOpts) (*DelegatecashRevokeDelegateIterator, error) {

	logs, sub, err := _Delegatecash.contract.FilterLogs(opts, "RevokeDelegate")
	if err != nil {
		return nil, err
	}
	return &DelegatecashRevokeDelegateIterator{contract: _Delegatecash.contract, event: "RevokeDelegate", logs: logs, sub: sub}, nil
}

// WatchRevokeDelegate is a free log subscription operation binding the contract event 0x3e34a3ee53064fb79c0ee57448f03774a627a9270b0c41286efb7d8e32dcde93.
//
// Solidity: event RevokeDelegate(address vault, address delegate)
func (_Delegatecash *DelegatecashFilterer) WatchRevokeDelegate(opts *bind.WatchOpts, sink chan<- *DelegatecashRevokeDelegate) (event.Subscription, error) {

	logs, sub, err := _Delegatecash.contract.WatchLogs(opts, "RevokeDelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatecashRevokeDelegate)
				if err := _Delegatecash.contract.UnpackLog(event, "RevokeDelegate", log); err != nil {
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

// ParseRevokeDelegate is a log parse operation binding the contract event 0x3e34a3ee53064fb79c0ee57448f03774a627a9270b0c41286efb7d8e32dcde93.
//
// Solidity: event RevokeDelegate(address vault, address delegate)
func (_Delegatecash *DelegatecashFilterer) ParseRevokeDelegate(log types.Log) (*DelegatecashRevokeDelegate, error) {
	event := new(DelegatecashRevokeDelegate)
	if err := _Delegatecash.contract.UnpackLog(event, "RevokeDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
