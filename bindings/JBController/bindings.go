// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package JBController

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

// JBCurrencyAmount is an auto generated low-level Go binding around an user-defined struct.
type JBCurrencyAmount struct {
	Amount   *big.Int
	Currency *big.Int
}

// JBFundAccessLimitGroup is an auto generated low-level Go binding around an user-defined struct.
type JBFundAccessLimitGroup struct {
	Terminal          common.Address
	Token             common.Address
	PayoutLimits      []JBCurrencyAmount
	SurplusAllowances []JBCurrencyAmount
}

// JBRuleset is an auto generated low-level Go binding around an user-defined struct.
type JBRuleset struct {
	CycleNumber  *big.Int
	Id           *big.Int
	BasedOnId    *big.Int
	Start        *big.Int
	Duration     *big.Int
	Weight       *big.Int
	DecayRate    *big.Int
	ApprovalHook common.Address
	Metadata     *big.Int
}

// JBRulesetConfig is an auto generated low-level Go binding around an user-defined struct.
type JBRulesetConfig struct {
	MustStartAtOrAfter    *big.Int
	Duration              *big.Int
	Weight                *big.Int
	DecayRate             *big.Int
	ApprovalHook          common.Address
	Metadata              JBRulesetMetadata
	SplitGroups           []JBSplitGroup
	FundAccessLimitGroups []JBFundAccessLimitGroup
}

// JBRulesetMetadata is an auto generated low-level Go binding around an user-defined struct.
type JBRulesetMetadata struct {
	ReservedRate                  *big.Int
	RedemptionRate                *big.Int
	BaseCurrency                  *big.Int
	PausePay                      bool
	PauseCreditTransfers          bool
	AllowOwnerMinting             bool
	AllowTerminalMigration        bool
	AllowSetTerminals             bool
	AllowControllerMigration      bool
	AllowSetController            bool
	HoldFees                      bool
	UseTotalSurplusForRedemptions bool
	UseDataHookForPay             bool
	UseDataHookForRedeem          bool
	DataHook                      common.Address
	Metadata                      *big.Int
}

// JBRulesetWithMetadata is an auto generated low-level Go binding around an user-defined struct.
type JBRulesetWithMetadata struct {
	Ruleset  JBRuleset
	Metadata JBRulesetMetadata
}

// JBSplit is an auto generated low-level Go binding around an user-defined struct.
type JBSplit struct {
	PreferAddToBalance bool
	Percent            *big.Int
	ProjectId          *big.Int
	Beneficiary        common.Address
	LockedUntil        *big.Int
	Hook               common.Address
}

// JBSplitGroup is an auto generated low-level Go binding around an user-defined struct.
type JBSplitGroup struct {
	GroupId *big.Int
	Splits  []JBSplit
}

// JBTerminalConfig is an auto generated low-level Go binding around an user-defined struct.
type JBTerminalConfig struct {
	Terminal       common.Address
	TokensToAccept []common.Address
}

// JBControllerMetaData contains all meta data concerning the JBController contract.
var JBControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"permissions\",\"type\":\"address\",\"internalType\":\"contractIJBPermissions\"},{\"name\":\"projects\",\"type\":\"address\",\"internalType\":\"contractIJBProjects\"},{\"name\":\"directory\",\"type\":\"address\",\"internalType\":\"contractIJBDirectory\"},{\"name\":\"rulesets\",\"type\":\"address\",\"internalType\":\"contractIJBRulesets\"},{\"name\":\"tokens\",\"type\":\"address\",\"internalType\":\"contractIJBTokens\"},{\"name\":\"splits\",\"type\":\"address\",\"internalType\":\"contractIJBSplits\"},{\"name\":\"fundAccessLimits\",\"type\":\"address\",\"internalType\":\"contractIJBFundAccessLimits\"},{\"name\":\"trustedForwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DIRECTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FUND_ACCESS_LIMITS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBFundAccessLimits\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PERMISSIONS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBPermissions\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROJECTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBProjects\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RULESETS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBRulesets\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SPLITS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBSplits\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKENS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBTokens\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnTokensOf\",\"inputs\":[{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimTokensFor\",\"inputs\":[{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentRulesetOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ruleset\",\"type\":\"tuple\",\"internalType\":\"structJBRuleset\",\"components\":[{\"name\":\"cycleNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"basedOnId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decayRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approvalHook\",\"type\":\"address\",\"internalType\":\"contractIJBRulesetApprovalHook\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structJBRulesetMetadata\",\"components\":[{\"name\":\"reservedRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redemptionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pausePay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"pauseCreditTransfers\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowOwnerMinting\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowTerminalMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetTerminals\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowControllerMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetController\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"holdFees\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useTotalSurplusForRedemptions\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForPay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForRedeem\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHook\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployERC20For\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIJBToken\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRulesetOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rulesetId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ruleset\",\"type\":\"tuple\",\"internalType\":\"structJBRuleset\",\"components\":[{\"name\":\"cycleNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"basedOnId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decayRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approvalHook\",\"type\":\"address\",\"internalType\":\"contractIJBRulesetApprovalHook\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structJBRulesetMetadata\",\"components\":[{\"name\":\"reservedRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redemptionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pausePay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"pauseCreditTransfers\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowOwnerMinting\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowTerminalMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetTerminals\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowControllerMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetController\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"holdFees\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useTotalSurplusForRedemptions\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForPay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForRedeem\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHook\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestQueuedRulesetOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ruleset\",\"type\":\"tuple\",\"internalType\":\"structJBRuleset\",\"components\":[{\"name\":\"cycleNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"basedOnId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decayRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approvalHook\",\"type\":\"address\",\"internalType\":\"contractIJBRulesetApprovalHook\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structJBRulesetMetadata\",\"components\":[{\"name\":\"reservedRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redemptionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pausePay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"pauseCreditTransfers\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowOwnerMinting\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowTerminalMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetTerminals\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowControllerMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetController\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"holdFees\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useTotalSurplusForRedemptions\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForPay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForRedeem\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHook\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approvalStatus\",\"type\":\"uint8\",\"internalType\":\"enumJBApprovalStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"launchProjectFor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rulesetConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structJBRulesetConfig[]\",\"components\":[{\"name\":\"mustStartAtOrAfter\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decayRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approvalHook\",\"type\":\"address\",\"internalType\":\"contractIJBRulesetApprovalHook\"},{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structJBRulesetMetadata\",\"components\":[{\"name\":\"reservedRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redemptionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pausePay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"pauseCreditTransfers\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowOwnerMinting\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowTerminalMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetTerminals\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowControllerMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetController\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"holdFees\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useTotalSurplusForRedemptions\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForPay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForRedeem\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHook\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"splitGroups\",\"type\":\"tuple[]\",\"internalType\":\"structJBSplitGroup[]\",\"components\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"splits\",\"type\":\"tuple[]\",\"internalType\":\"structJBSplit[]\",\"components\":[{\"name\":\"preferAddToBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"percent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"contractIJBSplitHook\"}]}]},{\"name\":\"fundAccessLimitGroups\",\"type\":\"tuple[]\",\"internalType\":\"structJBFundAccessLimitGroup[]\",\"components\":[{\"name\":\"terminal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payoutLimits\",\"type\":\"tuple[]\",\"internalType\":\"structJBCurrencyAmount[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"surplusAllowances\",\"type\":\"tuple[]\",\"internalType\":\"structJBCurrencyAmount[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]},{\"name\":\"terminalConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structJBTerminalConfig[]\",\"components\":[{\"name\":\"terminal\",\"type\":\"address\",\"internalType\":\"contractIJBTerminal\"},{\"name\":\"tokensToAccept\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"launchRulesetsFor\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rulesetConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structJBRulesetConfig[]\",\"components\":[{\"name\":\"mustStartAtOrAfter\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decayRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approvalHook\",\"type\":\"address\",\"internalType\":\"contractIJBRulesetApprovalHook\"},{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structJBRulesetMetadata\",\"components\":[{\"name\":\"reservedRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redemptionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pausePay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"pauseCreditTransfers\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowOwnerMinting\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowTerminalMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetTerminals\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowControllerMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetController\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"holdFees\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useTotalSurplusForRedemptions\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForPay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForRedeem\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHook\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"splitGroups\",\"type\":\"tuple[]\",\"internalType\":\"structJBSplitGroup[]\",\"components\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"splits\",\"type\":\"tuple[]\",\"internalType\":\"structJBSplit[]\",\"components\":[{\"name\":\"preferAddToBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"percent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"contractIJBSplitHook\"}]}]},{\"name\":\"fundAccessLimitGroups\",\"type\":\"tuple[]\",\"internalType\":\"structJBFundAccessLimitGroup[]\",\"components\":[{\"name\":\"terminal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payoutLimits\",\"type\":\"tuple[]\",\"internalType\":\"structJBCurrencyAmount[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"surplusAllowances\",\"type\":\"tuple[]\",\"internalType\":\"structJBCurrencyAmount[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]},{\"name\":\"terminalConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structJBTerminalConfig[]\",\"components\":[{\"name\":\"terminal\",\"type\":\"address\",\"internalType\":\"contractIJBTerminal\"},{\"name\":\"tokensToAccept\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"rulesetId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"migrateController\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"contractIJBMigratable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintTokensOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"useReservedRate\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"beneficiaryTokenCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payReservedTokenToTerminal\",\"inputs\":[{\"name\":\"terminal\",\"type\":\"address\",\"internalType\":\"contractIJBTerminal\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIJBToken\"},{\"name\":\"splitAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pendingReservedTokenBalanceOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueRulesetsOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rulesetConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structJBRulesetConfig[]\",\"components\":[{\"name\":\"mustStartAtOrAfter\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decayRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approvalHook\",\"type\":\"address\",\"internalType\":\"contractIJBRulesetApprovalHook\"},{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structJBRulesetMetadata\",\"components\":[{\"name\":\"reservedRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redemptionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pausePay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"pauseCreditTransfers\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowOwnerMinting\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowTerminalMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetTerminals\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowControllerMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetController\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"holdFees\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useTotalSurplusForRedemptions\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForPay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForRedeem\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHook\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"splitGroups\",\"type\":\"tuple[]\",\"internalType\":\"structJBSplitGroup[]\",\"components\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"splits\",\"type\":\"tuple[]\",\"internalType\":\"structJBSplit[]\",\"components\":[{\"name\":\"preferAddToBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"percent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"contractIJBSplitHook\"}]}]},{\"name\":\"fundAccessLimitGroups\",\"type\":\"tuple[]\",\"internalType\":\"structJBFundAccessLimitGroup[]\",\"components\":[{\"name\":\"terminal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payoutLimits\",\"type\":\"tuple[]\",\"internalType\":\"structJBCurrencyAmount[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"surplusAllowances\",\"type\":\"tuple[]\",\"internalType\":\"structJBCurrencyAmount[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"rulesetId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveMigrationFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"contractIERC165\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rulesetsOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startingId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"rulesets\",\"type\":\"tuple[]\",\"internalType\":\"structJBRulesetWithMetadata[]\",\"components\":[{\"name\":\"ruleset\",\"type\":\"tuple\",\"internalType\":\"structJBRuleset\",\"components\":[{\"name\":\"cycleNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"basedOnId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decayRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approvalHook\",\"type\":\"address\",\"internalType\":\"contractIJBRulesetApprovalHook\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structJBRulesetMetadata\",\"components\":[{\"name\":\"reservedRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redemptionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pausePay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"pauseCreditTransfers\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowOwnerMinting\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowTerminalMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetTerminals\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowControllerMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetController\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"holdFees\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useTotalSurplusForRedemptions\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForPay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForRedeem\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHook\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendReservedTokensToSplitsOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setControllerAllowed\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setSplitGroupsOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rulesetId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"splitGroups\",\"type\":\"tuple[]\",\"internalType\":\"structJBSplitGroup[]\",\"components\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"splits\",\"type\":\"tuple[]\",\"internalType\":\"structJBSplit[]\",\"components\":[{\"name\":\"preferAddToBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"percent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"contractIJBSplitHook\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTerminalsAllowed\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setTokenFor\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIJBToken\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUriOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalTokenSupplyWithReservedTokensOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferCreditsFrom\",\"inputs\":[{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trustedForwarder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upcomingRulesetOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ruleset\",\"type\":\"tuple\",\"internalType\":\"structJBRuleset\",\"components\":[{\"name\":\"cycleNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"basedOnId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decayRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approvalHook\",\"type\":\"address\",\"internalType\":\"contractIJBRulesetApprovalHook\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structJBRulesetMetadata\",\"components\":[{\"name\":\"reservedRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redemptionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pausePay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"pauseCreditTransfers\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowOwnerMinting\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowTerminalMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetTerminals\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowControllerMigration\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowSetController\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"holdFees\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useTotalSurplusForRedemptions\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForPay\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useDataHookForRedeem\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dataHook\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uriOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BurnTokens\",\"inputs\":[{\"name\":\"holder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LaunchProject\",\"inputs\":[{\"name\":\"rulesetId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"memo\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LaunchRulesets\",\"inputs\":[{\"name\":\"rulesetId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MigrateController\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIJBMigratable\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintTokens\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"beneficiaryTokenCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"reservedRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PrepMigration\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QueueRulesets\",\"inputs\":[{\"name\":\"rulesetId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReservedDistributionReverted\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"split\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structJBSplit\",\"components\":[{\"name\":\"preferAddToBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"percent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"contractIJBSplitHook\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendReservedTokensToSplit\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rulesetId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"split\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structJBSplit\",\"components\":[{\"name\":\"preferAddToBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"percent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"contractIJBSplitHook\"}]},{\"name\":\"tokenCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendReservedTokensToSplits\",\"inputs\":[{\"name\":\"rulesetId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rulesetCycleNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"beneficiaryTokenCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMetadata\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CONTROLLER_MIGRATION_NOT_ALLOWED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CREDIT_TRANSFERS_PAUSED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INVALID_BASE_CURRENCY\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INVALID_REDEMPTION_RATE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INVALID_RESERVED_RATE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MINT_NOT_ALLOWED_AND_NOT_TERMINAL_OR_HOOK\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NO_BURNABLE_TOKENS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NO_RESERVED_TOKENS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PRBMath_MulDiv_Overflow\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RULESETS_ALREADY_LAUNCHED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RULESETS_ARRAY_EMPTY\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UNAUTHORIZED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZERO_TOKENS_TO_MINT\",\"inputs\":[]}]",
}

// JBControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use JBControllerMetaData.ABI instead.
var JBControllerABI = JBControllerMetaData.ABI

// JBController is an auto generated Go binding around an Ethereum contract.
type JBController struct {
	JBControllerCaller     // Read-only binding to the contract
	JBControllerTransactor // Write-only binding to the contract
	JBControllerFilterer   // Log filterer for contract events
}

// JBControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type JBControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JBControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JBControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JBControllerSession struct {
	Contract     *JBController     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JBControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JBControllerCallerSession struct {
	Contract *JBControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// JBControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JBControllerTransactorSession struct {
	Contract     *JBControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// JBControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type JBControllerRaw struct {
	Contract *JBController // Generic contract binding to access the raw methods on
}

// JBControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JBControllerCallerRaw struct {
	Contract *JBControllerCaller // Generic read-only contract binding to access the raw methods on
}

// JBControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JBControllerTransactorRaw struct {
	Contract *JBControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJBController creates a new instance of JBController, bound to a specific deployed contract.
func NewJBController(address common.Address, backend bind.ContractBackend) (*JBController, error) {
	contract, err := bindJBController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JBController{JBControllerCaller: JBControllerCaller{contract: contract}, JBControllerTransactor: JBControllerTransactor{contract: contract}, JBControllerFilterer: JBControllerFilterer{contract: contract}}, nil
}

// NewJBControllerCaller creates a new read-only instance of JBController, bound to a specific deployed contract.
func NewJBControllerCaller(address common.Address, caller bind.ContractCaller) (*JBControllerCaller, error) {
	contract, err := bindJBController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JBControllerCaller{contract: contract}, nil
}

// NewJBControllerTransactor creates a new write-only instance of JBController, bound to a specific deployed contract.
func NewJBControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*JBControllerTransactor, error) {
	contract, err := bindJBController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JBControllerTransactor{contract: contract}, nil
}

// NewJBControllerFilterer creates a new log filterer instance of JBController, bound to a specific deployed contract.
func NewJBControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*JBControllerFilterer, error) {
	contract, err := bindJBController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JBControllerFilterer{contract: contract}, nil
}

// bindJBController binds a generic wrapper to an already deployed contract.
func bindJBController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JBControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JBController *JBControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JBController.Contract.JBControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JBController *JBControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBController.Contract.JBControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JBController *JBControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JBController.Contract.JBControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JBController *JBControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JBController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JBController *JBControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JBController *JBControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JBController.Contract.contract.Transact(opts, method, params...)
}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_JBController *JBControllerCaller) DIRECTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "DIRECTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_JBController *JBControllerSession) DIRECTORY() (common.Address, error) {
	return _JBController.Contract.DIRECTORY(&_JBController.CallOpts)
}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_JBController *JBControllerCallerSession) DIRECTORY() (common.Address, error) {
	return _JBController.Contract.DIRECTORY(&_JBController.CallOpts)
}

// FUNDACCESSLIMITS is a free data retrieval call binding the contract method 0xffa08244.
//
// Solidity: function FUND_ACCESS_LIMITS() view returns(address)
func (_JBController *JBControllerCaller) FUNDACCESSLIMITS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "FUND_ACCESS_LIMITS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNDACCESSLIMITS is a free data retrieval call binding the contract method 0xffa08244.
//
// Solidity: function FUND_ACCESS_LIMITS() view returns(address)
func (_JBController *JBControllerSession) FUNDACCESSLIMITS() (common.Address, error) {
	return _JBController.Contract.FUNDACCESSLIMITS(&_JBController.CallOpts)
}

// FUNDACCESSLIMITS is a free data retrieval call binding the contract method 0xffa08244.
//
// Solidity: function FUND_ACCESS_LIMITS() view returns(address)
func (_JBController *JBControllerCallerSession) FUNDACCESSLIMITS() (common.Address, error) {
	return _JBController.Contract.FUNDACCESSLIMITS(&_JBController.CallOpts)
}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_JBController *JBControllerCaller) PERMISSIONS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "PERMISSIONS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_JBController *JBControllerSession) PERMISSIONS() (common.Address, error) {
	return _JBController.Contract.PERMISSIONS(&_JBController.CallOpts)
}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_JBController *JBControllerCallerSession) PERMISSIONS() (common.Address, error) {
	return _JBController.Contract.PERMISSIONS(&_JBController.CallOpts)
}

// PROJECTS is a free data retrieval call binding the contract method 0x293c4999.
//
// Solidity: function PROJECTS() view returns(address)
func (_JBController *JBControllerCaller) PROJECTS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "PROJECTS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROJECTS is a free data retrieval call binding the contract method 0x293c4999.
//
// Solidity: function PROJECTS() view returns(address)
func (_JBController *JBControllerSession) PROJECTS() (common.Address, error) {
	return _JBController.Contract.PROJECTS(&_JBController.CallOpts)
}

// PROJECTS is a free data retrieval call binding the contract method 0x293c4999.
//
// Solidity: function PROJECTS() view returns(address)
func (_JBController *JBControllerCallerSession) PROJECTS() (common.Address, error) {
	return _JBController.Contract.PROJECTS(&_JBController.CallOpts)
}

// RULESETS is a free data retrieval call binding the contract method 0xd4a1b4b1.
//
// Solidity: function RULESETS() view returns(address)
func (_JBController *JBControllerCaller) RULESETS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "RULESETS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RULESETS is a free data retrieval call binding the contract method 0xd4a1b4b1.
//
// Solidity: function RULESETS() view returns(address)
func (_JBController *JBControllerSession) RULESETS() (common.Address, error) {
	return _JBController.Contract.RULESETS(&_JBController.CallOpts)
}

// RULESETS is a free data retrieval call binding the contract method 0xd4a1b4b1.
//
// Solidity: function RULESETS() view returns(address)
func (_JBController *JBControllerCallerSession) RULESETS() (common.Address, error) {
	return _JBController.Contract.RULESETS(&_JBController.CallOpts)
}

// SPLITS is a free data retrieval call binding the contract method 0x1f47ce69.
//
// Solidity: function SPLITS() view returns(address)
func (_JBController *JBControllerCaller) SPLITS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "SPLITS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPLITS is a free data retrieval call binding the contract method 0x1f47ce69.
//
// Solidity: function SPLITS() view returns(address)
func (_JBController *JBControllerSession) SPLITS() (common.Address, error) {
	return _JBController.Contract.SPLITS(&_JBController.CallOpts)
}

// SPLITS is a free data retrieval call binding the contract method 0x1f47ce69.
//
// Solidity: function SPLITS() view returns(address)
func (_JBController *JBControllerCallerSession) SPLITS() (common.Address, error) {
	return _JBController.Contract.SPLITS(&_JBController.CallOpts)
}

// TOKENS is a free data retrieval call binding the contract method 0x1d831d5c.
//
// Solidity: function TOKENS() view returns(address)
func (_JBController *JBControllerCaller) TOKENS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "TOKENS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENS is a free data retrieval call binding the contract method 0x1d831d5c.
//
// Solidity: function TOKENS() view returns(address)
func (_JBController *JBControllerSession) TOKENS() (common.Address, error) {
	return _JBController.Contract.TOKENS(&_JBController.CallOpts)
}

// TOKENS is a free data retrieval call binding the contract method 0x1d831d5c.
//
// Solidity: function TOKENS() view returns(address)
func (_JBController *JBControllerCallerSession) TOKENS() (common.Address, error) {
	return _JBController.Contract.TOKENS(&_JBController.CallOpts)
}

// CurrentRulesetOf is a free data retrieval call binding the contract method 0x41929626.
//
// Solidity: function currentRulesetOf(uint256 projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_JBController *JBControllerCaller) CurrentRulesetOf(opts *bind.CallOpts, projectId *big.Int) (struct {
	Ruleset  JBRuleset
	Metadata JBRulesetMetadata
}, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "currentRulesetOf", projectId)

	outstruct := new(struct {
		Ruleset  JBRuleset
		Metadata JBRulesetMetadata
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ruleset = *abi.ConvertType(out[0], new(JBRuleset)).(*JBRuleset)
	outstruct.Metadata = *abi.ConvertType(out[1], new(JBRulesetMetadata)).(*JBRulesetMetadata)

	return *outstruct, err

}

// CurrentRulesetOf is a free data retrieval call binding the contract method 0x41929626.
//
// Solidity: function currentRulesetOf(uint256 projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_JBController *JBControllerSession) CurrentRulesetOf(projectId *big.Int) (struct {
	Ruleset  JBRuleset
	Metadata JBRulesetMetadata
}, error) {
	return _JBController.Contract.CurrentRulesetOf(&_JBController.CallOpts, projectId)
}

