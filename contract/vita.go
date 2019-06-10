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

// VitaABI is the input ABI used to generate the binding from.
const VitaABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDecayHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cycleIncrementalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decayedIncrementalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPoolSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"donationPoolAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastClaimViewIDs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingPoolSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newVPS\",\"type\":\"address\"}],\"name\":\"setVPS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesisPoolAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vps\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDonationPoolClaimViewID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextDecayHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPoolAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"claimable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRewardPool\",\"type\":\"address\"}],\"name\":\"setRewardPoolAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"donationPoolSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastRewardPoolClaimViewID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incrementalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDonationPool\",\"type\":\"address\"}],\"name\":\"setDonationPoolAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_vps\",\"type\":\"address\"},{\"name\":\"_genesisPoolAddress\",\"type\":\"address\"},{\"name\":\"_rewardPoolAddress\",\"type\":\"address\"},{\"name\":\"_donationPoolAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"viewID\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"incremetnalSupply\",\"type\":\"uint256\"}],\"name\":\"Decay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"viewID\",\"type\":\"uint256\"}],\"name\":\"UpdateView\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// Vita is an auto generated Go binding around an Ethereum contract.
type Vita struct {
	VitaCaller     // Read-only binding to the contract
	VitaTransactor // Write-only binding to the contract
	VitaFilterer   // Log filterer for contract events
}

