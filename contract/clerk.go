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

// ClerkBin is the compiled bytecode used for deploying new contracts.
const ClerkBin = `{
	"linkReferences": {},
	"object": "608060405234801561001057600080fd5b5060405160208061101783398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610f53806100c46000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806324953eaa146100a9578063286dd3f514610127578063393d9bb3146101825780634e71d92d146101d95780637b9417c8146101f05780638da5cb5b1461024b5780639b19251a146102a2578063e2ec6ec3146102fd578063f2fde38b1461037b578063ffc3a769146103be575b600080fd5b3480156100b557600080fd5b5061010d60048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610411565b604051808215151515815260200191505060405180910390f35b34801561013357600080fd5b50610168600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104ba565b604051808215151515815260200191505060405180910390f35b34801561018e57600080fd5b5061019761062e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101e557600080fd5b506101ee610654565b005b3480156101fc57600080fd5b50610231600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106f4565b604051808215151515815260200191505060405180910390f35b34801561025757600080fd5b50610260610868565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102ae57600080fd5b506102e3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061088d565b604051808215151515815260200191505060405180910390f35b34801561030957600080fd5b50610361600480360381019080803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091929192905050506108ad565b604051808215151515815260200191505060405180910390f35b34801561038757600080fd5b506103bc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610956565b005b3480156103ca57600080fd5b5061040f600480360381019080803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390505050610aab565b005b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561046f57600080fd5b600090505b82518110156104b45761049d838281518110151561048e57fe5b906020019060200201516104ba565b156104a757600191505b8080600101915050610474565b50919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561051757600080fd5b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610629576000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507ff1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a82604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1600190505b919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634e71d92d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156106da57600080fd5b505af11580156106ee573d6000803e3d6000fd5b50505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561075157600080fd5b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156108635760018060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f82604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1600190505b919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915054906101000a900460ff1681565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561090b57600080fd5b600090505b825181101561095057610939838281518110151561092a57fe5b906020019060200201516106f4565b1561094357600191505b8080600101915050610910565b50919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109b157600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156109ed57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610b0657600080fd5b600086869050111515610b81576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f726563656976657273206973206f66206c656e6774682030000000000000000081525060200191505060405180910390fd5b8383905086869050141515610c24576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f746865206c656e67746873206f662072656365697665727320616e6420616d6f81526020017f756e7473206e6f7420657175616c00000000000000000000000000000000000081525060400191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610ce157600080fd5b505af1158015610cf5573d6000803e3d6000fd5b505050506040513d6020811015610d0b57600080fd5b81019080805190602001909291905050509150600090505b85859050811015610f1f578383828181101515610d3c57fe5b905060200201358210151515610dba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e73756666696369656e742062616c616e636500000000000000000000000081525060200191505060405180910390fd5b8383828181101515610dc857fe5b9050602002013582039150600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8787848181101515610e1f57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff168686858181101515610e4a57fe5b905060200201356040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610ed657600080fd5b505af1158015610eea573d6000803e3d6000fd5b505050506040513d6020811015610f0057600080fd5b8101908080519060200190929190505050508080600101915050610d23565b5050505050505600a165627a7a723058206c7b37483d0b320ee41d5c1bef74a572f611abde0a8bf31852b9d3de8b0a14df0029",
	"opcodes": "PUSH1 0x80 PUSH1 0x40 MSTORE CALLVALUE DUP1 ISZERO PUSH2 0x10 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH1 0x40 MLOAD PUSH1 0x20 DUP1 PUSH2 0x1017 DUP4 CODECOPY DUP2 ADD DUP1 PUSH1 0x40 MSTORE DUP2 ADD SWAP1 DUP1 DUP1 MLOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP CALLER PUSH1 0x0 DUP1 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP DUP1 PUSH1 0x2 PUSH1 0x0 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP POP PUSH2 0xF53 DUP1 PUSH2 0xC4 PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN STOP PUSH1 0x80 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH2 0xA4 JUMPI PUSH1 0x0 CALLDATALOAD PUSH29 0x100000000000000000000000000000000000000000000000000000000 SWAP1 DIV PUSH4 0xFFFFFFFF AND DUP1 PUSH4 0x24953EAA EQ PUSH2 0xA9 JUMPI DUP1 PUSH4 0x286DD3F5 EQ PUSH2 0x127 JUMPI DUP1 PUSH4 0x393D9BB3 EQ PUSH2 0x182 JUMPI DUP1 PUSH4 0x4E71D92D EQ PUSH2 0x1D9 JUMPI DUP1 PUSH4 0x7B9417C8 EQ PUSH2 0x1F0 JUMPI DUP1 PUSH4 0x8DA5CB5B EQ PUSH2 0x24B JUMPI DUP1 PUSH4 0x9B19251A EQ PUSH2 0x2A2 JUMPI DUP1 PUSH4 0xE2EC6EC3 EQ PUSH2 0x2FD JUMPI DUP1 PUSH4 0xF2FDE38B EQ PUSH2 0x37B JUMPI DUP1 PUSH4 0xFFC3A769 EQ PUSH2 0x3BE JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0xB5 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x10D PUSH1 0x4 DUP1 CALLDATASIZE SUB DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 DUP3 ADD DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 DUP1 DUP1 PUSH1 0x20 MUL PUSH1 0x20 ADD PUSH1 0x40 MLOAD SWAP1 DUP2 ADD PUSH1 0x40 MSTORE DUP1 SWAP4 SWAP3 SWAP2 SWAP1 DUP2 DUP2 MSTORE PUSH1 0x20 ADD DUP4 DUP4 PUSH1 0x20 MUL DUP1 DUP3 DUP5 CALLDATACOPY DUP3 ADD SWAP2 POP POP POP POP POP POP SWAP2 SWAP3 SWAP2 SWAP3 SWAP1 POP POP POP PUSH2 0x411 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 ISZERO ISZERO ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x133 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x168 PUSH1 0x4 DUP1 CALLDATASIZE SUB DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP PUSH2 0x4BA JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 ISZERO ISZERO ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x18E JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x197 PUSH2 0x62E JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x1E5 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x1EE PUSH2 0x654 JUMP JUMPDEST STOP JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x1FC JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x231 PUSH1 0x4 DUP1 CALLDATASIZE SUB DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP PUSH2 0x6F4 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 ISZERO ISZERO ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x257 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x260 PUSH2 0x868 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x2AE JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x2E3 PUSH1 0x4 DUP1 CALLDATASIZE SUB DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP PUSH2 0x88D JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 ISZERO ISZERO ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x309 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x361 PUSH1 0x4 DUP1 CALLDATASIZE SUB DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 DUP3 ADD DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 DUP1 DUP1 PUSH1 0x20 MUL PUSH1 0x20 ADD PUSH1 0x40 MLOAD SWAP1 DUP2 ADD PUSH1 0x40 MSTORE DUP1 SWAP4 SWAP3 SWAP2 SWAP1 DUP2 DUP2 MSTORE PUSH1 0x20 ADD DUP4 DUP4 PUSH1 0x20 MUL DUP1 DUP3 DUP5 CALLDATACOPY DUP3 ADD SWAP2 POP POP POP POP POP POP SWAP2 SWAP3 SWAP2 SWAP3 SWAP1 POP POP POP PUSH2 0x8AD JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 ISZERO ISZERO ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x387 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x3BC PUSH1 0x4 DUP1 CALLDATASIZE SUB DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP PUSH2 0x956 JUMP JUMPDEST STOP JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x3CA JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x40F PUSH1 0x4 DUP1 CALLDATASIZE SUB DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 DUP3 ADD DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP2 SWAP1 SWAP2 SWAP3 SWAP4 SWAP2 SWAP3 SWAP4 SWAP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 DUP3 ADD DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP2 SWAP1 SWAP2 SWAP3 SWAP4 SWAP2 SWAP3 SWAP4 SWAP1 POP POP POP PUSH2 0xAAB JUMP JUMPDEST STOP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x46F JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 SWAP1 POP JUMPDEST DUP3 MLOAD DUP2 LT ISZERO PUSH2 0x4B4 JUMPI PUSH2 0x49D DUP4 DUP3 DUP2 MLOAD DUP2 LT ISZERO ISZERO PUSH2 0x48E JUMPI INVALID JUMPDEST SWAP1 PUSH1 0x20 ADD SWAP1 PUSH1 0x20 MUL ADD MLOAD PUSH2 0x4BA JUMP JUMPDEST ISZERO PUSH2 0x4A7 JUMPI PUSH1 0x1 SWAP2 POP JUMPDEST DUP1 DUP1 PUSH1 0x1 ADD SWAP2 POP POP PUSH2 0x474 JUMP JUMPDEST POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x517 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x1 PUSH1 0x0 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH1 0xFF AND ISZERO PUSH2 0x629 JUMPI PUSH1 0x0 PUSH1 0x1 PUSH1 0x0 DUP5 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH1 0xFF MUL NOT AND SWAP1 DUP4 ISZERO ISZERO MUL OR SWAP1 SSTORE POP PUSH32 0xF1ABF01A1043B7C244D128E8595CF0C1D10743B022B03A02DFFD8CA3BF729F5A DUP3 PUSH1 0x40 MLOAD DUP1 DUP3 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 LOG1 PUSH1 0x1 SWAP1 POP JUMPDEST SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x2 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 JUMP JUMPDEST PUSH1 0x2 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH4 0x4E71D92D PUSH1 0x40 MLOAD DUP2 PUSH4 0xFFFFFFFF AND PUSH29 0x100000000000000000000000000000000000000000000000000000000 MUL DUP2 MSTORE PUSH1 0x4 ADD PUSH1 0x0 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 PUSH1 0x0 DUP8 DUP1 EXTCODESIZE ISZERO DUP1 ISZERO PUSH2 0x6DA JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP GAS CALL ISZERO DUP1 ISZERO PUSH2 0x6EE JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP POP JUMP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x751 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x1 PUSH1 0x0 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH1 0xFF AND ISZERO ISZERO PUSH2 0x863 JUMPI PUSH1 0x1 DUP1 PUSH1 0x0 DUP5 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH1 0xFF MUL NOT AND SWAP1 DUP4 ISZERO ISZERO MUL OR SWAP1 SSTORE POP PUSH32 0xD1BBA68C128CC3F427E5831B3C6F99F480B6EFA6B9E80C757768F6124158CC3F DUP3 PUSH1 0x40 MLOAD DUP1 DUP3 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 LOG1 PUSH1 0x1 SWAP1 POP JUMPDEST SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x20 MSTORE DUP1 PUSH1 0x0 MSTORE PUSH1 0x40 PUSH1 0x0 KECCAK256 PUSH1 0x0 SWAP2 POP SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH1 0xFF AND DUP2 JUMP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x90B JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 SWAP1 POP JUMPDEST DUP3 MLOAD DUP2 LT ISZERO PUSH2 0x950 JUMPI PUSH2 0x939 DUP4 DUP3 DUP2 MLOAD DUP2 LT ISZERO ISZERO PUSH2 0x92A JUMPI INVALID JUMPDEST SWAP1 PUSH1 0x20 ADD SWAP1 PUSH1 0x20 MUL ADD MLOAD PUSH2 0x6F4 JUMP JUMPDEST ISZERO PUSH2 0x943 JUMPI PUSH1 0x1 SWAP2 POP JUMPDEST DUP1 DUP1 PUSH1 0x1 ADD SWAP2 POP POP PUSH2 0x910 JUMP JUMPDEST POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x9B1 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO ISZERO PUSH2 0x9ED JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP1 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH32 0x8BE0079C531659141344CD1FD0A4F28419497F9722A3DAAFE3B4186F6B6457E0 PUSH1 0x40 MLOAD PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 LOG3 DUP1 PUSH1 0x0 DUP1 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP POP JUMP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x1 PUSH1 0x0 CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH1 0xFF AND ISZERO ISZERO PUSH2 0xB06 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 DUP7 DUP7 SWAP1 POP GT ISZERO ISZERO PUSH2 0xB81 JUMPI PUSH1 0x40 MLOAD PUSH32 0x8C379A000000000000000000000000000000000000000000000000000000000 DUP2 MSTORE PUSH1 0x4 ADD DUP1 DUP1 PUSH1 0x20 ADD DUP3 DUP2 SUB DUP3 MSTORE PUSH1 0x18 DUP2 MSTORE PUSH1 0x20 ADD DUP1 PUSH32 0x726563656976657273206973206F66206C656E67746820300000000000000000 DUP2 MSTORE POP PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 REVERT JUMPDEST DUP4 DUP4 SWAP1 POP DUP7 DUP7 SWAP1 POP EQ ISZERO ISZERO PUSH2 0xC24 JUMPI PUSH1 0x40 MLOAD PUSH32 0x8C379A000000000000000000000000000000000000000000000000000000000 DUP2 MSTORE PUSH1 0x4 ADD DUP1 DUP1 PUSH1 0x20 ADD DUP3 DUP2 SUB DUP3 MSTORE PUSH1 0x2E DUP2 MSTORE PUSH1 0x20 ADD DUP1 PUSH32 0x746865206C656E67746873206F662072656365697665727320616E6420616D6F DUP2 MSTORE PUSH1 0x20 ADD PUSH32 0x756E7473206E6F7420657175616C000000000000000000000000000000000000 DUP2 MSTORE POP PUSH1 0x40 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 REVERT JUMPDEST PUSH1 0x2 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH4 0x70A08231 ADDRESS PUSH1 0x40 MLOAD DUP3 PUSH4 0xFFFFFFFF AND PUSH29 0x100000000000000000000000000000000000000000000000000000000 MUL DUP2 MSTORE PUSH1 0x4 ADD DUP1 DUP3 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x20 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 PUSH1 0x0 DUP8 DUP1 EXTCODESIZE ISZERO DUP1 ISZERO PUSH2 0xCE1 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP GAS CALL ISZERO DUP1 ISZERO PUSH2 0xCF5 JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP POP PUSH1 0x40 MLOAD RETURNDATASIZE PUSH1 0x20 DUP2 LT ISZERO PUSH2 0xD0B JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP2 ADD SWAP1 DUP1 DUP1 MLOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP SWAP2 POP PUSH1 0x0 SWAP1 POP JUMPDEST DUP6 DUP6 SWAP1 POP DUP2 LT ISZERO PUSH2 0xF1F JUMPI DUP4 DUP4 DUP3 DUP2 DUP2 LT ISZERO ISZERO PUSH2 0xD3C JUMPI INVALID JUMPDEST SWAP1 POP PUSH1 0x20 MUL ADD CALLDATALOAD DUP3 LT ISZERO ISZERO ISZERO PUSH2 0xDBA JUMPI PUSH1 0x40 MLOAD PUSH32 0x8C379A000000000000000000000000000000000000000000000000000000000 DUP2 MSTORE PUSH1 0x4 ADD DUP1 DUP1 PUSH1 0x20 ADD DUP3 DUP2 SUB DUP3 MSTORE PUSH1 0x14 DUP2 MSTORE PUSH1 0x20 ADD DUP1 PUSH32 0x696E73756666696369656E742062616C616E6365000000000000000000000000 DUP2 MSTORE POP PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 REVERT JUMPDEST DUP4 DUP4 DUP3 DUP2 DUP2 LT ISZERO ISZERO PUSH2 0xDC8 JUMPI INVALID JUMPDEST SWAP1 POP PUSH1 0x20 MUL ADD CALLDATALOAD DUP3 SUB SWAP2 POP PUSH1 0x2 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH4 0xA9059CBB DUP8 DUP8 DUP5 DUP2 DUP2 LT ISZERO ISZERO PUSH2 0xE1F JUMPI INVALID JUMPDEST SWAP1 POP PUSH1 0x20 MUL ADD CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP7 DUP7 DUP6 DUP2 DUP2 LT ISZERO ISZERO PUSH2 0xE4A JUMPI INVALID JUMPDEST SWAP1 POP PUSH1 0x20 MUL ADD CALLDATALOAD PUSH1 0x40 MLOAD DUP4 PUSH4 0xFFFFFFFF AND PUSH29 0x100000000000000000000000000000000000000000000000000000000 MUL DUP2 MSTORE PUSH1 0x4 ADD DUP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD DUP3 DUP2 MSTORE PUSH1 0x20 ADD SWAP3 POP POP POP PUSH1 0x20 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 PUSH1 0x0 DUP8 DUP1 EXTCODESIZE ISZERO DUP1 ISZERO PUSH2 0xED6 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP GAS CALL ISZERO DUP1 ISZERO PUSH2 0xEEA JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP POP PUSH1 0x40 MLOAD RETURNDATASIZE PUSH1 0x20 DUP2 LT ISZERO PUSH2 0xF00 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP2 ADD SWAP1 DUP1 DUP1 MLOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP POP DUP1 DUP1 PUSH1 0x1 ADD SWAP2 POP POP PUSH2 0xD23 JUMP JUMPDEST POP POP POP POP POP POP JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 PUSH13 0x7B37483D0B320EE41D5C1BEF74 0xa5 PUSH19 0xF611ABDE0A8BF31852B9D3DE8B0A14DF002900 ",
	"sourceMap": "123:786:1:-;;;193:100;8:9:-1;5:2;;;30:1;27;20:12;5:2;193:100:1;;;;;;;;;;;;;;;;;;;;;;;;;;;;;670:10:4;662:5;;:18;;;;;;;;;;;;;;;;;;269:16:1;248:4;;:38;;;;;;;;;;;;;;;;;;193:100;123:786;;;;;;"
}`

// DeployClerk deploys a new Ethereum contract, binding an instance of Clerk to it.
func DeployClerk(auth *bind.TransactOpts, backend bind.ContractBackend, vitaTokenAddress common.Address) (common.Address, *types.Transaction, *Clerk, error) {
	parsed, err := abi.JSON(strings.NewReader(ClerkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClerkBin), backend, vitaTokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Clerk{ClerkCaller: ClerkCaller{contract: contract}, ClerkTransactor: ClerkTransactor{contract: contract}, ClerkFilterer: ClerkFilterer{contract: contract}}, nil
}

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
