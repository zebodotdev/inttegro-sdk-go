package payout

type Status string

const (
	StatusInitialized Status = "initialized"
	StatusScheduled   Status = "scheduled"
	StatusProcessing  Status = "processing"
	StatusExecuting   Status = "executing"
	StatusSucceeded   Status = "succeeded"
	StatusInvalid     Status = "invalid"
	StatusCanceled    Status = "canceled"
)
