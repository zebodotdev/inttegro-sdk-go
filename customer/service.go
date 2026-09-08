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

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
