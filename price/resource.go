package price

import (
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/product"
)

type Page struct {
	Number int        `json:"number,omitempty"`
	Size   int        `json:"size,omitempty"`
	Prices []Resource `json:"prices,omitempty"`
}

// Price is an inline price returned by the API.
type Inline struct {
	money.Amount
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
