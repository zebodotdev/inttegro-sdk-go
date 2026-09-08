package price

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
)

type PageParams struct {
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	ProductID  string `json:"product_id,omitempty"`
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

// PriceParams is an inline price supplied in a request. It embeds the amount
// fields because the API's price shape is {currency, value}, not
// {amount: {currency, value}}.
type InlineParams struct {
	money.AmountParams
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
