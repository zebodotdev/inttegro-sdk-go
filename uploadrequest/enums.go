package uploadrequest

type Status string

const (
	StatusPending   Status = "pending"
	StatusUploading Status = "uploading"
	StatusFulfilled Status = "fulfilled"
	StatusExpired   Status = "expired"
	StatusCanceled  Status = "canceled"
	StatusFailed    Status = "failed"
)

type ReviewDecision string

const (
	ReviewDecisionApproved ReviewDecision = "approved"
	ReviewDecisionRejected ReviewDecision = "rejected"
)

type ReviewType string

const (
	ReviewTypeAutomatic ReviewType = "automatic"
	ReviewTypeManual    ReviewType = "manual"
)
