package chime

// ChimeRecipient specifies who should receive a notification chime.
//
// Set Type and the corresponding field (Phone or Email).
//
// Example (SMS):
//
//	recipient := chime.Recipient{
//	    Type: chime.RecipientTypePhone,
//	    Name: "Jane Doe",
//	    Phone: &struct{Number string `json:"number"`}{Number: "+233244123456"},
//	}
type Recipient struct {
	// Type specifies how the recipient is identified (required).
	// Values: "phone" or "email"
	Type RecipientType `json:"type"`

	// Name is the recipient's display name (optional).
	// Used in email subject lines and SMS sender names.
	Name string `json:"name,omitempty"`

	// Phone contains the phone number (required when Type is "phone").
	// Must include country code. Example: "+233244123456"
	Phone *struct {
		Number string `json:"number"`
	} `json:"phone,omitempty"`

	// Email contains the email address (required when Type is "email").
	// Validated as RFC 5322 email format.
	Email *struct {
		Address string `json:"address"`
	} `json:"email,omitempty"`
}
