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

// AssetConfig is an auto generated low-level Go binding around an user-defined struct.
type AssetConfig struct {
	Name              string
	Symbol            string
	Uri               string
	SupplyCap         *big.Int
	TotalSupply       *big.Int
	IdentityProfileId uint32
	ComplianceModules []common.Address
	Issuer            common.Address
	Paused            bool
	Exists            bool
	AllowedCountries  []uint16
}

// IAssetManagerRegisterAssetParams is an auto generated low-level Go binding around an user-defined struct.
type IAssetManagerRegisterAssetParams struct {
	Name              string
	Symbol            string
	Uri               string
	SupplyCap         *big.Int
	IdentityProfileId uint32
	ComplianceModules []common.Address
	Issuer            common.Address
	AllowedCountries  []uint16
}

// AssetManagerFacetMetaData contains all meta data concerning the AssetManagerFacet contract.
var AssetManagerFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addComplianceModule\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"assetExists\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAssetConfig\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structAssetConfig\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"supplyCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"identityProfileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"complianceModules\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"allowedCountries\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getComplianceModules\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredTokenIds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextTokenId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerAsset\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structIAssetManager.RegisterAssetParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"supplyCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"identityProfileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"complianceModules\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedCountries\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}]}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeComplianceModule\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedCountries\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"countries\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAssetUri\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setComplianceModules\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"modules\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIdentityProfile\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIssuer\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSupplyCap\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AssetConfigUpdated\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AssetRegistered\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"issuer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"profileId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceModuleAdded\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"module\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceModuleRemoved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"module\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceModulesSet\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"modules\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"URI\",\"inputs\":[{\"name\":\"value\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AssetManagerFacet__EmptyString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AssetManagerFacet__ModuleAlreadyAdded\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AssetManagerFacet__ModuleNotFound\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AssetManagerFacet__NotRegistered\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AssetManagerFacet__TooManyModules\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AssetManagerFacet__Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AssetManagerFacet__ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LibDiamond__OnlyOwner\",\"inputs\":[]}]",
}

// AssetManagerFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetManagerFacetMetaData.ABI instead.
var AssetManagerFacetABI = AssetManagerFacetMetaData.ABI

// AssetManagerFacet is an auto generated Go binding around an Ethereum contract.
type AssetManagerFacet struct {
	AssetManagerFacetCaller     // Read-only binding to the contract
	AssetManagerFacetTransactor // Write-only binding to the contract
	AssetManagerFacetFilterer   // Log filterer for contract events
}

// AssetManagerFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetManagerFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagerFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetManagerFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagerFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetManagerFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagerFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetManagerFacetSession struct {
	Contract     *AssetManagerFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AssetManagerFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetManagerFacetCallerSession struct {
	Contract *AssetManagerFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AssetManagerFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetManagerFacetTransactorSession struct {
	Contract     *AssetManagerFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AssetManagerFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetManagerFacetRaw struct {
	Contract *AssetManagerFacet // Generic contract binding to access the raw methods on
}

// AssetManagerFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetManagerFacetCallerRaw struct {
	Contract *AssetManagerFacetCaller // Generic read-only contract binding to access the raw methods on
}

// AssetManagerFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetManagerFacetTransactorRaw struct {
	Contract *AssetManagerFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetManagerFacet creates a new instance of AssetManagerFacet, bound to a specific deployed contract.
func NewAssetManagerFacet(address common.Address, backend bind.ContractBackend) (*AssetManagerFacet, error) {
	contract, err := bindAssetManagerFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFacet{AssetManagerFacetCaller: AssetManagerFacetCaller{contract: contract}, AssetManagerFacetTransactor: AssetManagerFacetTransactor{contract: contract}, AssetManagerFacetFilterer: AssetManagerFacetFilterer{contract: contract}}, nil
}

// NewAssetManagerFacetCaller creates a new read-only instance of AssetManagerFacet, bound to a specific deployed contract.
func NewAssetManagerFacetCaller(address common.Address, caller bind.ContractCaller) (*AssetManagerFacetCaller, error) {
	contract, err := bindAssetManagerFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFacetCaller{contract: contract}, nil
}

// NewAssetManagerFacetTransactor creates a new write-only instance of AssetManagerFacet, bound to a specific deployed contract.
func NewAssetManagerFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetManagerFacetTransactor, error) {
	contract, err := bindAssetManagerFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFacetTransactor{contract: contract}, nil
}

