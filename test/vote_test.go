package test

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"nsure/vote/common"
	"nsure/vote/contract"
	"nsure/vote/rest"
	"nsure/vote/utils"
	"strings"
	"testing"
)

func TestGetAssessment(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/assessment?claimId=9", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestGetResultAssessment(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/assessmentResult?claimId=9", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestClaimVote(t *testing.T) {
	w := httptest.NewRecorder()
	reader := strings.NewReader(`{"claim_id":1,"arbiter_id":"0x1234","status":false,"sign_hash":"0x92cccddd"}`)
	req, _ := http.NewRequest("POST", "/api/vote", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestGetClaimApply(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/claimApply?userId=0x1230", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestAAAMessage(t *testing.T) {
	type AA struct {
		Amount int
	}
	type Person struct {
		name string
		Age  int
		BB   AA
	}
	A := make(map[string]interface{})
	A["Amount"] = 1

	mapInstance := make(map[string]interface{})
	mapInstance["name"] = "liang637210"
	mapInstance["Age"] = 28
	mapInstance["BB"] = A

	var person Person
	if err := mapstructure.Decode(mapInstance, &person); err != nil {
		fmt.Println(err)
	}
	fmt.Println(person)

	var person1 Person
	if err := utils.MessageToStruct1(&mapInstance, &person1); err != nil {
		fmt.Println(err)
	}
	fmt.Println(person1)
}

func TestSecondMessage(t *testing.T) {
	second := contract.DurationSecond(common.DurationClaim)
	fmt.Println(second)
}
