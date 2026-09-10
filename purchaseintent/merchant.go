package purchaseintent

// Merchant is the public merchant identity captured for the purchase intent.
type Merchant struct {
	AppName          string `json:"app_name,omitempty"`
	OrganizationID   string `json:"organization_id,omitempty"`
	OrganizationName string `json:"organization_name,omitempty"`
}
