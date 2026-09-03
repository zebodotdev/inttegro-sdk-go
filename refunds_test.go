package inttegro

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/zebodotdev/inttegro-sdk-go/v4/money"
)

func TestRefundsServiceUsesCanonicalContracts(t *testing.T) {
	type capturedRequest struct {
		path string
		body map[string]any
	}
	var captured []capturedRequest
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		captured = append(captured, capturedRequest{path: r.URL.Path, body: body})
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/refunds/page" {
			_, _ = io.WriteString(w, refundPageResponseJSON)
			return
		}
		_, _ = io.WriteString(w, refundResponseJSON)
	}))
	if client == nil {
		return
	}
	defer close()
	if client.Refunds == nil {
		t.Fatal("NewClient() did not initialize Refunds")
	}

	ctx := context.Background()
	createRequest := fullCreateRefundRequest()
	created, err := client.Refunds.Create(ctx, createRequest)
	if err != nil {
		t.Fatalf("Refunds.Create() error = %v", err)
	}
	canceled, err := client.Refunds.Cancel(ctx, CancelRefundRequest{
		RefundID: "rf_123", RequestMeta: &RequestMeta{IdempotencyKey: "cancel-refund-123"},
	})
	if err != nil {
		t.Fatalf("Refunds.Cancel() error = %v", err)
	}
	lookedUp, err := client.Refunds.Lookup(ctx, LookupRefundRequest{RefundID: "rf_123"})
	if err != nil {
		t.Fatalf("Refunds.Lookup() error = %v", err)
	}
	page, err := client.Refunds.Page(ctx, PageRefundsRequest{PageNumber: 2, PageSize: 25})
	if err != nil {
		t.Fatalf("Refunds.Page() error = %v", err)
	}

	wantPaths := []string{"/refunds/create", "/refunds/cancel", "/refunds/lookup", "/refunds/page"}
	if len(captured) != len(wantPaths) {
		t.Fatalf("captured requests = %#v", captured)
	}
	for i, want := range wantPaths {
		if captured[i].path != want {
			t.Fatalf("request %d path = %q, want %q", i, captured[i].path, want)
		}
	}
	assertJSONMapEqual(t, captured[0].body, map[string]any{
		"line_items": []any{map[string]any{
			"order_line_item_id": "oli_123",
			"refund_amount":      map[string]any{"currency": "ghs", "value": float64(2500)},
			"reason":             "item_damaged",
			"reason_details":     "damaged in transit",
		}},
		"order_id":       "or_123",
		"reason":         "requested_by_customer",
		"reason_details": "customer returned the item",
		"reference":      "RETURN-123",
		"custom_data":    map[string]any{"warehouse": "accra"},
		"request_meta":   map[string]any{"idempotency_key": "create-refund-123"},
	})
	assertJSONMapEqual(t, captured[1].body, map[string]any{
		"refund_id":    "rf_123",
		"request_meta": map[string]any{"idempotency_key": "cancel-refund-123"},
	})
	assertJSONMapEqual(t, captured[2].body, map[string]any{"refund_id": "rf_123"})
	assertJSONMapEqual(t, captured[3].body, map[string]any{
		"page_number": float64(2), "page_size": float64(25),
	})

	assertDecodedRefund(t, *created)
	assertDecodedRefund(t, *canceled)
	assertDecodedRefund(t, *lookedUp)
	if page.Number != 2 || page.Size != 1 || len(page.Refunds) != 1 {
		t.Fatalf("decoded refund page = %#v", page)
	}
	assertDecodedRefund(t, page.Refunds[0])
}

func TestRefundCreateRequestOmitsOptionalFields(t *testing.T) {
	raw, err := json.Marshal(CreateRefundRequest{
		OrderID: "or_123",
		Reason:  RefundReasonItemReturned,
		LineItems: []CreateRefundLineItem{{
			OrderLineItemID: "oli_123",
			RefundAmount:    money.AmountParams{Currency: money.GHS, Value: 100},
		}},
	})
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	var body map[string]any
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	for _, key := range []string{"custom_data", "reason_details", "reference", "request_meta"} {
		if _, exists := body[key]; exists {
			t.Fatalf("optional request property %q was encoded: %#v", key, body)
		}
	}
	line := body["line_items"].([]any)[0].(map[string]any)
	for _, key := range []string{"reason", "reason_details"} {
		if _, exists := line[key]; exists {
			t.Fatalf("optional line property %q was encoded: %#v", key, line)
		}
	}
}

func TestRefundOmitsOptionalResponseFields(t *testing.T) {
	raw, err := json.Marshal(Refund{
		ID:        "rf_123",
		OrderID:   "or_123",
		Status:    RefundStatusPending,
		Total:     money.Amount{Currency: money.GHS, Value: 100},
		LineItems: []RefundLineItem{},
		Reason:    RefundReasonItemReturned,
		CreatedAt: "2026-09-02T10:00:00Z",
	})
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	var body map[string]any
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	for _, key := range []string{
		"canceled_at", "custom_data", "failed_at", "processing_at",
		"reason_details", "reference", "succeeded_at",
	} {
		if _, exists := body[key]; exists {
			t.Fatalf("optional refund property %q was encoded: %#v", key, body)
		}
	}
	for _, key := range []string{"created_at", "id", "line_items", "order_id", "reason", "status", "total"} {
		if _, exists := body[key]; !exists {
			t.Fatalf("required refund property %q was omitted: %#v", key, body)
		}
	}
}

