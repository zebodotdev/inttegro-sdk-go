package inttegro

import "context"

// SchedulesService manages scheduled chimes.
type SchedulesService struct {
	client *Client
}

// Lookup retrieves scheduled chime details by schedule ID.
func (s *SchedulesService) Lookup(ctx context.Context, scheduleID string) (*ScheduleDetail, error) {
	var resp ScheduleLookupResponse
	if err := s.client.do(ctx, "POST", "/schedules/lookup", LookupScheduleParams{ScheduleID: scheduleID}, &resp); err != nil {
		return nil, err
	}
	return &resp.ScheduledChime, nil
}

// Cancel cancels a scheduled chime by schedule ID.
func (s *SchedulesService) Cancel(ctx context.Context, scheduleID string) (*ScheduleDetail, error) {
	var resp ScheduleCancelResponse
	if err := s.client.do(ctx, "POST", "/schedules/cancel", CancelScheduleParams{ScheduleID: scheduleID}, &resp); err != nil {
		return nil, err
	}
	return &resp.ScheduledChime, nil
}
