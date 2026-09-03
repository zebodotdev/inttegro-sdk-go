package inttegro

import (
	"encoding/json"
	"testing"

	"github.com/zebodotdev/inttegro-sdk-go/v4/money"
)

func TestPriceParamsEmbedsAmountOnTheWire(t *testing.T) {
	payload, err := json.Marshal(PriceParams{
		AmountParams: money.AmountParams{Currency: money.GHS, Value: 3005},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(payload), `{"currency":"ghs","value":3005}`; got != want {
		t.Fatalf("json.Marshal(PriceParams) = %s, want %s", got, want)
	}
}
