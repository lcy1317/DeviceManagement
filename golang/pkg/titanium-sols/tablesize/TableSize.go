// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tablesize

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

// TableSizeABI is the input ABI used to generate the binding from.
const TableSizeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"tablename\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"size\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"CreateResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"InsertResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"UpdateResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"RemoveResult\",\"type\":\"event\"}]"

// TableSizeBin is the compiled bytecode used for deploying new contracts.
var TableSizeBin = "0x608060405234801561001057600080fd5b506110016000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610727806100626000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063d2b5213014610046575b600080fd5b34801561005257600080fd5b5061006d6004803603610068919081019061048d565b610083565b60405161007a919061054d565b60405180910390f35b6000806000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c9876040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016100ff9190610568565b602060405180830381600087803b15801561011957600080fd5b505af115801561012d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610151919081019061043b565b92508273ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156101b757600080fd5b505af11580156101cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506101ef91908101906103e9565b91508273ffffffffffffffffffffffffffffffffffffffff1663e8434e3986846040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161024892919061058a565b602060405180830381600087803b15801561026257600080fd5b505af1158015610276573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061029a9190810190610412565b90508073ffffffffffffffffffffffffffffffffffffffff1663949d225d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561030057600080fd5b505af1158015610314573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103389190810190610464565b935050505092915050565b600061034f8251610648565b905092915050565b6000610363825161065a565b905092915050565b6000610377825161066c565b905092915050565b600061038b825161067e565b905092915050565b600082601f83011215156103a657600080fd5b81356103b96103b4826105e7565b6105ba565b915080825260208301602083018583830111156103d557600080fd5b6103e083828461069a565b50505092915050565b6000602082840312156103fb57600080fd5b600061040984828501610343565b91505092915050565b60006020828403121561042457600080fd5b600061043284828501610357565b91505092915050565b60006020828403121561044d57600080fd5b600061045b8482850161036b565b91505092915050565b60006020828403121561047657600080fd5b60006104848482850161037f565b91505092915050565b600080604083850312156104a057600080fd5b600083013567ffffffffffffffff8111156104ba57600080fd5b6104c685828601610393565b925050602083013567ffffffffffffffff8111156104e357600080fd5b6104ef85828601610393565b9150509250929050565b61050281610688565b82525050565b6105118161063e565b82525050565b600061052282610613565b8084526105368160208601602086016106a9565b61053f816106dc565b602085010191505092915050565b60006020820190506105626000830184610508565b92915050565b600060208201905081810360008301526105828184610517565b905092915050565b600060408201905081810360008301526105a48185610517565b90506105b360208301846104f9565b9392505050565b6000604051905081810181811067ffffffffffffffff821117156105dd57600080fd5b8060405250919050565b600067ffffffffffffffff8211156105fe57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006106538261061e565b9050919050565b60006106658261061e565b9050919050565b60006106778261061e565b9050919050565b6000819050919050565b60006106938261061e565b9050919050565b82818337600083830152505050565b60005b838110156106c75780820151818401526020810190506106ac565b838111156106d6576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a723058200f6e3c94389a49cf5193ead36b5e2f0b882326adac5afc6df437e70003a437e76c6578706572696d656e74616cf50037"

