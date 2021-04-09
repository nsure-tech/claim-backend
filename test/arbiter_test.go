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

func TestGetArbiter(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()
	// req, _ := http.NewRequest("POST", "/api/arbiter", nil)
	//arbiter?userId=0x666747ffD8417a735dFf70264FDf4e29076c775a
	req, _ := http.NewRequest("GET", "/api/arbiter?userId=0x666747ffD8417a735dFf70264FDf4e29076c775a", nil) // test
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestAddArbiter(t *testing.T) {
	w := httptest.NewRecorder()
	reader := strings.NewReader(`{"user_id":"0x1231","number":3}`)
	req, _ := http.NewRequest("POST", "/api/addArbiter", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestPendingArbiter(t *testing.T) {
	w := httptest.NewRecorder()
	reader := strings.NewReader(`{"user_id":"0x1234","number":2}`)
	req, _ := http.NewRequest("POST", "/api/pending", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}
