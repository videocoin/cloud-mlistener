// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package streamManager

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

// ManagerABI is the input ABI used to generate the binding from.
const ManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"},{\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"addInputChunkId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"refundAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"allowRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"v\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"revokeRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"v\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"endStream\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"createStream\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"},{\"name\":\"chunks\",\"type\":\"uint256[]\"}],\"name\":\"approveStreamCreation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"refund\",\"type\":\"bool\"},{\"name\":\"ended\",\"type\":\"bool\"},{\"name\":\"client\",\"type\":\"address\"},{\"name\":\"stream\",\"type\":\"address\"},{\"name\":\"streamId\",\"type\":\"uint256\"},{\"name\":\"RTMP\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"},{\"name\":\"RTMP\",\"type\":\"string\"},{\"name\":\"bitrates\",\"type\":\"uint256[]\"}],\"name\":\"requestStream\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"v\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"StreamRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"StreamApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"StreamCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"RefundAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"RefundRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"InputChunkAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"StreamEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Manager is an auto generated Go binding around an Ethereum contract.
type Manager struct {
	ManagerCaller     // Read-only binding to the contract
	ManagerTransactor // Write-only binding to the contract
	ManagerFilterer   // Log filterer for contract events
}

// ManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManagerSession struct {
	Contract     *Manager          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManagerCallerSession struct {
	Contract *ManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManagerTransactorSession struct {
	Contract     *ManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManagerRaw struct {
	Contract *Manager // Generic contract binding to access the raw methods on
}

// ManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManagerCallerRaw struct {
	Contract *ManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManagerTransactorRaw struct {
	Contract *ManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManager creates a new instance of Manager, bound to a specific deployed contract.
func NewManager(address common.Address, backend bind.ContractBackend) (*Manager, error) {
	contract, err := bindManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manager{ManagerCaller: ManagerCaller{contract: contract}, ManagerTransactor: ManagerTransactor{contract: contract}, ManagerFilterer: ManagerFilterer{contract: contract}}, nil
}

// NewManagerCaller creates a new read-only instance of Manager, bound to a specific deployed contract.
func NewManagerCaller(address common.Address, caller bind.ContractCaller) (*ManagerCaller, error) {
	contract, err := bindManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerCaller{contract: contract}, nil
}

// NewManagerTransactor creates a new write-only instance of Manager, bound to a specific deployed contract.
func NewManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ManagerTransactor, error) {
	contract, err := bindManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerTransactor{contract: contract}, nil
}

// NewManagerFilterer creates a new log filterer instance of Manager, bound to a specific deployed contract.
func NewManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ManagerFilterer, error) {
	contract, err := bindManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManagerFilterer{contract: contract}, nil
}

// bindManager binds a generic wrapper to an already deployed contract.
func bindManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.ManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Manager *ManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Manager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Manager *ManagerSession) IsOwner() (bool, error) {
	return _Manager.Contract.IsOwner(&_Manager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Manager *ManagerCallerSession) IsOwner() (bool, error) {
	return _Manager.Contract.IsOwner(&_Manager.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address v) constant returns(bool)
func (_Manager *ManagerCaller) IsValidator(opts *bind.CallOpts, v common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Manager.contract.Call(opts, out, "isValidator", v)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address v) constant returns(bool)
func (_Manager *ManagerSession) IsValidator(v common.Address) (bool, error) {
	return _Manager.Contract.IsValidator(&_Manager.CallOpts, v)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address v) constant returns(bool)
func (_Manager *ManagerCallerSession) IsValidator(v common.Address) (bool, error) {
	return _Manager.Contract.IsValidator(&_Manager.CallOpts, v)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Manager *ManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Manager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Manager *ManagerSession) Owner() (common.Address, error) {
	return _Manager.Contract.Owner(&_Manager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Manager *ManagerCallerSession) Owner() (common.Address, error) {
	return _Manager.Contract.Owner(&_Manager.CallOpts)
}

// RefundAllowed is a free data retrieval call binding the contract method 0x0f514717.
//
// Solidity: function refundAllowed(uint256 streamId) constant returns(bool)
func (_Manager *ManagerCaller) RefundAllowed(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Manager.contract.Call(opts, out, "refundAllowed", streamId)
	return *ret0, err
}

// RefundAllowed is a free data retrieval call binding the contract method 0x0f514717.
//
// Solidity: function refundAllowed(uint256 streamId) constant returns(bool)
func (_Manager *ManagerSession) RefundAllowed(streamId *big.Int) (bool, error) {
	return _Manager.Contract.RefundAllowed(&_Manager.CallOpts, streamId)
}

// RefundAllowed is a free data retrieval call binding the contract method 0x0f514717.
//
// Solidity: function refundAllowed(uint256 streamId) constant returns(bool)
func (_Manager *ManagerCallerSession) RefundAllowed(streamId *big.Int) (bool, error) {
	return _Manager.Contract.RefundAllowed(&_Manager.CallOpts, streamId)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) constant returns(bool approved, bool refund, bool ended, address client, address stream, uint256 streamId, string RTMP)
func (_Manager *ManagerCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Approved bool
	Refund   bool
	Ended    bool
	Client   common.Address
	Stream   common.Address
	StreamId *big.Int
	RTMP     string
}, error) {
	ret := new(struct {
		Approved bool
		Refund   bool
		Ended    bool
		Client   common.Address
		Stream   common.Address
		StreamId *big.Int
		RTMP     string
	})
	out := ret
	err := _Manager.contract.Call(opts, out, "requests", arg0)
	return *ret, err
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) constant returns(bool approved, bool refund, bool ended, address client, address stream, uint256 streamId, string RTMP)
func (_Manager *ManagerSession) Requests(arg0 *big.Int) (struct {
	Approved bool
	Refund   bool
	Ended    bool
	Client   common.Address
	Stream   common.Address
	StreamId *big.Int
	RTMP     string
}, error) {
	return _Manager.Contract.Requests(&_Manager.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) constant returns(bool approved, bool refund, bool ended, address client, address stream, uint256 streamId, string RTMP)
func (_Manager *ManagerCallerSession) Requests(arg0 *big.Int) (struct {
	Approved bool
	Refund   bool
	Ended    bool
	Client   common.Address
	Stream   common.Address
	StreamId *big.Int
	RTMP     string
}, error) {
	return _Manager.Contract.Requests(&_Manager.CallOpts, arg0)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x0a8c4c13.
//
// Solidity: function addInputChunkId(uint256 streamId, uint256 chunkId) returns()
func (_Manager *ManagerTransactor) AddInputChunkId(opts *bind.TransactOpts, streamId *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "addInputChunkId", streamId, chunkId)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x0a8c4c13.
//
// Solidity: function addInputChunkId(uint256 streamId, uint256 chunkId) returns()
func (_Manager *ManagerSession) AddInputChunkId(streamId *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.AddInputChunkId(&_Manager.TransactOpts, streamId, chunkId)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x0a8c4c13.
//
// Solidity: function addInputChunkId(uint256 streamId, uint256 chunkId) returns()
func (_Manager *ManagerTransactorSession) AddInputChunkId(streamId *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.AddInputChunkId(&_Manager.TransactOpts, streamId, chunkId)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address v) returns()
func (_Manager *ManagerTransactor) AddValidator(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "addValidator", v)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address v) returns()
func (_Manager *ManagerSession) AddValidator(v common.Address) (*types.Transaction, error) {
	return _Manager.Contract.AddValidator(&_Manager.TransactOpts, v)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address v) returns()
func (_Manager *ManagerTransactorSession) AddValidator(v common.Address) (*types.Transaction, error) {
	return _Manager.Contract.AddValidator(&_Manager.TransactOpts, v)
}

// AllowRefund is a paid mutator transaction binding the contract method 0x225f6541.
//
// Solidity: function allowRefund(uint256 streamId) returns()
func (_Manager *ManagerTransactor) AllowRefund(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "allowRefund", streamId)
}

// AllowRefund is a paid mutator transaction binding the contract method 0x225f6541.
//
// Solidity: function allowRefund(uint256 streamId) returns()
func (_Manager *ManagerSession) AllowRefund(streamId *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.AllowRefund(&_Manager.TransactOpts, streamId)
}

// AllowRefund is a paid mutator transaction binding the contract method 0x225f6541.
//
// Solidity: function allowRefund(uint256 streamId) returns()
func (_Manager *ManagerTransactorSession) AllowRefund(streamId *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.AllowRefund(&_Manager.TransactOpts, streamId)
}

// ApproveStreamCreation is a paid mutator transaction binding the contract method 0x70bdecb6.
//
// Solidity: function approveStreamCreation(uint256 streamId, uint256[] chunks) returns()
func (_Manager *ManagerTransactor) ApproveStreamCreation(opts *bind.TransactOpts, streamId *big.Int, chunks []*big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "approveStreamCreation", streamId, chunks)
}

// ApproveStreamCreation is a paid mutator transaction binding the contract method 0x70bdecb6.
//
// Solidity: function approveStreamCreation(uint256 streamId, uint256[] chunks) returns()
func (_Manager *ManagerSession) ApproveStreamCreation(streamId *big.Int, chunks []*big.Int) (*types.Transaction, error) {
	return _Manager.Contract.ApproveStreamCreation(&_Manager.TransactOpts, streamId, chunks)
}

// ApproveStreamCreation is a paid mutator transaction binding the contract method 0x70bdecb6.
//
// Solidity: function approveStreamCreation(uint256 streamId, uint256[] chunks) returns()
func (_Manager *ManagerTransactorSession) ApproveStreamCreation(streamId *big.Int, chunks []*big.Int) (*types.Transaction, error) {
	return _Manager.Contract.ApproveStreamCreation(&_Manager.TransactOpts, streamId, chunks)
}

// CreateStream is a paid mutator transaction binding the contract method 0x551479dd.
//
// Solidity: function createStream(uint256 streamId) returns(address)
func (_Manager *ManagerTransactor) CreateStream(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "createStream", streamId)
}

// CreateStream is a paid mutator transaction binding the contract method 0x551479dd.
//
// Solidity: function createStream(uint256 streamId) returns(address)
func (_Manager *ManagerSession) CreateStream(streamId *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.CreateStream(&_Manager.TransactOpts, streamId)
}

// CreateStream is a paid mutator transaction binding the contract method 0x551479dd.
//
// Solidity: function createStream(uint256 streamId) returns(address)
func (_Manager *ManagerTransactorSession) CreateStream(streamId *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.CreateStream(&_Manager.TransactOpts, streamId)
}

// EndStream is a paid mutator transaction binding the contract method 0x50d55afc.
//
// Solidity: function endStream(uint256 streamId) returns()
func (_Manager *ManagerTransactor) EndStream(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "endStream", streamId)
}

// EndStream is a paid mutator transaction binding the contract method 0x50d55afc.
//
// Solidity: function endStream(uint256 streamId) returns()
func (_Manager *ManagerSession) EndStream(streamId *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.EndStream(&_Manager.TransactOpts, streamId)
}

// EndStream is a paid mutator transaction binding the contract method 0x50d55afc.
//
// Solidity: function endStream(uint256 streamId) returns()
func (_Manager *ManagerTransactorSession) EndStream(streamId *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.EndStream(&_Manager.TransactOpts, streamId)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address v) returns()
func (_Manager *ManagerTransactor) RemoveValidator(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "removeValidator", v)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address v) returns()
func (_Manager *ManagerSession) RemoveValidator(v common.Address) (*types.Transaction, error) {
	return _Manager.Contract.RemoveValidator(&_Manager.TransactOpts, v)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address v) returns()
