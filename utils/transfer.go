package utils

import (
	"crypto/ecdsa"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
	"nsure/vote/config"
	"nsure/vote/log"
)

/*
func transfer(address string, amount decimal.Decimal) {
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.RawUrl)
	if err != nil {
		log.GetLog().Error("eth client.Dial", zap.Error(err))
	}
	contrAddress := ethcommon.HexToAddress(conf.ContractAddress)
	nSureContract, err := contract.NewContract(contrAddress, client)
	if err != nil {
		log.GetLog().Error("contract.NewContract", zap.Error(err))
	}

	key, fromAddress, err := getPrivateAddress()
	if err != nil {
		log.GetLog().Error("getPrivateAddress", zap.Error(err))
	}

	toAddress := ethcommon.HexToAddress(address)
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.GetLog().Error("NetworkID", zap.Error(err))
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainId)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.GetLog().Error("PendingNonceAt", zap.Error(err))
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.GetLog().Error("SuggestGasPrice", zap.Error(err))
	}
	value := amount.BigInt()
	tx, err := nSureContract.Transfer(&bind.TransactOpts{
		From:     fromAddress,
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   auth.Signer,
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		GasLimit: 80000,
	}, toAddress, value)
	fmt.Println(tx.Hash())

}*/

func getPrivateAddress() (*ecdsa.PrivateKey, ethcommon.Address, error) {
	privateKey, err := crypto.HexToECDSA(config.GetConfig().KeySecret)
	if err != nil {
		log.GetLog().Error("crypto.HexToECDSA", zap.Error(err))
		return nil, ethcommon.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.GetLog().Error("crypto.HexToECDSA", zap.Error(err))
		return nil, ethcommon.Address{}, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return privateKey, fromAddress, nil
}
