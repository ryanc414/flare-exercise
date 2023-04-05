// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package price_submitter

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

// PriceSubmitterMetaData contains all meta data concerning the PriceSubmitter contract.
var PriceSubmitterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"GovernanceCallTimelocked\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"selector\",\"internalType\":\"bytes4\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"encodedCall\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernanceInitialised\",\"inputs\":[{\"type\":\"address\",\"name\":\"initialGovernance\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernedProductionModeEntered\",\"inputs\":[{\"type\":\"address\",\"name\":\"governanceSettings\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HashSubmitted\",\"inputs\":[{\"type\":\"address\",\"name\":\"submitter\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"epochId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"hash\",\"internalType\":\"bytes32\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PricesRevealed\",\"inputs\":[{\"type\":\"address\",\"name\":\"voter\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"epochId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"address[]\",\"name\":\"ftsos\",\"internalType\":\"contractIFtsoGenesis[]\",\"indexed\":false},{\"type\":\"uint256[]\",\"name\":\"prices\",\"internalType\":\"uint256[]\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"random\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimelockedGovernanceCallCanceled\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"selector\",\"internalType\":\"bytes4\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimelockedGovernanceCallExecuted\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"selector\",\"internalType\":\"bytes4\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"MINIMAL_RANDOM\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"RANDOM_EPOCH_CYCLIC_BUFFER_SIZE\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"cancelGovernanceCall\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"_selector\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"executeGovernanceCall\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"_selector\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"_addressUpdater\",\"internalType\":\"address\"}],\"name\":\"getAddressUpdater\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentRandom\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIFtsoManagerGenesis\"}],\"name\":\"getFtsoManager\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIFtsoRegistryGenesis\"}],\"name\":\"getFtsoRegistry\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getRandom\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address[]\",\"name\":\"\",\"internalType\":\"address[]\"}],\"name\":\"getTrustedAddresses\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getVoterWhitelister\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"governance\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIGovernanceSettings\"}],\"name\":\"governanceSettings\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[],\"name\":\"initialise\",\"inputs\":[{\"type\":\"address\",\"name\":\"_governance\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"initialiseFixedAddress\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"productionMode\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"revealPrices\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"},{\"type\":\"uint256[]\",\"name\":\"_ftsoIndices\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"_prices\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256\",\"name\":\"_random\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setAddressUpdater\",\"inputs\":[{\"type\":\"address\",\"name\":\"_addressUpdater\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setTrustedAddresses\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"_trustedAddresses\",\"internalType\":\"address[]\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"submitHash\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"_hash\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"switchToProductionMode\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"encodedCall\",\"internalType\":\"bytes\"}],\"name\":\"timelockedCalls\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateContractAddresses\",\"inputs\":[{\"type\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"internalType\":\"bytes32[]\"},{\"type\":\"address[]\",\"name\":\"_contractAddresses\",\"internalType\":\"address[]\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"voterWhitelistBitmap\",\"inputs\":[{\"type\":\"address\",\"name\":\"_voter\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"voterWhitelisted\",\"inputs\":[{\"type\":\"address\",\"name\":\"_voter\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_ftsoIndex\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"votersRemovedFromWhitelist\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"_removedVoters\",\"internalType\":\"address[]\"},{\"type\":\"uint256\",\"name\":\"_ftsoIndex\",\"internalType\":\"uint256\"}]}]",
}

// PriceSubmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceSubmitterMetaData.ABI instead.
var PriceSubmitterABI = PriceSubmitterMetaData.ABI

// PriceSubmitter is an auto generated Go binding around an Ethereum contract.
type PriceSubmitter struct {
	PriceSubmitterCaller     // Read-only binding to the contract
	PriceSubmitterTransactor // Write-only binding to the contract
	PriceSubmitterFilterer   // Log filterer for contract events
}

// PriceSubmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceSubmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceSubmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceSubmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceSubmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceSubmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceSubmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceSubmitterSession struct {
	Contract     *PriceSubmitter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceSubmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceSubmitterCallerSession struct {
	Contract *PriceSubmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PriceSubmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceSubmitterTransactorSession struct {
	Contract     *PriceSubmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PriceSubmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceSubmitterRaw struct {
	Contract *PriceSubmitter // Generic contract binding to access the raw methods on
}

// PriceSubmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceSubmitterCallerRaw struct {
	Contract *PriceSubmitterCaller // Generic read-only contract binding to access the raw methods on
}

// PriceSubmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceSubmitterTransactorRaw struct {
	Contract *PriceSubmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceSubmitter creates a new instance of PriceSubmitter, bound to a specific deployed contract.
func NewPriceSubmitter(address common.Address, backend bind.ContractBackend) (*PriceSubmitter, error) {
	contract, err := bindPriceSubmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitter{PriceSubmitterCaller: PriceSubmitterCaller{contract: contract}, PriceSubmitterTransactor: PriceSubmitterTransactor{contract: contract}, PriceSubmitterFilterer: PriceSubmitterFilterer{contract: contract}}, nil
}

// NewPriceSubmitterCaller creates a new read-only instance of PriceSubmitter, bound to a specific deployed contract.
func NewPriceSubmitterCaller(address common.Address, caller bind.ContractCaller) (*PriceSubmitterCaller, error) {
	contract, err := bindPriceSubmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterCaller{contract: contract}, nil
}

// NewPriceSubmitterTransactor creates a new write-only instance of PriceSubmitter, bound to a specific deployed contract.
func NewPriceSubmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceSubmitterTransactor, error) {
	contract, err := bindPriceSubmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterTransactor{contract: contract}, nil
}

