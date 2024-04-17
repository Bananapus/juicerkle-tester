// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package JBMultiTerminal

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

// JBAccountingContext is an auto generated low-level Go binding around an user-defined struct.
type JBAccountingContext struct {
	Token    common.Address
	Decimals uint8
	Currency uint32
}

// JBAfterPayRecordedContext is an auto generated low-level Go binding around an user-defined struct.
type JBAfterPayRecordedContext struct {
	Payer             common.Address
	ProjectId         *big.Int
	RulesetId         *big.Int
	Amount            JBTokenAmount
	ForwardedAmount   JBTokenAmount
	Weight            *big.Int
	ProjectTokenCount *big.Int
	Beneficiary       common.Address
	HookMetadata      []byte
	PayerMetadata     []byte
}

// JBAfterRedeemRecordedContext is an auto generated low-level Go binding around an user-defined struct.
type JBAfterRedeemRecordedContext struct {
	Holder           common.Address
	ProjectId        *big.Int
	RulesetId        *big.Int
	RedeemCount      *big.Int
	ReclaimedAmount  JBTokenAmount
	ForwardedAmount  JBTokenAmount
	RedemptionRate   *big.Int
	Beneficiary      common.Address
	HookMetadata     []byte
	RedeemerMetadata []byte
}

// JBFee is an auto generated low-level Go binding around an user-defined struct.
type JBFee struct {
	Amount          *big.Int
	Beneficiary     common.Address
	UnlockTimestamp *big.Int
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

// JBTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type JBTokenAmount struct {
	Token    common.Address
	Value    *big.Int
	Decimals *big.Int
	Currency *big.Int
}

// JBMultiTerminalMetaData contains all meta data concerning the JBMultiTerminal contract.
var JBMultiTerminalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"permissions\",\"type\":\"address\",\"internalType\":\"contractIJBPermissions\"},{\"name\":\"projects\",\"type\":\"address\",\"internalType\":\"contractIJBProjects\"},{\"name\":\"directory\",\"type\":\"address\",\"internalType\":\"contractIJBDirectory\"},{\"name\":\"splits\",\"type\":\"address\",\"internalType\":\"contractIJBSplits\"},{\"name\":\"store\",\"type\":\"address\",\"internalType\":\"contractIJBTerminalStore\"},{\"name\":\"feelessAddresses\",\"type\":\"address\",\"internalType\":\"contractIJBFeelessAddresses\"},{\"name\":\"permit2\",\"type\":\"address\",\"internalType\":\"contractIPermit2\"},{\"name\":\"trustedForwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DIRECTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FEELESS_ADDRESSES\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBFeelessAddresses\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PERMISSIONS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBPermissions\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PERMIT2\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermit2\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROJECTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBProjects\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SPLITS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBSplits\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STORE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBTerminalStore\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountingContextForTokenOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structJBAccountingContext\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"currency\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountingContextsOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structJBAccountingContext[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"currency\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addAccountingContextsFor\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addToBalanceOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shouldReturnHeldFees\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"currentSurplusOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executePayout\",\"inputs\":[{\"name\":\"split\",\"type\":\"tuple\",\"internalType\":\"structJBSplit\",\"components\":[{\"name\":\"preferAddToBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"percent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"contractIJBSplitHook\"}]},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalMessageSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"netPayoutAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeProcessFee\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTerminal\",\"type\":\"address\",\"internalType\":\"contractIJBTerminal\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"heldFeesOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structJBFee[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"unlockTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateBalanceOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"contractIJBTerminal\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pay\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minReturnedTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"beneficiaryTokenCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"processHeldFeesOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemTokensOf\",\"inputs\":[{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenToReclaim\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"redeemCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minTokensReclaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"reclaimAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendPayoutsOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minTokensPaidOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amountPaidOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"trustedForwarder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"useAllowanceOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minTokensPaidOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"amountPaidOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddToBalance\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"unlockedFees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeReverted\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeProjectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HoldFee\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HookAfterRecordPay\",\"inputs\":[{\"name\":\"hook\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIJBPayHook\"},{\"name\":\"context\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structJBAfterPayRecordedContext\",\"components\":[{\"name\":\"payer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rulesetId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"tuple\",\"internalType\":\"structJBTokenAmount\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"forwardedAmount\",\"type\":\"tuple\",\"internalType\":\"structJBTokenAmount\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectTokenCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hookMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"payerMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"specificationAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HookAfterRecordRedeem\",\"inputs\":[{\"name\":\"hook\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIJBRedeemHook\"},{\"name\":\"context\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structJBAfterRedeemRecordedContext\",\"components\":[{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rulesetId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redeemCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reclaimedAmount\",\"type\":\"tuple\",\"internalType\":\"structJBTokenAmount\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"forwardedAmount\",\"type\":\"tuple\",\"internalType\":\"structJBTokenAmount\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"redemptionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"hookMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"redeemerMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"specificationAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MigrateTerminal\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIJBTerminal\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Pay\",\"inputs\":[{\"name\":\"rulesetId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rulesetCycleNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"payer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"beneficiaryTokenCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayoutReverted\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"split\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structJBSplit\",\"components\":[{\"name\":\"preferAddToBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"percent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"contractIJBSplitHook\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProcessFee\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"wasHeld\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedeemTokens\",\"inputs\":[{\"name\":\"rulesetId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rulesetCycleNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"redemptionRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reclaimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReturnHeldFees\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"returnedFees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendPayoutToSplit\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rulesetId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"group\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"split\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structJBSplit\",\"components\":[{\"name\":\"preferAddToBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"percent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hook\",\"type\":\"address\",\"internalType\":\"contractIJBSplitHook\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"netAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendPayouts\",\"inputs\":[{\"name\":\"rulesetId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rulesetCycleNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountPaidOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"beneficiaryDistributionAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAccountingContext\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"context\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structJBAccountingContext\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"currency\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UseAllowance\",\"inputs\":[{\"name\":\"rulesetId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rulesetCycleNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountPaidOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"netAmountPaidOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ACCOUNTING_CONTEXT_ALREADY_SET\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INADEQUATE_PAYOUT_AMOUNT\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INADEQUATE_RECLAIM_AMOUNT\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NO_MSG_VALUE_ALLOWED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OVERFLOW_ALERT\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PERMIT_ALLOWANCE_NOT_ENOUGH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PRBMath_MulDiv_Overflow\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TERMINAL_TOKENS_INCOMPATIBLE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TOKEN_NOT_ACCEPTED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UNAUTHORIZED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UNDER_MIN_RETURNED_TOKENS\",\"inputs\":[]}]",
}

// JBMultiTerminalABI is the input ABI used to generate the binding from.
// Deprecated: Use JBMultiTerminalMetaData.ABI instead.
var JBMultiTerminalABI = JBMultiTerminalMetaData.ABI

// JBMultiTerminal is an auto generated Go binding around an Ethereum contract.
type JBMultiTerminal struct {
	JBMultiTerminalCaller     // Read-only binding to the contract
	JBMultiTerminalTransactor // Write-only binding to the contract
	JBMultiTerminalFilterer   // Log filterer for contract events
}

// JBMultiTerminalCaller is an auto generated read-only Go binding around an Ethereum contract.
type JBMultiTerminalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBMultiTerminalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JBMultiTerminalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBMultiTerminalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JBMultiTerminalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBMultiTerminalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JBMultiTerminalSession struct {
	Contract     *JBMultiTerminal  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JBMultiTerminalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JBMultiTerminalCallerSession struct {
	Contract *JBMultiTerminalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// JBMultiTerminalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JBMultiTerminalTransactorSession struct {
	Contract     *JBMultiTerminalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// JBMultiTerminalRaw is an auto generated low-level Go binding around an Ethereum contract.
type JBMultiTerminalRaw struct {
	Contract *JBMultiTerminal // Generic contract binding to access the raw methods on
}

// JBMultiTerminalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JBMultiTerminalCallerRaw struct {
	Contract *JBMultiTerminalCaller // Generic read-only contract binding to access the raw methods on
}

// JBMultiTerminalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JBMultiTerminalTransactorRaw struct {
	Contract *JBMultiTerminalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJBMultiTerminal creates a new instance of JBMultiTerminal, bound to a specific deployed contract.
func NewJBMultiTerminal(address common.Address, backend bind.ContractBackend) (*JBMultiTerminal, error) {
	contract, err := bindJBMultiTerminal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminal{JBMultiTerminalCaller: JBMultiTerminalCaller{contract: contract}, JBMultiTerminalTransactor: JBMultiTerminalTransactor{contract: contract}, JBMultiTerminalFilterer: JBMultiTerminalFilterer{contract: contract}}, nil
}

// NewJBMultiTerminalCaller creates a new read-only instance of JBMultiTerminal, bound to a specific deployed contract.
func NewJBMultiTerminalCaller(address common.Address, caller bind.ContractCaller) (*JBMultiTerminalCaller, error) {
	contract, err := bindJBMultiTerminal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalCaller{contract: contract}, nil
}

// NewJBMultiTerminalTransactor creates a new write-only instance of JBMultiTerminal, bound to a specific deployed contract.
func NewJBMultiTerminalTransactor(address common.Address, transactor bind.ContractTransactor) (*JBMultiTerminalTransactor, error) {
	contract, err := bindJBMultiTerminal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalTransactor{contract: contract}, nil
}

// NewJBMultiTerminalFilterer creates a new log filterer instance of JBMultiTerminal, bound to a specific deployed contract.
func NewJBMultiTerminalFilterer(address common.Address, filterer bind.ContractFilterer) (*JBMultiTerminalFilterer, error) {
	contract, err := bindJBMultiTerminal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalFilterer{contract: contract}, nil
}

// bindJBMultiTerminal binds a generic wrapper to an already deployed contract.
func bindJBMultiTerminal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JBMultiTerminalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JBMultiTerminal *JBMultiTerminalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JBMultiTerminal.Contract.JBMultiTerminalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JBMultiTerminal *JBMultiTerminalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.JBMultiTerminalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JBMultiTerminal *JBMultiTerminalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.JBMultiTerminalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JBMultiTerminal *JBMultiTerminalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JBMultiTerminal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JBMultiTerminal *JBMultiTerminalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JBMultiTerminal *JBMultiTerminalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.contract.Transact(opts, method, params...)
}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCaller) DIRECTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "DIRECTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalSession) DIRECTORY() (common.Address, error) {
	return _JBMultiTerminal.Contract.DIRECTORY(&_JBMultiTerminal.CallOpts)
}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) DIRECTORY() (common.Address, error) {
	return _JBMultiTerminal.Contract.DIRECTORY(&_JBMultiTerminal.CallOpts)
}

