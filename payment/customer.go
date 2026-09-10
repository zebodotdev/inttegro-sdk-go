package payment

type Customer struct {
	ID              string   `json:"id"`
	EmailAddress    string   `json:"email_address,omitempty"`
	Guest           bool     `json:"guest"`
	Name            string   `json:"name"`
	PhoneNumber     string   `json:"phone_number,omitempty"`
	BillingAddress  *Address `json:"billing_address,omitempty"`
	ShippingAddress *Address `json:"shipping_address,omitempty"`
}

type Address struct {
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Line1       string `json:"line1,omitempty"`
	Line2       string `json:"line2,omitempty"`
	City        string `json:"city,omitempty"`
	Region      string `json:"region,omitempty"`
	PostCode    string `json:"post_code,omitempty"`
	Country     string `json:"country"`
}

type BillingDetails struct {
	Owner *PaymentMethodOwner `json:"owner,omitempty"`
}
