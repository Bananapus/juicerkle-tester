// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package JBPermissions

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

// JBPermissionsData is an auto generated low-level Go binding around an user-defined struct.
type JBPermissionsData struct {
	Operator      common.Address
	ProjectId     *big.Int
	PermissionIds []*big.Int
}

// JBPermissionsMetaData contains all meta data concerning the JBPermissions contract.
var JBPermissionsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"hasPermission\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permissionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasPermissions\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permissionIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionsOf\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPermissionsFor\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionsData\",\"type\":\"tuple\",\"internalType\":\"structJBPermissionsData\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permissionIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OperatorPermissionsSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"permissionIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"packed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"PERMISSION_ID_OUT_OF_BOUNDS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UNAUTHORIZED\",\"inputs\":[]}]",
}

// JBPermissionsABI is the input ABI used to generate the binding from.
// Deprecated: Use JBPermissionsMetaData.ABI instead.
var JBPermissionsABI = JBPermissionsMetaData.ABI

// JBPermissions is an auto generated Go binding around an Ethereum contract.
type JBPermissions struct {
	JBPermissionsCaller     // Read-only binding to the contract
	JBPermissionsTransactor // Write-only binding to the contract
	JBPermissionsFilterer   // Log filterer for contract events
}

// JBPermissionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type JBPermissionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBPermissionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JBPermissionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBPermissionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JBPermissionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBPermissionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JBPermissionsSession struct {
	Contract     *JBPermissions    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JBPermissionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JBPermissionsCallerSession struct {
	Contract *JBPermissionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// JBPermissionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JBPermissionsTransactorSession struct {
	Contract     *JBPermissionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// JBPermissionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type JBPermissionsRaw struct {
	Contract *JBPermissions // Generic contract binding to access the raw methods on
}

// JBPermissionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JBPermissionsCallerRaw struct {
	Contract *JBPermissionsCaller // Generic read-only contract binding to access the raw methods on
}

// JBPermissionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JBPermissionsTransactorRaw struct {
	Contract *JBPermissionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJBPermissions creates a new instance of JBPermissions, bound to a specific deployed contract.
func NewJBPermissions(address common.Address, backend bind.ContractBackend) (*JBPermissions, error) {
	contract, err := bindJBPermissions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JBPermissions{JBPermissionsCaller: JBPermissionsCaller{contract: contract}, JBPermissionsTransactor: JBPermissionsTransactor{contract: contract}, JBPermissionsFilterer: JBPermissionsFilterer{contract: contract}}, nil
}

// NewJBPermissionsCaller creates a new read-only instance of JBPermissions, bound to a specific deployed contract.
func NewJBPermissionsCaller(address common.Address, caller bind.ContractCaller) (*JBPermissionsCaller, error) {
	contract, err := bindJBPermissions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JBPermissionsCaller{contract: contract}, nil
}

// NewJBPermissionsTransactor creates a new write-only instance of JBPermissions, bound to a specific deployed contract.
func NewJBPermissionsTransactor(address common.Address, transactor bind.ContractTransactor) (*JBPermissionsTransactor, error) {
	contract, err := bindJBPermissions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JBPermissionsTransactor{contract: contract}, nil
}

// NewJBPermissionsFilterer creates a new log filterer instance of JBPermissions, bound to a specific deployed contract.
func NewJBPermissionsFilterer(address common.Address, filterer bind.ContractFilterer) (*JBPermissionsFilterer, error) {
	contract, err := bindJBPermissions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JBPermissionsFilterer{contract: contract}, nil
}

// bindJBPermissions binds a generic wrapper to an already deployed contract.
func bindJBPermissions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JBPermissionsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JBPermissions *JBPermissionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JBPermissions.Contract.JBPermissionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JBPermissions *JBPermissionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBPermissions.Contract.JBPermissionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JBPermissions *JBPermissionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JBPermissions.Contract.JBPermissionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JBPermissions *JBPermissionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JBPermissions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JBPermissions *JBPermissionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBPermissions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JBPermissions *JBPermissionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JBPermissions.Contract.contract.Transact(opts, method, params...)
}

