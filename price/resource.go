package price

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/money"
	"github.com/zebodotdev/inttegro-sdk-go/v6/product"
)

type Page struct {
	Number int     `json:"number,omitempty"`
	Size   int     `json:"size,omitempty"`
	Prices []Price `json:"prices,omitempty"`
}

// Price is an inline price returned by the API.
type Inline struct {
	money.Amount
}

// Price represents a catalog price.
type Price struct {
	ID         string           `json:"id,omitempty"`
	Label      string           `json:"label,omitempty"`
	About      string           `json:"about,omitempty"`
	Active     bool             `json:"active"`
	Nominal    *money.Amount    `json:"nominal,omitempty"`
	ProductID  string           `json:"product_id,omitempty"`
	Product    *product.Product `json:"product,omitempty"`
	CreatedAt  *time.Time       `json:"created_at,omitempty"`
	UpdatedAt  *time.Time       `json:"updated_at,omitempty"`
	ArchivedAt *time.Time       `json:"archived_at,omitempty"`
}
