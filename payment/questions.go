package payment

// IsPaid reports whether the payment has completed successfully.
func (p Payment) IsPaid() bool { return p.Status == StatusPaid }

// RequiresAction reports whether the payment is waiting for customer or merchant action.
func (p Payment) RequiresAction() bool { return p.Status == StatusRequiresAction }

// IsTerminal reports whether the payment has reached a final state.
func (p Payment) IsTerminal() bool {
	switch p.Status {
	case StatusPaid, StatusCanceled, StatusExpired, StatusFailed:
		return true
	default:
		return false
	}
}

// RequiredAction returns the action details only while the payment requires action.
func (p Payment) RequiredAction() (*NextAction, bool) {
	if !p.RequiresAction() || p.NextAction == nil {
		return nil, false
	}
	return p.NextAction, true
}
