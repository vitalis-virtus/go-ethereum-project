// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content string
	Status  bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"toggle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506111d38061005c5f395ff3fe608060405234801561000f575f80fd5b506004361061007b575f3560e01c80638da5cb5b116100595780638da5cb5b146100d55780639507d39a146100f3578063b0c8f9dc14610123578063f745630f1461013f5761007b565b80630f560cd71461007f5780634cc822151461009d578063751ef753146100b9575b5f80fd5b61008761015b565b6040516100949190610998565b60405180910390f35b6100b760048036038101906100b291906109fc565b6102b6565b005b6100d360048036038101906100ce91906109fc565b610416565b005b6100dd6104d9565b6040516100ea9190610a66565b60405180910390f35b61010d600480360381019061010891906109fc565b6104fc565b60405161011a9190610ab9565b60405180910390f35b61013d60048036038101906101389190610c05565b610635565b005b61015960048036038101906101549190610c4c565b610702565b005b60603373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101b3575f80fd5b6001805480602002602001604051908101604052809291908181526020015f905b828210156102ad578382905f5260205f2090600202016040518060400160405290815f8201805461020490610cd3565b80601f016020809104026020016040519081016040528092919081815260200182805461023090610cd3565b801561027b5780601f106102525761010080835404028352916020019161027b565b820191905f5260205f20905b81548152906001019060200180831161025e57829003601f168201915b50505050508152602001600182015f9054906101000a900460ff161515151581525050815260200190600101906101d4565b50505050905090565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461030c575f80fd5b5f8190505b600180805490506103229190610d30565b8110156103c857600180826103379190610d63565b8154811061034857610347610d96565b5b905f5260205f2090600202016001828154811061036857610367610d96565b5b905f5260205f2090600202015f8201815f0190816103869190610f75565b50600182015f9054906101000a900460ff16816001015f6101000a81548160ff02191690831515021790555090505080806103c09061105a565b915050610311565b5060018054806103db576103da6110a1565b5b600190038181905f5260205f2090600202015f8082015f6103fc919061078c565b600182015f6101000a81549060ff02191690555050905550565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461046c575f80fd5b600181815481106104805761047f610d96565b5b905f5260205f2090600202016001015f9054906101000a900460ff1615600182815481106104b1576104b0610d96565b5b905f5260205f2090600202016001015f6101000a81548160ff02191690831515021790555050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6105046107c9565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461055a575f80fd5b6001828154811061056e5761056d610d96565b5b905f5260205f2090600202016040518060400160405290815f8201805461059490610cd3565b80601f01602080910402602001604051908101604052809291908181526020018280546105c090610cd3565b801561060b5780601f106105e25761010080835404028352916020019161060b565b820191905f5260205f20905b8154815290600101906020018083116105ee57829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1615151515815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461068b575f80fd5b600160405180604001604052808381526020015f1515815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f0190816106dd91906110ce565b506020820151816001015f6101000a81548160ff021916908315150217905550505050565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610758575f80fd5b806001838154811061076d5761076c610d96565b5b905f5260205f2090600202015f01908161078791906110ce565b505050565b50805461079890610cd3565b5f825580601f106107a957506107c6565b601f0160209004905f5260205f20908101906107c591906107e4565b5b50565b6040518060400160405280606081526020015f151581525090565b5b808211156107fb575f815f9055506001016107e5565b5090565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561085f578082015181840152602081019050610844565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61088482610828565b61088e8185610832565b935061089e818560208601610842565b6108a78161086a565b840191505092915050565b5f8115159050919050565b6108c6816108b2565b82525050565b5f604083015f8301518482035f8601526108e6828261087a565b91505060208301516108fb60208601826108bd565b508091505092915050565b5f61091183836108cc565b905092915050565b5f602082019050919050565b5f61092f826107ff565b6109398185610809565b93508360208202850161094b85610819565b805f5b8581101561098657848403895281516109678582610906565b945061097283610919565b925060208a0199505060018101905061094e565b50829750879550505050505092915050565b5f6020820190508181035f8301526109b08184610925565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6109db816109c9565b81146109e5575f80fd5b50565b5f813590506109f6816109d2565b92915050565b5f60208284031215610a1157610a106109c1565b5b5f610a1e848285016109e8565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a5082610a27565b9050919050565b610a6081610a46565b82525050565b5f602082019050610a795f830184610a57565b92915050565b5f604083015f8301518482035f860152610a99828261087a565b9150506020830151610aae60208601826108bd565b508091505092915050565b5f6020820190508181035f830152610ad18184610a7f565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610b178261086a565b810181811067ffffffffffffffff82111715610b3657610b35610ae1565b5b80604052505050565b5f610b486109b8565b9050610b548282610b0e565b919050565b5f67ffffffffffffffff821115610b7357610b72610ae1565b5b610b7c8261086a565b9050602081019050919050565b828183375f83830152505050565b5f610ba9610ba484610b59565b610b3f565b905082815260208101848484011115610bc557610bc4610add565b5b610bd0848285610b89565b509392505050565b5f82601f830112610bec57610beb610ad9565b5b8135610bfc848260208601610b97565b91505092915050565b5f60208284031215610c1a57610c196109c1565b5b5f82013567ffffffffffffffff811115610c3757610c366109c5565b5b610c4384828501610bd8565b91505092915050565b5f8060408385031215610c6257610c616109c1565b5b5f610c6f858286016109e8565b925050602083013567ffffffffffffffff811115610c9057610c8f6109c5565b5b610c9c85828601610bd8565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610cea57607f821691505b602082108103610cfd57610cfc610ca6565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610d3a826109c9565b9150610d45836109c9565b9250828203905081811115610d5d57610d5c610d03565b5b92915050565b5f610d6d826109c9565b9150610d78836109c9565b9250828201905080821115610d9057610d8f610d03565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81549050610dd181610cd3565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610e347fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610df9565b610e3e8683610df9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610e79610e74610e6f846109c9565b610e56565b6109c9565b9050919050565b5f819050919050565b610e9283610e5f565b610ea6610e9e82610e80565b848454610e05565b825550505050565b5f90565b610eba610eae565b610ec5818484610e89565b505050565b5b81811015610ee857610edd5f82610eb2565b600181019050610ecb565b5050565b601f821115610f2d57610efe81610dd8565b610f0784610dea565b81016020851015610f16578190505b610f2a610f2285610dea565b830182610eca565b50505b505050565b5f82821c905092915050565b5f610f4d5f1984600802610f32565b1980831691505092915050565b5f610f658383610f3e565b9150826002028217905092915050565b818103610f83575050611058565b610f8c82610dc3565b67ffffffffffffffff811115610fa557610fa4610ae1565b5b610faf8254610cd3565b610fba828285610eec565b5f601f831160018114610fe7575f8415610fd5578287015490505b610fdf8582610f5a565b865550611051565b601f198416610ff587610dd8565b965061100086610dd8565b5f5b8281101561102757848901548255600182019150600185019450602081019050611002565b868310156110445784890154611040601f891682610f3e565b8355505b6001600288020188555050505b5050505050505b565b5f611064826109c9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361109657611095610d03565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b6110d782610828565b67ffffffffffffffff8111156110f0576110ef610ae1565b5b6110fa8254610cd3565b611105828285610eec565b5f60209050601f831160018114611136575f8415611124578287015190505b61112e8582610f5a565b865550611195565b601f19841661114486610dd8565b5f5b8281101561116b57848901518255600182019150602085019450602081019050611146565b868310156111885784890151611184601f891682610f3e565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220a4ada3e614ceb5dc72a78a97ce59acc54eb0582e6c62b0887c0b2dc93dcfa8ba64736f6c63430008140033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, _id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", _id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCallerSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactorSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactor) Remove(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "remove", _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoTransactor) Toggle(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "toggle", _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Toggle(&_Todo.TransactOpts, _id)
}

// Toggle is a paid mutator transaction binding the contract method 0x751ef753.
//
// Solidity: function toggle(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Toggle(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Toggle(&_Todo.TransactOpts, _id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactor) Update(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "update", _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactorSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}
