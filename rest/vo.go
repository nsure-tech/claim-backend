package rest

import (
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/shopspring/decimal"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/utils"
	"time"
)

type messageVo struct {
	Message string `json:"message"`
}

func newMessageVo(error error) *messageVo {
	return &messageVo{
		Message: error.Error(),
	}
}

type ClaimRequestMessage struct {
	UserId    string `json:"user_id"`
	Product   string `json:"product"`
	CoverId   string `json:"cover_id"`
	CoverHash string `json:"cover_hash"`
	Currency  string `json:"currency"`
	Amount    string `json:"amount"`
	Reward    string `json:"reward"` // 投保仲裁回报的金额
	BeginAt   string `json:"begin_at"`
	Period    int    `json:"period"`
	Desc      string `json:"desc"`
	Cred      string `json:"cred"`
	Loss      string `json:"loss"`
}

type ClaimMetamask struct {
	Method  string              `json:"method"`
	Content ClaimRequestMessage `json:"content"`
}

type ArbiterRequest struct {
	UserId string `json:"user_id"`
	Number int    `json:"number"`
}
type ArbiterMetamask struct {
	Method  string         `json:"method"`
	Content ArbiterRequest `json:"content"`
}

type ChaRequest struct {
	ClaimId     int64  `json:"claim_id"`
	ChallengeId string `json:"user_id"`
	CoverHash   string `json:"cover_hash"`
}

type ChallengeMetamask struct {
	Method  string     `json:"method"`
	Content ChaRequest `json:"content"`
}

type ChaVoteRequest struct {
	UserId  string `json:"user_id"`
	ClaimId int64  `json:"claim_id"`
	Status  string `json:"status"`
}

type ChaVoteMetamask struct {
	Method  string         `json:"method"`
	Content ChaVoteRequest `json:"content"`
}

type VoteRequestMessage struct {
	UserId    string `json:"user_id"`
	ClaimId   int64  `json:"claim_id"`
	CoverHash string `json:"cover_hash"`
	Status    string `json:"status"`
}

type VoteMetamask struct {
	Method  string             `json:"method"`
	Content VoteRequestMessage `json:"content"`
}

type WithDrawRequestMessage struct {
	UserId   string `json:"account"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	Nonce    string `json:"nonce"`
}

type WithdrawMetamask struct {
	Method  string                 `json:"method"`
	Content WithDrawRequestMessage `json:"content"`
}

type VoteRequest struct {
	ClaimId     int64  `json:"claim_id"`
	ClaimStatus bool   `json:"status"`
	ArbiterId   string `json:"arbiter_id"`
	SignHash    string `json:"sign_hash"`
	CoverHash   string `json:"cover_hash"`
}

type claimApply struct {
	UserId  string `json:"user_id"`
	ClaimId int64  `json:"claim_id"`
	Product string `json:"product"`
}
type claimApplyMetamask struct {
	Method  string     `json:"method"`
	Content claimApply `json:"content"`
}

type ClaimApply struct {
	ClaimId int64  `json:"claim_id"`
	Product string `json:"product"`
}

type ChallengeVoteRequest struct {
	ClaimId int64 `json:"claim_id"`
	Status  bool  `json:"status"`
}

type ChallengeRequest struct {
	ClaimId     int64  `json:"claim_id"`
	ChallengeId string `json:"user_id"`
}

type BillRequest struct {
	UserId   string  `json:"user_id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type LoginRequest struct {
	UserId string         `json:"user_id"`
	SigHex string         `json:"sig_hex"`
	Msg    core.TypedData `json:"msg"`
}

type ClaimRequest struct {
	UserId       string `json:"user_id"`
	Product      string `json:"product"`
	CoverId      string `json:"cover_id"`
	CoverHash    string `json:"cover_hash"`
	Currency     string `json:"currency"`
	Amount       string `json:"amount"`
	Cost         string `json:"cost"`
	Reward       string `json:"reward"`
	SubmitAt     string `json:"claim_begin_at"`
	SubmitEndAt  string `json:"claim_end_at"`
	CoverBeginAt string `json:"cover_begin_at"`
	CoverEndAt   string `json:"cover_end_at"`
	Desc         string `json:"desc"`
	Cred         string `json:"cred"`
	Loss         string `json:"loss"`
}
type claimVote struct {
	ClaimId int64  `json:"claim_id"`
	Status  string `json:"status"`
	ClaimRequest
}

