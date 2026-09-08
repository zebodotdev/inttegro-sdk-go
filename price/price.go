// Package price provides price resources and operations.
package price

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/product"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

// PricesService manages catalog prices.
type Service struct {
	client transport.Client
}

type PageParams struct {
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	ProductID  string `json:"product_id,omitempty"`
}

type Page struct {
	Number int        `json:"number,omitempty"`
	Size   int        `json:"size,omitempty"`
	Prices []Resource `json:"prices,omitempty"`
}

type ActionParams struct {
	PriceID string `json:"price_id"`
}

// AddToProductParams creates a new price for an existing product.
type AddToProductParams struct {
	ProductID    string             `json:"product_id"`
	Label        string             `json:"label,omitempty"`
	About        string             `json:"about,omitempty"`
	Amount       money.AmountParams `json:"amount"`
	SetAsDefault bool               `json:"set_as_default,omitempty"`
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

// PriceParams is an inline price supplied in a request. It embeds the amount
// fields because the API's price shape is {currency, value}, not
// {amount: {currency, value}}.
type InlineParams struct {
	money.AmountParams
}

// Price is an inline price returned by the API.
type Inline struct {
	money.Amount
}

// CatalogPriceParams creates a catalog price.
type CreateParams struct {
	ProductID string             `json:"product_id,omitempty"`
	Label     string             `json:"label,omitempty"`
	About     string             `json:"about,omitempty"`
	Amount    money.AmountParams `json:"amount"`
}

// LookupPriceParams looks up a price by ID.
type LookupParams struct {
	PriceID string `json:"price_id"`
}

// UpdatePriceParams updates price metadata.
type UpdateParams struct {
	PriceID   string `json:"price_id"`
	ProductID string `json:"product_id,omitempty"`
	Label     string `json:"label,omitempty"`
	About     string `json:"about,omitempty"`
}

// CatalogPrice represents a catalog price resource.
type Resource struct {
	ID         string            `json:"id,omitempty"`
	Label      string            `json:"label,omitempty"`
	About      string            `json:"about,omitempty"`
	Active     bool              `json:"active"`
	Nominal    *money.Amount     `json:"nominal,omitempty"`
	ProductID  string            `json:"product_id,omitempty"`
	Product    *product.Resource `json:"product,omitempty"`
	CreatedAt  string            `json:"created_at,omitempty"`
	UpdatedAt  string            `json:"updated_at,omitempty"`
	ArchivedAt string            `json:"archived_at,omitempty"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
