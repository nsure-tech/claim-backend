package models

import (
	"github.com/shopspring/decimal"
	"nsure/vote/common"
	"time"
)

type Config struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	KeyWord   string
	Val       string
}

type Metamask struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    string
	SigHex    string
	Msg       string
	Settled   bool
}

type Transfer struct {
	Id          int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Currency    string
	FromAddress string
	ToAddress   string
	Amount      decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Raw         string
	Status      common.TransferStatus
	Settled     bool
}

type Account struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    string          `gorm:"column:user_id;unique_index:idx_uid_currency"`
	Currency  string          `gorm:"column:currency;unique_index:idx_uid_currency"`
	Available decimal.Decimal `gorm:"column:available" sql:"type:decimal(64,0);"`
	Hold      decimal.Decimal `gorm:"column:hold" sql:"type:decimal(64,0);"`
}

type Bill struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    string
	Currency  string
	Available decimal.Decimal `sql:"type:decimal(64,0);"`
	Hold      decimal.Decimal `sql:"type:decimal(64,0);"`
	Type      common.BillType
	Settled   bool
	Notes     string
}

type Qualification struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ArbiterId string
	Available int
	Used      int
	Pending   int
	Closed    int
}

type Pending struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ArbiterId string
	SubmitAt  time.Time
	Pending   int
	Settled   bool
}

type Claim struct {
	Id            int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserId        string
	Product       string
	CoverId       string
	CoverHash     string
	Currency      string
	Amount        decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Cost          decimal.Decimal `gorm:"column:cost" sql:"type:decimal(64,0);"`
	Reward        decimal.Decimal `gorm:"column:reward" sql:"type:decimal(64,0);"`
	SubmitAt      time.Time
	ArbiterAt     time.Time
	VoteAt        time.Time
	ChallengeAt   time.Time
	CoverBeginAt  time.Time
	CoverEndAt    time.Time
	Status        common.ClaimStatus // `gorm:"column:claim_status"`
	PaymentStatus common.PaymentStatus
	ApplyNum      int
	VoteNum       int
	Challenged    bool
	Description   string
	Credential    string
	Settled       bool
	Notes         string
}

type Payment struct {
	Id          int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserId      string
	Product     string
	CoverId     string
	CoverHash   string
	Currency    string
	Amount      decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	ClaimId     int64
	ClaimStatus common.ClaimStatus
	Settled     bool
	Notes       string
}

type Challenge struct {
	Id          int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ChallengeId string
	ChallengeAt time.Time
	ClaimId     int64
	CoverId     string
	CoverHash   string
	Currency    string
	Amount      decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Reward      decimal.Decimal `gorm:"column:reward" sql:"type:decimal(64,0);"`
	Hold        decimal.Decimal `gorm:"column:hold" sql:"type:decimal(64,0);"`
	ClaimStatus common.ClaimStatus
	Status      common.ChallengeStatus
	Settled     bool
	Notes       string
}

type ChallengeFill struct {
	Id          int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ChallengeId string
	ChallengeAt time.Time
	ClaimId     int64
	CoverId     string
	CoverHash   string
	UserId      string
	Product     string
	Currency    string
	Amount      decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Reward      decimal.Decimal `gorm:"column:reward" sql:"type:decimal(64,0);"`
	Hold        decimal.Decimal `gorm:"column:hold" sql:"type:decimal(64,0);"`
	ClaimStatus common.ClaimStatus
	Status      common.ChallengeStatus
	Settled     bool
	Notes       string
}

type ClaimDesc struct {
	Id          int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ClaimId     int64
	Description string
}

type ClaimCred struct {
	Id          int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ClaimId     int64
	Credentials string
}

type Apply struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ClaimId   int64
	SubmitAt  time.Time
	CoverId   string
	CoverHash string
	UserId    string
	Product   string
	ApplyAt   time.Time
	ArbiterAt time.Time
	ApplyNum  int
	ArbiterId string
	Status    common.ApplyStatus
	Settled   bool
}

type Vote struct {
	Id           int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ClaimId      int64
	ArbiterId    string
	UserId       string
	Product      string
	CoverId      string
	CoverHash    string
	Currency     string
	Amount       decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Reward       decimal.Decimal `gorm:"column:reward" sql:"type:decimal(64,0);"`
	ArbiterAt    time.Time
	SubmitAt     time.Time
	CoverBeginAt time.Time
	Status       common.ClaimStatus
	SignHash     string
	Settled      bool
	Notes        string
}

type VoteFill struct {
	Id              int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	VoteId          int64
	ClaimId         int64
	CoverId         string
	CoverHash       string
	Currency        string
	Amount          decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Reward          decimal.Decimal `gorm:"column:reward" sql:"type:decimal(64,0);"`
	ArbiterId       string
	VoteAt          time.Time
	ClaimStatus     common.ClaimStatus
	VoteStatus      common.ClaimStatus
	PaymentStatus   common.PaymentStatus
	ChallengeStatus common.ChallengeStatus
	Status          common.FillStatus
	RewardNum       int
	Settled         bool
	Notes           string
}

type Reward struct {
	Id          int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	VoteId      int64
	ClaimId     int64
	CoverId     string
	CoverHash   string
	ArbiterId   string
	ClaimStatus common.ClaimStatus
	VoteStatus  common.ClaimStatus
	Currency    string
	Amount      decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Settled     bool
	Notes       string
}

type RewardCha struct {
	Id              int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ClaimId         int64
	CoverId         string
	CoverHash       string
	ChallengeId     string
	ClaimStatus     common.ClaimStatus
	ChallengeStatus common.ChallengeStatus
	Currency        string
	Amount          decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Settled         bool
	Notes           string
}

type PunishCha struct {
	Id              int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ClaimId         int64
	CoverId         string
	CoverHash       string
	ChallengeId     string
	ClaimStatus     common.ClaimStatus
	ChallengeStatus common.ChallengeStatus
	Currency        string
	Amount          decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Settled         bool
	Notes           string
}

type ChainClaim struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    string
	Currency  string
	Amount    decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Nonce     uint64
	Raw       string
	Status    common.WithdrawStatus
	Settled   bool
}

type Withdraw struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    string
	Currency  string
	Amount    decimal.Decimal `gorm:"column:amount" sql:"type:decimal(64,0);"`
	Nonce     uint64
	EndAt     time.Time
	Status    common.WithdrawStatus
	Settled   bool
}
