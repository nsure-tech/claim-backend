package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"net/http"
	"nsure/vote/common"
	"nsure/vote/contract"
	"nsure/vote/log"
	"nsure/vote/models"
	"nsure/vote/service"
	"nsure/vote/utils"
	"strconv"
	"time"
)

func PlaceClaim(ctx *gin.Context) {
	var metaMask ClaimMetamask
	userId, _, err := GetMetamaskMessage(ctx, &metaMask)
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

	beginAt, err := time.Parse(time.RFC3339, req.BeginAt)
	if err != nil {
		log.GetLog().Info("req.BeginAt", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid begin_at: %v", err)))
		return
	}
	if req.Period <= 0 {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid period: %v", req.Period)))
		return
	}

	endAt := beginAt.AddDate(0, 0, req.Period)

	submitAt := time.Now()
	if submitAt.After(endAt) {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid submit_at > end_at")))
		return
	}
	if beginAt.After(submitAt) {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid begin_at > submit_at")))
		return
	}

	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		log.GetLog().Info("req.Amount", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("amount %v:%v", req.Amount, err)))
		return
	}

	reward, err := decimal.NewFromString(req.Reward)
	if err != nil {
		log.GetLog().Info("req.Reward", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("reward %v:%v", req.Reward, err)))
		return
	}
	cost := reward

	userId = utils.Address(req.UserId)
	ret, err := service.AddClaim(userId, req.Product, req.CoverId, req.CoverHash, req.Currency, amount, cost, reward, submitAt, beginAt, endAt, req.Desc, req.Cred)
	if err != nil {
		log.GetLog().Info("AddClaim", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, newMessageDataOK(ret))
}

func PlaceClaimChain(ctx *gin.Context) {
	var metaMask ClaimMetamask
	userId, _, err := GetMetamaskMessage(ctx, &metaMask)
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
	buyClaim, err := contract.GetBuyClaim(big.NewInt(utils.StringToInt64(req.CoverId)))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("GetBuyClaim %v", err)))
		return
	}

	beginAt, err := time.Parse(time.RFC3339, req.BeginAt)
	if err != nil {
		log.GetLog().Info("req.BeginAt", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid begin_at: %v", err)))
		return
	}
	if !beginAt.Equal(buyClaim.CreateAt) {
		log.GetLog().Info("req.BeginAt err")
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid begin_at: %v %v", beginAt, buyClaim.CreateAt)))
		return
	}
	if req.Period <= 0 {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid period: %v", req.Period)))
		return
	}
	if uint64(req.Period) != buyClaim.Period {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid period: %v %v", req.Period, buyClaim.Period)))
		return
	}

	endAt := beginAt.AddDate(0, 0, req.Period)

	submitAt := time.Now()
	if submitAt.After(endAt) {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid submit_at > end_at")))
		return
	}
	if beginAt.After(submitAt) {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("invalid begin_at > submit_at")))
		return
	}

	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		log.GetLog().Info("req.Amount", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("amount %v:%v", req.Amount, err)))
		return
	}
	if !amount.Equal(buyClaim.Amount) {
		log.GetLog().Info("req.Amount err", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("amount %v:%v", req.Amount, buyClaim.Amount)))
		return
	}

	userId = utils.Address(req.UserId)
	if buyClaim.Buyer != userId {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("userId error %s != from %s", buyClaim.Buyer, userId)))
		return
	}

	ret, err := service.AddClaim(userId, req.Product, req.CoverId, req.CoverHash, req.Currency, amount, buyClaim.Cost, buyClaim.Reward, submitAt, beginAt, endAt, req.Desc, req.Cred)
	if err != nil {
		log.GetLog().Info("AddClaim", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, newMessageDataOK(ret))
}

func GetClaimListNew(ctx *gin.Context) {
	arbiterId := ctx.Query("userId")
	if len(arbiterId) != 0 {
		arbiterId = utils.Address(arbiterId)
	}
	before := utils.StringToInt64(ctx.Query("before"))
	after := utils.StringToInt64(ctx.Query("after"))
	limit := utils.StringToInt64(ctx.Query("limit"))
	//claims, err := service.GetClaim("", "", "", before, after, int(limit))
	if limit == 0 && after > before {
		limit = after - before
	}
	claims, err := service.GetClaimList("", "", "", int(before), int(limit))
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}

	var claimLists []*claimList
	for _, claim := range claims {
		claimLists = append(claimLists, newClaimList(claim))
	}
	qualification, err := service.GetQualificationByArbiterId(arbiterId)
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	if qualification == nil {
		qualification = &models.Qualification{}
	}

	total, err := service.GetClaimTotal()
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}

	listClaim := &claimListArbiter{
		ActiveClaim:   common.ClaimActive,
		ClosedClaim:   common.ClaimClosed,
		MyActiveClaim: qualification.Used,
		MyHonors:      qualification.Closed,
		Total:         total,
		ClaimList:     claimLists,
	}

	ctx.JSON(http.StatusOK, newMessageDataOK(listClaim))
}

