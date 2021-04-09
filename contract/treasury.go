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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nsure\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currency\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"SetClaimDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadlineDuration\",\"type\":\"uint256\"}],\"name\":\"SetDeadlineDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositMax\",\"type\":\"uint256\"}],\"name\":\"SetDepositMax\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"SetOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SetSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ePayouts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"eSetOperator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CLAIM_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHEREUM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Nsure\",\"outputs\":[{\"internalType\":\"contractINsure\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"addDivCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_burnUsers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"burnOuts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"divCurrencies\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"divCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDivCurrencyLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"myBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"payouts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setClaimDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setDeadlineDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setDepositMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// CLAIMTYPEHASH is a free data retrieval call binding the contract method 0x6b0509b1.
//
// Solidity: function CLAIM_TYPEHASH() view returns(bytes32)
func (_Contract *ContractCaller) CLAIMTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "CLAIM_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CLAIMTYPEHASH is a free data retrieval call binding the contract method 0x6b0509b1.
//
// Solidity: function CLAIM_TYPEHASH() view returns(bytes32)
func (_Contract *ContractSession) CLAIMTYPEHASH() ([32]byte, error) {
	return _Contract.Contract.CLAIMTYPEHASH(&_Contract.CallOpts)
}

// CLAIMTYPEHASH is a free data retrieval call binding the contract method 0x6b0509b1.
//
// Solidity: function CLAIM_TYPEHASH() view returns(bytes32)
func (_Contract *ContractCallerSession) CLAIMTYPEHASH() ([32]byte, error) {
	return _Contract.Contract.CLAIMTYPEHASH(&_Contract.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Contract *ContractCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Contract *ContractSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Contract.Contract.DOMAINTYPEHASH(&_Contract.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Contract *ContractCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Contract.Contract.DOMAINTYPEHASH(&_Contract.CallOpts)
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() view returns(address)
func (_Contract *ContractCaller) ETHEREUM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ETHEREUM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() view returns(address)
func (_Contract *ContractSession) ETHEREUM() (common.Address, error) {
	return _Contract.Contract.ETHEREUM(&_Contract.CallOpts)
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() view returns(address)
func (_Contract *ContractCallerSession) ETHEREUM() (common.Address, error) {
	return _Contract.Contract.ETHEREUM(&_Contract.CallOpts)
}

// Nsure is a free data retrieval call binding the contract method 0xb7027119.
//
// Solidity: function Nsure() view returns(address)
func (_Contract *ContractCaller) Nsure(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "Nsure")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nsure is a free data retrieval call binding the contract method 0xb7027119.
//
// Solidity: function Nsure() view returns(address)
func (_Contract *ContractSession) Nsure() (common.Address, error) {
	return _Contract.Contract.Nsure(&_Contract.CallOpts)
}

// Nsure is a free data retrieval call binding the contract method 0xb7027119.
//
// Solidity: function Nsure() view returns(address)
func (_Contract *ContractCallerSession) Nsure() (common.Address, error) {
	return _Contract.Contract.Nsure(&_Contract.CallOpts)
}

// WITHDRAWTYPEHASH is a free data retrieval call binding the contract method 0x76c5d758.
//
// Solidity: function WITHDRAW_TYPEHASH() view returns(bytes32)
func (_Contract *ContractCaller) WITHDRAWTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "WITHDRAW_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWTYPEHASH is a free data retrieval call binding the contract method 0x76c5d758.
//
// Solidity: function WITHDRAW_TYPEHASH() view returns(bytes32)
func (_Contract *ContractSession) WITHDRAWTYPEHASH() ([32]byte, error) {
	return _Contract.Contract.WITHDRAWTYPEHASH(&_Contract.CallOpts)
}

// WITHDRAWTYPEHASH is a free data retrieval call binding the contract method 0x76c5d758.
//
// Solidity: function WITHDRAW_TYPEHASH() view returns(bytes32)
func (_Contract *ContractCallerSession) WITHDRAWTYPEHASH() ([32]byte, error) {
	return _Contract.Contract.WITHDRAWTYPEHASH(&_Contract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, account)
}

// ClaimAt is a free data retrieval call binding the contract method 0x58a6a893.
//
// Solidity: function claimAt(address ) view returns(uint256)
func (_Contract *ContractCaller) ClaimAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "claimAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimAt is a free data retrieval call binding the contract method 0x58a6a893.
//
// Solidity: function claimAt(address ) view returns(uint256)
func (_Contract *ContractSession) ClaimAt(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.ClaimAt(&_Contract.CallOpts, arg0)
}

// ClaimAt is a free data retrieval call binding the contract method 0x58a6a893.
//
// Solidity: function claimAt(address ) view returns(uint256)
func (_Contract *ContractCallerSession) ClaimAt(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.ClaimAt(&_Contract.CallOpts, arg0)
}

// ClaimDuration is a free data retrieval call binding the contract method 0xab1a4d94.
//
// Solidity: function claimDuration() view returns(uint256)
func (_Contract *ContractCaller) ClaimDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "claimDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimDuration is a free data retrieval call binding the contract method 0xab1a4d94.
//
// Solidity: function claimDuration() view returns(uint256)
func (_Contract *ContractSession) ClaimDuration() (*big.Int, error) {
	return _Contract.Contract.ClaimDuration(&_Contract.CallOpts)
}

// ClaimDuration is a free data retrieval call binding the contract method 0xab1a4d94.
//
// Solidity: function claimDuration() view returns(uint256)
func (_Contract *ContractCallerSession) ClaimDuration() (*big.Int, error) {
	return _Contract.Contract.ClaimDuration(&_Contract.CallOpts)
}

// DeadlineDuration is a free data retrieval call binding the contract method 0xba71ae9f.
//
// Solidity: function deadlineDuration() view returns(uint256)
func (_Contract *ContractCaller) DeadlineDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "deadlineDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeadlineDuration is a free data retrieval call binding the contract method 0xba71ae9f.
//
// Solidity: function deadlineDuration() view returns(uint256)
func (_Contract *ContractSession) DeadlineDuration() (*big.Int, error) {
	return _Contract.Contract.DeadlineDuration(&_Contract.CallOpts)
}

// DeadlineDuration is a free data retrieval call binding the contract method 0xba71ae9f.
//
// Solidity: function deadlineDuration() view returns(uint256)
func (_Contract *ContractCallerSession) DeadlineDuration() (*big.Int, error) {
	return _Contract.Contract.DeadlineDuration(&_Contract.CallOpts)
}

// DepositMax is a free data retrieval call binding the contract method 0x3ea61cc0.
//
// Solidity: function depositMax() view returns(uint256)
func (_Contract *ContractCaller) DepositMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "depositMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositMax is a free data retrieval call binding the contract method 0x3ea61cc0.
//
// Solidity: function depositMax() view returns(uint256)
func (_Contract *ContractSession) DepositMax() (*big.Int, error) {
	return _Contract.Contract.DepositMax(&_Contract.CallOpts)
}

// DepositMax is a free data retrieval call binding the contract method 0x3ea61cc0.
//
// Solidity: function depositMax() view returns(uint256)
func (_Contract *ContractCallerSession) DepositMax() (*big.Int, error) {
	return _Contract.Contract.DepositMax(&_Contract.CallOpts)
}

// DivCurrencies is a free data retrieval call binding the contract method 0x064849f7.
//
// Solidity: function divCurrencies(uint256 ) view returns(address divCurrency, uint256 limit)
func (_Contract *ContractCaller) DivCurrencies(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DivCurrency common.Address
	Limit       *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "divCurrencies", arg0)

	outstruct := new(struct {
		DivCurrency common.Address
		Limit       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DivCurrency = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Limit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DivCurrencies is a free data retrieval call binding the contract method 0x064849f7.
//
// Solidity: function divCurrencies(uint256 ) view returns(address divCurrency, uint256 limit)
func (_Contract *ContractSession) DivCurrencies(arg0 *big.Int) (struct {
	DivCurrency common.Address
	Limit       *big.Int
}, error) {
	return _Contract.Contract.DivCurrencies(&_Contract.CallOpts, arg0)
}

// DivCurrencies is a free data retrieval call binding the contract method 0x064849f7.
//
// Solidity: function divCurrencies(uint256 ) view returns(address divCurrency, uint256 limit)
func (_Contract *ContractCallerSession) DivCurrencies(arg0 *big.Int) (struct {
	DivCurrency common.Address
	Limit       *big.Int
}, error) {
	return _Contract.Contract.DivCurrencies(&_Contract.CallOpts, arg0)
}

// GetDivCurrencyLength is a free data retrieval call binding the contract method 0xde31aa48.
//
// Solidity: function getDivCurrencyLength() view returns(uint256)
func (_Contract *ContractCaller) GetDivCurrencyLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDivCurrencyLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDivCurrencyLength is a free data retrieval call binding the contract method 0xde31aa48.
//
// Solidity: function getDivCurrencyLength() view returns(uint256)
func (_Contract *ContractSession) GetDivCurrencyLength() (*big.Int, error) {
	return _Contract.Contract.GetDivCurrencyLength(&_Contract.CallOpts)
}

// GetDivCurrencyLength is a free data retrieval call binding the contract method 0xde31aa48.
//
// Solidity: function getDivCurrencyLength() view returns(uint256)
func (_Contract *ContractCallerSession) GetDivCurrencyLength() (*big.Int, error) {
	return _Contract.Contract.GetDivCurrencyLength(&_Contract.CallOpts)
}

// MyBalanceOf is a free data retrieval call binding the contract method 0x055dbc28.
//
// Solidity: function myBalanceOf(address tokenAddress) view returns(uint256)
func (_Contract *ContractCaller) MyBalanceOf(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "myBalanceOf", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyBalanceOf is a free data retrieval call binding the contract method 0x055dbc28.
//
// Solidity: function myBalanceOf(address tokenAddress) view returns(uint256)
func (_Contract *ContractSession) MyBalanceOf(tokenAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.MyBalanceOf(&_Contract.CallOpts, tokenAddress)
}

// MyBalanceOf is a free data retrieval call binding the contract method 0x055dbc28.
//
// Solidity: function myBalanceOf(address tokenAddress) view returns(uint256)
func (_Contract *ContractCallerSession) MyBalanceOf(tokenAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.MyBalanceOf(&_Contract.CallOpts, tokenAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Contract *ContractCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Contract *ContractSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Nonces(&_Contract.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Nonces(&_Contract.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Contract *ContractCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Contract *ContractSession) Operator() (common.Address, error) {
	return _Contract.Contract.Operator(&_Contract.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Contract *ContractCallerSession) Operator() (common.Address, error) {
	return _Contract.Contract.Operator(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Contract *ContractCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Contract *ContractSession) Signer() (common.Address, error) {
	return _Contract.Contract.Signer(&_Contract.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Contract *ContractCallerSession) Signer() (common.Address, error) {
	return _Contract.Contract.Signer(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Contract *ContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Contract *ContractSession) Version() (string, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Contract *ContractCallerSession) Version() (string, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// AddDivCurrency is a paid mutator transaction binding the contract method 0x6a7a7601.
//
// Solidity: function addDivCurrency(address _currency, uint256 _limit) returns()
func (_Contract *ContractTransactor) AddDivCurrency(opts *bind.TransactOpts, _currency common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addDivCurrency", _currency, _limit)
}

// AddDivCurrency is a paid mutator transaction binding the contract method 0x6a7a7601.
//
// Solidity: function addDivCurrency(address _currency, uint256 _limit) returns()
func (_Contract *ContractSession) AddDivCurrency(_currency common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddDivCurrency(&_Contract.TransactOpts, _currency, _limit)
}

// AddDivCurrency is a paid mutator transaction binding the contract method 0x6a7a7601.
//
// Solidity: function addDivCurrency(address _currency, uint256 _limit) returns()
func (_Contract *ContractTransactorSession) AddDivCurrency(_currency common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddDivCurrency(&_Contract.TransactOpts, _currency, _limit)
}

// BurnOuts is a paid mutator transaction binding the contract method 0xa6a25c39.
//
// Solidity: function burnOuts(address[] _burnUsers, uint256[] _amounts) returns()
func (_Contract *ContractTransactor) BurnOuts(opts *bind.TransactOpts, _burnUsers []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "burnOuts", _burnUsers, _amounts)
}

// BurnOuts is a paid mutator transaction binding the contract method 0xa6a25c39.
//
// Solidity: function burnOuts(address[] _burnUsers, uint256[] _amounts) returns()
func (_Contract *ContractSession) BurnOuts(_burnUsers []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BurnOuts(&_Contract.TransactOpts, _burnUsers, _amounts)
}

// BurnOuts is a paid mutator transaction binding the contract method 0xa6a25c39.
//
// Solidity: function burnOuts(address[] _burnUsers, uint256[] _amounts) returns()
func (_Contract *ContractTransactorSession) BurnOuts(_burnUsers []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BurnOuts(&_Contract.TransactOpts, _burnUsers, _amounts)
}

// Claim is a paid mutator transaction binding the contract method 0x07d6b348.
//
// Solidity: function claim(uint256 _amount, uint256 currency, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactor) Claim(opts *bind.TransactOpts, _amount *big.Int, currency *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claim", _amount, currency, deadline, v, r, s)
}

// Claim is a paid mutator transaction binding the contract method 0x07d6b348.
//
// Solidity: function claim(uint256 _amount, uint256 currency, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractSession) Claim(_amount *big.Int, currency *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Claim(&_Contract.TransactOpts, _amount, currency, deadline, v, r, s)
}

// Claim is a paid mutator transaction binding the contract method 0x07d6b348.
//
// Solidity: function claim(uint256 _amount, uint256 currency, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactorSession) Claim(_amount *big.Int, currency *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Claim(&_Contract.TransactOpts, _amount, currency, deadline, v, r, s)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Contract *ContractSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Contract *ContractTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, amount)
}

// Payouts is a paid mutator transaction binding the contract method 0xb89f78ec.
//
// Solidity: function payouts(address _to, uint256 _amount, address token) returns()
func (_Contract *ContractTransactor) Payouts(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "payouts", _to, _amount, token)
}

// Payouts is a paid mutator transaction binding the contract method 0xb89f78ec.
//
// Solidity: function payouts(address _to, uint256 _amount, address token) returns()
func (_Contract *ContractSession) Payouts(_to common.Address, _amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Payouts(&_Contract.TransactOpts, _to, _amount, token)
}

// Payouts is a paid mutator transaction binding the contract method 0xb89f78ec.
//
// Solidity: function payouts(address _to, uint256 _amount, address token) returns()
func (_Contract *ContractTransactorSession) Payouts(_to common.Address, _amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Payouts(&_Contract.TransactOpts, _to, _amount, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetClaimDuration is a paid mutator transaction binding the contract method 0x97c8cd93.
//
// Solidity: function setClaimDuration(uint256 _duration) returns()
func (_Contract *ContractTransactor) SetClaimDuration(opts *bind.TransactOpts, _duration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setClaimDuration", _duration)
}

// SetClaimDuration is a paid mutator transaction binding the contract method 0x97c8cd93.
//
// Solidity: function setClaimDuration(uint256 _duration) returns()
func (_Contract *ContractSession) SetClaimDuration(_duration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetClaimDuration(&_Contract.TransactOpts, _duration)
}

// SetClaimDuration is a paid mutator transaction binding the contract method 0x97c8cd93.
//
// Solidity: function setClaimDuration(uint256 _duration) returns()
func (_Contract *ContractTransactorSession) SetClaimDuration(_duration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetClaimDuration(&_Contract.TransactOpts, _duration)
}

// SetDeadlineDuration is a paid mutator transaction binding the contract method 0x3edba6ee.
//
// Solidity: function setDeadlineDuration(uint256 _duration) returns()
func (_Contract *ContractTransactor) SetDeadlineDuration(opts *bind.TransactOpts, _duration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDeadlineDuration", _duration)
}

// SetDeadlineDuration is a paid mutator transaction binding the contract method 0x3edba6ee.
//
// Solidity: function setDeadlineDuration(uint256 _duration) returns()
func (_Contract *ContractSession) SetDeadlineDuration(_duration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetDeadlineDuration(&_Contract.TransactOpts, _duration)
}

// SetDeadlineDuration is a paid mutator transaction binding the contract method 0x3edba6ee.
//
// Solidity: function setDeadlineDuration(uint256 _duration) returns()
func (_Contract *ContractTransactorSession) SetDeadlineDuration(_duration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetDeadlineDuration(&_Contract.TransactOpts, _duration)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0xcdc35c67.
//
// Solidity: function setDepositMax(uint256 _max) returns()
func (_Contract *ContractTransactor) SetDepositMax(opts *bind.TransactOpts, _max *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDepositMax", _max)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0xcdc35c67.
//
// Solidity: function setDepositMax(uint256 _max) returns()
func (_Contract *ContractSession) SetDepositMax(_max *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetDepositMax(&_Contract.TransactOpts, _max)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0xcdc35c67.
//
// Solidity: function setDepositMax(uint256 _max) returns()
func (_Contract *ContractTransactorSession) SetDepositMax(_max *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetDepositMax(&_Contract.TransactOpts, _max)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Contract *ContractTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Contract *ContractSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOperator(&_Contract.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Contract *ContractTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOperator(&_Contract.TransactOpts, _operator)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_Contract *ContractTransactor) SetSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setSigner", _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_Contract *ContractSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetSigner(&_Contract.TransactOpts, _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_Contract *ContractTransactorSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetSigner(&_Contract.TransactOpts, _signer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30be79be.
//
// Solidity: function withdraw(uint256 _amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", _amount, deadline, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30be79be.
//
// Solidity: function withdraw(uint256 _amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractSession) Withdraw(_amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, _amount, deadline, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30be79be.
//
// Solidity: function withdraw(uint256 _amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactorSession) Withdraw(_amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, _amount, deadline, v, r, s)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Contract contract.
type ContractBurnIterator struct {
	Event *ContractBurn // Event containing the contract specifics and raw log

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
func (it *ContractBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBurn)
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
		it.Event = new(ContractBurn)
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
func (it *ContractBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBurn represents a Burn event raised by the Contract contract.
type ContractBurn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) FilterBurn(opts *bind.FilterOpts, user []common.Address) (*ContractBurnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractBurnIterator{contract: _Contract.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ContractBurn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBurn)
				if err := _Contract.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) ParseBurn(log types.Log) (*ContractBurn, error) {
	event := new(ContractBurn)
	if err := _Contract.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Contract contract.
type ContractClaimIterator struct {
	Event *ContractClaim // Event containing the contract specifics and raw log

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
func (it *ContractClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaim)
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
		it.Event = new(ContractClaim)
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
func (it *ContractClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaim represents a Claim event raised by the Contract contract.
type ContractClaim struct {
	User     common.Address
	Currency *big.Int
	Amount   *big.Int
	Nonce    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x45c072aa05b9853b5a993de7a28bc332ee01404a628cec1a23ce0f659f842ef1.
//
// Solidity: event Claim(address indexed user, uint256 currency, uint256 amount, uint256 nonce)
func (_Contract *ContractFilterer) FilterClaim(opts *bind.FilterOpts, user []common.Address) (*ContractClaimIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractClaimIterator{contract: _Contract.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x45c072aa05b9853b5a993de7a28bc332ee01404a628cec1a23ce0f659f842ef1.
//
// Solidity: event Claim(address indexed user, uint256 currency, uint256 amount, uint256 nonce)
func (_Contract *ContractFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *ContractClaim, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaim)
				if err := _Contract.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x45c072aa05b9853b5a993de7a28bc332ee01404a628cec1a23ce0f659f842ef1.
//
// Solidity: event Claim(address indexed user, uint256 currency, uint256 amount, uint256 nonce)
func (_Contract *ContractFilterer) ParseClaim(log types.Log) (*ContractClaim, error) {
	event := new(ContractClaim)
	if err := _Contract.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Contract contract.
type ContractDepositIterator struct {
	Event *ContractDeposit // Event containing the contract specifics and raw log

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
func (it *ContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeposit)
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
		it.Event = new(ContractDeposit)
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
func (it *ContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeposit represents a Deposit event raised by the Contract contract.
type ContractDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*ContractDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractDepositIterator{contract: _Contract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeposit)
				if err := _Contract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Contract *ContractFilterer) ParseDeposit(log types.Log) (*ContractDeposit, error) {
	event := new(ContractDeposit)
	if err := _Contract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetClaimDurationIterator is returned from FilterSetClaimDuration and is used to iterate over the raw logs and unpacked data for SetClaimDuration events raised by the Contract contract.
type ContractSetClaimDurationIterator struct {
	Event *ContractSetClaimDuration // Event containing the contract specifics and raw log

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
func (it *ContractSetClaimDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetClaimDuration)
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
		it.Event = new(ContractSetClaimDuration)
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
func (it *ContractSetClaimDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetClaimDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetClaimDuration represents a SetClaimDuration event raised by the Contract contract.
type ContractSetClaimDuration struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetClaimDuration is a free log retrieval operation binding the contract event 0x36c99a63aaaf7a98bb0eaaecb5fa7040272a5bb21656e4cda67af3f91c9d74ef.
//
// Solidity: event SetClaimDuration(uint256 duration)
func (_Contract *ContractFilterer) FilterSetClaimDuration(opts *bind.FilterOpts) (*ContractSetClaimDurationIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetClaimDuration")
	if err != nil {
		return nil, err
	}
	return &ContractSetClaimDurationIterator{contract: _Contract.contract, event: "SetClaimDuration", logs: logs, sub: sub}, nil
}

// WatchSetClaimDuration is a free log subscription operation binding the contract event 0x36c99a63aaaf7a98bb0eaaecb5fa7040272a5bb21656e4cda67af3f91c9d74ef.
//
// Solidity: event SetClaimDuration(uint256 duration)
func (_Contract *ContractFilterer) WatchSetClaimDuration(opts *bind.WatchOpts, sink chan<- *ContractSetClaimDuration) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetClaimDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetClaimDuration)
				if err := _Contract.contract.UnpackLog(event, "SetClaimDuration", log); err != nil {
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

// ParseSetClaimDuration is a log parse operation binding the contract event 0x36c99a63aaaf7a98bb0eaaecb5fa7040272a5bb21656e4cda67af3f91c9d74ef.
//
// Solidity: event SetClaimDuration(uint256 duration)
func (_Contract *ContractFilterer) ParseSetClaimDuration(log types.Log) (*ContractSetClaimDuration, error) {
	event := new(ContractSetClaimDuration)
	if err := _Contract.contract.UnpackLog(event, "SetClaimDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetDeadlineDurationIterator is returned from FilterSetDeadlineDuration and is used to iterate over the raw logs and unpacked data for SetDeadlineDuration events raised by the Contract contract.
type ContractSetDeadlineDurationIterator struct {
	Event *ContractSetDeadlineDuration // Event containing the contract specifics and raw log

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
func (it *ContractSetDeadlineDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetDeadlineDuration)
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
		it.Event = new(ContractSetDeadlineDuration)
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
func (it *ContractSetDeadlineDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetDeadlineDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetDeadlineDuration represents a SetDeadlineDuration event raised by the Contract contract.
type ContractSetDeadlineDuration struct {
	DeadlineDuration *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetDeadlineDuration is a free log retrieval operation binding the contract event 0x0f94ac0902239229bbd473ea5f9e71ea7b9314f5047d6e49a8359ce3e578e842.
//
// Solidity: event SetDeadlineDuration(uint256 deadlineDuration)
func (_Contract *ContractFilterer) FilterSetDeadlineDuration(opts *bind.FilterOpts) (*ContractSetDeadlineDurationIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetDeadlineDuration")
	if err != nil {
		return nil, err
	}
	return &ContractSetDeadlineDurationIterator{contract: _Contract.contract, event: "SetDeadlineDuration", logs: logs, sub: sub}, nil
}

// WatchSetDeadlineDuration is a free log subscription operation binding the contract event 0x0f94ac0902239229bbd473ea5f9e71ea7b9314f5047d6e49a8359ce3e578e842.
//
// Solidity: event SetDeadlineDuration(uint256 deadlineDuration)
func (_Contract *ContractFilterer) WatchSetDeadlineDuration(opts *bind.WatchOpts, sink chan<- *ContractSetDeadlineDuration) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetDeadlineDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetDeadlineDuration)
				if err := _Contract.contract.UnpackLog(event, "SetDeadlineDuration", log); err != nil {
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

// ParseSetDeadlineDuration is a log parse operation binding the contract event 0x0f94ac0902239229bbd473ea5f9e71ea7b9314f5047d6e49a8359ce3e578e842.
//
// Solidity: event SetDeadlineDuration(uint256 deadlineDuration)
func (_Contract *ContractFilterer) ParseSetDeadlineDuration(log types.Log) (*ContractSetDeadlineDuration, error) {
	event := new(ContractSetDeadlineDuration)
	if err := _Contract.contract.UnpackLog(event, "SetDeadlineDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetDepositMaxIterator is returned from FilterSetDepositMax and is used to iterate over the raw logs and unpacked data for SetDepositMax events raised by the Contract contract.
type ContractSetDepositMaxIterator struct {
	Event *ContractSetDepositMax // Event containing the contract specifics and raw log

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
func (it *ContractSetDepositMaxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetDepositMax)
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
		it.Event = new(ContractSetDepositMax)
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
func (it *ContractSetDepositMaxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetDepositMaxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetDepositMax represents a SetDepositMax event raised by the Contract contract.
type ContractSetDepositMax struct {
	DepositMax *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetDepositMax is a free log retrieval operation binding the contract event 0x923b217cac2834f6f30f205051ac2478961b05413782731a027539cc5ff82549.
//
// Solidity: event SetDepositMax(uint256 depositMax)
func (_Contract *ContractFilterer) FilterSetDepositMax(opts *bind.FilterOpts) (*ContractSetDepositMaxIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetDepositMax")
	if err != nil {
		return nil, err
	}
	return &ContractSetDepositMaxIterator{contract: _Contract.contract, event: "SetDepositMax", logs: logs, sub: sub}, nil
}

// WatchSetDepositMax is a free log subscription operation binding the contract event 0x923b217cac2834f6f30f205051ac2478961b05413782731a027539cc5ff82549.
//
// Solidity: event SetDepositMax(uint256 depositMax)
func (_Contract *ContractFilterer) WatchSetDepositMax(opts *bind.WatchOpts, sink chan<- *ContractSetDepositMax) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetDepositMax")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetDepositMax)
				if err := _Contract.contract.UnpackLog(event, "SetDepositMax", log); err != nil {
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

// ParseSetDepositMax is a log parse operation binding the contract event 0x923b217cac2834f6f30f205051ac2478961b05413782731a027539cc5ff82549.
//
// Solidity: event SetDepositMax(uint256 depositMax)
func (_Contract *ContractFilterer) ParseSetDepositMax(log types.Log) (*ContractSetDepositMax, error) {
	event := new(ContractSetDepositMax)
	if err := _Contract.contract.UnpackLog(event, "SetDepositMax", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetOperatorIterator is returned from FilterSetOperator and is used to iterate over the raw logs and unpacked data for SetOperator events raised by the Contract contract.
type ContractSetOperatorIterator struct {
	Event *ContractSetOperator // Event containing the contract specifics and raw log

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
func (it *ContractSetOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetOperator)
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
		it.Event = new(ContractSetOperator)
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
func (it *ContractSetOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetOperator represents a SetOperator event raised by the Contract contract.
type ContractSetOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOperator is a free log retrieval operation binding the contract event 0xdbebfba65bd6398fb722063efc10c99f624f9cd8ba657201056af918a676d5ee.
//
// Solidity: event SetOperator(address indexed operator)
func (_Contract *ContractFilterer) FilterSetOperator(opts *bind.FilterOpts, operator []common.Address) (*ContractSetOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetOperatorIterator{contract: _Contract.contract, event: "SetOperator", logs: logs, sub: sub}, nil
}

// WatchSetOperator is a free log subscription operation binding the contract event 0xdbebfba65bd6398fb722063efc10c99f624f9cd8ba657201056af918a676d5ee.
//
// Solidity: event SetOperator(address indexed operator)
func (_Contract *ContractFilterer) WatchSetOperator(opts *bind.WatchOpts, sink chan<- *ContractSetOperator, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetOperator)
				if err := _Contract.contract.UnpackLog(event, "SetOperator", log); err != nil {
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

// ParseSetOperator is a log parse operation binding the contract event 0xdbebfba65bd6398fb722063efc10c99f624f9cd8ba657201056af918a676d5ee.
//
// Solidity: event SetOperator(address indexed operator)
func (_Contract *ContractFilterer) ParseSetOperator(log types.Log) (*ContractSetOperator, error) {
	event := new(ContractSetOperator)
	if err := _Contract.contract.UnpackLog(event, "SetOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetSignerIterator is returned from FilterSetSigner and is used to iterate over the raw logs and unpacked data for SetSigner events raised by the Contract contract.
type ContractSetSignerIterator struct {
	Event *ContractSetSigner // Event containing the contract specifics and raw log

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
func (it *ContractSetSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetSigner)
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
		it.Event = new(ContractSetSigner)
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
func (it *ContractSetSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetSigner represents a SetSigner event raised by the Contract contract.
type ContractSetSigner struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSigner is a free log retrieval operation binding the contract event 0xbb10aee7ef5a307b8097c6a7f2892b909ff1736fd24a6a5260640c185f7153b6.
//
// Solidity: event SetSigner(address indexed signer)
func (_Contract *ContractFilterer) FilterSetSigner(opts *bind.FilterOpts, signer []common.Address) (*ContractSetSignerIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetSigner", signerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetSignerIterator{contract: _Contract.contract, event: "SetSigner", logs: logs, sub: sub}, nil
}

// WatchSetSigner is a free log subscription operation binding the contract event 0xbb10aee7ef5a307b8097c6a7f2892b909ff1736fd24a6a5260640c185f7153b6.
//
// Solidity: event SetSigner(address indexed signer)
func (_Contract *ContractFilterer) WatchSetSigner(opts *bind.WatchOpts, sink chan<- *ContractSetSigner, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetSigner", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetSigner)
				if err := _Contract.contract.UnpackLog(event, "SetSigner", log); err != nil {
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

// ParseSetSigner is a log parse operation binding the contract event 0xbb10aee7ef5a307b8097c6a7f2892b909ff1736fd24a6a5260640c185f7153b6.
//
// Solidity: event SetSigner(address indexed signer)
func (_Contract *ContractFilterer) ParseSetSigner(log types.Log) (*ContractSetSigner, error) {
	event := new(ContractSetSigner)
	if err := _Contract.contract.UnpackLog(event, "SetSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Contract contract.
type ContractWithdrawIterator struct {
	Event *ContractWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdraw)
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
		it.Event = new(ContractWithdraw)
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
func (it *ContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdraw represents a Withdraw event raised by the Contract contract.
type ContractWithdraw struct {
	User   common.Address
	Amount *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, uint256 nonce)
func (_Contract *ContractFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*ContractWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawIterator{contract: _Contract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, uint256 nonce)
func (_Contract *ContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ContractWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdraw)
				if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, uint256 nonce)
func (_Contract *ContractFilterer) ParseWithdraw(log types.Log) (*ContractWithdraw, error) {
	event := new(ContractWithdraw)
	if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEPayoutsIterator is returned from FilterEPayouts and is used to iterate over the raw logs and unpacked data for EPayouts events raised by the Contract contract.
type ContractEPayoutsIterator struct {
	Event *ContractEPayouts // Event containing the contract specifics and raw log

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
func (it *ContractEPayoutsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEPayouts)
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
		it.Event = new(ContractEPayouts)
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
func (it *ContractEPayoutsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEPayoutsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEPayouts represents a EPayouts event raised by the Contract contract.
type ContractEPayouts struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEPayouts is a free log retrieval operation binding the contract event 0x21f5d51767dbc29c660e79e98830844af5399b7623bbdfbe6ce5d6cb216bc2fe.
//
// Solidity: event ePayouts(address indexed to, uint256 amount)
func (_Contract *ContractFilterer) FilterEPayouts(opts *bind.FilterOpts, to []common.Address) (*ContractEPayoutsIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ePayouts", toRule)
	if err != nil {
		return nil, err
	}
	return &ContractEPayoutsIterator{contract: _Contract.contract, event: "ePayouts", logs: logs, sub: sub}, nil
}

// WatchEPayouts is a free log subscription operation binding the contract event 0x21f5d51767dbc29c660e79e98830844af5399b7623bbdfbe6ce5d6cb216bc2fe.
//
// Solidity: event ePayouts(address indexed to, uint256 amount)
func (_Contract *ContractFilterer) WatchEPayouts(opts *bind.WatchOpts, sink chan<- *ContractEPayouts, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ePayouts", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEPayouts)
				if err := _Contract.contract.UnpackLog(event, "ePayouts", log); err != nil {
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

// ParseEPayouts is a log parse operation binding the contract event 0x21f5d51767dbc29c660e79e98830844af5399b7623bbdfbe6ce5d6cb216bc2fe.
//
// Solidity: event ePayouts(address indexed to, uint256 amount)
func (_Contract *ContractFilterer) ParseEPayouts(log types.Log) (*ContractEPayouts, error) {
	event := new(ContractEPayouts)
	if err := _Contract.contract.UnpackLog(event, "ePayouts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractESetOperatorIterator is returned from FilterESetOperator and is used to iterate over the raw logs and unpacked data for ESetOperator events raised by the Contract contract.
type ContractESetOperatorIterator struct {
	Event *ContractESetOperator // Event containing the contract specifics and raw log

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
func (it *ContractESetOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractESetOperator)
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
		it.Event = new(ContractESetOperator)
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
func (it *ContractESetOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractESetOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractESetOperator represents a ESetOperator event raised by the Contract contract.
type ContractESetOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterESetOperator is a free log retrieval operation binding the contract event 0x22b76827cedebd97fbe4b16a026ed043bde58720fc1f4f0802457397016ceb1c.
//
// Solidity: event eSetOperator(address indexed operator)
func (_Contract *ContractFilterer) FilterESetOperator(opts *bind.FilterOpts, operator []common.Address) (*ContractESetOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "eSetOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractESetOperatorIterator{contract: _Contract.contract, event: "eSetOperator", logs: logs, sub: sub}, nil
}

// WatchESetOperator is a free log subscription operation binding the contract event 0x22b76827cedebd97fbe4b16a026ed043bde58720fc1f4f0802457397016ceb1c.
//
// Solidity: event eSetOperator(address indexed operator)
func (_Contract *ContractFilterer) WatchESetOperator(opts *bind.WatchOpts, sink chan<- *ContractESetOperator, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "eSetOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractESetOperator)
				if err := _Contract.contract.UnpackLog(event, "eSetOperator", log); err != nil {
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

// ParseESetOperator is a log parse operation binding the contract event 0x22b76827cedebd97fbe4b16a026ed043bde58720fc1f4f0802457397016ceb1c.
//
// Solidity: event eSetOperator(address indexed operator)
func (_Contract *ContractFilterer) ParseESetOperator(log types.Log) (*ContractESetOperator, error) {
	event := new(ContractESetOperator)
	if err := _Contract.contract.UnpackLog(event, "eSetOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
