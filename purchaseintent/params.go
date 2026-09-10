package purchaseintent

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/price"
)

type OriginalPriceParams struct {
	ID      string              `json:"id,omitempty"`
	Nominal *price.InlineParams `json:"nominal,omitempty"`
}

type CreateParams struct {
	Product   *ProductSelector `json:"product,omitempty"`
	ProductID string           `json:"product_id,omitempty"`
	Price     *PriceSelector   `json:"price,omitempty"`
	PriceID   string           `json:"price_id,omitempty"`
	Quantity  Quantity         `json:"quantity"`
	Usage     *Usage           `json:"usage,omitempty"`
	ExpiresAt *time.Time       `json:"expires_at,omitempty"`
}

type UpdateParams struct {
	ID         string    `json:"id"`
	Quantity   *Quantity `json:"quantity,omitempty"`
	ExpiresAt  any       `json:"expires_at,omitempty"`
	Reactivate *bool     `json:"reactivate,omitempty"`
}

type PageParams struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}