// NewPriceSubmitterFilterer creates a new log filterer instance of PriceSubmitter, bound to a specific deployed contract.
func NewPriceSubmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceSubmitterFilterer, error) {
	contract, err := bindPriceSubmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterFilterer{contract: contract}, nil
}

// bindPriceSubmitter binds a generic wrapper to an already deployed contract.
func bindPriceSubmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceSubmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceSubmitter *PriceSubmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceSubmitter.Contract.PriceSubmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceSubmitter *PriceSubmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.PriceSubmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceSubmitter *PriceSubmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.PriceSubmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceSubmitter *PriceSubmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceSubmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceSubmitter *PriceSubmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceSubmitter *PriceSubmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.contract.Transact(opts, method, params...)
}

// MINIMALRANDOM is a free data retrieval call binding the contract method 0x3b56f098.
//
// Solidity: function MINIMAL_RANDOM() view returns(uint256)
func (_PriceSubmitter *PriceSubmitterCaller) MINIMALRANDOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "MINIMAL_RANDOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMALRANDOM is a free data retrieval call binding the contract method 0x3b56f098.
//
// Solidity: function MINIMAL_RANDOM() view returns(uint256)
func (_PriceSubmitter *PriceSubmitterSession) MINIMALRANDOM() (*big.Int, error) {
	return _PriceSubmitter.Contract.MINIMALRANDOM(&_PriceSubmitter.CallOpts)
}

// MINIMALRANDOM is a free data retrieval call binding the contract method 0x3b56f098.
//
// Solidity: function MINIMAL_RANDOM() view returns(uint256)
func (_PriceSubmitter *PriceSubmitterCallerSession) MINIMALRANDOM() (*big.Int, error) {
	return _PriceSubmitter.Contract.MINIMALRANDOM(&_PriceSubmitter.CallOpts)
}

// RANDOMEPOCHCYCLICBUFFERSIZE is a free data retrieval call binding the contract method 0xc0156bcc.
//
// Solidity: function RANDOM_EPOCH_CYCLIC_BUFFER_SIZE() view returns(uint256)
func (_PriceSubmitter *PriceSubmitterCaller) RANDOMEPOCHCYCLICBUFFERSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "RANDOM_EPOCH_CYCLIC_BUFFER_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RANDOMEPOCHCYCLICBUFFERSIZE is a free data retrieval call binding the contract method 0xc0156bcc.
//
// Solidity: function RANDOM_EPOCH_CYCLIC_BUFFER_SIZE() view returns(uint256)
func (_PriceSubmitter *PriceSubmitterSession) RANDOMEPOCHCYCLICBUFFERSIZE() (*big.Int, error) {
	return _PriceSubmitter.Contract.RANDOMEPOCHCYCLICBUFFERSIZE(&_PriceSubmitter.CallOpts)
}

