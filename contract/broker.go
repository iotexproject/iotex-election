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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BrokerABI is the input ABI used to generate the binding from.
const BrokerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"round\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getTotalBidsValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeAddressesFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextBidToSettle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"settle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getNumBids\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addAddressesToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getAvailableVitaValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_round\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getBid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_vitaTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"claimedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"burnedAmount\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"VitaBidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vita\",\"type\":\"uint256\"}],\"name\":\"VitaBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"finished\",\"type\":\"bool\"}],\"name\":\"VitaBidsSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"CollateralWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Broker is an auto generated Go binding around an Ethereum contract.
type Broker struct {
	BrokerCaller     // Read-only binding to the contract
	BrokerTransactor // Write-only binding to the contract
	BrokerFilterer   // Log filterer for contract events
}

// BrokerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrokerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrokerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrokerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrokerSession struct {
	Contract     *Broker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrokerCallerSession struct {
	Contract *BrokerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BrokerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrokerTransactorSession struct {
	Contract     *BrokerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrokerRaw struct {
	Contract *Broker // Generic contract binding to access the raw methods on
}

// BrokerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrokerCallerRaw struct {
	Contract *BrokerCaller // Generic read-only contract binding to access the raw methods on
}

// BrokerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrokerTransactorRaw struct {
	Contract *BrokerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBroker creates a new instance of Broker, bound to a specific deployed contract.
func NewBroker(address common.Address, backend bind.ContractBackend) (*Broker, error) {
	contract, err := bindBroker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Broker{BrokerCaller: BrokerCaller{contract: contract}, BrokerTransactor: BrokerTransactor{contract: contract}, BrokerFilterer: BrokerFilterer{contract: contract}}, nil
}

// NewBrokerCaller creates a new read-only instance of Broker, bound to a specific deployed contract.
func NewBrokerCaller(address common.Address, caller bind.ContractCaller) (*BrokerCaller, error) {
	contract, err := bindBroker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerCaller{contract: contract}, nil
}

// NewBrokerTransactor creates a new write-only instance of Broker, bound to a specific deployed contract.
func NewBrokerTransactor(address common.Address, transactor bind.ContractTransactor) (*BrokerTransactor, error) {
	contract, err := bindBroker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerTransactor{contract: contract}, nil
}

// NewBrokerFilterer creates a new log filterer instance of Broker, bound to a specific deployed contract.
func NewBrokerFilterer(address common.Address, filterer bind.ContractFilterer) (*BrokerFilterer, error) {
	contract, err := bindBroker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrokerFilterer{contract: contract}, nil
}

// bindBroker binds a generic wrapper to an already deployed contract.
func bindBroker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrokerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Broker *BrokerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Broker.Contract.BrokerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Broker *BrokerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.Contract.BrokerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Broker *BrokerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Broker.Contract.BrokerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Broker *BrokerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Broker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Broker *BrokerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Broker *BrokerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Broker.Contract.contract.Transact(opts, method, params...)
}

