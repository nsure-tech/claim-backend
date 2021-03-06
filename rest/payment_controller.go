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
)

func PaymentAdmin(ctx *gin.Context) {
	var metaMask PaymentMetamask
	userId, _, err := GetMetamaskMessage(ctx, &metaMask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	req := &metaMask.Content
	if req.UserId != userId {
		log.GetLog().Info("ChallengeVote", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("req.UserId != userId %v!=%v", req.UserId, userId)))
		return
	}
	pay, err := decimal.NewFromString(req.Pay)
	if err != nil {
		log.GetLog().Info("req.Pay", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("amount %v:%v", req.Pay, err)))
		return
	}

	adminId := utils.Address(userId)
	ret, err := service.PaymentByAdmin(adminId, req.ClaimId, pay)
	if err != nil {
		log.GetLog().Info("PaymentByAdmin", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}

	if ret {
		ctx.JSON(http.StatusOK, newMessageOK())
	} else {
		ctx.JSON(http.StatusOK, newMessageFail())
	}
}

func GetClaimListPayment(ctx *gin.Context) {
	payments, err := service.GetPaymentUnsettled()
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}

	lists := make([]*claimList, len(payments))
	for i, payment := range payments {
		//lists = append(lists, newClaimListPayment(payment))
		lists[i] = newClaimListPayment(payment)
	}

	result := &claimResult{
		Method: "payment",
		List:   lists,
	}

	ctx.JSON(http.StatusOK, newMessageDataOK(result))
}
