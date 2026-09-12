// Package response contains response envelopes returned by opt-in SDK methods
// that preserve native HTTP response facts alongside decoded domain data.
package response

import "net/http"

// Meta is curated response metadata returned by the Inttegro API response body.
// It is distinct from native HTTP status and headers.
type Meta map[string]any

// Response contains a decoded SDK value plus response-only HTTP metadata.
type Response[T any] struct {
	// Data is the decoded domain value returned by the endpoint.
	Data T

	// StatusCode is the HTTP response status code.
	StatusCode int

	// Headers contains response headers only.
	Headers http.Header

	// Meta contains the API response_meta object, when present.
	Meta Meta
}

// RequestID returns the X-Request-Id response header, when present.
func (r *Response[T]) RequestID() string {
	if r == nil {
		return ""
	}
	return r.Headers.Get("X-Request-Id")
}

// RetryAfter returns the Retry-After response header, when present.
func (r *Response[T]) RetryAfter() string {
	if r == nil {
		return ""
	}
	return r.Headers.Get("Retry-After")
}

// WithData returns a response envelope with the same response metadata and new data.
func WithData[T any](base *Response[any], data T) *Response[T] {
	if base == nil {
		return &Response[T]{Data: data}
	}
	return &Response[T]{
		Data:       data,
		StatusCode: base.StatusCode,
		Headers:    base.Headers.Clone(),
		Meta:       cloneMeta(base.Meta),
	}
}

func cloneMeta(meta Meta) Meta {
	if len(meta) == 0 {
		return nil
	}
	clone := make(Meta, len(meta))
	for key, value := range meta {
		clone[key] = value
	}
	return clone
}
