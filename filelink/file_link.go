// Package filelink provides public file-link resources and operations.
package filelink

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service      = inttegro.FileLinksService
	Status       = inttegro.FileLinkStatus
	Kind         = inttegro.FileLinkKind
	DeliveryMode = inttegro.FileLinkDeliveryMode
	Actor        = inttegro.FileActor
	Delivery     = inttegro.FileLinkDelivery
	Access       = inttegro.FileLinkAccess
	CreateParams = inttegro.FileLinkCreateParams
	PageParams   = inttegro.FileLinkPageParams
	RevokeParams = inttegro.FileLinkRevokeParams
	Resource     = inttegro.FileLink
	Page         = inttegro.FileLinksPage
)

const (
	StatusActive   = inttegro.FileLinkStatusActive
	StatusRevoked  = inttegro.FileLinkStatusRevoked
	StatusExpired  = inttegro.FileLinkStatusExpired
	StatusDisabled = inttegro.FileLinkStatusDisabled

	KindPublic = inttegro.FileLinkKindPublic

	DeliveryModeRedirect = inttegro.FileLinkDeliveryModeRedirect
	DeliveryModeDownload = inttegro.FileLinkDeliveryModeDownload
	DeliveryModeInline   = inttegro.FileLinkDeliveryModeInline
)
