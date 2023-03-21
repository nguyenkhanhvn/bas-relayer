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

// TypesBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type TypesBlockHeader struct {
	BlockHash    [32]byte
	ParentHash   [32]byte
	BlockNumber  uint64
	Coinbase     common.Address
	ReceiptsRoot [32]byte
	TxsRoot      [32]byte
	StateRoot    [32]byte
}

// RelayHubMetaData contains all meta data concerning the RelayHub contract.
var RelayHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIProofVerificationFunction\",\"name\":\"defaultVerificationFunction\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"BridgeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"BridgeUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainReseted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validatorSet\",\"type\":\"address[]\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"checkValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"blockStart\",\"type\":\"uint256\"}],\"name\":\"checkEpochBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"blockProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"rawReceipt\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofSiblings\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofPath\",\"type\":\"bytes\"}],\"name\":\"checkReceiptProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"checkValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"blockStart\",\"type\":\"uint256\"}],\"name\":\"checkValidatorsAndQuorumReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"enableCrossChainBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLatestEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLatestValidatorSet\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIProofVerificationFunction\",\"name\":\"verificationFunction\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawRegisterBlock\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"epochLength\",\"type\":\"uint32\"}],\"name\":\"registerBridge\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.BlockHeader\",\"name\":\"blockHeader\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"initialValidatorSet\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"resetChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status_\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unregisterBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"blockProofs\",\"type\":\"bytes[]\"}],\"name\":\"updateValidatorSetUsingEpochBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RelayHubABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayHubMetaData.ABI instead.
var RelayHubABI = RelayHubMetaData.ABI

// RelayHub is an auto generated Go binding around an Ethereum contract.
type RelayHub struct {
	RelayHubCaller     // Read-only binding to the contract
	RelayHubTransactor // Write-only binding to the contract
	RelayHubFilterer   // Log filterer for contract events
}

// RelayHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayHubSession struct {
	Contract     *RelayHub         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayHubCallerSession struct {
	Contract *RelayHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RelayHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayHubTransactorSession struct {
	Contract     *RelayHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RelayHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayHubRaw struct {
	Contract *RelayHub // Generic contract binding to access the raw methods on
}

// RelayHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayHubCallerRaw struct {
	Contract *RelayHubCaller // Generic read-only contract binding to access the raw methods on
}

// RelayHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayHubTransactorRaw struct {
	Contract *RelayHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayHub creates a new instance of RelayHub, bound to a specific deployed contract.
func NewRelayHub(address common.Address, backend bind.ContractBackend) (*RelayHub, error) {
	contract, err := bindRelayHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayHub{RelayHubCaller: RelayHubCaller{contract: contract}, RelayHubTransactor: RelayHubTransactor{contract: contract}, RelayHubFilterer: RelayHubFilterer{contract: contract}}, nil
}

// NewRelayHubCaller creates a new read-only instance of RelayHub, bound to a specific deployed contract.
func NewRelayHubCaller(address common.Address, caller bind.ContractCaller) (*RelayHubCaller, error) {
	contract, err := bindRelayHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHubCaller{contract: contract}, nil
}

// NewRelayHubTransactor creates a new write-only instance of RelayHub, bound to a specific deployed contract.
func NewRelayHubTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayHubTransactor, error) {
	contract, err := bindRelayHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHubTransactor{contract: contract}, nil
}

// NewRelayHubFilterer creates a new log filterer instance of RelayHub, bound to a specific deployed contract.
func NewRelayHubFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayHubFilterer, error) {
	contract, err := bindRelayHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayHubFilterer{contract: contract}, nil
}

// bindRelayHub binds a generic wrapper to an already deployed contract.
func bindRelayHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHub *RelayHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayHub.Contract.RelayHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHub *RelayHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHub.Contract.RelayHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHub *RelayHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHub.Contract.RelayHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHub *RelayHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHub *RelayHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHub *RelayHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHub.Contract.contract.Transact(opts, method, params...)
}

