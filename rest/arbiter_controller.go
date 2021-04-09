package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"nsure/vote/log"
	"nsure/vote/models"
	"nsure/vote/service"
	"nsure/vote/utils"
)

func GetArbiter(ctx *gin.Context) {
	userId := ctx.Query("userId")
	if len(userId) != 0 {
		userId = utils.Address(userId)
	}
	qualification, account, err := service.GetQualificationAccount(userId)
	if err != nil {
		log.GetLog().Info("GetQualificationAccount", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, newMessageInternalServerError(err))
		return
	}

	if account == nil {
		account = &models.Account{}
	}
	if qualification == nil {
		qualification = &models.Qualification{}
	}

	myReward := service.GetRewardByArbiterId(userId)
	arbiterVo := newArbiterRewardVo(qualification, account, myReward)
	ctx.JSON(http.StatusOK, newMessageDataOK(arbiterVo))
}

func AddArbiter(ctx *gin.Context) {
	var metaMask ArbiterMetamask
	userId, _, err := GetMetamaskMessage(ctx, &metaMask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	req := &metaMask.Content
	if req.UserId != userId {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("req.UserId != userId %v!=%v", req.UserId, userId)))
		return
	}

	if req.Number <= 0 {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid number: %v", req.Number)))
		return
	}

	userId = utils.Address(userId)
	qualifications, account, err := service.ApplyQualification(userId, req.Number)
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	if account == nil || qualifications == nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(fmt.Errorf("invalid account or qualifications")))
		return
	}

	myReward := service.GetRewardByArbiterId(userId)
	arbiterVo := newArbiterRewardVo(qualifications, account, myReward)

	ctx.JSON(http.StatusOK, newMessageDataOK(arbiterVo))
}

func PendingArbiter(ctx *gin.Context) {
	var metaMask ArbiterMetamask
	userId, _, err := GetMetamaskMessage(ctx, &metaMask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	req := &metaMask.Content
	if req.UserId != userId {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("req.UserId != userId %v!=%v", req.UserId, userId)))
		return
	}

	if req.Number <= 0 {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid number: %v", req.Number)))
		return
	}
	userId = utils.Address(userId)
	qualifications, account, err := service.PendingQualifications(userId, req.Number)
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	if account == nil || qualifications == nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(fmt.Errorf("invalid account or qualifications")))
		return
	}
	myReward := service.GetRewardByArbiterId(userId)
	arbiterVo := newArbiterRewardVo(qualifications, account, myReward)

	ctx.JSON(http.StatusOK, newMessageDataOK(arbiterVo))

}

func AddArbiterOld(ctx *gin.Context) {
	var req ArbiterRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}
	if req.Number <= 0 {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid number: %v", req.Number)))
		return
	}
	//arbiterId := GetCurrentUser(ctx)
	qualifications, account, err := service.ApplyQualification(req.UserId, req.Number)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}
	if account == nil || qualifications == nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid account or qualifications")))
		return
	}

	arbiterVo := newArbiterVo(qualifications, account)
	ctx.JSON(http.StatusOK, arbiterVo)
}

func PendingArbiterOld(ctx *gin.Context) {
	var req ArbiterRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}
	if req.Number <= 0 {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid number: %v", req.Number)))
		return
	}
	//arbiterId := GetCurrentUser(ctx)
	qualifications, account, err := service.PendingQualifications(req.UserId, req.Number)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}
	if account == nil || qualifications == nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid account or qualifications")))
		return
	}

	arbiterVo := newArbiterVo(qualifications, account)
	ctx.JSON(http.StatusOK, arbiterVo)
}
