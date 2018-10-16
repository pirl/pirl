

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

// RedListABI is the input ABI used to generate the binding from.
const RedListABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getLastUpdate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isContainsAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_frodulentAccount\",\"type\":\"address\"}],\"name\":\"addNewAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"AddressAddedToRedlist\",\"type\":\"event\"}]"

// RedList is an auto generated Go binding around an Ethereum contract.
type RedList struct {
	RedListCaller     // Read-only binding to the contract
	RedListTransactor // Write-only binding to the contract
	RedListFilterer   // Log filterer for contract events
}

// RedListCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedListSession struct {
	Contract     *RedList          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedListCallerSession struct {
	Contract *RedListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RedListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedListTransactorSession struct {
	Contract     *RedListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RedListRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedListRaw struct {
	Contract *RedList // Generic contract binding to access the raw methods on
}

// RedListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedListCallerRaw struct {
	Contract *RedListCaller // Generic read-only contract binding to access the raw methods on
}

// RedListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedListTransactorRaw struct {
	Contract *RedListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedList creates a new instance of RedList, bound to a specific deployed contract.
func NewRedList(address common.Address, backend bind.ContractBackend) (*RedList, error) {
	contract, err := bindRedList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RedList{RedListCaller: RedListCaller{contract: contract}, RedListTransactor: RedListTransactor{contract: contract}, RedListFilterer: RedListFilterer{contract: contract}}, nil
}

// NewRedListCaller creates a new read-only instance of RedList, bound to a specific deployed contract.
func NewRedListCaller(address common.Address, caller bind.ContractCaller) (*RedListCaller, error) {
	contract, err := bindRedList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedListCaller{contract: contract}, nil
}

// NewRedListTransactor creates a new write-only instance of RedList, bound to a specific deployed contract.
func NewRedListTransactor(address common.Address, transactor bind.ContractTransactor) (*RedListTransactor, error) {
	contract, err := bindRedList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedListTransactor{contract: contract}, nil
}

// NewRedListFilterer creates a new log filterer instance of RedList, bound to a specific deployed contract.
func NewRedListFilterer(address common.Address, filterer bind.ContractFilterer) (*RedListFilterer, error) {
	contract, err := bindRedList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedListFilterer{contract: contract}, nil
}

// bindRedList binds a generic wrapper to an already deployed contract.
func bindRedList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RedListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedList *RedListRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RedList.Contract.RedListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedList *RedListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedList.Contract.RedListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedList *RedListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedList.Contract.RedListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedList *RedListCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RedList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedList *RedListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedList *RedListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedList.Contract.contract.Transact(opts, method, params...)
}

// GetAllAddresses is a free data retrieval call binding the contract method 0x9516a104.
//
// Solidity: function getAllAddresses() constant returns(address[])
func (_RedList *RedListCaller) GetAllAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _RedList.contract.Call(opts, out, "getAllAddresses")
	return *ret0, err
}

// GetAllAddresses is a free data retrieval call binding the contract method 0x9516a104.
//
// Solidity: function getAllAddresses() constant returns(address[])
func (_RedList *RedListSession) GetAllAddresses() ([]common.Address, error) {
	return _RedList.Contract.GetAllAddresses(&_RedList.CallOpts)
}

// GetAllAddresses is a free data retrieval call binding the contract method 0x9516a104.
//
// Solidity: function getAllAddresses() constant returns(address[])
func (_RedList *RedListCallerSession) GetAllAddresses() ([]common.Address, error) {
	return _RedList.Contract.GetAllAddresses(&_RedList.CallOpts)
}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() constant returns(uint256)
func (_RedList *RedListCaller) GetLastUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RedList.contract.Call(opts, out, "getLastUpdate")
	return *ret0, err
}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() constant returns(uint256)
func (_RedList *RedListSession) GetLastUpdate() (*big.Int, error) {
	return _RedList.Contract.GetLastUpdate(&_RedList.CallOpts)
}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() constant returns(uint256)
func (_RedList *RedListCallerSession) GetLastUpdate() (*big.Int, error) {
	return _RedList.Contract.GetLastUpdate(&_RedList.CallOpts)
}

