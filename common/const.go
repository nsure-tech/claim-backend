package common

const (
	LevelDebug int = 1
	LevelInfo  int = 2
	// LevelWarn int = 3 default
)

const (
	BlockNumberDeposit     = "block_number_deposit"
	BlockNumberClaim       = "block_number_claim"
	AdminAddressNum        = 8
	AdminAddressPrefix     = "admin_address_"
	ChallengeAddressNum    = 5
	ChallengeAddressPrefix = "challenge_address_"
	CurrencyPrefix         = "currency_"
)

const (
	ConfirmBlock     uint64 = 0
	UnitNSure        int64  = 1_000_000_000_000_000_000
	AccountNSure            = "0x123456aaa"
	AccountChallenge        = "0x123456ccc"
	AccountPayment          = "0x123456bbb"
	CurrencyNSure           = "Nsure"
	ArbiterNSure     int64  = 5000
	ChallengeTimes   int64  = 10
	CookieMaxAge            = 7 * 24 * 60 * 60
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
)

const (
	LoginLruSize = 1000
	MaxRandNum   = 1000_000

	PendingWorkChanNum       = 1000
	PendingWorkUnsettleCount = 1000
	PendingInspectorTime     = 60

	BillWorkChanNum       = 1000
	BillWorkUnsettleCount = 1000
	BillInspectorTime     = 30

	DepositInspectorTime      = 5
	WithdrawInspectorTime     = 5
	WithdrawBackInspectorTime = 5

	PaymentWorkChanNum   = 1000
	PaymentInspectorTime = 30

	RewardInspectorTime    = 30
	RewardChaInspectorTime = 30
	PunishChaInspectorTime = 30

	FillWorkChanNum       = 1000
	FillWorkUnsettleCount = 1000
	VoteFillInspectorTime = 30

	ChallengeWorkChanNum       = 1000
	ChallengeWorkUnsettleCount = 1000
	ChallengeInspectorTime     = 30

	ApplyWorkChanNum   = 1000
	ApplyInspectorTime = 1
	ClaimInspectorTime = 2

	VoteWorkChanNum   = 1000
	VoteInspectorTime = 1
)

const (
	ApplyMinute     uint  = 1440
	VoteMinute      uint  = 1440
	CloseMinute     uint  = 10 //
	RewardMinute    uint  = 10
	RewardChaMinute uint  = 10
	PunishChaMinute uint  = 10
	PaymentMinute   uint  = 10
	WithdrawBack    uint  = 0
	ApplyMaxNum     uint8 = 5  // 20 5
	ArbiterMaxNum         = 3  // 5  3
	VoteMinNum            = 2  // 3  2
	PendingTime     uint  = 10 //14 * 1440
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
	DurationWithdraw            = 30
)
