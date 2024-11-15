// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vanaDlpContract

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

// DataLiquidityPoolImplementationInitParams is an auto generated low-level Go binding around an user-defined struct.
type DataLiquidityPoolImplementationInitParams struct {
	OwnerAddress        common.Address
	TokenAddress        common.Address
	DataRegistryAddress common.Address
	TeePoolAddress      common.Address
	Name                string
	MasterKey           string
	ProofInstruction    string
	FileRewardFactor    *big.Int
}

// IDataLiquidityPoolContributorInfoResponse is an auto generated low-level Go binding around an user-defined struct.
type IDataLiquidityPoolContributorInfoResponse struct {
	ContributorAddress common.Address
	FilesListCount     *big.Int
}

// IDataLiquidityPoolFileResponse is an auto generated low-level Go binding around an user-defined struct.
type IDataLiquidityPoolFileResponse struct {
	FileId       *big.Int
	Timestamp    *big.Int
	ProofIndex   *big.Int
	RewardAmount *big.Int
}

// VanaDlpContractMetaData contains all meta data concerning the VanaDlpContract contract.
var VanaDlpContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"FileInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFileRewardFactor\",\"type\":\"uint256\"}],\"name\":\"FileRewardFactorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"FileValidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMasterKey\",\"type\":\"string\"}],\"name\":\"MasterKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newProofInstruction\",\"type\":\"string\"}],\"name\":\"ProofInstructionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proofIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"contributorsRewardAmount\",\"type\":\"uint256\"}],\"name\":\"addRewardsForContributors\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contributorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"contributorFiles\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIDataLiquidityPool.FileResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contributorAddress\",\"type\":\"address\"}],\"name\":\"contributorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contributorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"filesListCount\",\"type\":\"uint256\"}],\"internalType\":\"structIDataLiquidityPool.ContributorInfoResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"contributors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contributorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"filesListCount\",\"type\":\"uint256\"}],\"internalType\":\"structIDataLiquidityPool.ContributorInfoResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contributorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataRegistry\",\"outputs\":[{\"internalType\":\"contractIDataRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fileRewardFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"files\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIDataLiquidityPool.FileResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"filesListAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"filesListCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dataRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teePoolAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"masterKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proofInstruction\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fileRewardFactor\",\"type\":\"uint256\"}],\"internalType\":\"structDataLiquidityPoolImplementation.InitParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofInstruction\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofIndex\",\"type\":\"uint256\"}],\"name\":\"requestReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePool\",\"outputs\":[{\"internalType\":\"contractITeePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalContributorsRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFileRewardFactor\",\"type\":\"uint256\"}],\"name\":\"updateFileRewardFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newMasterKey\",\"type\":\"string\"}],\"name\":\"updateMasterKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newProofInstruction\",\"type\":\"string\"}],\"name\":\"updateProofInstruction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTeePool\",\"type\":\"address\"}],\"name\":\"updateTeePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// VanaDlpContractABI is the input ABI used to generate the binding from.
// Deprecated: Use VanaDlpContractMetaData.ABI instead.
var VanaDlpContractABI = VanaDlpContractMetaData.ABI

// VanaDlpContract is an auto generated Go binding around an Ethereum contract.
type VanaDlpContract struct {
	VanaDlpContractCaller     // Read-only binding to the contract
	VanaDlpContractTransactor // Write-only binding to the contract
	VanaDlpContractFilterer   // Log filterer for contract events
}

// VanaDlpContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VanaDlpContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VanaDlpContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VanaDlpContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VanaDlpContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VanaDlpContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VanaDlpContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VanaDlpContractSession struct {
	Contract     *VanaDlpContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VanaDlpContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VanaDlpContractCallerSession struct {
	Contract *VanaDlpContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VanaDlpContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VanaDlpContractTransactorSession struct {
	Contract     *VanaDlpContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VanaDlpContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VanaDlpContractRaw struct {
	Contract *VanaDlpContract // Generic contract binding to access the raw methods on
}

// VanaDlpContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VanaDlpContractCallerRaw struct {
	Contract *VanaDlpContractCaller // Generic read-only contract binding to access the raw methods on
}

// VanaDlpContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VanaDlpContractTransactorRaw struct {
	Contract *VanaDlpContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVanaDlpContract creates a new instance of VanaDlpContract, bound to a specific deployed contract.
func NewVanaDlpContract(address common.Address, backend bind.ContractBackend) (*VanaDlpContract, error) {
	contract, err := bindVanaDlpContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContract{VanaDlpContractCaller: VanaDlpContractCaller{contract: contract}, VanaDlpContractTransactor: VanaDlpContractTransactor{contract: contract}, VanaDlpContractFilterer: VanaDlpContractFilterer{contract: contract}}, nil
}

// NewVanaDlpContractCaller creates a new read-only instance of VanaDlpContract, bound to a specific deployed contract.
func NewVanaDlpContractCaller(address common.Address, caller bind.ContractCaller) (*VanaDlpContractCaller, error) {
	contract, err := bindVanaDlpContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractCaller{contract: contract}, nil
}

// NewVanaDlpContractTransactor creates a new write-only instance of VanaDlpContract, bound to a specific deployed contract.
func NewVanaDlpContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VanaDlpContractTransactor, error) {
	contract, err := bindVanaDlpContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractTransactor{contract: contract}, nil
}

// NewVanaDlpContractFilterer creates a new log filterer instance of VanaDlpContract, bound to a specific deployed contract.
func NewVanaDlpContractFilterer(address common.Address, filterer bind.ContractFilterer) (*VanaDlpContractFilterer, error) {
	contract, err := bindVanaDlpContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractFilterer{contract: contract}, nil
}

// bindVanaDlpContract binds a generic wrapper to an already deployed contract.
func bindVanaDlpContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VanaDlpContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VanaDlpContract *VanaDlpContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VanaDlpContract.Contract.VanaDlpContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VanaDlpContract *VanaDlpContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.VanaDlpContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VanaDlpContract *VanaDlpContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.VanaDlpContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VanaDlpContract *VanaDlpContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VanaDlpContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VanaDlpContract *VanaDlpContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VanaDlpContract *VanaDlpContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VanaDlpContract *VanaDlpContractCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VanaDlpContract *VanaDlpContractSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _VanaDlpContract.Contract.UPGRADEINTERFACEVERSION(&_VanaDlpContract.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VanaDlpContract *VanaDlpContractCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _VanaDlpContract.Contract.UPGRADEINTERFACEVERSION(&_VanaDlpContract.CallOpts)
}

