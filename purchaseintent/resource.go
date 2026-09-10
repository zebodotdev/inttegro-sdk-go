package purchaseintent

import "time"

// PurchaseIntent represents a customer's intent to purchase a product.
type PurchaseIntent struct {
	Activity      *ActivityLog `json:"activity,omitempty"`
	AllowVariants bool         `json:"allow_variants"`
	CreatedAt     time.Time    `json:"created_at"`
	ExpiresAt     *time.Time   `json:"expires_at,omitempty"`
	ID            string       `json:"id"`
	InactiveAt    *time.Time   `json:"inactive_at,omitempty"`
	Merchant      *Merchant    `json:"merchant,omitempty"`
	Price         *Price       `json:"price,omitempty"`
	Product       *Product     `json:"product,omitempty"`
	Quantity      Quantity     `json:"quantity"`
	Status        Status       `json:"status"`
	UpdatedAt     *time.Time   `json:"updated_at,omitempty"`
	Usage         Usage        `json:"usage"`
	VariantSet    *VariantSet  `json:"variant_set,omitempty"`
}

type Page struct {
	Number          int              `json:"number"`
	Size            int              `json:"size"`
	PurchaseIntents []PurchaseIntent `json:"purchase_intents"`
}