// CurrentRulesetOf is a free data retrieval call binding the contract method 0x41929626.
//
// Solidity: function currentRulesetOf(uint256 projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_JBController *JBControllerCallerSession) CurrentRulesetOf(projectId *big.Int) (struct {
	Ruleset  JBRuleset
	Metadata JBRulesetMetadata
}, error) {
	return _JBController.Contract.CurrentRulesetOf(&_JBController.CallOpts, projectId)
}

// GetRulesetOf is a free data retrieval call binding the contract method 0x25a61d5c.
//
// Solidity: function getRulesetOf(uint256 projectId, uint256 rulesetId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_JBController *JBControllerCaller) GetRulesetOf(opts *bind.CallOpts, projectId *big.Int, rulesetId *big.Int) (struct {
	Ruleset  JBRuleset
	Metadata JBRulesetMetadata
}, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "getRulesetOf", projectId, rulesetId)

	outstruct := new(struct {
		Ruleset  JBRuleset
		Metadata JBRulesetMetadata
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ruleset = *abi.ConvertType(out[0], new(JBRuleset)).(*JBRuleset)
	outstruct.Metadata = *abi.ConvertType(out[1], new(JBRulesetMetadata)).(*JBRulesetMetadata)

	return *outstruct, err

}

// GetRulesetOf is a free data retrieval call binding the contract method 0x25a61d5c.
//
// Solidity: function getRulesetOf(uint256 projectId, uint256 rulesetId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_JBController *JBControllerSession) GetRulesetOf(projectId *big.Int, rulesetId *big.Int) (struct {
	Ruleset  JBRuleset
	Metadata JBRulesetMetadata
}, error) {
	return _JBController.Contract.GetRulesetOf(&_JBController.CallOpts, projectId, rulesetId)
}

// GetRulesetOf is a free data retrieval call binding the contract method 0x25a61d5c.
//
// Solidity: function getRulesetOf(uint256 projectId, uint256 rulesetId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_JBController *JBControllerCallerSession) GetRulesetOf(projectId *big.Int, rulesetId *big.Int) (struct {
	Ruleset  JBRuleset
	Metadata JBRulesetMetadata
}, error) {
	return _JBController.Contract.GetRulesetOf(&_JBController.CallOpts, projectId, rulesetId)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_JBController *JBControllerCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_JBController *JBControllerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _JBController.Contract.IsTrustedForwarder(&_JBController.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_JBController *JBControllerCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _JBController.Contract.IsTrustedForwarder(&_JBController.CallOpts, forwarder)
}

// LatestQueuedRulesetOf is a free data retrieval call binding the contract method 0xc1ec61ee.
//
// Solidity: function latestQueuedRulesetOf(uint256 projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata, uint8 approvalStatus)
func (_JBController *JBControllerCaller) LatestQueuedRulesetOf(opts *bind.CallOpts, projectId *big.Int) (struct {
	Ruleset        JBRuleset
	Metadata       JBRulesetMetadata
	ApprovalStatus uint8
}, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "latestQueuedRulesetOf", projectId)

	outstruct := new(struct {
		Ruleset        JBRuleset
		Metadata       JBRulesetMetadata
		ApprovalStatus uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ruleset = *abi.ConvertType(out[0], new(JBRuleset)).(*JBRuleset)
	outstruct.Metadata = *abi.ConvertType(out[1], new(JBRulesetMetadata)).(*JBRulesetMetadata)
	outstruct.ApprovalStatus = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// LatestQueuedRulesetOf is a free data retrieval call binding the contract method 0xc1ec61ee.
//
// Solidity: function latestQueuedRulesetOf(uint256 projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata, uint8 approvalStatus)
func (_JBController *JBControllerSession) LatestQueuedRulesetOf(projectId *big.Int) (struct {
	Ruleset        JBRuleset
	Metadata       JBRulesetMetadata
	ApprovalStatus uint8
}, error) {
	return _JBController.Contract.LatestQueuedRulesetOf(&_JBController.CallOpts, projectId)
}

// LatestQueuedRulesetOf is a free data retrieval call binding the contract method 0xc1ec61ee.
//
// Solidity: function latestQueuedRulesetOf(uint256 projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata, uint8 approvalStatus)
func (_JBController *JBControllerCallerSession) LatestQueuedRulesetOf(projectId *big.Int) (struct {
	Ruleset        JBRuleset
	Metadata       JBRulesetMetadata
	ApprovalStatus uint8
}, error) {
	return _JBController.Contract.LatestQueuedRulesetOf(&_JBController.CallOpts, projectId)
}

// PendingReservedTokenBalanceOf is a free data retrieval call binding the contract method 0x39975571.
//
// Solidity: function pendingReservedTokenBalanceOf(uint256 projectId) view returns(uint256)
func (_JBController *JBControllerCaller) PendingReservedTokenBalanceOf(opts *bind.CallOpts, projectId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "pendingReservedTokenBalanceOf", projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReservedTokenBalanceOf is a free data retrieval call binding the contract method 0x39975571.
//
// Solidity: function pendingReservedTokenBalanceOf(uint256 projectId) view returns(uint256)
func (_JBController *JBControllerSession) PendingReservedTokenBalanceOf(projectId *big.Int) (*big.Int, error) {
	return _JBController.Contract.PendingReservedTokenBalanceOf(&_JBController.CallOpts, projectId)
}

// PendingReservedTokenBalanceOf is a free data retrieval call binding the contract method 0x39975571.
//
// Solidity: function pendingReservedTokenBalanceOf(uint256 projectId) view returns(uint256)
func (_JBController *JBControllerCallerSession) PendingReservedTokenBalanceOf(projectId *big.Int) (*big.Int, error) {
	return _JBController.Contract.PendingReservedTokenBalanceOf(&_JBController.CallOpts, projectId)
}

// RulesetsOf is a free data retrieval call binding the contract method 0xb51d6244.
//
// Solidity: function rulesetsOf(uint256 projectId, uint256 startingId, uint256 size) view returns(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256),(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256))[] rulesets)
func (_JBController *JBControllerCaller) RulesetsOf(opts *bind.CallOpts, projectId *big.Int, startingId *big.Int, size *big.Int) ([]JBRulesetWithMetadata, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "rulesetsOf", projectId, startingId, size)

	if err != nil {
		return *new([]JBRulesetWithMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]JBRulesetWithMetadata)).(*[]JBRulesetWithMetadata)

	return out0, err

}

// RulesetsOf is a free data retrieval call binding the contract method 0xb51d6244.
//
// Solidity: function rulesetsOf(uint256 projectId, uint256 startingId, uint256 size) view returns(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256),(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256))[] rulesets)
func (_JBController *JBControllerSession) RulesetsOf(projectId *big.Int, startingId *big.Int, size *big.Int) ([]JBRulesetWithMetadata, error) {
	return _JBController.Contract.RulesetsOf(&_JBController.CallOpts, projectId, startingId, size)
}

// RulesetsOf is a free data retrieval call binding the contract method 0xb51d6244.
//
// Solidity: function rulesetsOf(uint256 projectId, uint256 startingId, uint256 size) view returns(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256),(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256))[] rulesets)
func (_JBController *JBControllerCallerSession) RulesetsOf(projectId *big.Int, startingId *big.Int, size *big.Int) ([]JBRulesetWithMetadata, error) {
	return _JBController.Contract.RulesetsOf(&_JBController.CallOpts, projectId, startingId, size)
}

// SetControllerAllowed is a free data retrieval call binding the contract method 0x99d25a34.
//
// Solidity: function setControllerAllowed(uint256 projectId) view returns(bool)
func (_JBController *JBControllerCaller) SetControllerAllowed(opts *bind.CallOpts, projectId *big.Int) (bool, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "setControllerAllowed", projectId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SetControllerAllowed is a free data retrieval call binding the contract method 0x99d25a34.
//
// Solidity: function setControllerAllowed(uint256 projectId) view returns(bool)
func (_JBController *JBControllerSession) SetControllerAllowed(projectId *big.Int) (bool, error) {
	return _JBController.Contract.SetControllerAllowed(&_JBController.CallOpts, projectId)
}

// SetControllerAllowed is a free data retrieval call binding the contract method 0x99d25a34.
//
// Solidity: function setControllerAllowed(uint256 projectId) view returns(bool)
func (_JBController *JBControllerCallerSession) SetControllerAllowed(projectId *big.Int) (bool, error) {
	return _JBController.Contract.SetControllerAllowed(&_JBController.CallOpts, projectId)
}

// SetTerminalsAllowed is a free data retrieval call binding the contract method 0xb1a50e33.
//
// Solidity: function setTerminalsAllowed(uint256 projectId) view returns(bool)
func (_JBController *JBControllerCaller) SetTerminalsAllowed(opts *bind.CallOpts, projectId *big.Int) (bool, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "setTerminalsAllowed", projectId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SetTerminalsAllowed is a free data retrieval call binding the contract method 0xb1a50e33.
//
// Solidity: function setTerminalsAllowed(uint256 projectId) view returns(bool)
func (_JBController *JBControllerSession) SetTerminalsAllowed(projectId *big.Int) (bool, error) {
	return _JBController.Contract.SetTerminalsAllowed(&_JBController.CallOpts, projectId)
}

// SetTerminalsAllowed is a free data retrieval call binding the contract method 0xb1a50e33.
//
// Solidity: function setTerminalsAllowed(uint256 projectId) view returns(bool)
func (_JBController *JBControllerCallerSession) SetTerminalsAllowed(projectId *big.Int) (bool, error) {
	return _JBController.Contract.SetTerminalsAllowed(&_JBController.CallOpts, projectId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JBController *JBControllerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JBController *JBControllerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _JBController.Contract.SupportsInterface(&_JBController.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JBController *JBControllerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _JBController.Contract.SupportsInterface(&_JBController.CallOpts, interfaceId)
}

// TotalTokenSupplyWithReservedTokensOf is a free data retrieval call binding the contract method 0x4da781a9.
//
// Solidity: function totalTokenSupplyWithReservedTokensOf(uint256 projectId) view returns(uint256)
func (_JBController *JBControllerCaller) TotalTokenSupplyWithReservedTokensOf(opts *bind.CallOpts, projectId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "totalTokenSupplyWithReservedTokensOf", projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenSupplyWithReservedTokensOf is a free data retrieval call binding the contract method 0x4da781a9.
//
// Solidity: function totalTokenSupplyWithReservedTokensOf(uint256 projectId) view returns(uint256)
func (_JBController *JBControllerSession) TotalTokenSupplyWithReservedTokensOf(projectId *big.Int) (*big.Int, error) {
	return _JBController.Contract.TotalTokenSupplyWithReservedTokensOf(&_JBController.CallOpts, projectId)
}

// TotalTokenSupplyWithReservedTokensOf is a free data retrieval call binding the contract method 0x4da781a9.
//
// Solidity: function totalTokenSupplyWithReservedTokensOf(uint256 projectId) view returns(uint256)
func (_JBController *JBControllerCallerSession) TotalTokenSupplyWithReservedTokensOf(projectId *big.Int) (*big.Int, error) {
	return _JBController.Contract.TotalTokenSupplyWithReservedTokensOf(&_JBController.CallOpts, projectId)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_JBController *JBControllerCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_JBController *JBControllerSession) TrustedForwarder() (common.Address, error) {
	return _JBController.Contract.TrustedForwarder(&_JBController.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_JBController *JBControllerCallerSession) TrustedForwarder() (common.Address, error) {
	return _JBController.Contract.TrustedForwarder(&_JBController.CallOpts)
}

// UpcomingRulesetOf is a free data retrieval call binding the contract method 0xc02c63ad.
//
// Solidity: function upcomingRulesetOf(uint256 projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_JBController *JBControllerCaller) UpcomingRulesetOf(opts *bind.CallOpts, projectId *big.Int) (struct {
	Ruleset  JBRuleset
	Metadata JBRulesetMetadata
}, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "upcomingRulesetOf", projectId)

	outstruct := new(struct {
		Ruleset  JBRuleset
		Metadata JBRulesetMetadata
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ruleset = *abi.ConvertType(out[0], new(JBRuleset)).(*JBRuleset)
	outstruct.Metadata = *abi.ConvertType(out[1], new(JBRulesetMetadata)).(*JBRulesetMetadata)

	return *outstruct, err

}

// UpcomingRulesetOf is a free data retrieval call binding the contract method 0xc02c63ad.
//
// Solidity: function upcomingRulesetOf(uint256 projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_JBController *JBControllerSession) UpcomingRulesetOf(projectId *big.Int) (struct {
	Ruleset  JBRuleset
	Metadata JBRulesetMetadata
}, error) {
	return _JBController.Contract.UpcomingRulesetOf(&_JBController.CallOpts, projectId)
}

// UpcomingRulesetOf is a free data retrieval call binding the contract method 0xc02c63ad.
//
// Solidity: function upcomingRulesetOf(uint256 projectId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256) ruleset, (uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256) metadata)
func (_JBController *JBControllerCallerSession) UpcomingRulesetOf(projectId *big.Int) (struct {
	Ruleset  JBRuleset
	Metadata JBRulesetMetadata
}, error) {
	return _JBController.Contract.UpcomingRulesetOf(&_JBController.CallOpts, projectId)
}

// UriOf is a free data retrieval call binding the contract method 0xa312889b.
//
// Solidity: function uriOf(uint256 projectId) view returns(string)
func (_JBController *JBControllerCaller) UriOf(opts *bind.CallOpts, projectId *big.Int) (string, error) {
	var out []interface{}
	err := _JBController.contract.Call(opts, &out, "uriOf", projectId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UriOf is a free data retrieval call binding the contract method 0xa312889b.
//
// Solidity: function uriOf(uint256 projectId) view returns(string)
func (_JBController *JBControllerSession) UriOf(projectId *big.Int) (string, error) {
	return _JBController.Contract.UriOf(&_JBController.CallOpts, projectId)
}

// UriOf is a free data retrieval call binding the contract method 0xa312889b.
//
// Solidity: function uriOf(uint256 projectId) view returns(string)
func (_JBController *JBControllerCallerSession) UriOf(projectId *big.Int) (string, error) {
	return _JBController.Contract.UriOf(&_JBController.CallOpts, projectId)
}

// BurnTokensOf is a paid mutator transaction binding the contract method 0xa2d532e6.
//
// Solidity: function burnTokensOf(address holder, uint256 projectId, uint256 tokenCount, string memo) returns()
func (_JBController *JBControllerTransactor) BurnTokensOf(opts *bind.TransactOpts, holder common.Address, projectId *big.Int, tokenCount *big.Int, memo string) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "burnTokensOf", holder, projectId, tokenCount, memo)
}

// BurnTokensOf is a paid mutator transaction binding the contract method 0xa2d532e6.
//
// Solidity: function burnTokensOf(address holder, uint256 projectId, uint256 tokenCount, string memo) returns()
func (_JBController *JBControllerSession) BurnTokensOf(holder common.Address, projectId *big.Int, tokenCount *big.Int, memo string) (*types.Transaction, error) {
	return _JBController.Contract.BurnTokensOf(&_JBController.TransactOpts, holder, projectId, tokenCount, memo)
}

// BurnTokensOf is a paid mutator transaction binding the contract method 0xa2d532e6.
//
// Solidity: function burnTokensOf(address holder, uint256 projectId, uint256 tokenCount, string memo) returns()
func (_JBController *JBControllerTransactorSession) BurnTokensOf(holder common.Address, projectId *big.Int, tokenCount *big.Int, memo string) (*types.Transaction, error) {
	return _JBController.Contract.BurnTokensOf(&_JBController.TransactOpts, holder, projectId, tokenCount, memo)
}

// ClaimTokensFor is a paid mutator transaction binding the contract method 0x303f5dfa.
//
// Solidity: function claimTokensFor(address holder, uint256 projectId, uint256 amount, address beneficiary) returns()
func (_JBController *JBControllerTransactor) ClaimTokensFor(opts *bind.TransactOpts, holder common.Address, projectId *big.Int, amount *big.Int, beneficiary common.Address) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "claimTokensFor", holder, projectId, amount, beneficiary)
}

// ClaimTokensFor is a paid mutator transaction binding the contract method 0x303f5dfa.
//
// Solidity: function claimTokensFor(address holder, uint256 projectId, uint256 amount, address beneficiary) returns()
func (_JBController *JBControllerSession) ClaimTokensFor(holder common.Address, projectId *big.Int, amount *big.Int, beneficiary common.Address) (*types.Transaction, error) {
	return _JBController.Contract.ClaimTokensFor(&_JBController.TransactOpts, holder, projectId, amount, beneficiary)
}

// ClaimTokensFor is a paid mutator transaction binding the contract method 0x303f5dfa.
//
// Solidity: function claimTokensFor(address holder, uint256 projectId, uint256 amount, address beneficiary) returns()
func (_JBController *JBControllerTransactorSession) ClaimTokensFor(holder common.Address, projectId *big.Int, amount *big.Int, beneficiary common.Address) (*types.Transaction, error) {
	return _JBController.Contract.ClaimTokensFor(&_JBController.TransactOpts, holder, projectId, amount, beneficiary)
}

// DeployERC20For is a paid mutator transaction binding the contract method 0x58178191.
//
// Solidity: function deployERC20For(uint256 projectId, string name, string symbol, bytes32 salt) returns(address token)
func (_JBController *JBControllerTransactor) DeployERC20For(opts *bind.TransactOpts, projectId *big.Int, name string, symbol string, salt [32]byte) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "deployERC20For", projectId, name, symbol, salt)
}

// DeployERC20For is a paid mutator transaction binding the contract method 0x58178191.
//
// Solidity: function deployERC20For(uint256 projectId, string name, string symbol, bytes32 salt) returns(address token)
func (_JBController *JBControllerSession) DeployERC20For(projectId *big.Int, name string, symbol string, salt [32]byte) (*types.Transaction, error) {
	return _JBController.Contract.DeployERC20For(&_JBController.TransactOpts, projectId, name, symbol, salt)
}

// DeployERC20For is a paid mutator transaction binding the contract method 0x58178191.
//
// Solidity: function deployERC20For(uint256 projectId, string name, string symbol, bytes32 salt) returns(address token)
func (_JBController *JBControllerTransactorSession) DeployERC20For(projectId *big.Int, name string, symbol string, salt [32]byte) (*types.Transaction, error) {
	return _JBController.Contract.DeployERC20For(&_JBController.TransactOpts, projectId, name, symbol, salt)
}

// LaunchProjectFor is a paid mutator transaction binding the contract method 0x0634bae9.
//
// Solidity: function launchProjectFor(address owner, string projectUri, (uint256,uint256,uint256,uint256,address,(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256),(uint256,(bool,uint256,uint256,address,uint256,address)[])[],(address,address,(uint256,uint256)[],(uint256,uint256)[])[])[] rulesetConfigurations, (address,address[])[] terminalConfigurations, string memo) returns(uint256 projectId)
func (_JBController *JBControllerTransactor) LaunchProjectFor(opts *bind.TransactOpts, owner common.Address, projectUri string, rulesetConfigurations []JBRulesetConfig, terminalConfigurations []JBTerminalConfig, memo string) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "launchProjectFor", owner, projectUri, rulesetConfigurations, terminalConfigurations, memo)
}

