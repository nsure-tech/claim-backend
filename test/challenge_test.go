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

func TestChallengeApply(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()

	reader := strings.NewReader(`{"user_id":"0x772ac0631022d09E0a9134c3BA0D74AA5961a3E0","claim_id":9}`)
	req, _ := http.NewRequest("POST", "/api/challengeOld", reader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestChallengeVote(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()

	reader := strings.NewReader(`{"claim_id":7,"status":true}`)
	req, _ := http.NewRequest("POST", "/api/challengeResultOld", reader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestChallengePrice(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/challenge?claimId=9", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestAdmin(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/challenge?claimId=9", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestGetPayment(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/claimListPay", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}
