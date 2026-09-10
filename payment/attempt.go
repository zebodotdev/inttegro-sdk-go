package payment

import "time"

// Attempt is the latest execution attempt for a payment.
type Attempt struct {
	PaymentMethodType string        `json:"payment_method_type,omitempty"`
	PaymentMethodID   string        `json:"payment_method_id,omitempty"`
	Error             *AttemptError `json:"error,omitempty"`
	Reference         string        `json:"reference,omitempty"`
	Status            AttemptStatus `json:"status"`
	InitiatedAt       time.Time     `json:"initiated_at"`
	SucceededAt       *time.Time    `json:"succeeded_at,omitempty"`
}

type AttemptError struct {
	Message string `json:"message"`
}
