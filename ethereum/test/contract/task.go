// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// StdTasksABI is the input ABI used to generate the binding from.
const StdTasksABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"accept\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"serial\",\"type\":\"string\"}],\"name\":\"addTask\",\"outputs\":[{\"name\":\"taskId\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"confirm\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"serial\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"AddTask\",\"type\":\"event\"}]"

// StdTasksBin is the compiled bytecode used for deploying new contracts.
const StdTasksBin = `0x608060405234801561001057600080fd5b50610877806100206000396000f30060806040526004361061006c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166319b05f4981146100715780631d65e77e1461009d57806340e58ee5146101e1578063517c0734146101f9578063ba0179b5146102d3575b600080fd5b34801561007d57600080fd5b506100896004356102eb565b604080519115158252519081900360200190f35b3480156100a957600080fd5b506100b560043561036c565b604051808760048111156100c557fe5b60ff168152602001806020018060200186600160a060020a0316600160a060020a0316815260200185600160a060020a0316600160a060020a03168152602001848152602001838103835288818151815260200191508051906020019080838360005b83811015610140578181015183820152602001610128565b50505050905090810190601f16801561016d5780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b838110156101a0578181015183820152602001610188565b50505050905090810190601f1680156101cd5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156101ed57600080fd5b506100896004356104e8565b6040805160206004803580820135601f81018490048402850184019095528484526102c194369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506105869650505050505050565b60408051918252519081900360200190f35b3480156102df57600080fd5b5061008960043561071d565b6000818152600160208190526040822090815460ff16600481111561030c57fe5b14158061032557506003810154600160a060020a031633145b156103335760009150610366565b805460ff191660021781556004810180543373ffffffffffffffffffffffffffffffffffffffff19909116179055600191505b50919050565b6000818152600160208181526040808420805460038201546004830154600584015484880180548751601f60026000199c841615610100029c909c019092168b90049182018a90048a0281018a019098528088526060998a998c998a998a99909860ff90911697959693890195600160a060020a03918216959116939287919083018282801561043d5780601f106104125761010080835404028352916020019161043d565b820191906000526020600020905b81548152906001019060200180831161042057829003601f168201915b5050875460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959a50899450925084019050828280156104cb5780601f106104a0576101008083540402835291602001916104cb565b820191906000526020600020905b8154815290600101906020018083116104ae57829003601f168201915b505050505093509650965096509650965096505091939550919395565b6000818152600160208190526040822090815460ff16600481111561050957fe5b14158061052357506003810154600160a060020a03163314155b156105315760009150610366565b805460ff1916600417815560038101546005820154604051600160a060020a03909216916108fc82150291906000818181858888f1935050505015801561057c573d6000803e3d6000fd5b5060019392505050565b6000805460019081018083556040805160c0810182528381526020808201899052818301889052336060830152608082018690523460a083015283865284905293208351815492949391929091839160ff1916908360048111156105e657fe5b0217905550602082810151805161060392600185019201906107b0565b506040820151805161061f9160028401916020909101906107b0565b50606082810151600383018054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff1991821617909155608085015160048501805491909316911617905560a09092015160059091015560408051602080820185905282825285519282019290925284517fed2bff63ecccc37cad1d9b69cf6dda32fc00ea81193ca5e3780726d4beff83fa938693869392839283019186019080838360005b838110156106db5781810151838201526020016106c3565b50505050905090810190601f1680156107085780820380516001836020036101000a031916815260200191505b50935050505060405180910390a19392505050565b60008181526001602052604081206002815460ff16600481111561073d57fe5b14158061075757506003810154600160a060020a03163314155b156107655760009150610366565b805460ff1916600317815560048101546005820154604051600160a060020a03909216916108fc82150291906000818181858888f1935050505015801561057c573d6000803e3d6000fd5b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107f157805160ff191683800117855561081e565b8280016001018555821561081e579182015b8281111561081e578251825591602001919060010190610803565b5061082a92915061082e565b5090565b61084891905b8082111561082a5760008155600101610834565b905600a165627a7a72305820b5f73c73f799ea9995d2890a41378195bb714fcb394ec1a81e2490f29bd739920029`

