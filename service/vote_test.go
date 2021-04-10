package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/utils"
	"testing"
	"time"
)

var arbiters []string
var beforeAccount map[string]*models.Account
var beforeQualification map[string]*models.Qualification
var afterAccount map[string]*models.Account
var afterQualification map[string]*models.Qualification

func init() {
	arbiters = []string{"test01", "test02", "test03"}
	beforeAccount = make(map[string]*models.Account)
	beforeQualification = make(map[string]*models.Qualification)
	afterAccount = make(map[string]*models.Account)
	afterQualification = make(map[string]*models.Qualification)

	common.ChallengeAddress = make(map[string]struct{})
	fmt.Println("INIT")
	for i := 1; i <= common.AdminAddressNum; i++ {
		addressKey := common.ChallengeAddressPrefix + utils.IntToA(i)
		if val, err := GetConfig(addressKey); err == nil && len(val) > 0 {
			address := utils.Address(val)
			if len(address) != 0 {
				common.ChallengeAddress[address] = struct{}{}
			}
		}
	}
}

func TestVote2_1(t *testing.T) {
	claimId := int64(27)
	tGetAccountQua(beforeAccount, beforeQualification)

	claim, err := tApply(claimId)
	if err != nil {
		return
	}

	if _, err = ApplyVote(claimId, common.ClaimStatusPass, arbiters[0], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[0], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusPass, arbiters[1], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[1], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusDeny, arbiters[2], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[2], err)
		return
	}

	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusPass {
		fmt.Println("claim status isn't pass")
		return
	}

	time.Sleep(time.Minute)
	time.Sleep(time.Duration(common.CloseMinute) * time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusPassEnd {
		fmt.Println("claim status isn't pass end")
		return
	}

	time.Sleep(time.Minute)
	tGetAccountQua(afterAccount, afterQualification)

	for _, arbiter := range arbiters {
		a1 := beforeAccount[arbiter]
		a2 := afterAccount[arbiter]
		tAccount(t, a1, a2)

		q1 := beforeQualification[arbiter]
		q1.Closed++
		q2 := afterQualification[arbiter]
		tQualification(t, q1, q2)
	}
}

func TestVoteChaFail2_1(t *testing.T) {
	claimId := int64(25)
	chaId := "cha01"
	tGetAccountQua(beforeAccount, beforeQualification)

	claim, err := tApply(claimId)
	if err != nil {
		return
	}

	if _, err = ApplyVote(claimId, common.ClaimStatusPass, arbiters[0], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[0], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusPass, arbiters[1], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[1], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusDeny, arbiters[2], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[2], err)
		return
	}

	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusPass {
		fmt.Println("claim status isn't pass")
		return
	}

	if _, err = ApplyChallenge(chaId, claim.CoverHash, claimId); err != nil {
		fmt.Println("ApplyChallenge error", err)
		return
	}

	if _, err = ChallengeVoteByAdmin("0x2e9475c282069675fFAc22a8cd5038E4DAC01634", claimId, common.ChallengeStatusFail); err != nil {
		fmt.Println("ChallengeVoteByAdmin error", err)
		return
	}

	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusPassChaDeny {
		fmt.Println("claim status isn't pass cha deny")
		return
	}

	time.Sleep(time.Minute)
	tGetAccountQua(afterAccount, afterQualification)

	for _, arbiter := range arbiters {
		a1 := beforeAccount[arbiter]
		a2 := afterAccount[arbiter]
		tAccount(t, a1, a2)

		q1 := beforeQualification[arbiter]
		q1.Closed++
		q2 := afterQualification[arbiter]
		tQualification(t, q1, q2)
	}
}

