package messagetemplate

import "time"

// MessageTemplate represents a reusable chime message template.
type MessageTemplate struct {
	ID                    string        `json:"id"`
	About                 string        `json:"about,omitempty"`
	ArchivedAt            *time.Time    `json:"archived_at,omitempty"`
	Attachments           []string      `json:"attachments,omitempty"`
	Channel               Channel       `json:"channel"`
	CreatedAt             time.Time     `json:"created_at"`
	DraftVersion          int           `json:"draft_version"`
	Email                 *EmailContent `json:"email,omitempty"`
	HasUnpublishedChanges bool          `json:"has_unpublished_changes"`
	Locale                string        `json:"locale"`
	Name                  string        `json:"name"`
	PublishedAt           *time.Time    `json:"published_at,omitempty"`
	PublishedVersion      *int          `json:"published_version,omitempty"`
	Purpose               string        `json:"purpose"`
	SMS                   *SMSContent   `json:"sms,omitempty"`
	Status                Status        `json:"status"`
	UpdatedAt             time.Time     `json:"updated_at"`
	Variables             []Variable    `json:"variables,omitempty"`
	Version               int           `json:"version"`
}

type Page struct {
	Number           int               `json:"number"`
	Size             int               `json:"size"`
	MessageTemplates []MessageTemplate `json:"message_templates"`
}

type Reference struct {
	TemplateID string         `json:"template_id"`
	Variables  map[string]any `json:"variables,omitempty"`
}

type RenderPreviewOutput struct {
	MessageTemplate *MessageTemplate `json:"message_template,omitempty"`
	Rendered        *RenderedContent `json:"rendered,omitempty"`
}
