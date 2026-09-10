package paymentmethod

import (
	"testing"
	"time"
)

func TestPaymentMethodQuestions(t *testing.T) {
	verifiedAt := time.Now()
	method := PaymentMethod{Active: true, VerifiedAt: &verifiedAt}
	if method.IsArchived() || !method.IsVerified() || !method.IsReusable() {
		t.Fatal("reusable verified method reported the wrong state")
	}

	method.Ephemeral = true
	if method.IsReusable() {
		t.Fatal("ephemeral method should not be reusable")
	}
}
