package chime

type PageParams struct {
	CustomerID string `json:"customer_id,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	Recipient  string `json:"recipient,omitempty"`
}

// SendChimeParams sends a notification immediately.
//
// Chimes are transactional notifications sent to customers via SMS or email.
// Unlike invoice emails, chimes give you full control over message content.
//
// Example (SMS notification):
//
//	params := chime.SendParams{
//	    Recipient: chime.Recipient{
//	        Type: chime.RecipientTypePhone,
//	        Phone: &struct{Number string}{Number: "+233244123456"},
//	    },
//	    FullMessage: "Your order #12345 has shipped!",
//	    Transport:   chime.TransportSMS,
//	    Purpose:     "order_shipped",
//	    IdempotencyKey: "chime_order_12345_shipped",
//	}
type SendParams struct {
	// Recipient specifies who receives the chime (required).
	Recipient Recipient `json:"recipient"`

	// FullMessage is the complete message content (required).
	// For SMS: maximum 160 characters for single message, 1530 for concatenated.
	// For email: becomes the email body (plain text).
	FullMessage string `json:"full_message"`

	// Transport specifies delivery channel (optional, default: inferred from recipient type).
	// Values: "sms" or "email"
	// If omitted, uses SMS for phone recipients, email for email recipients.
	Transport Transport `json:"transport,omitempty"`

	// Sender is the sender ID shown to recipient (optional).
	// For SMS: sender name (max 11 characters). Example: "AcmeStore"
	// For email: sender name in "From" header. Example: "Acme Support"
	// If omitted, uses your business name from settings.
	Sender string `json:"sender,omitempty"`

	// Purpose categorizes the chime for analytics (optional).
	// Examples: "order_confirmation", "payment_reminder", "otp"
	// Use consistent values for reporting and filtering.
	Purpose string `json:"purpose,omitempty"`

	// CustomData holds arbitrary key-value custom data (optional).
	// Both keys and values must be strings. Maximum 25KB when serialized.
	// Useful for linking chimes to your internal records.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// IdempotencyKey prevents duplicate sends (optional but strongly recommended).
	// If the same key is used twice, the second request returns the original chime
	// instead of sending a duplicate. Keys can be reused if the original failed.
	// Example: "chime_order_12345_shipped"
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

// ScheduleChimeParams schedules a notification for future delivery.
//
// Use this to queue notifications to one or more recipient addresses.
// Useful for reminders, follow-ups, or coordinated messaging campaigns.
//
// Example:
//
//	params := chime.ScheduleParams{
//	    Recipients:  []string{"+233244123456", "user@example.com"},
//	    FullMessage: "Your subscription renews tomorrow.",
//	    SendAfter:   "2024-01-15T09:00:00Z",
//	    SenderID:    "YourBrand",
//	}
type ScheduleParams struct {
	// Recipients specifies all recipient addresses (required).
	Recipients []string `json:"recipients,omitempty"`

	// FullMessage is the complete message content (required).
	// For SMS: maximum 160 characters for single message.
	// For email: becomes the email body.
	FullMessage string `json:"full_message"`

	// SendAfter is when to send the chime (required).
	// ISO 8601 timestamp. Must be in the future.
	// Example: "2024-01-15T09:00:00Z"
	SendAfter string `json:"send_after,omitempty"`

	// SenderID is the sender identifier displayed to recipients (optional).
	SenderID string `json:"sender_id,omitempty"`

	// Purpose categorizes the chime for analytics (optional).
	Purpose string `json:"purpose,omitempty"`

	// IdempotencyKey prevents duplicate scheduling (optional but recommended).
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

// LookupChimeParams specifies which chime to retrieve.
type LookupParams struct {
	// ChimeID is the unique chime identifier (required).
	// Starts with "chm_". Example: "chm_abc123def456"
	ChimeID string `json:"chime_id"`
}
