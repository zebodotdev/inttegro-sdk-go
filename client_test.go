package inttegro

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

var uuidV7Pattern = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-7[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

// mockClient rewrites the base URL to the test server. If we cannot bind in
// this environment, skip instead of failing.
func newTestClient(t *testing.T, handler http.Handler) (*Client, func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Skipf("skipping: unable to start test server (%v)", r)
		}
	}()
	srv := httptest.NewServer(handler)
	client := NewClient("sk_test_123", WithBaseURL(srv.URL))
	return client, srv.Close
}

func TestDoSuccess(t *testing.T) {
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer sk_test_123" {
			t.Fatalf("expected auth header set, got %q", got)
		}
		if got := r.Header.Get("User-Agent"); got != "inttegro-sdk-go/"+Version {
			t.Fatalf("expected SDK user agent, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"ok":true}`)
	}))
	if client == nil {
		return
	}
	defer close()

	var resp struct {
		OK bool `json:"ok"`
	}
	if err := client.do(context.Background(), "GET", "/ping", nil, &resp); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.OK {
		t.Fatalf("expected ok=true")
	}
}

func TestDoAPIError(t *testing.T) {
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"type":"invalid_request_parameter","code":"invalid_payment_method","url":"https://studio.inttegro.com/e/invalid_payment_method","message":"invalid payment method","detail":"payment method not usable for this currency","fix_code":"change_request_parameters","cause":"validation_failure"}`)
	}))
	if client == nil {
		return
	}
	defer close()

	err := client.do(context.Background(), "GET", "/ping", nil, nil)
	if err == nil {
		t.Fatalf("expected error")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T", err)
	}
	if apiErr.Code != "invalid_payment_method" {
		t.Fatalf("unexpected code: %s", apiErr.Code)
	}
	if !strings.Contains(apiErr.Message, "invalid payment method") {
		t.Fatalf("unexpected message: %s", apiErr.Message)
	}
	if apiErr.URL == "" {
		t.Fatalf("expected url to be set")
	}
}

func TestTelemetryTracesLogicalOperationAndRedactsRequestData(t *testing.T) {
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	t.Cleanup(func() { _ = provider.Shutdown(context.Background()) })

	var traceparent string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceparent = r.Header.Get("traceparent")
		w.Header().Set("x-request-id", "req_telemetry")
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"order":{"id":"or_private_123"}}`)
	}))
	defer server.Close()

	client := NewClient(
		"sk_test_do_not_trace",
		WithBaseURL(server.URL),
		WithTracerProvider(provider),
		WithTextMapPropagator(propagation.TraceContext{}),
	)
	var response map[string]any
	if err := client.do(context.Background(), "POST", "/orders/lookup", map[string]string{"order_id": "or_private_123"}, &response); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if traceparent == "" {
		t.Fatal("expected traceparent to be propagated")
	}

	spans := recorder.Ended()
	if len(spans) != 1 {
		t.Fatalf("ended spans = %d, want 1", len(spans))
	}
	span := spans[0]
	if got := span.Name(); got != "inttegro.orders.lookup" {
		t.Fatalf("span name = %q, want inttegro.orders.lookup", got)
	}
	assertSpanAttribute(t, span, "inttegro.operation.name", "orders.lookup")
	assertSpanAttribute(t, span, "http.response.status_code", "200")
	assertSpanAttribute(t, span, "inttegro.request.id", "req_telemetry")
	assertSpanEvent(t, span, "inttegro.request.prepared")
	assertSpanEvent(t, span, "inttegro.http.attempt.started")
	assertSpanEvent(t, span, "inttegro.response.received")
	assertSpanEvent(t, span, "inttegro.response.decoded")

	telemetryText := spanTelemetryText(span)
	for _, secret := range []string{"sk_test_do_not_trace", "or_private_123"} {
		if strings.Contains(telemetryText, secret) {
			t.Fatalf("telemetry contained private value %q: %s", secret, telemetryText)
		}
	}
}

func TestTelemetryDoesNotNameUnknownRoutesFromResourceIDs(t *testing.T) {
	operation, route, serverAddress := telemetryRequestDetails(
		"https://api.inttegro.com",
		"/orders/or_private_123",
		"",
	)
	if operation != "http.request" || route != "" || serverAddress != "api.inttegro.com" {
		t.Fatalf("request details = (%q, %q, %q), want bounded fallback", operation, route, serverAddress)
	}
}

