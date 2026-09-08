// Package inttegro provides a Go client for the Inttegro API.
//
// The Inttegro API enables businesses to accept payments, manage payouts,
// tokenize payment methods, and send notifications across multiple payment
// rails including mobile money, bank accounts, and cards.
//
// # Resource packages
//
// Import the singular package for each API resource. These packages keep names
// short and domain-scoped: product.TypeDigital, purchaseintent.StatusActive,
// refund.CreateParams, and order.Order. Client services remain plural
// collections, such as Client.Products and Client.PurchaseIntents.
//
// The root inttegro package owns the client, transport options, API errors,
// and telemetry. Resource models, lifecycle values, request parameters, and
// services are defined by their resource packages.
//
// # Getting Started
//
// Create a client with your API key:
//
//	client := inttegro.NewClient("sk_live_...")
//
// For testing, use your test mode API key:
//
//	client := inttegro.NewClient("sk_test_...")
//
// # Authentication
//
// All requests require an API key passed as a Bearer token in the Authorization header.
// The client handles this automatically. Get your API keys from the Inttegro dashboard.
//
// # Error Handling
//
// API errors return an *APIError with structured error information:
//
//	order, err := client.Orders.Create(ctx, params)
//	if err != nil {
//	    if apiErr, ok := err.(*inttegro.APIError); ok {
//	        fmt.Printf("Error code: %s\n", apiErr.Code)
//	        fmt.Printf("Message: %s\n", apiErr.Message)
//	        fmt.Printf("Type: %s\n", apiErr.Type)
//	    }
//	    return err
//	}
//
// # Idempotency
//
// For write operations (creating orders, tokenizing payment methods, sending chimes),
// pass request.Meta.IdempotencyKey to safely retry requests without duplicating resources.
// The same idempotency key can be reused if the original request failed.
//
//	params := order.CreateParams{
//	    RequestMeta: &request.Meta{IdempotencyKey: "order_20231215_customer_123"},
//	    // ... other fields
//	}
//
// Learn more: https://studio.inttegro.com/idempotency
//
// # Context Support
//
// All service methods accept a context.Context for cancellation and timeouts:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	order, err := client.Orders.Create(ctx, params)
//
// # Custom HTTP Client
//
// Customize the underlying HTTP client for proxy support, custom timeouts, or retry logic:
//
//	httpClient := &http.Client{
//	    Timeout: 60 * time.Second,
//	    Transport: &http.Transport{
//	        Proxy: http.ProxyFromEnvironment,
//	    },
//	}
//	client := inttegro.NewClient("sk_live_...", inttegro.WithHTTPClient(httpClient))
package inttegro

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/app"
	"github.com/zebodotdev/inttegro-sdk-go/v6/balance"
	"github.com/zebodotdev/inttegro-sdk-go/v6/balancetransaction"
	"github.com/zebodotdev/inttegro-sdk-go/v6/broadcast"
	"github.com/zebodotdev/inttegro-sdk-go/v6/chime"
	"github.com/zebodotdev/inttegro-sdk-go/v6/customer"
	resourcefile "github.com/zebodotdev/inttegro-sdk-go/v6/file"
	"github.com/zebodotdev/inttegro-sdk-go/v6/filelink"
	"github.com/zebodotdev/inttegro-sdk-go/v6/filereference"
	"github.com/zebodotdev/inttegro-sdk-go/v6/financialaccount"
	"github.com/zebodotdev/inttegro-sdk-go/v6/messagetemplate"
	"github.com/zebodotdev/inttegro-sdk-go/v6/order"
	"github.com/zebodotdev/inttegro-sdk-go/v6/otp"
	"github.com/zebodotdev/inttegro-sdk-go/v6/paymentmethod"
	"github.com/zebodotdev/inttegro-sdk-go/v6/payout"
	"github.com/zebodotdev/inttegro-sdk-go/v6/price"
	"github.com/zebodotdev/inttegro-sdk-go/v6/product"
	"github.com/zebodotdev/inttegro-sdk-go/v6/purchaseintent"
	"github.com/zebodotdev/inttegro-sdk-go/v6/refund"
	"github.com/zebodotdev/inttegro-sdk-go/v6/schedule"
	"github.com/zebodotdev/inttegro-sdk-go/v6/secretkey"
	"github.com/zebodotdev/inttegro-sdk-go/v6/spec"
	"github.com/zebodotdev/inttegro-sdk-go/v6/uploadrequest"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	// DefaultBaseURL is the production Inttegro API base URL.
	DefaultBaseURL = "https://api.inttegro.com"
	defaultTimeout = 30 * time.Second
)

