// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package irohaics20transferbank

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

// ChannelCounterpartyData is an auto generated low-level Go binding around an user-defined struct.
type ChannelCounterpartyData struct {
	PortId    string
	ChannelId string
}

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// PacketData is an auto generated low-level Go binding around an user-defined struct.
type PacketData struct {
	Sequence           uint64
	SourcePort         string
	SourceChannel      string
	DestinationPort    string
	DestinationChannel string
	Data               []byte
	TimeoutHeight      HeightData
	TimeoutTimestamp   uint64
}

// Irohaics20transferbankABI is the input ABI used to generate the binding from.
const Irohaics20transferbankABI = "[{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host_\",\"type\":\"address\"},{\"internalType\":\"contractIBCHandler\",\"name\":\"ibcHandler_\",\"type\":\"address\"},{\"internalType\":\"contractIIrohaICS20Bank\",\"name\":\"bank_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"onAcknowledgementPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"onChanCloseConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"onChanCloseInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"counterpartyVersion\",\"type\":\"string\"}],\"name\":\"onChanOpenAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"onChanOpenConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumChannel.Order\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"onChanOpenInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumChannel.Order\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"onChanOpenTry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"source_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"source_channel\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination_channel\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"timeout_height\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"timeout_timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPacket.Data\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"onRecvPacket\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"srcAccountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destAccountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"amount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourcePort\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChannel\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"timeoutHeight\",\"type\":\"uint64\"}],\"name\":\"sendTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Irohaics20transferbank is an auto generated Go binding around an Ethereum contract.
type Irohaics20transferbank struct {
	Irohaics20transferbankCaller     // Read-only binding to the contract
	Irohaics20transferbankTransactor // Write-only binding to the contract
	Irohaics20transferbankFilterer   // Log filterer for contract events
}

// Irohaics20transferbankCaller is an auto generated read-only Go binding around an Ethereum contract.
type Irohaics20transferbankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Irohaics20transferbankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Irohaics20transferbankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Irohaics20transferbankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Irohaics20transferbankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Irohaics20transferbankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Irohaics20transferbankSession struct {
	Contract     *Irohaics20transferbank // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Irohaics20transferbankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Irohaics20transferbankCallerSession struct {
	Contract *Irohaics20transferbankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// Irohaics20transferbankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Irohaics20transferbankTransactorSession struct {
	Contract     *Irohaics20transferbankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// Irohaics20transferbankRaw is an auto generated low-level Go binding around an Ethereum contract.
type Irohaics20transferbankRaw struct {
	Contract *Irohaics20transferbank // Generic contract binding to access the raw methods on
}

// Irohaics20transferbankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Irohaics20transferbankCallerRaw struct {
	Contract *Irohaics20transferbankCaller // Generic read-only contract binding to access the raw methods on
}

// Irohaics20transferbankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Irohaics20transferbankTransactorRaw struct {
	Contract *Irohaics20transferbankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIrohaics20transferbank creates a new instance of Irohaics20transferbank, bound to a specific deployed contract.
func NewIrohaics20transferbank(address common.Address, backend bind.ContractBackend) (*Irohaics20transferbank, error) {
	contract, err := bindIrohaics20transferbank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Irohaics20transferbank{Irohaics20transferbankCaller: Irohaics20transferbankCaller{contract: contract}, Irohaics20transferbankTransactor: Irohaics20transferbankTransactor{contract: contract}, Irohaics20transferbankFilterer: Irohaics20transferbankFilterer{contract: contract}}, nil
}

// NewIrohaics20transferbankCaller creates a new read-only instance of Irohaics20transferbank, bound to a specific deployed contract.
func NewIrohaics20transferbankCaller(address common.Address, caller bind.ContractCaller) (*Irohaics20transferbankCaller, error) {
	contract, err := bindIrohaics20transferbank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Irohaics20transferbankCaller{contract: contract}, nil
}