// CheckEpochBlock is a free data retrieval call binding the contract method 0x516fd26f.
//
// Solidity: function checkEpochBlock(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHub *RelayHubCaller) CheckEpochBlock(opts *bind.CallOpts, chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	var out []interface{}
	err := _RelayHub.contract.Call(opts, &out, "checkEpochBlock", chainId, checkValidators, blockStart)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckEpochBlock is a free data retrieval call binding the contract method 0x516fd26f.
//
// Solidity: function checkEpochBlock(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHub *RelayHubSession) CheckEpochBlock(chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	return _RelayHub.Contract.CheckEpochBlock(&_RelayHub.CallOpts, chainId, checkValidators, blockStart)
}

// CheckEpochBlock is a free data retrieval call binding the contract method 0x516fd26f.
//
// Solidity: function checkEpochBlock(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHub *RelayHubCallerSession) CheckEpochBlock(chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	return _RelayHub.Contract.CheckEpochBlock(&_RelayHub.CallOpts, chainId, checkValidators, blockStart)
}

// CheckReceiptProof is a free data retrieval call binding the contract method 0x973ebcd8.
//
// Solidity: function checkReceiptProof(uint256 chainId, bytes[] blockProofs, bytes rawReceipt, bytes proofSiblings, bytes proofPath) view returns(bool)
func (_RelayHub *RelayHubCaller) CheckReceiptProof(opts *bind.CallOpts, chainId *big.Int, blockProofs [][]byte, rawReceipt []byte, proofSiblings []byte, proofPath []byte) (bool, error) {
	var out []interface{}
	err := _RelayHub.contract.Call(opts, &out, "checkReceiptProof", chainId, blockProofs, rawReceipt, proofSiblings, proofPath)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckReceiptProof is a free data retrieval call binding the contract method 0x973ebcd8.
//
// Solidity: function checkReceiptProof(uint256 chainId, bytes[] blockProofs, bytes rawReceipt, bytes proofSiblings, bytes proofPath) view returns(bool)
func (_RelayHub *RelayHubSession) CheckReceiptProof(chainId *big.Int, blockProofs [][]byte, rawReceipt []byte, proofSiblings []byte, proofPath []byte) (bool, error) {
	return _RelayHub.Contract.CheckReceiptProof(&_RelayHub.CallOpts, chainId, blockProofs, rawReceipt, proofSiblings, proofPath)
}

// CheckReceiptProof is a free data retrieval call binding the contract method 0x973ebcd8.
//
// Solidity: function checkReceiptProof(uint256 chainId, bytes[] blockProofs, bytes rawReceipt, bytes proofSiblings, bytes proofPath) view returns(bool)
func (_RelayHub *RelayHubCallerSession) CheckReceiptProof(chainId *big.Int, blockProofs [][]byte, rawReceipt []byte, proofSiblings []byte, proofPath []byte) (bool, error) {
	return _RelayHub.Contract.CheckReceiptProof(&_RelayHub.CallOpts, chainId, blockProofs, rawReceipt, proofSiblings, proofPath)
}

// CheckValidatorsAndQuorumReached is a free data retrieval call binding the contract method 0x7518ee25.
//
// Solidity: function checkValidatorsAndQuorumReached(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHub *RelayHubCaller) CheckValidatorsAndQuorumReached(opts *bind.CallOpts, chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	var out []interface{}
	err := _RelayHub.contract.Call(opts, &out, "checkValidatorsAndQuorumReached", chainId, checkValidators, blockStart)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckValidatorsAndQuorumReached is a free data retrieval call binding the contract method 0x7518ee25.
//
// Solidity: function checkValidatorsAndQuorumReached(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHub *RelayHubSession) CheckValidatorsAndQuorumReached(chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	return _RelayHub.Contract.CheckValidatorsAndQuorumReached(&_RelayHub.CallOpts, chainId, checkValidators, blockStart)
}