// NewAssetManagerFacetFilterer creates a new log filterer instance of AssetManagerFacet, bound to a specific deployed contract.
func NewAssetManagerFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetManagerFacetFilterer, error) {
	contract, err := bindAssetManagerFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFacetFilterer{contract: contract}, nil
}

// bindAssetManagerFacet binds a generic wrapper to an already deployed contract.
func bindAssetManagerFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssetManagerFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetManagerFacet *AssetManagerFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetManagerFacet.Contract.AssetManagerFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetManagerFacet *AssetManagerFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.AssetManagerFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetManagerFacet *AssetManagerFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.AssetManagerFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetManagerFacet *AssetManagerFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetManagerFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetManagerFacet *AssetManagerFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetManagerFacet *AssetManagerFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.contract.Transact(opts, method, params...)
}

// AssetExists is a free data retrieval call binding the contract method 0x518e2bdd.
//
// Solidity: function assetExists(uint256 tokenId) view returns(bool)
func (_AssetManagerFacet *AssetManagerFacetCaller) AssetExists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _AssetManagerFacet.contract.Call(opts, &out, "assetExists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetExists is a free data retrieval call binding the contract method 0x518e2bdd.
//
// Solidity: function assetExists(uint256 tokenId) view returns(bool)
func (_AssetManagerFacet *AssetManagerFacetSession) AssetExists(tokenId *big.Int) (bool, error) {
	return _AssetManagerFacet.Contract.AssetExists(&_AssetManagerFacet.CallOpts, tokenId)
}

// AssetExists is a free data retrieval call binding the contract method 0x518e2bdd.
//
// Solidity: function assetExists(uint256 tokenId) view returns(bool)
func (_AssetManagerFacet *AssetManagerFacetCallerSession) AssetExists(tokenId *big.Int) (bool, error) {
	return _AssetManagerFacet.Contract.AssetExists(&_AssetManagerFacet.CallOpts, tokenId)
}

// GetAssetConfig is a free data retrieval call binding the contract method 0xde31ea9f.
//
// Solidity: function getAssetConfig(uint256 tokenId) view returns((string,string,string,uint256,uint256,uint32,address[],address,bool,bool,uint16[]))
func (_AssetManagerFacet *AssetManagerFacetCaller) GetAssetConfig(opts *bind.CallOpts, tokenId *big.Int) (AssetConfig, error) {
	var out []interface{}
	err := _AssetManagerFacet.contract.Call(opts, &out, "getAssetConfig", tokenId)

	if err != nil {
		return *new(AssetConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(AssetConfig)).(*AssetConfig)

	return out0, err

}

// GetAssetConfig is a free data retrieval call binding the contract method 0xde31ea9f.
//
// Solidity: function getAssetConfig(uint256 tokenId) view returns((string,string,string,uint256,uint256,uint32,address[],address,bool,bool,uint16[]))
func (_AssetManagerFacet *AssetManagerFacetSession) GetAssetConfig(tokenId *big.Int) (AssetConfig, error) {
	return _AssetManagerFacet.Contract.GetAssetConfig(&_AssetManagerFacet.CallOpts, tokenId)
}

// GetAssetConfig is a free data retrieval call binding the contract method 0xde31ea9f.
//
// Solidity: function getAssetConfig(uint256 tokenId) view returns((string,string,string,uint256,uint256,uint32,address[],address,bool,bool,uint16[]))
func (_AssetManagerFacet *AssetManagerFacetCallerSession) GetAssetConfig(tokenId *big.Int) (AssetConfig, error) {
	return _AssetManagerFacet.Contract.GetAssetConfig(&_AssetManagerFacet.CallOpts, tokenId)
}

// GetComplianceModules is a free data retrieval call binding the contract method 0x71417690.
//
// Solidity: function getComplianceModules(uint256 tokenId) view returns(address[])
func (_AssetManagerFacet *AssetManagerFacetCaller) GetComplianceModules(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _AssetManagerFacet.contract.Call(opts, &out, "getComplianceModules", tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetComplianceModules is a free data retrieval call binding the contract method 0x71417690.
//
// Solidity: function getComplianceModules(uint256 tokenId) view returns(address[])
func (_AssetManagerFacet *AssetManagerFacetSession) GetComplianceModules(tokenId *big.Int) ([]common.Address, error) {
	return _AssetManagerFacet.Contract.GetComplianceModules(&_AssetManagerFacet.CallOpts, tokenId)
}

// GetComplianceModules is a free data retrieval call binding the contract method 0x71417690.
//
// Solidity: function getComplianceModules(uint256 tokenId) view returns(address[])
func (_AssetManagerFacet *AssetManagerFacetCallerSession) GetComplianceModules(tokenId *big.Int) ([]common.Address, error) {
	return _AssetManagerFacet.Contract.GetComplianceModules(&_AssetManagerFacet.CallOpts, tokenId)
}

// GetRegisteredTokenIds is a free data retrieval call binding the contract method 0x1df49c25.
//
// Solidity: function getRegisteredTokenIds() view returns(uint256[])
func (_AssetManagerFacet *AssetManagerFacetCaller) GetRegisteredTokenIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _AssetManagerFacet.contract.Call(opts, &out, "getRegisteredTokenIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRegisteredTokenIds is a free data retrieval call binding the contract method 0x1df49c25.
//
// Solidity: function getRegisteredTokenIds() view returns(uint256[])
func (_AssetManagerFacet *AssetManagerFacetSession) GetRegisteredTokenIds() ([]*big.Int, error) {
	return _AssetManagerFacet.Contract.GetRegisteredTokenIds(&_AssetManagerFacet.CallOpts)
}

// GetRegisteredTokenIds is a free data retrieval call binding the contract method 0x1df49c25.
//
// Solidity: function getRegisteredTokenIds() view returns(uint256[])
func (_AssetManagerFacet *AssetManagerFacetCallerSession) GetRegisteredTokenIds() ([]*big.Int, error) {
	return _AssetManagerFacet.Contract.GetRegisteredTokenIds(&_AssetManagerFacet.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AssetManagerFacet *AssetManagerFacetCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetManagerFacet.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AssetManagerFacet *AssetManagerFacetSession) NextTokenId() (*big.Int, error) {
	return _AssetManagerFacet.Contract.NextTokenId(&_AssetManagerFacet.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AssetManagerFacet *AssetManagerFacetCallerSession) NextTokenId() (*big.Int, error) {
	return _AssetManagerFacet.Contract.NextTokenId(&_AssetManagerFacet.CallOpts)
}

// AddComplianceModule is a paid mutator transaction binding the contract method 0x00c3169e.
//
// Solidity: function addComplianceModule(uint256 tokenId, address module) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactor) AddComplianceModule(opts *bind.TransactOpts, tokenId *big.Int, module common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.contract.Transact(opts, "addComplianceModule", tokenId, module)
}

// AddComplianceModule is a paid mutator transaction binding the contract method 0x00c3169e.
//
// Solidity: function addComplianceModule(uint256 tokenId, address module) returns()
func (_AssetManagerFacet *AssetManagerFacetSession) AddComplianceModule(tokenId *big.Int, module common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.AddComplianceModule(&_AssetManagerFacet.TransactOpts, tokenId, module)
}

// AddComplianceModule is a paid mutator transaction binding the contract method 0x00c3169e.
//
// Solidity: function addComplianceModule(uint256 tokenId, address module) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactorSession) AddComplianceModule(tokenId *big.Int, module common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.AddComplianceModule(&_AssetManagerFacet.TransactOpts, tokenId, module)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0xee48c026.
//
// Solidity: function registerAsset((string,string,string,uint256,uint32,address[],address,uint16[]) p) returns(uint256 tokenId)
func (_AssetManagerFacet *AssetManagerFacetTransactor) RegisterAsset(opts *bind.TransactOpts, p IAssetManagerRegisterAssetParams) (*types.Transaction, error) {
	return _AssetManagerFacet.contract.Transact(opts, "registerAsset", p)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0xee48c026.
//
// Solidity: function registerAsset((string,string,string,uint256,uint32,address[],address,uint16[]) p) returns(uint256 tokenId)
func (_AssetManagerFacet *AssetManagerFacetSession) RegisterAsset(p IAssetManagerRegisterAssetParams) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.RegisterAsset(&_AssetManagerFacet.TransactOpts, p)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0xee48c026.
//
// Solidity: function registerAsset((string,string,string,uint256,uint32,address[],address,uint16[]) p) returns(uint256 tokenId)
func (_AssetManagerFacet *AssetManagerFacetTransactorSession) RegisterAsset(p IAssetManagerRegisterAssetParams) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.RegisterAsset(&_AssetManagerFacet.TransactOpts, p)
}

// RemoveComplianceModule is a paid mutator transaction binding the contract method 0x17e648d4.
//
// Solidity: function removeComplianceModule(uint256 tokenId, address module) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactor) RemoveComplianceModule(opts *bind.TransactOpts, tokenId *big.Int, module common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.contract.Transact(opts, "removeComplianceModule", tokenId, module)
}

// RemoveComplianceModule is a paid mutator transaction binding the contract method 0x17e648d4.
//
// Solidity: function removeComplianceModule(uint256 tokenId, address module) returns()
func (_AssetManagerFacet *AssetManagerFacetSession) RemoveComplianceModule(tokenId *big.Int, module common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.RemoveComplianceModule(&_AssetManagerFacet.TransactOpts, tokenId, module)
}

// RemoveComplianceModule is a paid mutator transaction binding the contract method 0x17e648d4.
//
// Solidity: function removeComplianceModule(uint256 tokenId, address module) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactorSession) RemoveComplianceModule(tokenId *big.Int, module common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.RemoveComplianceModule(&_AssetManagerFacet.TransactOpts, tokenId, module)
}

// SetAllowedCountries is a paid mutator transaction binding the contract method 0x2445cbc3.
//
// Solidity: function setAllowedCountries(uint256 tokenId, uint16[] countries) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactor) SetAllowedCountries(opts *bind.TransactOpts, tokenId *big.Int, countries []uint16) (*types.Transaction, error) {
	return _AssetManagerFacet.contract.Transact(opts, "setAllowedCountries", tokenId, countries)
}

// SetAllowedCountries is a paid mutator transaction binding the contract method 0x2445cbc3.
//
// Solidity: function setAllowedCountries(uint256 tokenId, uint16[] countries) returns()
func (_AssetManagerFacet *AssetManagerFacetSession) SetAllowedCountries(tokenId *big.Int, countries []uint16) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetAllowedCountries(&_AssetManagerFacet.TransactOpts, tokenId, countries)
}

// SetAllowedCountries is a paid mutator transaction binding the contract method 0x2445cbc3.
//
// Solidity: function setAllowedCountries(uint256 tokenId, uint16[] countries) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactorSession) SetAllowedCountries(tokenId *big.Int, countries []uint16) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetAllowedCountries(&_AssetManagerFacet.TransactOpts, tokenId, countries)
}

// SetAssetUri is a paid mutator transaction binding the contract method 0x5f7b88a1.
//
// Solidity: function setAssetUri(uint256 tokenId, string uri) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactor) SetAssetUri(opts *bind.TransactOpts, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _AssetManagerFacet.contract.Transact(opts, "setAssetUri", tokenId, uri)
}