func newClaimVote(claim *models.Claim) *claimVote {
	claimVo := &claimVote{
		claim.Id,
		string(claim.Status),
		ClaimRequest{
			UserId:       claim.UserId,
			Product:      claim.Product,
			CoverId:      claim.CoverId,
			CoverHash:    claim.CoverHash,
			Currency:     claim.Currency,
			Amount:       utils.DToString(claim.Amount),
			Cost:         utils.DToString(claim.Cost),
			Reward:       utils.DToString(claim.Reward),
			SubmitAt:     claim.SubmitAt.Format(time.RFC3339),
			SubmitEndAt:  claim.SubmitAt.Add(time.Duration(common.ClaimMinute) * time.Minute).String(),
			CoverBeginAt: claim.CoverBeginAt.Format(time.RFC3339),
			CoverEndAt:   claim.CoverEndAt.Format(time.RFC3339),
			Desc:         claim.Description,
			Cred:         claim.Credential,
			Loss:         claim.Loss,
		},
	}

	return claimVo
}

type voteResult struct {
	ArbiterId string `json:"arbiter_id"`
	Status    string `json:"status"`
	SignHash  string `json:"sign_hash"`
}
type claimVoteResult struct {
	Claim *claimVote    `json:"claim"`
	Vote  []*voteResult `json:"result"`
}

func newClaimVoteResult(claim *models.Claim, votes []*models.Vote) *claimVoteResult {
	claimVote := newClaimVote(claim)
	claimVoteResult := &claimVoteResult{
		Claim: claimVote,
	}
	for _, vote := range votes {
		result := &voteResult{
			ArbiterId: vote.ArbiterId,
			Status:    string(vote.Status),
			SignHash:  vote.SignHash,
		}
		claimVoteResult.Vote = append(claimVoteResult.Vote, result)
	}
	return claimVoteResult
}

type claimList struct {
	ClaimId      int64  `json:"claim_id"`
	UserId       string `json:"user_id"`
	Product      string `json:"product"`
	CoverId      string `json:"cover_id"`
	CoverHash    string `json:"cover_hash"`
	Currency     string `json:"currency"`
	Amount       string `json:"amount"`
	SubmitAt     string `json:"submit_at"`
	ArbiterAt    string `json:"arbiter_at"`
	CoverBeginAt string `json:"begin_at"`
	CoverEndAt   string `json:"end_at"`
	Status       string `json:"status"`
	Notes        string `json:"notes"`
}

func newClaimList(claim *models.Claim) *claimList {
	return &claimList{
		ClaimId:      claim.Id,
		UserId:       claim.UserId,
		Product:      claim.Product,
		CoverId:      claim.CoverId,
		CoverHash:    claim.CoverHash,
		Currency:     claim.Currency,
		Amount:       utils.DToString(claim.Amount),
		SubmitAt:     claim.SubmitAt.Format(time.RFC3339),
		ArbiterAt:    claim.ArbiterAt.Format(time.RFC3339),
		CoverBeginAt: claim.CoverBeginAt.Format(time.RFC3339),
		CoverEndAt:   claim.CoverEndAt.Format(time.RFC3339),
		Status:       string(claim.Status),
		Notes:        claim.Notes,
	}
}

type claimListArbiter struct {
	ActiveClaim   int          `json:"active_claim"`
	MyActiveClaim int          `json:"my_active_claim"`
	ClosedClaim   int          `json:"closed_claim"`
	MyHonors      int          `json:"my_honors"`
	Total         int          `json:"total"`
	ClaimList     []*claimList `json:"claim_list"`
}

type applyList struct {
	ClaimId     int64  `json:"claim_id"`
	Product     string `json:"product"`
	UserId      string `json:"user_id"`
	CoverId     string `json:"cover_id"`
	CoverHash   string `json:"cover_hash"`
	SubmitAt    string `json:"submit_at"`
	EndAt       string `json:"end_at"`
	ResidueTime string `json:"residue_time"`
	Status      string `json:"status"`
}

type applyListArbiter struct {
	ActiveClaim   int          `json:"active_claim"`
	MyActiveClaim int          `json:"my_active_claim"`
	ClosedClaim   int          `json:"closed_claim"`
	MyHonors      int          `json:"my_honors"`
	Total         int          `json:"total"`
	ApplyList     []*applyList `json:"apply_list"`
}

func newApplyList(apply *models.Apply) *applyList {
	residua := "0s"
	endAt := time.Now()
	if apply.Status == common.ApplyStatusApply {
		endAt = apply.SubmitAt.Add(time.Duration(common.ClaimMinute) * time.Minute)
		if dur := time.Until(endAt); dur > 0 {
			residua = dur.String()
		}
	} else if apply.Status == common.ApplyStatusSuccess {
		endAt = apply.ArbiterAt.Add(time.Duration(common.VoteMinute) * time.Minute)
		if dur := time.Until(endAt); dur > 0 {
			residua = dur.String()
		}
	}

	return &applyList{
		ClaimId:     apply.ClaimId,
		Product:     apply.Product,
		CoverId:     apply.CoverId,
		UserId:      apply.UserId,
		CoverHash:   apply.CoverHash,
		SubmitAt:    apply.SubmitAt.Format(time.RFC3339),
		EndAt:       endAt.Format(time.RFC3339),
		ResidueTime: residua,
		Status:      string(apply.Status),
	}
}

