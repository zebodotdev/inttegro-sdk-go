package product

// CreateProductParams creates a catalog product.
type CreateParams struct {
	Type        Type              `json:"type"`
	Reference   string            `json:"reference,omitempty"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	About       string            `json:"about,omitempty"`
	TaxCode     string            `json:"tax_code,omitempty"`
	Category    *Category         `json:"category,omitempty"`
	Shipment    *ShipmentInput    `json:"shipment,omitempty"`
	Media       []MediaItem       `json:"media,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	CustomData  map[string]string `json:"custom_data,omitempty"`
}

// LookupProductParams looks up a product by ID.
type LookupParams struct {
	ProductID string `json:"product_id"`
}

// UpdateProductParams updates a product.
type UpdateParams struct {
	ProductID   string            `json:"product_id"`
	Reference   string            `json:"reference,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	About       string            `json:"about,omitempty"`
	TaxCode     string            `json:"tax_code,omitempty"`
	Category    *Category         `json:"category,omitempty"`
	Shipment    *ShipmentInput    `json:"shipment,omitempty"`
	Media       []MediaItem       `json:"media,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	CustomData  map[string]string `json:"custom_data,omitempty"`
}

// ProductActionParams performs an action on a product.
type ActionParams struct {
	ProductID string `json:"product_id"`
}

// SetDefaultUnitPriceParams changes the price used as a product's default.
type SetDefaultUnitPriceParams struct {
	ProductID string `json:"product_id"`
	PriceID   string `json:"price_id"`
}

// PageProductsParams pages through products.
type PageParams struct {
	PageNumber int `json:"page_number,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}