// CheckValidatorsAndQuorumReached is a free data retrieval call binding the contract method 0x7518ee25.
//
// Solidity: function checkValidatorsAndQuorumReached(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHub *RelayHubCallerSession) CheckValidatorsAndQuorumReached(chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	return _RelayHub.Contract.CheckValidatorsAndQuorumReached(&_RelayHub.CallOpts, chainId, checkValidators, blockStart)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0x0c46a0e1.
//
// Solidity: function getBridgeAddress(uint256 chainId) view returns(address)
func (_RelayHub *RelayHubCaller) GetBridgeAddress(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RelayHub.contract.Call(opts, &out, "getBridgeAddress", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeAddress is a free data retrieval call binding the contract method 0x0c46a0e1.
//
// Solidity: function getBridgeAddress(uint256 chainId) view returns(address)
func (_RelayHub *RelayHubSession) GetBridgeAddress(chainId *big.Int) (common.Address, error) {
	return _RelayHub.Contract.GetBridgeAddress(&_RelayHub.CallOpts, chainId)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0x0c46a0e1.
//
// Solidity: function getBridgeAddress(uint256 chainId) view returns(address)
func (_RelayHub *RelayHubCallerSession) GetBridgeAddress(chainId *big.Int) (common.Address, error) {
	return _RelayHub.Contract.GetBridgeAddress(&_RelayHub.CallOpts, chainId)
}

// GetLatestEpochNumber is a free data retrieval call binding the contract method 0x7e4b28db.
//
// Solidity: function getLatestEpochNumber(uint256 chainId) view returns(uint256)
func (_RelayHub *RelayHubCaller) GetLatestEpochNumber(opts *bind.CallOpts, chainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RelayHub.contract.Call(opts, &out, "getLatestEpochNumber", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestEpochNumber is a free data retrieval call binding the contract method 0x7e4b28db.
//
// Solidity: function getLatestEpochNumber(uint256 chainId) view returns(uint256)
func (_RelayHub *RelayHubSession) GetLatestEpochNumber(chainId *big.Int) (*big.Int, error) {
	return _RelayHub.Contract.GetLatestEpochNumber(&_RelayHub.CallOpts, chainId)
}

// GetLatestEpochNumber is a free data retrieval call binding the contract method 0x7e4b28db.
//
// Solidity: function getLatestEpochNumber(uint256 chainId) view returns(uint256)
func (_RelayHub *RelayHubCallerSession) GetLatestEpochNumber(chainId *big.Int) (*big.Int, error) {
	return _RelayHub.Contract.GetLatestEpochNumber(&_RelayHub.CallOpts, chainId)
}

// GetLatestValidatorSet is a free data retrieval call binding the contract method 0xf9b8165b.
//
// Solidity: function getLatestValidatorSet(uint256 chainId) view returns(address[])
func (_RelayHub *RelayHubCaller) GetLatestValidatorSet(opts *bind.CallOpts, chainId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _RelayHub.contract.Call(opts, &out, "getLatestValidatorSet", chainId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLatestValidatorSet is a free data retrieval call binding the contract method 0xf9b8165b.
//
// Solidity: function getLatestValidatorSet(uint256 chainId) view returns(address[])
func (_RelayHub *RelayHubSession) GetLatestValidatorSet(chainId *big.Int) ([]common.Address, error) {
	return _RelayHub.Contract.GetLatestValidatorSet(&_RelayHub.CallOpts, chainId)
}

// GetLatestValidatorSet is a free data retrieval call binding the contract method 0xf9b8165b.
//
// Solidity: function getLatestValidatorSet(uint256 chainId) view returns(address[])
func (_RelayHub *RelayHubCallerSession) GetLatestValidatorSet(chainId *big.Int) ([]common.Address, error) {
	return _RelayHub.Contract.GetLatestValidatorSet(&_RelayHub.CallOpts, chainId)
}

// EnableCrossChainBridge is a paid mutator transaction binding the contract method 0xd2e1a496.
//
// Solidity: function enableCrossChainBridge(uint256 chainId, address bridgeAddress) returns()
func (_RelayHub *RelayHubTransactor) EnableCrossChainBridge(opts *bind.TransactOpts, chainId *big.Int, bridgeAddress common.Address) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "enableCrossChainBridge", chainId, bridgeAddress)
}

// EnableCrossChainBridge is a paid mutator transaction binding the contract method 0xd2e1a496.
//
// Solidity: function enableCrossChainBridge(uint256 chainId, address bridgeAddress) returns()
func (_RelayHub *RelayHubSession) EnableCrossChainBridge(chainId *big.Int, bridgeAddress common.Address) (*types.Transaction, error) {
	return _RelayHub.Contract.EnableCrossChainBridge(&_RelayHub.TransactOpts, chainId, bridgeAddress)
}

// EnableCrossChainBridge is a paid mutator transaction binding the contract method 0xd2e1a496.
//
// Solidity: function enableCrossChainBridge(uint256 chainId, address bridgeAddress) returns()
func (_RelayHub *RelayHubTransactorSession) EnableCrossChainBridge(chainId *big.Int, bridgeAddress common.Address) (*types.Transaction, error) {
	return _RelayHub.Contract.EnableCrossChainBridge(&_RelayHub.TransactOpts, chainId, bridgeAddress)
}

// RegisterBridge is a paid mutator transaction binding the contract method 0xc80f0c20.
//
// Solidity: function registerBridge(uint256 chainId, address verificationFunction, bytes rawRegisterBlock, address bridgeAddress, uint32 epochLength) returns((bytes32,bytes32,uint64,address,bytes32,bytes32,bytes32) blockHeader, address[] initialValidatorSet)
func (_RelayHub *RelayHubTransactor) RegisterBridge(opts *bind.TransactOpts, chainId *big.Int, verificationFunction common.Address, rawRegisterBlock []byte, bridgeAddress common.Address, epochLength uint32) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "registerBridge", chainId, verificationFunction, rawRegisterBlock, bridgeAddress, epochLength)
}

// RegisterBridge is a paid mutator transaction binding the contract method 0xc80f0c20.
//
// Solidity: function registerBridge(uint256 chainId, address verificationFunction, bytes rawRegisterBlock, address bridgeAddress, uint32 epochLength) returns((bytes32,bytes32,uint64,address,bytes32,bytes32,bytes32) blockHeader, address[] initialValidatorSet)
func (_RelayHub *RelayHubSession) RegisterBridge(chainId *big.Int, verificationFunction common.Address, rawRegisterBlock []byte, bridgeAddress common.Address, epochLength uint32) (*types.Transaction, error) {
	return _RelayHub.Contract.RegisterBridge(&_RelayHub.TransactOpts, chainId, verificationFunction, rawRegisterBlock, bridgeAddress, epochLength)
}

// RegisterBridge is a paid mutator transaction binding the contract method 0xc80f0c20.
//
// Solidity: function registerBridge(uint256 chainId, address verificationFunction, bytes rawRegisterBlock, address bridgeAddress, uint32 epochLength) returns((bytes32,bytes32,uint64,address,bytes32,bytes32,bytes32) blockHeader, address[] initialValidatorSet)
func (_RelayHub *RelayHubTransactorSession) RegisterBridge(chainId *big.Int, verificationFunction common.Address, rawRegisterBlock []byte, bridgeAddress common.Address, epochLength uint32) (*types.Transaction, error) {
	return _RelayHub.Contract.RegisterBridge(&_RelayHub.TransactOpts, chainId, verificationFunction, rawRegisterBlock, bridgeAddress, epochLength)
}

// ResetChain is a paid mutator transaction binding the contract method 0x6f72020d.
//
// Solidity: function resetChain(uint256 chainId) returns()
func (_RelayHub *RelayHubTransactor) ResetChain(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "resetChain", chainId)
}

// ResetChain is a paid mutator transaction binding the contract method 0x6f72020d.
//
// Solidity: function resetChain(uint256 chainId) returns()
func (_RelayHub *RelayHubSession) ResetChain(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.Contract.ResetChain(&_RelayHub.TransactOpts, chainId)
}

// ResetChain is a paid mutator transaction binding the contract method 0x6f72020d.
//
// Solidity: function resetChain(uint256 chainId) returns()
func (_RelayHub *RelayHubTransactorSession) ResetChain(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.Contract.ResetChain(&_RelayHub.TransactOpts, chainId)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_RelayHub *RelayHubTransactor) SetOperator(opts *bind.TransactOpts, operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "setOperator", operator_, status_)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_RelayHub *RelayHubSession) SetOperator(operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _RelayHub.Contract.SetOperator(&_RelayHub.TransactOpts, operator_, status_)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_RelayHub *RelayHubTransactorSession) SetOperator(operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _RelayHub.Contract.SetOperator(&_RelayHub.TransactOpts, operator_, status_)
}

// UnregisterBridge is a paid mutator transaction binding the contract method 0x7018b4f9.
//
// Solidity: function unregisterBridge(uint256 chainId) returns()
func (_RelayHub *RelayHubTransactor) UnregisterBridge(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "unregisterBridge", chainId)
}