// LaunchProjectFor is a paid mutator transaction binding the contract method 0x0634bae9.
//
// Solidity: function launchProjectFor(address owner, string projectUri, (uint256,uint256,uint256,uint256,address,(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256),(uint256,(bool,uint256,uint256,address,uint256,address)[])[],(address,address,(uint256,uint256)[],(uint256,uint256)[])[])[] rulesetConfigurations, (address,address[])[] terminalConfigurations, string memo) returns(uint256 projectId)
func (_JBController *JBControllerSession) LaunchProjectFor(owner common.Address, projectUri string, rulesetConfigurations []JBRulesetConfig, terminalConfigurations []JBTerminalConfig, memo string) (*types.Transaction, error) {
	return _JBController.Contract.LaunchProjectFor(&_JBController.TransactOpts, owner, projectUri, rulesetConfigurations, terminalConfigurations, memo)
}

// LaunchProjectFor is a paid mutator transaction binding the contract method 0x0634bae9.
//
// Solidity: function launchProjectFor(address owner, string projectUri, (uint256,uint256,uint256,uint256,address,(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256),(uint256,(bool,uint256,uint256,address,uint256,address)[])[],(address,address,(uint256,uint256)[],(uint256,uint256)[])[])[] rulesetConfigurations, (address,address[])[] terminalConfigurations, string memo) returns(uint256 projectId)
func (_JBController *JBControllerTransactorSession) LaunchProjectFor(owner common.Address, projectUri string, rulesetConfigurations []JBRulesetConfig, terminalConfigurations []JBTerminalConfig, memo string) (*types.Transaction, error) {
	return _JBController.Contract.LaunchProjectFor(&_JBController.TransactOpts, owner, projectUri, rulesetConfigurations, terminalConfigurations, memo)
}

// LaunchRulesetsFor is a paid mutator transaction binding the contract method 0x1ea7bc8c.
//
// Solidity: function launchRulesetsFor(uint256 projectId, (uint256,uint256,uint256,uint256,address,(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256),(uint256,(bool,uint256,uint256,address,uint256,address)[])[],(address,address,(uint256,uint256)[],(uint256,uint256)[])[])[] rulesetConfigurations, (address,address[])[] terminalConfigurations, string memo) returns(uint256 rulesetId)
func (_JBController *JBControllerTransactor) LaunchRulesetsFor(opts *bind.TransactOpts, projectId *big.Int, rulesetConfigurations []JBRulesetConfig, terminalConfigurations []JBTerminalConfig, memo string) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "launchRulesetsFor", projectId, rulesetConfigurations, terminalConfigurations, memo)
}

// LaunchRulesetsFor is a paid mutator transaction binding the contract method 0x1ea7bc8c.
//
// Solidity: function launchRulesetsFor(uint256 projectId, (uint256,uint256,uint256,uint256,address,(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256),(uint256,(bool,uint256,uint256,address,uint256,address)[])[],(address,address,(uint256,uint256)[],(uint256,uint256)[])[])[] rulesetConfigurations, (address,address[])[] terminalConfigurations, string memo) returns(uint256 rulesetId)
func (_JBController *JBControllerSession) LaunchRulesetsFor(projectId *big.Int, rulesetConfigurations []JBRulesetConfig, terminalConfigurations []JBTerminalConfig, memo string) (*types.Transaction, error) {
	return _JBController.Contract.LaunchRulesetsFor(&_JBController.TransactOpts, projectId, rulesetConfigurations, terminalConfigurations, memo)
}

