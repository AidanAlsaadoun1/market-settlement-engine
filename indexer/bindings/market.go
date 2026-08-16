// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = time.Tick
	_ = context.Background
)

// MarketMetaData contains all meta data concerning the Market contract.
var MarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buy\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"yes\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimed\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createMarket\",\"inputs\":[{\"name\":\"question\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"marketCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"question\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outcome\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"yesPool\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"noPool\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noStakes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outcome\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"yesStakes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payout\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketCreated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"question\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketResolved\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"outcome\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SharesPurchased\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"yes\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// MarketABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketMetaData.ABI instead.
var MarketABI = MarketMetaData.ABI

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_Market *MarketCaller) Claimed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "claimed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_Market *MarketSession) Claimed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Market.Contract.Claimed(&_Market.CallOpts, arg0, arg1)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_Market *MarketCallerSession) Claimed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Market.Contract.Claimed(&_Market.CallOpts, arg0, arg1)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_Market *MarketCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "marketCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_Market *MarketSession) MarketCount() (*big.Int, error) {
	return _Market.Contract.MarketCount(&_Market.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_Market *MarketCallerSession) MarketCount() (*big.Int, error) {
	return _Market.Contract.MarketCount(&_Market.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0xb1283e77.
//
// Solidity: function markets(uint256 ) view returns(string question, uint256 closeTime, bool resolved, bool outcome, uint256 yesPool, uint256 noPool)
func (_Market *MarketCaller) Markets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Question  string
	CloseTime *big.Int
	Resolved  bool
	Outcome   bool
	YesPool   *big.Int
	NoPool    *big.Int
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		Question  string
		CloseTime *big.Int
		Resolved  bool
		Outcome   bool
		YesPool   *big.Int
		NoPool    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Question = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CloseTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Resolved = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Outcome = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.YesPool = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.NoPool = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0xb1283e77.
//
// Solidity: function markets(uint256 ) view returns(string question, uint256 closeTime, bool resolved, bool outcome, uint256 yesPool, uint256 noPool)
func (_Market *MarketSession) Markets(arg0 *big.Int) (struct {
	Question  string
	CloseTime *big.Int
	Resolved  bool
	Outcome   bool
	YesPool   *big.Int
	NoPool    *big.Int
}, error) {
	return _Market.Contract.Markets(&_Market.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0xb1283e77.
//
// Solidity: function markets(uint256 ) view returns(string question, uint256 closeTime, bool resolved, bool outcome, uint256 yesPool, uint256 noPool)
func (_Market *MarketCallerSession) Markets(arg0 *big.Int) (struct {
	Question  string
	CloseTime *big.Int
	Resolved  bool
	Outcome   bool
	YesPool   *big.Int
	NoPool    *big.Int
}, error) {
	return _Market.Contract.Markets(&_Market.CallOpts, arg0)
}

// NoStakes is a free data retrieval call binding the contract method 0xe06d8ad1.
//
// Solidity: function noStakes(uint256 , address ) view returns(uint256)
func (_Market *MarketCaller) NoStakes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "noStakes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoStakes is a free data retrieval call binding the contract method 0xe06d8ad1.
//
// Solidity: function noStakes(uint256 , address ) view returns(uint256)
func (_Market *MarketSession) NoStakes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Market.Contract.NoStakes(&_Market.CallOpts, arg0, arg1)
}

// NoStakes is a free data retrieval call binding the contract method 0xe06d8ad1.
//
// Solidity: function noStakes(uint256 , address ) view returns(uint256)
func (_Market *MarketCallerSession) NoStakes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Market.Contract.NoStakes(&_Market.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketCallerSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// YesStakes is a free data retrieval call binding the contract method 0x512f1b7a.
//
// Solidity: function yesStakes(uint256 , address ) view returns(uint256)
func (_Market *MarketCaller) YesStakes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "yesStakes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YesStakes is a free data retrieval call binding the contract method 0x512f1b7a.
//
// Solidity: function yesStakes(uint256 , address ) view returns(uint256)
func (_Market *MarketSession) YesStakes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Market.Contract.YesStakes(&_Market.CallOpts, arg0, arg1)
}

// YesStakes is a free data retrieval call binding the contract method 0x512f1b7a.
//
// Solidity: function yesStakes(uint256 , address ) view returns(uint256)
func (_Market *MarketCallerSession) YesStakes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Market.Contract.YesStakes(&_Market.CallOpts, arg0, arg1)
}

// Buy is a paid mutator transaction binding the contract method 0x31c26b11.
//
// Solidity: function buy(uint256 marketId, bool yes) payable returns()
func (_Market *MarketTransactor) Buy(opts *bind.TransactOpts, marketId *big.Int, yes bool) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "buy", marketId, yes)
}

// Buy is a paid mutator transaction binding the contract method 0x31c26b11.
//
// Solidity: function buy(uint256 marketId, bool yes) payable returns()
func (_Market *MarketSession) Buy(marketId *big.Int, yes bool) (*types.Transaction, error) {
	return _Market.Contract.Buy(&_Market.TransactOpts, marketId, yes)
}

// Buy is a paid mutator transaction binding the contract method 0x31c26b11.
//
// Solidity: function buy(uint256 marketId, bool yes) payable returns()
func (_Market *MarketTransactorSession) Buy(marketId *big.Int, yes bool) (*types.Transaction, error) {
	return _Market.Contract.Buy(&_Market.TransactOpts, marketId, yes)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 marketId) returns()
func (_Market *MarketTransactor) Claim(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "claim", marketId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 marketId) returns()
func (_Market *MarketSession) Claim(marketId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Claim(&_Market.TransactOpts, marketId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 marketId) returns()
func (_Market *MarketTransactorSession) Claim(marketId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Claim(&_Market.TransactOpts, marketId)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x883c84c1.
//
// Solidity: function createMarket(string question, uint256 closeTime) returns(uint256 id)
func (_Market *MarketTransactor) CreateMarket(opts *bind.TransactOpts, question string, closeTime *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "createMarket", question, closeTime)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x883c84c1.
//
// Solidity: function createMarket(string question, uint256 closeTime) returns(uint256 id)
func (_Market *MarketSession) CreateMarket(question string, closeTime *big.Int) (*types.Transaction, error) {
	return _Market.Contract.CreateMarket(&_Market.TransactOpts, question, closeTime)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x883c84c1.
//
// Solidity: function createMarket(string question, uint256 closeTime) returns(uint256 id)
func (_Market *MarketTransactorSession) CreateMarket(question string, closeTime *big.Int) (*types.Transaction, error) {
	return _Market.Contract.CreateMarket(&_Market.TransactOpts, question, closeTime)
}

// Resolve is a paid mutator transaction binding the contract method 0x52a34b05.
//
// Solidity: function resolve(uint256 marketId, bool outcome) returns()
func (_Market *MarketTransactor) Resolve(opts *bind.TransactOpts, marketId *big.Int, outcome bool) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "resolve", marketId, outcome)
}

// Resolve is a paid mutator transaction binding the contract method 0x52a34b05.
//
// Solidity: function resolve(uint256 marketId, bool outcome) returns()
func (_Market *MarketSession) Resolve(marketId *big.Int, outcome bool) (*types.Transaction, error) {
	return _Market.Contract.Resolve(&_Market.TransactOpts, marketId, outcome)
}

// Resolve is a paid mutator transaction binding the contract method 0x52a34b05.
//
// Solidity: function resolve(uint256 marketId, bool outcome) returns()
func (_Market *MarketTransactorSession) Resolve(marketId *big.Int, outcome bool) (*types.Transaction, error) {
	return _Market.Contract.Resolve(&_Market.TransactOpts, marketId, outcome)
}

// MarketClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Market contract.
type MarketClaimedIterator struct {
	Event *MarketClaimed // Event containing the contract specifics and raw log

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
func (it *MarketClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketClaimed)
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
		it.Event = new(MarketClaimed)
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
func (it *MarketClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketClaimed represents a Claimed event raised by the Market contract.
type MarketClaimed struct {
	MarketId *big.Int
	Claimer  common.Address
	Payout   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 indexed marketId, address indexed claimer, uint256 payout)
func (_Market *MarketFilterer) FilterClaimed(opts *bind.FilterOpts, marketId []*big.Int, claimer []common.Address) (*MarketClaimedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "Claimed", marketIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &MarketClaimedIterator{contract: _Market.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 indexed marketId, address indexed claimer, uint256 payout)
func (_Market *MarketFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *MarketClaimed, marketId []*big.Int, claimer []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "Claimed", marketIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketClaimed)
				if err := _Market.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 indexed marketId, address indexed claimer, uint256 payout)
func (_Market *MarketFilterer) ParseClaimed(log types.Log) (*MarketClaimed, error) {
	event := new(MarketClaimed)
	if err := _Market.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the Market contract.
type MarketMarketCreatedIterator struct {
	Event *MarketMarketCreated // Event containing the contract specifics and raw log

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
func (it *MarketMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketMarketCreated)
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
		it.Event = new(MarketMarketCreated)
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
func (it *MarketMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketMarketCreated represents a MarketCreated event raised by the Market contract.
type MarketMarketCreated struct {
	Id        *big.Int
	Question  string
	CloseTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0x2d1e9ad45dbe8e898e0374deebf2a661b2ae7c1855d2a29507d584c35d287c75.
//
// Solidity: event MarketCreated(uint256 indexed id, string question, uint256 closeTime)
func (_Market *MarketFilterer) FilterMarketCreated(opts *bind.FilterOpts, id []*big.Int) (*MarketMarketCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "MarketCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &MarketMarketCreatedIterator{contract: _Market.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0x2d1e9ad45dbe8e898e0374deebf2a661b2ae7c1855d2a29507d584c35d287c75.
//
// Solidity: event MarketCreated(uint256 indexed id, string question, uint256 closeTime)
func (_Market *MarketFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *MarketMarketCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "MarketCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketMarketCreated)
				if err := _Market.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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

// ParseMarketCreated is a log parse operation binding the contract event 0x2d1e9ad45dbe8e898e0374deebf2a661b2ae7c1855d2a29507d584c35d287c75.
//
// Solidity: event MarketCreated(uint256 indexed id, string question, uint256 closeTime)
func (_Market *MarketFilterer) ParseMarketCreated(log types.Log) (*MarketMarketCreated, error) {
	event := new(MarketMarketCreated)
	if err := _Market.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketMarketResolvedIterator is returned from FilterMarketResolved and is used to iterate over the raw logs and unpacked data for MarketResolved events raised by the Market contract.
type MarketMarketResolvedIterator struct {
	Event *MarketMarketResolved // Event containing the contract specifics and raw log

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
func (it *MarketMarketResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketMarketResolved)
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
		it.Event = new(MarketMarketResolved)
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
func (it *MarketMarketResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketMarketResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketMarketResolved represents a MarketResolved event raised by the Market contract.
type MarketMarketResolved struct {
	MarketId *big.Int
	Outcome  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketResolved is a free log retrieval operation binding the contract event 0x4927fe38919783250023d27e65a3e56b6b5c3e49364e51674a41ef08d62460d9.
//
// Solidity: event MarketResolved(uint256 indexed marketId, bool outcome)
func (_Market *MarketFilterer) FilterMarketResolved(opts *bind.FilterOpts, marketId []*big.Int) (*MarketMarketResolvedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "MarketResolved", marketIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketMarketResolvedIterator{contract: _Market.contract, event: "MarketResolved", logs: logs, sub: sub}, nil
}

// WatchMarketResolved is a free log subscription operation binding the contract event 0x4927fe38919783250023d27e65a3e56b6b5c3e49364e51674a41ef08d62460d9.
//
// Solidity: event MarketResolved(uint256 indexed marketId, bool outcome)
func (_Market *MarketFilterer) WatchMarketResolved(opts *bind.WatchOpts, sink chan<- *MarketMarketResolved, marketId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "MarketResolved", marketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketMarketResolved)
				if err := _Market.contract.UnpackLog(event, "MarketResolved", log); err != nil {
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

// ParseMarketResolved is a log parse operation binding the contract event 0x4927fe38919783250023d27e65a3e56b6b5c3e49364e51674a41ef08d62460d9.
//
// Solidity: event MarketResolved(uint256 indexed marketId, bool outcome)
func (_Market *MarketFilterer) ParseMarketResolved(log types.Log) (*MarketMarketResolved, error) {
	event := new(MarketMarketResolved)
	if err := _Market.contract.UnpackLog(event, "MarketResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketSharesPurchasedIterator is returned from FilterSharesPurchased and is used to iterate over the raw logs and unpacked data for SharesPurchased events raised by the Market contract.
type MarketSharesPurchasedIterator struct {
	Event *MarketSharesPurchased // Event containing the contract specifics and raw log

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
func (it *MarketSharesPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketSharesPurchased)
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
		it.Event = new(MarketSharesPurchased)
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
func (it *MarketSharesPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketSharesPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketSharesPurchased represents a SharesPurchased event raised by the Market contract.
type MarketSharesPurchased struct {
	MarketId *big.Int
	Buyer    common.Address
	Yes      bool
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesPurchased is a free log retrieval operation binding the contract event 0xb3a1ddef05fa52ff8a462662e00f5fd54a3b35e85417cafcaa671fb30434bd11.
//
// Solidity: event SharesPurchased(uint256 indexed marketId, address indexed buyer, bool yes, uint256 amount)
func (_Market *MarketFilterer) FilterSharesPurchased(opts *bind.FilterOpts, marketId []*big.Int, buyer []common.Address) (*MarketSharesPurchasedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "SharesPurchased", marketIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &MarketSharesPurchasedIterator{contract: _Market.contract, event: "SharesPurchased", logs: logs, sub: sub}, nil
}

// WatchSharesPurchased is a free log subscription operation binding the contract event 0xb3a1ddef05fa52ff8a462662e00f5fd54a3b35e85417cafcaa671fb30434bd11.
//
// Solidity: event SharesPurchased(uint256 indexed marketId, address indexed buyer, bool yes, uint256 amount)
func (_Market *MarketFilterer) WatchSharesPurchased(opts *bind.WatchOpts, sink chan<- *MarketSharesPurchased, marketId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "SharesPurchased", marketIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketSharesPurchased)
				if err := _Market.contract.UnpackLog(event, "SharesPurchased", log); err != nil {
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

// ParseSharesPurchased is a log parse operation binding the contract event 0xb3a1ddef05fa52ff8a462662e00f5fd54a3b35e85417cafcaa671fb30434bd11.
//
// Solidity: event SharesPurchased(uint256 indexed marketId, address indexed buyer, bool yes, uint256 amount)
func (_Market *MarketFilterer) ParseSharesPurchased(log types.Log) (*MarketSharesPurchased, error) {
	event := new(MarketSharesPurchased)
	if err := _Market.contract.UnpackLog(event, "SharesPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
