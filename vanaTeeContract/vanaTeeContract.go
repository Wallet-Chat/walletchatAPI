// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vanaTeeContract

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

// ITeePoolJob is an auto generated low-level Go binding around an user-defined struct.
type ITeePoolJob struct {
	FileId         *big.Int
	BidAmount      *big.Int
	Status         uint8
	AddedTimestamp *big.Int
	OwnerAddress   common.Address
	TeeAddress     common.Address
}

// ITeePoolTeeInfo is an auto generated low-level Go binding around an user-defined struct.
type ITeePoolTeeInfo struct {
	TeeAddress      common.Address
	Url             string
	Status          uint8
	Amount          *big.Int
	WithdrawnAmount *big.Int
	JobsCount       *big.Int
	PublicKey       string
}

// VanaTeeContractMetaData contains all meta data concerning the VanaTeeContract contract.
var VanaTeeContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CancelDelayNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJobStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJobTee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"JobCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActiveTee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotJobOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"JobCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"JobSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attestator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"ProofAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"}],\"name\":\"TeeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"}],\"name\":\"TeeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeTeeList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"activeTeeListAt\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumITeePool.TeeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobsCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"internalType\":\"structITeePool.TeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeTeesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dlpId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proofUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"instruction\",\"type\":\"string\"}],\"internalType\":\"structIDataRegistry.ProofData\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structIDataRegistry.Proof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"addProof\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"addTee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"cancelJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataRegistry\",\"outputs\":[{\"internalType\":\"contractIDataRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"fileJobIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dataRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialCancelDelay\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"}],\"name\":\"isTee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumITeePool.JobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"addedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"}],\"internalType\":\"structITeePool.Job\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"}],\"name\":\"removeTee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"requestContributionProof\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"submitJob\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"teeJobIdsPaginated\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"teeListAt\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumITeePool.TeeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobsCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"internalType\":\"structITeePool.TeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"}],\"name\":\"tees\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumITeePool.TeeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobsCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"internalType\":\"structITeePool.TeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCancelDelay\",\"type\":\"uint256\"}],\"name\":\"updateCancelDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDataRegistry\",\"name\":\"newDataRegistry\",\"type\":\"address\"}],\"name\":\"updateDataRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTeeFee\",\"type\":\"uint256\"}],\"name\":\"updateTeeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// VanaTeeContractABI is the input ABI used to generate the binding from.
// Deprecated: Use VanaTeeContractMetaData.ABI instead.
var VanaTeeContractABI = VanaTeeContractMetaData.ABI

// VanaTeeContract is an auto generated Go binding around an Ethereum contract.
type VanaTeeContract struct {
	VanaTeeContractCaller     // Read-only binding to the contract
	VanaTeeContractTransactor // Write-only binding to the contract
	VanaTeeContractFilterer   // Log filterer for contract events
}

// VanaTeeContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VanaTeeContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VanaTeeContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VanaTeeContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VanaTeeContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VanaTeeContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VanaTeeContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VanaTeeContractSession struct {
	Contract     *VanaTeeContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VanaTeeContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VanaTeeContractCallerSession struct {
	Contract *VanaTeeContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VanaTeeContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VanaTeeContractTransactorSession struct {
	Contract     *VanaTeeContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VanaTeeContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VanaTeeContractRaw struct {
	Contract *VanaTeeContract // Generic contract binding to access the raw methods on
}

// VanaTeeContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VanaTeeContractCallerRaw struct {
	Contract *VanaTeeContractCaller // Generic read-only contract binding to access the raw methods on
}

// VanaTeeContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VanaTeeContractTransactorRaw struct {
	Contract *VanaTeeContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVanaTeeContract creates a new instance of VanaTeeContract, bound to a specific deployed contract.
func NewVanaTeeContract(address common.Address, backend bind.ContractBackend) (*VanaTeeContract, error) {
	contract, err := bindVanaTeeContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContract{VanaTeeContractCaller: VanaTeeContractCaller{contract: contract}, VanaTeeContractTransactor: VanaTeeContractTransactor{contract: contract}, VanaTeeContractFilterer: VanaTeeContractFilterer{contract: contract}}, nil
}

// NewVanaTeeContractCaller creates a new read-only instance of VanaTeeContract, bound to a specific deployed contract.
func NewVanaTeeContractCaller(address common.Address, caller bind.ContractCaller) (*VanaTeeContractCaller, error) {
	contract, err := bindVanaTeeContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractCaller{contract: contract}, nil
}

// NewVanaTeeContractTransactor creates a new write-only instance of VanaTeeContract, bound to a specific deployed contract.
func NewVanaTeeContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VanaTeeContractTransactor, error) {
	contract, err := bindVanaTeeContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractTransactor{contract: contract}, nil
}

// NewVanaTeeContractFilterer creates a new log filterer instance of VanaTeeContract, bound to a specific deployed contract.
func NewVanaTeeContractFilterer(address common.Address, filterer bind.ContractFilterer) (*VanaTeeContractFilterer, error) {
	contract, err := bindVanaTeeContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractFilterer{contract: contract}, nil
}

// bindVanaTeeContract binds a generic wrapper to an already deployed contract.
func bindVanaTeeContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VanaTeeContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VanaTeeContract *VanaTeeContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VanaTeeContract.Contract.VanaTeeContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VanaTeeContract *VanaTeeContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.VanaTeeContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VanaTeeContract *VanaTeeContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.VanaTeeContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VanaTeeContract *VanaTeeContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VanaTeeContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VanaTeeContract *VanaTeeContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VanaTeeContract *VanaTeeContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VanaTeeContract *VanaTeeContractCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VanaTeeContract *VanaTeeContractSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _VanaTeeContract.Contract.UPGRADEINTERFACEVERSION(&_VanaTeeContract.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VanaTeeContract *VanaTeeContractCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _VanaTeeContract.Contract.UPGRADEINTERFACEVERSION(&_VanaTeeContract.CallOpts)
}

// ActiveTeeList is a free data retrieval call binding the contract method 0x10f3cd30.
//
// Solidity: function activeTeeList() view returns(address[])
func (_VanaTeeContract *VanaTeeContractCaller) ActiveTeeList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "activeTeeList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ActiveTeeList is a free data retrieval call binding the contract method 0x10f3cd30.
//
// Solidity: function activeTeeList() view returns(address[])
func (_VanaTeeContract *VanaTeeContractSession) ActiveTeeList() ([]common.Address, error) {
	return _VanaTeeContract.Contract.ActiveTeeList(&_VanaTeeContract.CallOpts)
}

// ActiveTeeList is a free data retrieval call binding the contract method 0x10f3cd30.
//
// Solidity: function activeTeeList() view returns(address[])
func (_VanaTeeContract *VanaTeeContractCallerSession) ActiveTeeList() ([]common.Address, error) {
	return _VanaTeeContract.Contract.ActiveTeeList(&_VanaTeeContract.CallOpts)
}

