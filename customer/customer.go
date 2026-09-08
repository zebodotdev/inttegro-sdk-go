// Package customer provides customer resources and operations.
package customer

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

// CustomersService manages customer records.
type Service struct {
	client transport.Client
}

// Create creates a customer record.
func (s *Service) Create(ctx context.Context, params CreateParams) (*Resource, error) {
	var resp struct {
		Customer Resource `json:"customer"`
	}
	if err := s.client.Do(ctx, "POST", "/customers/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Customer, nil
}

// Update replaces the supplied fields on a customer record.
func (s *Service) Update(ctx context.Context, params UpdateParams, opts ...request.Option) (*Resource, error) {
	var resp struct {
		Customer Resource `json:"customer"`
	}
	if err := s.client.DoJSON(ctx, "/customers/update", params, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.Customer, nil
}

// Lookup retrieves a customer by ID.
func (s *Service) Lookup(ctx context.Context, customerID string) (*Resource, error) {
	var resp struct {
		Customer Resource `json:"customer"`
	}
	if err := s.client.Do(ctx, "POST", "/customers/lookup", LookupParams{CustomerID: customerID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Customer, nil
}

// Page retrieves a page of customers.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/customers/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// CustomerData captures inline customer information for order creation.
//
// Use this when creating orders for new customers who don't have a customer ID yet.
// The API will create a customer record automatically and return it in the response.
//
// For existing customers, pass CustomerID instead to link the order to their record.
//
// Example:
//
//	customerData := &customer.Data{
//	    Name:        "Jane Doe",
//	    Email:       "jane@example.com",
//	    PhoneNumber: "+233244123456",
//	    Reference:   "customer_123_in_my_system",
//	}
type Data struct {
	// Name is the customer's full name (required).
	// Used for billing records and invoice generation.
	Name string `json:"name"`

	// Email is the customer's email address (optional but recommended).
	// Used for sending invoice links and payment notifications.
	// Validated as RFC 5322 email format.
	Email string `json:"email_address,omitempty"`

	// PhoneNumber is the customer's phone number with country code (required).
	// Used for SMS notifications and payment confirmations.
	// Example: "+233244123456"
	PhoneNumber string `json:"phone_number"`

	// Reference is your internal customer identifier (optional).
	// Use this to link Inttegro customer records to your system's users.
	// Maximum 255 characters. Must be unique across your customers.
	Reference string `json:"reference,omitempty"`

	// CustomData holds arbitrary key-value pairs about the customer (optional).
	// Both keys and values must be strings. Maximum 25KB when serialized.
	CustomData map[string]string `json:"custom_data,omitempty"`
}

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

// Customer represents a customer record.
type Resource struct {
	ID          string            `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Title       string            `json:"title,omitempty"`
	Suffix      string            `json:"suffix,omitempty"`
	Reference   string            `json:"reference,omitempty"`
	Email       string            `json:"email_address,omitempty"`
	PhoneNumber string            `json:"phone_number,omitempty"`
	CustomData  map[string]string `json:"custom_data,omitempty"`
	CreatedAt   string            `json:"created_at,omitempty"`
}

// CustomersPage holds a page of customers.
type Page struct {
	Number    int        `json:"number,omitempty"`
	Size      int        `json:"size,omitempty"`
	Customers []Resource `json:"customers,omitempty"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
