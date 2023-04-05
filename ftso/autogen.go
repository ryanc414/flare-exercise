// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ftso

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

// FtsoMetaData contains all meta data concerning the Ftso contract.
var FtsoMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"string\",\"name\":\"_symbol\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"_decimals\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_priceSubmitter\",\"internalType\":\"contractIPriceSubmitter\"},{\"type\":\"address\",\"name\":\"_wNat\",\"internalType\":\"contractIIVPToken\"},{\"type\":\"address\",\"name\":\"_ftsoManager\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_firstEpochStartTs\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_submitPeriodSeconds\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_revealPeriodSeconds\",\"internalType\":\"uint256\"},{\"type\":\"uint128\",\"name\":\"_initialPriceUSD\",\"internalType\":\"uint128\"},{\"type\":\"uint256\",\"name\":\"_priceDeviationThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_cyclicBufferSize\",\"internalType\":\"uint256\"}]},{\"type\":\"event\",\"name\":\"LowTurnout\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epochId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"natTurnout\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"lowNatTurnoutThresholdBIPS\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceEpochInitializedOnFtso\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epochId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"endTime\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFinalized\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epochId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"price\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bool\",\"name\":\"rewardedFtso\",\"internalType\":\"bool\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"lowRewardPrice\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"highRewardPrice\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint8\",\"name\":\"finalizationType\",\"internalType\":\"enumIFtso.PriceFinalizationType\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceRevealed\",\"inputs\":[{\"type\":\"address\",\"name\":\"voter\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"epochId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"price\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"votePowerNat\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"votePowerAsset\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"ASSET_PRICE_USD_DECIMALS\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"activateFtso\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_firstEpochStartTs\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_submitPeriodSeconds\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_revealPeriodSeconds\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"active\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIIFtso\"}],\"name\":\"assetFtsos\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIIVPToken\"}],\"name\":\"assets\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"configureEpochs\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_maxVotePowerNatThresholdFraction\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_maxVotePowerAssetThresholdFraction\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_lowAssetUSDThreshold\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_highAssetUSDThreshold\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_highAssetTurnoutThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_lowNatTurnoutThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"address[]\",\"name\":\"_trustedAddresses\",\"internalType\":\"address[]\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"deactivateFtso\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_maxVotePowerNatThresholdFraction\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_maxVotePowerAssetThresholdFraction\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_lowAssetUSDThreshold\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_highAssetUSDThreshold\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_highAssetTurnoutThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_lowNatTurnoutThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"address[]\",\"name\":\"_trustedAddresses\",\"internalType\":\"address[]\"}],\"name\":\"epochsConfiguration\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"fallbackFinalizePriceEpoch\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"address[]\",\"name\":\"_eligibleAddresses\",\"internalType\":\"address[]\"},{\"type\":\"uint256[]\",\"name\":\"_natWeights\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256\",\"name\":\"_natWeightsSum\",\"internalType\":\"uint256\"}],\"name\":\"finalizePriceEpoch\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"_returnRewardData\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"forceFinalizePriceEpoch\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"ftsoManager\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIIVPToken\"}],\"name\":\"getAsset\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address[]\",\"name\":\"\",\"internalType\":\"contractIIFtso[]\"}],\"name\":\"getAssetFtsos\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentEpochId\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_timestamp\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_priceTimestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"_priceFinalizationType\",\"internalType\":\"enumIFtso.PriceFinalizationType\"},{\"type\":\"uint256\",\"name\":\"_lastPriceEpochFinalizationTimestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"_lastPriceEpochFinalizationType\",\"internalType\":\"enumIFtso.PriceFinalizationType\"}],\"name\":\"getCurrentPriceDetails\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_timestamp\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentPriceFromTrustedProviders\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_timestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_assetPriceUsdDecimals\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentPriceWithDecimals\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_timestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_assetPriceUsdDecimals\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentPriceWithDecimalsFromTrustedProviders\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentRandom\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getEpochId\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_timestamp\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getEpochPrice\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getEpochPriceForVoter\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_voter\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_firstEpochStartTs\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_submitPeriodSeconds\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_revealPeriodSeconds\",\"internalType\":\"uint256\"}],\"name\":\"getPriceEpochConfiguration\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_epochSubmitEndTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_epochRevealEndTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_votePowerBlock\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"_fallbackMode\",\"internalType\":\"bool\"}],\"name\":\"getPriceEpochData\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getRandom\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address[]\",\"name\":\"_assets\",\"internalType\":\"contractIIVPToken[]\"},{\"type\":\"uint256[]\",\"name\":\"_assetMultipliers\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256\",\"name\":\"_totalVotePowerNat\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_totalVotePowerAsset\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_assetWeightRatio\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_votePowerBlock\",\"internalType\":\"uint256\"}],\"name\":\"getVoteWeightingParameters\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initializeCurrentEpochStateForReveal\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_circulatingSupplyNat\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"_fallbackMode\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"priceDeviationThresholdBIPS\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"priceEpochCyclicBufferSize\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIPriceSubmitter\"}],\"name\":\"priceSubmitter\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"revealPriceSubmitter\",\"inputs\":[{\"type\":\"address\",\"name\":\"_voter\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_voterWNatVP\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setAsset\",\"inputs\":[{\"type\":\"address\",\"name\":\"_asset\",\"internalType\":\"contractIIVPToken\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setAssetFtsos\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"_assetFtsos\",\"internalType\":\"contractIIFtso[]\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setVotePowerBlock\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_votePowerBlock\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateInitialPrice\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_initialPriceUSD\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_initialPriceTimestamp\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIIVPToken\"}],\"name\":\"wNat\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"wNatVotePowerCached\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_epochId\",\"internalType\":\"uint256\"}]}]",
}

// FtsoABI is the input ABI used to generate the binding from.
// Deprecated: Use FtsoMetaData.ABI instead.
var FtsoABI = FtsoMetaData.ABI

// Ftso is an auto generated Go binding around an Ethereum contract.
type Ftso struct {
	FtsoCaller     // Read-only binding to the contract
	FtsoTransactor // Write-only binding to the contract
	FtsoFilterer   // Log filterer for contract events
}

