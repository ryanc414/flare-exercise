// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ftso_manager

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

// IIFtsoManagerRewardEpochData is an auto generated low-level Go binding around an user-defined struct.
type IIFtsoManagerRewardEpochData struct {
	VotepowerBlock *big.Int
	StartBlock     *big.Int
	StartTimestamp *big.Int
}

// FtsoManagerMetaData contains all meta data concerning the FtsoManager contract.
var FtsoManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"_governance\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_flareDaemon\",\"internalType\":\"contractFlareDaemon\"},{\"type\":\"address\",\"name\":\"_addressUpdater\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_priceSubmitter\",\"internalType\":\"contractIIPriceSubmitter\"},{\"type\":\"address\",\"name\":\"_oldFtsoManager\",\"internalType\":\"contractIIFtsoManagerV1\"},{\"type\":\"uint256\",\"name\":\"_firstPriceEpochStartTs\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_priceEpochDurationSeconds\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_revealEpochDurationSeconds\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_firstRewardEpochStartTs\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_rewardEpochDurationSeconds\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_votePowerIntervalFraction\",\"internalType\":\"uint256\"}]},{\"type\":\"event\",\"name\":\"AccruingUnearnedRewardsFailed\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epochId\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CleanupBlockNumberManagerFailedForBlock\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"blockNumber\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClosingExpiredRewardEpochFailed\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"rewardEpoch\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractRevertError\",\"inputs\":[{\"type\":\"address\",\"name\":\"theContract\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"atBlock\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"string\",\"name\":\"theMessage\",\"internalType\":\"string\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributingRewardsFailed\",\"inputs\":[{\"type\":\"address\",\"name\":\"ftso\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"epochId\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FallbackMode\",\"inputs\":[{\"type\":\"bool\",\"name\":\"fallbackMode\",\"internalType\":\"bool\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalizingPriceEpochFailed\",\"inputs\":[{\"type\":\"address\",\"name\":\"ftso\",\"internalType\":\"contractIIFtso\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"epochId\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint8\",\"name\":\"failingType\",\"internalType\":\"enumIFtso.PriceFinalizationType\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FtsoAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"ftso\",\"internalType\":\"contractIIFtso\",\"indexed\":false},{\"type\":\"bool\",\"name\":\"add\",\"internalType\":\"bool\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FtsoDeactivationFailed\",\"inputs\":[{\"type\":\"address\",\"name\":\"ftso\",\"internalType\":\"contractIIFtso\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FtsoFallbackMode\",\"inputs\":[{\"type\":\"address\",\"name\":\"ftso\",\"internalType\":\"contractIIFtso\",\"indexed\":false},{\"type\":\"bool\",\"name\":\"fallbackMode\",\"internalType\":\"bool\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernanceCallTimelocked\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"selector\",\"internalType\":\"bytes4\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"encodedCall\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernanceInitialised\",\"inputs\":[{\"type\":\"address\",\"name\":\"initialGovernance\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernedProductionModeEntered\",\"inputs\":[{\"type\":\"address\",\"name\":\"governanceSettings\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitializingCurrentEpochStateForRevealFailed\",\"inputs\":[{\"type\":\"address\",\"name\":\"ftso\",\"internalType\":\"contractIIFtso\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"epochId\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceEpochFinalized\",\"inputs\":[{\"type\":\"address\",\"name\":\"chosenFtso\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"rewardEpochId\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardEpochFinalized\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"votepowerBlock\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"startBlock\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimelockedGovernanceCallCanceled\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"selector\",\"internalType\":\"bytes4\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimelockedGovernanceCallExecuted\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"selector\",\"internalType\":\"bytes4\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatingActiveValidatorsTriggerFailed\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"rewardEpoch\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"MAX_TRUSTED_ADDRESSES_LENGTH\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"activate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"active\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addFtso\",\"inputs\":[{\"type\":\"address\",\"name\":\"_ftso\",\"internalType\":\"contractIIFtso\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addFtsosBulk\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"_ftsos\",\"internalType\":\"contractIIFtso[]\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"cancelGovernanceCall\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"_selector\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractCleanupBlockNumberManager\"}],\"name\":\"cleanupBlockNumberManager\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"currentRewardEpochEnds\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"daemonize\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"deactivateFtsos\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"_ftsos\",\"internalType\":\"contractIIFtso[]\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint192\",\"name\":\"totalRevertedErrors\",\"internalType\":\"uint192\"},{\"type\":\"uint64\",\"name\":\"lastErrorTypeIndex\",\"internalType\":\"uint64\"}],\"name\":\"errorData\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"executeGovernanceCall\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"_selector\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractFlareDaemon\"}],\"name\":\"flareDaemon\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"ftsoRegistry\",\"internalType\":\"contractIIFtsoRegistry\"},{\"type\":\"address\",\"name\":\"voterWhitelister\",\"internalType\":\"contractIIVoterWhitelister\"}],\"name\":\"ftsoManagement\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIIFtsoRegistry\"}],\"name\":\"ftsoRegistry\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"_addressUpdater\",\"internalType\":\"address\"}],\"name\":\"getAddressUpdater\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"getContractName\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_priceEpochId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_priceEpochStartTimestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_priceEpochEndTimestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_priceEpochRevealEndTimestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_currentTimestamp\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentPriceEpochData\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_priceEpochId\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentPriceEpochId\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentRewardEpoch\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"_fallbackMode\",\"internalType\":\"bool\"},{\"type\":\"address[]\",\"name\":\"_ftsos\",\"internalType\":\"contractIIFtso[]\"},{\"type\":\"bool[]\",\"name\":\"_ftsoInFallbackMode\",\"internalType\":\"bool[]\"}],\"name\":\"getFallbackMode\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address[]\",\"name\":\"_ftsos\",\"internalType\":\"contractIIFtso[]\"}],\"name\":\"getFtsos\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_maxVotePowerNatThresholdFraction\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_maxVotePowerAssetThresholdFraction\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_lowAssetUSDThreshold\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_highAssetUSDThreshold\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_highAssetTurnoutThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_lowNatTurnoutThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_rewardExpiryOffsetSeconds\",\"internalType\":\"uint256\"},{\"type\":\"address[]\",\"name\":\"_trustedAddresses\",\"internalType\":\"address[]\"},{\"type\":\"bool\",\"name\":\"_initialized\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"_changed\",\"internalType\":\"bool\"}],\"name\":\"getGovernanceParameters\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_lastUnprocessedPriceEpoch\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_lastUnprocessedPriceEpochRevealEnds\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"_lastUnprocessedPriceEpochInitialized\",\"internalType\":\"bool\"}],\"name\":\"getLastUnprocessedPriceEpochData\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_firstPriceEpochStartTs\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_priceEpochDurationSeconds\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_revealEpochDurationSeconds\",\"internalType\":\"uint256\"}],\"name\":\"getPriceEpochConfiguration\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIIPriceSubmitter\"}],\"name\":\"getPriceSubmitter\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_firstRewardEpochStartTs\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_rewardEpochDurationSeconds\",\"internalType\":\"uint256\"}],\"name\":\"getRewardEpochConfiguration\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structIIFtsoManager.RewardEpochData\",\"components\":[{\"type\":\"uint256\",\"name\":\"votepowerBlock\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"startBlock\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"startTimestamp\",\"internalType\":\"uint256\"}]}],\"name\":\"getRewardEpochData\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_rewardEpochId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getRewardEpochToExpireNext\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_votepowerBlock\",\"internalType\":\"uint256\"}],\"name\":\"getRewardEpochVotePowerBlock\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_rewardEpoch\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getRewardExpiryOffsetSeconds\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getVotePowerIntervalFraction\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"governance\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIGovernanceSettings\"}],\"name\":\"governanceSettings\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialise\",\"inputs\":[{\"type\":\"address\",\"name\":\"_initialGovernance\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"lastRewardedFtsoAddress\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"notInitializedFtsos\",\"inputs\":[{\"type\":\"address\",\"name\":\"_ftso\",\"internalType\":\"contractIIFtso\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIIFtsoManagerV1\"}],\"name\":\"oldFtsoManager\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIIPriceSubmitter\"}],\"name\":\"priceSubmitter\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"productionMode\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeFtso\",\"inputs\":[{\"type\":\"address\",\"name\":\"_ftso\",\"internalType\":\"contractIIFtso\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"replaceFtso\",\"inputs\":[{\"type\":\"address\",\"name\":\"_ftsoToAdd\",\"internalType\":\"contractIIFtso\"},{\"type\":\"bool\",\"name\":\"_copyCurrentPrice\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"_copyAssetOrAssetFtsos\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"replaceFtsosBulk\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"_ftsosToAdd\",\"internalType\":\"contractIIFtso[]\"},{\"type\":\"bool\",\"name\":\"_copyCurrentPrice\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"_copyAssetOrAssetFtsos\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"rewardEpochDurationSeconds\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"_votepowerBlock\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_startBlock\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_startTimestamp\",\"internalType\":\"uint256\"}],\"name\":\"rewardEpochs\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_rewardEpochId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"rewardEpochsStartTs\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIIFtsoRewardManager\"}],\"name\":\"rewardManager\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setFallbackMode\",\"inputs\":[{\"type\":\"bool\",\"name\":\"_fallbackMode\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setFtsoAsset\",\"inputs\":[{\"type\":\"address\",\"name\":\"_ftso\",\"internalType\":\"contractIIFtso\"},{\"type\":\"address\",\"name\":\"_asset\",\"internalType\":\"contractIIVPToken\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setFtsoAssetFtsos\",\"inputs\":[{\"type\":\"address\",\"name\":\"_ftso\",\"internalType\":\"contractIIFtso\"},{\"type\":\"address[]\",\"name\":\"_assetFtsos\",\"internalType\":\"contractIIFtso[]\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setFtsoFallbackMode\",\"inputs\":[{\"type\":\"address\",\"name\":\"_ftso\",\"internalType\":\"contractIIFtso\"},{\"type\":\"bool\",\"name\":\"_fallbackMode\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setGovernanceParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_maxVotePowerNatThresholdFraction\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_maxVotePowerAssetThresholdFraction\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_lowAssetUSDThreshold\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_highAssetUSDThreshold\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_highAssetTurnoutThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_lowNatTurnoutThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_rewardExpiryOffsetSeconds\",\"internalType\":\"uint256\"},{\"type\":\"address[]\",\"name\":\"_trustedAddresses\",\"internalType\":\"address[]\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setInitialRewardData\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_nextRewardEpochToExpire\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_rewardEpochsLength\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_currentRewardEpochEnds\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setRewardEpochDurationSeconds\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_rewardEpochDurationSeconds\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setUpdateOnRewardEpochSwitchover\",\"inputs\":[{\"type\":\"address\",\"name\":\"_updateValidators\",\"internalType\":\"contractIUpdateValidators\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setVotePowerIntervalFraction\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_votePowerIntervalFraction\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"maxVotePowerNatThresholdFraction\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"maxVotePowerAssetThresholdFraction\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"lowAssetUSDThreshold\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"highAssetUSDThreshold\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"highAssetTurnoutThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"lowNatTurnoutThresholdBIPS\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"rewardExpiryOffsetSeconds\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"changed\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"initialized\",\"internalType\":\"bool\"}],\"name\":\"settings\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256[]\",\"name\":\"_lastErrorBlock\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"_numErrors\",\"internalType\":\"uint256[]\"},{\"type\":\"string[]\",\"name\":\"_errorString\",\"internalType\":\"string[]\"},{\"type\":\"address[]\",\"name\":\"_erroringContract\",\"internalType\":\"address[]\"},{\"type\":\"uint256\",\"name\":\"_totalRevertedErrors\",\"internalType\":\"uint256\"}],\"name\":\"showLastRevertedError\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256[]\",\"name\":\"_lastErrorBlock\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"_numErrors\",\"internalType\":\"uint256[]\"},{\"type\":\"string[]\",\"name\":\"_errorString\",\"internalType\":\"string[]\"},{\"type\":\"address[]\",\"name\":\"_erroringContract\",\"internalType\":\"address[]\"},{\"type\":\"uint256\",\"name\":\"_totalRevertedErrors\",\"internalType\":\"uint256\"}],\"name\":\"showRevertedErrors\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"startIndex\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"numErrorTypesToShow\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIISupply\"}],\"name\":\"supply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"switchToFallbackMode\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"switchToProductionMode\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"encodedCall\",\"internalType\":\"bytes\"}],\"name\":\"timelockedCalls\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateContractAddresses\",\"inputs\":[{\"type\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"internalType\":\"bytes32[]\"},{\"type\":\"address[]\",\"name\":\"_contractAddresses\",\"internalType\":\"address[]\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIUpdateValidators\"}],\"name\":\"updateOnRewardEpochSwitchover\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIIVoterWhitelister\"}],\"name\":\"voterWhitelister\",\"inputs\":[]}]",
}

// FtsoManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FtsoManagerMetaData.ABI instead.
var FtsoManagerABI = FtsoManagerMetaData.ABI

// FtsoManager is an auto generated Go binding around an Ethereum contract.
type FtsoManager struct {
	FtsoManagerCaller     // Read-only binding to the contract
	FtsoManagerTransactor // Write-only binding to the contract
	FtsoManagerFilterer   // Log filterer for contract events
}

// FtsoManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FtsoManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtsoManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FtsoManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtsoManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FtsoManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtsoManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FtsoManagerSession struct {
	Contract     *FtsoManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FtsoManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FtsoManagerCallerSession struct {
	Contract *FtsoManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FtsoManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FtsoManagerTransactorSession struct {
	Contract     *FtsoManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FtsoManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FtsoManagerRaw struct {
	Contract *FtsoManager // Generic contract binding to access the raw methods on
}

// FtsoManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FtsoManagerCallerRaw struct {
	Contract *FtsoManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FtsoManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FtsoManagerTransactorRaw struct {
	Contract *FtsoManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFtsoManager creates a new instance of FtsoManager, bound to a specific deployed contract.
func NewFtsoManager(address common.Address, backend bind.ContractBackend) (*FtsoManager, error) {
	contract, err := bindFtsoManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FtsoManager{FtsoManagerCaller: FtsoManagerCaller{contract: contract}, FtsoManagerTransactor: FtsoManagerTransactor{contract: contract}, FtsoManagerFilterer: FtsoManagerFilterer{contract: contract}}, nil
}

// NewFtsoManagerCaller creates a new read-only instance of FtsoManager, bound to a specific deployed contract.
func NewFtsoManagerCaller(address common.Address, caller bind.ContractCaller) (*FtsoManagerCaller, error) {
	contract, err := bindFtsoManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FtsoManagerCaller{contract: contract}, nil
}

// NewFtsoManagerTransactor creates a new write-only instance of FtsoManager, bound to a specific deployed contract.
func NewFtsoManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FtsoManagerTransactor, error) {
	contract, err := bindFtsoManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FtsoManagerTransactor{contract: contract}, nil
}

// NewFtsoManagerFilterer creates a new log filterer instance of FtsoManager, bound to a specific deployed contract.
func NewFtsoManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FtsoManagerFilterer, error) {
	contract, err := bindFtsoManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FtsoManagerFilterer{contract: contract}, nil
}

// bindFtsoManager binds a generic wrapper to an already deployed contract.
func bindFtsoManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FtsoManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FtsoManager *FtsoManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FtsoManager.Contract.FtsoManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FtsoManager *FtsoManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtsoManager.Contract.FtsoManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FtsoManager *FtsoManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FtsoManager.Contract.FtsoManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FtsoManager *FtsoManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FtsoManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FtsoManager *FtsoManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtsoManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FtsoManager *FtsoManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FtsoManager.Contract.contract.Transact(opts, method, params...)
}

// MAXTRUSTEDADDRESSESLENGTH is a free data retrieval call binding the contract method 0x69b11ac6.
//
// Solidity: function MAX_TRUSTED_ADDRESSES_LENGTH() view returns(uint256)
func (_FtsoManager *FtsoManagerCaller) MAXTRUSTEDADDRESSESLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "MAX_TRUSTED_ADDRESSES_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTRUSTEDADDRESSESLENGTH is a free data retrieval call binding the contract method 0x69b11ac6.
//
// Solidity: function MAX_TRUSTED_ADDRESSES_LENGTH() view returns(uint256)
func (_FtsoManager *FtsoManagerSession) MAXTRUSTEDADDRESSESLENGTH() (*big.Int, error) {
	return _FtsoManager.Contract.MAXTRUSTEDADDRESSESLENGTH(&_FtsoManager.CallOpts)
}

// MAXTRUSTEDADDRESSESLENGTH is a free data retrieval call binding the contract method 0x69b11ac6.
//
// Solidity: function MAX_TRUSTED_ADDRESSES_LENGTH() view returns(uint256)
func (_FtsoManager *FtsoManagerCallerSession) MAXTRUSTEDADDRESSESLENGTH() (*big.Int, error) {
	return _FtsoManager.Contract.MAXTRUSTEDADDRESSESLENGTH(&_FtsoManager.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_FtsoManager *FtsoManagerCaller) Active(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "active")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_FtsoManager *FtsoManagerSession) Active() (bool, error) {
	return _FtsoManager.Contract.Active(&_FtsoManager.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_FtsoManager *FtsoManagerCallerSession) Active() (bool, error) {
	return _FtsoManager.Contract.Active(&_FtsoManager.CallOpts)
}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_FtsoManager *FtsoManagerCaller) CleanupBlockNumberManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "cleanupBlockNumberManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_FtsoManager *FtsoManagerSession) CleanupBlockNumberManager() (common.Address, error) {
	return _FtsoManager.Contract.CleanupBlockNumberManager(&_FtsoManager.CallOpts)
}

