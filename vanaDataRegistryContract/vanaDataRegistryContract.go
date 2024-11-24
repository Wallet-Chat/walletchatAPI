// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vanaDataRegistryContract

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

// IDataRegistryFileResponse is an auto generated low-level Go binding around an user-defined struct.
type IDataRegistryFileResponse struct {
	Id           *big.Int
	OwnerAddress common.Address
	Url          string
	AddedAtBlock *big.Int
}

// IDataRegistryPermission is an auto generated low-level Go binding around an user-defined struct.
type IDataRegistryPermission struct {
	Account common.Address
	Key     string
}

// IDataRegistryProof is an auto generated low-level Go binding around an user-defined struct.
type IDataRegistryProof struct {
	Signature []byte
	Data      IDataRegistryProofData
}

// IDataRegistryProofData is an auto generated low-level Go binding around an user-defined struct.
type IDataRegistryProofData struct {
	Score       *big.Int
	DlpId       *big.Int
	Metadata    string
	ProofUrl    string
	Instruction string
}

// VanaDataRegistryContractMetaData contains all meta data concerning the VanaDataRegistryContract contract.
var VanaDataRegistryContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"FileAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PermissionGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proofIndex\",\"type\":\"uint256\"}],\"name\":\"ProofAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"addFile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"addFilePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"internalType\":\"structIDataRegistry.Permission[]\",\"name\":\"permissions\",\"type\":\"tuple[]\"}],\"name\":\"addFileWithPermissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dlpId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proofUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"instruction\",\"type\":\"string\"}],\"internalType\":\"structIDataRegistry.ProofData\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structIDataRegistry.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"addProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"filePermissions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"fileProofs\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dlpId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proofUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"instruction\",\"type\":\"string\"}],\"internalType\":\"structIDataRegistry.ProofData\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structIDataRegistry.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"files\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"addedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIDataRegistry.FileResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"filesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// VanaDataRegistryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use VanaDataRegistryContractMetaData.ABI instead.
var VanaDataRegistryContractABI = VanaDataRegistryContractMetaData.ABI

// VanaDataRegistryContract is an auto generated Go binding around an Ethereum contract.
type VanaDataRegistryContract struct {
	VanaDataRegistryContractCaller     // Read-only binding to the contract
	VanaDataRegistryContractTransactor // Write-only binding to the contract
	VanaDataRegistryContractFilterer   // Log filterer for contract events
}

// VanaDataRegistryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VanaDataRegistryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VanaDataRegistryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VanaDataRegistryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VanaDataRegistryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VanaDataRegistryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VanaDataRegistryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VanaDataRegistryContractSession struct {
	Contract     *VanaDataRegistryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VanaDataRegistryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VanaDataRegistryContractCallerSession struct {
	Contract *VanaDataRegistryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// VanaDataRegistryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VanaDataRegistryContractTransactorSession struct {
	Contract     *VanaDataRegistryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// VanaDataRegistryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VanaDataRegistryContractRaw struct {
	Contract *VanaDataRegistryContract // Generic contract binding to access the raw methods on
}

// VanaDataRegistryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VanaDataRegistryContractCallerRaw struct {
	Contract *VanaDataRegistryContractCaller // Generic read-only contract binding to access the raw methods on
}

// VanaDataRegistryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VanaDataRegistryContractTransactorRaw struct {
	Contract *VanaDataRegistryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVanaDataRegistryContract creates a new instance of VanaDataRegistryContract, bound to a specific deployed contract.
func NewVanaDataRegistryContract(address common.Address, backend bind.ContractBackend) (*VanaDataRegistryContract, error) {
	contract, err := bindVanaDataRegistryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContract{VanaDataRegistryContractCaller: VanaDataRegistryContractCaller{contract: contract}, VanaDataRegistryContractTransactor: VanaDataRegistryContractTransactor{contract: contract}, VanaDataRegistryContractFilterer: VanaDataRegistryContractFilterer{contract: contract}}, nil
}

// NewVanaDataRegistryContractCaller creates a new read-only instance of VanaDataRegistryContract, bound to a specific deployed contract.
func NewVanaDataRegistryContractCaller(address common.Address, caller bind.ContractCaller) (*VanaDataRegistryContractCaller, error) {
	contract, err := bindVanaDataRegistryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractCaller{contract: contract}, nil
}

