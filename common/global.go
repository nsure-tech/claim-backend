package common

import "github.com/shopspring/decimal"

var ClaimActive int
var ClaimClosed int
var RewardTotal decimal.Decimal

var VoteAddress map[string]struct{}
var ChallengeAddress map[string]struct{}
var PaymentAddress map[string]struct{}
