package common

const (
	LevelDebug int = 1
	LevelInfo  int = 2
	// LevelWarn int = 3 default
)

const (
	ApplyMaxNum   = 20 // 20 5
	ArbiterMaxNum = 5 // 5  3
	VoteMinNum    = 3 // 3  2
)

const (
	ApplyMinute   uint = 1440
	VoteMinute    uint = 3 * 1440
	CloseMinute   uint = 2 * 1440
	RewardMinute  uint = 1440
	BillMinute    uint = 1440
	PendingMinute uint = 14 * 1440

	RewardChaMinute uint = 1440
	PaymentMinute   uint = 1440
)

const (
	BlockNumberDeposit     = "block_number_deposit"
	BlockNumberClaim       = "block_number_claim"
	AdminAddressNum        = 8
	VoteAddressPrefix      = "vote_address_"
	ChallengeAddressPrefix = "challenge_address_"
	PaymentAddressPrefix   = "payment_address_"
	CurrencyPrefix         = "currency_"
)

const (
	ConfirmBlock       uint64 = 2
	WithdrawBackMinute uint   = 8

	UnitNSure        int64 = 1_000_000_000_000_000_000
	AccountNSure           = "0x123456aaa"
	AccountChallenge       = "0x123456ccc"
	AccountPayment         = "0x123456bbb"
	CurrencyNSure          = "Nsure"
	CurrencyETH            = "ETH"
	ArbiterNSure     int64 = 2000
	ChallengeTimes   int64 = 10
	CookieMaxAge           = 7 * 24 * 60 * 60
)

const (
	BillTypeArbiter          = BillType("Add Role")
	BillTypeDeposit          = BillType("Deposit")
	BillTypeWithdraw         = BillType("Withdraw")
	BillTypePending          = BillType("Deduct Role")
	BillTypeWithdrawLock     = BillType("WithdrawLock")
	BillTypeWithdrawBack     = BillType("WithdrawBack")
	BillTypeVotePunish       = BillType("Arbiter Punish")
	BillTypeChallenge        = BillType("Challenge")
	BillTypeChallengeSuccess = BillType("Challenge Success")
	BillTypeChallengeFail    = BillType("Challenge Fail")
	BillTypeReward           = BillType("Claim Reward")
	BillTypePayment          = BillType("Claim Pay")
)

const (
	LoginLruSize = 1000
	MaxRandNum   = 1000_000

	PendingWorkChanNum       = 100
	PendingWorkUnsettleCount = 1000
	PendingInspectorTime     = 10

	BillWorkChanNum       = 100
	BillWorkUnsettleCount = 1000
	BillInspectorTime     = 10
	BillWaitInspectorTime = 120

	DepositInspectorTime      = 120
	WithdrawInspectorTime     = 120
	WithdrawBackInspectorTime = 20

	PaymentInspectorTime = 30
	PaymentCount         = 1000

	RewardInspectorTime    = 30
	RewardChaInspectorTime = 30

	VoteFillWorkChanNum   = 100
	VoteFillInspectorTime = 30

	ChallengeWorkChanNum       = 100
	ChallengeWorkUnsettleCount = 1000
	ChallengeInspectorTime     = 30

	ApplyWorkChanNum   = 100
	ApplyInspectorTime = 10
	ClaimInspectorTime = 20

	VoteInspectorTime = 10
)

const (
	ClaimStatusNew         = ClaimStatus("New")
	ClaimStatusArbiter     = ClaimStatus("Arbiter")
	ClaimStatusApplyFail   = ClaimStatus("ApplyFail")
	ClaimStatusPass        = ClaimStatus("Pass")
	ClaimStatusPassEnd     = ClaimStatus("PassEnd")
	ClaimStatusDeny        = ClaimStatus("Deny")
	ClaimStatusDenyEnd     = ClaimStatus("DenyEnd")
	ClaimStatusArbiterFail = ClaimStatus("ArbiterFail")
	ClaimStatusPassCha     = ClaimStatus("PassCha")
	ClaimStatusDenyCha     = ClaimStatus("DenyCha")
	ClaimStatusPassChaPass = ClaimStatus("PassChaPass")
	ClaimStatusPassChaDeny = ClaimStatus("PassChaDeny")
	ClaimStatusDenyChaPass = ClaimStatus("DenyChaPass")
	ClaimStatusDenyChaDeny = ClaimStatus("DenyChaDeny")

	PaymentStatusCha  = PaymentStatus("Challenge")
	PaymentStatusPass = PaymentStatus("Pass")
	PaymentStatusFail = PaymentStatus("Fail")
)

const (
	ApplyStatusApply   = ApplyStatus("Apply")
	ApplyStatusSuccess = ApplyStatus("Arbiter")
	ApplyStatusFail    = ApplyStatus("Fail")
	ApplyStatusPass    = ApplyStatus("Pass")
	ApplyStatusDeny    = ApplyStatus("Deny")

	FillStatusEqual        = FillStatus("Equal")
	FillStatusDifferent    = FillStatus("Different")
	FillStatusChaEqual     = FillStatus("ChaEqual")
	FillStatusChaDifferent = FillStatus("ChaDifferent")
	FillStatusEmpty        = FillStatus("Empty")
	FillStatusAbstain      = FillStatus("Abstain")
)

const (
	ChallengeStatusApply   = ChallengeStatus("Apply")
	ChallengeStatusSuccess = ChallengeStatus("Success")
	ChallengeStatusFail    = ChallengeStatus("Fail")
)

const (
	WithdrawStatusApply   = WithdrawStatus("Apply")
	WithdrawStatusSuccess = WithdrawStatus("Success")
	WithdrawStatusFail    = WithdrawStatus("Fail")
)

const (
	TransferStatusDeposit = TransferStatus("Deposit")
)

const (
	TypedDataDomainNameClaim    = "Treasury"
	TypedDataDomainNameWithdraw = "Treasury"
	DurationClaim               = 28
)
