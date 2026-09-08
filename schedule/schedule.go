// Package schedule provides schedule resources and operations.
package schedule

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
)

// SchedulesService manages scheduled chimes.
type Service struct {
	client transport.Client
}

// Lookup retrieves scheduled chime details by schedule ID.
func (s *Service) Lookup(ctx context.Context, scheduleID string) (*Resource, error) {
	var resp struct {
		ScheduledChime Resource `json:"scheduled_chime"`
	}
	if err := s.client.Do(ctx, "POST", "/schedules/lookup", LookupParams{ScheduleID: scheduleID}, &resp); err != nil {
		return nil, err
	}
	return &resp.ScheduledChime, nil
}

// Cancel cancels a scheduled chime by schedule ID.
func (s *Service) Cancel(ctx context.Context, scheduleID string) (*Resource, error) {
	var resp struct {
		ScheduledChime Resource `json:"scheduled_chime"`
	}
	if err := s.client.Do(ctx, "POST", "/schedules/cancel", CancelParams{ScheduleID: scheduleID}, &resp); err != nil {
		return nil, err
	}
	return &resp.ScheduledChime, nil
}

// ScheduleDetail describes a scheduled chime and its execution state.
type Resource struct {
	ID         string   `json:"id,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
	Content    string   `json:"content,omitempty"`
	SenderID   string   `json:"sender_id,omitempty"`
	Purpose    *string  `json:"purpose,omitempty"`
	SendAfter  string   `json:"send_after,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty"`
	ExecutedAt *string  `json:"executed_at,omitempty"`
	CanceledAt *string  `json:"canceled_at,omitempty"`
	Errors     []Error  `json:"errors,omitempty"`
	ChimeIDs   []string `json:"chime_ids,omitempty"`
}

// ScheduleError reports a per-recipient scheduling error.
type Error struct {
	Recipient string `json:"recipient,omitempty"`
	FixCode   string `json:"fix_code,omitempty"`
	Type      string `json:"type,omitempty"`
}

// ScheduledChime is returned after scheduling a chime.
type ScheduledChime struct {
	ID          string   `json:"id,omitempty"`
	Recipients  []string `json:"recipients,omitempty"`
	FullMessage string   `json:"full_message,omitempty"`
	SenderID    string   `json:"sender_id,omitempty"`
	Purpose     *string  `json:"purpose,omitempty"`
	SendAfter   string   `json:"send_after,omitempty"`
	CreatedAt   string   `json:"created_at,omitempty"`
	ExecutedAt  *string  `json:"executed_at,omitempty"`
}

// LookupScheduleParams specifies which scheduled chime to retrieve.
type LookupParams struct {
	ScheduleID string `json:"schedule_id"`
}

// CancelScheduleParams specifies which scheduled chime to cancel.
type CancelParams struct {
	ScheduleID string `json:"schedule_id"`
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
