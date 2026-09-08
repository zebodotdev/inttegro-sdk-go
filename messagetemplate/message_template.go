// Package messagetemplate provides messagetemplate resources and operations.
package messagetemplate

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

type Channel string

const (
	ChannelSMS   Channel = "sms"
	ChannelEmail Channel = "email"
)

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPublished Status = "published"
	StatusArchived  Status = "archived"
)

type VariableType string

const (
	VariableTypeString   VariableType = "string"
	VariableTypeNumber   VariableType = "number"
	VariableTypeInteger  VariableType = "integer"
	VariableTypeBoolean  VariableType = "boolean"
	VariableTypeURL      VariableType = "url"
	VariableTypeEmail    VariableType = "email"
	VariableTypePhone    VariableType = "phone"
	VariableTypeDate     VariableType = "date"
	VariableTypeDatetime VariableType = "datetime"
	VariableTypeArray    VariableType = "array"
)

type VariableItemType string

const (
	VariableItemTypeString   VariableItemType = "string"
	VariableItemTypeNumber   VariableItemType = "number"
	VariableItemTypeInteger  VariableItemType = "integer"
	VariableItemTypeBoolean  VariableItemType = "boolean"
	VariableItemTypeURL      VariableItemType = "url"
	VariableItemTypeEmail    VariableItemType = "email"
	VariableItemTypePhone    VariableItemType = "phone"
	VariableItemTypeDate     VariableItemType = "date"
	VariableItemTypeDatetime VariableItemType = "datetime"
)

type ContentSafetyStatus string

const (
	ContentSafetyStatusAllowed     ContentSafetyStatus = "allowed"
	ContentSafetyStatusRejected    ContentSafetyStatus = "rejected"
	ContentSafetyStatusQuarantined ContentSafetyStatus = "quarantined"
)

type Variable struct {
	About    string         `json:"about,omitempty"`
	Default  any            `json:"default,omitempty"`
	Items    []VariableItem `json:"items,omitempty"`
	Name     string         `json:"name"`
	Required bool           `json:"required,omitempty"`
	Type     VariableType   `json:"type"`
}

type VariableItem struct {
	About    string           `json:"about,omitempty"`
	Default  any              `json:"default,omitempty"`
	Name     string           `json:"name"`
	Required bool             `json:"required,omitempty"`
	Type     VariableItemType `json:"type"`
}

type SMSContent struct {
	MessageTemplate string `json:"message_template"`
}

