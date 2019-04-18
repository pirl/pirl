// Copyright 2014 The go-ethereum Authors
// Copyright 2018 Pirl Sprl
// This file is part of the go-ethereum library modified with Pirl Security Protocol.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see http://www.gnu.org/licenses/.
// Package core implements the Ethereum consensus protocol modified with Pirl Security Protocol.

package redlist

import (
	"math/big"
	"git.pirl.io/community/pirl/accounts/abi"
	"git.pirl.io/community/pirl/accounts/abi/bind"
	"git.pirl.io/community/pirl/common"
	"git.pirl.io/community/pirl/core/types"
	"git.pirl.io/community/pirl/event"
	"strings"

	ethereum "git.pirl.io/community/pirl"
)

// RedlistABI is the input ABI used to generate the binding from.
const RedlistABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_frodulentAccount\",\"type\":\"address\"}],\"name\":\"addNewAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"AddressAddedToRedlist\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastUpdate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isContainsAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RedlistBin is the compiled bytecode used for deploying new contracts.
const RedlistBin = `6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561005057600080fd5b50610509806100606000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806341c0e1b5146100725780634c89867f146100895780635b687a61146100b45780639516a1041461010f578063a2f35f441461017b575b600080fd5b34801561007e57600080fd5b506100876101d6565b005b34801561009557600080fd5b5061009e610267565b6040518082815260200191505060405180910390f35b3480156100c057600080fd5b506100f5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610271565b604051808215151515815260200191505060405180910390f35b34801561011b57600080fd5b506101246102c7565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561016757808201518184015260208101905061014c565b505050509050019250505060405180910390f35b34801561018757600080fd5b506101bc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610355565b604051808215151515815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610265576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b565b6000600354905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6060600280548060200260200160405190810160405280929190818152602001828054801561034b57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610301575b5050505050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156104d85760018060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060028290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050426003819055507f2f203a3c141c116c8d0bf5c15a359f66807e949137f7a91ef8fca84331389f7582604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1600190505b9190505600a165627a7a723058204b18cefbdd4ee262f31affea829c8c25bcec9615dbb730def81dd91bba6b41950029`

// DeployRedlist deploys a new Ethereum contract, binding an instance of Redlist to it.
func DeployRedlist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Redlist, error) {
	parsed, err := abi.JSON(strings.NewReader(RedlistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RedlistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Redlist{RedlistCaller: RedlistCaller{contract: contract}, RedlistTransactor: RedlistTransactor{contract: contract}, RedlistFilterer: RedlistFilterer{contract: contract}}, nil
}

// Redlist is an auto generated Go binding around an Ethereum contract.
type Redlist struct {
	RedlistCaller     // Read-only binding to the contract
	RedlistTransactor // Write-only binding to the contract
	RedlistFilterer   // Log filterer for contract events
}

// RedlistCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedlistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedlistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedlistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedlistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedlistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedlistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedlistSession struct {
	Contract     *Redlist          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedlistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedlistCallerSession struct {
	Contract *RedlistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RedlistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedlistTransactorSession struct {
	Contract     *RedlistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RedlistRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedlistRaw struct {
	Contract *Redlist // Generic contract binding to access the raw methods on
}

// RedlistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedlistCallerRaw struct {
	Contract *RedlistCaller // Generic read-only contract binding to access the raw methods on
}

// RedlistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedlistTransactorRaw struct {
	Contract *RedlistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedlist creates a new instance of Redlist, bound to a specific deployed contract.
func NewRedlist(address common.Address, backend bind.ContractBackend) (*Redlist, error) {
	contract, err := bindRedlist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Redlist{RedlistCaller: RedlistCaller{contract: contract}, RedlistTransactor: RedlistTransactor{contract: contract}, RedlistFilterer: RedlistFilterer{contract: contract}}, nil
}

// NewRedlistCaller creates a new read-only instance of Redlist, bound to a specific deployed contract.
func NewRedlistCaller(address common.Address, caller bind.ContractCaller) (*RedlistCaller, error) {
	contract, err := bindRedlist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedlistCaller{contract: contract}, nil
}

// NewRedlistTransactor creates a new write-only instance of Redlist, bound to a specific deployed contract.
func NewRedlistTransactor(address common.Address, transactor bind.ContractTransactor) (*RedlistTransactor, error) {
	contract, err := bindRedlist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedlistTransactor{contract: contract}, nil
}

// NewRedlistFilterer creates a new log filterer instance of Redlist, bound to a specific deployed contract.
func NewRedlistFilterer(address common.Address, filterer bind.ContractFilterer) (*RedlistFilterer, error) {
	contract, err := bindRedlist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedlistFilterer{contract: contract}, nil
}

// bindRedlist binds a generic wrapper to an already deployed contract.
func bindRedlist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RedlistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redlist *RedlistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Redlist.Contract.RedlistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redlist *RedlistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redlist.Contract.RedlistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redlist *RedlistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redlist.Contract.RedlistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redlist *RedlistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Redlist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redlist *RedlistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redlist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redlist *RedlistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redlist.Contract.contract.Transact(opts, method, params...)
}

// GetAllAddresses is a free data retrieval call binding the contract method 0x9516a104.
//
// Solidity: function getAllAddresses() constant returns(address[])
func (_Redlist *RedlistCaller) GetAllAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Redlist.contract.Call(opts, out, "getAllAddresses")
	return *ret0, err
}

// GetAllAddresses is a free data retrieval call binding the contract method 0x9516a104.
//
// Solidity: function getAllAddresses() constant returns(address[])
func (_Redlist *RedlistSession) GetAllAddresses() ([]common.Address, error) {
	return _Redlist.Contract.GetAllAddresses(&_Redlist.CallOpts)
}

// GetAllAddresses is a free data retrieval call binding the contract method 0x9516a104.
//
// Solidity: function getAllAddresses() constant returns(address[])
func (_Redlist *RedlistCallerSession) GetAllAddresses() ([]common.Address, error) {
	return _Redlist.Contract.GetAllAddresses(&_Redlist.CallOpts)
}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() constant returns(uint256)
func (_Redlist *RedlistCaller) GetLastUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Redlist.contract.Call(opts, out, "getLastUpdate")
	return *ret0, err
}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() constant returns(uint256)
func (_Redlist *RedlistSession) GetLastUpdate() (*big.Int, error) {
	return _Redlist.Contract.GetLastUpdate(&_Redlist.CallOpts)
}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() constant returns(uint256)
func (_Redlist *RedlistCallerSession) GetLastUpdate() (*big.Int, error) {
	return _Redlist.Contract.GetLastUpdate(&_Redlist.CallOpts)
}