// FEE is a free data retrieval call binding the contract method 0xc57981b5.
//
// Solidity: function FEE() view returns(uint256)
func (_JBMultiTerminal *JBMultiTerminalCaller) FEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEE is a free data retrieval call binding the contract method 0xc57981b5.
//
// Solidity: function FEE() view returns(uint256)
func (_JBMultiTerminal *JBMultiTerminalSession) FEE() (*big.Int, error) {
	return _JBMultiTerminal.Contract.FEE(&_JBMultiTerminal.CallOpts)
}

// FEE is a free data retrieval call binding the contract method 0xc57981b5.
//
// Solidity: function FEE() view returns(uint256)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) FEE() (*big.Int, error) {
	return _JBMultiTerminal.Contract.FEE(&_JBMultiTerminal.CallOpts)
}

// FEELESSADDRESSES is a free data retrieval call binding the contract method 0x659a2047.
//
// Solidity: function FEELESS_ADDRESSES() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCaller) FEELESSADDRESSES(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "FEELESS_ADDRESSES")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FEELESSADDRESSES is a free data retrieval call binding the contract method 0x659a2047.
//
// Solidity: function FEELESS_ADDRESSES() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalSession) FEELESSADDRESSES() (common.Address, error) {
	return _JBMultiTerminal.Contract.FEELESSADDRESSES(&_JBMultiTerminal.CallOpts)
}

// FEELESSADDRESSES is a free data retrieval call binding the contract method 0x659a2047.
//
// Solidity: function FEELESS_ADDRESSES() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) FEELESSADDRESSES() (common.Address, error) {
	return _JBMultiTerminal.Contract.FEELESSADDRESSES(&_JBMultiTerminal.CallOpts)
}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCaller) PERMISSIONS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "PERMISSIONS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalSession) PERMISSIONS() (common.Address, error) {
	return _JBMultiTerminal.Contract.PERMISSIONS(&_JBMultiTerminal.CallOpts)
}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) PERMISSIONS() (common.Address, error) {
	return _JBMultiTerminal.Contract.PERMISSIONS(&_JBMultiTerminal.CallOpts)
}

// PERMIT2 is a free data retrieval call binding the contract method 0x6afdd850.
//
// Solidity: function PERMIT2() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCaller) PERMIT2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "PERMIT2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PERMIT2 is a free data retrieval call binding the contract method 0x6afdd850.
//
// Solidity: function PERMIT2() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalSession) PERMIT2() (common.Address, error) {
	return _JBMultiTerminal.Contract.PERMIT2(&_JBMultiTerminal.CallOpts)
}

// PERMIT2 is a free data retrieval call binding the contract method 0x6afdd850.
//
// Solidity: function PERMIT2() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) PERMIT2() (common.Address, error) {
	return _JBMultiTerminal.Contract.PERMIT2(&_JBMultiTerminal.CallOpts)
}

// PROJECTS is a free data retrieval call binding the contract method 0x293c4999.
//
// Solidity: function PROJECTS() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCaller) PROJECTS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "PROJECTS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROJECTS is a free data retrieval call binding the contract method 0x293c4999.
//
// Solidity: function PROJECTS() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalSession) PROJECTS() (common.Address, error) {
	return _JBMultiTerminal.Contract.PROJECTS(&_JBMultiTerminal.CallOpts)
}

// PROJECTS is a free data retrieval call binding the contract method 0x293c4999.
//
// Solidity: function PROJECTS() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) PROJECTS() (common.Address, error) {
	return _JBMultiTerminal.Contract.PROJECTS(&_JBMultiTerminal.CallOpts)
}

// SPLITS is a free data retrieval call binding the contract method 0x1f47ce69.
//
// Solidity: function SPLITS() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCaller) SPLITS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "SPLITS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPLITS is a free data retrieval call binding the contract method 0x1f47ce69.
//
// Solidity: function SPLITS() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalSession) SPLITS() (common.Address, error) {
	return _JBMultiTerminal.Contract.SPLITS(&_JBMultiTerminal.CallOpts)
}

// SPLITS is a free data retrieval call binding the contract method 0x1f47ce69.
//
// Solidity: function SPLITS() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) SPLITS() (common.Address, error) {
	return _JBMultiTerminal.Contract.SPLITS(&_JBMultiTerminal.CallOpts)
}

// STORE is a free data retrieval call binding the contract method 0x507f1465.
//
// Solidity: function STORE() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCaller) STORE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "STORE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STORE is a free data retrieval call binding the contract method 0x507f1465.
//
// Solidity: function STORE() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalSession) STORE() (common.Address, error) {
	return _JBMultiTerminal.Contract.STORE(&_JBMultiTerminal.CallOpts)
}

// STORE is a free data retrieval call binding the contract method 0x507f1465.
//
// Solidity: function STORE() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) STORE() (common.Address, error) {
	return _JBMultiTerminal.Contract.STORE(&_JBMultiTerminal.CallOpts)
}

// AccountingContextForTokenOf is a free data retrieval call binding the contract method 0x3a01714f.
//
// Solidity: function accountingContextForTokenOf(uint256 projectId, address token) view returns((address,uint8,uint32))
func (_JBMultiTerminal *JBMultiTerminalCaller) AccountingContextForTokenOf(opts *bind.CallOpts, projectId *big.Int, token common.Address) (JBAccountingContext, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "accountingContextForTokenOf", projectId, token)

	if err != nil {
		return *new(JBAccountingContext), err
	}

	out0 := *abi.ConvertType(out[0], new(JBAccountingContext)).(*JBAccountingContext)

	return out0, err

}

// AccountingContextForTokenOf is a free data retrieval call binding the contract method 0x3a01714f.
//
// Solidity: function accountingContextForTokenOf(uint256 projectId, address token) view returns((address,uint8,uint32))
func (_JBMultiTerminal *JBMultiTerminalSession) AccountingContextForTokenOf(projectId *big.Int, token common.Address) (JBAccountingContext, error) {
	return _JBMultiTerminal.Contract.AccountingContextForTokenOf(&_JBMultiTerminal.CallOpts, projectId, token)
}

// AccountingContextForTokenOf is a free data retrieval call binding the contract method 0x3a01714f.
//
// Solidity: function accountingContextForTokenOf(uint256 projectId, address token) view returns((address,uint8,uint32))
func (_JBMultiTerminal *JBMultiTerminalCallerSession) AccountingContextForTokenOf(projectId *big.Int, token common.Address) (JBAccountingContext, error) {
	return _JBMultiTerminal.Contract.AccountingContextForTokenOf(&_JBMultiTerminal.CallOpts, projectId, token)
}

// AccountingContextsOf is a free data retrieval call binding the contract method 0x515a9293.
//
// Solidity: function accountingContextsOf(uint256 projectId) view returns((address,uint8,uint32)[])
func (_JBMultiTerminal *JBMultiTerminalCaller) AccountingContextsOf(opts *bind.CallOpts, projectId *big.Int) ([]JBAccountingContext, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "accountingContextsOf", projectId)

	if err != nil {
		return *new([]JBAccountingContext), err
	}

	out0 := *abi.ConvertType(out[0], new([]JBAccountingContext)).(*[]JBAccountingContext)

	return out0, err

}

// AccountingContextsOf is a free data retrieval call binding the contract method 0x515a9293.
//
// Solidity: function accountingContextsOf(uint256 projectId) view returns((address,uint8,uint32)[])
func (_JBMultiTerminal *JBMultiTerminalSession) AccountingContextsOf(projectId *big.Int) ([]JBAccountingContext, error) {
	return _JBMultiTerminal.Contract.AccountingContextsOf(&_JBMultiTerminal.CallOpts, projectId)
}

// AccountingContextsOf is a free data retrieval call binding the contract method 0x515a9293.
//
// Solidity: function accountingContextsOf(uint256 projectId) view returns((address,uint8,uint32)[])
func (_JBMultiTerminal *JBMultiTerminalCallerSession) AccountingContextsOf(projectId *big.Int) ([]JBAccountingContext, error) {
	return _JBMultiTerminal.Contract.AccountingContextsOf(&_JBMultiTerminal.CallOpts, projectId)
}

// CurrentSurplusOf is a free data retrieval call binding the contract method 0xcc680127.
//
// Solidity: function currentSurplusOf(uint256 projectId, uint256 decimals, uint256 currency) view returns(uint256)
func (_JBMultiTerminal *JBMultiTerminalCaller) CurrentSurplusOf(opts *bind.CallOpts, projectId *big.Int, decimals *big.Int, currency *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "currentSurplusOf", projectId, decimals, currency)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSurplusOf is a free data retrieval call binding the contract method 0xcc680127.
//
// Solidity: function currentSurplusOf(uint256 projectId, uint256 decimals, uint256 currency) view returns(uint256)
func (_JBMultiTerminal *JBMultiTerminalSession) CurrentSurplusOf(projectId *big.Int, decimals *big.Int, currency *big.Int) (*big.Int, error) {
	return _JBMultiTerminal.Contract.CurrentSurplusOf(&_JBMultiTerminal.CallOpts, projectId, decimals, currency)
}