// VitaCaller is an auto generated read-only Go binding around an Ethereum contract.
type VitaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VitaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VitaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VitaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VitaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VitaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VitaSession struct {
	Contract     *Vita             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VitaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VitaCallerSession struct {
	Contract *VitaCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VitaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VitaTransactorSession struct {
	Contract     *VitaTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VitaRaw is an auto generated low-level Go binding around an Ethereum contract.
type VitaRaw struct {
	Contract *Vita // Generic contract binding to access the raw methods on
}

// VitaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VitaCallerRaw struct {
	Contract *VitaCaller // Generic read-only contract binding to access the raw methods on
}

// VitaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VitaTransactorRaw struct {
	Contract *VitaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVita creates a new instance of Vita, bound to a specific deployed contract.
func NewVita(address common.Address, backend bind.ContractBackend) (*Vita, error) {
	contract, err := bindVita(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vita{VitaCaller: VitaCaller{contract: contract}, VitaTransactor: VitaTransactor{contract: contract}, VitaFilterer: VitaFilterer{contract: contract}}, nil
}

// NewVitaCaller creates a new read-only instance of Vita, bound to a specific deployed contract.
func NewVitaCaller(address common.Address, caller bind.ContractCaller) (*VitaCaller, error) {
	contract, err := bindVita(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VitaCaller{contract: contract}, nil
}

// NewVitaTransactor creates a new write-only instance of Vita, bound to a specific deployed contract.
func NewVitaTransactor(address common.Address, transactor bind.ContractTransactor) (*VitaTransactor, error) {
	contract, err := bindVita(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VitaTransactor{contract: contract}, nil
}

// NewVitaFilterer creates a new log filterer instance of Vita, bound to a specific deployed contract.
func NewVitaFilterer(address common.Address, filterer bind.ContractFilterer) (*VitaFilterer, error) {
	contract, err := bindVita(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VitaFilterer{contract: contract}, nil
}

// bindVita binds a generic wrapper to an already deployed contract.
func bindVita(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VitaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vita *VitaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vita.Contract.VitaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vita *VitaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vita.Contract.VitaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vita *VitaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vita.Contract.VitaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vita *VitaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vita.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vita *VitaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vita.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vita *VitaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vita.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_Vita *VitaCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_Vita *VitaSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Vita.Contract.Allowance(&_Vita.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256)
func (_Vita *VitaCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Vita.Contract.Allowance(&_Vita.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Vita *VitaCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Vita *VitaSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Vita.Contract.BalanceOf(&_Vita.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Vita *VitaCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Vita.Contract.BalanceOf(&_Vita.CallOpts, _owner)
}

// Claimable is a free data retrieval call binding the contract method 0xaf38d757.
//
// Solidity: function claimable() constant returns(bool)
func (_Vita *VitaCaller) Claimable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "claimable")
	return *ret0, err
}

// Claimable is a free data retrieval call binding the contract method 0xaf38d757.
//
// Solidity: function claimable() constant returns(bool)
func (_Vita *VitaSession) Claimable() (bool, error) {
	return _Vita.Contract.Claimable(&_Vita.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0xaf38d757.
//
// Solidity: function claimable() constant returns(bool)
func (_Vita *VitaCallerSession) Claimable() (bool, error) {
	return _Vita.Contract.Claimable(&_Vita.CallOpts)
}

// CycleIncrementalSupply is a free data retrieval call binding the contract method 0x13d92623.
//
// Solidity: function cycleIncrementalSupply() constant returns(uint256)
func (_Vita *VitaCaller) CycleIncrementalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "cycleIncrementalSupply")
	return *ret0, err
}

// CycleIncrementalSupply is a free data retrieval call binding the contract method 0x13d92623.
//
// Solidity: function cycleIncrementalSupply() constant returns(uint256)
func (_Vita *VitaSession) CycleIncrementalSupply() (*big.Int, error) {
	return _Vita.Contract.CycleIncrementalSupply(&_Vita.CallOpts)
}

// CycleIncrementalSupply is a free data retrieval call binding the contract method 0x13d92623.
//
// Solidity: function cycleIncrementalSupply() constant returns(uint256)
func (_Vita *VitaCallerSession) CycleIncrementalSupply() (*big.Int, error) {
	return _Vita.Contract.CycleIncrementalSupply(&_Vita.CallOpts)
}

// DecayedIncrementalSupply is a free data retrieval call binding the contract method 0x1efd1a4c.
//
// Solidity: function decayedIncrementalSupply() constant returns(uint256)
func (_Vita *VitaCaller) DecayedIncrementalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "decayedIncrementalSupply")
	return *ret0, err
}

// DecayedIncrementalSupply is a free data retrieval call binding the contract method 0x1efd1a4c.
//
// Solidity: function decayedIncrementalSupply() constant returns(uint256)
func (_Vita *VitaSession) DecayedIncrementalSupply() (*big.Int, error) {
	return _Vita.Contract.DecayedIncrementalSupply(&_Vita.CallOpts)
}

// DecayedIncrementalSupply is a free data retrieval call binding the contract method 0x1efd1a4c.
//
// Solidity: function decayedIncrementalSupply() constant returns(uint256)
func (_Vita *VitaCallerSession) DecayedIncrementalSupply() (*big.Int, error) {
	return _Vita.Contract.DecayedIncrementalSupply(&_Vita.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Vita *VitaCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Vita *VitaSession) Decimals() (uint8, error) {
	return _Vita.Contract.Decimals(&_Vita.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Vita *VitaCallerSession) Decimals() (uint8, error) {
	return _Vita.Contract.Decimals(&_Vita.CallOpts)
}

// DonationPoolAddress is a free data retrieval call binding the contract method 0x280e6434.
//
// Solidity: function donationPoolAddress() constant returns(address)
func (_Vita *VitaCaller) DonationPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "donationPoolAddress")
	return *ret0, err
}

// DonationPoolAddress is a free data retrieval call binding the contract method 0x280e6434.
//
// Solidity: function donationPoolAddress() constant returns(address)
func (_Vita *VitaSession) DonationPoolAddress() (common.Address, error) {
	return _Vita.Contract.DonationPoolAddress(&_Vita.CallOpts)
}

// DonationPoolAddress is a free data retrieval call binding the contract method 0x280e6434.
//
// Solidity: function donationPoolAddress() constant returns(address)
func (_Vita *VitaCallerSession) DonationPoolAddress() (common.Address, error) {
	return _Vita.Contract.DonationPoolAddress(&_Vita.CallOpts)
}

// DonationPoolSize is a free data retrieval call binding the contract method 0xc7f3a874.
//
// Solidity: function donationPoolSize() constant returns(uint256)
func (_Vita *VitaCaller) DonationPoolSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "donationPoolSize")
	return *ret0, err
}

// DonationPoolSize is a free data retrieval call binding the contract method 0xc7f3a874.
//
// Solidity: function donationPoolSize() constant returns(uint256)
func (_Vita *VitaSession) DonationPoolSize() (*big.Int, error) {
	return _Vita.Contract.DonationPoolSize(&_Vita.CallOpts)
}

// DonationPoolSize is a free data retrieval call binding the contract method 0xc7f3a874.
//
// Solidity: function donationPoolSize() constant returns(uint256)
func (_Vita *VitaCallerSession) DonationPoolSize() (*big.Int, error) {
	return _Vita.Contract.DonationPoolSize(&_Vita.CallOpts)
}

// GenesisPoolAddress is a free data retrieval call binding the contract method 0x4063971c.
//
// Solidity: function genesisPoolAddress() constant returns(address)
func (_Vita *VitaCaller) GenesisPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "genesisPoolAddress")
	return *ret0, err
}

// GenesisPoolAddress is a free data retrieval call binding the contract method 0x4063971c.
//
// Solidity: function genesisPoolAddress() constant returns(address)
func (_Vita *VitaSession) GenesisPoolAddress() (common.Address, error) {
	return _Vita.Contract.GenesisPoolAddress(&_Vita.CallOpts)
}

// GenesisPoolAddress is a free data retrieval call binding the contract method 0x4063971c.
//
// Solidity: function genesisPoolAddress() constant returns(address)
func (_Vita *VitaCallerSession) GenesisPoolAddress() (common.Address, error) {
	return _Vita.Contract.GenesisPoolAddress(&_Vita.CallOpts)
}

// IncrementalSupply is a free data retrieval call binding the contract method 0xe26204d7.
//
// Solidity: function incrementalSupply() constant returns(uint256)
func (_Vita *VitaCaller) IncrementalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "incrementalSupply")
	return *ret0, err
}

