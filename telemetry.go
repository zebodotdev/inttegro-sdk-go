package inttegro

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type requestTelemetry struct {
	span       trace.Span
	propagator propagation.TextMapPropagator
}

var safeTelemetryResources = map[string]struct{}{
	"apps": {}, "balance_transactions": {}, "balances": {}, "broadcasts": {}, "checkout": {},
	"chimes": {}, "customers": {}, "file_links": {}, "file_references": {}, "files": {},
	"financial_accounts": {}, "keys": {}, "message_templates": {}, "orders": {}, "otp": {},
	"payment_methods": {}, "payouts": {}, "ping": {}, "prices": {}, "products": {},
	"purchase_intents": {}, "refunds": {}, "schedules": {}, "sessions": {}, "spec": {},
	"upload_requests": {},
}

var safeTelemetryActions = map[string]struct{}{
	"activate": {}, "add_price": {}, "archive": {}, "broadcast": {}, "cancel": {}, "complete": {},
	"confirm_payment": {}, "confirm_verification": {}, "connect": {}, "contents": {}, "countries": {},
	"create": {}, "deactivate": {}, "delete": {}, "destroy": {}, "disable": {}, "disable_fx": {},
	"disable_pull": {}, "disable_push": {}, "disactivate": {}, "disconnect": {}, "enable": {},
	"enable_fx": {}, "enable_pull": {}, "enable_push": {}, "finalize": {}, "generate": {},
	"initiate": {}, "lookup": {}, "new": {}, "open": {}, "page": {}, "pay": {}, "publish": {},
	"reconcile": {}, "reconnect": {}, "refund": {}, "render_preview": {}, "request_confirmation": {},
	"review": {}, "revoke": {}, "schedule": {}, "send": {}, "send_invoice": {}, "send_receipt": {},
	"set_default_unit_price": {}, "set_destinations": {}, "settings": {}, "tokenize": {},
	"unarchive": {}, "unpublish": {}, "update": {}, "upload": {}, "usage": {}, "verify": {},
}

func (c *Client) startRequestTelemetry(ctx context.Context, method, pathOrURL, explicitOperation string) (context.Context, requestTelemetry) {
	if !c.telemetryEnabled || c.tracer == nil {
		return ctx, requestTelemetry{}
	}

	operation, route, serverAddress := telemetryRequestDetails(c.BaseURL, pathOrURL, explicitOperation)
	ctx, span := c.tracer.Start(
		ctx,
		"inttegro."+operation,
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("inttegro.operation.name", operation),
			attribute.String("inttegro.sdk.language", "go"),
			attribute.String("inttegro.sdk.version", Version),
			attribute.String("http.request.method", strings.ToUpper(method)),
			attribute.String("server.address", serverAddress),
		),
	)
	if route != "" {
		span.SetAttributes(attribute.String("url.template", route))
	}
	span.AddEvent("inttegro.request.prepared")
	return ctx, requestTelemetry{span: span, propagator: c.propagator}
}

func (t requestTelemetry) inject(ctx context.Context, headers http.Header) {
	if t.span == nil || t.propagator == nil {
		return
	}
	carrier := propagation.HeaderCarrier(make(http.Header))
	t.propagator.Inject(ctx, carrier)
	for key, values := range http.Header(carrier) {
		if headers.Get(key) != "" {
			continue
		}
		for _, value := range values {
			headers.Add(key, value)
		}
	}
}

func (t requestTelemetry) attempt() {
	if t.span != nil {
		t.span.AddEvent("inttegro.http.attempt.started", trace.WithAttributes(attribute.Int("http.request.resend_count", 0)))
	}
}

func (t requestTelemetry) response(response *http.Response) {
	if t.span == nil || response == nil {
		return
	}
	t.span.SetAttributes(attribute.Int("http.response.status_code", response.StatusCode))
	if requestID := response.Header.Get("x-request-id"); requestID != "" {
		t.span.SetAttributes(attribute.String("inttegro.request.id", requestID))
	}
	t.span.AddEvent("inttegro.response.received")
}

func (t requestTelemetry) decoded() {
	if t.span != nil {
		t.span.AddEvent("inttegro.response.decoded")
	}
}

func (t requestTelemetry) fail(errorType string) {
	if t.span == nil || errorType == "" || errorType == "canceled" {
		return
	}
	t.span.SetAttributes(attribute.String("error.type", errorType))
	t.span.SetStatus(codes.Error, errorType)
	t.span.AddEvent("inttegro.request.failed", trace.WithAttributes(attribute.String("error.type", errorType)))
}

func (t requestTelemetry) end() {
	if t.span != nil {
		t.span.End()
	}
}

func telemetryRequestDetails(baseURL, pathOrURL, explicitOperation string) (operation, route, serverAddress string) {
	operation = explicitOperation
	if strings.HasPrefix(pathOrURL, "/") {
		if parsed, err := url.Parse(baseURL); err == nil {
			serverAddress = parsed.Hostname()
		}
		if isKnownTelemetryRoute(pathOrURL) {
			route = pathOrURL
		}
		if operation == "" && route != "" {
			operation = operationFromRoute(route)
		}
	} else if parsed, err := url.Parse(pathOrURL); err == nil {
		serverAddress = parsed.Hostname()
	}
	if operation == "" {
		operation = "http.request"
	}
	return operation, route, serverAddress
}

func isKnownTelemetryRoute(route string) bool {
	parts := strings.Split(strings.Trim(route, "/"), "/")
	if len(parts) == 0 || len(parts) > 2 {
		return false
	}
	if _, ok := safeTelemetryResources[parts[0]]; !ok {
		return false
	}
	if len(parts) == 2 {
		_, ok := safeTelemetryActions[parts[1]]
		return ok
	}
	return true
}

func operationFromRoute(route string) string {
	parts := strings.Split(strings.Trim(route, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		return "http.request"
	}
	if len(parts) == 1 {
		if parts[0] == "balances" {
			return "balances.lookup"
		}
		return parts[0] + ".request"
	}
	return parts[0] + "." + parts[len(parts)-1]
}

func classifyTelemetryError(err error, fallback string) string {
	if errors.Is(err, context.Canceled) {
		return "canceled"
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return "timeout"
	}
	var apiErr *APIError
	if errors.As(err, &apiErr) && apiErr.StatusCode > 0 {
		return "http_" + strconv.Itoa(apiErr.StatusCode)
	}
	var networkError net.Error
	if errors.As(err, &networkError) && networkError.Timeout() {
		return "timeout"
	}
	return fallback
}
