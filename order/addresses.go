package order

import (
	"github.com/zebodotdev/inttegro-sdk-go/v7/customer"
)

// BillingDetails captures billing contact information for an order.
//
// Required for all orders. This information appears on invoices and
// is used for payment authorization and dispute resolution.
//
// The billing phone number and email are used for payment notifications
// and confirmation codes.
type BillingDetails struct {
	// Name is the billing contact's full name (required).
	// Typically matches the customer name.
	Name string `json:"name"`

	// Email is the billing email address (required).
	// Used for invoice delivery and payment receipts.
	// Validated as RFC 5322 email format.
	Email string `json:"email_address"`

	// PhoneNumber is the billing contact phone with country code (required).
	// Used for payment confirmations and delivery of OTP codes.
	// Example: "+233244123456"
	PhoneNumber string `json:"phone_number"`

	// Address is the billing postal address (required).
	// Must be a complete, valid address.
	Address customer.Address `json:"address"`
}

// Shipping holds shipping information for physical goods.
//
// Only required when selling physical products that need delivery.
// Digital goods and services don't require shipping information.
type Shipping struct {
	// Address is the delivery postal address (required for physical goods).
	// Must be a complete, deliverable address.
	Address customer.Address `json:"address"`
}