// LaunchRulesetsFor is a paid mutator transaction binding the contract method 0x1ea7bc8c.
//
// Solidity: function launchRulesetsFor(uint256 projectId, (uint256,uint256,uint256,uint256,address,(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256),(uint256,(bool,uint256,uint256,address,uint256,address)[])[],(address,address,(uint256,uint256)[],(uint256,uint256)[])[])[] rulesetConfigurations, (address,address[])[] terminalConfigurations, string memo) returns(uint256 rulesetId)
func (_JBController *JBControllerTransactorSession) LaunchRulesetsFor(projectId *big.Int, rulesetConfigurations []JBRulesetConfig, terminalConfigurations []JBTerminalConfig, memo string) (*types.Transaction, error) {
	return _JBController.Contract.LaunchRulesetsFor(&_JBController.TransactOpts, projectId, rulesetConfigurations, terminalConfigurations, memo)
}

// MigrateController is a paid mutator transaction binding the contract method 0x5338fd9e.
//
// Solidity: function migrateController(uint256 projectId, address to) returns()
func (_JBController *JBControllerTransactor) MigrateController(opts *bind.TransactOpts, projectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "migrateController", projectId, to)
}

// MigrateController is a paid mutator transaction binding the contract method 0x5338fd9e.
//
// Solidity: function migrateController(uint256 projectId, address to) returns()
func (_JBController *JBControllerSession) MigrateController(projectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _JBController.Contract.MigrateController(&_JBController.TransactOpts, projectId, to)
}

// MigrateController is a paid mutator transaction binding the contract method 0x5338fd9e.
//
// Solidity: function migrateController(uint256 projectId, address to) returns()
func (_JBController *JBControllerTransactorSession) MigrateController(projectId *big.Int, to common.Address) (*types.Transaction, error) {
	return _JBController.Contract.MigrateController(&_JBController.TransactOpts, projectId, to)
}

// MintTokensOf is a paid mutator transaction binding the contract method 0xc7fb92de.
//
// Solidity: function mintTokensOf(uint256 projectId, uint256 tokenCount, address beneficiary, string memo, bool useReservedRate) returns(uint256 beneficiaryTokenCount)
func (_JBController *JBControllerTransactor) MintTokensOf(opts *bind.TransactOpts, projectId *big.Int, tokenCount *big.Int, beneficiary common.Address, memo string, useReservedRate bool) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "mintTokensOf", projectId, tokenCount, beneficiary, memo, useReservedRate)
}

// MintTokensOf is a paid mutator transaction binding the contract method 0xc7fb92de.
//
// Solidity: function mintTokensOf(uint256 projectId, uint256 tokenCount, address beneficiary, string memo, bool useReservedRate) returns(uint256 beneficiaryTokenCount)
func (_JBController *JBControllerSession) MintTokensOf(projectId *big.Int, tokenCount *big.Int, beneficiary common.Address, memo string, useReservedRate bool) (*types.Transaction, error) {
	return _JBController.Contract.MintTokensOf(&_JBController.TransactOpts, projectId, tokenCount, beneficiary, memo, useReservedRate)
}

// MintTokensOf is a paid mutator transaction binding the contract method 0xc7fb92de.
//
// Solidity: function mintTokensOf(uint256 projectId, uint256 tokenCount, address beneficiary, string memo, bool useReservedRate) returns(uint256 beneficiaryTokenCount)
func (_JBController *JBControllerTransactorSession) MintTokensOf(projectId *big.Int, tokenCount *big.Int, beneficiary common.Address, memo string, useReservedRate bool) (*types.Transaction, error) {
	return _JBController.Contract.MintTokensOf(&_JBController.TransactOpts, projectId, tokenCount, beneficiary, memo, useReservedRate)
}

// PayReservedTokenToTerminal is a paid mutator transaction binding the contract method 0x0a1181ed.
//
// Solidity: function payReservedTokenToTerminal(address terminal, uint256 projectId, address token, uint256 splitAmount, address beneficiary, bytes metadata) returns()
func (_JBController *JBControllerTransactor) PayReservedTokenToTerminal(opts *bind.TransactOpts, terminal common.Address, projectId *big.Int, token common.Address, splitAmount *big.Int, beneficiary common.Address, metadata []byte) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "payReservedTokenToTerminal", terminal, projectId, token, splitAmount, beneficiary, metadata)
}

// PayReservedTokenToTerminal is a paid mutator transaction binding the contract method 0x0a1181ed.
//
// Solidity: function payReservedTokenToTerminal(address terminal, uint256 projectId, address token, uint256 splitAmount, address beneficiary, bytes metadata) returns()
func (_JBController *JBControllerSession) PayReservedTokenToTerminal(terminal common.Address, projectId *big.Int, token common.Address, splitAmount *big.Int, beneficiary common.Address, metadata []byte) (*types.Transaction, error) {
	return _JBController.Contract.PayReservedTokenToTerminal(&_JBController.TransactOpts, terminal, projectId, token, splitAmount, beneficiary, metadata)
}

// PayReservedTokenToTerminal is a paid mutator transaction binding the contract method 0x0a1181ed.
//
// Solidity: function payReservedTokenToTerminal(address terminal, uint256 projectId, address token, uint256 splitAmount, address beneficiary, bytes metadata) returns()
func (_JBController *JBControllerTransactorSession) PayReservedTokenToTerminal(terminal common.Address, projectId *big.Int, token common.Address, splitAmount *big.Int, beneficiary common.Address, metadata []byte) (*types.Transaction, error) {
	return _JBController.Contract.PayReservedTokenToTerminal(&_JBController.TransactOpts, terminal, projectId, token, splitAmount, beneficiary, metadata)
}

// QueueRulesetsOf is a paid mutator transaction binding the contract method 0xa0d8e2a9.
//
// Solidity: function queueRulesetsOf(uint256 projectId, (uint256,uint256,uint256,uint256,address,(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256),(uint256,(bool,uint256,uint256,address,uint256,address)[])[],(address,address,(uint256,uint256)[],(uint256,uint256)[])[])[] rulesetConfigurations, string memo) returns(uint256 rulesetId)
func (_JBController *JBControllerTransactor) QueueRulesetsOf(opts *bind.TransactOpts, projectId *big.Int, rulesetConfigurations []JBRulesetConfig, memo string) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "queueRulesetsOf", projectId, rulesetConfigurations, memo)
}

// QueueRulesetsOf is a paid mutator transaction binding the contract method 0xa0d8e2a9.
//
// Solidity: function queueRulesetsOf(uint256 projectId, (uint256,uint256,uint256,uint256,address,(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256),(uint256,(bool,uint256,uint256,address,uint256,address)[])[],(address,address,(uint256,uint256)[],(uint256,uint256)[])[])[] rulesetConfigurations, string memo) returns(uint256 rulesetId)
func (_JBController *JBControllerSession) QueueRulesetsOf(projectId *big.Int, rulesetConfigurations []JBRulesetConfig, memo string) (*types.Transaction, error) {
	return _JBController.Contract.QueueRulesetsOf(&_JBController.TransactOpts, projectId, rulesetConfigurations, memo)
}

// QueueRulesetsOf is a paid mutator transaction binding the contract method 0xa0d8e2a9.
//
// Solidity: function queueRulesetsOf(uint256 projectId, (uint256,uint256,uint256,uint256,address,(uint256,uint256,uint256,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address,uint256),(uint256,(bool,uint256,uint256,address,uint256,address)[])[],(address,address,(uint256,uint256)[],(uint256,uint256)[])[])[] rulesetConfigurations, string memo) returns(uint256 rulesetId)
func (_JBController *JBControllerTransactorSession) QueueRulesetsOf(projectId *big.Int, rulesetConfigurations []JBRulesetConfig, memo string) (*types.Transaction, error) {
	return _JBController.Contract.QueueRulesetsOf(&_JBController.TransactOpts, projectId, rulesetConfigurations, memo)
}

// ReceiveMigrationFrom is a paid mutator transaction binding the contract method 0xc5b19b0e.
//
// Solidity: function receiveMigrationFrom(address from, uint256 projectId) returns()
func (_JBController *JBControllerTransactor) ReceiveMigrationFrom(opts *bind.TransactOpts, from common.Address, projectId *big.Int) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "receiveMigrationFrom", from, projectId)
}

// ReceiveMigrationFrom is a paid mutator transaction binding the contract method 0xc5b19b0e.
//
// Solidity: function receiveMigrationFrom(address from, uint256 projectId) returns()
func (_JBController *JBControllerSession) ReceiveMigrationFrom(from common.Address, projectId *big.Int) (*types.Transaction, error) {
	return _JBController.Contract.ReceiveMigrationFrom(&_JBController.TransactOpts, from, projectId)
}

// ReceiveMigrationFrom is a paid mutator transaction binding the contract method 0xc5b19b0e.
//
// Solidity: function receiveMigrationFrom(address from, uint256 projectId) returns()
func (_JBController *JBControllerTransactorSession) ReceiveMigrationFrom(from common.Address, projectId *big.Int) (*types.Transaction, error) {
	return _JBController.Contract.ReceiveMigrationFrom(&_JBController.TransactOpts, from, projectId)
}

// SendReservedTokensToSplitsOf is a paid mutator transaction binding the contract method 0x090db2f1.
//
// Solidity: function sendReservedTokensToSplitsOf(uint256 projectId) returns(uint256)
func (_JBController *JBControllerTransactor) SendReservedTokensToSplitsOf(opts *bind.TransactOpts, projectId *big.Int) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "sendReservedTokensToSplitsOf", projectId)
}

// SendReservedTokensToSplitsOf is a paid mutator transaction binding the contract method 0x090db2f1.
//
// Solidity: function sendReservedTokensToSplitsOf(uint256 projectId) returns(uint256)
func (_JBController *JBControllerSession) SendReservedTokensToSplitsOf(projectId *big.Int) (*types.Transaction, error) {
	return _JBController.Contract.SendReservedTokensToSplitsOf(&_JBController.TransactOpts, projectId)
}

// SendReservedTokensToSplitsOf is a paid mutator transaction binding the contract method 0x090db2f1.
//
// Solidity: function sendReservedTokensToSplitsOf(uint256 projectId) returns(uint256)
func (_JBController *JBControllerTransactorSession) SendReservedTokensToSplitsOf(projectId *big.Int) (*types.Transaction, error) {
	return _JBController.Contract.SendReservedTokensToSplitsOf(&_JBController.TransactOpts, projectId)
}

// SetSplitGroupsOf is a paid mutator transaction binding the contract method 0xdd5f214a.
//
// Solidity: function setSplitGroupsOf(uint256 projectId, uint256 rulesetId, (uint256,(bool,uint256,uint256,address,uint256,address)[])[] splitGroups) returns()
func (_JBController *JBControllerTransactor) SetSplitGroupsOf(opts *bind.TransactOpts, projectId *big.Int, rulesetId *big.Int, splitGroups []JBSplitGroup) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "setSplitGroupsOf", projectId, rulesetId, splitGroups)
}

// SetSplitGroupsOf is a paid mutator transaction binding the contract method 0xdd5f214a.
//
// Solidity: function setSplitGroupsOf(uint256 projectId, uint256 rulesetId, (uint256,(bool,uint256,uint256,address,uint256,address)[])[] splitGroups) returns()
func (_JBController *JBControllerSession) SetSplitGroupsOf(projectId *big.Int, rulesetId *big.Int, splitGroups []JBSplitGroup) (*types.Transaction, error) {
	return _JBController.Contract.SetSplitGroupsOf(&_JBController.TransactOpts, projectId, rulesetId, splitGroups)
}

// SetSplitGroupsOf is a paid mutator transaction binding the contract method 0xdd5f214a.
//
// Solidity: function setSplitGroupsOf(uint256 projectId, uint256 rulesetId, (uint256,(bool,uint256,uint256,address,uint256,address)[])[] splitGroups) returns()
func (_JBController *JBControllerTransactorSession) SetSplitGroupsOf(projectId *big.Int, rulesetId *big.Int, splitGroups []JBSplitGroup) (*types.Transaction, error) {
	return _JBController.Contract.SetSplitGroupsOf(&_JBController.TransactOpts, projectId, rulesetId, splitGroups)
}

