package order

import "github.com/zebodotdev/inttegro-sdk-go/v7/payment"

// IsPaid reports whether the order has recorded payment, including a paid order
// that has since advanced to completed.
func (o Order) IsPaid() bool { return o.Status == StatusPaid || o.PaidAt != nil }

// RequiresPayment reports whether the order is waiting for payment.
func (o Order) RequiresPayment() bool { return o.Status == StatusRequiresPayment }

// IsTerminal reports whether the order has reached a final state.
func (o Order) IsTerminal() bool {
	switch o.Status {
	case StatusPaid, StatusCompleted, StatusCanceled, StatusExpired:
		return true
	default:
		return false
	}
}

// RequiredPaymentAction returns the nested payment action without requiring
// callers to traverse the order's payment projection.
func (o Order) RequiredPaymentAction() (*payment.NextAction, bool) {
	if o.Payment == nil {
		return nil, false
	}
	return o.Payment.RequiredAction()
}