// IncrementalSupply is a free data retrieval call binding the contract method 0xe26204d7.
//
// Solidity: function incrementalSupply() constant returns(uint256)
func (_Vita *VitaSession) IncrementalSupply() (*big.Int, error) {
	return _Vita.Contract.IncrementalSupply(&_Vita.CallOpts)
}

// IncrementalSupply is a free data retrieval call binding the contract method 0xe26204d7.
//
// Solidity: function incrementalSupply() constant returns(uint256)
func (_Vita *VitaCallerSession) IncrementalSupply() (*big.Int, error) {
	return _Vita.Contract.IncrementalSupply(&_Vita.CallOpts)
}

// LastClaimViewIDs is a free data retrieval call binding the contract method 0x3609c052.
//
// Solidity: function lastClaimViewIDs(address ) constant returns(uint256)
func (_Vita *VitaCaller) LastClaimViewIDs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "lastClaimViewIDs", arg0)
	return *ret0, err
}

// LastClaimViewIDs is a free data retrieval call binding the contract method 0x3609c052.
//
// Solidity: function lastClaimViewIDs(address ) constant returns(uint256)
func (_Vita *VitaSession) LastClaimViewIDs(arg0 common.Address) (*big.Int, error) {
	return _Vita.Contract.LastClaimViewIDs(&_Vita.CallOpts, arg0)
}

// LastClaimViewIDs is a free data retrieval call binding the contract method 0x3609c052.
//
// Solidity: function lastClaimViewIDs(address ) constant returns(uint256)
func (_Vita *VitaCallerSession) LastClaimViewIDs(arg0 common.Address) (*big.Int, error) {
	return _Vita.Contract.LastClaimViewIDs(&_Vita.CallOpts, arg0)
}

// LastDecayHeight is a free data retrieval call binding the contract method 0x12fb4475.
//
// Solidity: function lastDecayHeight() constant returns(uint256)
func (_Vita *VitaCaller) LastDecayHeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "lastDecayHeight")
	return *ret0, err
}

// LastDecayHeight is a free data retrieval call binding the contract method 0x12fb4475.
//
// Solidity: function lastDecayHeight() constant returns(uint256)
func (_Vita *VitaSession) LastDecayHeight() (*big.Int, error) {
	return _Vita.Contract.LastDecayHeight(&_Vita.CallOpts)
}

// LastDecayHeight is a free data retrieval call binding the contract method 0x12fb4475.
//
// Solidity: function lastDecayHeight() constant returns(uint256)
func (_Vita *VitaCallerSession) LastDecayHeight() (*big.Int, error) {
	return _Vita.Contract.LastDecayHeight(&_Vita.CallOpts)
}

