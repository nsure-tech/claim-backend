package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
	lru "github.com/hashicorp/golang-lru"
	"math/big"
	"nsure/vote/common"
	"nsure/vote/config"
	"nsure/vote/utils"
	"time"
)

var loginCache *lru.Cache

func init() {
	var err error
	loginCache, err = lru.New(common.LoginLruSize)
	if err != nil {
		panic(err)
	}
}
func VoteAddress(address string) bool {
	if _, found := common.VoteAddress[address]; found {
		return true
	}
	return false
}
func ChallengeAddress(address string) bool {
	if _, found := common.ChallengeAddress[address]; found {
		return true
	}
	return false
}
func PaymentAddress(address string) bool {
	if _, found := common.PaymentAddress[address]; found {
		return true
	}
	return false
}

func GetUserRand(userId string) (string, error) {
	randInt, err := rand.Int(rand.Reader, big.NewInt(common.MaxRandNum))
	if err != nil {
		return "", err
	}

	randString := utils.I64ToA(randInt.Int64())
	loginCache.Add(userId, randString)
	return randString, nil
}

func GetTypedDataMessage(userId, sigHex string, msg core.TypedData) (*core.TypedDataMessage, error) {
	if err := verifySigMessage(userId, sigHex, msg); err != nil {
		return nil, err
	}
	return &msg.Message, nil
}

func RefreshAccessToken(userId, sigHex string, msg core.TypedData) (string, error) {
	if !verifySig(userId, sigHex, msg) {
		return "", errors.New("user sign error")
	}

	claim := jwt.MapClaims{
		"id":        userId,
		"expiredAt": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	fmt.Println(token)
	return token.SignedString([]byte(config.GetConfig().JwtSecret))

}

/*
func RefreshAccessToken(userId, sigHex, msg string) (string, error) {
	randString, ok := loginCache.Get(userId)
	if !ok {
		return "", errors.New("no login random num")
	}
	randInt, err := rand.Int(rand.Reader, big.NewInt(common.MaxRandNum))
	if err != nil {
		return "", err
	}

	randString := utils.I64ToA(randInt.Int64())
	loginCache.Add(userId, randString)
	return randString, nil
}*/

func verifySigMessage(from, sigHex string, msg core.TypedData) error {
	fromAddr := ethcommon.HexToAddress(from)
	sig := hexutil.MustDecode(sigHex)
	if sig[64] >= 27 {
		sig[64] -= 27
	}
	sigHash, err := signHash(msg)
	if err != nil {
		return err
	}

	pubKey, err := crypto.SigToPub(sigHash, sig)
	if err != nil {
		return err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if fromAddr == recoveredAddr {
		return nil
	}

	return fmt.Errorf("addr from!=recover: %v!=%v", fromAddr, recoveredAddr)
}

func verifySig(from, sigHex string, msg core.TypedData) bool {
	fromAddr := ethcommon.HexToAddress(from)
	sig := hexutil.MustDecode(sigHex)
	fmt.Println("sigHex:")
	fmt.Println(sigHex)
	fmt.Println(msg)
	fmt.Println(msg.Message)
	fmt.Println(msg.Types)
	sigHash, err := signHash(msg)
	if err != nil {
		fmt.Println(err)
		fmt.Println("111111111111111111")
		return false
	}
	if sig[64] >= 27 {
		sig[64] -= 27
	}
	pubKey, err := crypto.SigToPub(sigHash, sig)
	if err != nil {
		fmt.Println(err)
		fmt.Println("2222222222222222222")
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	fmt.Println(fromAddr)
	fmt.Println(recoveredAddr)
	return fromAddr == recoveredAddr
}

func signHash(msg core.TypedData) ([]byte, error) {
	domainSeparator, err := msg.HashStruct("EIP712Domain", msg.Domain.Map())
	if err != nil {
		return nil, err
	}

	typedDataHash, err := msg.HashStruct(msg.PrimaryType, msg.Message)
	if err != nil {
		return nil, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))

	return crypto.Keccak256(rawData), nil

}

func CheckToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JwtSecret), nil
	})
	if err != nil {
		return "", err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("cannot convert claim to MapClaims")
	}
	if !token.Valid {
		return "", errors.New("token is invalid")
	}

	idVal, found := claim["id"]
	if !found {
		return "", errors.New("bad token")
	}
	userId := idVal.(string)

	return userId, nil
}
