package paymentmethod

// PaymentMethodVerificationSession contains verification state and delivery details.
type VerificationSession struct {
	PaymentMethodID string                `json:"payment_method_id,omitempty"`
	Status          string                `json:"status,omitempty"`
	TokenSentAt     *string               `json:"token_sent_at,omitempty"`
	ExpiresAt       *string               `json:"expires_at,omitempty"`
	Delivery        *VerificationDelivery `json:"delivery,omitempty"`
}

// PaymentMethodVerificationDelivery describes where a verification token was sent.
type VerificationDelivery struct {
	Recipient string `json:"recipient,omitempty"`
	Channel   string `json:"channel,omitempty"`
	SenderID  string `json:"sender_id,omitempty"`
}
