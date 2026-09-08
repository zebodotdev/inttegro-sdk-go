package payment

type NextActionType string

const (
	NextActionTypeConfirmPayment NextActionType = "confirm_payment"
	NextActionTypeExecute        NextActionType = "execute"
	NextActionTypeRedirect       NextActionType = "redirect"
	NextActionTypeAuthorize      NextActionType = "authorize"
	NextActionTypeNone           NextActionType = "none"
)

type ConfirmationChannel string

const (
	ConfirmationChannelSMS   ConfirmationChannel = "sms"
	ConfirmationChannelEmail ConfirmationChannel = "email"
	ConfirmationChannelPush  ConfirmationChannel = "push"
)

type Status string

const (
	StatusInitiated      Status = "initiated"
	StatusRequiresAction Status = "requires_action"
	StatusOverdue        Status = "overdue"
	StatusExecuted       Status = "executed"
	StatusPaid           Status = "paid"
	StatusCanceled       Status = "canceled"
	StatusExpired        Status = "expired"
	StatusFailed         Status = "failed"
	StatusUnknown        Status = "unknown"
)

type AttemptStatus string

const (
	AttemptStatusInitiated AttemptStatus = "initiated"
	AttemptStatusExecuted  AttemptStatus = "executed"
	AttemptStatusSucceeded AttemptStatus = "succeeded"
	AttemptStatusCanceled  AttemptStatus = "canceled"
	AttemptStatusExpired   AttemptStatus = "expired"
	AttemptStatusFailed    AttemptStatus = "failed"
	AttemptStatusUnknown   AttemptStatus = "unknown"
)

type ResultStatus string

const (
	ResultStatusPending              ResultStatus = "pending"
	ResultStatusRequiresConfirmation ResultStatus = "requires_confirmation"
	ResultStatusProcessing           ResultStatus = "processing"
	ResultStatusSucceeded            ResultStatus = "succeeded"
	ResultStatusFailed               ResultStatus = "failed"
)
