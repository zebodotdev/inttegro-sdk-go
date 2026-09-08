package paymentmethod

// PaymentMethodDeletion confirms deletion.
type Deletion struct {
	// Deleted indicates whether deletion succeeded.
	Deleted bool `json:"deleted"`

	// PaymentMethodID is the deleted payment method's ID.
	PaymentMethodID string `json:"payment_method_id,omitempty"`
}
