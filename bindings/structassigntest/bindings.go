// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package structassigntest

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

// StructassigntestABI is the input ABI used to generate the binding from.
const StructassigntestABI = "[{\"inputs\":[],\"name\":\"setMany\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setSingle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"structs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StructassigntestBin is the compiled bytecode used for deploying new contracts.
var StructassigntestBin = "0x608060405234801561001057600080fd5b50610167806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80635534193814610046578063b89867d314610081578063d25347921461009d575b600080fd5b6100636004803603602081101561005c57600080fd5b50356100a5565b60408051938452602084019290925282820152519081900360600190f35b6100896100c6565b604080519115158252519081900360200190f35b6100896100ed565b60006020819052908152604090208054600182015460029092015490919083565b43600090815260208190526040902060018082556002828201819055600392019190915590565b60408051606081018252600180825260026020808401828152600385870190815243600090815292839052958220945185555192840192909255925191909201559056fea2646970667358221220f60d5d3a45f45aaf44ee6f2ddf39642b16a99e53ca744f07068b3884c8da53a364736f6c63430007030033"

// DeployStructassigntest deploys a new Ethereum contract, binding an instance of Structassigntest to it.
func DeployStructassigntest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Structassigntest, error) {
	parsed, err := abi.JSON(strings.NewReader(StructassigntestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StructassigntestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Structassigntest{StructassigntestCaller: StructassigntestCaller{contract: contract}, StructassigntestTransactor: StructassigntestTransactor{contract: contract}, StructassigntestFilterer: StructassigntestFilterer{contract: contract}}, nil
}

// Structassigntest is an auto generated Go binding around an Ethereum contract.
type Structassigntest struct {
	StructassigntestCaller     // Read-only binding to the contract
	StructassigntestTransactor // Write-only binding to the contract
	StructassigntestFilterer   // Log filterer for contract events
}

// StructassigntestCaller is an auto generated read-only Go binding around an Ethereum contract.
type StructassigntestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructassigntestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StructassigntestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructassigntestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StructassigntestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructassigntestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StructassigntestSession struct {
	Contract     *Structassigntest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StructassigntestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StructassigntestCallerSession struct {
	Contract *StructassigntestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StructassigntestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StructassigntestTransactorSession struct {
	Contract     *StructassigntestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StructassigntestRaw is an auto generated low-level Go binding around an Ethereum contract.
type StructassigntestRaw struct {
	Contract *Structassigntest // Generic contract binding to access the raw methods on
}

// StructassigntestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StructassigntestCallerRaw struct {
	Contract *StructassigntestCaller // Generic read-only contract binding to access the raw methods on
}

// StructassigntestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StructassigntestTransactorRaw struct {
	Contract *StructassigntestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStructassigntest creates a new instance of Structassigntest, bound to a specific deployed contract.
func NewStructassigntest(address common.Address, backend bind.ContractBackend) (*Structassigntest, error) {
	contract, err := bindStructassigntest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Structassigntest{StructassigntestCaller: StructassigntestCaller{contract: contract}, StructassigntestTransactor: StructassigntestTransactor{contract: contract}, StructassigntestFilterer: StructassigntestFilterer{contract: contract}}, nil
}

// NewStructassigntestCaller creates a new read-only instance of Structassigntest, bound to a specific deployed contract.
func NewStructassigntestCaller(address common.Address, caller bind.ContractCaller) (*StructassigntestCaller, error) {
	contract, err := bindStructassigntest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StructassigntestCaller{contract: contract}, nil
}

// NewStructassigntestTransactor creates a new write-only instance of Structassigntest, bound to a specific deployed contract.
func NewStructassigntestTransactor(address common.Address, transactor bind.ContractTransactor) (*StructassigntestTransactor, error) {
	contract, err := bindStructassigntest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StructassigntestTransactor{contract: contract}, nil
}

// NewStructassigntestFilterer creates a new log filterer instance of Structassigntest, bound to a specific deployed contract.
func NewStructassigntestFilterer(address common.Address, filterer bind.ContractFilterer) (*StructassigntestFilterer, error) {
	contract, err := bindStructassigntest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StructassigntestFilterer{contract: contract}, nil
}

// bindStructassigntest binds a generic wrapper to an already deployed contract.
func bindStructassigntest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StructassigntestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Structassigntest *StructassigntestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Structassigntest.Contract.StructassigntestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Structassigntest *StructassigntestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structassigntest.Contract.StructassigntestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Structassigntest *StructassigntestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Structassigntest.Contract.StructassigntestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Structassigntest *StructassigntestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Structassigntest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Structassigntest *StructassigntestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structassigntest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Structassigntest *StructassigntestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Structassigntest.Contract.contract.Transact(opts, method, params...)
}

// Structs is a free data retrieval call binding the contract method 0x55341938.
//
// Solidity: function structs(uint256 ) view returns(uint256 a, uint256 b, uint256 c)
func (_Structassigntest *StructassigntestCaller) Structs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	A *big.Int
	B *big.Int
	C *big.Int
}, error) {
	var out []interface{}
	err := _Structassigntest.contract.Call(opts, &out, "structs", arg0)

	outstruct := new(struct {
		A *big.Int
		B *big.Int
		C *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.A = out[0].(*big.Int)
	outstruct.B = out[1].(*big.Int)
	outstruct.C = out[2].(*big.Int)

	return *outstruct, err

}

// Structs is a free data retrieval call binding the contract method 0x55341938.
//
// Solidity: function structs(uint256 ) view returns(uint256 a, uint256 b, uint256 c)
func (_Structassigntest *StructassigntestSession) Structs(arg0 *big.Int) (struct {
	A *big.Int
	B *big.Int
	C *big.Int
}, error) {
	return _Structassigntest.Contract.Structs(&_Structassigntest.CallOpts, arg0)
}

// Structs is a free data retrieval call binding the contract method 0x55341938.
//
// Solidity: function structs(uint256 ) view returns(uint256 a, uint256 b, uint256 c)
func (_Structassigntest *StructassigntestCallerSession) Structs(arg0 *big.Int) (struct {
	A *big.Int
	B *big.Int
	C *big.Int
}, error) {
	return _Structassigntest.Contract.Structs(&_Structassigntest.CallOpts, arg0)
}

// SetMany is a paid mutator transaction binding the contract method 0xb89867d3.
//
// Solidity: function setMany() returns(bool)
func (_Structassigntest *StructassigntestTransactor) SetMany(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structassigntest.contract.Transact(opts, "setMany")
}

// SetMany is a paid mutator transaction binding the contract method 0xb89867d3.
//
// Solidity: function setMany() returns(bool)
func (_Structassigntest *StructassigntestSession) SetMany() (*types.Transaction, error) {
	return _Structassigntest.Contract.SetMany(&_Structassigntest.TransactOpts)
}

// SetMany is a paid mutator transaction binding the contract method 0xb89867d3.
//
// Solidity: function setMany() returns(bool)
func (_Structassigntest *StructassigntestTransactorSession) SetMany() (*types.Transaction, error) {
	return _Structassigntest.Contract.SetMany(&_Structassigntest.TransactOpts)
}

// SetSingle is a paid mutator transaction binding the contract method 0xd2534792.
//
// Solidity: function setSingle() returns(bool)
func (_Structassigntest *StructassigntestTransactor) SetSingle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Structassigntest.contract.Transact(opts, "setSingle")
}

// SetSingle is a paid mutator transaction binding the contract method 0xd2534792.
//
// Solidity: function setSingle() returns(bool)
func (_Structassigntest *StructassigntestSession) SetSingle() (*types.Transaction, error) {
	return _Structassigntest.Contract.SetSingle(&_Structassigntest.TransactOpts)
}

// SetSingle is a paid mutator transaction binding the contract method 0xd2534792.
//
// Solidity: function setSingle() returns(bool)
func (_Structassigntest *StructassigntestTransactorSession) SetSingle() (*types.Transaction, error) {
	return _Structassigntest.Contract.SetSingle(&_Structassigntest.TransactOpts)
}
