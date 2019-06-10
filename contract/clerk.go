// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ClerkABI is the input ABI used to generate the binding from.
const ClerkABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeAddressesFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vita\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addAddressesToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receivers\",\"type\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"vitaTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Clerk is an auto generated Go binding around an Ethereum contract.
type Clerk struct {
	ClerkCaller     // Read-only binding to the contract
	ClerkTransactor // Write-only binding to the contract
	ClerkFilterer   // Log filterer for contract events
}

// ClerkCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClerkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClerkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClerkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClerkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClerkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClerkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClerkSession struct {
	Contract     *Clerk            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClerkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClerkCallerSession struct {
	Contract *ClerkCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClerkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClerkTransactorSession struct {
	Contract     *ClerkTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClerkRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClerkRaw struct {
	Contract *Clerk // Generic contract binding to access the raw methods on
}

// ClerkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClerkCallerRaw struct {
	Contract *ClerkCaller // Generic read-only contract binding to access the raw methods on
}

// ClerkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClerkTransactorRaw struct {
	Contract *ClerkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClerk creates a new instance of Clerk, bound to a specific deployed contract.
func NewClerk(address common.Address, backend bind.ContractBackend) (*Clerk, error) {
	contract, err := bindClerk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clerk{ClerkCaller: ClerkCaller{contract: contract}, ClerkTransactor: ClerkTransactor{contract: contract}, ClerkFilterer: ClerkFilterer{contract: contract}}, nil
}

// NewClerkCaller creates a new read-only instance of Clerk, bound to a specific deployed contract.
func NewClerkCaller(address common.Address, caller bind.ContractCaller) (*ClerkCaller, error) {
	contract, err := bindClerk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClerkCaller{contract: contract}, nil
}

// NewClerkTransactor creates a new write-only instance of Clerk, bound to a specific deployed contract.
func NewClerkTransactor(address common.Address, transactor bind.ContractTransactor) (*ClerkTransactor, error) {
	contract, err := bindClerk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClerkTransactor{contract: contract}, nil
}

// NewClerkFilterer creates a new log filterer instance of Clerk, bound to a specific deployed contract.
func NewClerkFilterer(address common.Address, filterer bind.ContractFilterer) (*ClerkFilterer, error) {
	contract, err := bindClerk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClerkFilterer{contract: contract}, nil
}

// bindClerk binds a generic wrapper to an already deployed contract.
func bindClerk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClerkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clerk *ClerkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Clerk.Contract.ClerkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clerk *ClerkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clerk.Contract.ClerkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clerk *ClerkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clerk.Contract.ClerkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clerk *ClerkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Clerk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clerk *ClerkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clerk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clerk *ClerkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clerk.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Clerk *ClerkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Clerk.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Clerk *ClerkSession) Owner() (common.Address, error) {
	return _Clerk.Contract.Owner(&_Clerk.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Clerk *ClerkCallerSession) Owner() (common.Address, error) {
	return _Clerk.Contract.Owner(&_Clerk.CallOpts)
}

// Vita is a free data retrieval call binding the contract method 0x393d9bb3.
//
// Solidity: function vita() constant returns(address)
func (_Clerk *ClerkCaller) Vita(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Clerk.contract.Call(opts, out, "vita")
	return *ret0, err
}

// Vita is a free data retrieval call binding the contract method 0x393d9bb3.
//
// Solidity: function vita() constant returns(address)
func (_Clerk *ClerkSession) Vita() (common.Address, error) {
	return _Clerk.Contract.Vita(&_Clerk.CallOpts)
}

// Vita is a free data retrieval call binding the contract method 0x393d9bb3.
//
// Solidity: function vita() constant returns(address)
func (_Clerk *ClerkCallerSession) Vita() (common.Address, error) {
	return _Clerk.Contract.Vita(&_Clerk.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) constant returns(bool)
func (_Clerk *ClerkCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Clerk.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) constant returns(bool)
func (_Clerk *ClerkSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Clerk.Contract.Whitelist(&_Clerk.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) constant returns(bool)
func (_Clerk *ClerkCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Clerk.Contract.Whitelist(&_Clerk.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(address addr) returns(bool success)
func (_Clerk *ClerkTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Clerk.contract.Transact(opts, "addAddressToWhitelist", addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(address addr) returns(bool success)
func (_Clerk *ClerkSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Clerk.Contract.AddAddressToWhitelist(&_Clerk.TransactOpts, addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(address addr) returns(bool success)
func (_Clerk *ClerkTransactorSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Clerk.Contract.AddAddressToWhitelist(&_Clerk.TransactOpts, addr)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(address[] addrs) returns(bool success)
func (_Clerk *ClerkTransactor) AddAddressesToWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Clerk.contract.Transact(opts, "addAddressesToWhitelist", addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(address[] addrs) returns(bool success)
func (_Clerk *ClerkSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Clerk.Contract.AddAddressesToWhitelist(&_Clerk.TransactOpts, addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(address[] addrs) returns(bool success)
func (_Clerk *ClerkTransactorSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Clerk.Contract.AddAddressesToWhitelist(&_Clerk.TransactOpts, addrs)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Clerk *ClerkTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clerk.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Clerk *ClerkSession) Claim() (*types.Transaction, error) {
	return _Clerk.Contract.Claim(&_Clerk.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Clerk *ClerkTransactorSession) Claim() (*types.Transaction, error) {
	return _Clerk.Contract.Claim(&_Clerk.TransactOpts)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(address addr) returns(bool success)
func (_Clerk *ClerkTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Clerk.contract.Transact(opts, "removeAddressFromWhitelist", addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(address addr) returns(bool success)
func (_Clerk *ClerkSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Clerk.Contract.RemoveAddressFromWhitelist(&_Clerk.TransactOpts, addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(address addr) returns(bool success)
func (_Clerk *ClerkTransactorSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Clerk.Contract.RemoveAddressFromWhitelist(&_Clerk.TransactOpts, addr)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(address[] addrs) returns(bool success)
func (_Clerk *ClerkTransactor) RemoveAddressesFromWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Clerk.contract.Transact(opts, "removeAddressesFromWhitelist", addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(address[] addrs) returns(bool success)
func (_Clerk *ClerkSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Clerk.Contract.RemoveAddressesFromWhitelist(&_Clerk.TransactOpts, addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(address[] addrs) returns(bool success)
func (_Clerk *ClerkTransactorSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Clerk.Contract.RemoveAddressesFromWhitelist(&_Clerk.TransactOpts, addrs)
}

// Transfer is a paid mutator transaction binding the contract method 0xffc3a769.
//
// Solidity: function transfer(address[] receivers, uint256[] amounts) returns()
func (_Clerk *ClerkTransactor) Transfer(opts *bind.TransactOpts, receivers []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Clerk.contract.Transact(opts, "transfer", receivers, amounts)
}

// Transfer is a paid mutator transaction binding the contract method 0xffc3a769.
//
// Solidity: function transfer(address[] receivers, uint256[] amounts) returns()
func (_Clerk *ClerkSession) Transfer(receivers []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Clerk.Contract.Transfer(&_Clerk.TransactOpts, receivers, amounts)
}

// Transfer is a paid mutator transaction binding the contract method 0xffc3a769.
//
// Solidity: function transfer(address[] receivers, uint256[] amounts) returns()
func (_Clerk *ClerkTransactorSession) Transfer(receivers []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Clerk.Contract.Transfer(&_Clerk.TransactOpts, receivers, amounts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clerk *ClerkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Clerk.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clerk *ClerkSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Clerk.Contract.TransferOwnership(&_Clerk.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clerk *ClerkTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Clerk.Contract.TransferOwnership(&_Clerk.TransactOpts, newOwner)
}

// ClerkOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Clerk contract.
type ClerkOwnershipTransferredIterator struct {
	Event *ClerkOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClerkOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClerkOwnershipTransferred)
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
		it.Event = new(ClerkOwnershipTransferred)
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
func (it *ClerkOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClerkOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClerkOwnershipTransferred represents a OwnershipTransferred event raised by the Clerk contract.
type ClerkOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clerk *ClerkFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClerkOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clerk.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClerkOwnershipTransferredIterator{contract: _Clerk.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clerk *ClerkFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClerkOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clerk.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClerkOwnershipTransferred)
				if err := _Clerk.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ClerkWhitelistedAddressAddedIterator is returned from FilterWhitelistedAddressAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAddressAdded events raised by the Clerk contract.
type ClerkWhitelistedAddressAddedIterator struct {
	Event *ClerkWhitelistedAddressAdded // Event containing the contract specifics and raw log

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
func (it *ClerkWhitelistedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClerkWhitelistedAddressAdded)
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
		it.Event = new(ClerkWhitelistedAddressAdded)
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
func (it *ClerkWhitelistedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClerkWhitelistedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClerkWhitelistedAddressAdded represents a WhitelistedAddressAdded event raised by the Clerk contract.
type ClerkWhitelistedAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressAdded is a free log retrieval operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: event WhitelistedAddressAdded(address addr)
func (_Clerk *ClerkFilterer) FilterWhitelistedAddressAdded(opts *bind.FilterOpts) (*ClerkWhitelistedAddressAddedIterator, error) {

	logs, sub, err := _Clerk.contract.FilterLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return &ClerkWhitelistedAddressAddedIterator{contract: _Clerk.contract, event: "WhitelistedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressAdded is a free log subscription operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: event WhitelistedAddressAdded(address addr)
func (_Clerk *ClerkFilterer) WatchWhitelistedAddressAdded(opts *bind.WatchOpts, sink chan<- *ClerkWhitelistedAddressAdded) (event.Subscription, error) {

	logs, sub, err := _Clerk.contract.WatchLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClerkWhitelistedAddressAdded)
				if err := _Clerk.contract.UnpackLog(event, "WhitelistedAddressAdded", log); err != nil {
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

// ClerkWhitelistedAddressRemovedIterator is returned from FilterWhitelistedAddressRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedAddressRemoved events raised by the Clerk contract.
type ClerkWhitelistedAddressRemovedIterator struct {
	Event *ClerkWhitelistedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *ClerkWhitelistedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClerkWhitelistedAddressRemoved)
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
		it.Event = new(ClerkWhitelistedAddressRemoved)
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
func (it *ClerkWhitelistedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClerkWhitelistedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClerkWhitelistedAddressRemoved represents a WhitelistedAddressRemoved event raised by the Clerk contract.
type ClerkWhitelistedAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressRemoved is a free log retrieval operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(address addr)
func (_Clerk *ClerkFilterer) FilterWhitelistedAddressRemoved(opts *bind.FilterOpts) (*ClerkWhitelistedAddressRemovedIterator, error) {

	logs, sub, err := _Clerk.contract.FilterLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return &ClerkWhitelistedAddressRemovedIterator{contract: _Clerk.contract, event: "WhitelistedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressRemoved is a free log subscription operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(address addr)
func (_Clerk *ClerkFilterer) WatchWhitelistedAddressRemoved(opts *bind.WatchOpts, sink chan<- *ClerkWhitelistedAddressRemoved) (event.Subscription, error) {

	logs, sub, err := _Clerk.contract.WatchLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClerkWhitelistedAddressRemoved)
				if err := _Clerk.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
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