// ActiveTeeListAt is a free data retrieval call binding the contract method 0xca45a288.
//
// Solidity: function activeTeeListAt(uint256 index) view returns((address,string,uint8,uint256,uint256,uint256,string))
func (_VanaTeeContract *VanaTeeContractCaller) ActiveTeeListAt(opts *bind.CallOpts, index *big.Int) (ITeePoolTeeInfo, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "activeTeeListAt", index)

	if err != nil {
		return *new(ITeePoolTeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeePoolTeeInfo)).(*ITeePoolTeeInfo)

	return out0, err

}

// ActiveTeeListAt is a free data retrieval call binding the contract method 0xca45a288.
//
// Solidity: function activeTeeListAt(uint256 index) view returns((address,string,uint8,uint256,uint256,uint256,string))
func (_VanaTeeContract *VanaTeeContractSession) ActiveTeeListAt(index *big.Int) (ITeePoolTeeInfo, error) {
	return _VanaTeeContract.Contract.ActiveTeeListAt(&_VanaTeeContract.CallOpts, index)
}

// ActiveTeeListAt is a free data retrieval call binding the contract method 0xca45a288.
//
// Solidity: function activeTeeListAt(uint256 index) view returns((address,string,uint8,uint256,uint256,uint256,string))
func (_VanaTeeContract *VanaTeeContractCallerSession) ActiveTeeListAt(index *big.Int) (ITeePoolTeeInfo, error) {
	return _VanaTeeContract.Contract.ActiveTeeListAt(&_VanaTeeContract.CallOpts, index)
}

// ActiveTeesCount is a free data retrieval call binding the contract method 0xd402919f.
//
// Solidity: function activeTeesCount() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractCaller) ActiveTeesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "activeTeesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveTeesCount is a free data retrieval call binding the contract method 0xd402919f.
//
// Solidity: function activeTeesCount() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractSession) ActiveTeesCount() (*big.Int, error) {
	return _VanaTeeContract.Contract.ActiveTeesCount(&_VanaTeeContract.CallOpts)
}

// ActiveTeesCount is a free data retrieval call binding the contract method 0xd402919f.
//
// Solidity: function activeTeesCount() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractCallerSession) ActiveTeesCount() (*big.Int, error) {
	return _VanaTeeContract.Contract.ActiveTeesCount(&_VanaTeeContract.CallOpts)
}

// CancelDelay is a free data retrieval call binding the contract method 0x638a0f09.
//
// Solidity: function cancelDelay() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractCaller) CancelDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "cancelDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CancelDelay is a free data retrieval call binding the contract method 0x638a0f09.
//
// Solidity: function cancelDelay() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractSession) CancelDelay() (*big.Int, error) {
	return _VanaTeeContract.Contract.CancelDelay(&_VanaTeeContract.CallOpts)
}

// CancelDelay is a free data retrieval call binding the contract method 0x638a0f09.
//
// Solidity: function cancelDelay() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractCallerSession) CancelDelay() (*big.Int, error) {
	return _VanaTeeContract.Contract.CancelDelay(&_VanaTeeContract.CallOpts)
}

// DataRegistry is a free data retrieval call binding the contract method 0xa39c1d6b.
//
// Solidity: function dataRegistry() view returns(address)
func (_VanaTeeContract *VanaTeeContractCaller) DataRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "dataRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataRegistry is a free data retrieval call binding the contract method 0xa39c1d6b.
//
// Solidity: function dataRegistry() view returns(address)
func (_VanaTeeContract *VanaTeeContractSession) DataRegistry() (common.Address, error) {
	return _VanaTeeContract.Contract.DataRegistry(&_VanaTeeContract.CallOpts)
}

// DataRegistry is a free data retrieval call binding the contract method 0xa39c1d6b.
//
// Solidity: function dataRegistry() view returns(address)
func (_VanaTeeContract *VanaTeeContractCallerSession) DataRegistry() (common.Address, error) {
	return _VanaTeeContract.Contract.DataRegistry(&_VanaTeeContract.CallOpts)
}

// FileJobIds is a free data retrieval call binding the contract method 0x4c0eb9c9.
//
// Solidity: function fileJobIds(uint256 fileId) view returns(uint256[])
func (_VanaTeeContract *VanaTeeContractCaller) FileJobIds(opts *bind.CallOpts, fileId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "fileJobIds", fileId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// FileJobIds is a free data retrieval call binding the contract method 0x4c0eb9c9.
//
// Solidity: function fileJobIds(uint256 fileId) view returns(uint256[])
func (_VanaTeeContract *VanaTeeContractSession) FileJobIds(fileId *big.Int) ([]*big.Int, error) {
	return _VanaTeeContract.Contract.FileJobIds(&_VanaTeeContract.CallOpts, fileId)
}

// FileJobIds is a free data retrieval call binding the contract method 0x4c0eb9c9.
//
// Solidity: function fileJobIds(uint256 fileId) view returns(uint256[])
func (_VanaTeeContract *VanaTeeContractCallerSession) FileJobIds(fileId *big.Int) ([]*big.Int, error) {
	return _VanaTeeContract.Contract.FileJobIds(&_VanaTeeContract.CallOpts, fileId)
}

// IsTee is a free data retrieval call binding the contract method 0xa7016f70.
//
// Solidity: function isTee(address teeAddress) view returns(bool)
func (_VanaTeeContract *VanaTeeContractCaller) IsTee(opts *bind.CallOpts, teeAddress common.Address) (bool, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "isTee", teeAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTee is a free data retrieval call binding the contract method 0xa7016f70.
//
// Solidity: function isTee(address teeAddress) view returns(bool)
func (_VanaTeeContract *VanaTeeContractSession) IsTee(teeAddress common.Address) (bool, error) {
	return _VanaTeeContract.Contract.IsTee(&_VanaTeeContract.CallOpts, teeAddress)
}

// IsTee is a free data retrieval call binding the contract method 0xa7016f70.
//
// Solidity: function isTee(address teeAddress) view returns(bool)
func (_VanaTeeContract *VanaTeeContractCallerSession) IsTee(teeAddress common.Address) (bool, error) {
	return _VanaTeeContract.Contract.IsTee(&_VanaTeeContract.CallOpts, teeAddress)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 jobId) view returns((uint256,uint256,uint8,uint256,address,address))
func (_VanaTeeContract *VanaTeeContractCaller) Jobs(opts *bind.CallOpts, jobId *big.Int) (ITeePoolJob, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "jobs", jobId)

	if err != nil {
		return *new(ITeePoolJob), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeePoolJob)).(*ITeePoolJob)

	return out0, err

}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 jobId) view returns((uint256,uint256,uint8,uint256,address,address))
func (_VanaTeeContract *VanaTeeContractSession) Jobs(jobId *big.Int) (ITeePoolJob, error) {
	return _VanaTeeContract.Contract.Jobs(&_VanaTeeContract.CallOpts, jobId)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 jobId) view returns((uint256,uint256,uint8,uint256,address,address))