func (_Manager *ManagerTransactorSession) RemoveValidator(v common.Address) (*types.Transaction, error) {
	return _Manager.Contract.RemoveValidator(&_Manager.TransactOpts, v)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manager *ManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manager *ManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Manager.Contract.RenounceOwnership(&_Manager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manager *ManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Manager.Contract.RenounceOwnership(&_Manager.TransactOpts)
}

// RequestStream is a paid mutator transaction binding the contract method 0x8562faa2.
//
// Solidity: function requestStream(uint256 streamId, string RTMP, uint256[] bitrates) returns(uint256)
func (_Manager *ManagerTransactor) RequestStream(opts *bind.TransactOpts, streamId *big.Int, RTMP string, bitrates []*big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "requestStream", streamId, RTMP, bitrates)
}

// RequestStream is a paid mutator transaction binding the contract method 0x8562faa2.
//
// Solidity: function requestStream(uint256 streamId, string RTMP, uint256[] bitrates) returns(uint256)
func (_Manager *ManagerSession) RequestStream(streamId *big.Int, RTMP string, bitrates []*big.Int) (*types.Transaction, error) {
	return _Manager.Contract.RequestStream(&_Manager.TransactOpts, streamId, RTMP, bitrates)
}

// RequestStream is a paid mutator transaction binding the contract method 0x8562faa2.
//
// Solidity: function requestStream(uint256 streamId, string RTMP, uint256[] bitrates) returns(uint256)
func (_Manager *ManagerTransactorSession) RequestStream(streamId *big.Int, RTMP string, bitrates []*big.Int) (*types.Transaction, error) {
	return _Manager.Contract.RequestStream(&_Manager.TransactOpts, streamId, RTMP, bitrates)
}