// CleanupBlockNumberManager is a free data retrieval call binding the contract method 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) CleanupBlockNumberManager() (common.Address, error) {
	return _FtsoManager.Contract.CleanupBlockNumberManager(&_FtsoManager.CallOpts)
}

// CurrentRewardEpochEnds is a free data retrieval call binding the contract method 0xd89c39e6.
//
// Solidity: function currentRewardEpochEnds() view returns(uint256)
func (_FtsoManager *FtsoManagerCaller) CurrentRewardEpochEnds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "currentRewardEpochEnds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRewardEpochEnds is a free data retrieval call binding the contract method 0xd89c39e6.
//
// Solidity: function currentRewardEpochEnds() view returns(uint256)
func (_FtsoManager *FtsoManagerSession) CurrentRewardEpochEnds() (*big.Int, error) {
	return _FtsoManager.Contract.CurrentRewardEpochEnds(&_FtsoManager.CallOpts)
}

// CurrentRewardEpochEnds is a free data retrieval call binding the contract method 0xd89c39e6.
//
// Solidity: function currentRewardEpochEnds() view returns(uint256)
func (_FtsoManager *FtsoManagerCallerSession) CurrentRewardEpochEnds() (*big.Int, error) {
	return _FtsoManager.Contract.CurrentRewardEpochEnds(&_FtsoManager.CallOpts)
}

// ErrorData is a free data retrieval call binding the contract method 0xe371aef0.
//
// Solidity: function errorData() view returns(uint192 totalRevertedErrors, uint64 lastErrorTypeIndex)
func (_FtsoManager *FtsoManagerCaller) ErrorData(opts *bind.CallOpts) (struct {
	TotalRevertedErrors *big.Int
	LastErrorTypeIndex  uint64
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "errorData")

	outstruct := new(struct {
		TotalRevertedErrors *big.Int
		LastErrorTypeIndex  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalRevertedErrors = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastErrorTypeIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// ErrorData is a free data retrieval call binding the contract method 0xe371aef0.
//
// Solidity: function errorData() view returns(uint192 totalRevertedErrors, uint64 lastErrorTypeIndex)
func (_FtsoManager *FtsoManagerSession) ErrorData() (struct {
	TotalRevertedErrors *big.Int
	LastErrorTypeIndex  uint64
}, error) {
	return _FtsoManager.Contract.ErrorData(&_FtsoManager.CallOpts)
}

// ErrorData is a free data retrieval call binding the contract method 0xe371aef0.
//
// Solidity: function errorData() view returns(uint192 totalRevertedErrors, uint64 lastErrorTypeIndex)
func (_FtsoManager *FtsoManagerCallerSession) ErrorData() (struct {
	TotalRevertedErrors *big.Int
	LastErrorTypeIndex  uint64
}, error) {
	return _FtsoManager.Contract.ErrorData(&_FtsoManager.CallOpts)
}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FtsoManager *FtsoManagerCaller) FlareDaemon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "flareDaemon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FtsoManager *FtsoManagerSession) FlareDaemon() (common.Address, error) {
	return _FtsoManager.Contract.FlareDaemon(&_FtsoManager.CallOpts)
}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) FlareDaemon() (common.Address, error) {
	return _FtsoManager.Contract.FlareDaemon(&_FtsoManager.CallOpts)
}

