package filelink

type Status string

const (
	StatusActive   Status = "active"
	StatusRevoked  Status = "revoked"
	StatusExpired  Status = "expired"
	StatusDisabled Status = "disabled"
)

type Kind string

const KindPublic Kind = "public"

type DeliveryMode string

const (
	DeliveryModeRedirect DeliveryMode = "redirect"
	DeliveryModeDownload DeliveryMode = "download"
	DeliveryModeInline   DeliveryMode = "inline"
)
