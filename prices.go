package inttegro

import "context"

// PricesService manages catalog prices.
type PricesService struct {
	client *Client
}

type PricePageParams struct {
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	ProductID  string `json:"product_id,omitempty"`
}

type PricesPage struct {
	Number int     `json:"number,omitempty"`
	Size   int     `json:"size,omitempty"`
	Prices []Price `json:"prices,omitempty"`
}

type PriceActionParams struct {
	PriceID string `json:"price_id"`
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

// Page retrieves a paginated list of prices.
func (s *PricesService) Page(ctx context.Context, params PricePageParams) (*PricesPage, error) {
	var resp struct {
		Page PricesPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/prices/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
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

// Activate reactivates an inactive price.
func (s *PricesService) Activate(ctx context.Context, priceID string) (*Price, error) {
	return s.priceAction(ctx, "/prices/activate", priceID)
}

// Deactivate marks a price inactive.
func (s *PricesService) Deactivate(ctx context.Context, priceID string) (*Price, error) {
	return s.priceAction(ctx, "/prices/deactivate", priceID)
}

// Archive permanently archives a price and marks it inactive.
func (s *PricesService) Archive(ctx context.Context, priceID string, opts ...RequestOption) (*Price, error) {
	var resp struct {
		Price Price `json:"price"`
	}
	if err := s.client.doJSON(ctx, "/prices/archive", PriceActionParams{PriceID: priceID}, applyRequestOptions(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}

func (s *PricesService) priceAction(ctx context.Context, path, priceID string) (*Price, error) {
	var resp struct {
		Price Price `json:"price"`
	}
	if err := s.client.do(ctx, "POST", path, PriceActionParams{PriceID: priceID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}
