package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"net/http"
	"nsure/vote/log"
	"nsure/vote/service"
	"nsure/vote/utils"
	"strconv"
)

func PlaceWithdraw(ctx *gin.Context) {
	var metaMask WithdrawMetamask
	userId, _, err := GetMetamaskMessage(ctx, &metaMask)
	if err != nil {
		log.GetLog().Info("GetMetamaskMessage", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	req := &metaMask.Content
	log.GetLog().Debug("place withdraw", zap.String("userId", userId))

	if req.UserId != userId {
		log.GetLog().Error("userId err", zap.String("req userId", req.UserId), zap.String("userId", userId))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("req.UserId != userId %v!=%v", req.UserId, userId)))
		return
	}
	nonce, err := strconv.ParseUint(req.Nonce, 0, 0)
	if err != nil {
		log.GetLog().Error("ParseUint", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("nonce  %v:%v", req.Nonce, err)))
		return
	}
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		log.GetLog().Error("req Amount", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("amount  %v:%v", req.Amount, err)))
		return
	}
	if amount.LessThanOrEqual(decimal.Zero) {
		log.GetLog().Error("req Amount", zap.String("amount", amount.String()))
		ctx.JSON(http.StatusOK, newMessageBadRequest(fmt.Errorf("amount  %v <= 0", req.Amount)))
		return
	}

	userId = utils.Address(userId)
	ret, err := service.AddWithdraw(userId, req.Currency, amount, nonce)
	if err != nil {
		log.GetLog().Error("AddWithdraw", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, newMessageDataOK(ret))
}

func PlaceWithdrawTest(ctx *gin.Context) {
	var req WithDrawRequestMessage
	err := ctx.BindJSON(&req)
	if err != nil {
		log.GetLog().Info("BindJSON", zap.Error(err))
		return
	}
	nonce, err := strconv.ParseUint(req.Nonce, 0, 0)
	if err != nil {
		log.GetLog().Info("ParseUint", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("nonce  %v:%v", req.Nonce, err)))
		return
	}
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		log.GetLog().Info("req.Amount", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("amount %v:%v", req.Amount, err)))
		return
	}

	userId := utils.Address(req.UserId)
	ret, err := service.AddWithdraw(userId, req.Currency, amount, nonce)
	if err != nil {
		log.GetLog().Info("AddWithdraw", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, newMessageDataOK(ret))
}
