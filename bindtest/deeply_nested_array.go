// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindtest

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DeeplyNestedArrayABI is the input ABI used to generate the binding from.
const DeeplyNestedArrayABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"arr\",\"type\":\"uint64[3][4][5]\"}],\"name\":\"storeDeepUintArray\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"retrieveDeepArray\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64[3][4][5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deepUint64Array\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DeeplyNestedArrayBin is the compiled bytecode used for deploying new contracts.
const DeeplyNestedArrayBin = `0x6060604052341561000f57600080fd5b61059a8061001e6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166334424855811461005b5780638ed4573a1461010357806398ed185614610193575b600080fd5b341561006657600080fd5b610101600461078481600560a060405190810160405291906000835b828210156100f457610180820284016004608060405190810160405291906000835b828210156100e1578382606002016003806020026040519081016040529190828260608082843750505091835250506001909101906020016100a4565b5050505081526020019060010190610082565b50505050919050506101cc565b005b341561010e57600080fd5b6101166101dd565b6040516000826005835b81841015610183578284602002015160046000925b818410156101755782846020020151606080838360005b8381101561016457808201518382015260200161014c565b505050509050019260010192610135565b925050509260010192610120565b9250505091505060405180910390f35b341561019e57600080fd5b6101af6004356024356044356102b3565b60405167ffffffffffffffff909116815260200160405180910390f35b6101d96000826005610305565b5050565b6101e5610350565b6000600560a060405190810160405291906000835b828210156102a9576004808302850190608060405190810160405291906000835b828210156102965781840160036060604051908101604052919060608301826000855b82829054906101000a900467ffffffffffffffff1667ffffffffffffffff168152602001906008019060208260070104928301926001038202915080841161023e57905050505050508152602001906001019061021b565b50505050815260200190600101906101fa565b5050505090505b90565b600083600581106102c057fe5b60040201826004811015156102d157fe5b0181600381106102dd57fe5b6004918282040191900660080292509250509054906101000a900467ffffffffffffffff1681565b601483019183908215610340579160200282015b82811115610340578251610330908390600461037e565b5091602001919060040190610319565b5061034c9291506103c4565b5090565b6107806040519081016040526005815b6103686103e7565b8152602001906001900390816103605790505090565b82600481019282156103b8579160200282015b828111156103b85782516103a89083906003610415565b5091602001919060010190610391565b5061034c9291506104b9565b6102b091905b8082111561034c5760006103de82826104dc565b506004016103ca565b6101806040519081016040526004815b6103ff61051a565b8152602001906001900390816103f75790505090565b6001830191839082156104ad5791602002820160005b8382111561047757835183826101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550926020019260080160208160070104928301926001030261042b565b80156104ab5782816101000a81549067ffffffffffffffff0219169055600801602081600701049283019260010302610477565b505b5061034c929150610542565b6102b091905b8082111561034c5760006104d38282610567565b506001016104bf565b5060006104e98282610567565b5060010160006104f98282610567565b5060010160006105098282610567565b50610518906001016000610567565b565b60606040519081016040526003815b6000815260001990910190602001816105295790505090565b6102b091905b8082111561034c57805467ffffffffffffffff19168155600101610548565b50600090555600a165627a7a72305820bffabbb72cc4f911fc83c87ed454086fb39b516c32ceff116f004c16419d27e50029`

// DeployDeeplyNestedArray deploys a new Ethereum contract, binding an instance of DeeplyNestedArray to it.
func DeployDeeplyNestedArray(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DeeplyNestedArray, error) {
	parsed, err := abi.JSON(strings.NewReader(DeeplyNestedArrayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DeeplyNestedArrayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DeeplyNestedArray{DeeplyNestedArrayCaller: DeeplyNestedArrayCaller{contract: contract}, DeeplyNestedArrayTransactor: DeeplyNestedArrayTransactor{contract: contract}}, nil
}

// DeeplyNestedArray is an auto generated Go binding around an Ethereum contract.
type DeeplyNestedArray struct {
	DeeplyNestedArrayCaller     // Read-only binding to the contract
	DeeplyNestedArrayTransactor // Write-only binding to the contract
}

// DeeplyNestedArrayCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeeplyNestedArrayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeeplyNestedArrayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeeplyNestedArrayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeeplyNestedArraySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeeplyNestedArraySession struct {
	Contract     *DeeplyNestedArray // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DeeplyNestedArrayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeeplyNestedArrayCallerSession struct {
	Contract *DeeplyNestedArrayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DeeplyNestedArrayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeeplyNestedArrayTransactorSession struct {
	Contract     *DeeplyNestedArrayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DeeplyNestedArrayRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeeplyNestedArrayRaw struct {
	Contract *DeeplyNestedArray // Generic contract binding to access the raw methods on
}

// DeeplyNestedArrayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeeplyNestedArrayCallerRaw struct {
	Contract *DeeplyNestedArrayCaller // Generic read-only contract binding to access the raw methods on
}

// DeeplyNestedArrayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeeplyNestedArrayTransactorRaw struct {
	Contract *DeeplyNestedArrayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeeplyNestedArray creates a new instance of DeeplyNestedArray, bound to a specific deployed contract.
func NewDeeplyNestedArray(address common.Address, backend bind.ContractBackend) (*DeeplyNestedArray, error) {
	contract, err := bindDeeplyNestedArray(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeeplyNestedArray{DeeplyNestedArrayCaller: DeeplyNestedArrayCaller{contract: contract}, DeeplyNestedArrayTransactor: DeeplyNestedArrayTransactor{contract: contract}}, nil
}

// NewDeeplyNestedArrayCaller creates a new read-only instance of DeeplyNestedArray, bound to a specific deployed contract.
func NewDeeplyNestedArrayCaller(address common.Address, caller bind.ContractCaller) (*DeeplyNestedArrayCaller, error) {
	contract, err := bindDeeplyNestedArray(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DeeplyNestedArrayCaller{contract: contract}, nil
}

// NewDeeplyNestedArrayTransactor creates a new write-only instance of DeeplyNestedArray, bound to a specific deployed contract.
func NewDeeplyNestedArrayTransactor(address common.Address, transactor bind.ContractTransactor) (*DeeplyNestedArrayTransactor, error) {
	contract, err := bindDeeplyNestedArray(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DeeplyNestedArrayTransactor{contract: contract}, nil
}

// bindDeeplyNestedArray binds a generic wrapper to an already deployed contract.
func bindDeeplyNestedArray(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeeplyNestedArrayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeeplyNestedArray *DeeplyNestedArrayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DeeplyNestedArray.Contract.DeeplyNestedArrayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeeplyNestedArray *DeeplyNestedArrayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeeplyNestedArray.Contract.DeeplyNestedArrayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeeplyNestedArray *DeeplyNestedArrayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeeplyNestedArray.Contract.DeeplyNestedArrayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeeplyNestedArray *DeeplyNestedArrayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DeeplyNestedArray.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeeplyNestedArray *DeeplyNestedArrayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeeplyNestedArray.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeeplyNestedArray *DeeplyNestedArrayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeeplyNestedArray.Contract.contract.Transact(opts, method, params...)
}

// DeepUint64Array is a free data retrieval call binding the contract method 0x98ed1856.
//
// Solidity: function deepUint64Array( uint256,  uint256,  uint256) constant returns(uint64)
func (_DeeplyNestedArray *DeeplyNestedArrayCaller) DeepUint64Array(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _DeeplyNestedArray.contract.Call(opts, out, "deepUint64Array", arg0, arg1, arg2)
	return *ret0, err
}

// DeepUint64Array is a free data retrieval call binding the contract method 0x98ed1856.
//
// Solidity: function deepUint64Array( uint256,  uint256,  uint256) constant returns(uint64)
func (_DeeplyNestedArray *DeeplyNestedArraySession) DeepUint64Array(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (uint64, error) {
	return _DeeplyNestedArray.Contract.DeepUint64Array(&_DeeplyNestedArray.CallOpts, arg0, arg1, arg2)
}

// DeepUint64Array is a free data retrieval call binding the contract method 0x98ed1856.
//
// Solidity: function deepUint64Array( uint256,  uint256,  uint256) constant returns(uint64)
func (_DeeplyNestedArray *DeeplyNestedArrayCallerSession) DeepUint64Array(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (uint64, error) {
	return _DeeplyNestedArray.Contract.DeepUint64Array(&_DeeplyNestedArray.CallOpts, arg0, arg1, arg2)
}

// RetrieveDeepArray is a free data retrieval call binding the contract method 0x8ed4573a.
//
// Solidity: function retrieveDeepArray() constant returns(uint64[3][4][5])
func (_DeeplyNestedArray *DeeplyNestedArrayCaller) RetrieveDeepArray(opts *bind.CallOpts) ([5][4][3]uint64, error) {
	var (
		ret0 = new([5][4][3]uint64)
	)
	out := ret0
	err := _DeeplyNestedArray.contract.Call(opts, out, "retrieveDeepArray")
	return *ret0, err
}

// RetrieveDeepArray is a free data retrieval call binding the contract method 0x8ed4573a.
//
// Solidity: function retrieveDeepArray() constant returns(uint64[3][4][5])
func (_DeeplyNestedArray *DeeplyNestedArraySession) RetrieveDeepArray() ([5][4][3]uint64, error) {
	return _DeeplyNestedArray.Contract.RetrieveDeepArray(&_DeeplyNestedArray.CallOpts)
}

// RetrieveDeepArray is a free data retrieval call binding the contract method 0x8ed4573a.
//
// Solidity: function retrieveDeepArray() constant returns(uint64[3][4][5])
func (_DeeplyNestedArray *DeeplyNestedArrayCallerSession) RetrieveDeepArray() ([5][4][3]uint64, error) {
	return _DeeplyNestedArray.Contract.RetrieveDeepArray(&_DeeplyNestedArray.CallOpts)
}

// StoreDeepUintArray is a paid mutator transaction binding the contract method 0x34424855.
//
// Solidity: function storeDeepUintArray(arr uint64[3][4][5]) returns()
func (_DeeplyNestedArray *DeeplyNestedArrayTransactor) StoreDeepUintArray(opts *bind.TransactOpts, arr [5][4][3]uint64) (*types.Transaction, error) {
	return _DeeplyNestedArray.contract.Transact(opts, "storeDeepUintArray", arr)
}

// StoreDeepUintArray is a paid mutator transaction binding the contract method 0x34424855.
//
// Solidity: function storeDeepUintArray(arr uint64[3][4][5]) returns()
func (_DeeplyNestedArray *DeeplyNestedArraySession) StoreDeepUintArray(arr [5][4][3]uint64) (*types.Transaction, error) {
	return _DeeplyNestedArray.Contract.StoreDeepUintArray(&_DeeplyNestedArray.TransactOpts, arr)
}

// StoreDeepUintArray is a paid mutator transaction binding the contract method 0x34424855.
//
// Solidity: function storeDeepUintArray(arr uint64[3][4][5]) returns()
func (_DeeplyNestedArray *DeeplyNestedArrayTransactorSession) StoreDeepUintArray(arr [5][4][3]uint64) (*types.Transaction, error) {
	return _DeeplyNestedArray.Contract.StoreDeepUintArray(&_DeeplyNestedArray.TransactOpts, arr)
}
