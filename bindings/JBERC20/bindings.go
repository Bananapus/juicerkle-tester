// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package JBERC20

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

// CheckpointsCheckpoint208 is an auto generated low-level Go binding around an user-defined struct.
type CheckpointsCheckpoint208 struct {
	Key   *big.Int
	Value *big.Int
}

// JBERC20MetaData contains all meta data concerning the JBERC20 contract.
var JBERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pos\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCheckpoints.Checkpoint208\",\"components\":[{\"name\":\"_key\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"_value\",\"type\":\"uint208\",\"internalType\":\"uint208\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateBySig\",\"inputs\":[{\"name\":\"delegatee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegates\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPastTotalSupply\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPastVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numCheckpoints\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateChanged\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fromDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toDelegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateVotesChanged\",\"inputs\":[{\"name\":\"delegate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"previousVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20ExceededSafeSupply\",\"inputs\":[{\"name\":\"increasedSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC2612ExpiredSignature\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC2612InvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC5805FutureLookup\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clock\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"ERC6372InconsistentClock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccountNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"VotesExpiredSignature\",\"inputs\":[{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// JBERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use JBERC20MetaData.ABI instead.
var JBERC20ABI = JBERC20MetaData.ABI

// JBERC20 is an auto generated Go binding around an Ethereum contract.
type JBERC20 struct {
	JBERC20Caller     // Read-only binding to the contract
	JBERC20Transactor // Write-only binding to the contract
	JBERC20Filterer   // Log filterer for contract events
}

// JBERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type JBERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type JBERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JBERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JBERC20Session struct {
	Contract     *JBERC20          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JBERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JBERC20CallerSession struct {
	Contract *JBERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// JBERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JBERC20TransactorSession struct {
	Contract     *JBERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// JBERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type JBERC20Raw struct {
	Contract *JBERC20 // Generic contract binding to access the raw methods on
}

// JBERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JBERC20CallerRaw struct {
	Contract *JBERC20Caller // Generic read-only contract binding to access the raw methods on
}

// JBERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JBERC20TransactorRaw struct {
	Contract *JBERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewJBERC20 creates a new instance of JBERC20, bound to a specific deployed contract.
func NewJBERC20(address common.Address, backend bind.ContractBackend) (*JBERC20, error) {
	contract, err := bindJBERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JBERC20{JBERC20Caller: JBERC20Caller{contract: contract}, JBERC20Transactor: JBERC20Transactor{contract: contract}, JBERC20Filterer: JBERC20Filterer{contract: contract}}, nil
}

// NewJBERC20Caller creates a new read-only instance of JBERC20, bound to a specific deployed contract.
func NewJBERC20Caller(address common.Address, caller bind.ContractCaller) (*JBERC20Caller, error) {
	contract, err := bindJBERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JBERC20Caller{contract: contract}, nil
}

// NewJBERC20Transactor creates a new write-only instance of JBERC20, bound to a specific deployed contract.
func NewJBERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*JBERC20Transactor, error) {
	contract, err := bindJBERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JBERC20Transactor{contract: contract}, nil
}

// NewJBERC20Filterer creates a new log filterer instance of JBERC20, bound to a specific deployed contract.
func NewJBERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*JBERC20Filterer, error) {
	contract, err := bindJBERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JBERC20Filterer{contract: contract}, nil
}

