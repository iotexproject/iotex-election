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

// AgentABI is the input ABI used to generate the binding from.
const AgentABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vita\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AUTHENTICATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"claimFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimVita\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"digest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastClaimCycles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vita\",\"outputs\":[{\"internalType\":\"contractVita\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Agent is an auto generated Go binding around an Ethereum contract.
type Agent struct {
	AgentCaller     // Read-only binding to the contract
	AgentTransactor // Write-only binding to the contract
	AgentFilterer   // Log filterer for contract events
}

// AgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentSession struct {
	Contract     *Agent            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentCallerSession struct {
	Contract *AgentCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentTransactorSession struct {
	Contract     *AgentTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentRaw struct {
	Contract *Agent // Generic contract binding to access the raw methods on
}

// AgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentCallerRaw struct {
	Contract *AgentCaller // Generic read-only contract binding to access the raw methods on
}

// AgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentTransactorRaw struct {
	Contract *AgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgent creates a new instance of Agent, bound to a specific deployed contract.
func NewAgent(address common.Address, backend bind.ContractBackend) (*Agent, error) {
	contract, err := bindAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agent{AgentCaller: AgentCaller{contract: contract}, AgentTransactor: AgentTransactor{contract: contract}, AgentFilterer: AgentFilterer{contract: contract}}, nil
}

// NewAgentCaller creates a new read-only instance of Agent, bound to a specific deployed contract.
func NewAgentCaller(address common.Address, caller bind.ContractCaller) (*AgentCaller, error) {
	contract, err := bindAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentCaller{contract: contract}, nil
}

// NewAgentTransactor creates a new write-only instance of Agent, bound to a specific deployed contract.
func NewAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentTransactor, error) {
	contract, err := bindAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentTransactor{contract: contract}, nil
}

// NewAgentFilterer creates a new log filterer instance of Agent, bound to a specific deployed contract.
func NewAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentFilterer, error) {
	contract, err := bindAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentFilterer{contract: contract}, nil
}

// bindAgent binds a generic wrapper to an already deployed contract.
func bindAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agent *AgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agent.Contract.AgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agent *AgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.Contract.AgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agent *AgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agent.Contract.AgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agent *AgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agent *AgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agent *AgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agent.Contract.contract.Transact(opts, method, params...)
}

// AUTHENTICATIONTYPEHASH is a free data retrieval call binding the contract method 0xb2ae7d46.
//
// Solidity: function AUTHENTICATION_TYPEHASH() view returns(bytes32)
func (_Agent *AgentCaller) AUTHENTICATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "AUTHENTICATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUTHENTICATIONTYPEHASH is a free data retrieval call binding the contract method 0xb2ae7d46.
//
// Solidity: function AUTHENTICATION_TYPEHASH() view returns(bytes32)
func (_Agent *AgentSession) AUTHENTICATIONTYPEHASH() ([32]byte, error) {
	return _Agent.Contract.AUTHENTICATIONTYPEHASH(&_Agent.CallOpts)
}

