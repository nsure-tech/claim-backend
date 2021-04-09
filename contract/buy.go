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

// BuyOrder is an auto generated low-level Go binding around an user-defined struct.
type BuyOrder struct {
	Buyer     common.Address
	ProductId *big.Int
	Currency  *big.Int
	Premium   *big.Int
	Amount    *big.Int
	Period    *big.Int
	CreateAt  *big.Int
	State     uint8
	Nonce     *big.Int
	OrderId   *big.Int
}

// BuyABI is the input ABI used to generate the binding from.
const BuyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_underWriting\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_surplus\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_product\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"AddDivCurrency\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"DeleteDivCurrency\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"productId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createAt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structBuy.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"NewOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underWritingRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"surplusRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"treasuryRate\",\"type\":\"uint256\"}],\"name\":\"SetRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SetSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"surplus\",\"type\":\"address\"}],\"name\":\"SetSurplus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"}],\"name\":\"SetTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underWriting\",\"type\":\"address\"}],\"name\":\"SetUnderWriting\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BUY_INSURANCE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"addDivCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_productId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currency\",\"type\":\"uint256\"}],\"name\":\"buyInsurance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"delDivCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"divCurrencies\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDivCurrencyLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"insuranceOrders\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"productId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createAt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"product\",\"outputs\":[{\"internalType\":\"contractIProduct\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_underWritingRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_surplusRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_treasuryRate\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_surplusAddr\",\"type\":\"address\"}],\"name\":\"setSurplusAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddr\",\"type\":\"address\"}],\"name\":\"setTreasuryAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_underWritingAddr\",\"type\":\"address\"}],\"name\":\"setUnderWritingAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"surplus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"surplusRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underWriting\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underWritingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Buy is an auto generated Go binding around an Ethereum contract.
type Buy struct {
	BuyCaller     // Read-only binding to the contract
	BuyTransactor // Write-only binding to the contract
	BuyFilterer   // Log filterer for contract events
}

// BuyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuySession struct {
	Contract     *Buy              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyCallerSession struct {
	Contract *BuyCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BuyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyTransactorSession struct {
	Contract     *BuyTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyRaw struct {
	Contract *Buy // Generic contract binding to access the raw methods on
}

// BuyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyCallerRaw struct {
	Contract *BuyCaller // Generic read-only contract binding to access the raw methods on
}

// BuyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyTransactorRaw struct {
	Contract *BuyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuy creates a new instance of Buy, bound to a specific deployed contract.
func NewBuy(address common.Address, backend bind.ContractBackend) (*Buy, error) {
	contract, err := bindBuy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Buy{BuyCaller: BuyCaller{contract: contract}, BuyTransactor: BuyTransactor{contract: contract}, BuyFilterer: BuyFilterer{contract: contract}}, nil
}

// NewBuyCaller creates a new read-only instance of Buy, bound to a specific deployed contract.
func NewBuyCaller(address common.Address, caller bind.ContractCaller) (*BuyCaller, error) {
	contract, err := bindBuy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuyCaller{contract: contract}, nil
}

// NewBuyTransactor creates a new write-only instance of Buy, bound to a specific deployed contract.
func NewBuyTransactor(address common.Address, transactor bind.ContractTransactor) (*BuyTransactor, error) {
	contract, err := bindBuy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuyTransactor{contract: contract}, nil
}

// NewBuyFilterer creates a new log filterer instance of Buy, bound to a specific deployed contract.
func NewBuyFilterer(address common.Address, filterer bind.ContractFilterer) (*BuyFilterer, error) {
	contract, err := bindBuy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuyFilterer{contract: contract}, nil
}

// bindBuy binds a generic wrapper to an already deployed contract.
func bindBuy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BuyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buy *BuyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Buy.Contract.BuyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buy *BuyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.Contract.BuyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buy *BuyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buy.Contract.BuyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buy *BuyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Buy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buy *BuyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buy *BuyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buy.Contract.contract.Transact(opts, method, params...)
}

// BUYINSURANCETYPEHASH is a free data retrieval call binding the contract method 0x56345441.
//
// Solidity: function BUY_INSURANCE_TYPEHASH() view returns(bytes32)
func (_Buy *BuyCaller) BUYINSURANCETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "BUY_INSURANCE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BUYINSURANCETYPEHASH is a free data retrieval call binding the contract method 0x56345441.
//
// Solidity: function BUY_INSURANCE_TYPEHASH() view returns(bytes32)
func (_Buy *BuySession) BUYINSURANCETYPEHASH() ([32]byte, error) {
	return _Buy.Contract.BUYINSURANCETYPEHASH(&_Buy.CallOpts)
}

