package price

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

// PricesService manages catalog prices.
type Service struct {
	client transport.Client
}

// Create creates a price.
func (s *Service) Create(ctx context.Context, params CreateParams) (*Resource, error) {
	var resp struct {
		Price Resource `json:"price"`
	}
	if err := s.client.Do(ctx, "POST", "/prices/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}

// AddToProduct creates a price through the product-scoped endpoint.
func (s *Service) AddToProduct(ctx context.Context, params AddToProductParams) (*Resource, error) {
	var response struct {
		Price Resource `json:"price"`
	}
	if err := s.client.Do(ctx, "POST", "/products/add_price", params, &response); err != nil {
		return nil, err
	}
	return &response.Price, nil
}

// Lookup retrieves a price by ID.
func (s *Service) Lookup(ctx context.Context, priceID string) (*Resource, error) {
	var resp struct {
		Price Resource `json:"price"`
	}
	if err := s.client.Do(ctx, "POST", "/prices/lookup", LookupParams{PriceID: priceID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}

// Page retrieves a paginated list of prices.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/prices/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Update updates a price.
func (s *Service) Update(ctx context.Context, params UpdateParams) (*Resource, error) {
	var resp struct {
		Price Resource `json:"price"`
	}
	if err := s.client.Do(ctx, "POST", "/prices/update", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}

// Activate reactivates an inactive price.
func (s *Service) Activate(ctx context.Context, priceID string) (*Resource, error) {
	return s.priceAction(ctx, "/prices/activate", priceID)
}

// Deactivate marks a price inactive.
func (s *Service) Deactivate(ctx context.Context, priceID string) (*Resource, error) {
	return s.priceAction(ctx, "/prices/deactivate", priceID)
}

// Archive permanently archives a price and marks it inactive.
func (s *Service) Archive(ctx context.Context, priceID string, opts ...request.Option) (*Resource, error) {
	var resp struct {
		Price Resource `json:"price"`
	}
	if err := s.client.DoJSON(ctx, "/prices/archive", ActionParams{PriceID: priceID}, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}

func (s *Service) priceAction(ctx context.Context, path, priceID string) (*Resource, error) {
	var resp struct {
		Price Resource `json:"price"`
	}
	if err := s.client.Do(ctx, "POST", path, ActionParams{PriceID: priceID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