// Client is the main entry point for interacting with the Inttegro API.
//
// It provides access to all Inttegro API resources through service properties.
// Create a client using NewClient with your API key:
//
//	client := inttegro.NewClient("sk_live_...")
//
// The client automatically handles authentication, request serialization,
// and response deserialization. All service methods accept a context.Context
// for cancellation and timeouts.
type Client struct {
	// APIKey is your Inttegro API key (required).
	// Get your keys from the Inttegro dashboard.
	APIKey string

	// BaseURL is the API base URL. Defaults to DefaultBaseURL.
	// Override using WithBaseURL option for testing or custom environments.
	BaseURL string

	// HTTPClient is the underlying HTTP client used for requests.
	// Defaults to a client with 30 second timeout.
	// Customize using WithHTTPClient option for proxy, retry, or timeout control.
	HTTPClient *http.Client

	tracerProvider       trace.TracerProvider
	propagator           propagation.TextMapPropagator
	tracer               trace.Tracer
	telemetryEnabled     bool
	errorReporter        ErrorReporter
	errorReportingPolicy ErrorReportingPolicy

	// Orders provides access to order creation, payment, and lifecycle management.
	// See order.Service for available operations.
	Orders *order.Service

	// Refunds provides refund creation, lookup, cancellation, and paging.
	Refunds *refund.Service

	// Chimes provides access to notification sending and scheduling.
	// Send SMS or email notifications to customers.
	Chimes *chime.Service

	// Schedules provides access to scheduled chime lookups and cancellations.
	Schedules *schedule.Service

	// Broadcasts provides access to broadcast lookups and cancellations.
	Broadcasts *broadcast.Service

	// MessageTemplates provides access to reusable SMS and email templates.
	MessageTemplates *messagetemplate.Service

	// Otp provides access to one-time password initialization and verification.
	// Used for custom authentication flows.
	Otp *otp.Service

	// PaymentMethods provides payment method tokenization, verification, and management.
	// Save payment methods for repeat customers and verify ownership.
	PaymentMethods *paymentmethod.Service

	// Payouts provides payout configuration, scheduling, and listing.
	// Configure automatic or manual payout schedules and destination accounts.
	Payouts *payout.Service

	// Balances provides access to balance snapshots across currencies.
	Balances *balance.Service

	// BalanceTransactions provides access to balance transaction history.
	// View available and pending funds from completed payments.
	BalanceTransactions *balancetransaction.Service

	// FinancialAccounts provides financial account connection and management.
	// Connect mobile money, bank, or Dosh accounts for receiving payouts.
	FinancialAccounts *financialaccount.Service

	// Files provides file upload, lookup, download, paging, and deletion.
	Files *resourcefile.Service

	// FileLinks provides revocable public links for linkable files.
	FileLinks *filelink.Service

	// UploadRequests provides delegated public file upload requests.
	UploadRequests *uploadrequest.Service

	// Customers provides access to customer records.
	Customers *customer.Service

	// Products provides access to catalog products.
	Products *product.Service

	// Prices provides access to catalog prices.
	Prices *price.Service

	// Spec provides access to country specifications and supported features.
	// Query supported currencies, payment methods, and payout schedules by country.
	Spec *spec.Service

	// Apps provides access to application creation, lookup, and updates.
	Apps *app.Service

	// Keys provides access to secret key management.
	Keys *secretkey.Service

	// PurchaseIntents provides access to Buy link purchase intent management.
	PurchaseIntents *purchaseintent.Service

	// FileReferences provides access to file reference reconciliation.
	FileReferences *filereference.Service
}

