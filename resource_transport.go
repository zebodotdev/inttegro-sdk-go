package inttegro

import (
	"context"
	"io"
	"net/http"

	"github.com/zebodotdev/inttegro-sdk-go/v6/request"
)

// Do executes a JSON API request. It is exported for use by resource packages;
// applications should call the typed resource services on Client instead.
func (c *Client) Do(ctx context.Context, method, path string, body, out any) error {
	return c.do(ctx, method, path, body, out)
}

// DoJSON executes a JSON mutation with explicit request options.
func (c *Client) DoJSON(ctx context.Context, path string, body any, options request.Options, out any) error {
	return c.doJSON(ctx, path, body, requestOptions{IdempotencyKey: options.IdempotencyKey}, out)
}

// DoRaw executes a raw request for resource packages.
func (c *Client) DoRaw(ctx context.Context, method, pathOrURL string, body io.Reader, contentType, idempotencyKey string, authenticated bool, out any, operation string) error {
	return c.doRaw(ctx, method, pathOrURL, body, contentType, idempotencyKey, authenticated, out, operation)
}

// RawResponse executes a raw request and returns its HTTP response.
func (c *Client) RawResponse(ctx context.Context, method, pathOrURL string, body io.Reader, contentType, idempotencyKey string, authenticated bool, operation string) (*http.Response, error) {
	return c.rawResponse(ctx, method, pathOrURL, body, contentType, idempotencyKey, authenticated, operation)
}
