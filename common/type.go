package common

import (
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/shopspring/decimal"
	"time"
)

type ClaimStatus string
type ChallengeStatus string
type RewardStatus string
type PaymentStatus string
type FillStatus string
type ApplyStatus string
type BillType string
type TransferStatus string
type WithdrawStatus string
type MsgData core.TypedData

type ClaimId struct {
	ClaimId int64
}

type Reward struct {
	Reward decimal.Decimal
}

type BuyClaim struct {
	Buyer    string
	Amount   decimal.Decimal
	Cost     decimal.Decimal
	Reward   decimal.Decimal
	Period   uint64
	CreateAt time.Time
}

type WithdrawResult struct {
	R        string `json:"r"`
	S        string `json:"s"`
	V        string `json:"v"`
	DeadLine string `json:"deadline"`
}

type WithdrawResult11 struct {
	SigHash  string `json:"sig_hash"`
	DeadLine string `json:"dead_line"`
}
