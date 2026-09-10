package otp

// InitiateParams starts an OTP transaction and sends a verification code.
type InitiateParams struct {
	AsyncDelivery             *bool        `json:"async_delivery,omitempty"`
	MessageTemplate           string       `json:"message_template,omitempty"`
	Purpose                   string       `json:"purpose,omitempty"`
	Recipient                 string       `json:"recipient"`
	Sender                    string       `json:"sender,omitempty"`
	ServiceName               string       `json:"service_name"`
	TokenAlphabet             string       `json:"token_alphabet,omitempty"`
	TokenAlphabetType         AlphabetType `json:"token_alphabet_type,omitempty"`
	TokenSize                 int          `json:"token_size"`
	ValidityDurationInMinutes int          `json:"validity_duration_in_minutes,omitempty"`
}

// VerifyParams checks a token submitted for an OTP transaction.
type VerifyParams struct {
	Recipient     string `json:"recipient"`
	Token         string `json:"token"`
	TransactionID string `json:"transaction_id"`
}

// LookupParams identifies an OTP transaction to retrieve.
type LookupParams struct {
	TransactionID string `json:"transaction_id"`
}

// CancelParams identifies an OTP transaction to cancel.
type CancelParams struct {
	Reason        string `json:"reason,omitempty"`
	TransactionID string `json:"transaction_id"`
}