func (_VanaTeeContract *VanaTeeContractCallerSession) Jobs(jobId *big.Int) (ITeePoolJob, error) {
	return _VanaTeeContract.Contract.Jobs(&_VanaTeeContract.CallOpts, jobId)
}

// JobsCount is a free data retrieval call binding the contract method 0xd5fbcf42.
//
// Solidity: function jobsCount() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractCaller) JobsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "jobsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JobsCount is a free data retrieval call binding the contract method 0xd5fbcf42.
//
// Solidity: function jobsCount() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractSession) JobsCount() (*big.Int, error) {
	return _VanaTeeContract.Contract.JobsCount(&_VanaTeeContract.CallOpts)
}

// JobsCount is a free data retrieval call binding the contract method 0xd5fbcf42.
//
// Solidity: function jobsCount() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractCallerSession) JobsCount() (*big.Int, error) {
	return _VanaTeeContract.Contract.JobsCount(&_VanaTeeContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VanaTeeContract *VanaTeeContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VanaTeeContract *VanaTeeContractSession) Owner() (common.Address, error) {
	return _VanaTeeContract.Contract.Owner(&_VanaTeeContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VanaTeeContract *VanaTeeContractCallerSession) Owner() (common.Address, error) {
	return _VanaTeeContract.Contract.Owner(&_VanaTeeContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VanaTeeContract *VanaTeeContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VanaTeeContract *VanaTeeContractSession) Paused() (bool, error) {
	return _VanaTeeContract.Contract.Paused(&_VanaTeeContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VanaTeeContract *VanaTeeContractCallerSession) Paused() (bool, error) {
	return _VanaTeeContract.Contract.Paused(&_VanaTeeContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VanaTeeContract *VanaTeeContractCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VanaTeeContract *VanaTeeContractSession) PendingOwner() (common.Address, error) {
	return _VanaTeeContract.Contract.PendingOwner(&_VanaTeeContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_VanaTeeContract *VanaTeeContractCallerSession) PendingOwner() (common.Address, error) {
	return _VanaTeeContract.Contract.PendingOwner(&_VanaTeeContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VanaTeeContract *VanaTeeContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VanaTeeContract *VanaTeeContractSession) ProxiableUUID() ([32]byte, error) {
	return _VanaTeeContract.Contract.ProxiableUUID(&_VanaTeeContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VanaTeeContract *VanaTeeContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VanaTeeContract.Contract.ProxiableUUID(&_VanaTeeContract.CallOpts)
}

// TeeFee is a free data retrieval call binding the contract method 0xf5d548d8.
//
// Solidity: function teeFee() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractCaller) TeeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "teeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeeFee is a free data retrieval call binding the contract method 0xf5d548d8.
//
// Solidity: function teeFee() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractSession) TeeFee() (*big.Int, error) {
	return _VanaTeeContract.Contract.TeeFee(&_VanaTeeContract.CallOpts)
}

// TeeFee is a free data retrieval call binding the contract method 0xf5d548d8.
//
// Solidity: function teeFee() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractCallerSession) TeeFee() (*big.Int, error) {
	return _VanaTeeContract.Contract.TeeFee(&_VanaTeeContract.CallOpts)
}

// TeeJobIdsPaginated is a free data retrieval call binding the contract method 0x9f1b7a3c.
//
// Solidity: function teeJobIdsPaginated(address teeAddress, uint256 start, uint256 limit) view returns(uint256[])
func (_VanaTeeContract *VanaTeeContractCaller) TeeJobIdsPaginated(opts *bind.CallOpts, teeAddress common.Address, start *big.Int, limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "teeJobIdsPaginated", teeAddress, start, limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TeeJobIdsPaginated is a free data retrieval call binding the contract method 0x9f1b7a3c.
//
// Solidity: function teeJobIdsPaginated(address teeAddress, uint256 start, uint256 limit) view returns(uint256[])
func (_VanaTeeContract *VanaTeeContractSession) TeeJobIdsPaginated(teeAddress common.Address, start *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _VanaTeeContract.Contract.TeeJobIdsPaginated(&_VanaTeeContract.CallOpts, teeAddress, start, limit)
}

// TeeJobIdsPaginated is a free data retrieval call binding the contract method 0x9f1b7a3c.
//
// Solidity: function teeJobIdsPaginated(address teeAddress, uint256 start, uint256 limit) view returns(uint256[])
func (_VanaTeeContract *VanaTeeContractCallerSession) TeeJobIdsPaginated(teeAddress common.Address, start *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _VanaTeeContract.Contract.TeeJobIdsPaginated(&_VanaTeeContract.CallOpts, teeAddress, start, limit)
}

// TeeList is a free data retrieval call binding the contract method 0x5b818d36.
//
// Solidity: function teeList() view returns(address[])
func (_VanaTeeContract *VanaTeeContractCaller) TeeList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "teeList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TeeList is a free data retrieval call binding the contract method 0x5b818d36.
//
// Solidity: function teeList() view returns(address[])
func (_VanaTeeContract *VanaTeeContractSession) TeeList() ([]common.Address, error) {
	return _VanaTeeContract.Contract.TeeList(&_VanaTeeContract.CallOpts)
}

// TeeList is a free data retrieval call binding the contract method 0x5b818d36.
//
// Solidity: function teeList() view returns(address[])
func (_VanaTeeContract *VanaTeeContractCallerSession) TeeList() ([]common.Address, error) {
	return _VanaTeeContract.Contract.TeeList(&_VanaTeeContract.CallOpts)
}

// TeeListAt is a free data retrieval call binding the contract method 0x73c88396.
//
// Solidity: function teeListAt(uint256 index) view returns((address,string,uint8,uint256,uint256,uint256,string))
func (_VanaTeeContract *VanaTeeContractCaller) TeeListAt(opts *bind.CallOpts, index *big.Int) (ITeePoolTeeInfo, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "teeListAt", index)

	if err != nil {
		return *new(ITeePoolTeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeePoolTeeInfo)).(*ITeePoolTeeInfo)

	return out0, err

}

// TeeListAt is a free data retrieval call binding the contract method 0x73c88396.
//
// Solidity: function teeListAt(uint256 index) view returns((address,string,uint8,uint256,uint256,uint256,string))
func (_VanaTeeContract *VanaTeeContractSession) TeeListAt(index *big.Int) (ITeePoolTeeInfo, error) {
	return _VanaTeeContract.Contract.TeeListAt(&_VanaTeeContract.CallOpts, index)
}

// TeeListAt is a free data retrieval call binding the contract method 0x73c88396.
//
// Solidity: function teeListAt(uint256 index) view returns((address,string,uint8,uint256,uint256,uint256,string))
func (_VanaTeeContract *VanaTeeContractCallerSession) TeeListAt(index *big.Int) (ITeePoolTeeInfo, error) {
	return _VanaTeeContract.Contract.TeeListAt(&_VanaTeeContract.CallOpts, index)
}

// Tees is a free data retrieval call binding the contract method 0x961884aa.
//
// Solidity: function tees(address teeAddress) view returns((address,string,uint8,uint256,uint256,uint256,string))
func (_VanaTeeContract *VanaTeeContractCaller) Tees(opts *bind.CallOpts, teeAddress common.Address) (ITeePoolTeeInfo, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "tees", teeAddress)

	if err != nil {
		return *new(ITeePoolTeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeePoolTeeInfo)).(*ITeePoolTeeInfo)

	return out0, err

}

// Tees is a free data retrieval call binding the contract method 0x961884aa.
//
// Solidity: function tees(address teeAddress) view returns((address,string,uint8,uint256,uint256,uint256,string))
func (_VanaTeeContract *VanaTeeContractSession) Tees(teeAddress common.Address) (ITeePoolTeeInfo, error) {
	return _VanaTeeContract.Contract.Tees(&_VanaTeeContract.CallOpts, teeAddress)
}

// Tees is a free data retrieval call binding the contract method 0x961884aa.
//
// Solidity: function tees(address teeAddress) view returns((address,string,uint8,uint256,uint256,uint256,string))
func (_VanaTeeContract *VanaTeeContractCallerSession) Tees(teeAddress common.Address) (ITeePoolTeeInfo, error) {
	return _VanaTeeContract.Contract.Tees(&_VanaTeeContract.CallOpts, teeAddress)
}

// TeesCount is a free data retrieval call binding the contract method 0xeb07b9c6.
//
// Solidity: function teesCount() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractCaller) TeesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "teesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeesCount is a free data retrieval call binding the contract method 0xeb07b9c6.
//
// Solidity: function teesCount() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractSession) TeesCount() (*big.Int, error) {
	return _VanaTeeContract.Contract.TeesCount(&_VanaTeeContract.CallOpts)
}

// TeesCount is a free data retrieval call binding the contract method 0xeb07b9c6.
//
// Solidity: function teesCount() view returns(uint256)
func (_VanaTeeContract *VanaTeeContractCallerSession) TeesCount() (*big.Int, error) {
	return _VanaTeeContract.Contract.TeesCount(&_VanaTeeContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_VanaTeeContract *VanaTeeContractCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VanaTeeContract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_VanaTeeContract *VanaTeeContractSession) Version() (*big.Int, error) {
	return _VanaTeeContract.Contract.Version(&_VanaTeeContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_VanaTeeContract *VanaTeeContractCallerSession) Version() (*big.Int, error) {
	return _VanaTeeContract.Contract.Version(&_VanaTeeContract.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VanaTeeContract *VanaTeeContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VanaTeeContract *VanaTeeContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _VanaTeeContract.Contract.AcceptOwnership(&_VanaTeeContract.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VanaTeeContract.Contract.AcceptOwnership(&_VanaTeeContract.TransactOpts)
}

// AddProof is a paid mutator transaction binding the contract method 0xc26045f7.
//
// Solidity: function addProof(uint256 jobId, (bytes,(uint256,uint256,string,string,string)) proof) payable returns()
func (_VanaTeeContract *VanaTeeContractTransactor) AddProof(opts *bind.TransactOpts, jobId *big.Int, proof IDataRegistryProof) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "addProof", jobId, proof)
}

// AddProof is a paid mutator transaction binding the contract method 0xc26045f7.
//
// Solidity: function addProof(uint256 jobId, (bytes,(uint256,uint256,string,string,string)) proof) payable returns()
func (_VanaTeeContract *VanaTeeContractSession) AddProof(jobId *big.Int, proof IDataRegistryProof) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.AddProof(&_VanaTeeContract.TransactOpts, jobId, proof)
}

// AddProof is a paid mutator transaction binding the contract method 0xc26045f7.
//
// Solidity: function addProof(uint256 jobId, (bytes,(uint256,uint256,string,string,string)) proof) payable returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) AddProof(jobId *big.Int, proof IDataRegistryProof) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.AddProof(&_VanaTeeContract.TransactOpts, jobId, proof)
}

// AddTee is a paid mutator transaction binding the contract method 0x8acfae94.
//
// Solidity: function addTee(address teeAddress, string url, string publicKey) returns()
func (_VanaTeeContract *VanaTeeContractTransactor) AddTee(opts *bind.TransactOpts, teeAddress common.Address, url string, publicKey string) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "addTee", teeAddress, url, publicKey)
}

// AddTee is a paid mutator transaction binding the contract method 0x8acfae94.
//
// Solidity: function addTee(address teeAddress, string url, string publicKey) returns()
func (_VanaTeeContract *VanaTeeContractSession) AddTee(teeAddress common.Address, url string, publicKey string) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.AddTee(&_VanaTeeContract.TransactOpts, teeAddress, url, publicKey)
}

