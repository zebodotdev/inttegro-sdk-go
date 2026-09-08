// Package request defines controls shared by resource service calls.
package request

// Meta carries per-request controls that do not change an operation payload.
type Meta struct {
	// IdempotencyKey safely identifies retries for mutation requests.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

// Options contains controls passed alongside a request body.
type Options struct {
	IdempotencyKey string
}

// Option configures one resource request.
type Option func(*Options)

// WithIdempotencyKey sets the key used to safely retry a mutation.
func WithIdempotencyKey(key string) Option {
	return func(options *Options) {
		options.IdempotencyKey = key
	}
}

// Apply resolves a sequence of request options.
func Apply(options []Option) Options {
	var resolved Options
	for _, option := range options {
		if option != nil {
			option(&resolved)
		}
	}
	return resolved
}