func TestOrdersRefundMatchesCanonicalCreateContract(t *testing.T) {
	var paths []string
	var bodies []map[string]any
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		bodies = append(bodies, body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, refundResponseJSON)
	}))
	if client == nil {
		return
	}
	defer close()

	request := fullCreateRefundRequest()
	canonical, err := client.Refunds.Create(context.Background(), request)
	if err != nil {
		t.Fatalf("Refunds.Create() error = %v", err)
	}
	compatibility, err := client.Orders.Refund(context.Background(), request)
	if err != nil {
		t.Fatalf("Orders.Refund() error = %v", err)
	}
	if !reflect.DeepEqual(paths, []string{"/refunds/create", "/orders/refund"}) {
		t.Fatalf("paths = %#v", paths)
	}
	if len(bodies) != 2 || !reflect.DeepEqual(bodies[0], bodies[1]) {
		t.Fatalf("canonical and compatibility bodies differ: %#v", bodies)
	}
	if !reflect.DeepEqual(canonical, compatibility) {
		t.Fatalf("canonical and compatibility responses differ: %#v %#v", canonical, compatibility)
	}
}

func TestRefundMutationsGenerateRequestMeta(t *testing.T) {
	var bodies []map[string]any
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		bodies = append(bodies, body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, refundResponseJSON)
	}))
	if client == nil {
		return
	}
	defer close()

	request := fullCreateRefundRequest()
	request.RequestMeta = nil
	if _, err := client.Refunds.Create(context.Background(), request); err != nil {
		t.Fatalf("Refunds.Create() error = %v", err)
	}
	if _, err := client.Refunds.Cancel(context.Background(), CancelRefundRequest{RefundID: "rf_123"}); err != nil {
		t.Fatalf("Refunds.Cancel() error = %v", err)
	}

	if len(bodies) != 2 {
		t.Fatalf("captured bodies = %#v", bodies)
	}
	for i, body := range bodies {
		requestMeta, ok := body["request_meta"].(map[string]any)
		if !ok {
			t.Fatalf("request %d request_meta missing: %#v", i, body)
		}
		key, _ := requestMeta["idempotency_key"].(string)
		if !uuidV7Pattern.MatchString(key) {
			t.Fatalf("request %d idempotency key = %q, want UUIDv7", i, key)
		}
	}
}

func fullCreateRefundRequest() CreateRefundRequest {
	lineReason := RefundReasonItemDamaged
	return CreateRefundRequest{
		OrderID: "or_123",
		Reason:  RefundReasonRequestedByCustomer,
		LineItems: []CreateRefundLineItem{{
			OrderLineItemID: "oli_123",
			RefundAmount:    money.AmountParams{Currency: money.GHS, Value: 2500},
			Reason:          &lineReason,
			ReasonDetails:   "damaged in transit",
		}},
		ReasonDetails: "customer returned the item",
		Reference:     "RETURN-123",
		CustomData:    map[string]string{"warehouse": "accra"},
		RequestMeta:   &RequestMeta{IdempotencyKey: "create-refund-123"},
	}
}

func assertDecodedRefund(t *testing.T, refund Refund) {
	t.Helper()
	if refund.ID != "rf_123" || refund.OrderID != "or_123" ||
		refund.Status != RefundStatusProcessing || refund.Total.Currency != "ghs" ||
		refund.Total.Value != 2500 || refund.Reason != RefundReasonRequestedByCustomer ||
		refund.Reference != "RETURN-123" || refund.CustomData["warehouse"] != "accra" {
		t.Fatalf("decoded refund = %#v", refund)
	}
	if refund.ProcessingAt == nil || *refund.ProcessingAt != "2026-09-02T10:01:00Z" ||
		refund.SucceededAt != nil || refund.FailedAt != nil || refund.CanceledAt != nil {
		t.Fatalf("decoded lifecycle timestamps = %#v", refund)
	}
	if len(refund.LineItems) != 1 || refund.LineItems[0].ID != "rli_123" ||
		refund.LineItems[0].OrderLineItemID != "oli_123" ||
		refund.LineItems[0].OriginalAmountPaid.Value != 5000 ||
		refund.LineItems[0].RefundAmount.Value != 2500 ||
		refund.LineItems[0].Reason == nil ||
		*refund.LineItems[0].Reason != RefundReasonItemDamaged {
		t.Fatalf("decoded refund lines = %#v", refund.LineItems)
	}
}

func assertJSONMapEqual(t *testing.T, got, want map[string]any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request body = %#v, want %#v", got, want)
	}
}

const refundObjectJSON = `{
    "id": "rf_123",
    "order_id": "or_123",
    "status": "processing",
    "total": {"currency": "ghs", "value": 2500},
    "line_items": [{
      "id": "rli_123",
      "order_line_item_id": "oli_123",
      "original_amount_paid": {"currency": "ghs", "value": 5000},
      "refund_amount": {"currency": "ghs", "value": 2500},
      "reason": "item_damaged",
      "reason_details": "damaged in transit"
    }],
    "reason": "requested_by_customer",
    "reason_details": "customer returned the item",
    "reference": "RETURN-123",
    "custom_data": {"warehouse": "accra"},
    "created_at": "2026-09-02T10:00:00Z",
    "processing_at": "2026-09-02T10:01:00Z"
}`

const refundResponseJSON = `{"refund":` + refundObjectJSON + `}`

const refundPageResponseJSON = `{
  "page": {
    "number": 2,
    "size": 1,
    "refunds": [` + refundObjectJSON + `]
  }
}`