// FtsoManagement is a free data retrieval call binding the contract method 0xa7d2acfa.
//
// Solidity: function ftsoManagement() view returns(address ftsoRegistry, address voterWhitelister)
func (_FtsoManager *FtsoManagerCaller) FtsoManagement(opts *bind.CallOpts) (struct {
	FtsoRegistry     common.Address
	VoterWhitelister common.Address
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "ftsoManagement")

	outstruct := new(struct {
		FtsoRegistry     common.Address
		VoterWhitelister common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FtsoRegistry = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.VoterWhitelister = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// FtsoManagement is a free data retrieval call binding the contract method 0xa7d2acfa.
//
// Solidity: function ftsoManagement() view returns(address ftsoRegistry, address voterWhitelister)
func (_FtsoManager *FtsoManagerSession) FtsoManagement() (struct {
	FtsoRegistry     common.Address
	VoterWhitelister common.Address
}, error) {
	return _FtsoManager.Contract.FtsoManagement(&_FtsoManager.CallOpts)
}

// FtsoManagement is a free data retrieval call binding the contract method 0xa7d2acfa.
//
// Solidity: function ftsoManagement() view returns(address ftsoRegistry, address voterWhitelister)
func (_FtsoManager *FtsoManagerCallerSession) FtsoManagement() (struct {
	FtsoRegistry     common.Address
	VoterWhitelister common.Address
}, error) {
	return _FtsoManager.Contract.FtsoManagement(&_FtsoManager.CallOpts)
}

// FtsoRegistry is a free data retrieval call binding the contract method 0x38b5f869.
//
// Solidity: function ftsoRegistry() view returns(address)
func (_FtsoManager *FtsoManagerCaller) FtsoRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "ftsoRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FtsoRegistry is a free data retrieval call binding the contract method 0x38b5f869.
//
// Solidity: function ftsoRegistry() view returns(address)
func (_FtsoManager *FtsoManagerSession) FtsoRegistry() (common.Address, error) {
	return _FtsoManager.Contract.FtsoRegistry(&_FtsoManager.CallOpts)
}

// FtsoRegistry is a free data retrieval call binding the contract method 0x38b5f869.
//
// Solidity: function ftsoRegistry() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) FtsoRegistry() (common.Address, error) {
	return _FtsoManager.Contract.FtsoRegistry(&_FtsoManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FtsoManager *FtsoManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FtsoManager *FtsoManagerSession) GetAddressUpdater() (common.Address, error) {
	return _FtsoManager.Contract.GetAddressUpdater(&_FtsoManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FtsoManager *FtsoManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _FtsoManager.Contract.GetAddressUpdater(&_FtsoManager.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FtsoManager *FtsoManagerCaller) GetContractName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getContractName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FtsoManager *FtsoManagerSession) GetContractName() (string, error) {
	return _FtsoManager.Contract.GetContractName(&_FtsoManager.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FtsoManager *FtsoManagerCallerSession) GetContractName() (string, error) {
	return _FtsoManager.Contract.GetContractName(&_FtsoManager.CallOpts)
}

// GetCurrentPriceEpochData is a free data retrieval call binding the contract method 0x93a79025.
//
// Solidity: function getCurrentPriceEpochData() view returns(uint256 _priceEpochId, uint256 _priceEpochStartTimestamp, uint256 _priceEpochEndTimestamp, uint256 _priceEpochRevealEndTimestamp, uint256 _currentTimestamp)
func (_FtsoManager *FtsoManagerCaller) GetCurrentPriceEpochData(opts *bind.CallOpts) (struct {
	PriceEpochId                 *big.Int
	PriceEpochStartTimestamp     *big.Int
	PriceEpochEndTimestamp       *big.Int
	PriceEpochRevealEndTimestamp *big.Int
	CurrentTimestamp             *big.Int
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getCurrentPriceEpochData")

	outstruct := new(struct {
		PriceEpochId                 *big.Int
		PriceEpochStartTimestamp     *big.Int
		PriceEpochEndTimestamp       *big.Int
		PriceEpochRevealEndTimestamp *big.Int
		CurrentTimestamp             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PriceEpochId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PriceEpochStartTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PriceEpochEndTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PriceEpochRevealEndTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CurrentTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentPriceEpochData is a free data retrieval call binding the contract method 0x93a79025.
//
// Solidity: function getCurrentPriceEpochData() view returns(uint256 _priceEpochId, uint256 _priceEpochStartTimestamp, uint256 _priceEpochEndTimestamp, uint256 _priceEpochRevealEndTimestamp, uint256 _currentTimestamp)
func (_FtsoManager *FtsoManagerSession) GetCurrentPriceEpochData() (struct {
	PriceEpochId                 *big.Int
	PriceEpochStartTimestamp     *big.Int
	PriceEpochEndTimestamp       *big.Int
	PriceEpochRevealEndTimestamp *big.Int
	CurrentTimestamp             *big.Int
}, error) {
	return _FtsoManager.Contract.GetCurrentPriceEpochData(&_FtsoManager.CallOpts)
}

// GetCurrentPriceEpochData is a free data retrieval call binding the contract method 0x93a79025.
//
// Solidity: function getCurrentPriceEpochData() view returns(uint256 _priceEpochId, uint256 _priceEpochStartTimestamp, uint256 _priceEpochEndTimestamp, uint256 _priceEpochRevealEndTimestamp, uint256 _currentTimestamp)
func (_FtsoManager *FtsoManagerCallerSession) GetCurrentPriceEpochData() (struct {
	PriceEpochId                 *big.Int
	PriceEpochStartTimestamp     *big.Int
	PriceEpochEndTimestamp       *big.Int
	PriceEpochRevealEndTimestamp *big.Int
	CurrentTimestamp             *big.Int
}, error) {
	return _FtsoManager.Contract.GetCurrentPriceEpochData(&_FtsoManager.CallOpts)
}

// GetCurrentPriceEpochId is a free data retrieval call binding the contract method 0x08a7f402.
//
// Solidity: function getCurrentPriceEpochId() view returns(uint256 _priceEpochId)
func (_FtsoManager *FtsoManagerCaller) GetCurrentPriceEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getCurrentPriceEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPriceEpochId is a free data retrieval call binding the contract method 0x08a7f402.
//
// Solidity: function getCurrentPriceEpochId() view returns(uint256 _priceEpochId)
func (_FtsoManager *FtsoManagerSession) GetCurrentPriceEpochId() (*big.Int, error) {
	return _FtsoManager.Contract.GetCurrentPriceEpochId(&_FtsoManager.CallOpts)
}

// GetCurrentPriceEpochId is a free data retrieval call binding the contract method 0x08a7f402.
//
// Solidity: function getCurrentPriceEpochId() view returns(uint256 _priceEpochId)
func (_FtsoManager *FtsoManagerCallerSession) GetCurrentPriceEpochId() (*big.Int, error) {
	return _FtsoManager.Contract.GetCurrentPriceEpochId(&_FtsoManager.CallOpts)
}

// GetCurrentRewardEpoch is a free data retrieval call binding the contract method 0xe7c830d4.
//
// Solidity: function getCurrentRewardEpoch() view returns(uint256)
func (_FtsoManager *FtsoManagerCaller) GetCurrentRewardEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getCurrentRewardEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardEpoch is a free data retrieval call binding the contract method 0xe7c830d4.
//
// Solidity: function getCurrentRewardEpoch() view returns(uint256)
func (_FtsoManager *FtsoManagerSession) GetCurrentRewardEpoch() (*big.Int, error) {
	return _FtsoManager.Contract.GetCurrentRewardEpoch(&_FtsoManager.CallOpts)
}

// GetCurrentRewardEpoch is a free data retrieval call binding the contract method 0xe7c830d4.
//
// Solidity: function getCurrentRewardEpoch() view returns(uint256)
func (_FtsoManager *FtsoManagerCallerSession) GetCurrentRewardEpoch() (*big.Int, error) {
	return _FtsoManager.Contract.GetCurrentRewardEpoch(&_FtsoManager.CallOpts)
}

// GetFallbackMode is a free data retrieval call binding the contract method 0x4b48dd5e.
//
// Solidity: function getFallbackMode() view returns(bool _fallbackMode, address[] _ftsos, bool[] _ftsoInFallbackMode)
func (_FtsoManager *FtsoManagerCaller) GetFallbackMode(opts *bind.CallOpts) (struct {
	FallbackMode       bool
	Ftsos              []common.Address
	FtsoInFallbackMode []bool
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getFallbackMode")

	outstruct := new(struct {
		FallbackMode       bool
		Ftsos              []common.Address
		FtsoInFallbackMode []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FallbackMode = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Ftsos = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.FtsoInFallbackMode = *abi.ConvertType(out[2], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetFallbackMode is a free data retrieval call binding the contract method 0x4b48dd5e.
//
// Solidity: function getFallbackMode() view returns(bool _fallbackMode, address[] _ftsos, bool[] _ftsoInFallbackMode)
func (_FtsoManager *FtsoManagerSession) GetFallbackMode() (struct {
	FallbackMode       bool
	Ftsos              []common.Address
	FtsoInFallbackMode []bool
}, error) {
	return _FtsoManager.Contract.GetFallbackMode(&_FtsoManager.CallOpts)
}

// GetFallbackMode is a free data retrieval call binding the contract method 0x4b48dd5e.
//
// Solidity: function getFallbackMode() view returns(bool _fallbackMode, address[] _ftsos, bool[] _ftsoInFallbackMode)
func (_FtsoManager *FtsoManagerCallerSession) GetFallbackMode() (struct {
	FallbackMode       bool
	Ftsos              []common.Address
	FtsoInFallbackMode []bool
}, error) {
	return _FtsoManager.Contract.GetFallbackMode(&_FtsoManager.CallOpts)
}

// GetFtsos is a free data retrieval call binding the contract method 0xce69f833.
//
// Solidity: function getFtsos() view returns(address[] _ftsos)
func (_FtsoManager *FtsoManagerCaller) GetFtsos(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getFtsos")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFtsos is a free data retrieval call binding the contract method 0xce69f833.
//
// Solidity: function getFtsos() view returns(address[] _ftsos)
func (_FtsoManager *FtsoManagerSession) GetFtsos() ([]common.Address, error) {
	return _FtsoManager.Contract.GetFtsos(&_FtsoManager.CallOpts)
}

// GetFtsos is a free data retrieval call binding the contract method 0xce69f833.
//
// Solidity: function getFtsos() view returns(address[] _ftsos)
func (_FtsoManager *FtsoManagerCallerSession) GetFtsos() ([]common.Address, error) {
	return _FtsoManager.Contract.GetFtsos(&_FtsoManager.CallOpts)
}

// GetGovernanceParameters is a free data retrieval call binding the contract method 0x5835cf30.
//
// Solidity: function getGovernanceParameters() view returns(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, uint256 _rewardExpiryOffsetSeconds, address[] _trustedAddresses, bool _initialized, bool _changed)
func (_FtsoManager *FtsoManagerCaller) GetGovernanceParameters(opts *bind.CallOpts) (struct {
	MaxVotePowerNatThresholdFraction   *big.Int
	MaxVotePowerAssetThresholdFraction *big.Int
	LowAssetUSDThreshold               *big.Int
	HighAssetUSDThreshold              *big.Int
	HighAssetTurnoutThresholdBIPS      *big.Int
	LowNatTurnoutThresholdBIPS         *big.Int
	RewardExpiryOffsetSeconds          *big.Int
	TrustedAddresses                   []common.Address
	Initialized                        bool
	Changed                            bool
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getGovernanceParameters")

	outstruct := new(struct {
		MaxVotePowerNatThresholdFraction   *big.Int
		MaxVotePowerAssetThresholdFraction *big.Int
		LowAssetUSDThreshold               *big.Int
		HighAssetUSDThreshold              *big.Int
		HighAssetTurnoutThresholdBIPS      *big.Int
		LowNatTurnoutThresholdBIPS         *big.Int
		RewardExpiryOffsetSeconds          *big.Int
		TrustedAddresses                   []common.Address
		Initialized                        bool
		Changed                            bool
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
	outstruct.RewardExpiryOffsetSeconds = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TrustedAddresses = *abi.ConvertType(out[7], new([]common.Address)).(*[]common.Address)
	outstruct.Initialized = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.Changed = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// GetGovernanceParameters is a free data retrieval call binding the contract method 0x5835cf30.
//
// Solidity: function getGovernanceParameters() view returns(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, uint256 _rewardExpiryOffsetSeconds, address[] _trustedAddresses, bool _initialized, bool _changed)
func (_FtsoManager *FtsoManagerSession) GetGovernanceParameters() (struct {
	MaxVotePowerNatThresholdFraction   *big.Int
	MaxVotePowerAssetThresholdFraction *big.Int
	LowAssetUSDThreshold               *big.Int
	HighAssetUSDThreshold              *big.Int
	HighAssetTurnoutThresholdBIPS      *big.Int
	LowNatTurnoutThresholdBIPS         *big.Int
	RewardExpiryOffsetSeconds          *big.Int
	TrustedAddresses                   []common.Address
	Initialized                        bool
	Changed                            bool
}, error) {
	return _FtsoManager.Contract.GetGovernanceParameters(&_FtsoManager.CallOpts)
}

// GetGovernanceParameters is a free data retrieval call binding the contract method 0x5835cf30.
//
// Solidity: function getGovernanceParameters() view returns(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, uint256 _rewardExpiryOffsetSeconds, address[] _trustedAddresses, bool _initialized, bool _changed)
func (_FtsoManager *FtsoManagerCallerSession) GetGovernanceParameters() (struct {
	MaxVotePowerNatThresholdFraction   *big.Int
	MaxVotePowerAssetThresholdFraction *big.Int
	LowAssetUSDThreshold               *big.Int
	HighAssetUSDThreshold              *big.Int
	HighAssetTurnoutThresholdBIPS      *big.Int
	LowNatTurnoutThresholdBIPS         *big.Int
	RewardExpiryOffsetSeconds          *big.Int
	TrustedAddresses                   []common.Address
	Initialized                        bool
	Changed                            bool
}, error) {
	return _FtsoManager.Contract.GetGovernanceParameters(&_FtsoManager.CallOpts)
}

// GetLastUnprocessedPriceEpochData is a free data retrieval call binding the contract method 0x6ca051e6.
//
// Solidity: function getLastUnprocessedPriceEpochData() view returns(uint256 _lastUnprocessedPriceEpoch, uint256 _lastUnprocessedPriceEpochRevealEnds, bool _lastUnprocessedPriceEpochInitialized)
func (_FtsoManager *FtsoManagerCaller) GetLastUnprocessedPriceEpochData(opts *bind.CallOpts) (struct {
	LastUnprocessedPriceEpoch            *big.Int
	LastUnprocessedPriceEpochRevealEnds  *big.Int
	LastUnprocessedPriceEpochInitialized bool
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getLastUnprocessedPriceEpochData")

	outstruct := new(struct {
		LastUnprocessedPriceEpoch            *big.Int
		LastUnprocessedPriceEpochRevealEnds  *big.Int
		LastUnprocessedPriceEpochInitialized bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastUnprocessedPriceEpoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUnprocessedPriceEpochRevealEnds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastUnprocessedPriceEpochInitialized = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetLastUnprocessedPriceEpochData is a free data retrieval call binding the contract method 0x6ca051e6.
//
// Solidity: function getLastUnprocessedPriceEpochData() view returns(uint256 _lastUnprocessedPriceEpoch, uint256 _lastUnprocessedPriceEpochRevealEnds, bool _lastUnprocessedPriceEpochInitialized)
func (_FtsoManager *FtsoManagerSession) GetLastUnprocessedPriceEpochData() (struct {
	LastUnprocessedPriceEpoch            *big.Int
	LastUnprocessedPriceEpochRevealEnds  *big.Int
	LastUnprocessedPriceEpochInitialized bool
}, error) {
	return _FtsoManager.Contract.GetLastUnprocessedPriceEpochData(&_FtsoManager.CallOpts)
}

// GetLastUnprocessedPriceEpochData is a free data retrieval call binding the contract method 0x6ca051e6.
//
// Solidity: function getLastUnprocessedPriceEpochData() view returns(uint256 _lastUnprocessedPriceEpoch, uint256 _lastUnprocessedPriceEpochRevealEnds, bool _lastUnprocessedPriceEpochInitialized)
func (_FtsoManager *FtsoManagerCallerSession) GetLastUnprocessedPriceEpochData() (struct {
	LastUnprocessedPriceEpoch            *big.Int
	LastUnprocessedPriceEpochRevealEnds  *big.Int
	LastUnprocessedPriceEpochInitialized bool
}, error) {
	return _FtsoManager.Contract.GetLastUnprocessedPriceEpochData(&_FtsoManager.CallOpts)
}

// GetPriceEpochConfiguration is a free data retrieval call binding the contract method 0x144e1591.
//
// Solidity: function getPriceEpochConfiguration() view returns(uint256 _firstPriceEpochStartTs, uint256 _priceEpochDurationSeconds, uint256 _revealEpochDurationSeconds)
func (_FtsoManager *FtsoManagerCaller) GetPriceEpochConfiguration(opts *bind.CallOpts) (struct {
	FirstPriceEpochStartTs     *big.Int
	PriceEpochDurationSeconds  *big.Int
	RevealEpochDurationSeconds *big.Int
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getPriceEpochConfiguration")

	outstruct := new(struct {
		FirstPriceEpochStartTs     *big.Int
		PriceEpochDurationSeconds  *big.Int
		RevealEpochDurationSeconds *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FirstPriceEpochStartTs = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PriceEpochDurationSeconds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RevealEpochDurationSeconds = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPriceEpochConfiguration is a free data retrieval call binding the contract method 0x144e1591.
//
// Solidity: function getPriceEpochConfiguration() view returns(uint256 _firstPriceEpochStartTs, uint256 _priceEpochDurationSeconds, uint256 _revealEpochDurationSeconds)
func (_FtsoManager *FtsoManagerSession) GetPriceEpochConfiguration() (struct {
	FirstPriceEpochStartTs     *big.Int
	PriceEpochDurationSeconds  *big.Int
	RevealEpochDurationSeconds *big.Int
}, error) {
	return _FtsoManager.Contract.GetPriceEpochConfiguration(&_FtsoManager.CallOpts)
}

// GetPriceEpochConfiguration is a free data retrieval call binding the contract method 0x144e1591.
//
// Solidity: function getPriceEpochConfiguration() view returns(uint256 _firstPriceEpochStartTs, uint256 _priceEpochDurationSeconds, uint256 _revealEpochDurationSeconds)
func (_FtsoManager *FtsoManagerCallerSession) GetPriceEpochConfiguration() (struct {
	FirstPriceEpochStartTs     *big.Int
	PriceEpochDurationSeconds  *big.Int
	RevealEpochDurationSeconds *big.Int
}, error) {
	return _FtsoManager.Contract.GetPriceEpochConfiguration(&_FtsoManager.CallOpts)
}

// GetPriceSubmitter is a free data retrieval call binding the contract method 0x0e063d7d.
//
// Solidity: function getPriceSubmitter() view returns(address)
func (_FtsoManager *FtsoManagerCaller) GetPriceSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getPriceSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceSubmitter is a free data retrieval call binding the contract method 0x0e063d7d.
//
// Solidity: function getPriceSubmitter() view returns(address)
func (_FtsoManager *FtsoManagerSession) GetPriceSubmitter() (common.Address, error) {
	return _FtsoManager.Contract.GetPriceSubmitter(&_FtsoManager.CallOpts)
}

// GetPriceSubmitter is a free data retrieval call binding the contract method 0x0e063d7d.
//
// Solidity: function getPriceSubmitter() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) GetPriceSubmitter() (common.Address, error) {
	return _FtsoManager.Contract.GetPriceSubmitter(&_FtsoManager.CallOpts)
}

// GetRewardEpochConfiguration is a free data retrieval call binding the contract method 0x1cb513f7.
//
// Solidity: function getRewardEpochConfiguration() view returns(uint256 _firstRewardEpochStartTs, uint256 _rewardEpochDurationSeconds)
func (_FtsoManager *FtsoManagerCaller) GetRewardEpochConfiguration(opts *bind.CallOpts) (struct {
	FirstRewardEpochStartTs    *big.Int
	RewardEpochDurationSeconds *big.Int
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getRewardEpochConfiguration")

	outstruct := new(struct {
		FirstRewardEpochStartTs    *big.Int
		RewardEpochDurationSeconds *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FirstRewardEpochStartTs = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardEpochDurationSeconds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRewardEpochConfiguration is a free data retrieval call binding the contract method 0x1cb513f7.
//
// Solidity: function getRewardEpochConfiguration() view returns(uint256 _firstRewardEpochStartTs, uint256 _rewardEpochDurationSeconds)
func (_FtsoManager *FtsoManagerSession) GetRewardEpochConfiguration() (struct {
	FirstRewardEpochStartTs    *big.Int
	RewardEpochDurationSeconds *big.Int
}, error) {
	return _FtsoManager.Contract.GetRewardEpochConfiguration(&_FtsoManager.CallOpts)
}

// GetRewardEpochConfiguration is a free data retrieval call binding the contract method 0x1cb513f7.
//
// Solidity: function getRewardEpochConfiguration() view returns(uint256 _firstRewardEpochStartTs, uint256 _rewardEpochDurationSeconds)
func (_FtsoManager *FtsoManagerCallerSession) GetRewardEpochConfiguration() (struct {
	FirstRewardEpochStartTs    *big.Int
	RewardEpochDurationSeconds *big.Int
}, error) {
	return _FtsoManager.Contract.GetRewardEpochConfiguration(&_FtsoManager.CallOpts)
}

// GetRewardEpochData is a free data retrieval call binding the contract method 0xe5399da3.
//
// Solidity: function getRewardEpochData(uint256 _rewardEpochId) view returns((uint256,uint256,uint256))
func (_FtsoManager *FtsoManagerCaller) GetRewardEpochData(opts *bind.CallOpts, _rewardEpochId *big.Int) (IIFtsoManagerRewardEpochData, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getRewardEpochData", _rewardEpochId)

	if err != nil {
		return *new(IIFtsoManagerRewardEpochData), err
	}

	out0 := *abi.ConvertType(out[0], new(IIFtsoManagerRewardEpochData)).(*IIFtsoManagerRewardEpochData)

	return out0, err

}

// GetRewardEpochData is a free data retrieval call binding the contract method 0xe5399da3.
//
// Solidity: function getRewardEpochData(uint256 _rewardEpochId) view returns((uint256,uint256,uint256))
func (_FtsoManager *FtsoManagerSession) GetRewardEpochData(_rewardEpochId *big.Int) (IIFtsoManagerRewardEpochData, error) {
	return _FtsoManager.Contract.GetRewardEpochData(&_FtsoManager.CallOpts, _rewardEpochId)
}

// GetRewardEpochData is a free data retrieval call binding the contract method 0xe5399da3.
//
// Solidity: function getRewardEpochData(uint256 _rewardEpochId) view returns((uint256,uint256,uint256))
func (_FtsoManager *FtsoManagerCallerSession) GetRewardEpochData(_rewardEpochId *big.Int) (IIFtsoManagerRewardEpochData, error) {
	return _FtsoManager.Contract.GetRewardEpochData(&_FtsoManager.CallOpts, _rewardEpochId)
}

// GetRewardEpochToExpireNext is a free data retrieval call binding the contract method 0x3e7ff857.
//
// Solidity: function getRewardEpochToExpireNext() view returns(uint256)
func (_FtsoManager *FtsoManagerCaller) GetRewardEpochToExpireNext(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getRewardEpochToExpireNext")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardEpochToExpireNext is a free data retrieval call binding the contract method 0x3e7ff857.
//
// Solidity: function getRewardEpochToExpireNext() view returns(uint256)
func (_FtsoManager *FtsoManagerSession) GetRewardEpochToExpireNext() (*big.Int, error) {
	return _FtsoManager.Contract.GetRewardEpochToExpireNext(&_FtsoManager.CallOpts)
}

// GetRewardEpochToExpireNext is a free data retrieval call binding the contract method 0x3e7ff857.
//
// Solidity: function getRewardEpochToExpireNext() view returns(uint256)
func (_FtsoManager *FtsoManagerCallerSession) GetRewardEpochToExpireNext() (*big.Int, error) {
	return _FtsoManager.Contract.GetRewardEpochToExpireNext(&_FtsoManager.CallOpts)
}

// GetRewardEpochVotePowerBlock is a free data retrieval call binding the contract method 0xf2edab5a.
//
// Solidity: function getRewardEpochVotePowerBlock(uint256 _rewardEpoch) view returns(uint256 _votepowerBlock)
func (_FtsoManager *FtsoManagerCaller) GetRewardEpochVotePowerBlock(opts *bind.CallOpts, _rewardEpoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getRewardEpochVotePowerBlock", _rewardEpoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardEpochVotePowerBlock is a free data retrieval call binding the contract method 0xf2edab5a.
//
// Solidity: function getRewardEpochVotePowerBlock(uint256 _rewardEpoch) view returns(uint256 _votepowerBlock)
func (_FtsoManager *FtsoManagerSession) GetRewardEpochVotePowerBlock(_rewardEpoch *big.Int) (*big.Int, error) {
	return _FtsoManager.Contract.GetRewardEpochVotePowerBlock(&_FtsoManager.CallOpts, _rewardEpoch)
}

// GetRewardEpochVotePowerBlock is a free data retrieval call binding the contract method 0xf2edab5a.
//
// Solidity: function getRewardEpochVotePowerBlock(uint256 _rewardEpoch) view returns(uint256 _votepowerBlock)
func (_FtsoManager *FtsoManagerCallerSession) GetRewardEpochVotePowerBlock(_rewardEpoch *big.Int) (*big.Int, error) {
	return _FtsoManager.Contract.GetRewardEpochVotePowerBlock(&_FtsoManager.CallOpts, _rewardEpoch)
}

// GetRewardExpiryOffsetSeconds is a free data retrieval call binding the contract method 0xec31db0c.
//
// Solidity: function getRewardExpiryOffsetSeconds() view returns(uint256)
func (_FtsoManager *FtsoManagerCaller) GetRewardExpiryOffsetSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getRewardExpiryOffsetSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardExpiryOffsetSeconds is a free data retrieval call binding the contract method 0xec31db0c.
//
// Solidity: function getRewardExpiryOffsetSeconds() view returns(uint256)
func (_FtsoManager *FtsoManagerSession) GetRewardExpiryOffsetSeconds() (*big.Int, error) {
	return _FtsoManager.Contract.GetRewardExpiryOffsetSeconds(&_FtsoManager.CallOpts)
}

// GetRewardExpiryOffsetSeconds is a free data retrieval call binding the contract method 0xec31db0c.
//
// Solidity: function getRewardExpiryOffsetSeconds() view returns(uint256)
func (_FtsoManager *FtsoManagerCallerSession) GetRewardExpiryOffsetSeconds() (*big.Int, error) {
	return _FtsoManager.Contract.GetRewardExpiryOffsetSeconds(&_FtsoManager.CallOpts)
}

// GetVotePowerIntervalFraction is a free data retrieval call binding the contract method 0x60f2c5b2.
//
// Solidity: function getVotePowerIntervalFraction() view returns(uint256)
func (_FtsoManager *FtsoManagerCaller) GetVotePowerIntervalFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "getVotePowerIntervalFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotePowerIntervalFraction is a free data retrieval call binding the contract method 0x60f2c5b2.
//
// Solidity: function getVotePowerIntervalFraction() view returns(uint256)
func (_FtsoManager *FtsoManagerSession) GetVotePowerIntervalFraction() (*big.Int, error) {
	return _FtsoManager.Contract.GetVotePowerIntervalFraction(&_FtsoManager.CallOpts)
}

// GetVotePowerIntervalFraction is a free data retrieval call binding the contract method 0x60f2c5b2.
//
// Solidity: function getVotePowerIntervalFraction() view returns(uint256)
func (_FtsoManager *FtsoManagerCallerSession) GetVotePowerIntervalFraction() (*big.Int, error) {
	return _FtsoManager.Contract.GetVotePowerIntervalFraction(&_FtsoManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtsoManager *FtsoManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtsoManager *FtsoManagerSession) Governance() (common.Address, error) {
	return _FtsoManager.Contract.Governance(&_FtsoManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) Governance() (common.Address, error) {
	return _FtsoManager.Contract.Governance(&_FtsoManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtsoManager *FtsoManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtsoManager *FtsoManagerSession) GovernanceSettings() (common.Address, error) {
	return _FtsoManager.Contract.GovernanceSettings(&_FtsoManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _FtsoManager.Contract.GovernanceSettings(&_FtsoManager.CallOpts)
}

// LastRewardedFtsoAddress is a free data retrieval call binding the contract method 0x2fd8eb7d.
//
// Solidity: function lastRewardedFtsoAddress() view returns(address)
func (_FtsoManager *FtsoManagerCaller) LastRewardedFtsoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "lastRewardedFtsoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastRewardedFtsoAddress is a free data retrieval call binding the contract method 0x2fd8eb7d.
//
// Solidity: function lastRewardedFtsoAddress() view returns(address)
func (_FtsoManager *FtsoManagerSession) LastRewardedFtsoAddress() (common.Address, error) {
	return _FtsoManager.Contract.LastRewardedFtsoAddress(&_FtsoManager.CallOpts)
}

// LastRewardedFtsoAddress is a free data retrieval call binding the contract method 0x2fd8eb7d.
//
// Solidity: function lastRewardedFtsoAddress() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) LastRewardedFtsoAddress() (common.Address, error) {
	return _FtsoManager.Contract.LastRewardedFtsoAddress(&_FtsoManager.CallOpts)
}

// NotInitializedFtsos is a free data retrieval call binding the contract method 0x823033a9.
//
// Solidity: function notInitializedFtsos(address _ftso) view returns(bool)
func (_FtsoManager *FtsoManagerCaller) NotInitializedFtsos(opts *bind.CallOpts, _ftso common.Address) (bool, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "notInitializedFtsos", _ftso)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NotInitializedFtsos is a free data retrieval call binding the contract method 0x823033a9.
//
// Solidity: function notInitializedFtsos(address _ftso) view returns(bool)
func (_FtsoManager *FtsoManagerSession) NotInitializedFtsos(_ftso common.Address) (bool, error) {
	return _FtsoManager.Contract.NotInitializedFtsos(&_FtsoManager.CallOpts, _ftso)
}

// NotInitializedFtsos is a free data retrieval call binding the contract method 0x823033a9.
//
// Solidity: function notInitializedFtsos(address _ftso) view returns(bool)
func (_FtsoManager *FtsoManagerCallerSession) NotInitializedFtsos(_ftso common.Address) (bool, error) {
	return _FtsoManager.Contract.NotInitializedFtsos(&_FtsoManager.CallOpts, _ftso)
}

// OldFtsoManager is a free data retrieval call binding the contract method 0xb4dba0f3.
//
// Solidity: function oldFtsoManager() view returns(address)
func (_FtsoManager *FtsoManagerCaller) OldFtsoManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "oldFtsoManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OldFtsoManager is a free data retrieval call binding the contract method 0xb4dba0f3.
//
// Solidity: function oldFtsoManager() view returns(address)
func (_FtsoManager *FtsoManagerSession) OldFtsoManager() (common.Address, error) {
	return _FtsoManager.Contract.OldFtsoManager(&_FtsoManager.CallOpts)
}

// OldFtsoManager is a free data retrieval call binding the contract method 0xb4dba0f3.
//
// Solidity: function oldFtsoManager() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) OldFtsoManager() (common.Address, error) {
	return _FtsoManager.Contract.OldFtsoManager(&_FtsoManager.CallOpts)
}

// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
//
// Solidity: function priceSubmitter() view returns(address)
func (_FtsoManager *FtsoManagerCaller) PriceSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "priceSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
//
// Solidity: function priceSubmitter() view returns(address)
func (_FtsoManager *FtsoManagerSession) PriceSubmitter() (common.Address, error) {
	return _FtsoManager.Contract.PriceSubmitter(&_FtsoManager.CallOpts)
}

// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
//
// Solidity: function priceSubmitter() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) PriceSubmitter() (common.Address, error) {
	return _FtsoManager.Contract.PriceSubmitter(&_FtsoManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtsoManager *FtsoManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtsoManager *FtsoManagerSession) ProductionMode() (bool, error) {
	return _FtsoManager.Contract.ProductionMode(&_FtsoManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtsoManager *FtsoManagerCallerSession) ProductionMode() (bool, error) {
	return _FtsoManager.Contract.ProductionMode(&_FtsoManager.CallOpts)
}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint256)
func (_FtsoManager *FtsoManagerCaller) RewardEpochDurationSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "rewardEpochDurationSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint256)
func (_FtsoManager *FtsoManagerSession) RewardEpochDurationSeconds() (*big.Int, error) {
	return _FtsoManager.Contract.RewardEpochDurationSeconds(&_FtsoManager.CallOpts)
}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint256)
func (_FtsoManager *FtsoManagerCallerSession) RewardEpochDurationSeconds() (*big.Int, error) {
	return _FtsoManager.Contract.RewardEpochDurationSeconds(&_FtsoManager.CallOpts)
}

// RewardEpochs is a free data retrieval call binding the contract method 0xa795f409.
//
// Solidity: function rewardEpochs(uint256 _rewardEpochId) view returns(uint256 _votepowerBlock, uint256 _startBlock, uint256 _startTimestamp)
func (_FtsoManager *FtsoManagerCaller) RewardEpochs(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	VotepowerBlock *big.Int
	StartBlock     *big.Int
	StartTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "rewardEpochs", _rewardEpochId)

	outstruct := new(struct {
		VotepowerBlock *big.Int
		StartBlock     *big.Int
		StartTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VotepowerBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardEpochs is a free data retrieval call binding the contract method 0xa795f409.
//
// Solidity: function rewardEpochs(uint256 _rewardEpochId) view returns(uint256 _votepowerBlock, uint256 _startBlock, uint256 _startTimestamp)
func (_FtsoManager *FtsoManagerSession) RewardEpochs(_rewardEpochId *big.Int) (struct {
	VotepowerBlock *big.Int
	StartBlock     *big.Int
	StartTimestamp *big.Int
}, error) {
	return _FtsoManager.Contract.RewardEpochs(&_FtsoManager.CallOpts, _rewardEpochId)
}

// RewardEpochs is a free data retrieval call binding the contract method 0xa795f409.
//
// Solidity: function rewardEpochs(uint256 _rewardEpochId) view returns(uint256 _votepowerBlock, uint256 _startBlock, uint256 _startTimestamp)
func (_FtsoManager *FtsoManagerCallerSession) RewardEpochs(_rewardEpochId *big.Int) (struct {
	VotepowerBlock *big.Int
	StartBlock     *big.Int
	StartTimestamp *big.Int
}, error) {
	return _FtsoManager.Contract.RewardEpochs(&_FtsoManager.CallOpts, _rewardEpochId)
}

// RewardEpochsStartTs is a free data retrieval call binding the contract method 0xa578f55b.
//
// Solidity: function rewardEpochsStartTs() view returns(uint256)
func (_FtsoManager *FtsoManagerCaller) RewardEpochsStartTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "rewardEpochsStartTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardEpochsStartTs is a free data retrieval call binding the contract method 0xa578f55b.
//
// Solidity: function rewardEpochsStartTs() view returns(uint256)
func (_FtsoManager *FtsoManagerSession) RewardEpochsStartTs() (*big.Int, error) {
	return _FtsoManager.Contract.RewardEpochsStartTs(&_FtsoManager.CallOpts)
}

// RewardEpochsStartTs is a free data retrieval call binding the contract method 0xa578f55b.
//
// Solidity: function rewardEpochsStartTs() view returns(uint256)
func (_FtsoManager *FtsoManagerCallerSession) RewardEpochsStartTs() (*big.Int, error) {
	return _FtsoManager.Contract.RewardEpochsStartTs(&_FtsoManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FtsoManager *FtsoManagerCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FtsoManager *FtsoManagerSession) RewardManager() (common.Address, error) {
	return _FtsoManager.Contract.RewardManager(&_FtsoManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) RewardManager() (common.Address, error) {
	return _FtsoManager.Contract.RewardManager(&_FtsoManager.CallOpts)
}

// Settings is a free data retrieval call binding the contract method 0xe06174e4.
//
// Solidity: function settings() view returns(uint256 maxVotePowerNatThresholdFraction, uint256 maxVotePowerAssetThresholdFraction, uint256 lowAssetUSDThreshold, uint256 highAssetUSDThreshold, uint256 highAssetTurnoutThresholdBIPS, uint256 lowNatTurnoutThresholdBIPS, uint256 rewardExpiryOffsetSeconds, bool changed, bool initialized)
func (_FtsoManager *FtsoManagerCaller) Settings(opts *bind.CallOpts) (struct {
	MaxVotePowerNatThresholdFraction   *big.Int
	MaxVotePowerAssetThresholdFraction *big.Int
	LowAssetUSDThreshold               *big.Int
	HighAssetUSDThreshold              *big.Int
	HighAssetTurnoutThresholdBIPS      *big.Int
	LowNatTurnoutThresholdBIPS         *big.Int
	RewardExpiryOffsetSeconds          *big.Int
	Changed                            bool
	Initialized                        bool
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "settings")

	outstruct := new(struct {
		MaxVotePowerNatThresholdFraction   *big.Int
		MaxVotePowerAssetThresholdFraction *big.Int
		LowAssetUSDThreshold               *big.Int
		HighAssetUSDThreshold              *big.Int
		HighAssetTurnoutThresholdBIPS      *big.Int
		LowNatTurnoutThresholdBIPS         *big.Int
		RewardExpiryOffsetSeconds          *big.Int
		Changed                            bool
		Initialized                        bool
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
	outstruct.RewardExpiryOffsetSeconds = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Changed = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.Initialized = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Settings is a free data retrieval call binding the contract method 0xe06174e4.
//
// Solidity: function settings() view returns(uint256 maxVotePowerNatThresholdFraction, uint256 maxVotePowerAssetThresholdFraction, uint256 lowAssetUSDThreshold, uint256 highAssetUSDThreshold, uint256 highAssetTurnoutThresholdBIPS, uint256 lowNatTurnoutThresholdBIPS, uint256 rewardExpiryOffsetSeconds, bool changed, bool initialized)
func (_FtsoManager *FtsoManagerSession) Settings() (struct {
	MaxVotePowerNatThresholdFraction   *big.Int
	MaxVotePowerAssetThresholdFraction *big.Int
	LowAssetUSDThreshold               *big.Int
	HighAssetUSDThreshold              *big.Int
	HighAssetTurnoutThresholdBIPS      *big.Int
	LowNatTurnoutThresholdBIPS         *big.Int
	RewardExpiryOffsetSeconds          *big.Int
	Changed                            bool
	Initialized                        bool
}, error) {
	return _FtsoManager.Contract.Settings(&_FtsoManager.CallOpts)
}

// Settings is a free data retrieval call binding the contract method 0xe06174e4.
//
// Solidity: function settings() view returns(uint256 maxVotePowerNatThresholdFraction, uint256 maxVotePowerAssetThresholdFraction, uint256 lowAssetUSDThreshold, uint256 highAssetUSDThreshold, uint256 highAssetTurnoutThresholdBIPS, uint256 lowNatTurnoutThresholdBIPS, uint256 rewardExpiryOffsetSeconds, bool changed, bool initialized)
func (_FtsoManager *FtsoManagerCallerSession) Settings() (struct {
	MaxVotePowerNatThresholdFraction   *big.Int
	MaxVotePowerAssetThresholdFraction *big.Int
	LowAssetUSDThreshold               *big.Int
	HighAssetUSDThreshold              *big.Int
	HighAssetTurnoutThresholdBIPS      *big.Int
	LowNatTurnoutThresholdBIPS         *big.Int
	RewardExpiryOffsetSeconds          *big.Int
	Changed                            bool
	Initialized                        bool
}, error) {
	return _FtsoManager.Contract.Settings(&_FtsoManager.CallOpts)
}

// ShowLastRevertedError is a free data retrieval call binding the contract method 0x2b3c41a4.
//
// Solidity: function showLastRevertedError() view returns(uint256[] _lastErrorBlock, uint256[] _numErrors, string[] _errorString, address[] _erroringContract, uint256 _totalRevertedErrors)
func (_FtsoManager *FtsoManagerCaller) ShowLastRevertedError(opts *bind.CallOpts) (struct {
	LastErrorBlock      []*big.Int
	NumErrors           []*big.Int
	ErrorString         []string
	ErroringContract    []common.Address
	TotalRevertedErrors *big.Int
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "showLastRevertedError")

	outstruct := new(struct {
		LastErrorBlock      []*big.Int
		NumErrors           []*big.Int
		ErrorString         []string
		ErroringContract    []common.Address
		TotalRevertedErrors *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastErrorBlock = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.NumErrors = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.ErrorString = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.ErroringContract = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)
	outstruct.TotalRevertedErrors = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ShowLastRevertedError is a free data retrieval call binding the contract method 0x2b3c41a4.
//
// Solidity: function showLastRevertedError() view returns(uint256[] _lastErrorBlock, uint256[] _numErrors, string[] _errorString, address[] _erroringContract, uint256 _totalRevertedErrors)
func (_FtsoManager *FtsoManagerSession) ShowLastRevertedError() (struct {
	LastErrorBlock      []*big.Int
	NumErrors           []*big.Int
	ErrorString         []string
	ErroringContract    []common.Address
	TotalRevertedErrors *big.Int
}, error) {
	return _FtsoManager.Contract.ShowLastRevertedError(&_FtsoManager.CallOpts)
}

// ShowLastRevertedError is a free data retrieval call binding the contract method 0x2b3c41a4.
//
// Solidity: function showLastRevertedError() view returns(uint256[] _lastErrorBlock, uint256[] _numErrors, string[] _errorString, address[] _erroringContract, uint256 _totalRevertedErrors)
func (_FtsoManager *FtsoManagerCallerSession) ShowLastRevertedError() (struct {
	LastErrorBlock      []*big.Int
	NumErrors           []*big.Int
	ErrorString         []string
	ErroringContract    []common.Address
	TotalRevertedErrors *big.Int
}, error) {
	return _FtsoManager.Contract.ShowLastRevertedError(&_FtsoManager.CallOpts)
}

// ShowRevertedErrors is a free data retrieval call binding the contract method 0x6ea0aa31.
//
// Solidity: function showRevertedErrors(uint256 startIndex, uint256 numErrorTypesToShow) view returns(uint256[] _lastErrorBlock, uint256[] _numErrors, string[] _errorString, address[] _erroringContract, uint256 _totalRevertedErrors)
func (_FtsoManager *FtsoManagerCaller) ShowRevertedErrors(opts *bind.CallOpts, startIndex *big.Int, numErrorTypesToShow *big.Int) (struct {
	LastErrorBlock      []*big.Int
	NumErrors           []*big.Int
	ErrorString         []string
	ErroringContract    []common.Address
	TotalRevertedErrors *big.Int
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "showRevertedErrors", startIndex, numErrorTypesToShow)

	outstruct := new(struct {
		LastErrorBlock      []*big.Int
		NumErrors           []*big.Int
		ErrorString         []string
		ErroringContract    []common.Address
		TotalRevertedErrors *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastErrorBlock = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.NumErrors = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.ErrorString = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.ErroringContract = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)
	outstruct.TotalRevertedErrors = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ShowRevertedErrors is a free data retrieval call binding the contract method 0x6ea0aa31.
//
// Solidity: function showRevertedErrors(uint256 startIndex, uint256 numErrorTypesToShow) view returns(uint256[] _lastErrorBlock, uint256[] _numErrors, string[] _errorString, address[] _erroringContract, uint256 _totalRevertedErrors)
func (_FtsoManager *FtsoManagerSession) ShowRevertedErrors(startIndex *big.Int, numErrorTypesToShow *big.Int) (struct {
	LastErrorBlock      []*big.Int
	NumErrors           []*big.Int
	ErrorString         []string
	ErroringContract    []common.Address
	TotalRevertedErrors *big.Int
}, error) {
	return _FtsoManager.Contract.ShowRevertedErrors(&_FtsoManager.CallOpts, startIndex, numErrorTypesToShow)
}

// ShowRevertedErrors is a free data retrieval call binding the contract method 0x6ea0aa31.
//
// Solidity: function showRevertedErrors(uint256 startIndex, uint256 numErrorTypesToShow) view returns(uint256[] _lastErrorBlock, uint256[] _numErrors, string[] _errorString, address[] _erroringContract, uint256 _totalRevertedErrors)
func (_FtsoManager *FtsoManagerCallerSession) ShowRevertedErrors(startIndex *big.Int, numErrorTypesToShow *big.Int) (struct {
	LastErrorBlock      []*big.Int
	NumErrors           []*big.Int
	ErrorString         []string
	ErroringContract    []common.Address
	TotalRevertedErrors *big.Int
}, error) {
	return _FtsoManager.Contract.ShowRevertedErrors(&_FtsoManager.CallOpts, startIndex, numErrorTypesToShow)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(address)
func (_FtsoManager *FtsoManagerCaller) Supply(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(address)
func (_FtsoManager *FtsoManagerSession) Supply() (common.Address, error) {
	return _FtsoManager.Contract.Supply(&_FtsoManager.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) Supply() (common.Address, error) {
	return _FtsoManager.Contract.Supply(&_FtsoManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtsoManager *FtsoManagerCaller) TimelockedCalls(opts *bind.CallOpts, arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "timelockedCalls", arg0)

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
func (_FtsoManager *FtsoManagerSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FtsoManager.Contract.TimelockedCalls(&_FtsoManager.CallOpts, arg0)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtsoManager *FtsoManagerCallerSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FtsoManager.Contract.TimelockedCalls(&_FtsoManager.CallOpts, arg0)
}

// UpdateOnRewardEpochSwitchover is a free data retrieval call binding the contract method 0x5904089a.
//
// Solidity: function updateOnRewardEpochSwitchover() view returns(address)
func (_FtsoManager *FtsoManagerCaller) UpdateOnRewardEpochSwitchover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "updateOnRewardEpochSwitchover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpdateOnRewardEpochSwitchover is a free data retrieval call binding the contract method 0x5904089a.
//
// Solidity: function updateOnRewardEpochSwitchover() view returns(address)
func (_FtsoManager *FtsoManagerSession) UpdateOnRewardEpochSwitchover() (common.Address, error) {
	return _FtsoManager.Contract.UpdateOnRewardEpochSwitchover(&_FtsoManager.CallOpts)
}

// UpdateOnRewardEpochSwitchover is a free data retrieval call binding the contract method 0x5904089a.
//
// Solidity: function updateOnRewardEpochSwitchover() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) UpdateOnRewardEpochSwitchover() (common.Address, error) {
	return _FtsoManager.Contract.UpdateOnRewardEpochSwitchover(&_FtsoManager.CallOpts)
}

// VoterWhitelister is a free data retrieval call binding the contract method 0xc2b0d47b.
//
// Solidity: function voterWhitelister() view returns(address)
func (_FtsoManager *FtsoManagerCaller) VoterWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoManager.contract.Call(opts, &out, "voterWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterWhitelister is a free data retrieval call binding the contract method 0xc2b0d47b.
//
// Solidity: function voterWhitelister() view returns(address)
func (_FtsoManager *FtsoManagerSession) VoterWhitelister() (common.Address, error) {
	return _FtsoManager.Contract.VoterWhitelister(&_FtsoManager.CallOpts)
}

// VoterWhitelister is a free data retrieval call binding the contract method 0xc2b0d47b.
//
// Solidity: function voterWhitelister() view returns(address)
func (_FtsoManager *FtsoManagerCallerSession) VoterWhitelister() (common.Address, error) {
	return _FtsoManager.Contract.VoterWhitelister(&_FtsoManager.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_FtsoManager *FtsoManagerTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_FtsoManager *FtsoManagerSession) Activate() (*types.Transaction, error) {
	return _FtsoManager.Contract.Activate(&_FtsoManager.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_FtsoManager *FtsoManagerTransactorSession) Activate() (*types.Transaction, error) {
	return _FtsoManager.Contract.Activate(&_FtsoManager.TransactOpts)
}

// AddFtso is a paid mutator transaction binding the contract method 0x2663f1b4.
//
// Solidity: function addFtso(address _ftso) returns()
func (_FtsoManager *FtsoManagerTransactor) AddFtso(opts *bind.TransactOpts, _ftso common.Address) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "addFtso", _ftso)
}

// AddFtso is a paid mutator transaction binding the contract method 0x2663f1b4.
//
// Solidity: function addFtso(address _ftso) returns()
func (_FtsoManager *FtsoManagerSession) AddFtso(_ftso common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.AddFtso(&_FtsoManager.TransactOpts, _ftso)
}

// AddFtso is a paid mutator transaction binding the contract method 0x2663f1b4.
//
// Solidity: function addFtso(address _ftso) returns()
func (_FtsoManager *FtsoManagerTransactorSession) AddFtso(_ftso common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.AddFtso(&_FtsoManager.TransactOpts, _ftso)
}

// AddFtsosBulk is a paid mutator transaction binding the contract method 0xd429cfe5.
//
// Solidity: function addFtsosBulk(address[] _ftsos) returns()
func (_FtsoManager *FtsoManagerTransactor) AddFtsosBulk(opts *bind.TransactOpts, _ftsos []common.Address) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "addFtsosBulk", _ftsos)
}

// AddFtsosBulk is a paid mutator transaction binding the contract method 0xd429cfe5.
//
// Solidity: function addFtsosBulk(address[] _ftsos) returns()
func (_FtsoManager *FtsoManagerSession) AddFtsosBulk(_ftsos []common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.AddFtsosBulk(&_FtsoManager.TransactOpts, _ftsos)
}

// AddFtsosBulk is a paid mutator transaction binding the contract method 0xd429cfe5.
//
// Solidity: function addFtsosBulk(address[] _ftsos) returns()
func (_FtsoManager *FtsoManagerTransactorSession) AddFtsosBulk(_ftsos []common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.AddFtsosBulk(&_FtsoManager.TransactOpts, _ftsos)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtsoManager *FtsoManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtsoManager *FtsoManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtsoManager.Contract.CancelGovernanceCall(&_FtsoManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtsoManager *FtsoManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtsoManager.Contract.CancelGovernanceCall(&_FtsoManager.TransactOpts, _selector)
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FtsoManager *FtsoManagerTransactor) Daemonize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "daemonize")
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FtsoManager *FtsoManagerSession) Daemonize() (*types.Transaction, error) {
	return _FtsoManager.Contract.Daemonize(&_FtsoManager.TransactOpts)
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FtsoManager *FtsoManagerTransactorSession) Daemonize() (*types.Transaction, error) {
	return _FtsoManager.Contract.Daemonize(&_FtsoManager.TransactOpts)
}

// DeactivateFtsos is a paid mutator transaction binding the contract method 0x8de306b1.
//
// Solidity: function deactivateFtsos(address[] _ftsos) returns()
func (_FtsoManager *FtsoManagerTransactor) DeactivateFtsos(opts *bind.TransactOpts, _ftsos []common.Address) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "deactivateFtsos", _ftsos)
}

// DeactivateFtsos is a paid mutator transaction binding the contract method 0x8de306b1.
//
// Solidity: function deactivateFtsos(address[] _ftsos) returns()
func (_FtsoManager *FtsoManagerSession) DeactivateFtsos(_ftsos []common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.DeactivateFtsos(&_FtsoManager.TransactOpts, _ftsos)
}

// DeactivateFtsos is a paid mutator transaction binding the contract method 0x8de306b1.
//
// Solidity: function deactivateFtsos(address[] _ftsos) returns()
func (_FtsoManager *FtsoManagerTransactorSession) DeactivateFtsos(_ftsos []common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.DeactivateFtsos(&_FtsoManager.TransactOpts, _ftsos)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtsoManager *FtsoManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtsoManager *FtsoManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtsoManager.Contract.ExecuteGovernanceCall(&_FtsoManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtsoManager *FtsoManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtsoManager.Contract.ExecuteGovernanceCall(&_FtsoManager.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_FtsoManager *FtsoManagerTransactor) Initialise(opts *bind.TransactOpts, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "initialise", _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_FtsoManager *FtsoManagerSession) Initialise(_initialGovernance common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.Initialise(&_FtsoManager.TransactOpts, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_FtsoManager *FtsoManagerTransactorSession) Initialise(_initialGovernance common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.Initialise(&_FtsoManager.TransactOpts, _initialGovernance)
}

// RemoveFtso is a paid mutator transaction binding the contract method 0xa670ff87.
//
// Solidity: function removeFtso(address _ftso) returns()
func (_FtsoManager *FtsoManagerTransactor) RemoveFtso(opts *bind.TransactOpts, _ftso common.Address) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "removeFtso", _ftso)
}

// RemoveFtso is a paid mutator transaction binding the contract method 0xa670ff87.
//
// Solidity: function removeFtso(address _ftso) returns()
func (_FtsoManager *FtsoManagerSession) RemoveFtso(_ftso common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.RemoveFtso(&_FtsoManager.TransactOpts, _ftso)
}

// RemoveFtso is a paid mutator transaction binding the contract method 0xa670ff87.
//
// Solidity: function removeFtso(address _ftso) returns()
func (_FtsoManager *FtsoManagerTransactorSession) RemoveFtso(_ftso common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.RemoveFtso(&_FtsoManager.TransactOpts, _ftso)
}

// ReplaceFtso is a paid mutator transaction binding the contract method 0x3758e679.
//
// Solidity: function replaceFtso(address _ftsoToAdd, bool _copyCurrentPrice, bool _copyAssetOrAssetFtsos) returns()
func (_FtsoManager *FtsoManagerTransactor) ReplaceFtso(opts *bind.TransactOpts, _ftsoToAdd common.Address, _copyCurrentPrice bool, _copyAssetOrAssetFtsos bool) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "replaceFtso", _ftsoToAdd, _copyCurrentPrice, _copyAssetOrAssetFtsos)
}

// ReplaceFtso is a paid mutator transaction binding the contract method 0x3758e679.
//
// Solidity: function replaceFtso(address _ftsoToAdd, bool _copyCurrentPrice, bool _copyAssetOrAssetFtsos) returns()
func (_FtsoManager *FtsoManagerSession) ReplaceFtso(_ftsoToAdd common.Address, _copyCurrentPrice bool, _copyAssetOrAssetFtsos bool) (*types.Transaction, error) {
	return _FtsoManager.Contract.ReplaceFtso(&_FtsoManager.TransactOpts, _ftsoToAdd, _copyCurrentPrice, _copyAssetOrAssetFtsos)
}

// ReplaceFtso is a paid mutator transaction binding the contract method 0x3758e679.
//
// Solidity: function replaceFtso(address _ftsoToAdd, bool _copyCurrentPrice, bool _copyAssetOrAssetFtsos) returns()
func (_FtsoManager *FtsoManagerTransactorSession) ReplaceFtso(_ftsoToAdd common.Address, _copyCurrentPrice bool, _copyAssetOrAssetFtsos bool) (*types.Transaction, error) {
	return _FtsoManager.Contract.ReplaceFtso(&_FtsoManager.TransactOpts, _ftsoToAdd, _copyCurrentPrice, _copyAssetOrAssetFtsos)
}

// ReplaceFtsosBulk is a paid mutator transaction binding the contract method 0x758ff1da.
//
// Solidity: function replaceFtsosBulk(address[] _ftsosToAdd, bool _copyCurrentPrice, bool _copyAssetOrAssetFtsos) returns()
func (_FtsoManager *FtsoManagerTransactor) ReplaceFtsosBulk(opts *bind.TransactOpts, _ftsosToAdd []common.Address, _copyCurrentPrice bool, _copyAssetOrAssetFtsos bool) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "replaceFtsosBulk", _ftsosToAdd, _copyCurrentPrice, _copyAssetOrAssetFtsos)
}

// ReplaceFtsosBulk is a paid mutator transaction binding the contract method 0x758ff1da.
//
// Solidity: function replaceFtsosBulk(address[] _ftsosToAdd, bool _copyCurrentPrice, bool _copyAssetOrAssetFtsos) returns()
func (_FtsoManager *FtsoManagerSession) ReplaceFtsosBulk(_ftsosToAdd []common.Address, _copyCurrentPrice bool, _copyAssetOrAssetFtsos bool) (*types.Transaction, error) {
	return _FtsoManager.Contract.ReplaceFtsosBulk(&_FtsoManager.TransactOpts, _ftsosToAdd, _copyCurrentPrice, _copyAssetOrAssetFtsos)
}

// ReplaceFtsosBulk is a paid mutator transaction binding the contract method 0x758ff1da.
//
// Solidity: function replaceFtsosBulk(address[] _ftsosToAdd, bool _copyCurrentPrice, bool _copyAssetOrAssetFtsos) returns()
func (_FtsoManager *FtsoManagerTransactorSession) ReplaceFtsosBulk(_ftsosToAdd []common.Address, _copyCurrentPrice bool, _copyAssetOrAssetFtsos bool) (*types.Transaction, error) {
	return _FtsoManager.Contract.ReplaceFtsosBulk(&_FtsoManager.TransactOpts, _ftsosToAdd, _copyCurrentPrice, _copyAssetOrAssetFtsos)
}

// SetFallbackMode is a paid mutator transaction binding the contract method 0xff882fbb.
//
// Solidity: function setFallbackMode(bool _fallbackMode) returns()
func (_FtsoManager *FtsoManagerTransactor) SetFallbackMode(opts *bind.TransactOpts, _fallbackMode bool) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "setFallbackMode", _fallbackMode)
}

// SetFallbackMode is a paid mutator transaction binding the contract method 0xff882fbb.
//
// Solidity: function setFallbackMode(bool _fallbackMode) returns()
func (_FtsoManager *FtsoManagerSession) SetFallbackMode(_fallbackMode bool) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetFallbackMode(&_FtsoManager.TransactOpts, _fallbackMode)
}

// SetFallbackMode is a paid mutator transaction binding the contract method 0xff882fbb.
//
// Solidity: function setFallbackMode(bool _fallbackMode) returns()
func (_FtsoManager *FtsoManagerTransactorSession) SetFallbackMode(_fallbackMode bool) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetFallbackMode(&_FtsoManager.TransactOpts, _fallbackMode)
}

// SetFtsoAsset is a paid mutator transaction binding the contract method 0x6b65cc34.
//
// Solidity: function setFtsoAsset(address _ftso, address _asset) returns()
func (_FtsoManager *FtsoManagerTransactor) SetFtsoAsset(opts *bind.TransactOpts, _ftso common.Address, _asset common.Address) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "setFtsoAsset", _ftso, _asset)
}

// SetFtsoAsset is a paid mutator transaction binding the contract method 0x6b65cc34.
//
// Solidity: function setFtsoAsset(address _ftso, address _asset) returns()
func (_FtsoManager *FtsoManagerSession) SetFtsoAsset(_ftso common.Address, _asset common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetFtsoAsset(&_FtsoManager.TransactOpts, _ftso, _asset)
}

// SetFtsoAsset is a paid mutator transaction binding the contract method 0x6b65cc34.
//
// Solidity: function setFtsoAsset(address _ftso, address _asset) returns()
func (_FtsoManager *FtsoManagerTransactorSession) SetFtsoAsset(_ftso common.Address, _asset common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetFtsoAsset(&_FtsoManager.TransactOpts, _ftso, _asset)
}

// SetFtsoAssetFtsos is a paid mutator transaction binding the contract method 0xa93a6f42.
//
// Solidity: function setFtsoAssetFtsos(address _ftso, address[] _assetFtsos) returns()
func (_FtsoManager *FtsoManagerTransactor) SetFtsoAssetFtsos(opts *bind.TransactOpts, _ftso common.Address, _assetFtsos []common.Address) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "setFtsoAssetFtsos", _ftso, _assetFtsos)
}

// SetFtsoAssetFtsos is a paid mutator transaction binding the contract method 0xa93a6f42.
//
// Solidity: function setFtsoAssetFtsos(address _ftso, address[] _assetFtsos) returns()
func (_FtsoManager *FtsoManagerSession) SetFtsoAssetFtsos(_ftso common.Address, _assetFtsos []common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetFtsoAssetFtsos(&_FtsoManager.TransactOpts, _ftso, _assetFtsos)
}

// SetFtsoAssetFtsos is a paid mutator transaction binding the contract method 0xa93a6f42.
//
// Solidity: function setFtsoAssetFtsos(address _ftso, address[] _assetFtsos) returns()
func (_FtsoManager *FtsoManagerTransactorSession) SetFtsoAssetFtsos(_ftso common.Address, _assetFtsos []common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetFtsoAssetFtsos(&_FtsoManager.TransactOpts, _ftso, _assetFtsos)
}

// SetFtsoFallbackMode is a paid mutator transaction binding the contract method 0xaf946af7.
//
// Solidity: function setFtsoFallbackMode(address _ftso, bool _fallbackMode) returns()
func (_FtsoManager *FtsoManagerTransactor) SetFtsoFallbackMode(opts *bind.TransactOpts, _ftso common.Address, _fallbackMode bool) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "setFtsoFallbackMode", _ftso, _fallbackMode)
}

// SetFtsoFallbackMode is a paid mutator transaction binding the contract method 0xaf946af7.
//
// Solidity: function setFtsoFallbackMode(address _ftso, bool _fallbackMode) returns()
func (_FtsoManager *FtsoManagerSession) SetFtsoFallbackMode(_ftso common.Address, _fallbackMode bool) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetFtsoFallbackMode(&_FtsoManager.TransactOpts, _ftso, _fallbackMode)
}

// SetFtsoFallbackMode is a paid mutator transaction binding the contract method 0xaf946af7.
//
// Solidity: function setFtsoFallbackMode(address _ftso, bool _fallbackMode) returns()
func (_FtsoManager *FtsoManagerTransactorSession) SetFtsoFallbackMode(_ftso common.Address, _fallbackMode bool) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetFtsoFallbackMode(&_FtsoManager.TransactOpts, _ftso, _fallbackMode)
}

// SetGovernanceParameters is a paid mutator transaction binding the contract method 0x9131205b.
//
// Solidity: function setGovernanceParameters(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, uint256 _rewardExpiryOffsetSeconds, address[] _trustedAddresses) returns()
func (_FtsoManager *FtsoManagerTransactor) SetGovernanceParameters(opts *bind.TransactOpts, _maxVotePowerNatThresholdFraction *big.Int, _maxVotePowerAssetThresholdFraction *big.Int, _lowAssetUSDThreshold *big.Int, _highAssetUSDThreshold *big.Int, _highAssetTurnoutThresholdBIPS *big.Int, _lowNatTurnoutThresholdBIPS *big.Int, _rewardExpiryOffsetSeconds *big.Int, _trustedAddresses []common.Address) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "setGovernanceParameters", _maxVotePowerNatThresholdFraction, _maxVotePowerAssetThresholdFraction, _lowAssetUSDThreshold, _highAssetUSDThreshold, _highAssetTurnoutThresholdBIPS, _lowNatTurnoutThresholdBIPS, _rewardExpiryOffsetSeconds, _trustedAddresses)
}

// SetGovernanceParameters is a paid mutator transaction binding the contract method 0x9131205b.
//
// Solidity: function setGovernanceParameters(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, uint256 _rewardExpiryOffsetSeconds, address[] _trustedAddresses) returns()
func (_FtsoManager *FtsoManagerSession) SetGovernanceParameters(_maxVotePowerNatThresholdFraction *big.Int, _maxVotePowerAssetThresholdFraction *big.Int, _lowAssetUSDThreshold *big.Int, _highAssetUSDThreshold *big.Int, _highAssetTurnoutThresholdBIPS *big.Int, _lowNatTurnoutThresholdBIPS *big.Int, _rewardExpiryOffsetSeconds *big.Int, _trustedAddresses []common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetGovernanceParameters(&_FtsoManager.TransactOpts, _maxVotePowerNatThresholdFraction, _maxVotePowerAssetThresholdFraction, _lowAssetUSDThreshold, _highAssetUSDThreshold, _highAssetTurnoutThresholdBIPS, _lowNatTurnoutThresholdBIPS, _rewardExpiryOffsetSeconds, _trustedAddresses)
}

// SetGovernanceParameters is a paid mutator transaction binding the contract method 0x9131205b.
//
// Solidity: function setGovernanceParameters(uint256 _maxVotePowerNatThresholdFraction, uint256 _maxVotePowerAssetThresholdFraction, uint256 _lowAssetUSDThreshold, uint256 _highAssetUSDThreshold, uint256 _highAssetTurnoutThresholdBIPS, uint256 _lowNatTurnoutThresholdBIPS, uint256 _rewardExpiryOffsetSeconds, address[] _trustedAddresses) returns()
func (_FtsoManager *FtsoManagerTransactorSession) SetGovernanceParameters(_maxVotePowerNatThresholdFraction *big.Int, _maxVotePowerAssetThresholdFraction *big.Int, _lowAssetUSDThreshold *big.Int, _highAssetUSDThreshold *big.Int, _highAssetTurnoutThresholdBIPS *big.Int, _lowNatTurnoutThresholdBIPS *big.Int, _rewardExpiryOffsetSeconds *big.Int, _trustedAddresses []common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetGovernanceParameters(&_FtsoManager.TransactOpts, _maxVotePowerNatThresholdFraction, _maxVotePowerAssetThresholdFraction, _lowAssetUSDThreshold, _highAssetUSDThreshold, _highAssetTurnoutThresholdBIPS, _lowNatTurnoutThresholdBIPS, _rewardExpiryOffsetSeconds, _trustedAddresses)
}

// SetInitialRewardData is a paid mutator transaction binding the contract method 0xe080a970.
//
// Solidity: function setInitialRewardData(uint256 _nextRewardEpochToExpire, uint256 _rewardEpochsLength, uint256 _currentRewardEpochEnds) returns()
func (_FtsoManager *FtsoManagerTransactor) SetInitialRewardData(opts *bind.TransactOpts, _nextRewardEpochToExpire *big.Int, _rewardEpochsLength *big.Int, _currentRewardEpochEnds *big.Int) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "setInitialRewardData", _nextRewardEpochToExpire, _rewardEpochsLength, _currentRewardEpochEnds)
}

// SetInitialRewardData is a paid mutator transaction binding the contract method 0xe080a970.
//
// Solidity: function setInitialRewardData(uint256 _nextRewardEpochToExpire, uint256 _rewardEpochsLength, uint256 _currentRewardEpochEnds) returns()
func (_FtsoManager *FtsoManagerSession) SetInitialRewardData(_nextRewardEpochToExpire *big.Int, _rewardEpochsLength *big.Int, _currentRewardEpochEnds *big.Int) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetInitialRewardData(&_FtsoManager.TransactOpts, _nextRewardEpochToExpire, _rewardEpochsLength, _currentRewardEpochEnds)
}

// SetInitialRewardData is a paid mutator transaction binding the contract method 0xe080a970.
//
// Solidity: function setInitialRewardData(uint256 _nextRewardEpochToExpire, uint256 _rewardEpochsLength, uint256 _currentRewardEpochEnds) returns()
func (_FtsoManager *FtsoManagerTransactorSession) SetInitialRewardData(_nextRewardEpochToExpire *big.Int, _rewardEpochsLength *big.Int, _currentRewardEpochEnds *big.Int) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetInitialRewardData(&_FtsoManager.TransactOpts, _nextRewardEpochToExpire, _rewardEpochsLength, _currentRewardEpochEnds)
}

// SetRewardEpochDurationSeconds is a paid mutator transaction binding the contract method 0x132c7e1f.
//
// Solidity: function setRewardEpochDurationSeconds(uint256 _rewardEpochDurationSeconds) returns()
func (_FtsoManager *FtsoManagerTransactor) SetRewardEpochDurationSeconds(opts *bind.TransactOpts, _rewardEpochDurationSeconds *big.Int) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "setRewardEpochDurationSeconds", _rewardEpochDurationSeconds)
}

// SetRewardEpochDurationSeconds is a paid mutator transaction binding the contract method 0x132c7e1f.
//
// Solidity: function setRewardEpochDurationSeconds(uint256 _rewardEpochDurationSeconds) returns()
func (_FtsoManager *FtsoManagerSession) SetRewardEpochDurationSeconds(_rewardEpochDurationSeconds *big.Int) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetRewardEpochDurationSeconds(&_FtsoManager.TransactOpts, _rewardEpochDurationSeconds)
}

// SetRewardEpochDurationSeconds is a paid mutator transaction binding the contract method 0x132c7e1f.
//
// Solidity: function setRewardEpochDurationSeconds(uint256 _rewardEpochDurationSeconds) returns()
func (_FtsoManager *FtsoManagerTransactorSession) SetRewardEpochDurationSeconds(_rewardEpochDurationSeconds *big.Int) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetRewardEpochDurationSeconds(&_FtsoManager.TransactOpts, _rewardEpochDurationSeconds)
}

// SetUpdateOnRewardEpochSwitchover is a paid mutator transaction binding the contract method 0x3fdeb7e1.
//
// Solidity: function setUpdateOnRewardEpochSwitchover(address _updateValidators) returns()
func (_FtsoManager *FtsoManagerTransactor) SetUpdateOnRewardEpochSwitchover(opts *bind.TransactOpts, _updateValidators common.Address) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "setUpdateOnRewardEpochSwitchover", _updateValidators)
}

// SetUpdateOnRewardEpochSwitchover is a paid mutator transaction binding the contract method 0x3fdeb7e1.
//
// Solidity: function setUpdateOnRewardEpochSwitchover(address _updateValidators) returns()
func (_FtsoManager *FtsoManagerSession) SetUpdateOnRewardEpochSwitchover(_updateValidators common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetUpdateOnRewardEpochSwitchover(&_FtsoManager.TransactOpts, _updateValidators)
}

// SetUpdateOnRewardEpochSwitchover is a paid mutator transaction binding the contract method 0x3fdeb7e1.
//
// Solidity: function setUpdateOnRewardEpochSwitchover(address _updateValidators) returns()
func (_FtsoManager *FtsoManagerTransactorSession) SetUpdateOnRewardEpochSwitchover(_updateValidators common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetUpdateOnRewardEpochSwitchover(&_FtsoManager.TransactOpts, _updateValidators)
}

// SetVotePowerIntervalFraction is a paid mutator transaction binding the contract method 0x361b5459.
//
// Solidity: function setVotePowerIntervalFraction(uint256 _votePowerIntervalFraction) returns()
func (_FtsoManager *FtsoManagerTransactor) SetVotePowerIntervalFraction(opts *bind.TransactOpts, _votePowerIntervalFraction *big.Int) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "setVotePowerIntervalFraction", _votePowerIntervalFraction)
}

// SetVotePowerIntervalFraction is a paid mutator transaction binding the contract method 0x361b5459.
//
// Solidity: function setVotePowerIntervalFraction(uint256 _votePowerIntervalFraction) returns()
func (_FtsoManager *FtsoManagerSession) SetVotePowerIntervalFraction(_votePowerIntervalFraction *big.Int) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetVotePowerIntervalFraction(&_FtsoManager.TransactOpts, _votePowerIntervalFraction)
}

// SetVotePowerIntervalFraction is a paid mutator transaction binding the contract method 0x361b5459.
//
// Solidity: function setVotePowerIntervalFraction(uint256 _votePowerIntervalFraction) returns()
func (_FtsoManager *FtsoManagerTransactorSession) SetVotePowerIntervalFraction(_votePowerIntervalFraction *big.Int) (*types.Transaction, error) {
	return _FtsoManager.Contract.SetVotePowerIntervalFraction(&_FtsoManager.TransactOpts, _votePowerIntervalFraction)
}

// SwitchToFallbackMode is a paid mutator transaction binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() returns(bool)
func (_FtsoManager *FtsoManagerTransactor) SwitchToFallbackMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "switchToFallbackMode")
}

// SwitchToFallbackMode is a paid mutator transaction binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() returns(bool)
func (_FtsoManager *FtsoManagerSession) SwitchToFallbackMode() (*types.Transaction, error) {
	return _FtsoManager.Contract.SwitchToFallbackMode(&_FtsoManager.TransactOpts)
}

// SwitchToFallbackMode is a paid mutator transaction binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() returns(bool)
func (_FtsoManager *FtsoManagerTransactorSession) SwitchToFallbackMode() (*types.Transaction, error) {
	return _FtsoManager.Contract.SwitchToFallbackMode(&_FtsoManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtsoManager *FtsoManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtsoManager *FtsoManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FtsoManager.Contract.SwitchToProductionMode(&_FtsoManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtsoManager *FtsoManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FtsoManager.Contract.SwitchToProductionMode(&_FtsoManager.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FtsoManager *FtsoManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FtsoManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FtsoManager *FtsoManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.UpdateContractAddresses(&_FtsoManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FtsoManager *FtsoManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FtsoManager.Contract.UpdateContractAddresses(&_FtsoManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// FtsoManagerAccruingUnearnedRewardsFailedIterator is returned from FilterAccruingUnearnedRewardsFailed and is used to iterate over the raw logs and unpacked data for AccruingUnearnedRewardsFailed events raised by the FtsoManager contract.
type FtsoManagerAccruingUnearnedRewardsFailedIterator struct {
	Event *FtsoManagerAccruingUnearnedRewardsFailed // Event containing the contract specifics and raw log

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
func (it *FtsoManagerAccruingUnearnedRewardsFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerAccruingUnearnedRewardsFailed)
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
		it.Event = new(FtsoManagerAccruingUnearnedRewardsFailed)
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
func (it *FtsoManagerAccruingUnearnedRewardsFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerAccruingUnearnedRewardsFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerAccruingUnearnedRewardsFailed represents a AccruingUnearnedRewardsFailed event raised by the FtsoManager contract.
type FtsoManagerAccruingUnearnedRewardsFailed struct {
	EpochId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccruingUnearnedRewardsFailed is a free log retrieval operation binding the contract event 0x8eb60f903ef61e0e490d7d7ba6e5b85cd949ebece7a5e5b3346eb046c041413f.
//
// Solidity: event AccruingUnearnedRewardsFailed(uint256 epochId)
func (_FtsoManager *FtsoManagerFilterer) FilterAccruingUnearnedRewardsFailed(opts *bind.FilterOpts) (*FtsoManagerAccruingUnearnedRewardsFailedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "AccruingUnearnedRewardsFailed")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerAccruingUnearnedRewardsFailedIterator{contract: _FtsoManager.contract, event: "AccruingUnearnedRewardsFailed", logs: logs, sub: sub}, nil
}

// WatchAccruingUnearnedRewardsFailed is a free log subscription operation binding the contract event 0x8eb60f903ef61e0e490d7d7ba6e5b85cd949ebece7a5e5b3346eb046c041413f.
//
// Solidity: event AccruingUnearnedRewardsFailed(uint256 epochId)
func (_FtsoManager *FtsoManagerFilterer) WatchAccruingUnearnedRewardsFailed(opts *bind.WatchOpts, sink chan<- *FtsoManagerAccruingUnearnedRewardsFailed) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "AccruingUnearnedRewardsFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerAccruingUnearnedRewardsFailed)
				if err := _FtsoManager.contract.UnpackLog(event, "AccruingUnearnedRewardsFailed", log); err != nil {
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

// ParseAccruingUnearnedRewardsFailed is a log parse operation binding the contract event 0x8eb60f903ef61e0e490d7d7ba6e5b85cd949ebece7a5e5b3346eb046c041413f.
//
// Solidity: event AccruingUnearnedRewardsFailed(uint256 epochId)
func (_FtsoManager *FtsoManagerFilterer) ParseAccruingUnearnedRewardsFailed(log types.Log) (*FtsoManagerAccruingUnearnedRewardsFailed, error) {
	event := new(FtsoManagerAccruingUnearnedRewardsFailed)
	if err := _FtsoManager.contract.UnpackLog(event, "AccruingUnearnedRewardsFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerCleanupBlockNumberManagerFailedForBlockIterator is returned from FilterCleanupBlockNumberManagerFailedForBlock and is used to iterate over the raw logs and unpacked data for CleanupBlockNumberManagerFailedForBlock events raised by the FtsoManager contract.
type FtsoManagerCleanupBlockNumberManagerFailedForBlockIterator struct {
	Event *FtsoManagerCleanupBlockNumberManagerFailedForBlock // Event containing the contract specifics and raw log

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
func (it *FtsoManagerCleanupBlockNumberManagerFailedForBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerCleanupBlockNumberManagerFailedForBlock)
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
		it.Event = new(FtsoManagerCleanupBlockNumberManagerFailedForBlock)
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
func (it *FtsoManagerCleanupBlockNumberManagerFailedForBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerCleanupBlockNumberManagerFailedForBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerCleanupBlockNumberManagerFailedForBlock represents a CleanupBlockNumberManagerFailedForBlock event raised by the FtsoManager contract.
type FtsoManagerCleanupBlockNumberManagerFailedForBlock struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCleanupBlockNumberManagerFailedForBlock is a free log retrieval operation binding the contract event 0x9f874ea08c7014cce74622bfe71434f81aba7598ad65126a6aea86945bdfa18d.
//
// Solidity: event CleanupBlockNumberManagerFailedForBlock(uint256 blockNumber)
func (_FtsoManager *FtsoManagerFilterer) FilterCleanupBlockNumberManagerFailedForBlock(opts *bind.FilterOpts) (*FtsoManagerCleanupBlockNumberManagerFailedForBlockIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "CleanupBlockNumberManagerFailedForBlock")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerCleanupBlockNumberManagerFailedForBlockIterator{contract: _FtsoManager.contract, event: "CleanupBlockNumberManagerFailedForBlock", logs: logs, sub: sub}, nil
}

// WatchCleanupBlockNumberManagerFailedForBlock is a free log subscription operation binding the contract event 0x9f874ea08c7014cce74622bfe71434f81aba7598ad65126a6aea86945bdfa18d.
//
// Solidity: event CleanupBlockNumberManagerFailedForBlock(uint256 blockNumber)
func (_FtsoManager *FtsoManagerFilterer) WatchCleanupBlockNumberManagerFailedForBlock(opts *bind.WatchOpts, sink chan<- *FtsoManagerCleanupBlockNumberManagerFailedForBlock) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "CleanupBlockNumberManagerFailedForBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerCleanupBlockNumberManagerFailedForBlock)
				if err := _FtsoManager.contract.UnpackLog(event, "CleanupBlockNumberManagerFailedForBlock", log); err != nil {
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

// ParseCleanupBlockNumberManagerFailedForBlock is a log parse operation binding the contract event 0x9f874ea08c7014cce74622bfe71434f81aba7598ad65126a6aea86945bdfa18d.
//
// Solidity: event CleanupBlockNumberManagerFailedForBlock(uint256 blockNumber)
func (_FtsoManager *FtsoManagerFilterer) ParseCleanupBlockNumberManagerFailedForBlock(log types.Log) (*FtsoManagerCleanupBlockNumberManagerFailedForBlock, error) {
	event := new(FtsoManagerCleanupBlockNumberManagerFailedForBlock)
	if err := _FtsoManager.contract.UnpackLog(event, "CleanupBlockNumberManagerFailedForBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerClosingExpiredRewardEpochFailedIterator is returned from FilterClosingExpiredRewardEpochFailed and is used to iterate over the raw logs and unpacked data for ClosingExpiredRewardEpochFailed events raised by the FtsoManager contract.
type FtsoManagerClosingExpiredRewardEpochFailedIterator struct {
	Event *FtsoManagerClosingExpiredRewardEpochFailed // Event containing the contract specifics and raw log

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
func (it *FtsoManagerClosingExpiredRewardEpochFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerClosingExpiredRewardEpochFailed)
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
		it.Event = new(FtsoManagerClosingExpiredRewardEpochFailed)
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
func (it *FtsoManagerClosingExpiredRewardEpochFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerClosingExpiredRewardEpochFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerClosingExpiredRewardEpochFailed represents a ClosingExpiredRewardEpochFailed event raised by the FtsoManager contract.
type FtsoManagerClosingExpiredRewardEpochFailed struct {
	RewardEpoch *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClosingExpiredRewardEpochFailed is a free log retrieval operation binding the contract event 0xa819a21065ad87bdde9e6d398d3213e0d3634afd87aceb7092236483f5d7ca8d.
//
// Solidity: event ClosingExpiredRewardEpochFailed(uint256 rewardEpoch)
func (_FtsoManager *FtsoManagerFilterer) FilterClosingExpiredRewardEpochFailed(opts *bind.FilterOpts) (*FtsoManagerClosingExpiredRewardEpochFailedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "ClosingExpiredRewardEpochFailed")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerClosingExpiredRewardEpochFailedIterator{contract: _FtsoManager.contract, event: "ClosingExpiredRewardEpochFailed", logs: logs, sub: sub}, nil
}

// WatchClosingExpiredRewardEpochFailed is a free log subscription operation binding the contract event 0xa819a21065ad87bdde9e6d398d3213e0d3634afd87aceb7092236483f5d7ca8d.
//
// Solidity: event ClosingExpiredRewardEpochFailed(uint256 rewardEpoch)
func (_FtsoManager *FtsoManagerFilterer) WatchClosingExpiredRewardEpochFailed(opts *bind.WatchOpts, sink chan<- *FtsoManagerClosingExpiredRewardEpochFailed) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "ClosingExpiredRewardEpochFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerClosingExpiredRewardEpochFailed)
				if err := _FtsoManager.contract.UnpackLog(event, "ClosingExpiredRewardEpochFailed", log); err != nil {
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

// ParseClosingExpiredRewardEpochFailed is a log parse operation binding the contract event 0xa819a21065ad87bdde9e6d398d3213e0d3634afd87aceb7092236483f5d7ca8d.
//
// Solidity: event ClosingExpiredRewardEpochFailed(uint256 rewardEpoch)
func (_FtsoManager *FtsoManagerFilterer) ParseClosingExpiredRewardEpochFailed(log types.Log) (*FtsoManagerClosingExpiredRewardEpochFailed, error) {
	event := new(FtsoManagerClosingExpiredRewardEpochFailed)
	if err := _FtsoManager.contract.UnpackLog(event, "ClosingExpiredRewardEpochFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerContractRevertErrorIterator is returned from FilterContractRevertError and is used to iterate over the raw logs and unpacked data for ContractRevertError events raised by the FtsoManager contract.
type FtsoManagerContractRevertErrorIterator struct {
	Event *FtsoManagerContractRevertError // Event containing the contract specifics and raw log

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
func (it *FtsoManagerContractRevertErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerContractRevertError)
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
		it.Event = new(FtsoManagerContractRevertError)
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
func (it *FtsoManagerContractRevertErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerContractRevertErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerContractRevertError represents a ContractRevertError event raised by the FtsoManager contract.
type FtsoManagerContractRevertError struct {
	TheContract common.Address
	AtBlock     *big.Int
	TheMessage  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractRevertError is a free log retrieval operation binding the contract event 0x1a601cf5e0efbd558b2778b7389af04741d1c49bcab104c40daa2da194593617.
//
// Solidity: event ContractRevertError(address theContract, uint256 atBlock, string theMessage)
func (_FtsoManager *FtsoManagerFilterer) FilterContractRevertError(opts *bind.FilterOpts) (*FtsoManagerContractRevertErrorIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "ContractRevertError")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerContractRevertErrorIterator{contract: _FtsoManager.contract, event: "ContractRevertError", logs: logs, sub: sub}, nil
}

// WatchContractRevertError is a free log subscription operation binding the contract event 0x1a601cf5e0efbd558b2778b7389af04741d1c49bcab104c40daa2da194593617.
//
// Solidity: event ContractRevertError(address theContract, uint256 atBlock, string theMessage)
func (_FtsoManager *FtsoManagerFilterer) WatchContractRevertError(opts *bind.WatchOpts, sink chan<- *FtsoManagerContractRevertError) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "ContractRevertError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerContractRevertError)
				if err := _FtsoManager.contract.UnpackLog(event, "ContractRevertError", log); err != nil {
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

// ParseContractRevertError is a log parse operation binding the contract event 0x1a601cf5e0efbd558b2778b7389af04741d1c49bcab104c40daa2da194593617.
//
// Solidity: event ContractRevertError(address theContract, uint256 atBlock, string theMessage)
func (_FtsoManager *FtsoManagerFilterer) ParseContractRevertError(log types.Log) (*FtsoManagerContractRevertError, error) {
	event := new(FtsoManagerContractRevertError)
	if err := _FtsoManager.contract.UnpackLog(event, "ContractRevertError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerDistributingRewardsFailedIterator is returned from FilterDistributingRewardsFailed and is used to iterate over the raw logs and unpacked data for DistributingRewardsFailed events raised by the FtsoManager contract.
type FtsoManagerDistributingRewardsFailedIterator struct {
	Event *FtsoManagerDistributingRewardsFailed // Event containing the contract specifics and raw log

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
func (it *FtsoManagerDistributingRewardsFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerDistributingRewardsFailed)
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
		it.Event = new(FtsoManagerDistributingRewardsFailed)
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
func (it *FtsoManagerDistributingRewardsFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerDistributingRewardsFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerDistributingRewardsFailed represents a DistributingRewardsFailed event raised by the FtsoManager contract.
type FtsoManagerDistributingRewardsFailed struct {
	Ftso    common.Address
	EpochId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDistributingRewardsFailed is a free log retrieval operation binding the contract event 0x175a1d13d190d6a1e14461c214b3ecf6118b828797750b7bffd7c4f2c1eba54c.
//
// Solidity: event DistributingRewardsFailed(address ftso, uint256 epochId)
func (_FtsoManager *FtsoManagerFilterer) FilterDistributingRewardsFailed(opts *bind.FilterOpts) (*FtsoManagerDistributingRewardsFailedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "DistributingRewardsFailed")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerDistributingRewardsFailedIterator{contract: _FtsoManager.contract, event: "DistributingRewardsFailed", logs: logs, sub: sub}, nil
}

// WatchDistributingRewardsFailed is a free log subscription operation binding the contract event 0x175a1d13d190d6a1e14461c214b3ecf6118b828797750b7bffd7c4f2c1eba54c.
//
// Solidity: event DistributingRewardsFailed(address ftso, uint256 epochId)
func (_FtsoManager *FtsoManagerFilterer) WatchDistributingRewardsFailed(opts *bind.WatchOpts, sink chan<- *FtsoManagerDistributingRewardsFailed) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "DistributingRewardsFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerDistributingRewardsFailed)
				if err := _FtsoManager.contract.UnpackLog(event, "DistributingRewardsFailed", log); err != nil {
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

// ParseDistributingRewardsFailed is a log parse operation binding the contract event 0x175a1d13d190d6a1e14461c214b3ecf6118b828797750b7bffd7c4f2c1eba54c.
//
// Solidity: event DistributingRewardsFailed(address ftso, uint256 epochId)
func (_FtsoManager *FtsoManagerFilterer) ParseDistributingRewardsFailed(log types.Log) (*FtsoManagerDistributingRewardsFailed, error) {
	event := new(FtsoManagerDistributingRewardsFailed)
	if err := _FtsoManager.contract.UnpackLog(event, "DistributingRewardsFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerFallbackModeIterator is returned from FilterFallbackMode and is used to iterate over the raw logs and unpacked data for FallbackMode events raised by the FtsoManager contract.
type FtsoManagerFallbackModeIterator struct {
	Event *FtsoManagerFallbackMode // Event containing the contract specifics and raw log

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
func (it *FtsoManagerFallbackModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerFallbackMode)
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
		it.Event = new(FtsoManagerFallbackMode)
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
func (it *FtsoManagerFallbackModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerFallbackModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerFallbackMode represents a FallbackMode event raised by the FtsoManager contract.
type FtsoManagerFallbackMode struct {
	FallbackMode bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFallbackMode is a free log retrieval operation binding the contract event 0x217a37a37fc40a97159886f80c3d45986e6fc4330ce6ad7283478b5e5ab705bc.
//
// Solidity: event FallbackMode(bool fallbackMode)
func (_FtsoManager *FtsoManagerFilterer) FilterFallbackMode(opts *bind.FilterOpts) (*FtsoManagerFallbackModeIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "FallbackMode")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerFallbackModeIterator{contract: _FtsoManager.contract, event: "FallbackMode", logs: logs, sub: sub}, nil
}

// WatchFallbackMode is a free log subscription operation binding the contract event 0x217a37a37fc40a97159886f80c3d45986e6fc4330ce6ad7283478b5e5ab705bc.
//
// Solidity: event FallbackMode(bool fallbackMode)
func (_FtsoManager *FtsoManagerFilterer) WatchFallbackMode(opts *bind.WatchOpts, sink chan<- *FtsoManagerFallbackMode) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "FallbackMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerFallbackMode)
				if err := _FtsoManager.contract.UnpackLog(event, "FallbackMode", log); err != nil {
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

// ParseFallbackMode is a log parse operation binding the contract event 0x217a37a37fc40a97159886f80c3d45986e6fc4330ce6ad7283478b5e5ab705bc.
//
// Solidity: event FallbackMode(bool fallbackMode)
func (_FtsoManager *FtsoManagerFilterer) ParseFallbackMode(log types.Log) (*FtsoManagerFallbackMode, error) {
	event := new(FtsoManagerFallbackMode)
	if err := _FtsoManager.contract.UnpackLog(event, "FallbackMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerFinalizingPriceEpochFailedIterator is returned from FilterFinalizingPriceEpochFailed and is used to iterate over the raw logs and unpacked data for FinalizingPriceEpochFailed events raised by the FtsoManager contract.
type FtsoManagerFinalizingPriceEpochFailedIterator struct {
	Event *FtsoManagerFinalizingPriceEpochFailed // Event containing the contract specifics and raw log

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
func (it *FtsoManagerFinalizingPriceEpochFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerFinalizingPriceEpochFailed)
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
		it.Event = new(FtsoManagerFinalizingPriceEpochFailed)
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
func (it *FtsoManagerFinalizingPriceEpochFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerFinalizingPriceEpochFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerFinalizingPriceEpochFailed represents a FinalizingPriceEpochFailed event raised by the FtsoManager contract.
type FtsoManagerFinalizingPriceEpochFailed struct {
	Ftso        common.Address
	EpochId     *big.Int
	FailingType uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinalizingPriceEpochFailed is a free log retrieval operation binding the contract event 0x79f4c7cc43bfb79f5a3aad0d92f75b6fed7db061bb5cc2580a01c8132711b881.
//
// Solidity: event FinalizingPriceEpochFailed(address ftso, uint256 epochId, uint8 failingType)
func (_FtsoManager *FtsoManagerFilterer) FilterFinalizingPriceEpochFailed(opts *bind.FilterOpts) (*FtsoManagerFinalizingPriceEpochFailedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "FinalizingPriceEpochFailed")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerFinalizingPriceEpochFailedIterator{contract: _FtsoManager.contract, event: "FinalizingPriceEpochFailed", logs: logs, sub: sub}, nil
}

// WatchFinalizingPriceEpochFailed is a free log subscription operation binding the contract event 0x79f4c7cc43bfb79f5a3aad0d92f75b6fed7db061bb5cc2580a01c8132711b881.
//
// Solidity: event FinalizingPriceEpochFailed(address ftso, uint256 epochId, uint8 failingType)
func (_FtsoManager *FtsoManagerFilterer) WatchFinalizingPriceEpochFailed(opts *bind.WatchOpts, sink chan<- *FtsoManagerFinalizingPriceEpochFailed) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "FinalizingPriceEpochFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerFinalizingPriceEpochFailed)
				if err := _FtsoManager.contract.UnpackLog(event, "FinalizingPriceEpochFailed", log); err != nil {
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

// ParseFinalizingPriceEpochFailed is a log parse operation binding the contract event 0x79f4c7cc43bfb79f5a3aad0d92f75b6fed7db061bb5cc2580a01c8132711b881.
//
// Solidity: event FinalizingPriceEpochFailed(address ftso, uint256 epochId, uint8 failingType)
func (_FtsoManager *FtsoManagerFilterer) ParseFinalizingPriceEpochFailed(log types.Log) (*FtsoManagerFinalizingPriceEpochFailed, error) {
	event := new(FtsoManagerFinalizingPriceEpochFailed)
	if err := _FtsoManager.contract.UnpackLog(event, "FinalizingPriceEpochFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerFtsoAddedIterator is returned from FilterFtsoAdded and is used to iterate over the raw logs and unpacked data for FtsoAdded events raised by the FtsoManager contract.
type FtsoManagerFtsoAddedIterator struct {
	Event *FtsoManagerFtsoAdded // Event containing the contract specifics and raw log

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
func (it *FtsoManagerFtsoAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerFtsoAdded)
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
		it.Event = new(FtsoManagerFtsoAdded)
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
func (it *FtsoManagerFtsoAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerFtsoAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerFtsoAdded represents a FtsoAdded event raised by the FtsoManager contract.
type FtsoManagerFtsoAdded struct {
	Ftso common.Address
	Add  bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFtsoAdded is a free log retrieval operation binding the contract event 0xa0985424f2efdcae4b57a7c84bbf0a0b19f93054f21e9eb1cfcd5a59813fe1da.
//
// Solidity: event FtsoAdded(address ftso, bool add)
func (_FtsoManager *FtsoManagerFilterer) FilterFtsoAdded(opts *bind.FilterOpts) (*FtsoManagerFtsoAddedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "FtsoAdded")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerFtsoAddedIterator{contract: _FtsoManager.contract, event: "FtsoAdded", logs: logs, sub: sub}, nil
}

// WatchFtsoAdded is a free log subscription operation binding the contract event 0xa0985424f2efdcae4b57a7c84bbf0a0b19f93054f21e9eb1cfcd5a59813fe1da.
//
// Solidity: event FtsoAdded(address ftso, bool add)
func (_FtsoManager *FtsoManagerFilterer) WatchFtsoAdded(opts *bind.WatchOpts, sink chan<- *FtsoManagerFtsoAdded) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "FtsoAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerFtsoAdded)
				if err := _FtsoManager.contract.UnpackLog(event, "FtsoAdded", log); err != nil {
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

// ParseFtsoAdded is a log parse operation binding the contract event 0xa0985424f2efdcae4b57a7c84bbf0a0b19f93054f21e9eb1cfcd5a59813fe1da.
//
// Solidity: event FtsoAdded(address ftso, bool add)
func (_FtsoManager *FtsoManagerFilterer) ParseFtsoAdded(log types.Log) (*FtsoManagerFtsoAdded, error) {
	event := new(FtsoManagerFtsoAdded)
	if err := _FtsoManager.contract.UnpackLog(event, "FtsoAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerFtsoDeactivationFailedIterator is returned from FilterFtsoDeactivationFailed and is used to iterate over the raw logs and unpacked data for FtsoDeactivationFailed events raised by the FtsoManager contract.
type FtsoManagerFtsoDeactivationFailedIterator struct {
	Event *FtsoManagerFtsoDeactivationFailed // Event containing the contract specifics and raw log

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
func (it *FtsoManagerFtsoDeactivationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerFtsoDeactivationFailed)
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
		it.Event = new(FtsoManagerFtsoDeactivationFailed)
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
func (it *FtsoManagerFtsoDeactivationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerFtsoDeactivationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerFtsoDeactivationFailed represents a FtsoDeactivationFailed event raised by the FtsoManager contract.
type FtsoManagerFtsoDeactivationFailed struct {
	Ftso common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFtsoDeactivationFailed is a free log retrieval operation binding the contract event 0xa2109c8ecab2fa994cb3c911158dfa5c776e8f9052b49782465fc324d5b5c5f9.
//
// Solidity: event FtsoDeactivationFailed(address ftso)
func (_FtsoManager *FtsoManagerFilterer) FilterFtsoDeactivationFailed(opts *bind.FilterOpts) (*FtsoManagerFtsoDeactivationFailedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "FtsoDeactivationFailed")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerFtsoDeactivationFailedIterator{contract: _FtsoManager.contract, event: "FtsoDeactivationFailed", logs: logs, sub: sub}, nil
}

// WatchFtsoDeactivationFailed is a free log subscription operation binding the contract event 0xa2109c8ecab2fa994cb3c911158dfa5c776e8f9052b49782465fc324d5b5c5f9.
//
// Solidity: event FtsoDeactivationFailed(address ftso)
func (_FtsoManager *FtsoManagerFilterer) WatchFtsoDeactivationFailed(opts *bind.WatchOpts, sink chan<- *FtsoManagerFtsoDeactivationFailed) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "FtsoDeactivationFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerFtsoDeactivationFailed)
				if err := _FtsoManager.contract.UnpackLog(event, "FtsoDeactivationFailed", log); err != nil {
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

// ParseFtsoDeactivationFailed is a log parse operation binding the contract event 0xa2109c8ecab2fa994cb3c911158dfa5c776e8f9052b49782465fc324d5b5c5f9.
//
// Solidity: event FtsoDeactivationFailed(address ftso)
func (_FtsoManager *FtsoManagerFilterer) ParseFtsoDeactivationFailed(log types.Log) (*FtsoManagerFtsoDeactivationFailed, error) {
	event := new(FtsoManagerFtsoDeactivationFailed)
	if err := _FtsoManager.contract.UnpackLog(event, "FtsoDeactivationFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerFtsoFallbackModeIterator is returned from FilterFtsoFallbackMode and is used to iterate over the raw logs and unpacked data for FtsoFallbackMode events raised by the FtsoManager contract.
type FtsoManagerFtsoFallbackModeIterator struct {
	Event *FtsoManagerFtsoFallbackMode // Event containing the contract specifics and raw log

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
func (it *FtsoManagerFtsoFallbackModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerFtsoFallbackMode)
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
		it.Event = new(FtsoManagerFtsoFallbackMode)
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
func (it *FtsoManagerFtsoFallbackModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerFtsoFallbackModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerFtsoFallbackMode represents a FtsoFallbackMode event raised by the FtsoManager contract.
type FtsoManagerFtsoFallbackMode struct {
	Ftso         common.Address
	FallbackMode bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFtsoFallbackMode is a free log retrieval operation binding the contract event 0x24462ede4d3e8e5a69fecec6290d42a311016ca752216d9a3d681e284791b7ac.
//
// Solidity: event FtsoFallbackMode(address ftso, bool fallbackMode)
func (_FtsoManager *FtsoManagerFilterer) FilterFtsoFallbackMode(opts *bind.FilterOpts) (*FtsoManagerFtsoFallbackModeIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "FtsoFallbackMode")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerFtsoFallbackModeIterator{contract: _FtsoManager.contract, event: "FtsoFallbackMode", logs: logs, sub: sub}, nil
}

// WatchFtsoFallbackMode is a free log subscription operation binding the contract event 0x24462ede4d3e8e5a69fecec6290d42a311016ca752216d9a3d681e284791b7ac.
//
// Solidity: event FtsoFallbackMode(address ftso, bool fallbackMode)
func (_FtsoManager *FtsoManagerFilterer) WatchFtsoFallbackMode(opts *bind.WatchOpts, sink chan<- *FtsoManagerFtsoFallbackMode) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "FtsoFallbackMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerFtsoFallbackMode)
				if err := _FtsoManager.contract.UnpackLog(event, "FtsoFallbackMode", log); err != nil {
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

// ParseFtsoFallbackMode is a log parse operation binding the contract event 0x24462ede4d3e8e5a69fecec6290d42a311016ca752216d9a3d681e284791b7ac.
//
// Solidity: event FtsoFallbackMode(address ftso, bool fallbackMode)
func (_FtsoManager *FtsoManagerFilterer) ParseFtsoFallbackMode(log types.Log) (*FtsoManagerFtsoFallbackMode, error) {
	event := new(FtsoManagerFtsoFallbackMode)
	if err := _FtsoManager.contract.UnpackLog(event, "FtsoFallbackMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the FtsoManager contract.
type FtsoManagerGovernanceCallTimelockedIterator struct {
	Event *FtsoManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *FtsoManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerGovernanceCallTimelocked)
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
		it.Event = new(FtsoManagerGovernanceCallTimelocked)
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
func (it *FtsoManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FtsoManager contract.
type FtsoManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtsoManager *FtsoManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*FtsoManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerGovernanceCallTimelockedIterator{contract: _FtsoManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtsoManager *FtsoManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *FtsoManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerGovernanceCallTimelocked)
				if err := _FtsoManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_FtsoManager *FtsoManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*FtsoManagerGovernanceCallTimelocked, error) {
	event := new(FtsoManagerGovernanceCallTimelocked)
	if err := _FtsoManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FtsoManager contract.
type FtsoManagerGovernanceInitialisedIterator struct {
	Event *FtsoManagerGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *FtsoManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerGovernanceInitialised)
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
		it.Event = new(FtsoManagerGovernanceInitialised)
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
func (it *FtsoManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerGovernanceInitialised represents a GovernanceInitialised event raised by the FtsoManager contract.
type FtsoManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FtsoManager *FtsoManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FtsoManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerGovernanceInitialisedIterator{contract: _FtsoManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FtsoManager *FtsoManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FtsoManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerGovernanceInitialised)
				if err := _FtsoManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_FtsoManager *FtsoManagerFilterer) ParseGovernanceInitialised(log types.Log) (*FtsoManagerGovernanceInitialised, error) {
	event := new(FtsoManagerGovernanceInitialised)
	if err := _FtsoManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the FtsoManager contract.
type FtsoManagerGovernedProductionModeEnteredIterator struct {
	Event *FtsoManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *FtsoManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerGovernedProductionModeEntered)
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
		it.Event = new(FtsoManagerGovernedProductionModeEntered)
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
func (it *FtsoManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FtsoManager contract.
type FtsoManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FtsoManager *FtsoManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*FtsoManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerGovernedProductionModeEnteredIterator{contract: _FtsoManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FtsoManager *FtsoManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *FtsoManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerGovernedProductionModeEntered)
				if err := _FtsoManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_FtsoManager *FtsoManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*FtsoManagerGovernedProductionModeEntered, error) {
	event := new(FtsoManagerGovernedProductionModeEntered)
	if err := _FtsoManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerInitializingCurrentEpochStateForRevealFailedIterator is returned from FilterInitializingCurrentEpochStateForRevealFailed and is used to iterate over the raw logs and unpacked data for InitializingCurrentEpochStateForRevealFailed events raised by the FtsoManager contract.
type FtsoManagerInitializingCurrentEpochStateForRevealFailedIterator struct {
	Event *FtsoManagerInitializingCurrentEpochStateForRevealFailed // Event containing the contract specifics and raw log

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
func (it *FtsoManagerInitializingCurrentEpochStateForRevealFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerInitializingCurrentEpochStateForRevealFailed)
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
		it.Event = new(FtsoManagerInitializingCurrentEpochStateForRevealFailed)
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
func (it *FtsoManagerInitializingCurrentEpochStateForRevealFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerInitializingCurrentEpochStateForRevealFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerInitializingCurrentEpochStateForRevealFailed represents a InitializingCurrentEpochStateForRevealFailed event raised by the FtsoManager contract.
type FtsoManagerInitializingCurrentEpochStateForRevealFailed struct {
	Ftso    common.Address
	EpochId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitializingCurrentEpochStateForRevealFailed is a free log retrieval operation binding the contract event 0x61156899176547b8075bfa81fa2996c6057ce9c71320884b11c0179d9dc2e462.
//
// Solidity: event InitializingCurrentEpochStateForRevealFailed(address ftso, uint256 epochId)
func (_FtsoManager *FtsoManagerFilterer) FilterInitializingCurrentEpochStateForRevealFailed(opts *bind.FilterOpts) (*FtsoManagerInitializingCurrentEpochStateForRevealFailedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "InitializingCurrentEpochStateForRevealFailed")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerInitializingCurrentEpochStateForRevealFailedIterator{contract: _FtsoManager.contract, event: "InitializingCurrentEpochStateForRevealFailed", logs: logs, sub: sub}, nil
}

// WatchInitializingCurrentEpochStateForRevealFailed is a free log subscription operation binding the contract event 0x61156899176547b8075bfa81fa2996c6057ce9c71320884b11c0179d9dc2e462.
//
// Solidity: event InitializingCurrentEpochStateForRevealFailed(address ftso, uint256 epochId)
func (_FtsoManager *FtsoManagerFilterer) WatchInitializingCurrentEpochStateForRevealFailed(opts *bind.WatchOpts, sink chan<- *FtsoManagerInitializingCurrentEpochStateForRevealFailed) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "InitializingCurrentEpochStateForRevealFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerInitializingCurrentEpochStateForRevealFailed)
				if err := _FtsoManager.contract.UnpackLog(event, "InitializingCurrentEpochStateForRevealFailed", log); err != nil {
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

// ParseInitializingCurrentEpochStateForRevealFailed is a log parse operation binding the contract event 0x61156899176547b8075bfa81fa2996c6057ce9c71320884b11c0179d9dc2e462.
//
// Solidity: event InitializingCurrentEpochStateForRevealFailed(address ftso, uint256 epochId)
func (_FtsoManager *FtsoManagerFilterer) ParseInitializingCurrentEpochStateForRevealFailed(log types.Log) (*FtsoManagerInitializingCurrentEpochStateForRevealFailed, error) {
	event := new(FtsoManagerInitializingCurrentEpochStateForRevealFailed)
	if err := _FtsoManager.contract.UnpackLog(event, "InitializingCurrentEpochStateForRevealFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerPriceEpochFinalizedIterator is returned from FilterPriceEpochFinalized and is used to iterate over the raw logs and unpacked data for PriceEpochFinalized events raised by the FtsoManager contract.
type FtsoManagerPriceEpochFinalizedIterator struct {
	Event *FtsoManagerPriceEpochFinalized // Event containing the contract specifics and raw log

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
func (it *FtsoManagerPriceEpochFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerPriceEpochFinalized)
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
		it.Event = new(FtsoManagerPriceEpochFinalized)
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
func (it *FtsoManagerPriceEpochFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerPriceEpochFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerPriceEpochFinalized represents a PriceEpochFinalized event raised by the FtsoManager contract.
type FtsoManagerPriceEpochFinalized struct {
	ChosenFtso    common.Address
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPriceEpochFinalized is a free log retrieval operation binding the contract event 0x98b050a4042fbd1b89934ef40b9342e593f15081a348af940573a0179031f4ad.
//
// Solidity: event PriceEpochFinalized(address chosenFtso, uint256 rewardEpochId)
func (_FtsoManager *FtsoManagerFilterer) FilterPriceEpochFinalized(opts *bind.FilterOpts) (*FtsoManagerPriceEpochFinalizedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "PriceEpochFinalized")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerPriceEpochFinalizedIterator{contract: _FtsoManager.contract, event: "PriceEpochFinalized", logs: logs, sub: sub}, nil
}

// WatchPriceEpochFinalized is a free log subscription operation binding the contract event 0x98b050a4042fbd1b89934ef40b9342e593f15081a348af940573a0179031f4ad.
//
// Solidity: event PriceEpochFinalized(address chosenFtso, uint256 rewardEpochId)
func (_FtsoManager *FtsoManagerFilterer) WatchPriceEpochFinalized(opts *bind.WatchOpts, sink chan<- *FtsoManagerPriceEpochFinalized) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "PriceEpochFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerPriceEpochFinalized)
				if err := _FtsoManager.contract.UnpackLog(event, "PriceEpochFinalized", log); err != nil {
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

// ParsePriceEpochFinalized is a log parse operation binding the contract event 0x98b050a4042fbd1b89934ef40b9342e593f15081a348af940573a0179031f4ad.
//
// Solidity: event PriceEpochFinalized(address chosenFtso, uint256 rewardEpochId)
func (_FtsoManager *FtsoManagerFilterer) ParsePriceEpochFinalized(log types.Log) (*FtsoManagerPriceEpochFinalized, error) {
	event := new(FtsoManagerPriceEpochFinalized)
	if err := _FtsoManager.contract.UnpackLog(event, "PriceEpochFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerRewardEpochFinalizedIterator is returned from FilterRewardEpochFinalized and is used to iterate over the raw logs and unpacked data for RewardEpochFinalized events raised by the FtsoManager contract.
type FtsoManagerRewardEpochFinalizedIterator struct {
	Event *FtsoManagerRewardEpochFinalized // Event containing the contract specifics and raw log

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
func (it *FtsoManagerRewardEpochFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerRewardEpochFinalized)
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
		it.Event = new(FtsoManagerRewardEpochFinalized)
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
func (it *FtsoManagerRewardEpochFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerRewardEpochFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerRewardEpochFinalized represents a RewardEpochFinalized event raised by the FtsoManager contract.
type FtsoManagerRewardEpochFinalized struct {
	VotepowerBlock *big.Int
	StartBlock     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardEpochFinalized is a free log retrieval operation binding the contract event 0x1813f880dc24666c8b69c9d771a487ea620a27fde1514be3112847056c0c5322.
//
// Solidity: event RewardEpochFinalized(uint256 votepowerBlock, uint256 startBlock)
func (_FtsoManager *FtsoManagerFilterer) FilterRewardEpochFinalized(opts *bind.FilterOpts) (*FtsoManagerRewardEpochFinalizedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "RewardEpochFinalized")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerRewardEpochFinalizedIterator{contract: _FtsoManager.contract, event: "RewardEpochFinalized", logs: logs, sub: sub}, nil
}

// WatchRewardEpochFinalized is a free log subscription operation binding the contract event 0x1813f880dc24666c8b69c9d771a487ea620a27fde1514be3112847056c0c5322.
//
// Solidity: event RewardEpochFinalized(uint256 votepowerBlock, uint256 startBlock)
func (_FtsoManager *FtsoManagerFilterer) WatchRewardEpochFinalized(opts *bind.WatchOpts, sink chan<- *FtsoManagerRewardEpochFinalized) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "RewardEpochFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerRewardEpochFinalized)
				if err := _FtsoManager.contract.UnpackLog(event, "RewardEpochFinalized", log); err != nil {
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

// ParseRewardEpochFinalized is a log parse operation binding the contract event 0x1813f880dc24666c8b69c9d771a487ea620a27fde1514be3112847056c0c5322.
//
// Solidity: event RewardEpochFinalized(uint256 votepowerBlock, uint256 startBlock)
func (_FtsoManager *FtsoManagerFilterer) ParseRewardEpochFinalized(log types.Log) (*FtsoManagerRewardEpochFinalized, error) {
	event := new(FtsoManagerRewardEpochFinalized)
	if err := _FtsoManager.contract.UnpackLog(event, "RewardEpochFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the FtsoManager contract.
type FtsoManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *FtsoManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *FtsoManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerTimelockedGovernanceCallCanceled)
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
		it.Event = new(FtsoManagerTimelockedGovernanceCallCanceled)
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
func (it *FtsoManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FtsoManager contract.
type FtsoManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FtsoManager *FtsoManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*FtsoManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerTimelockedGovernanceCallCanceledIterator{contract: _FtsoManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FtsoManager *FtsoManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *FtsoManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerTimelockedGovernanceCallCanceled)
				if err := _FtsoManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_FtsoManager *FtsoManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*FtsoManagerTimelockedGovernanceCallCanceled, error) {
	event := new(FtsoManagerTimelockedGovernanceCallCanceled)
	if err := _FtsoManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the FtsoManager contract.
type FtsoManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *FtsoManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *FtsoManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerTimelockedGovernanceCallExecuted)
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
		it.Event = new(FtsoManagerTimelockedGovernanceCallExecuted)
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
func (it *FtsoManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FtsoManager contract.
type FtsoManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FtsoManager *FtsoManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*FtsoManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerTimelockedGovernanceCallExecutedIterator{contract: _FtsoManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FtsoManager *FtsoManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *FtsoManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerTimelockedGovernanceCallExecuted)
				if err := _FtsoManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_FtsoManager *FtsoManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*FtsoManagerTimelockedGovernanceCallExecuted, error) {
	event := new(FtsoManagerTimelockedGovernanceCallExecuted)
	if err := _FtsoManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoManagerUpdatingActiveValidatorsTriggerFailedIterator is returned from FilterUpdatingActiveValidatorsTriggerFailed and is used to iterate over the raw logs and unpacked data for UpdatingActiveValidatorsTriggerFailed events raised by the FtsoManager contract.
type FtsoManagerUpdatingActiveValidatorsTriggerFailedIterator struct {
	Event *FtsoManagerUpdatingActiveValidatorsTriggerFailed // Event containing the contract specifics and raw log

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
func (it *FtsoManagerUpdatingActiveValidatorsTriggerFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoManagerUpdatingActiveValidatorsTriggerFailed)
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
		it.Event = new(FtsoManagerUpdatingActiveValidatorsTriggerFailed)
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
func (it *FtsoManagerUpdatingActiveValidatorsTriggerFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoManagerUpdatingActiveValidatorsTriggerFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoManagerUpdatingActiveValidatorsTriggerFailed represents a UpdatingActiveValidatorsTriggerFailed event raised by the FtsoManager contract.
type FtsoManagerUpdatingActiveValidatorsTriggerFailed struct {
	RewardEpoch *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatingActiveValidatorsTriggerFailed is a free log retrieval operation binding the contract event 0xf7c7d6681321cc290eb89e8c96dba504436073b8bb277945cc32177e5181dd84.
//
// Solidity: event UpdatingActiveValidatorsTriggerFailed(uint256 rewardEpoch)
func (_FtsoManager *FtsoManagerFilterer) FilterUpdatingActiveValidatorsTriggerFailed(opts *bind.FilterOpts) (*FtsoManagerUpdatingActiveValidatorsTriggerFailedIterator, error) {

	logs, sub, err := _FtsoManager.contract.FilterLogs(opts, "UpdatingActiveValidatorsTriggerFailed")
	if err != nil {
		return nil, err
	}
	return &FtsoManagerUpdatingActiveValidatorsTriggerFailedIterator{contract: _FtsoManager.contract, event: "UpdatingActiveValidatorsTriggerFailed", logs: logs, sub: sub}, nil
}

// WatchUpdatingActiveValidatorsTriggerFailed is a free log subscription operation binding the contract event 0xf7c7d6681321cc290eb89e8c96dba504436073b8bb277945cc32177e5181dd84.
//
// Solidity: event UpdatingActiveValidatorsTriggerFailed(uint256 rewardEpoch)
func (_FtsoManager *FtsoManagerFilterer) WatchUpdatingActiveValidatorsTriggerFailed(opts *bind.WatchOpts, sink chan<- *FtsoManagerUpdatingActiveValidatorsTriggerFailed) (event.Subscription, error) {

	logs, sub, err := _FtsoManager.contract.WatchLogs(opts, "UpdatingActiveValidatorsTriggerFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoManagerUpdatingActiveValidatorsTriggerFailed)
				if err := _FtsoManager.contract.UnpackLog(event, "UpdatingActiveValidatorsTriggerFailed", log); err != nil {
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

// ParseUpdatingActiveValidatorsTriggerFailed is a log parse operation binding the contract event 0xf7c7d6681321cc290eb89e8c96dba504436073b8bb277945cc32177e5181dd84.
//
// Solidity: event UpdatingActiveValidatorsTriggerFailed(uint256 rewardEpoch)
func (_FtsoManager *FtsoManagerFilterer) ParseUpdatingActiveValidatorsTriggerFailed(log types.Log) (*FtsoManagerUpdatingActiveValidatorsTriggerFailed, error) {
	event := new(FtsoManagerUpdatingActiveValidatorsTriggerFailed)
	if err := _FtsoManager.contract.UnpackLog(event, "UpdatingActiveValidatorsTriggerFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
