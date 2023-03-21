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
// already define in relay_hub.go
// type TypesBlockHeader struct {
// 	BlockHash    [32]byte
// 	ParentHash   [32]byte
// 	BlockNumber  uint64
// 	Coinbase     common.Address
// 	ReceiptsRoot [32]byte
// 	TxsRoot      [32]byte
// 	StateRoot    [32]byte
// }

// RelayHubBetaMetaData contains all meta data concerning the RelayHubBeta contract.
var RelayHubBetaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIProofVerificationFunction\",\"name\":\"defaultVerificationFunction\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"BridgeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"BridgeUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainReseted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validatorSet\",\"type\":\"address[]\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"checkValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"blockStart\",\"type\":\"uint256\"}],\"name\":\"checkEpochBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"blockProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"rawReceipt\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofSiblings\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofPath\",\"type\":\"bytes\"}],\"name\":\"checkReceiptProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"checkValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"blockStart\",\"type\":\"uint256\"}],\"name\":\"checkValidatorsAndQuorumReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"enableCrossChainBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLatestEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLatestValidatorSet\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"getValidatorSetForEpoch\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIProofVerificationFunction\",\"name\":\"verificationFunction\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawRegisterBlock\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"epochLength\",\"type\":\"uint32\"}],\"name\":\"registerBridge\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.BlockHeader\",\"name\":\"blockHeader\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"initialValidatorSet\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"resetChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status_\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unregisterBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"blockProofs\",\"type\":\"bytes[]\"}],\"name\":\"updateValidatorSetUsingEpochBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RelayHubBetaABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayHubBetaMetaData.ABI instead.
var RelayHubBetaABI = RelayHubBetaMetaData.ABI

// RelayHubBeta is an auto generated Go binding around an Ethereum contract.
type RelayHubBeta struct {
	RelayHubBetaCaller     // Read-only binding to the contract
	RelayHubBetaTransactor // Write-only binding to the contract
	RelayHubBetaFilterer   // Log filterer for contract events
}

// RelayHubBetaCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayHubBetaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubBetaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayHubBetaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubBetaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayHubBetaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubBetaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayHubBetaSession struct {
	Contract     *RelayHubBeta     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayHubBetaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayHubBetaCallerSession struct {
	Contract *RelayHubBetaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RelayHubBetaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayHubBetaTransactorSession struct {
	Contract     *RelayHubBetaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RelayHubBetaRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayHubBetaRaw struct {
	Contract *RelayHubBeta // Generic contract binding to access the raw methods on
}

// RelayHubBetaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayHubBetaCallerRaw struct {
	Contract *RelayHubBetaCaller // Generic read-only contract binding to access the raw methods on
}

// RelayHubBetaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayHubBetaTransactorRaw struct {
	Contract *RelayHubBetaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayHubBeta creates a new instance of RelayHubBeta, bound to a specific deployed contract.
func NewRelayHubBeta(address common.Address, backend bind.ContractBackend) (*RelayHubBeta, error) {
	contract, err := bindRelayHubBeta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayHubBeta{RelayHubBetaCaller: RelayHubBetaCaller{contract: contract}, RelayHubBetaTransactor: RelayHubBetaTransactor{contract: contract}, RelayHubBetaFilterer: RelayHubBetaFilterer{contract: contract}}, nil
}

// NewRelayHubBetaCaller creates a new read-only instance of RelayHubBeta, bound to a specific deployed contract.
func NewRelayHubBetaCaller(address common.Address, caller bind.ContractCaller) (*RelayHubBetaCaller, error) {
	contract, err := bindRelayHubBeta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHubBetaCaller{contract: contract}, nil
}

// NewRelayHubBetaTransactor creates a new write-only instance of RelayHubBeta, bound to a specific deployed contract.
func NewRelayHubBetaTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayHubBetaTransactor, error) {
	contract, err := bindRelayHubBeta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHubBetaTransactor{contract: contract}, nil
}

