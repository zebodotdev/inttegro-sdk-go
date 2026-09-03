package inttegro

import "context"

// RefundsService creates and manages refunds against paid order line items.
type RefundsService struct {
	client *Client
}

// Create starts an asynchronous refund for one or more paid order line items.
func (s *RefundsService) Create(
	ctx context.Context,
	request CreateRefundRequest,
) (*Refund, error) {
	return createRefund(ctx, s.client, "/refunds/create", request)
}

// Cancel cancels a pending refund before provider processing begins.
func (s *RefundsService) Cancel(
	ctx context.Context,
	request CancelRefundRequest,
) (*Refund, error) {
	var response struct {
		Refund Refund `json:"refund"`
	}
	if err := s.client.do(ctx, "POST", "/refunds/cancel", request, &response); err != nil {
		return nil, err
	}
	return &response.Refund, nil
}

// Lookup retrieves the current state of one refund.
func (s *RefundsService) Lookup(
	ctx context.Context,
	request LookupRefundRequest,
) (*Refund, error) {
	var response struct {
		Refund Refund `json:"refund"`
	}
	if err := s.client.do(ctx, "POST", "/refunds/lookup", request, &response); err != nil {
		return nil, err
	}
	return &response.Refund, nil
}

// Page returns one page of refunds, newest first.
func (s *RefundsService) Page(
	ctx context.Context,
	request PageRefundsRequest,
) (*RefundPage, error) {
	var response struct {
		Page RefundPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/refunds/page", request, &response); err != nil {
		return nil, err
	}
	return &response.Page, nil
}

func createRefund(
	ctx context.Context,
	client *Client,
	path string,
	request CreateRefundRequest,
) (*Refund, error) {
	var response struct {
		Refund Refund `json:"refund"`
	}
	if err := client.do(ctx, "POST", path, request, &response); err != nil {
		return nil, err
	}
	return &response.Refund, nil
}