// NewVanaDataRegistryContractTransactor creates a new write-only instance of VanaDataRegistryContract, bound to a specific deployed contract.
func NewVanaDataRegistryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VanaDataRegistryContractTransactor, error) {
	contract, err := bindVanaDataRegistryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractTransactor{contract: contract}, nil
}

// NewVanaDataRegistryContractFilterer creates a new log filterer instance of VanaDataRegistryContract, bound to a specific deployed contract.
func NewVanaDataRegistryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*VanaDataRegistryContractFilterer, error) {
	contract, err := bindVanaDataRegistryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractFilterer{contract: contract}, nil
}

// bindVanaDataRegistryContract binds a generic wrapper to an already deployed contract.
func bindVanaDataRegistryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VanaDataRegistryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VanaDataRegistryContract *VanaDataRegistryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VanaDataRegistryContract.Contract.VanaDataRegistryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VanaDataRegistryContract *VanaDataRegistryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.VanaDataRegistryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VanaDataRegistryContract *VanaDataRegistryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.VanaDataRegistryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VanaDataRegistryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VanaDataRegistryContract *VanaDataRegistryContractCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VanaDataRegistryContract.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _VanaDataRegistryContract.Contract.UPGRADEINTERFACEVERSION(&_VanaDataRegistryContract.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _VanaDataRegistryContract.Contract.UPGRADEINTERFACEVERSION(&_VanaDataRegistryContract.CallOpts)
}

// FilePermissions is a free data retrieval call binding the contract method 0x60f1c43a.
//
// Solidity: function filePermissions(uint256 fileId, address account) view returns(string)
func (_VanaDataRegistryContract *VanaDataRegistryContractCaller) FilePermissions(opts *bind.CallOpts, fileId *big.Int, account common.Address) (string, error) {
	var out []interface{}
	err := _VanaDataRegistryContract.contract.Call(opts, &out, "filePermissions", fileId, account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FilePermissions is a free data retrieval call binding the contract method 0x60f1c43a.
//
// Solidity: function filePermissions(uint256 fileId, address account) view returns(string)
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) FilePermissions(fileId *big.Int, account common.Address) (string, error) {
	return _VanaDataRegistryContract.Contract.FilePermissions(&_VanaDataRegistryContract.CallOpts, fileId, account)
}

// FilePermissions is a free data retrieval call binding the contract method 0x60f1c43a.
//
// Solidity: function filePermissions(uint256 fileId, address account) view returns(string)
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerSession) FilePermissions(fileId *big.Int, account common.Address) (string, error) {
	return _VanaDataRegistryContract.Contract.FilePermissions(&_VanaDataRegistryContract.CallOpts, fileId, account)
}

// FileProofs is a free data retrieval call binding the contract method 0xdbda68db.
//
// Solidity: function fileProofs(uint256 fileId, uint256 index) view returns((bytes,(uint256,uint256,string,string,string)))
func (_VanaDataRegistryContract *VanaDataRegistryContractCaller) FileProofs(opts *bind.CallOpts, fileId *big.Int, index *big.Int) (IDataRegistryProof, error) {
	var out []interface{}
	err := _VanaDataRegistryContract.contract.Call(opts, &out, "fileProofs", fileId, index)

	if err != nil {
		return *new(IDataRegistryProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IDataRegistryProof)).(*IDataRegistryProof)

	return out0, err

}

// FileProofs is a free data retrieval call binding the contract method 0xdbda68db.
//
// Solidity: function fileProofs(uint256 fileId, uint256 index) view returns((bytes,(uint256,uint256,string,string,string)))
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) FileProofs(fileId *big.Int, index *big.Int) (IDataRegistryProof, error) {
	return _VanaDataRegistryContract.Contract.FileProofs(&_VanaDataRegistryContract.CallOpts, fileId, index)
}

// FileProofs is a free data retrieval call binding the contract method 0xdbda68db.
//
// Solidity: function fileProofs(uint256 fileId, uint256 index) view returns((bytes,(uint256,uint256,string,string,string)))
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerSession) FileProofs(fileId *big.Int, index *big.Int) (IDataRegistryProof, error) {
	return _VanaDataRegistryContract.Contract.FileProofs(&_VanaDataRegistryContract.CallOpts, fileId, index)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 fileId) view returns((uint256,address,string,uint256))