// AddTee is a paid mutator transaction binding the contract method 0x8acfae94.
//
// Solidity: function addTee(address teeAddress, string url, string publicKey) returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) AddTee(teeAddress common.Address, url string, publicKey string) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.AddTee(&_VanaTeeContract.TransactOpts, teeAddress, url, publicKey)
}

// CancelJob is a paid mutator transaction binding the contract method 0x1dffa3dc.
//
// Solidity: function cancelJob(uint256 jobId) returns()
func (_VanaTeeContract *VanaTeeContractTransactor) CancelJob(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "cancelJob", jobId)
}

// CancelJob is a paid mutator transaction binding the contract method 0x1dffa3dc.
//
// Solidity: function cancelJob(uint256 jobId) returns()
func (_VanaTeeContract *VanaTeeContractSession) CancelJob(jobId *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.CancelJob(&_VanaTeeContract.TransactOpts, jobId)
}

// CancelJob is a paid mutator transaction binding the contract method 0x1dffa3dc.
//
// Solidity: function cancelJob(uint256 jobId) returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) CancelJob(jobId *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.CancelJob(&_VanaTeeContract.TransactOpts, jobId)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_VanaTeeContract *VanaTeeContractTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_VanaTeeContract *VanaTeeContractSession) Claim() (*types.Transaction, error) {
	return _VanaTeeContract.Contract.Claim(&_VanaTeeContract.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) Claim() (*types.Transaction, error) {
	return _VanaTeeContract.Contract.Claim(&_VanaTeeContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address ownerAddress, address dataRegistryAddress, uint256 initialCancelDelay) returns()
func (_VanaTeeContract *VanaTeeContractTransactor) Initialize(opts *bind.TransactOpts, ownerAddress common.Address, dataRegistryAddress common.Address, initialCancelDelay *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "initialize", ownerAddress, dataRegistryAddress, initialCancelDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address ownerAddress, address dataRegistryAddress, uint256 initialCancelDelay) returns()
func (_VanaTeeContract *VanaTeeContractSession) Initialize(ownerAddress common.Address, dataRegistryAddress common.Address, initialCancelDelay *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.Initialize(&_VanaTeeContract.TransactOpts, ownerAddress, dataRegistryAddress, initialCancelDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address ownerAddress, address dataRegistryAddress, uint256 initialCancelDelay) returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) Initialize(ownerAddress common.Address, dataRegistryAddress common.Address, initialCancelDelay *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.Initialize(&_VanaTeeContract.TransactOpts, ownerAddress, dataRegistryAddress, initialCancelDelay)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_VanaTeeContract *VanaTeeContractTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_VanaTeeContract *VanaTeeContractSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.Multicall(&_VanaTeeContract.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_VanaTeeContract *VanaTeeContractTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.Multicall(&_VanaTeeContract.TransactOpts, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VanaTeeContract *VanaTeeContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VanaTeeContract *VanaTeeContractSession) Pause() (*types.Transaction, error) {
	return _VanaTeeContract.Contract.Pause(&_VanaTeeContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) Pause() (*types.Transaction, error) {
	return _VanaTeeContract.Contract.Pause(&_VanaTeeContract.TransactOpts)
}

// RemoveTee is a paid mutator transaction binding the contract method 0x288fffac.
//
// Solidity: function removeTee(address teeAddress) returns()
func (_VanaTeeContract *VanaTeeContractTransactor) RemoveTee(opts *bind.TransactOpts, teeAddress common.Address) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "removeTee", teeAddress)
}

// RemoveTee is a paid mutator transaction binding the contract method 0x288fffac.
//
// Solidity: function removeTee(address teeAddress) returns()
func (_VanaTeeContract *VanaTeeContractSession) RemoveTee(teeAddress common.Address) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.RemoveTee(&_VanaTeeContract.TransactOpts, teeAddress)
}

// RemoveTee is a paid mutator transaction binding the contract method 0x288fffac.
//
// Solidity: function removeTee(address teeAddress) returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) RemoveTee(teeAddress common.Address) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.RemoveTee(&_VanaTeeContract.TransactOpts, teeAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VanaTeeContract *VanaTeeContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VanaTeeContract *VanaTeeContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _VanaTeeContract.Contract.RenounceOwnership(&_VanaTeeContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VanaTeeContract.Contract.RenounceOwnership(&_VanaTeeContract.TransactOpts)
}

