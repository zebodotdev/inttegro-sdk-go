// Package chime provides chime resources and operations.
package chime

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/broadcast"
	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/schedule"
)

// ChimesService sends and manages notification messages (SMS and email).
//
// Chimes are transactional notifications you send to customers. Use them for:
//   - Order confirmations and shipping updates
//   - Payment receipts and reminders
//   - OTP codes and security alerts
//   - Custom notifications
//
// Example:
//
//	chime, err := client.Chimes.Send(ctx, chime.SendParams{
//	    Recipient: chime.Recipient{
//	        Type: chime.RecipientTypePhone,
//	        Phone: &struct{Number string}{Number: "+233244123456"},
//	    },
//	    FullMessage: "Your order #12345 has shipped!",
//	    Transport:   chime.TransportSMS,
//	    IdempotencyKey: "chime_order_12345_shipped",
//	})
type Service struct {
	client transport.Client
}

type PageParams struct {
	CustomerID string `json:"customer_id,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	Recipient  string `json:"recipient,omitempty"`
}

type Page struct {
	Number int        `json:"number,omitempty"`
	Size   int        `json:"size,omitempty"`
	Chimes []Resource `json:"chimes,omitempty"`
}

// Send dispatches a notification immediately.
//
// Sends SMS or email to a customer. Returns immediately with chime ID—
// delivery happens asynchronously.
func (s *Service) Send(ctx context.Context, params SendParams) (*Resource, error) {
	var resp struct {
		Chime Resource `json:"chime"`
	}
	if err := s.client.Do(ctx, "POST", "/chimes/send", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Chime, nil
}

// Lookup retrieves chime details and delivery status by ID.
func (s *Service) Lookup(ctx context.Context, chimeID string) (*Resource, error) {
	var resp struct {
		Chime Resource `json:"chime"`
	}
	if err := s.client.Do(ctx, "POST", "/chimes/lookup", LookupParams{ChimeID: chimeID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Chime, nil
}

// Page retrieves a paginated list of chimes.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/chimes/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Schedule enqueues a notification for delivery at a specific time.
//
// Useful for reminders, follow-ups, or coordinated campaigns.
func (s *Service) Schedule(ctx context.Context, params ScheduleParams) (*schedule.ScheduledChime, error) {
	var resp struct {
		ScheduledChime schedule.ScheduledChime `json:"scheduled_chime"`
	}
	if err := s.client.Do(ctx, "POST", "/chimes/schedule", params, &resp); err != nil {
		return nil, err
	}
	return &resp.ScheduledChime, nil
}

// Broadcast sends a chime to many recipients.
func (s *Service) Broadcast(ctx context.Context, params broadcast.CreateParams) (*broadcast.Creation, error) {
	var resp struct {
		Broadcast broadcast.Creation `json:"broadcast"`
	}
	if err := s.client.Do(ctx, "POST", "/chimes/broadcast", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Broadcast, nil
}

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

// ChimeRecipient specifies who should receive a notification chime.
//
// Set Type and the corresponding field (Phone or Email).
//
// Example (SMS):
//
//	recipient := chime.Recipient{
//	    Type: chime.RecipientTypePhone,
//	    Name: "Jane Doe",
//	    Phone: &struct{Number string `json:"number"`}{Number: "+233244123456"},
//	}
type Recipient struct {
	// Type specifies how the recipient is identified (required).
	// Values: "phone" or "email"
	Type RecipientType `json:"type"`

	// Name is the recipient's display name (optional).
	// Used in email subject lines and SMS sender names.
	Name string `json:"name,omitempty"`

	// Phone contains the phone number (required when Type is "phone").
	// Must include country code. Example: "+233244123456"
	Phone *struct {
		Number string `json:"number"`
	} `json:"phone,omitempty"`

	// Email contains the email address (required when Type is "email").
	// Validated as RFC 5322 email format.
	Email *struct {
		Address string `json:"address"`
	} `json:"email,omitempty"`
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

// Chime represents a notification message.
//
// Chimes track delivery status, transmission details, and any errors
// that occurred during sending.
type Resource struct {
	// ID is the unique chime identifier (read-only).
	// Starts with "chm_". Example: "chm_abc123def456"
	ID string `json:"id,omitempty"`

	// CreatedAt is when the chime was created (ISO 8601, read-only).
	CreatedAt string `json:"created_at,omitempty"`

	// FullMessage is the message content that was sent.
	FullMessage string `json:"full_message,omitempty"`

	// Recipient contains the recipient details.
	Recipient *Recipient `json:"recipient,omitempty"`

	// SenderID is the sender name shown to the recipient.
	SenderID string `json:"sender_id,omitempty"`

	// Purpose is the chime category.
	Purpose string `json:"purpose,omitempty"`

	// CustomData contains attached custom data.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// Delivery contains delivery status and timestamps.
	// Tracks whether the message was successfully delivered.
	Delivery any `json:"delivery,omitempty"`

	// Transmission contains low-level transmission details.
	// Includes carrier responses, error codes, etc.
	Transmission any `json:"transmission,omitempty"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
