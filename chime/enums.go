package chime

type RecipientType string

const (
	RecipientTypePhone RecipientType = "phone"
	RecipientTypeEmail RecipientType = "email"
)

type Transport string

const (
	TransportSMS   Transport = "sms"
	TransportEmail Transport = "email"
)

type EmailSchemaKind string

const (
	EmailSchemaKindGmailViewAction  EmailSchemaKind = "gmail_view_action"
	EmailSchemaKindSchemaOrgOrder   EmailSchemaKind = "schema_org_order"
	EmailSchemaKindSchemaOrgInvoice EmailSchemaKind = "schema_org_invoice"
)