// bindJBERC20 binds a generic wrapper to an already deployed contract.
func bindJBERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JBERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JBERC20 *JBERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JBERC20.Contract.JBERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JBERC20 *JBERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBERC20.Contract.JBERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JBERC20 *JBERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JBERC20.Contract.JBERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JBERC20 *JBERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JBERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JBERC20 *JBERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JBERC20 *JBERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JBERC20.Contract.contract.Transact(opts, method, params...)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_JBERC20 *JBERC20Caller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_JBERC20 *JBERC20Session) CLOCKMODE() (string, error) {
	return _JBERC20.Contract.CLOCKMODE(&_JBERC20.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_JBERC20 *JBERC20CallerSession) CLOCKMODE() (string, error) {
	return _JBERC20.Contract.CLOCKMODE(&_JBERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JBERC20 *JBERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JBERC20 *JBERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _JBERC20.Contract.DOMAINSEPARATOR(&_JBERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_JBERC20 *JBERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _JBERC20.Contract.DOMAINSEPARATOR(&_JBERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_JBERC20 *JBERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_JBERC20 *JBERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _JBERC20.Contract.Allowance(&_JBERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_JBERC20 *JBERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _JBERC20.Contract.Allowance(&_JBERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_JBERC20 *JBERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_JBERC20 *JBERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _JBERC20.Contract.BalanceOf(&_JBERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_JBERC20 *JBERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _JBERC20.Contract.BalanceOf(&_JBERC20.CallOpts, account)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (_JBERC20 *JBERC20Caller) Checkpoints(opts *bind.CallOpts, account common.Address, pos uint32) (CheckpointsCheckpoint208, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "checkpoints", account, pos)

	if err != nil {
		return *new(CheckpointsCheckpoint208), err
	}

	out0 := *abi.ConvertType(out[0], new(CheckpointsCheckpoint208)).(*CheckpointsCheckpoint208)

	return out0, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (_JBERC20 *JBERC20Session) Checkpoints(account common.Address, pos uint32) (CheckpointsCheckpoint208, error) {
	return _JBERC20.Contract.Checkpoints(&_JBERC20.CallOpts, account, pos)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint48,uint208))
func (_JBERC20 *JBERC20CallerSession) Checkpoints(account common.Address, pos uint32) (CheckpointsCheckpoint208, error) {
	return _JBERC20.Contract.Checkpoints(&_JBERC20.CallOpts, account, pos)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_JBERC20 *JBERC20Caller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_JBERC20 *JBERC20Session) Clock() (*big.Int, error) {
	return _JBERC20.Contract.Clock(&_JBERC20.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_JBERC20 *JBERC20CallerSession) Clock() (*big.Int, error) {
	return _JBERC20.Contract.Clock(&_JBERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JBERC20 *JBERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JBERC20 *JBERC20Session) Decimals() (uint8, error) {
	return _JBERC20.Contract.Decimals(&_JBERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_JBERC20 *JBERC20CallerSession) Decimals() (uint8, error) {
	return _JBERC20.Contract.Decimals(&_JBERC20.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_JBERC20 *JBERC20Caller) Delegates(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "delegates", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_JBERC20 *JBERC20Session) Delegates(account common.Address) (common.Address, error) {
	return _JBERC20.Contract.Delegates(&_JBERC20.CallOpts, account)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_JBERC20 *JBERC20CallerSession) Delegates(account common.Address) (common.Address, error) {
	return _JBERC20.Contract.Delegates(&_JBERC20.CallOpts, account)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_JBERC20 *JBERC20Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_JBERC20 *JBERC20Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _JBERC20.Contract.Eip712Domain(&_JBERC20.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_JBERC20 *JBERC20CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _JBERC20.Contract.Eip712Domain(&_JBERC20.CallOpts)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (_JBERC20 *JBERC20Caller) GetPastTotalSupply(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "getPastTotalSupply", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (_JBERC20 *JBERC20Session) GetPastTotalSupply(timepoint *big.Int) (*big.Int, error) {
	return _JBERC20.Contract.GetPastTotalSupply(&_JBERC20.CallOpts, timepoint)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timepoint) view returns(uint256)
func (_JBERC20 *JBERC20CallerSession) GetPastTotalSupply(timepoint *big.Int) (*big.Int, error) {
	return _JBERC20.Contract.GetPastTotalSupply(&_JBERC20.CallOpts, timepoint)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (_JBERC20 *JBERC20Caller) GetPastVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "getPastVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (_JBERC20 *JBERC20Session) GetPastVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _JBERC20.Contract.GetPastVotes(&_JBERC20.CallOpts, account, timepoint)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timepoint) view returns(uint256)
func (_JBERC20 *JBERC20CallerSession) GetPastVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _JBERC20.Contract.GetPastVotes(&_JBERC20.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_JBERC20 *JBERC20Caller) GetVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_JBERC20 *JBERC20Session) GetVotes(account common.Address) (*big.Int, error) {
	return _JBERC20.Contract.GetVotes(&_JBERC20.CallOpts, account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_JBERC20 *JBERC20CallerSession) GetVotes(account common.Address) (*big.Int, error) {
	return _JBERC20.Contract.GetVotes(&_JBERC20.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JBERC20 *JBERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JBERC20 *JBERC20Session) Name() (string, error) {
	return _JBERC20.Contract.Name(&_JBERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_JBERC20 *JBERC20CallerSession) Name() (string, error) {
	return _JBERC20.Contract.Name(&_JBERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_JBERC20 *JBERC20Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_JBERC20 *JBERC20Session) Nonces(owner common.Address) (*big.Int, error) {
	return _JBERC20.Contract.Nonces(&_JBERC20.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_JBERC20 *JBERC20CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _JBERC20.Contract.Nonces(&_JBERC20.CallOpts, owner)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_JBERC20 *JBERC20Caller) NumCheckpoints(opts *bind.CallOpts, account common.Address) (uint32, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "numCheckpoints", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_JBERC20 *JBERC20Session) NumCheckpoints(account common.Address) (uint32, error) {
	return _JBERC20.Contract.NumCheckpoints(&_JBERC20.CallOpts, account)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_JBERC20 *JBERC20CallerSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _JBERC20.Contract.NumCheckpoints(&_JBERC20.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JBERC20 *JBERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JBERC20 *JBERC20Session) Owner() (common.Address, error) {
	return _JBERC20.Contract.Owner(&_JBERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JBERC20 *JBERC20CallerSession) Owner() (common.Address, error) {
	return _JBERC20.Contract.Owner(&_JBERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JBERC20 *JBERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JBERC20 *JBERC20Session) Symbol() (string, error) {
	return _JBERC20.Contract.Symbol(&_JBERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_JBERC20 *JBERC20CallerSession) Symbol() (string, error) {
	return _JBERC20.Contract.Symbol(&_JBERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JBERC20 *JBERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JBERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JBERC20 *JBERC20Session) TotalSupply() (*big.Int, error) {
	return _JBERC20.Contract.TotalSupply(&_JBERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_JBERC20 *JBERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _JBERC20.Contract.TotalSupply(&_JBERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_JBERC20 *JBERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_JBERC20 *JBERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _JBERC20.Contract.Approve(&_JBERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_JBERC20 *JBERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _JBERC20.Contract.Approve(&_JBERC20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_JBERC20 *JBERC20Transactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_JBERC20 *JBERC20Session) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JBERC20.Contract.Burn(&_JBERC20.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_JBERC20 *JBERC20TransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JBERC20.Contract.Burn(&_JBERC20.TransactOpts, account, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_JBERC20 *JBERC20Transactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_JBERC20 *JBERC20Session) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _JBERC20.Contract.Delegate(&_JBERC20.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_JBERC20 *JBERC20TransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _JBERC20.Contract.Delegate(&_JBERC20.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_JBERC20 *JBERC20Transactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_JBERC20 *JBERC20Session) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JBERC20.Contract.DelegateBySig(&_JBERC20.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_JBERC20 *JBERC20TransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JBERC20.Contract.DelegateBySig(&_JBERC20.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name_, string symbol_, address owner) returns()
func (_JBERC20 *JBERC20Transactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, owner common.Address) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "initialize", name_, symbol_, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name_, string symbol_, address owner) returns()
func (_JBERC20 *JBERC20Session) Initialize(name_ string, symbol_ string, owner common.Address) (*types.Transaction, error) {
	return _JBERC20.Contract.Initialize(&_JBERC20.TransactOpts, name_, symbol_, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name_, string symbol_, address owner) returns()
func (_JBERC20 *JBERC20TransactorSession) Initialize(name_ string, symbol_ string, owner common.Address) (*types.Transaction, error) {
	return _JBERC20.Contract.Initialize(&_JBERC20.TransactOpts, name_, symbol_, owner)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_JBERC20 *JBERC20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_JBERC20 *JBERC20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JBERC20.Contract.Mint(&_JBERC20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_JBERC20 *JBERC20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JBERC20.Contract.Mint(&_JBERC20.TransactOpts, account, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JBERC20 *JBERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JBERC20 *JBERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JBERC20.Contract.Permit(&_JBERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_JBERC20 *JBERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _JBERC20.Contract.Permit(&_JBERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JBERC20 *JBERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JBERC20 *JBERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _JBERC20.Contract.RenounceOwnership(&_JBERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JBERC20 *JBERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _JBERC20.Contract.RenounceOwnership(&_JBERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_JBERC20 *JBERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_JBERC20 *JBERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JBERC20.Contract.Transfer(&_JBERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_JBERC20 *JBERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JBERC20.Contract.Transfer(&_JBERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_JBERC20 *JBERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_JBERC20 *JBERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JBERC20.Contract.TransferFrom(&_JBERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_JBERC20 *JBERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _JBERC20.Contract.TransferFrom(&_JBERC20.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JBERC20 *JBERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _JBERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JBERC20 *JBERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _JBERC20.Contract.TransferOwnership(&_JBERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JBERC20 *JBERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _JBERC20.Contract.TransferOwnership(&_JBERC20.TransactOpts, newOwner)
}

// JBERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the JBERC20 contract.
type JBERC20ApprovalIterator struct {
	Event *JBERC20Approval // Event containing the contract specifics and raw log

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
func (it *JBERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBERC20Approval)
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
		it.Event = new(JBERC20Approval)
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
func (it *JBERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBERC20Approval represents a Approval event raised by the JBERC20 contract.
type JBERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JBERC20 *JBERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*JBERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _JBERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &JBERC20ApprovalIterator{contract: _JBERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JBERC20 *JBERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *JBERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _JBERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBERC20Approval)
				if err := _JBERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_JBERC20 *JBERC20Filterer) ParseApproval(log types.Log) (*JBERC20Approval, error) {
	event := new(JBERC20Approval)
	if err := _JBERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBERC20DelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the JBERC20 contract.
type JBERC20DelegateChangedIterator struct {
	Event *JBERC20DelegateChanged // Event containing the contract specifics and raw log

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
func (it *JBERC20DelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBERC20DelegateChanged)
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
		it.Event = new(JBERC20DelegateChanged)
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
func (it *JBERC20DelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBERC20DelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBERC20DelegateChanged represents a DelegateChanged event raised by the JBERC20 contract.
type JBERC20DelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_JBERC20 *JBERC20Filterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*JBERC20DelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _JBERC20.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &JBERC20DelegateChangedIterator{contract: _JBERC20.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_JBERC20 *JBERC20Filterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *JBERC20DelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _JBERC20.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBERC20DelegateChanged)
				if err := _JBERC20.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_JBERC20 *JBERC20Filterer) ParseDelegateChanged(log types.Log) (*JBERC20DelegateChanged, error) {
	event := new(JBERC20DelegateChanged)
	if err := _JBERC20.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBERC20DelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the JBERC20 contract.
type JBERC20DelegateVotesChangedIterator struct {
	Event *JBERC20DelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *JBERC20DelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBERC20DelegateVotesChanged)
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
		it.Event = new(JBERC20DelegateVotesChanged)
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
func (it *JBERC20DelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBERC20DelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBERC20DelegateVotesChanged represents a DelegateVotesChanged event raised by the JBERC20 contract.
type JBERC20DelegateVotesChanged struct {
	Delegate      common.Address
	PreviousVotes *big.Int
	NewVotes      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_JBERC20 *JBERC20Filterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*JBERC20DelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _JBERC20.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &JBERC20DelegateVotesChangedIterator{contract: _JBERC20.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_JBERC20 *JBERC20Filterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *JBERC20DelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _JBERC20.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBERC20DelegateVotesChanged)
				if err := _JBERC20.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousVotes, uint256 newVotes)
func (_JBERC20 *JBERC20Filterer) ParseDelegateVotesChanged(log types.Log) (*JBERC20DelegateVotesChanged, error) {
	event := new(JBERC20DelegateVotesChanged)
	if err := _JBERC20.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBERC20EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the JBERC20 contract.
type JBERC20EIP712DomainChangedIterator struct {
	Event *JBERC20EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *JBERC20EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBERC20EIP712DomainChanged)
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
		it.Event = new(JBERC20EIP712DomainChanged)
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
func (it *JBERC20EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBERC20EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBERC20EIP712DomainChanged represents a EIP712DomainChanged event raised by the JBERC20 contract.
type JBERC20EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_JBERC20 *JBERC20Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*JBERC20EIP712DomainChangedIterator, error) {

	logs, sub, err := _JBERC20.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &JBERC20EIP712DomainChangedIterator{contract: _JBERC20.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_JBERC20 *JBERC20Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *JBERC20EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _JBERC20.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBERC20EIP712DomainChanged)
				if err := _JBERC20.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_JBERC20 *JBERC20Filterer) ParseEIP712DomainChanged(log types.Log) (*JBERC20EIP712DomainChanged, error) {
	event := new(JBERC20EIP712DomainChanged)
	if err := _JBERC20.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the JBERC20 contract.
type JBERC20OwnershipTransferredIterator struct {
	Event *JBERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *JBERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBERC20OwnershipTransferred)
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
		it.Event = new(JBERC20OwnershipTransferred)
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
func (it *JBERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBERC20OwnershipTransferred represents a OwnershipTransferred event raised by the JBERC20 contract.
type JBERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JBERC20 *JBERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JBERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JBERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JBERC20OwnershipTransferredIterator{contract: _JBERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JBERC20 *JBERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *JBERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JBERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBERC20OwnershipTransferred)
				if err := _JBERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_JBERC20 *JBERC20Filterer) ParseOwnershipTransferred(log types.Log) (*JBERC20OwnershipTransferred, error) {
	event := new(JBERC20OwnershipTransferred)
	if err := _JBERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the JBERC20 contract.
type JBERC20TransferIterator struct {
	Event *JBERC20Transfer // Event containing the contract specifics and raw log

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
func (it *JBERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBERC20Transfer)
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
		it.Event = new(JBERC20Transfer)
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
func (it *JBERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBERC20Transfer represents a Transfer event raised by the JBERC20 contract.
type JBERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JBERC20 *JBERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*JBERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JBERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &JBERC20TransferIterator{contract: _JBERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JBERC20 *JBERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *JBERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _JBERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBERC20Transfer)
				if err := _JBERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_JBERC20 *JBERC20Filterer) ParseTransfer(log types.Log) (*JBERC20Transfer, error) {
	event := new(JBERC20Transfer)
	if err := _JBERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
