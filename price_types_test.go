package inttegro

import (
	"encoding/json"
	"testing"

	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/price"
)

func TestPriceParamsEmbedsAmountOnTheWire(t *testing.T) {
	payload, err := json.Marshal(price.InlineParams{
		AmountParams: money.AmountParams{Currency: money.GHS, Value: 3005},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(payload), `{"currency":"ghs","value":3005}`; got != want {
		t.Fatalf("json.Marshal(PriceParams) = %s, want %s", got, want)
	}
}

func TestCatalogPriceRetainsReferencedProductID(t *testing.T) {
	var price price.Resource
	if err := json.Unmarshal([]byte(`{"id":"pr_123","active":true,"nominal":{"currency":"ghs","value":3005},"product_id":"prod_123","created_at":"2026-09-02T12:00:00Z"}`), &price); err != nil {
		t.Fatal(err)
	}
	if got, want := price.ProductID, "prod_123"; got != want {
		t.Fatalf("CatalogPrice.ProductID = %q, want %q", got, want)
	}
}
