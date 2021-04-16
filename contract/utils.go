package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"nsure/vote/common"
	"nsure/vote/config"
	"nsure/vote/log"
	"sync"
	"time"
)

var clientOnce sync.Once
var chinaIdOnce sync.Once
var keyOnce sync.Once
var client *ethclient.Client
var chainID *big.Int
var key *ecdsa.PrivateKey

func GetClient() *ethclient.Client {
	clientOnce.Do(func() {
		var err error
		client, err = ethclient.Dial(config.GetConfig().RawUrl)
		if err != nil {
			panic(err)
		}
	})
	return client
}

func GetChainId() *big.Int {
	chinaIdOnce.Do(func() {
		var err error
		chainID, err = GetClient().NetworkID(context.Background())
		if err != nil {
			panic(err)
		}
	})
	return chainID
}

func GetNonceByAddress(address string) (uint64, error) {
	contrAddress := ethcommon.HexToAddress(config.GetConfig().ContractTreasuryAddress)
	treasuryContract, err := NewContract(contrAddress, GetClient())
	if err != nil {
		log.GetLog().Error("contract Treasury", zap.Error(err))
		return 0, err
	}
	argAddress := ethcommon.HexToAddress(address)
	nonce, err := treasuryContract.Nonces(nil, argAddress)
	if err != nil {
		log.GetLog().Error("contract Nonce", zap.Error(err))
		return 0, err
	}
	return nonce.Uint64(), nil
}

func GetBuyClaim(coverId *big.Int) (*common.BuyClaim, error) {
	contrAddress := ethcommon.HexToAddress(config.GetConfig().ContractBuyAddress)
	buyContract, err := NewBuy(contrAddress, GetClient())
	if err != nil {
		log.GetLog().Error("contract NewBuy", zap.Error(err))
		return nil, err
	}
	buy, err := buyContract.InsuranceOrders(nil, coverId)
	if err != nil {
		log.GetLog().Error("contract InsuranceOrders", zap.Error(err))
		return nil, err
	}
	if buy.Period.Int64() <= 0 {
		log.GetLog().Error("contract error coverId", zap.Any("coverId", coverId))
		return nil, fmt.Errorf("coverId:%v", coverId)
	}
	return &common.BuyClaim{
		Buyer:    buy.Buyer.String(),
		Amount:   decimal.NewFromBigInt(buy.Amount, 0),
		Cost:     decimal.NewFromBigInt(buy.Premium, 0),
		Reward:   decimal.NewFromBigInt(buy.Premium.Div(buy.Premium, big.NewInt(10)), 0),
		Period:   buy.Period.Uint64(),
		CreateAt: time.Unix(buy.CreateAt.Int64(), 0),
	}, nil
}

func getPrivateAddress() *ecdsa.PrivateKey {
	keyOnce.Do(func() {
		var err error
		key, err = crypto.HexToECDSA(config.GetConfig().KeySecret)
		if err != nil {
			panic(err)
		}
	})
	return key
}

func GetFromByTxHash(strHash string) (string, error) {
	txHash := ethcommon.HexToHash(strHash)
	tx, _, err := GetClient().TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.GetLog().Error("TransactionByHash", zap.Error(err))
		return "", err
	}

	if msg, err := tx.AsMessage(types.NewEIP155Signer(GetChainId())); err != nil {
		log.GetLog().Error("AsMessage", zap.Error(err))
		return "", err
	} else {
		return msg.From().String(), nil
	}
}

func SignClaim(account, currency, amount, nonce, deadline string) (hexutil.Bytes, hexutil.Bytes, error) {
	claim := CreateDataClaim(config.GetConfig().ChainId, config.GetConfig().ContractTreasuryAddress,
		account, currency, amount, nonce, deadline)
	return GetSign(&claim)
}
func SignWithdraw(account, amount, nonce, deadline string) (hexutil.Bytes, hexutil.Bytes, error) {
	withdraw := CreateDataWithdraw(config.GetConfig().ChainId, config.GetConfig().ContractTreasuryAddress,
		account, amount, nonce, deadline)
	return GetSign(&withdraw)
}

func GetSign(data *core.TypedData) (hexutil.Bytes, hexutil.Bytes, error) {
	key := getPrivateAddress()
	return Sign(key, data)
}

func Sign(key *ecdsa.PrivateKey, data *core.TypedData) (hexutil.Bytes, hexutil.Bytes, error) {
	domainSeparator, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return nil, nil, err
	}
	typedDataHash, err := data.HashStruct(data.PrimaryType, data.Message)
	if err != nil {
		return nil, nil, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	dataHash := crypto.Keccak256(rawData)

	signature, err := crypto.Sign(dataHash, key)
	if err != nil {
		return nil, nil, err
	}
	return signature, dataHash, nil
}

func CreateDataClaim(chainId int64, contract string, account, currency, amount, nonce, deadline string) core.TypedData {
	return core.TypedData{
		Types:       claimType,
		PrimaryType: claimPrimaryType,
		Domain:      createDomainClaim(chainId, contract),
		Message:     createMessageClaim(account, currency, amount, nonce, deadline),
	}
}

func CreateDataWithdraw(chainId int64, contract string, account, amount, nonce, deadline string) core.TypedData {
	return core.TypedData{
		Types:       withDrawType,
		PrimaryType: withdrawPrimaryType,
		Domain:      createDomainWithdraw(chainId, contract),
		Message:     createMessageWithdraw(account, amount, nonce, deadline),
	}
}

func createDomainClaim(chainId int64, contract string) core.TypedDataDomain {
	return core.TypedDataDomain{
		Name:              common.TypedDataDomainNameClaim,
		Version:           "1",
		ChainId:           math.NewHexOrDecimal256(chainId),
		VerifyingContract: contract,
		Salt:              "",
	}
}

func createDomainWithdraw(chainId int64, contract string) core.TypedDataDomain {
	return core.TypedDataDomain{
		Name:              common.TypedDataDomainNameWithdraw,
		Version:           "1",
		ChainId:           math.NewHexOrDecimal256(chainId),
		VerifyingContract: contract,
		Salt:              "",
	}
}

func createMessageClaim(account, currency, amount, nonce, deadline string) map[string]interface{} {
	return map[string]interface{}{
		"account":  account,
		"currency": currency,
		"amount":   amount,
		"nonce":    nonce,
		"deadline": deadline,
	}
}
func createMessageWithdraw(account, amount, nonce, deadline string) map[string]interface{} {
	return map[string]interface{}{
		"account":  account,
		"amount":   amount,
		"nonce":    nonce,
		"deadline": deadline,
	}
}

func DurationSecond(duration int) int64 {
	return time.Now().Add(time.Duration(duration) * time.Minute).Unix()
}