func (_VanaDataRegistryContract *VanaDataRegistryContractCaller) Files(opts *bind.CallOpts, fileId *big.Int) (IDataRegistryFileResponse, error) {
	var out []interface{}
	err := _VanaDataRegistryContract.contract.Call(opts, &out, "files", fileId)

	if err != nil {
		return *new(IDataRegistryFileResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(IDataRegistryFileResponse)).(*IDataRegistryFileResponse)

	return out0, err

}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 fileId) view returns((uint256,address,string,uint256))
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) Files(fileId *big.Int) (IDataRegistryFileResponse, error) {
	return _VanaDataRegistryContract.Contract.Files(&_VanaDataRegistryContract.CallOpts, fileId)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 fileId) view returns((uint256,address,string,uint256))
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerSession) Files(fileId *big.Int) (IDataRegistryFileResponse, error) {
	return _VanaDataRegistryContract.Contract.Files(&_VanaDataRegistryContract.CallOpts, fileId)
}

// FilesCount is a free data retrieval call binding the contract method 0xdfb5e9ba.
//
// Solidity: function filesCount() view returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractCaller) FilesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaDataRegistryContract.contract.Call(opts, &out, "filesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FilesCount is a free data retrieval call binding the contract method 0xdfb5e9ba.
//
// Solidity: function filesCount() view returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) FilesCount() (*big.Int, error) {
	return _VanaDataRegistryContract.Contract.FilesCount(&_VanaDataRegistryContract.CallOpts)
}

// FilesCount is a free data retrieval call binding the contract method 0xdfb5e9ba.
//
// Solidity: function filesCount() view returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerSession) FilesCount() (*big.Int, error) {
	return _VanaDataRegistryContract.Contract.FilesCount(&_VanaDataRegistryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VanaDataRegistryContract *VanaDataRegistryContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaDataRegistryContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) Owner() (common.Address, error) {
	return _VanaDataRegistryContract.Contract.Owner(&_VanaDataRegistryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerSession) Owner() (common.Address, error) {
	return _VanaDataRegistryContract.Contract.Owner(&_VanaDataRegistryContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VanaDataRegistryContract *VanaDataRegistryContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VanaDataRegistryContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) Paused() (bool, error) {
	return _VanaDataRegistryContract.Contract.Paused(&_VanaDataRegistryContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerSession) Paused() (bool, error) {
	return _VanaDataRegistryContract.Contract.Paused(&_VanaDataRegistryContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VanaDataRegistryContract *VanaDataRegistryContractCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaDataRegistryContract.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) PendingOwner() (common.Address, error) {
	return _VanaDataRegistryContract.Contract.PendingOwner(&_VanaDataRegistryContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerSession) PendingOwner() (common.Address, error) {
	return _VanaDataRegistryContract.Contract.PendingOwner(&_VanaDataRegistryContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VanaDataRegistryContract *VanaDataRegistryContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VanaDataRegistryContract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) ProxiableUUID() ([32]byte, error) {
	return _VanaDataRegistryContract.Contract.ProxiableUUID(&_VanaDataRegistryContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VanaDataRegistryContract.Contract.ProxiableUUID(&_VanaDataRegistryContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaDataRegistryContract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) Version() (*big.Int, error) {
	return _VanaDataRegistryContract.Contract.Version(&_VanaDataRegistryContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractCallerSession) Version() (*big.Int, error) {
	return _VanaDataRegistryContract.Contract.Version(&_VanaDataRegistryContract.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.AcceptOwnership(&_VanaDataRegistryContract.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.AcceptOwnership(&_VanaDataRegistryContract.TransactOpts)
}

// AddFile is a paid mutator transaction binding the contract method 0xeb9b9b64.
//
// Solidity: function addFile(string url) returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) AddFile(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "addFile", url)
}

// AddFile is a paid mutator transaction binding the contract method 0xeb9b9b64.
//
// Solidity: function addFile(string url) returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) AddFile(url string) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.AddFile(&_VanaDataRegistryContract.TransactOpts, url)
}

// AddFile is a paid mutator transaction binding the contract method 0xeb9b9b64.
//
// Solidity: function addFile(string url) returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) AddFile(url string) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.AddFile(&_VanaDataRegistryContract.TransactOpts, url)
}

// AddFilePermission is a paid mutator transaction binding the contract method 0xf75cf867.
//
// Solidity: function addFilePermission(uint256 fileId, address account, string key) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) AddFilePermission(opts *bind.TransactOpts, fileId *big.Int, account common.Address, key string) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "addFilePermission", fileId, account, key)
}

// AddFilePermission is a paid mutator transaction binding the contract method 0xf75cf867.
//
// Solidity: function addFilePermission(uint256 fileId, address account, string key) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) AddFilePermission(fileId *big.Int, account common.Address, key string) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.AddFilePermission(&_VanaDataRegistryContract.TransactOpts, fileId, account, key)
}

// AddFilePermission is a paid mutator transaction binding the contract method 0xf75cf867.
//
// Solidity: function addFilePermission(uint256 fileId, address account, string key) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) AddFilePermission(fileId *big.Int, account common.Address, key string) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.AddFilePermission(&_VanaDataRegistryContract.TransactOpts, fileId, account, key)
}

// AddFileWithPermissions is a paid mutator transaction binding the contract method 0xafbfc156.
//
// Solidity: function addFileWithPermissions(string url, address ownerAddress, (address,string)[] permissions) returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) AddFileWithPermissions(opts *bind.TransactOpts, url string, ownerAddress common.Address, permissions []IDataRegistryPermission) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "addFileWithPermissions", url, ownerAddress, permissions)
}

// AddFileWithPermissions is a paid mutator transaction binding the contract method 0xafbfc156.
//
// Solidity: function addFileWithPermissions(string url, address ownerAddress, (address,string)[] permissions) returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) AddFileWithPermissions(url string, ownerAddress common.Address, permissions []IDataRegistryPermission) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.AddFileWithPermissions(&_VanaDataRegistryContract.TransactOpts, url, ownerAddress, permissions)
}

// AddFileWithPermissions is a paid mutator transaction binding the contract method 0xafbfc156.
//
// Solidity: function addFileWithPermissions(string url, address ownerAddress, (address,string)[] permissions) returns(uint256)
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) AddFileWithPermissions(url string, ownerAddress common.Address, permissions []IDataRegistryPermission) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.AddFileWithPermissions(&_VanaDataRegistryContract.TransactOpts, url, ownerAddress, permissions)
}

// AddProof is a paid mutator transaction binding the contract method 0xc26045f7.
//
// Solidity: function addProof(uint256 fileId, (bytes,(uint256,uint256,string,string,string)) proof) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) AddProof(opts *bind.TransactOpts, fileId *big.Int, proof IDataRegistryProof) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "addProof", fileId, proof)
}