// AUTHENTICATIONTYPEHASH is a free data retrieval call binding the contract method 0xb2ae7d46.
//
// Solidity: function AUTHENTICATION_TYPEHASH() view returns(bytes32)
func (_Agent *AgentCallerSession) AUTHENTICATIONTYPEHASH() ([32]byte, error) {
	return _Agent.Contract.AUTHENTICATIONTYPEHASH(&_Agent.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Agent *AgentCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Agent *AgentSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Agent.Contract.DOMAINSEPARATOR(&_Agent.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Agent *AgentCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Agent.Contract.DOMAINSEPARATOR(&_Agent.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_Agent *AgentCaller) EIP712DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "EIP712DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_Agent *AgentSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _Agent.Contract.EIP712DOMAINTYPEHASH(&_Agent.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_Agent *AgentCallerSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _Agent.Contract.EIP712DOMAINTYPEHASH(&_Agent.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address _user) view returns(uint256, uint256, bool)
func (_Agent *AgentCaller) Claimed(opts *bind.CallOpts, _user common.Address) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "claimed", _user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address _user) view returns(uint256, uint256, bool)
func (_Agent *AgentSession) Claimed(_user common.Address) (*big.Int, *big.Int, bool, error) {
	return _Agent.Contract.Claimed(&_Agent.CallOpts, _user)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address _user) view returns(uint256, uint256, bool)
func (_Agent *AgentCallerSession) Claimed(_user common.Address) (*big.Int, *big.Int, bool, error) {
	return _Agent.Contract.Claimed(&_Agent.CallOpts, _user)
}

// Digest is a free data retrieval call binding the contract method 0xfe54bd15.
//
// Solidity: function digest(address wallet, uint256 amount, uint256 deadline) view returns(bytes32)
func (_Agent *AgentCaller) Digest(opts *bind.CallOpts, wallet common.Address, amount *big.Int, deadline *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "digest", wallet, amount, deadline)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Digest is a free data retrieval call binding the contract method 0xfe54bd15.
//
// Solidity: function digest(address wallet, uint256 amount, uint256 deadline) view returns(bytes32)
func (_Agent *AgentSession) Digest(wallet common.Address, amount *big.Int, deadline *big.Int) ([32]byte, error) {
	return _Agent.Contract.Digest(&_Agent.CallOpts, wallet, amount, deadline)
}

// Digest is a free data retrieval call binding the contract method 0xfe54bd15.
//
// Solidity: function digest(address wallet, uint256 amount, uint256 deadline) view returns(bytes32)
func (_Agent *AgentCallerSession) Digest(wallet common.Address, amount *big.Int, deadline *big.Int) ([32]byte, error) {
	return _Agent.Contract.Digest(&_Agent.CallOpts, wallet, amount, deadline)
}

// LastClaimCycles is a free data retrieval call binding the contract method 0xb1629241.
//
// Solidity: function lastClaimCycles(address ) view returns(uint256)
func (_Agent *AgentCaller) LastClaimCycles(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "lastClaimCycles", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastClaimCycles is a free data retrieval call binding the contract method 0xb1629241.
//
// Solidity: function lastClaimCycles(address ) view returns(uint256)
func (_Agent *AgentSession) LastClaimCycles(arg0 common.Address) (*big.Int, error) {
	return _Agent.Contract.LastClaimCycles(&_Agent.CallOpts, arg0)
}

// LastClaimCycles is a free data retrieval call binding the contract method 0xb1629241.
//
// Solidity: function lastClaimCycles(address ) view returns(uint256)
func (_Agent *AgentCallerSession) LastClaimCycles(arg0 common.Address) (*big.Int, error) {
	return _Agent.Contract.LastClaimCycles(&_Agent.CallOpts, arg0)
}

// Notary is a free data retrieval call binding the contract method 0x9d54c79d.
//
// Solidity: function notary() view returns(address)
func (_Agent *AgentCaller) Notary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "notary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Notary is a free data retrieval call binding the contract method 0x9d54c79d.
//
// Solidity: function notary() view returns(address)
func (_Agent *AgentSession) Notary() (common.Address, error) {
	return _Agent.Contract.Notary(&_Agent.CallOpts)
}

// Notary is a free data retrieval call binding the contract method 0x9d54c79d.
//
// Solidity: function notary() view returns(address)
func (_Agent *AgentCallerSession) Notary() (common.Address, error) {
	return _Agent.Contract.Notary(&_Agent.CallOpts)
}

// PoolSize is a free data retrieval call binding the contract method 0x4ec18db9.
//
// Solidity: function poolSize() view returns(uint256)
func (_Agent *AgentCaller) PoolSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "poolSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolSize is a free data retrieval call binding the contract method 0x4ec18db9.
//
// Solidity: function poolSize() view returns(uint256)
func (_Agent *AgentSession) PoolSize() (*big.Int, error) {
	return _Agent.Contract.PoolSize(&_Agent.CallOpts)
}

// PoolSize is a free data retrieval call binding the contract method 0x4ec18db9.
//
// Solidity: function poolSize() view returns(uint256)
func (_Agent *AgentCallerSession) PoolSize() (*big.Int, error) {
	return _Agent.Contract.PoolSize(&_Agent.CallOpts)
}

// Vita is a free data retrieval call binding the contract method 0x393d9bb3.
//
// Solidity: function vita() view returns(address)
func (_Agent *AgentCaller) Vita(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "vita")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vita is a free data retrieval call binding the contract method 0x393d9bb3.
//
// Solidity: function vita() view returns(address)
func (_Agent *AgentSession) Vita() (common.Address, error) {
	return _Agent.Contract.Vita(&_Agent.CallOpts)
}

// Vita is a free data retrieval call binding the contract method 0x393d9bb3.
//
// Solidity: function vita() view returns(address)
func (_Agent *AgentCallerSession) Vita() (common.Address, error) {
	return _Agent.Contract.Vita(&_Agent.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xfa6b3b53.
//
// Solidity: function claim(uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Agent *AgentTransactor) Claim(opts *bind.TransactOpts, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "claim", _amount, _deadline, _v, _r, _s)
}

// Claim is a paid mutator transaction binding the contract method 0xfa6b3b53.
//
// Solidity: function claim(uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Agent *AgentSession) Claim(_amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Agent.Contract.Claim(&_Agent.TransactOpts, _amount, _deadline, _v, _r, _s)
}

// Claim is a paid mutator transaction binding the contract method 0xfa6b3b53.
//
// Solidity: function claim(uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Agent *AgentTransactorSession) Claim(_amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Agent.Contract.Claim(&_Agent.TransactOpts, _amount, _deadline, _v, _r, _s)
}

// ClaimFor is a paid mutator transaction binding the contract method 0xa34010a8.
//
// Solidity: function claimFor(address _owner, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Agent *AgentTransactor) ClaimFor(opts *bind.TransactOpts, _owner common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "claimFor", _owner, _amount, _deadline, _v, _r, _s)
}

// ClaimFor is a paid mutator transaction binding the contract method 0xa34010a8.
//
// Solidity: function claimFor(address _owner, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Agent *AgentSession) ClaimFor(_owner common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Agent.Contract.ClaimFor(&_Agent.TransactOpts, _owner, _amount, _deadline, _v, _r, _s)
}

// ClaimFor is a paid mutator transaction binding the contract method 0xa34010a8.
//
// Solidity: function claimFor(address _owner, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Agent *AgentTransactorSession) ClaimFor(_owner common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Agent.Contract.ClaimFor(&_Agent.TransactOpts, _owner, _amount, _deadline, _v, _r, _s)
}

// ClaimVita is a paid mutator transaction binding the contract method 0x77a7af06.
//
// Solidity: function claimVita() returns()
func (_Agent *AgentTransactor) ClaimVita(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "claimVita")
}

// ClaimVita is a paid mutator transaction binding the contract method 0x77a7af06.
//
// Solidity: function claimVita() returns()
func (_Agent *AgentSession) ClaimVita() (*types.Transaction, error) {
	return _Agent.Contract.ClaimVita(&_Agent.TransactOpts)
}

// ClaimVita is a paid mutator transaction binding the contract method 0x77a7af06.
//
// Solidity: function claimVita() returns()
func (_Agent *AgentTransactorSession) ClaimVita() (*types.Transaction, error) {
	return _Agent.Contract.ClaimVita(&_Agent.TransactOpts)
}
