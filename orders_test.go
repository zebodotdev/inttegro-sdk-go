package inttegro

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zebodotdev/inttegro-sdk-go/v6/customer"
	"github.com/zebodotdev/inttegro-sdk-go/v6/order"
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
				"document_url":  "https://pages.inttegro.com/invoices/or_123",
				"sent_channels": []string{"sms"},
			},
		})
	}))
	if client == nil {
		return
	}
	defer close()

	ctx := context.Background()
	invoice, err := client.Orders.SendInvoice(ctx, order.SendInvoiceParams{OrderID: "or_123"})
	if err != nil {
		t.Fatal(err)
	}
	if invoice.Delivery.DocumentURL == "" {
		t.Fatalf("expected invoice delivery document URL")
	}
	if _, err := client.Orders.SendReceipt(ctx, order.SendReceiptParams{OrderID: "or_123"}); err != nil {
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

func TestOrdersPayReturnsOrder(t *testing.T) {
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/orders/pay" {
			t.Fatalf("path = %q, want /orders/pay", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"order": map[string]any{
				"id":     "or_123",
				"status": "requires_payment",
				"payment": map[string]any{
					"id":     "py_123",
					"status": "requires_action",
				},
			},
		})
	}))
	if client == nil {
		return
	}
	defer close()

	order, err := client.Orders.Pay(context.Background(), order.PayParams{OrderID: "or_123"})
	if err != nil {
		t.Fatal(err)
	}
	if order.ID != "or_123" {
		t.Fatalf("order.ID = %q, want or_123", order.ID)
	}
	if order.Payment == nil || order.Payment.ID != "py_123" {
		t.Fatalf("order.Payment = %#v, want payment py_123", order.Payment)
	}
}

func TestOrderCreateParamsOmitZeroBillingDetails(t *testing.T) {
	params := order.CreateParams{
		CustomerData: &customer.Data{Name: "Akua Mensah", Email: "akua@example.com", PhoneNumber: "+233544998605"},
		LineItems:    []order.LineItemParams{{Type: order.LineItemTypeProduct}},
	}
	if err := params.Validate(); err != nil {
		t.Fatalf("zero billing details should be optional: %v", err)
	}
	payload, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatal(err)
	}
	if _, exists := decoded["billing_details"]; exists {
		t.Fatalf("zero billing details must be omitted: %s", payload)
	}

	params.BillingDetails = order.BillingDetails{Name: "Akua Mensah"}
	if err := params.Validate(); err == nil {
		t.Fatal("partially supplied billing details must be rejected")
	}
}

func TestOrdersRequestConfirmationReturnsOrder(t *testing.T) {
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/orders/request_confirmation" {
			t.Fatalf("path = %q, want /orders/request_confirmation", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"order": map[string]any{
				"id":     "or_123",
				"status": "requires_payment",
				"payment": map[string]any{
					"id":     "py_123",
					"status": "requires_action",
					"next_action": map[string]any{
						"type": "confirm_payment",
					},
				},
			},
		})
	}))
	if client == nil {
		return
	}
	defer close()

	order, err := client.Orders.RequestConfirmation(context.Background(), "or_123")
	if err != nil {
		t.Fatal(err)
	}
	if order.ID != "or_123" {
		t.Fatalf("order.ID = %q, want or_123", order.ID)
	}
	if order.Payment == nil || order.Payment.NextAction == nil {
		t.Fatalf("order.Payment.NextAction missing: %#v", order.Payment)
	}
}
