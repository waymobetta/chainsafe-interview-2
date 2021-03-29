// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"}],\"name\":\"addCID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cidArray\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCIDArray\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StoreFuncSigs maps the 4-byte function signature to its string representation.
var StoreFuncSigs = map[string]string{
	"fa0f1346": "addCID(string)",
	"b02f5dd4": "cidArray(uint256)",
	"8247b8ab": "getCIDArray()",
	"8da5cb5b": "owner()",
}

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610559806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80638247b8ab146100515780638da5cb5b1461006f578063b02f5dd41461009a578063fa0f1346146100ba575b600080fd5b6100596100cf565b6040516100669190610457565b60405180910390f35b600054610082906001600160a01b031681565b6040516001600160a01b039091168152602001610066565b6100ad6100a83660046103f4565b6101a8565b60405161006691906104b8565b6100cd6100c836600461034a565b610254565b005b60606001805480602002602001604051908101604052809291908181526020016000905b8282101561019f578382906000526020600020018054610112906104d2565b80601f016020809104026020016040519081016040528092919081815260200182805461013e906104d2565b801561018b5780601f106101605761010080835404028352916020019161018b565b820191906000526020600020905b81548152906001019060200180831161016e57829003601f168201915b5050505050815260200190600101906100f3565b50505050905090565b600181815481106101b857600080fd5b9060005260206000200160009150905080546101d3906104d2565b80601f01602080910402602001604051908101604052809291908181526020018280546101ff906104d2565b801561024c5780601f106102215761010080835404028352916020019161024c565b820191906000526020600020905b81548152906001019060200180831161022f57829003601f168201915b505050505081565b6000546001600160a01b0316331461026b57600080fd5b60018054808201825560009190915281516102ad917fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6019060208401906102b1565b5050565b8280546102bd906104d2565b90600052602060002090601f0160209004810192826102df5760008555610325565b82601f106102f857805160ff1916838001178555610325565b82800160010185558215610325579182015b8281111561032557825182559160200191906001019061030a565b50610331929150610335565b5090565b5b808211156103315760008155600101610336565b60006020828403121561035b578081fd5b813567ffffffffffffffff80821115610372578283fd5b818401915084601f830112610385578283fd5b8135818111156103975761039761050d565b604051601f8201601f19908116603f011681019083821181831017156103bf576103bf61050d565b816040528281528760208487010111156103d7578586fd5b826020860160208301379182016020019490945295945050505050565b600060208284031215610405578081fd5b5035919050565b60008151808452815b8181101561043157602081850181015186830182015201610415565b818111156104425782602083870101525b50601f01601f19169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b8701019250838701855b828110156104ab57603f1988860301845261049985835161040c565b9450928501929085019060010161047d565b5092979650505050505050565b6000602082526104cb602083018461040c565b9392505050565b600181811c908216806104e657607f821691505b6020821081141561050757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea264697066735822122006950e9aa664a23ed3c59c9b3bcf0e3d0372083af47a115a6147b50af9f6ab5e64736f6c63430008030033"

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// CidArray is a free data retrieval call binding the contract method 0xb02f5dd4.
//
// Solidity: function cidArray(uint256 ) view returns(string)
func (_Store *StoreCaller) CidArray(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "cidArray", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CidArray is a free data retrieval call binding the contract method 0xb02f5dd4.
//
// Solidity: function cidArray(uint256 ) view returns(string)
func (_Store *StoreSession) CidArray(arg0 *big.Int) (string, error) {
	return _Store.Contract.CidArray(&_Store.CallOpts, arg0)
}

// CidArray is a free data retrieval call binding the contract method 0xb02f5dd4.
//
// Solidity: function cidArray(uint256 ) view returns(string)
func (_Store *StoreCallerSession) CidArray(arg0 *big.Int) (string, error) {
	return _Store.Contract.CidArray(&_Store.CallOpts, arg0)
}

// GetCIDArray is a free data retrieval call binding the contract method 0x8247b8ab.
//
// Solidity: function getCIDArray() view returns(string[])
func (_Store *StoreCaller) GetCIDArray(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getCIDArray")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetCIDArray is a free data retrieval call binding the contract method 0x8247b8ab.
//
// Solidity: function getCIDArray() view returns(string[])
func (_Store *StoreSession) GetCIDArray() ([]string, error) {
	return _Store.Contract.GetCIDArray(&_Store.CallOpts)
}

// GetCIDArray is a free data retrieval call binding the contract method 0x8247b8ab.
//
// Solidity: function getCIDArray() view returns(string[])
func (_Store *StoreCallerSession) GetCIDArray() ([]string, error) {
	return _Store.Contract.GetCIDArray(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// AddCID is a paid mutator transaction binding the contract method 0xfa0f1346.
//
// Solidity: function addCID(string _cid) returns()
func (_Store *StoreTransactor) AddCID(opts *bind.TransactOpts, _cid string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addCID", _cid)
}

// AddCID is a paid mutator transaction binding the contract method 0xfa0f1346.
//
// Solidity: function addCID(string _cid) returns()
func (_Store *StoreSession) AddCID(_cid string) (*types.Transaction, error) {
	return _Store.Contract.AddCID(&_Store.TransactOpts, _cid)
}

// AddCID is a paid mutator transaction binding the contract method 0xfa0f1346.
//
// Solidity: function addCID(string _cid) returns()
func (_Store *StoreTransactorSession) AddCID(_cid string) (*types.Transaction, error) {
	return _Store.Contract.AddCID(&_Store.TransactOpts, _cid)
}
