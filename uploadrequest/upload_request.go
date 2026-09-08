// Package uploadrequest provides delegated upload-request resources and operations.
package uploadrequest

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service        = inttegro.UploadRequestsService
	Status         = inttegro.UploadRequestStatus
	ReviewDecision = inttegro.UploadReviewDecision
	ReviewType     = inttegro.UploadReviewType
	Constraints    = inttegro.UploadRequestConstraints
	Display        = inttegro.UploadRequestDisplay
	Party          = inttegro.FileParty
	FileResource   = inttegro.FileResource
	CreateParams   = inttegro.UploadRequestCreateParams
	PageParams     = inttegro.UploadRequestPageParams
	CancelParams   = inttegro.UploadRequestCancelParams
	ReviewReason   = inttegro.UploadRequestReviewReason
	ReviewParams   = inttegro.UploadRequestReviewParams
	FulfillParams  = inttegro.UploadRequestFulfillParams
	Resource       = inttegro.UploadRequest
	Page           = inttegro.UploadRequestsPage
)

const (
	StatusPending   = inttegro.UploadRequestStatusPending
	StatusUploading = inttegro.UploadRequestStatusUploading
	StatusFulfilled = inttegro.UploadRequestStatusFulfilled
	StatusExpired   = inttegro.UploadRequestStatusExpired
	StatusCanceled  = inttegro.UploadRequestStatusCanceled
	StatusFailed    = inttegro.UploadRequestStatusFailed

	ReviewDecisionApproved = inttegro.UploadReviewDecisionApproved
	ReviewDecisionRejected = inttegro.UploadReviewDecisionRejected

	ReviewTypeAutomatic = inttegro.UploadReviewTypeAutomatic
	ReviewTypeManual    = inttegro.UploadReviewTypeManual
)
