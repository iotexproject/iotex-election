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

// RotatableVPSABI is the input ABI used to generate the binding from.
const RotatableVPSABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_voters\",\"type\":\"address[]\"},{\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"updateVotingPowers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"powerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeAddressesFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newViewID\",\"type\":\"uint256\"}],\"name\":\"rotate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inactiveViewID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"viewID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activeVPS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activeVPSIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inactiveVPS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalPower\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addAddressesToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"voters\",\"outputs\":[{\"name\":\"voters_\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voters\",\"type\":\"address[]\"}],\"name\":\"powersOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addrs\",\"type\":\"address[]\"},{\"name\":\"_viewIDs\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"SetVotingPower\",\"type\":\"event\"}]"

// RotatableVPS is an auto generated Go binding around an Ethereum contract.
type RotatableVPS struct {
	RotatableVPSCaller     // Read-only binding to the contract
	RotatableVPSTransactor // Write-only binding to the contract
	RotatableVPSFilterer   // Log filterer for contract events
}

// RotatableVPSCaller is an auto generated read-only Go binding around an Ethereum contract.
type RotatableVPSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RotatableVPSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RotatableVPSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RotatableVPSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RotatableVPSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RotatableVPSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RotatableVPSSession struct {
	Contract     *RotatableVPS     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RotatableVPSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RotatableVPSCallerSession struct {
	Contract *RotatableVPSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RotatableVPSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RotatableVPSTransactorSession struct {
	Contract     *RotatableVPSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RotatableVPSRaw is an auto generated low-level Go binding around an Ethereum contract.
type RotatableVPSRaw struct {
	Contract *RotatableVPS // Generic contract binding to access the raw methods on
}

// RotatableVPSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RotatableVPSCallerRaw struct {
	Contract *RotatableVPSCaller // Generic read-only contract binding to access the raw methods on
}

// RotatableVPSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RotatableVPSTransactorRaw struct {
	Contract *RotatableVPSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRotatableVPS creates a new instance of RotatableVPS, bound to a specific deployed contract.
func NewRotatableVPS(address common.Address, backend bind.ContractBackend) (*RotatableVPS, error) {
	contract, err := bindRotatableVPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RotatableVPS{RotatableVPSCaller: RotatableVPSCaller{contract: contract}, RotatableVPSTransactor: RotatableVPSTransactor{contract: contract}, RotatableVPSFilterer: RotatableVPSFilterer{contract: contract}}, nil
}

// NewRotatableVPSCaller creates a new read-only instance of RotatableVPS, bound to a specific deployed contract.
func NewRotatableVPSCaller(address common.Address, caller bind.ContractCaller) (*RotatableVPSCaller, error) {
	contract, err := bindRotatableVPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RotatableVPSCaller{contract: contract}, nil
}

// NewRotatableVPSTransactor creates a new write-only instance of RotatableVPS, bound to a specific deployed contract.
func NewRotatableVPSTransactor(address common.Address, transactor bind.ContractTransactor) (*RotatableVPSTransactor, error) {
	contract, err := bindRotatableVPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RotatableVPSTransactor{contract: contract}, nil
}

// NewRotatableVPSFilterer creates a new log filterer instance of RotatableVPS, bound to a specific deployed contract.
func NewRotatableVPSFilterer(address common.Address, filterer bind.ContractFilterer) (*RotatableVPSFilterer, error) {
	contract, err := bindRotatableVPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RotatableVPSFilterer{contract: contract}, nil
}

// bindRotatableVPS binds a generic wrapper to an already deployed contract.
func bindRotatableVPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RotatableVPSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RotatableVPS *RotatableVPSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RotatableVPS.Contract.RotatableVPSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RotatableVPS *RotatableVPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RotatableVPS.Contract.RotatableVPSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RotatableVPS *RotatableVPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RotatableVPS.Contract.RotatableVPSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RotatableVPS *RotatableVPSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RotatableVPS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RotatableVPS *RotatableVPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RotatableVPS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RotatableVPS *RotatableVPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RotatableVPS.Contract.contract.Transact(opts, method, params...)
}

// ActiveVPS is a free data retrieval call binding the contract method 0x90332e4d.
//
// Solidity: function activeVPS() constant returns(address)
func (_RotatableVPS *RotatableVPSCaller) ActiveVPS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "activeVPS")
	return *ret0, err
}