// ContributorFiles is a free data retrieval call binding the contract method 0xb3aa4e7b.
//
// Solidity: function contributorFiles(address contributorAddress, uint256 index) view returns((uint256,uint256,uint256,uint256))
func (_VanaDlpContract *VanaDlpContractCaller) ContributorFiles(opts *bind.CallOpts, contributorAddress common.Address, index *big.Int) (IDataLiquidityPoolFileResponse, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "contributorFiles", contributorAddress, index)

	if err != nil {
		return *new(IDataLiquidityPoolFileResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(IDataLiquidityPoolFileResponse)).(*IDataLiquidityPoolFileResponse)

	return out0, err

}

// ContributorFiles is a free data retrieval call binding the contract method 0xb3aa4e7b.
//
// Solidity: function contributorFiles(address contributorAddress, uint256 index) view returns((uint256,uint256,uint256,uint256))
func (_VanaDlpContract *VanaDlpContractSession) ContributorFiles(contributorAddress common.Address, index *big.Int) (IDataLiquidityPoolFileResponse, error) {
	return _VanaDlpContract.Contract.ContributorFiles(&_VanaDlpContract.CallOpts, contributorAddress, index)
}

// ContributorFiles is a free data retrieval call binding the contract method 0xb3aa4e7b.
//
// Solidity: function contributorFiles(address contributorAddress, uint256 index) view returns((uint256,uint256,uint256,uint256))
func (_VanaDlpContract *VanaDlpContractCallerSession) ContributorFiles(contributorAddress common.Address, index *big.Int) (IDataLiquidityPoolFileResponse, error) {
	return _VanaDlpContract.Contract.ContributorFiles(&_VanaDlpContract.CallOpts, contributorAddress, index)
}

// ContributorInfo is a free data retrieval call binding the contract method 0x4b545f3a.
//
// Solidity: function contributorInfo(address contributorAddress) view returns((address,uint256))
func (_VanaDlpContract *VanaDlpContractCaller) ContributorInfo(opts *bind.CallOpts, contributorAddress common.Address) (IDataLiquidityPoolContributorInfoResponse, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "contributorInfo", contributorAddress)

	if err != nil {
		return *new(IDataLiquidityPoolContributorInfoResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(IDataLiquidityPoolContributorInfoResponse)).(*IDataLiquidityPoolContributorInfoResponse)

	return out0, err

}

// ContributorInfo is a free data retrieval call binding the contract method 0x4b545f3a.
//
// Solidity: function contributorInfo(address contributorAddress) view returns((address,uint256))
func (_VanaDlpContract *VanaDlpContractSession) ContributorInfo(contributorAddress common.Address) (IDataLiquidityPoolContributorInfoResponse, error) {
	return _VanaDlpContract.Contract.ContributorInfo(&_VanaDlpContract.CallOpts, contributorAddress)
}

// ContributorInfo is a free data retrieval call binding the contract method 0x4b545f3a.
//
// Solidity: function contributorInfo(address contributorAddress) view returns((address,uint256))
func (_VanaDlpContract *VanaDlpContractCallerSession) ContributorInfo(contributorAddress common.Address) (IDataLiquidityPoolContributorInfoResponse, error) {
	return _VanaDlpContract.Contract.ContributorInfo(&_VanaDlpContract.CallOpts, contributorAddress)
}

// Contributors is a free data retrieval call binding the contract method 0x3cb5d100.
//
// Solidity: function contributors(uint256 index) view returns((address,uint256))
func (_VanaDlpContract *VanaDlpContractCaller) Contributors(opts *bind.CallOpts, index *big.Int) (IDataLiquidityPoolContributorInfoResponse, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "contributors", index)

	if err != nil {
		return *new(IDataLiquidityPoolContributorInfoResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(IDataLiquidityPoolContributorInfoResponse)).(*IDataLiquidityPoolContributorInfoResponse)

	return out0, err

}

// Contributors is a free data retrieval call binding the contract method 0x3cb5d100.
//
// Solidity: function contributors(uint256 index) view returns((address,uint256))
func (_VanaDlpContract *VanaDlpContractSession) Contributors(index *big.Int) (IDataLiquidityPoolContributorInfoResponse, error) {
	return _VanaDlpContract.Contract.Contributors(&_VanaDlpContract.CallOpts, index)
}

// Contributors is a free data retrieval call binding the contract method 0x3cb5d100.
//
// Solidity: function contributors(uint256 index) view returns((address,uint256))
func (_VanaDlpContract *VanaDlpContractCallerSession) Contributors(index *big.Int) (IDataLiquidityPoolContributorInfoResponse, error) {
	return _VanaDlpContract.Contract.Contributors(&_VanaDlpContract.CallOpts, index)
}

// ContributorsCount is a free data retrieval call binding the contract method 0x7569b3d7.
//
// Solidity: function contributorsCount() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractCaller) ContributorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "contributorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContributorsCount is a free data retrieval call binding the contract method 0x7569b3d7.
//
// Solidity: function contributorsCount() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractSession) ContributorsCount() (*big.Int, error) {
	return _VanaDlpContract.Contract.ContributorsCount(&_VanaDlpContract.CallOpts)
}

// ContributorsCount is a free data retrieval call binding the contract method 0x7569b3d7.
//
// Solidity: function contributorsCount() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractCallerSession) ContributorsCount() (*big.Int, error) {
	return _VanaDlpContract.Contract.ContributorsCount(&_VanaDlpContract.CallOpts)
}

// DataRegistry is a free data retrieval call binding the contract method 0xa39c1d6b.
//
// Solidity: function dataRegistry() view returns(address)
func (_VanaDlpContract *VanaDlpContractCaller) DataRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "dataRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataRegistry is a free data retrieval call binding the contract method 0xa39c1d6b.
//
// Solidity: function dataRegistry() view returns(address)
func (_VanaDlpContract *VanaDlpContractSession) DataRegistry() (common.Address, error) {
	return _VanaDlpContract.Contract.DataRegistry(&_VanaDlpContract.CallOpts)
}

// DataRegistry is a free data retrieval call binding the contract method 0xa39c1d6b.
//
// Solidity: function dataRegistry() view returns(address)
func (_VanaDlpContract *VanaDlpContractCallerSession) DataRegistry() (common.Address, error) {
	return _VanaDlpContract.Contract.DataRegistry(&_VanaDlpContract.CallOpts)
}