// UnregisterBridge is a paid mutator transaction binding the contract method 0x7018b4f9.
//
// Solidity: function unregisterBridge(uint256 chainId) returns()
func (_RelayHub *RelayHubSession) UnregisterBridge(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.Contract.UnregisterBridge(&_RelayHub.TransactOpts, chainId)
}

// UnregisterBridge is a paid mutator transaction binding the contract method 0x7018b4f9.
//
// Solidity: function unregisterBridge(uint256 chainId) returns()
func (_RelayHub *RelayHubTransactorSession) UnregisterBridge(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.Contract.UnregisterBridge(&_RelayHub.TransactOpts, chainId)
}

// UpdateValidatorSetUsingEpochBlocks is a paid mutator transaction binding the contract method 0x383d0ad9.
//
// Solidity: function updateValidatorSetUsingEpochBlocks(uint256 chainId, bytes[] blockProofs) returns()
func (_RelayHub *RelayHubTransactor) UpdateValidatorSetUsingEpochBlocks(opts *bind.TransactOpts, chainId *big.Int, blockProofs [][]byte) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "updateValidatorSetUsingEpochBlocks", chainId, blockProofs)
}

// UpdateValidatorSetUsingEpochBlocks is a paid mutator transaction binding the contract method 0x383d0ad9.
//
// Solidity: function updateValidatorSetUsingEpochBlocks(uint256 chainId, bytes[] blockProofs) returns()
func (_RelayHub *RelayHubSession) UpdateValidatorSetUsingEpochBlocks(chainId *big.Int, blockProofs [][]byte) (*types.Transaction, error) {
	return _RelayHub.Contract.UpdateValidatorSetUsingEpochBlocks(&_RelayHub.TransactOpts, chainId, blockProofs)
}

