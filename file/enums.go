package file

type Status string

const (
	StatusUploading  Status = "uploading"
	StatusProcessing Status = "processing"
	StatusAvailable  Status = "available"
	StatusFailed     Status = "failed"
	StatusDeleted    Status = "deleted"
)

type Disposition string

const (
	DispositionAttachment Disposition = "attachment"
	DispositionInline     Disposition = "inline"
)

type Delivery string

const (
	DeliveryStream   Delivery = "stream"
	DeliveryRedirect Delivery = "redirect"
)

type ScanStatus string

const (
	ScanStatusPending ScanStatus = "pending"
	ScanStatusPassed  ScanStatus = "passed"
	ScanStatusFailed  ScanStatus = "failed"
	ScanStatusSkipped ScanStatus = "skipped"
)

type SourceType string

const (
	SourceTypeDirect        SourceType = "direct"
	SourceTypeUploadRequest SourceType = "upload_request"
	SourceTypeService       SourceType = "service"
)

type StorageEncoding string

const (
	StorageEncodingIdentity StorageEncoding = "identity"
	StorageEncodingBrotli   StorageEncoding = "br"
)
