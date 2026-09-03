package money

import (
	"encoding/json"
	"testing"
)

func TestAmountParamsUsesCurrencyWireValue(t *testing.T) {
	payload, err := json.Marshal(AmountParams{Currency: GHS, Value: 3005})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(payload), `{"currency":"ghs","value":3005}`; got != want {
		t.Fatalf("payload = %s, want %s", got, want)
	}
}
