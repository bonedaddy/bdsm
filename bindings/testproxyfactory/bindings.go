// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testproxyfactory

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

// TestproxyfactoryABI is the input ABI used to generate the binding from.
const TestproxyfactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializerData\",\"type\":\"bytes\"}],\"name\":\"deployInstance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializerData\",\"type\":\"bytes\"}],\"name\":\"deployProxyContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numProxies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proxies\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TestproxyfactoryBin is the compiled bytecode used for deploying new contracts.
var TestproxyfactoryBin = "0x608060405234801561001057600080fd5b50604051610020602082016100c9565b6020820181038252601f19601f820116604052506100478161006d60201b6103011760201c565b51600080546001600160a01b0319166001600160a01b03909216919091179055506100e8565b6100756100d6565b6000610080836100aa565b905061008a6100d6565b5060408051602081019091526001600160a01b0390911681529050919050565b6000808251602084016000f09050803b6100c357600080fd5b92915050565b61018780610e5a83390190565b60408051602081019091526000815290565b610d63806100f76000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631028f90a1461005157806370bf749714610123578063769af69d1461013d578063abd90f85146101f3575b600080fd5b6101076004803603604081101561006757600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561009257600080fd5b8201836020820111156100a457600080fd5b803590602001918460018302840111640100000000831117156100c657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610210945050505050565b604080516001600160a01b039092168252519081900360200190f35b61012b6102bc565b60408051918252519081900360200190f35b6101076004803603604081101561015357600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561017e57600080fd5b82018360208201111561019057600080fd5b803590602001918460018302840111640100000000831117156101b257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102c2945050505050565b6101076004803603602081101561020957600080fd5b50356102d7565b60008061021f6000858561033e565b604080516001600160a01b038316815290519192507efffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349919081900360200190a160028054600180820183556000929092527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b03939093169290921790915580548101905592915050565b60015481565b60006102d06000848461033e565b9392505050565b600281815481106102e757600080fd5b6000918252602090912001546001600160a01b0316905081565b610309610432565b600061031483610413565b905061031e610432565b5060408051602081019091526001600160a01b0390911681529050919050565b825460405160009182916001600160a01b03909116908590859061036190610444565b80846001600160a01b03168152602001836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103b95781810151838201526020016103a1565b50505050905090810190601f1680156103e65780820380516001836020036101000a031916815260200191505b50945050505050604051809103906000f080158015610409573d6000803e3d6000fd5b5095945050505050565b6000808251602084016000f09050803b61042c57600080fd5b92915050565b60408051602081019091526000815290565b6108dc806104528339019056fe60806040526040516108dc3803806108dc8339818101604052606081101561002657600080fd5b8151602083015160408085018051915193959294830192918464010000000082111561005157600080fd5b90830190602082018581111561006657600080fd5b825164010000000081118282018810171561008057600080fd5b82525081516020918201929091019080838360005b838110156100ad578181015183820152602001610095565b50505050905090810190601f1680156100da5780820380516001836020036101000a031916815260200191505b50604052508491508290506100ee826101bf565b8051156101a6576000826001600160a01b0316826040518082805190602001908083835b602083106101315780518252601f199092019160209182019101610112565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610191576040519150601f19603f3d011682016040523d82523d6000602084013e610196565b606091505b50509050806101a457600080fd5b505b506101ae9050565b6101b782610231565b50505061025b565b6101d28161025560201b6103b41760201c565b61020d5760405162461bcd60e51b81526004018080602001828103825260368152602001806108a66036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b3b151590565b61063c8061026a6000396000f3fe60806040526004361061004e5760003560e01c80633659cfe6146100655780634f1ef286146100985780635c60da1b146101185780638f28397014610149578063f851a4401461017c5761005d565b3661005d5761005b610191565b005b61005b610191565b34801561007157600080fd5b5061005b6004803603602081101561008857600080fd5b50356001600160a01b03166101ab565b61005b600480360360408110156100ae57600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100d957600080fd5b8201836020820111156100eb57600080fd5b8035906020019184600183028401116401000000008311171561010d57600080fd5b5090925090506101e5565b34801561012457600080fd5b5061012d610292565b604080516001600160a01b039092168252519081900360200190f35b34801561015557600080fd5b5061005b6004803603602081101561016c57600080fd5b50356001600160a01b03166102cf565b34801561018857600080fd5b5061012d610389565b6101996103ba565b6101a96101a461041a565b61043f565b565b6101b3610463565b6001600160a01b0316336001600160a01b031614156101da576101d581610488565b6101e2565b6101e2610191565b50565b6101ed610463565b6001600160a01b0316336001600160a01b031614156102855761020f83610488565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d806000811461026c576040519150601f19603f3d011682016040523d82523d6000602084013e610271565b606091505b505090508061027f57600080fd5b5061028d565b61028d610191565b505050565b600061029c610463565b6001600160a01b0316336001600160a01b031614156102c4576102bd61041a565b90506102cc565b6102cc610191565b90565b6102d7610463565b6001600160a01b0316336001600160a01b031614156101da576001600160a01b0381166103355760405162461bcd60e51b815260040180806020018281038252603a815260200180610555603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61035e610463565b604080516001600160a01b03928316815291841660208301528051918290030190a16101d5816104c8565b6000610393610463565b6001600160a01b0316336001600160a01b031614156102c4576102bd610463565b3b151590565b6103c2610463565b6001600160a01b0316336001600160a01b031614156104125760405162461bcd60e51b81526004018080602001828103825260428152602001806105c56042913960600191505060405180910390fd5b6101a96101a9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e80801561045e573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b610491816104ec565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b6104f5816103b4565b6105305760405162461bcd60e51b815260040180806020018281038252603681526020018061058f6036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5556fe5472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f20616464726573735570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a2646970667358221220cac8759456cf94da768b83d81b29189b50849b642b839f7a932597dfe603e6b964736f6c634300070400335570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374a2646970667358221220049ce6acd5a9050e54c79181896c44e8733f77e257af259783d3970d8731f1ab64736f6c63430007040033608060405234801561001057600080fd5b50610167806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80635534193814610046578063b89867d314610081578063d25347921461009d575b600080fd5b6100636004803603602081101561005c57600080fd5b50356100a5565b60408051938452602084019290925282820152519081900360600190f35b6100896100c6565b604080519115158252519081900360200190f35b6100896100ed565b60006020819052908152604090208054600182015460029092015490919083565b43600090815260208190526040902060018082556002828201819055600392019190915590565b60408051606081018252600180825260026020808401828152600385870190815243600090815292839052958220945185555192840192909255925191909201559056fea2646970667358221220e66c5b548fe9005de184d43dc7a83b00a32cf9f1e598e33aae97164457bba3ee64736f6c63430007040033"

