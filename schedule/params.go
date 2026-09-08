package schedule

// LookupScheduleParams specifies which scheduled chime to retrieve.
type LookupParams struct {
	ScheduleID string `json:"schedule_id"`
}

// CancelScheduleParams specifies which scheduled chime to cancel.
type CancelParams struct {
	ScheduleID string `json:"schedule_id"`
}
