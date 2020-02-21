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

// PyggStakingABI is the input ABI used to generate the binding from.
const PyggStakingABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_pyggIndex\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_canName\",\"type\":\"bytes12\"},{\"name\":\"_stakeDuration\",\"type\":\"uint256\"},{\"name\":\"_nonDecay\",\"type\":\"bool\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"createPygg\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxPyggsPerAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeAddressesFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeholders\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerDay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prevIndex\",\"type\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getActivePyggCreateTimes\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"},{\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"name\":\"createTimes\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pyggIndex\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"storeToPygg\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxStakeDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pyggIndex\",\"type\":\"uint256\"},{\"name\":\"_stakeDuration\",\"type\":\"uint256\"},{\"name\":\"_nonDecay\",\"type\":\"bool\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"restake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pyggIndex\",\"type\":\"uint256\"},{\"name\":\"_newOwner\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferOwnershipOfPygg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prevIndex\",\"type\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getActivePyggIdx\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"},{\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unStakeDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pyggIndex\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"unstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pyggs\",\"outputs\":[{\"name\":\"canName\",\"type\":\"bytes12\"},{\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"name\":\"stakeDuration\",\"type\":\"uint256\"},{\"name\":\"stakeStartTime\",\"type\":\"uint256\"},{\"name\":\"nonDecay\",\"type\":\"bool\"},{\"name\":\"unstakeStartTime\",\"type\":\"uint256\"},{\"name\":\"pyggOwner\",\"type\":\"address\"},{\"name\":\"createTime\",\"type\":\"uint256\"},{\"name\":\"prev\",\"type\":\"uint256\"},{\"name\":\"next\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getPyggIndexesByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pyggIndex\",\"type\":\"uint256\"},{\"name\":\"_canName\",\"type\":\"bytes12\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"revote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_prevIndex\",\"type\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getActivePyggs\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"},{\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"name\":\"stakeStartTimes\",\"type\":\"uint256[]\"},{\"name\":\"stakeDurations\",\"type\":\"uint256[]\"},{\"name\":\"decays\",\"type\":\"bool[]\"},{\"name\":\"stakedAmounts\",\"type\":\"uint256[]\"},{\"name\":\"canNames\",\"type\":\"bytes12[]\"},{\"name\":\"owners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addAddressesToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_minStakeAmount\",\"type\":\"uint256\"},{\"name\":\"_maxPyggsPerAddr\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pyggIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"canName\",\"type\":\"bytes12\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakeDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakeStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonDecay\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"pyggOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"PyggCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pyggIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"canName\",\"type\":\"bytes12\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakeDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakeStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonDecay\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"pyggOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"PyggUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pyggIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"canName\",\"type\":\"bytes12\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"PyggUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pyggIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"canName\",\"type\":\"bytes12\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"PyggWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// PyggStaking is an auto generated Go binding around an Ethereum contract.
type PyggStaking struct {
	PyggStakingCaller     // Read-only binding to the contract
	PyggStakingTransactor // Write-only binding to the contract
	PyggStakingFilterer   // Log filterer for contract events
}

// PyggStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PyggStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PyggStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PyggStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PyggStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PyggStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PyggStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PyggStakingSession struct {
	Contract     *PyggStaking      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PyggStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PyggStakingCallerSession struct {
	Contract *PyggStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PyggStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PyggStakingTransactorSession struct {
	Contract     *PyggStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PyggStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PyggStakingRaw struct {
	Contract *PyggStaking // Generic contract binding to access the raw methods on
}

// PyggStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PyggStakingCallerRaw struct {
	Contract *PyggStakingCaller // Generic read-only contract binding to access the raw methods on
}

// PyggStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PyggStakingTransactorRaw struct {
	Contract *PyggStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPyggStaking creates a new instance of PyggStaking, bound to a specific deployed contract.
func NewPyggStaking(address common.Address, backend bind.ContractBackend) (*PyggStaking, error) {
	contract, err := bindPyggStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PyggStaking{PyggStakingCaller: PyggStakingCaller{contract: contract}, PyggStakingTransactor: PyggStakingTransactor{contract: contract}, PyggStakingFilterer: PyggStakingFilterer{contract: contract}}, nil
}

// NewPyggStakingCaller creates a new read-only instance of PyggStaking, bound to a specific deployed contract.
func NewPyggStakingCaller(address common.Address, caller bind.ContractCaller) (*PyggStakingCaller, error) {
	contract, err := bindPyggStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PyggStakingCaller{contract: contract}, nil
}

// NewPyggStakingTransactor creates a new write-only instance of PyggStaking, bound to a specific deployed contract.
func NewPyggStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*PyggStakingTransactor, error) {
	contract, err := bindPyggStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PyggStakingTransactor{contract: contract}, nil
}

// NewPyggStakingFilterer creates a new log filterer instance of PyggStaking, bound to a specific deployed contract.
func NewPyggStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*PyggStakingFilterer, error) {
	contract, err := bindPyggStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PyggStakingFilterer{contract: contract}, nil
}

// bindPyggStaking binds a generic wrapper to an already deployed contract.
func bindPyggStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PyggStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PyggStaking *PyggStakingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PyggStaking.Contract.PyggStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PyggStaking *PyggStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PyggStaking.Contract.PyggStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PyggStaking *PyggStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PyggStaking.Contract.PyggStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PyggStaking *PyggStakingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PyggStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PyggStaking *PyggStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PyggStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PyggStaking *PyggStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PyggStaking.Contract.contract.Transact(opts, method, params...)
}

// GetActivePyggCreateTimes is a free data retrieval call binding the contract method 0x6c0a5ebd.
//
// Solidity: function getActivePyggCreateTimes(uint256 _prevIndex, uint256 _limit) constant returns(uint256 count, uint256[] indexes, uint256[] createTimes)
func (_PyggStaking *PyggStakingCaller) GetActivePyggCreateTimes(opts *bind.CallOpts, _prevIndex *big.Int, _limit *big.Int) (struct {
	Count       *big.Int
	Indexes     []*big.Int
	CreateTimes []*big.Int
}, error) {
	ret := new(struct {
		Count       *big.Int
		Indexes     []*big.Int
		CreateTimes []*big.Int
	})
	out := ret
	err := _PyggStaking.contract.Call(opts, out, "getActivePyggCreateTimes", _prevIndex, _limit)
	return *ret, err
}

