// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind_nested_array

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TestContractABI is the input ABI used to generate the binding from.
const TestContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"arr\",\"type\":\"uint64[3][4][5]\"}],\"name\":\"storeDeepUintArray\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"retrieveDeepArray\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64[3][4][5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deepUint64Array\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TestContractBin is the compiled bytecode used for deploying new contracts.
const TestContractBin = `0x6060604052341561000f57600080fd5b61059a8061001e6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166334424855811461005b5780638ed4573a1461010357806398ed185614610193575b600080fd5b341561006657600080fd5b610101600461078481600560a060405190810160405291906000835b828210156100f457610180820284016004608060405190810160405291906000835b828210156100e1578382606002016003806020026040519081016040529190828260608082843750505091835250506001909101906020016100a4565b5050505081526020019060010190610082565b50505050919050506101cc565b005b341561010e57600080fd5b6101166101dd565b6040516000826005835b81841015610183578284602002015160046000925b818410156101755782846020020151606080838360005b8381101561016457808201518382015260200161014c565b505050509050019260010192610135565b925050509260010192610120565b9250505091505060405180910390f35b341561019e57600080fd5b6101af6004356024356044356102b3565b60405167ffffffffffffffff909116815260200160405180910390f35b6101d96000826005610305565b5050565b6101e5610350565b6000600560a060405190810160405291906000835b828210156102a9576004808302850190608060405190810160405291906000835b828210156102965781840160036060604051908101604052919060608301826000855b82829054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001906008019060208260070104928301926001038202915080841161023e57905050505050508152602001906001019061021b565b50505050815260200190600101906101fa565b5050505090505b90565b600083600581106102c057fe5b60040201826004811015156102d157fe5b0181600381106102dd57fe5b6004918282040191900660080292509250509054906101000a900467ffffffffffffffff1681565b601483019183908215610340579160200282015b82811115610340578251610330908390600461037e565b5091602001919060040190610319565b5061034c9291506103c4565b5090565b6107806040519081016040526005815b6103686103e7565b8152602001906001900390816103605790505090565b82600481019282156103b8579160200282015b828111156103b85782516103a89083906003610415565b5091602001919060010190610391565b5061034c9291506104b9565b6102b091905b8082111561034c5760006103de82826104dc565b506004016103ca565b6101806040519081016040526004815b6103ff61051a565b8152602001906001900390816103f75790505090565b6001830191839082156104ad5791602002820160005b8382111561047757835183826101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550926020019260080160208160070104928301926001030261042b565b80156104ab5782816101000a81549067ffffffffffffffff0219169055600801602081600701049283019260010302610477565b505b5061034c929150610542565b6102b091905b8082111561034c5760006104d38282610567565b506001016104bf565b5060006104e98282610567565b5060010160006104f98282610567565b5060010160006105098282610567565b50610518906001016000610567565b565b60606040519081016040526003815b6000815260001990910190602001816105295790505090565b6102b091905b8082111561034c57805467ffffffffffffffff19168155600101610548565b50600090555600a165627a7a723058209ca2279bce44aafbc811fc7d943983c66a456788b6d5c0313b90458b8012fe9e0029`

// DeployTestContract deploys a new Ethereum contract, binding an instance of TestContract to it.
func DeployTestContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestContract{TestContractCaller: TestContractCaller{contract: contract}, TestContractTransactor: TestContractTransactor{contract: contract}}, nil
}

// TestContract is an auto generated Go binding around an Ethereum contract.
type TestContract struct {
	TestContractCaller     // Read-only binding to the contract
	TestContractTransactor // Write-only binding to the contract
}

// TestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestContractSession struct {
	Contract     *TestContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestContractCallerSession struct {
	Contract *TestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestContractTransactorSession struct {
	Contract     *TestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestContractRaw struct {
	Contract *TestContract // Generic contract binding to access the raw methods on
}

// TestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestContractCallerRaw struct {
	Contract *TestContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestContractTransactorRaw struct {
	Contract *TestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestContract creates a new instance of TestContract, bound to a specific deployed contract.
func NewTestContract(address common.Address, backend bind.ContractBackend) (*TestContract, error) {
	contract, err := bindTestContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestContract{TestContractCaller: TestContractCaller{contract: contract}, TestContractTransactor: TestContractTransactor{contract: contract}}, nil
}

// NewTestContractCaller creates a new read-only instance of TestContract, bound to a specific deployed contract.
func NewTestContractCaller(address common.Address, caller bind.ContractCaller) (*TestContractCaller, error) {
	contract, err := bindTestContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractCaller{contract: contract}, nil
}

// NewTestContractTransactor creates a new write-only instance of TestContract, bound to a specific deployed contract.
func NewTestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestContractTransactor, error) {
	contract, err := bindTestContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TestContractTransactor{contract: contract}, nil
}

// bindTestContract binds a generic wrapper to an already deployed contract.
func bindTestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContract *TestContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestContract.Contract.TestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContract *TestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContract.Contract.TestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContract *TestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContract.Contract.TestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContract *TestContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContract *TestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContract *TestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContract.Contract.contract.Transact(opts, method, params...)
}

// DeepUint64Array is a free data retrieval call binding the contract method 0x98ed1856.
//
// Solidity: function deepUint64Array( uint256,  uint256,  uint256) constant returns(uint64)
func (_TestContract *TestContractCaller) DeepUint64Array(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _TestContract.contract.Call(opts, out, "deepUint64Array", arg0, arg1, arg2)
	return *ret0, err
}

// DeepUint64Array is a free data retrieval call binding the contract method 0x98ed1856.
//
// Solidity: function deepUint64Array( uint256,  uint256,  uint256) constant returns(uint64)
func (_TestContract *TestContractSession) DeepUint64Array(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (uint64, error) {
	return _TestContract.Contract.DeepUint64Array(&_TestContract.CallOpts, arg0, arg1, arg2)
}

// DeepUint64Array is a free data retrieval call binding the contract method 0x98ed1856.
//
// Solidity: function deepUint64Array( uint256,  uint256,  uint256) constant returns(uint64)
func (_TestContract *TestContractCallerSession) DeepUint64Array(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (uint64, error) {
	return _TestContract.Contract.DeepUint64Array(&_TestContract.CallOpts, arg0, arg1, arg2)
}

// RetrieveDeepArray is a free data retrieval call binding the contract method 0x8ed4573a.
//
// Solidity: function retrieveDeepArray() constant returns(uint64[3][4][5])
func (_TestContract *TestContractCaller) RetrieveDeepArray(opts *bind.CallOpts) ([5][4][3]uint64, error) {
	var (
		ret0 = new([5][4][3]uint64)
	)
	out := ret0
	err := _TestContract.contract.Call(opts, out, "retrieveDeepArray")
	return *ret0, err
}

// RetrieveDeepArray is a free data retrieval call binding the contract method 0x8ed4573a.
//
// Solidity: function retrieveDeepArray() constant returns(uint64[3][4][5])
func (_TestContract *TestContractSession) RetrieveDeepArray() ([5][4][3]uint64, error) {
	return _TestContract.Contract.RetrieveDeepArray(&_TestContract.CallOpts)
}

// RetrieveDeepArray is a free data retrieval call binding the contract method 0x8ed4573a.
//
// Solidity: function retrieveDeepArray() constant returns(uint64[3][4][5])
func (_TestContract *TestContractCallerSession) RetrieveDeepArray() ([5][4][3]uint64, error) {
	return _TestContract.Contract.RetrieveDeepArray(&_TestContract.CallOpts)
}

// StoreDeepUintArray is a paid mutator transaction binding the contract method 0x34424855.
//
// Solidity: function storeDeepUintArray(arr uint64[3][4][5]) returns()
func (_TestContract *TestContractTransactor) StoreDeepUintArray(opts *bind.TransactOpts, arr [5][4][3]uint64) (*types.Transaction, error) {
	return _TestContract.contract.Transact(opts, "storeDeepUintArray", arr)
}

// StoreDeepUintArray is a paid mutator transaction binding the contract method 0x34424855.
//
// Solidity: function storeDeepUintArray(arr uint64[3][4][5]) returns()
func (_TestContract *TestContractSession) StoreDeepUintArray(arr [5][4][3]uint64) (*types.Transaction, error) {
	return _TestContract.Contract.StoreDeepUintArray(&_TestContract.TransactOpts, arr)
}

// StoreDeepUintArray is a paid mutator transaction binding the contract method 0x34424855.
//
// Solidity: function storeDeepUintArray(arr uint64[3][4][5]) returns()
func (_TestContract *TestContractTransactorSession) StoreDeepUintArray(arr [5][4][3]uint64) (*types.Transaction, error) {
	return _TestContract.Contract.StoreDeepUintArray(&_TestContract.TransactOpts, arr)
}