// NewRelayHubBetaFilterer creates a new log filterer instance of RelayHubBeta, bound to a specific deployed contract.
func NewRelayHubBetaFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayHubBetaFilterer, error) {
	contract, err := bindRelayHubBeta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayHubBetaFilterer{contract: contract}, nil
}

// bindRelayHubBeta binds a generic wrapper to an already deployed contract.
func bindRelayHubBeta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayHubBetaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHubBeta *RelayHubBetaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayHubBeta.Contract.RelayHubBetaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHubBeta *RelayHubBetaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.RelayHubBetaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHubBeta *RelayHubBetaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.RelayHubBetaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHubBeta *RelayHubBetaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayHubBeta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHubBeta *RelayHubBetaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHubBeta *RelayHubBetaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.contract.Transact(opts, method, params...)
}

// CheckEpochBlock is a free data retrieval call binding the contract method 0x516fd26f.
//
// Solidity: function checkEpochBlock(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHubBeta *RelayHubBetaCaller) CheckEpochBlock(opts *bind.CallOpts, chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	var out []interface{}
	err := _RelayHubBeta.contract.Call(opts, &out, "checkEpochBlock", chainId, checkValidators, blockStart)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckEpochBlock is a free data retrieval call binding the contract method 0x516fd26f.
//
// Solidity: function checkEpochBlock(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHubBeta *RelayHubBetaSession) CheckEpochBlock(chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	return _RelayHubBeta.Contract.CheckEpochBlock(&_RelayHubBeta.CallOpts, chainId, checkValidators, blockStart)
}

// CheckEpochBlock is a free data retrieval call binding the contract method 0x516fd26f.
//
// Solidity: function checkEpochBlock(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHubBeta *RelayHubBetaCallerSession) CheckEpochBlock(chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	return _RelayHubBeta.Contract.CheckEpochBlock(&_RelayHubBeta.CallOpts, chainId, checkValidators, blockStart)
}

// CheckReceiptProof is a free data retrieval call binding the contract method 0x973ebcd8.
//
// Solidity: function checkReceiptProof(uint256 chainId, bytes[] blockProofs, bytes rawReceipt, bytes proofSiblings, bytes proofPath) view returns(bool)
func (_RelayHubBeta *RelayHubBetaCaller) CheckReceiptProof(opts *bind.CallOpts, chainId *big.Int, blockProofs [][]byte, rawReceipt []byte, proofSiblings []byte, proofPath []byte) (bool, error) {
	var out []interface{}
	err := _RelayHubBeta.contract.Call(opts, &out, "checkReceiptProof", chainId, blockProofs, rawReceipt, proofSiblings, proofPath)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckReceiptProof is a free data retrieval call binding the contract method 0x973ebcd8.
//
// Solidity: function checkReceiptProof(uint256 chainId, bytes[] blockProofs, bytes rawReceipt, bytes proofSiblings, bytes proofPath) view returns(bool)
func (_RelayHubBeta *RelayHubBetaSession) CheckReceiptProof(chainId *big.Int, blockProofs [][]byte, rawReceipt []byte, proofSiblings []byte, proofPath []byte) (bool, error) {
	return _RelayHubBeta.Contract.CheckReceiptProof(&_RelayHubBeta.CallOpts, chainId, blockProofs, rawReceipt, proofSiblings, proofPath)
}

// CheckReceiptProof is a free data retrieval call binding the contract method 0x973ebcd8.
//
// Solidity: function checkReceiptProof(uint256 chainId, bytes[] blockProofs, bytes rawReceipt, bytes proofSiblings, bytes proofPath) view returns(bool)
func (_RelayHubBeta *RelayHubBetaCallerSession) CheckReceiptProof(chainId *big.Int, blockProofs [][]byte, rawReceipt []byte, proofSiblings []byte, proofPath []byte) (bool, error) {
	return _RelayHubBeta.Contract.CheckReceiptProof(&_RelayHubBeta.CallOpts, chainId, blockProofs, rawReceipt, proofSiblings, proofPath)
}

// CheckValidatorsAndQuorumReached is a free data retrieval call binding the contract method 0x7518ee25.
//
// Solidity: function checkValidatorsAndQuorumReached(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHubBeta *RelayHubBetaCaller) CheckValidatorsAndQuorumReached(opts *bind.CallOpts, chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	var out []interface{}
	err := _RelayHubBeta.contract.Call(opts, &out, "checkValidatorsAndQuorumReached", chainId, checkValidators, blockStart)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckValidatorsAndQuorumReached is a free data retrieval call binding the contract method 0x7518ee25.
//
// Solidity: function checkValidatorsAndQuorumReached(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHubBeta *RelayHubBetaSession) CheckValidatorsAndQuorumReached(chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	return _RelayHubBeta.Contract.CheckValidatorsAndQuorumReached(&_RelayHubBeta.CallOpts, chainId, checkValidators, blockStart)
}

