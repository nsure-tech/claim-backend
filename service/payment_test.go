package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestPayment(t *testing.T) {
	pay := decimal.NewFromInt(10000)
	if ret, err := PaymentByAdmin("0x11", 25, pay); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
}
