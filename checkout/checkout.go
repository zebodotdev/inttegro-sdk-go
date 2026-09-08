// Package checkout provides checkout resources and operations.
package checkout

type OrderStatus string

const (
	OrderStatusPreparing OrderStatus = "preparing"
)

const (
	OrderStatusRequiresPayment OrderStatus = "requires_payment"
)

const (
	OrderStatusCompleted OrderStatus = "completed"
)

const (
	OrderStatusCanceled OrderStatus = "canceled"
)

const (
	OrderStatusExpired OrderStatus = "expired"
)

type PaymentStatus string

const (
	PaymentStatusRequiresAction PaymentStatus = "requires_action"
)

const (
	PaymentStatusProcessing PaymentStatus = "processing"
)

const (
	PaymentStatusSucceeded PaymentStatus = "succeeded"
)

const (
	PaymentStatusFailed PaymentStatus = "failed"
)

const (
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

// CheckoutSettings configures checkout page behavior and redirect URLs.
//
// When you finalize an order (finalize: true), Inttegro
// generates a hosted checkout page. These settings control where customers
// are redirected after completing or canceling payment.
type Settings struct {
	// RedirectURL is where customers go after successful payment (optional).
	// Must be HTTPS in production. Can be HTTP for testing.
	// If omitted, customers see a generic success page.
	// Example: "https://example.com/orders/thank-you"
	RedirectURL string `json:"redirect_url,omitempty"`

	// CancelURL is where customers go if they cancel payment (optional).
	// Must be HTTPS in production. Can be HTTP for testing.
	// If omitted, customers see a generic cancellation page.
	// Example: "https://example.com/cart"
	CancelURL string `json:"cancel_url,omitempty"`
}