// RequestContributionProof is a paid mutator transaction binding the contract method 0xd2ef73d9.
//
// Solidity: function requestContributionProof(uint256 fileId) payable returns()
func (_VanaTeeContract *VanaTeeContractTransactor) RequestContributionProof(opts *bind.TransactOpts, fileId *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "requestContributionProof", fileId)
}

// RequestContributionProof is a paid mutator transaction binding the contract method 0xd2ef73d9.
//
// Solidity: function requestContributionProof(uint256 fileId) payable returns()
func (_VanaTeeContract *VanaTeeContractSession) RequestContributionProof(fileId *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.RequestContributionProof(&_VanaTeeContract.TransactOpts, fileId)
}

// RequestContributionProof is a paid mutator transaction binding the contract method 0xd2ef73d9.
//
// Solidity: function requestContributionProof(uint256 fileId) payable returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) RequestContributionProof(fileId *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.RequestContributionProof(&_VanaTeeContract.TransactOpts, fileId)
}

// SubmitJob is a paid mutator transaction binding the contract method 0xe7aa0194.
//
// Solidity: function submitJob(uint256 fileId) payable returns()
func (_VanaTeeContract *VanaTeeContractTransactor) SubmitJob(opts *bind.TransactOpts, fileId *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "submitJob", fileId)
}

// SubmitJob is a paid mutator transaction binding the contract method 0xe7aa0194.
//
// Solidity: function submitJob(uint256 fileId) payable returns()
func (_VanaTeeContract *VanaTeeContractSession) SubmitJob(fileId *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.SubmitJob(&_VanaTeeContract.TransactOpts, fileId)
}

// SubmitJob is a paid mutator transaction binding the contract method 0xe7aa0194.
//
// Solidity: function submitJob(uint256 fileId) payable returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) SubmitJob(fileId *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.SubmitJob(&_VanaTeeContract.TransactOpts, fileId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VanaTeeContract *VanaTeeContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VanaTeeContract *VanaTeeContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.TransferOwnership(&_VanaTeeContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.TransferOwnership(&_VanaTeeContract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VanaTeeContract *VanaTeeContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VanaTeeContract *VanaTeeContractSession) Unpause() (*types.Transaction, error) {
	return _VanaTeeContract.Contract.Unpause(&_VanaTeeContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _VanaTeeContract.Contract.Unpause(&_VanaTeeContract.TransactOpts)
}

// UpdateCancelDelay is a paid mutator transaction binding the contract method 0x576d6d4d.
//
// Solidity: function updateCancelDelay(uint256 newCancelDelay) returns()
func (_VanaTeeContract *VanaTeeContractTransactor) UpdateCancelDelay(opts *bind.TransactOpts, newCancelDelay *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "updateCancelDelay", newCancelDelay)
}

// UpdateCancelDelay is a paid mutator transaction binding the contract method 0x576d6d4d.
//
// Solidity: function updateCancelDelay(uint256 newCancelDelay) returns()
func (_VanaTeeContract *VanaTeeContractSession) UpdateCancelDelay(newCancelDelay *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.UpdateCancelDelay(&_VanaTeeContract.TransactOpts, newCancelDelay)
}

// UpdateCancelDelay is a paid mutator transaction binding the contract method 0x576d6d4d.
//
// Solidity: function updateCancelDelay(uint256 newCancelDelay) returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) UpdateCancelDelay(newCancelDelay *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.UpdateCancelDelay(&_VanaTeeContract.TransactOpts, newCancelDelay)
}

// UpdateDataRegistry is a paid mutator transaction binding the contract method 0x0e2a14a3.
//
// Solidity: function updateDataRegistry(address newDataRegistry) returns()
func (_VanaTeeContract *VanaTeeContractTransactor) UpdateDataRegistry(opts *bind.TransactOpts, newDataRegistry common.Address) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "updateDataRegistry", newDataRegistry)
}

