package purchaseintent

type VariantSet struct {
	Active           bool          `json:"active"`
	DefaultProductID string        `json:"default_product_id,omitempty"`
	Description      string        `json:"description,omitempty"`
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Reference        string        `json:"reference,omitempty"`
	VariantAxes      []VariantAxis `json:"variant_axes"`
	Variants         []Variant     `json:"variants"`
}

type VariantAxis struct {
	Key      string `json:"key"`
	Label    string `json:"label"`
	Position int    `json:"position"`
}

type Variant struct {
	Active        bool              `json:"active"`
	Position      int               `json:"position,omitempty"`
	Price         *Price            `json:"price,omitempty"`
	Product       *Product          `json:"product,omitempty"`
	ProductID     string            `json:"product_id"`
	VariantValues map[string]string `json:"variant_values"`
}
