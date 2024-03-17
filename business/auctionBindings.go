// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package business

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

// SimpleAuctionMetaData contains all meta data concerning the SimpleAuction contract.
var SimpleAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_biddingTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auctionEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610b1b380380610b1b8339818101604052810190610031919061011a565b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550814261007c9190610185565b60018190555050506101b8565b5f80fd5b5f819050919050565b61009f8161008d565b81146100a9575f80fd5b50565b5f815190506100ba81610096565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100e9826100c0565b9050919050565b6100f9816100df565b8114610103575f80fd5b50565b5f81519050610114816100f0565b92915050565b5f80604083850312156101305761012f610089565b5b5f61013d858286016100ac565b925050602061014e85828601610106565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61018f8261008d565b915061019a8361008d565b92508282019050808211156101b2576101b1610158565b5b92915050565b610956806101c55f395ff3fe60806040526004361061006f575f3560e01c80633ccfd60b1161004d5780633ccfd60b146100bd5780634b449cba146100e757806391f9015714610111578063d57bde791461013b5761006f565b80631998aeef146100735780632a24f46c1461007d57806338af3eed14610093575b5f80fd5b61007b610165565b005b348015610088575f80fd5b506100916102f0565b005b34801561009e575f80fd5b506100a7610461565b6040516100b4919061060d565b60405180910390f35b3480156100c8575f80fd5b506100d1610484565b6040516100de9190610640565b60405180910390f35b3480156100f2575f80fd5b506100fb61059d565b6040516101089190610671565b60405180910390f35b34801561011c575f80fd5b506101256105a3565b60405161013291906106aa565b60405180910390f35b348015610146575f80fd5b5061014f6105c8565b60405161015c9190610671565b60405180910390f35b6001544211156101aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101a19061071d565b60405180910390fd5b60035434116101ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e590610785565b60405180910390fd5b5f6003541461026e5760035460045f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461026691906107d0565b925050819055505b3360025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550346003819055507ff4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad30033346040516102e6929190610803565b60405180910390a1565b600154421015610335576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032c90610874565b60405180910390fd5b60055f9054906101000a900460ff1615610384576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037b90610902565b60405180910390fd5b600160055f6101000a81548160ff0219169083151502179055507fdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166003546040516103f2929190610803565b60405180910390a15f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc60035490811502906040515f60405180830381858888f1935050505015801561045e573d5f803e3d5ffd5b50565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f811115610594575f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050610593578060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f91505061059a565b5b60019150505b90565b60015481565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6105f7826105ce565b9050919050565b610607816105ed565b82525050565b5f6020820190506106205f8301846105fe565b92915050565b5f8115159050919050565b61063a81610626565b82525050565b5f6020820190506106535f830184610631565b92915050565b5f819050919050565b61066b81610659565b82525050565b5f6020820190506106845f830184610662565b92915050565b5f610694826105ce565b9050919050565b6106a48161068a565b82525050565b5f6020820190506106bd5f83018461069b565b92915050565b5f82825260208201905092915050565b7f41756374696f6e20616c726561647920656e64656400000000000000000000005f82015250565b5f6107076015836106c3565b9150610712826106d3565b602082019050919050565b5f6020820190508181035f830152610734816106fb565b9050919050565b7f546865726520616c7265616479206973206120686967686572206269640000005f82015250565b5f61076f601d836106c3565b915061077a8261073b565b602082019050919050565b5f6020820190508181035f83015261079c81610763565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6107da82610659565b91506107e583610659565b92508282019050808211156107fd576107fc6107a3565b5b92915050565b5f6040820190506108165f83018561069b565b6108236020830184610662565b9392505050565b7f41756374696f6e206e6f742079657420656e64656400000000000000000000005f82015250565b5f61085e6015836106c3565b91506108698261082a565b602082019050919050565b5f6020820190508181035f83015261088b81610852565b9050919050565b7f61756374696f6e456e642068617320616c7265616479206265656e2063616c6c5f8201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b5f6108ec6022836106c3565b91506108f782610892565b604082019050919050565b5f6020820190508181035f830152610919816108e0565b905091905056fea26469706673582212208127f5e0921fcc06af819fcc274f0f4a15d606cd9d6aee134a126d0971eb86c564736f6c63430008190033",
}

// SimpleAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleAuctionMetaData.ABI instead.
var SimpleAuctionABI = SimpleAuctionMetaData.ABI

// SimpleAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleAuctionMetaData.Bin instead.
var SimpleAuctionBin = SimpleAuctionMetaData.Bin

// DeploySimpleAuction deploys a new Ethereum contract, binding an instance of SimpleAuction to it.
func DeploySimpleAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _biddingTime *big.Int, _beneficiary common.Address) (common.Address, *types.Transaction, *SimpleAuction, error) {
	parsed, err := SimpleAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleAuctionBin), backend, _biddingTime, _beneficiary)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleAuction{SimpleAuctionCaller: SimpleAuctionCaller{contract: contract}, SimpleAuctionTransactor: SimpleAuctionTransactor{contract: contract}, SimpleAuctionFilterer: SimpleAuctionFilterer{contract: contract}}, nil
}

// SimpleAuction is an auto generated Go binding around an Ethereum contract.
type SimpleAuction struct {
	SimpleAuctionCaller     // Read-only binding to the contract
	SimpleAuctionTransactor // Write-only binding to the contract
	SimpleAuctionFilterer   // Log filterer for contract events
}

// SimpleAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleAuctionSession struct {
	Contract     *SimpleAuction    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleAuctionCallerSession struct {
	Contract *SimpleAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleAuctionTransactorSession struct {
	Contract     *SimpleAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleAuctionRaw struct {
	Contract *SimpleAuction // Generic contract binding to access the raw methods on
}

// SimpleAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleAuctionCallerRaw struct {
	Contract *SimpleAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleAuctionTransactorRaw struct {
	Contract *SimpleAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleAuction creates a new instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuction(address common.Address, backend bind.ContractBackend) (*SimpleAuction, error) {
	contract, err := bindSimpleAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleAuction{SimpleAuctionCaller: SimpleAuctionCaller{contract: contract}, SimpleAuctionTransactor: SimpleAuctionTransactor{contract: contract}, SimpleAuctionFilterer: SimpleAuctionFilterer{contract: contract}}, nil
}

// NewSimpleAuctionCaller creates a new read-only instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuctionCaller(address common.Address, caller bind.ContractCaller) (*SimpleAuctionCaller, error) {
	contract, err := bindSimpleAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionCaller{contract: contract}, nil
}

// NewSimpleAuctionTransactor creates a new write-only instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleAuctionTransactor, error) {
	contract, err := bindSimpleAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionTransactor{contract: contract}, nil
}

// NewSimpleAuctionFilterer creates a new log filterer instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleAuctionFilterer, error) {
	contract, err := bindSimpleAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionFilterer{contract: contract}, nil
}