// SetAssetUri is a paid mutator transaction binding the contract method 0x5f7b88a1.
//
// Solidity: function setAssetUri(uint256 tokenId, string uri) returns()
func (_AssetManagerFacet *AssetManagerFacetSession) SetAssetUri(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetAssetUri(&_AssetManagerFacet.TransactOpts, tokenId, uri)
}

// SetAssetUri is a paid mutator transaction binding the contract method 0x5f7b88a1.
//
// Solidity: function setAssetUri(uint256 tokenId, string uri) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactorSession) SetAssetUri(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetAssetUri(&_AssetManagerFacet.TransactOpts, tokenId, uri)
}

// SetComplianceModules is a paid mutator transaction binding the contract method 0x74f691e2.
//
// Solidity: function setComplianceModules(uint256 tokenId, address[] modules) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactor) SetComplianceModules(opts *bind.TransactOpts, tokenId *big.Int, modules []common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.contract.Transact(opts, "setComplianceModules", tokenId, modules)
}

// SetComplianceModules is a paid mutator transaction binding the contract method 0x74f691e2.
//
// Solidity: function setComplianceModules(uint256 tokenId, address[] modules) returns()
func (_AssetManagerFacet *AssetManagerFacetSession) SetComplianceModules(tokenId *big.Int, modules []common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetComplianceModules(&_AssetManagerFacet.TransactOpts, tokenId, modules)
}

