package schedule

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v7/internal/transport"
)

// SchedulesService manages scheduled chimes.
type Service struct {
	client transport.Client
}

// Lookup retrieves scheduled chime details by schedule ID.
func (s *Service) Lookup(ctx context.Context, scheduleID string) (*Schedule, error) {
	var resp struct {
		ScheduledChime Schedule `json:"scheduled_chime"`
	}
	if err := s.client.Do(ctx, "POST", "/schedules/lookup", LookupParams{ScheduleID: scheduleID}, &resp); err != nil {
		return nil, err
	}
	return &resp.ScheduledChime, nil
}

// Cancel cancels a scheduled chime by schedule ID.
func (s *Service) Cancel(ctx context.Context, scheduleID string) (*Schedule, error) {
	var resp struct {
		ScheduledChime Schedule `json:"scheduled_chime"`
	}
	if err := s.client.Do(ctx, "POST", "/schedules/cancel", CancelParams{ScheduleID: scheduleID}, &resp); err != nil {
		return nil, err
	}
	return &resp.ScheduledChime, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