// AddProof is a paid mutator transaction binding the contract method 0xc26045f7.
//
// Solidity: function addProof(uint256 fileId, (bytes,(uint256,uint256,string,string,string)) proof) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) AddProof(fileId *big.Int, proof IDataRegistryProof) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.AddProof(&_VanaDataRegistryContract.TransactOpts, fileId, proof)
}

// AddProof is a paid mutator transaction binding the contract method 0xc26045f7.
//
// Solidity: function addProof(uint256 fileId, (bytes,(uint256,uint256,string,string,string)) proof) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) AddProof(fileId *big.Int, proof IDataRegistryProof) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.AddProof(&_VanaDataRegistryContract.TransactOpts, fileId, proof)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ownerAddress) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) Initialize(opts *bind.TransactOpts, ownerAddress common.Address) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "initialize", ownerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ownerAddress) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) Initialize(ownerAddress common.Address) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.Initialize(&_VanaDataRegistryContract.TransactOpts, ownerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ownerAddress) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) Initialize(ownerAddress common.Address) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.Initialize(&_VanaDataRegistryContract.TransactOpts, ownerAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) Pause() (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.Pause(&_VanaDataRegistryContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) Pause() (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.Pause(&_VanaDataRegistryContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.RenounceOwnership(&_VanaDataRegistryContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.RenounceOwnership(&_VanaDataRegistryContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.TransferOwnership(&_VanaDataRegistryContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.TransferOwnership(&_VanaDataRegistryContract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) Unpause() (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.Unpause(&_VanaDataRegistryContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.Unpause(&_VanaDataRegistryContract.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VanaDataRegistryContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.UpgradeToAndCall(&_VanaDataRegistryContract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VanaDataRegistryContract *VanaDataRegistryContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VanaDataRegistryContract.Contract.UpgradeToAndCall(&_VanaDataRegistryContract.TransactOpts, newImplementation, data)
}