// CurrentSurplusOf is a free data retrieval call binding the contract method 0xcc680127.
//
// Solidity: function currentSurplusOf(uint256 projectId, uint256 decimals, uint256 currency) view returns(uint256)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) CurrentSurplusOf(projectId *big.Int, decimals *big.Int, currency *big.Int) (*big.Int, error) {
	return _JBMultiTerminal.Contract.CurrentSurplusOf(&_JBMultiTerminal.CallOpts, projectId, decimals, currency)
}

// HeldFeesOf is a free data retrieval call binding the contract method 0x33b30610.
//
// Solidity: function heldFeesOf(uint256 projectId, address token) view returns((uint256,address,uint256)[])
func (_JBMultiTerminal *JBMultiTerminalCaller) HeldFeesOf(opts *bind.CallOpts, projectId *big.Int, token common.Address) ([]JBFee, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "heldFeesOf", projectId, token)

	if err != nil {
		return *new([]JBFee), err
	}

	out0 := *abi.ConvertType(out[0], new([]JBFee)).(*[]JBFee)

	return out0, err

}

// HeldFeesOf is a free data retrieval call binding the contract method 0x33b30610.
//
// Solidity: function heldFeesOf(uint256 projectId, address token) view returns((uint256,address,uint256)[])
func (_JBMultiTerminal *JBMultiTerminalSession) HeldFeesOf(projectId *big.Int, token common.Address) ([]JBFee, error) {
	return _JBMultiTerminal.Contract.HeldFeesOf(&_JBMultiTerminal.CallOpts, projectId, token)
}

// HeldFeesOf is a free data retrieval call binding the contract method 0x33b30610.
//
// Solidity: function heldFeesOf(uint256 projectId, address token) view returns((uint256,address,uint256)[])
func (_JBMultiTerminal *JBMultiTerminalCallerSession) HeldFeesOf(projectId *big.Int, token common.Address) ([]JBFee, error) {
	return _JBMultiTerminal.Contract.HeldFeesOf(&_JBMultiTerminal.CallOpts, projectId, token)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_JBMultiTerminal *JBMultiTerminalCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_JBMultiTerminal *JBMultiTerminalSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _JBMultiTerminal.Contract.IsTrustedForwarder(&_JBMultiTerminal.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _JBMultiTerminal.Contract.IsTrustedForwarder(&_JBMultiTerminal.CallOpts, forwarder)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JBMultiTerminal *JBMultiTerminalCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JBMultiTerminal *JBMultiTerminalSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _JBMultiTerminal.Contract.SupportsInterface(&_JBMultiTerminal.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _JBMultiTerminal.Contract.SupportsInterface(&_JBMultiTerminal.CallOpts, interfaceId)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBMultiTerminal.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalSession) TrustedForwarder() (common.Address, error) {
	return _JBMultiTerminal.Contract.TrustedForwarder(&_JBMultiTerminal.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_JBMultiTerminal *JBMultiTerminalCallerSession) TrustedForwarder() (common.Address, error) {
	return _JBMultiTerminal.Contract.TrustedForwarder(&_JBMultiTerminal.CallOpts)
}

// AddAccountingContextsFor is a paid mutator transaction binding the contract method 0x98ea47fc.
//
// Solidity: function addAccountingContextsFor(uint256 projectId, address[] tokens) returns()
func (_JBMultiTerminal *JBMultiTerminalTransactor) AddAccountingContextsFor(opts *bind.TransactOpts, projectId *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.contract.Transact(opts, "addAccountingContextsFor", projectId, tokens)
}

// AddAccountingContextsFor is a paid mutator transaction binding the contract method 0x98ea47fc.
//
// Solidity: function addAccountingContextsFor(uint256 projectId, address[] tokens) returns()
func (_JBMultiTerminal *JBMultiTerminalSession) AddAccountingContextsFor(projectId *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.AddAccountingContextsFor(&_JBMultiTerminal.TransactOpts, projectId, tokens)
}

// AddAccountingContextsFor is a paid mutator transaction binding the contract method 0x98ea47fc.
//
// Solidity: function addAccountingContextsFor(uint256 projectId, address[] tokens) returns()
func (_JBMultiTerminal *JBMultiTerminalTransactorSession) AddAccountingContextsFor(projectId *big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.AddAccountingContextsFor(&_JBMultiTerminal.TransactOpts, projectId, tokens)
}

// AddToBalanceOf is a paid mutator transaction binding the contract method 0x9e6eec05.
//
// Solidity: function addToBalanceOf(uint256 projectId, address token, uint256 amount, bool shouldReturnHeldFees, string memo, bytes metadata) payable returns()
func (_JBMultiTerminal *JBMultiTerminalTransactor) AddToBalanceOf(opts *bind.TransactOpts, projectId *big.Int, token common.Address, amount *big.Int, shouldReturnHeldFees bool, memo string, metadata []byte) (*types.Transaction, error) {
	return _JBMultiTerminal.contract.Transact(opts, "addToBalanceOf", projectId, token, amount, shouldReturnHeldFees, memo, metadata)
}

// AddToBalanceOf is a paid mutator transaction binding the contract method 0x9e6eec05.
//
// Solidity: function addToBalanceOf(uint256 projectId, address token, uint256 amount, bool shouldReturnHeldFees, string memo, bytes metadata) payable returns()
func (_JBMultiTerminal *JBMultiTerminalSession) AddToBalanceOf(projectId *big.Int, token common.Address, amount *big.Int, shouldReturnHeldFees bool, memo string, metadata []byte) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.AddToBalanceOf(&_JBMultiTerminal.TransactOpts, projectId, token, amount, shouldReturnHeldFees, memo, metadata)
}

// AddToBalanceOf is a paid mutator transaction binding the contract method 0x9e6eec05.
//
// Solidity: function addToBalanceOf(uint256 projectId, address token, uint256 amount, bool shouldReturnHeldFees, string memo, bytes metadata) payable returns()
func (_JBMultiTerminal *JBMultiTerminalTransactorSession) AddToBalanceOf(projectId *big.Int, token common.Address, amount *big.Int, shouldReturnHeldFees bool, memo string, metadata []byte) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.AddToBalanceOf(&_JBMultiTerminal.TransactOpts, projectId, token, amount, shouldReturnHeldFees, memo, metadata)
}

// ExecutePayout is a paid mutator transaction binding the contract method 0x43a89aaf.
//
// Solidity: function executePayout((bool,uint256,uint256,address,uint256,address) split, uint256 projectId, address token, uint256 amount, address originalMessageSender) returns(uint256 netPayoutAmount)
func (_JBMultiTerminal *JBMultiTerminalTransactor) ExecutePayout(opts *bind.TransactOpts, split JBSplit, projectId *big.Int, token common.Address, amount *big.Int, originalMessageSender common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.contract.Transact(opts, "executePayout", split, projectId, token, amount, originalMessageSender)
}

// ExecutePayout is a paid mutator transaction binding the contract method 0x43a89aaf.
//
// Solidity: function executePayout((bool,uint256,uint256,address,uint256,address) split, uint256 projectId, address token, uint256 amount, address originalMessageSender) returns(uint256 netPayoutAmount)
func (_JBMultiTerminal *JBMultiTerminalSession) ExecutePayout(split JBSplit, projectId *big.Int, token common.Address, amount *big.Int, originalMessageSender common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.ExecutePayout(&_JBMultiTerminal.TransactOpts, split, projectId, token, amount, originalMessageSender)
}

// ExecutePayout is a paid mutator transaction binding the contract method 0x43a89aaf.
//
// Solidity: function executePayout((bool,uint256,uint256,address,uint256,address) split, uint256 projectId, address token, uint256 amount, address originalMessageSender) returns(uint256 netPayoutAmount)
func (_JBMultiTerminal *JBMultiTerminalTransactorSession) ExecutePayout(split JBSplit, projectId *big.Int, token common.Address, amount *big.Int, originalMessageSender common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.ExecutePayout(&_JBMultiTerminal.TransactOpts, split, projectId, token, amount, originalMessageSender)
}

// ExecuteProcessFee is a paid mutator transaction binding the contract method 0x5f912e56.
//
// Solidity: function executeProcessFee(uint256 projectId, address token, uint256 amount, address beneficiary, address feeTerminal) returns()
func (_JBMultiTerminal *JBMultiTerminalTransactor) ExecuteProcessFee(opts *bind.TransactOpts, projectId *big.Int, token common.Address, amount *big.Int, beneficiary common.Address, feeTerminal common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.contract.Transact(opts, "executeProcessFee", projectId, token, amount, beneficiary, feeTerminal)
}

// ExecuteProcessFee is a paid mutator transaction binding the contract method 0x5f912e56.
//
// Solidity: function executeProcessFee(uint256 projectId, address token, uint256 amount, address beneficiary, address feeTerminal) returns()
func (_JBMultiTerminal *JBMultiTerminalSession) ExecuteProcessFee(projectId *big.Int, token common.Address, amount *big.Int, beneficiary common.Address, feeTerminal common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.ExecuteProcessFee(&_JBMultiTerminal.TransactOpts, projectId, token, amount, beneficiary, feeTerminal)
}

// ExecuteProcessFee is a paid mutator transaction binding the contract method 0x5f912e56.
//
// Solidity: function executeProcessFee(uint256 projectId, address token, uint256 amount, address beneficiary, address feeTerminal) returns()
func (_JBMultiTerminal *JBMultiTerminalTransactorSession) ExecuteProcessFee(projectId *big.Int, token common.Address, amount *big.Int, beneficiary common.Address, feeTerminal common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.ExecuteProcessFee(&_JBMultiTerminal.TransactOpts, projectId, token, amount, beneficiary, feeTerminal)
}

// MigrateBalanceOf is a paid mutator transaction binding the contract method 0xe28b5411.
//
// Solidity: function migrateBalanceOf(uint256 projectId, address token, address to) returns(uint256 balance)
func (_JBMultiTerminal *JBMultiTerminalTransactor) MigrateBalanceOf(opts *bind.TransactOpts, projectId *big.Int, token common.Address, to common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.contract.Transact(opts, "migrateBalanceOf", projectId, token, to)
}

// MigrateBalanceOf is a paid mutator transaction binding the contract method 0xe28b5411.
//
// Solidity: function migrateBalanceOf(uint256 projectId, address token, address to) returns(uint256 balance)
func (_JBMultiTerminal *JBMultiTerminalSession) MigrateBalanceOf(projectId *big.Int, token common.Address, to common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.MigrateBalanceOf(&_JBMultiTerminal.TransactOpts, projectId, token, to)
}

// MigrateBalanceOf is a paid mutator transaction binding the contract method 0xe28b5411.
//
// Solidity: function migrateBalanceOf(uint256 projectId, address token, address to) returns(uint256 balance)
func (_JBMultiTerminal *JBMultiTerminalTransactorSession) MigrateBalanceOf(projectId *big.Int, token common.Address, to common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.MigrateBalanceOf(&_JBMultiTerminal.TransactOpts, projectId, token, to)
}

// Pay is a paid mutator transaction binding the contract method 0xfef43257.
//
// Solidity: function pay(uint256 projectId, address token, uint256 amount, address beneficiary, uint256 minReturnedTokens, string memo, bytes metadata) payable returns(uint256 beneficiaryTokenCount)
func (_JBMultiTerminal *JBMultiTerminalTransactor) Pay(opts *bind.TransactOpts, projectId *big.Int, token common.Address, amount *big.Int, beneficiary common.Address, minReturnedTokens *big.Int, memo string, metadata []byte) (*types.Transaction, error) {
	return _JBMultiTerminal.contract.Transact(opts, "pay", projectId, token, amount, beneficiary, minReturnedTokens, memo, metadata)
}

// Pay is a paid mutator transaction binding the contract method 0xfef43257.
//
// Solidity: function pay(uint256 projectId, address token, uint256 amount, address beneficiary, uint256 minReturnedTokens, string memo, bytes metadata) payable returns(uint256 beneficiaryTokenCount)
func (_JBMultiTerminal *JBMultiTerminalSession) Pay(projectId *big.Int, token common.Address, amount *big.Int, beneficiary common.Address, minReturnedTokens *big.Int, memo string, metadata []byte) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.Pay(&_JBMultiTerminal.TransactOpts, projectId, token, amount, beneficiary, minReturnedTokens, memo, metadata)
}

