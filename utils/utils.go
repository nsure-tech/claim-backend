package utils

import (
	"encoding/json"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/goinggo/mapstructure"
	"github.com/shopspring/decimal"
	"strconv"
)

func I64ToA(i int64) string {
	return strconv.FormatInt(i, 10)
}

func U64ToA(i uint64) string {
	return strconv.FormatUint(i, 10)
}
func IntToA(i int) string {
	return strconv.Itoa(i)
}

func Address(userId string) string {
	address := ethcommon.HexToAddress(userId)
	return address.String()
}

func StringToInt64(str string) int64 {
	if i, err := strconv.ParseInt(str, 0, 0); err != nil {
		return 0
	} else {
		return i
	}
}

func StringToUint64(str string) uint64 {
	if i, err := strconv.ParseUint(str, 0, 0); err != nil {
		return 0
	} else {
		return i
	}
}

func DToF64(d decimal.Decimal) float64 {
	f, _ := d.Float64()
	return f
}

func DToString(d decimal.Decimal) string {
	return d.String()
}

func DecimalDiv(d decimal.Decimal, i int64) decimal.Decimal {
	di := decimal.NewFromInt(i)
	return d.DivRound(di, 0)
}

func MessageToStruct111(message *core.TypedDataMessage, val interface{}) error {
	if err := mapstructure.Decode(message, val); err != nil {
		return err
	}
	return nil
}

func MessageToStruct(message *core.TypedDataMessage, val interface{}) error {
	if msgJson, err := json.Marshal(message); err != nil {
		return err
	} else {
		return json.Unmarshal(msgJson, val)
	}
}

func MessageToStruct1(message *map[string]interface{}, val interface{}) error {
	if err := mapstructure.Decode(*message, val); err != nil {
		return err
	}
	return nil
}