// FileRewardFactor is a free data retrieval call binding the contract method 0xab049ffd.
//
// Solidity: function fileRewardFactor() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractCaller) FileRewardFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "fileRewardFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FileRewardFactor is a free data retrieval call binding the contract method 0xab049ffd.
//
// Solidity: function fileRewardFactor() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractSession) FileRewardFactor() (*big.Int, error) {
	return _VanaDlpContract.Contract.FileRewardFactor(&_VanaDlpContract.CallOpts)
}

// FileRewardFactor is a free data retrieval call binding the contract method 0xab049ffd.
//
// Solidity: function fileRewardFactor() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractCallerSession) FileRewardFactor() (*big.Int, error) {
	return _VanaDlpContract.Contract.FileRewardFactor(&_VanaDlpContract.CallOpts)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 fileId) view returns((uint256,uint256,uint256,uint256))
func (_VanaDlpContract *VanaDlpContractCaller) Files(opts *bind.CallOpts, fileId *big.Int) (IDataLiquidityPoolFileResponse, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "files", fileId)

	if err != nil {
		return *new(IDataLiquidityPoolFileResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(IDataLiquidityPoolFileResponse)).(*IDataLiquidityPoolFileResponse)

	return out0, err

}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 fileId) view returns((uint256,uint256,uint256,uint256))
func (_VanaDlpContract *VanaDlpContractSession) Files(fileId *big.Int) (IDataLiquidityPoolFileResponse, error) {
	return _VanaDlpContract.Contract.Files(&_VanaDlpContract.CallOpts, fileId)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 fileId) view returns((uint256,uint256,uint256,uint256))
func (_VanaDlpContract *VanaDlpContractCallerSession) Files(fileId *big.Int) (IDataLiquidityPoolFileResponse, error) {
	return _VanaDlpContract.Contract.Files(&_VanaDlpContract.CallOpts, fileId)
}

// FilesListAt is a free data retrieval call binding the contract method 0x3b3cd378.
//
// Solidity: function filesListAt(uint256 index) view returns(uint256)
func (_VanaDlpContract *VanaDlpContractCaller) FilesListAt(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "filesListAt", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FilesListAt is a free data retrieval call binding the contract method 0x3b3cd378.
//
// Solidity: function filesListAt(uint256 index) view returns(uint256)
func (_VanaDlpContract *VanaDlpContractSession) FilesListAt(index *big.Int) (*big.Int, error) {
	return _VanaDlpContract.Contract.FilesListAt(&_VanaDlpContract.CallOpts, index)
}

// FilesListAt is a free data retrieval call binding the contract method 0x3b3cd378.
//
// Solidity: function filesListAt(uint256 index) view returns(uint256)
func (_VanaDlpContract *VanaDlpContractCallerSession) FilesListAt(index *big.Int) (*big.Int, error) {
	return _VanaDlpContract.Contract.FilesListAt(&_VanaDlpContract.CallOpts, index)
}

// FilesListCount is a free data retrieval call binding the contract method 0x7ccf35a6.
//
// Solidity: function filesListCount() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractCaller) FilesListCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "filesListCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FilesListCount is a free data retrieval call binding the contract method 0x7ccf35a6.
//
// Solidity: function filesListCount() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractSession) FilesListCount() (*big.Int, error) {
	return _VanaDlpContract.Contract.FilesListCount(&_VanaDlpContract.CallOpts)
}

// FilesListCount is a free data retrieval call binding the contract method 0x7ccf35a6.
//
// Solidity: function filesListCount() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractCallerSession) FilesListCount() (*big.Int, error) {
	return _VanaDlpContract.Contract.FilesListCount(&_VanaDlpContract.CallOpts)
}

// MasterKey is a free data retrieval call binding the contract method 0x8afe35b9.
//
// Solidity: function masterKey() view returns(string)
func (_VanaDlpContract *VanaDlpContractCaller) MasterKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "masterKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MasterKey is a free data retrieval call binding the contract method 0x8afe35b9.
//
// Solidity: function masterKey() view returns(string)
func (_VanaDlpContract *VanaDlpContractSession) MasterKey() (string, error) {
	return _VanaDlpContract.Contract.MasterKey(&_VanaDlpContract.CallOpts)
}