// Pay is a paid mutator transaction binding the contract method 0xfef43257.
//
// Solidity: function pay(uint256 projectId, address token, uint256 amount, address beneficiary, uint256 minReturnedTokens, string memo, bytes metadata) payable returns(uint256 beneficiaryTokenCount)
func (_JBMultiTerminal *JBMultiTerminalTransactorSession) Pay(projectId *big.Int, token common.Address, amount *big.Int, beneficiary common.Address, minReturnedTokens *big.Int, memo string, metadata []byte) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.Pay(&_JBMultiTerminal.TransactOpts, projectId, token, amount, beneficiary, minReturnedTokens, memo, metadata)
}

// ProcessHeldFeesOf is a paid mutator transaction binding the contract method 0x62ef3036.
//
// Solidity: function processHeldFeesOf(uint256 projectId, address token) returns()
func (_JBMultiTerminal *JBMultiTerminalTransactor) ProcessHeldFeesOf(opts *bind.TransactOpts, projectId *big.Int, token common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.contract.Transact(opts, "processHeldFeesOf", projectId, token)
}

// ProcessHeldFeesOf is a paid mutator transaction binding the contract method 0x62ef3036.
//
// Solidity: function processHeldFeesOf(uint256 projectId, address token) returns()
func (_JBMultiTerminal *JBMultiTerminalSession) ProcessHeldFeesOf(projectId *big.Int, token common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.ProcessHeldFeesOf(&_JBMultiTerminal.TransactOpts, projectId, token)
}

// ProcessHeldFeesOf is a paid mutator transaction binding the contract method 0x62ef3036.
//
// Solidity: function processHeldFeesOf(uint256 projectId, address token) returns()
func (_JBMultiTerminal *JBMultiTerminalTransactorSession) ProcessHeldFeesOf(projectId *big.Int, token common.Address) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.ProcessHeldFeesOf(&_JBMultiTerminal.TransactOpts, projectId, token)
}

// RedeemTokensOf is a paid mutator transaction binding the contract method 0xbb7214fe.
//
// Solidity: function redeemTokensOf(address holder, uint256 projectId, address tokenToReclaim, uint256 redeemCount, uint256 minTokensReclaimed, address beneficiary, bytes metadata) returns(uint256 reclaimAmount)
func (_JBMultiTerminal *JBMultiTerminalTransactor) RedeemTokensOf(opts *bind.TransactOpts, holder common.Address, projectId *big.Int, tokenToReclaim common.Address, redeemCount *big.Int, minTokensReclaimed *big.Int, beneficiary common.Address, metadata []byte) (*types.Transaction, error) {
	return _JBMultiTerminal.contract.Transact(opts, "redeemTokensOf", holder, projectId, tokenToReclaim, redeemCount, minTokensReclaimed, beneficiary, metadata)
}

// RedeemTokensOf is a paid mutator transaction binding the contract method 0xbb7214fe.
//
// Solidity: function redeemTokensOf(address holder, uint256 projectId, address tokenToReclaim, uint256 redeemCount, uint256 minTokensReclaimed, address beneficiary, bytes metadata) returns(uint256 reclaimAmount)
func (_JBMultiTerminal *JBMultiTerminalSession) RedeemTokensOf(holder common.Address, projectId *big.Int, tokenToReclaim common.Address, redeemCount *big.Int, minTokensReclaimed *big.Int, beneficiary common.Address, metadata []byte) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.RedeemTokensOf(&_JBMultiTerminal.TransactOpts, holder, projectId, tokenToReclaim, redeemCount, minTokensReclaimed, beneficiary, metadata)
}

// RedeemTokensOf is a paid mutator transaction binding the contract method 0xbb7214fe.
//
// Solidity: function redeemTokensOf(address holder, uint256 projectId, address tokenToReclaim, uint256 redeemCount, uint256 minTokensReclaimed, address beneficiary, bytes metadata) returns(uint256 reclaimAmount)
func (_JBMultiTerminal *JBMultiTerminalTransactorSession) RedeemTokensOf(holder common.Address, projectId *big.Int, tokenToReclaim common.Address, redeemCount *big.Int, minTokensReclaimed *big.Int, beneficiary common.Address, metadata []byte) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.RedeemTokensOf(&_JBMultiTerminal.TransactOpts, holder, projectId, tokenToReclaim, redeemCount, minTokensReclaimed, beneficiary, metadata)
}

// SendPayoutsOf is a paid mutator transaction binding the contract method 0xcfaf5839.
//
// Solidity: function sendPayoutsOf(uint256 projectId, address token, uint256 amount, uint256 currency, uint256 minTokensPaidOut) returns(uint256 amountPaidOut)
func (_JBMultiTerminal *JBMultiTerminalTransactor) SendPayoutsOf(opts *bind.TransactOpts, projectId *big.Int, token common.Address, amount *big.Int, currency *big.Int, minTokensPaidOut *big.Int) (*types.Transaction, error) {
	return _JBMultiTerminal.contract.Transact(opts, "sendPayoutsOf", projectId, token, amount, currency, minTokensPaidOut)
}

// SendPayoutsOf is a paid mutator transaction binding the contract method 0xcfaf5839.
//
// Solidity: function sendPayoutsOf(uint256 projectId, address token, uint256 amount, uint256 currency, uint256 minTokensPaidOut) returns(uint256 amountPaidOut)
func (_JBMultiTerminal *JBMultiTerminalSession) SendPayoutsOf(projectId *big.Int, token common.Address, amount *big.Int, currency *big.Int, minTokensPaidOut *big.Int) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.SendPayoutsOf(&_JBMultiTerminal.TransactOpts, projectId, token, amount, currency, minTokensPaidOut)
}

// SendPayoutsOf is a paid mutator transaction binding the contract method 0xcfaf5839.
//
// Solidity: function sendPayoutsOf(uint256 projectId, address token, uint256 amount, uint256 currency, uint256 minTokensPaidOut) returns(uint256 amountPaidOut)
func (_JBMultiTerminal *JBMultiTerminalTransactorSession) SendPayoutsOf(projectId *big.Int, token common.Address, amount *big.Int, currency *big.Int, minTokensPaidOut *big.Int) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.SendPayoutsOf(&_JBMultiTerminal.TransactOpts, projectId, token, amount, currency, minTokensPaidOut)
}

// UseAllowanceOf is a paid mutator transaction binding the contract method 0xb62d17cf.
//
// Solidity: function useAllowanceOf(uint256 projectId, address token, uint256 amount, uint256 currency, uint256 minTokensPaidOut, address beneficiary, string memo) returns(uint256 amountPaidOut)
func (_JBMultiTerminal *JBMultiTerminalTransactor) UseAllowanceOf(opts *bind.TransactOpts, projectId *big.Int, token common.Address, amount *big.Int, currency *big.Int, minTokensPaidOut *big.Int, beneficiary common.Address, memo string) (*types.Transaction, error) {
	return _JBMultiTerminal.contract.Transact(opts, "useAllowanceOf", projectId, token, amount, currency, minTokensPaidOut, beneficiary, memo)
}

// UseAllowanceOf is a paid mutator transaction binding the contract method 0xb62d17cf.
//
// Solidity: function useAllowanceOf(uint256 projectId, address token, uint256 amount, uint256 currency, uint256 minTokensPaidOut, address beneficiary, string memo) returns(uint256 amountPaidOut)
func (_JBMultiTerminal *JBMultiTerminalSession) UseAllowanceOf(projectId *big.Int, token common.Address, amount *big.Int, currency *big.Int, minTokensPaidOut *big.Int, beneficiary common.Address, memo string) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.UseAllowanceOf(&_JBMultiTerminal.TransactOpts, projectId, token, amount, currency, minTokensPaidOut, beneficiary, memo)
}

