package service

import (
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
)

func GetConfig(keyWord string) (string, error) {
	if config, err := mysql.SharedStore().GetConfig(keyWord); err != nil {
		return "", err
	} else {
		return config.Val, nil
	}
}

func GetCurrency(currency string) (string, error) {
	return GetConfig(common.CurrencyPrefix + currency)
}

func UpdateConfig(keyWork string, val string) error {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()

	if config, err := db.GetConfigForUpdate(keyWork); err != nil {
		return err
	} else {
		config.Val = val
		if err := db.UpdateConfig(config); err != nil {
			return err
		}
	}

	return db.CommitTx()
}

func AddMetamask(userId, sigHex, msg string) {
	metamask := &models.Metamask{
		UserId: userId,
		SigHex: sigHex,
		Msg:    msg,
	}
	mysql.SharedStore().AddMetamask(metamask)
}
