package order

import (
	"testing"
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v7/payment"
)

func TestOrderQuestions(t *testing.T) {
	action := &payment.NextAction{Type: payment.NextActionTypeConfirmPayment}
	order := Order{
		Status: StatusRequiresPayment,
		Payment: &payment.Payment{
			Status:     payment.StatusRequiresAction,
			NextAction: action,
		},
	}

	if order.IsPaid() || !order.RequiresPayment() || order.IsTerminal() {
		t.Fatal("requires-payment order reported the wrong lifecycle state")
	}
	if got, ok := order.RequiredPaymentAction(); !ok || got != action {
		t.Fatalf("RequiredPaymentAction() = (%#v, %t), want (%#v, true)", got, ok, action)
	}

	paidAt := time.Now()
	order.Status = StatusCompleted
	order.PaidAt = &paidAt
	if !order.IsPaid() || !order.IsTerminal() {
		t.Fatal("completed paid order should be paid and terminal")
	}
}
