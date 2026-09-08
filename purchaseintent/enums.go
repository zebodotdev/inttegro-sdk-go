package purchaseintent

type Status string

const (
	StatusActive   Status = "active"
	StatusExpired  Status = "expired"
	StatusInactive Status = "inactive"
	StatusUsed     Status = "used"
)

type ActivityType string

const (
	ActivityTypeExpiredViewed  ActivityType = "expired_viewed"
	ActivityTypeOrderCreated   ActivityType = "order_created"
	ActivityTypePaymentFailed  ActivityType = "payment_failed"
	ActivityTypePaymentStarted ActivityType = "payment_started"
	ActivityTypeViewed         ActivityType = "viewed"
)