// BUYINSURANCETYPEHASH is a free data retrieval call binding the contract method 0x56345441.
//
// Solidity: function BUY_INSURANCE_TYPEHASH() view returns(bytes32)
func (_Buy *BuyCallerSession) BUYINSURANCETYPEHASH() ([32]byte, error) {
	return _Buy.Contract.BUYINSURANCETYPEHASH(&_Buy.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Buy *BuyCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Buy *BuySession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Buy.Contract.DOMAINTYPEHASH(&_Buy.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Buy *BuyCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Buy.Contract.DOMAINTYPEHASH(&_Buy.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Buy *BuyCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Buy *BuySession) WETH() (common.Address, error) {
	return _Buy.Contract.WETH(&_Buy.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Buy *BuyCallerSession) WETH() (common.Address, error) {
	return _Buy.Contract.WETH(&_Buy.CallOpts)
}

// DivCurrencies is a free data retrieval call binding the contract method 0x064849f7.
//
// Solidity: function divCurrencies(uint256 ) view returns(address)
func (_Buy *BuyCaller) DivCurrencies(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "divCurrencies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DivCurrencies is a free data retrieval call binding the contract method 0x064849f7.
//
// Solidity: function divCurrencies(uint256 ) view returns(address)
func (_Buy *BuySession) DivCurrencies(arg0 *big.Int) (common.Address, error) {
	return _Buy.Contract.DivCurrencies(&_Buy.CallOpts, arg0)
}

// DivCurrencies is a free data retrieval call binding the contract method 0x064849f7.
//
// Solidity: function divCurrencies(uint256 ) view returns(address)
func (_Buy *BuyCallerSession) DivCurrencies(arg0 *big.Int) (common.Address, error) {
	return _Buy.Contract.DivCurrencies(&_Buy.CallOpts, arg0)
}

// GetDivCurrencyLength is a free data retrieval call binding the contract method 0xde31aa48.
//
// Solidity: function getDivCurrencyLength() view returns(uint256)
func (_Buy *BuyCaller) GetDivCurrencyLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "getDivCurrencyLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDivCurrencyLength is a free data retrieval call binding the contract method 0xde31aa48.
//
// Solidity: function getDivCurrencyLength() view returns(uint256)
func (_Buy *BuySession) GetDivCurrencyLength() (*big.Int, error) {
	return _Buy.Contract.GetDivCurrencyLength(&_Buy.CallOpts)
}

// GetDivCurrencyLength is a free data retrieval call binding the contract method 0xde31aa48.
//
// Solidity: function getDivCurrencyLength() view returns(uint256)
func (_Buy *BuyCallerSession) GetDivCurrencyLength() (*big.Int, error) {
	return _Buy.Contract.GetDivCurrencyLength(&_Buy.CallOpts)
}

// InsuranceOrders is a free data retrieval call binding the contract method 0xd30b8915.
//
// Solidity: function insuranceOrders(uint256 ) view returns(address buyer, uint256 productId, uint256 currency, uint256 premium, uint256 amount, uint256 period, uint256 createAt, uint8 state, uint256 nonce, uint256 orderId)
func (_Buy *BuyCaller) InsuranceOrders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Buyer     common.Address
	ProductId *big.Int
	Currency  *big.Int
	Premium   *big.Int
	Amount    *big.Int
	Period    *big.Int
	CreateAt  *big.Int
	State     uint8
	Nonce     *big.Int
	OrderId   *big.Int
}, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "insuranceOrders", arg0)

	outstruct := new(struct {
		Buyer     common.Address
		ProductId *big.Int
		Currency  *big.Int
		Premium   *big.Int
		Amount    *big.Int
		Period    *big.Int
		CreateAt  *big.Int
		State     uint8
		Nonce     *big.Int
		OrderId   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ProductId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Currency = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Premium = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Period = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.CreateAt = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.Nonce = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.OrderId = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// InsuranceOrders is a free data retrieval call binding the contract method 0xd30b8915.
//
// Solidity: function insuranceOrders(uint256 ) view returns(address buyer, uint256 productId, uint256 currency, uint256 premium, uint256 amount, uint256 period, uint256 createAt, uint8 state, uint256 nonce, uint256 orderId)
func (_Buy *BuySession) InsuranceOrders(arg0 *big.Int) (struct {
	Buyer     common.Address
	ProductId *big.Int
	Currency  *big.Int
	Premium   *big.Int
	Amount    *big.Int
	Period    *big.Int
	CreateAt  *big.Int
	State     uint8
	Nonce     *big.Int
	OrderId   *big.Int
}, error) {
	return _Buy.Contract.InsuranceOrders(&_Buy.CallOpts, arg0)
}

// InsuranceOrders is a free data retrieval call binding the contract method 0xd30b8915.
//
// Solidity: function insuranceOrders(uint256 ) view returns(address buyer, uint256 productId, uint256 currency, uint256 premium, uint256 amount, uint256 period, uint256 createAt, uint8 state, uint256 nonce, uint256 orderId)
func (_Buy *BuyCallerSession) InsuranceOrders(arg0 *big.Int) (struct {
	Buyer     common.Address
	ProductId *big.Int
	Currency  *big.Int
	Premium   *big.Int
	Amount    *big.Int
	Period    *big.Int
	CreateAt  *big.Int
	State     uint8
	Nonce     *big.Int
	OrderId   *big.Int
}, error) {
	return _Buy.Contract.InsuranceOrders(&_Buy.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Buy *BuyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Buy *BuySession) Name() (string, error) {
	return _Buy.Contract.Name(&_Buy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Buy *BuyCallerSession) Name() (string, error) {
	return _Buy.Contract.Name(&_Buy.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Buy *BuyCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Buy *BuySession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Buy.Contract.Nonces(&_Buy.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Buy *BuyCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Buy.Contract.Nonces(&_Buy.CallOpts, arg0)
}

// OrderIndex is a free data retrieval call binding the contract method 0x30471fee.
//
// Solidity: function orderIndex() view returns(uint256)
func (_Buy *BuyCaller) OrderIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "orderIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderIndex is a free data retrieval call binding the contract method 0x30471fee.
//
// Solidity: function orderIndex() view returns(uint256)
func (_Buy *BuySession) OrderIndex() (*big.Int, error) {
	return _Buy.Contract.OrderIndex(&_Buy.CallOpts)
}

// OrderIndex is a free data retrieval call binding the contract method 0x30471fee.
//
// Solidity: function orderIndex() view returns(uint256)
func (_Buy *BuyCallerSession) OrderIndex() (*big.Int, error) {
	return _Buy.Contract.OrderIndex(&_Buy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Buy *BuyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Buy *BuySession) Owner() (common.Address, error) {
	return _Buy.Contract.Owner(&_Buy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Buy *BuyCallerSession) Owner() (common.Address, error) {
	return _Buy.Contract.Owner(&_Buy.CallOpts)
}

// Product is a free data retrieval call binding the contract method 0xbf9ce952.
//
// Solidity: function product() view returns(address)
func (_Buy *BuyCaller) Product(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "product")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Product is a free data retrieval call binding the contract method 0xbf9ce952.
//
// Solidity: function product() view returns(address)
func (_Buy *BuySession) Product() (common.Address, error) {
	return _Buy.Contract.Product(&_Buy.CallOpts)
}

// Product is a free data retrieval call binding the contract method 0xbf9ce952.
//
// Solidity: function product() view returns(address)
func (_Buy *BuyCallerSession) Product() (common.Address, error) {
	return _Buy.Contract.Product(&_Buy.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Buy *BuyCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Buy *BuySession) Signer() (common.Address, error) {
	return _Buy.Contract.Signer(&_Buy.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Buy *BuyCallerSession) Signer() (common.Address, error) {
	return _Buy.Contract.Signer(&_Buy.CallOpts)
}

// Surplus is a free data retrieval call binding the contract method 0x13888565.
//
// Solidity: function surplus() view returns(address)
func (_Buy *BuyCaller) Surplus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "surplus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Surplus is a free data retrieval call binding the contract method 0x13888565.
//
// Solidity: function surplus() view returns(address)
func (_Buy *BuySession) Surplus() (common.Address, error) {
	return _Buy.Contract.Surplus(&_Buy.CallOpts)
}

// Surplus is a free data retrieval call binding the contract method 0x13888565.
//
// Solidity: function surplus() view returns(address)
func (_Buy *BuyCallerSession) Surplus() (common.Address, error) {
	return _Buy.Contract.Surplus(&_Buy.CallOpts)
}

// SurplusRate is a free data retrieval call binding the contract method 0x49102af9.
//
// Solidity: function surplusRate() view returns(uint256)
func (_Buy *BuyCaller) SurplusRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "surplusRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SurplusRate is a free data retrieval call binding the contract method 0x49102af9.
//
// Solidity: function surplusRate() view returns(uint256)
func (_Buy *BuySession) SurplusRate() (*big.Int, error) {
	return _Buy.Contract.SurplusRate(&_Buy.CallOpts)
}

// SurplusRate is a free data retrieval call binding the contract method 0x49102af9.
//
// Solidity: function surplusRate() view returns(uint256)
func (_Buy *BuyCallerSession) SurplusRate() (*big.Int, error) {
	return _Buy.Contract.SurplusRate(&_Buy.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Buy *BuyCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Buy *BuySession) Treasury() (common.Address, error) {
	return _Buy.Contract.Treasury(&_Buy.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Buy *BuyCallerSession) Treasury() (common.Address, error) {
	return _Buy.Contract.Treasury(&_Buy.CallOpts)
}

// TreasuryRate is a free data retrieval call binding the contract method 0xe4b72516.
//
// Solidity: function treasuryRate() view returns(uint256)
func (_Buy *BuyCaller) TreasuryRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "treasuryRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryRate is a free data retrieval call binding the contract method 0xe4b72516.
//
// Solidity: function treasuryRate() view returns(uint256)
func (_Buy *BuySession) TreasuryRate() (*big.Int, error) {
	return _Buy.Contract.TreasuryRate(&_Buy.CallOpts)
}

// TreasuryRate is a free data retrieval call binding the contract method 0xe4b72516.
//
// Solidity: function treasuryRate() view returns(uint256)
func (_Buy *BuyCallerSession) TreasuryRate() (*big.Int, error) {
	return _Buy.Contract.TreasuryRate(&_Buy.CallOpts)
}

// UnderWriting is a free data retrieval call binding the contract method 0x91e5795d.
//
// Solidity: function underWriting() view returns(address)
func (_Buy *BuyCaller) UnderWriting(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "underWriting")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderWriting is a free data retrieval call binding the contract method 0x91e5795d.
//
// Solidity: function underWriting() view returns(address)
func (_Buy *BuySession) UnderWriting() (common.Address, error) {
	return _Buy.Contract.UnderWriting(&_Buy.CallOpts)
}

// UnderWriting is a free data retrieval call binding the contract method 0x91e5795d.
//
// Solidity: function underWriting() view returns(address)
func (_Buy *BuyCallerSession) UnderWriting() (common.Address, error) {
	return _Buy.Contract.UnderWriting(&_Buy.CallOpts)
}

// UnderWritingRate is a free data retrieval call binding the contract method 0xec499350.
//
// Solidity: function underWritingRate() view returns(uint256)
func (_Buy *BuyCaller) UnderWritingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "underWritingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderWritingRate is a free data retrieval call binding the contract method 0xec499350.
//
// Solidity: function underWritingRate() view returns(uint256)
func (_Buy *BuySession) UnderWritingRate() (*big.Int, error) {
	return _Buy.Contract.UnderWritingRate(&_Buy.CallOpts)
}

// UnderWritingRate is a free data retrieval call binding the contract method 0xec499350.
//
// Solidity: function underWritingRate() view returns(uint256)
func (_Buy *BuyCallerSession) UnderWritingRate() (*big.Int, error) {
	return _Buy.Contract.UnderWritingRate(&_Buy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Buy *BuyCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Buy *BuySession) Version() (string, error) {
	return _Buy.Contract.Version(&_Buy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Buy *BuyCallerSession) Version() (string, error) {
	return _Buy.Contract.Version(&_Buy.CallOpts)
}

// AddDivCurrency is a paid mutator transaction binding the contract method 0xb48015c0.
//
// Solidity: function addDivCurrency(address currency) returns()
func (_Buy *BuyTransactor) AddDivCurrency(opts *bind.TransactOpts, currency common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "addDivCurrency", currency)
}

// AddDivCurrency is a paid mutator transaction binding the contract method 0xb48015c0.
//
// Solidity: function addDivCurrency(address currency) returns()
func (_Buy *BuySession) AddDivCurrency(currency common.Address) (*types.Transaction, error) {
	return _Buy.Contract.AddDivCurrency(&_Buy.TransactOpts, currency)
}

// AddDivCurrency is a paid mutator transaction binding the contract method 0xb48015c0.
//
// Solidity: function addDivCurrency(address currency) returns()
func (_Buy *BuyTransactorSession) AddDivCurrency(currency common.Address) (*types.Transaction, error) {
	return _Buy.Contract.AddDivCurrency(&_Buy.TransactOpts, currency)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x2fb9cb78.
//
// Solidity: function buyInsurance(uint256 _productId, uint256 _amount, uint256 _cost, uint256 period, uint8 v, bytes32 r, bytes32 s, uint256 deadline, uint256 currency) payable returns()
func (_Buy *BuyTransactor) BuyInsurance(opts *bind.TransactOpts, _productId *big.Int, _amount *big.Int, _cost *big.Int, period *big.Int, v uint8, r [32]byte, s [32]byte, deadline *big.Int, currency *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "buyInsurance", _productId, _amount, _cost, period, v, r, s, deadline, currency)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x2fb9cb78.
//
// Solidity: function buyInsurance(uint256 _productId, uint256 _amount, uint256 _cost, uint256 period, uint8 v, bytes32 r, bytes32 s, uint256 deadline, uint256 currency) payable returns()
func (_Buy *BuySession) BuyInsurance(_productId *big.Int, _amount *big.Int, _cost *big.Int, period *big.Int, v uint8, r [32]byte, s [32]byte, deadline *big.Int, currency *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.BuyInsurance(&_Buy.TransactOpts, _productId, _amount, _cost, period, v, r, s, deadline, currency)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x2fb9cb78.
//
// Solidity: function buyInsurance(uint256 _productId, uint256 _amount, uint256 _cost, uint256 period, uint8 v, bytes32 r, bytes32 s, uint256 deadline, uint256 currency) payable returns()
func (_Buy *BuyTransactorSession) BuyInsurance(_productId *big.Int, _amount *big.Int, _cost *big.Int, period *big.Int, v uint8, r [32]byte, s [32]byte, deadline *big.Int, currency *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.BuyInsurance(&_Buy.TransactOpts, _productId, _amount, _cost, period, v, r, s, deadline, currency)
}

// DelDivCurrency is a paid mutator transaction binding the contract method 0xbc30ae4f.
//
// Solidity: function delDivCurrency(address currency) returns()
func (_Buy *BuyTransactor) DelDivCurrency(opts *bind.TransactOpts, currency common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "delDivCurrency", currency)
}

// DelDivCurrency is a paid mutator transaction binding the contract method 0xbc30ae4f.
//
// Solidity: function delDivCurrency(address currency) returns()
func (_Buy *BuySession) DelDivCurrency(currency common.Address) (*types.Transaction, error) {
	return _Buy.Contract.DelDivCurrency(&_Buy.TransactOpts, currency)
}

// DelDivCurrency is a paid mutator transaction binding the contract method 0xbc30ae4f.
//
// Solidity: function delDivCurrency(address currency) returns()
func (_Buy *BuyTransactorSession) DelDivCurrency(currency common.Address) (*types.Transaction, error) {
	return _Buy.Contract.DelDivCurrency(&_Buy.TransactOpts, currency)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Buy *BuyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Buy *BuySession) RenounceOwnership() (*types.Transaction, error) {
	return _Buy.Contract.RenounceOwnership(&_Buy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Buy *BuyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Buy.Contract.RenounceOwnership(&_Buy.TransactOpts)
}

// SetRate is a paid mutator transaction binding the contract method 0x3989c666.
//
// Solidity: function setRate(uint256 _underWritingRate, uint256 _surplusRate, uint256 _treasuryRate) returns()
func (_Buy *BuyTransactor) SetRate(opts *bind.TransactOpts, _underWritingRate *big.Int, _surplusRate *big.Int, _treasuryRate *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setRate", _underWritingRate, _surplusRate, _treasuryRate)
}

// SetRate is a paid mutator transaction binding the contract method 0x3989c666.
//
// Solidity: function setRate(uint256 _underWritingRate, uint256 _surplusRate, uint256 _treasuryRate) returns()
func (_Buy *BuySession) SetRate(_underWritingRate *big.Int, _surplusRate *big.Int, _treasuryRate *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.SetRate(&_Buy.TransactOpts, _underWritingRate, _surplusRate, _treasuryRate)
}

// SetRate is a paid mutator transaction binding the contract method 0x3989c666.
//
// Solidity: function setRate(uint256 _underWritingRate, uint256 _surplusRate, uint256 _treasuryRate) returns()
func (_Buy *BuyTransactorSession) SetRate(_underWritingRate *big.Int, _surplusRate *big.Int, _treasuryRate *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.SetRate(&_Buy.TransactOpts, _underWritingRate, _surplusRate, _treasuryRate)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_Buy *BuyTransactor) SetSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setSigner", _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_Buy *BuySession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetSigner(&_Buy.TransactOpts, _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_Buy *BuyTransactorSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetSigner(&_Buy.TransactOpts, _signer)
}

// SetSurplusAddr is a paid mutator transaction binding the contract method 0x43f3fdc1.
//
// Solidity: function setSurplusAddr(address _surplusAddr) returns()
func (_Buy *BuyTransactor) SetSurplusAddr(opts *bind.TransactOpts, _surplusAddr common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setSurplusAddr", _surplusAddr)
}

// SetSurplusAddr is a paid mutator transaction binding the contract method 0x43f3fdc1.
//
// Solidity: function setSurplusAddr(address _surplusAddr) returns()
func (_Buy *BuySession) SetSurplusAddr(_surplusAddr common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetSurplusAddr(&_Buy.TransactOpts, _surplusAddr)
}

// SetSurplusAddr is a paid mutator transaction binding the contract method 0x43f3fdc1.
//
// Solidity: function setSurplusAddr(address _surplusAddr) returns()
func (_Buy *BuyTransactorSession) SetSurplusAddr(_surplusAddr common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetSurplusAddr(&_Buy.TransactOpts, _surplusAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Buy *BuyTransactor) SetTreasuryAddr(opts *bind.TransactOpts, _treasuryAddr common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setTreasuryAddr", _treasuryAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Buy *BuySession) SetTreasuryAddr(_treasuryAddr common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetTreasuryAddr(&_Buy.TransactOpts, _treasuryAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Buy *BuyTransactorSession) SetTreasuryAddr(_treasuryAddr common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetTreasuryAddr(&_Buy.TransactOpts, _treasuryAddr)
}

// SetUnderWritingAddr is a paid mutator transaction binding the contract method 0x8f6bb4b7.
//
// Solidity: function setUnderWritingAddr(address _underWritingAddr) returns()
func (_Buy *BuyTransactor) SetUnderWritingAddr(opts *bind.TransactOpts, _underWritingAddr common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setUnderWritingAddr", _underWritingAddr)
}

// SetUnderWritingAddr is a paid mutator transaction binding the contract method 0x8f6bb4b7.
//
// Solidity: function setUnderWritingAddr(address _underWritingAddr) returns()
func (_Buy *BuySession) SetUnderWritingAddr(_underWritingAddr common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetUnderWritingAddr(&_Buy.TransactOpts, _underWritingAddr)
}

// SetUnderWritingAddr is a paid mutator transaction binding the contract method 0x8f6bb4b7.
//
// Solidity: function setUnderWritingAddr(address _underWritingAddr) returns()
func (_Buy *BuyTransactorSession) SetUnderWritingAddr(_underWritingAddr common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetUnderWritingAddr(&_Buy.TransactOpts, _underWritingAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Buy *BuyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Buy *BuySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Buy.Contract.TransferOwnership(&_Buy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Buy *BuyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Buy.Contract.TransferOwnership(&_Buy.TransactOpts, newOwner)
}

// BuyAddDivCurrencyIterator is returned from FilterAddDivCurrency and is used to iterate over the raw logs and unpacked data for AddDivCurrency events raised by the Buy contract.
type BuyAddDivCurrencyIterator struct {
	Event *BuyAddDivCurrency // Event containing the contract specifics and raw log

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
func (it *BuyAddDivCurrencyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyAddDivCurrency)
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
		it.Event = new(BuyAddDivCurrency)
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
func (it *BuyAddDivCurrencyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyAddDivCurrencyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyAddDivCurrency represents a AddDivCurrency event raised by the Buy contract.
type BuyAddDivCurrency struct {
	Currency common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddDivCurrency is a free log retrieval operation binding the contract event 0x6cbce25c0c275b55de54cf93bc357dd0ef22cfeedba7a0d73b4ed3ed6bb8b35e.
//
// Solidity: event AddDivCurrency(address indexed currency)
func (_Buy *BuyFilterer) FilterAddDivCurrency(opts *bind.FilterOpts, currency []common.Address) (*BuyAddDivCurrencyIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "AddDivCurrency", currencyRule)
	if err != nil {
		return nil, err
	}
	return &BuyAddDivCurrencyIterator{contract: _Buy.contract, event: "AddDivCurrency", logs: logs, sub: sub}, nil
}

// WatchAddDivCurrency is a free log subscription operation binding the contract event 0x6cbce25c0c275b55de54cf93bc357dd0ef22cfeedba7a0d73b4ed3ed6bb8b35e.
//
// Solidity: event AddDivCurrency(address indexed currency)
func (_Buy *BuyFilterer) WatchAddDivCurrency(opts *bind.WatchOpts, sink chan<- *BuyAddDivCurrency, currency []common.Address) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "AddDivCurrency", currencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyAddDivCurrency)
				if err := _Buy.contract.UnpackLog(event, "AddDivCurrency", log); err != nil {
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

// ParseAddDivCurrency is a log parse operation binding the contract event 0x6cbce25c0c275b55de54cf93bc357dd0ef22cfeedba7a0d73b4ed3ed6bb8b35e.
//
// Solidity: event AddDivCurrency(address indexed currency)
func (_Buy *BuyFilterer) ParseAddDivCurrency(log types.Log) (*BuyAddDivCurrency, error) {
	event := new(BuyAddDivCurrency)
	if err := _Buy.contract.UnpackLog(event, "AddDivCurrency", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyDeleteDivCurrencyIterator is returned from FilterDeleteDivCurrency and is used to iterate over the raw logs and unpacked data for DeleteDivCurrency events raised by the Buy contract.
type BuyDeleteDivCurrencyIterator struct {
	Event *BuyDeleteDivCurrency // Event containing the contract specifics and raw log

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
func (it *BuyDeleteDivCurrencyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyDeleteDivCurrency)
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
		it.Event = new(BuyDeleteDivCurrency)
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
func (it *BuyDeleteDivCurrencyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyDeleteDivCurrencyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyDeleteDivCurrency represents a DeleteDivCurrency event raised by the Buy contract.
type BuyDeleteDivCurrency struct {
	Currency common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteDivCurrency is a free log retrieval operation binding the contract event 0x1c7c8a8c8a9b613bdba2dce2cbe6d748eba23cb6d907221c9de511ca43777b61.
//
// Solidity: event DeleteDivCurrency(address indexed currency)
func (_Buy *BuyFilterer) FilterDeleteDivCurrency(opts *bind.FilterOpts, currency []common.Address) (*BuyDeleteDivCurrencyIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "DeleteDivCurrency", currencyRule)
	if err != nil {
		return nil, err
	}
	return &BuyDeleteDivCurrencyIterator{contract: _Buy.contract, event: "DeleteDivCurrency", logs: logs, sub: sub}, nil
}

// WatchDeleteDivCurrency is a free log subscription operation binding the contract event 0x1c7c8a8c8a9b613bdba2dce2cbe6d748eba23cb6d907221c9de511ca43777b61.
//
// Solidity: event DeleteDivCurrency(address indexed currency)
func (_Buy *BuyFilterer) WatchDeleteDivCurrency(opts *bind.WatchOpts, sink chan<- *BuyDeleteDivCurrency, currency []common.Address) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "DeleteDivCurrency", currencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyDeleteDivCurrency)
				if err := _Buy.contract.UnpackLog(event, "DeleteDivCurrency", log); err != nil {
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

// ParseDeleteDivCurrency is a log parse operation binding the contract event 0x1c7c8a8c8a9b613bdba2dce2cbe6d748eba23cb6d907221c9de511ca43777b61.
//
// Solidity: event DeleteDivCurrency(address indexed currency)
func (_Buy *BuyFilterer) ParseDeleteDivCurrency(log types.Log) (*BuyDeleteDivCurrency, error) {
	event := new(BuyDeleteDivCurrency)
	if err := _Buy.contract.UnpackLog(event, "DeleteDivCurrency", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyNewOrderIterator is returned from FilterNewOrder and is used to iterate over the raw logs and unpacked data for NewOrder events raised by the Buy contract.
type BuyNewOrderIterator struct {
	Event *BuyNewOrder // Event containing the contract specifics and raw log

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
func (it *BuyNewOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyNewOrder)
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
		it.Event = new(BuyNewOrder)
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
func (it *BuyNewOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyNewOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyNewOrder represents a NewOrder event raised by the Buy contract.
type BuyNewOrder struct {
	Arg0 BuyOrder
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewOrder is a free log retrieval operation binding the contract event 0x9e83cecc3f2665c32cab8e7ee3becaec6263745877c15c84e9e56f877cd4e768.
//
// Solidity: event NewOrder((address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint256,uint256) arg0)
func (_Buy *BuyFilterer) FilterNewOrder(opts *bind.FilterOpts) (*BuyNewOrderIterator, error) {

	logs, sub, err := _Buy.contract.FilterLogs(opts, "NewOrder")
	if err != nil {
		return nil, err
	}
	return &BuyNewOrderIterator{contract: _Buy.contract, event: "NewOrder", logs: logs, sub: sub}, nil
}

// WatchNewOrder is a free log subscription operation binding the contract event 0x9e83cecc3f2665c32cab8e7ee3becaec6263745877c15c84e9e56f877cd4e768.
//
// Solidity: event NewOrder((address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint256,uint256) arg0)
func (_Buy *BuyFilterer) WatchNewOrder(opts *bind.WatchOpts, sink chan<- *BuyNewOrder) (event.Subscription, error) {

	logs, sub, err := _Buy.contract.WatchLogs(opts, "NewOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyNewOrder)
				if err := _Buy.contract.UnpackLog(event, "NewOrder", log); err != nil {
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

// ParseNewOrder is a log parse operation binding the contract event 0x9e83cecc3f2665c32cab8e7ee3becaec6263745877c15c84e9e56f877cd4e768.
//
// Solidity: event NewOrder((address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint256,uint256) arg0)
func (_Buy *BuyFilterer) ParseNewOrder(log types.Log) (*BuyNewOrder, error) {
	event := new(BuyNewOrder)
	if err := _Buy.contract.UnpackLog(event, "NewOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Buy contract.
type BuyOwnershipTransferredIterator struct {
	Event *BuyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BuyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyOwnershipTransferred)
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
		it.Event = new(BuyOwnershipTransferred)
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
func (it *BuyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyOwnershipTransferred represents a OwnershipTransferred event raised by the Buy contract.
type BuyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Buy *BuyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BuyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BuyOwnershipTransferredIterator{contract: _Buy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Buy *BuyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BuyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyOwnershipTransferred)
				if err := _Buy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Buy *BuyFilterer) ParseOwnershipTransferred(log types.Log) (*BuyOwnershipTransferred, error) {
	event := new(BuyOwnershipTransferred)
	if err := _Buy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuySetRateIterator is returned from FilterSetRate and is used to iterate over the raw logs and unpacked data for SetRate events raised by the Buy contract.
type BuySetRateIterator struct {
	Event *BuySetRate // Event containing the contract specifics and raw log

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
func (it *BuySetRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuySetRate)
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
		it.Event = new(BuySetRate)
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
func (it *BuySetRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuySetRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuySetRate represents a SetRate event raised by the Buy contract.
type BuySetRate struct {
	UnderWritingRate *big.Int
	SurplusRate      *big.Int
	TreasuryRate     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetRate is a free log retrieval operation binding the contract event 0x52543716810f73c3fa9bca74622aecb6d3614ca4991472f3e999d531c2f6afb8.
//
// Solidity: event SetRate(uint256 underWritingRate, uint256 surplusRate, uint256 treasuryRate)
func (_Buy *BuyFilterer) FilterSetRate(opts *bind.FilterOpts) (*BuySetRateIterator, error) {

	logs, sub, err := _Buy.contract.FilterLogs(opts, "SetRate")
	if err != nil {
		return nil, err
	}
	return &BuySetRateIterator{contract: _Buy.contract, event: "SetRate", logs: logs, sub: sub}, nil
}

// WatchSetRate is a free log subscription operation binding the contract event 0x52543716810f73c3fa9bca74622aecb6d3614ca4991472f3e999d531c2f6afb8.
//
// Solidity: event SetRate(uint256 underWritingRate, uint256 surplusRate, uint256 treasuryRate)
func (_Buy *BuyFilterer) WatchSetRate(opts *bind.WatchOpts, sink chan<- *BuySetRate) (event.Subscription, error) {

	logs, sub, err := _Buy.contract.WatchLogs(opts, "SetRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuySetRate)
				if err := _Buy.contract.UnpackLog(event, "SetRate", log); err != nil {
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

// ParseSetRate is a log parse operation binding the contract event 0x52543716810f73c3fa9bca74622aecb6d3614ca4991472f3e999d531c2f6afb8.
//
// Solidity: event SetRate(uint256 underWritingRate, uint256 surplusRate, uint256 treasuryRate)
func (_Buy *BuyFilterer) ParseSetRate(log types.Log) (*BuySetRate, error) {
	event := new(BuySetRate)
	if err := _Buy.contract.UnpackLog(event, "SetRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuySetSignerIterator is returned from FilterSetSigner and is used to iterate over the raw logs and unpacked data for SetSigner events raised by the Buy contract.
type BuySetSignerIterator struct {
	Event *BuySetSigner // Event containing the contract specifics and raw log

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
func (it *BuySetSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuySetSigner)
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
		it.Event = new(BuySetSigner)
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
func (it *BuySetSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuySetSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuySetSigner represents a SetSigner event raised by the Buy contract.
type BuySetSigner struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSigner is a free log retrieval operation binding the contract event 0xbb10aee7ef5a307b8097c6a7f2892b909ff1736fd24a6a5260640c185f7153b6.
//
// Solidity: event SetSigner(address indexed signer)
func (_Buy *BuyFilterer) FilterSetSigner(opts *bind.FilterOpts, signer []common.Address) (*BuySetSignerIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "SetSigner", signerRule)
	if err != nil {
		return nil, err
	}
	return &BuySetSignerIterator{contract: _Buy.contract, event: "SetSigner", logs: logs, sub: sub}, nil
}

// WatchSetSigner is a free log subscription operation binding the contract event 0xbb10aee7ef5a307b8097c6a7f2892b909ff1736fd24a6a5260640c185f7153b6.
//
// Solidity: event SetSigner(address indexed signer)
func (_Buy *BuyFilterer) WatchSetSigner(opts *bind.WatchOpts, sink chan<- *BuySetSigner, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "SetSigner", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuySetSigner)
				if err := _Buy.contract.UnpackLog(event, "SetSigner", log); err != nil {
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
func (_Buy *BuyFilterer) ParseSetSigner(log types.Log) (*BuySetSigner, error) {
	event := new(BuySetSigner)
	if err := _Buy.contract.UnpackLog(event, "SetSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuySetSurplusIterator is returned from FilterSetSurplus and is used to iterate over the raw logs and unpacked data for SetSurplus events raised by the Buy contract.
type BuySetSurplusIterator struct {
	Event *BuySetSurplus // Event containing the contract specifics and raw log

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
func (it *BuySetSurplusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuySetSurplus)
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
		it.Event = new(BuySetSurplus)
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
func (it *BuySetSurplusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuySetSurplusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuySetSurplus represents a SetSurplus event raised by the Buy contract.
type BuySetSurplus struct {
	Surplus common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetSurplus is a free log retrieval operation binding the contract event 0x27f2e1ae96e62d9358b7a2ca44b3af435ece8555438ac9c146019591b4988ecf.
//
// Solidity: event SetSurplus(address indexed surplus)
func (_Buy *BuyFilterer) FilterSetSurplus(opts *bind.FilterOpts, surplus []common.Address) (*BuySetSurplusIterator, error) {

	var surplusRule []interface{}
	for _, surplusItem := range surplus {
		surplusRule = append(surplusRule, surplusItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "SetSurplus", surplusRule)
	if err != nil {
		return nil, err
	}
	return &BuySetSurplusIterator{contract: _Buy.contract, event: "SetSurplus", logs: logs, sub: sub}, nil
}

// WatchSetSurplus is a free log subscription operation binding the contract event 0x27f2e1ae96e62d9358b7a2ca44b3af435ece8555438ac9c146019591b4988ecf.
//
// Solidity: event SetSurplus(address indexed surplus)
func (_Buy *BuyFilterer) WatchSetSurplus(opts *bind.WatchOpts, sink chan<- *BuySetSurplus, surplus []common.Address) (event.Subscription, error) {

	var surplusRule []interface{}
	for _, surplusItem := range surplus {
		surplusRule = append(surplusRule, surplusItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "SetSurplus", surplusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuySetSurplus)
				if err := _Buy.contract.UnpackLog(event, "SetSurplus", log); err != nil {
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

// ParseSetSurplus is a log parse operation binding the contract event 0x27f2e1ae96e62d9358b7a2ca44b3af435ece8555438ac9c146019591b4988ecf.
//
// Solidity: event SetSurplus(address indexed surplus)
func (_Buy *BuyFilterer) ParseSetSurplus(log types.Log) (*BuySetSurplus, error) {
	event := new(BuySetSurplus)
	if err := _Buy.contract.UnpackLog(event, "SetSurplus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuySetTreasuryIterator is returned from FilterSetTreasury and is used to iterate over the raw logs and unpacked data for SetTreasury events raised by the Buy contract.
type BuySetTreasuryIterator struct {
	Event *BuySetTreasury // Event containing the contract specifics and raw log

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
func (it *BuySetTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuySetTreasury)
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
		it.Event = new(BuySetTreasury)
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
func (it *BuySetTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuySetTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuySetTreasury represents a SetTreasury event raised by the Buy contract.
type BuySetTreasury struct {
	Treasury common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetTreasury is a free log retrieval operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address indexed treasury)
func (_Buy *BuyFilterer) FilterSetTreasury(opts *bind.FilterOpts, treasury []common.Address) (*BuySetTreasuryIterator, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "SetTreasury", treasuryRule)
	if err != nil {
		return nil, err
	}
	return &BuySetTreasuryIterator{contract: _Buy.contract, event: "SetTreasury", logs: logs, sub: sub}, nil
}

// WatchSetTreasury is a free log subscription operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address indexed treasury)
func (_Buy *BuyFilterer) WatchSetTreasury(opts *bind.WatchOpts, sink chan<- *BuySetTreasury, treasury []common.Address) (event.Subscription, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "SetTreasury", treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuySetTreasury)
				if err := _Buy.contract.UnpackLog(event, "SetTreasury", log); err != nil {
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

// ParseSetTreasury is a log parse operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address indexed treasury)
func (_Buy *BuyFilterer) ParseSetTreasury(log types.Log) (*BuySetTreasury, error) {
	event := new(BuySetTreasury)
	if err := _Buy.contract.UnpackLog(event, "SetTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuySetUnderWritingIterator is returned from FilterSetUnderWriting and is used to iterate over the raw logs and unpacked data for SetUnderWriting events raised by the Buy contract.
type BuySetUnderWritingIterator struct {
	Event *BuySetUnderWriting // Event containing the contract specifics and raw log

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
func (it *BuySetUnderWritingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuySetUnderWriting)
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
		it.Event = new(BuySetUnderWriting)
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
func (it *BuySetUnderWritingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuySetUnderWritingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuySetUnderWriting represents a SetUnderWriting event raised by the Buy contract.
type BuySetUnderWriting struct {
	UnderWriting common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetUnderWriting is a free log retrieval operation binding the contract event 0xb0a416724e35cbe1d796ba3a4b16b5406be2b9f06147393a81a10ab3d8f0a070.
//
// Solidity: event SetUnderWriting(address indexed underWriting)
func (_Buy *BuyFilterer) FilterSetUnderWriting(opts *bind.FilterOpts, underWriting []common.Address) (*BuySetUnderWritingIterator, error) {

	var underWritingRule []interface{}
	for _, underWritingItem := range underWriting {
		underWritingRule = append(underWritingRule, underWritingItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "SetUnderWriting", underWritingRule)
	if err != nil {
		return nil, err
	}
	return &BuySetUnderWritingIterator{contract: _Buy.contract, event: "SetUnderWriting", logs: logs, sub: sub}, nil
}

// WatchSetUnderWriting is a free log subscription operation binding the contract event 0xb0a416724e35cbe1d796ba3a4b16b5406be2b9f06147393a81a10ab3d8f0a070.
//
// Solidity: event SetUnderWriting(address indexed underWriting)
func (_Buy *BuyFilterer) WatchSetUnderWriting(opts *bind.WatchOpts, sink chan<- *BuySetUnderWriting, underWriting []common.Address) (event.Subscription, error) {

	var underWritingRule []interface{}
	for _, underWritingItem := range underWriting {
		underWritingRule = append(underWritingRule, underWritingItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "SetUnderWriting", underWritingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuySetUnderWriting)
				if err := _Buy.contract.UnpackLog(event, "SetUnderWriting", log); err != nil {
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

// ParseSetUnderWriting is a log parse operation binding the contract event 0xb0a416724e35cbe1d796ba3a4b16b5406be2b9f06147393a81a10ab3d8f0a070.
//
// Solidity: event SetUnderWriting(address indexed underWriting)
func (_Buy *BuyFilterer) ParseSetUnderWriting(log types.Log) (*BuySetUnderWriting, error) {
	event := new(BuySetUnderWriting)
	if err := _Buy.contract.UnpackLog(event, "SetUnderWriting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