// UpdateDataRegistry is a paid mutator transaction binding the contract method 0x0e2a14a3.
//
// Solidity: function updateDataRegistry(address newDataRegistry) returns()
func (_VanaTeeContract *VanaTeeContractSession) UpdateDataRegistry(newDataRegistry common.Address) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.UpdateDataRegistry(&_VanaTeeContract.TransactOpts, newDataRegistry)
}

// UpdateDataRegistry is a paid mutator transaction binding the contract method 0x0e2a14a3.
//
// Solidity: function updateDataRegistry(address newDataRegistry) returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) UpdateDataRegistry(newDataRegistry common.Address) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.UpdateDataRegistry(&_VanaTeeContract.TransactOpts, newDataRegistry)
}

// UpdateTeeFee is a paid mutator transaction binding the contract method 0x69e8de56.
//
// Solidity: function updateTeeFee(uint256 newTeeFee) returns()
func (_VanaTeeContract *VanaTeeContractTransactor) UpdateTeeFee(opts *bind.TransactOpts, newTeeFee *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "updateTeeFee", newTeeFee)
}

// UpdateTeeFee is a paid mutator transaction binding the contract method 0x69e8de56.
//
// Solidity: function updateTeeFee(uint256 newTeeFee) returns()
func (_VanaTeeContract *VanaTeeContractSession) UpdateTeeFee(newTeeFee *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.UpdateTeeFee(&_VanaTeeContract.TransactOpts, newTeeFee)
}

// UpdateTeeFee is a paid mutator transaction binding the contract method 0x69e8de56.
//
// Solidity: function updateTeeFee(uint256 newTeeFee) returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) UpdateTeeFee(newTeeFee *big.Int) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.UpdateTeeFee(&_VanaTeeContract.TransactOpts, newTeeFee)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VanaTeeContract *VanaTeeContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VanaTeeContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VanaTeeContract *VanaTeeContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.UpgradeToAndCall(&_VanaTeeContract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VanaTeeContract *VanaTeeContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VanaTeeContract.Contract.UpgradeToAndCall(&_VanaTeeContract.TransactOpts, newImplementation, data)
}

// VanaTeeContractClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the VanaTeeContract contract.
type VanaTeeContractClaimedIterator struct {
	Event *VanaTeeContractClaimed // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractClaimed)
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
		it.Event = new(VanaTeeContractClaimed)
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
func (it *VanaTeeContractClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractClaimed represents a Claimed event raised by the VanaTeeContract contract.
type VanaTeeContractClaimed struct {
	TeeAddress common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed teeAddress, uint256 amount)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterClaimed(opts *bind.FilterOpts, teeAddress []common.Address) (*VanaTeeContractClaimedIterator, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "Claimed", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractClaimedIterator{contract: _VanaTeeContract.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed teeAddress, uint256 amount)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *VanaTeeContractClaimed, teeAddress []common.Address) (event.Subscription, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "Claimed", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractClaimed)
				if err := _VanaTeeContract.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed teeAddress, uint256 amount)
func (_VanaTeeContract *VanaTeeContractFilterer) ParseClaimed(log types.Log) (*VanaTeeContractClaimed, error) {
	event := new(VanaTeeContractClaimed)
	if err := _VanaTeeContract.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VanaTeeContract contract.
type VanaTeeContractInitializedIterator struct {
	Event *VanaTeeContractInitialized // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractInitialized)
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
		it.Event = new(VanaTeeContractInitialized)
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
func (it *VanaTeeContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractInitialized represents a Initialized event raised by the VanaTeeContract contract.
type VanaTeeContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*VanaTeeContractInitializedIterator, error) {

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractInitializedIterator{contract: _VanaTeeContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VanaTeeContractInitialized) (event.Subscription, error) {

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractInitialized)
				if err := _VanaTeeContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VanaTeeContract *VanaTeeContractFilterer) ParseInitialized(log types.Log) (*VanaTeeContractInitialized, error) {
	event := new(VanaTeeContractInitialized)
	if err := _VanaTeeContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractJobCanceledIterator is returned from FilterJobCanceled and is used to iterate over the raw logs and unpacked data for JobCanceled events raised by the VanaTeeContract contract.
type VanaTeeContractJobCanceledIterator struct {
	Event *VanaTeeContractJobCanceled // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractJobCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractJobCanceled)
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
		it.Event = new(VanaTeeContractJobCanceled)
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
func (it *VanaTeeContractJobCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractJobCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractJobCanceled represents a JobCanceled event raised by the VanaTeeContract contract.
type VanaTeeContractJobCanceled struct {
	JobId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterJobCanceled is a free log retrieval operation binding the contract event 0x68a66a704ceacc38da3f12c63779e47866d9d72e875ec5d43237777adc666d65.
//
// Solidity: event JobCanceled(uint256 indexed jobId)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterJobCanceled(opts *bind.FilterOpts, jobId []*big.Int) (*VanaTeeContractJobCanceledIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "JobCanceled", jobIdRule)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractJobCanceledIterator{contract: _VanaTeeContract.contract, event: "JobCanceled", logs: logs, sub: sub}, nil
}

// WatchJobCanceled is a free log subscription operation binding the contract event 0x68a66a704ceacc38da3f12c63779e47866d9d72e875ec5d43237777adc666d65.
//
// Solidity: event JobCanceled(uint256 indexed jobId)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchJobCanceled(opts *bind.WatchOpts, sink chan<- *VanaTeeContractJobCanceled, jobId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "JobCanceled", jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractJobCanceled)
				if err := _VanaTeeContract.contract.UnpackLog(event, "JobCanceled", log); err != nil {
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

// ParseJobCanceled is a log parse operation binding the contract event 0x68a66a704ceacc38da3f12c63779e47866d9d72e875ec5d43237777adc666d65.
//
// Solidity: event JobCanceled(uint256 indexed jobId)
func (_VanaTeeContract *VanaTeeContractFilterer) ParseJobCanceled(log types.Log) (*VanaTeeContractJobCanceled, error) {
	event := new(VanaTeeContractJobCanceled)
	if err := _VanaTeeContract.contract.UnpackLog(event, "JobCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractJobSubmittedIterator is returned from FilterJobSubmitted and is used to iterate over the raw logs and unpacked data for JobSubmitted events raised by the VanaTeeContract contract.
type VanaTeeContractJobSubmittedIterator struct {
	Event *VanaTeeContractJobSubmitted // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractJobSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractJobSubmitted)
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
		it.Event = new(VanaTeeContractJobSubmitted)
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
func (it *VanaTeeContractJobSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractJobSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractJobSubmitted represents a JobSubmitted event raised by the VanaTeeContract contract.
type VanaTeeContractJobSubmitted struct {
	JobId      *big.Int
	FileId     *big.Int
	TeeAddress common.Address
	BidAmount  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterJobSubmitted is a free log retrieval operation binding the contract event 0x489e923ac6ba31ef19d89736ac6998000e3fe5b1d1f26174a2c12e02393aa67e.
//
// Solidity: event JobSubmitted(uint256 indexed jobId, uint256 indexed fileId, address teeAddress, uint256 bidAmount)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterJobSubmitted(opts *bind.FilterOpts, jobId []*big.Int, fileId []*big.Int) (*VanaTeeContractJobSubmittedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "JobSubmitted", jobIdRule, fileIdRule)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractJobSubmittedIterator{contract: _VanaTeeContract.contract, event: "JobSubmitted", logs: logs, sub: sub}, nil
}

// WatchJobSubmitted is a free log subscription operation binding the contract event 0x489e923ac6ba31ef19d89736ac6998000e3fe5b1d1f26174a2c12e02393aa67e.
//
// Solidity: event JobSubmitted(uint256 indexed jobId, uint256 indexed fileId, address teeAddress, uint256 bidAmount)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchJobSubmitted(opts *bind.WatchOpts, sink chan<- *VanaTeeContractJobSubmitted, jobId []*big.Int, fileId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "JobSubmitted", jobIdRule, fileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractJobSubmitted)
				if err := _VanaTeeContract.contract.UnpackLog(event, "JobSubmitted", log); err != nil {
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

// ParseJobSubmitted is a log parse operation binding the contract event 0x489e923ac6ba31ef19d89736ac6998000e3fe5b1d1f26174a2c12e02393aa67e.
//
// Solidity: event JobSubmitted(uint256 indexed jobId, uint256 indexed fileId, address teeAddress, uint256 bidAmount)
func (_VanaTeeContract *VanaTeeContractFilterer) ParseJobSubmitted(log types.Log) (*VanaTeeContractJobSubmitted, error) {
	event := new(VanaTeeContractJobSubmitted)
	if err := _VanaTeeContract.contract.UnpackLog(event, "JobSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the VanaTeeContract contract.
type VanaTeeContractOwnershipTransferStartedIterator struct {
	Event *VanaTeeContractOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractOwnershipTransferStarted)
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
		it.Event = new(VanaTeeContractOwnershipTransferStarted)
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
func (it *VanaTeeContractOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the VanaTeeContract contract.
type VanaTeeContractOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VanaTeeContractOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractOwnershipTransferStartedIterator{contract: _VanaTeeContract.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *VanaTeeContractOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractOwnershipTransferStarted)
				if err := _VanaTeeContract.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_VanaTeeContract *VanaTeeContractFilterer) ParseOwnershipTransferStarted(log types.Log) (*VanaTeeContractOwnershipTransferStarted, error) {
	event := new(VanaTeeContractOwnershipTransferStarted)
	if err := _VanaTeeContract.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VanaTeeContract contract.
type VanaTeeContractOwnershipTransferredIterator struct {
	Event *VanaTeeContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractOwnershipTransferred)
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
		it.Event = new(VanaTeeContractOwnershipTransferred)
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
func (it *VanaTeeContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractOwnershipTransferred represents a OwnershipTransferred event raised by the VanaTeeContract contract.
type VanaTeeContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VanaTeeContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractOwnershipTransferredIterator{contract: _VanaTeeContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VanaTeeContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractOwnershipTransferred)
				if err := _VanaTeeContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VanaTeeContract *VanaTeeContractFilterer) ParseOwnershipTransferred(log types.Log) (*VanaTeeContractOwnershipTransferred, error) {
	event := new(VanaTeeContractOwnershipTransferred)
	if err := _VanaTeeContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VanaTeeContract contract.
type VanaTeeContractPausedIterator struct {
	Event *VanaTeeContractPaused // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractPaused)
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
		it.Event = new(VanaTeeContractPaused)
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
func (it *VanaTeeContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractPaused represents a Paused event raised by the VanaTeeContract contract.
type VanaTeeContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterPaused(opts *bind.FilterOpts) (*VanaTeeContractPausedIterator, error) {

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractPausedIterator{contract: _VanaTeeContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VanaTeeContractPaused) (event.Subscription, error) {

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractPaused)
				if err := _VanaTeeContract.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_VanaTeeContract *VanaTeeContractFilterer) ParsePaused(log types.Log) (*VanaTeeContractPaused, error) {
	event := new(VanaTeeContractPaused)
	if err := _VanaTeeContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractProofAddedIterator is returned from FilterProofAdded and is used to iterate over the raw logs and unpacked data for ProofAdded events raised by the VanaTeeContract contract.
type VanaTeeContractProofAddedIterator struct {
	Event *VanaTeeContractProofAdded // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractProofAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractProofAdded)
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
		it.Event = new(VanaTeeContractProofAdded)
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
func (it *VanaTeeContractProofAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractProofAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractProofAdded represents a ProofAdded event raised by the VanaTeeContract contract.
type VanaTeeContractProofAdded struct {
	Attestator common.Address
	JobId      *big.Int
	FileId     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProofAdded is a free log retrieval operation binding the contract event 0x41eb9d6336d189fb4bf444a6b5056628c51aa4405c87f7389a107686e8057263.
//
// Solidity: event ProofAdded(address indexed attestator, uint256 indexed jobId, uint256 indexed fileId)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterProofAdded(opts *bind.FilterOpts, attestator []common.Address, jobId []*big.Int, fileId []*big.Int) (*VanaTeeContractProofAddedIterator, error) {

	var attestatorRule []interface{}
	for _, attestatorItem := range attestator {
		attestatorRule = append(attestatorRule, attestatorItem)
	}
	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "ProofAdded", attestatorRule, jobIdRule, fileIdRule)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractProofAddedIterator{contract: _VanaTeeContract.contract, event: "ProofAdded", logs: logs, sub: sub}, nil
}

// WatchProofAdded is a free log subscription operation binding the contract event 0x41eb9d6336d189fb4bf444a6b5056628c51aa4405c87f7389a107686e8057263.
//
// Solidity: event ProofAdded(address indexed attestator, uint256 indexed jobId, uint256 indexed fileId)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchProofAdded(opts *bind.WatchOpts, sink chan<- *VanaTeeContractProofAdded, attestator []common.Address, jobId []*big.Int, fileId []*big.Int) (event.Subscription, error) {

	var attestatorRule []interface{}
	for _, attestatorItem := range attestator {
		attestatorRule = append(attestatorRule, attestatorItem)
	}
	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "ProofAdded", attestatorRule, jobIdRule, fileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractProofAdded)
				if err := _VanaTeeContract.contract.UnpackLog(event, "ProofAdded", log); err != nil {
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

// ParseProofAdded is a log parse operation binding the contract event 0x41eb9d6336d189fb4bf444a6b5056628c51aa4405c87f7389a107686e8057263.
//
// Solidity: event ProofAdded(address indexed attestator, uint256 indexed jobId, uint256 indexed fileId)
func (_VanaTeeContract *VanaTeeContractFilterer) ParseProofAdded(log types.Log) (*VanaTeeContractProofAdded, error) {
	event := new(VanaTeeContractProofAdded)
	if err := _VanaTeeContract.contract.UnpackLog(event, "ProofAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractTeeAddedIterator is returned from FilterTeeAdded and is used to iterate over the raw logs and unpacked data for TeeAdded events raised by the VanaTeeContract contract.
type VanaTeeContractTeeAddedIterator struct {
	Event *VanaTeeContractTeeAdded // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractTeeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractTeeAdded)
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
		it.Event = new(VanaTeeContractTeeAdded)
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
func (it *VanaTeeContractTeeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractTeeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractTeeAdded represents a TeeAdded event raised by the VanaTeeContract contract.
type VanaTeeContractTeeAdded struct {
	TeeAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeeAdded is a free log retrieval operation binding the contract event 0xcc7ab74a9b11b56487e4c7356320d0be894972f6132884ccd04e04baa284fb79.
//
// Solidity: event TeeAdded(address indexed teeAddress)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterTeeAdded(opts *bind.FilterOpts, teeAddress []common.Address) (*VanaTeeContractTeeAddedIterator, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "TeeAdded", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractTeeAddedIterator{contract: _VanaTeeContract.contract, event: "TeeAdded", logs: logs, sub: sub}, nil
}

// WatchTeeAdded is a free log subscription operation binding the contract event 0xcc7ab74a9b11b56487e4c7356320d0be894972f6132884ccd04e04baa284fb79.
//
// Solidity: event TeeAdded(address indexed teeAddress)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchTeeAdded(opts *bind.WatchOpts, sink chan<- *VanaTeeContractTeeAdded, teeAddress []common.Address) (event.Subscription, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "TeeAdded", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractTeeAdded)
				if err := _VanaTeeContract.contract.UnpackLog(event, "TeeAdded", log); err != nil {
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

// ParseTeeAdded is a log parse operation binding the contract event 0xcc7ab74a9b11b56487e4c7356320d0be894972f6132884ccd04e04baa284fb79.
//
// Solidity: event TeeAdded(address indexed teeAddress)
func (_VanaTeeContract *VanaTeeContractFilterer) ParseTeeAdded(log types.Log) (*VanaTeeContractTeeAdded, error) {
	event := new(VanaTeeContractTeeAdded)
	if err := _VanaTeeContract.contract.UnpackLog(event, "TeeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractTeeRemovedIterator is returned from FilterTeeRemoved and is used to iterate over the raw logs and unpacked data for TeeRemoved events raised by the VanaTeeContract contract.
type VanaTeeContractTeeRemovedIterator struct {
	Event *VanaTeeContractTeeRemoved // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractTeeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractTeeRemoved)
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
		it.Event = new(VanaTeeContractTeeRemoved)
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
func (it *VanaTeeContractTeeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractTeeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractTeeRemoved represents a TeeRemoved event raised by the VanaTeeContract contract.
type VanaTeeContractTeeRemoved struct {
	TeeAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeeRemoved is a free log retrieval operation binding the contract event 0x18a003b29a5d1b1d6f0603956003fdc00a42ea8db3825c24a6bf532d5bae0bdf.
//
// Solidity: event TeeRemoved(address indexed teeAddress)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterTeeRemoved(opts *bind.FilterOpts, teeAddress []common.Address) (*VanaTeeContractTeeRemovedIterator, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "TeeRemoved", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractTeeRemovedIterator{contract: _VanaTeeContract.contract, event: "TeeRemoved", logs: logs, sub: sub}, nil
}

// WatchTeeRemoved is a free log subscription operation binding the contract event 0x18a003b29a5d1b1d6f0603956003fdc00a42ea8db3825c24a6bf532d5bae0bdf.
//
// Solidity: event TeeRemoved(address indexed teeAddress)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchTeeRemoved(opts *bind.WatchOpts, sink chan<- *VanaTeeContractTeeRemoved, teeAddress []common.Address) (event.Subscription, error) {

	var teeAddressRule []interface{}
	for _, teeAddressItem := range teeAddress {
		teeAddressRule = append(teeAddressRule, teeAddressItem)
	}

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "TeeRemoved", teeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractTeeRemoved)
				if err := _VanaTeeContract.contract.UnpackLog(event, "TeeRemoved", log); err != nil {
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

// ParseTeeRemoved is a log parse operation binding the contract event 0x18a003b29a5d1b1d6f0603956003fdc00a42ea8db3825c24a6bf532d5bae0bdf.
//
// Solidity: event TeeRemoved(address indexed teeAddress)
func (_VanaTeeContract *VanaTeeContractFilterer) ParseTeeRemoved(log types.Log) (*VanaTeeContractTeeRemoved, error) {
	event := new(VanaTeeContractTeeRemoved)
	if err := _VanaTeeContract.contract.UnpackLog(event, "TeeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VanaTeeContract contract.
type VanaTeeContractUnpausedIterator struct {
	Event *VanaTeeContractUnpaused // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractUnpaused)
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
		it.Event = new(VanaTeeContractUnpaused)
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
func (it *VanaTeeContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractUnpaused represents a Unpaused event raised by the VanaTeeContract contract.
type VanaTeeContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VanaTeeContractUnpausedIterator, error) {

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractUnpausedIterator{contract: _VanaTeeContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VanaTeeContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractUnpaused)
				if err := _VanaTeeContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_VanaTeeContract *VanaTeeContractFilterer) ParseUnpaused(log types.Log) (*VanaTeeContractUnpaused, error) {
	event := new(VanaTeeContractUnpaused)
	if err := _VanaTeeContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VanaTeeContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VanaTeeContract contract.
type VanaTeeContractUpgradedIterator struct {
	Event *VanaTeeContractUpgraded // Event containing the contract specifics and raw log

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
func (it *VanaTeeContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VanaTeeContractUpgraded)
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
		it.Event = new(VanaTeeContractUpgraded)
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
func (it *VanaTeeContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VanaTeeContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VanaTeeContractUpgraded represents a Upgraded event raised by the VanaTeeContract contract.
type VanaTeeContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VanaTeeContract *VanaTeeContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VanaTeeContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VanaTeeContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VanaTeeContractUpgradedIterator{contract: _VanaTeeContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VanaTeeContract *VanaTeeContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VanaTeeContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VanaTeeContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VanaTeeContractUpgraded)
				if err := _VanaTeeContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_VanaTeeContract *VanaTeeContractFilterer) ParseUpgraded(log types.Log) (*VanaTeeContractUpgraded, error) {
	event := new(VanaTeeContractUpgraded)
	if err := _VanaTeeContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
