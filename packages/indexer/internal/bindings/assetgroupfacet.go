// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.ConvertType
)

// AssetGroup is an auto generated low-level Go binding around an user-defined struct.
type AssetGroup struct {
	ParentTokenId *big.Int
	Name          string
	MaxUnits      *big.Int
	UnitCount     *big.Int
	NextUnitIndex *big.Int
	Exists        bool
}

// IAssetGroupCreateGroupParams is an auto generated low-level Go binding around an user-defined struct.
type IAssetGroupCreateGroupParams struct {
	Name          string
	ParentTokenId *big.Int
	MaxUnits      *big.Int
}

// IAssetGroupMintUnitParams is an auto generated low-level Go binding around an user-defined struct.
type IAssetGroupMintUnitParams struct {
	GroupId   *big.Int
	Name      string
	Symbol    string
	Uri       string
	SupplyCap *big.Int
	Investor  common.Address
	Amount    *big.Int
}

// AssetGroupFacetMetaData contains all meta data concerning the AssetGroupFacet contract.
var AssetGroupFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"createGroup\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAssetGroup.CreateGroupParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"parentTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUnits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getChildGroup\",\"inputs\":[{\"name\":\"childTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGroup\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structAssetGroup\",\"components\":[{\"name\":\"parentTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxUnits\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unitCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nextUnitIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGroupChildren\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredGroupIds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"groupExists\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintUnit\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAssetGroup.MintUnitParams\",\"components\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"supplyCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"investor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"childTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintUnitBatch\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAssetGroup.MintUnitParams[]\",\"components\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"supplyCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"investor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"childTokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"GroupCreated\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"parentTokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"maxUnits\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferSingle\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnitMinted\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"childTokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"investor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AssetGroupFacet__ChildTokenIdCollision\",\"inputs\":[{\"name\":\"childTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AssetGroupFacet__EmptyBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AssetGroupFacet__EmptyName\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AssetGroupFacet__GroupNotFound\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AssetGroupFacet__MaxUnitsReached\",\"inputs\":[{\"name\":\"groupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUnits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AssetGroupFacet__MintToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AssetGroupFacet__ParentNotRegistered\",\"inputs\":[{\"name\":\"parentTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AssetGroupFacet__ProtocolPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AssetGroupFacet__ReceiverFrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AssetGroupFacet__Unauthorized\",\"inputs\":[]}]",
}

// AssetGroupFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetGroupFacetMetaData.ABI instead.
var AssetGroupFacetABI = AssetGroupFacetMetaData.ABI

// AssetGroupFacet is an auto generated Go binding around an Ethereum contract.
type AssetGroupFacet struct {
	AssetGroupFacetCaller     // Read-only binding to the contract
	AssetGroupFacetTransactor // Write-only binding to the contract
	AssetGroupFacetFilterer   // Log filterer for contract events
}

// AssetGroupFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetGroupFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetGroupFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetGroupFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetGroupFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetGroupFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetGroupFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetGroupFacetSession struct {
	Contract     *AssetGroupFacet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetGroupFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetGroupFacetCallerSession struct {
	Contract *AssetGroupFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AssetGroupFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetGroupFacetTransactorSession struct {
	Contract     *AssetGroupFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AssetGroupFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetGroupFacetRaw struct {
	Contract *AssetGroupFacet // Generic contract binding to access the raw methods on
}

// AssetGroupFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetGroupFacetCallerRaw struct {
	Contract *AssetGroupFacetCaller // Generic read-only contract binding to access the raw methods on
}

// AssetGroupFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetGroupFacetTransactorRaw struct {
	Contract *AssetGroupFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetGroupFacet creates a new instance of AssetGroupFacet, bound to a specific deployed contract.
func NewAssetGroupFacet(address common.Address, backend bind.ContractBackend) (*AssetGroupFacet, error) {
	contract, err := bindAssetGroupFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetGroupFacet{AssetGroupFacetCaller: AssetGroupFacetCaller{contract: contract}, AssetGroupFacetTransactor: AssetGroupFacetTransactor{contract: contract}, AssetGroupFacetFilterer: AssetGroupFacetFilterer{contract: contract}}, nil
}

// NewAssetGroupFacetCaller creates a new read-only instance of AssetGroupFacet, bound to a specific deployed contract.
func NewAssetGroupFacetCaller(address common.Address, caller bind.ContractCaller) (*AssetGroupFacetCaller, error) {
	contract, err := bindAssetGroupFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetGroupFacetCaller{contract: contract}, nil
}

// NewAssetGroupFacetTransactor creates a new write-only instance of AssetGroupFacet, bound to a specific deployed contract.
func NewAssetGroupFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetGroupFacetTransactor, error) {
	contract, err := bindAssetGroupFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetGroupFacetTransactor{contract: contract}, nil
}

