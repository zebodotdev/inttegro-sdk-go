package commerce

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestOrderDocumentDeliveryEndpointsMatchSpec(t *testing.T) {
	var paths []string
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"order": map[string]any{"id": "or_123"},
			"delivery": map[string]any{
				"document_kind": "invoice",
				"document_url":  "https://pages.zebo.dev/invoices/or_123",
				"sent_channels": []string{"sms"},
			},
		})
	}))
	if client == nil {
		return
	}
	defer close()

	ctx := context.Background()
	invoice, err := client.Orders.SendInvoice(ctx, OrderSendInvoiceParams{OrderID: "or_123"})
	if err != nil {
		t.Fatal(err)
	}
	if invoice.Delivery.DocumentURL == "" {
		t.Fatalf("expected invoice delivery document URL")
	}
	if _, err := client.Orders.SendReceipt(ctx, OrderSendReceiptParams{OrderID: "or_123"}); err != nil {
		t.Fatal(err)
	}

	want := []string{"/orders/send_invoice", "/orders/send_receipt"}
	if len(paths) != len(want) {
		t.Fatalf("got paths %v, want %v", paths, want)
	}
	for i := range want {
		if paths[i] != want[i] {
			t.Fatalf("path %d: got %q, want %q", i, paths[i], want[i])
		}
	}
}