// SetComplianceModules is a paid mutator transaction binding the contract method 0x74f691e2.
//
// Solidity: function setComplianceModules(uint256 tokenId, address[] modules) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactorSession) SetComplianceModules(tokenId *big.Int, modules []common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetComplianceModules(&_AssetManagerFacet.TransactOpts, tokenId, modules)
}

// SetIdentityProfile is a paid mutator transaction binding the contract method 0x3590bd78.
//
// Solidity: function setIdentityProfile(uint256 tokenId, uint32 profileId) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactor) SetIdentityProfile(opts *bind.TransactOpts, tokenId *big.Int, profileId uint32) (*types.Transaction, error) {
	return _AssetManagerFacet.contract.Transact(opts, "setIdentityProfile", tokenId, profileId)
}

// SetIdentityProfile is a paid mutator transaction binding the contract method 0x3590bd78.
//
// Solidity: function setIdentityProfile(uint256 tokenId, uint32 profileId) returns()
func (_AssetManagerFacet *AssetManagerFacetSession) SetIdentityProfile(tokenId *big.Int, profileId uint32) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetIdentityProfile(&_AssetManagerFacet.TransactOpts, tokenId, profileId)
}

// SetIdentityProfile is a paid mutator transaction binding the contract method 0x3590bd78.
//
// Solidity: function setIdentityProfile(uint256 tokenId, uint32 profileId) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactorSession) SetIdentityProfile(tokenId *big.Int, profileId uint32) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetIdentityProfile(&_AssetManagerFacet.TransactOpts, tokenId, profileId)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x074b38ae.
//
// Solidity: function setIssuer(uint256 tokenId, address issuer) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactor) SetIssuer(opts *bind.TransactOpts, tokenId *big.Int, issuer common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.contract.Transact(opts, "setIssuer", tokenId, issuer)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x074b38ae.
//
// Solidity: function setIssuer(uint256 tokenId, address issuer) returns()
func (_AssetManagerFacet *AssetManagerFacetSession) SetIssuer(tokenId *big.Int, issuer common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetIssuer(&_AssetManagerFacet.TransactOpts, tokenId, issuer)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x074b38ae.
//
// Solidity: function setIssuer(uint256 tokenId, address issuer) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactorSession) SetIssuer(tokenId *big.Int, issuer common.Address) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetIssuer(&_AssetManagerFacet.TransactOpts, tokenId, issuer)
}