// DeployStdTasks deploys a new Ethereum contract, binding an instance of StdTasks to it.
func DeployStdTasks(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StdTasks, error) {
	parsed, err := abi.JSON(strings.NewReader(StdTasksABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StdTasksBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StdTasks{StdTasksCaller: StdTasksCaller{contract: contract}, StdTasksTransactor: StdTasksTransactor{contract: contract}, StdTasksFilterer: StdTasksFilterer{contract: contract}}, nil
}

// StdTasks is an auto generated Go binding around an Ethereum contract.
type StdTasks struct {
	StdTasksCaller     // Read-only binding to the contract
	StdTasksTransactor // Write-only binding to the contract
	StdTasksFilterer   // Log filterer for contract events
}

// StdTasksCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdTasksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdTasksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdTasksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdTasksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdTasksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdTasksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdTasksSession struct {
	Contract     *StdTasks         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdTasksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdTasksCallerSession struct {
	Contract *StdTasksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StdTasksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdTasksTransactorSession struct {
	Contract     *StdTasksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StdTasksRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdTasksRaw struct {
	Contract *StdTasks // Generic contract binding to access the raw methods on
}

// StdTasksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdTasksCallerRaw struct {
	Contract *StdTasksCaller // Generic read-only contract binding to access the raw methods on
}

// StdTasksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdTasksTransactorRaw struct {
	Contract *StdTasksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdTasks creates a new instance of StdTasks, bound to a specific deployed contract.
func NewStdTasks(address common.Address, backend bind.ContractBackend) (*StdTasks, error) {
	contract, err := bindStdTasks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdTasks{StdTasksCaller: StdTasksCaller{contract: contract}, StdTasksTransactor: StdTasksTransactor{contract: contract}, StdTasksFilterer: StdTasksFilterer{contract: contract}}, nil
}

// NewStdTasksCaller creates a new read-only instance of StdTasks, bound to a specific deployed contract.
func NewStdTasksCaller(address common.Address, caller bind.ContractCaller) (*StdTasksCaller, error) {
	contract, err := bindStdTasks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdTasksCaller{contract: contract}, nil
}

// NewStdTasksTransactor creates a new write-only instance of StdTasks, bound to a specific deployed contract.
func NewStdTasksTransactor(address common.Address, transactor bind.ContractTransactor) (*StdTasksTransactor, error) {
	contract, err := bindStdTasks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdTasksTransactor{contract: contract}, nil
}

// NewStdTasksFilterer creates a new log filterer instance of StdTasks, bound to a specific deployed contract.
func NewStdTasksFilterer(address common.Address, filterer bind.ContractFilterer) (*StdTasksFilterer, error) {
	contract, err := bindStdTasks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdTasksFilterer{contract: contract}, nil
}

// bindStdTasks binds a generic wrapper to an already deployed contract.
func bindStdTasks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StdTasksABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdTasks *StdTasksRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StdTasks.Contract.StdTasksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdTasks *StdTasksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdTasks.Contract.StdTasksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdTasks *StdTasksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdTasks.Contract.StdTasksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdTasks *StdTasksCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StdTasks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdTasks *StdTasksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdTasks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdTasks *StdTasksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdTasks.Contract.contract.Transact(opts, method, params...)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(taskId uint256) constant returns(uint8, string, string, address, address, uint256)
func (_StdTasks *StdTasksCaller) GetTask(opts *bind.CallOpts, taskId *big.Int) (uint8, string, string, common.Address, common.Address, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(common.Address)
		ret4 = new(common.Address)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _StdTasks.contract.Call(opts, out, "getTask", taskId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(taskId uint256) constant returns(uint8, string, string, address, address, uint256)
func (_StdTasks *StdTasksSession) GetTask(taskId *big.Int) (uint8, string, string, common.Address, common.Address, *big.Int, error) {
	return _StdTasks.Contract.GetTask(&_StdTasks.CallOpts, taskId)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(taskId uint256) constant returns(uint8, string, string, address, address, uint256)
func (_StdTasks *StdTasksCallerSession) GetTask(taskId *big.Int) (uint8, string, string, common.Address, common.Address, *big.Int, error) {
	return _StdTasks.Contract.GetTask(&_StdTasks.CallOpts, taskId)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(taskId uint256) returns(bool)
func (_StdTasks *StdTasksTransactor) Accept(opts *bind.TransactOpts, taskId *big.Int) (*types.Transaction, error) {
	return _StdTasks.contract.Transact(opts, "accept", taskId)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(taskId uint256) returns(bool)
func (_StdTasks *StdTasksSession) Accept(taskId *big.Int) (*types.Transaction, error) {
	return _StdTasks.Contract.Accept(&_StdTasks.TransactOpts, taskId)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(taskId uint256) returns(bool)
func (_StdTasks *StdTasksTransactorSession) Accept(taskId *big.Int) (*types.Transaction, error) {
	return _StdTasks.Contract.Accept(&_StdTasks.TransactOpts, taskId)
}

// AddTask is a paid mutator transaction binding the contract method 0x517c0734.
//
// Solidity: function addTask(name string, description string, serial string) returns(taskId uint256)
func (_StdTasks *StdTasksTransactor) AddTask(opts *bind.TransactOpts, name string, description string, serial string) (*types.Transaction, error) {
	return _StdTasks.contract.Transact(opts, "addTask", name, description, serial)
}

// AddTask is a paid mutator transaction binding the contract method 0x517c0734.
//
// Solidity: function addTask(name string, description string, serial string) returns(taskId uint256)
func (_StdTasks *StdTasksSession) AddTask(name string, description string, serial string) (*types.Transaction, error) {
	return _StdTasks.Contract.AddTask(&_StdTasks.TransactOpts, name, description, serial)
}

// AddTask is a paid mutator transaction binding the contract method 0x517c0734.
//
// Solidity: function addTask(name string, description string, serial string) returns(taskId uint256)
func (_StdTasks *StdTasksTransactorSession) AddTask(name string, description string, serial string) (*types.Transaction, error) {
	return _StdTasks.Contract.AddTask(&_StdTasks.TransactOpts, name, description, serial)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(taskId uint256) returns(bool)
func (_StdTasks *StdTasksTransactor) Cancel(opts *bind.TransactOpts, taskId *big.Int) (*types.Transaction, error) {
	return _StdTasks.contract.Transact(opts, "cancel", taskId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(taskId uint256) returns(bool)
func (_StdTasks *StdTasksSession) Cancel(taskId *big.Int) (*types.Transaction, error) {
	return _StdTasks.Contract.Cancel(&_StdTasks.TransactOpts, taskId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(taskId uint256) returns(bool)
func (_StdTasks *StdTasksTransactorSession) Cancel(taskId *big.Int) (*types.Transaction, error) {
	return _StdTasks.Contract.Cancel(&_StdTasks.TransactOpts, taskId)
}

// Confirm is a paid mutator transaction binding the contract method 0xba0179b5.
//
// Solidity: function confirm(taskId uint256) returns(bool)
func (_StdTasks *StdTasksTransactor) Confirm(opts *bind.TransactOpts, taskId *big.Int) (*types.Transaction, error) {
	return _StdTasks.contract.Transact(opts, "confirm", taskId)
}

// Confirm is a paid mutator transaction binding the contract method 0xba0179b5.
//
// Solidity: function confirm(taskId uint256) returns(bool)
func (_StdTasks *StdTasksSession) Confirm(taskId *big.Int) (*types.Transaction, error) {
	return _StdTasks.Contract.Confirm(&_StdTasks.TransactOpts, taskId)
}

// Confirm is a paid mutator transaction binding the contract method 0xba0179b5.
//
// Solidity: function confirm(taskId uint256) returns(bool)
func (_StdTasks *StdTasksTransactorSession) Confirm(taskId *big.Int) (*types.Transaction, error) {
	return _StdTasks.Contract.Confirm(&_StdTasks.TransactOpts, taskId)
}

// StdTasksAddTaskIterator is returned from FilterAddTask and is used to iterate over the raw logs and unpacked data for AddTask events raised by the StdTasks contract.
type StdTasksAddTaskIterator struct {
	Event *StdTasksAddTask // Event containing the contract specifics and raw log

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
func (it *StdTasksAddTaskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StdTasksAddTask)
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
		it.Event = new(StdTasksAddTask)
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
func (it *StdTasksAddTaskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StdTasksAddTaskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StdTasksAddTask represents a AddTask event raised by the StdTasks contract.
type StdTasksAddTask struct {
	Serial string
	TaskId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddTask is a free log retrieval operation binding the contract event 0xed2bff63ecccc37cad1d9b69cf6dda32fc00ea81193ca5e3780726d4beff83fa.
//
// Solidity: e AddTask(serial string, taskId uint256)
func (_StdTasks *StdTasksFilterer) FilterAddTask(opts *bind.FilterOpts) (*StdTasksAddTaskIterator, error) {

	logs, sub, err := _StdTasks.contract.FilterLogs(opts, "AddTask")
	if err != nil {
		return nil, err
	}
	return &StdTasksAddTaskIterator{contract: _StdTasks.contract, event: "AddTask", logs: logs, sub: sub}, nil
}

// WatchAddTask is a free log subscription operation binding the contract event 0xed2bff63ecccc37cad1d9b69cf6dda32fc00ea81193ca5e3780726d4beff83fa.
//
// Solidity: e AddTask(serial string, taskId uint256)
func (_StdTasks *StdTasksFilterer) WatchAddTask(opts *bind.WatchOpts, sink chan<- *StdTasksAddTask) (event.Subscription, error) {

	logs, sub, err := _StdTasks.contract.WatchLogs(opts, "AddTask")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StdTasksAddTask)
				if err := _StdTasks.contract.UnpackLog(event, "AddTask", log); err != nil {
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