func newApplyListByClaim(claim *models.Claim) *applyList {
	residua := "0s"
	endAt := time.Now()
	if claim.Status == common.ClaimStatusArbiter {
		endAt = claim.ArbiterAt.Add(time.Duration(common.VoteMinute) * time.Minute)
		if dur := time.Until(endAt); dur > 0 {
			residua = dur.String()
		}
	}

	return &applyList{
		ClaimId:     claim.Id,
		Product:     claim.Product,
		CoverId:     claim.CoverId,
		UserId:      claim.UserId,
		CoverHash:   claim.CoverHash,
		SubmitAt:    claim.SubmitAt.Format(time.RFC3339),
		EndAt:       endAt.Format(time.RFC3339),
		ResidueTime: residua,
		Status:      string(claim.Status),
	}
}

func newClaimListVote(vote *models.Vote) *claimList {
	return &claimList{
		ClaimId:      vote.Id,
		UserId:       vote.UserId,
		Product:      vote.Product,
		CoverId:      vote.CoverId,
		CoverHash:    vote.CoverHash,
		Currency:     vote.Currency,
		Amount:       utils.DToString(vote.Amount),
		SubmitAt:     vote.SubmitAt.Format(time.RFC3339),
		ArbiterAt:    vote.ArbiterAt.Format(time.RFC3339),
		CoverBeginAt: vote.CoverBeginAt.Format(time.RFC3339),
		Status:       string(vote.Status),
		Notes:        vote.Notes,
	}
}

type claimVo struct {
	ClaimId int64  `json:"claim_id"`
	Product string `json:"product"`
}

func newClaimVo(claim *models.Claim) *claimVo {
	return &claimVo{
		ClaimId: claim.Id,
		Product: claim.Product,
	}
}

type AccountVo struct {
	UserId   string  `json:"user_id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

func newAccountVo(account *models.Account) *AccountVo {
	return &AccountVo{
		UserId:   account.UserId,
		Currency: account.Currency,
		Amount:   utils.DToF64(account.Available.Add(account.Hold)),
	}
}

type arbiterVo struct {
	ArbiterId  string `json:"user_id"`
	Available  string `json:"available"`
	Hold       string `json:"hold"`
	ArbiterSum int    `json:"arbiter_sum"`
	Pending    int    `json:"pending"`
}

func newArbiterVo(qualification *models.Qualification, account *models.Account) *arbiterVo {
	return &arbiterVo{
		ArbiterId:  account.UserId,
		Available:  utils.DToString(account.Available),
		Hold:       utils.DToString(account.Hold),
		ArbiterSum: qualification.Available + qualification.Used,
		Pending:    qualification.Pending,
	}
}

type arbiterRewardVo struct {
	ArbiterId        string `json:"user_id"`
	Available        string `json:"available"`
	Hold             string `json:"hold"`
	ArbiterAvailable int    `json:"arbiter_sum"`
	Pending          int    `json:"pending"`
	TotalReward      string `json:"total_reward"`
	MyReward         string `json:"my_reward"`
}

func newArbiterRewardVo(qualification *models.Qualification, account *models.Account, myReward decimal.Decimal) *arbiterRewardVo {
	return &arbiterRewardVo{
		ArbiterId:        account.UserId,
		Available:        utils.DToString(account.Available),
		Hold:             utils.DToString(account.Hold),
		ArbiterAvailable: qualification.Available,
		Pending:          qualification.Pending + qualification.Used,
		TotalReward:      utils.DToString(common.RewardTotal),
		MyReward:         utils.DToString(myReward),
	}
}

type billVo struct {
	Currency  string `json:"asset"`
	Available string `json:"amount"`
	Date      string `json:"date"`
	Action    string `json:"action"`
}

type AssetVo struct {
	Total    int       `json:"total"`
	BillList []*billVo `json:"asset_list"`
}

func newBillVo(bill *models.Bill) *billVo {
	return &billVo{
		Currency:  bill.Currency,
		Available: bill.Available.String(),
		Date:      bill.CreatedAt.Format(time.RFC3339),
		Action:    string(bill.Type),
	}
}

type ChaAmountVo struct {
	ChaAmount string `json:"chaAmount"`
	Available string `json:"available"`
}
