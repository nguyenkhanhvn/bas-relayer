// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// TypesTokenMetadata is an auto generated low-level Go binding around an user-defined struct.
type TypesTokenMetadata struct {
	Name          string
	Symbol        string
	OriginChain   *big.Int
	OriginAddress common.Address
}

// CrossChainBridgeMetaData contains all meta data concerning the CrossChainBridge contract.
var CrossChainBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRelayHub\",\"name\":\"relayHub\",\"type\":\"address\"},{\"internalType\":\"contractBridgeRouter\",\"name\":\"bridgeRouter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nativeTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nativeTokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"BridgeRouterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"originChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structTypes.TokenMetadata\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"DepositToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"originChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structTypes.TokenMetadata\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"WithdrawToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NATIVE_TOKEN_METADATA\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"originChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBridgeRouter\",\"name\":\"otherRouter\",\"type\":\"address\"}],\"name\":\"changeRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"}],\"name\":\"getPeggedTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayHub\",\"outputs\":[{\"internalType\":\"contractIRelayHub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status_\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"blockProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"rawReceipt\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofPath\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofSiblings\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CrossChainBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossChainBridgeMetaData.ABI instead.
var CrossChainBridgeABI = CrossChainBridgeMetaData.ABI

// CrossChainBridge is an auto generated Go binding around an Ethereum contract.
type CrossChainBridge struct {
	CrossChainBridgeCaller     // Read-only binding to the contract
	CrossChainBridgeTransactor // Write-only binding to the contract
	CrossChainBridgeFilterer   // Log filterer for contract events
}

// CrossChainBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossChainBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossChainBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossChainBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossChainBridgeSession struct {
	Contract     *CrossChainBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossChainBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossChainBridgeCallerSession struct {
	Contract *CrossChainBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CrossChainBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossChainBridgeTransactorSession struct {
	Contract     *CrossChainBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CrossChainBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossChainBridgeRaw struct {
	Contract *CrossChainBridge // Generic contract binding to access the raw methods on
}

// CrossChainBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossChainBridgeCallerRaw struct {
	Contract *CrossChainBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// CrossChainBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossChainBridgeTransactorRaw struct {
	Contract *CrossChainBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossChainBridge creates a new instance of CrossChainBridge, bound to a specific deployed contract.
func NewCrossChainBridge(address common.Address, backend bind.ContractBackend) (*CrossChainBridge, error) {
	contract, err := bindCrossChainBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChainBridge{CrossChainBridgeCaller: CrossChainBridgeCaller{contract: contract}, CrossChainBridgeTransactor: CrossChainBridgeTransactor{contract: contract}, CrossChainBridgeFilterer: CrossChainBridgeFilterer{contract: contract}}, nil
}

// NewCrossChainBridgeCaller creates a new read-only instance of CrossChainBridge, bound to a specific deployed contract.
func NewCrossChainBridgeCaller(address common.Address, caller bind.ContractCaller) (*CrossChainBridgeCaller, error) {
	contract, err := bindCrossChainBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainBridgeCaller{contract: contract}, nil
}

// NewCrossChainBridgeTransactor creates a new write-only instance of CrossChainBridge, bound to a specific deployed contract.
func NewCrossChainBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainBridgeTransactor, error) {
	contract, err := bindCrossChainBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainBridgeTransactor{contract: contract}, nil
}

// NewCrossChainBridgeFilterer creates a new log filterer instance of CrossChainBridge, bound to a specific deployed contract.
func NewCrossChainBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainBridgeFilterer, error) {
	contract, err := bindCrossChainBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainBridgeFilterer{contract: contract}, nil
}

// bindCrossChainBridge binds a generic wrapper to an already deployed contract.
func bindCrossChainBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossChainBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainBridge *CrossChainBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainBridge.Contract.CrossChainBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainBridge *CrossChainBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.CrossChainBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainBridge *CrossChainBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.CrossChainBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainBridge *CrossChainBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainBridge *CrossChainBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainBridge *CrossChainBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.contract.Transact(opts, method, params...)
}