// SetTokenFor is a paid mutator transaction binding the contract method 0xf12b64a5.
//
// Solidity: function setTokenFor(uint256 projectId, address token) returns()
func (_JBController *JBControllerTransactor) SetTokenFor(opts *bind.TransactOpts, projectId *big.Int, token common.Address) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "setTokenFor", projectId, token)
}

// SetTokenFor is a paid mutator transaction binding the contract method 0xf12b64a5.
//
// Solidity: function setTokenFor(uint256 projectId, address token) returns()
func (_JBController *JBControllerSession) SetTokenFor(projectId *big.Int, token common.Address) (*types.Transaction, error) {
	return _JBController.Contract.SetTokenFor(&_JBController.TransactOpts, projectId, token)
}

// SetTokenFor is a paid mutator transaction binding the contract method 0xf12b64a5.
//
// Solidity: function setTokenFor(uint256 projectId, address token) returns()
func (_JBController *JBControllerTransactorSession) SetTokenFor(projectId *big.Int, token common.Address) (*types.Transaction, error) {
	return _JBController.Contract.SetTokenFor(&_JBController.TransactOpts, projectId, token)
}

// SetUriOf is a paid mutator transaction binding the contract method 0x702a3977.
//
// Solidity: function setUriOf(uint256 projectId, string metadata) returns()
func (_JBController *JBControllerTransactor) SetUriOf(opts *bind.TransactOpts, projectId *big.Int, metadata string) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "setUriOf", projectId, metadata)
}

// SetUriOf is a paid mutator transaction binding the contract method 0x702a3977.
//
// Solidity: function setUriOf(uint256 projectId, string metadata) returns()
func (_JBController *JBControllerSession) SetUriOf(projectId *big.Int, metadata string) (*types.Transaction, error) {
	return _JBController.Contract.SetUriOf(&_JBController.TransactOpts, projectId, metadata)
}

// SetUriOf is a paid mutator transaction binding the contract method 0x702a3977.
//
// Solidity: function setUriOf(uint256 projectId, string metadata) returns()
func (_JBController *JBControllerTransactorSession) SetUriOf(projectId *big.Int, metadata string) (*types.Transaction, error) {
	return _JBController.Contract.SetUriOf(&_JBController.TransactOpts, projectId, metadata)
}

// TransferCreditsFrom is a paid mutator transaction binding the contract method 0xb1e6d2a1.
//
// Solidity: function transferCreditsFrom(address holder, uint256 projectId, address recipient, uint256 amount) returns()
func (_JBController *JBControllerTransactor) TransferCreditsFrom(opts *bind.TransactOpts, holder common.Address, projectId *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JBController.contract.Transact(opts, "transferCreditsFrom", holder, projectId, recipient, amount)
}

// TransferCreditsFrom is a paid mutator transaction binding the contract method 0xb1e6d2a1.
//
// Solidity: function transferCreditsFrom(address holder, uint256 projectId, address recipient, uint256 amount) returns()
func (_JBController *JBControllerSession) TransferCreditsFrom(holder common.Address, projectId *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JBController.Contract.TransferCreditsFrom(&_JBController.TransactOpts, holder, projectId, recipient, amount)
}

// TransferCreditsFrom is a paid mutator transaction binding the contract method 0xb1e6d2a1.
//
// Solidity: function transferCreditsFrom(address holder, uint256 projectId, address recipient, uint256 amount) returns()
func (_JBController *JBControllerTransactorSession) TransferCreditsFrom(holder common.Address, projectId *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JBController.Contract.TransferCreditsFrom(&_JBController.TransactOpts, holder, projectId, recipient, amount)
}

