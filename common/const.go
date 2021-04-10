package common

const (
	LevelDebug int = 1
	LevelInfo  int = 2
	// LevelWarn int = 3 default
)

const (
	ApplyMaxNum   = 5 // 20 5
	ArbiterMaxNum = 3 // 5  3
	VoteMinNum    = 2 // 3  2
)

const (
	ApplyMinute     uint = 1440
	VoteMinute      uint = 1440
	CloseMinute     uint = 3 * 1440
	RewardMinute    uint = 1440
	RewardChaMinute uint = 1440
	PaymentMinute   uint = 1440
	PendingMinute   uint = 14 * 1440
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
	ClaimStatusNew         = ClaimStatus("new")
	ClaimStatusArbiter     = ClaimStatus("arbiter")
	ClaimStatusApplyFail   = ClaimStatus("apply_fail")
	ClaimStatusPass        = ClaimStatus("pass")
	ClaimStatusPassEnd     = ClaimStatus("pass_end")
	ClaimStatusDeny        = ClaimStatus("deny")
	ClaimStatusDenyEnd     = ClaimStatus("deny_end")
	ClaimStatusArbiterFail = ClaimStatus("arbiter_fail")
	ClaimStatusPassCha     = ClaimStatus("pass_cha")
	ClaimStatusDenyCha     = ClaimStatus("deny_cha")
	ClaimStatusPassChaPass = ClaimStatus("pass_cha_pass")
	ClaimStatusPassChaDeny = ClaimStatus("pass_cha_deny")
	ClaimStatusDenyChaPass = ClaimStatus("deny_cha_pass")
	ClaimStatusDenyChaDeny = ClaimStatus("deny_cha_deny")

	PaymentStatusCha  = PaymentStatus("challenge")
	PaymentStatusPass = PaymentStatus("pass")
	PaymentStatusFail = PaymentStatus("fail")
)

const (
	ApplyStatusApply   = ApplyStatus("apply")
	ApplyStatusSuccess = ApplyStatus("arbiter")
	ApplyStatusFail    = ApplyStatus("fail")
	ApplyStatusPass    = ApplyStatus("pass")
	ApplyStatusDeny    = ApplyStatus("deny")

	FillStatusEqual        = FillStatus("equal")
	FillStatusDifferent    = FillStatus("different")
	FillStatusChaEqual     = FillStatus("cha_equal")
	FillStatusChaDifferent = FillStatus("cha_different")
	FillStatusEmpty        = FillStatus("empty")
	FillStatusAbstain      = FillStatus("abstain")
)

const (
	ChallengeStatusApply   = ChallengeStatus("apply")
	ChallengeStatusSuccess = ChallengeStatus("success")
	ChallengeStatusFail    = ChallengeStatus("fail")
)

const (
	WithdrawStatusApply   = WithdrawStatus("apply")
	WithdrawStatusSuccess = WithdrawStatus("success")
	WithdrawStatusFail    = WithdrawStatus("fail")
)

const (
	TransferStatusDeposit = TransferStatus("deposit")
)

const (
	TypedDataDomainNameClaim    = "Treasury"
	TypedDataDomainNameWithdraw = "Treasury"
	DurationClaim               = 30
)