// HasPermission is a free data retrieval call binding the contract method 0xc161c93f.
//
// Solidity: function hasPermission(address operator, address account, uint256 projectId, uint256 permissionId) view returns(bool)
func (_JBPermissions *JBPermissionsCaller) HasPermission(opts *bind.CallOpts, operator common.Address, account common.Address, projectId *big.Int, permissionId *big.Int) (bool, error) {
	var out []interface{}
	err := _JBPermissions.contract.Call(opts, &out, "hasPermission", operator, account, projectId, permissionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0xc161c93f.
//
// Solidity: function hasPermission(address operator, address account, uint256 projectId, uint256 permissionId) view returns(bool)
func (_JBPermissions *JBPermissionsSession) HasPermission(operator common.Address, account common.Address, projectId *big.Int, permissionId *big.Int) (bool, error) {
	return _JBPermissions.Contract.HasPermission(&_JBPermissions.CallOpts, operator, account, projectId, permissionId)
}

// HasPermission is a free data retrieval call binding the contract method 0xc161c93f.
//
// Solidity: function hasPermission(address operator, address account, uint256 projectId, uint256 permissionId) view returns(bool)
func (_JBPermissions *JBPermissionsCallerSession) HasPermission(operator common.Address, account common.Address, projectId *big.Int, permissionId *big.Int) (bool, error) {
	return _JBPermissions.Contract.HasPermission(&_JBPermissions.CallOpts, operator, account, projectId, permissionId)
}

// HasPermissions is a free data retrieval call binding the contract method 0x0f5932f0.
//
// Solidity: function hasPermissions(address operator, address account, uint256 projectId, uint256[] permissionIds) view returns(bool)
func (_JBPermissions *JBPermissionsCaller) HasPermissions(opts *bind.CallOpts, operator common.Address, account common.Address, projectId *big.Int, permissionIds []*big.Int) (bool, error) {
	var out []interface{}
	err := _JBPermissions.contract.Call(opts, &out, "hasPermissions", operator, account, projectId, permissionIds)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermissions is a free data retrieval call binding the contract method 0x0f5932f0.
//
// Solidity: function hasPermissions(address operator, address account, uint256 projectId, uint256[] permissionIds) view returns(bool)
func (_JBPermissions *JBPermissionsSession) HasPermissions(operator common.Address, account common.Address, projectId *big.Int, permissionIds []*big.Int) (bool, error) {
	return _JBPermissions.Contract.HasPermissions(&_JBPermissions.CallOpts, operator, account, projectId, permissionIds)
}

// HasPermissions is a free data retrieval call binding the contract method 0x0f5932f0.
//
// Solidity: function hasPermissions(address operator, address account, uint256 projectId, uint256[] permissionIds) view returns(bool)
func (_JBPermissions *JBPermissionsCallerSession) HasPermissions(operator common.Address, account common.Address, projectId *big.Int, permissionIds []*big.Int) (bool, error) {
	return _JBPermissions.Contract.HasPermissions(&_JBPermissions.CallOpts, operator, account, projectId, permissionIds)
}

// PermissionsOf is a free data retrieval call binding the contract method 0x80deb230.
//
// Solidity: function permissionsOf(address operator, address account, uint256 projectId) view returns(uint256)
func (_JBPermissions *JBPermissionsCaller) PermissionsOf(opts *bind.CallOpts, operator common.Address, account common.Address, projectId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _JBPermissions.contract.Call(opts, &out, "permissionsOf", operator, account, projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PermissionsOf is a free data retrieval call binding the contract method 0x80deb230.
//
// Solidity: function permissionsOf(address operator, address account, uint256 projectId) view returns(uint256)
func (_JBPermissions *JBPermissionsSession) PermissionsOf(operator common.Address, account common.Address, projectId *big.Int) (*big.Int, error) {
	return _JBPermissions.Contract.PermissionsOf(&_JBPermissions.CallOpts, operator, account, projectId)
}

// PermissionsOf is a free data retrieval call binding the contract method 0x80deb230.
//
// Solidity: function permissionsOf(address operator, address account, uint256 projectId) view returns(uint256)
func (_JBPermissions *JBPermissionsCallerSession) PermissionsOf(operator common.Address, account common.Address, projectId *big.Int) (*big.Int, error) {
	return _JBPermissions.Contract.PermissionsOf(&_JBPermissions.CallOpts, operator, account, projectId)
}

// SetPermissionsFor is a paid mutator transaction binding the contract method 0x05d2912e.
//
// Solidity: function setPermissionsFor(address account, (address,uint256,uint256[]) permissionsData) returns()
func (_JBPermissions *JBPermissionsTransactor) SetPermissionsFor(opts *bind.TransactOpts, account common.Address, permissionsData JBPermissionsData) (*types.Transaction, error) {
	return _JBPermissions.contract.Transact(opts, "setPermissionsFor", account, permissionsData)
}

// SetPermissionsFor is a paid mutator transaction binding the contract method 0x05d2912e.
//
// Solidity: function setPermissionsFor(address account, (address,uint256,uint256[]) permissionsData) returns()
func (_JBPermissions *JBPermissionsSession) SetPermissionsFor(account common.Address, permissionsData JBPermissionsData) (*types.Transaction, error) {
	return _JBPermissions.Contract.SetPermissionsFor(&_JBPermissions.TransactOpts, account, permissionsData)
}

// SetPermissionsFor is a paid mutator transaction binding the contract method 0x05d2912e.
//
// Solidity: function setPermissionsFor(address account, (address,uint256,uint256[]) permissionsData) returns()
func (_JBPermissions *JBPermissionsTransactorSession) SetPermissionsFor(account common.Address, permissionsData JBPermissionsData) (*types.Transaction, error) {
	return _JBPermissions.Contract.SetPermissionsFor(&_JBPermissions.TransactOpts, account, permissionsData)
}

// JBPermissionsOperatorPermissionsSetIterator is returned from FilterOperatorPermissionsSet and is used to iterate over the raw logs and unpacked data for OperatorPermissionsSet events raised by the JBPermissions contract.
type JBPermissionsOperatorPermissionsSetIterator struct {
	Event *JBPermissionsOperatorPermissionsSet // Event containing the contract specifics and raw log

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
func (it *JBPermissionsOperatorPermissionsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBPermissionsOperatorPermissionsSet)
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
		it.Event = new(JBPermissionsOperatorPermissionsSet)
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
func (it *JBPermissionsOperatorPermissionsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBPermissionsOperatorPermissionsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBPermissionsOperatorPermissionsSet represents a OperatorPermissionsSet event raised by the JBPermissions contract.
type JBPermissionsOperatorPermissionsSet struct {
	Operator      common.Address
	Account       common.Address
	ProjectId     *big.Int
	PermissionIds []*big.Int
	Packed        *big.Int
	Caller        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorPermissionsSet is a free log retrieval operation binding the contract event 0xf8cdf10a0ce6189c6d069e97320ef146e2bcbd32d6c0b86dcbfbf3bd8578fa70.
//
// Solidity: event OperatorPermissionsSet(address indexed operator, address indexed account, uint256 indexed projectId, uint256[] permissionIds, uint256 packed, address caller)
func (_JBPermissions *JBPermissionsFilterer) FilterOperatorPermissionsSet(opts *bind.FilterOpts, operator []common.Address, account []common.Address, projectId []*big.Int) (*JBPermissionsOperatorPermissionsSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBPermissions.contract.FilterLogs(opts, "OperatorPermissionsSet", operatorRule, accountRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return &JBPermissionsOperatorPermissionsSetIterator{contract: _JBPermissions.contract, event: "OperatorPermissionsSet", logs: logs, sub: sub}, nil
}

// WatchOperatorPermissionsSet is a free log subscription operation binding the contract event 0xf8cdf10a0ce6189c6d069e97320ef146e2bcbd32d6c0b86dcbfbf3bd8578fa70.
//
// Solidity: event OperatorPermissionsSet(address indexed operator, address indexed account, uint256 indexed projectId, uint256[] permissionIds, uint256 packed, address caller)
func (_JBPermissions *JBPermissionsFilterer) WatchOperatorPermissionsSet(opts *bind.WatchOpts, sink chan<- *JBPermissionsOperatorPermissionsSet, operator []common.Address, account []common.Address, projectId []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _JBPermissions.contract.WatchLogs(opts, "OperatorPermissionsSet", operatorRule, accountRule, projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBPermissionsOperatorPermissionsSet)
				if err := _JBPermissions.contract.UnpackLog(event, "OperatorPermissionsSet", log); err != nil {
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

// ParseOperatorPermissionsSet is a log parse operation binding the contract event 0xf8cdf10a0ce6189c6d069e97320ef146e2bcbd32d6c0b86dcbfbf3bd8578fa70.
//
// Solidity: event OperatorPermissionsSet(address indexed operator, address indexed account, uint256 indexed projectId, uint256[] permissionIds, uint256 packed, address caller)
func (_JBPermissions *JBPermissionsFilterer) ParseOperatorPermissionsSet(log types.Log) (*JBPermissionsOperatorPermissionsSet, error) {
	event := new(JBPermissionsOperatorPermissionsSet)
	if err := _JBPermissions.contract.UnpackLog(event, "OperatorPermissionsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