// MasterKey is a free data retrieval call binding the contract method 0x8afe35b9.
//
// Solidity: function masterKey() view returns(string)
func (_VanaDlpContract *VanaDlpContractCallerSession) MasterKey() (string, error) {
	return _VanaDlpContract.Contract.MasterKey(&_VanaDlpContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VanaDlpContract *VanaDlpContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VanaDlpContract *VanaDlpContractSession) Name() (string, error) {
	return _VanaDlpContract.Contract.Name(&_VanaDlpContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VanaDlpContract *VanaDlpContractCallerSession) Name() (string, error) {
	return _VanaDlpContract.Contract.Name(&_VanaDlpContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VanaDlpContract *VanaDlpContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VanaDlpContract *VanaDlpContractSession) Owner() (common.Address, error) {
	return _VanaDlpContract.Contract.Owner(&_VanaDlpContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VanaDlpContract *VanaDlpContractCallerSession) Owner() (common.Address, error) {
	return _VanaDlpContract.Contract.Owner(&_VanaDlpContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VanaDlpContract *VanaDlpContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VanaDlpContract *VanaDlpContractSession) Paused() (bool, error) {
	return _VanaDlpContract.Contract.Paused(&_VanaDlpContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VanaDlpContract *VanaDlpContractCallerSession) Paused() (bool, error) {
	return _VanaDlpContract.Contract.Paused(&_VanaDlpContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VanaDlpContract *VanaDlpContractCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VanaDlpContract *VanaDlpContractSession) PendingOwner() (common.Address, error) {
	return _VanaDlpContract.Contract.PendingOwner(&_VanaDlpContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VanaDlpContract *VanaDlpContractCallerSession) PendingOwner() (common.Address, error) {
	return _VanaDlpContract.Contract.PendingOwner(&_VanaDlpContract.CallOpts)
}

// ProofInstruction is a free data retrieval call binding the contract method 0x084a09da.
//
// Solidity: function proofInstruction() view returns(string)
func (_VanaDlpContract *VanaDlpContractCaller) ProofInstruction(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "proofInstruction")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ProofInstruction is a free data retrieval call binding the contract method 0x084a09da.
//
// Solidity: function proofInstruction() view returns(string)
func (_VanaDlpContract *VanaDlpContractSession) ProofInstruction() (string, error) {
	return _VanaDlpContract.Contract.ProofInstruction(&_VanaDlpContract.CallOpts)
}

// ProofInstruction is a free data retrieval call binding the contract method 0x084a09da.
//
// Solidity: function proofInstruction() view returns(string)
func (_VanaDlpContract *VanaDlpContractCallerSession) ProofInstruction() (string, error) {
	return _VanaDlpContract.Contract.ProofInstruction(&_VanaDlpContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractSession) ProxiableUUID() ([32]byte, error) {
	return _VanaDlpContract.Contract.ProxiableUUID(&_VanaDlpContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VanaDlpContract.Contract.ProxiableUUID(&_VanaDlpContract.CallOpts)
}

// TeePool is a free data retrieval call binding the contract method 0xd503d4e4.
//
// Solidity: function teePool() view returns(address)
func (_VanaDlpContract *VanaDlpContractCaller) TeePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "teePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePool is a free data retrieval call binding the contract method 0xd503d4e4.
//
// Solidity: function teePool() view returns(address)
func (_VanaDlpContract *VanaDlpContractSession) TeePool() (common.Address, error) {
	return _VanaDlpContract.Contract.TeePool(&_VanaDlpContract.CallOpts)
}

// TeePool is a free data retrieval call binding the contract method 0xd503d4e4.
//
// Solidity: function teePool() view returns(address)
func (_VanaDlpContract *VanaDlpContractCallerSession) TeePool() (common.Address, error) {
	return _VanaDlpContract.Contract.TeePool(&_VanaDlpContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VanaDlpContract *VanaDlpContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VanaDlpContract *VanaDlpContractSession) Token() (common.Address, error) {
	return _VanaDlpContract.Contract.Token(&_VanaDlpContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VanaDlpContract *VanaDlpContractCallerSession) Token() (common.Address, error) {
	return _VanaDlpContract.Contract.Token(&_VanaDlpContract.CallOpts)
}

// TotalContributorsRewardAmount is a free data retrieval call binding the contract method 0xc41d3b63.
//
// Solidity: function totalContributorsRewardAmount() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractCaller) TotalContributorsRewardAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "totalContributorsRewardAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalContributorsRewardAmount is a free data retrieval call binding the contract method 0xc41d3b63.
//
// Solidity: function totalContributorsRewardAmount() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractSession) TotalContributorsRewardAmount() (*big.Int, error) {
	return _VanaDlpContract.Contract.TotalContributorsRewardAmount(&_VanaDlpContract.CallOpts)
}

// TotalContributorsRewardAmount is a free data retrieval call binding the contract method 0xc41d3b63.
//
// Solidity: function totalContributorsRewardAmount() view returns(uint256)
func (_VanaDlpContract *VanaDlpContractCallerSession) TotalContributorsRewardAmount() (*big.Int, error) {
	return _VanaDlpContract.Contract.TotalContributorsRewardAmount(&_VanaDlpContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_VanaDlpContract *VanaDlpContractCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_VanaDlpContract *VanaDlpContractSession) Version() (*big.Int, error) {
	return _VanaDlpContract.Contract.Version(&_VanaDlpContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_VanaDlpContract *VanaDlpContractCallerSession) Version() (*big.Int, error) {
	return _VanaDlpContract.Contract.Version(&_VanaDlpContract.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VanaDlpContract *VanaDlpContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VanaDlpContract *VanaDlpContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _VanaDlpContract.Contract.AcceptOwnership(&_VanaDlpContract.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VanaDlpContract.Contract.AcceptOwnership(&_VanaDlpContract.TransactOpts)
}

// AddRewardsForContributors is a paid mutator transaction binding the contract method 0x1201c547.
//
// Solidity: function addRewardsForContributors(uint256 contributorsRewardAmount) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) AddRewardsForContributors(opts *bind.TransactOpts, contributorsRewardAmount *big.Int) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "addRewardsForContributors", contributorsRewardAmount)
}

// AddRewardsForContributors is a paid mutator transaction binding the contract method 0x1201c547.
//
// Solidity: function addRewardsForContributors(uint256 contributorsRewardAmount) returns()
func (_VanaDlpContract *VanaDlpContractSession) AddRewardsForContributors(contributorsRewardAmount *big.Int) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.AddRewardsForContributors(&_VanaDlpContract.TransactOpts, contributorsRewardAmount)
}

// AddRewardsForContributors is a paid mutator transaction binding the contract method 0x1201c547.
//
// Solidity: function addRewardsForContributors(uint256 contributorsRewardAmount) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) AddRewardsForContributors(contributorsRewardAmount *big.Int) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.AddRewardsForContributors(&_VanaDlpContract.TransactOpts, contributorsRewardAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x3f76c7cb.
//
// Solidity: function initialize((address,address,address,address,string,string,string,uint256) params) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) Initialize(opts *bind.TransactOpts, params DataLiquidityPoolImplementationInitParams) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "initialize", params)
}

// Initialize is a paid mutator transaction binding the contract method 0x3f76c7cb.
//
// Solidity: function initialize((address,address,address,address,string,string,string,uint256) params) returns()
func (_VanaDlpContract *VanaDlpContractSession) Initialize(params DataLiquidityPoolImplementationInitParams) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.Initialize(&_VanaDlpContract.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x3f76c7cb.
//
// Solidity: function initialize((address,address,address,address,string,string,string,uint256) params) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) Initialize(params DataLiquidityPoolImplementationInitParams) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.Initialize(&_VanaDlpContract.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_VanaDlpContract *VanaDlpContractTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_VanaDlpContract *VanaDlpContractSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.Multicall(&_VanaDlpContract.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_VanaDlpContract *VanaDlpContractTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.Multicall(&_VanaDlpContract.TransactOpts, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VanaDlpContract *VanaDlpContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VanaDlpContract *VanaDlpContractSession) Pause() (*types.Transaction, error) {
	return _VanaDlpContract.Contract.Pause(&_VanaDlpContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) Pause() (*types.Transaction, error) {
	return _VanaDlpContract.Contract.Pause(&_VanaDlpContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VanaDlpContract *VanaDlpContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VanaDlpContract *VanaDlpContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _VanaDlpContract.Contract.RenounceOwnership(&_VanaDlpContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VanaDlpContract.Contract.RenounceOwnership(&_VanaDlpContract.TransactOpts)
}

// RequestReward is a paid mutator transaction binding the contract method 0x5062bc5a.
//
// Solidity: function requestReward(uint256 fileId, uint256 proofIndex) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) RequestReward(opts *bind.TransactOpts, fileId *big.Int, proofIndex *big.Int) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "requestReward", fileId, proofIndex)
}

// RequestReward is a paid mutator transaction binding the contract method 0x5062bc5a.
//
// Solidity: function requestReward(uint256 fileId, uint256 proofIndex) returns()
func (_VanaDlpContract *VanaDlpContractSession) RequestReward(fileId *big.Int, proofIndex *big.Int) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.RequestReward(&_VanaDlpContract.TransactOpts, fileId, proofIndex)
}

// RequestReward is a paid mutator transaction binding the contract method 0x5062bc5a.
//
// Solidity: function requestReward(uint256 fileId, uint256 proofIndex) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) RequestReward(fileId *big.Int, proofIndex *big.Int) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.RequestReward(&_VanaDlpContract.TransactOpts, fileId, proofIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VanaDlpContract *VanaDlpContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.TransferOwnership(&_VanaDlpContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.TransferOwnership(&_VanaDlpContract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VanaDlpContract *VanaDlpContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VanaDlpContract *VanaDlpContractSession) Unpause() (*types.Transaction, error) {
	return _VanaDlpContract.Contract.Unpause(&_VanaDlpContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _VanaDlpContract.Contract.Unpause(&_VanaDlpContract.TransactOpts)
}

// UpdateFileRewardFactor is a paid mutator transaction binding the contract method 0x1a8bcb1d.
//
// Solidity: function updateFileRewardFactor(uint256 newFileRewardFactor) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) UpdateFileRewardFactor(opts *bind.TransactOpts, newFileRewardFactor *big.Int) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "updateFileRewardFactor", newFileRewardFactor)
}

// UpdateFileRewardFactor is a paid mutator transaction binding the contract method 0x1a8bcb1d.
//
// Solidity: function updateFileRewardFactor(uint256 newFileRewardFactor) returns()
func (_VanaDlpContract *VanaDlpContractSession) UpdateFileRewardFactor(newFileRewardFactor *big.Int) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdateFileRewardFactor(&_VanaDlpContract.TransactOpts, newFileRewardFactor)
}

// UpdateFileRewardFactor is a paid mutator transaction binding the contract method 0x1a8bcb1d.
//
// Solidity: function updateFileRewardFactor(uint256 newFileRewardFactor) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) UpdateFileRewardFactor(newFileRewardFactor *big.Int) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdateFileRewardFactor(&_VanaDlpContract.TransactOpts, newFileRewardFactor)
}

// UpdateMasterKey is a paid mutator transaction binding the contract method 0x62baff9f.
//
// Solidity: function updateMasterKey(string newMasterKey) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) UpdateMasterKey(opts *bind.TransactOpts, newMasterKey string) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "updateMasterKey", newMasterKey)
}

// UpdateMasterKey is a paid mutator transaction binding the contract method 0x62baff9f.
//
// Solidity: function updateMasterKey(string newMasterKey) returns()
func (_VanaDlpContract *VanaDlpContractSession) UpdateMasterKey(newMasterKey string) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdateMasterKey(&_VanaDlpContract.TransactOpts, newMasterKey)
}

// UpdateMasterKey is a paid mutator transaction binding the contract method 0x62baff9f.
//
// Solidity: function updateMasterKey(string newMasterKey) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) UpdateMasterKey(newMasterKey string) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdateMasterKey(&_VanaDlpContract.TransactOpts, newMasterKey)
}

// UpdateProofInstruction is a paid mutator transaction binding the contract method 0x7ba24aa5.
//
// Solidity: function updateProofInstruction(string newProofInstruction) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) UpdateProofInstruction(opts *bind.TransactOpts, newProofInstruction string) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "updateProofInstruction", newProofInstruction)
}

// UpdateProofInstruction is a paid mutator transaction binding the contract method 0x7ba24aa5.
//
// Solidity: function updateProofInstruction(string newProofInstruction) returns()
func (_VanaDlpContract *VanaDlpContractSession) UpdateProofInstruction(newProofInstruction string) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdateProofInstruction(&_VanaDlpContract.TransactOpts, newProofInstruction)
}

// UpdateProofInstruction is a paid mutator transaction binding the contract method 0x7ba24aa5.
//
// Solidity: function updateProofInstruction(string newProofInstruction) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) UpdateProofInstruction(newProofInstruction string) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdateProofInstruction(&_VanaDlpContract.TransactOpts, newProofInstruction)
}

// UpdateTeePool is a paid mutator transaction binding the contract method 0x4f33f79d.
//
// Solidity: function updateTeePool(address newTeePool) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) UpdateTeePool(opts *bind.TransactOpts, newTeePool common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "updateTeePool", newTeePool)
}

// UpdateTeePool is a paid mutator transaction binding the contract method 0x4f33f79d.
//
// Solidity: function updateTeePool(address newTeePool) returns()
func (_VanaDlpContract *VanaDlpContractSession) UpdateTeePool(newTeePool common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdateTeePool(&_VanaDlpContract.TransactOpts, newTeePool)
}

// UpdateTeePool is a paid mutator transaction binding the contract method 0x4f33f79d.
//
// Solidity: function updateTeePool(address newTeePool) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) UpdateTeePool(newTeePool common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdateTeePool(&_VanaDlpContract.TransactOpts, newTeePool)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VanaDlpContract *VanaDlpContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VanaDlpContract *VanaDlpContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpgradeToAndCall(&_VanaDlpContract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpgradeToAndCall(&_VanaDlpContract.TransactOpts, newImplementation, data)
}

// VanaDlpContractFileInvalidatedIterator is returned from FilterFileInvalidated and is used to iterate over the raw logs and unpacked data for FileInvalidated events raised by the VanaDlpContract contract.
type VanaDlpContractFileInvalidatedIterator struct {
	Event *VanaDlpContractFileInvalidated // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractFileInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractFileInvalidated)
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
		it.Event = new(VanaDlpContractFileInvalidated)
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
func (it *VanaDlpContractFileInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractFileInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractFileInvalidated represents a FileInvalidated event raised by the VanaDlpContract contract.
type VanaDlpContractFileInvalidated struct {
	FileId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFileInvalidated is a free log retrieval operation binding the contract event 0x4952d59ce08b26bac55ad736c35678f175e3be46988d459a78e1a9d25c9f7fcf.
//
// Solidity: event FileInvalidated(uint256 indexed fileId)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterFileInvalidated(opts *bind.FilterOpts, fileId []*big.Int) (*VanaDlpContractFileInvalidatedIterator, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "FileInvalidated", fileIdRule)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractFileInvalidatedIterator{contract: _VanaDlpContract.contract, event: "FileInvalidated", logs: logs, sub: sub}, nil
}

// WatchFileInvalidated is a free log subscription operation binding the contract event 0x4952d59ce08b26bac55ad736c35678f175e3be46988d459a78e1a9d25c9f7fcf.
//
// Solidity: event FileInvalidated(uint256 indexed fileId)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchFileInvalidated(opts *bind.WatchOpts, sink chan<- *VanaDlpContractFileInvalidated, fileId []*big.Int) (event.Subscription, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "FileInvalidated", fileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractFileInvalidated)
				if err := _VanaDlpContract.contract.UnpackLog(event, "FileInvalidated", log); err != nil {
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

// ParseFileInvalidated is a log parse operation binding the contract event 0x4952d59ce08b26bac55ad736c35678f175e3be46988d459a78e1a9d25c9f7fcf.
//
// Solidity: event FileInvalidated(uint256 indexed fileId)
func (_VanaDlpContract *VanaDlpContractFilterer) ParseFileInvalidated(log types.Log) (*VanaDlpContractFileInvalidated, error) {
	event := new(VanaDlpContractFileInvalidated)
	if err := _VanaDlpContract.contract.UnpackLog(event, "FileInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractFileRewardFactorUpdatedIterator is returned from FilterFileRewardFactorUpdated and is used to iterate over the raw logs and unpacked data for FileRewardFactorUpdated events raised by the VanaDlpContract contract.
type VanaDlpContractFileRewardFactorUpdatedIterator struct {
	Event *VanaDlpContractFileRewardFactorUpdated // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractFileRewardFactorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractFileRewardFactorUpdated)
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
		it.Event = new(VanaDlpContractFileRewardFactorUpdated)
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
func (it *VanaDlpContractFileRewardFactorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractFileRewardFactorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractFileRewardFactorUpdated represents a FileRewardFactorUpdated event raised by the VanaDlpContract contract.
type VanaDlpContractFileRewardFactorUpdated struct {
	NewFileRewardFactor *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFileRewardFactorUpdated is a free log retrieval operation binding the contract event 0x49713882dce7cebaf1e95b21d928f97374855c5adfabf73652b8230de06e4779.
//
// Solidity: event FileRewardFactorUpdated(uint256 newFileRewardFactor)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterFileRewardFactorUpdated(opts *bind.FilterOpts) (*VanaDlpContractFileRewardFactorUpdatedIterator, error) {

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "FileRewardFactorUpdated")
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractFileRewardFactorUpdatedIterator{contract: _VanaDlpContract.contract, event: "FileRewardFactorUpdated", logs: logs, sub: sub}, nil
}

// WatchFileRewardFactorUpdated is a free log subscription operation binding the contract event 0x49713882dce7cebaf1e95b21d928f97374855c5adfabf73652b8230de06e4779.
//
// Solidity: event FileRewardFactorUpdated(uint256 newFileRewardFactor)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchFileRewardFactorUpdated(opts *bind.WatchOpts, sink chan<- *VanaDlpContractFileRewardFactorUpdated) (event.Subscription, error) {

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "FileRewardFactorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractFileRewardFactorUpdated)
				if err := _VanaDlpContract.contract.UnpackLog(event, "FileRewardFactorUpdated", log); err != nil {
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

// ParseFileRewardFactorUpdated is a log parse operation binding the contract event 0x49713882dce7cebaf1e95b21d928f97374855c5adfabf73652b8230de06e4779.
//
// Solidity: event FileRewardFactorUpdated(uint256 newFileRewardFactor)
func (_VanaDlpContract *VanaDlpContractFilterer) ParseFileRewardFactorUpdated(log types.Log) (*VanaDlpContractFileRewardFactorUpdated, error) {
	event := new(VanaDlpContractFileRewardFactorUpdated)
	if err := _VanaDlpContract.contract.UnpackLog(event, "FileRewardFactorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractFileValidatedIterator is returned from FilterFileValidated and is used to iterate over the raw logs and unpacked data for FileValidated events raised by the VanaDlpContract contract.
type VanaDlpContractFileValidatedIterator struct {
	Event *VanaDlpContractFileValidated // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractFileValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractFileValidated)
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
		it.Event = new(VanaDlpContractFileValidated)
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
func (it *VanaDlpContractFileValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractFileValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractFileValidated represents a FileValidated event raised by the VanaDlpContract contract.
type VanaDlpContractFileValidated struct {
	FileId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFileValidated is a free log retrieval operation binding the contract event 0xea71c44822b8c59d3d22cdb2ff9063104576c822412dd4aba7fcc6f066f68c3f.
//
// Solidity: event FileValidated(uint256 indexed fileId)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterFileValidated(opts *bind.FilterOpts, fileId []*big.Int) (*VanaDlpContractFileValidatedIterator, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "FileValidated", fileIdRule)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractFileValidatedIterator{contract: _VanaDlpContract.contract, event: "FileValidated", logs: logs, sub: sub}, nil
}

// WatchFileValidated is a free log subscription operation binding the contract event 0xea71c44822b8c59d3d22cdb2ff9063104576c822412dd4aba7fcc6f066f68c3f.
//
// Solidity: event FileValidated(uint256 indexed fileId)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchFileValidated(opts *bind.WatchOpts, sink chan<- *VanaDlpContractFileValidated, fileId []*big.Int) (event.Subscription, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "FileValidated", fileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractFileValidated)
				if err := _VanaDlpContract.contract.UnpackLog(event, "FileValidated", log); err != nil {
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

// ParseFileValidated is a log parse operation binding the contract event 0xea71c44822b8c59d3d22cdb2ff9063104576c822412dd4aba7fcc6f066f68c3f.
//
// Solidity: event FileValidated(uint256 indexed fileId)
func (_VanaDlpContract *VanaDlpContractFilterer) ParseFileValidated(log types.Log) (*VanaDlpContractFileValidated, error) {
	event := new(VanaDlpContractFileValidated)
	if err := _VanaDlpContract.contract.UnpackLog(event, "FileValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VanaDlpContract contract.
type VanaDlpContractInitializedIterator struct {
	Event *VanaDlpContractInitialized // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractInitialized)
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
		it.Event = new(VanaDlpContractInitialized)
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
func (it *VanaDlpContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractInitialized represents a Initialized event raised by the VanaDlpContract contract.
type VanaDlpContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*VanaDlpContractInitializedIterator, error) {

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractInitializedIterator{contract: _VanaDlpContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VanaDlpContractInitialized) (event.Subscription, error) {

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractInitialized)
				if err := _VanaDlpContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VanaDlpContract *VanaDlpContractFilterer) ParseInitialized(log types.Log) (*VanaDlpContractInitialized, error) {
	event := new(VanaDlpContractInitialized)
	if err := _VanaDlpContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractMasterKeyUpdatedIterator is returned from FilterMasterKeyUpdated and is used to iterate over the raw logs and unpacked data for MasterKeyUpdated events raised by the VanaDlpContract contract.
type VanaDlpContractMasterKeyUpdatedIterator struct {
	Event *VanaDlpContractMasterKeyUpdated // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractMasterKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractMasterKeyUpdated)
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
		it.Event = new(VanaDlpContractMasterKeyUpdated)
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
func (it *VanaDlpContractMasterKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractMasterKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractMasterKeyUpdated represents a MasterKeyUpdated event raised by the VanaDlpContract contract.
type VanaDlpContractMasterKeyUpdated struct {
	NewMasterKey string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMasterKeyUpdated is a free log retrieval operation binding the contract event 0xcf6a9cc9de7d6a2f9488bac950d70cd0eff6bd12d21289198406960d6003de21.
//
// Solidity: event MasterKeyUpdated(string newMasterKey)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterMasterKeyUpdated(opts *bind.FilterOpts) (*VanaDlpContractMasterKeyUpdatedIterator, error) {

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "MasterKeyUpdated")
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractMasterKeyUpdatedIterator{contract: _VanaDlpContract.contract, event: "MasterKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchMasterKeyUpdated is a free log subscription operation binding the contract event 0xcf6a9cc9de7d6a2f9488bac950d70cd0eff6bd12d21289198406960d6003de21.
//
// Solidity: event MasterKeyUpdated(string newMasterKey)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchMasterKeyUpdated(opts *bind.WatchOpts, sink chan<- *VanaDlpContractMasterKeyUpdated) (event.Subscription, error) {

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "MasterKeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractMasterKeyUpdated)
				if err := _VanaDlpContract.contract.UnpackLog(event, "MasterKeyUpdated", log); err != nil {
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

// ParseMasterKeyUpdated is a log parse operation binding the contract event 0xcf6a9cc9de7d6a2f9488bac950d70cd0eff6bd12d21289198406960d6003de21.
//
// Solidity: event MasterKeyUpdated(string newMasterKey)
func (_VanaDlpContract *VanaDlpContractFilterer) ParseMasterKeyUpdated(log types.Log) (*VanaDlpContractMasterKeyUpdated, error) {
	event := new(VanaDlpContractMasterKeyUpdated)
	if err := _VanaDlpContract.contract.UnpackLog(event, "MasterKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the VanaDlpContract contract.
type VanaDlpContractOwnershipTransferStartedIterator struct {
	Event *VanaDlpContractOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractOwnershipTransferStarted)
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
		it.Event = new(VanaDlpContractOwnershipTransferStarted)
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
func (it *VanaDlpContractOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the VanaDlpContract contract.
type VanaDlpContractOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VanaDlpContractOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractOwnershipTransferStartedIterator{contract: _VanaDlpContract.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *VanaDlpContractOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractOwnershipTransferStarted)
				if err := _VanaDlpContract.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_VanaDlpContract *VanaDlpContractFilterer) ParseOwnershipTransferStarted(log types.Log) (*VanaDlpContractOwnershipTransferStarted, error) {
	event := new(VanaDlpContractOwnershipTransferStarted)
	if err := _VanaDlpContract.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VanaDlpContract contract.
type VanaDlpContractOwnershipTransferredIterator struct {
	Event *VanaDlpContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractOwnershipTransferred)
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
		it.Event = new(VanaDlpContractOwnershipTransferred)
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
func (it *VanaDlpContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractOwnershipTransferred represents a OwnershipTransferred event raised by the VanaDlpContract contract.
type VanaDlpContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VanaDlpContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractOwnershipTransferredIterator{contract: _VanaDlpContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VanaDlpContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractOwnershipTransferred)
				if err := _VanaDlpContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VanaDlpContract *VanaDlpContractFilterer) ParseOwnershipTransferred(log types.Log) (*VanaDlpContractOwnershipTransferred, error) {
	event := new(VanaDlpContractOwnershipTransferred)
	if err := _VanaDlpContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VanaDlpContract contract.
type VanaDlpContractPausedIterator struct {
	Event *VanaDlpContractPaused // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractPaused)
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
		it.Event = new(VanaDlpContractPaused)
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
func (it *VanaDlpContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractPaused represents a Paused event raised by the VanaDlpContract contract.
type VanaDlpContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterPaused(opts *bind.FilterOpts) (*VanaDlpContractPausedIterator, error) {

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractPausedIterator{contract: _VanaDlpContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VanaDlpContractPaused) (event.Subscription, error) {

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractPaused)
				if err := _VanaDlpContract.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_VanaDlpContract *VanaDlpContractFilterer) ParsePaused(log types.Log) (*VanaDlpContractPaused, error) {
	event := new(VanaDlpContractPaused)
	if err := _VanaDlpContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractProofInstructionUpdatedIterator is returned from FilterProofInstructionUpdated and is used to iterate over the raw logs and unpacked data for ProofInstructionUpdated events raised by the VanaDlpContract contract.
type VanaDlpContractProofInstructionUpdatedIterator struct {
	Event *VanaDlpContractProofInstructionUpdated // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractProofInstructionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractProofInstructionUpdated)
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
		it.Event = new(VanaDlpContractProofInstructionUpdated)
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
func (it *VanaDlpContractProofInstructionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractProofInstructionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractProofInstructionUpdated represents a ProofInstructionUpdated event raised by the VanaDlpContract contract.
type VanaDlpContractProofInstructionUpdated struct {
	NewProofInstruction string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterProofInstructionUpdated is a free log retrieval operation binding the contract event 0xba416802d8f8c88a69872ea24a6897e9c3b8bdf72487c62ab999954782a87731.
//
// Solidity: event ProofInstructionUpdated(string newProofInstruction)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterProofInstructionUpdated(opts *bind.FilterOpts) (*VanaDlpContractProofInstructionUpdatedIterator, error) {

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "ProofInstructionUpdated")
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractProofInstructionUpdatedIterator{contract: _VanaDlpContract.contract, event: "ProofInstructionUpdated", logs: logs, sub: sub}, nil
}

// WatchProofInstructionUpdated is a free log subscription operation binding the contract event 0xba416802d8f8c88a69872ea24a6897e9c3b8bdf72487c62ab999954782a87731.
//
// Solidity: event ProofInstructionUpdated(string newProofInstruction)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchProofInstructionUpdated(opts *bind.WatchOpts, sink chan<- *VanaDlpContractProofInstructionUpdated) (event.Subscription, error) {

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "ProofInstructionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractProofInstructionUpdated)
				if err := _VanaDlpContract.contract.UnpackLog(event, "ProofInstructionUpdated", log); err != nil {
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

// ParseProofInstructionUpdated is a log parse operation binding the contract event 0xba416802d8f8c88a69872ea24a6897e9c3b8bdf72487c62ab999954782a87731.
//
// Solidity: event ProofInstructionUpdated(string newProofInstruction)
func (_VanaDlpContract *VanaDlpContractFilterer) ParseProofInstructionUpdated(log types.Log) (*VanaDlpContractProofInstructionUpdated, error) {
	event := new(VanaDlpContractProofInstructionUpdated)
	if err := _VanaDlpContract.contract.UnpackLog(event, "ProofInstructionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractRewardRequestedIterator is returned from FilterRewardRequested and is used to iterate over the raw logs and unpacked data for RewardRequested events raised by the VanaDlpContract contract.
type VanaDlpContractRewardRequestedIterator struct {
	Event *VanaDlpContractRewardRequested // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractRewardRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractRewardRequested)
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
		it.Event = new(VanaDlpContractRewardRequested)
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
func (it *VanaDlpContractRewardRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractRewardRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractRewardRequested represents a RewardRequested event raised by the VanaDlpContract contract.
type VanaDlpContractRewardRequested struct {
	ContributorAddress common.Address
	FileId             *big.Int
	ProofIndex         *big.Int
	RewardAmount       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardRequested is a free log retrieval operation binding the contract event 0xc79b2d4186d26cf06f3e1a252e47fb477392229c39da486f16d9285961fdd9fe.
//
// Solidity: event RewardRequested(address indexed contributorAddress, uint256 indexed fileId, uint256 indexed proofIndex, uint256 rewardAmount)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterRewardRequested(opts *bind.FilterOpts, contributorAddress []common.Address, fileId []*big.Int, proofIndex []*big.Int) (*VanaDlpContractRewardRequestedIterator, error) {

	var contributorAddressRule []interface{}
	for _, contributorAddressItem := range contributorAddress {
		contributorAddressRule = append(contributorAddressRule, contributorAddressItem)
	}
	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var proofIndexRule []interface{}
	for _, proofIndexItem := range proofIndex {
		proofIndexRule = append(proofIndexRule, proofIndexItem)
	}

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "RewardRequested", contributorAddressRule, fileIdRule, proofIndexRule)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractRewardRequestedIterator{contract: _VanaDlpContract.contract, event: "RewardRequested", logs: logs, sub: sub}, nil
}

// WatchRewardRequested is a free log subscription operation binding the contract event 0xc79b2d4186d26cf06f3e1a252e47fb477392229c39da486f16d9285961fdd9fe.
//
// Solidity: event RewardRequested(address indexed contributorAddress, uint256 indexed fileId, uint256 indexed proofIndex, uint256 rewardAmount)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchRewardRequested(opts *bind.WatchOpts, sink chan<- *VanaDlpContractRewardRequested, contributorAddress []common.Address, fileId []*big.Int, proofIndex []*big.Int) (event.Subscription, error) {

	var contributorAddressRule []interface{}
	for _, contributorAddressItem := range contributorAddress {
		contributorAddressRule = append(contributorAddressRule, contributorAddressItem)
	}
	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var proofIndexRule []interface{}
	for _, proofIndexItem := range proofIndex {
		proofIndexRule = append(proofIndexRule, proofIndexItem)
	}

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "RewardRequested", contributorAddressRule, fileIdRule, proofIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractRewardRequested)
				if err := _VanaDlpContract.contract.UnpackLog(event, "RewardRequested", log); err != nil {
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

// ParseRewardRequested is a log parse operation binding the contract event 0xc79b2d4186d26cf06f3e1a252e47fb477392229c39da486f16d9285961fdd9fe.
//
// Solidity: event RewardRequested(address indexed contributorAddress, uint256 indexed fileId, uint256 indexed proofIndex, uint256 rewardAmount)
func (_VanaDlpContract *VanaDlpContractFilterer) ParseRewardRequested(log types.Log) (*VanaDlpContractRewardRequested, error) {
	event := new(VanaDlpContractRewardRequested)
	if err := _VanaDlpContract.contract.UnpackLog(event, "RewardRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VanaDlpContract contract.
type VanaDlpContractUnpausedIterator struct {
	Event *VanaDlpContractUnpaused // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractUnpaused)
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
		it.Event = new(VanaDlpContractUnpaused)
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
func (it *VanaDlpContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractUnpaused represents a Unpaused event raised by the VanaDlpContract contract.
type VanaDlpContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VanaDlpContractUnpausedIterator, error) {

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractUnpausedIterator{contract: _VanaDlpContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VanaDlpContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractUnpaused)
				if err := _VanaDlpContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_VanaDlpContract *VanaDlpContractFilterer) ParseUnpaused(log types.Log) (*VanaDlpContractUnpaused, error) {
	event := new(VanaDlpContractUnpaused)
	if err := _VanaDlpContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VanaDlpContract contract.
type VanaDlpContractUpgradedIterator struct {
	Event *VanaDlpContractUpgraded // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractUpgraded)
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
		it.Event = new(VanaDlpContractUpgraded)
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
func (it *VanaDlpContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractUpgraded represents a Upgraded event raised by the VanaDlpContract contract.
type VanaDlpContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VanaDlpContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractUpgradedIterator{contract: _VanaDlpContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VanaDlpContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractUpgraded)
				if err := _VanaDlpContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_VanaDlpContract *VanaDlpContractFilterer) ParseUpgraded(log types.Log) (*VanaDlpContractUpgraded, error) {
	event := new(VanaDlpContractUpgraded)
	if err := _VanaDlpContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