func GetClaimListAdmin(ctx *gin.Context) {
	arbiterId := ctx.Query("userId")
	if len(arbiterId) != 0 {
		arbiterId = utils.Address(arbiterId)
	}

	claims, err := service.GetClaimByEndVote()
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}

	var claimLists []*claimList
	for _, claim := range claims {
		claimLists = append(claimLists, newClaimList(claim))
	}
	qualification, err := service.GetQualificationByArbiterId(arbiterId)
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	if qualification == nil {
		qualification = &models.Qualification{}
	}

	listClaim := &claimListArbiter{
		ActiveClaim:   common.ClaimActive,
		ClosedClaim:   common.ClaimClosed,
		MyActiveClaim: qualification.Used,
		MyHonors:      qualification.Closed,
		Total:         0,
		ClaimList:     claimLists,
	}

	ctx.JSON(http.StatusOK, newMessageDataOK(listClaim))
}

func GetClaimListAdminNew(ctx *gin.Context) {
	arbiterId := ctx.Query("userId")
	if len(arbiterId) != 0 {
		arbiterId = utils.Address(arbiterId)
	}

	claims, err := service.GetClaimByEndVote()
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	var applyLists []*applyList
	for _, claim := range claims {
		applyLists = append(applyLists, newApplyListByClaim(claim))
	}

	qualification, err := service.GetQualificationByArbiterId(arbiterId)
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	if qualification == nil {
		qualification = &models.Qualification{}
	}

	listApply := &applyListArbiter{
		ActiveClaim:   common.ClaimActive,
		ClosedClaim:   common.ClaimClosed,
		MyActiveClaim: qualification.Used,
		MyHonors:      qualification.Closed,
		Total:         0,
		ApplyList:     applyLists,
	}

	ctx.JSON(http.StatusOK, newMessageDataOK(listApply))
}

func GetClaimListDown(ctx *gin.Context) {
	arbiterId := ctx.Query("userId")
	if len(arbiterId) == 0 {
		log.GetLog().Error("error: userId is null")
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("error: userId is null")))
		return
	} else {
		arbiterId = utils.Address(arbiterId)
	}
	before := utils.StringToInt64(ctx.Query("before"))
	after := utils.StringToInt64(ctx.Query("after"))
	limit := utils.StringToInt64(ctx.Query("limit"))
	if limit == 0 && after > before {
		limit = after - before
	}

	applies, err := service.GetApplyList(arbiterId, "", int(before), int(limit))
	if err != nil {
		log.GetLog().Error("service GetApplyList")
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	total, err := service.GetApplyCount(arbiterId, "")
	if err != nil {
		log.GetLog().Error("service GetApplyCount")
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}

	var applyLists []*applyList
	for _, apply := range applies {
		applyLists = append(applyLists, newApplyList(apply))
	}
	qualification, err := service.GetQualificationByArbiterId(arbiterId)
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	if qualification == nil {
		qualification = &models.Qualification{}
	}

	listApply := &applyListArbiter{
		ActiveClaim:   common.ClaimActive,
		ClosedClaim:   common.ClaimClosed,
		MyActiveClaim: qualification.Used,
		MyHonors:      qualification.Closed,
		Total:         total,
		ApplyList:     applyLists,
	}

	ctx.JSON(http.StatusOK, newMessageDataOK(listApply))

}

func GetClaimList(ctx *gin.Context) {
	userId := ctx.Query("userId")
	product := ctx.Query("product")
	claims, err := service.GetClaim(userId, product, "", 0, 0, 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageInternalServerError(err))
		return
	}

	var claimLists []*claimList
	for _, claim := range claims {
		claimLists = append(claimLists, newClaimList(claim))
	}

	if claimLists == nil {
		ctx.JSON(http.StatusOK, newMessageDataOK(""))
	} else {
		ctx.JSON(http.StatusOK, newMessageDataOK(claimLists))
	}
}