func TestVoteChaSuccess2_1(t *testing.T) {
	claimId := int64(19)
	chaId := "cha02"
	tGetAccountQua(beforeAccount, beforeQualification)

	claim, err := tApply(claimId)
	if err != nil {
		return
	}

	if _, err = ApplyVote(claimId, common.ClaimStatusPass, arbiters[0], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[0], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusPass, arbiters[1], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[1], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusDeny, arbiters[2], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[2], err)
		return
	}

	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusPass {
		fmt.Println("claim status isn't pass")
		return
	}

	if _, err = ApplyChallenge(chaId, claim.CoverHash, claimId); err != nil {
		fmt.Println("ApplyChallenge error", err)
		return
	}

	if _, err = ChallengeVoteByAdmin("0x2e9475c282069675fFAc22a8cd5038E4DAC01634", claimId, common.ChallengeStatusSuccess); err != nil {
		fmt.Println("ChallengeVoteByAdmin error", err)
		return
	}

	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusPassChaPass {
		fmt.Println("claim status isn't pass cha pass")
		return
	}

	time.Sleep(time.Minute)
	tGetAccountQua(afterAccount, afterQualification)

	arbiter := arbiters[0]
	a1 := beforeAccount[arbiter]
	a1.Hold = a1.Hold.Sub(utils.ArbiterNSure())
	a2 := afterAccount[arbiter]
	tAccount(t, a1, a2)

	q1 := beforeQualification[arbiter]
	q1.Closed++
	q1.Available--
	q2 := afterQualification[arbiter]
	tQualification(t, q1, q2)

	arbiter = arbiters[2]
	a1 = beforeAccount[arbiter]
	a2 = afterAccount[arbiter]
	tAccount(t, a1, a2)

	q1 = beforeQualification[arbiter]
	q1.Closed++
	q2 = afterQualification[arbiter]
	tQualification(t, q1, q2)
}

func TestVote1_2(t *testing.T) {
	claimId := int64(14)
	tGetAccountQua(beforeAccount, beforeQualification)

	claim, err := tApply(claimId)
	if err != nil {
		return
	}

	if _, err = ApplyVote(claimId, common.ClaimStatusPass, arbiters[0], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[0], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusDeny, arbiters[1], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[1], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusDeny, arbiters[2], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[2], err)
		return
	}

	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusDeny {
		fmt.Println("claim status isn't deny")
		return
	}

	time.Sleep(time.Minute)
	time.Sleep(time.Duration(common.CloseMinute) * time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusDenyEnd {
		fmt.Println("claim status isn't deny end")
		return
	}

	time.Sleep(time.Minute)
	tGetAccountQua(afterAccount, afterQualification)

	for _, arbiter := range arbiters {
		a1 := beforeAccount[arbiter]
		a2 := afterAccount[arbiter]
		tAccount(t, a1, a2)

		q1 := beforeQualification[arbiter]
		q1.Closed++
		q2 := afterQualification[arbiter]
		tQualification(t, q1, q2)
	}
}

func TestVoteChaFail1_2(t *testing.T) {
	claimId := int64(21)
	chaId := "cha03"
	tGetAccountQua(beforeAccount, beforeQualification)

	claim, err := tApply(claimId)
	if err != nil {
		return
	}

	if _, err = ApplyVote(claimId, common.ClaimStatusPass, arbiters[0], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[0], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusDeny, arbiters[1], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[1], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusDeny, arbiters[2], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[2], err)
		return
	}

	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusDeny {
		fmt.Println("claim status isn't deny")
		return
	}

	if _, err = ApplyChallenge(chaId, claim.CoverHash, claimId); err != nil {
		fmt.Println("ApplyChallenge error", err)
		return
	}

	if _, err = ChallengeVoteByAdmin("0x2e9475c282069675fFAc22a8cd5038E4DAC01634", claimId, common.ChallengeStatusFail); err != nil {
		fmt.Println("ChallengeVoteByAdmin error", err)
		return
	}

	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusDenyChaDeny {
		fmt.Println("claim status isn't deny cha deny")
		return
	}

	time.Sleep(time.Minute)
	tGetAccountQua(afterAccount, afterQualification)

	for _, arbiter := range arbiters {
		a1 := beforeAccount[arbiter]
		a2 := afterAccount[arbiter]
		tAccount(t, a1, a2)

		q1 := beforeQualification[arbiter]
		q1.Closed++
		q2 := afterQualification[arbiter]
		tQualification(t, q1, q2)
	}
}

