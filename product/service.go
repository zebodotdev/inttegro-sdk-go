package product

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v7/internal/transport"
)

// ProductsService manages catalog products.
type Service struct {
	client transport.Client
}

// Create creates a product.
func (s *Service) Create(ctx context.Context, params CreateParams) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// SetDefaultUnitPrice sets an existing product price as the product's default unit price.
func (s *Service) SetDefaultUnitPrice(ctx context.Context, params SetDefaultUnitPriceParams) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/set_default_unit_price", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Lookup retrieves a product by ID.
func (s *Service) Lookup(ctx context.Context, productID string) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/lookup", LookupParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Page retrieves a page of products.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/products/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Update updates a product.
func (s *Service) Update(ctx context.Context, params UpdateParams) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/update", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Publish publishes a product.
func (s *Service) Publish(ctx context.Context, productID string) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/publish", ActionParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Unpublish unpublishes a product.
func (s *Service) Unpublish(ctx context.Context, productID string) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/unpublish", ActionParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Archive archives a product.
func (s *Service) Archive(ctx context.Context, productID string) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.Do(ctx, "POST", "/products/archive", ActionParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
