package purchaseintent

import (
	"github.com/zebodotdev/inttegro-sdk-go/v7/price"
)

type Quantity struct {
	Min int `json:"min"`
	Max int `json:"max,omitempty"`
}

type ProductSelector struct {
	ID           string `json:"id"`
	VariantSetID string `json:"variant_set_id,omitempty"`
}

type PriceSelector struct {
	ID         string               `json:"id,omitempty"`
	Nominal    *price.InlineParams  `json:"nominal,omitempty"`
	Original   *OriginalPriceParams `json:"original,omitempty"`
	OriginalID string               `json:"original_id,omitempty"`
}