// NATIVETOKENMETADATA is a free data retrieval call binding the contract method 0xdcdc12c5.
//
// Solidity: function NATIVE_TOKEN_METADATA() view returns(string name, string symbol, uint256 originChain, address originAddress)
func (_CrossChainBridge *CrossChainBridgeCaller) NATIVETOKENMETADATA(opts *bind.CallOpts) (struct {
	Name          string
	Symbol        string
	OriginChain   *big.Int
	OriginAddress common.Address
}, error) {
	var out []interface{}
	err := _CrossChainBridge.contract.Call(opts, &out, "NATIVE_TOKEN_METADATA")

	outstruct := new(struct {
		Name          string
		Symbol        string
		OriginChain   *big.Int
		OriginAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.OriginChain = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OriginAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// NATIVETOKENMETADATA is a free data retrieval call binding the contract method 0xdcdc12c5.
//
// Solidity: function NATIVE_TOKEN_METADATA() view returns(string name, string symbol, uint256 originChain, address originAddress)
func (_CrossChainBridge *CrossChainBridgeSession) NATIVETOKENMETADATA() (struct {
	Name          string
	Symbol        string
	OriginChain   *big.Int
	OriginAddress common.Address
}, error) {
	return _CrossChainBridge.Contract.NATIVETOKENMETADATA(&_CrossChainBridge.CallOpts)
}

// NATIVETOKENMETADATA is a free data retrieval call binding the contract method 0xdcdc12c5.
//
// Solidity: function NATIVE_TOKEN_METADATA() view returns(string name, string symbol, uint256 originChain, address originAddress)
func (_CrossChainBridge *CrossChainBridgeCallerSession) NATIVETOKENMETADATA() (struct {
	Name          string
	Symbol        string
	OriginChain   *big.Int
	OriginAddress common.Address
}, error) {
	return _CrossChainBridge.Contract.NATIVETOKENMETADATA(&_CrossChainBridge.CallOpts)
}

// GetPeggedTokenAddress is a free data retrieval call binding the contract method 0x53cc8084.
//
// Solidity: function getPeggedTokenAddress(address fromToken) view returns(address)
func (_CrossChainBridge *CrossChainBridgeCaller) GetPeggedTokenAddress(opts *bind.CallOpts, fromToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _CrossChainBridge.contract.Call(opts, &out, "getPeggedTokenAddress", fromToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPeggedTokenAddress is a free data retrieval call binding the contract method 0x53cc8084.
//
// Solidity: function getPeggedTokenAddress(address fromToken) view returns(address)
func (_CrossChainBridge *CrossChainBridgeSession) GetPeggedTokenAddress(fromToken common.Address) (common.Address, error) {
	return _CrossChainBridge.Contract.GetPeggedTokenAddress(&_CrossChainBridge.CallOpts, fromToken)
}

// GetPeggedTokenAddress is a free data retrieval call binding the contract method 0x53cc8084.
//
// Solidity: function getPeggedTokenAddress(address fromToken) view returns(address)
func (_CrossChainBridge *CrossChainBridgeCallerSession) GetPeggedTokenAddress(fromToken common.Address) (common.Address, error) {
	return _CrossChainBridge.Contract.GetPeggedTokenAddress(&_CrossChainBridge.CallOpts, fromToken)
}

// GetRelayHub is a free data retrieval call binding the contract method 0x7bdf2ec7.
//
// Solidity: function getRelayHub() view returns(address)
func (_CrossChainBridge *CrossChainBridgeCaller) GetRelayHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainBridge.contract.Call(opts, &out, "getRelayHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayHub is a free data retrieval call binding the contract method 0x7bdf2ec7.
//
// Solidity: function getRelayHub() view returns(address)
func (_CrossChainBridge *CrossChainBridgeSession) GetRelayHub() (common.Address, error) {
	return _CrossChainBridge.Contract.GetRelayHub(&_CrossChainBridge.CallOpts)
}

// GetRelayHub is a free data retrieval call binding the contract method 0x7bdf2ec7.
//
// Solidity: function getRelayHub() view returns(address)
func (_CrossChainBridge *CrossChainBridgeCallerSession) GetRelayHub() (common.Address, error) {
	return _CrossChainBridge.Contract.GetRelayHub(&_CrossChainBridge.CallOpts)
}

// GetTokenImplementation is a free data retrieval call binding the contract method 0x709bc7f3.
//
// Solidity: function getTokenImplementation() view returns(address)
func (_CrossChainBridge *CrossChainBridgeCaller) GetTokenImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainBridge.contract.Call(opts, &out, "getTokenImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenImplementation is a free data retrieval call binding the contract method 0x709bc7f3.
//
// Solidity: function getTokenImplementation() view returns(address)
func (_CrossChainBridge *CrossChainBridgeSession) GetTokenImplementation() (common.Address, error) {
	return _CrossChainBridge.Contract.GetTokenImplementation(&_CrossChainBridge.CallOpts)
}

// GetTokenImplementation is a free data retrieval call binding the contract method 0x709bc7f3.
//
// Solidity: function getTokenImplementation() view returns(address)
func (_CrossChainBridge *CrossChainBridgeCallerSession) GetTokenImplementation() (common.Address, error) {
	return _CrossChainBridge.Contract.GetTokenImplementation(&_CrossChainBridge.CallOpts)
}

// ChangeRouter is a paid mutator transaction binding the contract method 0x340ac20f.
//
// Solidity: function changeRouter(address otherRouter) returns()
func (_CrossChainBridge *CrossChainBridgeTransactor) ChangeRouter(opts *bind.TransactOpts, otherRouter common.Address) (*types.Transaction, error) {
	return _CrossChainBridge.contract.Transact(opts, "changeRouter", otherRouter)
}

// ChangeRouter is a paid mutator transaction binding the contract method 0x340ac20f.
//
// Solidity: function changeRouter(address otherRouter) returns()
func (_CrossChainBridge *CrossChainBridgeSession) ChangeRouter(otherRouter common.Address) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.ChangeRouter(&_CrossChainBridge.TransactOpts, otherRouter)
}

// ChangeRouter is a paid mutator transaction binding the contract method 0x340ac20f.
//
// Solidity: function changeRouter(address otherRouter) returns()
func (_CrossChainBridge *CrossChainBridgeTransactorSession) ChangeRouter(otherRouter common.Address) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.ChangeRouter(&_CrossChainBridge.TransactOpts, otherRouter)
}

// DepositNative is a paid mutator transaction binding the contract method 0xf0194945.
//
// Solidity: function depositNative(uint256 toChain, address toAddress) payable returns()
func (_CrossChainBridge *CrossChainBridgeTransactor) DepositNative(opts *bind.TransactOpts, toChain *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _CrossChainBridge.contract.Transact(opts, "depositNative", toChain, toAddress)
}

// DepositNative is a paid mutator transaction binding the contract method 0xf0194945.
//
// Solidity: function depositNative(uint256 toChain, address toAddress) payable returns()
func (_CrossChainBridge *CrossChainBridgeSession) DepositNative(toChain *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.DepositNative(&_CrossChainBridge.TransactOpts, toChain, toAddress)
}

// DepositNative is a paid mutator transaction binding the contract method 0xf0194945.
//
// Solidity: function depositNative(uint256 toChain, address toAddress) payable returns()
func (_CrossChainBridge *CrossChainBridgeTransactorSession) DepositNative(toChain *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.DepositNative(&_CrossChainBridge.TransactOpts, toChain, toAddress)
}

// DepositToken is a paid mutator transaction binding the contract method 0x6701ccb5.
//
// Solidity: function depositToken(address fromToken, uint256 toChain, address toAddress, uint256 totalAmount) returns()
func (_CrossChainBridge *CrossChainBridgeTransactor) DepositToken(opts *bind.TransactOpts, fromToken common.Address, toChain *big.Int, toAddress common.Address, totalAmount *big.Int) (*types.Transaction, error) {
	return _CrossChainBridge.contract.Transact(opts, "depositToken", fromToken, toChain, toAddress, totalAmount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x6701ccb5.
//
// Solidity: function depositToken(address fromToken, uint256 toChain, address toAddress, uint256 totalAmount) returns()
func (_CrossChainBridge *CrossChainBridgeSession) DepositToken(fromToken common.Address, toChain *big.Int, toAddress common.Address, totalAmount *big.Int) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.DepositToken(&_CrossChainBridge.TransactOpts, fromToken, toChain, toAddress, totalAmount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x6701ccb5.
//
// Solidity: function depositToken(address fromToken, uint256 toChain, address toAddress, uint256 totalAmount) returns()
func (_CrossChainBridge *CrossChainBridgeTransactorSession) DepositToken(fromToken common.Address, toChain *big.Int, toAddress common.Address, totalAmount *big.Int) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.DepositToken(&_CrossChainBridge.TransactOpts, fromToken, toChain, toAddress, totalAmount)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_CrossChainBridge *CrossChainBridgeTransactor) SetOperator(opts *bind.TransactOpts, operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _CrossChainBridge.contract.Transact(opts, "setOperator", operator_, status_)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_CrossChainBridge *CrossChainBridgeSession) SetOperator(operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.SetOperator(&_CrossChainBridge.TransactOpts, operator_, status_)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_CrossChainBridge *CrossChainBridgeTransactorSession) SetOperator(operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.SetOperator(&_CrossChainBridge.TransactOpts, operator_, status_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x662c7ba5.
//
// Solidity: function withdraw(bytes[] blockProofs, bytes rawReceipt, bytes proofPath, bytes proofSiblings) returns()
func (_CrossChainBridge *CrossChainBridgeTransactor) Withdraw(opts *bind.TransactOpts, blockProofs [][]byte, rawReceipt []byte, proofPath []byte, proofSiblings []byte) (*types.Transaction, error) {
	return _CrossChainBridge.contract.Transact(opts, "withdraw", blockProofs, rawReceipt, proofPath, proofSiblings)
}

// Withdraw is a paid mutator transaction binding the contract method 0x662c7ba5.
//
// Solidity: function withdraw(bytes[] blockProofs, bytes rawReceipt, bytes proofPath, bytes proofSiblings) returns()
func (_CrossChainBridge *CrossChainBridgeSession) Withdraw(blockProofs [][]byte, rawReceipt []byte, proofPath []byte, proofSiblings []byte) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.Withdraw(&_CrossChainBridge.TransactOpts, blockProofs, rawReceipt, proofPath, proofSiblings)
}

// Withdraw is a paid mutator transaction binding the contract method 0x662c7ba5.
//
// Solidity: function withdraw(bytes[] blockProofs, bytes rawReceipt, bytes proofPath, bytes proofSiblings) returns()
func (_CrossChainBridge *CrossChainBridgeTransactorSession) Withdraw(blockProofs [][]byte, rawReceipt []byte, proofPath []byte, proofSiblings []byte) (*types.Transaction, error) {
	return _CrossChainBridge.Contract.Withdraw(&_CrossChainBridge.TransactOpts, blockProofs, rawReceipt, proofPath, proofSiblings)
}

// CrossChainBridgeBridgeRouterChangedIterator is returned from FilterBridgeRouterChanged and is used to iterate over the raw logs and unpacked data for BridgeRouterChanged events raised by the CrossChainBridge contract.
type CrossChainBridgeBridgeRouterChangedIterator struct {
	Event *CrossChainBridgeBridgeRouterChanged // Event containing the contract specifics and raw log

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
func (it *CrossChainBridgeBridgeRouterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainBridgeBridgeRouterChanged)
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
		it.Event = new(CrossChainBridgeBridgeRouterChanged)
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
func (it *CrossChainBridgeBridgeRouterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainBridgeBridgeRouterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainBridgeBridgeRouterChanged represents a BridgeRouterChanged event raised by the CrossChainBridge contract.
type CrossChainBridgeBridgeRouterChanged struct {
	OldValue common.Address
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeRouterChanged is a free log retrieval operation binding the contract event 0x3e9aaaa36d36b050584098ca9083e20442f6c87221948cb8a6d79c9ef82c5165.
//
// Solidity: event BridgeRouterChanged(address oldValue, address newValue)
func (_CrossChainBridge *CrossChainBridgeFilterer) FilterBridgeRouterChanged(opts *bind.FilterOpts) (*CrossChainBridgeBridgeRouterChangedIterator, error) {

	logs, sub, err := _CrossChainBridge.contract.FilterLogs(opts, "BridgeRouterChanged")
	if err != nil {
		return nil, err
	}
	return &CrossChainBridgeBridgeRouterChangedIterator{contract: _CrossChainBridge.contract, event: "BridgeRouterChanged", logs: logs, sub: sub}, nil
}

// WatchBridgeRouterChanged is a free log subscription operation binding the contract event 0x3e9aaaa36d36b050584098ca9083e20442f6c87221948cb8a6d79c9ef82c5165.
//
// Solidity: event BridgeRouterChanged(address oldValue, address newValue)
func (_CrossChainBridge *CrossChainBridgeFilterer) WatchBridgeRouterChanged(opts *bind.WatchOpts, sink chan<- *CrossChainBridgeBridgeRouterChanged) (event.Subscription, error) {

	logs, sub, err := _CrossChainBridge.contract.WatchLogs(opts, "BridgeRouterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainBridgeBridgeRouterChanged)
				if err := _CrossChainBridge.contract.UnpackLog(event, "BridgeRouterChanged", log); err != nil {
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

// ParseBridgeRouterChanged is a log parse operation binding the contract event 0x3e9aaaa36d36b050584098ca9083e20442f6c87221948cb8a6d79c9ef82c5165.
//
// Solidity: event BridgeRouterChanged(address oldValue, address newValue)
func (_CrossChainBridge *CrossChainBridgeFilterer) ParseBridgeRouterChanged(log types.Log) (*CrossChainBridgeBridgeRouterChanged, error) {
	event := new(CrossChainBridgeBridgeRouterChanged)
	if err := _CrossChainBridge.contract.UnpackLog(event, "BridgeRouterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainBridgeDepositTokenIterator is returned from FilterDepositToken and is used to iterate over the raw logs and unpacked data for DepositToken events raised by the CrossChainBridge contract.
type CrossChainBridgeDepositTokenIterator struct {
	Event *CrossChainBridgeDepositToken // Event containing the contract specifics and raw log

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
func (it *CrossChainBridgeDepositTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainBridgeDepositToken)
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
		it.Event = new(CrossChainBridgeDepositToken)
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
func (it *CrossChainBridgeDepositTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainBridgeDepositTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainBridgeDepositToken represents a DepositToken event raised by the CrossChainBridge contract.
type CrossChainBridgeDepositToken struct {
	FromChain   *big.Int
	ToChain     *big.Int
	FromAddress common.Address
	ToAddress   common.Address
	FromToken   common.Address
	ToToken     common.Address
	TotalAmount *big.Int
	Arg7        TypesTokenMetadata
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositToken is a free log retrieval operation binding the contract event 0x52e2a71d8eeeaf64c3fa0a51a00f5cd6c4d0e7847d213c8d68d823a4b4288d7e.
//
// Solidity: event DepositToken(uint256 fromChain, uint256 indexed toChain, address indexed fromAddress, address indexed toAddress, address fromToken, address toToken, uint256 totalAmount, (string,string,uint256,address) arg7)
func (_CrossChainBridge *CrossChainBridgeFilterer) FilterDepositToken(opts *bind.FilterOpts, toChain []*big.Int, fromAddress []common.Address, toAddress []common.Address) (*CrossChainBridgeDepositTokenIterator, error) {

	var toChainRule []interface{}
	for _, toChainItem := range toChain {
		toChainRule = append(toChainRule, toChainItem)
	}
	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CrossChainBridge.contract.FilterLogs(opts, "DepositToken", toChainRule, fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainBridgeDepositTokenIterator{contract: _CrossChainBridge.contract, event: "DepositToken", logs: logs, sub: sub}, nil
}

// WatchDepositToken is a free log subscription operation binding the contract event 0x52e2a71d8eeeaf64c3fa0a51a00f5cd6c4d0e7847d213c8d68d823a4b4288d7e.
//
// Solidity: event DepositToken(uint256 fromChain, uint256 indexed toChain, address indexed fromAddress, address indexed toAddress, address fromToken, address toToken, uint256 totalAmount, (string,string,uint256,address) arg7)
func (_CrossChainBridge *CrossChainBridgeFilterer) WatchDepositToken(opts *bind.WatchOpts, sink chan<- *CrossChainBridgeDepositToken, toChain []*big.Int, fromAddress []common.Address, toAddress []common.Address) (event.Subscription, error) {

	var toChainRule []interface{}
	for _, toChainItem := range toChain {
		toChainRule = append(toChainRule, toChainItem)
	}
	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CrossChainBridge.contract.WatchLogs(opts, "DepositToken", toChainRule, fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainBridgeDepositToken)
				if err := _CrossChainBridge.contract.UnpackLog(event, "DepositToken", log); err != nil {
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

// ParseDepositToken is a log parse operation binding the contract event 0x52e2a71d8eeeaf64c3fa0a51a00f5cd6c4d0e7847d213c8d68d823a4b4288d7e.
//
// Solidity: event DepositToken(uint256 fromChain, uint256 indexed toChain, address indexed fromAddress, address indexed toAddress, address fromToken, address toToken, uint256 totalAmount, (string,string,uint256,address) arg7)
func (_CrossChainBridge *CrossChainBridgeFilterer) ParseDepositToken(log types.Log) (*CrossChainBridgeDepositToken, error) {
	event := new(CrossChainBridgeDepositToken)
	if err := _CrossChainBridge.contract.UnpackLog(event, "DepositToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainBridgeWithdrawTokenIterator is returned from FilterWithdrawToken and is used to iterate over the raw logs and unpacked data for WithdrawToken events raised by the CrossChainBridge contract.
type CrossChainBridgeWithdrawTokenIterator struct {
	Event *CrossChainBridgeWithdrawToken // Event containing the contract specifics and raw log

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
func (it *CrossChainBridgeWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainBridgeWithdrawToken)
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
		it.Event = new(CrossChainBridgeWithdrawToken)
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
func (it *CrossChainBridgeWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainBridgeWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainBridgeWithdrawToken represents a WithdrawToken event raised by the CrossChainBridge contract.
type CrossChainBridgeWithdrawToken struct {
	FromChain   *big.Int
	FromAddress common.Address
	ToAddress   common.Address
	FromToken   common.Address
	ToToken     common.Address
	TotalAmount *big.Int
	Arg6        TypesTokenMetadata
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawToken is a free log retrieval operation binding the contract event 0x2295be0bafe4991e6241da8918d557775630cdaaec17b6b5d8eae6a2b3d217d4.
//
// Solidity: event WithdrawToken(uint256 indexed fromChain, address indexed fromAddress, address indexed toAddress, address fromToken, address toToken, uint256 totalAmount, (string,string,uint256,address) arg6)
func (_CrossChainBridge *CrossChainBridgeFilterer) FilterWithdrawToken(opts *bind.FilterOpts, fromChain []*big.Int, fromAddress []common.Address, toAddress []common.Address) (*CrossChainBridgeWithdrawTokenIterator, error) {

	var fromChainRule []interface{}
	for _, fromChainItem := range fromChain {
		fromChainRule = append(fromChainRule, fromChainItem)
	}
	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CrossChainBridge.contract.FilterLogs(opts, "WithdrawToken", fromChainRule, fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainBridgeWithdrawTokenIterator{contract: _CrossChainBridge.contract, event: "WithdrawToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawToken is a free log subscription operation binding the contract event 0x2295be0bafe4991e6241da8918d557775630cdaaec17b6b5d8eae6a2b3d217d4.
//
// Solidity: event WithdrawToken(uint256 indexed fromChain, address indexed fromAddress, address indexed toAddress, address fromToken, address toToken, uint256 totalAmount, (string,string,uint256,address) arg6)
func (_CrossChainBridge *CrossChainBridgeFilterer) WatchWithdrawToken(opts *bind.WatchOpts, sink chan<- *CrossChainBridgeWithdrawToken, fromChain []*big.Int, fromAddress []common.Address, toAddress []common.Address) (event.Subscription, error) {

	var fromChainRule []interface{}
	for _, fromChainItem := range fromChain {
		fromChainRule = append(fromChainRule, fromChainItem)
	}
	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CrossChainBridge.contract.WatchLogs(opts, "WithdrawToken", fromChainRule, fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainBridgeWithdrawToken)
				if err := _CrossChainBridge.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
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

// ParseWithdrawToken is a log parse operation binding the contract event 0x2295be0bafe4991e6241da8918d557775630cdaaec17b6b5d8eae6a2b3d217d4.
//
// Solidity: event WithdrawToken(uint256 indexed fromChain, address indexed fromAddress, address indexed toAddress, address fromToken, address toToken, uint256 totalAmount, (string,string,uint256,address) arg6)
func (_CrossChainBridge *CrossChainBridgeFilterer) ParseWithdrawToken(log types.Log) (*CrossChainBridgeWithdrawToken, error) {
	event := new(CrossChainBridgeWithdrawToken)
	if err := _CrossChainBridge.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