// bindSimpleAuction binds a generic wrapper to an already deployed contract.
func bindSimpleAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAuction *SimpleAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAuction.Contract.SimpleAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAuction *SimpleAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.Contract.SimpleAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAuction *SimpleAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAuction.Contract.SimpleAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAuction *SimpleAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAuction *SimpleAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAuction *SimpleAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAuction.Contract.contract.Transact(opts, method, params...)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_SimpleAuction *SimpleAuctionCaller) AuctionEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleAuction.contract.Call(opts, &out, "auctionEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_SimpleAuction *SimpleAuctionSession) AuctionEndTime() (*big.Int, error) {
	return _SimpleAuction.Contract.AuctionEndTime(&_SimpleAuction.CallOpts)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_SimpleAuction *SimpleAuctionCallerSession) AuctionEndTime() (*big.Int, error) {
	return _SimpleAuction.Contract.AuctionEndTime(&_SimpleAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_SimpleAuction *SimpleAuctionCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleAuction.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_SimpleAuction *SimpleAuctionSession) Beneficiary() (common.Address, error) {
	return _SimpleAuction.Contract.Beneficiary(&_SimpleAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_SimpleAuction *SimpleAuctionCallerSession) Beneficiary() (common.Address, error) {
	return _SimpleAuction.Contract.Beneficiary(&_SimpleAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_SimpleAuction *SimpleAuctionCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleAuction.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_SimpleAuction *SimpleAuctionSession) HighestBid() (*big.Int, error) {
	return _SimpleAuction.Contract.HighestBid(&_SimpleAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_SimpleAuction *SimpleAuctionCallerSession) HighestBid() (*big.Int, error) {
	return _SimpleAuction.Contract.HighestBid(&_SimpleAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_SimpleAuction *SimpleAuctionCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleAuction.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_SimpleAuction *SimpleAuctionSession) HighestBidder() (common.Address, error) {
	return _SimpleAuction.Contract.HighestBidder(&_SimpleAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_SimpleAuction *SimpleAuctionCallerSession) HighestBidder() (common.Address, error) {
	return _SimpleAuction.Contract.HighestBidder(&_SimpleAuction.CallOpts)
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_SimpleAuction *SimpleAuctionTransactor) AuctionEnd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.contract.Transact(opts, "auctionEnd")
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_SimpleAuction *SimpleAuctionSession) AuctionEnd() (*types.Transaction, error) {
	return _SimpleAuction.Contract.AuctionEnd(&_SimpleAuction.TransactOpts)
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_SimpleAuction *SimpleAuctionTransactorSession) AuctionEnd() (*types.Transaction, error) {
	return _SimpleAuction.Contract.AuctionEnd(&_SimpleAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_SimpleAuction *SimpleAuctionTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_SimpleAuction *SimpleAuctionSession) Bid() (*types.Transaction, error) {
	return _SimpleAuction.Contract.Bid(&_SimpleAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_SimpleAuction *SimpleAuctionTransactorSession) Bid() (*types.Transaction, error) {
	return _SimpleAuction.Contract.Bid(&_SimpleAuction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_SimpleAuction *SimpleAuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_SimpleAuction *SimpleAuctionSession) Withdraw() (*types.Transaction, error) {
	return _SimpleAuction.Contract.Withdraw(&_SimpleAuction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_SimpleAuction *SimpleAuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _SimpleAuction.Contract.Withdraw(&_SimpleAuction.TransactOpts)
}

// SimpleAuctionAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the SimpleAuction contract.
type SimpleAuctionAuctionEndedIterator struct {
	Event *SimpleAuctionAuctionEnded // Event containing the contract specifics and raw log

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
func (it *SimpleAuctionAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleAuctionAuctionEnded)
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
		it.Event = new(SimpleAuctionAuctionEnded)
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
func (it *SimpleAuctionAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleAuctionAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleAuctionAuctionEnded represents a AuctionEnded event raised by the SimpleAuction contract.
type SimpleAuctionAuctionEnded struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*SimpleAuctionAuctionEndedIterator, error) {

	logs, sub, err := _SimpleAuction.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionAuctionEndedIterator{contract: _SimpleAuction.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *SimpleAuctionAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _SimpleAuction.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleAuctionAuctionEnded)
				if err := _SimpleAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) ParseAuctionEnded(log types.Log) (*SimpleAuctionAuctionEnded, error) {
	event := new(SimpleAuctionAuctionEnded)
	if err := _SimpleAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the SimpleAuction contract.
type SimpleAuctionHighestBidIncreasedIterator struct {
	Event *SimpleAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *SimpleAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleAuctionHighestBidIncreased)
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
		it.Event = new(SimpleAuctionHighestBidIncreased)
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
func (it *SimpleAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the SimpleAuction contract.
type SimpleAuctionHighestBidIncreased struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*SimpleAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _SimpleAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionHighestBidIncreasedIterator{contract: _SimpleAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *SimpleAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _SimpleAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleAuctionHighestBidIncreased)
				if err := _SimpleAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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

// ParseHighestBidIncreased is a log parse operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*SimpleAuctionHighestBidIncreased, error) {
	event := new(SimpleAuctionHighestBidIncreased)
	if err := _SimpleAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