// SetSupplyCap is a paid mutator transaction binding the contract method 0xccb3cf07.
//
// Solidity: function setSupplyCap(uint256 tokenId, uint256 cap) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactor) SetSupplyCap(opts *bind.TransactOpts, tokenId *big.Int, cap *big.Int) (*types.Transaction, error) {
	return _AssetManagerFacet.contract.Transact(opts, "setSupplyCap", tokenId, cap)
}

// SetSupplyCap is a paid mutator transaction binding the contract method 0xccb3cf07.
//
// Solidity: function setSupplyCap(uint256 tokenId, uint256 cap) returns()
func (_AssetManagerFacet *AssetManagerFacetSession) SetSupplyCap(tokenId *big.Int, cap *big.Int) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetSupplyCap(&_AssetManagerFacet.TransactOpts, tokenId, cap)
}

// SetSupplyCap is a paid mutator transaction binding the contract method 0xccb3cf07.
//
// Solidity: function setSupplyCap(uint256 tokenId, uint256 cap) returns()
func (_AssetManagerFacet *AssetManagerFacetTransactorSession) SetSupplyCap(tokenId *big.Int, cap *big.Int) (*types.Transaction, error) {
	return _AssetManagerFacet.Contract.SetSupplyCap(&_AssetManagerFacet.TransactOpts, tokenId, cap)
}

