package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"nsure/vote/common"
	"nsure/vote/models"
	"testing"
	"time"
)

var applies []string

func init() {
	applies = []string{"test01", "test02", "test03", "test04", "test05", "test06", "test07", "test08", "test09"}
}
func TestApply02(t *testing.T) {
	beforeQua := make(map[string]*models.Qualification)
	afterQua := make(map[string]*models.Qualification)
	testQualification(applies, beforeQua)
	claimId, _, err := testNewApply("0x0000006")
	if err != nil {
		fmt.Println("testNewApply error", err)
		return
	}

	time.Sleep(time.Minute)
	claim, err := GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err", err)
		return
	}
	if claim.Status != common.ClaimStatusArbiter {
		fmt.Println("claim status isn't arbiter")
		return
	}

	votes, err := GetVoteByClaimId(claimId)
	if err != nil {
		fmt.Println("GetVoteByClaimId error", err)
		return
	}
	var votesId []string
	for _, vote := range votes {
		votesId = append(votesId, vote.ArbiterId)
	}

	testQualification(applies, afterQua)
	mapApply := make(map[string]struct{})
	for _, apply := range votesId {
		fmt.Println("test apply ok:", apply)
		q1 := beforeQua[apply]
		q1.Used++
		q1.Available--
		q2 := afterQua[apply]
		tQualification(t, q1, q2)
		mapApply[apply] = struct{}{}
	}
	for _, apply := range applies {
		if _, found := mapApply[apply]; !found {
			fmt.Println("test apply fail:", apply)
			q1 := beforeQua[apply]
			q2 := afterQua[apply]
			tQualification(t, q1, q2)
		}
	}
}

func TestApply01(t *testing.T) {
	beforeQua := make(map[string]*models.Qualification)
	afterQua := make(map[string]*models.Qualification)
	testQualification(applies, beforeQua)
	_, applyOk, err := testNewApply("0x0000001")
	if err != nil {
		fmt.Println("testNewApply error", err)
		return
	}
	testQualification(applies, afterQua)
	mapApply := make(map[string]struct{})
	for _, apply := range applyOk {
		fmt.Println("test apply ok:", apply)
		q1 := beforeQua[apply]
		q1.Used++
		q1.Available--
		q2 := afterQua[apply]
		tQualification(t, q1, q2)
		mapApply[apply] = struct{}{}
	}
	for _, apply := range applies {
		if _, found := mapApply[apply]; !found {
			fmt.Println("test apply fail:", apply)
			q1 := beforeQua[apply]
			q2 := afterQua[apply]
			tQualification(t, q1, q2)
		}
	}
}

func testNewApply(coverHash string) (int64, []string, error) {
	claimId, err := testNewClaim(coverHash)
	if err != nil {
		return 0, nil, err
	}
	claim, err := GetClaimById(claimId)
	if err != nil {
		fmt.Println("get claim err")
		return 0, nil, err
	}
	if claim.Status != common.ClaimStatusNew {
		fmt.Println("claim status isn't new")
		return 0, nil, fmt.Errorf("claim status isn't new")
	}
	var applyOk []string
	for _, apply := range applies {
		if _, err = ClaimApply(claimId, apply); err != nil {
			fmt.Println("apply err", apply, err)
		} else {
			fmt.Println("apply ok", apply)
			applyOk = append(applyOk, apply)
		}
	}
	return claimId, applyOk, nil
}

func testNewClaim(coverHash string) (int64, error) {
	userId := "user01"
	coverId := "001"
	amount := decimal.New(10000, 0)
	cost := decimal.New(1000, 0)
	reward := decimal.New(100, 0)
	submitAt := time.Now()
	beginAt := submitAt.AddDate(0, 0, -10)
	endAt := beginAt.AddDate(0, 0, 30)

	if claimId, err := AddClaim(userId, "ProductTest", coverId, coverHash, "ETH", amount, cost,
		reward, submitAt, beginAt, endAt, "desc", "cred", "loss"); err != nil {
		return 0, err
	} else {
		return claimId, nil
	}
}

func testQualification(userIds []string, mapQua map[string]*models.Qualification) {
	for _, userId := range userIds {
		if qua, err := GetQualificationByArbiterId(userId); err == nil {
			mapQua[userId] = qua
		}
	}
}
