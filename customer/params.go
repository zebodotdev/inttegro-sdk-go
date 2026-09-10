package customer

import (
	"github.com/zebodotdev/inttegro-sdk-go/v7/request"
)

// CreateCustomerParams creates a customer record.
type CreateParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// Name is the customer's full name (required).
	Name string `json:"name"`

	// Title is an honorific or title (optional).
	Title string `json:"title,omitempty"`

	// Suffix is a name suffix like Jr. (optional).
	Suffix string `json:"suffix,omitempty"`

	// Reference is your internal customer identifier (optional).
	Reference string `json:"reference,omitempty"`

	// Email is the customer's email address (optional).
	Email string `json:"email_address,omitempty"`

	// PhoneNumber is the customer's phone number (optional).
	PhoneNumber string `json:"phone_number,omitempty"`

	// CustomData holds arbitrary key-value pairs (optional).
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// UpdateCustomerParams replaces the supplied fields on an existing customer.
// Omitted fields remain unchanged.
type UpdateParams struct {
	CustomerID      string            `json:"customer_id"`
	BillingAddress  *Address          `json:"billing_address,omitempty"`
	CustomData      map[string]string `json:"custom_data,omitempty"`
	Email           string            `json:"email_address,omitempty"`
	Name            string            `json:"name,omitempty"`
	PhoneNumber     string            `json:"phone_number,omitempty"`
	Reference       string            `json:"reference,omitempty"`
	ShippingAddress *Address          `json:"shipping_address,omitempty"`
	Suffix          string            `json:"suffix,omitempty"`
	Title           string            `json:"title,omitempty"`
}

// LookupCustomerParams looks up a customer by ID.
type LookupParams struct {
	CustomerID string `json:"customer_id"`
}

// PageCustomersParams pages through customers.
type PageParams struct {
	PageNumber int `json:"page_number,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}
