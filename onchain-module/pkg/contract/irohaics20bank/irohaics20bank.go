// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package irohaics20bank

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

// Irohaics20bankABI is the input ABI used to generate the binding from.
const Irohaics20bankABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"srcAccountId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"amount\",\"type\":\"string\"}],\"name\":\"BurnRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destAccountId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"amount\",\"type\":\"string\"}],\"name\":\"MintRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"BANK_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"ICS20_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setIcs20Contract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"accountId\",\"type\":\"string\"}],\"name\":\"setBank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"setNextBurnRequestId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"setNextMintRequestId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"srcAccountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"amount\",\"type\":\"string\"}],\"name\":\"requestBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"destAccountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"amount\",\"type\":\"string\"}],\"name\":\"requestMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Irohaics20bank is an auto generated Go binding around an Ethereum contract.
type Irohaics20bank struct {
	Irohaics20bankCaller     // Read-only binding to the contract
	Irohaics20bankTransactor // Write-only binding to the contract
	Irohaics20bankFilterer   // Log filterer for contract events
}

// Irohaics20bankCaller is an auto generated read-only Go binding around an Ethereum contract.
type Irohaics20bankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Irohaics20bankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Irohaics20bankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Irohaics20bankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Irohaics20bankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Irohaics20bankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Irohaics20bankSession struct {
	Contract     *Irohaics20bank   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Irohaics20bankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Irohaics20bankCallerSession struct {
	Contract *Irohaics20bankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Irohaics20bankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Irohaics20bankTransactorSession struct {
	Contract     *Irohaics20bankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Irohaics20bankRaw is an auto generated low-level Go binding around an Ethereum contract.
type Irohaics20bankRaw struct {
	Contract *Irohaics20bank // Generic contract binding to access the raw methods on
}

// Irohaics20bankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Irohaics20bankCallerRaw struct {
	Contract *Irohaics20bankCaller // Generic read-only contract binding to access the raw methods on
}

// Irohaics20bankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Irohaics20bankTransactorRaw struct {
	Contract *Irohaics20bankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIrohaics20bank creates a new instance of Irohaics20bank, bound to a specific deployed contract.
func NewIrohaics20bank(address common.Address, backend bind.ContractBackend) (*Irohaics20bank, error) {
	contract, err := bindIrohaics20bank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Irohaics20bank{Irohaics20bankCaller: Irohaics20bankCaller{contract: contract}, Irohaics20bankTransactor: Irohaics20bankTransactor{contract: contract}, Irohaics20bankFilterer: Irohaics20bankFilterer{contract: contract}}, nil
}

// NewIrohaics20bankCaller creates a new read-only instance of Irohaics20bank, bound to a specific deployed contract.
func NewIrohaics20bankCaller(address common.Address, caller bind.ContractCaller) (*Irohaics20bankCaller, error) {
	contract, err := bindIrohaics20bank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Irohaics20bankCaller{contract: contract}, nil
}

// NewIrohaics20bankTransactor creates a new write-only instance of Irohaics20bank, bound to a specific deployed contract.
func NewIrohaics20bankTransactor(address common.Address, transactor bind.ContractTransactor) (*Irohaics20bankTransactor, error) {
	contract, err := bindIrohaics20bank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Irohaics20bankTransactor{contract: contract}, nil
}

// NewIrohaics20bankFilterer creates a new log filterer instance of Irohaics20bank, bound to a specific deployed contract.
func NewIrohaics20bankFilterer(address common.Address, filterer bind.ContractFilterer) (*Irohaics20bankFilterer, error) {
	contract, err := bindIrohaics20bank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Irohaics20bankFilterer{contract: contract}, nil
}

// bindIrohaics20bank binds a generic wrapper to an already deployed contract.
func bindIrohaics20bank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Irohaics20bankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Irohaics20bank *Irohaics20bankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Irohaics20bank.Contract.Irohaics20bankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Irohaics20bank *Irohaics20bankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.Irohaics20bankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Irohaics20bank *Irohaics20bankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.Irohaics20bankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Irohaics20bank *Irohaics20bankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Irohaics20bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Irohaics20bank *Irohaics20bankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Irohaics20bank *Irohaics20bankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irohaics20bank.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankSession) ADMINROLE() ([32]byte, error) {
	return _Irohaics20bank.Contract.ADMINROLE(&_Irohaics20bank.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankCallerSession) ADMINROLE() ([32]byte, error) {
	return _Irohaics20bank.Contract.ADMINROLE(&_Irohaics20bank.CallOpts)
}

// BANKROLE is a free data retrieval call binding the contract method 0x0eb71432.
//
// Solidity: function BANK_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankCaller) BANKROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irohaics20bank.contract.Call(opts, &out, "BANK_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANKROLE is a free data retrieval call binding the contract method 0x0eb71432.
//
// Solidity: function BANK_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankSession) BANKROLE() ([32]byte, error) {
	return _Irohaics20bank.Contract.BANKROLE(&_Irohaics20bank.CallOpts)
}

// BANKROLE is a free data retrieval call binding the contract method 0x0eb71432.
//
// Solidity: function BANK_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankCallerSession) BANKROLE() ([32]byte, error) {
	return _Irohaics20bank.Contract.BANKROLE(&_Irohaics20bank.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irohaics20bank.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Irohaics20bank.Contract.DEFAULTADMINROLE(&_Irohaics20bank.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Irohaics20bank.Contract.DEFAULTADMINROLE(&_Irohaics20bank.CallOpts)
}

// ICS20ROLE is a free data retrieval call binding the contract method 0xf014966e.
//
// Solidity: function ICS20_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankCaller) ICS20ROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Irohaics20bank.contract.Call(opts, &out, "ICS20_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ICS20ROLE is a free data retrieval call binding the contract method 0xf014966e.
//
// Solidity: function ICS20_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankSession) ICS20ROLE() ([32]byte, error) {
	return _Irohaics20bank.Contract.ICS20ROLE(&_Irohaics20bank.CallOpts)
}

// ICS20ROLE is a free data retrieval call binding the contract method 0xf014966e.
//
// Solidity: function ICS20_ROLE() view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankCallerSession) ICS20ROLE() ([32]byte, error) {
	return _Irohaics20bank.Contract.ICS20ROLE(&_Irohaics20bank.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Irohaics20bank.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Irohaics20bank.Contract.GetRoleAdmin(&_Irohaics20bank.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Irohaics20bank *Irohaics20bankCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Irohaics20bank.Contract.GetRoleAdmin(&_Irohaics20bank.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Irohaics20bank *Irohaics20bankCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Irohaics20bank.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Irohaics20bank *Irohaics20bankSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Irohaics20bank.Contract.HasRole(&_Irohaics20bank.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Irohaics20bank *Irohaics20bankCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Irohaics20bank.Contract.HasRole(&_Irohaics20bank.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Irohaics20bank *Irohaics20bankCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Irohaics20bank.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Irohaics20bank *Irohaics20bankSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Irohaics20bank.Contract.SupportsInterface(&_Irohaics20bank.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Irohaics20bank *Irohaics20bankCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Irohaics20bank.Contract.SupportsInterface(&_Irohaics20bank.CallOpts, interfaceId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) Burn(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "burn", requestId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankSession) Burn(requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.Burn(&_Irohaics20bank.TransactOpts, requestId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) Burn(requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.Burn(&_Irohaics20bank.TransactOpts, requestId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Irohaics20bank *Irohaics20bankSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.GrantRole(&_Irohaics20bank.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.GrantRole(&_Irohaics20bank.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) Mint(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "mint", requestId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankSession) Mint(requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.Mint(&_Irohaics20bank.TransactOpts, requestId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) Mint(requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.Mint(&_Irohaics20bank.TransactOpts, requestId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Irohaics20bank *Irohaics20bankSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.RenounceRole(&_Irohaics20bank.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.RenounceRole(&_Irohaics20bank.TransactOpts, role, account)
}

// RequestBurn is a paid mutator transaction binding the contract method 0x22dd65b3.
//
// Solidity: function requestBurn(string srcAccountId, string assetId, string description, string amount) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) RequestBurn(opts *bind.TransactOpts, srcAccountId string, assetId string, description string, amount string) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "requestBurn", srcAccountId, assetId, description, amount)
}

// RequestBurn is a paid mutator transaction binding the contract method 0x22dd65b3.
//
// Solidity: function requestBurn(string srcAccountId, string assetId, string description, string amount) returns()
func (_Irohaics20bank *Irohaics20bankSession) RequestBurn(srcAccountId string, assetId string, description string, amount string) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.RequestBurn(&_Irohaics20bank.TransactOpts, srcAccountId, assetId, description, amount)
}

// RequestBurn is a paid mutator transaction binding the contract method 0x22dd65b3.
//
// Solidity: function requestBurn(string srcAccountId, string assetId, string description, string amount) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) RequestBurn(srcAccountId string, assetId string, description string, amount string) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.RequestBurn(&_Irohaics20bank.TransactOpts, srcAccountId, assetId, description, amount)
}

// RequestMint is a paid mutator transaction binding the contract method 0xc1ee91a4.
//
// Solidity: function requestMint(string destAccountId, string assetId, string description, string amount) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) RequestMint(opts *bind.TransactOpts, destAccountId string, assetId string, description string, amount string) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "requestMint", destAccountId, assetId, description, amount)
}

// RequestMint is a paid mutator transaction binding the contract method 0xc1ee91a4.
//
// Solidity: function requestMint(string destAccountId, string assetId, string description, string amount) returns()
func (_Irohaics20bank *Irohaics20bankSession) RequestMint(destAccountId string, assetId string, description string, amount string) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.RequestMint(&_Irohaics20bank.TransactOpts, destAccountId, assetId, description, amount)
}

// RequestMint is a paid mutator transaction binding the contract method 0xc1ee91a4.
//
// Solidity: function requestMint(string destAccountId, string assetId, string description, string amount) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) RequestMint(destAccountId string, assetId string, description string, amount string) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.RequestMint(&_Irohaics20bank.TransactOpts, destAccountId, assetId, description, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Irohaics20bank *Irohaics20bankSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.RevokeRole(&_Irohaics20bank.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.RevokeRole(&_Irohaics20bank.TransactOpts, role, account)
}

// SetBank is a paid mutator transaction binding the contract method 0xde0697d1.
//
// Solidity: function setBank(string accountId) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) SetBank(opts *bind.TransactOpts, accountId string) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "setBank", accountId)
}

// SetBank is a paid mutator transaction binding the contract method 0xde0697d1.
//
// Solidity: function setBank(string accountId) returns()
func (_Irohaics20bank *Irohaics20bankSession) SetBank(accountId string) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.SetBank(&_Irohaics20bank.TransactOpts, accountId)
}

// SetBank is a paid mutator transaction binding the contract method 0xde0697d1.
//
// Solidity: function setBank(string accountId) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) SetBank(accountId string) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.SetBank(&_Irohaics20bank.TransactOpts, accountId)
}

// SetIcs20Contract is a paid mutator transaction binding the contract method 0x47c9e18a.
//
// Solidity: function setIcs20Contract(address addr) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) SetIcs20Contract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "setIcs20Contract", addr)
}

// SetIcs20Contract is a paid mutator transaction binding the contract method 0x47c9e18a.
//
// Solidity: function setIcs20Contract(address addr) returns()
func (_Irohaics20bank *Irohaics20bankSession) SetIcs20Contract(addr common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.SetIcs20Contract(&_Irohaics20bank.TransactOpts, addr)
}

// SetIcs20Contract is a paid mutator transaction binding the contract method 0x47c9e18a.
//
// Solidity: function setIcs20Contract(address addr) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) SetIcs20Contract(addr common.Address) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.SetIcs20Contract(&_Irohaics20bank.TransactOpts, addr)
}

// SetNextBurnRequestId is a paid mutator transaction binding the contract method 0x7821a935.
//
// Solidity: function setNextBurnRequestId(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) SetNextBurnRequestId(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "setNextBurnRequestId", requestId)
}

// SetNextBurnRequestId is a paid mutator transaction binding the contract method 0x7821a935.
//
// Solidity: function setNextBurnRequestId(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankSession) SetNextBurnRequestId(requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.SetNextBurnRequestId(&_Irohaics20bank.TransactOpts, requestId)
}

// SetNextBurnRequestId is a paid mutator transaction binding the contract method 0x7821a935.
//
// Solidity: function setNextBurnRequestId(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) SetNextBurnRequestId(requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.SetNextBurnRequestId(&_Irohaics20bank.TransactOpts, requestId)
}

