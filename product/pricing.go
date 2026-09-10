package product

import (
	"github.com/zebodotdev/inttegro-sdk-go/v7/money"
)

// ProductPriceSummary represents a product price listed with a product response.
type PriceSummary struct {
	ID      string       `json:"id"`
	Label   string       `json:"label,omitempty"`
	Nominal money.Amount `json:"nominal"`
	Active  bool         `json:"active"`
}