// DeployTestproxyfactory deploys a new Ethereum contract, binding an instance of Testproxyfactory to it.
func DeployTestproxyfactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Testproxyfactory, error) {
	parsed, err := abi.JSON(strings.NewReader(TestproxyfactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestproxyfactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Testproxyfactory{TestproxyfactoryCaller: TestproxyfactoryCaller{contract: contract}, TestproxyfactoryTransactor: TestproxyfactoryTransactor{contract: contract}, TestproxyfactoryFilterer: TestproxyfactoryFilterer{contract: contract}}, nil
}

// Testproxyfactory is an auto generated Go binding around an Ethereum contract.
type Testproxyfactory struct {
	TestproxyfactoryCaller     // Read-only binding to the contract
	TestproxyfactoryTransactor // Write-only binding to the contract
	TestproxyfactoryFilterer   // Log filterer for contract events
}

// TestproxyfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestproxyfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestproxyfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestproxyfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestproxyfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestproxyfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestproxyfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestproxyfactorySession struct {
	Contract     *Testproxyfactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestproxyfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestproxyfactoryCallerSession struct {
	Contract *TestproxyfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TestproxyfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestproxyfactoryTransactorSession struct {
	Contract     *TestproxyfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TestproxyfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestproxyfactoryRaw struct {
	Contract *Testproxyfactory // Generic contract binding to access the raw methods on
}

// TestproxyfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestproxyfactoryCallerRaw struct {
	Contract *TestproxyfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TestproxyfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestproxyfactoryTransactorRaw struct {
	Contract *TestproxyfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestproxyfactory creates a new instance of Testproxyfactory, bound to a specific deployed contract.
func NewTestproxyfactory(address common.Address, backend bind.ContractBackend) (*Testproxyfactory, error) {
	contract, err := bindTestproxyfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Testproxyfactory{TestproxyfactoryCaller: TestproxyfactoryCaller{contract: contract}, TestproxyfactoryTransactor: TestproxyfactoryTransactor{contract: contract}, TestproxyfactoryFilterer: TestproxyfactoryFilterer{contract: contract}}, nil
}

// NewTestproxyfactoryCaller creates a new read-only instance of Testproxyfactory, bound to a specific deployed contract.
func NewTestproxyfactoryCaller(address common.Address, caller bind.ContractCaller) (*TestproxyfactoryCaller, error) {
	contract, err := bindTestproxyfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestproxyfactoryCaller{contract: contract}, nil
}

// NewTestproxyfactoryTransactor creates a new write-only instance of Testproxyfactory, bound to a specific deployed contract.
func NewTestproxyfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TestproxyfactoryTransactor, error) {
	contract, err := bindTestproxyfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestproxyfactoryTransactor{contract: contract}, nil
}

// NewTestproxyfactoryFilterer creates a new log filterer instance of Testproxyfactory, bound to a specific deployed contract.
func NewTestproxyfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TestproxyfactoryFilterer, error) {
	contract, err := bindTestproxyfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestproxyfactoryFilterer{contract: contract}, nil
}

// bindTestproxyfactory binds a generic wrapper to an already deployed contract.
func bindTestproxyfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestproxyfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testproxyfactory *TestproxyfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testproxyfactory.Contract.TestproxyfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testproxyfactory *TestproxyfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testproxyfactory.Contract.TestproxyfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testproxyfactory *TestproxyfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testproxyfactory.Contract.TestproxyfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testproxyfactory *TestproxyfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testproxyfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testproxyfactory *TestproxyfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testproxyfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testproxyfactory *TestproxyfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testproxyfactory.Contract.contract.Transact(opts, method, params...)
}

// NumProxies is a free data retrieval call binding the contract method 0x70bf7497.
//
// Solidity: function numProxies() view returns(uint256)
func (_Testproxyfactory *TestproxyfactoryCaller) NumProxies(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Testproxyfactory.contract.Call(opts, &out, "numProxies")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumProxies is a free data retrieval call binding the contract method 0x70bf7497.
//
// Solidity: function numProxies() view returns(uint256)
func (_Testproxyfactory *TestproxyfactorySession) NumProxies() (*big.Int, error) {
	return _Testproxyfactory.Contract.NumProxies(&_Testproxyfactory.CallOpts)
}

// NumProxies is a free data retrieval call binding the contract method 0x70bf7497.
//
// Solidity: function numProxies() view returns(uint256)
func (_Testproxyfactory *TestproxyfactoryCallerSession) NumProxies() (*big.Int, error) {
	return _Testproxyfactory.Contract.NumProxies(&_Testproxyfactory.CallOpts)
}

// Proxies is a free data retrieval call binding the contract method 0xabd90f85.
//
// Solidity: function proxies(uint256 ) view returns(address)
func (_Testproxyfactory *TestproxyfactoryCaller) Proxies(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Testproxyfactory.contract.Call(opts, &out, "proxies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxies is a free data retrieval call binding the contract method 0xabd90f85.
//
// Solidity: function proxies(uint256 ) view returns(address)
func (_Testproxyfactory *TestproxyfactorySession) Proxies(arg0 *big.Int) (common.Address, error) {
	return _Testproxyfactory.Contract.Proxies(&_Testproxyfactory.CallOpts, arg0)
}

// Proxies is a free data retrieval call binding the contract method 0xabd90f85.
//
// Solidity: function proxies(uint256 ) view returns(address)
func (_Testproxyfactory *TestproxyfactoryCallerSession) Proxies(arg0 *big.Int) (common.Address, error) {
	return _Testproxyfactory.Contract.Proxies(&_Testproxyfactory.CallOpts, arg0)
}

// DeployInstance is a paid mutator transaction binding the contract method 0x769af69d.
//
// Solidity: function deployInstance(address admin, bytes initializerData) returns(address)
func (_Testproxyfactory *TestproxyfactoryTransactor) DeployInstance(opts *bind.TransactOpts, admin common.Address, initializerData []byte) (*types.Transaction, error) {
	return _Testproxyfactory.contract.Transact(opts, "deployInstance", admin, initializerData)
}

// DeployInstance is a paid mutator transaction binding the contract method 0x769af69d.
//
// Solidity: function deployInstance(address admin, bytes initializerData) returns(address)
func (_Testproxyfactory *TestproxyfactorySession) DeployInstance(admin common.Address, initializerData []byte) (*types.Transaction, error) {
	return _Testproxyfactory.Contract.DeployInstance(&_Testproxyfactory.TransactOpts, admin, initializerData)
}

// DeployInstance is a paid mutator transaction binding the contract method 0x769af69d.
//
// Solidity: function deployInstance(address admin, bytes initializerData) returns(address)
func (_Testproxyfactory *TestproxyfactoryTransactorSession) DeployInstance(admin common.Address, initializerData []byte) (*types.Transaction, error) {
	return _Testproxyfactory.Contract.DeployInstance(&_Testproxyfactory.TransactOpts, admin, initializerData)
}

// DeployProxyContract is a paid mutator transaction binding the contract method 0x1028f90a.
//
// Solidity: function deployProxyContract(address admin, bytes initializerData) returns(address)
func (_Testproxyfactory *TestproxyfactoryTransactor) DeployProxyContract(opts *bind.TransactOpts, admin common.Address, initializerData []byte) (*types.Transaction, error) {
	return _Testproxyfactory.contract.Transact(opts, "deployProxyContract", admin, initializerData)
}

// DeployProxyContract is a paid mutator transaction binding the contract method 0x1028f90a.
//
// Solidity: function deployProxyContract(address admin, bytes initializerData) returns(address)
func (_Testproxyfactory *TestproxyfactorySession) DeployProxyContract(admin common.Address, initializerData []byte) (*types.Transaction, error) {
	return _Testproxyfactory.Contract.DeployProxyContract(&_Testproxyfactory.TransactOpts, admin, initializerData)
}

// DeployProxyContract is a paid mutator transaction binding the contract method 0x1028f90a.
//
// Solidity: function deployProxyContract(address admin, bytes initializerData) returns(address)
func (_Testproxyfactory *TestproxyfactoryTransactorSession) DeployProxyContract(admin common.Address, initializerData []byte) (*types.Transaction, error) {
	return _Testproxyfactory.Contract.DeployProxyContract(&_Testproxyfactory.TransactOpts, admin, initializerData)
}

// TestproxyfactoryProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the Testproxyfactory contract.
type TestproxyfactoryProxyCreatedIterator struct {
	Event *TestproxyfactoryProxyCreated // Event containing the contract specifics and raw log

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
func (it *TestproxyfactoryProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestproxyfactoryProxyCreated)
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
		it.Event = new(TestproxyfactoryProxyCreated)
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
func (it *TestproxyfactoryProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestproxyfactoryProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestproxyfactoryProxyCreated represents a ProxyCreated event raised by the Testproxyfactory contract.
type TestproxyfactoryProxyCreated struct {
	Proxy common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address _proxy)
func (_Testproxyfactory *TestproxyfactoryFilterer) FilterProxyCreated(opts *bind.FilterOpts) (*TestproxyfactoryProxyCreatedIterator, error) {

	logs, sub, err := _Testproxyfactory.contract.FilterLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return &TestproxyfactoryProxyCreatedIterator{contract: _Testproxyfactory.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address _proxy)
func (_Testproxyfactory *TestproxyfactoryFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *TestproxyfactoryProxyCreated) (event.Subscription, error) {

	logs, sub, err := _Testproxyfactory.contract.WatchLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestproxyfactoryProxyCreated)
				if err := _Testproxyfactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: event ProxyCreated(address _proxy)
func (_Testproxyfactory *TestproxyfactoryFilterer) ParseProxyCreated(log types.Log) (*TestproxyfactoryProxyCreated, error) {
	event := new(TestproxyfactoryProxyCreated)
	if err := _Testproxyfactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}
