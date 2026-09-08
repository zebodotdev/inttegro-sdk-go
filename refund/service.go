package refund

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v6/internal/transport"
)

// RefundsService creates and manages refunds against paid order line items.
type Service struct {
	client transport.Client
}

// Create starts an asynchronous refund for one or more paid order line items.
func (s *Service) Create(
	ctx context.Context,
	request CreateParams,
) (*Refund, error) {
	return createRefund(ctx, s.client, "/refunds/create", request)
}

// Cancel cancels a pending refund before provider processing begins.
func (s *Service) Cancel(
	ctx context.Context,
	request CancelParams,
) (*Refund, error) {
	var response struct {
		Refund Refund `json:"refund"`
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
) (*Refund, error) {
	var response struct {
		Refund Refund `json:"refund"`
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
) (*Refund, error) {
	var response struct {
		Refund Refund `json:"refund"`
	}
	if err := client.Do(ctx, "POST", path, request, &response); err != nil {
		return nil, err
	}
	return &response.Refund, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
