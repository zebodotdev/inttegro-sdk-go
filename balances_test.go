package inttegro

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBalancesGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/balances" {
			t.Fatalf("expected /balances, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"balances": map[string]any{
				"ghs": map[string]any{
					"available":                    map[string]any{"amount": 1000},
					"pending":                      map[string]any{"amount": 200},
					"includes_transactions_before": "2024-01-01T00:00:00Z",
				},
			},
		})
	}))
	defer srv.Close()

	client := NewClient("sk_test", WithBaseURL(srv.URL))
	resp, err := client.Balances.Get(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Balances["ghs"].Available.Amount != 1000 {
		t.Fatalf("unexpected amount: %+v", resp.Balances["ghs"])
	}
}
