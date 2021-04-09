package rest

import (
	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru"
	"net/http"
	"nsure/vote/common"
	"nsure/vote/service"
)

var loginCache *lru.Cache

func init() {
	var err error
	loginCache, err = lru.New(common.LoginLruSize)
	if err != nil {
		panic(err)
	}
}

func GetUserRandom(ctx *gin.Context) {
	userId := ctx.Query("userId")
	userRand, err := service.GetUserRand(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}
	ctx.JSON(http.StatusOK, userRand)
}

func SignIn(ctx *gin.Context) {
	var req1 LoginRequest
	GetMetamaskMessage(ctx, req1)
	var req LoginRequest
	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}
	token, err := service.RefreshAccessToken(req.UserId, req.SigHex, req.Msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}

	ctx.SetCookie("accessToken", token, common.CookieMaxAge, "/", "", true, false)
	ctx.JSON(http.StatusOK, token)
}

func LoginTest(ctx *gin.Context) {
	userId := GetCurrentUser(ctx)
	ctx.JSON(http.StatusOK, userId)
}
