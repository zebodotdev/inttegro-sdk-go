package inttegro

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestEnumConstantsSerializeAsWireValues(t *testing.T) {
	payload := struct {
		Product ProductType         `json:"product"`
		Refund  RefundReason        `json:"refund"`
		Status  UploadRequestStatus `json:"status"`
	}{
		Product: ProductTypeDigital,
		Refund:  RefundReasonRequestedByCustomer,
		Status:  UploadRequestStatusPending,
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