func GetClaimListByArbiter(ctx *gin.Context) {
	userId := ctx.Query("userId")
	product := ctx.Query("product")
	votes, qualification, err := service.GetClaimByArbiter(userId, product, "", 0, 0, 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageInternalServerError(err))
		return
	}

	var claimLists []*claimList
	for _, vote := range votes {
		claimLists = append(claimLists, newClaimListVote(vote))
	}

	var listClaim = &claimListArbiter{
		ActiveClaim: common.ClaimActive,
		ClosedClaim: common.ClaimClosed,
		ClaimList:   claimLists,
	}
	if qualification != nil {
		listClaim.MyActiveClaim = qualification.Used
		listClaim.MyHonors = qualification.Closed
	}

	ctx.JSON(http.StatusOK, newMessageDataOK(listClaim))
}

func GetClaimVote(ctx *gin.Context) {
	strClaimId := ctx.Query("claimId")
	claimId := utils.StringToInt64(strClaimId)
	if claimId == 0 {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("claimId error:%v", strClaimId)))
		return
	}
	claim, err := service.GetClaimById(claimId)
	if err != nil {
		log.GetLog().Debug("service.GetClaimVote", zap.Error(err))
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	if claim == nil {
		log.GetLog().Debug("service.GetClaimVote claim nil")
		ctx.JSON(http.StatusOK, newMessageInternalServerError(fmt.Errorf("claim nil where id %v", strClaimId)))
		return
	}

	ctx.JSON(http.StatusOK, newMessageDataOK(newClaimVote(claim)))
}

func GetClaimResult(ctx *gin.Context) {
	claimId := ctx.Query("claimId")
	iClaimId, err := strconv.Atoi(claimId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(err))
		return
	}
	claim, votes, err := service.GetClaimResult(int64(iClaimId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageInternalServerError(err))
		return
	}
	if claim == nil {
		ctx.JSON(http.StatusInternalServerError, newMessageInternalServerError(err))
		return
	}

	claimVotes := newClaimVoteResult(claim, votes)
	ctx.JSON(http.StatusOK, newMessageDataOK(claimVotes))
}

func PlaceClaimOld(ctx *gin.Context) {
	var req ClaimRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}
	amount, _ := decimal.NewFromString(req.Amount)
	cost, _ := decimal.NewFromString(req.Cost)
	reward, _ := decimal.NewFromString(req.Reward)
	beginAt, err := time.Parse(time.RFC3339, req.CoverBeginAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid begin_at: %v", err)))
		return
	}
	endAt, err := time.Parse(time.RFC3339, req.CoverEndAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid end_at: %v", err)))
		return
	}
	if beginAt.After(endAt) {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid begin_at > end_at")))
		return
	}

	submitAt := time.Now()
	if submitAt.After(endAt) {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid submit_at > end_at")))
		return
	}
	if beginAt.After(submitAt) {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid begin_at > submit_at")))
		return
	}
	//if *GetCurrentUser(ctx) != req.UserId {
	//	ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid user: #{req.UserId}")))
	//	return
	//}
	ret, err := service.AddClaim(req.UserId, req.Product, req.CoverId, req.CoverHash, req.Currency, amount, cost, reward, submitAt, beginAt, endAt, req.Desc, req.Cred)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func GetClaimVoteOld(ctx *gin.Context) {
	claimId := ctx.Query("claimId")
	iClaimId, err := strconv.Atoi(claimId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}
	claim, err := service.GetClaimById(int64(iClaimId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}

	ctx.JSON(http.StatusOK, newClaimVote(claim))
}

func GetClaimByApply(ctx *gin.Context) {
	claims, err := service.GetClaimByApply()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}

	var claimVos []*claimVo
	for _, claim := range claims {
		claimVos = append(claimVos, newClaimVo(claim))
	}

	ctx.JSON(http.StatusOK, claimVos)
}

func GetClaimByArbiterId(ctx *gin.Context) {
	arbiterId := ctx.Query("userId")
	claims, err := service.GetClaimByArbiterId(arbiterId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}

	var claimVos []*claimVo
	for _, claim := range claims {
		claimVos = append(claimVos, newClaimVo(claim))
	}

	ctx.JSON(http.StatusOK, claimVos)
}

func ApplyClaimByArbiterId(ctx *gin.Context) {
	arbiterId := ctx.Query("userId")
	var req claimVo
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}
	ret, err := service.ClaimApply(req.ClaimId, arbiterId)
	if err != nil {
		ctx.JSON(http.StatusOK, newMessageVo(err))
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
