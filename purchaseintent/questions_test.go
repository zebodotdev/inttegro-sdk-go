package purchaseintent

import "testing"

func TestPurchaseIntentQuestions(t *testing.T) {
	singleUse := true
	intent := PurchaseIntent{
		Status: StatusUsed,
		Usage:  Usage{SingleUse: &singleUse, Order: &UsageOrder{ID: "or_123"}},
	}

	if intent.IsActive() || !intent.IsSingleUse() {
		t.Fatal("used single-use intent reported the wrong state")
	}
	if got, ok := intent.UsedOrderID(); !ok || got != "or_123" {
		t.Fatalf("UsedOrderID() = (%q, %t), want (%q, true)", got, ok, "or_123")
	}
}
