package commerce

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
		io.WriteString(w, `{"type":"invalid_request_parameter","code":"invalid_payment_method","url":"https://studio.zebo.dev/e/invalid_payment_method","message":"invalid payment method","detail":"payment method not usable for this currency","fix_code":"change_request_parameters","cause":"validation_failure"}`)
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
