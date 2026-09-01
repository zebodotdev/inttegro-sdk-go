package inttegro

import (
	"encoding/json"
	"testing"
)

func TestBalanceTransactionDeserializesSemanticSources(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantType   BalanceTransactionType
		wantSource string
	}{
		{
			name:       "payment",
			body:       `{"id":"bt_payment","type":"payment","payment_id":"py_123","order_id":"or_123","amount":{"currency":"GHS","value":2500},"created_at":"2026-08-31T12:00:00Z"}`,
			wantType:   BalanceTransactionTypePayment,
			wantSource: "py_123",
		},
		{
			name:       "refund",
			body:       `{"id":"bt_refund","type":"refund","refund_id":"rf_123","order_id":"or_123","amount":{"currency":"GHS","value":500},"created_at":"2026-08-31T12:01:00Z"}`,
			wantType:   BalanceTransactionTypeRefund,
			wantSource: "rf_123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var txn BalanceTransaction
			if err := json.Unmarshal([]byte(tt.body), &txn); err != nil {
				t.Fatalf("unmarshal balance transaction: %v", err)
			}
			if txn.Type != tt.wantType {
				t.Fatalf("type = %q, want %q", txn.Type, tt.wantType)
			}
			if sourceID, ok := txn.SourceID(); !ok || sourceID != tt.wantSource {
				t.Fatalf("SourceID() = (%q, %t), want (%q, true)", sourceID, ok, tt.wantSource)
			}
			if txn.ID == "" || txn.OrderID == "" || txn.Amount.Currency == "" || txn.CreatedAt == "" {
				t.Fatalf("required fields were not decoded: %#v", txn)
			}
		})
	}
}

func TestBalanceTransactionSourceIDRejectsContradictoryReferences(t *testing.T) {
	txn := BalanceTransaction{
		Type:      BalanceTransactionTypeRefund,
		PaymentID: "py_123",
		RefundID:  "rf_123",
	}
	if sourceID, ok := txn.SourceID(); ok || sourceID != "" {
		t.Fatalf("SourceID() = (%q, %t), want empty/false", sourceID, ok)
	}
}

func TestPaymentDeserializesCanonicalEmbeddedBalanceTransaction(t *testing.T) {
	var payment Payment
	body := `{"id":"py_123","balance_transaction":{"id":"bt_payment","type":"payment","payment_id":"py_123","order_id":"or_123","amount":{"currency":"GHS","value":2500},"created_at":"2026-08-31T12:00:00Z"}}`
	if err := json.Unmarshal([]byte(body), &payment); err != nil {
		t.Fatalf("unmarshal payment: %v", err)
	}
	if payment.BalanceTransaction == nil {
		t.Fatal("embedded balance transaction is nil")
	}
	if sourceID, ok := payment.BalanceTransaction.SourceID(); !ok || sourceID != payment.ID {
		t.Fatalf("embedded SourceID() = (%q, %t), want (%q, true)", sourceID, ok, payment.ID)
	}
}
