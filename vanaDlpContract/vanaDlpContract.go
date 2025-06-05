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
	TrustedForwarder    common.Address
	OwnerAddress        common.Address
	TokenAddress        common.Address
	DataRegistryAddress common.Address
	TeePoolAddress      common.Address
	Name                string
	PublicKey           string
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidScore\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"FileInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFileRewardFactor\",\"type\":\"uint256\"}],\"name\":\"FileRewardFactorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"FileValidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newProofInstruction\",\"type\":\"string\"}],\"name\":\"ProofInstructionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newPublicKey\",\"type\":\"string\"}],\"name\":\"PublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proofIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTeePool\",\"type\":\"address\"}],\"name\":\"TeePoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAINTAINER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"contributorsRewardAmount\",\"type\":\"uint256\"}],\"name\":\"addRewardsForContributors\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contributorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"contributorFiles\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIDataLiquidityPool.FileResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contributorAddress\",\"type\":\"address\"}],\"name\":\"contributorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contributorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"filesListCount\",\"type\":\"uint256\"}],\"internalType\":\"structIDataLiquidityPool.ContributorInfoResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"contributors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contributorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"filesListCount\",\"type\":\"uint256\"}],\"internalType\":\"structIDataLiquidityPool.ContributorInfoResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contributorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataRegistry\",\"outputs\":[{\"internalType\":\"contractIDataRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fileRewardFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"files\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIDataLiquidityPool.FileResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"filesListAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"filesListCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trustedForwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dataRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teePoolAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proofInstruction\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fileRewardFactor\",\"type\":\"uint256\"}],\"internalType\":\"structDataLiquidityPoolImplementation.InitParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofInstruction\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofIndex\",\"type\":\"uint256\"}],\"name\":\"requestReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePool\",\"outputs\":[{\"internalType\":\"contractITeePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalContributorsRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFileRewardFactor\",\"type\":\"uint256\"}],\"name\":\"updateFileRewardFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newProofInstruction\",\"type\":\"string\"}],\"name\":\"updateProofInstruction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newPublicKey\",\"type\":\"string\"}],\"name\":\"updatePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTeePool\",\"type\":\"address\"}],\"name\":\"updateTeePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedForwarderAddress\",\"type\":\"address\"}],\"name\":\"updateTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VanaDlpContract.Contract.DEFAULTADMINROLE(&_VanaDlpContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VanaDlpContract.Contract.DEFAULTADMINROLE(&_VanaDlpContract.CallOpts)
}

// MAINTAINERROLE is a free data retrieval call binding the contract method 0xf8742254.
//
// Solidity: function MAINTAINER_ROLE() view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractCaller) MAINTAINERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "MAINTAINER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAINTAINERROLE is a free data retrieval call binding the contract method 0xf8742254.
//
// Solidity: function MAINTAINER_ROLE() view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractSession) MAINTAINERROLE() ([32]byte, error) {
	return _VanaDlpContract.Contract.MAINTAINERROLE(&_VanaDlpContract.CallOpts)
}

