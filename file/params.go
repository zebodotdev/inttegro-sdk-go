package file

import "time"

type CreateParams struct {
	File           string            `json:"-"`
	Purpose        string            `json:"purpose"`
	Title          string            `json:"title,omitempty"`
	CustomData     map[string]string `json:"custom_data,omitempty"`
	IdempotencyKey string            `json:"-"`
}

type PageParams struct {
	CreatedAfter  *time.Time `json:"created_after,omitempty"`
	CreatedBefore *time.Time `json:"created_before,omitempty"`
	PageNumber    int        `json:"page_number,omitempty"`
	PageSize      int        `json:"page_size,omitempty"`
	Purpose       string     `json:"purpose,omitempty"`
	Status        Status     `json:"status,omitempty"`
}

type ContentsParams struct {
	FileID      string      `json:"file_id"`
	Disposition Disposition `json:"disposition,omitempty"`
}