// LastDonationPoolClaimViewID is a free data retrieval call binding the contract method 0x64a23687.
//
// Solidity: function lastDonationPoolClaimViewID() constant returns(uint256)
func (_Vita *VitaCaller) LastDonationPoolClaimViewID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "lastDonationPoolClaimViewID")
	return *ret0, err
}

// LastDonationPoolClaimViewID is a free data retrieval call binding the contract method 0x64a23687.
//
// Solidity: function lastDonationPoolClaimViewID() constant returns(uint256)
func (_Vita *VitaSession) LastDonationPoolClaimViewID() (*big.Int, error) {
	return _Vita.Contract.LastDonationPoolClaimViewID(&_Vita.CallOpts)
}

// LastDonationPoolClaimViewID is a free data retrieval call binding the contract method 0x64a23687.
//
// Solidity: function lastDonationPoolClaimViewID() constant returns(uint256)
func (_Vita *VitaCallerSession) LastDonationPoolClaimViewID() (*big.Int, error) {
	return _Vita.Contract.LastDonationPoolClaimViewID(&_Vita.CallOpts)
}

// LastRewardPoolClaimViewID is a free data retrieval call binding the contract method 0xca8153aa.
//
// Solidity: function lastRewardPoolClaimViewID() constant returns(uint256)
func (_Vita *VitaCaller) LastRewardPoolClaimViewID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "lastRewardPoolClaimViewID")
	return *ret0, err
}

// LastRewardPoolClaimViewID is a free data retrieval call binding the contract method 0xca8153aa.
//
// Solidity: function lastRewardPoolClaimViewID() constant returns(uint256)
func (_Vita *VitaSession) LastRewardPoolClaimViewID() (*big.Int, error) {
	return _Vita.Contract.LastRewardPoolClaimViewID(&_Vita.CallOpts)
}

// LastRewardPoolClaimViewID is a free data retrieval call binding the contract method 0xca8153aa.
//
// Solidity: function lastRewardPoolClaimViewID() constant returns(uint256)
func (_Vita *VitaCallerSession) LastRewardPoolClaimViewID() (*big.Int, error) {
	return _Vita.Contract.LastRewardPoolClaimViewID(&_Vita.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Vita *VitaCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Vita *VitaSession) Name() (string, error) {
	return _Vita.Contract.Name(&_Vita.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Vita *VitaCallerSession) Name() (string, error) {
	return _Vita.Contract.Name(&_Vita.CallOpts)
}

// NextDecayHeight is a free data retrieval call binding the contract method 0x80ad454e.
//
// Solidity: function nextDecayHeight() constant returns(uint256)
func (_Vita *VitaCaller) NextDecayHeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "nextDecayHeight")
	return *ret0, err
}

// NextDecayHeight is a free data retrieval call binding the contract method 0x80ad454e.
//
// Solidity: function nextDecayHeight() constant returns(uint256)
func (_Vita *VitaSession) NextDecayHeight() (*big.Int, error) {
	return _Vita.Contract.NextDecayHeight(&_Vita.CallOpts)
}

// NextDecayHeight is a free data retrieval call binding the contract method 0x80ad454e.
//
// Solidity: function nextDecayHeight() constant returns(uint256)
func (_Vita *VitaCallerSession) NextDecayHeight() (*big.Int, error) {
	return _Vita.Contract.NextDecayHeight(&_Vita.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vita *VitaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vita *VitaSession) Owner() (common.Address, error) {
	return _Vita.Contract.Owner(&_Vita.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vita *VitaCallerSession) Owner() (common.Address, error) {
	return _Vita.Contract.Owner(&_Vita.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Vita *VitaCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Vita *VitaSession) Paused() (bool, error) {
	return _Vita.Contract.Paused(&_Vita.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Vita *VitaCallerSession) Paused() (bool, error) {
	return _Vita.Contract.Paused(&_Vita.CallOpts)
}

// RewardPoolAddress is a free data retrieval call binding the contract method 0x845a51ec.
//
// Solidity: function rewardPoolAddress() constant returns(address)
func (_Vita *VitaCaller) RewardPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "rewardPoolAddress")
	return *ret0, err
}

// RewardPoolAddress is a free data retrieval call binding the contract method 0x845a51ec.
//
// Solidity: function rewardPoolAddress() constant returns(address)
func (_Vita *VitaSession) RewardPoolAddress() (common.Address, error) {
	return _Vita.Contract.RewardPoolAddress(&_Vita.CallOpts)
}

// RewardPoolAddress is a free data retrieval call binding the contract method 0x845a51ec.
//
// Solidity: function rewardPoolAddress() constant returns(address)
func (_Vita *VitaCallerSession) RewardPoolAddress() (common.Address, error) {
	return _Vita.Contract.RewardPoolAddress(&_Vita.CallOpts)
}

// RewardPoolSize is a free data retrieval call binding the contract method 0x211b827d.
//
// Solidity: function rewardPoolSize() constant returns(uint256)
func (_Vita *VitaCaller) RewardPoolSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "rewardPoolSize")
	return *ret0, err
}

