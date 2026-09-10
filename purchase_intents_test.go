package inttegro

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPurchaseIntentLookupReturnsTypedResource(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/purchase_intents/lookup" {
			t.Fatalf("expected /purchase_intents/lookup, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"purchase_intent": map[string]any{
				"activity": map[string]any{"recent": []any{map[string]any{
					"created_at":         "2026-09-09T12:01:00Z",
					"id":                 "saleevt_123",
					"purchase_intent_id": "sale_123",
					"type":               "viewed",
					"visitor":            map[string]any{"ip_address": "203.0.113.7"},
				}}},
				"allow_variants": false,
				"created_at":     "2026-09-09T12:00:00Z",
				"id":             "sale_123",
				"merchant":       map[string]any{"organization_name": "Tea House Ltd"},
				"product": map[string]any{
					"active":     true,
					"created_at": "2026-09-09T11:00:00Z",
					"dimensions": map[string]any{"digital": map[string]any{"bytes": 1024}},
					"id":         "prod_123",
					"name":       "Tea guide",
					"type":       "digital",
				},
				"quantity": map[string]any{"min": 1},
				"status":   "active",
				"usage": map[string]any{
					"order":      map[string]any{"created_at": "2026-09-09T12:02:00Z", "id": "or_123"},
					"single_use": true,
				},
			},
		})
	}))
	defer srv.Close()

	client := NewClient("sk_test", WithBaseURL(srv.URL))
	intent, err := client.PurchaseIntents.Lookup(context.Background(), "sale_123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := intent.Activity.Recent[0].Visitor.IPAddress; got != "203.0.113.7" {
		t.Fatalf("visitor IP address = %q", got)
	}
	if got := intent.Merchant.OrganizationName; got != "Tea House Ltd" {
		t.Fatalf("merchant organization name = %q", got)
	}
	if got := intent.Product.Dimensions.Digital.Bytes; got != 1024 {
		t.Fatalf("product byte size = %v", got)
	}
	if got := intent.Usage.Order.ID; got != "or_123" {
		t.Fatalf("usage order ID = %q", got)
	}
}