// AssetManagerFacetAssetConfigUpdatedIterator is returned from FilterAssetConfigUpdated and is used to iterate over the raw logs and unpacked data for AssetConfigUpdated events raised by the AssetManagerFacet contract.
type AssetManagerFacetAssetConfigUpdatedIterator struct {
	Event *AssetManagerFacetAssetConfigUpdated // Event containing the contract specifics and raw log

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
func (it *AssetManagerFacetAssetConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerFacetAssetConfigUpdated)
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
		it.Event = new(AssetManagerFacetAssetConfigUpdated)
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
func (it *AssetManagerFacetAssetConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerFacetAssetConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerFacetAssetConfigUpdated represents a AssetConfigUpdated event raised by the AssetManagerFacet contract.
type AssetManagerFacetAssetConfigUpdated struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetConfigUpdated is a free log retrieval operation binding the contract event 0x05555224d5e2cf9b763227d0940b583bf752950c7f14f5c808ab1fdc5df2d3e3.
//
// Solidity: event AssetConfigUpdated(uint256 indexed tokenId)
func (_AssetManagerFacet *AssetManagerFacetFilterer) FilterAssetConfigUpdated(opts *bind.FilterOpts, tokenId []*big.Int) (*AssetManagerFacetAssetConfigUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.FilterLogs(opts, "AssetConfigUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFacetAssetConfigUpdatedIterator{contract: _AssetManagerFacet.contract, event: "AssetConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetConfigUpdated is a free log subscription operation binding the contract event 0x05555224d5e2cf9b763227d0940b583bf752950c7f14f5c808ab1fdc5df2d3e3.
//
// Solidity: event AssetConfigUpdated(uint256 indexed tokenId)
func (_AssetManagerFacet *AssetManagerFacetFilterer) WatchAssetConfigUpdated(opts *bind.WatchOpts, sink chan<- *AssetManagerFacetAssetConfigUpdated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.WatchLogs(opts, "AssetConfigUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerFacetAssetConfigUpdated)
				if err := _AssetManagerFacet.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
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

// ParseAssetConfigUpdated is a log parse operation binding the contract event 0x05555224d5e2cf9b763227d0940b583bf752950c7f14f5c808ab1fdc5df2d3e3.
//
// Solidity: event AssetConfigUpdated(uint256 indexed tokenId)
func (_AssetManagerFacet *AssetManagerFacetFilterer) ParseAssetConfigUpdated(log types.Log) (*AssetManagerFacetAssetConfigUpdated, error) {
	event := new(AssetManagerFacetAssetConfigUpdated)
	if err := _AssetManagerFacet.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerFacetAssetRegisteredIterator is returned from FilterAssetRegistered and is used to iterate over the raw logs and unpacked data for AssetRegistered events raised by the AssetManagerFacet contract.
type AssetManagerFacetAssetRegisteredIterator struct {
	Event *AssetManagerFacetAssetRegistered // Event containing the contract specifics and raw log

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
func (it *AssetManagerFacetAssetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerFacetAssetRegistered)
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
		it.Event = new(AssetManagerFacetAssetRegistered)
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
func (it *AssetManagerFacetAssetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerFacetAssetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerFacetAssetRegistered represents a AssetRegistered event raised by the AssetManagerFacet contract.
type AssetManagerFacetAssetRegistered struct {
	TokenId   *big.Int
	Issuer    common.Address
	ProfileId uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssetRegistered is a free log retrieval operation binding the contract event 0x8b249cee60df2ced7ffea85d48eef23c7ddfd72be05a3bc50ee7960dbe6e6c8d.
//
// Solidity: event AssetRegistered(uint256 indexed tokenId, address indexed issuer, uint32 profileId)
func (_AssetManagerFacet *AssetManagerFacetFilterer) FilterAssetRegistered(opts *bind.FilterOpts, tokenId []*big.Int, issuer []common.Address) (*AssetManagerFacetAssetRegisteredIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.FilterLogs(opts, "AssetRegistered", tokenIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFacetAssetRegisteredIterator{contract: _AssetManagerFacet.contract, event: "AssetRegistered", logs: logs, sub: sub}, nil
}

// WatchAssetRegistered is a free log subscription operation binding the contract event 0x8b249cee60df2ced7ffea85d48eef23c7ddfd72be05a3bc50ee7960dbe6e6c8d.
//
// Solidity: event AssetRegistered(uint256 indexed tokenId, address indexed issuer, uint32 profileId)
func (_AssetManagerFacet *AssetManagerFacetFilterer) WatchAssetRegistered(opts *bind.WatchOpts, sink chan<- *AssetManagerFacetAssetRegistered, tokenId []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.WatchLogs(opts, "AssetRegistered", tokenIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerFacetAssetRegistered)
				if err := _AssetManagerFacet.contract.UnpackLog(event, "AssetRegistered", log); err != nil {
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

// ParseAssetRegistered is a log parse operation binding the contract event 0x8b249cee60df2ced7ffea85d48eef23c7ddfd72be05a3bc50ee7960dbe6e6c8d.
//
// Solidity: event AssetRegistered(uint256 indexed tokenId, address indexed issuer, uint32 profileId)
func (_AssetManagerFacet *AssetManagerFacetFilterer) ParseAssetRegistered(log types.Log) (*AssetManagerFacetAssetRegistered, error) {
	event := new(AssetManagerFacetAssetRegistered)
	if err := _AssetManagerFacet.contract.UnpackLog(event, "AssetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerFacetComplianceModuleAddedIterator is returned from FilterComplianceModuleAdded and is used to iterate over the raw logs and unpacked data for ComplianceModuleAdded events raised by the AssetManagerFacet contract.
type AssetManagerFacetComplianceModuleAddedIterator struct {
	Event *AssetManagerFacetComplianceModuleAdded // Event containing the contract specifics and raw log

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
func (it *AssetManagerFacetComplianceModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerFacetComplianceModuleAdded)
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
		it.Event = new(AssetManagerFacetComplianceModuleAdded)
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
func (it *AssetManagerFacetComplianceModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerFacetComplianceModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerFacetComplianceModuleAdded represents a ComplianceModuleAdded event raised by the AssetManagerFacet contract.
type AssetManagerFacetComplianceModuleAdded struct {
	TokenId *big.Int
	Module  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterComplianceModuleAdded is a free log retrieval operation binding the contract event 0xabcc35061552383bd7e261c95e3528941e4c157881d874c30daf186176e110a9.
//
// Solidity: event ComplianceModuleAdded(uint256 indexed tokenId, address indexed module)
func (_AssetManagerFacet *AssetManagerFacetFilterer) FilterComplianceModuleAdded(opts *bind.FilterOpts, tokenId []*big.Int, module []common.Address) (*AssetManagerFacetComplianceModuleAddedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.FilterLogs(opts, "ComplianceModuleAdded", tokenIdRule, moduleRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFacetComplianceModuleAddedIterator{contract: _AssetManagerFacet.contract, event: "ComplianceModuleAdded", logs: logs, sub: sub}, nil
}

// WatchComplianceModuleAdded is a free log subscription operation binding the contract event 0xabcc35061552383bd7e261c95e3528941e4c157881d874c30daf186176e110a9.
//
// Solidity: event ComplianceModuleAdded(uint256 indexed tokenId, address indexed module)
func (_AssetManagerFacet *AssetManagerFacetFilterer) WatchComplianceModuleAdded(opts *bind.WatchOpts, sink chan<- *AssetManagerFacetComplianceModuleAdded, tokenId []*big.Int, module []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.WatchLogs(opts, "ComplianceModuleAdded", tokenIdRule, moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerFacetComplianceModuleAdded)
				if err := _AssetManagerFacet.contract.UnpackLog(event, "ComplianceModuleAdded", log); err != nil {
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

// ParseComplianceModuleAdded is a log parse operation binding the contract event 0xabcc35061552383bd7e261c95e3528941e4c157881d874c30daf186176e110a9.
//
// Solidity: event ComplianceModuleAdded(uint256 indexed tokenId, address indexed module)
func (_AssetManagerFacet *AssetManagerFacetFilterer) ParseComplianceModuleAdded(log types.Log) (*AssetManagerFacetComplianceModuleAdded, error) {
	event := new(AssetManagerFacetComplianceModuleAdded)
	if err := _AssetManagerFacet.contract.UnpackLog(event, "ComplianceModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerFacetComplianceModuleRemovedIterator is returned from FilterComplianceModuleRemoved and is used to iterate over the raw logs and unpacked data for ComplianceModuleRemoved events raised by the AssetManagerFacet contract.
type AssetManagerFacetComplianceModuleRemovedIterator struct {
	Event *AssetManagerFacetComplianceModuleRemoved // Event containing the contract specifics and raw log

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
func (it *AssetManagerFacetComplianceModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerFacetComplianceModuleRemoved)
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
		it.Event = new(AssetManagerFacetComplianceModuleRemoved)
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
func (it *AssetManagerFacetComplianceModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerFacetComplianceModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerFacetComplianceModuleRemoved represents a ComplianceModuleRemoved event raised by the AssetManagerFacet contract.
type AssetManagerFacetComplianceModuleRemoved struct {
	TokenId *big.Int
	Module  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterComplianceModuleRemoved is a free log retrieval operation binding the contract event 0x05c5c1bbcf90a4dce5c8a68fee0b15a588bcb04538a4e1c3ae42c7174202946d.
//
// Solidity: event ComplianceModuleRemoved(uint256 indexed tokenId, address indexed module)
func (_AssetManagerFacet *AssetManagerFacetFilterer) FilterComplianceModuleRemoved(opts *bind.FilterOpts, tokenId []*big.Int, module []common.Address) (*AssetManagerFacetComplianceModuleRemovedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.FilterLogs(opts, "ComplianceModuleRemoved", tokenIdRule, moduleRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFacetComplianceModuleRemovedIterator{contract: _AssetManagerFacet.contract, event: "ComplianceModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchComplianceModuleRemoved is a free log subscription operation binding the contract event 0x05c5c1bbcf90a4dce5c8a68fee0b15a588bcb04538a4e1c3ae42c7174202946d.
//
// Solidity: event ComplianceModuleRemoved(uint256 indexed tokenId, address indexed module)
func (_AssetManagerFacet *AssetManagerFacetFilterer) WatchComplianceModuleRemoved(opts *bind.WatchOpts, sink chan<- *AssetManagerFacetComplianceModuleRemoved, tokenId []*big.Int, module []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.WatchLogs(opts, "ComplianceModuleRemoved", tokenIdRule, moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerFacetComplianceModuleRemoved)
				if err := _AssetManagerFacet.contract.UnpackLog(event, "ComplianceModuleRemoved", log); err != nil {
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

// ParseComplianceModuleRemoved is a log parse operation binding the contract event 0x05c5c1bbcf90a4dce5c8a68fee0b15a588bcb04538a4e1c3ae42c7174202946d.
//
// Solidity: event ComplianceModuleRemoved(uint256 indexed tokenId, address indexed module)
func (_AssetManagerFacet *AssetManagerFacetFilterer) ParseComplianceModuleRemoved(log types.Log) (*AssetManagerFacetComplianceModuleRemoved, error) {
	event := new(AssetManagerFacetComplianceModuleRemoved)
	if err := _AssetManagerFacet.contract.UnpackLog(event, "ComplianceModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerFacetComplianceModulesSetIterator is returned from FilterComplianceModulesSet and is used to iterate over the raw logs and unpacked data for ComplianceModulesSet events raised by the AssetManagerFacet contract.
type AssetManagerFacetComplianceModulesSetIterator struct {
	Event *AssetManagerFacetComplianceModulesSet // Event containing the contract specifics and raw log

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
func (it *AssetManagerFacetComplianceModulesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerFacetComplianceModulesSet)
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
		it.Event = new(AssetManagerFacetComplianceModulesSet)
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
func (it *AssetManagerFacetComplianceModulesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerFacetComplianceModulesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerFacetComplianceModulesSet represents a ComplianceModulesSet event raised by the AssetManagerFacet contract.
type AssetManagerFacetComplianceModulesSet struct {
	TokenId *big.Int
	Modules []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterComplianceModulesSet is a free log retrieval operation binding the contract event 0xa827c786814c67851536e5bf50f965c6c5f3799aea5a8674ef22641d3897dab5.
//
// Solidity: event ComplianceModulesSet(uint256 indexed tokenId, address[] modules)
func (_AssetManagerFacet *AssetManagerFacetFilterer) FilterComplianceModulesSet(opts *bind.FilterOpts, tokenId []*big.Int) (*AssetManagerFacetComplianceModulesSetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.FilterLogs(opts, "ComplianceModulesSet", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFacetComplianceModulesSetIterator{contract: _AssetManagerFacet.contract, event: "ComplianceModulesSet", logs: logs, sub: sub}, nil
}

// WatchComplianceModulesSet is a free log subscription operation binding the contract event 0xa827c786814c67851536e5bf50f965c6c5f3799aea5a8674ef22641d3897dab5.
//
// Solidity: event ComplianceModulesSet(uint256 indexed tokenId, address[] modules)
func (_AssetManagerFacet *AssetManagerFacetFilterer) WatchComplianceModulesSet(opts *bind.WatchOpts, sink chan<- *AssetManagerFacetComplianceModulesSet, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.WatchLogs(opts, "ComplianceModulesSet", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerFacetComplianceModulesSet)
				if err := _AssetManagerFacet.contract.UnpackLog(event, "ComplianceModulesSet", log); err != nil {
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

// ParseComplianceModulesSet is a log parse operation binding the contract event 0xa827c786814c67851536e5bf50f965c6c5f3799aea5a8674ef22641d3897dab5.
//
// Solidity: event ComplianceModulesSet(uint256 indexed tokenId, address[] modules)
func (_AssetManagerFacet *AssetManagerFacetFilterer) ParseComplianceModulesSet(log types.Log) (*AssetManagerFacetComplianceModulesSet, error) {
	event := new(AssetManagerFacetComplianceModulesSet)
	if err := _AssetManagerFacet.contract.UnpackLog(event, "ComplianceModulesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerFacetURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the AssetManagerFacet contract.
type AssetManagerFacetURIIterator struct {
	Event *AssetManagerFacetURI // Event containing the contract specifics and raw log

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
func (it *AssetManagerFacetURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerFacetURI)
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
		it.Event = new(AssetManagerFacetURI)
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
func (it *AssetManagerFacetURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerFacetURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerFacetURI represents a URI event raised by the AssetManagerFacet contract.
type AssetManagerFacetURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_AssetManagerFacet *AssetManagerFacetFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*AssetManagerFacetURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFacetURIIterator{contract: _AssetManagerFacet.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_AssetManagerFacet *AssetManagerFacetFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *AssetManagerFacetURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AssetManagerFacet.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerFacetURI)
				if err := _AssetManagerFacet.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_AssetManagerFacet *AssetManagerFacetFilterer) ParseURI(log types.Log) (*AssetManagerFacetURI, error) {
	event := new(AssetManagerFacetURI)
	if err := _AssetManagerFacet.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
