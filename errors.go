package commerce

import (
	"fmt"
)

// APIError represents a structured error response from the Commerce API.
//
// All API errors (HTTP status >= 400) are returned as *APIError. The error
// contains a machine-readable code, human-readable message, and additional context.
//
// Error structure:
//   - Type: Broad error category (e.g., "invalid_request_parameter")
//   - Code: Unique error identifier (e.g., "invalid_payment_method")
//   - URL: Link to error reference documentation
//   - Message: Human-readable summary of the error
//   - Detail: Comprehensive explanation with guidance
//   - FixCode: Machine-readable suggestion for how to resolve the error
//   - Cause: Underlying category of failure (e.g., "validation_failure")
//   - StatusCode: HTTP status code
//
// Example:
//
//	order, err := client.Orders.Create(ctx, params)
//	if err != nil {
//	    if apiErr, ok := err.(*commerce.APIError); ok {
//	        fmt.Printf("Error code: %s\n", apiErr.Code)
//	        fmt.Printf("Type: %s\n", apiErr.Type)
//	        fmt.Printf("Message: %s\n", apiErr.Message)
//	        if apiErr.URL != "" {
//	            fmt.Printf("Docs: %s\n", apiErr.URL)
//	        }
//	    }
//	    return err
//	}
//
// Learn more: https://studio.zebo.dev/errors
type APIError struct {
	// StatusCode is the HTTP status code (400-599).
	StatusCode int `json:"-"`

	// Type is the broad error category.
	// Example: "invalid_request_parameter", "transient_error"
	Type string `json:"type"`

	// Code is the machine-readable error identifier.
	// Example: "invalid_payment_method"
	Code string `json:"code"`

	// URL links to error reference documentation.
	// Example: "https://studio.zebo.dev/e/invalid_payment_method"
	URL string `json:"url"`

	// Message is a concise, human-readable summary of the error.
	Message string `json:"message"`

	// Detail is a comprehensive explanation with guidance.
	Detail string `json:"detail"`

	// FixCode is a machine-readable suggestion for resolution.
	// Example: "change_request_parameters"
	FixCode string `json:"fix_code"`

	// Cause is the underlying category of failure.
	// Example: "validation_failure"
	Cause string `json:"cause"`

	// Body contains the raw response body for debugging.
	// Useful when error parsing fails or for logging.
	Body []byte `json:"-"`
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	code := e.Code
	if code == "" {
		code = e.Type
	}
	if code == "" {
		code = fmt.Sprintf("status_%d", e.StatusCode)
	}
	msg := e.Message
	if msg == "" {
		msg = e.Detail
	}
	if msg != "" {
		return fmt.Sprintf("%s: %s", code, msg)
	}
	return code
}

// errorEnvelope matches legacy error response shapes.
type errorEnvelope struct {
	Error *APIError `json:"error"`
}
