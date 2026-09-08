package inttegro

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type requestOptions struct {
	IdempotencyKey string
}

// Review records a manual approval or rejection for an upload attempt.

func (c *Client) doJSON(ctx context.Context, path string, body any, opts requestOptions, out any) error {
	raw, err := c.jsonRequestBody("POST", path, body, opts.IdempotencyKey)
	if err != nil {
		return err
	}
	return c.doRaw(ctx, "POST", path, bytes.NewReader(raw), "application/json", opts.IdempotencyKey, true, out, "")
}

func (c *Client) doRaw(ctx context.Context, method, pathOrURL string, body io.Reader, contentType, idempotencyKey string, authenticated bool, out any, operation string) error {
	resp, err := c.rawResponse(ctx, method, pathOrURL, body, contentType, idempotencyKey, authenticated, operation)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if out == nil {
		return nil
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respBytes) == 0 {
		return nil
	}
	return json.Unmarshal(respBytes, out)
}

func (c *Client) rawResponse(ctx context.Context, method, pathOrURL string, body io.Reader, contentType, idempotencyKey string, authenticated bool, operation string) (*http.Response, error) {
	ctx, telemetry := c.startRequestTelemetry(ctx, method, pathOrURL, operation)
	defer telemetry.end()

	url := pathOrURL
	if len(pathOrURL) == 0 || pathOrURL[0] == '/' {
		url = c.BaseURL + pathOrURL
	}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		telemetry.failAndReport(ctx, err, "request_error")
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if authenticated {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if idempotencyKey != "" {
		req.Header.Set("Idempotency-Key", idempotencyKey)
	} else if authenticated && strings.EqualFold(method, "POST") && isIdempotentMutationPath(pathOrURL) && !strings.HasPrefix(contentType, "application/json") {
		req.Header.Set("Idempotency-Key", generateIdempotencyKey())
	}
	telemetry.inject(ctx, req.Header)
	telemetry.attempt()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		telemetry.failAndReport(ctx, err, "transport_error")
		return nil, err
	}
	telemetry.response(resp)
	if resp.StatusCode < 400 {
		return resp, nil
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		telemetry.failAndReport(ctx, err, "read_error")
		return nil, err
	}
	apiErr := &APIError{StatusCode: resp.StatusCode, Body: respBytes}
	var parsed APIError
	if err := json.Unmarshal(respBytes, &parsed); err == nil && hasAPIErrorDetails(&parsed) {
		copyAPIError(apiErr, &parsed)
	} else if len(respBytes) > 0 {
		apiErr.Message = string(respBytes)
	}
	apiErr.RequestID = resp.Header.Get("x-request-id")
	telemetry.failAndReport(ctx, apiErr, fmt.Sprintf("http_%d", resp.StatusCode))
	return nil, fmt.Errorf("inttegro api error: %w", apiErr)
}
