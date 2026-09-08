package customer

// Address represents a postal address used for billing and shipping.
type Address struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2,omitempty"`
	Town        string `json:"town"`
	Region      string `json:"region,omitempty"`
	District    string `json:"district,omitempty"`
	Country     string `json:"country"`
	PostCode    string `json:"post_code,omitempty"`
}