// JBControllerBurnTokensIterator is returned from FilterBurnTokens and is used to iterate over the raw logs and unpacked data for BurnTokens events raised by the JBController contract.
type JBControllerBurnTokensIterator struct {
	Event *JBControllerBurnTokens // Event containing the contract specifics and raw log

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
func (it *JBControllerBurnTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerBurnTokens)
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
		it.Event = new(JBControllerBurnTokens)
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
func (it *JBControllerBurnTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerBurnTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerBurnTokens represents a BurnTokens event raised by the JBController contract.
type JBControllerBurnTokens struct {
	Holder     common.Address
	ProjectId  *big.Int
	TokenCount *big.Int
	Memo       string
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurnTokens is a free log retrieval operation binding the contract event 0xdf04e13ee4fcd48a81ab2fd114757093740a3efa9b6475d86e05878b4c59d079.
//
// Solidity: event BurnTokens(address indexed holder, uint256 indexed projectId, uint256 tokenCount, string memo, address caller)
func (_JBController *JBControllerFilterer) FilterBurnTokens(opts *bind.FilterOpts, holder []common.Address, projectId []*big.Int) (*JBControllerBurnTokensIterator, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.FilterLogs(opts, "BurnTokens", holderRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBControllerBurnTokensIterator{contract: _JBController.contract, event: "BurnTokens", logs: logs, sub: sub}, nil
}

// WatchBurnTokens is a free log subscription operation binding the contract event 0xdf04e13ee4fcd48a81ab2fd114757093740a3efa9b6475d86e05878b4c59d079.
//
// Solidity: event BurnTokens(address indexed holder, uint256 indexed projectId, uint256 tokenCount, string memo, address caller)
func (_JBController *JBControllerFilterer) WatchBurnTokens(opts *bind.WatchOpts, sink chan<- *JBControllerBurnTokens, holder []common.Address, projectId []*big.Int) (event.Subscription, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.WatchLogs(opts, "BurnTokens", holderRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerBurnTokens)
				if err := _JBController.contract.UnpackLog(event, "BurnTokens", log); err != nil {
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

// ParseBurnTokens is a log parse operation binding the contract event 0xdf04e13ee4fcd48a81ab2fd114757093740a3efa9b6475d86e05878b4c59d079.
//
// Solidity: event BurnTokens(address indexed holder, uint256 indexed projectId, uint256 tokenCount, string memo, address caller)
func (_JBController *JBControllerFilterer) ParseBurnTokens(log types.Log) (*JBControllerBurnTokens, error) {
	event := new(JBControllerBurnTokens)
	if err := _JBController.contract.UnpackLog(event, "BurnTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBControllerLaunchProjectIterator is returned from FilterLaunchProject and is used to iterate over the raw logs and unpacked data for LaunchProject events raised by the JBController contract.
type JBControllerLaunchProjectIterator struct {
	Event *JBControllerLaunchProject // Event containing the contract specifics and raw log

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
func (it *JBControllerLaunchProjectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerLaunchProject)
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
		it.Event = new(JBControllerLaunchProject)
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
func (it *JBControllerLaunchProjectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerLaunchProjectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerLaunchProject represents a LaunchProject event raised by the JBController contract.
type JBControllerLaunchProject struct {
	RulesetId *big.Int
	ProjectId *big.Int
	Metadata  string
	Memo      string
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLaunchProject is a free log retrieval operation binding the contract event 0x8dac501f24f52bda9ebdfa6a1789878c1d1e23823c771f7d52b5ba41261b0f45.
//
// Solidity: event LaunchProject(uint256 rulesetId, uint256 projectId, string metadata, string memo, address caller)
func (_JBController *JBControllerFilterer) FilterLaunchProject(opts *bind.FilterOpts) (*JBControllerLaunchProjectIterator, error) {

	logs, sub, err := _JBController.contract.FilterLogs(opts, "LaunchProject")
	if err != nil {
		return nil, err
	}
	return &JBControllerLaunchProjectIterator{contract: _JBController.contract, event: "LaunchProject", logs: logs, sub: sub}, nil
}

// WatchLaunchProject is a free log subscription operation binding the contract event 0x8dac501f24f52bda9ebdfa6a1789878c1d1e23823c771f7d52b5ba41261b0f45.
//
// Solidity: event LaunchProject(uint256 rulesetId, uint256 projectId, string metadata, string memo, address caller)
func (_JBController *JBControllerFilterer) WatchLaunchProject(opts *bind.WatchOpts, sink chan<- *JBControllerLaunchProject) (event.Subscription, error) {

	logs, sub, err := _JBController.contract.WatchLogs(opts, "LaunchProject")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerLaunchProject)
				if err := _JBController.contract.UnpackLog(event, "LaunchProject", log); err != nil {
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

// ParseLaunchProject is a log parse operation binding the contract event 0x8dac501f24f52bda9ebdfa6a1789878c1d1e23823c771f7d52b5ba41261b0f45.
//
// Solidity: event LaunchProject(uint256 rulesetId, uint256 projectId, string metadata, string memo, address caller)
func (_JBController *JBControllerFilterer) ParseLaunchProject(log types.Log) (*JBControllerLaunchProject, error) {
	event := new(JBControllerLaunchProject)
	if err := _JBController.contract.UnpackLog(event, "LaunchProject", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBControllerLaunchRulesetsIterator is returned from FilterLaunchRulesets and is used to iterate over the raw logs and unpacked data for LaunchRulesets events raised by the JBController contract.
type JBControllerLaunchRulesetsIterator struct {
	Event *JBControllerLaunchRulesets // Event containing the contract specifics and raw log

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
func (it *JBControllerLaunchRulesetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerLaunchRulesets)
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
		it.Event = new(JBControllerLaunchRulesets)
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
func (it *JBControllerLaunchRulesetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerLaunchRulesetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerLaunchRulesets represents a LaunchRulesets event raised by the JBController contract.
type JBControllerLaunchRulesets struct {
	RulesetId *big.Int
	ProjectId *big.Int
	Memo      string
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLaunchRulesets is a free log retrieval operation binding the contract event 0x23164adda696b19106c2eef50ba44461997682bf5aeb9fdb383117fa9952cc75.
//
// Solidity: event LaunchRulesets(uint256 rulesetId, uint256 projectId, string memo, address caller)
func (_JBController *JBControllerFilterer) FilterLaunchRulesets(opts *bind.FilterOpts) (*JBControllerLaunchRulesetsIterator, error) {

	logs, sub, err := _JBController.contract.FilterLogs(opts, "LaunchRulesets")
	if err != nil {
		return nil, err
	}
	return &JBControllerLaunchRulesetsIterator{contract: _JBController.contract, event: "LaunchRulesets", logs: logs, sub: sub}, nil
}

// WatchLaunchRulesets is a free log subscription operation binding the contract event 0x23164adda696b19106c2eef50ba44461997682bf5aeb9fdb383117fa9952cc75.
//
// Solidity: event LaunchRulesets(uint256 rulesetId, uint256 projectId, string memo, address caller)
func (_JBController *JBControllerFilterer) WatchLaunchRulesets(opts *bind.WatchOpts, sink chan<- *JBControllerLaunchRulesets) (event.Subscription, error) {

	logs, sub, err := _JBController.contract.WatchLogs(opts, "LaunchRulesets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerLaunchRulesets)
				if err := _JBController.contract.UnpackLog(event, "LaunchRulesets", log); err != nil {
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

// ParseLaunchRulesets is a log parse operation binding the contract event 0x23164adda696b19106c2eef50ba44461997682bf5aeb9fdb383117fa9952cc75.
//
// Solidity: event LaunchRulesets(uint256 rulesetId, uint256 projectId, string memo, address caller)
func (_JBController *JBControllerFilterer) ParseLaunchRulesets(log types.Log) (*JBControllerLaunchRulesets, error) {
	event := new(JBControllerLaunchRulesets)
	if err := _JBController.contract.UnpackLog(event, "LaunchRulesets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBControllerMigrateControllerIterator is returned from FilterMigrateController and is used to iterate over the raw logs and unpacked data for MigrateController events raised by the JBController contract.
type JBControllerMigrateControllerIterator struct {
	Event *JBControllerMigrateController // Event containing the contract specifics and raw log

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
func (it *JBControllerMigrateControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerMigrateController)
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
		it.Event = new(JBControllerMigrateController)
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
func (it *JBControllerMigrateControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerMigrateControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerMigrateController represents a MigrateController event raised by the JBController contract.
type JBControllerMigrateController struct {
	ProjectId *big.Int
	To        common.Address
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMigrateController is a free log retrieval operation binding the contract event 0x6d2afe69329448f0491521e53921343cc696596a49f64a5502a27935e6a482f9.
//
// Solidity: event MigrateController(uint256 indexed projectId, address to, address caller)
func (_JBController *JBControllerFilterer) FilterMigrateController(opts *bind.FilterOpts, projectId []*big.Int) (*JBControllerMigrateControllerIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.FilterLogs(opts, "MigrateController", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBControllerMigrateControllerIterator{contract: _JBController.contract, event: "MigrateController", logs: logs, sub: sub}, nil
}

// WatchMigrateController is a free log subscription operation binding the contract event 0x6d2afe69329448f0491521e53921343cc696596a49f64a5502a27935e6a482f9.
//
// Solidity: event MigrateController(uint256 indexed projectId, address to, address caller)
func (_JBController *JBControllerFilterer) WatchMigrateController(opts *bind.WatchOpts, sink chan<- *JBControllerMigrateController, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.WatchLogs(opts, "MigrateController", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerMigrateController)
				if err := _JBController.contract.UnpackLog(event, "MigrateController", log); err != nil {
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

// ParseMigrateController is a log parse operation binding the contract event 0x6d2afe69329448f0491521e53921343cc696596a49f64a5502a27935e6a482f9.
//
// Solidity: event MigrateController(uint256 indexed projectId, address to, address caller)
func (_JBController *JBControllerFilterer) ParseMigrateController(log types.Log) (*JBControllerMigrateController, error) {
	event := new(JBControllerMigrateController)
	if err := _JBController.contract.UnpackLog(event, "MigrateController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBControllerMintTokensIterator is returned from FilterMintTokens and is used to iterate over the raw logs and unpacked data for MintTokens events raised by the JBController contract.
type JBControllerMintTokensIterator struct {
	Event *JBControllerMintTokens // Event containing the contract specifics and raw log

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
func (it *JBControllerMintTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerMintTokens)
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
		it.Event = new(JBControllerMintTokens)
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
func (it *JBControllerMintTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerMintTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerMintTokens represents a MintTokens event raised by the JBController contract.
type JBControllerMintTokens struct {
	Beneficiary           common.Address
	ProjectId             *big.Int
	TokenCount            *big.Int
	BeneficiaryTokenCount *big.Int
	Memo                  string
	ReservedRate          *big.Int
	Caller                common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMintTokens is a free log retrieval operation binding the contract event 0xe6fee9c572244c0c2238c3112ac12d411750a7ee00eeebd32521c3e5a666c14b.
//
// Solidity: event MintTokens(address indexed beneficiary, uint256 indexed projectId, uint256 tokenCount, uint256 beneficiaryTokenCount, string memo, uint256 reservedRate, address caller)
func (_JBController *JBControllerFilterer) FilterMintTokens(opts *bind.FilterOpts, beneficiary []common.Address, projectId []*big.Int) (*JBControllerMintTokensIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.FilterLogs(opts, "MintTokens", beneficiaryRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBControllerMintTokensIterator{contract: _JBController.contract, event: "MintTokens", logs: logs, sub: sub}, nil
}

// WatchMintTokens is a free log subscription operation binding the contract event 0xe6fee9c572244c0c2238c3112ac12d411750a7ee00eeebd32521c3e5a666c14b.
//
// Solidity: event MintTokens(address indexed beneficiary, uint256 indexed projectId, uint256 tokenCount, uint256 beneficiaryTokenCount, string memo, uint256 reservedRate, address caller)
func (_JBController *JBControllerFilterer) WatchMintTokens(opts *bind.WatchOpts, sink chan<- *JBControllerMintTokens, beneficiary []common.Address, projectId []*big.Int) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.WatchLogs(opts, "MintTokens", beneficiaryRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerMintTokens)
				if err := _JBController.contract.UnpackLog(event, "MintTokens", log); err != nil {
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

// ParseMintTokens is a log parse operation binding the contract event 0xe6fee9c572244c0c2238c3112ac12d411750a7ee00eeebd32521c3e5a666c14b.
//
// Solidity: event MintTokens(address indexed beneficiary, uint256 indexed projectId, uint256 tokenCount, uint256 beneficiaryTokenCount, string memo, uint256 reservedRate, address caller)
func (_JBController *JBControllerFilterer) ParseMintTokens(log types.Log) (*JBControllerMintTokens, error) {
	event := new(JBControllerMintTokens)
	if err := _JBController.contract.UnpackLog(event, "MintTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBControllerPrepMigrationIterator is returned from FilterPrepMigration and is used to iterate over the raw logs and unpacked data for PrepMigration events raised by the JBController contract.
type JBControllerPrepMigrationIterator struct {
	Event *JBControllerPrepMigration // Event containing the contract specifics and raw log

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
func (it *JBControllerPrepMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerPrepMigration)
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
		it.Event = new(JBControllerPrepMigration)
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
func (it *JBControllerPrepMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerPrepMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerPrepMigration represents a PrepMigration event raised by the JBController contract.
type JBControllerPrepMigration struct {
	ProjectId *big.Int
	From      common.Address
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrepMigration is a free log retrieval operation binding the contract event 0xf389f4f5d01fe4903d6a7a63b8790b7bf80d374b6afed808c03795c3b323d4d3.
//
// Solidity: event PrepMigration(uint256 indexed projectId, address from, address caller)
func (_JBController *JBControllerFilterer) FilterPrepMigration(opts *bind.FilterOpts, projectId []*big.Int) (*JBControllerPrepMigrationIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.FilterLogs(opts, "PrepMigration", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBControllerPrepMigrationIterator{contract: _JBController.contract, event: "PrepMigration", logs: logs, sub: sub}, nil
}

// WatchPrepMigration is a free log subscription operation binding the contract event 0xf389f4f5d01fe4903d6a7a63b8790b7bf80d374b6afed808c03795c3b323d4d3.
//
// Solidity: event PrepMigration(uint256 indexed projectId, address from, address caller)
func (_JBController *JBControllerFilterer) WatchPrepMigration(opts *bind.WatchOpts, sink chan<- *JBControllerPrepMigration, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.WatchLogs(opts, "PrepMigration", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerPrepMigration)
				if err := _JBController.contract.UnpackLog(event, "PrepMigration", log); err != nil {
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

// ParsePrepMigration is a log parse operation binding the contract event 0xf389f4f5d01fe4903d6a7a63b8790b7bf80d374b6afed808c03795c3b323d4d3.
//
// Solidity: event PrepMigration(uint256 indexed projectId, address from, address caller)
func (_JBController *JBControllerFilterer) ParsePrepMigration(log types.Log) (*JBControllerPrepMigration, error) {
	event := new(JBControllerPrepMigration)
	if err := _JBController.contract.UnpackLog(event, "PrepMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBControllerQueueRulesetsIterator is returned from FilterQueueRulesets and is used to iterate over the raw logs and unpacked data for QueueRulesets events raised by the JBController contract.
type JBControllerQueueRulesetsIterator struct {
	Event *JBControllerQueueRulesets // Event containing the contract specifics and raw log

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
func (it *JBControllerQueueRulesetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerQueueRulesets)
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
		it.Event = new(JBControllerQueueRulesets)
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
func (it *JBControllerQueueRulesetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerQueueRulesetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerQueueRulesets represents a QueueRulesets event raised by the JBController contract.
type JBControllerQueueRulesets struct {
	RulesetId *big.Int
	ProjectId *big.Int
	Memo      string
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQueueRulesets is a free log retrieval operation binding the contract event 0x24d02d612b06648bfa2000859f3de7e6f336139eaf5877c24b0d213206252862.
//
// Solidity: event QueueRulesets(uint256 rulesetId, uint256 projectId, string memo, address caller)
func (_JBController *JBControllerFilterer) FilterQueueRulesets(opts *bind.FilterOpts) (*JBControllerQueueRulesetsIterator, error) {

	logs, sub, err := _JBController.contract.FilterLogs(opts, "QueueRulesets")
	if err != nil {
		return nil, err
	}
	return &JBControllerQueueRulesetsIterator{contract: _JBController.contract, event: "QueueRulesets", logs: logs, sub: sub}, nil
}

// WatchQueueRulesets is a free log subscription operation binding the contract event 0x24d02d612b06648bfa2000859f3de7e6f336139eaf5877c24b0d213206252862.
//
// Solidity: event QueueRulesets(uint256 rulesetId, uint256 projectId, string memo, address caller)
func (_JBController *JBControllerFilterer) WatchQueueRulesets(opts *bind.WatchOpts, sink chan<- *JBControllerQueueRulesets) (event.Subscription, error) {

	logs, sub, err := _JBController.contract.WatchLogs(opts, "QueueRulesets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerQueueRulesets)
				if err := _JBController.contract.UnpackLog(event, "QueueRulesets", log); err != nil {
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

// ParseQueueRulesets is a log parse operation binding the contract event 0x24d02d612b06648bfa2000859f3de7e6f336139eaf5877c24b0d213206252862.
//
// Solidity: event QueueRulesets(uint256 rulesetId, uint256 projectId, string memo, address caller)
func (_JBController *JBControllerFilterer) ParseQueueRulesets(log types.Log) (*JBControllerQueueRulesets, error) {
	event := new(JBControllerQueueRulesets)
	if err := _JBController.contract.UnpackLog(event, "QueueRulesets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBControllerReservedDistributionRevertedIterator is returned from FilterReservedDistributionReverted and is used to iterate over the raw logs and unpacked data for ReservedDistributionReverted events raised by the JBController contract.
type JBControllerReservedDistributionRevertedIterator struct {
	Event *JBControllerReservedDistributionReverted // Event containing the contract specifics and raw log

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
func (it *JBControllerReservedDistributionRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerReservedDistributionReverted)
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
		it.Event = new(JBControllerReservedDistributionReverted)
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
func (it *JBControllerReservedDistributionRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerReservedDistributionRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerReservedDistributionReverted represents a ReservedDistributionReverted event raised by the JBController contract.
type JBControllerReservedDistributionReverted struct {
	ProjectId *big.Int
	Split     JBSplit
	Amount    *big.Int
	Reason    []byte
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReservedDistributionReverted is a free log retrieval operation binding the contract event 0xeb1c69f596057b3b93b14d290bc52c6ddc7b7fa47184a8be88a30c656eac83e8.
//
// Solidity: event ReservedDistributionReverted(uint256 indexed projectId, (bool,uint256,uint256,address,uint256,address) split, uint256 amount, bytes reason, address caller)
func (_JBController *JBControllerFilterer) FilterReservedDistributionReverted(opts *bind.FilterOpts, projectId []*big.Int) (*JBControllerReservedDistributionRevertedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.FilterLogs(opts, "ReservedDistributionReverted", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBControllerReservedDistributionRevertedIterator{contract: _JBController.contract, event: "ReservedDistributionReverted", logs: logs, sub: sub}, nil
}

// WatchReservedDistributionReverted is a free log subscription operation binding the contract event 0xeb1c69f596057b3b93b14d290bc52c6ddc7b7fa47184a8be88a30c656eac83e8.
//
// Solidity: event ReservedDistributionReverted(uint256 indexed projectId, (bool,uint256,uint256,address,uint256,address) split, uint256 amount, bytes reason, address caller)
func (_JBController *JBControllerFilterer) WatchReservedDistributionReverted(opts *bind.WatchOpts, sink chan<- *JBControllerReservedDistributionReverted, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.WatchLogs(opts, "ReservedDistributionReverted", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerReservedDistributionReverted)
				if err := _JBController.contract.UnpackLog(event, "ReservedDistributionReverted", log); err != nil {
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

// ParseReservedDistributionReverted is a log parse operation binding the contract event 0xeb1c69f596057b3b93b14d290bc52c6ddc7b7fa47184a8be88a30c656eac83e8.
//
// Solidity: event ReservedDistributionReverted(uint256 indexed projectId, (bool,uint256,uint256,address,uint256,address) split, uint256 amount, bytes reason, address caller)
func (_JBController *JBControllerFilterer) ParseReservedDistributionReverted(log types.Log) (*JBControllerReservedDistributionReverted, error) {
	event := new(JBControllerReservedDistributionReverted)
	if err := _JBController.contract.UnpackLog(event, "ReservedDistributionReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBControllerSendReservedTokensToSplitIterator is returned from FilterSendReservedTokensToSplit and is used to iterate over the raw logs and unpacked data for SendReservedTokensToSplit events raised by the JBController contract.
type JBControllerSendReservedTokensToSplitIterator struct {
	Event *JBControllerSendReservedTokensToSplit // Event containing the contract specifics and raw log

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
func (it *JBControllerSendReservedTokensToSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerSendReservedTokensToSplit)
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
		it.Event = new(JBControllerSendReservedTokensToSplit)
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
func (it *JBControllerSendReservedTokensToSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerSendReservedTokensToSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerSendReservedTokensToSplit represents a SendReservedTokensToSplit event raised by the JBController contract.
type JBControllerSendReservedTokensToSplit struct {
	ProjectId  *big.Int
	RulesetId  *big.Int
	Group      *big.Int
	Split      JBSplit
	TokenCount *big.Int
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendReservedTokensToSplit is a free log retrieval operation binding the contract event 0x5b9f38b2a849d80d4925a023ac038581d16642209624dc7b8353077d36a76824.
//
// Solidity: event SendReservedTokensToSplit(uint256 indexed projectId, uint256 indexed rulesetId, uint256 indexed group, (bool,uint256,uint256,address,uint256,address) split, uint256 tokenCount, address caller)
func (_JBController *JBControllerFilterer) FilterSendReservedTokensToSplit(opts *bind.FilterOpts, projectId []*big.Int, rulesetId []*big.Int, group []*big.Int) (*JBControllerSendReservedTokensToSplitIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var rulesetIdRule []interface{}
	for _, rulesetIdItem := range rulesetId {
		rulesetIdRule = append(rulesetIdRule, rulesetIdItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _JBController.contract.FilterLogs(opts, "SendReservedTokensToSplit", projectIdRule, rulesetIdRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &JBControllerSendReservedTokensToSplitIterator{contract: _JBController.contract, event: "SendReservedTokensToSplit", logs: logs, sub: sub}, nil
}

// WatchSendReservedTokensToSplit is a free log subscription operation binding the contract event 0x5b9f38b2a849d80d4925a023ac038581d16642209624dc7b8353077d36a76824.
//
// Solidity: event SendReservedTokensToSplit(uint256 indexed projectId, uint256 indexed rulesetId, uint256 indexed group, (bool,uint256,uint256,address,uint256,address) split, uint256 tokenCount, address caller)
func (_JBController *JBControllerFilterer) WatchSendReservedTokensToSplit(opts *bind.WatchOpts, sink chan<- *JBControllerSendReservedTokensToSplit, projectId []*big.Int, rulesetId []*big.Int, group []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var rulesetIdRule []interface{}
	for _, rulesetIdItem := range rulesetId {
		rulesetIdRule = append(rulesetIdRule, rulesetIdItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _JBController.contract.WatchLogs(opts, "SendReservedTokensToSplit", projectIdRule, rulesetIdRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerSendReservedTokensToSplit)
				if err := _JBController.contract.UnpackLog(event, "SendReservedTokensToSplit", log); err != nil {
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

// ParseSendReservedTokensToSplit is a log parse operation binding the contract event 0x5b9f38b2a849d80d4925a023ac038581d16642209624dc7b8353077d36a76824.
//
// Solidity: event SendReservedTokensToSplit(uint256 indexed projectId, uint256 indexed rulesetId, uint256 indexed group, (bool,uint256,uint256,address,uint256,address) split, uint256 tokenCount, address caller)
func (_JBController *JBControllerFilterer) ParseSendReservedTokensToSplit(log types.Log) (*JBControllerSendReservedTokensToSplit, error) {
	event := new(JBControllerSendReservedTokensToSplit)
	if err := _JBController.contract.UnpackLog(event, "SendReservedTokensToSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBControllerSendReservedTokensToSplitsIterator is returned from FilterSendReservedTokensToSplits and is used to iterate over the raw logs and unpacked data for SendReservedTokensToSplits events raised by the JBController contract.
type JBControllerSendReservedTokensToSplitsIterator struct {
	Event *JBControllerSendReservedTokensToSplits // Event containing the contract specifics and raw log

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
func (it *JBControllerSendReservedTokensToSplitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerSendReservedTokensToSplits)
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
		it.Event = new(JBControllerSendReservedTokensToSplits)
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
func (it *JBControllerSendReservedTokensToSplitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerSendReservedTokensToSplitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerSendReservedTokensToSplits represents a SendReservedTokensToSplits event raised by the JBController contract.
type JBControllerSendReservedTokensToSplits struct {
	RulesetId             *big.Int
	RulesetCycleNumber    *big.Int
	ProjectId             *big.Int
	Beneficiary           common.Address
	TokenCount            *big.Int
	BeneficiaryTokenCount *big.Int
	Caller                common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSendReservedTokensToSplits is a free log retrieval operation binding the contract event 0x32411cac611c47cfe73afc187645c9cf3aec828d5f91780138d8421378fc0edb.
//
// Solidity: event SendReservedTokensToSplits(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address beneficiary, uint256 tokenCount, uint256 beneficiaryTokenCount, address caller)
func (_JBController *JBControllerFilterer) FilterSendReservedTokensToSplits(opts *bind.FilterOpts, rulesetId []*big.Int, rulesetCycleNumber []*big.Int, projectId []*big.Int) (*JBControllerSendReservedTokensToSplitsIterator, error) {

	var rulesetIdRule []interface{}
	for _, rulesetIdItem := range rulesetId {
		rulesetIdRule = append(rulesetIdRule, rulesetIdItem)
	}
	var rulesetCycleNumberRule []interface{}
	for _, rulesetCycleNumberItem := range rulesetCycleNumber {
		rulesetCycleNumberRule = append(rulesetCycleNumberRule, rulesetCycleNumberItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.FilterLogs(opts, "SendReservedTokensToSplits", rulesetIdRule, rulesetCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBControllerSendReservedTokensToSplitsIterator{contract: _JBController.contract, event: "SendReservedTokensToSplits", logs: logs, sub: sub}, nil
}

// WatchSendReservedTokensToSplits is a free log subscription operation binding the contract event 0x32411cac611c47cfe73afc187645c9cf3aec828d5f91780138d8421378fc0edb.
//
// Solidity: event SendReservedTokensToSplits(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address beneficiary, uint256 tokenCount, uint256 beneficiaryTokenCount, address caller)
func (_JBController *JBControllerFilterer) WatchSendReservedTokensToSplits(opts *bind.WatchOpts, sink chan<- *JBControllerSendReservedTokensToSplits, rulesetId []*big.Int, rulesetCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

	var rulesetIdRule []interface{}
	for _, rulesetIdItem := range rulesetId {
		rulesetIdRule = append(rulesetIdRule, rulesetIdItem)
	}
	var rulesetCycleNumberRule []interface{}
	for _, rulesetCycleNumberItem := range rulesetCycleNumber {
		rulesetCycleNumberRule = append(rulesetCycleNumberRule, rulesetCycleNumberItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.WatchLogs(opts, "SendReservedTokensToSplits", rulesetIdRule, rulesetCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerSendReservedTokensToSplits)
				if err := _JBController.contract.UnpackLog(event, "SendReservedTokensToSplits", log); err != nil {
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

// ParseSendReservedTokensToSplits is a log parse operation binding the contract event 0x32411cac611c47cfe73afc187645c9cf3aec828d5f91780138d8421378fc0edb.
//
// Solidity: event SendReservedTokensToSplits(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address beneficiary, uint256 tokenCount, uint256 beneficiaryTokenCount, address caller)
func (_JBController *JBControllerFilterer) ParseSendReservedTokensToSplits(log types.Log) (*JBControllerSendReservedTokensToSplits, error) {
	event := new(JBControllerSendReservedTokensToSplits)
	if err := _JBController.contract.UnpackLog(event, "SendReservedTokensToSplits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBControllerSetMetadataIterator is returned from FilterSetMetadata and is used to iterate over the raw logs and unpacked data for SetMetadata events raised by the JBController contract.
type JBControllerSetMetadataIterator struct {
	Event *JBControllerSetMetadata // Event containing the contract specifics and raw log

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
func (it *JBControllerSetMetadataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBControllerSetMetadata)
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
		it.Event = new(JBControllerSetMetadata)
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
func (it *JBControllerSetMetadataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBControllerSetMetadataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBControllerSetMetadata represents a SetMetadata event raised by the JBController contract.
type JBControllerSetMetadata struct {
	ProjectId *big.Int
	Metadata  string
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetMetadata is a free log retrieval operation binding the contract event 0x76f445b3718cf71aa0c056105beab456eba31f72f5cc5a60fde060fd14ea1684.
//
// Solidity: event SetMetadata(uint256 indexed projectId, string metadata, address caller)
func (_JBController *JBControllerFilterer) FilterSetMetadata(opts *bind.FilterOpts, projectId []*big.Int) (*JBControllerSetMetadataIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.FilterLogs(opts, "SetMetadata", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBControllerSetMetadataIterator{contract: _JBController.contract, event: "SetMetadata", logs: logs, sub: sub}, nil
}

// WatchSetMetadata is a free log subscription operation binding the contract event 0x76f445b3718cf71aa0c056105beab456eba31f72f5cc5a60fde060fd14ea1684.
//
// Solidity: event SetMetadata(uint256 indexed projectId, string metadata, address caller)
func (_JBController *JBControllerFilterer) WatchSetMetadata(opts *bind.WatchOpts, sink chan<- *JBControllerSetMetadata, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBController.contract.WatchLogs(opts, "SetMetadata", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBControllerSetMetadata)
				if err := _JBController.contract.UnpackLog(event, "SetMetadata", log); err != nil {
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

// ParseSetMetadata is a log parse operation binding the contract event 0x76f445b3718cf71aa0c056105beab456eba31f72f5cc5a60fde060fd14ea1684.
//
// Solidity: event SetMetadata(uint256 indexed projectId, string metadata, address caller)
func (_JBController *JBControllerFilterer) ParseSetMetadata(log types.Log) (*JBControllerSetMetadata, error) {
	event := new(JBControllerSetMetadata)
	if err := _JBController.contract.UnpackLog(event, "SetMetadata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