// RANDOMEPOCHCYCLICBUFFERSIZE is a free data retrieval call binding the contract method 0xc0156bcc.
//
// Solidity: function RANDOM_EPOCH_CYCLIC_BUFFER_SIZE() view returns(uint256)
func (_PriceSubmitter *PriceSubmitterCallerSession) RANDOMEPOCHCYCLICBUFFERSIZE() (*big.Int, error) {
	return _PriceSubmitter.Contract.RANDOMEPOCHCYCLICBUFFERSIZE(&_PriceSubmitter.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_PriceSubmitter *PriceSubmitterCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_PriceSubmitter *PriceSubmitterSession) GetAddressUpdater() (common.Address, error) {
	return _PriceSubmitter.Contract.GetAddressUpdater(&_PriceSubmitter.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_PriceSubmitter *PriceSubmitterCallerSession) GetAddressUpdater() (common.Address, error) {
	return _PriceSubmitter.Contract.GetAddressUpdater(&_PriceSubmitter.CallOpts)
}

// GetCurrentRandom is a free data retrieval call binding the contract method 0xd89601fd.
//
// Solidity: function getCurrentRandom() view returns(uint256)
func (_PriceSubmitter *PriceSubmitterCaller) GetCurrentRandom(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getCurrentRandom")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRandom is a free data retrieval call binding the contract method 0xd89601fd.
//
// Solidity: function getCurrentRandom() view returns(uint256)
func (_PriceSubmitter *PriceSubmitterSession) GetCurrentRandom() (*big.Int, error) {
	return _PriceSubmitter.Contract.GetCurrentRandom(&_PriceSubmitter.CallOpts)
}

// GetCurrentRandom is a free data retrieval call binding the contract method 0xd89601fd.
//
// Solidity: function getCurrentRandom() view returns(uint256)
func (_PriceSubmitter *PriceSubmitterCallerSession) GetCurrentRandom() (*big.Int, error) {
	return _PriceSubmitter.Contract.GetCurrentRandom(&_PriceSubmitter.CallOpts)
}

// GetFtsoManager is a free data retrieval call binding the contract method 0xb39c6858.
//
// Solidity: function getFtsoManager() view returns(address)
func (_PriceSubmitter *PriceSubmitterCaller) GetFtsoManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getFtsoManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFtsoManager is a free data retrieval call binding the contract method 0xb39c6858.
//
// Solidity: function getFtsoManager() view returns(address)
func (_PriceSubmitter *PriceSubmitterSession) GetFtsoManager() (common.Address, error) {
	return _PriceSubmitter.Contract.GetFtsoManager(&_PriceSubmitter.CallOpts)
}

// GetFtsoManager is a free data retrieval call binding the contract method 0xb39c6858.
//
// Solidity: function getFtsoManager() view returns(address)
func (_PriceSubmitter *PriceSubmitterCallerSession) GetFtsoManager() (common.Address, error) {
	return _PriceSubmitter.Contract.GetFtsoManager(&_PriceSubmitter.CallOpts)
}

// GetFtsoRegistry is a free data retrieval call binding the contract method 0x8c9d28b6.
//
// Solidity: function getFtsoRegistry() view returns(address)
func (_PriceSubmitter *PriceSubmitterCaller) GetFtsoRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getFtsoRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFtsoRegistry is a free data retrieval call binding the contract method 0x8c9d28b6.
//
// Solidity: function getFtsoRegistry() view returns(address)
func (_PriceSubmitter *PriceSubmitterSession) GetFtsoRegistry() (common.Address, error) {
	return _PriceSubmitter.Contract.GetFtsoRegistry(&_PriceSubmitter.CallOpts)
}

// GetFtsoRegistry is a free data retrieval call binding the contract method 0x8c9d28b6.
//
// Solidity: function getFtsoRegistry() view returns(address)
func (_PriceSubmitter *PriceSubmitterCallerSession) GetFtsoRegistry() (common.Address, error) {
	return _PriceSubmitter.Contract.GetFtsoRegistry(&_PriceSubmitter.CallOpts)
}

// GetRandom is a free data retrieval call binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(uint256 _epochId) view returns(uint256)
func (_PriceSubmitter *PriceSubmitterCaller) GetRandom(opts *bind.CallOpts, _epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getRandom", _epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRandom is a free data retrieval call binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(uint256 _epochId) view returns(uint256)
func (_PriceSubmitter *PriceSubmitterSession) GetRandom(_epochId *big.Int) (*big.Int, error) {
	return _PriceSubmitter.Contract.GetRandom(&_PriceSubmitter.CallOpts, _epochId)
}

// GetRandom is a free data retrieval call binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(uint256 _epochId) view returns(uint256)
func (_PriceSubmitter *PriceSubmitterCallerSession) GetRandom(_epochId *big.Int) (*big.Int, error) {
	return _PriceSubmitter.Contract.GetRandom(&_PriceSubmitter.CallOpts, _epochId)
}

// GetTrustedAddresses is a free data retrieval call binding the contract method 0xffacb84e.
//
// Solidity: function getTrustedAddresses() view returns(address[])
func (_PriceSubmitter *PriceSubmitterCaller) GetTrustedAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getTrustedAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTrustedAddresses is a free data retrieval call binding the contract method 0xffacb84e.
//
// Solidity: function getTrustedAddresses() view returns(address[])
func (_PriceSubmitter *PriceSubmitterSession) GetTrustedAddresses() ([]common.Address, error) {
	return _PriceSubmitter.Contract.GetTrustedAddresses(&_PriceSubmitter.CallOpts)
}

// GetTrustedAddresses is a free data retrieval call binding the contract method 0xffacb84e.
//
// Solidity: function getTrustedAddresses() view returns(address[])
func (_PriceSubmitter *PriceSubmitterCallerSession) GetTrustedAddresses() ([]common.Address, error) {
	return _PriceSubmitter.Contract.GetTrustedAddresses(&_PriceSubmitter.CallOpts)
}

// GetVoterWhitelister is a free data retrieval call binding the contract method 0x71e1fad9.
//
// Solidity: function getVoterWhitelister() view returns(address)
func (_PriceSubmitter *PriceSubmitterCaller) GetVoterWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getVoterWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoterWhitelister is a free data retrieval call binding the contract method 0x71e1fad9.
//
// Solidity: function getVoterWhitelister() view returns(address)
func (_PriceSubmitter *PriceSubmitterSession) GetVoterWhitelister() (common.Address, error) {
	return _PriceSubmitter.Contract.GetVoterWhitelister(&_PriceSubmitter.CallOpts)
}

// GetVoterWhitelister is a free data retrieval call binding the contract method 0x71e1fad9.
//
// Solidity: function getVoterWhitelister() view returns(address)
func (_PriceSubmitter *PriceSubmitterCallerSession) GetVoterWhitelister() (common.Address, error) {
	return _PriceSubmitter.Contract.GetVoterWhitelister(&_PriceSubmitter.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PriceSubmitter *PriceSubmitterCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PriceSubmitter *PriceSubmitterSession) Governance() (common.Address, error) {
	return _PriceSubmitter.Contract.Governance(&_PriceSubmitter.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PriceSubmitter *PriceSubmitterCallerSession) Governance() (common.Address, error) {
	return _PriceSubmitter.Contract.Governance(&_PriceSubmitter.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_PriceSubmitter *PriceSubmitterCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_PriceSubmitter *PriceSubmitterSession) GovernanceSettings() (common.Address, error) {
	return _PriceSubmitter.Contract.GovernanceSettings(&_PriceSubmitter.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_PriceSubmitter *PriceSubmitterCallerSession) GovernanceSettings() (common.Address, error) {
	return _PriceSubmitter.Contract.GovernanceSettings(&_PriceSubmitter.CallOpts)
}

// Initialise is a free data retrieval call binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) pure returns()
func (_PriceSubmitter *PriceSubmitterCaller) Initialise(opts *bind.CallOpts, _governance common.Address) error {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "initialise", _governance)

	if err != nil {
		return err
	}

	return err

}

// Initialise is a free data retrieval call binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) pure returns()
func (_PriceSubmitter *PriceSubmitterSession) Initialise(_governance common.Address) error {
	return _PriceSubmitter.Contract.Initialise(&_PriceSubmitter.CallOpts, _governance)
}

// Initialise is a free data retrieval call binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) pure returns()
func (_PriceSubmitter *PriceSubmitterCallerSession) Initialise(_governance common.Address) error {
	return _PriceSubmitter.Contract.Initialise(&_PriceSubmitter.CallOpts, _governance)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_PriceSubmitter *PriceSubmitterCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_PriceSubmitter *PriceSubmitterSession) ProductionMode() (bool, error) {
	return _PriceSubmitter.Contract.ProductionMode(&_PriceSubmitter.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_PriceSubmitter *PriceSubmitterCallerSession) ProductionMode() (bool, error) {
	return _PriceSubmitter.Contract.ProductionMode(&_PriceSubmitter.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_PriceSubmitter *PriceSubmitterCaller) TimelockedCalls(opts *bind.CallOpts, arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "timelockedCalls", arg0)

	outstruct := new(struct {
		AllowedAfterTimestamp *big.Int
		EncodedCall           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowedAfterTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_PriceSubmitter *PriceSubmitterSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _PriceSubmitter.Contract.TimelockedCalls(&_PriceSubmitter.CallOpts, arg0)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_PriceSubmitter *PriceSubmitterCallerSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _PriceSubmitter.Contract.TimelockedCalls(&_PriceSubmitter.CallOpts, arg0)
}

// VoterWhitelistBitmap is a free data retrieval call binding the contract method 0x7ac420ad.
//
// Solidity: function voterWhitelistBitmap(address _voter) view returns(uint256)
func (_PriceSubmitter *PriceSubmitterCaller) VoterWhitelistBitmap(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "voterWhitelistBitmap", _voter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoterWhitelistBitmap is a free data retrieval call binding the contract method 0x7ac420ad.
//
// Solidity: function voterWhitelistBitmap(address _voter) view returns(uint256)
func (_PriceSubmitter *PriceSubmitterSession) VoterWhitelistBitmap(_voter common.Address) (*big.Int, error) {
	return _PriceSubmitter.Contract.VoterWhitelistBitmap(&_PriceSubmitter.CallOpts, _voter)
}

// VoterWhitelistBitmap is a free data retrieval call binding the contract method 0x7ac420ad.
//
// Solidity: function voterWhitelistBitmap(address _voter) view returns(uint256)
func (_PriceSubmitter *PriceSubmitterCallerSession) VoterWhitelistBitmap(_voter common.Address) (*big.Int, error) {
	return _PriceSubmitter.Contract.VoterWhitelistBitmap(&_PriceSubmitter.CallOpts, _voter)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_PriceSubmitter *PriceSubmitterTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_PriceSubmitter *PriceSubmitterSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.CancelGovernanceCall(&_PriceSubmitter.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_PriceSubmitter *PriceSubmitterTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.CancelGovernanceCall(&_PriceSubmitter.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_PriceSubmitter *PriceSubmitterTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_PriceSubmitter *PriceSubmitterSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.ExecuteGovernanceCall(&_PriceSubmitter.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_PriceSubmitter *PriceSubmitterTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.ExecuteGovernanceCall(&_PriceSubmitter.TransactOpts, _selector)
}

// InitialiseFixedAddress is a paid mutator transaction binding the contract method 0xc9f960eb.
//
// Solidity: function initialiseFixedAddress() returns(address)
func (_PriceSubmitter *PriceSubmitterTransactor) InitialiseFixedAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "initialiseFixedAddress")
}

// InitialiseFixedAddress is a paid mutator transaction binding the contract method 0xc9f960eb.
//
// Solidity: function initialiseFixedAddress() returns(address)
func (_PriceSubmitter *PriceSubmitterSession) InitialiseFixedAddress() (*types.Transaction, error) {
	return _PriceSubmitter.Contract.InitialiseFixedAddress(&_PriceSubmitter.TransactOpts)
}

// InitialiseFixedAddress is a paid mutator transaction binding the contract method 0xc9f960eb.
//
// Solidity: function initialiseFixedAddress() returns(address)
func (_PriceSubmitter *PriceSubmitterTransactorSession) InitialiseFixedAddress() (*types.Transaction, error) {
	return _PriceSubmitter.Contract.InitialiseFixedAddress(&_PriceSubmitter.TransactOpts)
}

// RevealPrices is a paid mutator transaction binding the contract method 0xe2db5a52.
//
// Solidity: function revealPrices(uint256 _epochId, uint256[] _ftsoIndices, uint256[] _prices, uint256 _random) returns()
func (_PriceSubmitter *PriceSubmitterTransactor) RevealPrices(opts *bind.TransactOpts, _epochId *big.Int, _ftsoIndices []*big.Int, _prices []*big.Int, _random *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "revealPrices", _epochId, _ftsoIndices, _prices, _random)
}

// RevealPrices is a paid mutator transaction binding the contract method 0xe2db5a52.
//
// Solidity: function revealPrices(uint256 _epochId, uint256[] _ftsoIndices, uint256[] _prices, uint256 _random) returns()
func (_PriceSubmitter *PriceSubmitterSession) RevealPrices(_epochId *big.Int, _ftsoIndices []*big.Int, _prices []*big.Int, _random *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.RevealPrices(&_PriceSubmitter.TransactOpts, _epochId, _ftsoIndices, _prices, _random)
}

// RevealPrices is a paid mutator transaction binding the contract method 0xe2db5a52.
//
// Solidity: function revealPrices(uint256 _epochId, uint256[] _ftsoIndices, uint256[] _prices, uint256 _random) returns()
func (_PriceSubmitter *PriceSubmitterTransactorSession) RevealPrices(_epochId *big.Int, _ftsoIndices []*big.Int, _prices []*big.Int, _random *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.RevealPrices(&_PriceSubmitter.TransactOpts, _epochId, _ftsoIndices, _prices, _random)
}

// SetAddressUpdater is a paid mutator transaction binding the contract method 0xaea36b53.
//
// Solidity: function setAddressUpdater(address _addressUpdater) returns()
func (_PriceSubmitter *PriceSubmitterTransactor) SetAddressUpdater(opts *bind.TransactOpts, _addressUpdater common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "setAddressUpdater", _addressUpdater)
}

// SetAddressUpdater is a paid mutator transaction binding the contract method 0xaea36b53.
//
// Solidity: function setAddressUpdater(address _addressUpdater) returns()
func (_PriceSubmitter *PriceSubmitterSession) SetAddressUpdater(_addressUpdater common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SetAddressUpdater(&_PriceSubmitter.TransactOpts, _addressUpdater)
}

// SetAddressUpdater is a paid mutator transaction binding the contract method 0xaea36b53.
//
// Solidity: function setAddressUpdater(address _addressUpdater) returns()
func (_PriceSubmitter *PriceSubmitterTransactorSession) SetAddressUpdater(_addressUpdater common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SetAddressUpdater(&_PriceSubmitter.TransactOpts, _addressUpdater)
}

// SetTrustedAddresses is a paid mutator transaction binding the contract method 0x9ec2b581.
//
// Solidity: function setTrustedAddresses(address[] _trustedAddresses) returns()
func (_PriceSubmitter *PriceSubmitterTransactor) SetTrustedAddresses(opts *bind.TransactOpts, _trustedAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "setTrustedAddresses", _trustedAddresses)
}

// SetTrustedAddresses is a paid mutator transaction binding the contract method 0x9ec2b581.
//
// Solidity: function setTrustedAddresses(address[] _trustedAddresses) returns()
func (_PriceSubmitter *PriceSubmitterSession) SetTrustedAddresses(_trustedAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SetTrustedAddresses(&_PriceSubmitter.TransactOpts, _trustedAddresses)
}

// SetTrustedAddresses is a paid mutator transaction binding the contract method 0x9ec2b581.
//
// Solidity: function setTrustedAddresses(address[] _trustedAddresses) returns()
func (_PriceSubmitter *PriceSubmitterTransactorSession) SetTrustedAddresses(_trustedAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SetTrustedAddresses(&_PriceSubmitter.TransactOpts, _trustedAddresses)
}

// SubmitHash is a paid mutator transaction binding the contract method 0x8fc6f667.
//
// Solidity: function submitHash(uint256 _epochId, bytes32 _hash) returns()
func (_PriceSubmitter *PriceSubmitterTransactor) SubmitHash(opts *bind.TransactOpts, _epochId *big.Int, _hash [32]byte) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "submitHash", _epochId, _hash)
}

// SubmitHash is a paid mutator transaction binding the contract method 0x8fc6f667.
//
// Solidity: function submitHash(uint256 _epochId, bytes32 _hash) returns()
func (_PriceSubmitter *PriceSubmitterSession) SubmitHash(_epochId *big.Int, _hash [32]byte) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SubmitHash(&_PriceSubmitter.TransactOpts, _epochId, _hash)
}

// SubmitHash is a paid mutator transaction binding the contract method 0x8fc6f667.
//
// Solidity: function submitHash(uint256 _epochId, bytes32 _hash) returns()
func (_PriceSubmitter *PriceSubmitterTransactorSession) SubmitHash(_epochId *big.Int, _hash [32]byte) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SubmitHash(&_PriceSubmitter.TransactOpts, _epochId, _hash)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_PriceSubmitter *PriceSubmitterTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_PriceSubmitter *PriceSubmitterSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SwitchToProductionMode(&_PriceSubmitter.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_PriceSubmitter *PriceSubmitterTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SwitchToProductionMode(&_PriceSubmitter.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_PriceSubmitter *PriceSubmitterTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_PriceSubmitter *PriceSubmitterSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.UpdateContractAddresses(&_PriceSubmitter.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_PriceSubmitter *PriceSubmitterTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.UpdateContractAddresses(&_PriceSubmitter.TransactOpts, _contractNameHashes, _contractAddresses)
}

// VoterWhitelisted is a paid mutator transaction binding the contract method 0x9d986f91.
//
// Solidity: function voterWhitelisted(address _voter, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *PriceSubmitterTransactor) VoterWhitelisted(opts *bind.TransactOpts, _voter common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "voterWhitelisted", _voter, _ftsoIndex)
}

// VoterWhitelisted is a paid mutator transaction binding the contract method 0x9d986f91.
//
// Solidity: function voterWhitelisted(address _voter, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *PriceSubmitterSession) VoterWhitelisted(_voter common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.VoterWhitelisted(&_PriceSubmitter.TransactOpts, _voter, _ftsoIndex)
}

// VoterWhitelisted is a paid mutator transaction binding the contract method 0x9d986f91.
//
// Solidity: function voterWhitelisted(address _voter, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *PriceSubmitterTransactorSession) VoterWhitelisted(_voter common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.VoterWhitelisted(&_PriceSubmitter.TransactOpts, _voter, _ftsoIndex)
}

// VotersRemovedFromWhitelist is a paid mutator transaction binding the contract method 0x76794efb.
//
// Solidity: function votersRemovedFromWhitelist(address[] _removedVoters, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *PriceSubmitterTransactor) VotersRemovedFromWhitelist(opts *bind.TransactOpts, _removedVoters []common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "votersRemovedFromWhitelist", _removedVoters, _ftsoIndex)
}

// VotersRemovedFromWhitelist is a paid mutator transaction binding the contract method 0x76794efb.
//
// Solidity: function votersRemovedFromWhitelist(address[] _removedVoters, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *PriceSubmitterSession) VotersRemovedFromWhitelist(_removedVoters []common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.VotersRemovedFromWhitelist(&_PriceSubmitter.TransactOpts, _removedVoters, _ftsoIndex)
}

// VotersRemovedFromWhitelist is a paid mutator transaction binding the contract method 0x76794efb.
//
// Solidity: function votersRemovedFromWhitelist(address[] _removedVoters, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *PriceSubmitterTransactorSession) VotersRemovedFromWhitelist(_removedVoters []common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.VotersRemovedFromWhitelist(&_PriceSubmitter.TransactOpts, _removedVoters, _ftsoIndex)
}

// PriceSubmitterGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the PriceSubmitter contract.
type PriceSubmitterGovernanceCallTimelockedIterator struct {
	Event *PriceSubmitterGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterGovernanceCallTimelocked)
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
		it.Event = new(PriceSubmitterGovernanceCallTimelocked)
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
func (it *PriceSubmitterGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the PriceSubmitter contract.
type PriceSubmitterGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_PriceSubmitter *PriceSubmitterFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*PriceSubmitterGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterGovernanceCallTimelockedIterator{contract: _PriceSubmitter.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_PriceSubmitter *PriceSubmitterFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *PriceSubmitterGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterGovernanceCallTimelocked)
				if err := _PriceSubmitter.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_PriceSubmitter *PriceSubmitterFilterer) ParseGovernanceCallTimelocked(log types.Log) (*PriceSubmitterGovernanceCallTimelocked, error) {
	event := new(PriceSubmitterGovernanceCallTimelocked)
	if err := _PriceSubmitter.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceSubmitterGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the PriceSubmitter contract.
type PriceSubmitterGovernanceInitialisedIterator struct {
	Event *PriceSubmitterGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterGovernanceInitialised)
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
		it.Event = new(PriceSubmitterGovernanceInitialised)
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
func (it *PriceSubmitterGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterGovernanceInitialised represents a GovernanceInitialised event raised by the PriceSubmitter contract.
type PriceSubmitterGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_PriceSubmitter *PriceSubmitterFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*PriceSubmitterGovernanceInitialisedIterator, error) {

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterGovernanceInitialisedIterator{contract: _PriceSubmitter.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_PriceSubmitter *PriceSubmitterFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *PriceSubmitterGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterGovernanceInitialised)
				if err := _PriceSubmitter.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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

// ParseGovernanceInitialised is a log parse operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_PriceSubmitter *PriceSubmitterFilterer) ParseGovernanceInitialised(log types.Log) (*PriceSubmitterGovernanceInitialised, error) {
	event := new(PriceSubmitterGovernanceInitialised)
	if err := _PriceSubmitter.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceSubmitterGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the PriceSubmitter contract.
type PriceSubmitterGovernedProductionModeEnteredIterator struct {
	Event *PriceSubmitterGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterGovernedProductionModeEntered)
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
		it.Event = new(PriceSubmitterGovernedProductionModeEntered)
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
func (it *PriceSubmitterGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the PriceSubmitter contract.
type PriceSubmitterGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_PriceSubmitter *PriceSubmitterFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*PriceSubmitterGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterGovernedProductionModeEnteredIterator{contract: _PriceSubmitter.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_PriceSubmitter *PriceSubmitterFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *PriceSubmitterGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterGovernedProductionModeEntered)
				if err := _PriceSubmitter.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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

// ParseGovernedProductionModeEntered is a log parse operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_PriceSubmitter *PriceSubmitterFilterer) ParseGovernedProductionModeEntered(log types.Log) (*PriceSubmitterGovernedProductionModeEntered, error) {
	event := new(PriceSubmitterGovernedProductionModeEntered)
	if err := _PriceSubmitter.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceSubmitterHashSubmittedIterator is returned from FilterHashSubmitted and is used to iterate over the raw logs and unpacked data for HashSubmitted events raised by the PriceSubmitter contract.
type PriceSubmitterHashSubmittedIterator struct {
	Event *PriceSubmitterHashSubmitted // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterHashSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterHashSubmitted)
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
		it.Event = new(PriceSubmitterHashSubmitted)
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
func (it *PriceSubmitterHashSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterHashSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterHashSubmitted represents a HashSubmitted event raised by the PriceSubmitter contract.
type PriceSubmitterHashSubmitted struct {
	Submitter common.Address
	EpochId   *big.Int
	Hash      [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHashSubmitted is a free log retrieval operation binding the contract event 0x5e2f64e70eafef31c2f48c8ef140b36406531c36ab0faaede30843202c16f6a8.
//
// Solidity: event HashSubmitted(address indexed submitter, uint256 indexed epochId, bytes32 hash, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) FilterHashSubmitted(opts *bind.FilterOpts, submitter []common.Address, epochId []*big.Int) (*PriceSubmitterHashSubmittedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "HashSubmitted", submitterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterHashSubmittedIterator{contract: _PriceSubmitter.contract, event: "HashSubmitted", logs: logs, sub: sub}, nil
}

// WatchHashSubmitted is a free log subscription operation binding the contract event 0x5e2f64e70eafef31c2f48c8ef140b36406531c36ab0faaede30843202c16f6a8.
//
// Solidity: event HashSubmitted(address indexed submitter, uint256 indexed epochId, bytes32 hash, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) WatchHashSubmitted(opts *bind.WatchOpts, sink chan<- *PriceSubmitterHashSubmitted, submitter []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "HashSubmitted", submitterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterHashSubmitted)
				if err := _PriceSubmitter.contract.UnpackLog(event, "HashSubmitted", log); err != nil {
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

// ParseHashSubmitted is a log parse operation binding the contract event 0x5e2f64e70eafef31c2f48c8ef140b36406531c36ab0faaede30843202c16f6a8.
//
// Solidity: event HashSubmitted(address indexed submitter, uint256 indexed epochId, bytes32 hash, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) ParseHashSubmitted(log types.Log) (*PriceSubmitterHashSubmitted, error) {
	event := new(PriceSubmitterHashSubmitted)
	if err := _PriceSubmitter.contract.UnpackLog(event, "HashSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceSubmitterPricesRevealedIterator is returned from FilterPricesRevealed and is used to iterate over the raw logs and unpacked data for PricesRevealed events raised by the PriceSubmitter contract.
type PriceSubmitterPricesRevealedIterator struct {
	Event *PriceSubmitterPricesRevealed // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterPricesRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterPricesRevealed)
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
		it.Event = new(PriceSubmitterPricesRevealed)
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
func (it *PriceSubmitterPricesRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterPricesRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterPricesRevealed represents a PricesRevealed event raised by the PriceSubmitter contract.
type PriceSubmitterPricesRevealed struct {
	Voter     common.Address
	EpochId   *big.Int
	Ftsos     []common.Address
	Prices    []*big.Int
	Random    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPricesRevealed is a free log retrieval operation binding the contract event 0xafffa539ac1cad89751c875d871abadc6deb7bd51bf6baea004fc71ca0a48fa5.
//
// Solidity: event PricesRevealed(address indexed voter, uint256 indexed epochId, address[] ftsos, uint256[] prices, uint256 random, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) FilterPricesRevealed(opts *bind.FilterOpts, voter []common.Address, epochId []*big.Int) (*PriceSubmitterPricesRevealedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "PricesRevealed", voterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterPricesRevealedIterator{contract: _PriceSubmitter.contract, event: "PricesRevealed", logs: logs, sub: sub}, nil
}

// WatchPricesRevealed is a free log subscription operation binding the contract event 0xafffa539ac1cad89751c875d871abadc6deb7bd51bf6baea004fc71ca0a48fa5.
//
// Solidity: event PricesRevealed(address indexed voter, uint256 indexed epochId, address[] ftsos, uint256[] prices, uint256 random, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) WatchPricesRevealed(opts *bind.WatchOpts, sink chan<- *PriceSubmitterPricesRevealed, voter []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "PricesRevealed", voterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterPricesRevealed)
				if err := _PriceSubmitter.contract.UnpackLog(event, "PricesRevealed", log); err != nil {
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

// ParsePricesRevealed is a log parse operation binding the contract event 0xafffa539ac1cad89751c875d871abadc6deb7bd51bf6baea004fc71ca0a48fa5.
//
// Solidity: event PricesRevealed(address indexed voter, uint256 indexed epochId, address[] ftsos, uint256[] prices, uint256 random, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) ParsePricesRevealed(log types.Log) (*PriceSubmitterPricesRevealed, error) {
	event := new(PriceSubmitterPricesRevealed)
	if err := _PriceSubmitter.contract.UnpackLog(event, "PricesRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceSubmitterTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the PriceSubmitter contract.
type PriceSubmitterTimelockedGovernanceCallCanceledIterator struct {
	Event *PriceSubmitterTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterTimelockedGovernanceCallCanceled)
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
		it.Event = new(PriceSubmitterTimelockedGovernanceCallCanceled)
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
func (it *PriceSubmitterTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the PriceSubmitter contract.
type PriceSubmitterTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*PriceSubmitterTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterTimelockedGovernanceCallCanceledIterator{contract: _PriceSubmitter.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *PriceSubmitterTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterTimelockedGovernanceCallCanceled)
				if err := _PriceSubmitter.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*PriceSubmitterTimelockedGovernanceCallCanceled, error) {
	event := new(PriceSubmitterTimelockedGovernanceCallCanceled)
	if err := _PriceSubmitter.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceSubmitterTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the PriceSubmitter contract.
type PriceSubmitterTimelockedGovernanceCallExecutedIterator struct {
	Event *PriceSubmitterTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterTimelockedGovernanceCallExecuted)
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
		it.Event = new(PriceSubmitterTimelockedGovernanceCallExecuted)
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
func (it *PriceSubmitterTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the PriceSubmitter contract.
type PriceSubmitterTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*PriceSubmitterTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterTimelockedGovernanceCallExecutedIterator{contract: _PriceSubmitter.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *PriceSubmitterTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterTimelockedGovernanceCallExecuted)
				if err := _PriceSubmitter.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_PriceSubmitter *PriceSubmitterFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*PriceSubmitterTimelockedGovernanceCallExecuted, error) {
	event := new(PriceSubmitterTimelockedGovernanceCallExecuted)
	if err := _PriceSubmitter.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