// GetActivePyggCreateTimes is a free data retrieval call binding the contract method 0x6c0a5ebd.
//
// Solidity: function getActivePyggCreateTimes(uint256 _prevIndex, uint256 _limit) constant returns(uint256 count, uint256[] indexes, uint256[] createTimes)
func (_PyggStaking *PyggStakingSession) GetActivePyggCreateTimes(_prevIndex *big.Int, _limit *big.Int) (struct {
	Count       *big.Int
	Indexes     []*big.Int
	CreateTimes []*big.Int
}, error) {
	return _PyggStaking.Contract.GetActivePyggCreateTimes(&_PyggStaking.CallOpts, _prevIndex, _limit)
}

// GetActivePyggCreateTimes is a free data retrieval call binding the contract method 0x6c0a5ebd.
//
// Solidity: function getActivePyggCreateTimes(uint256 _prevIndex, uint256 _limit) constant returns(uint256 count, uint256[] indexes, uint256[] createTimes)
func (_PyggStaking *PyggStakingCallerSession) GetActivePyggCreateTimes(_prevIndex *big.Int, _limit *big.Int) (struct {
	Count       *big.Int
	Indexes     []*big.Int
	CreateTimes []*big.Int
}, error) {
	return _PyggStaking.Contract.GetActivePyggCreateTimes(&_PyggStaking.CallOpts, _prevIndex, _limit)
}

// GetActivePyggIdx is a free data retrieval call binding the contract method 0x94a9c0f9.
//
// Solidity: function getActivePyggIdx(uint256 _prevIndex, uint256 _limit) constant returns(uint256 count, uint256[] indexes)
func (_PyggStaking *PyggStakingCaller) GetActivePyggIdx(opts *bind.CallOpts, _prevIndex *big.Int, _limit *big.Int) (struct {
	Count   *big.Int
	Indexes []*big.Int
}, error) {
	ret := new(struct {
		Count   *big.Int
		Indexes []*big.Int
	})
	out := ret
	err := _PyggStaking.contract.Call(opts, out, "getActivePyggIdx", _prevIndex, _limit)
	return *ret, err
}

// GetActivePyggIdx is a free data retrieval call binding the contract method 0x94a9c0f9.
//
// Solidity: function getActivePyggIdx(uint256 _prevIndex, uint256 _limit) constant returns(uint256 count, uint256[] indexes)
func (_PyggStaking *PyggStakingSession) GetActivePyggIdx(_prevIndex *big.Int, _limit *big.Int) (struct {
	Count   *big.Int
	Indexes []*big.Int
}, error) {
	return _PyggStaking.Contract.GetActivePyggIdx(&_PyggStaking.CallOpts, _prevIndex, _limit)
}

// GetActivePyggIdx is a free data retrieval call binding the contract method 0x94a9c0f9.
//
// Solidity: function getActivePyggIdx(uint256 _prevIndex, uint256 _limit) constant returns(uint256 count, uint256[] indexes)
func (_PyggStaking *PyggStakingCallerSession) GetActivePyggIdx(_prevIndex *big.Int, _limit *big.Int) (struct {
	Count   *big.Int
	Indexes []*big.Int
}, error) {
	return _PyggStaking.Contract.GetActivePyggIdx(&_PyggStaking.CallOpts, _prevIndex, _limit)
}

// GetActivePyggs is a free data retrieval call binding the contract method 0xdf43a94e.
//
// Solidity: function getActivePyggs(uint256 _prevIndex, uint256 _limit) constant returns(uint256 count, uint256[] indexes, uint256[] stakeStartTimes, uint256[] stakeDurations, bool[] decays, uint256[] stakedAmounts, bytes12[] canNames, address[] owners)
func (_PyggStaking *PyggStakingCaller) GetActivePyggs(opts *bind.CallOpts, _prevIndex *big.Int, _limit *big.Int) (struct {
	Count           *big.Int
	Indexes         []*big.Int
	StakeStartTimes []*big.Int
	StakeDurations  []*big.Int
	Decays          []bool
	StakedAmounts   []*big.Int
	CanNames        [][12]byte
	Owners          []common.Address
}, error) {
	ret := new(struct {
		Count           *big.Int
		Indexes         []*big.Int
		StakeStartTimes []*big.Int
		StakeDurations  []*big.Int
		Decays          []bool
		StakedAmounts   []*big.Int
		CanNames        [][12]byte
		Owners          []common.Address
	})
	out := ret
	err := _PyggStaking.contract.Call(opts, out, "getActivePyggs", _prevIndex, _limit)
	return *ret, err
}

// GetActivePyggs is a free data retrieval call binding the contract method 0xdf43a94e.
//
// Solidity: function getActivePyggs(uint256 _prevIndex, uint256 _limit) constant returns(uint256 count, uint256[] indexes, uint256[] stakeStartTimes, uint256[] stakeDurations, bool[] decays, uint256[] stakedAmounts, bytes12[] canNames, address[] owners)
func (_PyggStaking *PyggStakingSession) GetActivePyggs(_prevIndex *big.Int, _limit *big.Int) (struct {
	Count           *big.Int
	Indexes         []*big.Int
	StakeStartTimes []*big.Int
	StakeDurations  []*big.Int
	Decays          []bool
	StakedAmounts   []*big.Int
	CanNames        [][12]byte
	Owners          []common.Address
}, error) {
	return _PyggStaking.Contract.GetActivePyggs(&_PyggStaking.CallOpts, _prevIndex, _limit)
}

// GetActivePyggs is a free data retrieval call binding the contract method 0xdf43a94e.
//
// Solidity: function getActivePyggs(uint256 _prevIndex, uint256 _limit) constant returns(uint256 count, uint256[] indexes, uint256[] stakeStartTimes, uint256[] stakeDurations, bool[] decays, uint256[] stakedAmounts, bytes12[] canNames, address[] owners)
func (_PyggStaking *PyggStakingCallerSession) GetActivePyggs(_prevIndex *big.Int, _limit *big.Int) (struct {
	Count           *big.Int
	Indexes         []*big.Int
	StakeStartTimes []*big.Int
	StakeDurations  []*big.Int
	Decays          []bool
	StakedAmounts   []*big.Int
	CanNames        [][12]byte
	Owners          []common.Address
}, error) {
	return _PyggStaking.Contract.GetActivePyggs(&_PyggStaking.CallOpts, _prevIndex, _limit)
}

