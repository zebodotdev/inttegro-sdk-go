package payment

import "testing"

func TestPaymentQuestions(t *testing.T) {
	action := &NextAction{Type: NextActionTypeRedirect}
	payment := Payment{Status: StatusRequiresAction, NextAction: action}

	if payment.IsPaid() {
		t.Fatal("IsPaid() = true, want false")
	}
	if !payment.RequiresAction() {
		t.Fatal("RequiresAction() = false, want true")
	}
	if payment.IsTerminal() {
		t.Fatal("IsTerminal() = true, want false")
	}
	if got, ok := payment.RequiredAction(); !ok || got != action {
		t.Fatalf("RequiredAction() = (%#v, %t), want (%#v, true)", got, ok, action)
	}

	payment.Status = StatusPaid
	if !payment.IsPaid() || !payment.IsTerminal() {
		t.Fatal("paid payment should be paid and terminal")
	}
	if got, ok := payment.RequiredAction(); ok || got != nil {
		t.Fatalf("RequiredAction() = (%#v, %t), want (nil, false)", got, ok)
	}
}
