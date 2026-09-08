// Package refund provides refund resources and operations.
package refund

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/money"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

type Reason string

const (
	ReasonRequestedByCustomer Reason = "requested_by_customer"
)

const (
	ReasonDuplicate Reason = "duplicate"
)

const (
	ReasonFraudulent Reason = "fraudulent"
)

const (
	ReasonOrderCanceled Reason = "order_canceled"
)

const (
	ReasonItemReturned Reason = "item_returned"
)

const (
	ReasonItemDamaged Reason = "item_damaged"
)

const (
	ReasonItemNotReceived Reason = "item_not_received"
)

const (
	ReasonItemNotAsDescribed Reason = "item_not_as_described"
)

const (
	ReasonCustom Reason = "custom"
)

type Status string

const (
	StatusCanceled Status = "canceled"
)

const (
	StatusFailed Status = "failed"
)

const (
	StatusPending Status = "pending"
)

const (
	StatusProcessing Status = "processing"
)

const (
	StatusSucceeded Status = "succeeded"
)

// RefundsService creates and manages refunds against paid order line items.
type Service struct {
	client transport.Client
}

// Create starts an asynchronous refund for one or more paid order line items.
func (s *Service) Create(
	ctx context.Context,
	request CreateParams,
) (*Resource, error) {
	return createRefund(ctx, s.client, "/refunds/create", request)
}

// Cancel cancels a pending refund before provider processing begins.
func (s *Service) Cancel(
	ctx context.Context,
	request CancelParams,
) (*Resource, error) {
	var response struct {
		Refund Resource `json:"refund"`
	}
	if err := s.client.Do(ctx, "POST", "/refunds/cancel", request, &response); err != nil {
		return nil, err
	}
	return &response.Refund, nil
}

// Lookup retrieves the current state of one refund.
func (s *Service) Lookup(
	ctx context.Context,
	request LookupParams,
) (*Resource, error) {
	var response struct {
		Refund Resource `json:"refund"`
	}
	if err := s.client.Do(ctx, "POST", "/refunds/lookup", request, &response); err != nil {
		return nil, err
	}
	return &response.Refund, nil
}

// Page returns one page of refunds, newest first.
func (s *Service) Page(
	ctx context.Context,
	request PageParams,
) (*Page, error) {
	var response struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/refunds/page", request, &response); err != nil {
		return nil, err
	}
	return &response.Page, nil
}

func createRefund(
	ctx context.Context,
	client transport.Client, path string,
	request CreateParams,
) (*Resource, error) {
	var response struct {
		Refund Resource `json:"refund"`
	}
	if err := client.Do(ctx, "POST", path, request, &response); err != nil {
		return nil, err
	}
	return &response.Refund, nil
}

// CreateRefundLineItem requests a refund allocation against one paid order
// line item. Reason and ReasonDetails are independent from the overall reason.
type CreateLineItem struct {
	OrderLineItemID string             `json:"order_line_item_id"`
	RefundAmount    money.AmountParams `json:"refund_amount"`
	Reason          *Reason            `json:"reason,omitempty"`
	ReasonDetails   string             `json:"reason_details,omitempty"`
}

// CreateRefundRequest starts a refund for 1 to 64 paid order line items.
type CreateParams struct {
	LineItems     []CreateLineItem  `json:"line_items"`
	OrderID       string            `json:"order_id"`
	Reason        Reason            `json:"reason"`
	CustomData    map[string]string `json:"custom_data,omitempty"`
	ReasonDetails string            `json:"reason_details,omitempty"`
	Reference     string            `json:"reference,omitempty"`
	RequestMeta   *request.Meta     `json:"request_meta,omitempty"`
}

// CancelRefundRequest cancels a refund that has not begun processing.
type CancelParams struct {
	RefundID    string        `json:"refund_id"`
	RequestMeta *request.Meta `json:"request_meta,omitempty"`
}

// LookupRefundRequest identifies the refund to retrieve.
type LookupParams struct {
	RefundID string `json:"refund_id"`
}

// PageRefundsRequest selects a one-based refund page. PageNumber is required.
type PageParams struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size,omitempty"`
}

// RefundLineItem is one immutable order-line allocation in a refund.
type LineItem struct {
	ID                 string       `json:"id"`
	OrderLineItemID    string       `json:"order_line_item_id"`
	OriginalAmountPaid money.Amount `json:"original_amount_paid"`
	RefundAmount       money.Amount `json:"refund_amount"`
	Reason             *Reason      `json:"reason,omitempty"`
	ReasonDetails      string       `json:"reason_details,omitempty"`
}

// Refund is the canonical refund object embedded in order responses.
type Resource struct {
	ID            string            `json:"id"`
	OrderID       string            `json:"order_id"`
	Status        Status            `json:"status"`
	Total         money.Amount      `json:"total"`
	LineItems     []LineItem        `json:"line_items"`
	Reason        Reason            `json:"reason"`
	ReasonDetails string            `json:"reason_details,omitempty"`
	Reference     string            `json:"reference,omitempty"`
	CustomData    map[string]string `json:"custom_data,omitempty"`
	CreatedAt     string            `json:"created_at"`
	ProcessingAt  *string           `json:"processing_at,omitempty"`
	SucceededAt   *string           `json:"succeeded_at,omitempty"`
	FailedAt      *string           `json:"failed_at,omitempty"`
	CanceledAt    *string           `json:"canceled_at,omitempty"`
}

// RefundPage contains one page of refunds.
type Page struct {
	Number  int        `json:"number"`
	Refunds []Resource `json:"refunds"`
	Size    int        `json:"size"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