// NewAssetGroupFacetFilterer creates a new log filterer instance of AssetGroupFacet, bound to a specific deployed contract.
func NewAssetGroupFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetGroupFacetFilterer, error) {
	contract, err := bindAssetGroupFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetGroupFacetFilterer{contract: contract}, nil
}

// bindAssetGroupFacet binds a generic wrapper to an already deployed contract.
func bindAssetGroupFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssetGroupFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetGroupFacet *AssetGroupFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetGroupFacet.Contract.AssetGroupFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetGroupFacet *AssetGroupFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetGroupFacet.Contract.AssetGroupFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetGroupFacet *AssetGroupFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetGroupFacet.Contract.AssetGroupFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetGroupFacet *AssetGroupFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetGroupFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetGroupFacet *AssetGroupFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetGroupFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetGroupFacet *AssetGroupFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetGroupFacet.Contract.contract.Transact(opts, method, params...)
}

// GetChildGroup is a free data retrieval call binding the contract method 0x4f1a3894.
//
// Solidity: function getChildGroup(uint256 childTokenId) view returns(uint256)
func (_AssetGroupFacet *AssetGroupFacetCaller) GetChildGroup(opts *bind.CallOpts, childTokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssetGroupFacet.contract.Call(opts, &out, "getChildGroup", childTokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChildGroup is a free data retrieval call binding the contract method 0x4f1a3894.
//
// Solidity: function getChildGroup(uint256 childTokenId) view returns(uint256)
func (_AssetGroupFacet *AssetGroupFacetSession) GetChildGroup(childTokenId *big.Int) (*big.Int, error) {
	return _AssetGroupFacet.Contract.GetChildGroup(&_AssetGroupFacet.CallOpts, childTokenId)
}

// GetChildGroup is a free data retrieval call binding the contract method 0x4f1a3894.
//
// Solidity: function getChildGroup(uint256 childTokenId) view returns(uint256)
func (_AssetGroupFacet *AssetGroupFacetCallerSession) GetChildGroup(childTokenId *big.Int) (*big.Int, error) {
	return _AssetGroupFacet.Contract.GetChildGroup(&_AssetGroupFacet.CallOpts, childTokenId)
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 groupId) view returns((uint256,string,uint256,uint256,uint256,bool))
func (_AssetGroupFacet *AssetGroupFacetCaller) GetGroup(opts *bind.CallOpts, groupId *big.Int) (AssetGroup, error) {
	var out []interface{}
	err := _AssetGroupFacet.contract.Call(opts, &out, "getGroup", groupId)

	if err != nil {
		return *new(AssetGroup), err
	}

	out0 := *abi.ConvertType(out[0], new(AssetGroup)).(*AssetGroup)

	return out0, err

}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 groupId) view returns((uint256,string,uint256,uint256,uint256,bool))
func (_AssetGroupFacet *AssetGroupFacetSession) GetGroup(groupId *big.Int) (AssetGroup, error) {
	return _AssetGroupFacet.Contract.GetGroup(&_AssetGroupFacet.CallOpts, groupId)
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 groupId) view returns((uint256,string,uint256,uint256,uint256,bool))
func (_AssetGroupFacet *AssetGroupFacetCallerSession) GetGroup(groupId *big.Int) (AssetGroup, error) {
	return _AssetGroupFacet.Contract.GetGroup(&_AssetGroupFacet.CallOpts, groupId)
}

// GetGroupChildren is a free data retrieval call binding the contract method 0x10fc7992.
//
// Solidity: function getGroupChildren(uint256 groupId) view returns(uint256[])
func (_AssetGroupFacet *AssetGroupFacetCaller) GetGroupChildren(opts *bind.CallOpts, groupId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _AssetGroupFacet.contract.Call(opts, &out, "getGroupChildren", groupId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetGroupChildren is a free data retrieval call binding the contract method 0x10fc7992.
//
// Solidity: function getGroupChildren(uint256 groupId) view returns(uint256[])
func (_AssetGroupFacet *AssetGroupFacetSession) GetGroupChildren(groupId *big.Int) ([]*big.Int, error) {
	return _AssetGroupFacet.Contract.GetGroupChildren(&_AssetGroupFacet.CallOpts, groupId)
}

// GetGroupChildren is a free data retrieval call binding the contract method 0x10fc7992.
//
// Solidity: function getGroupChildren(uint256 groupId) view returns(uint256[])
func (_AssetGroupFacet *AssetGroupFacetCallerSession) GetGroupChildren(groupId *big.Int) ([]*big.Int, error) {
	return _AssetGroupFacet.Contract.GetGroupChildren(&_AssetGroupFacet.CallOpts, groupId)
}

// GetRegisteredGroupIds is a free data retrieval call binding the contract method 0x7dc51657.
//
// Solidity: function getRegisteredGroupIds() view returns(uint256[])
func (_AssetGroupFacet *AssetGroupFacetCaller) GetRegisteredGroupIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _AssetGroupFacet.contract.Call(opts, &out, "getRegisteredGroupIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRegisteredGroupIds is a free data retrieval call binding the contract method 0x7dc51657.
//
// Solidity: function getRegisteredGroupIds() view returns(uint256[])
func (_AssetGroupFacet *AssetGroupFacetSession) GetRegisteredGroupIds() ([]*big.Int, error) {
	return _AssetGroupFacet.Contract.GetRegisteredGroupIds(&_AssetGroupFacet.CallOpts)
}

// GetRegisteredGroupIds is a free data retrieval call binding the contract method 0x7dc51657.
//
// Solidity: function getRegisteredGroupIds() view returns(uint256[])
func (_AssetGroupFacet *AssetGroupFacetCallerSession) GetRegisteredGroupIds() ([]*big.Int, error) {
	return _AssetGroupFacet.Contract.GetRegisteredGroupIds(&_AssetGroupFacet.CallOpts)
}

// GroupExists is a free data retrieval call binding the contract method 0xbd5263d8.
//
// Solidity: function groupExists(uint256 groupId) view returns(bool)
func (_AssetGroupFacet *AssetGroupFacetCaller) GroupExists(opts *bind.CallOpts, groupId *big.Int) (bool, error) {
	var out []interface{}
	err := _AssetGroupFacet.contract.Call(opts, &out, "groupExists", groupId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GroupExists is a free data retrieval call binding the contract method 0xbd5263d8.
//
// Solidity: function groupExists(uint256 groupId) view returns(bool)
func (_AssetGroupFacet *AssetGroupFacetSession) GroupExists(groupId *big.Int) (bool, error) {
	return _AssetGroupFacet.Contract.GroupExists(&_AssetGroupFacet.CallOpts, groupId)
}

// GroupExists is a free data retrieval call binding the contract method 0xbd5263d8.
//
// Solidity: function groupExists(uint256 groupId) view returns(bool)
func (_AssetGroupFacet *AssetGroupFacetCallerSession) GroupExists(groupId *big.Int) (bool, error) {
	return _AssetGroupFacet.Contract.GroupExists(&_AssetGroupFacet.CallOpts, groupId)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x320df3d7.
//
// Solidity: function createGroup((string,uint256,uint256) params) returns(uint256 groupId)
func (_AssetGroupFacet *AssetGroupFacetTransactor) CreateGroup(opts *bind.TransactOpts, params IAssetGroupCreateGroupParams) (*types.Transaction, error) {
	return _AssetGroupFacet.contract.Transact(opts, "createGroup", params)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x320df3d7.
//
// Solidity: function createGroup((string,uint256,uint256) params) returns(uint256 groupId)
func (_AssetGroupFacet *AssetGroupFacetSession) CreateGroup(params IAssetGroupCreateGroupParams) (*types.Transaction, error) {
	return _AssetGroupFacet.Contract.CreateGroup(&_AssetGroupFacet.TransactOpts, params)
}

// CreateGroup is a paid mutator transaction binding the contract method 0x320df3d7.
//
// Solidity: function createGroup((string,uint256,uint256) params) returns(uint256 groupId)
func (_AssetGroupFacet *AssetGroupFacetTransactorSession) CreateGroup(params IAssetGroupCreateGroupParams) (*types.Transaction, error) {
	return _AssetGroupFacet.Contract.CreateGroup(&_AssetGroupFacet.TransactOpts, params)
}

// MintUnit is a paid mutator transaction binding the contract method 0x3dea3338.
//
// Solidity: function mintUnit((uint256,string,string,string,uint256,address,uint256) params) returns(uint256 childTokenId)
func (_AssetGroupFacet *AssetGroupFacetTransactor) MintUnit(opts *bind.TransactOpts, params IAssetGroupMintUnitParams) (*types.Transaction, error) {
	return _AssetGroupFacet.contract.Transact(opts, "mintUnit", params)
}

// MintUnit is a paid mutator transaction binding the contract method 0x3dea3338.
//
// Solidity: function mintUnit((uint256,string,string,string,uint256,address,uint256) params) returns(uint256 childTokenId)
func (_AssetGroupFacet *AssetGroupFacetSession) MintUnit(params IAssetGroupMintUnitParams) (*types.Transaction, error) {
	return _AssetGroupFacet.Contract.MintUnit(&_AssetGroupFacet.TransactOpts, params)
}

// MintUnit is a paid mutator transaction binding the contract method 0x3dea3338.
//
// Solidity: function mintUnit((uint256,string,string,string,uint256,address,uint256) params) returns(uint256 childTokenId)
func (_AssetGroupFacet *AssetGroupFacetTransactorSession) MintUnit(params IAssetGroupMintUnitParams) (*types.Transaction, error) {
	return _AssetGroupFacet.Contract.MintUnit(&_AssetGroupFacet.TransactOpts, params)
}

// MintUnitBatch is a paid mutator transaction binding the contract method 0x524bac33.
//
// Solidity: function mintUnitBatch((uint256,string,string,string,uint256,address,uint256)[] params) returns(uint256[] childTokenIds)
func (_AssetGroupFacet *AssetGroupFacetTransactor) MintUnitBatch(opts *bind.TransactOpts, params []IAssetGroupMintUnitParams) (*types.Transaction, error) {
	return _AssetGroupFacet.contract.Transact(opts, "mintUnitBatch", params)
}

// MintUnitBatch is a paid mutator transaction binding the contract method 0x524bac33.
//
// Solidity: function mintUnitBatch((uint256,string,string,string,uint256,address,uint256)[] params) returns(uint256[] childTokenIds)
func (_AssetGroupFacet *AssetGroupFacetSession) MintUnitBatch(params []IAssetGroupMintUnitParams) (*types.Transaction, error) {
	return _AssetGroupFacet.Contract.MintUnitBatch(&_AssetGroupFacet.TransactOpts, params)
}

// MintUnitBatch is a paid mutator transaction binding the contract method 0x524bac33.
//
// Solidity: function mintUnitBatch((uint256,string,string,string,uint256,address,uint256)[] params) returns(uint256[] childTokenIds)
func (_AssetGroupFacet *AssetGroupFacetTransactorSession) MintUnitBatch(params []IAssetGroupMintUnitParams) (*types.Transaction, error) {
	return _AssetGroupFacet.Contract.MintUnitBatch(&_AssetGroupFacet.TransactOpts, params)
}

// AssetGroupFacetGroupCreatedIterator is returned from FilterGroupCreated and is used to iterate over the raw logs and unpacked data for GroupCreated events raised by the AssetGroupFacet contract.
type AssetGroupFacetGroupCreatedIterator struct {
	Event *AssetGroupFacetGroupCreated // Event containing the contract specifics and raw log

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
func (it *AssetGroupFacetGroupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetGroupFacetGroupCreated)
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
		it.Event = new(AssetGroupFacetGroupCreated)
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
func (it *AssetGroupFacetGroupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetGroupFacetGroupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetGroupFacetGroupCreated represents a GroupCreated event raised by the AssetGroupFacet contract.
type AssetGroupFacetGroupCreated struct {
	GroupId       *big.Int
	ParentTokenId *big.Int
	Name          string
	MaxUnits      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGroupCreated is a free log retrieval operation binding the contract event 0x535283d48414ecbe62a0d4913d8fcda1b2d27bf95c663f0ae62b2117488bfe11.
//
// Solidity: event GroupCreated(uint256 indexed groupId, uint256 indexed parentTokenId, string name, uint256 maxUnits)
func (_AssetGroupFacet *AssetGroupFacetFilterer) FilterGroupCreated(opts *bind.FilterOpts, groupId []*big.Int, parentTokenId []*big.Int) (*AssetGroupFacetGroupCreatedIterator, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var parentTokenIdRule []interface{}
	for _, parentTokenIdItem := range parentTokenId {
		parentTokenIdRule = append(parentTokenIdRule, parentTokenIdItem)
	}

	logs, sub, err := _AssetGroupFacet.contract.FilterLogs(opts, "GroupCreated", groupIdRule, parentTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetGroupFacetGroupCreatedIterator{contract: _AssetGroupFacet.contract, event: "GroupCreated", logs: logs, sub: sub}, nil
}

// WatchGroupCreated is a free log subscription operation binding the contract event 0x535283d48414ecbe62a0d4913d8fcda1b2d27bf95c663f0ae62b2117488bfe11.
//
// Solidity: event GroupCreated(uint256 indexed groupId, uint256 indexed parentTokenId, string name, uint256 maxUnits)
func (_AssetGroupFacet *AssetGroupFacetFilterer) WatchGroupCreated(opts *bind.WatchOpts, sink chan<- *AssetGroupFacetGroupCreated, groupId []*big.Int, parentTokenId []*big.Int) (event.Subscription, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var parentTokenIdRule []interface{}
	for _, parentTokenIdItem := range parentTokenId {
		parentTokenIdRule = append(parentTokenIdRule, parentTokenIdItem)
	}

	logs, sub, err := _AssetGroupFacet.contract.WatchLogs(opts, "GroupCreated", groupIdRule, parentTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetGroupFacetGroupCreated)
				if err := _AssetGroupFacet.contract.UnpackLog(event, "GroupCreated", log); err != nil {
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

// ParseGroupCreated is a log parse operation binding the contract event 0x535283d48414ecbe62a0d4913d8fcda1b2d27bf95c663f0ae62b2117488bfe11.
//
// Solidity: event GroupCreated(uint256 indexed groupId, uint256 indexed parentTokenId, string name, uint256 maxUnits)
func (_AssetGroupFacet *AssetGroupFacetFilterer) ParseGroupCreated(log types.Log) (*AssetGroupFacetGroupCreated, error) {
	event := new(AssetGroupFacetGroupCreated)
	if err := _AssetGroupFacet.contract.UnpackLog(event, "GroupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetGroupFacetTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the AssetGroupFacet contract.
type AssetGroupFacetTransferSingleIterator struct {
	Event *AssetGroupFacetTransferSingle // Event containing the contract specifics and raw log

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
func (it *AssetGroupFacetTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetGroupFacetTransferSingle)
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
		it.Event = new(AssetGroupFacetTransferSingle)
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
func (it *AssetGroupFacetTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetGroupFacetTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetGroupFacetTransferSingle represents a TransferSingle event raised by the AssetGroupFacet contract.
type AssetGroupFacetTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_AssetGroupFacet *AssetGroupFacetFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*AssetGroupFacetTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AssetGroupFacet.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AssetGroupFacetTransferSingleIterator{contract: _AssetGroupFacet.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_AssetGroupFacet *AssetGroupFacetFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *AssetGroupFacetTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AssetGroupFacet.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetGroupFacetTransferSingle)
				if err := _AssetGroupFacet.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_AssetGroupFacet *AssetGroupFacetFilterer) ParseTransferSingle(log types.Log) (*AssetGroupFacetTransferSingle, error) {
	event := new(AssetGroupFacetTransferSingle)
	if err := _AssetGroupFacet.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetGroupFacetUnitMintedIterator is returned from FilterUnitMinted and is used to iterate over the raw logs and unpacked data for UnitMinted events raised by the AssetGroupFacet contract.
type AssetGroupFacetUnitMintedIterator struct {
	Event *AssetGroupFacetUnitMinted // Event containing the contract specifics and raw log

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
func (it *AssetGroupFacetUnitMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetGroupFacetUnitMinted)
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
		it.Event = new(AssetGroupFacetUnitMinted)
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
func (it *AssetGroupFacetUnitMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetGroupFacetUnitMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetGroupFacetUnitMinted represents a UnitMinted event raised by the AssetGroupFacet contract.
type AssetGroupFacetUnitMinted struct {
	GroupId      *big.Int
	ChildTokenId *big.Int
	Investor     common.Address
	Name         string
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnitMinted is a free log retrieval operation binding the contract event 0x627584d9a00e2cf9c2dac2edb85bff33b81fbdabb77881e7cafcc72c2c95d7ec.
//
// Solidity: event UnitMinted(uint256 indexed groupId, uint256 indexed childTokenId, address indexed investor, string name, uint256 amount)
func (_AssetGroupFacet *AssetGroupFacetFilterer) FilterUnitMinted(opts *bind.FilterOpts, groupId []*big.Int, childTokenId []*big.Int, investor []common.Address) (*AssetGroupFacetUnitMintedIterator, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var childTokenIdRule []interface{}
	for _, childTokenIdItem := range childTokenId {
		childTokenIdRule = append(childTokenIdRule, childTokenIdItem)
	}
	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}

	logs, sub, err := _AssetGroupFacet.contract.FilterLogs(opts, "UnitMinted", groupIdRule, childTokenIdRule, investorRule)
	if err != nil {
		return nil, err
	}
	return &AssetGroupFacetUnitMintedIterator{contract: _AssetGroupFacet.contract, event: "UnitMinted", logs: logs, sub: sub}, nil
}

// WatchUnitMinted is a free log subscription operation binding the contract event 0x627584d9a00e2cf9c2dac2edb85bff33b81fbdabb77881e7cafcc72c2c95d7ec.
//
// Solidity: event UnitMinted(uint256 indexed groupId, uint256 indexed childTokenId, address indexed investor, string name, uint256 amount)
func (_AssetGroupFacet *AssetGroupFacetFilterer) WatchUnitMinted(opts *bind.WatchOpts, sink chan<- *AssetGroupFacetUnitMinted, groupId []*big.Int, childTokenId []*big.Int, investor []common.Address) (event.Subscription, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var childTokenIdRule []interface{}
	for _, childTokenIdItem := range childTokenId {
		childTokenIdRule = append(childTokenIdRule, childTokenIdItem)
	}
	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}

	logs, sub, err := _AssetGroupFacet.contract.WatchLogs(opts, "UnitMinted", groupIdRule, childTokenIdRule, investorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetGroupFacetUnitMinted)
				if err := _AssetGroupFacet.contract.UnpackLog(event, "UnitMinted", log); err != nil {
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

// ParseUnitMinted is a log parse operation binding the contract event 0x627584d9a00e2cf9c2dac2edb85bff33b81fbdabb77881e7cafcc72c2c95d7ec.
//
// Solidity: event UnitMinted(uint256 indexed groupId, uint256 indexed childTokenId, address indexed investor, string name, uint256 amount)
func (_AssetGroupFacet *AssetGroupFacetFilterer) ParseUnitMinted(log types.Log) (*AssetGroupFacetUnitMinted, error) {
	event := new(AssetGroupFacetUnitMinted)
	if err := _AssetGroupFacet.contract.UnpackLog(event, "UnitMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