type Mailbox struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type EmailContent struct {
	From    *Mailbox          `json:"from,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	HTML    string            `json:"html"`
	ReplyTo *Mailbox          `json:"reply_to,omitempty"`
	Subject string            `json:"subject"`
}

type Resource struct {
	ID                    string        `json:"id"`
	About                 string        `json:"about,omitempty"`
	ArchivedAt            string        `json:"archived_at,omitempty"`
	Attachments           []string      `json:"attachments,omitempty"`
	Channel               Channel       `json:"channel"`
	CreatedAt             string        `json:"created_at"`
	DraftVersion          int           `json:"draft_version"`
	Email                 *EmailContent `json:"email,omitempty"`
	HasUnpublishedChanges bool          `json:"has_unpublished_changes"`
	Locale                string        `json:"locale"`
	Name                  string        `json:"name"`
	PublishedAt           string        `json:"published_at,omitempty"`
	PublishedVersion      *int          `json:"published_version,omitempty"`
	Purpose               string        `json:"purpose"`
	SMS                   *SMSContent   `json:"sms,omitempty"`
	Status                Status        `json:"status"`
	UpdatedAt             string        `json:"updated_at"`
	Variables             []Variable    `json:"variables,omitempty"`
	Version               int           `json:"version"`
}

type CreateParams struct {
	IdempotencyKey string        `json:"-"`
	About          string        `json:"about,omitempty"`
	Attachments    []string      `json:"attachments,omitempty"`
	Channel        Channel       `json:"channel"`
	Email          *EmailContent `json:"email,omitempty"`
	Locale         string        `json:"locale,omitempty"`
	Name           string        `json:"name"`
	Purpose        string        `json:"purpose"`
	SMS            *SMSContent   `json:"sms,omitempty"`
	Variables      []Variable    `json:"variables,omitempty"`
}

type UpdateParams struct {
	IdempotencyKey string        `json:"-"`
	ID             string        `json:"id"`
	About          string        `json:"about,omitempty"`
	Attachments    []string      `json:"attachments,omitempty"`
	Channel        Channel       `json:"channel,omitempty"`
	Email          *EmailContent `json:"email,omitempty"`
	Locale         string        `json:"locale,omitempty"`
	Name           string        `json:"name,omitempty"`
	Purpose        string        `json:"purpose,omitempty"`
	SMS            *SMSContent   `json:"sms,omitempty"`
	Variables      []Variable    `json:"variables,omitempty"`
}

type PageParams struct {
	Channel Channel `json:"channel,omitempty"`
	Locale  string  `json:"locale,omitempty"`
	Page    int     `json:"page,omitempty"`
	Purpose string  `json:"purpose,omitempty"`
	Size    int     `json:"size,omitempty"`
	Status  Status  `json:"status,omitempty"`
}

type Page struct {
	Number           int        `json:"number"`
	Size             int        `json:"size"`
	MessageTemplates []Resource `json:"message_templates"`
}

type Reference struct {
	TemplateID string         `json:"template_id"`
	Variables  map[string]any `json:"variables,omitempty"`
}

type RenderPreviewParams struct {
	MessageTemplate Reference `json:"message_template"`
}

type RenderedContent struct {
	Channel Channel        `json:"channel"`
	Email   map[string]any `json:"email,omitempty"`
	SMS     map[string]any `json:"sms,omitempty"`
}

type RenderPreviewOutput struct {
	MessageTemplate *Resource        `json:"message_template,omitempty"`
	Rendered        *RenderedContent `json:"rendered,omitempty"`
}

type Service struct {
	client transport.Client
}

func (s *Service) Create(ctx context.Context, params CreateParams) (*Resource, error) {
	var resp struct {
		MessageTemplate Resource `json:"message_template"`
	}
	opts := request.Options{IdempotencyKey: params.IdempotencyKey}
	if err := s.client.DoJSON(ctx, "/message_templates/create", params, opts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *Service) Update(ctx context.Context, params UpdateParams) (*Resource, error) {
	var resp struct {
		MessageTemplate Resource `json:"message_template"`
	}
	opts := request.Options{IdempotencyKey: params.IdempotencyKey}
	if err := s.client.DoJSON(ctx, "/message_templates/update", params, opts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *Service) Publish(ctx context.Context, id string, opts ...request.Option) (*Resource, error) {
	var resp struct {
		MessageTemplate Resource `json:"message_template"`
	}
	requestOpts := request.Apply(opts)
	if err := s.client.DoJSON(ctx, "/message_templates/publish", map[string]string{"id": id}, requestOpts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *Service) Archive(ctx context.Context, id string, opts ...request.Option) (*Resource, error) {
	var resp struct {
		MessageTemplate Resource `json:"message_template"`
	}
	requestOpts := request.Apply(opts)
	if err := s.client.DoJSON(ctx, "/message_templates/archive", map[string]string{"id": id}, requestOpts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *Service) Lookup(ctx context.Context, id string) (*Resource, error) {
	var resp struct {
		MessageTemplate Resource `json:"message_template"`
	}
	if err := s.client.Do(ctx, "POST", "/message_templates/lookup", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/message_templates/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

func (s *Service) RenderPreview(ctx context.Context, params RenderPreviewParams) (*RenderPreviewOutput, error) {
	var resp RenderPreviewOutput
	if err := s.client.Do(ctx, "POST", "/message_templates/render_preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