// NewIrohaics20transferbankTransactor creates a new write-only instance of Irohaics20transferbank, bound to a specific deployed contract.
func NewIrohaics20transferbankTransactor(address common.Address, transactor bind.ContractTransactor) (*Irohaics20transferbankTransactor, error) {
	contract, err := bindIrohaics20transferbank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Irohaics20transferbankTransactor{contract: contract}, nil
}

// NewIrohaics20transferbankFilterer creates a new log filterer instance of Irohaics20transferbank, bound to a specific deployed contract.
func NewIrohaics20transferbankFilterer(address common.Address, filterer bind.ContractFilterer) (*Irohaics20transferbankFilterer, error) {
	contract, err := bindIrohaics20transferbank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Irohaics20transferbankFilterer{contract: contract}, nil
}

// bindIrohaics20transferbank binds a generic wrapper to an already deployed contract.
func bindIrohaics20transferbank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Irohaics20transferbankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Irohaics20transferbank *Irohaics20transferbankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Irohaics20transferbank.Contract.Irohaics20transferbankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Irohaics20transferbank *Irohaics20transferbankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.Irohaics20transferbankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Irohaics20transferbank *Irohaics20transferbankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.Irohaics20transferbankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Irohaics20transferbank *Irohaics20transferbankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Irohaics20transferbank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Irohaics20transferbank *Irohaics20transferbankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Irohaics20transferbank *Irohaics20transferbankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.contract.Transact(opts, method, params...)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xda7b08a7.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet PacketData, acknowledgement []byte) (*types.Transaction, error) {
	return _Irohaics20transferbank.contract.Transact(opts, "onAcknowledgementPacket", packet, acknowledgement)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xda7b08a7.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement) returns()
func (_Irohaics20transferbank *Irohaics20transferbankSession) OnAcknowledgementPacket(packet PacketData, acknowledgement []byte) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnAcknowledgementPacket(&_Irohaics20transferbank.TransactOpts, packet, acknowledgement)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0xda7b08a7.
//
// Solidity: function onAcknowledgementPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet, bytes acknowledgement) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactorSession) OnAcknowledgementPacket(packet PacketData, acknowledgement []byte) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnAcknowledgementPacket(&_Irohaics20transferbank.TransactOpts, packet, acknowledgement)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0xef4776d2.
//
// Solidity: function onChanCloseConfirm(string portId, string channelId) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, portId string, channelId string) (*types.Transaction, error) {
	return _Irohaics20transferbank.contract.Transact(opts, "onChanCloseConfirm", portId, channelId)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0xef4776d2.