// ClientOption allows customizing the client during construction.
// Options are passed to NewClient to override defaults.
type ClientOption func(*Client)

// WithBaseURL overrides the default API base URL.
//
// Use this for testing against a local or staging environment:
//
//	client := inttegro.NewClient(apiKey, inttegro.WithBaseURL("http://localhost:8080"))
//
// The URL is automatically trimmed of trailing slashes.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		if baseURL != "" {
			c.BaseURL = strings.TrimRight(baseURL, "/")
		}
	}
}

// WithHTTPClient uses a custom http.Client for requests.
//
// Use this to customize timeouts, proxy settings, or retry behavior:
//
//	httpClient := &http.Client{
//	    Timeout: 60 * time.Second,
//	    Transport: &http.Transport{
//	        Proxy: http.ProxyFromEnvironment,
//	        MaxIdleConns: 100,
//	    },
//	}
//	client := inttegro.NewClient(apiKey, inttegro.WithHTTPClient(httpClient))
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		if httpClient != nil {
			c.HTTPClient = httpClient
		}
	}
}

// WithTracerProvider uses provider for Inttegro client spans. When omitted,
// the OpenTelemetry global tracer provider is used.
func WithTracerProvider(provider trace.TracerProvider) ClientOption {
	return func(c *Client) {
		if provider != nil {
			c.tracerProvider = provider
		}
	}
}

// WithTextMapPropagator uses propagator to inject distributed-tracing context.
// When omitted, the OpenTelemetry global text-map propagator is used.
func WithTextMapPropagator(propagator propagation.TextMapPropagator) ClientOption {
	return func(c *Client) {
		if propagator != nil {
			c.propagator = propagator
		}
	}
}

// WithTelemetryEnabled controls Inttegro's OpenTelemetry instrumentation.
// Instrumentation is enabled by default but remains a no-op unless the
// application configures an OpenTelemetry tracer provider.
func WithTelemetryEnabled(enabled bool) ClientOption {
	return func(c *Client) {
		c.telemetryEnabled = enabled
	}
}

// WithErrorReporter configures an application-owned destination for privacy-safe
// final-failure reports. Without this option, the SDK performs no report preparation.
func WithErrorReporter(reporter ErrorReporter) ClientOption {
	return func(c *Client) {
		c.errorReporter = reporter
	}
}

// WithErrorReportingPolicy controls whether expected API errors are reported.
// The default is ErrorReportingUnexpected.
func WithErrorReportingPolicy(policy ErrorReportingPolicy) ClientOption {
	return func(c *Client) {
		if policy == ErrorReportingAll || policy == ErrorReportingUnexpected {
			c.errorReportingPolicy = policy
		}
	}
}