// FtsoCaller is an auto generated read-only Go binding around an Ethereum contract.
type FtsoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtsoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FtsoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtsoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FtsoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtsoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FtsoSession struct {
	Contract     *Ftso             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FtsoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FtsoCallerSession struct {
	Contract *FtsoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FtsoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FtsoTransactorSession struct {
	Contract     *FtsoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FtsoRaw is an auto generated low-level Go binding around an Ethereum contract.
type FtsoRaw struct {
	Contract *Ftso // Generic contract binding to access the raw methods on
}

// FtsoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FtsoCallerRaw struct {
	Contract *FtsoCaller // Generic read-only contract binding to access the raw methods on
}

// FtsoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FtsoTransactorRaw struct {
	Contract *FtsoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFtso creates a new instance of Ftso, bound to a specific deployed contract.
func NewFtso(address common.Address, backend bind.ContractBackend) (*Ftso, error) {
	contract, err := bindFtso(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ftso{FtsoCaller: FtsoCaller{contract: contract}, FtsoTransactor: FtsoTransactor{contract: contract}, FtsoFilterer: FtsoFilterer{contract: contract}}, nil
}

// NewFtsoCaller creates a new read-only instance of Ftso, bound to a specific deployed contract.
func NewFtsoCaller(address common.Address, caller bind.ContractCaller) (*FtsoCaller, error) {
	contract, err := bindFtso(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FtsoCaller{contract: contract}, nil
}

// NewFtsoTransactor creates a new write-only instance of Ftso, bound to a specific deployed contract.
func NewFtsoTransactor(address common.Address, transactor bind.ContractTransactor) (*FtsoTransactor, error) {
	contract, err := bindFtso(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FtsoTransactor{contract: contract}, nil
}

// NewFtsoFilterer creates a new log filterer instance of Ftso, bound to a specific deployed contract.
func NewFtsoFilterer(address common.Address, filterer bind.ContractFilterer) (*FtsoFilterer, error) {
	contract, err := bindFtso(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FtsoFilterer{contract: contract}, nil
}

// bindFtso binds a generic wrapper to an already deployed contract.
func bindFtso(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FtsoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ftso *FtsoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ftso.Contract.FtsoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ftso *FtsoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ftso.Contract.FtsoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ftso *FtsoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ftso.Contract.FtsoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ftso *FtsoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ftso.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ftso *FtsoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ftso.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ftso *FtsoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ftso.Contract.contract.Transact(opts, method, params...)
}

// ASSETPRICEUSDDECIMALS is a free data retrieval call binding the contract method 0x6b52b242.
//
// Solidity: function ASSET_PRICE_USD_DECIMALS() view returns(uint256)
func (_Ftso *FtsoCaller) ASSETPRICEUSDDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "ASSET_PRICE_USD_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ASSETPRICEUSDDECIMALS is a free data retrieval call binding the contract method 0x6b52b242.
//
// Solidity: function ASSET_PRICE_USD_DECIMALS() view returns(uint256)
func (_Ftso *FtsoSession) ASSETPRICEUSDDECIMALS() (*big.Int, error) {
	return _Ftso.Contract.ASSETPRICEUSDDECIMALS(&_Ftso.CallOpts)
}

// ASSETPRICEUSDDECIMALS is a free data retrieval call binding the contract method 0x6b52b242.
//
// Solidity: function ASSET_PRICE_USD_DECIMALS() view returns(uint256)
func (_Ftso *FtsoCallerSession) ASSETPRICEUSDDECIMALS() (*big.Int, error) {
	return _Ftso.Contract.ASSETPRICEUSDDECIMALS(&_Ftso.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Ftso *FtsoCaller) Active(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "active")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Ftso *FtsoSession) Active() (bool, error) {
	return _Ftso.Contract.Active(&_Ftso.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Ftso *FtsoCallerSession) Active() (bool, error) {
	return _Ftso.Contract.Active(&_Ftso.CallOpts)
}

// AssetFtsos is a free data retrieval call binding the contract method 0x826cc76b.
//
// Solidity: function assetFtsos(uint256 ) view returns(address)
func (_Ftso *FtsoCaller) AssetFtsos(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "assetFtsos", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetFtsos is a free data retrieval call binding the contract method 0x826cc76b.
//
// Solidity: function assetFtsos(uint256 ) view returns(address)
func (_Ftso *FtsoSession) AssetFtsos(arg0 *big.Int) (common.Address, error) {
	return _Ftso.Contract.AssetFtsos(&_Ftso.CallOpts, arg0)
}

// AssetFtsos is a free data retrieval call binding the contract method 0x826cc76b.
//
// Solidity: function assetFtsos(uint256 ) view returns(address)
func (_Ftso *FtsoCallerSession) AssetFtsos(arg0 *big.Int) (common.Address, error) {
	return _Ftso.Contract.AssetFtsos(&_Ftso.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_Ftso *FtsoCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "assets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_Ftso *FtsoSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Ftso.Contract.Assets(&_Ftso.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_Ftso *FtsoCallerSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Ftso.Contract.Assets(&_Ftso.CallOpts, arg0)
}

// EpochsConfiguration is a free data retrieval call binding the contract method 0xe3749e0c.
//
// Solidity: function epochsConfiguration() view returns(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, address[] _trustedAddresses)
func (_Ftso *FtsoCaller) EpochsConfiguration(opts *bind.CallOpts) (struct {
	MaxVotePowerNatThresholdFraction   *big.Int
	MaxVotePowerAssetThresholdFraction *big.Int
	LowAssetUSDThreshold               *big.Int
	HighAssetUSDThreshold              *big.Int
	HighAssetTurnoutThresholdBIPS      *big.Int
	LowNatTurnoutThresholdBIPS         *big.Int
	TrustedAddresses                   []common.Address
}, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "epochsConfiguration")

	outstruct := new(struct {
		MaxVotePowerNatThresholdFraction   *big.Int
		MaxVotePowerAssetThresholdFraction *big.Int
		LowAssetUSDThreshold               *big.Int
		HighAssetUSDThreshold              *big.Int
		HighAssetTurnoutThresholdBIPS      *big.Int
		LowNatTurnoutThresholdBIPS         *big.Int
		TrustedAddresses                   []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxVotePowerNatThresholdFraction = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxVotePowerAssetThresholdFraction = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LowAssetUSDThreshold = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.HighAssetUSDThreshold = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.HighAssetTurnoutThresholdBIPS = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LowNatTurnoutThresholdBIPS = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TrustedAddresses = *abi.ConvertType(out[6], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// EpochsConfiguration is a free data retrieval call binding the contract method 0xe3749e0c.
//
// Solidity: function epochsConfiguration() view returns(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, address[] _trustedAddresses)
func (_Ftso *FtsoSession) EpochsConfiguration() (struct {
	MaxVotePowerNatThresholdFraction   *big.Int
	MaxVotePowerAssetThresholdFraction *big.Int
	LowAssetUSDThreshold               *big.Int
	HighAssetUSDThreshold              *big.Int
	HighAssetTurnoutThresholdBIPS      *big.Int
	LowNatTurnoutThresholdBIPS         *big.Int
	TrustedAddresses                   []common.Address
}, error) {
	return _Ftso.Contract.EpochsConfiguration(&_Ftso.CallOpts)
}

// EpochsConfiguration is a free data retrieval call binding the contract method 0xe3749e0c.
//
// Solidity: function epochsConfiguration() view returns(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, address[] _trustedAddresses)
func (_Ftso *FtsoCallerSession) EpochsConfiguration() (struct {
	MaxVotePowerNatThresholdFraction   *big.Int
	MaxVotePowerAssetThresholdFraction *big.Int
	LowAssetUSDThreshold               *big.Int
	HighAssetUSDThreshold              *big.Int
	HighAssetTurnoutThresholdBIPS      *big.Int
	LowNatTurnoutThresholdBIPS         *big.Int
	TrustedAddresses                   []common.Address
}, error) {
	return _Ftso.Contract.EpochsConfiguration(&_Ftso.CallOpts)
}

// FtsoManager is a free data retrieval call binding the contract method 0x11a7aaaa.
//
// Solidity: function ftsoManager() view returns(address)
func (_Ftso *FtsoCaller) FtsoManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "ftsoManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FtsoManager is a free data retrieval call binding the contract method 0x11a7aaaa.
//
// Solidity: function ftsoManager() view returns(address)
func (_Ftso *FtsoSession) FtsoManager() (common.Address, error) {
	return _Ftso.Contract.FtsoManager(&_Ftso.CallOpts)
}

// FtsoManager is a free data retrieval call binding the contract method 0x11a7aaaa.
//
// Solidity: function ftsoManager() view returns(address)
func (_Ftso *FtsoCallerSession) FtsoManager() (common.Address, error) {
	return _Ftso.Contract.FtsoManager(&_Ftso.CallOpts)
}

// GetAsset is a free data retrieval call binding the contract method 0x5c222bad.
//
// Solidity: function getAsset() view returns(address)
func (_Ftso *FtsoCaller) GetAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAsset is a free data retrieval call binding the contract method 0x5c222bad.
//
// Solidity: function getAsset() view returns(address)
func (_Ftso *FtsoSession) GetAsset() (common.Address, error) {
	return _Ftso.Contract.GetAsset(&_Ftso.CallOpts)
}

// GetAsset is a free data retrieval call binding the contract method 0x5c222bad.
//
// Solidity: function getAsset() view returns(address)
func (_Ftso *FtsoCallerSession) GetAsset() (common.Address, error) {
	return _Ftso.Contract.GetAsset(&_Ftso.CallOpts)
}

// GetAssetFtsos is a free data retrieval call binding the contract method 0x18931c35.
//
// Solidity: function getAssetFtsos() view returns(address[])
func (_Ftso *FtsoCaller) GetAssetFtsos(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getAssetFtsos")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetFtsos is a free data retrieval call binding the contract method 0x18931c35.
//
// Solidity: function getAssetFtsos() view returns(address[])
func (_Ftso *FtsoSession) GetAssetFtsos() ([]common.Address, error) {
	return _Ftso.Contract.GetAssetFtsos(&_Ftso.CallOpts)
}

// GetAssetFtsos is a free data retrieval call binding the contract method 0x18931c35.
//
// Solidity: function getAssetFtsos() view returns(address[])
func (_Ftso *FtsoCallerSession) GetAssetFtsos() ([]common.Address, error) {
	return _Ftso.Contract.GetAssetFtsos(&_Ftso.CallOpts)
}

// GetCurrentEpochId is a free data retrieval call binding the contract method 0xa29a839f.
//
// Solidity: function getCurrentEpochId() view returns(uint256)
func (_Ftso *FtsoCaller) GetCurrentEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getCurrentEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpochId is a free data retrieval call binding the contract method 0xa29a839f.
//
// Solidity: function getCurrentEpochId() view returns(uint256)
func (_Ftso *FtsoSession) GetCurrentEpochId() (*big.Int, error) {
	return _Ftso.Contract.GetCurrentEpochId(&_Ftso.CallOpts)
}

// GetCurrentEpochId is a free data retrieval call binding the contract method 0xa29a839f.
//
// Solidity: function getCurrentEpochId() view returns(uint256)
func (_Ftso *FtsoCallerSession) GetCurrentEpochId() (*big.Int, error) {
	return _Ftso.Contract.GetCurrentEpochId(&_Ftso.CallOpts)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xeb91d37e.
//
// Solidity: function getCurrentPrice() view returns(uint256 _price, uint256 _timestamp)
func (_Ftso *FtsoCaller) GetCurrentPrice(opts *bind.CallOpts) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getCurrentPrice")

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xeb91d37e.
//
// Solidity: function getCurrentPrice() view returns(uint256 _price, uint256 _timestamp)
func (_Ftso *FtsoSession) GetCurrentPrice() (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _Ftso.Contract.GetCurrentPrice(&_Ftso.CallOpts)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xeb91d37e.
//
// Solidity: function getCurrentPrice() view returns(uint256 _price, uint256 _timestamp)
func (_Ftso *FtsoCallerSession) GetCurrentPrice() (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _Ftso.Contract.GetCurrentPrice(&_Ftso.CallOpts)
}

// GetCurrentPriceDetails is a free data retrieval call binding the contract method 0x040d73b8.
//
// Solidity: function getCurrentPriceDetails() view returns(uint256 _price, uint256 _priceTimestamp, uint8 _priceFinalizationType, uint256 _lastPriceEpochFinalizationTimestamp, uint8 _lastPriceEpochFinalizationType)
func (_Ftso *FtsoCaller) GetCurrentPriceDetails(opts *bind.CallOpts) (struct {
	Price                               *big.Int
	PriceTimestamp                      *big.Int
	PriceFinalizationType               uint8
	LastPriceEpochFinalizationTimestamp *big.Int
	LastPriceEpochFinalizationType      uint8
}, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getCurrentPriceDetails")

	outstruct := new(struct {
		Price                               *big.Int
		PriceTimestamp                      *big.Int
		PriceFinalizationType               uint8
		LastPriceEpochFinalizationTimestamp *big.Int
		LastPriceEpochFinalizationType      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PriceTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PriceFinalizationType = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.LastPriceEpochFinalizationTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastPriceEpochFinalizationType = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetCurrentPriceDetails is a free data retrieval call binding the contract method 0x040d73b8.
//
// Solidity: function getCurrentPriceDetails() view returns(uint256 _price, uint256 _priceTimestamp, uint8 _priceFinalizationType, uint256 _lastPriceEpochFinalizationTimestamp, uint8 _lastPriceEpochFinalizationType)
func (_Ftso *FtsoSession) GetCurrentPriceDetails() (struct {
	Price                               *big.Int
	PriceTimestamp                      *big.Int
	PriceFinalizationType               uint8
	LastPriceEpochFinalizationTimestamp *big.Int
	LastPriceEpochFinalizationType      uint8
}, error) {
	return _Ftso.Contract.GetCurrentPriceDetails(&_Ftso.CallOpts)
}

// GetCurrentPriceDetails is a free data retrieval call binding the contract method 0x040d73b8.
//
// Solidity: function getCurrentPriceDetails() view returns(uint256 _price, uint256 _priceTimestamp, uint8 _priceFinalizationType, uint256 _lastPriceEpochFinalizationTimestamp, uint8 _lastPriceEpochFinalizationType)
func (_Ftso *FtsoCallerSession) GetCurrentPriceDetails() (struct {
	Price                               *big.Int
	PriceTimestamp                      *big.Int
	PriceFinalizationType               uint8
	LastPriceEpochFinalizationTimestamp *big.Int
	LastPriceEpochFinalizationType      uint8
}, error) {
	return _Ftso.Contract.GetCurrentPriceDetails(&_Ftso.CallOpts)
}

// GetCurrentPriceFromTrustedProviders is a free data retrieval call binding the contract method 0xaf52df08.
//
// Solidity: function getCurrentPriceFromTrustedProviders() view returns(uint256 _price, uint256 _timestamp)
func (_Ftso *FtsoCaller) GetCurrentPriceFromTrustedProviders(opts *bind.CallOpts) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getCurrentPriceFromTrustedProviders")

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentPriceFromTrustedProviders is a free data retrieval call binding the contract method 0xaf52df08.
//
// Solidity: function getCurrentPriceFromTrustedProviders() view returns(uint256 _price, uint256 _timestamp)
func (_Ftso *FtsoSession) GetCurrentPriceFromTrustedProviders() (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _Ftso.Contract.GetCurrentPriceFromTrustedProviders(&_Ftso.CallOpts)
}

// GetCurrentPriceFromTrustedProviders is a free data retrieval call binding the contract method 0xaf52df08.
//
// Solidity: function getCurrentPriceFromTrustedProviders() view returns(uint256 _price, uint256 _timestamp)
func (_Ftso *FtsoCallerSession) GetCurrentPriceFromTrustedProviders() (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _Ftso.Contract.GetCurrentPriceFromTrustedProviders(&_Ftso.CallOpts)
}

// GetCurrentPriceWithDecimals is a free data retrieval call binding the contract method 0x65f5cd86.
//
// Solidity: function getCurrentPriceWithDecimals() view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_Ftso *FtsoCaller) GetCurrentPriceWithDecimals(opts *bind.CallOpts) (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getCurrentPriceWithDecimals")

	outstruct := new(struct {
		Price                 *big.Int
		Timestamp             *big.Int
		AssetPriceUsdDecimals *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AssetPriceUsdDecimals = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentPriceWithDecimals is a free data retrieval call binding the contract method 0x65f5cd86.
//
// Solidity: function getCurrentPriceWithDecimals() view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_Ftso *FtsoSession) GetCurrentPriceWithDecimals() (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	return _Ftso.Contract.GetCurrentPriceWithDecimals(&_Ftso.CallOpts)
}

// GetCurrentPriceWithDecimals is a free data retrieval call binding the contract method 0x65f5cd86.
//
// Solidity: function getCurrentPriceWithDecimals() view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_Ftso *FtsoCallerSession) GetCurrentPriceWithDecimals() (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	return _Ftso.Contract.GetCurrentPriceWithDecimals(&_Ftso.CallOpts)
}

// GetCurrentPriceWithDecimalsFromTrustedProviders is a free data retrieval call binding the contract method 0x3cacb3ae.
//
// Solidity: function getCurrentPriceWithDecimalsFromTrustedProviders() view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_Ftso *FtsoCaller) GetCurrentPriceWithDecimalsFromTrustedProviders(opts *bind.CallOpts) (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getCurrentPriceWithDecimalsFromTrustedProviders")

	outstruct := new(struct {
		Price                 *big.Int
		Timestamp             *big.Int
		AssetPriceUsdDecimals *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AssetPriceUsdDecimals = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentPriceWithDecimalsFromTrustedProviders is a free data retrieval call binding the contract method 0x3cacb3ae.
//
// Solidity: function getCurrentPriceWithDecimalsFromTrustedProviders() view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_Ftso *FtsoSession) GetCurrentPriceWithDecimalsFromTrustedProviders() (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	return _Ftso.Contract.GetCurrentPriceWithDecimalsFromTrustedProviders(&_Ftso.CallOpts)
}

// GetCurrentPriceWithDecimalsFromTrustedProviders is a free data retrieval call binding the contract method 0x3cacb3ae.
//
// Solidity: function getCurrentPriceWithDecimalsFromTrustedProviders() view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_Ftso *FtsoCallerSession) GetCurrentPriceWithDecimalsFromTrustedProviders() (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	return _Ftso.Contract.GetCurrentPriceWithDecimalsFromTrustedProviders(&_Ftso.CallOpts)
}

// GetCurrentRandom is a free data retrieval call binding the contract method 0xd89601fd.
//
// Solidity: function getCurrentRandom() view returns(uint256)
func (_Ftso *FtsoCaller) GetCurrentRandom(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getCurrentRandom")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRandom is a free data retrieval call binding the contract method 0xd89601fd.
//
// Solidity: function getCurrentRandom() view returns(uint256)
func (_Ftso *FtsoSession) GetCurrentRandom() (*big.Int, error) {
	return _Ftso.Contract.GetCurrentRandom(&_Ftso.CallOpts)
}

// GetCurrentRandom is a free data retrieval call binding the contract method 0xd89601fd.
//
// Solidity: function getCurrentRandom() view returns(uint256)
func (_Ftso *FtsoCallerSession) GetCurrentRandom() (*big.Int, error) {
	return _Ftso.Contract.GetCurrentRandom(&_Ftso.CallOpts)
}

// GetEpochId is a free data retrieval call binding the contract method 0x5303548b.
//
// Solidity: function getEpochId(uint256 _timestamp) view returns(uint256)
func (_Ftso *FtsoCaller) GetEpochId(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getEpochId", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochId is a free data retrieval call binding the contract method 0x5303548b.
//
// Solidity: function getEpochId(uint256 _timestamp) view returns(uint256)
func (_Ftso *FtsoSession) GetEpochId(_timestamp *big.Int) (*big.Int, error) {
	return _Ftso.Contract.GetEpochId(&_Ftso.CallOpts, _timestamp)
}

// GetEpochId is a free data retrieval call binding the contract method 0x5303548b.
//
// Solidity: function getEpochId(uint256 _timestamp) view returns(uint256)
func (_Ftso *FtsoCallerSession) GetEpochId(_timestamp *big.Int) (*big.Int, error) {
	return _Ftso.Contract.GetEpochId(&_Ftso.CallOpts, _timestamp)
}

// GetEpochPrice is a free data retrieval call binding the contract method 0x7d1d6f12.
//
// Solidity: function getEpochPrice(uint256 _epochId) view returns(uint256)
func (_Ftso *FtsoCaller) GetEpochPrice(opts *bind.CallOpts, _epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getEpochPrice", _epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochPrice is a free data retrieval call binding the contract method 0x7d1d6f12.
//
// Solidity: function getEpochPrice(uint256 _epochId) view returns(uint256)
func (_Ftso *FtsoSession) GetEpochPrice(_epochId *big.Int) (*big.Int, error) {
	return _Ftso.Contract.GetEpochPrice(&_Ftso.CallOpts, _epochId)
}

// GetEpochPrice is a free data retrieval call binding the contract method 0x7d1d6f12.
//
// Solidity: function getEpochPrice(uint256 _epochId) view returns(uint256)
func (_Ftso *FtsoCallerSession) GetEpochPrice(_epochId *big.Int) (*big.Int, error) {
	return _Ftso.Contract.GetEpochPrice(&_Ftso.CallOpts, _epochId)
}

// GetEpochPriceForVoter is a free data retrieval call binding the contract method 0xc5d8b9e7.
//
// Solidity: function getEpochPriceForVoter(uint256 _epochId, address _voter) view returns(uint256)
func (_Ftso *FtsoCaller) GetEpochPriceForVoter(opts *bind.CallOpts, _epochId *big.Int, _voter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getEpochPriceForVoter", _epochId, _voter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochPriceForVoter is a free data retrieval call binding the contract method 0xc5d8b9e7.
//
// Solidity: function getEpochPriceForVoter(uint256 _epochId, address _voter) view returns(uint256)
func (_Ftso *FtsoSession) GetEpochPriceForVoter(_epochId *big.Int, _voter common.Address) (*big.Int, error) {
	return _Ftso.Contract.GetEpochPriceForVoter(&_Ftso.CallOpts, _epochId, _voter)
}

// GetEpochPriceForVoter is a free data retrieval call binding the contract method 0xc5d8b9e7.
//
// Solidity: function getEpochPriceForVoter(uint256 _epochId, address _voter) view returns(uint256)
func (_Ftso *FtsoCallerSession) GetEpochPriceForVoter(_epochId *big.Int, _voter common.Address) (*big.Int, error) {
	return _Ftso.Contract.GetEpochPriceForVoter(&_Ftso.CallOpts, _epochId, _voter)
}

// GetPriceEpochConfiguration is a free data retrieval call binding the contract method 0x144e1591.
//
// Solidity: function getPriceEpochConfiguration() view returns(uint256 _firstEpochStartTs, uint256 _submitPeriodSeconds, uint256 _revealPeriodSeconds)
func (_Ftso *FtsoCaller) GetPriceEpochConfiguration(opts *bind.CallOpts) (struct {
	FirstEpochStartTs   *big.Int
	SubmitPeriodSeconds *big.Int
	RevealPeriodSeconds *big.Int
}, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getPriceEpochConfiguration")

	outstruct := new(struct {
		FirstEpochStartTs   *big.Int
		SubmitPeriodSeconds *big.Int
		RevealPeriodSeconds *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FirstEpochStartTs = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SubmitPeriodSeconds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RevealPeriodSeconds = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPriceEpochConfiguration is a free data retrieval call binding the contract method 0x144e1591.
//
// Solidity: function getPriceEpochConfiguration() view returns(uint256 _firstEpochStartTs, uint256 _submitPeriodSeconds, uint256 _revealPeriodSeconds)
func (_Ftso *FtsoSession) GetPriceEpochConfiguration() (struct {
	FirstEpochStartTs   *big.Int
	SubmitPeriodSeconds *big.Int
	RevealPeriodSeconds *big.Int
}, error) {
	return _Ftso.Contract.GetPriceEpochConfiguration(&_Ftso.CallOpts)
}

// GetPriceEpochConfiguration is a free data retrieval call binding the contract method 0x144e1591.
//
// Solidity: function getPriceEpochConfiguration() view returns(uint256 _firstEpochStartTs, uint256 _submitPeriodSeconds, uint256 _revealPeriodSeconds)
func (_Ftso *FtsoCallerSession) GetPriceEpochConfiguration() (struct {
	FirstEpochStartTs   *big.Int
	SubmitPeriodSeconds *big.Int
	RevealPeriodSeconds *big.Int
}, error) {
	return _Ftso.Contract.GetPriceEpochConfiguration(&_Ftso.CallOpts)
}

// GetPriceEpochData is a free data retrieval call binding the contract method 0xe3b3a3b3.
//
// Solidity: function getPriceEpochData() view returns(uint256 _epochId, uint256 _epochSubmitEndTime, uint256 _epochRevealEndTime, uint256 _votePowerBlock, bool _fallbackMode)
func (_Ftso *FtsoCaller) GetPriceEpochData(opts *bind.CallOpts) (struct {
	EpochId            *big.Int
	EpochSubmitEndTime *big.Int
	EpochRevealEndTime *big.Int
	VotePowerBlock     *big.Int
	FallbackMode       bool
}, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getPriceEpochData")

	outstruct := new(struct {
		EpochId            *big.Int
		EpochSubmitEndTime *big.Int
		EpochRevealEndTime *big.Int
		VotePowerBlock     *big.Int
		FallbackMode       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EpochId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochSubmitEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochRevealEndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VotePowerBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FallbackMode = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// GetPriceEpochData is a free data retrieval call binding the contract method 0xe3b3a3b3.
//
// Solidity: function getPriceEpochData() view returns(uint256 _epochId, uint256 _epochSubmitEndTime, uint256 _epochRevealEndTime, uint256 _votePowerBlock, bool _fallbackMode)
func (_Ftso *FtsoSession) GetPriceEpochData() (struct {
	EpochId            *big.Int
	EpochSubmitEndTime *big.Int
	EpochRevealEndTime *big.Int
	VotePowerBlock     *big.Int
	FallbackMode       bool
}, error) {
	return _Ftso.Contract.GetPriceEpochData(&_Ftso.CallOpts)
}

// GetPriceEpochData is a free data retrieval call binding the contract method 0xe3b3a3b3.
//
// Solidity: function getPriceEpochData() view returns(uint256 _epochId, uint256 _epochSubmitEndTime, uint256 _epochRevealEndTime, uint256 _votePowerBlock, bool _fallbackMode)
func (_Ftso *FtsoCallerSession) GetPriceEpochData() (struct {
	EpochId            *big.Int
	EpochSubmitEndTime *big.Int
	EpochRevealEndTime *big.Int
	VotePowerBlock     *big.Int
	FallbackMode       bool
}, error) {
	return _Ftso.Contract.GetPriceEpochData(&_Ftso.CallOpts)
}

// GetRandom is a free data retrieval call binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(uint256 _epochId) view returns(uint256)
func (_Ftso *FtsoCaller) GetRandom(opts *bind.CallOpts, _epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getRandom", _epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRandom is a free data retrieval call binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(uint256 _epochId) view returns(uint256)
func (_Ftso *FtsoSession) GetRandom(_epochId *big.Int) (*big.Int, error) {
	return _Ftso.Contract.GetRandom(&_Ftso.CallOpts, _epochId)
}

// GetRandom is a free data retrieval call binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(uint256 _epochId) view returns(uint256)
func (_Ftso *FtsoCallerSession) GetRandom(_epochId *big.Int) (*big.Int, error) {
	return _Ftso.Contract.GetRandom(&_Ftso.CallOpts, _epochId)
}

// GetVoteWeightingParameters is a free data retrieval call binding the contract method 0x8357d08c.
//
// Solidity: function getVoteWeightingParameters() view returns(address[] _assets, uint256[] _assetMultipliers, uint256 _totalVotePowerNat, uint256 _totalVotePowerAsset, uint256 _assetWeightRatio, uint256 _votePowerBlock)
func (_Ftso *FtsoCaller) GetVoteWeightingParameters(opts *bind.CallOpts) (struct {
	Assets              []common.Address
	AssetMultipliers    []*big.Int
	TotalVotePowerNat   *big.Int
	TotalVotePowerAsset *big.Int
	AssetWeightRatio    *big.Int
	VotePowerBlock      *big.Int
}, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "getVoteWeightingParameters")

	outstruct := new(struct {
		Assets              []common.Address
		AssetMultipliers    []*big.Int
		TotalVotePowerNat   *big.Int
		TotalVotePowerAsset *big.Int
		AssetWeightRatio    *big.Int
		VotePowerBlock      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AssetMultipliers = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.TotalVotePowerNat = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalVotePowerAsset = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AssetWeightRatio = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.VotePowerBlock = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVoteWeightingParameters is a free data retrieval call binding the contract method 0x8357d08c.
//
// Solidity: function getVoteWeightingParameters() view returns(address[] _assets, uint256[] _assetMultipliers, uint256 _totalVotePowerNat, uint256 _totalVotePowerAsset, uint256 _assetWeightRatio, uint256 _votePowerBlock)
func (_Ftso *FtsoSession) GetVoteWeightingParameters() (struct {
	Assets              []common.Address
	AssetMultipliers    []*big.Int
	TotalVotePowerNat   *big.Int
	TotalVotePowerAsset *big.Int
	AssetWeightRatio    *big.Int
	VotePowerBlock      *big.Int
}, error) {
	return _Ftso.Contract.GetVoteWeightingParameters(&_Ftso.CallOpts)
}

// GetVoteWeightingParameters is a free data retrieval call binding the contract method 0x8357d08c.
//
// Solidity: function getVoteWeightingParameters() view returns(address[] _assets, uint256[] _assetMultipliers, uint256 _totalVotePowerNat, uint256 _totalVotePowerAsset, uint256 _assetWeightRatio, uint256 _votePowerBlock)
func (_Ftso *FtsoCallerSession) GetVoteWeightingParameters() (struct {
	Assets              []common.Address
	AssetMultipliers    []*big.Int
	TotalVotePowerNat   *big.Int
	TotalVotePowerAsset *big.Int
	AssetWeightRatio    *big.Int
	VotePowerBlock      *big.Int
}, error) {
	return _Ftso.Contract.GetVoteWeightingParameters(&_Ftso.CallOpts)
}

// PriceDeviationThresholdBIPS is a free data retrieval call binding the contract method 0xcc245ab5.
//
// Solidity: function priceDeviationThresholdBIPS() view returns(uint256)
func (_Ftso *FtsoCaller) PriceDeviationThresholdBIPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "priceDeviationThresholdBIPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceDeviationThresholdBIPS is a free data retrieval call binding the contract method 0xcc245ab5.
//
// Solidity: function priceDeviationThresholdBIPS() view returns(uint256)
func (_Ftso *FtsoSession) PriceDeviationThresholdBIPS() (*big.Int, error) {
	return _Ftso.Contract.PriceDeviationThresholdBIPS(&_Ftso.CallOpts)
}

// PriceDeviationThresholdBIPS is a free data retrieval call binding the contract method 0xcc245ab5.
//
// Solidity: function priceDeviationThresholdBIPS() view returns(uint256)
func (_Ftso *FtsoCallerSession) PriceDeviationThresholdBIPS() (*big.Int, error) {
	return _Ftso.Contract.PriceDeviationThresholdBIPS(&_Ftso.CallOpts)
}

// PriceEpochCyclicBufferSize is a free data retrieval call binding the contract method 0xf025bf66.
//
// Solidity: function priceEpochCyclicBufferSize() view returns(uint256)
func (_Ftso *FtsoCaller) PriceEpochCyclicBufferSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "priceEpochCyclicBufferSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceEpochCyclicBufferSize is a free data retrieval call binding the contract method 0xf025bf66.
//
// Solidity: function priceEpochCyclicBufferSize() view returns(uint256)
func (_Ftso *FtsoSession) PriceEpochCyclicBufferSize() (*big.Int, error) {
	return _Ftso.Contract.PriceEpochCyclicBufferSize(&_Ftso.CallOpts)
}

// PriceEpochCyclicBufferSize is a free data retrieval call binding the contract method 0xf025bf66.
//
// Solidity: function priceEpochCyclicBufferSize() view returns(uint256)
func (_Ftso *FtsoCallerSession) PriceEpochCyclicBufferSize() (*big.Int, error) {
	return _Ftso.Contract.PriceEpochCyclicBufferSize(&_Ftso.CallOpts)
}

// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
//
// Solidity: function priceSubmitter() view returns(address)
func (_Ftso *FtsoCaller) PriceSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "priceSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
//
// Solidity: function priceSubmitter() view returns(address)
func (_Ftso *FtsoSession) PriceSubmitter() (common.Address, error) {
	return _Ftso.Contract.PriceSubmitter(&_Ftso.CallOpts)
}

// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
//
// Solidity: function priceSubmitter() view returns(address)
func (_Ftso *FtsoCallerSession) PriceSubmitter() (common.Address, error) {
	return _Ftso.Contract.PriceSubmitter(&_Ftso.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ftso *FtsoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ftso *FtsoSession) Symbol() (string, error) {
	return _Ftso.Contract.Symbol(&_Ftso.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ftso *FtsoCallerSession) Symbol() (string, error) {
	return _Ftso.Contract.Symbol(&_Ftso.CallOpts)
}

// WNat is a free data retrieval call binding the contract method 0x9edbf007.
//
// Solidity: function wNat() view returns(address)
func (_Ftso *FtsoCaller) WNat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ftso.contract.Call(opts, &out, "wNat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WNat is a free data retrieval call binding the contract method 0x9edbf007.
//
// Solidity: function wNat() view returns(address)
func (_Ftso *FtsoSession) WNat() (common.Address, error) {
	return _Ftso.Contract.WNat(&_Ftso.CallOpts)
}

// WNat is a free data retrieval call binding the contract method 0x9edbf007.
//
// Solidity: function wNat() view returns(address)
func (_Ftso *FtsoCallerSession) WNat() (common.Address, error) {
	return _Ftso.Contract.WNat(&_Ftso.CallOpts)
}

// ActivateFtso is a paid mutator transaction binding the contract method 0x2f0a6f3c.
//
// Solidity: function activateFtso(uint256 _firstEpochStartTs, uint256 _submitPeriodSeconds, uint256 _revealPeriodSeconds) returns()
func (_Ftso *FtsoTransactor) ActivateFtso(opts *bind.TransactOpts, _firstEpochStartTs *big.Int, _submitPeriodSeconds *big.Int, _revealPeriodSeconds *big.Int) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "activateFtso", _firstEpochStartTs, _submitPeriodSeconds, _revealPeriodSeconds)
}

// ActivateFtso is a paid mutator transaction binding the contract method 0x2f0a6f3c.
//
// Solidity: function activateFtso(uint256 _firstEpochStartTs, uint256 _submitPeriodSeconds, uint256 _revealPeriodSeconds) returns()
func (_Ftso *FtsoSession) ActivateFtso(_firstEpochStartTs *big.Int, _submitPeriodSeconds *big.Int, _revealPeriodSeconds *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.ActivateFtso(&_Ftso.TransactOpts, _firstEpochStartTs, _submitPeriodSeconds, _revealPeriodSeconds)
}

// ActivateFtso is a paid mutator transaction binding the contract method 0x2f0a6f3c.
//
// Solidity: function activateFtso(uint256 _firstEpochStartTs, uint256 _submitPeriodSeconds, uint256 _revealPeriodSeconds) returns()
func (_Ftso *FtsoTransactorSession) ActivateFtso(_firstEpochStartTs *big.Int, _submitPeriodSeconds *big.Int, _revealPeriodSeconds *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.ActivateFtso(&_Ftso.TransactOpts, _firstEpochStartTs, _submitPeriodSeconds, _revealPeriodSeconds)
}

// ConfigureEpochs is a paid mutator transaction binding the contract method 0xf7dba1f5.
//
// Solidity: function configureEpochs(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, address[] _trustedAddresses) returns()
func (_Ftso *FtsoTransactor) ConfigureEpochs(opts *bind.TransactOpts, _maxVotePowerNatThresholdFraction *big.Int, _maxVotePowerAssetThresholdFraction *big.Int, _lowAssetUSDThreshold *big.Int, _highAssetUSDThreshold *big.Int, _highAssetTurnoutThresholdBIPS *big.Int, _lowNatTurnoutThresholdBIPS *big.Int, _trustedAddresses []common.Address) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "configureEpochs", _maxVotePowerNatThresholdFraction, _maxVotePowerAssetThresholdFraction, _lowAssetUSDThreshold, _highAssetUSDThreshold, _highAssetTurnoutThresholdBIPS, _lowNatTurnoutThresholdBIPS, _trustedAddresses)
}

// ConfigureEpochs is a paid mutator transaction binding the contract method 0xf7dba1f5.
//
// Solidity: function configureEpochs(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, address[] _trustedAddresses) returns()
func (_Ftso *FtsoSession) ConfigureEpochs(_maxVotePowerNatThresholdFraction *big.Int, _maxVotePowerAssetThresholdFraction *big.Int, _lowAssetUSDThreshold *big.Int, _highAssetUSDThreshold *big.Int, _highAssetTurnoutThresholdBIPS *big.Int, _lowNatTurnoutThresholdBIPS *big.Int, _trustedAddresses []common.Address) (*types.Transaction, error) {
	return _Ftso.Contract.ConfigureEpochs(&_Ftso.TransactOpts, _maxVotePowerNatThresholdFraction, _maxVotePowerAssetThresholdFraction, _lowAssetUSDThreshold, _highAssetUSDThreshold, _highAssetTurnoutThresholdBIPS, _lowNatTurnoutThresholdBIPS, _trustedAddresses)
}

// ConfigureEpochs is a paid mutator transaction binding the contract method 0xf7dba1f5.
//
// Solidity: function configureEpochs(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, address[] _trustedAddresses) returns()
func (_Ftso *FtsoTransactorSession) ConfigureEpochs(_maxVotePowerNatThresholdFraction *big.Int, _maxVotePowerAssetThresholdFraction *big.Int, _lowAssetUSDThreshold *big.Int, _highAssetUSDThreshold *big.Int, _highAssetTurnoutThresholdBIPS *big.Int, _lowNatTurnoutThresholdBIPS *big.Int, _trustedAddresses []common.Address) (*types.Transaction, error) {
	return _Ftso.Contract.ConfigureEpochs(&_Ftso.TransactOpts, _maxVotePowerNatThresholdFraction, _maxVotePowerAssetThresholdFraction, _lowAssetUSDThreshold, _highAssetUSDThreshold, _highAssetTurnoutThresholdBIPS, _lowNatTurnoutThresholdBIPS, _trustedAddresses)
}

// DeactivateFtso is a paid mutator transaction binding the contract method 0x555989da.
//
// Solidity: function deactivateFtso() returns()
func (_Ftso *FtsoTransactor) DeactivateFtso(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "deactivateFtso")
}

// DeactivateFtso is a paid mutator transaction binding the contract method 0x555989da.
//
// Solidity: function deactivateFtso() returns()
func (_Ftso *FtsoSession) DeactivateFtso() (*types.Transaction, error) {
	return _Ftso.Contract.DeactivateFtso(&_Ftso.TransactOpts)
}

// DeactivateFtso is a paid mutator transaction binding the contract method 0x555989da.
//
// Solidity: function deactivateFtso() returns()
func (_Ftso *FtsoTransactorSession) DeactivateFtso() (*types.Transaction, error) {
	return _Ftso.Contract.DeactivateFtso(&_Ftso.TransactOpts)
}

// FallbackFinalizePriceEpoch is a paid mutator transaction binding the contract method 0x4afd5102.
//
// Solidity: function fallbackFinalizePriceEpoch(uint256 _epochId) returns()
func (_Ftso *FtsoTransactor) FallbackFinalizePriceEpoch(opts *bind.TransactOpts, _epochId *big.Int) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "fallbackFinalizePriceEpoch", _epochId)
}

// FallbackFinalizePriceEpoch is a paid mutator transaction binding the contract method 0x4afd5102.
//
// Solidity: function fallbackFinalizePriceEpoch(uint256 _epochId) returns()
func (_Ftso *FtsoSession) FallbackFinalizePriceEpoch(_epochId *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.FallbackFinalizePriceEpoch(&_Ftso.TransactOpts, _epochId)
}

// FallbackFinalizePriceEpoch is a paid mutator transaction binding the contract method 0x4afd5102.
//
// Solidity: function fallbackFinalizePriceEpoch(uint256 _epochId) returns()
func (_Ftso *FtsoTransactorSession) FallbackFinalizePriceEpoch(_epochId *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.FallbackFinalizePriceEpoch(&_Ftso.TransactOpts, _epochId)
}

// FinalizePriceEpoch is a paid mutator transaction binding the contract method 0x40462a2d.
//
// Solidity: function finalizePriceEpoch(uint256 _epochId, bool _returnRewardData) returns(address[] _eligibleAddresses, uint256[] _natWeights, uint256 _natWeightsSum)
func (_Ftso *FtsoTransactor) FinalizePriceEpoch(opts *bind.TransactOpts, _epochId *big.Int, _returnRewardData bool) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "finalizePriceEpoch", _epochId, _returnRewardData)
}

// FinalizePriceEpoch is a paid mutator transaction binding the contract method 0x40462a2d.
//
// Solidity: function finalizePriceEpoch(uint256 _epochId, bool _returnRewardData) returns(address[] _eligibleAddresses, uint256[] _natWeights, uint256 _natWeightsSum)
func (_Ftso *FtsoSession) FinalizePriceEpoch(_epochId *big.Int, _returnRewardData bool) (*types.Transaction, error) {
	return _Ftso.Contract.FinalizePriceEpoch(&_Ftso.TransactOpts, _epochId, _returnRewardData)
}

// FinalizePriceEpoch is a paid mutator transaction binding the contract method 0x40462a2d.
//
// Solidity: function finalizePriceEpoch(uint256 _epochId, bool _returnRewardData) returns(address[] _eligibleAddresses, uint256[] _natWeights, uint256 _natWeightsSum)
func (_Ftso *FtsoTransactorSession) FinalizePriceEpoch(_epochId *big.Int, _returnRewardData bool) (*types.Transaction, error) {
	return _Ftso.Contract.FinalizePriceEpoch(&_Ftso.TransactOpts, _epochId, _returnRewardData)
}

// ForceFinalizePriceEpoch is a paid mutator transaction binding the contract method 0x974d7a6b.
//
// Solidity: function forceFinalizePriceEpoch(uint256 _epochId) returns()
func (_Ftso *FtsoTransactor) ForceFinalizePriceEpoch(opts *bind.TransactOpts, _epochId *big.Int) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "forceFinalizePriceEpoch", _epochId)
}

// ForceFinalizePriceEpoch is a paid mutator transaction binding the contract method 0x974d7a6b.
//
// Solidity: function forceFinalizePriceEpoch(uint256 _epochId) returns()
func (_Ftso *FtsoSession) ForceFinalizePriceEpoch(_epochId *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.ForceFinalizePriceEpoch(&_Ftso.TransactOpts, _epochId)
}

// ForceFinalizePriceEpoch is a paid mutator transaction binding the contract method 0x974d7a6b.
//
// Solidity: function forceFinalizePriceEpoch(uint256 _epochId) returns()
func (_Ftso *FtsoTransactorSession) ForceFinalizePriceEpoch(_epochId *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.ForceFinalizePriceEpoch(&_Ftso.TransactOpts, _epochId)
}

// InitializeCurrentEpochStateForReveal is a paid mutator transaction binding the contract method 0xf670ebe3.
//
// Solidity: function initializeCurrentEpochStateForReveal(uint256 _circulatingSupplyNat, bool _fallbackMode) returns()
func (_Ftso *FtsoTransactor) InitializeCurrentEpochStateForReveal(opts *bind.TransactOpts, _circulatingSupplyNat *big.Int, _fallbackMode bool) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "initializeCurrentEpochStateForReveal", _circulatingSupplyNat, _fallbackMode)
}

// InitializeCurrentEpochStateForReveal is a paid mutator transaction binding the contract method 0xf670ebe3.
//
// Solidity: function initializeCurrentEpochStateForReveal(uint256 _circulatingSupplyNat, bool _fallbackMode) returns()
func (_Ftso *FtsoSession) InitializeCurrentEpochStateForReveal(_circulatingSupplyNat *big.Int, _fallbackMode bool) (*types.Transaction, error) {
	return _Ftso.Contract.InitializeCurrentEpochStateForReveal(&_Ftso.TransactOpts, _circulatingSupplyNat, _fallbackMode)
}

// InitializeCurrentEpochStateForReveal is a paid mutator transaction binding the contract method 0xf670ebe3.
//
// Solidity: function initializeCurrentEpochStateForReveal(uint256 _circulatingSupplyNat, bool _fallbackMode) returns()
func (_Ftso *FtsoTransactorSession) InitializeCurrentEpochStateForReveal(_circulatingSupplyNat *big.Int, _fallbackMode bool) (*types.Transaction, error) {
	return _Ftso.Contract.InitializeCurrentEpochStateForReveal(&_Ftso.TransactOpts, _circulatingSupplyNat, _fallbackMode)
}

// RevealPriceSubmitter is a paid mutator transaction binding the contract method 0xc1f6c36e.
//
// Solidity: function revealPriceSubmitter(address _voter, uint256 _epochId, uint256 _price, uint256 _voterWNatVP) returns()
func (_Ftso *FtsoTransactor) RevealPriceSubmitter(opts *bind.TransactOpts, _voter common.Address, _epochId *big.Int, _price *big.Int, _voterWNatVP *big.Int) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "revealPriceSubmitter", _voter, _epochId, _price, _voterWNatVP)
}

// RevealPriceSubmitter is a paid mutator transaction binding the contract method 0xc1f6c36e.
//
// Solidity: function revealPriceSubmitter(address _voter, uint256 _epochId, uint256 _price, uint256 _voterWNatVP) returns()
func (_Ftso *FtsoSession) RevealPriceSubmitter(_voter common.Address, _epochId *big.Int, _price *big.Int, _voterWNatVP *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.RevealPriceSubmitter(&_Ftso.TransactOpts, _voter, _epochId, _price, _voterWNatVP)
}

// RevealPriceSubmitter is a paid mutator transaction binding the contract method 0xc1f6c36e.
//
// Solidity: function revealPriceSubmitter(address _voter, uint256 _epochId, uint256 _price, uint256 _voterWNatVP) returns()
func (_Ftso *FtsoTransactorSession) RevealPriceSubmitter(_voter common.Address, _epochId *big.Int, _price *big.Int, _voterWNatVP *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.RevealPriceSubmitter(&_Ftso.TransactOpts, _voter, _epochId, _price, _voterWNatVP)
}

// SetAsset is a paid mutator transaction binding the contract method 0xd0d552dd.
//
// Solidity: function setAsset(address _asset) returns()
func (_Ftso *FtsoTransactor) SetAsset(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "setAsset", _asset)
}

// SetAsset is a paid mutator transaction binding the contract method 0xd0d552dd.
//
// Solidity: function setAsset(address _asset) returns()
func (_Ftso *FtsoSession) SetAsset(_asset common.Address) (*types.Transaction, error) {
	return _Ftso.Contract.SetAsset(&_Ftso.TransactOpts, _asset)
}

// SetAsset is a paid mutator transaction binding the contract method 0xd0d552dd.
//
// Solidity: function setAsset(address _asset) returns()
func (_Ftso *FtsoTransactorSession) SetAsset(_asset common.Address) (*types.Transaction, error) {
	return _Ftso.Contract.SetAsset(&_Ftso.TransactOpts, _asset)
}

// SetAssetFtsos is a paid mutator transaction binding the contract method 0x131fdee2.
//
// Solidity: function setAssetFtsos(address[] _assetFtsos) returns()
func (_Ftso *FtsoTransactor) SetAssetFtsos(opts *bind.TransactOpts, _assetFtsos []common.Address) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "setAssetFtsos", _assetFtsos)
}

// SetAssetFtsos is a paid mutator transaction binding the contract method 0x131fdee2.
//
// Solidity: function setAssetFtsos(address[] _assetFtsos) returns()
func (_Ftso *FtsoSession) SetAssetFtsos(_assetFtsos []common.Address) (*types.Transaction, error) {
	return _Ftso.Contract.SetAssetFtsos(&_Ftso.TransactOpts, _assetFtsos)
}

// SetAssetFtsos is a paid mutator transaction binding the contract method 0x131fdee2.
//
// Solidity: function setAssetFtsos(address[] _assetFtsos) returns()
func (_Ftso *FtsoTransactorSession) SetAssetFtsos(_assetFtsos []common.Address) (*types.Transaction, error) {
	return _Ftso.Contract.SetAssetFtsos(&_Ftso.TransactOpts, _assetFtsos)
}

// SetVotePowerBlock is a paid mutator transaction binding the contract method 0xe536f396.
//
// Solidity: function setVotePowerBlock(uint256 _votePowerBlock) returns()
func (_Ftso *FtsoTransactor) SetVotePowerBlock(opts *bind.TransactOpts, _votePowerBlock *big.Int) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "setVotePowerBlock", _votePowerBlock)
}

// SetVotePowerBlock is a paid mutator transaction binding the contract method 0xe536f396.
//
// Solidity: function setVotePowerBlock(uint256 _votePowerBlock) returns()
func (_Ftso *FtsoSession) SetVotePowerBlock(_votePowerBlock *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.SetVotePowerBlock(&_Ftso.TransactOpts, _votePowerBlock)
}

// SetVotePowerBlock is a paid mutator transaction binding the contract method 0xe536f396.
//
// Solidity: function setVotePowerBlock(uint256 _votePowerBlock) returns()
func (_Ftso *FtsoTransactorSession) SetVotePowerBlock(_votePowerBlock *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.SetVotePowerBlock(&_Ftso.TransactOpts, _votePowerBlock)
}

// UpdateInitialPrice is a paid mutator transaction binding the contract method 0x306ba253.
//
// Solidity: function updateInitialPrice(uint256 _initialPriceUSD, uint256 _initialPriceTimestamp) returns()
func (_Ftso *FtsoTransactor) UpdateInitialPrice(opts *bind.TransactOpts, _initialPriceUSD *big.Int, _initialPriceTimestamp *big.Int) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "updateInitialPrice", _initialPriceUSD, _initialPriceTimestamp)
}

// UpdateInitialPrice is a paid mutator transaction binding the contract method 0x306ba253.
//
// Solidity: function updateInitialPrice(uint256 _initialPriceUSD, uint256 _initialPriceTimestamp) returns()
func (_Ftso *FtsoSession) UpdateInitialPrice(_initialPriceUSD *big.Int, _initialPriceTimestamp *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.UpdateInitialPrice(&_Ftso.TransactOpts, _initialPriceUSD, _initialPriceTimestamp)
}

// UpdateInitialPrice is a paid mutator transaction binding the contract method 0x306ba253.
//
// Solidity: function updateInitialPrice(uint256 _initialPriceUSD, uint256 _initialPriceTimestamp) returns()
func (_Ftso *FtsoTransactorSession) UpdateInitialPrice(_initialPriceUSD *big.Int, _initialPriceTimestamp *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.UpdateInitialPrice(&_Ftso.TransactOpts, _initialPriceUSD, _initialPriceTimestamp)
}

// WNatVotePowerCached is a paid mutator transaction binding the contract method 0xf72cab28.
//
// Solidity: function wNatVotePowerCached(address _owner, uint256 _epochId) returns(uint256)
func (_Ftso *FtsoTransactor) WNatVotePowerCached(opts *bind.TransactOpts, _owner common.Address, _epochId *big.Int) (*types.Transaction, error) {
	return _Ftso.contract.Transact(opts, "wNatVotePowerCached", _owner, _epochId)
}

// WNatVotePowerCached is a paid mutator transaction binding the contract method 0xf72cab28.
//
// Solidity: function wNatVotePowerCached(address _owner, uint256 _epochId) returns(uint256)
func (_Ftso *FtsoSession) WNatVotePowerCached(_owner common.Address, _epochId *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.WNatVotePowerCached(&_Ftso.TransactOpts, _owner, _epochId)
}

// WNatVotePowerCached is a paid mutator transaction binding the contract method 0xf72cab28.
//
// Solidity: function wNatVotePowerCached(address _owner, uint256 _epochId) returns(uint256)
func (_Ftso *FtsoTransactorSession) WNatVotePowerCached(_owner common.Address, _epochId *big.Int) (*types.Transaction, error) {
	return _Ftso.Contract.WNatVotePowerCached(&_Ftso.TransactOpts, _owner, _epochId)
}

// FtsoLowTurnoutIterator is returned from FilterLowTurnout and is used to iterate over the raw logs and unpacked data for LowTurnout events raised by the Ftso contract.
type FtsoLowTurnoutIterator struct {
	Event *FtsoLowTurnout // Event containing the contract specifics and raw log

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
func (it *FtsoLowTurnoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoLowTurnout)
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
		it.Event = new(FtsoLowTurnout)
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
func (it *FtsoLowTurnoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoLowTurnoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoLowTurnout represents a LowTurnout event raised by the Ftso contract.
type FtsoLowTurnout struct {
	EpochId                    *big.Int
	NatTurnout                 *big.Int
	LowNatTurnoutThresholdBIPS *big.Int
	Timestamp                  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLowTurnout is a free log retrieval operation binding the contract event 0xbece8aa526cdc5e528cdaa56c1d03edc19da51e41845aa146f64a7071f74c65a.
//
// Solidity: event LowTurnout(uint256 indexed epochId, uint256 natTurnout, uint256 lowNatTurnoutThresholdBIPS, uint256 timestamp)
func (_Ftso *FtsoFilterer) FilterLowTurnout(opts *bind.FilterOpts, epochId []*big.Int) (*FtsoLowTurnoutIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Ftso.contract.FilterLogs(opts, "LowTurnout", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &FtsoLowTurnoutIterator{contract: _Ftso.contract, event: "LowTurnout", logs: logs, sub: sub}, nil
}

// WatchLowTurnout is a free log subscription operation binding the contract event 0xbece8aa526cdc5e528cdaa56c1d03edc19da51e41845aa146f64a7071f74c65a.
//
// Solidity: event LowTurnout(uint256 indexed epochId, uint256 natTurnout, uint256 lowNatTurnoutThresholdBIPS, uint256 timestamp)
func (_Ftso *FtsoFilterer) WatchLowTurnout(opts *bind.WatchOpts, sink chan<- *FtsoLowTurnout, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Ftso.contract.WatchLogs(opts, "LowTurnout", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoLowTurnout)
				if err := _Ftso.contract.UnpackLog(event, "LowTurnout", log); err != nil {
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

// ParseLowTurnout is a log parse operation binding the contract event 0xbece8aa526cdc5e528cdaa56c1d03edc19da51e41845aa146f64a7071f74c65a.
//
// Solidity: event LowTurnout(uint256 indexed epochId, uint256 natTurnout, uint256 lowNatTurnoutThresholdBIPS, uint256 timestamp)
func (_Ftso *FtsoFilterer) ParseLowTurnout(log types.Log) (*FtsoLowTurnout, error) {
	event := new(FtsoLowTurnout)
	if err := _Ftso.contract.UnpackLog(event, "LowTurnout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoPriceEpochInitializedOnFtsoIterator is returned from FilterPriceEpochInitializedOnFtso and is used to iterate over the raw logs and unpacked data for PriceEpochInitializedOnFtso events raised by the Ftso contract.
type FtsoPriceEpochInitializedOnFtsoIterator struct {
	Event *FtsoPriceEpochInitializedOnFtso // Event containing the contract specifics and raw log

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
func (it *FtsoPriceEpochInitializedOnFtsoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoPriceEpochInitializedOnFtso)
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
		it.Event = new(FtsoPriceEpochInitializedOnFtso)
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
func (it *FtsoPriceEpochInitializedOnFtsoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoPriceEpochInitializedOnFtsoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoPriceEpochInitializedOnFtso represents a PriceEpochInitializedOnFtso event raised by the Ftso contract.
type FtsoPriceEpochInitializedOnFtso struct {
	EpochId   *big.Int
	EndTime   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceEpochInitializedOnFtso is a free log retrieval operation binding the contract event 0xc0eaa359541c7c642d9947c9496507c134f3e4f8e1fd433313eb27dc48cb1fb7.
//
// Solidity: event PriceEpochInitializedOnFtso(uint256 indexed epochId, uint256 endTime, uint256 timestamp)
func (_Ftso *FtsoFilterer) FilterPriceEpochInitializedOnFtso(opts *bind.FilterOpts, epochId []*big.Int) (*FtsoPriceEpochInitializedOnFtsoIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Ftso.contract.FilterLogs(opts, "PriceEpochInitializedOnFtso", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &FtsoPriceEpochInitializedOnFtsoIterator{contract: _Ftso.contract, event: "PriceEpochInitializedOnFtso", logs: logs, sub: sub}, nil
}

// WatchPriceEpochInitializedOnFtso is a free log subscription operation binding the contract event 0xc0eaa359541c7c642d9947c9496507c134f3e4f8e1fd433313eb27dc48cb1fb7.
//
// Solidity: event PriceEpochInitializedOnFtso(uint256 indexed epochId, uint256 endTime, uint256 timestamp)
func (_Ftso *FtsoFilterer) WatchPriceEpochInitializedOnFtso(opts *bind.WatchOpts, sink chan<- *FtsoPriceEpochInitializedOnFtso, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Ftso.contract.WatchLogs(opts, "PriceEpochInitializedOnFtso", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoPriceEpochInitializedOnFtso)
				if err := _Ftso.contract.UnpackLog(event, "PriceEpochInitializedOnFtso", log); err != nil {
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

// ParsePriceEpochInitializedOnFtso is a log parse operation binding the contract event 0xc0eaa359541c7c642d9947c9496507c134f3e4f8e1fd433313eb27dc48cb1fb7.
//
// Solidity: event PriceEpochInitializedOnFtso(uint256 indexed epochId, uint256 endTime, uint256 timestamp)
func (_Ftso *FtsoFilterer) ParsePriceEpochInitializedOnFtso(log types.Log) (*FtsoPriceEpochInitializedOnFtso, error) {
	event := new(FtsoPriceEpochInitializedOnFtso)
	if err := _Ftso.contract.UnpackLog(event, "PriceEpochInitializedOnFtso", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoPriceFinalizedIterator is returned from FilterPriceFinalized and is used to iterate over the raw logs and unpacked data for PriceFinalized events raised by the Ftso contract.
type FtsoPriceFinalizedIterator struct {
	Event *FtsoPriceFinalized // Event containing the contract specifics and raw log

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
func (it *FtsoPriceFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoPriceFinalized)
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
		it.Event = new(FtsoPriceFinalized)
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
func (it *FtsoPriceFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoPriceFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoPriceFinalized represents a PriceFinalized event raised by the Ftso contract.
type FtsoPriceFinalized struct {
	EpochId          *big.Int
	Price            *big.Int
	RewardedFtso     bool
	LowRewardPrice   *big.Int
	HighRewardPrice  *big.Int
	FinalizationType uint8
	Timestamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPriceFinalized is a free log retrieval operation binding the contract event 0xfe8865c1fe85bbf124b9e0f16cccfeeb6f330454fd79475a31261c8fa250bc30.
//
// Solidity: event PriceFinalized(uint256 indexed epochId, uint256 price, bool rewardedFtso, uint256 lowRewardPrice, uint256 highRewardPrice, uint8 finalizationType, uint256 timestamp)
func (_Ftso *FtsoFilterer) FilterPriceFinalized(opts *bind.FilterOpts, epochId []*big.Int) (*FtsoPriceFinalizedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Ftso.contract.FilterLogs(opts, "PriceFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &FtsoPriceFinalizedIterator{contract: _Ftso.contract, event: "PriceFinalized", logs: logs, sub: sub}, nil
}

// WatchPriceFinalized is a free log subscription operation binding the contract event 0xfe8865c1fe85bbf124b9e0f16cccfeeb6f330454fd79475a31261c8fa250bc30.
//
// Solidity: event PriceFinalized(uint256 indexed epochId, uint256 price, bool rewardedFtso, uint256 lowRewardPrice, uint256 highRewardPrice, uint8 finalizationType, uint256 timestamp)
func (_Ftso *FtsoFilterer) WatchPriceFinalized(opts *bind.WatchOpts, sink chan<- *FtsoPriceFinalized, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Ftso.contract.WatchLogs(opts, "PriceFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoPriceFinalized)
				if err := _Ftso.contract.UnpackLog(event, "PriceFinalized", log); err != nil {
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

// ParsePriceFinalized is a log parse operation binding the contract event 0xfe8865c1fe85bbf124b9e0f16cccfeeb6f330454fd79475a31261c8fa250bc30.
//
// Solidity: event PriceFinalized(uint256 indexed epochId, uint256 price, bool rewardedFtso, uint256 lowRewardPrice, uint256 highRewardPrice, uint8 finalizationType, uint256 timestamp)
func (_Ftso *FtsoFilterer) ParsePriceFinalized(log types.Log) (*FtsoPriceFinalized, error) {
	event := new(FtsoPriceFinalized)
	if err := _Ftso.contract.UnpackLog(event, "PriceFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoPriceRevealedIterator is returned from FilterPriceRevealed and is used to iterate over the raw logs and unpacked data for PriceRevealed events raised by the Ftso contract.
type FtsoPriceRevealedIterator struct {
	Event *FtsoPriceRevealed // Event containing the contract specifics and raw log

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
func (it *FtsoPriceRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoPriceRevealed)
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
		it.Event = new(FtsoPriceRevealed)
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
func (it *FtsoPriceRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoPriceRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoPriceRevealed represents a PriceRevealed event raised by the Ftso contract.
type FtsoPriceRevealed struct {
	Voter          common.Address
	EpochId        *big.Int
	Price          *big.Int
	Timestamp      *big.Int
	VotePowerNat   *big.Int
	VotePowerAsset *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPriceRevealed is a free log retrieval operation binding the contract event 0xc1b1d37612887c207efe8cb44f4069b7f10c45edaf6e8405648a94e02b6e9ec7.
//
// Solidity: event PriceRevealed(address indexed voter, uint256 indexed epochId, uint256 price, uint256 timestamp, uint256 votePowerNat, uint256 votePowerAsset)
func (_Ftso *FtsoFilterer) FilterPriceRevealed(opts *bind.FilterOpts, voter []common.Address, epochId []*big.Int) (*FtsoPriceRevealedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Ftso.contract.FilterLogs(opts, "PriceRevealed", voterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &FtsoPriceRevealedIterator{contract: _Ftso.contract, event: "PriceRevealed", logs: logs, sub: sub}, nil
}

// WatchPriceRevealed is a free log subscription operation binding the contract event 0xc1b1d37612887c207efe8cb44f4069b7f10c45edaf6e8405648a94e02b6e9ec7.
//
// Solidity: event PriceRevealed(address indexed voter, uint256 indexed epochId, uint256 price, uint256 timestamp, uint256 votePowerNat, uint256 votePowerAsset)
func (_Ftso *FtsoFilterer) WatchPriceRevealed(opts *bind.WatchOpts, sink chan<- *FtsoPriceRevealed, voter []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Ftso.contract.WatchLogs(opts, "PriceRevealed", voterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoPriceRevealed)
				if err := _Ftso.contract.UnpackLog(event, "PriceRevealed", log); err != nil {
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

// ParsePriceRevealed is a log parse operation binding the contract event 0xc1b1d37612887c207efe8cb44f4069b7f10c45edaf6e8405648a94e02b6e9ec7.
//
// Solidity: event PriceRevealed(address indexed voter, uint256 indexed epochId, uint256 price, uint256 timestamp, uint256 votePowerNat, uint256 votePowerAsset)
func (_Ftso *FtsoFilterer) ParsePriceRevealed(log types.Log) (*FtsoPriceRevealed, error) {
	event := new(FtsoPriceRevealed)
	if err := _Ftso.contract.UnpackLog(event, "PriceRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
