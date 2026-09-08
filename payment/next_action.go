package payment

// PaymentNextAction describes the next step required to complete payment.
//
// When a payment requires customer action (OTP confirmation, redirect to
// bank page, etc), this field describes what needs to happen.
//
// Check Type to determine the required action:
//   - "confirm_payment": Customer must provide OTP
//   - "redirect": Customer must visit a URL
//   - "execute": Internal processing (no action needed)
type NextAction struct {
	// Type specifies the action category.
	// Values: "confirm_payment", "redirect", "execute"
	Type NextActionType `json:"type"`

	// ConfirmPayment contains OTP confirmation details.
	// Only present when Type is "confirm_payment".
	ConfirmPayment *struct {
		// ExpiresAt is when the OTP expires (ISO 8601).
		// Customer must confirm before this time.
		ExpiresAt string `json:"expires_at"`

		// Scheme describes the confirmation method.
		// Example: "otp"
		Scheme string `json:"scheme,omitempty"`

		// Request contains OTP delivery details.
		Request *struct {
			// ID is the OTP request identifier.
			ID string `json:"id"`

			// Recipient is where the OTP was sent.
			// For SMS: the phone number.
			Recipient string `json:"recipient"`

			// SentVia is the delivery channel.
			// Values: "sms", "email"
			SentVia ConfirmationChannel `json:"sent_via"`

			// TokenSize is the number of OTP digits.
			// Typically 4 or 6.
			TokenSize int `json:"token_size"`

			// SenderID is the SMS sender name.
			SenderID string `json:"sender_id"`
		} `json:"request,omitempty"`
	} `json:"confirm_payment,omitempty"`

	// Execute is present when Type is "execute" (internal processing).
	Execute any `json:"execute,omitempty"`

	// Redirect contains redirect details when Type is "redirect".
	Redirect *struct {
		// URL is where the customer should be redirected.
		// Open this URL in a browser for the customer to complete payment.
		URL string `json:"url"`
	} `json:"redirect,omitempty"`
}
