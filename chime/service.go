package chime

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v6/broadcast"
	"github.com/zebodotdev/inttegro-sdk-go/v6/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v6/schedule"
)

// ChimesService sends and manages notification messages (SMS and email).
//
// Chimes are transactional notifications you send to customers. Use them for:
//   - Order confirmations and shipping updates
//   - Payment receipts and reminders
//   - OTP codes and security alerts
//   - Custom notifications
//
// Example:
//
//	chime, err := client.Chimes.Send(ctx, chime.SendParams{
//	    Recipient: chime.Recipient{
//	        Type: chime.RecipientTypePhone,
//	        Phone: &struct{Number string}{Number: "+233244123456"},
//	    },
//	    FullMessage: "Your order #12345 has shipped!",
//	    Transport:   chime.TransportSMS,
//	    IdempotencyKey: "chime_order_12345_shipped",
//	})
type Service struct {
	client transport.Client
}

// Send dispatches a notification immediately.
//
// Sends SMS or email to a customer. Returns immediately with chime ID—
// delivery happens asynchronously.
func (s *Service) Send(ctx context.Context, params SendParams) (*Chime, error) {
	var resp struct {
		Chime Chime `json:"chime"`
	}
	if err := s.client.Do(ctx, "POST", "/chimes/send", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Chime, nil
}

// Lookup retrieves chime details and delivery status by ID.
func (s *Service) Lookup(ctx context.Context, chimeID string) (*Chime, error) {
	var resp struct {
		Chime Chime `json:"chime"`
	}
	if err := s.client.Do(ctx, "POST", "/chimes/lookup", LookupParams{ChimeID: chimeID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Chime, nil
}

// Page retrieves a paginated list of chimes.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/chimes/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Schedule enqueues a notification for delivery at a specific time.
//
// Useful for reminders, follow-ups, or coordinated campaigns.
func (s *Service) Schedule(ctx context.Context, params ScheduleParams) (*schedule.ScheduledChime, error) {
	var resp struct {
		ScheduledChime schedule.ScheduledChime `json:"scheduled_chime"`
	}
	if err := s.client.Do(ctx, "POST", "/chimes/schedule", params, &resp); err != nil {
		return nil, err
	}
	return &resp.ScheduledChime, nil
}

// Broadcast sends a chime to many recipients.
func (s *Service) Broadcast(ctx context.Context, params broadcast.CreateParams) (*broadcast.Creation, error) {
	var resp struct {
		Broadcast broadcast.Creation `json:"broadcast"`
	}
	if err := s.client.Do(ctx, "POST", "/chimes/broadcast", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Broadcast, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