// DeployTableSize deploys a new contract, binding an instance of TableSize to it.
func DeployTableSize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TableSize, error) {
	parsed, err := abi.JSON(strings.NewReader(TableSizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TableSizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TableSize{TableSizeCaller: TableSizeCaller{contract: contract}, TableSizeTransactor: TableSizeTransactor{contract: contract}, TableSizeFilterer: TableSizeFilterer{contract: contract}}, nil
}

func AsyncDeployTableSize(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TableSizeABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(TableSizeBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// TableSize is an auto generated Go binding around a Solidity contract.
type TableSize struct {
	TableSizeCaller     // Read-only binding to the contract
	TableSizeTransactor // Write-only binding to the contract
	TableSizeFilterer   // Log filterer for contract events
}

// TableSizeCaller is an auto generated read-only Go binding around a Solidity contract.
type TableSizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableSizeTransactor is an auto generated write-only Go binding around a Solidity contract.
type TableSizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableSizeFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TableSizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableSizeSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TableSizeSession struct {
	Contract     *TableSize        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TableSizeCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TableSizeCallerSession struct {
	Contract *TableSizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TableSizeTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TableSizeTransactorSession struct {
	Contract     *TableSizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TableSizeRaw is an auto generated low-level Go binding around a Solidity contract.
type TableSizeRaw struct {
	Contract *TableSize // Generic contract binding to access the raw methods on
}

// TableSizeCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TableSizeCallerRaw struct {
	Contract *TableSizeCaller // Generic read-only contract binding to access the raw methods on
}

// TableSizeTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TableSizeTransactorRaw struct {
	Contract *TableSizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTableSize creates a new instance of TableSize, bound to a specific deployed contract.
func NewTableSize(address common.Address, backend bind.ContractBackend) (*TableSize, error) {
	contract, err := bindTableSize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TableSize{TableSizeCaller: TableSizeCaller{contract: contract}, TableSizeTransactor: TableSizeTransactor{contract: contract}, TableSizeFilterer: TableSizeFilterer{contract: contract}}, nil
}

// NewTableSizeCaller creates a new read-only instance of TableSize, bound to a specific deployed contract.
func NewTableSizeCaller(address common.Address, caller bind.ContractCaller) (*TableSizeCaller, error) {
	contract, err := bindTableSize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TableSizeCaller{contract: contract}, nil
}

// NewTableSizeTransactor creates a new write-only instance of TableSize, bound to a specific deployed contract.
func NewTableSizeTransactor(address common.Address, transactor bind.ContractTransactor) (*TableSizeTransactor, error) {
	contract, err := bindTableSize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TableSizeTransactor{contract: contract}, nil
}

// NewTableSizeFilterer creates a new log filterer instance of TableSize, bound to a specific deployed contract.
func NewTableSizeFilterer(address common.Address, filterer bind.ContractFilterer) (*TableSizeFilterer, error) {
	contract, err := bindTableSize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TableSizeFilterer{contract: contract}, nil
}

// bindTableSize binds a generic wrapper to an already deployed contract.
func bindTableSize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TableSizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TableSize *TableSizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TableSize.Contract.TableSizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TableSize *TableSizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TableSize.Contract.TableSizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TableSize *TableSizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TableSize.Contract.TableSizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TableSize *TableSizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TableSize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TableSize *TableSizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TableSize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TableSize *TableSizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TableSize.Contract.contract.Transact(opts, method, params...)
}

// Size is a free data retrieval call binding the contract method 0xd2b52130.
//
// Solidity: function size(string tablename, string name) constant returns(int256)
func (_TableSize *TableSizeCaller) Size(opts *bind.CallOpts, tablename string, name string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TableSize.contract.Call(opts, out, "size", tablename, name)
	return *ret0, err
}

// Size is a free data retrieval call binding the contract method 0xd2b52130.
//
// Solidity: function size(string tablename, string name) constant returns(int256)
func (_TableSize *TableSizeSession) Size(tablename string, name string) (*big.Int, error) {
	return _TableSize.Contract.Size(&_TableSize.CallOpts, tablename, name)
}

// Size is a free data retrieval call binding the contract method 0xd2b52130.
//
// Solidity: function size(string tablename, string name) constant returns(int256)
func (_TableSize *TableSizeCallerSession) Size(tablename string, name string) (*big.Int, error) {
	return _TableSize.Contract.Size(&_TableSize.CallOpts, tablename, name)
}

// TableSizeCreateResultIterator is returned from FilterCreateResult and is used to iterate over the raw logs and unpacked data for CreateResult events raised by the TableSize contract.
type TableSizeCreateResultIterator struct {
	Event *TableSizeCreateResult // Event containing the contract specifics and raw log

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
func (it *TableSizeCreateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TableSizeCreateResult)
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
		it.Event = new(TableSizeCreateResult)
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
func (it *TableSizeCreateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TableSizeCreateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TableSizeCreateResult represents a CreateResult event raised by the TableSize contract.
type TableSizeCreateResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreateResult is a free log retrieval operation binding the contract event 0xb5636cd912a73dcdb5b570dbe331dfa3e6435c93e029e642def2c8e40dacf210.
//
// Solidity: event CreateResult(int256 count)
func (_TableSize *TableSizeFilterer) FilterCreateResult(opts *bind.FilterOpts) (*TableSizeCreateResultIterator, error) {

	logs, sub, err := _TableSize.contract.FilterLogs(opts, "CreateResult")
	if err != nil {
		return nil, err
	}
	return &TableSizeCreateResultIterator{contract: _TableSize.contract, event: "CreateResult", logs: logs, sub: sub}, nil
}

// WatchCreateResult is a free log subscription operation binding the contract event 0xb5636cd912a73dcdb5b570dbe331dfa3e6435c93e029e642def2c8e40dacf210.
//
// Solidity: event CreateResult(int256 count)
func (_TableSize *TableSizeFilterer) WatchCreateResult(opts *bind.WatchOpts, sink chan<- *TableSizeCreateResult) (event.Subscription, error) {

	logs, sub, err := _TableSize.contract.WatchLogs(opts, "CreateResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TableSizeCreateResult)
				if err := _TableSize.contract.UnpackLog(event, "CreateResult", log); err != nil {
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

// ParseCreateResult is a log parse operation binding the contract event 0xb5636cd912a73dcdb5b570dbe331dfa3e6435c93e029e642def2c8e40dacf210.
//
// Solidity: event CreateResult(int256 count)
func (_TableSize *TableSizeFilterer) ParseCreateResult(log types.Log) (*TableSizeCreateResult, error) {
	event := new(TableSizeCreateResult)
	if err := _TableSize.contract.UnpackLog(event, "CreateResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TableSizeInsertResultIterator is returned from FilterInsertResult and is used to iterate over the raw logs and unpacked data for InsertResult events raised by the TableSize contract.
type TableSizeInsertResultIterator struct {
	Event *TableSizeInsertResult // Event containing the contract specifics and raw log

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
func (it *TableSizeInsertResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TableSizeInsertResult)
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
		it.Event = new(TableSizeInsertResult)
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
func (it *TableSizeInsertResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TableSizeInsertResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TableSizeInsertResult represents a InsertResult event raised by the TableSize contract.
type TableSizeInsertResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInsertResult is a free log retrieval operation binding the contract event 0xc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce.
//
// Solidity: event InsertResult(int256 count)
func (_TableSize *TableSizeFilterer) FilterInsertResult(opts *bind.FilterOpts) (*TableSizeInsertResultIterator, error) {

	logs, sub, err := _TableSize.contract.FilterLogs(opts, "InsertResult")
	if err != nil {
		return nil, err
	}
	return &TableSizeInsertResultIterator{contract: _TableSize.contract, event: "InsertResult", logs: logs, sub: sub}, nil
}

// WatchInsertResult is a free log subscription operation binding the contract event 0xc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce.
//
// Solidity: event InsertResult(int256 count)
func (_TableSize *TableSizeFilterer) WatchInsertResult(opts *bind.WatchOpts, sink chan<- *TableSizeInsertResult) (event.Subscription, error) {

	logs, sub, err := _TableSize.contract.WatchLogs(opts, "InsertResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TableSizeInsertResult)
				if err := _TableSize.contract.UnpackLog(event, "InsertResult", log); err != nil {
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

// ParseInsertResult is a log parse operation binding the contract event 0xc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce.
//
// Solidity: event InsertResult(int256 count)
func (_TableSize *TableSizeFilterer) ParseInsertResult(log types.Log) (*TableSizeInsertResult, error) {
	event := new(TableSizeInsertResult)
	if err := _TableSize.contract.UnpackLog(event, "InsertResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TableSizeRemoveResultIterator is returned from FilterRemoveResult and is used to iterate over the raw logs and unpacked data for RemoveResult events raised by the TableSize contract.
type TableSizeRemoveResultIterator struct {
	Event *TableSizeRemoveResult // Event containing the contract specifics and raw log

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
func (it *TableSizeRemoveResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TableSizeRemoveResult)
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
		it.Event = new(TableSizeRemoveResult)
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
func (it *TableSizeRemoveResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TableSizeRemoveResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TableSizeRemoveResult represents a RemoveResult event raised by the TableSize contract.
type TableSizeRemoveResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemoveResult is a free log retrieval operation binding the contract event 0x4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112.
//
// Solidity: event RemoveResult(int256 count)
func (_TableSize *TableSizeFilterer) FilterRemoveResult(opts *bind.FilterOpts) (*TableSizeRemoveResultIterator, error) {

	logs, sub, err := _TableSize.contract.FilterLogs(opts, "RemoveResult")
	if err != nil {
		return nil, err
	}
	return &TableSizeRemoveResultIterator{contract: _TableSize.contract, event: "RemoveResult", logs: logs, sub: sub}, nil
}

// WatchRemoveResult is a free log subscription operation binding the contract event 0x4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112.
//
// Solidity: event RemoveResult(int256 count)
func (_TableSize *TableSizeFilterer) WatchRemoveResult(opts *bind.WatchOpts, sink chan<- *TableSizeRemoveResult) (event.Subscription, error) {

	logs, sub, err := _TableSize.contract.WatchLogs(opts, "RemoveResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TableSizeRemoveResult)
				if err := _TableSize.contract.UnpackLog(event, "RemoveResult", log); err != nil {
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

// ParseRemoveResult is a log parse operation binding the contract event 0x4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112.
//
// Solidity: event RemoveResult(int256 count)
func (_TableSize *TableSizeFilterer) ParseRemoveResult(log types.Log) (*TableSizeRemoveResult, error) {
	event := new(TableSizeRemoveResult)
	if err := _TableSize.contract.UnpackLog(event, "RemoveResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TableSizeUpdateResultIterator is returned from FilterUpdateResult and is used to iterate over the raw logs and unpacked data for UpdateResult events raised by the TableSize contract.
type TableSizeUpdateResultIterator struct {
	Event *TableSizeUpdateResult // Event containing the contract specifics and raw log

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
func (it *TableSizeUpdateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TableSizeUpdateResult)
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
		it.Event = new(TableSizeUpdateResult)
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
func (it *TableSizeUpdateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TableSizeUpdateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TableSizeUpdateResult represents a UpdateResult event raised by the TableSize contract.
type TableSizeUpdateResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateResult is a free log retrieval operation binding the contract event 0x8e5890af40fc24a059396aca2f83d6ce41fcef086876548fa4fb8ec27e9d292a.
//
// Solidity: event UpdateResult(int256 count)
func (_TableSize *TableSizeFilterer) FilterUpdateResult(opts *bind.FilterOpts) (*TableSizeUpdateResultIterator, error) {

	logs, sub, err := _TableSize.contract.FilterLogs(opts, "UpdateResult")
	if err != nil {
		return nil, err
	}
	return &TableSizeUpdateResultIterator{contract: _TableSize.contract, event: "UpdateResult", logs: logs, sub: sub}, nil
}

// WatchUpdateResult is a free log subscription operation binding the contract event 0x8e5890af40fc24a059396aca2f83d6ce41fcef086876548fa4fb8ec27e9d292a.
//
// Solidity: event UpdateResult(int256 count)
func (_TableSize *TableSizeFilterer) WatchUpdateResult(opts *bind.WatchOpts, sink chan<- *TableSizeUpdateResult) (event.Subscription, error) {

	logs, sub, err := _TableSize.contract.WatchLogs(opts, "UpdateResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TableSizeUpdateResult)
				if err := _TableSize.contract.UnpackLog(event, "UpdateResult", log); err != nil {
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

// ParseUpdateResult is a log parse operation binding the contract event 0x8e5890af40fc24a059396aca2f83d6ce41fcef086876548fa4fb8ec27e9d292a.
//
// Solidity: event UpdateResult(int256 count)
func (_TableSize *TableSizeFilterer) ParseUpdateResult(log types.Log) (*TableSizeUpdateResult, error) {
	event := new(TableSizeUpdateResult)
	if err := _TableSize.contract.UnpackLog(event, "UpdateResult", log); err != nil {
		return nil, err
	}
	return event, nil
}