// ActiveVPS is a free data retrieval call binding the contract method 0x90332e4d.
//
// Solidity: function activeVPS() constant returns(address)
func (_RotatableVPS *RotatableVPSSession) ActiveVPS() (common.Address, error) {
	return _RotatableVPS.Contract.ActiveVPS(&_RotatableVPS.CallOpts)
}

// ActiveVPS is a free data retrieval call binding the contract method 0x90332e4d.
//
// Solidity: function activeVPS() constant returns(address)
func (_RotatableVPS *RotatableVPSCallerSession) ActiveVPS() (common.Address, error) {
	return _RotatableVPS.Contract.ActiveVPS(&_RotatableVPS.CallOpts)
}

// ActiveVPSIndex is a free data retrieval call binding the contract method 0xcec455c7.
//
// Solidity: function activeVPSIndex() constant returns(uint256)
func (_RotatableVPS *RotatableVPSCaller) ActiveVPSIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "activeVPSIndex")
	return *ret0, err
}

// ActiveVPSIndex is a free data retrieval call binding the contract method 0xcec455c7.
//
// Solidity: function activeVPSIndex() constant returns(uint256)
func (_RotatableVPS *RotatableVPSSession) ActiveVPSIndex() (*big.Int, error) {
	return _RotatableVPS.Contract.ActiveVPSIndex(&_RotatableVPS.CallOpts)
}

// ActiveVPSIndex is a free data retrieval call binding the contract method 0xcec455c7.
//
// Solidity: function activeVPSIndex() constant returns(uint256)
func (_RotatableVPS *RotatableVPSCallerSession) ActiveVPSIndex() (*big.Int, error) {
	return _RotatableVPS.Contract.ActiveVPSIndex(&_RotatableVPS.CallOpts)
}

// InactiveVPS is a free data retrieval call binding the contract method 0xd5a86bf2.
//
// Solidity: function inactiveVPS() constant returns(address)
func (_RotatableVPS *RotatableVPSCaller) InactiveVPS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "inactiveVPS")
	return *ret0, err
}

// InactiveVPS is a free data retrieval call binding the contract method 0xd5a86bf2.
//
// Solidity: function inactiveVPS() constant returns(address)
func (_RotatableVPS *RotatableVPSSession) InactiveVPS() (common.Address, error) {
	return _RotatableVPS.Contract.InactiveVPS(&_RotatableVPS.CallOpts)
}

// InactiveVPS is a free data retrieval call binding the contract method 0xd5a86bf2.
//
// Solidity: function inactiveVPS() constant returns(address)
func (_RotatableVPS *RotatableVPSCallerSession) InactiveVPS() (common.Address, error) {
	return _RotatableVPS.Contract.InactiveVPS(&_RotatableVPS.CallOpts)
}

// InactiveViewID is a free data retrieval call binding the contract method 0x63aedb0a.
//
// Solidity: function inactiveViewID() constant returns(uint256)
func (_RotatableVPS *RotatableVPSCaller) InactiveViewID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "inactiveViewID")
	return *ret0, err
}

// InactiveViewID is a free data retrieval call binding the contract method 0x63aedb0a.
//
// Solidity: function inactiveViewID() constant returns(uint256)
func (_RotatableVPS *RotatableVPSSession) InactiveViewID() (*big.Int, error) {
	return _RotatableVPS.Contract.InactiveViewID(&_RotatableVPS.CallOpts)
}

