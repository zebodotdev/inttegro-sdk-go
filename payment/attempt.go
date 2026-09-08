package payment

// PaymentAttempt captures details of a single payment attempt.
//
// Each payment may have multiple attempts if initial attempts fail.
// This tracks the most recent attempt's status and timing.
type Attempt struct {
	// PaymentMethodType is the type of payment method used.
	PaymentMethodType string `json:"payment_method_type,omitempty"`

	// PaymentMethodID is the ID of the payment method charged.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// Reference is the external transaction reference from the payment provider.
	// Use this when contacting support or investigating payment issues.
	Reference string `json:"reference,omitempty"`

	// Status is the attempt's current state.
	// Values: "initiated", "succeeded", "failed"
	Status AttemptStatus `json:"status,omitempty"`

	// InitiatedAt is when the attempt started (ISO 8601).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// SucceededAt is when the attempt succeeded (ISO 8601).
	// Nil if not yet succeeded.
	SucceededAt *string `json:"succeeded_at,omitempty"`

	// FailedAt is when the attempt failed (ISO 8601).
	// Nil if not yet failed.
	FailedAt *string `json:"failed_at,omitempty"`
}