//
// Solidity: function onChanCloseConfirm(string portId, string channelId) returns()
func (_Irohaics20transferbank *Irohaics20transferbankSession) OnChanCloseConfirm(portId string, channelId string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanCloseConfirm(&_Irohaics20transferbank.TransactOpts, portId, channelId)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0xef4776d2.
//
// Solidity: function onChanCloseConfirm(string portId, string channelId) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactorSession) OnChanCloseConfirm(portId string, channelId string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanCloseConfirm(&_Irohaics20transferbank.TransactOpts, portId, channelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0xe74a1ac2.
//
// Solidity: function onChanCloseInit(string portId, string channelId) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactor) OnChanCloseInit(opts *bind.TransactOpts, portId string, channelId string) (*types.Transaction, error) {
	return _Irohaics20transferbank.contract.Transact(opts, "onChanCloseInit", portId, channelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0xe74a1ac2.
//
// Solidity: function onChanCloseInit(string portId, string channelId) returns()
func (_Irohaics20transferbank *Irohaics20transferbankSession) OnChanCloseInit(portId string, channelId string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanCloseInit(&_Irohaics20transferbank.TransactOpts, portId, channelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0xe74a1ac2.
//
// Solidity: function onChanCloseInit(string portId, string channelId) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactorSession) OnChanCloseInit(portId string, channelId string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanCloseInit(&_Irohaics20transferbank.TransactOpts, portId, channelId)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x4942d1ac.
//
// Solidity: function onChanOpenAck(string portId, string channelId, string counterpartyVersion) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactor) OnChanOpenAck(opts *bind.TransactOpts, portId string, channelId string, counterpartyVersion string) (*types.Transaction, error) {
	return _Irohaics20transferbank.contract.Transact(opts, "onChanOpenAck", portId, channelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x4942d1ac.
//
// Solidity: function onChanOpenAck(string portId, string channelId, string counterpartyVersion) returns()
func (_Irohaics20transferbank *Irohaics20transferbankSession) OnChanOpenAck(portId string, channelId string, counterpartyVersion string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanOpenAck(&_Irohaics20transferbank.TransactOpts, portId, channelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0x4942d1ac.
//
// Solidity: function onChanOpenAck(string portId, string channelId, string counterpartyVersion) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactorSession) OnChanOpenAck(portId string, channelId string, counterpartyVersion string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanOpenAck(&_Irohaics20transferbank.TransactOpts, portId, channelId, counterpartyVersion)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xa113e411.
//
// Solidity: function onChanOpenConfirm(string portId, string channelId) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, portId string, channelId string) (*types.Transaction, error) {
	return _Irohaics20transferbank.contract.Transact(opts, "onChanOpenConfirm", portId, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xa113e411.
//
// Solidity: function onChanOpenConfirm(string portId, string channelId) returns()
func (_Irohaics20transferbank *Irohaics20transferbankSession) OnChanOpenConfirm(portId string, channelId string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanOpenConfirm(&_Irohaics20transferbank.TransactOpts, portId, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xa113e411.
//
// Solidity: function onChanOpenConfirm(string portId, string channelId) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactorSession) OnChanOpenConfirm(portId string, channelId string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanOpenConfirm(&_Irohaics20transferbank.TransactOpts, portId, channelId)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x44dd9638.
//
// Solidity: function onChanOpenInit(uint8 , string[] , string , string channelId, (string,string) , string ) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactor) OnChanOpenInit(opts *bind.TransactOpts, arg0 uint8, arg1 []string, arg2 string, channelId string, arg4 ChannelCounterpartyData, arg5 string) (*types.Transaction, error) {
	return _Irohaics20transferbank.contract.Transact(opts, "onChanOpenInit", arg0, arg1, arg2, channelId, arg4, arg5)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x44dd9638.
//
// Solidity: function onChanOpenInit(uint8 , string[] , string , string channelId, (string,string) , string ) returns()
func (_Irohaics20transferbank *Irohaics20transferbankSession) OnChanOpenInit(arg0 uint8, arg1 []string, arg2 string, channelId string, arg4 ChannelCounterpartyData, arg5 string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanOpenInit(&_Irohaics20transferbank.TransactOpts, arg0, arg1, arg2, channelId, arg4, arg5)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x44dd9638.
//
// Solidity: function onChanOpenInit(uint8 , string[] , string , string channelId, (string,string) , string ) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactorSession) OnChanOpenInit(arg0 uint8, arg1 []string, arg2 string, channelId string, arg4 ChannelCounterpartyData, arg5 string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanOpenInit(&_Irohaics20transferbank.TransactOpts, arg0, arg1, arg2, channelId, arg4, arg5)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x981389f2.
//
// Solidity: function onChanOpenTry(uint8 , string[] , string , string channelId, (string,string) , string , string ) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactor) OnChanOpenTry(opts *bind.TransactOpts, arg0 uint8, arg1 []string, arg2 string, channelId string, arg4 ChannelCounterpartyData, arg5 string, arg6 string) (*types.Transaction, error) {
	return _Irohaics20transferbank.contract.Transact(opts, "onChanOpenTry", arg0, arg1, arg2, channelId, arg4, arg5, arg6)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x981389f2.
//
// Solidity: function onChanOpenTry(uint8 , string[] , string , string channelId, (string,string) , string , string ) returns()
func (_Irohaics20transferbank *Irohaics20transferbankSession) OnChanOpenTry(arg0 uint8, arg1 []string, arg2 string, channelId string, arg4 ChannelCounterpartyData, arg5 string, arg6 string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanOpenTry(&_Irohaics20transferbank.TransactOpts, arg0, arg1, arg2, channelId, arg4, arg5, arg6)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x981389f2.
//
// Solidity: function onChanOpenTry(uint8 , string[] , string , string channelId, (string,string) , string , string ) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactorSession) OnChanOpenTry(arg0 uint8, arg1 []string, arg2 string, channelId string, arg4 ChannelCounterpartyData, arg5 string, arg6 string) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnChanOpenTry(&_Irohaics20transferbank.TransactOpts, arg0, arg1, arg2, channelId, arg4, arg5, arg6)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x5550b656.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns(bytes acknowledgement)
func (_Irohaics20transferbank *Irohaics20transferbankTransactor) OnRecvPacket(opts *bind.TransactOpts, packet PacketData) (*types.Transaction, error) {
	return _Irohaics20transferbank.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x5550b656.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns(bytes acknowledgement)
func (_Irohaics20transferbank *Irohaics20transferbankSession) OnRecvPacket(packet PacketData) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnRecvPacket(&_Irohaics20transferbank.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x5550b656.
//
// Solidity: function onRecvPacket((uint64,string,string,string,string,bytes,(uint64,uint64),uint64) packet) returns(bytes acknowledgement)
func (_Irohaics20transferbank *Irohaics20transferbankTransactorSession) OnRecvPacket(packet PacketData) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.OnRecvPacket(&_Irohaics20transferbank.TransactOpts, packet)
}

// SendTransfer is a paid mutator transaction binding the contract method 0xaec17f99.
//
// Solidity: function sendTransfer(string srcAccountId, string destAccountId, string assetId, string description, string amount, string sourcePort, string sourceChannel, uint64 timeoutHeight) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactor) SendTransfer(opts *bind.TransactOpts, srcAccountId string, destAccountId string, assetId string, description string, amount string, sourcePort string, sourceChannel string, timeoutHeight uint64) (*types.Transaction, error) {
	return _Irohaics20transferbank.contract.Transact(opts, "sendTransfer", srcAccountId, destAccountId, assetId, description, amount, sourcePort, sourceChannel, timeoutHeight)
}

// SendTransfer is a paid mutator transaction binding the contract method 0xaec17f99.
//
// Solidity: function sendTransfer(string srcAccountId, string destAccountId, string assetId, string description, string amount, string sourcePort, string sourceChannel, uint64 timeoutHeight) returns()
func (_Irohaics20transferbank *Irohaics20transferbankSession) SendTransfer(srcAccountId string, destAccountId string, assetId string, description string, amount string, sourcePort string, sourceChannel string, timeoutHeight uint64) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.SendTransfer(&_Irohaics20transferbank.TransactOpts, srcAccountId, destAccountId, assetId, description, amount, sourcePort, sourceChannel, timeoutHeight)
}

// SendTransfer is a paid mutator transaction binding the contract method 0xaec17f99.
//
// Solidity: function sendTransfer(string srcAccountId, string destAccountId, string assetId, string description, string amount, string sourcePort, string sourceChannel, uint64 timeoutHeight) returns()
func (_Irohaics20transferbank *Irohaics20transferbankTransactorSession) SendTransfer(srcAccountId string, destAccountId string, assetId string, description string, amount string, sourcePort string, sourceChannel string, timeoutHeight uint64) (*types.Transaction, error) {
	return _Irohaics20transferbank.Contract.SendTransfer(&_Irohaics20transferbank.TransactOpts, srcAccountId, destAccountId, assetId, description, amount, sourcePort, sourceChannel, timeoutHeight)
}