// InactiveViewID is a free data retrieval call binding the contract method 0x63aedb0a.
//
// Solidity: function inactiveViewID() constant returns(uint256)
func (_RotatableVPS *RotatableVPSCallerSession) InactiveViewID() (*big.Int, error) {
	return _RotatableVPS.Contract.InactiveViewID(&_RotatableVPS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RotatableVPS *RotatableVPSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RotatableVPS *RotatableVPSSession) Owner() (common.Address, error) {
	return _RotatableVPS.Contract.Owner(&_RotatableVPS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RotatableVPS *RotatableVPSCallerSession) Owner() (common.Address, error) {
	return _RotatableVPS.Contract.Owner(&_RotatableVPS.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RotatableVPS *RotatableVPSCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RotatableVPS *RotatableVPSSession) Paused() (bool, error) {
	return _RotatableVPS.Contract.Paused(&_RotatableVPS.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RotatableVPS *RotatableVPSCallerSession) Paused() (bool, error) {
	return _RotatableVPS.Contract.Paused(&_RotatableVPS.CallOpts)
}

// PowerOf is a free data retrieval call binding the contract method 0x1ac84690.
//
// Solidity: function powerOf(_voter address) constant returns(uint256)
func (_RotatableVPS *RotatableVPSCaller) PowerOf(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "powerOf", _voter)
	return *ret0, err
}

// PowerOf is a free data retrieval call binding the contract method 0x1ac84690.
//
// Solidity: function powerOf(_voter address) constant returns(uint256)
func (_RotatableVPS *RotatableVPSSession) PowerOf(_voter common.Address) (*big.Int, error) {
	return _RotatableVPS.Contract.PowerOf(&_RotatableVPS.CallOpts, _voter)
}

// PowerOf is a free data retrieval call binding the contract method 0x1ac84690.
//
// Solidity: function powerOf(_voter address) constant returns(uint256)
func (_RotatableVPS *RotatableVPSCallerSession) PowerOf(_voter common.Address) (*big.Int, error) {
	return _RotatableVPS.Contract.PowerOf(&_RotatableVPS.CallOpts, _voter)
}

// PowersOf is a free data retrieval call binding the contract method 0xff82c4ca.
//
// Solidity: function powersOf(_voters address[]) constant returns(uint256[])
func (_RotatableVPS *RotatableVPSCaller) PowersOf(opts *bind.CallOpts, _voters []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "powersOf", _voters)
	return *ret0, err
}

// PowersOf is a free data retrieval call binding the contract method 0xff82c4ca.
//
// Solidity: function powersOf(_voters address[]) constant returns(uint256[])
func (_RotatableVPS *RotatableVPSSession) PowersOf(_voters []common.Address) ([]*big.Int, error) {
	return _RotatableVPS.Contract.PowersOf(&_RotatableVPS.CallOpts, _voters)
}

// PowersOf is a free data retrieval call binding the contract method 0xff82c4ca.
//
// Solidity: function powersOf(_voters address[]) constant returns(uint256[])
func (_RotatableVPS *RotatableVPSCallerSession) PowersOf(_voters []common.Address) ([]*big.Int, error) {
	return _RotatableVPS.Contract.PowersOf(&_RotatableVPS.CallOpts, _voters)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() constant returns(uint256)
func (_RotatableVPS *RotatableVPSCaller) TotalPower(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "totalPower")
	return *ret0, err
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() constant returns(uint256)
func (_RotatableVPS *RotatableVPSSession) TotalPower() (*big.Int, error) {
	return _RotatableVPS.Contract.TotalPower(&_RotatableVPS.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() constant returns(uint256)
func (_RotatableVPS *RotatableVPSCallerSession) TotalPower() (*big.Int, error) {
	return _RotatableVPS.Contract.TotalPower(&_RotatableVPS.CallOpts)
}

// ViewID is a free data retrieval call binding the contract method 0x8280264b.
//
// Solidity: function viewID() constant returns(uint256)
func (_RotatableVPS *RotatableVPSCaller) ViewID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "viewID")
	return *ret0, err
}

// ViewID is a free data retrieval call binding the contract method 0x8280264b.
//
// Solidity: function viewID() constant returns(uint256)
func (_RotatableVPS *RotatableVPSSession) ViewID() (*big.Int, error) {
	return _RotatableVPS.Contract.ViewID(&_RotatableVPS.CallOpts)
}

// ViewID is a free data retrieval call binding the contract method 0x8280264b.
//
// Solidity: function viewID() constant returns(uint256)
func (_RotatableVPS *RotatableVPSCallerSession) ViewID() (*big.Int, error) {
	return _RotatableVPS.Contract.ViewID(&_RotatableVPS.CallOpts)
}

// Voters is a free data retrieval call binding the contract method 0xfba00cbd.
//
// Solidity: function voters(_offset uint256, _limit uint256) constant returns(voters_ address[])
func (_RotatableVPS *RotatableVPSCaller) Voters(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "voters", _offset, _limit)
	return *ret0, err
}

// Voters is a free data retrieval call binding the contract method 0xfba00cbd.
//
// Solidity: function voters(_offset uint256, _limit uint256) constant returns(voters_ address[])
func (_RotatableVPS *RotatableVPSSession) Voters(_offset *big.Int, _limit *big.Int) ([]common.Address, error) {
	return _RotatableVPS.Contract.Voters(&_RotatableVPS.CallOpts, _offset, _limit)
}

// Voters is a free data retrieval call binding the contract method 0xfba00cbd.
//
// Solidity: function voters(_offset uint256, _limit uint256) constant returns(voters_ address[])
func (_RotatableVPS *RotatableVPSCallerSession) Voters(_offset *big.Int, _limit *big.Int) ([]common.Address, error) {
	return _RotatableVPS.Contract.Voters(&_RotatableVPS.CallOpts, _offset, _limit)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_RotatableVPS *RotatableVPSCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RotatableVPS.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_RotatableVPS *RotatableVPSSession) Whitelist(arg0 common.Address) (bool, error) {
	return _RotatableVPS.Contract.Whitelist(&_RotatableVPS.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_RotatableVPS *RotatableVPSCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _RotatableVPS.Contract.Whitelist(&_RotatableVPS.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_RotatableVPS *RotatableVPSTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _RotatableVPS.contract.Transact(opts, "addAddressToWhitelist", addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_RotatableVPS *RotatableVPSSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _RotatableVPS.Contract.AddAddressToWhitelist(&_RotatableVPS.TransactOpts, addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_RotatableVPS *RotatableVPSTransactorSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _RotatableVPS.Contract.AddAddressToWhitelist(&_RotatableVPS.TransactOpts, addr)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_RotatableVPS *RotatableVPSTransactor) AddAddressesToWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _RotatableVPS.contract.Transact(opts, "addAddressesToWhitelist", addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_RotatableVPS *RotatableVPSSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _RotatableVPS.Contract.AddAddressesToWhitelist(&_RotatableVPS.TransactOpts, addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_RotatableVPS *RotatableVPSTransactorSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _RotatableVPS.Contract.AddAddressesToWhitelist(&_RotatableVPS.TransactOpts, addrs)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_RotatableVPS *RotatableVPSTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _RotatableVPS.contract.Transact(opts, "removeAddressFromWhitelist", addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_RotatableVPS *RotatableVPSSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _RotatableVPS.Contract.RemoveAddressFromWhitelist(&_RotatableVPS.TransactOpts, addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_RotatableVPS *RotatableVPSTransactorSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _RotatableVPS.Contract.RemoveAddressFromWhitelist(&_RotatableVPS.TransactOpts, addr)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_RotatableVPS *RotatableVPSTransactor) RemoveAddressesFromWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _RotatableVPS.contract.Transact(opts, "removeAddressesFromWhitelist", addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_RotatableVPS *RotatableVPSSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _RotatableVPS.Contract.RemoveAddressesFromWhitelist(&_RotatableVPS.TransactOpts, addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_RotatableVPS *RotatableVPSTransactorSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _RotatableVPS.Contract.RemoveAddressesFromWhitelist(&_RotatableVPS.TransactOpts, addrs)
}

// Rotate is a paid mutator transaction binding the contract method 0x3852f4b0.
//
// Solidity: function rotate(newViewID uint256) returns()
func (_RotatableVPS *RotatableVPSTransactor) Rotate(opts *bind.TransactOpts, newViewID *big.Int) (*types.Transaction, error) {
	return _RotatableVPS.contract.Transact(opts, "rotate", newViewID)
}

// Rotate is a paid mutator transaction binding the contract method 0x3852f4b0.
//
// Solidity: function rotate(newViewID uint256) returns()
func (_RotatableVPS *RotatableVPSSession) Rotate(newViewID *big.Int) (*types.Transaction, error) {
	return _RotatableVPS.Contract.Rotate(&_RotatableVPS.TransactOpts, newViewID)
}

// Rotate is a paid mutator transaction binding the contract method 0x3852f4b0.
//
// Solidity: function rotate(newViewID uint256) returns()
func (_RotatableVPS *RotatableVPSTransactorSession) Rotate(newViewID *big.Int) (*types.Transaction, error) {
	return _RotatableVPS.Contract.Rotate(&_RotatableVPS.TransactOpts, newViewID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RotatableVPS *RotatableVPSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RotatableVPS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RotatableVPS *RotatableVPSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RotatableVPS.Contract.TransferOwnership(&_RotatableVPS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RotatableVPS *RotatableVPSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RotatableVPS.Contract.TransferOwnership(&_RotatableVPS.TransactOpts, newOwner)
}

// UpdateVotingPowers is a paid mutator transaction binding the contract method 0x03083dcf.
//
// Solidity: function updateVotingPowers(_voters address[], _powers uint256[]) returns()
func (_RotatableVPS *RotatableVPSTransactor) UpdateVotingPowers(opts *bind.TransactOpts, _voters []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _RotatableVPS.contract.Transact(opts, "updateVotingPowers", _voters, _powers)
}

// UpdateVotingPowers is a paid mutator transaction binding the contract method 0x03083dcf.
//
// Solidity: function updateVotingPowers(_voters address[], _powers uint256[]) returns()
func (_RotatableVPS *RotatableVPSSession) UpdateVotingPowers(_voters []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _RotatableVPS.Contract.UpdateVotingPowers(&_RotatableVPS.TransactOpts, _voters, _powers)
}

// UpdateVotingPowers is a paid mutator transaction binding the contract method 0x03083dcf.
//
// Solidity: function updateVotingPowers(_voters address[], _powers uint256[]) returns()
func (_RotatableVPS *RotatableVPSTransactorSession) UpdateVotingPowers(_voters []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _RotatableVPS.Contract.UpdateVotingPowers(&_RotatableVPS.TransactOpts, _voters, _powers)
}

// RotatableVPSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RotatableVPS contract.
type RotatableVPSOwnershipTransferredIterator struct {
	Event *RotatableVPSOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RotatableVPSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RotatableVPSOwnershipTransferred)
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
		it.Event = new(RotatableVPSOwnershipTransferred)
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
func (it *RotatableVPSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RotatableVPSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RotatableVPSOwnershipTransferred represents a OwnershipTransferred event raised by the RotatableVPS contract.
type RotatableVPSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RotatableVPS *RotatableVPSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RotatableVPSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RotatableVPS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RotatableVPSOwnershipTransferredIterator{contract: _RotatableVPS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RotatableVPS *RotatableVPSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RotatableVPSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RotatableVPS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RotatableVPSOwnershipTransferred)
				if err := _RotatableVPS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RotatableVPSSetVotingPowerIterator is returned from FilterSetVotingPower and is used to iterate over the raw logs and unpacked data for SetVotingPower events raised by the RotatableVPS contract.
type RotatableVPSSetVotingPowerIterator struct {
	Event *RotatableVPSSetVotingPower // Event containing the contract specifics and raw log

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
func (it *RotatableVPSSetVotingPowerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RotatableVPSSetVotingPower)
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
		it.Event = new(RotatableVPSSetVotingPower)
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
func (it *RotatableVPSSetVotingPowerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RotatableVPSSetVotingPowerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RotatableVPSSetVotingPower represents a SetVotingPower event raised by the RotatableVPS contract.
type RotatableVPSSetVotingPower struct {
	Voter common.Address
	Power *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetVotingPower is a free log retrieval operation binding the contract event 0x048ebfbc43eddab7a7a576fb6a22551fb89c9fa0bb520948a32ec8b1546640d3.
//
// Solidity: e SetVotingPower(voter address, power uint256)
func (_RotatableVPS *RotatableVPSFilterer) FilterSetVotingPower(opts *bind.FilterOpts) (*RotatableVPSSetVotingPowerIterator, error) {

	logs, sub, err := _RotatableVPS.contract.FilterLogs(opts, "SetVotingPower")
	if err != nil {
		return nil, err
	}
	return &RotatableVPSSetVotingPowerIterator{contract: _RotatableVPS.contract, event: "SetVotingPower", logs: logs, sub: sub}, nil
}

// WatchSetVotingPower is a free log subscription operation binding the contract event 0x048ebfbc43eddab7a7a576fb6a22551fb89c9fa0bb520948a32ec8b1546640d3.
//
// Solidity: e SetVotingPower(voter address, power uint256)
func (_RotatableVPS *RotatableVPSFilterer) WatchSetVotingPower(opts *bind.WatchOpts, sink chan<- *RotatableVPSSetVotingPower) (event.Subscription, error) {

	logs, sub, err := _RotatableVPS.contract.WatchLogs(opts, "SetVotingPower")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RotatableVPSSetVotingPower)
				if err := _RotatableVPS.contract.UnpackLog(event, "SetVotingPower", log); err != nil {
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

// RotatableVPSWhitelistedAddressAddedIterator is returned from FilterWhitelistedAddressAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAddressAdded events raised by the RotatableVPS contract.
type RotatableVPSWhitelistedAddressAddedIterator struct {
	Event *RotatableVPSWhitelistedAddressAdded // Event containing the contract specifics and raw log

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
func (it *RotatableVPSWhitelistedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RotatableVPSWhitelistedAddressAdded)
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
		it.Event = new(RotatableVPSWhitelistedAddressAdded)
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
func (it *RotatableVPSWhitelistedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RotatableVPSWhitelistedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RotatableVPSWhitelistedAddressAdded represents a WhitelistedAddressAdded event raised by the RotatableVPS contract.
type RotatableVPSWhitelistedAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressAdded is a free log retrieval operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: e WhitelistedAddressAdded(addr address)
func (_RotatableVPS *RotatableVPSFilterer) FilterWhitelistedAddressAdded(opts *bind.FilterOpts) (*RotatableVPSWhitelistedAddressAddedIterator, error) {

	logs, sub, err := _RotatableVPS.contract.FilterLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return &RotatableVPSWhitelistedAddressAddedIterator{contract: _RotatableVPS.contract, event: "WhitelistedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressAdded is a free log subscription operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: e WhitelistedAddressAdded(addr address)
func (_RotatableVPS *RotatableVPSFilterer) WatchWhitelistedAddressAdded(opts *bind.WatchOpts, sink chan<- *RotatableVPSWhitelistedAddressAdded) (event.Subscription, error) {

	logs, sub, err := _RotatableVPS.contract.WatchLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RotatableVPSWhitelistedAddressAdded)
				if err := _RotatableVPS.contract.UnpackLog(event, "WhitelistedAddressAdded", log); err != nil {
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

// RotatableVPSWhitelistedAddressRemovedIterator is returned from FilterWhitelistedAddressRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedAddressRemoved events raised by the RotatableVPS contract.
type RotatableVPSWhitelistedAddressRemovedIterator struct {
	Event *RotatableVPSWhitelistedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *RotatableVPSWhitelistedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RotatableVPSWhitelistedAddressRemoved)
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
		it.Event = new(RotatableVPSWhitelistedAddressRemoved)
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
func (it *RotatableVPSWhitelistedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RotatableVPSWhitelistedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RotatableVPSWhitelistedAddressRemoved represents a WhitelistedAddressRemoved event raised by the RotatableVPS contract.
type RotatableVPSWhitelistedAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressRemoved is a free log retrieval operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: e WhitelistedAddressRemoved(addr address)
func (_RotatableVPS *RotatableVPSFilterer) FilterWhitelistedAddressRemoved(opts *bind.FilterOpts) (*RotatableVPSWhitelistedAddressRemovedIterator, error) {

	logs, sub, err := _RotatableVPS.contract.FilterLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return &RotatableVPSWhitelistedAddressRemovedIterator{contract: _RotatableVPS.contract, event: "WhitelistedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressRemoved is a free log subscription operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: e WhitelistedAddressRemoved(addr address)
func (_RotatableVPS *RotatableVPSFilterer) WatchWhitelistedAddressRemoved(opts *bind.WatchOpts, sink chan<- *RotatableVPSWhitelistedAddressRemoved) (event.Subscription, error) {

	logs, sub, err := _RotatableVPS.contract.WatchLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RotatableVPSWhitelistedAddressRemoved)
				if err := _RotatableVPS.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
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