// GetPyggIndexesByAddress is a free data retrieval call binding the contract method 0xd09daa99.
//
// Solidity: function getPyggIndexesByAddress(address _owner) constant returns(uint256[])
func (_PyggStaking *PyggStakingCaller) GetPyggIndexesByAddress(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "getPyggIndexesByAddress", _owner)
	return *ret0, err
}

// GetPyggIndexesByAddress is a free data retrieval call binding the contract method 0xd09daa99.
//
// Solidity: function getPyggIndexesByAddress(address _owner) constant returns(uint256[])
func (_PyggStaking *PyggStakingSession) GetPyggIndexesByAddress(_owner common.Address) ([]*big.Int, error) {
	return _PyggStaking.Contract.GetPyggIndexesByAddress(&_PyggStaking.CallOpts, _owner)
}

// GetPyggIndexesByAddress is a free data retrieval call binding the contract method 0xd09daa99.
//
// Solidity: function getPyggIndexesByAddress(address _owner) constant returns(uint256[])
func (_PyggStaking *PyggStakingCallerSession) GetPyggIndexesByAddress(_owner common.Address) ([]*big.Int, error) {
	return _PyggStaking.Contract.GetPyggIndexesByAddress(&_PyggStaking.CallOpts, _owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _address) constant returns(bool)
func (_PyggStaking *PyggStakingCaller) IsOwner(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "isOwner", _address)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _address) constant returns(bool)
func (_PyggStaking *PyggStakingSession) IsOwner(_address common.Address) (bool, error) {
	return _PyggStaking.Contract.IsOwner(&_PyggStaking.CallOpts, _address)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _address) constant returns(bool)
func (_PyggStaking *PyggStakingCallerSession) IsOwner(_address common.Address) (bool, error) {
	return _PyggStaking.Contract.IsOwner(&_PyggStaking.CallOpts, _address)
}

// MaxPyggsPerAddr is a free data retrieval call binding the contract method 0x1b0690ed.
//
// Solidity: function maxPyggsPerAddr() constant returns(uint256)
func (_PyggStaking *PyggStakingCaller) MaxPyggsPerAddr(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "maxPyggsPerAddr")
	return *ret0, err
}

// MaxPyggsPerAddr is a free data retrieval call binding the contract method 0x1b0690ed.
//
// Solidity: function maxPyggsPerAddr() constant returns(uint256)
func (_PyggStaking *PyggStakingSession) MaxPyggsPerAddr() (*big.Int, error) {
	return _PyggStaking.Contract.MaxPyggsPerAddr(&_PyggStaking.CallOpts)
}

// MaxPyggsPerAddr is a free data retrieval call binding the contract method 0x1b0690ed.
//
// Solidity: function maxPyggsPerAddr() constant returns(uint256)
func (_PyggStaking *PyggStakingCallerSession) MaxPyggsPerAddr() (*big.Int, error) {
	return _PyggStaking.Contract.MaxPyggsPerAddr(&_PyggStaking.CallOpts)
}

// MaxStakeDuration is a free data retrieval call binding the contract method 0x76f70003.
//
// Solidity: function maxStakeDuration() constant returns(uint256)
func (_PyggStaking *PyggStakingCaller) MaxStakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "maxStakeDuration")
	return *ret0, err
}

// MaxStakeDuration is a free data retrieval call binding the contract method 0x76f70003.
//
// Solidity: function maxStakeDuration() constant returns(uint256)
func (_PyggStaking *PyggStakingSession) MaxStakeDuration() (*big.Int, error) {
	return _PyggStaking.Contract.MaxStakeDuration(&_PyggStaking.CallOpts)
}

