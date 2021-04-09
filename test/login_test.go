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

func TestSignIn(t *testing.T) {
	reader := strings.NewReader(`{"user_id":"0x96216849c49358B10257cb55b28eA603c874b05E", "sig_hex":"0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301", "msg":"hello"}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}
func TestSignTTT(t *testing.T) {
	reader := strings.NewReader(`{"user_id":"0x11f4d0A3c12e86B4b5F39B213F7E19D048276DAe", "sig_hex":"0x30755ed65396facf86c53e6217c52b4daebe72aa4941d89635409de4c9c7f9466d4e9aaec7977f05e923889b33c0d0dd27d7226b6e6f56ce737465c5cfd04be400", "msg":"Hello world"}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestSignAA(t *testing.T) {
	msg := string(`hello`)
	fmt.Println(msg)
	reader := strings.NewReader(`{"user_id":"0x7302509251F3bdA6428365bAe08206D4C3D3173C", "sig_hex":"0xc5ce206fa055aa76a94781017b262256d3017cdcb42ce077bd9a4442a57f5f643ab8137f004d941d465d4438d114524bf7a1a78b550b134d911fea1f3c84d0331b", "msg":"` + msg + `"}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}
func TestSign0309(t *testing.T) {
	msg := string(`{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Mail":[{"name":"from","type":"Person"},{"name":"to","type":"Person"},{"name":"contents","type":"string"}],"Person":[{"name":"name","type":"string"},{"name":"wallet","type":"address"}]},"primaryType":"Mail","domain":{"name":"Ether Mail","version":"1","chainId":"0x1","verifyingContract":"0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC","salt":""},"message":{"contents":"Hello, Bob!","from":{"name":"Cow","wallet":"0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826"},"to":{"name":"Bob","wallet":"0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB"}}}`)
	fmt.Println(msg)
	reader := strings.NewReader(`{"user_id":"0x96216849c49358B10257cb55b28eA603c874b05E", "sig_hex":"0x5bef8c3bda3554ae658577689d5fb8eaf9b6a25ca850cdc051a00ecdd1e29e1909e954241e40ae0b98acfcc45f25887b8eb79e3c72a8dce0adca8ba3d8960f0501", "msg":` + msg + `}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}
func TestSignII(t *testing.T) {
	msg := string(`{"domain":{"name":"unStake","chainId":"42"},"message":{"to":[{"products":"1","amounts":1.23e+21}]},"primaryType":"unStake","types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"chainId","type":"uint256"}],"unStake":[{"to":[{"products":"string","amounts":"uint256"}]}]}}`)
	msg = string(`{"domain":{"name":"unStake","chainId":"0x2a"},"message":{"to":[{"products":"1","amounts":1.23e+21}]},"primaryType":"unStake","types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"chainId","type":"uint256"}],"Person":[{"name":"products","type":"string"},{"name":"amounts","type":"uint256"}],"unStake":[{"name":"to","type":"Person"}]}}`)
	fmt.Println(msg)
	reader := strings.NewReader(`{"user_id":"0x7302509251F3bdA6428365bAe08206D4C3D3173C", "sig_hex":"0xc5ce206fa055aa76a94781017b262256d3017cdcb42ce077bd9a4442a57f5f643ab8137f004d941d465d4438d114524bf7a1a78b550b134d911fea1f3c84d0331b", "msg":` + msg + `}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestSignTest(t *testing.T) {
	reader := strings.NewReader(`{"user_id":"0x12345678", "sig_hex":"kk", "msg":"123"}`)

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/api/login", reader)
	router := rest.SetupRouter()
	router.ServeHTTP(w1, req1)

	assert.Equal(t, 200, w1.Code)
	fmt.Println(w1.Body.String())

	token := w1.Body.String()

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/api/loginTest", nil)
	//cookie1 := &http.Cookie{Name: "accessToken",Value: token, Path:"/", Domain:"aa", HttpOnly: true}
	cookie1 := &http.Cookie{Name: "accessToken", Value: token}
	//ctx.SetCookie("accessToken", token, 7*24*60*60, "/", "", true, false)
	req2.AddCookie(cookie1)

	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
	fmt.Println(w2.Body.String())

}

func TestLoginRandom(t *testing.T) {
	router := rest.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/login?userId=0x1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestLogin(t *testing.T) {
	reader := strings.NewReader(`{"user_id":"0x1"}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", reader)

	router := rest.SetupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestLogin11(t *testing.T) {
	router := rest.SetupRouter()
	req1, _ := http.NewRequest("GET", "/api/login?userId=0x123", nil)
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	assert.Equal(t, 200, w1.Code)
	fmt.Println(w1.Body.String())

	reader := strings.NewReader(`{"user_id":"0x123"}`)
	req2, _ := http.NewRequest("POST", "/api/login", reader)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
	fmt.Println(w2.Body.String())
}

func TestLogin22(t *testing.T) {
	router := rest.SetupRouter()
	req1, _ := http.NewRequest("GET", "/api/login?userId=0x1234", nil)
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	assert.Equal(t, 200, w1.Code)
	fmt.Println(w1.Body.String())

	reader := strings.NewReader(`{"user_id":"0x1234"}`)
	req2, _ := http.NewRequest("POST", "/api/login", reader)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
	fmt.Println(w2.Body.String())

	reader3 := strings.NewReader(`{"user_id":"0x1234"}`)
	req3, _ := http.NewRequest("POST", "/api/login", reader3)
	w3 := httptest.NewRecorder()
	router.ServeHTTP(w3, req3)

	assert.Equal(t, 400, w3.Code)
	fmt.Println(w3.Body.String())
}
