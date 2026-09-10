// Package transport defines the private boundary between resource services and
// the root HTTP client.
package transport

import (
	"context"
	"io"
	"net/http"

	"github.com/zebodotdev/inttegro-sdk-go/v7/request"
)

// Client is implemented by the root inttegro.Client.
type Client interface {
	Do(context.Context, string, string, any, any) error
	DoJSON(context.Context, string, any, request.Options, any) error
	DoRaw(context.Context, string, string, io.Reader, string, string, bool, any, string) error
	RawResponse(context.Context, string, string, io.Reader, string, string, bool, string) (*http.Response, error)
}
