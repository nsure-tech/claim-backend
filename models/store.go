package models

import (
	"nsure/vote/common"
)

type Store interface {
	BeginTx() (Store, error)
	Rollback() error
	CommitTx() error

	AddMetamask(metamask *Metamask) error

	GetAccount(userId string, currency string) (*Account, error)
	GetAccountsByArbiterId(userId string) ([]*Account, error)
	GetAccountForUpdate(userId string, currency string) (*Account, error)
	AddAccount(account *Account) error
	UpdateAccount(account *Account) error

	AddBills(bills []*Bill) error
	AddBill(bill *Bill) error
	GetUnsettledBills(count int) ([]*Bill, error)
	UpdateBill(bill *Bill) error
	GetBillsCountByUserId(userId string, statuses []common.BillType) (int, error)
	GetBillsByUserId(userId string, statuses []common.BillType, offset, limit int) ([]*Bill, error)

	AddQualification(qualifications *Qualification) error
	GetQualificationByArbiterId(arbiterId string) (*Qualification, error)
	GetQualificationForUpdate(arbiterId string) (*Qualification, error)
	UpdateQualification(qualifications *Qualification) error

	AddPending(pending *Pending) error
	GetPendingByArbiterId(arbiterId string) (*Pending, error)
	GetPendingForUpdate(arbiterId string) (*Pending, error)
	UpdatePending(pending *Pending) error
	GetUnsettledPends(pendingTime uint, count int) ([]*Pending, error)

	AddApply(apply *Apply) error
	UpdateApply(apply *Apply) error
	GetApplyCountByArbiterId(arbiterId string, status common.ApplyStatus) (int, error)
	GetApplyListByArbiterId(arbiterId string, status common.ApplyStatus, begin, limit int) ([]*Apply, error)
	GetApplyByArbiterId(arbiterId string, status *common.ApplyStatus) ([]*Apply, error)
	GetApplyByClaimId(claimId int64, status *common.ApplyStatus) ([]*Apply, error)
	GetApplyByApplyNum(applyNum uint8, status *common.ApplyStatus) ([]*Apply, error)
	GetApplyForUpdate(claimId int64, arbiterId string) (*Apply, error)

	AddVote(vote *Vote) error
	UpdateVote(vote *Vote) error
	GetVoteByClaimId(claimId int64) ([]*Vote, error)
	GetVoteForUpdate(claimId int64, arbiterId string) (*Vote, error)
	GetVote(arbiterId string, product string, status common.ClaimStatus, beforeId, afterId int64, limit int) ([]*Vote, error)
	GetVoteByVoteNum(statuses []common.ClaimStatus, voteNum uint8) ([]*common.ClaimId, error)
	GetVoteByClaimIdStatus(claimId int64, statuses []common.ClaimStatus) ([]*Vote, error)

	AddVoteFill(voteFill *VoteFill) error
	UpdateVoteFill(voteFill *VoteFill) error
	GetVoteFillsByClaimId(claimId int64) ([]*VoteFill, error)
	GetUnsettledVoteFills(count int) ([]*VoteFill, error)

	AddReward(reward *Reward) error
	UpdateReward(reward *Reward) error
	GetRewardByEnd(endTime uint) ([]*Reward, error)
	GetRewardForUpdate(id int64) (*Reward, error)
	GetUnsettledRewards(count int) ([]*Reward, error)
	GetRewardByArbiterId(arbiterId string) (*common.Reward, error)
	GetRewardTotal() (*common.Reward, error)

	AddRewardCha(rewardCha *RewardCha) error
	UpdateRewardCha(rewardCha *RewardCha) error
	GetRewardChaByEnd(endTime uint) ([]*RewardCha, error)
	GetRewardChaForUpdate(id int64) (*RewardCha, error)
	AddPunishCha(punishCha *PunishCha) error
	UpdatePunishCha(punishCha *PunishCha) error
	GetPunishChaByEnd(endTime uint) ([]*PunishCha, error)
	GetPunishChaForUpdate(id int64) (*PunishCha, error)

	AddPayment(payment *Payment) error
	UpdatePayment(payment *Payment) error
	GetPaymentByEnd(endTime uint) ([]*Payment, error)
	GetPaymentForUpdate(id int64) (*Payment, error)
	GetUnsettledPayments(count int) ([]*Payment, error)

	GetClaim(userId string, product string, status common.ClaimStatus, beforeId, afterId int64, limit int) ([]*Claim, error)
	GetClaimList(userId string, product string, status common.ClaimStatus, offset, limit int) ([]*Claim, error)
	GetClaimByApply(status common.ClaimStatus, applyTime uint, applyNum uint) ([]*Claim, error)
	GetClaimByEndApply(status common.ClaimStatus, applyTime uint) ([]*Claim, error)
	GetClaimByEndVote(status common.ClaimStatus, voteTime uint) ([]*Claim, error)
	GetClaimClose(statuses []common.ClaimStatus, paymentTime uint) ([]*Claim, error)
	GetClaimForUpdate(claimId int64) (*Claim, error)
	GetClaimByUserId(userId string) ([]*Claim, error)
	GetClaimById(claimId int64) (*Claim, error)
	GetClaimByHash(coverHash string) (*Claim, error)
	UpdateClaim(claim *Claim) error
	AddClaim(claim *Claim) error
	GetClaimTotal() (int, error)
	GetClaimCount(settled bool) (int, error)

	AddChallenge(challenge *Challenge) error
	UpdateChallenge(challenge *Challenge) error
	GetChallengeForUpdate(claimId int64) (*Challenge, error)
	GetChallengeByStatus(status common.ChallengeStatus, count int) ([]*Challenge, error)

	AddChallengeFill(challengeFill *ChallengeFill) error
	UpdateChallengeFill(challengeFill *ChallengeFill) error
	GetChallengeFillForUpdate(claimId int64) (*ChallengeFill, error)
	GetChallengeFillByStatus(status common.ChallengeStatus, count int) ([]*ChallengeFill, error)

	GetConfig(keyWord string) (*Config, error)
	GetConfigForUpdate(keyWord string) (*Config, error)
	UpdateConfig(config *Config) error

	AddTransfer(transfer *Transfer) error
	UpdateTransfer(transfer *Transfer) error

	GetWithdrawByUserNonce(userId string, nonce uint64) (*Withdraw, error)
	GetWithdrawForUpdate(userId string, nonce uint64) (*Withdraw, error)
	GetWithdrawsByEndAt(status common.WithdrawStatus, endTime uint) ([]*Withdraw, error)
	AddWithdraw(withdraw *Withdraw) error
	UpdateWithdraw(withdraw *Withdraw) error
	AddChainClaim(chainClaim *ChainClaim) error
	UpdateChainClaim(chainClaim *ChainClaim) error
}