// NewClient constructs a Inttegro API client.
//
// The apiKey parameter is required and should be your Inttegro API key
// from the dashboard (starts with sk_live_ or sk_test_).
//
// Options can be passed to customize the client behavior:
//
//	// Basic client with defaults
//	client := inttegro.NewClient("sk_live_...")
//
//	// Client with custom timeout and base URL
//	client := inttegro.NewClient(
//	    "sk_test_...",
//	    inttegro.WithBaseURL("https://api.staging.inttegro.com"),
//	    inttegro.WithHTTPClient(&http.Client{Timeout: 60*time.Second}),
//	)
//
// The client is safe for concurrent use by multiple goroutines.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		APIKey:               apiKey,
		BaseURL:              strings.TrimRight(DefaultBaseURL, "/"),
		HTTPClient:           &http.Client{Timeout: defaultTimeout},
		tracerProvider:       otel.GetTracerProvider(),
		propagator:           otel.GetTextMapPropagator(),
		telemetryEnabled:     true,
		errorReportingPolicy: ErrorReportingUnexpected,
	}
	for _, opt := range opts {
		opt(c)
	}
	c.tracer = c.tracerProvider.Tracer("inttegro", trace.WithInstrumentationVersion(Version))

	c.Orders = order.NewService(c)
	c.Refunds = refund.NewService(c)
	c.Chimes = chime.NewService(c)
	c.Schedules = schedule.NewService(c)
	c.Broadcasts = broadcast.NewService(c)
	c.MessageTemplates = messagetemplate.NewService(c)
	c.Otp = otp.NewService(c)
	c.PaymentMethods = paymentmethod.NewService(c)
	c.Payouts = payout.NewService(c)
	c.Balances = balance.NewService(c)
	c.BalanceTransactions = balancetransaction.NewService(c)
	c.FinancialAccounts = financialaccount.NewService(c)
	c.Files = resourcefile.NewService(c)
	c.FileLinks = filelink.NewService(c)
	c.UploadRequests = uploadrequest.NewService(c)
	c.Customers = customer.NewService(c)
	c.Products = product.NewService(c)
	c.Prices = price.NewService(c)
	c.Spec = spec.NewService(c)
	c.Apps = app.NewService(c)
	c.Keys = secretkey.NewService(c)
	c.PurchaseIntents = purchaseintent.NewService(c)
	c.FileReferences = filereference.NewService(c)

	return c
}

// do executes an HTTP request to the Inttegro API.
//
// This is an internal method used by all service methods. It handles:
// - Request serialization (JSON encoding)
// - Authentication (Bearer token)
// - Response deserialization
// - Error parsing and structured error responses
// - Context cancellation and timeouts
//
// The method parameter specifies the HTTP method (GET, POST, etc).
// The path parameter is relative to BaseURL (e.g., "/orders/new").
// The body parameter is JSON-encoded if not nil.
// The out parameter receives the decoded response if not nil.
//
// Returns an *APIError for HTTP errors (status >= 400).
func (c *Client) do(ctx context.Context, method, path string, body any, out any) error {
	if c.APIKey == "" {
		return errors.New("api key is required")
	}

	ctx, telemetry := c.startRequestTelemetry(ctx, method, path, "")
	defer telemetry.end()

	var reqBody io.Reader
	if body != nil {
		raw, err := c.jsonRequestBody(method, path, body, "")
		if err != nil {
			wrapped := fmt.Errorf("encode request body: %w", err)
			telemetry.failAndReport(ctx, wrapped, "encode_error")
			return wrapped
		}
		reqBody = bytes.NewReader(raw)
	}

	url := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		wrapped := fmt.Errorf("create request: %w", err)
		telemetry.failAndReport(ctx, wrapped, "request_error")
		return wrapped
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "inttegro-sdk-go/"+Version)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	telemetry.inject(ctx, req.Header)
	telemetry.attempt()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		wrapped := fmt.Errorf("execute request: %w", err)
		telemetry.failAndReport(ctx, wrapped, "transport_error")
		return wrapped
	}
	defer resp.Body.Close()
	telemetry.response(resp)

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		wrapped := fmt.Errorf("read response: %w", err)
		telemetry.failAndReport(ctx, wrapped, "read_error")
		return wrapped
	}

	if resp.StatusCode >= 400 {
		apiErr := &APIError{StatusCode: resp.StatusCode, Body: respBytes}
		var parsed APIError
		if err := json.Unmarshal(respBytes, &parsed); err == nil && hasAPIErrorDetails(&parsed) {
			copyAPIError(apiErr, &parsed)
		} else {
			var env errorEnvelope
			if err := json.Unmarshal(respBytes, &env); err == nil && env.Error != nil {
				copyAPIError(apiErr, env.Error)
			} else if len(respBytes) > 0 {
				apiErr.Message = string(respBytes)
			}
		}
		if apiErr.Message == "" && len(respBytes) > 0 {
			apiErr.Message = string(respBytes)
		}
		apiErr.RequestID = resp.Header.Get("x-request-id")
		telemetry.failAndReport(ctx, apiErr, fmt.Sprintf("http_%d", resp.StatusCode))
		return apiErr
	}

	if out != nil && len(respBytes) > 0 {
		if err := json.Unmarshal(respBytes, out); err != nil {
			wrapped := fmt.Errorf("decode response: %w", err)
			telemetry.failAndReport(ctx, wrapped, "decode_error")
			return wrapped
		}
		telemetry.decoded()
	}

	return nil
}

