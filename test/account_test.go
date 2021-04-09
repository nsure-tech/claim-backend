package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"nsure/vote/rest"
	"strings"
	"testing"
)

func TestAccountDeposit(t *testing.T) {
	w := httptest.NewRecorder()
	// 1000000000000000000
	//reader := strings.NewReader(`{"user_id":"0x1230","currency":"NSURE","amount":200000000}`)
	reader := strings.NewReader(`{"user_id":"0x8","currency":"NSURE","amount":1000000000000000000000000000000000000000000000}`)
	req, _ := http.NewRequest("POST", "/api/deposit", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestAccountWithdraw(t *testing.T) {
	w := httptest.NewRecorder()
	reader := strings.NewReader(`{"user_id":"0x2","currency":"NSURE","amount":1000}`)
	req, _ := http.NewRequest("POST", "/api/withdraw", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestClaimWithdraw(t *testing.T) {
	w := httptest.NewRecorder()
	reader := strings.NewReader(`{"account":"0x2e9475c282069675fFAc22a8cd5038E4DAC01634", "currency":"1","amount":"1234","nonce":"0"}`)
	req, _ := http.NewRequest("POST", "/api/withdrawTest", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestAssetHistory(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/asset?userId=0x666747ffD8417a735dFf70264FDf4e29076c775a&offset=2&limit=2", nil)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}
