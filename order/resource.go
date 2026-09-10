package order

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/payment"
	"github.com/zebodotdev/inttegro-sdk-go/v6/refund"
)

// Order is the complete public order projection returned by the API.
type Order struct {
	ID               string            `json:"id"`
	Status           Status            `json:"status"`
	Number           string            `json:"number,omitempty"`
	ReceiptNumber    string            `json:"receipt_number,omitempty"`
	Reference        string            `json:"reference,omitempty"`
	Customer         Customer          `json:"customer"`
	CheckoutSettings *CheckoutSettings `json:"checkout_settings,omitempty"`
	InvoiceSettings  *InvoiceSettings  `json:"invoice_settings,omitempty"`
	Invoice          *Invoice          `json:"invoice,omitempty"`
	LineItemGroup    *LineItemGroup    `json:"line_item_group,omitempty"`
	Payment          *payment.Payment  `json:"payment,omitempty"`
	CustomData       map[string]string `json:"custom_data,omitempty"`
	CreatedFrom      *CreatedFrom      `json:"created_from,omitempty"`
	InitiatedAt      time.Time         `json:"initiated_at"`
	CompletedAt      *time.Time        `json:"completed_at,omitempty"`
	SealedAt         *time.Time        `json:"sealed_at,omitempty"`
	PaidAt           *time.Time        `json:"paid_at,omitempty"`
	CanceledAt       *time.Time        `json:"canceled_at,omitempty"`
	ExpiresAt        *time.Time        `json:"expires_at,omitempty"`
	PaymentDueAt     *time.Time        `json:"payment_due_at,omitempty"`
	Refunds          []refund.Refund   `json:"refunds,omitempty"`
}

type CheckoutSettings struct {
	RedirectURL string `json:"redirect_url,omitempty"`
	CancelURL   string `json:"cancel_url,omitempty"`
}

type CreatedFrom struct {
	Source       string                  `json:"source,omitempty"`
	ResourceType CreatedFromResourceType `json:"resource_type,omitempty"`
	ResourceID   string                  `json:"resource_id,omitempty"`
}