// GetAvailableVitaValue is a free data retrieval call binding the contract method 0xe80804be.
//
// Solidity: function getAvailableVitaValue(uint256 _round) view returns(uint256)
func (_Broker *BrokerCaller) GetAvailableVitaValue(opts *bind.CallOpts, _round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getAvailableVitaValue", _round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableVitaValue is a free data retrieval call binding the contract method 0xe80804be.
//
// Solidity: function getAvailableVitaValue(uint256 _round) view returns(uint256)
func (_Broker *BrokerSession) GetAvailableVitaValue(_round *big.Int) (*big.Int, error) {
	return _Broker.Contract.GetAvailableVitaValue(&_Broker.CallOpts, _round)
}

// GetAvailableVitaValue is a free data retrieval call binding the contract method 0xe80804be.
//
// Solidity: function getAvailableVitaValue(uint256 _round) view returns(uint256)
func (_Broker *BrokerCallerSession) GetAvailableVitaValue(_round *big.Int) (*big.Int, error) {
	return _Broker.Contract.GetAvailableVitaValue(&_Broker.CallOpts, _round)
}

// GetBid is a free data retrieval call binding the contract method 0xeba1b60b.
//
// Solidity: function getBid(uint256 _round, address _address) view returns(uint256)
func (_Broker *BrokerCaller) GetBid(opts *bind.CallOpts, _round *big.Int, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getBid", _round, _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBid is a free data retrieval call binding the contract method 0xeba1b60b.
//
// Solidity: function getBid(uint256 _round, address _address) view returns(uint256)
func (_Broker *BrokerSession) GetBid(_round *big.Int, _address common.Address) (*big.Int, error) {
	return _Broker.Contract.GetBid(&_Broker.CallOpts, _round, _address)
}

// GetBid is a free data retrieval call binding the contract method 0xeba1b60b.
//
// Solidity: function getBid(uint256 _round, address _address) view returns(uint256)
func (_Broker *BrokerCallerSession) GetBid(_round *big.Int, _address common.Address) (*big.Int, error) {
	return _Broker.Contract.GetBid(&_Broker.CallOpts, _round, _address)
}

// GetNumBids is a free data retrieval call binding the contract method 0x9e0a673f.
//
// Solidity: function getNumBids(uint256 _round) view returns(uint256)
func (_Broker *BrokerCaller) GetNumBids(opts *bind.CallOpts, _round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getNumBids", _round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumBids is a free data retrieval call binding the contract method 0x9e0a673f.
//
// Solidity: function getNumBids(uint256 _round) view returns(uint256)
func (_Broker *BrokerSession) GetNumBids(_round *big.Int) (*big.Int, error) {
	return _Broker.Contract.GetNumBids(&_Broker.CallOpts, _round)
}

// GetNumBids is a free data retrieval call binding the contract method 0x9e0a673f.
//
// Solidity: function getNumBids(uint256 _round) view returns(uint256)
func (_Broker *BrokerCallerSession) GetNumBids(_round *big.Int) (*big.Int, error) {
	return _Broker.Contract.GetNumBids(&_Broker.CallOpts, _round)
}

// GetTotalBidsValue is a free data retrieval call binding the contract method 0x23702b34.
//
// Solidity: function getTotalBidsValue(uint256 _round) view returns(uint256)
func (_Broker *BrokerCaller) GetTotalBidsValue(opts *bind.CallOpts, _round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getTotalBidsValue", _round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBidsValue is a free data retrieval call binding the contract method 0x23702b34.
//
// Solidity: function getTotalBidsValue(uint256 _round) view returns(uint256)
func (_Broker *BrokerSession) GetTotalBidsValue(_round *big.Int) (*big.Int, error) {
	return _Broker.Contract.GetTotalBidsValue(&_Broker.CallOpts, _round)
}

// GetTotalBidsValue is a free data retrieval call binding the contract method 0x23702b34.
//
// Solidity: function getTotalBidsValue(uint256 _round) view returns(uint256)
func (_Broker *BrokerCallerSession) GetTotalBidsValue(_round *big.Int) (*big.Int, error) {
	return _Broker.Contract.GetTotalBidsValue(&_Broker.CallOpts, _round)
}

// NextBidToSettle is a free data retrieval call binding the contract method 0x650cc7cc.
//
// Solidity: function nextBidToSettle() view returns(uint256)
func (_Broker *BrokerCaller) NextBidToSettle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "nextBidToSettle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBidToSettle is a free data retrieval call binding the contract method 0x650cc7cc.
//
// Solidity: function nextBidToSettle() view returns(uint256)
func (_Broker *BrokerSession) NextBidToSettle() (*big.Int, error) {
	return _Broker.Contract.NextBidToSettle(&_Broker.CallOpts)
}

// NextBidToSettle is a free data retrieval call binding the contract method 0x650cc7cc.
//
// Solidity: function nextBidToSettle() view returns(uint256)
func (_Broker *BrokerCallerSession) NextBidToSettle() (*big.Int, error) {
	return _Broker.Contract.NextBidToSettle(&_Broker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Broker *BrokerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Broker *BrokerSession) Owner() (common.Address, error) {
	return _Broker.Contract.Owner(&_Broker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Broker *BrokerCallerSession) Owner() (common.Address, error) {
	return _Broker.Contract.Owner(&_Broker.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Broker *BrokerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Broker *BrokerSession) Paused() (bool, error) {
	return _Broker.Contract.Paused(&_Broker.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Broker *BrokerCallerSession) Paused() (bool, error) {
	return _Broker.Contract.Paused(&_Broker.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Broker *BrokerCaller) Round(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "round")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Broker *BrokerSession) Round() (*big.Int, error) {
	return _Broker.Contract.Round(&_Broker.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Broker *BrokerCallerSession) Round() (*big.Int, error) {
	return _Broker.Contract.Round(&_Broker.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Broker *BrokerCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Broker *BrokerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Broker.Contract.Whitelist(&_Broker.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Broker *BrokerCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Broker.Contract.Whitelist(&_Broker.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(address addr) returns(bool success)
func (_Broker *BrokerTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "addAddressToWhitelist", addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(address addr) returns(bool success)
func (_Broker *BrokerSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Broker.Contract.AddAddressToWhitelist(&_Broker.TransactOpts, addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(address addr) returns(bool success)
func (_Broker *BrokerTransactorSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Broker.Contract.AddAddressToWhitelist(&_Broker.TransactOpts, addr)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(address[] addrs) returns(bool success)
func (_Broker *BrokerTransactor) AddAddressesToWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "addAddressesToWhitelist", addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(address[] addrs) returns(bool success)
func (_Broker *BrokerSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Broker.Contract.AddAddressesToWhitelist(&_Broker.TransactOpts, addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(address[] addrs) returns(bool success)
func (_Broker *BrokerTransactorSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Broker.Contract.AddAddressesToWhitelist(&_Broker.TransactOpts, addrs)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Broker *BrokerTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Broker *BrokerSession) Bid() (*types.Transaction, error) {
	return _Broker.Contract.Bid(&_Broker.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Broker *BrokerTransactorSession) Bid() (*types.Transaction, error) {
	return _Broker.Contract.Bid(&_Broker.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Broker *BrokerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Broker *BrokerSession) Pause() (*types.Transaction, error) {
	return _Broker.Contract.Pause(&_Broker.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Broker *BrokerTransactorSession) Pause() (*types.Transaction, error) {
	return _Broker.Contract.Pause(&_Broker.TransactOpts)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(address addr) returns(bool success)
func (_Broker *BrokerTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "removeAddressFromWhitelist", addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(address addr) returns(bool success)
func (_Broker *BrokerSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Broker.Contract.RemoveAddressFromWhitelist(&_Broker.TransactOpts, addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(address addr) returns(bool success)
func (_Broker *BrokerTransactorSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Broker.Contract.RemoveAddressFromWhitelist(&_Broker.TransactOpts, addr)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(address[] addrs) returns(bool success)
func (_Broker *BrokerTransactor) RemoveAddressesFromWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "removeAddressesFromWhitelist", addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(address[] addrs) returns(bool success)
func (_Broker *BrokerSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Broker.Contract.RemoveAddressesFromWhitelist(&_Broker.TransactOpts, addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(address[] addrs) returns(bool success)
func (_Broker *BrokerTransactorSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Broker.Contract.RemoveAddressesFromWhitelist(&_Broker.TransactOpts, addrs)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Broker *BrokerTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Broker *BrokerSession) Reset() (*types.Transaction, error) {
	return _Broker.Contract.Reset(&_Broker.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Broker *BrokerTransactorSession) Reset() (*types.Transaction, error) {
	return _Broker.Contract.Reset(&_Broker.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _count) returns()
func (_Broker *BrokerTransactor) Settle(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "settle", _count)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _count) returns()
func (_Broker *BrokerSession) Settle(_count *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.Settle(&_Broker.TransactOpts, _count)
}

// Settle is a paid mutator transaction binding the contract method 0x8df82800.
//
// Solidity: function settle(uint256 _count) returns()
func (_Broker *BrokerTransactorSession) Settle(_count *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.Settle(&_Broker.TransactOpts, _count)
}

// TransferCollateral is a paid mutator transaction binding the contract method 0x816b1e8f.
//
// Solidity: function transferCollateral(address _receiver, uint256 _amount) returns()
func (_Broker *BrokerTransactor) TransferCollateral(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "transferCollateral", _receiver, _amount)
}

// TransferCollateral is a paid mutator transaction binding the contract method 0x816b1e8f.
//
// Solidity: function transferCollateral(address _receiver, uint256 _amount) returns()
func (_Broker *BrokerSession) TransferCollateral(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.TransferCollateral(&_Broker.TransactOpts, _receiver, _amount)
}

// TransferCollateral is a paid mutator transaction binding the contract method 0x816b1e8f.
//
// Solidity: function transferCollateral(address _receiver, uint256 _amount) returns()
func (_Broker *BrokerTransactorSession) TransferCollateral(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.TransferCollateral(&_Broker.TransactOpts, _receiver, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Broker *BrokerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Broker *BrokerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Broker.Contract.TransferOwnership(&_Broker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Broker *BrokerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Broker.Contract.TransferOwnership(&_Broker.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Broker *BrokerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Broker *BrokerSession) Unpause() (*types.Transaction, error) {
	return _Broker.Contract.Unpause(&_Broker.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Broker *BrokerTransactorSession) Unpause() (*types.Transaction, error) {
	return _Broker.Contract.Unpause(&_Broker.TransactOpts)
}

// BrokerCollateralWithdrawnIterator is returned from FilterCollateralWithdrawn and is used to iterate over the raw logs and unpacked data for CollateralWithdrawn events raised by the Broker contract.
type BrokerCollateralWithdrawnIterator struct {
	Event *BrokerCollateralWithdrawn // Event containing the contract specifics and raw log

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
func (it *BrokerCollateralWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerCollateralWithdrawn)
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
		it.Event = new(BrokerCollateralWithdrawn)
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
func (it *BrokerCollateralWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerCollateralWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerCollateralWithdrawn represents a CollateralWithdrawn event raised by the Broker contract.
type BrokerCollateralWithdrawn struct {
	Receiver   common.Address
	Collateral *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralWithdrawn is a free log retrieval operation binding the contract event 0xc30fcfbcaac9e0deffa719714eaa82396ff506a0d0d0eebe170830177288715d.
//
// Solidity: event CollateralWithdrawn(address indexed receiver, uint256 collateral)
func (_Broker *BrokerFilterer) FilterCollateralWithdrawn(opts *bind.FilterOpts, receiver []common.Address) (*BrokerCollateralWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "CollateralWithdrawn", receiverRule)
	if err != nil {
		return nil, err
	}
	return &BrokerCollateralWithdrawnIterator{contract: _Broker.contract, event: "CollateralWithdrawn", logs: logs, sub: sub}, nil
}

// WatchCollateralWithdrawn is a free log subscription operation binding the contract event 0xc30fcfbcaac9e0deffa719714eaa82396ff506a0d0d0eebe170830177288715d.
//
// Solidity: event CollateralWithdrawn(address indexed receiver, uint256 collateral)
func (_Broker *BrokerFilterer) WatchCollateralWithdrawn(opts *bind.WatchOpts, sink chan<- *BrokerCollateralWithdrawn, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "CollateralWithdrawn", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerCollateralWithdrawn)
				if err := _Broker.contract.UnpackLog(event, "CollateralWithdrawn", log); err != nil {
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

// ParseCollateralWithdrawn is a log parse operation binding the contract event 0xc30fcfbcaac9e0deffa719714eaa82396ff506a0d0d0eebe170830177288715d.
//
// Solidity: event CollateralWithdrawn(address indexed receiver, uint256 collateral)
func (_Broker *BrokerFilterer) ParseCollateralWithdrawn(log types.Log) (*BrokerCollateralWithdrawn, error) {
	event := new(BrokerCollateralWithdrawn)
	if err := _Broker.contract.UnpackLog(event, "CollateralWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the Broker contract.
type BrokerNewRoundIterator struct {
	Event *BrokerNewRound // Event containing the contract specifics and raw log

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
func (it *BrokerNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerNewRound)
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
		it.Event = new(BrokerNewRound)
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
func (it *BrokerNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerNewRound represents a NewRound event raised by the Broker contract.
type BrokerNewRound struct {
	Round         *big.Int
	ClaimedAmount *big.Int
	BurnedAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x5aec57d81928b24d30b1a2aec0d23d693412c37d7ec106b5d8259413716bb1f4.
//
// Solidity: event NewRound(uint256 round, uint256 claimedAmount, uint256 burnedAmount)
func (_Broker *BrokerFilterer) FilterNewRound(opts *bind.FilterOpts) (*BrokerNewRoundIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return &BrokerNewRoundIterator{contract: _Broker.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x5aec57d81928b24d30b1a2aec0d23d693412c37d7ec106b5d8259413716bb1f4.
//
// Solidity: event NewRound(uint256 round, uint256 claimedAmount, uint256 burnedAmount)
func (_Broker *BrokerFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *BrokerNewRound) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "NewRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerNewRound)
				if err := _Broker.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0x5aec57d81928b24d30b1a2aec0d23d693412c37d7ec106b5d8259413716bb1f4.
//
// Solidity: event NewRound(uint256 round, uint256 claimedAmount, uint256 burnedAmount)
func (_Broker *BrokerFilterer) ParseNewRound(log types.Log) (*BrokerNewRound, error) {
	event := new(BrokerNewRound)
	if err := _Broker.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Broker contract.
type BrokerOwnershipTransferredIterator struct {
	Event *BrokerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BrokerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerOwnershipTransferred)
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
		it.Event = new(BrokerOwnershipTransferred)
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
func (it *BrokerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerOwnershipTransferred represents a OwnershipTransferred event raised by the Broker contract.
type BrokerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Broker *BrokerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BrokerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BrokerOwnershipTransferredIterator{contract: _Broker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Broker *BrokerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BrokerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerOwnershipTransferred)
				if err := _Broker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Broker *BrokerFilterer) ParseOwnershipTransferred(log types.Log) (*BrokerOwnershipTransferred, error) {
	event := new(BrokerOwnershipTransferred)
	if err := _Broker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Broker contract.
type BrokerPauseIterator struct {
	Event *BrokerPause // Event containing the contract specifics and raw log

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
func (it *BrokerPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerPause)
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
		it.Event = new(BrokerPause)
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
func (it *BrokerPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerPause represents a Pause event raised by the Broker contract.
type BrokerPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Broker *BrokerFilterer) FilterPause(opts *bind.FilterOpts) (*BrokerPauseIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &BrokerPauseIterator{contract: _Broker.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Broker *BrokerFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *BrokerPause) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerPause)
				if err := _Broker.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Broker *BrokerFilterer) ParsePause(log types.Log) (*BrokerPause, error) {
	event := new(BrokerPause)
	if err := _Broker.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Broker contract.
type BrokerUnpauseIterator struct {
	Event *BrokerUnpause // Event containing the contract specifics and raw log

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
func (it *BrokerUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerUnpause)
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
		it.Event = new(BrokerUnpause)
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
func (it *BrokerUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerUnpause represents a Unpause event raised by the Broker contract.
type BrokerUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Broker *BrokerFilterer) FilterUnpause(opts *bind.FilterOpts) (*BrokerUnpauseIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &BrokerUnpauseIterator{contract: _Broker.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Broker *BrokerFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *BrokerUnpause) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerUnpause)
				if err := _Broker.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Broker *BrokerFilterer) ParseUnpause(log types.Log) (*BrokerUnpause, error) {
	event := new(BrokerUnpause)
	if err := _Broker.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerVitaBiddenIterator is returned from FilterVitaBidden and is used to iterate over the raw logs and unpacked data for VitaBidden events raised by the Broker contract.
type BrokerVitaBiddenIterator struct {
	Event *BrokerVitaBidden // Event containing the contract specifics and raw log

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
func (it *BrokerVitaBiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerVitaBidden)
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
		it.Event = new(BrokerVitaBidden)
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
func (it *BrokerVitaBiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerVitaBiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerVitaBidden represents a VitaBidden event raised by the Broker contract.
type BrokerVitaBidden struct {
	Round      *big.Int
	Sender     common.Address
	Collateral *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVitaBidden is a free log retrieval operation binding the contract event 0x79a65306bcf4df035a08de81edeed738ecd4ed22db73a8ec95d953b82e1611dd.
//
// Solidity: event VitaBidden(uint256 indexed round, address indexed sender, uint256 collateral)
func (_Broker *BrokerFilterer) FilterVitaBidden(opts *bind.FilterOpts, round []*big.Int, sender []common.Address) (*BrokerVitaBiddenIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "VitaBidden", roundRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BrokerVitaBiddenIterator{contract: _Broker.contract, event: "VitaBidden", logs: logs, sub: sub}, nil
}

// WatchVitaBidden is a free log subscription operation binding the contract event 0x79a65306bcf4df035a08de81edeed738ecd4ed22db73a8ec95d953b82e1611dd.
//
// Solidity: event VitaBidden(uint256 indexed round, address indexed sender, uint256 collateral)
func (_Broker *BrokerFilterer) WatchVitaBidden(opts *bind.WatchOpts, sink chan<- *BrokerVitaBidden, round []*big.Int, sender []common.Address) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "VitaBidden", roundRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerVitaBidden)
				if err := _Broker.contract.UnpackLog(event, "VitaBidden", log); err != nil {
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

// ParseVitaBidden is a log parse operation binding the contract event 0x79a65306bcf4df035a08de81edeed738ecd4ed22db73a8ec95d953b82e1611dd.
//
// Solidity: event VitaBidden(uint256 indexed round, address indexed sender, uint256 collateral)
func (_Broker *BrokerFilterer) ParseVitaBidden(log types.Log) (*BrokerVitaBidden, error) {
	event := new(BrokerVitaBidden)
	if err := _Broker.contract.UnpackLog(event, "VitaBidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerVitaBidsSettledIterator is returned from FilterVitaBidsSettled and is used to iterate over the raw logs and unpacked data for VitaBidsSettled events raised by the Broker contract.
type BrokerVitaBidsSettledIterator struct {
	Event *BrokerVitaBidsSettled // Event containing the contract specifics and raw log

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
func (it *BrokerVitaBidsSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerVitaBidsSettled)
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
		it.Event = new(BrokerVitaBidsSettled)
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
func (it *BrokerVitaBidsSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerVitaBidsSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerVitaBidsSettled represents a VitaBidsSettled event raised by the Broker contract.
type BrokerVitaBidsSettled struct {
	Finished bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVitaBidsSettled is a free log retrieval operation binding the contract event 0x80fd533e8e20c803bc5ea74913bea6a24c31cf5d3c8773c55b35692291a4f95d.
//
// Solidity: event VitaBidsSettled(bool finished)
func (_Broker *BrokerFilterer) FilterVitaBidsSettled(opts *bind.FilterOpts) (*BrokerVitaBidsSettledIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "VitaBidsSettled")
	if err != nil {
		return nil, err
	}
	return &BrokerVitaBidsSettledIterator{contract: _Broker.contract, event: "VitaBidsSettled", logs: logs, sub: sub}, nil
}

// WatchVitaBidsSettled is a free log subscription operation binding the contract event 0x80fd533e8e20c803bc5ea74913bea6a24c31cf5d3c8773c55b35692291a4f95d.
//
// Solidity: event VitaBidsSettled(bool finished)
func (_Broker *BrokerFilterer) WatchVitaBidsSettled(opts *bind.WatchOpts, sink chan<- *BrokerVitaBidsSettled) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "VitaBidsSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerVitaBidsSettled)
				if err := _Broker.contract.UnpackLog(event, "VitaBidsSettled", log); err != nil {
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

// ParseVitaBidsSettled is a log parse operation binding the contract event 0x80fd533e8e20c803bc5ea74913bea6a24c31cf5d3c8773c55b35692291a4f95d.
//
// Solidity: event VitaBidsSettled(bool finished)
func (_Broker *BrokerFilterer) ParseVitaBidsSettled(log types.Log) (*BrokerVitaBidsSettled, error) {
	event := new(BrokerVitaBidsSettled)
	if err := _Broker.contract.UnpackLog(event, "VitaBidsSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerVitaBoughtIterator is returned from FilterVitaBought and is used to iterate over the raw logs and unpacked data for VitaBought events raised by the Broker contract.
type BrokerVitaBoughtIterator struct {
	Event *BrokerVitaBought // Event containing the contract specifics and raw log

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
func (it *BrokerVitaBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerVitaBought)
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
		it.Event = new(BrokerVitaBought)
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
func (it *BrokerVitaBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerVitaBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerVitaBought represents a VitaBought event raised by the Broker contract.
type BrokerVitaBought struct {
	Round      *big.Int
	Sender     common.Address
	Collateral *big.Int
	Vita       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVitaBought is a free log retrieval operation binding the contract event 0x2788583a28c987267b8994b86aa4aca25db0bb25d1a742da84afd7845af92af6.
//
// Solidity: event VitaBought(uint256 indexed round, address indexed sender, uint256 collateral, uint256 vita)
func (_Broker *BrokerFilterer) FilterVitaBought(opts *bind.FilterOpts, round []*big.Int, sender []common.Address) (*BrokerVitaBoughtIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "VitaBought", roundRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BrokerVitaBoughtIterator{contract: _Broker.contract, event: "VitaBought", logs: logs, sub: sub}, nil
}

// WatchVitaBought is a free log subscription operation binding the contract event 0x2788583a28c987267b8994b86aa4aca25db0bb25d1a742da84afd7845af92af6.
//
// Solidity: event VitaBought(uint256 indexed round, address indexed sender, uint256 collateral, uint256 vita)
func (_Broker *BrokerFilterer) WatchVitaBought(opts *bind.WatchOpts, sink chan<- *BrokerVitaBought, round []*big.Int, sender []common.Address) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "VitaBought", roundRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerVitaBought)
				if err := _Broker.contract.UnpackLog(event, "VitaBought", log); err != nil {
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

// ParseVitaBought is a log parse operation binding the contract event 0x2788583a28c987267b8994b86aa4aca25db0bb25d1a742da84afd7845af92af6.
//
// Solidity: event VitaBought(uint256 indexed round, address indexed sender, uint256 collateral, uint256 vita)
func (_Broker *BrokerFilterer) ParseVitaBought(log types.Log) (*BrokerVitaBought, error) {
	event := new(BrokerVitaBought)
	if err := _Broker.contract.UnpackLog(event, "VitaBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerWhitelistedAddressAddedIterator is returned from FilterWhitelistedAddressAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAddressAdded events raised by the Broker contract.
type BrokerWhitelistedAddressAddedIterator struct {
	Event *BrokerWhitelistedAddressAdded // Event containing the contract specifics and raw log

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
func (it *BrokerWhitelistedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerWhitelistedAddressAdded)
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
		it.Event = new(BrokerWhitelistedAddressAdded)
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
func (it *BrokerWhitelistedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerWhitelistedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerWhitelistedAddressAdded represents a WhitelistedAddressAdded event raised by the Broker contract.
type BrokerWhitelistedAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressAdded is a free log retrieval operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: event WhitelistedAddressAdded(address addr)
func (_Broker *BrokerFilterer) FilterWhitelistedAddressAdded(opts *bind.FilterOpts) (*BrokerWhitelistedAddressAddedIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return &BrokerWhitelistedAddressAddedIterator{contract: _Broker.contract, event: "WhitelistedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressAdded is a free log subscription operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: event WhitelistedAddressAdded(address addr)
func (_Broker *BrokerFilterer) WatchWhitelistedAddressAdded(opts *bind.WatchOpts, sink chan<- *BrokerWhitelistedAddressAdded) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerWhitelistedAddressAdded)
				if err := _Broker.contract.UnpackLog(event, "WhitelistedAddressAdded", log); err != nil {
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

// ParseWhitelistedAddressAdded is a log parse operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: event WhitelistedAddressAdded(address addr)
func (_Broker *BrokerFilterer) ParseWhitelistedAddressAdded(log types.Log) (*BrokerWhitelistedAddressAdded, error) {
	event := new(BrokerWhitelistedAddressAdded)
	if err := _Broker.contract.UnpackLog(event, "WhitelistedAddressAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerWhitelistedAddressRemovedIterator is returned from FilterWhitelistedAddressRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedAddressRemoved events raised by the Broker contract.
type BrokerWhitelistedAddressRemovedIterator struct {
	Event *BrokerWhitelistedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *BrokerWhitelistedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerWhitelistedAddressRemoved)
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
		it.Event = new(BrokerWhitelistedAddressRemoved)
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
func (it *BrokerWhitelistedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerWhitelistedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerWhitelistedAddressRemoved represents a WhitelistedAddressRemoved event raised by the Broker contract.
type BrokerWhitelistedAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressRemoved is a free log retrieval operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(address addr)
func (_Broker *BrokerFilterer) FilterWhitelistedAddressRemoved(opts *bind.FilterOpts) (*BrokerWhitelistedAddressRemovedIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return &BrokerWhitelistedAddressRemovedIterator{contract: _Broker.contract, event: "WhitelistedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressRemoved is a free log subscription operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(address addr)
func (_Broker *BrokerFilterer) WatchWhitelistedAddressRemoved(opts *bind.WatchOpts, sink chan<- *BrokerWhitelistedAddressRemoved) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerWhitelistedAddressRemoved)
				if err := _Broker.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
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

// ParseWhitelistedAddressRemoved is a log parse operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(address addr)
func (_Broker *BrokerFilterer) ParseWhitelistedAddressRemoved(log types.Log) (*BrokerWhitelistedAddressRemoved, error) {
	event := new(BrokerWhitelistedAddressRemoved)
	if err := _Broker.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