// IsContainsAddress is a free data retrieval call binding the contract method 0x5b687a61.
//
// Solidity: function isContainsAddress(_address address) constant returns(bool)
func (_Redlist *RedlistCaller) IsContainsAddress(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Redlist.contract.Call(opts, out, "isContainsAddress", _address)
	return *ret0, err
}

// IsContainsAddress is a free data retrieval call binding the contract method 0x5b687a61.
//
// Solidity: function isContainsAddress(_address address) constant returns(bool)
func (_Redlist *RedlistSession) IsContainsAddress(_address common.Address) (bool, error) {
	return _Redlist.Contract.IsContainsAddress(&_Redlist.CallOpts, _address)
}

// IsContainsAddress is a free data retrieval call binding the contract method 0x5b687a61.
//
// Solidity: function isContainsAddress(_address address) constant returns(bool)
func (_Redlist *RedlistCallerSession) IsContainsAddress(_address common.Address) (bool, error) {
	return _Redlist.Contract.IsContainsAddress(&_Redlist.CallOpts, _address)
}

// AddNewAddress is a paid mutator transaction binding the contract method 0xa2f35f44.
//
// Solidity: function addNewAddress(_frodulentAccount address) returns(bool)
func (_Redlist *RedlistTransactor) AddNewAddress(opts *bind.TransactOpts, _frodulentAccount common.Address) (*types.Transaction, error) {
	return _Redlist.contract.Transact(opts, "addNewAddress", _frodulentAccount)
}

// AddNewAddress is a paid mutator transaction binding the contract method 0xa2f35f44.
//
// Solidity: function addNewAddress(_frodulentAccount address) returns(bool)
func (_Redlist *RedlistSession) AddNewAddress(_frodulentAccount common.Address) (*types.Transaction, error) {
	return _Redlist.Contract.AddNewAddress(&_Redlist.TransactOpts, _frodulentAccount)
}

// AddNewAddress is a paid mutator transaction binding the contract method 0xa2f35f44.
//
// Solidity: function addNewAddress(_frodulentAccount address) returns(bool)
func (_Redlist *RedlistTransactorSession) AddNewAddress(_frodulentAccount common.Address) (*types.Transaction, error) {
	return _Redlist.Contract.AddNewAddress(&_Redlist.TransactOpts, _frodulentAccount)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Redlist *RedlistTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redlist.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Redlist *RedlistSession) Kill() (*types.Transaction, error) {
	return _Redlist.Contract.Kill(&_Redlist.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Redlist *RedlistTransactorSession) Kill() (*types.Transaction, error) {
	return _Redlist.Contract.Kill(&_Redlist.TransactOpts)
}

// RedlistAddressAddedToRedlistIterator is returned from FilterAddressAddedToRedlist and is used to iterate over the raw logs and unpacked data for AddressAddedToRedlist events raised by the Redlist contract.
type RedlistAddressAddedToRedlistIterator struct {
	Event *RedlistAddressAddedToRedlist // Event containing the contract specifics and raw log

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
func (it *RedlistAddressAddedToRedlistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedlistAddressAddedToRedlist)
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
		it.Event = new(RedlistAddressAddedToRedlist)
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
func (it *RedlistAddressAddedToRedlistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedlistAddressAddedToRedlistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedlistAddressAddedToRedlist represents a AddressAddedToRedlist event raised by the Redlist contract.
type RedlistAddressAddedToRedlist struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressAddedToRedlist is a free log retrieval operation binding the contract event 0x2f203a3c141c116c8d0bf5c15a359f66807e949137f7a91ef8fca84331389f75.
//
// Solidity: event AddressAddedToRedlist(_address address)
func (_Redlist *RedlistFilterer) FilterAddressAddedToRedlist(opts *bind.FilterOpts) (*RedlistAddressAddedToRedlistIterator, error) {

	logs, sub, err := _Redlist.contract.FilterLogs(opts, "AddressAddedToRedlist")
	if err != nil {
		return nil, err
	}
	return &RedlistAddressAddedToRedlistIterator{contract: _Redlist.contract, event: "AddressAddedToRedlist", logs: logs, sub: sub}, nil
}

// WatchAddressAddedToRedlist is a free log subscription operation binding the contract event 0x2f203a3c141c116c8d0bf5c15a359f66807e949137f7a91ef8fca84331389f75.
//
// Solidity: event AddressAddedToRedlist(_address address)
func (_Redlist *RedlistFilterer) WatchAddressAddedToRedlist(opts *bind.WatchOpts, sink chan<- *RedlistAddressAddedToRedlist) (event.Subscription, error) {

	logs, sub, err := _Redlist.contract.WatchLogs(opts, "AddressAddedToRedlist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedlistAddressAddedToRedlist)
				if err := _Redlist.contract.UnpackLog(event, "AddressAddedToRedlist", log); err != nil {
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
