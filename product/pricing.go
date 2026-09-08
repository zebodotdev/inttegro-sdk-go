package product

import (
	"github.com/zebodotdev/inttegro-sdk-go/v6/money"
)

// ProductDefaultUnitPrice represents a product's loaded default unit price.
type DefaultUnitPrice struct {
	ID         string        `json:"id,omitempty"`
	ProductID  string        `json:"product_id,omitempty"`
	Label      string        `json:"label,omitempty"`
	About      string        `json:"about,omitempty"`
	Nominal    *money.Amount `json:"nominal,omitempty"`
	CreatedAt  string        `json:"created_at,omitempty"`
	UpdatedAt  string        `json:"updated_at,omitempty"`
	ArchivedAt string        `json:"archived_at,omitempty"`
}

// ProductPriceSummary represents a product price listed with a product response.
type PriceSummary struct {
	ID      string        `json:"id,omitempty"`
	Label   string        `json:"label,omitempty"`
	Nominal *money.Amount `json:"nominal,omitempty"`
	Active  bool          `json:"active,omitempty"`
}
