package commerce

import "context"

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
//	chime, err := client.Chimes.Send(ctx, commerce.SendChimeParams{
//	    Recipient: commerce.ChimeRecipient{
//	        Type: commerce.ChimeRecipientTypePhone,
//	        Phone: &struct{Number string}{Number: "+233244123456"},
//	    },
//	    FullMessage: "Your order #12345 has shipped!",
//	    Transport:   commerce.ChimeTransportSMS,
//	    IdempotencyKey: "chime_order_12345_shipped",
//	})
type ChimesService struct {
	client *Client
}

// Send dispatches a notification immediately.
//
// Sends SMS or email to a customer. Returns immediately with chime ID—
// delivery happens asynchronously.
func (s *ChimesService) Send(ctx context.Context, params SendChimeParams) (*Chime, error) {
	var resp struct {
		Chime Chime `json:"chime"`
	}
	if err := s.client.do(ctx, "POST", "/chimes/send", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Chime, nil
}

// Lookup retrieves chime details and delivery status by ID.
func (s *ChimesService) Lookup(ctx context.Context, chimeID string) (*Chime, error) {
	var resp struct {
		Chime Chime `json:"chime"`
	}
	if err := s.client.do(ctx, "POST", "/chimes/lookup", LookupChimeParams{ChimeID: chimeID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Chime, nil
}

// Schedule enqueues a notification for delivery at a specific time.
//
// Useful for reminders, follow-ups, or coordinated campaigns.
func (s *ChimesService) Schedule(ctx context.Context, params ScheduleChimeParams) (*ScheduledChime, error) {
	var resp ScheduleResponse
	if err := s.client.do(ctx, "POST", "/chimes/schedule", params, &resp); err != nil {
		return nil, err
	}
	return &resp.ScheduledChime, nil
}

// Broadcast sends a chime to many recipients.
func (s *ChimesService) Broadcast(ctx context.Context, params BroadcastChimeParams) (*BroadcastResponse, error) {
	var resp BroadcastResponse
	if err := s.client.do(ctx, "POST", "/chimes/broadcast", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
