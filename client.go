// Package commerce provides a Go client for the Commerce API.
//
// The Commerce API enables businesses to accept payments, manage payouts,
// tokenize payment methods, and send notifications across multiple payment
// rails including mobile money, bank accounts, and cards.
//
// # Getting Started
//
// Create a client with your API key:
//
//	client := commerce.NewClient("sk_live_...")
//
// For testing, use your test mode API key:
//
//	client := commerce.NewClient("sk_test_...")
//
// # Authentication
//
// All requests require an API key passed as a Bearer token in the Authorization header.
// The client handles this automatically. Get your API keys from the Commerce dashboard.
//
// # Error Handling
//
// API errors return an *APIError with structured error information:
//
//	order, err := client.Orders.Create(ctx, params)
//	if err != nil {
//	    if apiErr, ok := err.(*commerce.APIError); ok {
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
// pass RequestMeta.IdempotencyKey to safely retry requests without duplicating resources.
// The same idempotency key can be reused if the original request failed.
//
//	params := commerce.OrderCreateParams{
//	    RequestMeta: &commerce.RequestMeta{IdempotencyKey: "order_20231215_customer_123"},
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
//	client := commerce.NewClient("sk_live_...", commerce.WithHTTPClient(httpClient))
package commerce

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
)

const (
	// DefaultBaseURL is the production Commerce API base URL.
	DefaultBaseURL = "https://api.inttegro.com"
	defaultTimeout = 30 * time.Second
)

// Client is the main entry point for interacting with the Commerce API.
//
// It provides access to all Commerce API resources through service properties.
// Create a client using NewClient with your API key:
//
//	client := commerce.NewClient("sk_live_...")
//
// The client automatically handles authentication, request serialization,
// and response deserialization. All service methods accept a context.Context
// for cancellation and timeouts.
type Client struct {
	// APIKey is your Commerce API key (required).
	// Get your keys from the Commerce dashboard.
	APIKey string

	// BaseURL is the API base URL. Defaults to DefaultBaseURL.
	// Override using WithBaseURL option for testing or custom environments.
	BaseURL string

	// HTTPClient is the underlying HTTP client used for requests.
	// Defaults to a client with 30 second timeout.
	// Customize using WithHTTPClient option for proxy, retry, or timeout control.
	HTTPClient *http.Client

	// Orders provides access to order creation, payment, and lifecycle management.
	// See OrdersService for available operations.
	Orders *OrdersService

	// Chimes provides access to notification sending and scheduling.
	// Send SMS or email notifications to customers.
	Chimes *ChimesService

	// Schedules provides access to scheduled chime lookups and cancellations.
	Schedules *SchedulesService

	// Broadcasts provides access to broadcast lookups and cancellations.
	Broadcasts *BroadcastsService

	// MessageTemplates provides access to reusable SMS and email templates.
	MessageTemplates *MessageTemplatesService

	// Otp provides access to one-time password initialization and verification.
	// Used for custom authentication flows.
	Otp *OtpService

	// PaymentMethods provides payment method tokenization, verification, and management.
	// Save payment methods for repeat customers and verify ownership.
	PaymentMethods *PaymentMethodsService

	// Payouts provides payout configuration, scheduling, and listing.
	// Configure automatic or manual payout schedules and destination accounts.
	Payouts *PayoutsService

	// Balances provides access to balance snapshots across currencies.
	Balances *BalancesService

	// BalanceTransactions provides access to balance transaction history.
	// View available and pending funds from completed payments.
	BalanceTransactions *BalanceTransactionsService

	// FinancialAccounts provides financial account connection and management.
	// Connect mobile money, bank, or Dosh accounts for receiving payouts.
	FinancialAccounts *FinancialAccountsService

	// Files provides file upload, lookup, download, paging, and deletion.
	Files *FilesService

	// FileLinks provides revocable public links for linkable files.
	FileLinks *FileLinksService

	// UploadRequests provides delegated public file upload requests.
	UploadRequests *UploadRequestsService

	// Customers provides access to customer records.
	Customers *CustomersService

	// Products provides access to catalog products.
	Products *ProductsService

	// Prices provides access to catalog prices.
	Prices *PricesService

	// Spec provides access to country specifications and supported features.
	// Query supported currencies, payment methods, and payout schedules by country.
	Spec *SpecService

	// Apps provides access to application creation, lookup, and updates.
	Apps *AppsService
}

// ClientOption allows customizing the client during construction.
// Options are passed to NewClient to override defaults.
type ClientOption func(*Client)

// WithBaseURL overrides the default API base URL.
//
// Use this for testing against a local or staging environment:
//
//	client := commerce.NewClient(apiKey, commerce.WithBaseURL("http://localhost:8080"))
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
//	client := commerce.NewClient(apiKey, commerce.WithHTTPClient(httpClient))
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		if httpClient != nil {
			c.HTTPClient = httpClient
		}
	}
}

// NewClient constructs a Commerce API client.
//
// The apiKey parameter is required and should be your Commerce API key
// from the dashboard (starts with sk_live_ or sk_test_).
//
// Options can be passed to customize the client behavior:
//
//	// Basic client with defaults
//	client := commerce.NewClient("sk_live_...")
//
//	// Client with custom timeout and base URL
//	client := commerce.NewClient(
//	    "sk_test_...",
//	    commerce.WithBaseURL("https://api.staging.inttegro.com"),
//	    commerce.WithHTTPClient(&http.Client{Timeout: 60*time.Second}),
//	)
//
// The client is safe for concurrent use by multiple goroutines.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		APIKey:     apiKey,
		BaseURL:    strings.TrimRight(DefaultBaseURL, "/"),
		HTTPClient: &http.Client{Timeout: defaultTimeout},
	}
	for _, opt := range opts {
		opt(c)
	}

	c.Orders = &OrdersService{client: c}
	c.Chimes = &ChimesService{client: c}
	c.Schedules = &SchedulesService{client: c}
	c.Broadcasts = &BroadcastsService{client: c}
	c.MessageTemplates = &MessageTemplatesService{client: c}
	c.Otp = &OtpService{client: c}
	c.PaymentMethods = &PaymentMethodsService{client: c}
	c.Payouts = &PayoutsService{client: c}
	c.Balances = &BalancesService{client: c}
	c.BalanceTransactions = &BalanceTransactionsService{client: c}
	c.FinancialAccounts = &FinancialAccountsService{client: c}
	c.Files = &FilesService{client: c}
	c.FileLinks = &FileLinksService{client: c}
	c.UploadRequests = &UploadRequestsService{client: c}
	c.Customers = &CustomersService{client: c}
	c.Products = &ProductsService{client: c}
	c.Prices = &PricesService{client: c}
	c.Spec = &SpecService{client: c}
	c.Apps = &AppsService{client: c}

	return c
}

// do executes an HTTP request to the Commerce API.
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

	var reqBody io.Reader
	if body != nil {
		raw, err := c.jsonRequestBody(method, path, body, "")
		if err != nil {
			return fmt.Errorf("encode request body: %w", err)
		}
		reqBody = bytes.NewReader(raw)
	}

	url := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
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
		return apiErr
	}

	if out != nil && len(respBytes) > 0 {
		if err := json.Unmarshal(respBytes, out); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
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
	case "lookup", "page", "settings", "countries", "contents", "balances", "render_preview":
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
