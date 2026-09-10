package uploadrequest

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v6/filelink"
)

type CreateParams struct {
	Constraints Constraints       `json:"constraints"`
	Display     Display           `json:"display"`
	ExpiresAt   *time.Time        `json:"expires_at,omitempty"`
	CustomData  map[string]string `json:"custom_data,omitempty"`
	Purpose     string            `json:"purpose"`
	Recipient   Party             `json:"recipient"`
	Requester   filelink.Actor    `json:"requester"`
	Resource    FileResource      `json:"resource"`
	Subject     Party             `json:"subject"`
}

type PageParams struct {
	PageNumber int          `json:"page_number,omitempty"`
	PageSize   int          `json:"page_size,omitempty"`
	Purpose    string       `json:"purpose,omitempty"`
	Resource   FileResource `json:"resource"`
	Status     Status       `json:"status,omitempty"`
}

type CancelParams struct {
	CanceledBy filelink.Actor `json:"canceled_by"`
	ID         string         `json:"id"`
}

type ReviewParams struct {
	AttemptID      string         `json:"attempt_id,omitempty"`
	AttemptOrdinal int64          `json:"attempt_ordinal,omitempty"`
	Decision       string         `json:"decision"`
	ID             string         `json:"id"`
	PublicMessage  string         `json:"public_message,omitempty"`
	Reasons        []ReviewReason `json:"reasons,omitempty"`
}

type FulfillParams struct {
	File      string
	UploadURL string
}
