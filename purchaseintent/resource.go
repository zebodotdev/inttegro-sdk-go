package purchaseintent

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/product"
)

type Resource struct {
	ID                 string            `json:"id"`
	ProductID          string            `json:"product_id"`
	PriceID            string            `json:"price_id"`
	Quantity           Quantity          `json:"quantity"`
	AdjustableQuantity bool              `json:"adjustable_quantity"`
	AllowVariants      bool              `json:"allow_variants"`
	Status             Status            `json:"status"`
	CreatedAt          string            `json:"created_at"`
	UpdatedAt          string            `json:"updated_at,omitempty"`
	Activity           *ActivityLog      `json:"activity,omitempty"`
	Product            *product.Resource `json:"product,omitempty"`
	Price              *Price            `json:"price,omitempty"`
}

type Page struct {
	Number          int        `json:"number,omitempty"`
	Size            int        `json:"size,omitempty"`
	PurchaseIntents []Resource `json:"purchase_intents,omitempty"`
}