// UseAllowanceOf is a paid mutator transaction binding the contract method 0xb62d17cf.
//
// Solidity: function useAllowanceOf(uint256 projectId, address token, uint256 amount, uint256 currency, uint256 minTokensPaidOut, address beneficiary, string memo) returns(uint256 amountPaidOut)
func (_JBMultiTerminal *JBMultiTerminalTransactorSession) UseAllowanceOf(projectId *big.Int, token common.Address, amount *big.Int, currency *big.Int, minTokensPaidOut *big.Int, beneficiary common.Address, memo string) (*types.Transaction, error) {
	return _JBMultiTerminal.Contract.UseAllowanceOf(&_JBMultiTerminal.TransactOpts, projectId, token, amount, currency, minTokensPaidOut, beneficiary, memo)
}

// JBMultiTerminalAddToBalanceIterator is returned from FilterAddToBalance and is used to iterate over the raw logs and unpacked data for AddToBalance events raised by the JBMultiTerminal contract.
type JBMultiTerminalAddToBalanceIterator struct {
	Event *JBMultiTerminalAddToBalance // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalAddToBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalAddToBalance)
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
		it.Event = new(JBMultiTerminalAddToBalance)
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
func (it *JBMultiTerminalAddToBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalAddToBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalAddToBalance represents a AddToBalance event raised by the JBMultiTerminal contract.
type JBMultiTerminalAddToBalance struct {
	ProjectId    *big.Int
	Amount       *big.Int
	UnlockedFees *big.Int
	Memo         string
	Metadata     []byte
	Caller       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddToBalance is a free log retrieval operation binding the contract event 0x9ecaf7fc3dfffd6867c175d6e684b1f1e3aef019398ba8db2c1ffab4a09db253.
//
// Solidity: event AddToBalance(uint256 indexed projectId, uint256 amount, uint256 unlockedFees, string memo, bytes metadata, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterAddToBalance(opts *bind.FilterOpts, projectId []*big.Int) (*JBMultiTerminalAddToBalanceIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "AddToBalance", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalAddToBalanceIterator{contract: _JBMultiTerminal.contract, event: "AddToBalance", logs: logs, sub: sub}, nil
}

// WatchAddToBalance is a free log subscription operation binding the contract event 0x9ecaf7fc3dfffd6867c175d6e684b1f1e3aef019398ba8db2c1ffab4a09db253.
//
// Solidity: event AddToBalance(uint256 indexed projectId, uint256 amount, uint256 unlockedFees, string memo, bytes metadata, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchAddToBalance(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalAddToBalance, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "AddToBalance", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalAddToBalance)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "AddToBalance", log); err != nil {
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

// ParseAddToBalance is a log parse operation binding the contract event 0x9ecaf7fc3dfffd6867c175d6e684b1f1e3aef019398ba8db2c1ffab4a09db253.
//
// Solidity: event AddToBalance(uint256 indexed projectId, uint256 amount, uint256 unlockedFees, string memo, bytes metadata, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseAddToBalance(log types.Log) (*JBMultiTerminalAddToBalance, error) {
	event := new(JBMultiTerminalAddToBalance)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "AddToBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalFeeRevertedIterator is returned from FilterFeeReverted and is used to iterate over the raw logs and unpacked data for FeeReverted events raised by the JBMultiTerminal contract.
type JBMultiTerminalFeeRevertedIterator struct {
	Event *JBMultiTerminalFeeReverted // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalFeeRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalFeeReverted)
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
		it.Event = new(JBMultiTerminalFeeReverted)
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
func (it *JBMultiTerminalFeeRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalFeeRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalFeeReverted represents a FeeReverted event raised by the JBMultiTerminal contract.
type JBMultiTerminalFeeReverted struct {
	ProjectId    *big.Int
	Token        common.Address
	FeeProjectId *big.Int
	Amount       *big.Int
	Reason       []byte
	Caller       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeReverted is a free log retrieval operation binding the contract event 0x4b73f3c5cfb174c9d0494138d07ff8ee3aa400df46327d9893ed5ac122dd558a.
//
// Solidity: event FeeReverted(uint256 indexed projectId, address indexed token, uint256 indexed feeProjectId, uint256 amount, bytes reason, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterFeeReverted(opts *bind.FilterOpts, projectId []*big.Int, token []common.Address, feeProjectId []*big.Int) (*JBMultiTerminalFeeRevertedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var feeProjectIdRule []interface{}
	for _, feeProjectIdItem := range feeProjectId {
		feeProjectIdRule = append(feeProjectIdRule, feeProjectIdItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "FeeReverted", projectIdRule, tokenRule, feeProjectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalFeeRevertedIterator{contract: _JBMultiTerminal.contract, event: "FeeReverted", logs: logs, sub: sub}, nil
}

// WatchFeeReverted is a free log subscription operation binding the contract event 0x4b73f3c5cfb174c9d0494138d07ff8ee3aa400df46327d9893ed5ac122dd558a.
//
// Solidity: event FeeReverted(uint256 indexed projectId, address indexed token, uint256 indexed feeProjectId, uint256 amount, bytes reason, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchFeeReverted(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalFeeReverted, projectId []*big.Int, token []common.Address, feeProjectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var feeProjectIdRule []interface{}
	for _, feeProjectIdItem := range feeProjectId {
		feeProjectIdRule = append(feeProjectIdRule, feeProjectIdItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "FeeReverted", projectIdRule, tokenRule, feeProjectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalFeeReverted)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "FeeReverted", log); err != nil {
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

// ParseFeeReverted is a log parse operation binding the contract event 0x4b73f3c5cfb174c9d0494138d07ff8ee3aa400df46327d9893ed5ac122dd558a.
//
// Solidity: event FeeReverted(uint256 indexed projectId, address indexed token, uint256 indexed feeProjectId, uint256 amount, bytes reason, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseFeeReverted(log types.Log) (*JBMultiTerminalFeeReverted, error) {
	event := new(JBMultiTerminalFeeReverted)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "FeeReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalHoldFeeIterator is returned from FilterHoldFee and is used to iterate over the raw logs and unpacked data for HoldFee events raised by the JBMultiTerminal contract.
type JBMultiTerminalHoldFeeIterator struct {
	Event *JBMultiTerminalHoldFee // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalHoldFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalHoldFee)
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
		it.Event = new(JBMultiTerminalHoldFee)
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
func (it *JBMultiTerminalHoldFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalHoldFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalHoldFee represents a HoldFee event raised by the JBMultiTerminal contract.
type JBMultiTerminalHoldFee struct {
	ProjectId   *big.Int
	Token       common.Address
	Amount      *big.Int
	Fee         *big.Int
	Beneficiary common.Address
	Caller      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHoldFee is a free log retrieval operation binding the contract event 0xef9cd8fc3b61b64b44042ada0a8a3ede1649bc1b112673da228120f13bf1381e.
//
// Solidity: event HoldFee(uint256 indexed projectId, address indexed token, uint256 indexed amount, uint256 fee, address beneficiary, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterHoldFee(opts *bind.FilterOpts, projectId []*big.Int, token []common.Address, amount []*big.Int) (*JBMultiTerminalHoldFeeIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "HoldFee", projectIdRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalHoldFeeIterator{contract: _JBMultiTerminal.contract, event: "HoldFee", logs: logs, sub: sub}, nil
}

// WatchHoldFee is a free log subscription operation binding the contract event 0xef9cd8fc3b61b64b44042ada0a8a3ede1649bc1b112673da228120f13bf1381e.
//
// Solidity: event HoldFee(uint256 indexed projectId, address indexed token, uint256 indexed amount, uint256 fee, address beneficiary, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchHoldFee(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalHoldFee, projectId []*big.Int, token []common.Address, amount []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "HoldFee", projectIdRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalHoldFee)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "HoldFee", log); err != nil {
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

// ParseHoldFee is a log parse operation binding the contract event 0xef9cd8fc3b61b64b44042ada0a8a3ede1649bc1b112673da228120f13bf1381e.
//
// Solidity: event HoldFee(uint256 indexed projectId, address indexed token, uint256 indexed amount, uint256 fee, address beneficiary, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseHoldFee(log types.Log) (*JBMultiTerminalHoldFee, error) {
	event := new(JBMultiTerminalHoldFee)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "HoldFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalHookAfterRecordPayIterator is returned from FilterHookAfterRecordPay and is used to iterate over the raw logs and unpacked data for HookAfterRecordPay events raised by the JBMultiTerminal contract.
type JBMultiTerminalHookAfterRecordPayIterator struct {
	Event *JBMultiTerminalHookAfterRecordPay // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalHookAfterRecordPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalHookAfterRecordPay)
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
		it.Event = new(JBMultiTerminalHookAfterRecordPay)
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
func (it *JBMultiTerminalHookAfterRecordPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalHookAfterRecordPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalHookAfterRecordPay represents a HookAfterRecordPay event raised by the JBMultiTerminal contract.
type JBMultiTerminalHookAfterRecordPay struct {
	Hook                common.Address
	Context             JBAfterPayRecordedContext
	SpecificationAmount *big.Int
	Caller              common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHookAfterRecordPay is a free log retrieval operation binding the contract event 0x092bf36eb737817b64e5cea9e9b69305c423a987b5c3066516b046cad5c0bc36.
//
// Solidity: event HookAfterRecordPay(address indexed hook, (address,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),uint256,uint256,address,bytes,bytes) context, uint256 specificationAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterHookAfterRecordPay(opts *bind.FilterOpts, hook []common.Address) (*JBMultiTerminalHookAfterRecordPayIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "HookAfterRecordPay", hookRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalHookAfterRecordPayIterator{contract: _JBMultiTerminal.contract, event: "HookAfterRecordPay", logs: logs, sub: sub}, nil
}

// WatchHookAfterRecordPay is a free log subscription operation binding the contract event 0x092bf36eb737817b64e5cea9e9b69305c423a987b5c3066516b046cad5c0bc36.
//
// Solidity: event HookAfterRecordPay(address indexed hook, (address,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),uint256,uint256,address,bytes,bytes) context, uint256 specificationAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchHookAfterRecordPay(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalHookAfterRecordPay, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "HookAfterRecordPay", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalHookAfterRecordPay)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "HookAfterRecordPay", log); err != nil {
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

// ParseHookAfterRecordPay is a log parse operation binding the contract event 0x092bf36eb737817b64e5cea9e9b69305c423a987b5c3066516b046cad5c0bc36.
//
// Solidity: event HookAfterRecordPay(address indexed hook, (address,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),uint256,uint256,address,bytes,bytes) context, uint256 specificationAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseHookAfterRecordPay(log types.Log) (*JBMultiTerminalHookAfterRecordPay, error) {
	event := new(JBMultiTerminalHookAfterRecordPay)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "HookAfterRecordPay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalHookAfterRecordRedeemIterator is returned from FilterHookAfterRecordRedeem and is used to iterate over the raw logs and unpacked data for HookAfterRecordRedeem events raised by the JBMultiTerminal contract.
type JBMultiTerminalHookAfterRecordRedeemIterator struct {
	Event *JBMultiTerminalHookAfterRecordRedeem // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalHookAfterRecordRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalHookAfterRecordRedeem)
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
		it.Event = new(JBMultiTerminalHookAfterRecordRedeem)
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
func (it *JBMultiTerminalHookAfterRecordRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalHookAfterRecordRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalHookAfterRecordRedeem represents a HookAfterRecordRedeem event raised by the JBMultiTerminal contract.
type JBMultiTerminalHookAfterRecordRedeem struct {
	Hook                common.Address
	Context             JBAfterRedeemRecordedContext
	SpecificationAmount *big.Int
	Fee                 *big.Int
	Caller              common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHookAfterRecordRedeem is a free log retrieval operation binding the contract event 0x5d2ae7ac07b3a26000acaa692b550c84cbb87300557e0359791cfec9efe1615a.
//
// Solidity: event HookAfterRecordRedeem(address indexed hook, (address,uint256,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),uint256,address,bytes,bytes) context, uint256 specificationAmount, uint256 fee, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterHookAfterRecordRedeem(opts *bind.FilterOpts, hook []common.Address) (*JBMultiTerminalHookAfterRecordRedeemIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "HookAfterRecordRedeem", hookRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalHookAfterRecordRedeemIterator{contract: _JBMultiTerminal.contract, event: "HookAfterRecordRedeem", logs: logs, sub: sub}, nil
}

// WatchHookAfterRecordRedeem is a free log subscription operation binding the contract event 0x5d2ae7ac07b3a26000acaa692b550c84cbb87300557e0359791cfec9efe1615a.
//
// Solidity: event HookAfterRecordRedeem(address indexed hook, (address,uint256,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),uint256,address,bytes,bytes) context, uint256 specificationAmount, uint256 fee, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchHookAfterRecordRedeem(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalHookAfterRecordRedeem, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "HookAfterRecordRedeem", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalHookAfterRecordRedeem)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "HookAfterRecordRedeem", log); err != nil {
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

// ParseHookAfterRecordRedeem is a log parse operation binding the contract event 0x5d2ae7ac07b3a26000acaa692b550c84cbb87300557e0359791cfec9efe1615a.
//
// Solidity: event HookAfterRecordRedeem(address indexed hook, (address,uint256,uint256,uint256,(address,uint256,uint256,uint256),(address,uint256,uint256,uint256),uint256,address,bytes,bytes) context, uint256 specificationAmount, uint256 fee, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseHookAfterRecordRedeem(log types.Log) (*JBMultiTerminalHookAfterRecordRedeem, error) {
	event := new(JBMultiTerminalHookAfterRecordRedeem)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "HookAfterRecordRedeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalMigrateTerminalIterator is returned from FilterMigrateTerminal and is used to iterate over the raw logs and unpacked data for MigrateTerminal events raised by the JBMultiTerminal contract.
type JBMultiTerminalMigrateTerminalIterator struct {
	Event *JBMultiTerminalMigrateTerminal // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalMigrateTerminalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalMigrateTerminal)
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
		it.Event = new(JBMultiTerminalMigrateTerminal)
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
func (it *JBMultiTerminalMigrateTerminalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalMigrateTerminalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalMigrateTerminal represents a MigrateTerminal event raised by the JBMultiTerminal contract.
type JBMultiTerminalMigrateTerminal struct {
	ProjectId *big.Int
	Token     common.Address
	To        common.Address
	Amount    *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMigrateTerminal is a free log retrieval operation binding the contract event 0xf0aa72bb9c2a07cf43281bfe8c525ba320fdc71e18b54c9f93ab22ce074af7f6.
//
// Solidity: event MigrateTerminal(uint256 indexed projectId, address indexed token, address indexed to, uint256 amount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterMigrateTerminal(opts *bind.FilterOpts, projectId []*big.Int, token []common.Address, to []common.Address) (*JBMultiTerminalMigrateTerminalIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "MigrateTerminal", projectIdRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalMigrateTerminalIterator{contract: _JBMultiTerminal.contract, event: "MigrateTerminal", logs: logs, sub: sub}, nil
}

// WatchMigrateTerminal is a free log subscription operation binding the contract event 0xf0aa72bb9c2a07cf43281bfe8c525ba320fdc71e18b54c9f93ab22ce074af7f6.
//
// Solidity: event MigrateTerminal(uint256 indexed projectId, address indexed token, address indexed to, uint256 amount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchMigrateTerminal(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalMigrateTerminal, projectId []*big.Int, token []common.Address, to []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "MigrateTerminal", projectIdRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalMigrateTerminal)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "MigrateTerminal", log); err != nil {
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

// ParseMigrateTerminal is a log parse operation binding the contract event 0xf0aa72bb9c2a07cf43281bfe8c525ba320fdc71e18b54c9f93ab22ce074af7f6.
//
// Solidity: event MigrateTerminal(uint256 indexed projectId, address indexed token, address indexed to, uint256 amount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseMigrateTerminal(log types.Log) (*JBMultiTerminalMigrateTerminal, error) {
	event := new(JBMultiTerminalMigrateTerminal)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "MigrateTerminal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalPayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the JBMultiTerminal contract.
type JBMultiTerminalPayIterator struct {
	Event *JBMultiTerminalPay // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalPay)
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
		it.Event = new(JBMultiTerminalPay)
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
func (it *JBMultiTerminalPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalPay represents a Pay event raised by the JBMultiTerminal contract.
type JBMultiTerminalPay struct {
	RulesetId             *big.Int
	RulesetCycleNumber    *big.Int
	ProjectId             *big.Int
	Payer                 common.Address
	Beneficiary           common.Address
	Amount                *big.Int
	BeneficiaryTokenCount *big.Int
	Memo                  string
	Metadata              []byte
	Caller                common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x133161f1c9161488f777ab9a26aae91d47c0d9a3fafb398960f138db02c73797.
//
// Solidity: event Pay(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address payer, address beneficiary, uint256 amount, uint256 beneficiaryTokenCount, string memo, bytes metadata, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterPay(opts *bind.FilterOpts, rulesetId []*big.Int, rulesetCycleNumber []*big.Int, projectId []*big.Int) (*JBMultiTerminalPayIterator, error) {

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

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "Pay", rulesetIdRule, rulesetCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalPayIterator{contract: _JBMultiTerminal.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x133161f1c9161488f777ab9a26aae91d47c0d9a3fafb398960f138db02c73797.
//
// Solidity: event Pay(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address payer, address beneficiary, uint256 amount, uint256 beneficiaryTokenCount, string memo, bytes metadata, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalPay, rulesetId []*big.Int, rulesetCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "Pay", rulesetIdRule, rulesetCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalPay)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "Pay", log); err != nil {
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

// ParsePay is a log parse operation binding the contract event 0x133161f1c9161488f777ab9a26aae91d47c0d9a3fafb398960f138db02c73797.
//
// Solidity: event Pay(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address payer, address beneficiary, uint256 amount, uint256 beneficiaryTokenCount, string memo, bytes metadata, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParsePay(log types.Log) (*JBMultiTerminalPay, error) {
	event := new(JBMultiTerminalPay)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalPayoutRevertedIterator is returned from FilterPayoutReverted and is used to iterate over the raw logs and unpacked data for PayoutReverted events raised by the JBMultiTerminal contract.
type JBMultiTerminalPayoutRevertedIterator struct {
	Event *JBMultiTerminalPayoutReverted // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalPayoutRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalPayoutReverted)
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
		it.Event = new(JBMultiTerminalPayoutReverted)
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
func (it *JBMultiTerminalPayoutRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalPayoutRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalPayoutReverted represents a PayoutReverted event raised by the JBMultiTerminal contract.
type JBMultiTerminalPayoutReverted struct {
	ProjectId *big.Int
	Split     JBSplit
	Amount    *big.Int
	Reason    []byte
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayoutReverted is a free log retrieval operation binding the contract event 0x37b65da09d1a43465d8f823dcc858aa187792b35585b56afe707d57d7108d228.
//
// Solidity: event PayoutReverted(uint256 indexed projectId, (bool,uint256,uint256,address,uint256,address) split, uint256 amount, bytes reason, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterPayoutReverted(opts *bind.FilterOpts, projectId []*big.Int) (*JBMultiTerminalPayoutRevertedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "PayoutReverted", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalPayoutRevertedIterator{contract: _JBMultiTerminal.contract, event: "PayoutReverted", logs: logs, sub: sub}, nil
}

// WatchPayoutReverted is a free log subscription operation binding the contract event 0x37b65da09d1a43465d8f823dcc858aa187792b35585b56afe707d57d7108d228.
//
// Solidity: event PayoutReverted(uint256 indexed projectId, (bool,uint256,uint256,address,uint256,address) split, uint256 amount, bytes reason, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchPayoutReverted(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalPayoutReverted, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "PayoutReverted", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalPayoutReverted)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "PayoutReverted", log); err != nil {
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

// ParsePayoutReverted is a log parse operation binding the contract event 0x37b65da09d1a43465d8f823dcc858aa187792b35585b56afe707d57d7108d228.
//
// Solidity: event PayoutReverted(uint256 indexed projectId, (bool,uint256,uint256,address,uint256,address) split, uint256 amount, bytes reason, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParsePayoutReverted(log types.Log) (*JBMultiTerminalPayoutReverted, error) {
	event := new(JBMultiTerminalPayoutReverted)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "PayoutReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalProcessFeeIterator is returned from FilterProcessFee and is used to iterate over the raw logs and unpacked data for ProcessFee events raised by the JBMultiTerminal contract.
type JBMultiTerminalProcessFeeIterator struct {
	Event *JBMultiTerminalProcessFee // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalProcessFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalProcessFee)
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
		it.Event = new(JBMultiTerminalProcessFee)
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
func (it *JBMultiTerminalProcessFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalProcessFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalProcessFee represents a ProcessFee event raised by the JBMultiTerminal contract.
type JBMultiTerminalProcessFee struct {
	ProjectId   *big.Int
	Token       common.Address
	Amount      *big.Int
	WasHeld     bool
	Beneficiary common.Address
	Caller      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProcessFee is a free log retrieval operation binding the contract event 0xb514e730b3f8ad3aa94b6857bcc5ff4a46954bdcf8c4b0346705b1d0ac7a4325.
//
// Solidity: event ProcessFee(uint256 indexed projectId, address indexed token, uint256 indexed amount, bool wasHeld, address beneficiary, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterProcessFee(opts *bind.FilterOpts, projectId []*big.Int, token []common.Address, amount []*big.Int) (*JBMultiTerminalProcessFeeIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "ProcessFee", projectIdRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalProcessFeeIterator{contract: _JBMultiTerminal.contract, event: "ProcessFee", logs: logs, sub: sub}, nil
}

// WatchProcessFee is a free log subscription operation binding the contract event 0xb514e730b3f8ad3aa94b6857bcc5ff4a46954bdcf8c4b0346705b1d0ac7a4325.
//
// Solidity: event ProcessFee(uint256 indexed projectId, address indexed token, uint256 indexed amount, bool wasHeld, address beneficiary, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchProcessFee(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalProcessFee, projectId []*big.Int, token []common.Address, amount []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "ProcessFee", projectIdRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalProcessFee)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "ProcessFee", log); err != nil {
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

// ParseProcessFee is a log parse operation binding the contract event 0xb514e730b3f8ad3aa94b6857bcc5ff4a46954bdcf8c4b0346705b1d0ac7a4325.
//
// Solidity: event ProcessFee(uint256 indexed projectId, address indexed token, uint256 indexed amount, bool wasHeld, address beneficiary, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseProcessFee(log types.Log) (*JBMultiTerminalProcessFee, error) {
	event := new(JBMultiTerminalProcessFee)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "ProcessFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalRedeemTokensIterator is returned from FilterRedeemTokens and is used to iterate over the raw logs and unpacked data for RedeemTokens events raised by the JBMultiTerminal contract.
type JBMultiTerminalRedeemTokensIterator struct {
	Event *JBMultiTerminalRedeemTokens // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalRedeemTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalRedeemTokens)
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
		it.Event = new(JBMultiTerminalRedeemTokens)
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
func (it *JBMultiTerminalRedeemTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalRedeemTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalRedeemTokens represents a RedeemTokens event raised by the JBMultiTerminal contract.
type JBMultiTerminalRedeemTokens struct {
	RulesetId          *big.Int
	RulesetCycleNumber *big.Int
	ProjectId          *big.Int
	Holder             common.Address
	Beneficiary        common.Address
	TokenCount         *big.Int
	RedemptionRate     *big.Int
	ReclaimedAmount    *big.Int
	Metadata           []byte
	Caller             common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRedeemTokens is a free log retrieval operation binding the contract event 0x9565fc8bdd4f45a73582d1163dbe522a4fb0c4dbb882b5e66269f70101a2509a.
//
// Solidity: event RedeemTokens(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address holder, address beneficiary, uint256 tokenCount, uint256 redemptionRate, uint256 reclaimedAmount, bytes metadata, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterRedeemTokens(opts *bind.FilterOpts, rulesetId []*big.Int, rulesetCycleNumber []*big.Int, projectId []*big.Int) (*JBMultiTerminalRedeemTokensIterator, error) {

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

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "RedeemTokens", rulesetIdRule, rulesetCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalRedeemTokensIterator{contract: _JBMultiTerminal.contract, event: "RedeemTokens", logs: logs, sub: sub}, nil
}

// WatchRedeemTokens is a free log subscription operation binding the contract event 0x9565fc8bdd4f45a73582d1163dbe522a4fb0c4dbb882b5e66269f70101a2509a.
//
// Solidity: event RedeemTokens(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address holder, address beneficiary, uint256 tokenCount, uint256 redemptionRate, uint256 reclaimedAmount, bytes metadata, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchRedeemTokens(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalRedeemTokens, rulesetId []*big.Int, rulesetCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "RedeemTokens", rulesetIdRule, rulesetCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalRedeemTokens)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "RedeemTokens", log); err != nil {
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

// ParseRedeemTokens is a log parse operation binding the contract event 0x9565fc8bdd4f45a73582d1163dbe522a4fb0c4dbb882b5e66269f70101a2509a.
//
// Solidity: event RedeemTokens(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address holder, address beneficiary, uint256 tokenCount, uint256 redemptionRate, uint256 reclaimedAmount, bytes metadata, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseRedeemTokens(log types.Log) (*JBMultiTerminalRedeemTokens, error) {
	event := new(JBMultiTerminalRedeemTokens)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "RedeemTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalReturnHeldFeesIterator is returned from FilterReturnHeldFees and is used to iterate over the raw logs and unpacked data for ReturnHeldFees events raised by the JBMultiTerminal contract.
type JBMultiTerminalReturnHeldFeesIterator struct {
	Event *JBMultiTerminalReturnHeldFees // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalReturnHeldFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalReturnHeldFees)
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
		it.Event = new(JBMultiTerminalReturnHeldFees)
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
func (it *JBMultiTerminalReturnHeldFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalReturnHeldFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalReturnHeldFees represents a ReturnHeldFees event raised by the JBMultiTerminal contract.
type JBMultiTerminalReturnHeldFees struct {
	ProjectId      *big.Int
	Token          common.Address
	Amount         *big.Int
	ReturnedFees   *big.Int
	LeftoverAmount *big.Int
	Caller         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReturnHeldFees is a free log retrieval operation binding the contract event 0xedc5125503c021a73fd7d8ed0c2456b296be2c88f39fbed068bd091cef7ca43d.
//
// Solidity: event ReturnHeldFees(uint256 indexed projectId, address indexed token, uint256 indexed amount, uint256 returnedFees, uint256 leftoverAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterReturnHeldFees(opts *bind.FilterOpts, projectId []*big.Int, token []common.Address, amount []*big.Int) (*JBMultiTerminalReturnHeldFeesIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "ReturnHeldFees", projectIdRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalReturnHeldFeesIterator{contract: _JBMultiTerminal.contract, event: "ReturnHeldFees", logs: logs, sub: sub}, nil
}

// WatchReturnHeldFees is a free log subscription operation binding the contract event 0xedc5125503c021a73fd7d8ed0c2456b296be2c88f39fbed068bd091cef7ca43d.
//
// Solidity: event ReturnHeldFees(uint256 indexed projectId, address indexed token, uint256 indexed amount, uint256 returnedFees, uint256 leftoverAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchReturnHeldFees(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalReturnHeldFees, projectId []*big.Int, token []common.Address, amount []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "ReturnHeldFees", projectIdRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalReturnHeldFees)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "ReturnHeldFees", log); err != nil {
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

// ParseReturnHeldFees is a log parse operation binding the contract event 0xedc5125503c021a73fd7d8ed0c2456b296be2c88f39fbed068bd091cef7ca43d.
//
// Solidity: event ReturnHeldFees(uint256 indexed projectId, address indexed token, uint256 indexed amount, uint256 returnedFees, uint256 leftoverAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseReturnHeldFees(log types.Log) (*JBMultiTerminalReturnHeldFees, error) {
	event := new(JBMultiTerminalReturnHeldFees)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "ReturnHeldFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalSendPayoutToSplitIterator is returned from FilterSendPayoutToSplit and is used to iterate over the raw logs and unpacked data for SendPayoutToSplit events raised by the JBMultiTerminal contract.
type JBMultiTerminalSendPayoutToSplitIterator struct {
	Event *JBMultiTerminalSendPayoutToSplit // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalSendPayoutToSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalSendPayoutToSplit)
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
		it.Event = new(JBMultiTerminalSendPayoutToSplit)
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
func (it *JBMultiTerminalSendPayoutToSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalSendPayoutToSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalSendPayoutToSplit represents a SendPayoutToSplit event raised by the JBMultiTerminal contract.
type JBMultiTerminalSendPayoutToSplit struct {
	ProjectId *big.Int
	RulesetId *big.Int
	Group     *big.Int
	Split     JBSplit
	Amount    *big.Int
	NetAmount *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSendPayoutToSplit is a free log retrieval operation binding the contract event 0x39305f11865947ab83bc6548c8c77e5890ed5bc4a6ea54e6fe7b22f4f406f4b8.
//
// Solidity: event SendPayoutToSplit(uint256 indexed projectId, uint256 indexed rulesetId, uint256 indexed group, (bool,uint256,uint256,address,uint256,address) split, uint256 amount, uint256 netAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterSendPayoutToSplit(opts *bind.FilterOpts, projectId []*big.Int, rulesetId []*big.Int, group []*big.Int) (*JBMultiTerminalSendPayoutToSplitIterator, error) {

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

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "SendPayoutToSplit", projectIdRule, rulesetIdRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalSendPayoutToSplitIterator{contract: _JBMultiTerminal.contract, event: "SendPayoutToSplit", logs: logs, sub: sub}, nil
}

// WatchSendPayoutToSplit is a free log subscription operation binding the contract event 0x39305f11865947ab83bc6548c8c77e5890ed5bc4a6ea54e6fe7b22f4f406f4b8.
//
// Solidity: event SendPayoutToSplit(uint256 indexed projectId, uint256 indexed rulesetId, uint256 indexed group, (bool,uint256,uint256,address,uint256,address) split, uint256 amount, uint256 netAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchSendPayoutToSplit(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalSendPayoutToSplit, projectId []*big.Int, rulesetId []*big.Int, group []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "SendPayoutToSplit", projectIdRule, rulesetIdRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalSendPayoutToSplit)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "SendPayoutToSplit", log); err != nil {
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

// ParseSendPayoutToSplit is a log parse operation binding the contract event 0x39305f11865947ab83bc6548c8c77e5890ed5bc4a6ea54e6fe7b22f4f406f4b8.
//
// Solidity: event SendPayoutToSplit(uint256 indexed projectId, uint256 indexed rulesetId, uint256 indexed group, (bool,uint256,uint256,address,uint256,address) split, uint256 amount, uint256 netAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseSendPayoutToSplit(log types.Log) (*JBMultiTerminalSendPayoutToSplit, error) {
	event := new(JBMultiTerminalSendPayoutToSplit)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "SendPayoutToSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalSendPayoutsIterator is returned from FilterSendPayouts and is used to iterate over the raw logs and unpacked data for SendPayouts events raised by the JBMultiTerminal contract.
type JBMultiTerminalSendPayoutsIterator struct {
	Event *JBMultiTerminalSendPayouts // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalSendPayoutsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalSendPayouts)
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
		it.Event = new(JBMultiTerminalSendPayouts)
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
func (it *JBMultiTerminalSendPayoutsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalSendPayoutsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalSendPayouts represents a SendPayouts event raised by the JBMultiTerminal contract.
type JBMultiTerminalSendPayouts struct {
	RulesetId                     *big.Int
	RulesetCycleNumber            *big.Int
	ProjectId                     *big.Int
	Beneficiary                   common.Address
	Amount                        *big.Int
	AmountPaidOut                 *big.Int
	Fee                           *big.Int
	BeneficiaryDistributionAmount *big.Int
	Caller                        common.Address
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterSendPayouts is a free log retrieval operation binding the contract event 0xda42c5c18ef97738d10c32811895add569670a6fcbe9627ea198d446f3be6dea.
//
// Solidity: event SendPayouts(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 amountPaidOut, uint256 fee, uint256 beneficiaryDistributionAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterSendPayouts(opts *bind.FilterOpts, rulesetId []*big.Int, rulesetCycleNumber []*big.Int, projectId []*big.Int) (*JBMultiTerminalSendPayoutsIterator, error) {

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

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "SendPayouts", rulesetIdRule, rulesetCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalSendPayoutsIterator{contract: _JBMultiTerminal.contract, event: "SendPayouts", logs: logs, sub: sub}, nil
}

// WatchSendPayouts is a free log subscription operation binding the contract event 0xda42c5c18ef97738d10c32811895add569670a6fcbe9627ea198d446f3be6dea.
//
// Solidity: event SendPayouts(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 amountPaidOut, uint256 fee, uint256 beneficiaryDistributionAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchSendPayouts(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalSendPayouts, rulesetId []*big.Int, rulesetCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "SendPayouts", rulesetIdRule, rulesetCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalSendPayouts)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "SendPayouts", log); err != nil {
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

// ParseSendPayouts is a log parse operation binding the contract event 0xda42c5c18ef97738d10c32811895add569670a6fcbe9627ea198d446f3be6dea.
//
// Solidity: event SendPayouts(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 amountPaidOut, uint256 fee, uint256 beneficiaryDistributionAmount, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseSendPayouts(log types.Log) (*JBMultiTerminalSendPayouts, error) {
	event := new(JBMultiTerminalSendPayouts)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "SendPayouts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalSetAccountingContextIterator is returned from FilterSetAccountingContext and is used to iterate over the raw logs and unpacked data for SetAccountingContext events raised by the JBMultiTerminal contract.
type JBMultiTerminalSetAccountingContextIterator struct {
	Event *JBMultiTerminalSetAccountingContext // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalSetAccountingContextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalSetAccountingContext)
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
		it.Event = new(JBMultiTerminalSetAccountingContext)
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
func (it *JBMultiTerminalSetAccountingContextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalSetAccountingContextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalSetAccountingContext represents a SetAccountingContext event raised by the JBMultiTerminal contract.
type JBMultiTerminalSetAccountingContext struct {
	ProjectId *big.Int
	Token     common.Address
	Context   JBAccountingContext
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetAccountingContext is a free log retrieval operation binding the contract event 0x9570de6fc0456f642bd0a001b24cd31676388c24596df64a17190c24310690a1.
//
// Solidity: event SetAccountingContext(uint256 indexed projectId, address indexed token, (address,uint8,uint32) context, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterSetAccountingContext(opts *bind.FilterOpts, projectId []*big.Int, token []common.Address) (*JBMultiTerminalSetAccountingContextIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "SetAccountingContext", projectIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalSetAccountingContextIterator{contract: _JBMultiTerminal.contract, event: "SetAccountingContext", logs: logs, sub: sub}, nil
}

// WatchSetAccountingContext is a free log subscription operation binding the contract event 0x9570de6fc0456f642bd0a001b24cd31676388c24596df64a17190c24310690a1.
//
// Solidity: event SetAccountingContext(uint256 indexed projectId, address indexed token, (address,uint8,uint32) context, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchSetAccountingContext(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalSetAccountingContext, projectId []*big.Int, token []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "SetAccountingContext", projectIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalSetAccountingContext)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "SetAccountingContext", log); err != nil {
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

// ParseSetAccountingContext is a log parse operation binding the contract event 0x9570de6fc0456f642bd0a001b24cd31676388c24596df64a17190c24310690a1.
//
// Solidity: event SetAccountingContext(uint256 indexed projectId, address indexed token, (address,uint8,uint32) context, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseSetAccountingContext(log types.Log) (*JBMultiTerminalSetAccountingContext, error) {
	event := new(JBMultiTerminalSetAccountingContext)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "SetAccountingContext", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBMultiTerminalUseAllowanceIterator is returned from FilterUseAllowance and is used to iterate over the raw logs and unpacked data for UseAllowance events raised by the JBMultiTerminal contract.
type JBMultiTerminalUseAllowanceIterator struct {
	Event *JBMultiTerminalUseAllowance // Event containing the contract specifics and raw log

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
func (it *JBMultiTerminalUseAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBMultiTerminalUseAllowance)
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
		it.Event = new(JBMultiTerminalUseAllowance)
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
func (it *JBMultiTerminalUseAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBMultiTerminalUseAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBMultiTerminalUseAllowance represents a UseAllowance event raised by the JBMultiTerminal contract.
type JBMultiTerminalUseAllowance struct {
	RulesetId          *big.Int
	RulesetCycleNumber *big.Int
	ProjectId          *big.Int
	Beneficiary        common.Address
	Amount             *big.Int
	AmountPaidOut      *big.Int
	NetAmountPaidOut   *big.Int
	Memo               string
	Caller             common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUseAllowance is a free log retrieval operation binding the contract event 0x8657a0c05a68a912c23c1bd00124afaa8c669063b046bd9bfd22b21d573c5e6d.
//
// Solidity: event UseAllowance(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 amountPaidOut, uint256 netAmountPaidOut, string memo, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) FilterUseAllowance(opts *bind.FilterOpts, rulesetId []*big.Int, rulesetCycleNumber []*big.Int, projectId []*big.Int) (*JBMultiTerminalUseAllowanceIterator, error) {

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

	logs, sub, err := _JBMultiTerminal.contract.FilterLogs(opts, "UseAllowance", rulesetIdRule, rulesetCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBMultiTerminalUseAllowanceIterator{contract: _JBMultiTerminal.contract, event: "UseAllowance", logs: logs, sub: sub}, nil
}

// WatchUseAllowance is a free log subscription operation binding the contract event 0x8657a0c05a68a912c23c1bd00124afaa8c669063b046bd9bfd22b21d573c5e6d.
//
// Solidity: event UseAllowance(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 amountPaidOut, uint256 netAmountPaidOut, string memo, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) WatchUseAllowance(opts *bind.WatchOpts, sink chan<- *JBMultiTerminalUseAllowance, rulesetId []*big.Int, rulesetCycleNumber []*big.Int, projectId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _JBMultiTerminal.contract.WatchLogs(opts, "UseAllowance", rulesetIdRule, rulesetCycleNumberRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBMultiTerminalUseAllowance)
				if err := _JBMultiTerminal.contract.UnpackLog(event, "UseAllowance", log); err != nil {
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

// ParseUseAllowance is a log parse operation binding the contract event 0x8657a0c05a68a912c23c1bd00124afaa8c669063b046bd9bfd22b21d573c5e6d.
//
// Solidity: event UseAllowance(uint256 indexed rulesetId, uint256 indexed rulesetCycleNumber, uint256 indexed projectId, address beneficiary, uint256 amount, uint256 amountPaidOut, uint256 netAmountPaidOut, string memo, address caller)
func (_JBMultiTerminal *JBMultiTerminalFilterer) ParseUseAllowance(log types.Log) (*JBMultiTerminalUseAllowance, error) {
	event := new(JBMultiTerminalUseAllowance)
	if err := _JBMultiTerminal.contract.UnpackLog(event, "UseAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