// MAINTAINERROLE is a free data retrieval call binding the contract method 0xf8742254.
//
// Solidity: function MAINTAINER_ROLE() view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractCallerSession) MAINTAINERROLE() ([32]byte, error) {
	return _VanaDlpContract.Contract.MAINTAINERROLE(&_VanaDlpContract.CallOpts)
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

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VanaDlpContract.Contract.GetRoleAdmin(&_VanaDlpContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VanaDlpContract *VanaDlpContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VanaDlpContract.Contract.GetRoleAdmin(&_VanaDlpContract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VanaDlpContract *VanaDlpContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VanaDlpContract *VanaDlpContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VanaDlpContract.Contract.HasRole(&_VanaDlpContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VanaDlpContract *VanaDlpContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VanaDlpContract.Contract.HasRole(&_VanaDlpContract.CallOpts, role, account)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_VanaDlpContract *VanaDlpContractCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_VanaDlpContract *VanaDlpContractSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _VanaDlpContract.Contract.IsTrustedForwarder(&_VanaDlpContract.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_VanaDlpContract *VanaDlpContractCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _VanaDlpContract.Contract.IsTrustedForwarder(&_VanaDlpContract.CallOpts, forwarder)
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

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() view returns(string)
func (_VanaDlpContract *VanaDlpContractCaller) PublicKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "publicKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() view returns(string)
func (_VanaDlpContract *VanaDlpContractSession) PublicKey() (string, error) {
	return _VanaDlpContract.Contract.PublicKey(&_VanaDlpContract.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() view returns(string)
func (_VanaDlpContract *VanaDlpContractCallerSession) PublicKey() (string, error) {
	return _VanaDlpContract.Contract.PublicKey(&_VanaDlpContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VanaDlpContract *VanaDlpContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VanaDlpContract *VanaDlpContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VanaDlpContract.Contract.SupportsInterface(&_VanaDlpContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VanaDlpContract *VanaDlpContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VanaDlpContract.Contract.SupportsInterface(&_VanaDlpContract.CallOpts, interfaceId)
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

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_VanaDlpContract *VanaDlpContractCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaDlpContract.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_VanaDlpContract *VanaDlpContractSession) TrustedForwarder() (common.Address, error) {
	return _VanaDlpContract.Contract.TrustedForwarder(&_VanaDlpContract.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_VanaDlpContract *VanaDlpContractCallerSession) TrustedForwarder() (common.Address, error) {
	return _VanaDlpContract.Contract.TrustedForwarder(&_VanaDlpContract.CallOpts)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VanaDlpContract *VanaDlpContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.GrantRole(&_VanaDlpContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.GrantRole(&_VanaDlpContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x3b149013.
//
// Solidity: function initialize((address,address,address,address,address,string,string,string,uint256) params) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) Initialize(opts *bind.TransactOpts, params DataLiquidityPoolImplementationInitParams) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "initialize", params)
}

// Initialize is a paid mutator transaction binding the contract method 0x3b149013.
//
// Solidity: function initialize((address,address,address,address,address,string,string,string,uint256) params) returns()
func (_VanaDlpContract *VanaDlpContractSession) Initialize(params DataLiquidityPoolImplementationInitParams) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.Initialize(&_VanaDlpContract.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x3b149013.
//
// Solidity: function initialize((address,address,address,address,address,string,string,string,uint256) params) returns()
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

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VanaDlpContract *VanaDlpContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.RenounceRole(&_VanaDlpContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.RenounceRole(&_VanaDlpContract.TransactOpts, role, callerConfirmation)
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

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VanaDlpContract *VanaDlpContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.RevokeRole(&_VanaDlpContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.RevokeRole(&_VanaDlpContract.TransactOpts, role, account)
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

// UpdatePublicKey is a paid mutator transaction binding the contract method 0x1dc6fa5e.
//
// Solidity: function updatePublicKey(string newPublicKey) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) UpdatePublicKey(opts *bind.TransactOpts, newPublicKey string) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "updatePublicKey", newPublicKey)
}

// UpdatePublicKey is a paid mutator transaction binding the contract method 0x1dc6fa5e.
//
// Solidity: function updatePublicKey(string newPublicKey) returns()
func (_VanaDlpContract *VanaDlpContractSession) UpdatePublicKey(newPublicKey string) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdatePublicKey(&_VanaDlpContract.TransactOpts, newPublicKey)
}

// UpdatePublicKey is a paid mutator transaction binding the contract method 0x1dc6fa5e.
//
// Solidity: function updatePublicKey(string newPublicKey) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) UpdatePublicKey(newPublicKey string) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdatePublicKey(&_VanaDlpContract.TransactOpts, newPublicKey)
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

// UpdateTrustedForwarder is a paid mutator transaction binding the contract method 0xf90b0311.
//
// Solidity: function updateTrustedForwarder(address trustedForwarderAddress) returns()
func (_VanaDlpContract *VanaDlpContractTransactor) UpdateTrustedForwarder(opts *bind.TransactOpts, trustedForwarderAddress common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.contract.Transact(opts, "updateTrustedForwarder", trustedForwarderAddress)
}

// UpdateTrustedForwarder is a paid mutator transaction binding the contract method 0xf90b0311.
//
// Solidity: function updateTrustedForwarder(address trustedForwarderAddress) returns()
func (_VanaDlpContract *VanaDlpContractSession) UpdateTrustedForwarder(trustedForwarderAddress common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdateTrustedForwarder(&_VanaDlpContract.TransactOpts, trustedForwarderAddress)
}

// UpdateTrustedForwarder is a paid mutator transaction binding the contract method 0xf90b0311.
//
// Solidity: function updateTrustedForwarder(address trustedForwarderAddress) returns()
func (_VanaDlpContract *VanaDlpContractTransactorSession) UpdateTrustedForwarder(trustedForwarderAddress common.Address) (*types.Transaction, error) {
	return _VanaDlpContract.Contract.UpdateTrustedForwarder(&_VanaDlpContract.TransactOpts, trustedForwarderAddress)
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

// VanaDlpContractPublicKeyUpdatedIterator is returned from FilterPublicKeyUpdated and is used to iterate over the raw logs and unpacked data for PublicKeyUpdated events raised by the VanaDlpContract contract.
type VanaDlpContractPublicKeyUpdatedIterator struct {
	Event *VanaDlpContractPublicKeyUpdated // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractPublicKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractPublicKeyUpdated)
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
		it.Event = new(VanaDlpContractPublicKeyUpdated)
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
func (it *VanaDlpContractPublicKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractPublicKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractPublicKeyUpdated represents a PublicKeyUpdated event raised by the VanaDlpContract contract.
type VanaDlpContractPublicKeyUpdated struct {
	NewPublicKey string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyUpdated is a free log retrieval operation binding the contract event 0x0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e.
//
// Solidity: event PublicKeyUpdated(string newPublicKey)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterPublicKeyUpdated(opts *bind.FilterOpts) (*VanaDlpContractPublicKeyUpdatedIterator, error) {

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "PublicKeyUpdated")
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractPublicKeyUpdatedIterator{contract: _VanaDlpContract.contract, event: "PublicKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicKeyUpdated is a free log subscription operation binding the contract event 0x0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e.
//
// Solidity: event PublicKeyUpdated(string newPublicKey)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchPublicKeyUpdated(opts *bind.WatchOpts, sink chan<- *VanaDlpContractPublicKeyUpdated) (event.Subscription, error) {

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "PublicKeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractPublicKeyUpdated)
				if err := _VanaDlpContract.contract.UnpackLog(event, "PublicKeyUpdated", log); err != nil {
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

// ParsePublicKeyUpdated is a log parse operation binding the contract event 0x0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e.
//
// Solidity: event PublicKeyUpdated(string newPublicKey)
func (_VanaDlpContract *VanaDlpContractFilterer) ParsePublicKeyUpdated(log types.Log) (*VanaDlpContractPublicKeyUpdated, error) {
	event := new(VanaDlpContractPublicKeyUpdated)
	if err := _VanaDlpContract.contract.UnpackLog(event, "PublicKeyUpdated", log); err != nil {
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

// VanaDlpContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VanaDlpContract contract.
type VanaDlpContractRoleAdminChangedIterator struct {
	Event *VanaDlpContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractRoleAdminChanged)
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
		it.Event = new(VanaDlpContractRoleAdminChanged)
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
func (it *VanaDlpContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractRoleAdminChanged represents a RoleAdminChanged event raised by the VanaDlpContract contract.
type VanaDlpContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VanaDlpContractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractRoleAdminChangedIterator{contract: _VanaDlpContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VanaDlpContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractRoleAdminChanged)
				if err := _VanaDlpContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VanaDlpContract *VanaDlpContractFilterer) ParseRoleAdminChanged(log types.Log) (*VanaDlpContractRoleAdminChanged, error) {
	event := new(VanaDlpContractRoleAdminChanged)
	if err := _VanaDlpContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VanaDlpContract contract.
type VanaDlpContractRoleGrantedIterator struct {
	Event *VanaDlpContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractRoleGranted)
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
		it.Event = new(VanaDlpContractRoleGranted)
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
func (it *VanaDlpContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractRoleGranted represents a RoleGranted event raised by the VanaDlpContract contract.
type VanaDlpContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VanaDlpContractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractRoleGrantedIterator{contract: _VanaDlpContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VanaDlpContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractRoleGranted)
				if err := _VanaDlpContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VanaDlpContract *VanaDlpContractFilterer) ParseRoleGranted(log types.Log) (*VanaDlpContractRoleGranted, error) {
	event := new(VanaDlpContractRoleGranted)
	if err := _VanaDlpContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VanaDlpContract contract.
type VanaDlpContractRoleRevokedIterator struct {
	Event *VanaDlpContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractRoleRevoked)
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
		it.Event = new(VanaDlpContractRoleRevoked)
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
func (it *VanaDlpContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractRoleRevoked represents a RoleRevoked event raised by the VanaDlpContract contract.
type VanaDlpContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VanaDlpContractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractRoleRevokedIterator{contract: _VanaDlpContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VanaDlpContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractRoleRevoked)
				if err := _VanaDlpContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VanaDlpContract *VanaDlpContractFilterer) ParseRoleRevoked(log types.Log) (*VanaDlpContractRoleRevoked, error) {
	event := new(VanaDlpContractRoleRevoked)
	if err := _VanaDlpContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaDlpContractTeePoolUpdatedIterator is returned from FilterTeePoolUpdated and is used to iterate over the raw logs and unpacked data for TeePoolUpdated events raised by the VanaDlpContract contract.
type VanaDlpContractTeePoolUpdatedIterator struct {
	Event *VanaDlpContractTeePoolUpdated // Event containing the contract specifics and raw log

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
func (it *VanaDlpContractTeePoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaDlpContractTeePoolUpdated)
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
		it.Event = new(VanaDlpContractTeePoolUpdated)
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
func (it *VanaDlpContractTeePoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaDlpContractTeePoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaDlpContractTeePoolUpdated represents a TeePoolUpdated event raised by the VanaDlpContract contract.
type VanaDlpContractTeePoolUpdated struct {
	NewTeePool common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeePoolUpdated is a free log retrieval operation binding the contract event 0x496c74a108597a960a3e602b872823b2e82840913e9b0fc10bc921bc68412b7c.
//
// Solidity: event TeePoolUpdated(address newTeePool)
func (_VanaDlpContract *VanaDlpContractFilterer) FilterTeePoolUpdated(opts *bind.FilterOpts) (*VanaDlpContractTeePoolUpdatedIterator, error) {

	logs, sub, err := _VanaDlpContract.contract.FilterLogs(opts, "TeePoolUpdated")
	if err != nil {
		return nil, err
	}
	return &VanaDlpContractTeePoolUpdatedIterator{contract: _VanaDlpContract.contract, event: "TeePoolUpdated", logs: logs, sub: sub}, nil
}

// WatchTeePoolUpdated is a free log subscription operation binding the contract event 0x496c74a108597a960a3e602b872823b2e82840913e9b0fc10bc921bc68412b7c.
//
// Solidity: event TeePoolUpdated(address newTeePool)
func (_VanaDlpContract *VanaDlpContractFilterer) WatchTeePoolUpdated(opts *bind.WatchOpts, sink chan<- *VanaDlpContractTeePoolUpdated) (event.Subscription, error) {

	logs, sub, err := _VanaDlpContract.contract.WatchLogs(opts, "TeePoolUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaDlpContractTeePoolUpdated)
				if err := _VanaDlpContract.contract.UnpackLog(event, "TeePoolUpdated", log); err != nil {
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

// ParseTeePoolUpdated is a log parse operation binding the contract event 0x496c74a108597a960a3e602b872823b2e82840913e9b0fc10bc921bc68412b7c.
//
// Solidity: event TeePoolUpdated(address newTeePool)
func (_VanaDlpContract *VanaDlpContractFilterer) ParseTeePoolUpdated(log types.Log) (*VanaDlpContractTeePoolUpdated, error) {
	event := new(VanaDlpContractTeePoolUpdated)
	if err := _VanaDlpContract.contract.UnpackLog(event, "TeePoolUpdated", log); err != nil {
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