// RewardPoolSize is a free data retrieval call binding the contract method 0x211b827d.
//
// Solidity: function rewardPoolSize() constant returns(uint256)
func (_Vita *VitaSession) RewardPoolSize() (*big.Int, error) {
	return _Vita.Contract.RewardPoolSize(&_Vita.CallOpts)
}

// RewardPoolSize is a free data retrieval call binding the contract method 0x211b827d.
//
// Solidity: function rewardPoolSize() constant returns(uint256)
func (_Vita *VitaCallerSession) RewardPoolSize() (*big.Int, error) {
	return _Vita.Contract.RewardPoolSize(&_Vita.CallOpts)
}

// StakingPoolSize is a free data retrieval call binding the contract method 0x3bcc14d4.
//
// Solidity: function stakingPoolSize() constant returns(uint256)
func (_Vita *VitaCaller) StakingPoolSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "stakingPoolSize")
	return *ret0, err
}

// StakingPoolSize is a free data retrieval call binding the contract method 0x3bcc14d4.
//
// Solidity: function stakingPoolSize() constant returns(uint256)
func (_Vita *VitaSession) StakingPoolSize() (*big.Int, error) {
	return _Vita.Contract.StakingPoolSize(&_Vita.CallOpts)
}

// StakingPoolSize is a free data retrieval call binding the contract method 0x3bcc14d4.
//
// Solidity: function stakingPoolSize() constant returns(uint256)
func (_Vita *VitaCallerSession) StakingPoolSize() (*big.Int, error) {
	return _Vita.Contract.StakingPoolSize(&_Vita.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Vita *VitaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Vita *VitaSession) Symbol() (string, error) {
	return _Vita.Contract.Symbol(&_Vita.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Vita *VitaCallerSession) Symbol() (string, error) {
	return _Vita.Contract.Symbol(&_Vita.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Vita *VitaCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Vita *VitaSession) TotalSupply() (*big.Int, error) {
	return _Vita.Contract.TotalSupply(&_Vita.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Vita *VitaCallerSession) TotalSupply() (*big.Int, error) {
	return _Vita.Contract.TotalSupply(&_Vita.CallOpts)
}

// Vps is a free data retrieval call binding the contract method 0x56702d6e.
//
// Solidity: function vps() constant returns(address)
func (_Vita *VitaCaller) Vps(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vita.contract.Call(opts, out, "vps")
	return *ret0, err
}

// Vps is a free data retrieval call binding the contract method 0x56702d6e.
//
// Solidity: function vps() constant returns(address)
func (_Vita *VitaSession) Vps() (common.Address, error) {
	return _Vita.Contract.Vps(&_Vita.CallOpts)
}

// Vps is a free data retrieval call binding the contract method 0x56702d6e.
//
// Solidity: function vps() constant returns(address)
func (_Vita *VitaCallerSession) Vps() (common.Address, error) {
	return _Vita.Contract.Vps(&_Vita.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Vita *VitaTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Vita *VitaSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.Approve(&_Vita.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Vita *VitaTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.Approve(&_Vita.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Vita *VitaTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Vita *VitaSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.Burn(&_Vita.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Vita *VitaTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.Burn(&_Vita.TransactOpts, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vita *VitaTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vita *VitaSession) Claim() (*types.Transaction, error) {
	return _Vita.Contract.Claim(&_Vita.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vita *VitaTransactorSession) Claim() (*types.Transaction, error) {
	return _Vita.Contract.Claim(&_Vita.TransactOpts)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_Vita *VitaTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_Vita *VitaSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.DecreaseApproval(&_Vita.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_Vita *VitaTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.DecreaseApproval(&_Vita.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_Vita *VitaTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_Vita *VitaSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.IncreaseApproval(&_Vita.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_Vita *VitaTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.IncreaseApproval(&_Vita.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vita *VitaTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vita *VitaSession) Pause() (*types.Transaction, error) {
	return _Vita.Contract.Pause(&_Vita.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vita *VitaTransactorSession) Pause() (*types.Transaction, error) {
	return _Vita.Contract.Pause(&_Vita.TransactOpts)
}

// SetDonationPoolAddress is a paid mutator transaction binding the contract method 0xed644211.
//
// Solidity: function setDonationPoolAddress(address _newDonationPool) returns()
func (_Vita *VitaTransactor) SetDonationPoolAddress(opts *bind.TransactOpts, _newDonationPool common.Address) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "setDonationPoolAddress", _newDonationPool)
}

// SetDonationPoolAddress is a paid mutator transaction binding the contract method 0xed644211.
//
// Solidity: function setDonationPoolAddress(address _newDonationPool) returns()
func (_Vita *VitaSession) SetDonationPoolAddress(_newDonationPool common.Address) (*types.Transaction, error) {
	return _Vita.Contract.SetDonationPoolAddress(&_Vita.TransactOpts, _newDonationPool)
}

// SetDonationPoolAddress is a paid mutator transaction binding the contract method 0xed644211.
//
// Solidity: function setDonationPoolAddress(address _newDonationPool) returns()
func (_Vita *VitaTransactorSession) SetDonationPoolAddress(_newDonationPool common.Address) (*types.Transaction, error) {
	return _Vita.Contract.SetDonationPoolAddress(&_Vita.TransactOpts, _newDonationPool)
}

// SetRewardPoolAddress is a paid mutator transaction binding the contract method 0xb24cf5d7.
//
// Solidity: function setRewardPoolAddress(address _newRewardPool) returns()
func (_Vita *VitaTransactor) SetRewardPoolAddress(opts *bind.TransactOpts, _newRewardPool common.Address) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "setRewardPoolAddress", _newRewardPool)
}

// SetRewardPoolAddress is a paid mutator transaction binding the contract method 0xb24cf5d7.
//
// Solidity: function setRewardPoolAddress(address _newRewardPool) returns()
func (_Vita *VitaSession) SetRewardPoolAddress(_newRewardPool common.Address) (*types.Transaction, error) {
	return _Vita.Contract.SetRewardPoolAddress(&_Vita.TransactOpts, _newRewardPool)
}

// SetRewardPoolAddress is a paid mutator transaction binding the contract method 0xb24cf5d7.
//
// Solidity: function setRewardPoolAddress(address _newRewardPool) returns()
func (_Vita *VitaTransactorSession) SetRewardPoolAddress(_newRewardPool common.Address) (*types.Transaction, error) {
	return _Vita.Contract.SetRewardPoolAddress(&_Vita.TransactOpts, _newRewardPool)
}

// SetVPS is a paid mutator transaction binding the contract method 0x3d8384b6.
//
// Solidity: function setVPS(address _newVPS) returns()
func (_Vita *VitaTransactor) SetVPS(opts *bind.TransactOpts, _newVPS common.Address) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "setVPS", _newVPS)
}

// SetVPS is a paid mutator transaction binding the contract method 0x3d8384b6.
//
// Solidity: function setVPS(address _newVPS) returns()
func (_Vita *VitaSession) SetVPS(_newVPS common.Address) (*types.Transaction, error) {
	return _Vita.Contract.SetVPS(&_Vita.TransactOpts, _newVPS)
}

// SetVPS is a paid mutator transaction binding the contract method 0x3d8384b6.
//
// Solidity: function setVPS(address _newVPS) returns()
func (_Vita *VitaTransactorSession) SetVPS(_newVPS common.Address) (*types.Transaction, error) {
	return _Vita.Contract.SetVPS(&_Vita.TransactOpts, _newVPS)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Vita *VitaTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Vita *VitaSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.Transfer(&_Vita.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Vita *VitaTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.Transfer(&_Vita.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Vita *VitaTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Vita *VitaSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.TransferFrom(&_Vita.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Vita *VitaTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Vita.Contract.TransferFrom(&_Vita.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vita *VitaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vita *VitaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vita.Contract.TransferOwnership(&_Vita.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vita *VitaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vita.Contract.TransferOwnership(&_Vita.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vita *VitaTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vita.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vita *VitaSession) Unpause() (*types.Transaction, error) {
	return _Vita.Contract.Unpause(&_Vita.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vita *VitaTransactorSession) Unpause() (*types.Transaction, error) {
	return _Vita.Contract.Unpause(&_Vita.TransactOpts)
}

// VitaApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vita contract.
type VitaApprovalIterator struct {
	Event *VitaApproval // Event containing the contract specifics and raw log

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
func (it *VitaApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VitaApproval)
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
		it.Event = new(VitaApproval)
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
func (it *VitaApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VitaApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VitaApproval represents a Approval event raised by the Vita contract.
type VitaApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Vita *VitaFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VitaApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vita.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VitaApprovalIterator{contract: _Vita.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Vita *VitaFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VitaApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vita.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VitaApproval)
				if err := _Vita.contract.UnpackLog(event, "Approval", log); err != nil {
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

// VitaClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Vita contract.
type VitaClaimIterator struct {
	Event *VitaClaim // Event containing the contract specifics and raw log

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
func (it *VitaClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VitaClaim)
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
		it.Event = new(VitaClaim)
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
func (it *VitaClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VitaClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VitaClaim represents a Claim event raised by the Vita contract.
type VitaClaim struct {
	Claimer common.Address
	Amount  *big.Int
	ViewID  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address claimer, uint256 amount, uint256 viewID)
func (_Vita *VitaFilterer) FilterClaim(opts *bind.FilterOpts) (*VitaClaimIterator, error) {

	logs, sub, err := _Vita.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &VitaClaimIterator{contract: _Vita.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address claimer, uint256 amount, uint256 viewID)
func (_Vita *VitaFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *VitaClaim) (event.Subscription, error) {

	logs, sub, err := _Vita.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VitaClaim)
				if err := _Vita.contract.UnpackLog(event, "Claim", log); err != nil {
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

// VitaDecayIterator is returned from FilterDecay and is used to iterate over the raw logs and unpacked data for Decay events raised by the Vita contract.
type VitaDecayIterator struct {
	Event *VitaDecay // Event containing the contract specifics and raw log

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
func (it *VitaDecayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VitaDecay)
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
		it.Event = new(VitaDecay)
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
func (it *VitaDecayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VitaDecayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VitaDecay represents a Decay event raised by the Vita contract.
type VitaDecay struct {
	Height            *big.Int
	IncremetnalSupply *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDecay is a free log retrieval operation binding the contract event 0xb78fecc4804d4e8de58f70e839d77affe1cd72a59c58f87bafad35eb08813679.
//
// Solidity: event Decay(uint256 height, uint256 incremetnalSupply)
func (_Vita *VitaFilterer) FilterDecay(opts *bind.FilterOpts) (*VitaDecayIterator, error) {

	logs, sub, err := _Vita.contract.FilterLogs(opts, "Decay")
	if err != nil {
		return nil, err
	}
	return &VitaDecayIterator{contract: _Vita.contract, event: "Decay", logs: logs, sub: sub}, nil
}

// WatchDecay is a free log subscription operation binding the contract event 0xb78fecc4804d4e8de58f70e839d77affe1cd72a59c58f87bafad35eb08813679.
//
// Solidity: event Decay(uint256 height, uint256 incremetnalSupply)
func (_Vita *VitaFilterer) WatchDecay(opts *bind.WatchOpts, sink chan<- *VitaDecay) (event.Subscription, error) {

	logs, sub, err := _Vita.contract.WatchLogs(opts, "Decay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VitaDecay)
				if err := _Vita.contract.UnpackLog(event, "Decay", log); err != nil {
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

// VitaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vita contract.
type VitaOwnershipTransferredIterator struct {
	Event *VitaOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VitaOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VitaOwnershipTransferred)
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
		it.Event = new(VitaOwnershipTransferred)
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
func (it *VitaOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VitaOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VitaOwnershipTransferred represents a OwnershipTransferred event raised by the Vita contract.
type VitaOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vita *VitaFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VitaOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vita.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VitaOwnershipTransferredIterator{contract: _Vita.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vita *VitaFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VitaOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vita.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VitaOwnershipTransferred)
				if err := _Vita.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// VitaPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Vita contract.
type VitaPauseIterator struct {
	Event *VitaPause // Event containing the contract specifics and raw log

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
func (it *VitaPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VitaPause)
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
		it.Event = new(VitaPause)
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
func (it *VitaPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VitaPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VitaPause represents a Pause event raised by the Vita contract.
type VitaPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Vita *VitaFilterer) FilterPause(opts *bind.FilterOpts) (*VitaPauseIterator, error) {

	logs, sub, err := _Vita.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &VitaPauseIterator{contract: _Vita.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Vita *VitaFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *VitaPause) (event.Subscription, error) {

	logs, sub, err := _Vita.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VitaPause)
				if err := _Vita.contract.UnpackLog(event, "Pause", log); err != nil {
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

// VitaTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vita contract.
type VitaTransferIterator struct {
	Event *VitaTransfer // Event containing the contract specifics and raw log

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
func (it *VitaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VitaTransfer)
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
		it.Event = new(VitaTransfer)
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
func (it *VitaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VitaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VitaTransfer represents a Transfer event raised by the Vita contract.
type VitaTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Vita *VitaFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VitaTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vita.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VitaTransferIterator{contract: _Vita.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Vita *VitaFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VitaTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vita.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VitaTransfer)
				if err := _Vita.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// VitaUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Vita contract.
type VitaUnpauseIterator struct {
	Event *VitaUnpause // Event containing the contract specifics and raw log

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
func (it *VitaUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VitaUnpause)
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
		it.Event = new(VitaUnpause)
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
func (it *VitaUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VitaUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VitaUnpause represents a Unpause event raised by the Vita contract.
type VitaUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Vita *VitaFilterer) FilterUnpause(opts *bind.FilterOpts) (*VitaUnpauseIterator, error) {

	logs, sub, err := _Vita.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &VitaUnpauseIterator{contract: _Vita.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Vita *VitaFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *VitaUnpause) (event.Subscription, error) {

	logs, sub, err := _Vita.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VitaUnpause)
				if err := _Vita.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// VitaUpdateViewIterator is returned from FilterUpdateView and is used to iterate over the raw logs and unpacked data for UpdateView events raised by the Vita contract.
type VitaUpdateViewIterator struct {
	Event *VitaUpdateView // Event containing the contract specifics and raw log

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
func (it *VitaUpdateViewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VitaUpdateView)
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
		it.Event = new(VitaUpdateView)
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
func (it *VitaUpdateViewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VitaUpdateViewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VitaUpdateView represents a UpdateView event raised by the Vita contract.
type VitaUpdateView struct {
	ViewID *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateView is a free log retrieval operation binding the contract event 0x2087ba3f49fc996df288fe78f79bfa9d28edef119f970f375e2c7305fc4e2bad.
//
// Solidity: event UpdateView(uint256 viewID)
func (_Vita *VitaFilterer) FilterUpdateView(opts *bind.FilterOpts) (*VitaUpdateViewIterator, error) {

	logs, sub, err := _Vita.contract.FilterLogs(opts, "UpdateView")
	if err != nil {
		return nil, err
	}
	return &VitaUpdateViewIterator{contract: _Vita.contract, event: "UpdateView", logs: logs, sub: sub}, nil
}

// WatchUpdateView is a free log subscription operation binding the contract event 0x2087ba3f49fc996df288fe78f79bfa9d28edef119f970f375e2c7305fc4e2bad.
//
// Solidity: event UpdateView(uint256 viewID)
func (_Vita *VitaFilterer) WatchUpdateView(opts *bind.WatchOpts, sink chan<- *VitaUpdateView) (event.Subscription, error) {

	logs, sub, err := _Vita.contract.WatchLogs(opts, "UpdateView")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VitaUpdateView)
				if err := _Vita.contract.UnpackLog(event, "UpdateView", log); err != nil {
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
