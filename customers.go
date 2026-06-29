package commerce

import "context"

// CustomersService manages customer records.
type CustomersService struct {
	client *Client
}

// Create creates a customer record.
func (s *CustomersService) Create(ctx context.Context, params CreateCustomerParams) (*Customer, error) {
	var resp struct {
		Customer Customer `json:"customer"`
	}
	if err := s.client.do(ctx, "POST", "/customers/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Customer, nil
}

// Lookup retrieves a customer by ID.
func (s *CustomersService) Lookup(ctx context.Context, customerID string) (*Customer, error) {
	var resp struct {
		Customer Customer `json:"customer"`
	}
	if err := s.client.do(ctx, "POST", "/customers/lookup", LookupCustomerParams{CustomerID: customerID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Customer, nil
}

// Page retrieves a page of customers.
func (s *CustomersService) Page(ctx context.Context, params PageCustomersParams) (*CustomersPage, error) {
	var resp struct {
		Page CustomersPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/customers/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}
