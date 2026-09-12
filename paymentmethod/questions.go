package paymentmethod

// IsArchived reports whether the payment method is archived.
func (p PaymentMethod) IsArchived() bool { return p.ArchivedAt != nil }

// IsVerified reports whether payment-method ownership has been verified.
func (p PaymentMethod) IsVerified() bool { return p.VerifiedAt != nil }

// IsReusable reports whether the payment method may be reused in new payment flows.
func (p PaymentMethod) IsReusable() bool {
	return p.Active && !p.IsArchived() && !p.Ephemeral
}
