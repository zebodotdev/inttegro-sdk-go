package commerce

import "context"

// PricesService manages catalog prices.
type PricesService struct {
	client *Client
}

// Create creates a price.
func (s *PricesService) Create(ctx context.Context, params CreatePriceParams) (*Price, error) {
	var resp struct {
		Price Price `json:"price"`
	}
	if err := s.client.do(ctx, "POST", "/prices/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}

// Lookup retrieves a price by ID.
func (s *PricesService) Lookup(ctx context.Context, priceID string) (*Price, error) {
	var resp struct {
		Price Price `json:"price"`
	}
	if err := s.client.do(ctx, "POST", "/prices/lookup", LookupPriceParams{PriceID: priceID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}

// Update updates a price.
func (s *PricesService) Update(ctx context.Context, params UpdatePriceParams) (*Price, error) {
	var resp struct {
		Price Price `json:"price"`
	}
	if err := s.client.do(ctx, "POST", "/prices/update", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}
