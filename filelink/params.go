package filelink

import "time"

type CreateParams struct {
	Access     Access            `json:"access"`
	CreatedBy  Actor             `json:"created_by"`
	Delivery   Delivery          `json:"delivery"`
	ExpiresAt  *time.Time        `json:"expires_at,omitempty"`
	FileID     string            `json:"file_id"`
	CustomData map[string]string `json:"custom_data,omitempty"`
}

type PageParams struct {
	FileID     string `json:"file_id,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	Status     Status `json:"status,omitempty"`
}

type RevokeParams struct {
	ID        string `json:"id"`
	RevokedBy Actor  `json:"revoked_by"`
}
