package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"net/http"
	"nsure/vote/common"
	"nsure/vote/log"
	"nsure/vote/service"
	"nsure/vote/utils"
)

func init() {
	common.ClaimActive, _ = service.GetClaimCount(false)
	common.ClaimClosed, _ = service.GetClaimCount(true)
	common.RewardTotal = service.GetRewardTotal()
	common.VoteAddress = make(map[string]struct{})
	for i := 1; i <= common.AdminAddressNum; i++ {
		addressKey := common.VoteAddressPrefix + utils.IntToA(i)
		if val, err := service.GetConfig(addressKey); err == nil && len(val) > 0 {
			address := utils.Address(val)
			if len(address) != 0 {
				common.VoteAddress[address] = struct{}{}
			}
		}
	}

	common.ChallengeAddress = make(map[string]struct{})
	for i := 1; i <= common.AdminAddressNum; i++ {
		addressKey := common.ChallengeAddressPrefix + utils.IntToA(i)
		if val, err := service.GetConfig(addressKey); err == nil && len(val) > 0 {
			address := utils.Address(val)
			if len(address) != 0 {
				common.ChallengeAddress[address] = struct{}{}
			}
		}
	}

	common.PaymentAddress = make(map[string]struct{})
	for i := 1; i <= common.AdminAddressNum; i++ {
		addressKey := common.PaymentAddressPrefix + utils.IntToA(i)
		if val, err := service.GetConfig(addressKey); err == nil && len(val) > 0 {
			address := utils.Address(val)
			if len(address) != 0 {
				common.PaymentAddress[address] = struct{}{}
			}
		}
	}
}

func GetMetamaskMessage(ctx *gin.Context, val interface{}) (string, string, error) {
	var req MetamaskRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		log.GetLog().Info("BindJSON", zap.Error(err))
		return "", "", err
	}
	msg, err := json.Marshal(req.Msg)
	if err != nil {
		log.GetLog().Info("Marshal", zap.Error(err))
		return "", "", err
	}
	service.AddMetamask(req.UserId, req.SigHex, string(msg))
	message, err := service.GetTypedDataMessage(req.UserId, req.SigHex, req.Msg)
	if err != nil {
		log.GetLog().Info("GetTypedDataMessage", zap.Error(err))
		return "", "", err
	}
	return req.UserId, req.SigHex, utils.MessageToStruct(message, val)
}

func AssetHistory(ctx *gin.Context) {
	userId := ctx.Query("userId")
	if len(userId) == 0 {
		log.GetLog().Error("error: userId is null")
		ctx.JSON(http.StatusBadRequest, newMessageBadRequest(fmt.Errorf("error: userId is null")))
		return
	} else {
		userId = utils.Address(userId)
	}
	offset := utils.StringToInt64(ctx.Query("offset"))
	limit := utils.StringToInt64(ctx.Query("limit"))
	//statuses := []common.BillType{common.BillTypeWithdraw, common.BillTypeDeposit}
	var statuses []common.BillType

	bills, err := service.GetBillsByUserId(userId, statuses, int(offset), int(limit))
	if err != nil {
		log.GetLog().Error("service GetBillsByUserId")
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	total, err := service.GetBillsCountByUserId(userId, statuses)
	if err != nil {
		log.GetLog().Error("service GetBillsCountByUserId")
		ctx.JSON(http.StatusOK, newMessageInternalServerError(err))
		return
	}
	var billVos []*billVo
	for _, bill := range bills {
		billVos = append(billVos, newBillVo(bill))
	}
	ret := &AssetVo{
		Total:    total,
		BillList: billVos,
	}
	ctx.JSON(http.StatusOK, newMessageDataOK(ret))
}

func Deposit(ctx *gin.Context) {
	var req BillRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}

	amount := decimal.NewFromFloat(req.Amount)
	if amount.LessThanOrEqual(decimal.Zero) {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("amount less than 0: #{req.Amount}")))
		return
	}

	//if *GetCurrentUser(ctx) != req.UserId {
	//	ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid user: #{req.UserId}")))
	//	return
	//}
	ret, err := service.AccountDeposit(req.UserId, req.Currency, amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}

	ctx.JSON(http.StatusOK, newAccountVo(ret))
}

func Withdraw(ctx *gin.Context) {
	var req BillRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newMessageVo(err))
		return
	}

	amount := decimal.NewFromFloat(req.Amount)
	if amount.LessThanOrEqual(decimal.Zero) {
		ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("amount less than 0: #{req.Amount}")))
		return
	}

	//if *GetCurrentUser(ctx) != req.UserId {
	//	ctx.JSON(http.StatusBadRequest, newMessageVo(fmt.Errorf("invalid user: #{req.UserId}")))
	//	return
	//}
	ret, err := service.AccountWithdraw(req.UserId, req.Currency, amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}

	ctx.JSON(http.StatusOK, newAccountVo(ret))
}
