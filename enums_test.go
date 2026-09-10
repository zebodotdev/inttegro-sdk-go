package inttegro

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/zebodotdev/inttegro-sdk-go/v7/product"
	"github.com/zebodotdev/inttegro-sdk-go/v7/refund"
	"github.com/zebodotdev/inttegro-sdk-go/v7/uploadrequest"
)

func TestEnumConstantsSerializeAsWireValues(t *testing.T) {
	payload := struct {
		Product product.Type         `json:"product"`
		Refund  refund.Reason        `json:"refund"`
		Status  uploadrequest.Status `json:"status"`
	}{
		Product: product.TypeDigital,
		Refund:  refund.ReasonRequestedByCustomer,
		Status:  uploadrequest.StatusPending,
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal enum constants: %v", err)
	}
	for _, value := range []string{"digital", "requested_by_customer", "pending"} {
		if !strings.Contains(string(encoded), `"`+value+`"`) {
			t.Fatalf("expected %q in %s", value, encoded)
		}
	}
}