// CheckValidatorsAndQuorumReached is a free data retrieval call binding the contract method 0x7518ee25.
//
// Solidity: function checkValidatorsAndQuorumReached(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHubBeta *RelayHubBetaCallerSession) CheckValidatorsAndQuorumReached(chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	return _RelayHubBeta.Contract.CheckValidatorsAndQuorumReached(&_RelayHubBeta.CallOpts, chainId, checkValidators, blockStart)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0x0c46a0e1.
//
// Solidity: function getBridgeAddress(uint256 chainId) view returns(address)
func (_RelayHubBeta *RelayHubBetaCaller) GetBridgeAddress(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RelayHubBeta.contract.Call(opts, &out, "getBridgeAddress", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeAddress is a free data retrieval call binding the contract method 0x0c46a0e1.
//
// Solidity: function getBridgeAddress(uint256 chainId) view returns(address)
func (_RelayHubBeta *RelayHubBetaSession) GetBridgeAddress(chainId *big.Int) (common.Address, error) {
	return _RelayHubBeta.Contract.GetBridgeAddress(&_RelayHubBeta.CallOpts, chainId)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0x0c46a0e1.
//
// Solidity: function getBridgeAddress(uint256 chainId) view returns(address)
func (_RelayHubBeta *RelayHubBetaCallerSession) GetBridgeAddress(chainId *big.Int) (common.Address, error) {
	return _RelayHubBeta.Contract.GetBridgeAddress(&_RelayHubBeta.CallOpts, chainId)
}

// GetLatestEpochNumber is a free data retrieval call binding the contract method 0x7e4b28db.
//
// Solidity: function getLatestEpochNumber(uint256 chainId) view returns(uint256)
func (_RelayHubBeta *RelayHubBetaCaller) GetLatestEpochNumber(opts *bind.CallOpts, chainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RelayHubBeta.contract.Call(opts, &out, "getLatestEpochNumber", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestEpochNumber is a free data retrieval call binding the contract method 0x7e4b28db.
//
// Solidity: function getLatestEpochNumber(uint256 chainId) view returns(uint256)
func (_RelayHubBeta *RelayHubBetaSession) GetLatestEpochNumber(chainId *big.Int) (*big.Int, error) {
	return _RelayHubBeta.Contract.GetLatestEpochNumber(&_RelayHubBeta.CallOpts, chainId)
}

// GetLatestEpochNumber is a free data retrieval call binding the contract method 0x7e4b28db.
//
// Solidity: function getLatestEpochNumber(uint256 chainId) view returns(uint256)
func (_RelayHubBeta *RelayHubBetaCallerSession) GetLatestEpochNumber(chainId *big.Int) (*big.Int, error) {
	return _RelayHubBeta.Contract.GetLatestEpochNumber(&_RelayHubBeta.CallOpts, chainId)
}

// GetLatestValidatorSet is a free data retrieval call binding the contract method 0xf9b8165b.
//
// Solidity: function getLatestValidatorSet(uint256 chainId) view returns(address[])
func (_RelayHubBeta *RelayHubBetaCaller) GetLatestValidatorSet(opts *bind.CallOpts, chainId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _RelayHubBeta.contract.Call(opts, &out, "getLatestValidatorSet", chainId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLatestValidatorSet is a free data retrieval call binding the contract method 0xf9b8165b.
//
// Solidity: function getLatestValidatorSet(uint256 chainId) view returns(address[])
func (_RelayHubBeta *RelayHubBetaSession) GetLatestValidatorSet(chainId *big.Int) ([]common.Address, error) {
	return _RelayHubBeta.Contract.GetLatestValidatorSet(&_RelayHubBeta.CallOpts, chainId)
}

// GetLatestValidatorSet is a free data retrieval call binding the contract method 0xf9b8165b.
//
// Solidity: function getLatestValidatorSet(uint256 chainId) view returns(address[])
func (_RelayHubBeta *RelayHubBetaCallerSession) GetLatestValidatorSet(chainId *big.Int) ([]common.Address, error) {
	return _RelayHubBeta.Contract.GetLatestValidatorSet(&_RelayHubBeta.CallOpts, chainId)
}

// GetValidatorSetForEpoch is a free data retrieval call binding the contract method 0x209262db.
//
// Solidity: function getValidatorSetForEpoch(uint256 chainId, uint256 epochNumber) view returns(address[])
func (_RelayHubBeta *RelayHubBetaCaller) GetValidatorSetForEpoch(opts *bind.CallOpts, chainId *big.Int, epochNumber *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _RelayHubBeta.contract.Call(opts, &out, "getValidatorSetForEpoch", chainId, epochNumber)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorSetForEpoch is a free data retrieval call binding the contract method 0x209262db.
//
// Solidity: function getValidatorSetForEpoch(uint256 chainId, uint256 epochNumber) view returns(address[])
func (_RelayHubBeta *RelayHubBetaSession) GetValidatorSetForEpoch(chainId *big.Int, epochNumber *big.Int) ([]common.Address, error) {
	return _RelayHubBeta.Contract.GetValidatorSetForEpoch(&_RelayHubBeta.CallOpts, chainId, epochNumber)
}

// GetValidatorSetForEpoch is a free data retrieval call binding the contract method 0x209262db.
//
// Solidity: function getValidatorSetForEpoch(uint256 chainId, uint256 epochNumber) view returns(address[])
func (_RelayHubBeta *RelayHubBetaCallerSession) GetValidatorSetForEpoch(chainId *big.Int, epochNumber *big.Int) ([]common.Address, error) {
	return _RelayHubBeta.Contract.GetValidatorSetForEpoch(&_RelayHubBeta.CallOpts, chainId, epochNumber)
}

// EnableCrossChainBridge is a paid mutator transaction binding the contract method 0xd2e1a496.
//
// Solidity: function enableCrossChainBridge(uint256 chainId, address bridgeAddress) returns()
func (_RelayHubBeta *RelayHubBetaTransactor) EnableCrossChainBridge(opts *bind.TransactOpts, chainId *big.Int, bridgeAddress common.Address) (*types.Transaction, error) {
	return _RelayHubBeta.contract.Transact(opts, "enableCrossChainBridge", chainId, bridgeAddress)
}

// EnableCrossChainBridge is a paid mutator transaction binding the contract method 0xd2e1a496.
//
// Solidity: function enableCrossChainBridge(uint256 chainId, address bridgeAddress) returns()
func (_RelayHubBeta *RelayHubBetaSession) EnableCrossChainBridge(chainId *big.Int, bridgeAddress common.Address) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.EnableCrossChainBridge(&_RelayHubBeta.TransactOpts, chainId, bridgeAddress)
}

// EnableCrossChainBridge is a paid mutator transaction binding the contract method 0xd2e1a496.
//
// Solidity: function enableCrossChainBridge(uint256 chainId, address bridgeAddress) returns()
func (_RelayHubBeta *RelayHubBetaTransactorSession) EnableCrossChainBridge(chainId *big.Int, bridgeAddress common.Address) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.EnableCrossChainBridge(&_RelayHubBeta.TransactOpts, chainId, bridgeAddress)
}

// RegisterBridge is a paid mutator transaction binding the contract method 0xc80f0c20.
//
// Solidity: function registerBridge(uint256 chainId, address verificationFunction, bytes rawRegisterBlock, address bridgeAddress, uint32 epochLength) returns((bytes32,bytes32,uint64,address,bytes32,bytes32,bytes32) blockHeader, address[] initialValidatorSet)
func (_RelayHubBeta *RelayHubBetaTransactor) RegisterBridge(opts *bind.TransactOpts, chainId *big.Int, verificationFunction common.Address, rawRegisterBlock []byte, bridgeAddress common.Address, epochLength uint32) (*types.Transaction, error) {
	return _RelayHubBeta.contract.Transact(opts, "registerBridge", chainId, verificationFunction, rawRegisterBlock, bridgeAddress, epochLength)
}

// RegisterBridge is a paid mutator transaction binding the contract method 0xc80f0c20.
//
// Solidity: function registerBridge(uint256 chainId, address verificationFunction, bytes rawRegisterBlock, address bridgeAddress, uint32 epochLength) returns((bytes32,bytes32,uint64,address,bytes32,bytes32,bytes32) blockHeader, address[] initialValidatorSet)
func (_RelayHubBeta *RelayHubBetaSession) RegisterBridge(chainId *big.Int, verificationFunction common.Address, rawRegisterBlock []byte, bridgeAddress common.Address, epochLength uint32) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.RegisterBridge(&_RelayHubBeta.TransactOpts, chainId, verificationFunction, rawRegisterBlock, bridgeAddress, epochLength)
}

// RegisterBridge is a paid mutator transaction binding the contract method 0xc80f0c20.
//
// Solidity: function registerBridge(uint256 chainId, address verificationFunction, bytes rawRegisterBlock, address bridgeAddress, uint32 epochLength) returns((bytes32,bytes32,uint64,address,bytes32,bytes32,bytes32) blockHeader, address[] initialValidatorSet)
func (_RelayHubBeta *RelayHubBetaTransactorSession) RegisterBridge(chainId *big.Int, verificationFunction common.Address, rawRegisterBlock []byte, bridgeAddress common.Address, epochLength uint32) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.RegisterBridge(&_RelayHubBeta.TransactOpts, chainId, verificationFunction, rawRegisterBlock, bridgeAddress, epochLength)
}

// ResetChain is a paid mutator transaction binding the contract method 0x6f72020d.
//
// Solidity: function resetChain(uint256 chainId) returns()
func (_RelayHubBeta *RelayHubBetaTransactor) ResetChain(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _RelayHubBeta.contract.Transact(opts, "resetChain", chainId)
}

// ResetChain is a paid mutator transaction binding the contract method 0x6f72020d.
//
// Solidity: function resetChain(uint256 chainId) returns()
func (_RelayHubBeta *RelayHubBetaSession) ResetChain(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.ResetChain(&_RelayHubBeta.TransactOpts, chainId)
}

// ResetChain is a paid mutator transaction binding the contract method 0x6f72020d.
//
// Solidity: function resetChain(uint256 chainId) returns()
func (_RelayHubBeta *RelayHubBetaTransactorSession) ResetChain(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.ResetChain(&_RelayHubBeta.TransactOpts, chainId)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_RelayHubBeta *RelayHubBetaTransactor) SetOperator(opts *bind.TransactOpts, operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _RelayHubBeta.contract.Transact(opts, "setOperator", operator_, status_)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_RelayHubBeta *RelayHubBetaSession) SetOperator(operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.SetOperator(&_RelayHubBeta.TransactOpts, operator_, status_)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_RelayHubBeta *RelayHubBetaTransactorSession) SetOperator(operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.SetOperator(&_RelayHubBeta.TransactOpts, operator_, status_)
}

// UnregisterBridge is a paid mutator transaction binding the contract method 0x7018b4f9.
//
// Solidity: function unregisterBridge(uint256 chainId) returns()
func (_RelayHubBeta *RelayHubBetaTransactor) UnregisterBridge(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _RelayHubBeta.contract.Transact(opts, "unregisterBridge", chainId)
}

// UnregisterBridge is a paid mutator transaction binding the contract method 0x7018b4f9.
//
// Solidity: function unregisterBridge(uint256 chainId) returns()
func (_RelayHubBeta *RelayHubBetaSession) UnregisterBridge(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.UnregisterBridge(&_RelayHubBeta.TransactOpts, chainId)
}

// UnregisterBridge is a paid mutator transaction binding the contract method 0x7018b4f9.
//
// Solidity: function unregisterBridge(uint256 chainId) returns()
func (_RelayHubBeta *RelayHubBetaTransactorSession) UnregisterBridge(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.UnregisterBridge(&_RelayHubBeta.TransactOpts, chainId)
}

// UpdateValidatorSetUsingEpochBlocks is a paid mutator transaction binding the contract method 0x383d0ad9.
//
// Solidity: function updateValidatorSetUsingEpochBlocks(uint256 chainId, bytes[] blockProofs) returns()
func (_RelayHubBeta *RelayHubBetaTransactor) UpdateValidatorSetUsingEpochBlocks(opts *bind.TransactOpts, chainId *big.Int, blockProofs [][]byte) (*types.Transaction, error) {
	return _RelayHubBeta.contract.Transact(opts, "updateValidatorSetUsingEpochBlocks", chainId, blockProofs)
}

// UpdateValidatorSetUsingEpochBlocks is a paid mutator transaction binding the contract method 0x383d0ad9.
//
// Solidity: function updateValidatorSetUsingEpochBlocks(uint256 chainId, bytes[] blockProofs) returns()
func (_RelayHubBeta *RelayHubBetaSession) UpdateValidatorSetUsingEpochBlocks(chainId *big.Int, blockProofs [][]byte) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.UpdateValidatorSetUsingEpochBlocks(&_RelayHubBeta.TransactOpts, chainId, blockProofs)
}

// UpdateValidatorSetUsingEpochBlocks is a paid mutator transaction binding the contract method 0x383d0ad9.
//
// Solidity: function updateValidatorSetUsingEpochBlocks(uint256 chainId, bytes[] blockProofs) returns()
func (_RelayHubBeta *RelayHubBetaTransactorSession) UpdateValidatorSetUsingEpochBlocks(chainId *big.Int, blockProofs [][]byte) (*types.Transaction, error) {
	return _RelayHubBeta.Contract.UpdateValidatorSetUsingEpochBlocks(&_RelayHubBeta.TransactOpts, chainId, blockProofs)
}

// RelayHubBetaBridgeRegisteredIterator is returned from FilterBridgeRegistered and is used to iterate over the raw logs and unpacked data for BridgeRegistered events raised by the RelayHubBeta contract.
type RelayHubBetaBridgeRegisteredIterator struct {
	Event *RelayHubBetaBridgeRegistered // Event containing the contract specifics and raw log

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
func (it *RelayHubBetaBridgeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubBetaBridgeRegistered)
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
		it.Event = new(RelayHubBetaBridgeRegistered)
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
func (it *RelayHubBetaBridgeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubBetaBridgeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubBetaBridgeRegistered represents a BridgeRegistered event raised by the RelayHubBeta contract.
type RelayHubBetaBridgeRegistered struct {
	ChainId       *big.Int
	BridgeAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRegistered is a free log retrieval operation binding the contract event 0xb72d37f67f55309a7342120ad674525186207bb1835dd860c16280472ea6fc50.
//
// Solidity: event BridgeRegistered(uint256 indexed chainId, address indexed bridgeAddress)
func (_RelayHubBeta *RelayHubBetaFilterer) FilterBridgeRegistered(opts *bind.FilterOpts, chainId []*big.Int, bridgeAddress []common.Address) (*RelayHubBetaBridgeRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var bridgeAddressRule []interface{}
	for _, bridgeAddressItem := range bridgeAddress {
		bridgeAddressRule = append(bridgeAddressRule, bridgeAddressItem)
	}

	logs, sub, err := _RelayHubBeta.contract.FilterLogs(opts, "BridgeRegistered", chainIdRule, bridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubBetaBridgeRegisteredIterator{contract: _RelayHubBeta.contract, event: "BridgeRegistered", logs: logs, sub: sub}, nil
}

// WatchBridgeRegistered is a free log subscription operation binding the contract event 0xb72d37f67f55309a7342120ad674525186207bb1835dd860c16280472ea6fc50.
//
// Solidity: event BridgeRegistered(uint256 indexed chainId, address indexed bridgeAddress)
func (_RelayHubBeta *RelayHubBetaFilterer) WatchBridgeRegistered(opts *bind.WatchOpts, sink chan<- *RelayHubBetaBridgeRegistered, chainId []*big.Int, bridgeAddress []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var bridgeAddressRule []interface{}
	for _, bridgeAddressItem := range bridgeAddress {
		bridgeAddressRule = append(bridgeAddressRule, bridgeAddressItem)
	}

	logs, sub, err := _RelayHubBeta.contract.WatchLogs(opts, "BridgeRegistered", chainIdRule, bridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubBetaBridgeRegistered)
				if err := _RelayHubBeta.contract.UnpackLog(event, "BridgeRegistered", log); err != nil {
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
func (_RelayHubBeta *RelayHubBetaFilterer) ParseBridgeRegistered(log types.Log) (*RelayHubBetaBridgeRegistered, error) {
	event := new(RelayHubBetaBridgeRegistered)
	if err := _RelayHubBeta.contract.UnpackLog(event, "BridgeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayHubBetaBridgeUnregisteredIterator is returned from FilterBridgeUnregistered and is used to iterate over the raw logs and unpacked data for BridgeUnregistered events raised by the RelayHubBeta contract.
type RelayHubBetaBridgeUnregisteredIterator struct {
	Event *RelayHubBetaBridgeUnregistered // Event containing the contract specifics and raw log

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
func (it *RelayHubBetaBridgeUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubBetaBridgeUnregistered)
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
		it.Event = new(RelayHubBetaBridgeUnregistered)
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
func (it *RelayHubBetaBridgeUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubBetaBridgeUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubBetaBridgeUnregistered represents a BridgeUnregistered event raised by the RelayHubBeta contract.
type RelayHubBetaBridgeUnregistered struct {
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeUnregistered is a free log retrieval operation binding the contract event 0x5ed0f717c1d62eb20db7f8f67a7f08a656c882be0834c2975e4991b546faac87.
//
// Solidity: event BridgeUnregistered(uint256 indexed chainId)
func (_RelayHubBeta *RelayHubBetaFilterer) FilterBridgeUnregistered(opts *bind.FilterOpts, chainId []*big.Int) (*RelayHubBetaBridgeUnregisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHubBeta.contract.FilterLogs(opts, "BridgeUnregistered", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubBetaBridgeUnregisteredIterator{contract: _RelayHubBeta.contract, event: "BridgeUnregistered", logs: logs, sub: sub}, nil
}

// WatchBridgeUnregistered is a free log subscription operation binding the contract event 0x5ed0f717c1d62eb20db7f8f67a7f08a656c882be0834c2975e4991b546faac87.
//
// Solidity: event BridgeUnregistered(uint256 indexed chainId)
func (_RelayHubBeta *RelayHubBetaFilterer) WatchBridgeUnregistered(opts *bind.WatchOpts, sink chan<- *RelayHubBetaBridgeUnregistered, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHubBeta.contract.WatchLogs(opts, "BridgeUnregistered", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubBetaBridgeUnregistered)
				if err := _RelayHubBeta.contract.UnpackLog(event, "BridgeUnregistered", log); err != nil {
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
func (_RelayHubBeta *RelayHubBetaFilterer) ParseBridgeUnregistered(log types.Log) (*RelayHubBetaBridgeUnregistered, error) {
	event := new(RelayHubBetaBridgeUnregistered)
	if err := _RelayHubBeta.contract.UnpackLog(event, "BridgeUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayHubBetaChainResetedIterator is returned from FilterChainReseted and is used to iterate over the raw logs and unpacked data for ChainReseted events raised by the RelayHubBeta contract.
type RelayHubBetaChainResetedIterator struct {
	Event *RelayHubBetaChainReseted // Event containing the contract specifics and raw log

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
func (it *RelayHubBetaChainResetedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubBetaChainReseted)
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
		it.Event = new(RelayHubBetaChainReseted)
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
func (it *RelayHubBetaChainResetedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubBetaChainResetedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubBetaChainReseted represents a ChainReseted event raised by the RelayHubBeta contract.
type RelayHubBetaChainReseted struct {
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainReseted is a free log retrieval operation binding the contract event 0x1b3e6308933a8c83e369ce09eb5dc180027a75240e78b85b9832061d9f1716ed.
//
// Solidity: event ChainReseted(uint256 indexed chainId)
func (_RelayHubBeta *RelayHubBetaFilterer) FilterChainReseted(opts *bind.FilterOpts, chainId []*big.Int) (*RelayHubBetaChainResetedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHubBeta.contract.FilterLogs(opts, "ChainReseted", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubBetaChainResetedIterator{contract: _RelayHubBeta.contract, event: "ChainReseted", logs: logs, sub: sub}, nil
}

// WatchChainReseted is a free log subscription operation binding the contract event 0x1b3e6308933a8c83e369ce09eb5dc180027a75240e78b85b9832061d9f1716ed.
//
// Solidity: event ChainReseted(uint256 indexed chainId)
func (_RelayHubBeta *RelayHubBetaFilterer) WatchChainReseted(opts *bind.WatchOpts, sink chan<- *RelayHubBetaChainReseted, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHubBeta.contract.WatchLogs(opts, "ChainReseted", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubBetaChainReseted)
				if err := _RelayHubBeta.contract.UnpackLog(event, "ChainReseted", log); err != nil {
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
func (_RelayHubBeta *RelayHubBetaFilterer) ParseChainReseted(log types.Log) (*RelayHubBetaChainReseted, error) {
	event := new(RelayHubBetaChainReseted)
	if err := _RelayHubBeta.contract.UnpackLog(event, "ChainReseted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayHubBetaValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the RelayHubBeta contract.
type RelayHubBetaValidatorSetUpdatedIterator struct {
	Event *RelayHubBetaValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *RelayHubBetaValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubBetaValidatorSetUpdated)
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
		it.Event = new(RelayHubBetaValidatorSetUpdated)
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
func (it *RelayHubBetaValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubBetaValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubBetaValidatorSetUpdated represents a ValidatorSetUpdated event raised by the RelayHubBeta contract.
type RelayHubBetaValidatorSetUpdated struct {
	ChainId      *big.Int
	ValidatorSet []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed chainId, address[] validatorSet)
func (_RelayHubBeta *RelayHubBetaFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*RelayHubBetaValidatorSetUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHubBeta.contract.FilterLogs(opts, "ValidatorSetUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubBetaValidatorSetUpdatedIterator{contract: _RelayHubBeta.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed chainId, address[] validatorSet)
func (_RelayHubBeta *RelayHubBetaFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *RelayHubBetaValidatorSetUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHubBeta.contract.WatchLogs(opts, "ValidatorSetUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubBetaValidatorSetUpdated)
				if err := _RelayHubBeta.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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
func (_RelayHubBeta *RelayHubBetaFilterer) ParseValidatorSetUpdated(log types.Log) (*RelayHubBetaValidatorSetUpdated, error) {
	event := new(RelayHubBetaValidatorSetUpdated)
	if err := _RelayHubBeta.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
