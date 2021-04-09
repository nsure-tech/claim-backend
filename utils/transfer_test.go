package utils

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestTransMessage(t *testing.T) {
	toAddress := "0x4897284C3Faf8d8764F77a3d0e81F85B33936d7E"
	amount := decimal.New(123000000000000, 0)
	transfer(toAddress, amount)
}
