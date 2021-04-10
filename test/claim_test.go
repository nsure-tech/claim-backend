package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"nsure/vote/rest"
	"nsure/vote/service"
	"nsure/vote/utils"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := rest.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestClaimNew(t *testing.T) {
	w := httptest.NewRecorder()
	reader := strings.NewReader(`{"user_id":"0x923129","product":"product1",
    "currency":"ETH","amount":"1", "reward":"103",
    "cover_id":"126", "cover_hash":"0x11135",
    "cover_begin_at":"2021-04-04T08:08:08+08:00",
    "cover_end_at":"2021-08-04T08:08:08+08:00", "desc":"a1", "cred":"p1"}`)
	req, _ := http.NewRequest("POST", "/claim/new", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println("xxxxxxxxxxxx")
	fmt.Println(w.Body.String())
}

func TestGetNewClaim(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/claimListNew?userId=0x7302509251F3bdA6428365bAe08206D4C3D3173C", nil)
	//req, _ := http.NewRequest("GET", "/api/claimListNew?userId=0x666747ffD8417a735dFf70264FDf4e29076c775a&before=20&after=20&limit=20", nil)
	//req, _ := http.NewRequest("GET", "/api/claimListNew?userId=0x666747ffD8417a735dFf70264FDf4e29076c775a", nil)
	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}
func TestGetClaimUserId(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/claimListUser?userId=0x02Ec1090D59cbAA9D58EAeBb3328dF56A6dEd2D3", nil)
	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestGetDownClaim(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/claimListDown?userId=0x11111", nil)
	//req, _ := http.NewRequest("GET", "/api/claimListNew?userId=0x666747ffD8417a735dFf70264FDf4e29076c775a&before=20&after=20&limit=20", nil)
	//req, _ := http.NewRequest("GET", "/api/claimListNew?userId=0x666747ffD8417a735dFf70264FDf4e29076c775a", nil)
	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestGetClaim(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/claim", nil)
	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestClaimByArbiter(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/claimArbiter?userId=0x1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestOldClaimApply(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()

	reader := strings.NewReader(`{"claim_id":10,"product":"product1"}`)
	req, _ := http.NewRequest("POST", "/api/applyArbiter?userId=0xb78E36E627138D918C09648a25Ca328D0aeA083a", reader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestClaimList(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/claimList?userId=0x923", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestAClaimList(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/claimListArbiter?userId=0x923", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestGetClaimVote(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/claimVote?claimId=4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestServiceVote(t *testing.T) {
	claims, _ := service.GetVoteByVoteNum(1)
	for _, claim := range claims {
		fmt.Println(claim.ClaimId)
	}
}

func TestPrice(t *testing.T) {
	if price, err := utils.EthNSurePrice(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(price)
	}
}