// MaxStakeDuration is a free data retrieval call binding the contract method 0x76f70003.
//
// Solidity: function maxStakeDuration() constant returns(uint256)
func (_PyggStaking *PyggStakingCallerSession) MaxStakeDuration() (*big.Int, error) {
	return _PyggStaking.Contract.MaxStakeDuration(&_PyggStaking.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() constant returns(uint256)
func (_PyggStaking *PyggStakingCaller) MinStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "minStakeAmount")
	return *ret0, err
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() constant returns(uint256)
func (_PyggStaking *PyggStakingSession) MinStakeAmount() (*big.Int, error) {
	return _PyggStaking.Contract.MinStakeAmount(&_PyggStaking.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() constant returns(uint256)
func (_PyggStaking *PyggStakingCallerSession) MinStakeAmount() (*big.Int, error) {
	return _PyggStaking.Contract.MinStakeAmount(&_PyggStaking.CallOpts)
}

// MinStakeDuration is a free data retrieval call binding the contract method 0x5fec5c64.
//
// Solidity: function minStakeDuration() constant returns(uint256)
func (_PyggStaking *PyggStakingCaller) MinStakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "minStakeDuration")
	return *ret0, err
}

// MinStakeDuration is a free data retrieval call binding the contract method 0x5fec5c64.
//
// Solidity: function minStakeDuration() constant returns(uint256)
func (_PyggStaking *PyggStakingSession) MinStakeDuration() (*big.Int, error) {
	return _PyggStaking.Contract.MinStakeDuration(&_PyggStaking.CallOpts)
}

// MinStakeDuration is a free data retrieval call binding the contract method 0x5fec5c64.
//
// Solidity: function minStakeDuration() constant returns(uint256)
func (_PyggStaking *PyggStakingCallerSession) MinStakeDuration() (*big.Int, error) {
	return _PyggStaking.Contract.MinStakeDuration(&_PyggStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PyggStaking *PyggStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PyggStaking *PyggStakingSession) Owner() (common.Address, error) {
	return _PyggStaking.Contract.Owner(&_PyggStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PyggStaking *PyggStakingCallerSession) Owner() (common.Address, error) {
	return _PyggStaking.Contract.Owner(&_PyggStaking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PyggStaking *PyggStakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PyggStaking *PyggStakingSession) Paused() (bool, error) {
	return _PyggStaking.Contract.Paused(&_PyggStaking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PyggStaking *PyggStakingCallerSession) Paused() (bool, error) {
	return _PyggStaking.Contract.Paused(&_PyggStaking.CallOpts)
}

// Pyggs is a free data retrieval call binding the contract method 0xccfafd5c.
//
// Solidity: function pyggs(uint256 ) constant returns(bytes12 canName, uint256 stakedAmount, uint256 stakeDuration, uint256 stakeStartTime, bool nonDecay, uint256 unstakeStartTime, address pyggOwner, uint256 createTime, uint256 prev, uint256 next)
func (_PyggStaking *PyggStakingCaller) Pyggs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CanName          [12]byte
	StakedAmount     *big.Int
	StakeDuration    *big.Int
	StakeStartTime   *big.Int
	NonDecay         bool
	UnstakeStartTime *big.Int
	PyggOwner        common.Address
	CreateTime       *big.Int
	Prev             *big.Int
	Next             *big.Int
}, error) {
	ret := new(struct {
		CanName          [12]byte
		StakedAmount     *big.Int
		StakeDuration    *big.Int
		StakeStartTime   *big.Int
		NonDecay         bool
		UnstakeStartTime *big.Int
		PyggOwner        common.Address
		CreateTime       *big.Int
		Prev             *big.Int
		Next             *big.Int
	})
	out := ret
	err := _PyggStaking.contract.Call(opts, out, "pyggs", arg0)
	return *ret, err
}

// Pyggs is a free data retrieval call binding the contract method 0xccfafd5c.
//
// Solidity: function pyggs(uint256 ) constant returns(bytes12 canName, uint256 stakedAmount, uint256 stakeDuration, uint256 stakeStartTime, bool nonDecay, uint256 unstakeStartTime, address pyggOwner, uint256 createTime, uint256 prev, uint256 next)
func (_PyggStaking *PyggStakingSession) Pyggs(arg0 *big.Int) (struct {
	CanName          [12]byte
	StakedAmount     *big.Int
	StakeDuration    *big.Int
	StakeStartTime   *big.Int
	NonDecay         bool
	UnstakeStartTime *big.Int
	PyggOwner        common.Address
	CreateTime       *big.Int
	Prev             *big.Int
	Next             *big.Int
}, error) {
	return _PyggStaking.Contract.Pyggs(&_PyggStaking.CallOpts, arg0)
}

// Pyggs is a free data retrieval call binding the contract method 0xccfafd5c.
//
// Solidity: function pyggs(uint256 ) constant returns(bytes12 canName, uint256 stakedAmount, uint256 stakeDuration, uint256 stakeStartTime, bool nonDecay, uint256 unstakeStartTime, address pyggOwner, uint256 createTime, uint256 prev, uint256 next)
func (_PyggStaking *PyggStakingCallerSession) Pyggs(arg0 *big.Int) (struct {
	CanName          [12]byte
	StakedAmount     *big.Int
	StakeDuration    *big.Int
	StakeStartTime   *big.Int
	NonDecay         bool
	UnstakeStartTime *big.Int
	PyggOwner        common.Address
	CreateTime       *big.Int
	Prev             *big.Int
	Next             *big.Int
}, error) {
	return _PyggStaking.Contract.Pyggs(&_PyggStaking.CallOpts, arg0)
}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() constant returns(uint256)
func (_PyggStaking *PyggStakingCaller) SecondsPerDay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "secondsPerDay")
	return *ret0, err
}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() constant returns(uint256)
func (_PyggStaking *PyggStakingSession) SecondsPerDay() (*big.Int, error) {
	return _PyggStaking.Contract.SecondsPerDay(&_PyggStaking.CallOpts)
}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() constant returns(uint256)
func (_PyggStaking *PyggStakingCallerSession) SecondsPerDay() (*big.Int, error) {
	return _PyggStaking.Contract.SecondsPerDay(&_PyggStaking.CallOpts)
}

// Stakeholders is a free data retrieval call binding the contract method 0x423ce1ae.
//
// Solidity: function stakeholders(address , uint256 ) constant returns(uint256)
func (_PyggStaking *PyggStakingCaller) Stakeholders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "stakeholders", arg0, arg1)
	return *ret0, err
}

// Stakeholders is a free data retrieval call binding the contract method 0x423ce1ae.
//
// Solidity: function stakeholders(address , uint256 ) constant returns(uint256)
func (_PyggStaking *PyggStakingSession) Stakeholders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _PyggStaking.Contract.Stakeholders(&_PyggStaking.CallOpts, arg0, arg1)
}

// Stakeholders is a free data retrieval call binding the contract method 0x423ce1ae.
//
// Solidity: function stakeholders(address , uint256 ) constant returns(uint256)
func (_PyggStaking *PyggStakingCallerSession) Stakeholders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _PyggStaking.Contract.Stakeholders(&_PyggStaking.CallOpts, arg0, arg1)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_PyggStaking *PyggStakingCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "totalStaked")
	return *ret0, err
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_PyggStaking *PyggStakingSession) TotalStaked() (*big.Int, error) {
	return _PyggStaking.Contract.TotalStaked(&_PyggStaking.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_PyggStaking *PyggStakingCallerSession) TotalStaked() (*big.Int, error) {
	return _PyggStaking.Contract.TotalStaked(&_PyggStaking.CallOpts)
}

// UnStakeDuration is a free data retrieval call binding the contract method 0xc698d495.
//
// Solidity: function unStakeDuration() constant returns(uint256)
func (_PyggStaking *PyggStakingCaller) UnStakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "unStakeDuration")
	return *ret0, err
}

// UnStakeDuration is a free data retrieval call binding the contract method 0xc698d495.
//
// Solidity: function unStakeDuration() constant returns(uint256)
func (_PyggStaking *PyggStakingSession) UnStakeDuration() (*big.Int, error) {
	return _PyggStaking.Contract.UnStakeDuration(&_PyggStaking.CallOpts)
}