func TestTelemetryRecordsSafeHTTPFailure(t *testing.T) {
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	t.Cleanup(func() { _ = provider.Shutdown(context.Background()) })

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"message":"sk_live_private must never be traced"}`)
	}))
	defer server.Close()

	client := NewClient("sk_live_private", WithBaseURL(server.URL), WithTracerProvider(provider))
	if err := client.do(context.Background(), "POST", "/orders/lookup", map[string]string{"order_id": "or_private"}, nil); err == nil {
		t.Fatal("expected API error")
	}

	spans := recorder.Ended()
	if len(spans) != 1 {
		t.Fatalf("ended spans = %d, want 1", len(spans))
	}
	span := spans[0]
	assertSpanAttribute(t, span, "error.type", "http_401")
	assertSpanEvent(t, span, "inttegro.request.failed")
	telemetryText := spanTelemetryText(span)
	if strings.Contains(telemetryText, "sk_live_private") || strings.Contains(telemetryText, "or_private") {
		t.Fatalf("failure telemetry contained private request data: %s", telemetryText)
	}
}

func assertSpanAttribute(t *testing.T, span sdktrace.ReadOnlySpan, key, want string) {
	t.Helper()
	for _, attr := range span.Attributes() {
		if string(attr.Key) == key && attr.Value.Emit() == want {
			return
		}
	}
	t.Fatalf("span attribute %q = %q not found", key, want)
}

func assertSpanEvent(t *testing.T, span sdktrace.ReadOnlySpan, want string) {
	t.Helper()
	for _, event := range span.Events() {
		if event.Name == want {
			return
		}
	}
	t.Fatalf("span event %q not found", want)
}

func spanTelemetryText(span sdktrace.ReadOnlySpan) string {
	var builder strings.Builder
	for _, attr := range span.Attributes() {
		builder.WriteString(string(attr.Key))
		builder.WriteString(attr.Value.Emit())
	}
	for _, event := range span.Events() {
		builder.WriteString(event.Name)
		for _, attr := range event.Attributes {
			builder.WriteString(string(attr.Key))
			builder.WriteString(attr.Value.Emit())
		}
	}
	return builder.String()
}

func TestDoGeneratesRequestMetaIdempotencyKeyForMutations(t *testing.T) {
	var body map[string]any
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"ok":true}`)
	}))
	if client == nil {
		return
	}
	defer close()

	err := client.do(context.Background(), "POST", "/orders/new", map[string]any{
		"number":          "ORDER-1",
		"idempotency_key": "legacy",
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := body["idempotency_key"]; ok {
		t.Fatalf("top-level idempotency_key should not be sent: %#v", body)
	}
	requestMeta, ok := body["request_meta"].(map[string]any)
	if !ok {
		t.Fatalf("request_meta missing: %#v", body)
	}
	key, _ := requestMeta["idempotency_key"].(string)
	if !uuidV7Pattern.MatchString(key) {
		t.Fatalf("idempotency key %q is not UUIDv7", key)
	}
}

func TestDoSkipsIdempotencyForReadStylePosts(t *testing.T) {
	var body map[string]any
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"ok":true}`)
	}))
	if client == nil {
		return
	}
	defer close()

	err := client.do(context.Background(), "POST", "/orders/lookup", map[string]any{
		"order_id":        "or_123",
		"idempotency_key": "legacy",
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := body["idempotency_key"]; ok {
		t.Fatalf("top-level idempotency_key should not be sent: %#v", body)
	}
	if _, ok := body["request_meta"]; ok {
		t.Fatalf("request_meta should not be sent for lookup: %#v", body)
	}
}

func TestMessageTemplatesCreateUsesRequestMetaIdempotencyByDefault(t *testing.T) {
	var body map[string]any
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Path; got != "/message_templates/create" {
			t.Fatalf("path = %q, want /message_templates/create", got)
		}
		if got := r.Header.Get("Idempotency-Key"); got != "" {
			t.Fatalf("Idempotency-Key header should not be generated for JSON mutations, got %q", got)
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"message_template":{"id":"mt_123"}}`)
	}))
	if client == nil {
		return
	}
	defer close()

	_, err := client.MessageTemplates.Create(context.Background(), MessageTemplateCreateParams{
		Name:    "welcome_sms",
		Channel: "sms",
		Purpose: "marketing",
		SMS:     &MessageTemplateSMSContent{MessageTemplate: "Welcome {{name}}"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	requestMeta, ok := body["request_meta"].(map[string]any)
	if !ok {
		t.Fatalf("request_meta missing: %#v", body)
	}
	key, _ := requestMeta["idempotency_key"].(string)
	if !uuidV7Pattern.MatchString(key) {
		t.Fatalf("idempotency key %q is not UUIDv7", key)
	}
}
