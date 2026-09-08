package order

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/checkout"
	"github.com/zebodotdev/inttegro-sdk-go/v5/customer"
	"github.com/zebodotdev/inttegro-sdk-go/v5/invoice"
	"github.com/zebodotdev/inttegro-sdk-go/v5/payment"
	"github.com/zebodotdev/inttegro-sdk-go/v5/refund"
)

// Order represents a complete order object.
//
// Orders are the central resource in Inttegro, representing a purchase
// transaction from cart to fulfillment. Orders go through several states:
//
// 1. preparing: Created but not finalized
// 2. requires_payment: Finalized and ready for payment
// 3. paid: Payment succeeded
// 4. completed: Fulfilled and settled
// 5. canceled: Permanently canceled
// 6. expired: Payment window elapsed
type Resource struct {
	// ID is the unique order identifier (read-only).
	// Starts with "or_". Example: "or_abc123def456"
	ID string `json:"id"`

	// Status is the order's current state (read-only).
	// Values: "preparing", "requires_payment", "paid", "completed", "canceled", "expired", "unknown"
	Status Status `json:"status"`

	// Number is the human-readable order number.
	// Auto-generated if not provided during creation.
	// Example: "ORD-2023-00123"
	Number string `json:"number,omitempty"`

	// CustomerID is the owning customer's ID.
	CustomerID string `json:"customer_id,omitempty"`

	// Customer contains full customer details.
	// Only populated in responses when customer exists.
	Customer *customer.Data `json:"customer,omitempty"`

	// BillingDetails holds billing contact and address.
	BillingDetails *BillingDetails `json:"billing_details,omitempty"`

	// Shipping holds delivery address for physical goods.
	// Nil for digital-only orders.
	Shipping *Shipping `json:"shipping,omitempty"`

	// LineItems is the list of products, fees, and shipping charges.
	LineItems []LineItem `json:"line_items,omitempty"`

	// LineItemGroup contains line items grouped with totals.
	// Useful for displaying cart breakdowns.
	LineItemGroup *LineItemGroup `json:"line_item_group,omitempty"`

	// Payment contains payment details and status.
	// Nil if order hasn't been charged yet.
	Payment *payment.Resource `json:"payment,omitempty"`

	// PaymentStatus is a summary of payment state (read-only).
	// Values: "unpaid", "requires_action", "processing", "paid", "failed"
	PaymentStatus string `json:"payment_status,omitempty"`

	// PaymentMethodID is the attached payment method's ID.
	// May be set but not yet charged.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// StatementDescriptor appears on payment statements.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// CheckoutSettings contains hosted checkout configuration.
	CheckoutSettings *checkout.Settings `json:"checkout_settings,omitempty"`

	// InitiatedAt is when the order was created (ISO 8601, read-only).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// SealedAt is when the order was finalized (ISO 8601, read-only).
	// Nil if not yet finalized.
	SealedAt *string `json:"sealed_at,omitempty"`

	// CompletedAt is when the order was marked complete (ISO 8601, read-only).
	// Nil if not yet completed.
	CompletedAt *string `json:"completed_at,omitempty"`

	// ExpiresAt is when the order expires if unpaid (ISO 8601, read-only).
	// Nil for completed orders.
	ExpiresAt *string `json:"expires_at,omitempty"`

	// CreatedAt is the creation timestamp (ISO 8601, read-only).
	CreatedAt *string `json:"created_at,omitempty"`

	// UpdatedAt is the last modification timestamp (ISO 8601, read-only).
	UpdatedAt *string `json:"updated_at,omitempty"`

	// PaidAt is when payment was confirmed (ISO 8601, read-only).
	// Nil if not yet paid.
	PaidAt *string `json:"paid_at,omitempty"`

	// CancelledAt is when the order was canceled (ISO 8601, read-only).
	// Nil if not canceled.
	CancelledAt *string `json:"cancelled_at,omitempty"`

	// Invoice contains invoice document links and delivery status.
	// Nil if order not finalized.
	Invoice *invoice.Resource `json:"invoice,omitempty"`

	// Refunds contains every refund issued for this order, newest first.
	// It is omitted when no refunds exist.
	Refunds []refund.Resource `json:"refunds,omitempty"`
}