// IsContainsAddress is a free data retrieval call binding the contract method 0x5b687a61.
//
// Solidity: function isContainsAddress(_address address) constant returns(bool)
func (_RedList *RedListCaller) IsContainsAddress(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RedList.contract.Call(opts, out, "isContainsAddress", _address)
	return *ret0, err
}

// IsContainsAddress is a free data retrieval call binding the contract method 0x5b687a61.
//
// Solidity: function isContainsAddress(_address address) constant returns(bool)
func (_RedList *RedListSession) IsContainsAddress(_address common.Address) (bool, error) {
	return _RedList.Contract.IsContainsAddress(&_RedList.CallOpts, _address)
}

// IsContainsAddress is a free data retrieval call binding the contract method 0x5b687a61.
//
// Solidity: function isContainsAddress(_address address) constant returns(bool)
func (_RedList *RedListCallerSession) IsContainsAddress(_address common.Address) (bool, error) {
	return _RedList.Contract.IsContainsAddress(&_RedList.CallOpts, _address)
}

// AddNewAddress is a paid mutator transaction binding the contract method 0xa2f35f44.
//
// Solidity: function addNewAddress(_frodulentAccount address) returns(bool)
func (_RedList *RedListTransactor) AddNewAddress(opts *bind.TransactOpts, _frodulentAccount common.Address) (*types.Transaction, error) {
	return _RedList.contract.Transact(opts, "addNewAddress", _frodulentAccount)
}

// AddNewAddress is a paid mutator transaction binding the contract method 0xa2f35f44.
//
// Solidity: function addNewAddress(_frodulentAccount address) returns(bool)
func (_RedList *RedListSession) AddNewAddress(_frodulentAccount common.Address) (*types.Transaction, error) {
	return _RedList.Contract.AddNewAddress(&_RedList.TransactOpts, _frodulentAccount)
}

// AddNewAddress is a paid mutator transaction binding the contract method 0xa2f35f44.
//
// Solidity: function addNewAddress(_frodulentAccount address) returns(bool)
func (_RedList *RedListTransactorSession) AddNewAddress(_frodulentAccount common.Address) (*types.Transaction, error) {
	return _RedList.Contract.AddNewAddress(&_RedList.TransactOpts, _frodulentAccount)
}

// RedListAddressAddedToRedlistIterator is returned from FilterAddressAddedToRedlist and is used to iterate over the raw logs and unpacked data for AddressAddedToRedlist events raised by the RedList contract.
type RedListAddressAddedToRedlistIterator struct {
	Event *RedListAddressAddedToRedlist // Event containing the contract specifics and raw log

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
func (it *RedListAddressAddedToRedlistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedListAddressAddedToRedlist)
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
		it.Event = new(RedListAddressAddedToRedlist)
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
func (it *RedListAddressAddedToRedlistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedListAddressAddedToRedlistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedListAddressAddedToRedlist represents a AddressAddedToRedlist event raised by the RedList contract.
type RedListAddressAddedToRedlist struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressAddedToRedlist is a free log retrieval operation binding the contract event 0x2f203a3c141c116c8d0bf5c15a359f66807e949137f7a91ef8fca84331389f75.
//
// Solidity: event AddressAddedToRedlist(_address address)
func (_RedList *RedListFilterer) FilterAddressAddedToRedlist(opts *bind.FilterOpts) (*RedListAddressAddedToRedlistIterator, error) {

	logs, sub, err := _RedList.contract.FilterLogs(opts, "AddressAddedToRedlist")
	if err != nil {
		return nil, err
	}
	return &RedListAddressAddedToRedlistIterator{contract: _RedList.contract, event: "AddressAddedToRedlist", logs: logs, sub: sub}, nil
}

// WatchAddressAddedToRedlist is a free log subscription operation binding the contract event 0x2f203a3c141c116c8d0bf5c15a359f66807e949137f7a91ef8fca84331389f75.
//
// Solidity: event AddressAddedToRedlist(_address address)
func (_RedList *RedListFilterer) WatchAddressAddedToRedlist(opts *bind.WatchOpts, sink chan<- *RedListAddressAddedToRedlist) (event.Subscription, error) {

	logs, sub, err := _RedList.contract.WatchLogs(opts, "AddressAddedToRedlist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedListAddressAddedToRedlist)
				if err := _RedList.contract.UnpackLog(event, "AddressAddedToRedlist", log); err != nil {
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