// SetNextMintRequestId is a paid mutator transaction binding the contract method 0x997188a8.
//
// Solidity: function setNextMintRequestId(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankTransactor) SetNextMintRequestId(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.contract.Transact(opts, "setNextMintRequestId", requestId)
}

// SetNextMintRequestId is a paid mutator transaction binding the contract method 0x997188a8.
//
// Solidity: function setNextMintRequestId(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankSession) SetNextMintRequestId(requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.SetNextMintRequestId(&_Irohaics20bank.TransactOpts, requestId)
}

// SetNextMintRequestId is a paid mutator transaction binding the contract method 0x997188a8.
//
// Solidity: function setNextMintRequestId(uint256 requestId) returns()
func (_Irohaics20bank *Irohaics20bankTransactorSession) SetNextMintRequestId(requestId *big.Int) (*types.Transaction, error) {
	return _Irohaics20bank.Contract.SetNextMintRequestId(&_Irohaics20bank.TransactOpts, requestId)
}

// Irohaics20bankBurnRequestedIterator is returned from FilterBurnRequested and is used to iterate over the raw logs and unpacked data for BurnRequested events raised by the Irohaics20bank contract.
type Irohaics20bankBurnRequestedIterator struct {
	Event *Irohaics20bankBurnRequested // Event containing the contract specifics and raw log

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
func (it *Irohaics20bankBurnRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Irohaics20bankBurnRequested)
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
		it.Event = new(Irohaics20bankBurnRequested)
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
func (it *Irohaics20bankBurnRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Irohaics20bankBurnRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Irohaics20bankBurnRequested represents a BurnRequested event raised by the Irohaics20bank contract.
type Irohaics20bankBurnRequested struct {
	Id           *big.Int
	SrcAccountId string
	AssetId      string
	Description  string
	Amount       string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurnRequested is a free log retrieval operation binding the contract event 0xb1b0f000252bdaf51b9d8a3b02b12c0f07b278c2d25c1767a70728c522668510.
//
// Solidity: event BurnRequested(uint256 id, string srcAccountId, string assetId, string description, string amount)
func (_Irohaics20bank *Irohaics20bankFilterer) FilterBurnRequested(opts *bind.FilterOpts) (*Irohaics20bankBurnRequestedIterator, error) {

	logs, sub, err := _Irohaics20bank.contract.FilterLogs(opts, "BurnRequested")
	if err != nil {
		return nil, err
	}
	return &Irohaics20bankBurnRequestedIterator{contract: _Irohaics20bank.contract, event: "BurnRequested", logs: logs, sub: sub}, nil
}

// WatchBurnRequested is a free log subscription operation binding the contract event 0xb1b0f000252bdaf51b9d8a3b02b12c0f07b278c2d25c1767a70728c522668510.
//
// Solidity: event BurnRequested(uint256 id, string srcAccountId, string assetId, string description, string amount)
func (_Irohaics20bank *Irohaics20bankFilterer) WatchBurnRequested(opts *bind.WatchOpts, sink chan<- *Irohaics20bankBurnRequested) (event.Subscription, error) {

	logs, sub, err := _Irohaics20bank.contract.WatchLogs(opts, "BurnRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Irohaics20bankBurnRequested)
				if err := _Irohaics20bank.contract.UnpackLog(event, "BurnRequested", log); err != nil {
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

// ParseBurnRequested is a log parse operation binding the contract event 0xb1b0f000252bdaf51b9d8a3b02b12c0f07b278c2d25c1767a70728c522668510.
//
// Solidity: event BurnRequested(uint256 id, string srcAccountId, string assetId, string description, string amount)
func (_Irohaics20bank *Irohaics20bankFilterer) ParseBurnRequested(log types.Log) (*Irohaics20bankBurnRequested, error) {
	event := new(Irohaics20bankBurnRequested)
	if err := _Irohaics20bank.contract.UnpackLog(event, "BurnRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Irohaics20bankMintRequestedIterator is returned from FilterMintRequested and is used to iterate over the raw logs and unpacked data for MintRequested events raised by the Irohaics20bank contract.
type Irohaics20bankMintRequestedIterator struct {
	Event *Irohaics20bankMintRequested // Event containing the contract specifics and raw log

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
func (it *Irohaics20bankMintRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Irohaics20bankMintRequested)
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
		it.Event = new(Irohaics20bankMintRequested)
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
func (it *Irohaics20bankMintRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Irohaics20bankMintRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Irohaics20bankMintRequested represents a MintRequested event raised by the Irohaics20bank contract.
type Irohaics20bankMintRequested struct {
	Id            *big.Int
	DestAccountId string
	AssetId       string
	Description   string
	Amount        string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMintRequested is a free log retrieval operation binding the contract event 0xd0c2e6a3cbb4beac7c4f2ed15e8ebb1b0f63b4c0a731577bded43cee8287a49e.
//
// Solidity: event MintRequested(uint256 id, string destAccountId, string assetId, string description, string amount)
func (_Irohaics20bank *Irohaics20bankFilterer) FilterMintRequested(opts *bind.FilterOpts) (*Irohaics20bankMintRequestedIterator, error) {

	logs, sub, err := _Irohaics20bank.contract.FilterLogs(opts, "MintRequested")
	if err != nil {
		return nil, err
	}
	return &Irohaics20bankMintRequestedIterator{contract: _Irohaics20bank.contract, event: "MintRequested", logs: logs, sub: sub}, nil
}

// WatchMintRequested is a free log subscription operation binding the contract event 0xd0c2e6a3cbb4beac7c4f2ed15e8ebb1b0f63b4c0a731577bded43cee8287a49e.
//
// Solidity: event MintRequested(uint256 id, string destAccountId, string assetId, string description, string amount)
func (_Irohaics20bank *Irohaics20bankFilterer) WatchMintRequested(opts *bind.WatchOpts, sink chan<- *Irohaics20bankMintRequested) (event.Subscription, error) {

	logs, sub, err := _Irohaics20bank.contract.WatchLogs(opts, "MintRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Irohaics20bankMintRequested)
				if err := _Irohaics20bank.contract.UnpackLog(event, "MintRequested", log); err != nil {
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

// ParseMintRequested is a log parse operation binding the contract event 0xd0c2e6a3cbb4beac7c4f2ed15e8ebb1b0f63b4c0a731577bded43cee8287a49e.
//
// Solidity: event MintRequested(uint256 id, string destAccountId, string assetId, string description, string amount)
func (_Irohaics20bank *Irohaics20bankFilterer) ParseMintRequested(log types.Log) (*Irohaics20bankMintRequested, error) {
	event := new(Irohaics20bankMintRequested)
	if err := _Irohaics20bank.contract.UnpackLog(event, "MintRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Irohaics20bankRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Irohaics20bank contract.
type Irohaics20bankRoleAdminChangedIterator struct {
	Event *Irohaics20bankRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *Irohaics20bankRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Irohaics20bankRoleAdminChanged)
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
		it.Event = new(Irohaics20bankRoleAdminChanged)
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
func (it *Irohaics20bankRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Irohaics20bankRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Irohaics20bankRoleAdminChanged represents a RoleAdminChanged event raised by the Irohaics20bank contract.
type Irohaics20bankRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Irohaics20bank *Irohaics20bankFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Irohaics20bankRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Irohaics20bank.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Irohaics20bankRoleAdminChangedIterator{contract: _Irohaics20bank.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Irohaics20bank *Irohaics20bankFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Irohaics20bankRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Irohaics20bank.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Irohaics20bankRoleAdminChanged)
				if err := _Irohaics20bank.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Irohaics20bank *Irohaics20bankFilterer) ParseRoleAdminChanged(log types.Log) (*Irohaics20bankRoleAdminChanged, error) {
	event := new(Irohaics20bankRoleAdminChanged)
	if err := _Irohaics20bank.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Irohaics20bankRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Irohaics20bank contract.
type Irohaics20bankRoleGrantedIterator struct {
	Event *Irohaics20bankRoleGranted // Event containing the contract specifics and raw log

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
func (it *Irohaics20bankRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Irohaics20bankRoleGranted)
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
		it.Event = new(Irohaics20bankRoleGranted)
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
func (it *Irohaics20bankRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Irohaics20bankRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Irohaics20bankRoleGranted represents a RoleGranted event raised by the Irohaics20bank contract.
type Irohaics20bankRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Irohaics20bank *Irohaics20bankFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Irohaics20bankRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Irohaics20bank.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Irohaics20bankRoleGrantedIterator{contract: _Irohaics20bank.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Irohaics20bank *Irohaics20bankFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Irohaics20bankRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Irohaics20bank.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Irohaics20bankRoleGranted)
				if err := _Irohaics20bank.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Irohaics20bank *Irohaics20bankFilterer) ParseRoleGranted(log types.Log) (*Irohaics20bankRoleGranted, error) {
	event := new(Irohaics20bankRoleGranted)
	if err := _Irohaics20bank.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Irohaics20bankRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Irohaics20bank contract.
type Irohaics20bankRoleRevokedIterator struct {
	Event *Irohaics20bankRoleRevoked // Event containing the contract specifics and raw log

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
func (it *Irohaics20bankRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Irohaics20bankRoleRevoked)
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
		it.Event = new(Irohaics20bankRoleRevoked)
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
func (it *Irohaics20bankRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Irohaics20bankRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Irohaics20bankRoleRevoked represents a RoleRevoked event raised by the Irohaics20bank contract.
type Irohaics20bankRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Irohaics20bank *Irohaics20bankFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Irohaics20bankRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Irohaics20bank.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Irohaics20bankRoleRevokedIterator{contract: _Irohaics20bank.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Irohaics20bank *Irohaics20bankFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Irohaics20bankRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Irohaics20bank.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Irohaics20bankRoleRevoked)
				if err := _Irohaics20bank.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Irohaics20bank *Irohaics20bankFilterer) ParseRoleRevoked(log types.Log) (*Irohaics20bankRoleRevoked, error) {
	event := new(Irohaics20bankRoleRevoked)
	if err := _Irohaics20bank.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
