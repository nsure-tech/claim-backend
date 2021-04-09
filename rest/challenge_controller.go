package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"nsure/vote/common"
	"nsure/vote/log"
	"nsure/vote/service"
	"nsure/vote/utils"
	"strconv"
)

func ApplyChallenge(ctx *gin.Context) {
	var metaMask ChallengeMetamask
	userId, _, err := GetMetamaskMessage(ctx, &metaMask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	req := &metaMask.Content
	if req.ChallengeId != userId {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("req.UserId != userId %v!=%v", req.ChallengeId, userId)))
		return
	}

	userId = utils.Address(userId)
	ret, err := service.ApplyChallenge(userId, req.CoverHash, req.ClaimId)
	if err != nil {
		log.GetLog().Info("ApplyChallenge", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	if ret {
		ctx.JSON(http.StatusOK, newMessageOK())
	} else {
		ctx.JSON(http.StatusOK, newMessageFail())
	}
}

func ChallengeVote(ctx *gin.Context) {
	var metaMask ChaVoteMetamask
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
	var status common.ChallengeStatus
	if req.Status == "true" {
		status = common.ChallengeStatusSuccess
	} else {
		status = common.ChallengeStatusFail
	}

	adminId := utils.Address(userId)
	ret, err := service.ChallengeVoteByAdmin(adminId, req.ClaimId, status)
	if err != nil {
		log.GetLog().Info("ChallengeVoteByUserId", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}

	if ret {
		ctx.JSON(http.StatusOK, newMessageOK())
	} else {
		ctx.JSON(http.StatusOK, newMessageFail())
	}
}

func ApplyChallengeOld(ctx *gin.Context) {
	var req ChallengeRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}

	ret, err := service.ApplyChallengeOld(req.ChallengeId, req.ClaimId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
func GetChallengeChaAmount(ctx *gin.Context) {
	userId := ctx.Query("userId")
	strClaimId := ctx.Query("claimId")
	claimId := utils.StringToInt64(strClaimId)
	if claimId == 0 {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("claimId error:%v", strClaimId)))
		return
	}

	claim, err := service.GetClaimById(claimId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	account, err := service.GetBalanceByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	chaAmount, err := utils.ChallengeNSure(claim.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	chaAmountVo := &ChaAmountVo{
		ChaAmount: chaAmount.String(),
		Available: account.Available.String(),
	}
	ctx.JSON(http.StatusOK, newMessageDataOK(chaAmountVo))
}

func GetChallengePrice(ctx *gin.Context) {
	strClaimId := ctx.Query("claimId")
	claimId, err := strconv.Atoi(strClaimId)
	claim, err := service.GetClaimById(int64(claimId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	chaAmount, err := utils.ChallengeNSure(claim.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}

	ctx.JSON(http.StatusOK, newMessageDataOK(chaAmount))
}

func ChallengeVoteOld(ctx *gin.Context) {
	var req ChallengeVoteRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}

	ret, err := service.ChallengeVote(req.ClaimId, req.Status)
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageVo(err))
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