func TestVoteChaSuccess1_2(t *testing.T) {
	claimId := int64(34)
	chaId := "cha04"
	tGetAccountQua(beforeAccount, beforeQualification)

	claim, err := tApply(claimId)
	if err != nil {
		return
	}

	if _, err = ApplyVote(claimId, common.ClaimStatusPass, arbiters[0], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[0], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusDeny, arbiters[1], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[1], err)
		return
	}
	if _, err = ApplyVote(claimId, common.ClaimStatusDeny, arbiters[2], claim.CoverHash, "111"); err != nil {
		fmt.Println("ApplyVote error", arbiters[2], err)
		return
	}

	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusDeny {
		fmt.Println("claim status isn't deny")
		return
	}

	if _, err = ApplyChallenge(chaId, claim.CoverHash, claimId); err != nil {
		fmt.Println("ApplyChallenge error", err)
		return
	}

	if _, err = ChallengeVoteByAdmin("0x2e9475c282069675fFAc22a8cd5038E4DAC01634", claimId, common.ChallengeStatusSuccess); err != nil {
		fmt.Println("ChallengeVoteByAdmin error", err)
		return
	}

	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return
	}
	if claim.Status != common.ClaimStatusDenyChaPass {
		fmt.Println("claim status isn't deny cha pass")
		return
	}

	time.Sleep(3*time.Minute)
	time.Sleep(time.Duration(common.RewardMinute)*time.Minute)
	tGetAccountQua(afterAccount, afterQualification)

	arbiter := arbiters[2]
	a1 := beforeAccount[arbiter]
	a1.Hold = a1.Hold.Sub(utils.ArbiterNSure())
	a2 := afterAccount[arbiter]
	tAccount(t, a1, a2)

	q1 := beforeQualification[arbiter]
	q1.Closed++
	q1.Available--
	q2 := afterQualification[arbiter]
	tQualification(t, q1, q2)

	arbiter = arbiters[0]
	a1 = beforeAccount[arbiter]
	a2 = afterAccount[arbiter]
	tAccount(t, a1, a2)

	q1 = beforeQualification[arbiter]
	q1.Closed++
	q2 = afterQualification[arbiter]
	tQualification(t, q1, q2)

}

func TestAccountQualification(t *testing.T) {
	tGetAccountQua(beforeAccount, beforeQualification)
	tGetAccountQua(afterAccount, afterQualification)

	for _, arbiter := range arbiters {
		a1 := beforeAccount[arbiter]
		a2 := afterAccount[arbiter]
		tAccount(t, a1, a2)

		q1 := beforeQualification[arbiter]
		q2 := afterQualification[arbiter]
		tQualification(t, q1, q2)
	}
}

func tGetAccountQua(mapAccount map[string]*models.Account, mapQua map[string]*models.Qualification) {
	for _, arbiter := range arbiters {
		if qua, account, err := GetQualificationAccount(arbiter); err == nil {
			mapAccount[arbiter] = account
			mapQua[arbiter] = qua
		}
	}
}

func tAccount(t *testing.T, a1, a2 *models.Account) {
	assert.Equal(t, a1.Available, a2.Available)
	assert.Equal(t, a1.Hold, a2.Hold)
}

func tQualification(t *testing.T, q1, q2 *models.Qualification) {
	assert.Equal(t, q1.Available, q2.Available)
	assert.Equal(t, q1.Used, q2.Used)
	assert.Equal(t, q1.Closed, q2.Closed)
}

func tApply(claimId int64) (*models.Claim, error) {
	claim, err := GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return nil, err
	}
	if claim.Status != common.ClaimStatusNew {
		fmt.Println("claim status isn't new")
		return nil, fmt.Errorf("claim status isn't new")
	}
	for _, arbiter := range arbiters {
		if _, err = ClaimApply(claimId, arbiter); err != nil {
			fmt.Println("arbiter", arbiter, err)
		}
	}
	time.Sleep(time.Minute)
	claim, err = GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return nil, err
	}
	if claim.Status != common.ClaimStatusArbiter {
		fmt.Println("claim status isn't arbiter")
		return nil, fmt.Errorf("claim status isn't arbiter")
	}
	return claim, nil
}
