// Package file provides file resources, downloads, and operations.
package file

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service         = inttegro.FilesService
	Status          = inttegro.FileStatus
	Disposition     = inttegro.FileDisposition
	Delivery        = inttegro.FileDelivery
	ScanStatus      = inttegro.FileScanStatus
	SourceType      = inttegro.FileSourceType
	StorageEncoding = inttegro.FileStorageEncoding
	CreateParams    = inttegro.FileCreateParams
	PageParams      = inttegro.FilePageParams
	ContentsParams  = inttegro.FileContentsParams
	Resource        = inttegro.File
	Page            = inttegro.FilesPage
	Download        = inttegro.FileDownload
)

const (
	StatusUploading  = inttegro.FileStatusUploading
	StatusProcessing = inttegro.FileStatusProcessing
	StatusAvailable  = inttegro.FileStatusAvailable
	StatusFailed     = inttegro.FileStatusFailed
	StatusDeleted    = inttegro.FileStatusDeleted

	DispositionAttachment = inttegro.FileDispositionAttachment
	DispositionInline     = inttegro.FileDispositionInline

	DeliveryStream   = inttegro.FileDeliveryStream
	DeliveryRedirect = inttegro.FileDeliveryRedirect

	ScanStatusPending = inttegro.FileScanStatusPending
	ScanStatusPassed  = inttegro.FileScanStatusPassed
	ScanStatusFailed  = inttegro.FileScanStatusFailed
	ScanStatusSkipped = inttegro.FileScanStatusSkipped

	SourceTypeDirect        = inttegro.FileSourceTypeDirect
	SourceTypeUploadRequest = inttegro.FileSourceTypeUploadRequest
	SourceTypeService       = inttegro.FileSourceTypeService

	StorageEncodingIdentity = inttegro.FileStorageEncodingIdentity
	StorageEncodingBrotli   = inttegro.FileStorageEncodingBrotli
)
