package product

import "time"

// Product is a product returned by the catalog API.
type Product struct {
	ID          string            `json:"id"`
	About       string            `json:"about,omitempty"`
	Active      bool              `json:"active"`
	ArchivedAt  *time.Time        `json:"archived_at,omitempty"`
	Attributes  []Attribute       `json:"attributes,omitempty"`
	Category    string            `json:"category,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	CustomData  map[string]string `json:"custom_data,omitempty"`
	Description string            `json:"description,omitempty"`
	Dimensions  *Dimensions       `json:"dimensions,omitempty"`
	Media       *Media            `json:"media,omitempty"`
	Name        string            `json:"name"`
	Prices      []PriceSummary    `json:"prices,omitempty"`
	PublishedAt *time.Time        `json:"published_at,omitempty"`
	Reference   string            `json:"reference,omitempty"`
	Shipment    *Shipment         `json:"shipment,omitempty"`
	TaxCode     string            `json:"tax_code,omitempty"`
	Type        Type              `json:"type"`
	UnitDim     string            `json:"unit_dim,omitempty"`
	UpdatedAt   *time.Time        `json:"updated_at,omitempty"`
}

// Page holds a page of products.
type Page struct {
	Number   int       `json:"number"`
	Size     int       `json:"size"`
	Products []Product `json:"products"`
}