// VanaDataRegistryContractFileAddedIterator is returned from FilterFileAdded and is used to iterate over the raw logs and unpacked data for FileAdded events raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractFileAddedIterator struct {
	Event *VanaDataRegistryContractFileAdded // Event containing the contract specifics and raw log

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
func (it *VanaDataRegistryContractFileAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDataRegistryContractFileAdded)
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
		it.Event = new(VanaDataRegistryContractFileAdded)
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
func (it *VanaDataRegistryContractFileAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDataRegistryContractFileAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDataRegistryContractFileAdded represents a FileAdded event raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractFileAdded struct {
	FileId       *big.Int
	OwnerAddress common.Address
	Url          string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFileAdded is a free log retrieval operation binding the contract event 0x2980b13036c0d9df6b71ea9e9c09536efd24d7dd1868e29aa730f0c3fbd1a0dc.
//
// Solidity: event FileAdded(uint256 indexed fileId, address indexed ownerAddress, string url)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) FilterFileAdded(opts *bind.FilterOpts, fileId []*big.Int, ownerAddress []common.Address) (*VanaDataRegistryContractFileAddedIterator, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.FilterLogs(opts, "FileAdded", fileIdRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractFileAddedIterator{contract: _VanaDataRegistryContract.contract, event: "FileAdded", logs: logs, sub: sub}, nil
}

// WatchFileAdded is a free log subscription operation binding the contract event 0x2980b13036c0d9df6b71ea9e9c09536efd24d7dd1868e29aa730f0c3fbd1a0dc.
//
// Solidity: event FileAdded(uint256 indexed fileId, address indexed ownerAddress, string url)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) WatchFileAdded(opts *bind.WatchOpts, sink chan<- *VanaDataRegistryContractFileAdded, fileId []*big.Int, ownerAddress []common.Address) (event.Subscription, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.WatchLogs(opts, "FileAdded", fileIdRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDataRegistryContractFileAdded)
				if err := _VanaDataRegistryContract.contract.UnpackLog(event, "FileAdded", log); err != nil {
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

// ParseFileAdded is a log parse operation binding the contract event 0x2980b13036c0d9df6b71ea9e9c09536efd24d7dd1868e29aa730f0c3fbd1a0dc.
//
// Solidity: event FileAdded(uint256 indexed fileId, address indexed ownerAddress, string url)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) ParseFileAdded(log types.Log) (*VanaDataRegistryContractFileAdded, error) {
	event := new(VanaDataRegistryContractFileAdded)
	if err := _VanaDataRegistryContract.contract.UnpackLog(event, "FileAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDataRegistryContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractInitializedIterator struct {
	Event *VanaDataRegistryContractInitialized // Event containing the contract specifics and raw log

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
func (it *VanaDataRegistryContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDataRegistryContractInitialized)
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
		it.Event = new(VanaDataRegistryContractInitialized)
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
func (it *VanaDataRegistryContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDataRegistryContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDataRegistryContractInitialized represents a Initialized event raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*VanaDataRegistryContractInitializedIterator, error) {

	logs, sub, err := _VanaDataRegistryContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractInitializedIterator{contract: _VanaDataRegistryContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VanaDataRegistryContractInitialized) (event.Subscription, error) {

	logs, sub, err := _VanaDataRegistryContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDataRegistryContractInitialized)
				if err := _VanaDataRegistryContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) ParseInitialized(log types.Log) (*VanaDataRegistryContractInitialized, error) {
	event := new(VanaDataRegistryContractInitialized)
	if err := _VanaDataRegistryContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDataRegistryContractOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractOwnershipTransferStartedIterator struct {
	Event *VanaDataRegistryContractOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *VanaDataRegistryContractOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDataRegistryContractOwnershipTransferStarted)
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
		it.Event = new(VanaDataRegistryContractOwnershipTransferStarted)
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
func (it *VanaDataRegistryContractOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDataRegistryContractOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDataRegistryContractOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VanaDataRegistryContractOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractOwnershipTransferStartedIterator{contract: _VanaDataRegistryContract.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *VanaDataRegistryContractOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDataRegistryContractOwnershipTransferStarted)
				if err := _VanaDataRegistryContract.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) ParseOwnershipTransferStarted(log types.Log) (*VanaDataRegistryContractOwnershipTransferStarted, error) {
	event := new(VanaDataRegistryContractOwnershipTransferStarted)
	if err := _VanaDataRegistryContract.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDataRegistryContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractOwnershipTransferredIterator struct {
	Event *VanaDataRegistryContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VanaDataRegistryContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDataRegistryContractOwnershipTransferred)
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
		it.Event = new(VanaDataRegistryContractOwnershipTransferred)
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
func (it *VanaDataRegistryContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDataRegistryContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDataRegistryContractOwnershipTransferred represents a OwnershipTransferred event raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VanaDataRegistryContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractOwnershipTransferredIterator{contract: _VanaDataRegistryContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VanaDataRegistryContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDataRegistryContractOwnershipTransferred)
				if err := _VanaDataRegistryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) ParseOwnershipTransferred(log types.Log) (*VanaDataRegistryContractOwnershipTransferred, error) {
	event := new(VanaDataRegistryContractOwnershipTransferred)
	if err := _VanaDataRegistryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDataRegistryContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractPausedIterator struct {
	Event *VanaDataRegistryContractPaused // Event containing the contract specifics and raw log

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
func (it *VanaDataRegistryContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDataRegistryContractPaused)
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
		it.Event = new(VanaDataRegistryContractPaused)
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
func (it *VanaDataRegistryContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDataRegistryContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDataRegistryContractPaused represents a Paused event raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) FilterPaused(opts *bind.FilterOpts) (*VanaDataRegistryContractPausedIterator, error) {

	logs, sub, err := _VanaDataRegistryContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractPausedIterator{contract: _VanaDataRegistryContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VanaDataRegistryContractPaused) (event.Subscription, error) {

	logs, sub, err := _VanaDataRegistryContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDataRegistryContractPaused)
				if err := _VanaDataRegistryContract.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) ParsePaused(log types.Log) (*VanaDataRegistryContractPaused, error) {
	event := new(VanaDataRegistryContractPaused)
	if err := _VanaDataRegistryContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDataRegistryContractPermissionGrantedIterator is returned from FilterPermissionGranted and is used to iterate over the raw logs and unpacked data for PermissionGranted events raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractPermissionGrantedIterator struct {
	Event *VanaDataRegistryContractPermissionGranted // Event containing the contract specifics and raw log

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
func (it *VanaDataRegistryContractPermissionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDataRegistryContractPermissionGranted)
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
		it.Event = new(VanaDataRegistryContractPermissionGranted)
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
func (it *VanaDataRegistryContractPermissionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDataRegistryContractPermissionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDataRegistryContractPermissionGranted represents a PermissionGranted event raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractPermissionGranted struct {
	FileId  *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPermissionGranted is a free log retrieval operation binding the contract event 0x6e0403ef2b13328247bf9260ad8dd9c18fb5a18b95ea25de817ca991da05929b.
//
// Solidity: event PermissionGranted(uint256 indexed fileId, address indexed account)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) FilterPermissionGranted(opts *bind.FilterOpts, fileId []*big.Int, account []common.Address) (*VanaDataRegistryContractPermissionGrantedIterator, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.FilterLogs(opts, "PermissionGranted", fileIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractPermissionGrantedIterator{contract: _VanaDataRegistryContract.contract, event: "PermissionGranted", logs: logs, sub: sub}, nil
}

// WatchPermissionGranted is a free log subscription operation binding the contract event 0x6e0403ef2b13328247bf9260ad8dd9c18fb5a18b95ea25de817ca991da05929b.
//
// Solidity: event PermissionGranted(uint256 indexed fileId, address indexed account)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) WatchPermissionGranted(opts *bind.WatchOpts, sink chan<- *VanaDataRegistryContractPermissionGranted, fileId []*big.Int, account []common.Address) (event.Subscription, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.WatchLogs(opts, "PermissionGranted", fileIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDataRegistryContractPermissionGranted)
				if err := _VanaDataRegistryContract.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
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

// ParsePermissionGranted is a log parse operation binding the contract event 0x6e0403ef2b13328247bf9260ad8dd9c18fb5a18b95ea25de817ca991da05929b.
//
// Solidity: event PermissionGranted(uint256 indexed fileId, address indexed account)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) ParsePermissionGranted(log types.Log) (*VanaDataRegistryContractPermissionGranted, error) {
	event := new(VanaDataRegistryContractPermissionGranted)
	if err := _VanaDataRegistryContract.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDataRegistryContractProofAddedIterator is returned from FilterProofAdded and is used to iterate over the raw logs and unpacked data for ProofAdded events raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractProofAddedIterator struct {
	Event *VanaDataRegistryContractProofAdded // Event containing the contract specifics and raw log

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
func (it *VanaDataRegistryContractProofAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDataRegistryContractProofAdded)
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
		it.Event = new(VanaDataRegistryContractProofAdded)
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
func (it *VanaDataRegistryContractProofAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDataRegistryContractProofAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDataRegistryContractProofAdded represents a ProofAdded event raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractProofAdded struct {
	FileId     *big.Int
	ProofIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProofAdded is a free log retrieval operation binding the contract event 0xaf40333be7583051bc0470d0b829c7717be60e0b467a5f5fd722c30e1a8b00c6.
//
// Solidity: event ProofAdded(uint256 indexed fileId, uint256 indexed proofIndex)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) FilterProofAdded(opts *bind.FilterOpts, fileId []*big.Int, proofIndex []*big.Int) (*VanaDataRegistryContractProofAddedIterator, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var proofIndexRule []interface{}
	for _, proofIndexItem := range proofIndex {
		proofIndexRule = append(proofIndexRule, proofIndexItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.FilterLogs(opts, "ProofAdded", fileIdRule, proofIndexRule)
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractProofAddedIterator{contract: _VanaDataRegistryContract.contract, event: "ProofAdded", logs: logs, sub: sub}, nil
}

// WatchProofAdded is a free log subscription operation binding the contract event 0xaf40333be7583051bc0470d0b829c7717be60e0b467a5f5fd722c30e1a8b00c6.
//
// Solidity: event ProofAdded(uint256 indexed fileId, uint256 indexed proofIndex)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) WatchProofAdded(opts *bind.WatchOpts, sink chan<- *VanaDataRegistryContractProofAdded, fileId []*big.Int, proofIndex []*big.Int) (event.Subscription, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var proofIndexRule []interface{}
	for _, proofIndexItem := range proofIndex {
		proofIndexRule = append(proofIndexRule, proofIndexItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.WatchLogs(opts, "ProofAdded", fileIdRule, proofIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDataRegistryContractProofAdded)
				if err := _VanaDataRegistryContract.contract.UnpackLog(event, "ProofAdded", log); err != nil {
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

// ParseProofAdded is a log parse operation binding the contract event 0xaf40333be7583051bc0470d0b829c7717be60e0b467a5f5fd722c30e1a8b00c6.
//
// Solidity: event ProofAdded(uint256 indexed fileId, uint256 indexed proofIndex)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) ParseProofAdded(log types.Log) (*VanaDataRegistryContractProofAdded, error) {
	event := new(VanaDataRegistryContractProofAdded)
	if err := _VanaDataRegistryContract.contract.UnpackLog(event, "ProofAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDataRegistryContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractUnpausedIterator struct {
	Event *VanaDataRegistryContractUnpaused // Event containing the contract specifics and raw log

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
func (it *VanaDataRegistryContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDataRegistryContractUnpaused)
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
		it.Event = new(VanaDataRegistryContractUnpaused)
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
func (it *VanaDataRegistryContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDataRegistryContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDataRegistryContractUnpaused represents a Unpaused event raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VanaDataRegistryContractUnpausedIterator, error) {

	logs, sub, err := _VanaDataRegistryContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractUnpausedIterator{contract: _VanaDataRegistryContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VanaDataRegistryContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _VanaDataRegistryContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDataRegistryContractUnpaused)
				if err := _VanaDataRegistryContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) ParseUnpaused(log types.Log) (*VanaDataRegistryContractUnpaused, error) {
	event := new(VanaDataRegistryContractUnpaused)
	if err := _VanaDataRegistryContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDataRegistryContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractUpgradedIterator struct {
	Event *VanaDataRegistryContractUpgraded // Event containing the contract specifics and raw log

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
func (it *VanaDataRegistryContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDataRegistryContractUpgraded)
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
		it.Event = new(VanaDataRegistryContractUpgraded)
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
func (it *VanaDataRegistryContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDataRegistryContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDataRegistryContractUpgraded represents a Upgraded event raised by the VanaDataRegistryContract contract.
type VanaDataRegistryContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VanaDataRegistryContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VanaDataRegistryContractUpgradedIterator{contract: _VanaDataRegistryContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VanaDataRegistryContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VanaDataRegistryContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDataRegistryContractUpgraded)
				if err := _VanaDataRegistryContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VanaDataRegistryContract *VanaDataRegistryContractFilterer) ParseUpgraded(log types.Log) (*VanaDataRegistryContractUpgraded, error) {
	event := new(VanaDataRegistryContractUpgraded)
	if err := _VanaDataRegistryContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
