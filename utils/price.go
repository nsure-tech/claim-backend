package utils

import (
	"encoding/json"
	"errors"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"nsure/vote/common"
)

type Ethereum struct {
	Eth Usd `json:"ethereum"`
}

type NSure struct {
	NSure Usd `json:"nsure-network"`
}

type Usd struct {
	Usd float64 `json:"usd"`
}

func EthPrice() (decimal.Decimal, error) {
	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd")
	if err != nil {
		return decimal.Zero, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return decimal.Zero, err
	}
	var eth Ethereum
	err = json.Unmarshal(body, &eth)
	if err != nil {
		return decimal.Zero, err
	}
	return decimal.NewFromFloat(eth.Eth.Usd), nil
}

func NSurePrice() (decimal.Decimal, error) {
	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=nsure-network&vs_currencies=usd")
	if err != nil {
		return decimal.Zero, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return decimal.Zero, err
	}
	var nSure NSure
	err = json.Unmarshal(body, &nSure)
	if err != nil {
		return decimal.Zero, err
	}
	return decimal.NewFromFloat(nSure.NSure.Usd), nil
}

func EthNSurePrice() (decimal.Decimal, error) {
	eth, err := EthPrice()
	if err != nil {
		return decimal.Zero, err
	}
	nSure, err := NSurePrice()
	if err != nil {
		return decimal.Zero, err
	}
	if eth == decimal.Zero || nSure == decimal.Zero {
		return decimal.Zero, errors.New("eth or nsure price error")
	}

	return eth.Div(nSure), nil
}

func ChallengeNSure(amount decimal.Decimal) (decimal.Decimal, error) {
	price, err := EthNSurePrice()
	if err != nil {
		return decimal.Zero, err
	}
	nSure := price.Mul(amount).Mul(decimal.NewFromInt(common.ChallengeTimes))

	return nSure.Round(0), nil
}

func ArbiterNSure() decimal.Decimal {
	return decimal.NewFromInt(common.ArbiterNSure).Mul(UnitNSure())
}

func UnitNSure() decimal.Decimal {
	return decimal.NewFromInt(common.UnitNSure)
}
