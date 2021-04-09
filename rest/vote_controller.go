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
)

func PlaceVote(ctx *gin.Context) {
	var metaMask VoteMetamask
	userId, signHash, err := GetMetamaskMessage(ctx, &metaMask)
	if err != nil {
		log.GetLog().Info("GetMetamaskMessage", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	req := &metaMask.Content
	log.GetLog().Debug("place claim", zap.String("userId", userId))

	if req.UserId != userId {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("req.UserId != userId %v!=%v", req.UserId, userId)))
		return
	}
	var status common.ClaimStatus
	if req.Status == "Pass" {
		status = common.ClaimStatusPass
	} else {
		status = common.ClaimStatusDeny
	}

	userId = utils.Address(userId)
	ret, err := service.ApplyVote(req.ClaimId, status, userId, req.CoverHash, signHash)
	if err != nil {
		log.GetLog().Info("ApplyVote", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, newMessageDataOK(ret))
}

func PlaceVoteAdmin(ctx *gin.Context) {
	var metaMask VoteMetamask
	userId, signHash, err := GetMetamaskMessage(ctx, &metaMask)
	if err != nil {
		log.GetLog().Info("GetMetamaskMessage", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	req := &metaMask.Content
	log.GetLog().Debug("place claim admin", zap.String("userId", userId))

	if req.UserId != userId {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("req.UserId != userId %v!=%v", req.UserId, userId)))
		return
	}
	var status common.ClaimStatus
	if req.Status == "Pass" {
		status = common.ClaimStatusPass
	} else {
		status = common.ClaimStatusDeny
	}

	userId = utils.Address(userId)
	ret, err := service.ApplyVoteAdmin(req.ClaimId, status, userId, req.CoverHash, signHash)
	if err != nil {
		log.GetLog().Info("ApplyVoteAdmin", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, newMessageDataOK(ret))
}

func PlaceVoteOld(ctx *gin.Context) {
	var req VoteRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}

	var status common.ClaimStatus
	if req.ClaimStatus {
		status = common.ClaimStatusPass
	} else {
		status = common.ClaimStatusDeny
	}
	ret, err := service.ApplyVote(req.ClaimId, status, req.ArbiterId, req.CoverHash, req.SignHash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func PlaceClaimApply(ctx *gin.Context) {
	var metaMask claimApplyMetamask
	userId, _, err := GetMetamaskMessage(ctx, &metaMask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	req := &metaMask.Content
	if req.UserId != userId {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("req.UserId!=userId %v!=%v", req.UserId, userId)))
		return
	}

	userId = utils.Address(req.UserId)
	ret, err := service.ClaimApply(req.ClaimId, userId)
	if err != nil {
		log.GetLog().Debug("ClaimApply", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, newMessageDataOK(ret))
}

func GetClaimApply(ctx *gin.Context) {
	arbiterId := ctx.Query("userId")
	claims, err := service.GetClaimByArbiterId(arbiterId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageInternalServerError(err))
		return
	}

	var claimVos []*claimVo
	for _, claim := range claims {
		claimVos = append(claimVos, newClaimVo(claim))
	}

	if claimVos == nil {
		ctx.JSON(http.StatusOK, newMessageDataOK(""))
	} else {
		ctx.JSON(http.StatusOK, newMessageDataOK(claimVos))
	}

}