func (c *Client) jsonRequestBody(method, path string, body any, explicitIdempotencyKey string) ([]byte, error) {
	raw, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	if !strings.EqualFold(method, "POST") || !isIdempotentMutationPath(path) {
		return withoutTopLevelIdempotencyKey(raw)
	}
	return withRequestMetaIdempotency(raw, explicitIdempotencyKey == "")
}

func withRequestMetaIdempotency(raw []byte, generate bool) ([]byte, error) {
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil || payload == nil {
		return raw, nil
	}
	delete(payload, "idempotency_key")
	requestMeta, _ := payload["request_meta"].(map[string]any)
	if requestMeta == nil {
		requestMeta = map[string]any{}
	}
	if existing, ok := requestMeta["idempotency_key"].(string); ok && strings.TrimSpace(existing) != "" {
		payload["request_meta"] = requestMeta
		return json.Marshal(payload)
	}
	if generate {
		requestMeta["idempotency_key"] = generateIdempotencyKey()
		payload["request_meta"] = requestMeta
	}
	return json.Marshal(payload)
}

func withoutTopLevelIdempotencyKey(raw []byte) ([]byte, error) {
	return withRequestMetaIdempotency(raw, false)
}

func isIdempotentMutationPath(pathOrURL string) bool {
	path := pathOrURL
	if strings.HasPrefix(pathOrURL, "http://") || strings.HasPrefix(pathOrURL, "https://") {
		parsed, err := url.Parse(pathOrURL)
		if err != nil {
			return false
		}
		path = parsed.Path
	}
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 || parts[len(parts)-1] == "" {
		return false
	}
	switch parts[len(parts)-1] {
	case "lookup", "page", "settings", "countries", "contents", "balances", "render_preview", "usage":
		return false
	default:
		return true
	}
}

func generateIdempotencyKey() string {
	var b [16]byte
	timestamp := uint64(time.Now().UnixMilli()) & ((uint64(1) << 48) - 1)
	b[0] = byte(timestamp >> 40)
	b[1] = byte(timestamp >> 32)
	b[2] = byte(timestamp >> 24)
	b[3] = byte(timestamp >> 16)
	b[4] = byte(timestamp >> 8)
	b[5] = byte(timestamp)
	if _, err := rand.Read(b[6:]); err != nil {
		fallback := uint64(time.Now().UnixNano())
		for i := 6; i < len(b); i++ {
			fallback = fallback*6364136223846793005 + 1
			b[i] = byte(fallback >> 56)
		}
	}
	b[6] = (b[6] & 0x0f) | 0x70
	b[8] = (b[8] & 0x3f) | 0x80
	hex := fmt.Sprintf("%x", b)
	return hex[0:8] + "-" + hex[8:12] + "-" + hex[12:16] + "-" + hex[16:20] + "-" + hex[20:32]
}

func hasAPIErrorDetails(err *APIError) bool {
	if err == nil {
		return false
	}
	return err.Code != "" ||
		err.Type != "" ||
		err.URL != "" ||
		err.Message != "" ||
		err.Detail != "" ||
		err.FixCode != "" ||
		err.Cause != ""
}

func copyAPIError(dst *APIError, src *APIError) {
	if dst == nil || src == nil {
		return
	}
	dst.Code = src.Code
	dst.Type = src.Type
	dst.URL = src.URL
	dst.Message = src.Message
	dst.Detail = src.Detail
	dst.FixCode = src.FixCode
	dst.Cause = src.Cause
}
