package inttegro

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// ErrorReportingPolicy controls which final SDK failures are reported.
type ErrorReportingPolicy string

const (
	// ErrorReportingUnexpected reports SDK, transport, decoding, and server failures.
	ErrorReportingUnexpected ErrorReportingPolicy = "unexpected"
	// ErrorReportingAll reports every final SDK failure except cancellation.
	ErrorReportingAll ErrorReportingPolicy = "all"
)

// SDKReportContext identifies the SDK that generated an error report.
type SDKReportContext struct {
	Language string `json:"language"`
	Version  string `json:"version"`
}

// HTTPReportContext contains bounded, privacy-safe request metadata.
type HTTPReportContext struct {
	Method        string `json:"method"`
	Route         string `json:"route,omitempty"`
	ServerAddress string `json:"serverAddress"`
	StatusCode    int    `json:"statusCode,omitempty"`
	RequestID     string `json:"requestId,omitempty"`
	DurationMS    int64  `json:"durationMs"`
}

// APIErrorReportContext contains machine-readable API failure metadata.
type APIErrorReportContext struct {
	Type    string `json:"type,omitempty"`
	Code    string `json:"code,omitempty"`
	FixCode string `json:"fixCode,omitempty"`
}

// TraceReportContext links a report to application-owned tracing.
type TraceReportContext struct {
	TraceID string `json:"traceId"`
	SpanID  string `json:"spanId"`
}

// ErrorReport is a privacy-safe description of a failed Inttegro SDK operation.
type ErrorReport struct {
	SchemaVersion int                    `json:"schemaVersion"`
	EventID       string                 `json:"eventId"`
	OccurredAt    time.Time              `json:"occurredAt"`
	Severity      string                 `json:"severity"`
	Category      string                 `json:"category"`
	Operation     string                 `json:"operation"`
	SDK           SDKReportContext       `json:"sdk"`
	HTTP          HTTPReportContext      `json:"http"`
	APIError      *APIErrorReportContext `json:"apiError,omitempty"`
	Trace         *TraceReportContext    `json:"trace,omitempty"`
	ExceptionType string                 `json:"exceptionType"`
	Fingerprint   string                 `json:"fingerprint"`
}

// ErrorReporter receives one report after a logical SDK operation finally fails.
// Implementations should enqueue quickly; reporter panics are isolated by the SDK.
type ErrorReporter func(context.Context, ErrorReport)

func (t requestTelemetry) failAndReport(ctx context.Context, err error, fallback string) {
	if t.span == nil && t.errorReporter == nil {
		return
	}
	defer func() { _ = recover() }()
	category := classifyTelemetryError(err, fallback)
	t.fail(category)
	if t.errorReporter == nil || category == "canceled" {
		return
	}
	var apiErr *APIError
	_ = errors.As(err, &apiErr)
	if t.errorReportingPolicy == ErrorReportingUnexpected && apiErr != nil && apiErr.StatusCode < 500 && apiErr.Type != "unknown_error" {
		return
	}

	var apiContext *APIErrorReportContext
	if apiErr != nil && (apiErr.Type != "" || apiErr.Code != "" || apiErr.FixCode != "") {
		apiContext = &APIErrorReportContext{Type: apiErr.Type, Code: apiErr.Code, FixCode: apiErr.FixCode}
	}
	var traceContext *TraceReportContext
	if t.span != nil {
		spanContext := t.span.SpanContext()
		if spanContext.IsValid() {
			traceContext = &TraceReportContext{
				TraceID: spanContext.TraceID().String(),
				SpanID:  spanContext.SpanID().String(),
			}
		}
	}
	statusCode := 0
	requestID := ""
	fingerprintStatus := "none"
	if apiErr != nil {
		statusCode = apiErr.StatusCode
		requestID = apiErr.RequestID
		fingerprintStatus = fmt.Sprintf("%d", statusCode)
	}
	report := &ErrorReport{
		SchemaVersion: 1,
		EventID:       generateIdempotencyKey(),
		OccurredAt:    time.Now().UTC(),
		Severity:      "error",
		Category:      category,
		Operation:     t.operation,
		SDK:           SDKReportContext{Language: "go", Version: Version},
		HTTP: HTTPReportContext{
			Method:        strings.ToUpper(t.method),
			Route:         t.route,
			ServerAddress: t.serverAddress,
			StatusCode:    statusCode,
			RequestID:     requestID,
			DurationMS:    max(time.Since(t.startedAt).Milliseconds(), 0),
		},
		APIError:      apiContext,
		Trace:         traceContext,
		ExceptionType: exceptionType(err),
		Fingerprint: strings.Join([]string{
			"inttegro", "go", t.operation, category, fingerprintStatus,
		}, ":"),
	}
	if apiErr != nil {
		apiErr.Report = report
	}
	callErrorReporter(t.errorReporter, ctx, *report)
}

func exceptionType(err error) string {
	if err == nil {
		return "error"
	}
	typeOfError := reflect.TypeOf(err)
	for typeOfError.Kind() == reflect.Pointer {
		typeOfError = typeOfError.Elem()
	}
	if name := typeOfError.Name(); name != "" {
		return name
	}
	return "error"
}

func callErrorReporter(reporter ErrorReporter, ctx context.Context, report ErrorReport) {
	defer func() { _ = recover() }()
	reporter(ctx, report)
}