// UpdateValidatorSetUsingEpochBlocks is a paid mutator transaction binding the contract method 0x383d0ad9.
//
// Solidity: function updateValidatorSetUsingEpochBlocks(uint256 chainId, bytes[] blockProofs) returns()
func (_RelayHub *RelayHubTransactorSession) UpdateValidatorSetUsingEpochBlocks(chainId *big.Int, blockProofs [][]byte) (*types.Transaction, error) {
	return _RelayHub.Contract.UpdateValidatorSetUsingEpochBlocks(&_RelayHub.TransactOpts, chainId, blockProofs)
}

// RelayHubBridgeRegisteredIterator is returned from FilterBridgeRegistered and is used to iterate over the raw logs and unpacked data for BridgeRegistered events raised by the RelayHub contract.
type RelayHubBridgeRegisteredIterator struct {
	Event *RelayHubBridgeRegistered // Event containing the contract specifics and raw log

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
func (it *RelayHubBridgeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubBridgeRegistered)
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
		it.Event = new(RelayHubBridgeRegistered)
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
func (it *RelayHubBridgeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubBridgeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubBridgeRegistered represents a BridgeRegistered event raised by the RelayHub contract.
type RelayHubBridgeRegistered struct {
	ChainId       *big.Int
	BridgeAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRegistered is a free log retrieval operation binding the contract event 0xb72d37f67f55309a7342120ad674525186207bb1835dd860c16280472ea6fc50.
//
// Solidity: event BridgeRegistered(uint256 indexed chainId, address indexed bridgeAddress)
func (_RelayHub *RelayHubFilterer) FilterBridgeRegistered(opts *bind.FilterOpts, chainId []*big.Int, bridgeAddress []common.Address) (*RelayHubBridgeRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var bridgeAddressRule []interface{}
	for _, bridgeAddressItem := range bridgeAddress {
		bridgeAddressRule = append(bridgeAddressRule, bridgeAddressItem)
	}

	logs, sub, err := _RelayHub.contract.FilterLogs(opts, "BridgeRegistered", chainIdRule, bridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubBridgeRegisteredIterator{contract: _RelayHub.contract, event: "BridgeRegistered", logs: logs, sub: sub}, nil
}

// WatchBridgeRegistered is a free log subscription operation binding the contract event 0xb72d37f67f55309a7342120ad674525186207bb1835dd860c16280472ea6fc50.
//
// Solidity: event BridgeRegistered(uint256 indexed chainId, address indexed bridgeAddress)
func (_RelayHub *RelayHubFilterer) WatchBridgeRegistered(opts *bind.WatchOpts, sink chan<- *RelayHubBridgeRegistered, chainId []*big.Int, bridgeAddress []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var bridgeAddressRule []interface{}
	for _, bridgeAddressItem := range bridgeAddress {
		bridgeAddressRule = append(bridgeAddressRule, bridgeAddressItem)
	}

	logs, sub, err := _RelayHub.contract.WatchLogs(opts, "BridgeRegistered", chainIdRule, bridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubBridgeRegistered)
				if err := _RelayHub.contract.UnpackLog(event, "BridgeRegistered", log); err != nil {
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

// ParseBridgeRegistered is a log parse operation binding the contract event 0xb72d37f67f55309a7342120ad674525186207bb1835dd860c16280472ea6fc50.
//
// Solidity: event BridgeRegistered(uint256 indexed chainId, address indexed bridgeAddress)
func (_RelayHub *RelayHubFilterer) ParseBridgeRegistered(log types.Log) (*RelayHubBridgeRegistered, error) {
	event := new(RelayHubBridgeRegistered)
	if err := _RelayHub.contract.UnpackLog(event, "BridgeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayHubBridgeUnregisteredIterator is returned from FilterBridgeUnregistered and is used to iterate over the raw logs and unpacked data for BridgeUnregistered events raised by the RelayHub contract.
type RelayHubBridgeUnregisteredIterator struct {
	Event *RelayHubBridgeUnregistered // Event containing the contract specifics and raw log

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
func (it *RelayHubBridgeUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubBridgeUnregistered)
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
		it.Event = new(RelayHubBridgeUnregistered)
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
func (it *RelayHubBridgeUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubBridgeUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubBridgeUnregistered represents a BridgeUnregistered event raised by the RelayHub contract.
type RelayHubBridgeUnregistered struct {
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeUnregistered is a free log retrieval operation binding the contract event 0x5ed0f717c1d62eb20db7f8f67a7f08a656c882be0834c2975e4991b546faac87.
//
// Solidity: event BridgeUnregistered(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) FilterBridgeUnregistered(opts *bind.FilterOpts, chainId []*big.Int) (*RelayHubBridgeUnregisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.FilterLogs(opts, "BridgeUnregistered", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubBridgeUnregisteredIterator{contract: _RelayHub.contract, event: "BridgeUnregistered", logs: logs, sub: sub}, nil
}

// WatchBridgeUnregistered is a free log subscription operation binding the contract event 0x5ed0f717c1d62eb20db7f8f67a7f08a656c882be0834c2975e4991b546faac87.
//
// Solidity: event BridgeUnregistered(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) WatchBridgeUnregistered(opts *bind.WatchOpts, sink chan<- *RelayHubBridgeUnregistered, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.WatchLogs(opts, "BridgeUnregistered", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubBridgeUnregistered)
				if err := _RelayHub.contract.UnpackLog(event, "BridgeUnregistered", log); err != nil {
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

// ParseBridgeUnregistered is a log parse operation binding the contract event 0x5ed0f717c1d62eb20db7f8f67a7f08a656c882be0834c2975e4991b546faac87.
//
// Solidity: event BridgeUnregistered(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) ParseBridgeUnregistered(log types.Log) (*RelayHubBridgeUnregistered, error) {
	event := new(RelayHubBridgeUnregistered)
	if err := _RelayHub.contract.UnpackLog(event, "BridgeUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayHubChainResetedIterator is returned from FilterChainReseted and is used to iterate over the raw logs and unpacked data for ChainReseted events raised by the RelayHub contract.
type RelayHubChainResetedIterator struct {
	Event *RelayHubChainReseted // Event containing the contract specifics and raw log

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
func (it *RelayHubChainResetedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubChainReseted)
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
		it.Event = new(RelayHubChainReseted)
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
func (it *RelayHubChainResetedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubChainResetedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubChainReseted represents a ChainReseted event raised by the RelayHub contract.
type RelayHubChainReseted struct {
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainReseted is a free log retrieval operation binding the contract event 0x1b3e6308933a8c83e369ce09eb5dc180027a75240e78b85b9832061d9f1716ed.
//
// Solidity: event ChainReseted(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) FilterChainReseted(opts *bind.FilterOpts, chainId []*big.Int) (*RelayHubChainResetedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.FilterLogs(opts, "ChainReseted", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubChainResetedIterator{contract: _RelayHub.contract, event: "ChainReseted", logs: logs, sub: sub}, nil
}

// WatchChainReseted is a free log subscription operation binding the contract event 0x1b3e6308933a8c83e369ce09eb5dc180027a75240e78b85b9832061d9f1716ed.
//
// Solidity: event ChainReseted(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) WatchChainReseted(opts *bind.WatchOpts, sink chan<- *RelayHubChainReseted, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.WatchLogs(opts, "ChainReseted", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubChainReseted)
				if err := _RelayHub.contract.UnpackLog(event, "ChainReseted", log); err != nil {
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

// ParseChainReseted is a log parse operation binding the contract event 0x1b3e6308933a8c83e369ce09eb5dc180027a75240e78b85b9832061d9f1716ed.
//
// Solidity: event ChainReseted(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) ParseChainReseted(log types.Log) (*RelayHubChainReseted, error) {
	event := new(RelayHubChainReseted)
	if err := _RelayHub.contract.UnpackLog(event, "ChainReseted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayHubValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the RelayHub contract.
type RelayHubValidatorSetUpdatedIterator struct {
	Event *RelayHubValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *RelayHubValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubValidatorSetUpdated)
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
		it.Event = new(RelayHubValidatorSetUpdated)
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
func (it *RelayHubValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubValidatorSetUpdated represents a ValidatorSetUpdated event raised by the RelayHub contract.
type RelayHubValidatorSetUpdated struct {
	ChainId      *big.Int
	ValidatorSet []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed chainId, address[] validatorSet)
func (_RelayHub *RelayHubFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*RelayHubValidatorSetUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.FilterLogs(opts, "ValidatorSetUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubValidatorSetUpdatedIterator{contract: _RelayHub.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed chainId, address[] validatorSet)
func (_RelayHub *RelayHubFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *RelayHubValidatorSetUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.WatchLogs(opts, "ValidatorSetUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubValidatorSetUpdated)
				if err := _RelayHub.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed chainId, address[] validatorSet)
func (_RelayHub *RelayHubFilterer) ParseValidatorSetUpdated(log types.Log) (*RelayHubValidatorSetUpdated, error) {
	event := new(RelayHubValidatorSetUpdated)
	if err := _RelayHub.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
