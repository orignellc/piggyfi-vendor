// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package celo

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/event"
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
)

// WhitelistMetaData contains all meta data concerning the Whitelist contract.
var WhitelistMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"maxAllowanceToMint\",\"type\":\"uint64\"}],\"name\":\"checkInWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"updateMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516105e23803806105e28339818101604052810190610032919061007a565b80600081905550506100a7565b600080fd5b6000819050919050565b61005781610044565b811461006257600080fd5b50565b6000815190506100748161004e565b92915050565b6000602082840312156100905761008f61003f565b5b600061009e84828501610065565b91505092915050565b61052c806100b66000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632eb4a7ab146100465780634783f0ef14610064578063e10d3be214610080575b600080fd5b61004e6100b0565b60405161005b9190610214565b60405180910390f35b61007e60048036038101906100799190610265565b6100b6565b005b61009a60048036038101906100959190610337565b6100c0565b6040516100a791906103b2565b60405180910390f35b60005481565b8060008190555050565b60008033836040516020016100d692919061041d565b604051602081830303815290604052805190602001209050600061013e868680806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050506000548461014c565b905080925050509392505050565b6000826101598584610163565b1490509392505050565b60008082905060005b84518110156101ae576101998286838151811061018c5761018b610446565b5b60200260200101516101b9565b915080806101a6906104ae565b91505061016c565b508091505092915050565b60008183106101d1576101cc82846101e4565b6101dc565b6101db83836101e4565b5b905092915050565b600082600052816020526040600020905092915050565b6000819050919050565b61020e816101fb565b82525050565b60006020820190506102296000830184610205565b92915050565b600080fd5b600080fd5b610242816101fb565b811461024d57600080fd5b50565b60008135905061025f81610239565b92915050565b60006020828403121561027b5761027a61022f565b5b600061028984828501610250565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126102b7576102b6610292565b5b8235905067ffffffffffffffff8111156102d4576102d3610297565b5b6020830191508360208202830111156102f0576102ef61029c565b5b9250929050565b600067ffffffffffffffff82169050919050565b610314816102f7565b811461031f57600080fd5b50565b6000813590506103318161030b565b92915050565b6000806000604084860312156103505761034f61022f565b5b600084013567ffffffffffffffff81111561036e5761036d610234565b5b61037a868287016102a1565b9350935050602061038d86828701610322565b9150509250925092565b60008115159050919050565b6103ac81610397565b82525050565b60006020820190506103c760008301846103a3565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103f8826103cd565b9050919050565b610408816103ed565b82525050565b610417816102f7565b82525050565b600060408201905061043260008301856103ff565b61043f602083018461040e565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b60006104b9826104a4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036104eb576104ea610475565b5b60018201905091905056fea2646970667358221220fc8efbc478a709fd6e9ccea2de90c829dd034f60c88cddd1ff8c5dffcbecf8fd64736f6c634300080f0033",
}

// WhitelistABI is the input ABI used to generate the binding from.
// Deprecated: Use WhitelistMetaData.ABI instead.
var WhitelistABI = WhitelistMetaData.ABI

// WhitelistBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WhitelistMetaData.Bin instead.
var WhitelistBin = WhitelistMetaData.Bin

// DeployWhitelist deploys a new Ethereum contract, binding an instance of Whitelist to it.
func DeployWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend, _merkleRoot [32]byte) (common.Address, *types.Transaction, *Whitelist, error) {
	parsed, err := WhitelistMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WhitelistBin), backend, _merkleRoot)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// Whitelist is an auto generated Go binding around an Ethereum contract.
type Whitelist struct {
	WhitelistCaller     // Read-only binding to the contract
	WhitelistTransactor // Write-only binding to the contract
	WhitelistFilterer   // Log filterer for contract events
}

// WhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistSession struct {
	Contract     *Whitelist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistCallerSession struct {
	Contract *WhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistTransactorSession struct {
	Contract     *WhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistRaw struct {
	Contract *Whitelist // Generic contract binding to access the raw methods on
}

// WhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistCallerRaw struct {
	Contract *WhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistTransactorRaw struct {
	Contract *WhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelist creates a new instance of Whitelist, bound to a specific deployed contract.
func NewWhitelist(address common.Address, backend bind.ContractBackend) (*Whitelist, error) {
	contract, err := bindWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// NewWhitelistCaller creates a new read-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistCaller(address common.Address, caller bind.ContractCaller) (*WhitelistCaller, error) {
	contract, err := bindWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistCaller{contract: contract}, nil
}

// NewWhitelistTransactor creates a new write-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistTransactor, error) {
	contract, err := bindWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistTransactor{contract: contract}, nil
}

// NewWhitelistFilterer creates a new log filterer instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistFilterer, error) {
	contract, err := bindWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistFilterer{contract: contract}, nil
}

// bindWhitelist binds a generic wrapper to an already deployed contract.
func bindWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseWhitelistABI parses the ABI
func ParseWhitelistABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.WhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transact(opts, method, params...)
}

// CheckInWhitelist is a free data retrieval call binding the contract method 0xe10d3be2.
//
// Solidity: function checkInWhitelist(bytes32[] proof, uint64 maxAllowanceToMint) view returns(bool)
func (_Whitelist *WhitelistCaller) CheckInWhitelist(opts *bind.CallOpts, proof [][32]byte, maxAllowanceToMint uint64) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "checkInWhitelist", proof, maxAllowanceToMint)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckInWhitelist is a free data retrieval call binding the contract method 0xe10d3be2.
//
// Solidity: function checkInWhitelist(bytes32[] proof, uint64 maxAllowanceToMint) view returns(bool)
func (_Whitelist *WhitelistSession) CheckInWhitelist(proof [][32]byte, maxAllowanceToMint uint64) (bool, error) {
	return _Whitelist.Contract.CheckInWhitelist(&_Whitelist.CallOpts, proof, maxAllowanceToMint)
}

// CheckInWhitelist is a free data retrieval call binding the contract method 0xe10d3be2.
//
// Solidity: function checkInWhitelist(bytes32[] proof, uint64 maxAllowanceToMint) view returns(bool)
func (_Whitelist *WhitelistCallerSession) CheckInWhitelist(proof [][32]byte, maxAllowanceToMint uint64) (bool, error) {
	return _Whitelist.Contract.CheckInWhitelist(&_Whitelist.CallOpts, proof, maxAllowanceToMint)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Whitelist *WhitelistCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Whitelist *WhitelistSession) MerkleRoot() ([32]byte, error) {
	return _Whitelist.Contract.MerkleRoot(&_Whitelist.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Whitelist *WhitelistCallerSession) MerkleRoot() ([32]byte, error) {
	return _Whitelist.Contract.MerkleRoot(&_Whitelist.CallOpts)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 root) returns()
func (_Whitelist *WhitelistTransactor) UpdateMerkleRoot(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "updateMerkleRoot", root)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 root) returns()
func (_Whitelist *WhitelistSession) UpdateMerkleRoot(root [32]byte) (*types.Transaction, error) {
	return _Whitelist.Contract.UpdateMerkleRoot(&_Whitelist.TransactOpts, root)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 root) returns()
func (_Whitelist *WhitelistTransactorSession) UpdateMerkleRoot(root [32]byte) (*types.Transaction, error) {
	return _Whitelist.Contract.UpdateMerkleRoot(&_Whitelist.TransactOpts, root)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Whitelist *WhitelistFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Whitelist.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}