// UnStakeDuration is a free data retrieval call binding the contract method 0xc698d495.
//
// Solidity: function unStakeDuration() constant returns(uint256)
func (_PyggStaking *PyggStakingCallerSession) UnStakeDuration() (*big.Int, error) {
	return _PyggStaking.Contract.UnStakeDuration(&_PyggStaking.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) constant returns(bool)
func (_PyggStaking *PyggStakingCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PyggStaking.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) constant returns(bool)
func (_PyggStaking *PyggStakingSession) Whitelist(arg0 common.Address) (bool, error) {
	return _PyggStaking.Contract.Whitelist(&_PyggStaking.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) constant returns(bool)
func (_PyggStaking *PyggStakingCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _PyggStaking.Contract.Whitelist(&_PyggStaking.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(address addr) returns(bool success)
func (_PyggStaking *PyggStakingTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "addAddressToWhitelist", addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(address addr) returns(bool success)
func (_PyggStaking *PyggStakingSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _PyggStaking.Contract.AddAddressToWhitelist(&_PyggStaking.TransactOpts, addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(address addr) returns(bool success)
func (_PyggStaking *PyggStakingTransactorSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _PyggStaking.Contract.AddAddressToWhitelist(&_PyggStaking.TransactOpts, addr)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(address[] addrs) returns(bool success)
func (_PyggStaking *PyggStakingTransactor) AddAddressesToWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "addAddressesToWhitelist", addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(address[] addrs) returns(bool success)
func (_PyggStaking *PyggStakingSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _PyggStaking.Contract.AddAddressesToWhitelist(&_PyggStaking.TransactOpts, addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(address[] addrs) returns(bool success)
func (_PyggStaking *PyggStakingTransactorSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _PyggStaking.Contract.AddAddressesToWhitelist(&_PyggStaking.TransactOpts, addrs)
}

// CreatePygg is a paid mutator transaction binding the contract method 0x07c35fc0.
//
// Solidity: function createPygg(bytes12 _canName, uint256 _stakeDuration, bool _nonDecay, bytes _data) returns(uint256)
func (_PyggStaking *PyggStakingTransactor) CreatePygg(opts *bind.TransactOpts, _canName [12]byte, _stakeDuration *big.Int, _nonDecay bool, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "createPygg", _canName, _stakeDuration, _nonDecay, _data)
}

// CreatePygg is a paid mutator transaction binding the contract method 0x07c35fc0.
//
// Solidity: function createPygg(bytes12 _canName, uint256 _stakeDuration, bool _nonDecay, bytes _data) returns(uint256)
func (_PyggStaking *PyggStakingSession) CreatePygg(_canName [12]byte, _stakeDuration *big.Int, _nonDecay bool, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.CreatePygg(&_PyggStaking.TransactOpts, _canName, _stakeDuration, _nonDecay, _data)
}

// CreatePygg is a paid mutator transaction binding the contract method 0x07c35fc0.
//
// Solidity: function createPygg(bytes12 _canName, uint256 _stakeDuration, bool _nonDecay, bytes _data) returns(uint256)
func (_PyggStaking *PyggStakingTransactorSession) CreatePygg(_canName [12]byte, _stakeDuration *big.Int, _nonDecay bool, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.CreatePygg(&_PyggStaking.TransactOpts, _canName, _stakeDuration, _nonDecay, _data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PyggStaking *PyggStakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PyggStaking *PyggStakingSession) Pause() (*types.Transaction, error) {
	return _PyggStaking.Contract.Pause(&_PyggStaking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PyggStaking *PyggStakingTransactorSession) Pause() (*types.Transaction, error) {
	return _PyggStaking.Contract.Pause(&_PyggStaking.TransactOpts)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(address addr) returns(bool success)
func (_PyggStaking *PyggStakingTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "removeAddressFromWhitelist", addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(address addr) returns(bool success)
func (_PyggStaking *PyggStakingSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _PyggStaking.Contract.RemoveAddressFromWhitelist(&_PyggStaking.TransactOpts, addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(address addr) returns(bool success)
func (_PyggStaking *PyggStakingTransactorSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _PyggStaking.Contract.RemoveAddressFromWhitelist(&_PyggStaking.TransactOpts, addr)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(address[] addrs) returns(bool success)
func (_PyggStaking *PyggStakingTransactor) RemoveAddressesFromWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "removeAddressesFromWhitelist", addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(address[] addrs) returns(bool success)
func (_PyggStaking *PyggStakingSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _PyggStaking.Contract.RemoveAddressesFromWhitelist(&_PyggStaking.TransactOpts, addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(address[] addrs) returns(bool success)
func (_PyggStaking *PyggStakingTransactorSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _PyggStaking.Contract.RemoveAddressesFromWhitelist(&_PyggStaking.TransactOpts, addrs)
}

// Restake is a paid mutator transaction binding the contract method 0x7b24a5fd.
//
// Solidity: function restake(uint256 _pyggIndex, uint256 _stakeDuration, bool _nonDecay, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactor) Restake(opts *bind.TransactOpts, _pyggIndex *big.Int, _stakeDuration *big.Int, _nonDecay bool, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "restake", _pyggIndex, _stakeDuration, _nonDecay, _data)
}

// Restake is a paid mutator transaction binding the contract method 0x7b24a5fd.
//
// Solidity: function restake(uint256 _pyggIndex, uint256 _stakeDuration, bool _nonDecay, bytes _data) returns()
func (_PyggStaking *PyggStakingSession) Restake(_pyggIndex *big.Int, _stakeDuration *big.Int, _nonDecay bool, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.Restake(&_PyggStaking.TransactOpts, _pyggIndex, _stakeDuration, _nonDecay, _data)
}

// Restake is a paid mutator transaction binding the contract method 0x7b24a5fd.
//
// Solidity: function restake(uint256 _pyggIndex, uint256 _stakeDuration, bool _nonDecay, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactorSession) Restake(_pyggIndex *big.Int, _stakeDuration *big.Int, _nonDecay bool, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.Restake(&_PyggStaking.TransactOpts, _pyggIndex, _stakeDuration, _nonDecay, _data)
}

// Revote is a paid mutator transaction binding the contract method 0xd3e41fd2.
//
// Solidity: function revote(uint256 _pyggIndex, bytes12 _canName, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactor) Revote(opts *bind.TransactOpts, _pyggIndex *big.Int, _canName [12]byte, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "revote", _pyggIndex, _canName, _data)
}

// Revote is a paid mutator transaction binding the contract method 0xd3e41fd2.
//
// Solidity: function revote(uint256 _pyggIndex, bytes12 _canName, bytes _data) returns()
func (_PyggStaking *PyggStakingSession) Revote(_pyggIndex *big.Int, _canName [12]byte, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.Revote(&_PyggStaking.TransactOpts, _pyggIndex, _canName, _data)
}

// Revote is a paid mutator transaction binding the contract method 0xd3e41fd2.
//
// Solidity: function revote(uint256 _pyggIndex, bytes12 _canName, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactorSession) Revote(_pyggIndex *big.Int, _canName [12]byte, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.Revote(&_PyggStaking.TransactOpts, _pyggIndex, _canName, _data)
}

// StoreToPygg is a paid mutator transaction binding the contract method 0x6e7b3017.
//
// Solidity: function storeToPygg(uint256 _pyggIndex, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactor) StoreToPygg(opts *bind.TransactOpts, _pyggIndex *big.Int, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "storeToPygg", _pyggIndex, _data)
}

// StoreToPygg is a paid mutator transaction binding the contract method 0x6e7b3017.
//
// Solidity: function storeToPygg(uint256 _pyggIndex, bytes _data) returns()
func (_PyggStaking *PyggStakingSession) StoreToPygg(_pyggIndex *big.Int, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.StoreToPygg(&_PyggStaking.TransactOpts, _pyggIndex, _data)
}

// StoreToPygg is a paid mutator transaction binding the contract method 0x6e7b3017.
//
// Solidity: function storeToPygg(uint256 _pyggIndex, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactorSession) StoreToPygg(_pyggIndex *big.Int, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.StoreToPygg(&_PyggStaking.TransactOpts, _pyggIndex, _data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PyggStaking *PyggStakingTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PyggStaking *PyggStakingSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PyggStaking.Contract.TransferOwnership(&_PyggStaking.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PyggStaking *PyggStakingTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PyggStaking.Contract.TransferOwnership(&_PyggStaking.TransactOpts, _newOwner)
}

// TransferOwnershipOfPygg is a paid mutator transaction binding the contract method 0x7d564937.
//
// Solidity: function transferOwnershipOfPygg(uint256 _pyggIndex, address _newOwner, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactor) TransferOwnershipOfPygg(opts *bind.TransactOpts, _pyggIndex *big.Int, _newOwner common.Address, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "transferOwnershipOfPygg", _pyggIndex, _newOwner, _data)
}

// TransferOwnershipOfPygg is a paid mutator transaction binding the contract method 0x7d564937.
//
// Solidity: function transferOwnershipOfPygg(uint256 _pyggIndex, address _newOwner, bytes _data) returns()
func (_PyggStaking *PyggStakingSession) TransferOwnershipOfPygg(_pyggIndex *big.Int, _newOwner common.Address, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.TransferOwnershipOfPygg(&_PyggStaking.TransactOpts, _pyggIndex, _newOwner, _data)
}

// TransferOwnershipOfPygg is a paid mutator transaction binding the contract method 0x7d564937.
//
// Solidity: function transferOwnershipOfPygg(uint256 _pyggIndex, address _newOwner, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactorSession) TransferOwnershipOfPygg(_pyggIndex *big.Int, _newOwner common.Address, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.TransferOwnershipOfPygg(&_PyggStaking.TransactOpts, _pyggIndex, _newOwner, _data)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PyggStaking *PyggStakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PyggStaking *PyggStakingSession) Unpause() (*types.Transaction, error) {
	return _PyggStaking.Contract.Unpause(&_PyggStaking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PyggStaking *PyggStakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _PyggStaking.Contract.Unpause(&_PyggStaking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0xc8fd6ed0.
//
// Solidity: function unstake(uint256 _pyggIndex, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactor) Unstake(opts *bind.TransactOpts, _pyggIndex *big.Int, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "unstake", _pyggIndex, _data)
}

// Unstake is a paid mutator transaction binding the contract method 0xc8fd6ed0.
//
// Solidity: function unstake(uint256 _pyggIndex, bytes _data) returns()
func (_PyggStaking *PyggStakingSession) Unstake(_pyggIndex *big.Int, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.Unstake(&_PyggStaking.TransactOpts, _pyggIndex, _data)
}

// Unstake is a paid mutator transaction binding the contract method 0xc8fd6ed0.
//
// Solidity: function unstake(uint256 _pyggIndex, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactorSession) Unstake(_pyggIndex *big.Int, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.Unstake(&_PyggStaking.TransactOpts, _pyggIndex, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x030ba25d.
//
// Solidity: function withdraw(uint256 _pyggIndex, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactor) Withdraw(opts *bind.TransactOpts, _pyggIndex *big.Int, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.contract.Transact(opts, "withdraw", _pyggIndex, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x030ba25d.
//
// Solidity: function withdraw(uint256 _pyggIndex, bytes _data) returns()
func (_PyggStaking *PyggStakingSession) Withdraw(_pyggIndex *big.Int, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.Withdraw(&_PyggStaking.TransactOpts, _pyggIndex, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x030ba25d.
//
// Solidity: function withdraw(uint256 _pyggIndex, bytes _data) returns()
func (_PyggStaking *PyggStakingTransactorSession) Withdraw(_pyggIndex *big.Int, _data []byte) (*types.Transaction, error) {
	return _PyggStaking.Contract.Withdraw(&_PyggStaking.TransactOpts, _pyggIndex, _data)
}

// PyggStakingPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PyggStaking contract.
type PyggStakingPauseIterator struct {
	Event *PyggStakingPause // Event containing the contract specifics and raw log

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
func (it *PyggStakingPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PyggStakingPause)
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
		it.Event = new(PyggStakingPause)
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
func (it *PyggStakingPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PyggStakingPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PyggStakingPause represents a Pause event raised by the PyggStaking contract.
type PyggStakingPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PyggStaking *PyggStakingFilterer) FilterPause(opts *bind.FilterOpts) (*PyggStakingPauseIterator, error) {

	logs, sub, err := _PyggStaking.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PyggStakingPauseIterator{contract: _PyggStaking.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PyggStaking *PyggStakingFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PyggStakingPause) (event.Subscription, error) {

	logs, sub, err := _PyggStaking.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PyggStakingPause)
				if err := _PyggStaking.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PyggStakingPyggCreatedIterator is returned from FilterPyggCreated and is used to iterate over the raw logs and unpacked data for PyggCreated events raised by the PyggStaking contract.
type PyggStakingPyggCreatedIterator struct {
	Event *PyggStakingPyggCreated // Event containing the contract specifics and raw log

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
func (it *PyggStakingPyggCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PyggStakingPyggCreated)
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
		it.Event = new(PyggStakingPyggCreated)
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
func (it *PyggStakingPyggCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PyggStakingPyggCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PyggStakingPyggCreated represents a PyggCreated event raised by the PyggStaking contract.
type PyggStakingPyggCreated struct {
	PyggIndex      *big.Int
	CanName        [12]byte
	Amount         *big.Int
	StakeDuration  *big.Int
	StakeStartTime *big.Int
	NonDecay       bool
	PyggOwner      common.Address
	Data           []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPyggCreated is a free log retrieval operation binding the contract event 0xd7812fae7f8126d2df0f5449a2cc0744d2e9d3fc8c161de6193bc4df6c68d365.
//
// Solidity: event PyggCreated(uint256 pyggIndex, bytes12 canName, uint256 amount, uint256 stakeDuration, uint256 stakeStartTime, bool nonDecay, address pyggOwner, bytes data)
func (_PyggStaking *PyggStakingFilterer) FilterPyggCreated(opts *bind.FilterOpts) (*PyggStakingPyggCreatedIterator, error) {

	logs, sub, err := _PyggStaking.contract.FilterLogs(opts, "PyggCreated")
	if err != nil {
		return nil, err
	}
	return &PyggStakingPyggCreatedIterator{contract: _PyggStaking.contract, event: "PyggCreated", logs: logs, sub: sub}, nil
}

// WatchPyggCreated is a free log subscription operation binding the contract event 0xd7812fae7f8126d2df0f5449a2cc0744d2e9d3fc8c161de6193bc4df6c68d365.
//
// Solidity: event PyggCreated(uint256 pyggIndex, bytes12 canName, uint256 amount, uint256 stakeDuration, uint256 stakeStartTime, bool nonDecay, address pyggOwner, bytes data)
func (_PyggStaking *PyggStakingFilterer) WatchPyggCreated(opts *bind.WatchOpts, sink chan<- *PyggStakingPyggCreated) (event.Subscription, error) {

	logs, sub, err := _PyggStaking.contract.WatchLogs(opts, "PyggCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PyggStakingPyggCreated)
				if err := _PyggStaking.contract.UnpackLog(event, "PyggCreated", log); err != nil {
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

// PyggStakingPyggUnstakeIterator is returned from FilterPyggUnstake and is used to iterate over the raw logs and unpacked data for PyggUnstake events raised by the PyggStaking contract.
type PyggStakingPyggUnstakeIterator struct {
	Event *PyggStakingPyggUnstake // Event containing the contract specifics and raw log

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
func (it *PyggStakingPyggUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PyggStakingPyggUnstake)
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
		it.Event = new(PyggStakingPyggUnstake)
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
func (it *PyggStakingPyggUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PyggStakingPyggUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PyggStakingPyggUnstake represents a PyggUnstake event raised by the PyggStaking contract.
type PyggStakingPyggUnstake struct {
	PyggIndex *big.Int
	CanName   [12]byte
	Amount    *big.Int
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPyggUnstake is a free log retrieval operation binding the contract event 0x9954bdedc474e937b39bbb080fc136e2edf1cef61f0906d36203267f4930762e.
//
// Solidity: event PyggUnstake(uint256 pyggIndex, bytes12 canName, uint256 amount, bytes data)
func (_PyggStaking *PyggStakingFilterer) FilterPyggUnstake(opts *bind.FilterOpts) (*PyggStakingPyggUnstakeIterator, error) {

	logs, sub, err := _PyggStaking.contract.FilterLogs(opts, "PyggUnstake")
	if err != nil {
		return nil, err
	}
	return &PyggStakingPyggUnstakeIterator{contract: _PyggStaking.contract, event: "PyggUnstake", logs: logs, sub: sub}, nil
}

// WatchPyggUnstake is a free log subscription operation binding the contract event 0x9954bdedc474e937b39bbb080fc136e2edf1cef61f0906d36203267f4930762e.
//
// Solidity: event PyggUnstake(uint256 pyggIndex, bytes12 canName, uint256 amount, bytes data)
func (_PyggStaking *PyggStakingFilterer) WatchPyggUnstake(opts *bind.WatchOpts, sink chan<- *PyggStakingPyggUnstake) (event.Subscription, error) {

	logs, sub, err := _PyggStaking.contract.WatchLogs(opts, "PyggUnstake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PyggStakingPyggUnstake)
				if err := _PyggStaking.contract.UnpackLog(event, "PyggUnstake", log); err != nil {
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

// PyggStakingPyggUpdatedIterator is returned from FilterPyggUpdated and is used to iterate over the raw logs and unpacked data for PyggUpdated events raised by the PyggStaking contract.
type PyggStakingPyggUpdatedIterator struct {
	Event *PyggStakingPyggUpdated // Event containing the contract specifics and raw log

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
func (it *PyggStakingPyggUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PyggStakingPyggUpdated)
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
		it.Event = new(PyggStakingPyggUpdated)
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
func (it *PyggStakingPyggUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PyggStakingPyggUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PyggStakingPyggUpdated represents a PyggUpdated event raised by the PyggStaking contract.
type PyggStakingPyggUpdated struct {
	PyggIndex      *big.Int
	CanName        [12]byte
	Amount         *big.Int
	StakeDuration  *big.Int
	StakeStartTime *big.Int
	NonDecay       bool
	PyggOwner      common.Address
	Data           []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPyggUpdated is a free log retrieval operation binding the contract event 0x0b074423c8a0f26c131cd7c88b19ef6adf084b812c97bdd1fb9dcf339ee9a387.
//
// Solidity: event PyggUpdated(uint256 pyggIndex, bytes12 canName, uint256 amount, uint256 stakeDuration, uint256 stakeStartTime, bool nonDecay, address pyggOwner, bytes data)
func (_PyggStaking *PyggStakingFilterer) FilterPyggUpdated(opts *bind.FilterOpts) (*PyggStakingPyggUpdatedIterator, error) {

	logs, sub, err := _PyggStaking.contract.FilterLogs(opts, "PyggUpdated")
	if err != nil {
		return nil, err
	}
	return &PyggStakingPyggUpdatedIterator{contract: _PyggStaking.contract, event: "PyggUpdated", logs: logs, sub: sub}, nil
}

// WatchPyggUpdated is a free log subscription operation binding the contract event 0x0b074423c8a0f26c131cd7c88b19ef6adf084b812c97bdd1fb9dcf339ee9a387.
//
// Solidity: event PyggUpdated(uint256 pyggIndex, bytes12 canName, uint256 amount, uint256 stakeDuration, uint256 stakeStartTime, bool nonDecay, address pyggOwner, bytes data)
func (_PyggStaking *PyggStakingFilterer) WatchPyggUpdated(opts *bind.WatchOpts, sink chan<- *PyggStakingPyggUpdated) (event.Subscription, error) {

	logs, sub, err := _PyggStaking.contract.WatchLogs(opts, "PyggUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PyggStakingPyggUpdated)
				if err := _PyggStaking.contract.UnpackLog(event, "PyggUpdated", log); err != nil {
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

// PyggStakingPyggWithdrawIterator is returned from FilterPyggWithdraw and is used to iterate over the raw logs and unpacked data for PyggWithdraw events raised by the PyggStaking contract.
type PyggStakingPyggWithdrawIterator struct {
	Event *PyggStakingPyggWithdraw // Event containing the contract specifics and raw log

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
func (it *PyggStakingPyggWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PyggStakingPyggWithdraw)
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
		it.Event = new(PyggStakingPyggWithdraw)
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
func (it *PyggStakingPyggWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PyggStakingPyggWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PyggStakingPyggWithdraw represents a PyggWithdraw event raised by the PyggStaking contract.
type PyggStakingPyggWithdraw struct {
	PyggIndex *big.Int
	CanName   [12]byte
	Amount    *big.Int
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPyggWithdraw is a free log retrieval operation binding the contract event 0xf99c0736fafe9102d41ec0b56c187b26a6e35ae50415dcbecedf73112d0ec763.
//
// Solidity: event PyggWithdraw(uint256 pyggIndex, bytes12 canName, uint256 amount, bytes data)
func (_PyggStaking *PyggStakingFilterer) FilterPyggWithdraw(opts *bind.FilterOpts) (*PyggStakingPyggWithdrawIterator, error) {

	logs, sub, err := _PyggStaking.contract.FilterLogs(opts, "PyggWithdraw")
	if err != nil {
		return nil, err
	}
	return &PyggStakingPyggWithdrawIterator{contract: _PyggStaking.contract, event: "PyggWithdraw", logs: logs, sub: sub}, nil
}

// WatchPyggWithdraw is a free log subscription operation binding the contract event 0xf99c0736fafe9102d41ec0b56c187b26a6e35ae50415dcbecedf73112d0ec763.
//
// Solidity: event PyggWithdraw(uint256 pyggIndex, bytes12 canName, uint256 amount, bytes data)
func (_PyggStaking *PyggStakingFilterer) WatchPyggWithdraw(opts *bind.WatchOpts, sink chan<- *PyggStakingPyggWithdraw) (event.Subscription, error) {

	logs, sub, err := _PyggStaking.contract.WatchLogs(opts, "PyggWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PyggStakingPyggWithdraw)
				if err := _PyggStaking.contract.UnpackLog(event, "PyggWithdraw", log); err != nil {
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

// PyggStakingUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PyggStaking contract.
type PyggStakingUnpauseIterator struct {
	Event *PyggStakingUnpause // Event containing the contract specifics and raw log

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
func (it *PyggStakingUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PyggStakingUnpause)
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
		it.Event = new(PyggStakingUnpause)
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
func (it *PyggStakingUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PyggStakingUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PyggStakingUnpause represents a Unpause event raised by the PyggStaking contract.
type PyggStakingUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PyggStaking *PyggStakingFilterer) FilterUnpause(opts *bind.FilterOpts) (*PyggStakingUnpauseIterator, error) {

	logs, sub, err := _PyggStaking.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PyggStakingUnpauseIterator{contract: _PyggStaking.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PyggStaking *PyggStakingFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PyggStakingUnpause) (event.Subscription, error) {

	logs, sub, err := _PyggStaking.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PyggStakingUnpause)
				if err := _PyggStaking.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// PyggStakingWhitelistedAddressAddedIterator is returned from FilterWhitelistedAddressAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAddressAdded events raised by the PyggStaking contract.
type PyggStakingWhitelistedAddressAddedIterator struct {
	Event *PyggStakingWhitelistedAddressAdded // Event containing the contract specifics and raw log

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
func (it *PyggStakingWhitelistedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PyggStakingWhitelistedAddressAdded)
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
		it.Event = new(PyggStakingWhitelistedAddressAdded)
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
func (it *PyggStakingWhitelistedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PyggStakingWhitelistedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PyggStakingWhitelistedAddressAdded represents a WhitelistedAddressAdded event raised by the PyggStaking contract.
type PyggStakingWhitelistedAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressAdded is a free log retrieval operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: event WhitelistedAddressAdded(address addr)
func (_PyggStaking *PyggStakingFilterer) FilterWhitelistedAddressAdded(opts *bind.FilterOpts) (*PyggStakingWhitelistedAddressAddedIterator, error) {

	logs, sub, err := _PyggStaking.contract.FilterLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return &PyggStakingWhitelistedAddressAddedIterator{contract: _PyggStaking.contract, event: "WhitelistedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressAdded is a free log subscription operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: event WhitelistedAddressAdded(address addr)
func (_PyggStaking *PyggStakingFilterer) WatchWhitelistedAddressAdded(opts *bind.WatchOpts, sink chan<- *PyggStakingWhitelistedAddressAdded) (event.Subscription, error) {

	logs, sub, err := _PyggStaking.contract.WatchLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PyggStakingWhitelistedAddressAdded)
				if err := _PyggStaking.contract.UnpackLog(event, "WhitelistedAddressAdded", log); err != nil {
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

// PyggStakingWhitelistedAddressRemovedIterator is returned from FilterWhitelistedAddressRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedAddressRemoved events raised by the PyggStaking contract.
type PyggStakingWhitelistedAddressRemovedIterator struct {
	Event *PyggStakingWhitelistedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *PyggStakingWhitelistedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PyggStakingWhitelistedAddressRemoved)
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
		it.Event = new(PyggStakingWhitelistedAddressRemoved)
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
func (it *PyggStakingWhitelistedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PyggStakingWhitelistedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PyggStakingWhitelistedAddressRemoved represents a WhitelistedAddressRemoved event raised by the PyggStaking contract.
type PyggStakingWhitelistedAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressRemoved is a free log retrieval operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(address addr)
func (_PyggStaking *PyggStakingFilterer) FilterWhitelistedAddressRemoved(opts *bind.FilterOpts) (*PyggStakingWhitelistedAddressRemovedIterator, error) {

	logs, sub, err := _PyggStaking.contract.FilterLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return &PyggStakingWhitelistedAddressRemovedIterator{contract: _PyggStaking.contract, event: "WhitelistedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressRemoved is a free log subscription operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(address addr)
func (_PyggStaking *PyggStakingFilterer) WatchWhitelistedAddressRemoved(opts *bind.WatchOpts, sink chan<- *PyggStakingWhitelistedAddressRemoved) (event.Subscription, error) {

	logs, sub, err := _PyggStaking.contract.WatchLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PyggStakingWhitelistedAddressRemoved)
				if err := _PyggStaking.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
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