// RevokeRefund is a paid mutator transaction binding the contract method 0x4b51438c.
//
// Solidity: function revokeRefund(uint256 streamId) returns()
func (_Manager *ManagerTransactor) RevokeRefund(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "revokeRefund", streamId)
}

// RevokeRefund is a paid mutator transaction binding the contract method 0x4b51438c.
//
// Solidity: function revokeRefund(uint256 streamId) returns()
func (_Manager *ManagerSession) RevokeRefund(streamId *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.RevokeRefund(&_Manager.TransactOpts, streamId)
}

// RevokeRefund is a paid mutator transaction binding the contract method 0x4b51438c.
//
// Solidity: function revokeRefund(uint256 streamId) returns()
func (_Manager *ManagerTransactorSession) RevokeRefund(streamId *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.RevokeRefund(&_Manager.TransactOpts, streamId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manager *ManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manager *ManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Manager.Contract.TransferOwnership(&_Manager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manager *ManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Manager.Contract.TransferOwnership(&_Manager.TransactOpts, newOwner)
}

// ManagerInputChunkAddedIterator is returned from FilterInputChunkAdded and is used to iterate over the raw logs and unpacked data for InputChunkAdded events raised by the Manager contract.
type ManagerInputChunkAddedIterator struct {
	Event *ManagerInputChunkAdded // Event containing the contract specifics and raw log

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
func (it *ManagerInputChunkAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerInputChunkAdded)
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
		it.Event = new(ManagerInputChunkAdded)
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
func (it *ManagerInputChunkAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerInputChunkAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerInputChunkAdded represents a InputChunkAdded event raised by the Manager contract.
type ManagerInputChunkAdded struct {
	StreamId *big.Int
	ChunkId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInputChunkAdded is a free log retrieval operation binding the contract event 0x79f21a91842ed307e9fade5e30c0d0322ee9f1f64d3918e3d8b2815f02ce85d6.
//
// Solidity: event InputChunkAdded(uint256 indexed streamId, uint256 indexed chunkId)
func (_Manager *ManagerFilterer) FilterInputChunkAdded(opts *bind.FilterOpts, streamId []*big.Int, chunkId []*big.Int) (*ManagerInputChunkAddedIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "InputChunkAdded", streamIdRule, chunkIdRule)
	if err != nil {
		return nil, err
	}
	return &ManagerInputChunkAddedIterator{contract: _Manager.contract, event: "InputChunkAdded", logs: logs, sub: sub}, nil
}

// WatchInputChunkAdded is a free log subscription operation binding the contract event 0x79f21a91842ed307e9fade5e30c0d0322ee9f1f64d3918e3d8b2815f02ce85d6.
//
// Solidity: event InputChunkAdded(uint256 indexed streamId, uint256 indexed chunkId)
func (_Manager *ManagerFilterer) WatchInputChunkAdded(opts *bind.WatchOpts, sink chan<- *ManagerInputChunkAdded, streamId []*big.Int, chunkId []*big.Int) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "InputChunkAdded", streamIdRule, chunkIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerInputChunkAdded)
				if err := _Manager.contract.UnpackLog(event, "InputChunkAdded", log); err != nil {
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

// ManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Manager contract.
type ManagerOwnershipTransferredIterator struct {
	Event *ManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerOwnershipTransferred)
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
		it.Event = new(ManagerOwnershipTransferred)
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
func (it *ManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerOwnershipTransferred represents a OwnershipTransferred event raised by the Manager contract.
type ManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Manager *ManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ManagerOwnershipTransferredIterator{contract: _Manager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Manager *ManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerOwnershipTransferred)
				if err := _Manager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ManagerRefundAllowedIterator is returned from FilterRefundAllowed and is used to iterate over the raw logs and unpacked data for RefundAllowed events raised by the Manager contract.
type ManagerRefundAllowedIterator struct {
	Event *ManagerRefundAllowed // Event containing the contract specifics and raw log

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
func (it *ManagerRefundAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerRefundAllowed)
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
		it.Event = new(ManagerRefundAllowed)
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
func (it *ManagerRefundAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerRefundAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerRefundAllowed represents a RefundAllowed event raised by the Manager contract.
type ManagerRefundAllowed struct {
	StreamId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefundAllowed is a free log retrieval operation binding the contract event 0x48030861b818deafd7def8853bcc8a6ec6bab746521a9546b79b8baab82dce6b.
//
// Solidity: event RefundAllowed(uint256 indexed streamId)
func (_Manager *ManagerFilterer) FilterRefundAllowed(opts *bind.FilterOpts, streamId []*big.Int) (*ManagerRefundAllowedIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "RefundAllowed", streamIdRule)
	if err != nil {
		return nil, err
	}
	return &ManagerRefundAllowedIterator{contract: _Manager.contract, event: "RefundAllowed", logs: logs, sub: sub}, nil
}

// WatchRefundAllowed is a free log subscription operation binding the contract event 0x48030861b818deafd7def8853bcc8a6ec6bab746521a9546b79b8baab82dce6b.
//
// Solidity: event RefundAllowed(uint256 indexed streamId)
func (_Manager *ManagerFilterer) WatchRefundAllowed(opts *bind.WatchOpts, sink chan<- *ManagerRefundAllowed, streamId []*big.Int) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "RefundAllowed", streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerRefundAllowed)
				if err := _Manager.contract.UnpackLog(event, "RefundAllowed", log); err != nil {
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

// ManagerRefundRevokedIterator is returned from FilterRefundRevoked and is used to iterate over the raw logs and unpacked data for RefundRevoked events raised by the Manager contract.
type ManagerRefundRevokedIterator struct {
	Event *ManagerRefundRevoked // Event containing the contract specifics and raw log

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
func (it *ManagerRefundRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerRefundRevoked)
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
		it.Event = new(ManagerRefundRevoked)
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
func (it *ManagerRefundRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerRefundRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerRefundRevoked represents a RefundRevoked event raised by the Manager contract.
type ManagerRefundRevoked struct {
	StreamId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefundRevoked is a free log retrieval operation binding the contract event 0xf12d3e50d9ed05d035c02b9f29e1da20004bce112d4fbe2cc0f0ba272ee5d977.
//
// Solidity: event RefundRevoked(uint256 indexed streamId)
func (_Manager *ManagerFilterer) FilterRefundRevoked(opts *bind.FilterOpts, streamId []*big.Int) (*ManagerRefundRevokedIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "RefundRevoked", streamIdRule)
	if err != nil {
		return nil, err
	}
	return &ManagerRefundRevokedIterator{contract: _Manager.contract, event: "RefundRevoked", logs: logs, sub: sub}, nil
}

// WatchRefundRevoked is a free log subscription operation binding the contract event 0xf12d3e50d9ed05d035c02b9f29e1da20004bce112d4fbe2cc0f0ba272ee5d977.
//
// Solidity: event RefundRevoked(uint256 indexed streamId)
func (_Manager *ManagerFilterer) WatchRefundRevoked(opts *bind.WatchOpts, sink chan<- *ManagerRefundRevoked, streamId []*big.Int) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "RefundRevoked", streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerRefundRevoked)
				if err := _Manager.contract.UnpackLog(event, "RefundRevoked", log); err != nil {
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

// ManagerStreamApprovedIterator is returned from FilterStreamApproved and is used to iterate over the raw logs and unpacked data for StreamApproved events raised by the Manager contract.
type ManagerStreamApprovedIterator struct {
	Event *ManagerStreamApproved // Event containing the contract specifics and raw log

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
func (it *ManagerStreamApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerStreamApproved)
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
		it.Event = new(ManagerStreamApproved)
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
func (it *ManagerStreamApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerStreamApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerStreamApproved represents a StreamApproved event raised by the Manager contract.
type ManagerStreamApproved struct {
	StreamId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStreamApproved is a free log retrieval operation binding the contract event 0x8fd81831397822207cf7571ac8b61f51ae7de628b05492a586bc63c3fd23bf8c.
//
// Solidity: event StreamApproved(uint256 indexed streamId)
func (_Manager *ManagerFilterer) FilterStreamApproved(opts *bind.FilterOpts, streamId []*big.Int) (*ManagerStreamApprovedIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "StreamApproved", streamIdRule)
	if err != nil {
		return nil, err
	}
	return &ManagerStreamApprovedIterator{contract: _Manager.contract, event: "StreamApproved", logs: logs, sub: sub}, nil
}

// WatchStreamApproved is a free log subscription operation binding the contract event 0x8fd81831397822207cf7571ac8b61f51ae7de628b05492a586bc63c3fd23bf8c.
//
// Solidity: event StreamApproved(uint256 indexed streamId)
func (_Manager *ManagerFilterer) WatchStreamApproved(opts *bind.WatchOpts, sink chan<- *ManagerStreamApproved, streamId []*big.Int) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "StreamApproved", streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerStreamApproved)
				if err := _Manager.contract.UnpackLog(event, "StreamApproved", log); err != nil {
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

// ManagerStreamCreatedIterator is returned from FilterStreamCreated and is used to iterate over the raw logs and unpacked data for StreamCreated events raised by the Manager contract.
type ManagerStreamCreatedIterator struct {
	Event *ManagerStreamCreated // Event containing the contract specifics and raw log

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
func (it *ManagerStreamCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerStreamCreated)
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
		it.Event = new(ManagerStreamCreated)
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
func (it *ManagerStreamCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerStreamCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerStreamCreated represents a StreamCreated event raised by the Manager contract.
type ManagerStreamCreated struct {
	StreamAddress common.Address
	StreamId      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStreamCreated is a free log retrieval operation binding the contract event 0x1bd63527042f119292b792487cee2f3e2f788737aa8ce9c0b5e79a2e17bd6bab.
//
// Solidity: event StreamCreated(address indexed streamAddress, uint256 indexed streamId)
func (_Manager *ManagerFilterer) FilterStreamCreated(opts *bind.FilterOpts, streamAddress []common.Address, streamId []*big.Int) (*ManagerStreamCreatedIterator, error) {

	var streamAddressRule []interface{}
	for _, streamAddressItem := range streamAddress {
		streamAddressRule = append(streamAddressRule, streamAddressItem)
	}
	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "StreamCreated", streamAddressRule, streamIdRule)
	if err != nil {
		return nil, err
	}
	return &ManagerStreamCreatedIterator{contract: _Manager.contract, event: "StreamCreated", logs: logs, sub: sub}, nil
}

// WatchStreamCreated is a free log subscription operation binding the contract event 0x1bd63527042f119292b792487cee2f3e2f788737aa8ce9c0b5e79a2e17bd6bab.
//
// Solidity: event StreamCreated(address indexed streamAddress, uint256 indexed streamId)
func (_Manager *ManagerFilterer) WatchStreamCreated(opts *bind.WatchOpts, sink chan<- *ManagerStreamCreated, streamAddress []common.Address, streamId []*big.Int) (event.Subscription, error) {

	var streamAddressRule []interface{}
	for _, streamAddressItem := range streamAddress {
		streamAddressRule = append(streamAddressRule, streamAddressItem)
	}
	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "StreamCreated", streamAddressRule, streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerStreamCreated)
				if err := _Manager.contract.UnpackLog(event, "StreamCreated", log); err != nil {
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

// ManagerStreamEndedIterator is returned from FilterStreamEnded and is used to iterate over the raw logs and unpacked data for StreamEnded events raised by the Manager contract.
type ManagerStreamEndedIterator struct {
	Event *ManagerStreamEnded // Event containing the contract specifics and raw log

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
func (it *ManagerStreamEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerStreamEnded)
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
		it.Event = new(ManagerStreamEnded)
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
func (it *ManagerStreamEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerStreamEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerStreamEnded represents a StreamEnded event raised by the Manager contract.
type ManagerStreamEnded struct {
	StreamId *big.Int
	Caller   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStreamEnded is a free log retrieval operation binding the contract event 0xde05e689a2a03aa3267e2f457184c649b080aaa00013ed27d21b85bcb04901ff.
//
// Solidity: event StreamEnded(uint256 indexed streamId, address indexed caller)
func (_Manager *ManagerFilterer) FilterStreamEnded(opts *bind.FilterOpts, streamId []*big.Int, caller []common.Address) (*ManagerStreamEndedIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "StreamEnded", streamIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &ManagerStreamEndedIterator{contract: _Manager.contract, event: "StreamEnded", logs: logs, sub: sub}, nil
}

// WatchStreamEnded is a free log subscription operation binding the contract event 0xde05e689a2a03aa3267e2f457184c649b080aaa00013ed27d21b85bcb04901ff.
//
// Solidity: event StreamEnded(uint256 indexed streamId, address indexed caller)
func (_Manager *ManagerFilterer) WatchStreamEnded(opts *bind.WatchOpts, sink chan<- *ManagerStreamEnded, streamId []*big.Int, caller []common.Address) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "StreamEnded", streamIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerStreamEnded)
				if err := _Manager.contract.UnpackLog(event, "StreamEnded", log); err != nil {
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

// ManagerStreamRequestedIterator is returned from FilterStreamRequested and is used to iterate over the raw logs and unpacked data for StreamRequested events raised by the Manager contract.
type ManagerStreamRequestedIterator struct {
	Event *ManagerStreamRequested // Event containing the contract specifics and raw log

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
func (it *ManagerStreamRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerStreamRequested)
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
		it.Event = new(ManagerStreamRequested)
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
func (it *ManagerStreamRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerStreamRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerStreamRequested represents a StreamRequested event raised by the Manager contract.
type ManagerStreamRequested struct {
	Client   common.Address
	StreamId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStreamRequested is a free log retrieval operation binding the contract event 0xcf93cb8f3f726dd083e429df912bf338cc7f82cc3344b3ab5fa960ef0357e321.
//
// Solidity: event StreamRequested(address indexed client, uint256 indexed streamId)
func (_Manager *ManagerFilterer) FilterStreamRequested(opts *bind.FilterOpts, client []common.Address, streamId []*big.Int) (*ManagerStreamRequestedIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "StreamRequested", clientRule, streamIdRule)
	if err != nil {
		return nil, err
	}
	return &ManagerStreamRequestedIterator{contract: _Manager.contract, event: "StreamRequested", logs: logs, sub: sub}, nil
}

// WatchStreamRequested is a free log subscription operation binding the contract event 0xcf93cb8f3f726dd083e429df912bf338cc7f82cc3344b3ab5fa960ef0357e321.
//
// Solidity: event StreamRequested(address indexed client, uint256 indexed streamId)
func (_Manager *ManagerFilterer) WatchStreamRequested(opts *bind.WatchOpts, sink chan<- *ManagerStreamRequested, client []common.Address, streamId []*big.Int) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "StreamRequested", clientRule, streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerStreamRequested)
				if err := _Manager.contract.UnpackLog(event, "StreamRequested", log); err != nil {
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

// ManagerValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the Manager contract.
type ManagerValidatorAddedIterator struct {
	Event *ManagerValidatorAdded // Event containing the contract specifics and raw log

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
func (it *ManagerValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerValidatorAdded)
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
		it.Event = new(ManagerValidatorAdded)
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
func (it *ManagerValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerValidatorAdded represents a ValidatorAdded event raised by the Manager contract.
type ManagerValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_Manager *ManagerFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validator []common.Address) (*ManagerValidatorAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ManagerValidatorAddedIterator{contract: _Manager.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_Manager *ManagerFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ManagerValidatorAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerValidatorAdded)
				if err := _Manager.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ManagerValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the Manager contract.
type ManagerValidatorRemovedIterator struct {
	Event *ManagerValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *ManagerValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerValidatorRemoved)
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
		it.Event = new(ManagerValidatorRemoved)
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
func (it *ManagerValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerValidatorRemoved represents a ValidatorRemoved event raised by the Manager contract.
type ManagerValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_Manager *ManagerFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validator []common.Address) (*ManagerValidatorRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ManagerValidatorRemovedIterator{contract: _Manager.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_Manager *ManagerFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ManagerValidatorRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerValidatorRemoved)
				if err := _Manager.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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
